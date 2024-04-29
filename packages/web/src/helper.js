import { ref, watch, onMounted, onUnmounted } from 'vue'
import axios from 'axios';

// const useLocalStorage = (key, defaultValue) => {
//   const value = ref(defaultValue)
//   const read = () => {
//     const v = window.localStorage.getItem(key)
//     if (v != null) value.value = JSON.parse(v)
//   }

//   read()

//   onMounted(() => {
//     window.addEventListener('storage', read)
//   })
//   onUnmounted(() => {
//     window.removeEventListener('storage', read)
//   })

//   const write = () => {
//     window.localStorage.setItem(key, JSON.stringify(value.value))
//   }
//   watch([value], write, { deep: true })

//   return value
// }



const useRemoteStorage = (defaultValue) => {
  const url = '/api';

  const valueRef = ref(defaultValue);
  
  const read = async ()=>{

    const res = await axios.get(`${url}/list`);
    var valueStore = {
      items: res.data,
      write: null,
    
      create: (note)=>{
        return async () => {
          console.log(`create a new note, ${note.title}}`);
          await axios.post(`${url}/create`, note);
        }
      },

      delete: (note) => {
        return async()=>{
          console.log(`delete note, ${note.title}`)
          await axios.post(`${url}/delete`, note);
        }
      },

      update: (note) => {
        return async()=>{
          console.log(`update note, ${note.title}`)
          await axios.post(`${url}/update`, note);
        }
      },

    }  

    valueRef.value = valueStore;
    console.log('read notes from remote server, ', res.data);
  }

  
  onMounted(async () => {
    window.addEventListener('notes-write', read)
  })
  onUnmounted(() => {
    window.removeEventListener('notes-write', read)
  })


  const write = async () => {
    if (valueRef.value.write){
      await valueRef.value.write();
      window.dispatchEvent(new Event('notes-write'));
    }
  }

  watch([valueRef], write, { deep: true })

  read();

  return valueRef;
}

export const useLocalNotes = () => {
  return useRemoteStorage({items:[]});
}

