<template>
  <div class="events container">
    <div class="header-party">
    <h2 class="subtitle is-3 all-party">ปาร์ตี้ทั้งหมด</h2>
    <router-link to="/addParty">
    <button class="subtitle is-6 btn-add-party">เพิ่มปาร์ตี้ใหม่</button>
    </router-link>
    </div>
    <div class="columns is-multiline">
      <div v-for="event in events" :event="event" :key="event.id" class="column is-one-quarter">
        <router-link :to="`/event/${event.id}`">
          <PartyCard :event="event" />
        </router-link>
      </div>
    </div>
  </div>
</template>
<script>
import PartyCard from "@/components/PartyCard";
import EventService from '@/services/EventService.js';
export default {
  name: "PartyList",
  components: {
    PartyCard
  },
  data() {
    return {
      event: {},
      events: []
    };
  },
  created() {
    this.getEventsData(); // call getEventData() when the instance is created
  },
  methods: {
    async getEventsData() {
      // Use the eventService to call the getEvents() method
      EventService.getEvents()
      .then(
        (events => {
          this.$set(this, "events", events);
        }).bind(this)
      );
    },

  }
};
</script>
<style lang="scss" scoped>
.events {
  margin-top: 100px;
  text-align: center;
}
.btn-add-party {
    background-color: #48c78e;
    border-color: transparent;
    color: #fff;
    cursor: pointer;
    justify-content: center;
    padding-bottom: calc(.5em - 1px);
    padding-left: 1em;
    padding-right: 1em;
    padding-top: calc(.5em - 1px);
    text-align: center;
    white-space: nowrap;
    border-radius: .375em;
    box-shadow: none;
    display: inline-flex;
    font-size: 1rem;
    height: 2.5em;
}
.header-party{
    display: flex;
    justify-content: space-between;
    margin-bottom: 10px;
}
.container {
    padding: 0 30px;
}

@media only screen and (max-width: 480px) {
  .all-party {
    font-size: 1rem !important;
  }
}

</style>