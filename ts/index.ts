'use strict';

import axios from 'axios';

const uuid:string = (<HTMLMetaElement>document.querySelector("meta[name='uuid']")).content;
const download:HTMLElement = document.getElementById('download');
console.log(uuid);

download.addEventListener('click', () => {
  const url:string = '/IsDownload?uuid=' + uuid;
  const DownloadFileURL:string = '/startDownload?uuid=' + uuid;
  const removeFileURL:string = '/remove?uuid=' + uuid;

  axios.get(
    url
  ).then(response => {

    if (response.data === true) {
      get(DownloadFileURL, removeFileURL)
    }else{
      alert('アップロードされたファイルが見つかりません');
    }

  }).catch(err => {
    console.log(err);
    alert('error');
  });

});

function get(url:string, remove:string) {
  const {data} :any = axios.get(
    url,
    { responseType: 'arraybuffer',
      headers: { Accept: 'application/zip' },
    }
  )

  const blob = new Blob([data], { type: 'application/zip' })

  const uri = URL.createObjectURL(blob)
  const link = document.createElement('a')
  link.download = uuid + '.zip'
  link.href = uri
  link.click()
  axios.get(remove)
}