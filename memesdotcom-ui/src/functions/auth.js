import { ref } from 'vue'
import { startLoading, stopLoading } from '@/functions/spinner'
import router from '@/router'

const localStorageValue = localStorage.getItem('isUserLogged')

export const isUserLogged = ref(localStorageValue === 'true' || false)

export function login () {
  startLoading()
  setTimeout(async () => {
    localStorage.setItem('isUserLogged', 'true')
    isUserLogged.value = true
    await router.push({ name: 'feed' })
    stopLoading()
  }, 1000)
}

export function logout () {
  startLoading()
  setTimeout(async () => {
    localStorage.setItem('isUserLogged', 'false')
    isUserLogged.value = false
    await router.push({ name: 'login' })
    stopLoading()
  }, 1000)
}
