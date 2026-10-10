<script>
  let userId = "";
  let isLoggedIn = false;
  let viewMode = "shop"; // "shop" or "checkout"
  
  let products = [];
  let cart = { items: [], total: 0 };

  const API_URL = "http://localhost:8080/api";

  function handleLogin() {
    if (!userId.trim()) return;
    isLoggedIn = true;
    fetchData();
  }

  async function fetchData() {
    await fetchInventory();
    await fetchCart();
  }

  async function fetchInventory() {
    const res = await fetch(`${API_URL}/inventory`);
    products = await res.json();
  }

  async function fetchCart() {
    const res = await fetch(`${API_URL}/cart?user_id=${userId}`);
    cart = await res.json();
  }

  function getCartItemQty(productId) {
    if (!cart.items) return 0;
    const item = cart.items.find(i => i.product.id === productId);
    return item ? item.quantity : 0;
  }

  async function updateCart(productID, delta) {
    await fetch(`${API_URL}/cart`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ user_id: userId, product_id: productID, quantity: delta })
    });
    fetchCart();
  }
</script>

<style>
  :global(body) {
    background-color: #f8f9fa;
    margin: 0;
    font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  }

  .container {
    max-width: 1000px;
    margin: auto;
    padding: 2rem;
  }

  .login-box {
    background: white;
    padding: 2.5rem;
    border-radius: 12px;
    box-shadow: 0 4px 12px rgba(0,0,0,0.05);
    max-width: 400px;
    margin: 4rem auto;
    text-align: center;
  }

  .header-bar {
    display: flex;
    justify-content: space-between;
    align-items: center;
    background: white;
    padding: 1rem 1.5rem;
    border-radius: 12px;
    box-shadow: 0 2px 8px rgba(0,0,0,0.04);
    margin-bottom: 2rem;
  }

  .nav-buttons {
    display: flex;
    gap: 0.5rem;
  }

  .grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(260px, 1fr));
    gap: 1.5rem;
    margin-bottom: 2rem;
  }

  .card {
    background: white;
    border-radius: 12px;
    overflow: hidden;
    box-shadow: 0 4px 10px rgba(0,0,0,0.05);
    position: relative;
    display: flex;
    flex-direction: column;
  }

  .card-badge {
    position: absolute;
    top: 12px;
    right: 12px;
    background: #ff4757;
    color: white;
    padding: 0.25rem 0.6rem;
    border-radius: 20px;
    font-weight: bold;
    font-size: 0.85rem;
    box-shadow: 0 2px 6px rgba(255, 71, 87, 0.4);
    z-index: 2;
  }

  .card img {
    width: 100%;
    height: 150px;
    object-fit: cover;
  }

  .card-body {
    padding: 1rem;
    flex: 1;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
  }

  .card-title {
    font-size: 1.1rem;
    font-weight: bold;
    margin: 0 0 0.5rem 0;
    color: #333;
  }

  .card-price {
    color: #2ed573;
    font-weight: bold;
    font-size: 1.15rem;
    margin-bottom: 1rem;
  }

  .qty-controls {
    display: flex;
    align-items: center;
    justify-content: space-between;
    background: #f1f2f6;
    border-radius: 8px;
    padding: 0.4rem 0.8rem;
  }

  .qty-btn {
    background: white;
    border: 1px solid #dcdde1;
    color: #2f3542;
    width: 32px;
    height: 32px;
    border-radius: 6px;
    font-weight: bold;
    font-size: 1.1rem;
    cursor: pointer;
    box-shadow: 0 2px 4px rgba(0,0,0,0.05);
    display: flex;
    align-items: center;
    justify-content: center;
    transition: background 0.2s;
  }

  .qty-btn:hover {
    background: #dfe4ea;
  }

  input, button {
    padding: 0.6rem 1rem;
    border-radius: 6px;
    border: 1px solid #dcdde1;
    font-size: 0.95rem;
  }

  button {
    background: #2f3542;
    color: white;
    border: none;
    cursor: pointer;
    font-weight: 600;
  }

  button:hover {
    background: #57606f;
  }

  .checkout-container {
    background: white;
    padding: 2rem;
    border-radius: 12px;
    box-shadow: 0 4px 10px rgba(0,0,0,0.05);
    margin-bottom: 2rem;
  }

  .checkout-title {
    color: #2f3542 !important;
    font-size: 1.8rem;
    margin-top: 0;
  }

  .checkout-item {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0.8rem 0;
    border-bottom: 1px solid #f1f2f6;
    color: #333;
  }
