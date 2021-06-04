<template>
  <div class="event-single">
    <div class="modal" v-if="showModalPaymeny" style="display: block">
      <div class="modal-background is-info is-light">
        <div class="modal-content">
          <img
            class="scb"
            src="https://m.thaiware.com/upload_misc/software/2018_10/thumbnails/4877_181009110327yj.jpg"
            alt=""
          />
          กำลังทำรายการ
        </div>
      </div>
    </div>
    <section class="hero is-primary">
      <div class="hero-body">
        <div class="container cont">
          <div>
            <h1 class="title">{{ event.name }}</h1>
            <h2 class="subtitle">
              <strong>Date:</strong>
              {{ event.date }}
              <br />
              <strong>Time:</strong>
              {{ event.time }}
            </h2>
          </div>
          <span class="btn-join">
            <button
              v-if="isJoin == false"
              class="button is-warning is-light"
              @click="joinParty()"
            >
              เข้าร่วม Party
            </button>
            <button
              v-else-if="isPaymentSuccess == false && isJoin"
              class="button is-success is-light"
            >
              กำลังดำเนินการ
            </button>
            <button v-else class="button is-success is-light">
              เข้าร่วม Party เรียบร้อยแล้ว
            </button>
          </span>
        </div>
      </div>
    </section>
    <section class="event-content">
      <div class="container">
        <p class="is-size-4 description">{{ event.description }}</p>
        <p class="is-size-5">
          <strong>Location:</strong>
          {{ event.location }}
        </p>
        <p class="is-size-5">
          <strong>Category:</strong>
          {{ event.category }}
        </p>
        <div class="event-images columns is-multiline has-text-centered">
          <div
            v-for="image in event.images"
            :key="image.id"
            class="column is-one-third"
          >
            <img :src="`${image}`" :alt="`${event.name}`" />
          </div>
        </div>
      </div>
    </section>
  </div>
</template>
<script>
// import EventService
import EventService from "@/services/EventService.js";
export default {
  name: "EventSingle",
  data() {
    // initialize the event object
    return {
      event: {},
      isJoin: false,
      isPaymentSuccess: false,
      showModalPaymeny: false,
    };
  },
  created() {
    this.getEventData();
  },
  methods: {
    async getEventData() {
      // Get the access token from the auth wrapper
      const accessToken = await this.$auth.getTokenSilently();

      // Use the eventService to call the getEventSingle method
      EventService.getEventSingle(
        this.$route.params.id,
        this.$auth.user.email,
        accessToken
      ).then(
        ((event) => {
          this.$set(this, "event", event);
        }).bind(this)
      );
    },
    async joinParty() {
      let obj = {
        partyID: parseInt(this.$route.params.id),
        userID: this.$auth.user.email,
      };
      const accessToken = await this.$auth.getTokenSilently();
      try {
        await EventService.joinParty(obj, accessToken);
        this.showModalPaymeny = true;
        setTimeout(async () => {
          await EventService.payment(
            this.event.name,
            this.$auth.user.email,
            accessToken
          );
          this.isPaymentSuccess = true;
          this.showModalPaymeny = false;
        }, 3000);

        this.isJoin = true;
      } catch (error) {
       throw(error)
      }
    },
  },
};
</script>
<style scoped>
.cont {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
@media only screen and (max-width: 480px) {
  .cont {
    display: block;
  }
}
.btn-join {
  margin-right: 40px;
}
/* .modal-background {
  background-color: white;
  width: 80%;
  margin: auto;
  padding: 10px;
} */
.modal-content,
.modal-card {
  height: 20%;
  /* margin: 0 20px; */
  max-height: calc(100vh - 160px);
  overflow: auto;
  position: relative;
  width: 100%;
  margin: auto;
  padding: 10px;
  width: 50%;
  background: wheat;
  display: flex;
  justify-content: space-around;
  align-items: center;
}
.scb {
  width: 100px;
}
.container{
  padding: 0 20px;
}
</style>