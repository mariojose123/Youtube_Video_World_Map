var tag = document.createElement('script');
        
              tag.src = "https://www.youtube.com/iframe_api";
              var firstScriptTag = document.getElementsByTagName('script')[0];
              firstScriptTag.parentNode.insertBefore(tag, firstScriptTag);
        
              // 3. This function creates an <iframe> (and YouTube player)
              //    after the API code downloads.
              var player;
              function onYouTubeIframeAPIReady() {
                player = new YT.Player('player', {
                  height: '390',
                  width: '640',
                  videoId: 'M7lc1UVf-VE',
                  playerVars: {
                    'playsinline': 1
                  },
                  events: {
                    'onReady': onPlayerReady,
                    'onStateChange': onPlayerStateChange
                  }
                });
              }
        
              // 4. The API will call this function when the video player is ready.
              function onPlayerReady(event) {
                event.target.playVideo();
              }
        
              // 5. The API calls this function when the player's state changes.
              //    The function indicates that when playing a video (state=1),
              //    the player should play for six seconds and then stop.
              var done = false;
              function onPlayerStateChange(event) {
                if (event.data == YT.PlayerState.PLAYING && !done) {
                  done = true;
                }
              }
              function stopVideo() {
                player.stopVideo();
              }

    
      google.charts.load('current', {
        'packages':['geochart'],
      });
      google.charts.setOnLoadCallback(drawRegionsMap);


      var countriesdata = [['Country'],
      ['AF']
      ,['AX']
      ,['AL']
      ,['DZ']
      ,['AS']
      ,['AD']
      ,['AO']
      ,['AI']
      ,['AQ']
      ,['AG']
      ,['AR']
      ,['AM']
      ,['AW']
      ,['AU']
      ,['AT']
      ,['AZ']
      ,['BS']
      ,['BH']
      ,['BD']
      ,['BB']
      ,['BY']
      ,['BE']
      ,['BZ']
      ,['BJ']
      ,['BM']
      ,['BT']
      ,['BO']
      ,['BA']
      ,['BW']
      ,['BV']
      ,['BR']
      ,['IO']
      ,['BN']
      ,['BG']
      ,['BF']
      ,['BI']
      ,['KH']
      ,['CM']
      ,['CA']
      ,['CV']
      ,['KY']
      ,['CF']
      ,['TD']
      ,['CL']
      ,['CN']
      ,['CX']
      ,['CC']
      ,['CO']
      ,['KM']
      ,['CG']
      ,['CD']
      ,['CK']
      ,['CR']
      ,['CI']
      ,['HR']
      ,['CU']
      ,['CY']
      ,['CZ']
      ,['DK']
      ,['DJ']
      ,['DM']
      ,['DO']
      ,['EC']
      ,['EG']
      ,['SV']
      ,['GQ']
      ,['ER']
      ,['EE']
      ,['ET']
      ,['FK']
      ,['FO']
      ,['FJ']
      ,['FI']
      ,['FR']
      ,['GF']
      ,['PF']
      ,['TF']
      ,['GA']
      ,['GM']
      ,['GE']
      ,['DE']
      ,['GH']
      ,['GI']
      ,['GR']
      ,['GL']
      ,['GD']
      ,['GP']
      ,['GU']
      ,['GT']
      ,['GG']
      ,['GN']
      ,['GW']
      ,['GY']
      ,['HT']
      ,['HM']
      ,['VA']
      ,['HN']
      ,['HK']
      ,['HU']
      ,['IS']
      ,['IN']
      ,['ID']
      ,['IR']
      ,['IQ']
      ,['IE']
      ,['IM']
      ,['IL']
      ,['IT']
      ,['JM']
      ,['JP']
      ,['JE']
      ,['JO']
      ,['KZ']
      ,['KE']
      ,['KI']
      ,['KR']
      ,['KW']
      ,['KG']
      ,['LA']
      ,['LV']
      ,['LB']
      ,['LS']
      ,['LR']
      ,['LY']
      ,['LI']
      ,['LT']
      ,['LU']
      ,['MO']
      ,['MK']
      ,['MG']
      ,['MW']
      ,['MY']
      ,['MV']
      ,['ML']
      ,['MT']
      ,['MH']
      ,['MQ']
      ,['MR']
      ,['MU']
      ,['YT']
      ,['MX']
      ,['FM']
      ,['MD']
      ,['MC']
      ,['MN']
      ,['ME']
      ,['MS']
      ,['MA']
      ,['MZ']
      ,['MM']
      ,['NA']
      ,['NR']
      ,['NP']
      ,['NL']
      ,['AN']
      ,['NC']
      ,['NZ']
      ,['NI']
      ,['NE']
      ,['NG']
      ,['NU']
      ,['NF']
      ,['MP']
      ,['NO']
      ,['OM']
      ,['PK']
      ,['PW']
      ,['PS']
      ,['PA']
      ,['PG']
      ,['PY']
      ,['PE']
      ,['PH']
      ,['PN']
      ,['PL']
      ,['PT']
      ,['PR']
      ,['QA']
      ,['RE']
      ,['RO']
      ,['RU']
      ,['RW']
      ,['BL']
      ,['SH']
      ,['KN']
      ,['LC']
      ,['MF']
      ,['PM']
      ,['VC']
      ,['WS']
      ,['SM']
      ,['ST']
      ,['SA']
      ,['SN']
      ,['RS']
      ,['SC']
      ,['SL']
      ,['SG']
      ,['SK']
      ,['SI']
      ,['SB']
      ,['SO']
      ,['ZA']
      ,['GS']
      ,['ES']
      ,['LK']
      ,['SD']
      ,['SR']
      ,['SJ']
      ,['SZ']
      ,['SE']
      ,['CH']
      ,['SY']
      ,['TW']
      ,['TJ']
      ,['TZ']
      ,['TH']
      ,['TL']
      ,['TG']
      ,['TK']
      ,['TO']
      ,['TT']
      ,['TN']
      ,['TR']
      ,['TM']
      ,['TC']
      ,['TV']
      ,['UG']
      ,['UA']
      ,['AE']
      ,['GB']
      ,['US']
      ,['UM']
      ,['UY']
      ,['UZ']
      ,['VU']
      ,['VE']
      ,['VN']
      ,['VG']
      ,['VI']
      ,['WF']
      ,['EH']
      ,['YE']
      ,['ZM']
      ,['ZW']]
      function drawRegionsMap() {
       
        var selectcountry = function(selectedCountry){
          return function(country){
          if(country=='Country'){
            return country.push("selected")
          }
          if(country==selectedCountry){
            return country.push(1)
          }
          return country.push(0)
        }
        }
        countriesdata.every(selectcountry())
        var data = google.visualization.arrayToDataTable(countriesdata
        );

        var options = {colors:["white","green"],legend: 'none'};

        var chart = new google.visualization.GeoChart(document.getElementById('regions_div'));
        
        var previousSelectedCountry = ""
        var indexpreviousSelected 
        initcache()
        addcache()
        async function selectHandler() {
          var selectedItem = chart.getSelection()[0];
          if (selectedItem) {    
              
            var country = data.getValue(selectedItem.row, 0);
            if(previousSelectedCountry!=""){
              data.setValue(indexpreviousSelected,1,0)
            }
            indexpreviousSelected = selectedItem.row
            previousSelectedCountry = country
            data.setValue(selectedItem.row,1,1)
            if(sessionStorage.getItem(country)==[]){
            try{
            const videojson = await fetchcountryvideos(country)
            var randomindex  = Math.floor(Math.random() * Object.keys(videojson).length);
            player.loadVideoById(videojson[randomindex]['VideoID'])
            }catch{
            }
          }else{
            player.loadVideoById(sessionStorage.getItem(country))
              addcache()
          }
            chart.draw(data,options)
          }
        }
        google.visualization.events.addListener(chart, 'select', selectHandler);    
        chart.draw(data, options);
      
      }

      function initcache(){
        countriesdata.forEach((item) => {
          sessionStorage.setItem(item,[])
        })
      }
      
      //addvideos to cache
      async function addcache(){
        videos = await fetchEveryCountryvideos()
        videos.forEach((video) => {
          sessionStorage.setItem(video["Iso2"],video["VideoID"] )
        })
      }

      //get 1 random video from every country to get cache
      async function fetchEveryCountryvideos(){
        return fetch("/game/everycountryvideo", {
          method: 'GET',
          mode: 'same-origin'
        }).then(res=> {
          return res.json()})
        }


        async function fetchcountryvideos(country){
          return fetch('/game/'+country+"/50", {
          method: 'GET',
          mode: 'same-origin'
        }).then(res=> {
          return res.json()})
        }
  