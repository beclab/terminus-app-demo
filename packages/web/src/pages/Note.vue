<script>
import { computed, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useLocalNotes } from 'src/helper'
import Container from 'src/components/Container.vue'

export default {
  components: { Container },
  setup() {
    const notes = useLocalNotes()
    const route = useRoute()
    const noteId = computed(() => parseInt(route.params.id))
    const editing = ref(false)

    const router = useRouter()
    const remove = () => {
      notes.value.write = notes.value.delete(notes.value.items[noteId.value]);
      router.push('/');
    }

    const update = () => {
      notes.value.write = notes.value.update(notes.value.items[noteId.value]);
      editing.value = false
    }

    return { notes, editing, noteId, remove, update }
  }
}
</script>

<template>
  <q-page padding>
    <Container>
      <div v-if="editing">
        <form @submit="update">
          <q-input v-model="notes.items[noteId].title" label="Title" filled />
          <q-input
            v-model="notes.items[noteId].description"
            label="Description"
            filled
            class="q-mt-sm"
            dense
          />

          <q-card flat bordered class="q-mt-sm">
            <q-editor v-model="notes.items[noteId].content" min-height="5rem" />
          </q-card>

          <div class="q-mt-md">
            <q-btn class="q-ml-sm" color="positive" type="submit"> Done </q-btn>
          </div>
        </form>
      </div>

      <div v-else>
        <div class="row items-center justify-between">
          <h3 class="q-mb-md q-mt-md">{{ notes.items[noteId]?.title }}</h3>
          <div>
            <q-btn
              round
              color="secondary"
              icon="edit"
              @click="editing = true"
            />
            <q-btn
              class="q-ml-sm"
              round
              color="red"
              icon="delete"
              @click="remove"
            />
          </div>
        </div>
        <div>{{ notes.items[noteId]?.description }}</div>
        <div class="q-mt-md" v-html="notes.items[noteId]?.content" />
      </div>
    </Container>
  </q-page>
</template>
