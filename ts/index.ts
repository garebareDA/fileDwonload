'use strict';

import axios from 'axios';

const uuid:string = (<HTMLMetaElement>document.querySelector("meta[name='uuid']")).content;
const download:HTMLElement = document.getElementById('download');
console.log(uuid);

download.addEventListener('click', () => {
  const url:string = '/IsDownload?uuid=' + uuid;
  const downloadURL:string = '/zip/' + uuid + '.zip';
  const removeFileURL:string = '/remove?uuid=' + uuid;

  axios.get(
    url
  ).then(response => {

    if (response.data === true) {
      const link:HTMLElement = document.getElementById('link');
      const a:HTMLAnchorElement = document.createElement('a')
      a.href = downloadURL;
      a.innerText = 'ダウンロード';
      link.appendChild(a);
      a.addEventListener('click', () =>{get(removeFileURL)})
    }else{
      alert('アップロードされたファイルが見つかりません');
    }

  }).catch(err => {
    console.log(err);
    alert('error');
  });

});

function get(url:string) {
  axios.get(
    url
  ).then(response => {
    console.log(response)
  }).catch(err => {
    console.log(err)
  });
}