<template>
  <div class="home">
    <section class="hero is-dark">
      <div class="hero-body"></div>
    </section>
    <PartyList />
  </div>
</template>
<script>
import PartyList from "../components/PartyList";

export default {
  name: "home",
  components: {
    PartyList,
  },
   data() {
    return {
      message: "",
    }
  },
  methods: {
    // Log the user in
    login() {
      this.$auth.loginWithRedirect();
    },
  },
};
</script>
<style lang="scss" scoped>
.hero {
  text-align: center;
  background-image: url("https://scontent.fbkk28-1.fna.fbcdn.net/v/t1.6435-9/143543101_467388584641766_3583936560886207390_n.png?_nc_cat=104&ccb=1-3&_nc_sid=e3f864&_nc_eui2=AeHHhX5KduC0ih-XZEC_EYpOdh0MCQXBO_F2HQwJBcE78WiTk9U63XYfTSELAHNQPrCROfISijPjqeiXjvBdubwo&_nc_ohc=CB7onCFoPHsAX_ltBRP&_nc_ht=scontent.fbkk28-1.fna&oh=10571919e5c97c029054d4434a185f4d&oe=60D5A162");
  background-size: cover;
  background-position: center;
  background-repeat: no-repeat;
  min-height: 100px;
  height: 620px;
}
.hero-body .title {
  text-shadow: 4px 4px 4px rgba(0, 0, 0, 0.6);
  padding: 40px 0 20px 0;
  font-size: 60px;
}
.subtitle {
  text-shadow: 4px 4px 4px rgba(0, 0, 0, 0.7);
  font-size: 30px;
}
.button-block {
  text-align: center;
  margin-left: auto;
  margin-right: auto;
  width: 100%;
  position: absolute;
  bottom: -150px;
  .button {
    margin-right: 50px;
    padding-left: 50px;
    padding-right: 50px;
  }
  .welcome {
    width: 400px;
    padding: 10px;
    margin-left: auto;
    margin-right: auto;
  }
}
.is-xl {
  font-size: 1.7rem;
}
</style>