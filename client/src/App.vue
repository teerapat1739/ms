<template>
  <div id="app">
    <vue-modaltor :visible="isOpenModal" @hideModal="hideModal">
      <template #header class="header-modal">
        <!--    add your custom header     -->
        <div>
          <i @click="hideModal" style="cursor: pointer" class="closeicon">x</i>
          <p>แจ้งเตือนการจ่ายเงิน</p>
        </div>
      </template>
      <template #body>
        <p>
          {{ message }}
        </p>
      </template>
    </vue-modaltor>
    <Nav />
    <router-view />
  </div>
</template>
<script>
import Nav from "./components/partials/Nav.vue";
export default {
  name: "app",
  components: {
    Nav,
  },
  data() {
    return {
      isOpenModal: false,
      message: "",
    };
  },
  created() {
    try {
      const ws = new WebSocket(`ws://${process.env.WS_SOCKET_NOTI_URL || 'localhost:5555'}/`);
      ws.onmessage = ({ data }) => {
        this.message = data;
        if (this.$auth.isAuthenticated) {
          this.isOpenModal = true;
        }
      };
    } catch (err) {
      throw(err)
    }
  },
  methods: {
    hideModal() {
      this.isOpenModal = false;
      this.message = "";
    },
  },
};
</script>
<style lang="scss">
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}
.header-modal {
  padding: 20px;
}
.closeicon {
  cursor: pointer;
  position: absolute;
  right: 10px;
}
</style>