</style>

<div class="container">
  {#if !isLoggedIn}
    <div class="login-box">
      <h2>🍔 Quick Delivery</h2>
      <p style="color: #666; margin-bottom: 1.5rem;">Enter your User ID to get started</p>
      <div style="display: flex; gap: 0.5rem; justify-content: center;">
        <input type="text" bind:value={userId} placeholder="e.g. user-1" />
        <button on:click={handleLogin}>Login</button>
      </div>
    </div>
  {:else}
    <div class="header-bar">
      <div>
        <h2 style="margin:0; color: #2f3542;">🍔 Food & Grocery Store</h2>
        <small style="color: #747d8c;">Logged in as: <strong>{userId}</strong></small>
      </div>
      <div class="nav-buttons">
        <button on:click={() => viewMode = "shop"} style="background: {viewMode === 'shop' ? '#2f3542' : '#a4b0be'};">Shop</button>
        <button on:click={() => viewMode = "checkout"} style="background: {viewMode === 'checkout' ? '#2ed573' : '#a4b0be'};">
          Checkout 🛒 ({cart.items ? cart.items.reduce((sum, i) => sum + i.quantity, 0) : 0})
        </button>
        <button on:click={() => isLoggedIn = false} style="background: #ff4757;">Logout</button>
      </div>
    </div>

    {#if viewMode === "shop"}
      <!-- Product Card Grid -->
      <h3 style="margin-bottom: 1rem;">Available Products</h3>
      <div class="grid">
        {#each products as prod}
          {@const qty = getCartItemQty(prod.id)}
          <div class="card">
            {#if qty > 0}
              <div class="card-badge">{qty} in cart</div>
            {/if}
            <img src={prod.image || "https://images.unsplash.com/photo-1546069901-ba9599a7e63c"} alt={prod.name} />
            
            <div class="card-body">
              <div>
                <div class="card-title">{prod.name}</div>
                <div class="card-price">${prod.price.toFixed(2)}</div>
              </div>

              <div class="qty-controls">
                <button class="qty-btn" on:click={() => updateCart(prod.id, -1)}>-</button>
                <span style="font-weight: bold; font-size: 1.1rem; color: #2f3542;">{qty}</span>
                <button class="qty-btn" on:click={() => updateCart(prod.id, 1)}>+</button>
              </div>
            </div>
          </div>
        {/each}
      </div>

    {:else}
      <!-- Checkout Screen -->
      <div class="checkout-container">
        <h2 class="checkout-title">🛍️ Order Checkout Summary</h2>
        
        {#if cart.items && cart.items.length > 0}
          <div style="margin-top: 1.5rem;">
            {#each cart.items as item}
              <div class="checkout-item">
                <div>
                  <strong style="color: #2f3542; font-size: 1.1rem;">{item.product.name}</strong>
                  <div style="color: #747d8c; font-size: 0.9rem;">Unit Price: ${item.product.price.toFixed(2)} × {item.quantity}</div>
                </div>
                <div style="font-weight: bold; font-size: 1.1rem; color: #2f3542;">
                  ${(item.product.price * item.quantity).toFixed(2)}
                </div>
              </div>
            {/each}
            
            <div style="display: flex; justify-content: space-between; align-items: center; margin-top: 2rem; padding-top: 1rem; border-top: 2px solid #2f3542;">
              <h2 style="margin:0; color: #2f3542;">Total Amount:</h2>
              <h2 style="margin:0; color: #2ed573;">${cart.total.toFixed(2)}</h2>
            </div>

            <button style="width: 100%; margin-top: 2rem; background: #2ed573; padding: 1rem; font-size: 1.1rem;" on:click={() => alert('Order placed successfully!')}>
              Confirm and Pay
            </button>
          </div>
        {:else}
          <div style="text-align: center; padding: 3rem 0;">
            <p style="color: #747d8c; font-size: 1.1rem; margin-bottom: 1.5rem;">Your cart is currently empty.</p>
            <button on:click={() => viewMode = "shop"} style="background: #2f3542; padding: 0.8rem 1.5rem;">
              Go Back to Shop
            </button>
          </div>
        {/if}
      </div>
    {/if}
  {/if}
</div>