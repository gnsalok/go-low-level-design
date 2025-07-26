# Rate Limiter Design

A **rate limiter** is used to control how frequently a particular action can be performed. In the context of HTTP servers or APIs, it helps prevent abuse by limiting the number of requests a client can make in a given time window.

## How the Rate Limiter Works

The typical Go rate limiter uses one of these strategies:

- **Token Bucket:** Allows a burst of requests up to a certain limit, then refills at a steady rate.
- **Leaky Bucket:** Processes requests at a fixed rate, queuing excess requests.
- **Fixed Window Counter:** Counts requests in fixed time windows.
- **Sliding Window:** Smooths out the fixed window approach for more accurate limiting.

### Example: Token Bucket Implementation

### **Explanation:**

1. **Capture the Current Time:**
   ```go
   now := time.Now()
   ```
   - The current time is stored in the `now` variable.

2. **Calculate Elapsed Time:**
   ```go
   elapsed := now.Sub(rl.lastRefill)
   ```
   - The `elapsed` variable represents the time that has passed since the last refill. This is calculated by subtracting `rl.lastRefill` (the time when the bucket was last refilled) from `now`.

3. **Check if Refill is Needed:**
   ```go
   if elapsed >= rl.refillInterval {
   ```
   - The function checks if the elapsed time is greater than or equal to the `refillInterval`. The `refillInterval` defines how often tokens should be added to the bucket (e.g., every 1 second).

4. **Calculate Refill Tokens:**
   ```go
   refillTokens := int(elapsed / rl.refillInterval) * rl.refillRate
   ```
   - This line calculates how many tokens should be added based on the elapsed time.
   - `elapsed / rl.refillInterval` gives the number of intervals that have passed since the last refill.
   - `refillRate` is the number of tokens added per interval.
   - `refillTokens` is the total number of tokens that should be added.

   **Example:**
   - Suppose the `refillInterval` is 1 second, and the `refillRate` is 2 tokens per second.
   - If 3 seconds have elapsed since the last refill, `refillTokens` will be calculated as:
     ```go
     refillTokens = int(3 * time.Second / 1 * time.Second) * 2 = 6 tokens
     ```

5. **Add Tokens to the Bucket:**
   ```go
   rl.tokens = min(rl.capacity, rl.tokens+refillTokens)
   ```
   - The new token count is calculated by adding `refillTokens` to the current `tokens`.
   - However, the bucket’s capacity should not be exceeded, so `min(rl.capacity, rl.tokens+refillTokens)` ensures that the token count does not exceed the maximum capacity.

6. **Update Last Refill Time:**
   ```go
   rl.lastRefill = now
   ```
   - Finally, the `lastRefill` timestamp is updated to the current time (`now`).
