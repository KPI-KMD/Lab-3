const https = require('request');

const Client = baseURL => {
  return {
    get: path => {
      return new Promise((resolve, reject)=>{
        request(`${baseUrl}${path}`, {json: true}, (err, res, body) => {
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
        request.post(`${baseUrl}${path}`, {json: true, body: data}, (err, res, body) => {
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
