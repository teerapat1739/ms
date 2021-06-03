import axios from "axios"

export default {
  async getEvents() {
    let res = await axios.get(`http://${process.env.API_GATEWAY_URL || "localhost:8000"}/public/api-party/party`);
    return res.data.data || [];
  },
  async getPrivateEvents(email, accessToken) {
    let res = await axios({
      method: 'get',
      url: `http://${process.env.API_GATEWAY_URL || "localhost:8000"}/private/api-party/party`,
      headers: {
        Authorization: `Bearer ${accessToken}`,
        email: email
      }
    });
    return res.data.data || [];
  },
  async getEventSingle(eventId, email, accessToken) {
    let { data } = await axios.get(`http://${process.env.API_GATEWAY_URL || "localhost:8000"}/private/api-party/party/` + eventId, {
      headers: {
        Authorization: `Bearer ${accessToken}`,
        email: email,
      }
    });
    return data.data;
  },
  async addParty(data, accessToken) {
    let res = await axios({
      method: 'post',
      url: `http://${process.env.API_GATEWAY_URL || "localhost:8000"}/private/api-party/party`,
      headers: {
        Authorization: `Bearer ${accessToken}`
      },
      data
    });
    return res.data;
  },
  async joinParty(data, accessToken) {
    let res = await axios({
      method: 'post',
      url: `http://${process.env.API_GATEWAY_URL || "localhost:8000"}/private/api-party/joinParty`,
      headers: {
        Authorization: `Bearer ${accessToken}`
      },
      data
    });
    return res.data;
  },
  async payment(eventName, email, accessToken) {
    let res = await axios({
      method: 'post',
      url: `http://${process.env.API_GATEWAY_URL || "localhost:8000"}/private/api-payment/payment`,
      headers: {
        Authorization: `Bearer ${accessToken}`
      },
      data: {
        email: email,
        id: eventName
      }
    });
    return res.data;
  }
}