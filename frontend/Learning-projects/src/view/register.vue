<template>
  <div class="auth-shell">
    <aside class="brand-panel">
      <div class="brand-mark">Schedula</div>

      <div class="clock-wrap" aria-hidden="true">
        <svg viewBox="0 0 240 240" class="clock-svg">
          <circle cx="120" cy="120" r="108" class="clock-ring" />
          <circle
            v-for="n in 12"
            :key="n"
            :cx="120 + 92 * Math.sin((n * Math.PI) / 6)"
            :cy="120 - 92 * Math.cos((n * Math.PI) / 6)"
            r="2.5"
            class="clock-tick"
          />
          <line x1="120" y1="120" x2="120" y2="62" class="clock-hand hour" />
          <line x1="120" y1="120" x2="168" y2="120" class="clock-hand minute" />
          <circle cx="120" cy="120" r="4" class="clock-pivot" />
        </svg>

        <div class="slot-card">
          <span class="slot-label">New this week</span>
          <span class="slot-time">32 bookings confirmed</span>
        </div>
      </div>

      <p class="brand-copy">
        Set up your business in minutes. Add services, invite your team, start booking.
      </p>
    </aside>

    <main class="form-panel">
      <div class="form-card">
        <h1 class="form-title">Create your account</h1>
        <p class="form-subtitle">Start taking bookings today.</p>

        <form @submit.prevent="handleRegister" novalidate>
          <div class="field">
            <label for="name">Full name</label>
            <input
              id="name"
              v-model="name"
              type="text"
              autocomplete="name"
              placeholder="Your name"
              :class="{ invalid: touched.name && !isNameValid }"
              @blur="touched.name = true"
            />
            <span v-if="touched.name && !isNameValid" class="field-error">
              Enter your full name.
            </span>
          </div>

          <div class="field">
            <label for="email">Email</label>
            <input
              id="email"
              v-model="email"
              type="email"
              autocomplete="email"
              placeholder="you@business.com"
              :class="{ invalid: touched.email && !isEmailValid }"
              @blur="touched.email = true"
            />
            <span v-if="touched.email && !isEmailValid" class="field-error">
              Enter a valid email address.
            </span>
          </div>

          <div class="field">
            <label for="password">Password</label>
            <div class="password-input">
              <input
                id="password"
                v-model="password"
                :type="showPassword ? 'text' : 'password'"
                autocomplete="new-password"
                placeholder="At least 6 characters"
                :class="{ invalid: touched.password && !isPasswordValid }"
                @blur="touched.password = true"
              />
              <button
                type="button"
                class="toggle-visibility"
                :aria-label="showPassword ? 'Hide password' : 'Show password'"
                @click="showPassword = !showPassword"
              >
                {{ showPassword ? 'Hide' : 'Show' }}
              </button>
            </div>
            <span v-if="touched.password && !isPasswordValid" class="field-error">
              Password must be at least 6 characters.
            </span>
          </div>

          <div class="field">
            <label for="confirm">Confirm password</label>
            <input
              id="confirm"
              v-model="confirmPassword"
              :type="showPassword ? 'text' : 'password'"
              autocomplete="new-password"
              placeholder="Re-enter password"
              :class="{ invalid: touched.confirm && !doPasswordsMatch }"
              @blur="touched.confirm = true"
            />
            <span v-if="touched.confirm && !doPasswordsMatch" class="field-error">
              Passwords don't match.
            </span>
          </div>

          <p v-if="errorMessage" class="form-error" role="alert">{{ errorMessage }}</p>

          <button type="submit" class="submit-btn" :disabled="isSubmitting || !canSubmit">
            <span v-if="isSubmitting" class="spinner" aria-hidden="true"></span>
            {{ isSubmitting ? 'Creating account…' : 'Create account' }}
          </button>
        </form>

        <p class="switch-line">
          Already have an account?
          <RouterLink to="/login" class="inline-link strong">Sign in</RouterLink>
        </p>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, reactive } from 'vue'
import { useAuthStore } from '@/store/auth'
import {useRouter} from 'vue-router'

const authStore = useAuthStore()
const router=useRouter()
const name = ref('')
const email = ref('')
const password = ref('')
const confirmPassword = ref('')
const showPassword = ref(false)
const isSubmitting = ref(false)
const errorMessage = ref('')
const touched = reactive({ name: false, email: false, password: false, confirm: false })

const isNameValid = computed(() => name.value.trim().length > 1)
const isEmailValid = computed(() => /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email.value))
const isPasswordValid = computed(() => password.value.length >= 6)
const doPasswordsMatch = computed(() => password.value === confirmPassword.value && confirmPassword.value.length > 0)
const canSubmit = computed(
  () => isNameValid.value && isEmailValid.value && isPasswordValid.value && doPasswordsMatch.value
)

const handleRegister = async () => {
  touched.name = true
  touched.email = true
  touched.password = true
  touched.confirm = true
  errorMessage.value = ''

  if (!canSubmit.value) return

  isSubmitting.value = true
  try {
    // NOTE: assumes authStore exposes a `register` action, e.g.
    // async register(payload) { ... same pattern as login ... }
    await authStore.register({ name: name.value, email: email.value, password: password.value })
    router.push('/dashboard')
  } catch (err) {
    const error = err as { response?: { data?: { message?: string } } }
    errorMessage.value = error.response?.data?.message || 'Could not create your account. Please try again.'
  } finally {
    isSubmitting.value = false
  }
}
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Fraunces:opsz,wght@9..144,500;9..144,600&family=Inter:wght@400;500;600&display=swap');

