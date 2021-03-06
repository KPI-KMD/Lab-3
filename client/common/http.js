const request = require('request');

const Client = baseURL => {
  return {
    get: (tabletid, path) => {
      return new Promise((resolve, reject)=>{
        request(`${baseURL}${path}`, {json: true, body: tabletid}, (err, res, body) => {
          if(err) {
            reject(err);
          } else {
            resolve(body);
          }
        });
      });
    },
    post: (path, data) => {
      return new Promise((resolve, reject)=> {
        request.post(`${baseURL}${path}`, {json: true, body: data}, (err, res, body) => {
          if(err) {
            reject(err);
          } else {
            resolve(body);
          };
        })
      })
    }
  }
}

module.exports = {Client};
