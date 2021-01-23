import { ref } from 'vue'

export const isLoading = ref(false)

export function startLoading () {
  isLoading.value = true
}

export function stopLoading () {
  isLoading.value = false
}