.auth-shell {
  min-height: 100vh;
  display: grid;
  grid-template-columns: 1fr 1fr;
  font-family: 'Inter', system-ui, sans-serif;
  background: #faf9f7;
}

.brand-panel {
  background: #14213d;
  color: #f4f2ee;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  padding: 3rem 2.5rem;
  position: relative;
  overflow: hidden;
}

.brand-mark {
  position: absolute;
  top: 2.5rem;
  left: 2.5rem;
  font-family: 'Fraunces', serif;
  font-size: 1.35rem;
  font-weight: 600;
  letter-spacing: 0.01em;
}

.clock-wrap {
  position: relative;
  width: 240px;
  height: 240px;
  margin: 1.5rem 0 2.5rem;
}

.clock-svg { width: 100%; height: 100%; }

.clock-ring {
  fill: none;
  stroke: rgba(244, 242, 238, 0.25);
  stroke-width: 1.5;
}

.clock-tick { fill: rgba(244, 242, 238, 0.55); }

.clock-hand {
  stroke: #e8a33d;
  stroke-width: 3.5;
  stroke-linecap: round;
  transform-origin: 120px 120px;
  animation: sweep 12s linear infinite;
}

.clock-hand.hour { animation-duration: 43200s; opacity: 0.85; }
.clock-hand.minute { animation-duration: 3600s; }

@keyframes sweep {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}

.clock-pivot { fill: #e8a33d; }

.slot-card {
  position: absolute;
  bottom: -0.5rem;
  right: -2.25rem;
  background: #f4f2ee;
  color: #14213d;
  border-radius: 10px;
  padding: 0.7rem 1rem;
  display: flex;
  flex-direction: column;
  gap: 0.15rem;
  box-shadow: 0 12px 24px rgba(0, 0, 0, 0.25);
  min-width: 168px;
}

.slot-label { font-size: 0.7rem; color: #5b6472; }
.slot-time { font-size: 0.85rem; font-weight: 600; }

.brand-copy {
  max-width: 26ch;
  text-align: center;
  font-family: 'Fraunces', serif;
  font-size: 1.3rem;
  line-height: 1.4;
  font-weight: 500;
  color: #f4f2ee;
}

.form-panel {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 2rem;
}

.form-card { width: 100%; max-width: 380px; }

.form-title {
  font-family: 'Fraunces', serif;
  font-size: 2rem;
  font-weight: 600;
  color: #14213d;
  margin: 0 0 0.4rem;
}

.form-subtitle {
  color: #5b6472;
  margin: 0 0 2rem;
  font-size: 0.95rem;
}

.field { margin-bottom: 1.15rem; }

label {
  display: block;
  font-size: 0.85rem;
  font-weight: 500;
  color: #14213d;
  margin-bottom: 0.4rem;
}

input[type='email'],
input[type='password'],
input[type='text'] {
  width: 100%;
  padding: 0.7rem 0.85rem;
  border: 1px solid #e4e2dc;
  border-radius: 8px;
  font-size: 0.95rem;
  font-family: inherit;
  background: #fff;
  box-sizing: border-box;
  transition: border-color 0.15s ease;
}

input:focus {
  outline: none;
  border-color: #14213d;
  box-shadow: 0 0 0 3px rgba(20, 33, 61, 0.08);
}

input.invalid { border-color: #c0392b; }

.password-input { position: relative; }
.password-input input { padding-right: 3.5rem; }

.toggle-visibility {
  position: absolute;
  right: 0.6rem;
  top: 50%;
  transform: translateY(-50%);
  background: none;
  border: none;
  color: #5b6472;
  font-size: 0.8rem;
  font-weight: 500;
  cursor: pointer;
  padding: 0.25rem 0.4rem;
}

.toggle-visibility:hover { color: #14213d; }

.field-error {
  display: block;
  color: #c0392b;
  font-size: 0.78rem;
  margin-top: 0.35rem;
}

.form-error {
  background: #fdecea;
  color: #c0392b;
  padding: 0.65rem 0.85rem;
  border-radius: 8px;
  font-size: 0.85rem;
  margin: 0 0 1.1rem;
}

.inline-link {
  color: #14213d;
  font-size: 0.82rem;
  text-decoration: none;
  border-bottom: 1px solid transparent;
}

.inline-link:hover { border-bottom-color: currentColor; }
.inline-link.strong { font-weight: 600; }

.submit-btn {
  width: 100%;
  background: #14213d;
  color: #f4f2ee;
  border: none;
  border-radius: 8px;
  padding: 0.8rem;
  font-size: 0.95rem;
  font-weight: 600;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.6rem;
  transition: background 0.15s ease, opacity 0.15s ease;
}

.submit-btn:hover:not(:disabled) { background: #1c2d52; }
.submit-btn:disabled { opacity: 0.55; cursor: not-allowed; }

.spinner {
  width: 15px;
  height: 15px;
  border: 2px solid rgba(244, 242, 238, 0.4);
  border-top-color: #f4f2ee;
  border-radius: 50%;
  animation: spin 0.7s linear infinite;
}

@keyframes spin { to { transform: rotate(360deg); } }

.switch-line {
  text-align: center;
  margin-top: 1.6rem;
  font-size: 0.88rem;
  color: #5b6472;
}

@media (max-width: 860px) {
  .auth-shell { grid-template-columns: 1fr; }
  .brand-panel { display: none; }
}

@media (prefers-reduced-motion: reduce) {
  .clock-hand { animation: none; }
}
</style>