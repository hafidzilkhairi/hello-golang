<h1>Tugas 1 Netpro</h1>

<center>
<h2>Anggota Kelompok</h2>
<ol>
    <li>Hafidzil Khairi (1301160171)</li>
    <li>Yola Adipratama</li>
</ol>
<br/>
<br/>
<h2>Jawaban</h2>
</center>
<ol>

<h4><li>Penjelasan Finite State </li></h4>
<ul>
<li>langkah awal dari starting point adalah send SYN </li>
<li> ada dua kemungkinan, menerima flag syn atau terabaikan </li>
<li> jika menerima flag syn dan juga menerima ack maka koneksi terhubung dengan mengirimkan flag ack juga </li>
<li> jika hanya menerima flag syn maka akan masuk ke state active close dan time wait </li>
<li> ketika koneksi yang terhubung telah selesai melakukan data transfer maka juga akan memasuki state active close dan time wait </li>
<li> setelah time wait dan mencapai timeout, maka koneksi terputus </li>
</ul>

<h4><li>For loop dan conditional if/else </li></h4>
<ul>
<li>foor loop melakukan perulangan dengan jumlah yang telah di tentukan, dan memiliki titik henti lain, yaitu ketika ada break</li>
<li>conditional syntax adalah if else, yaitu menerima sebuah perbandingan logika dan melakukan baris kode sesuai hasil perbandingan tersebut</li>
</ul>
<center>
<img src="https://github.com/hafidzilkhairi/hello-golang/blob/task/Screenshots/for-loop.png?raw=true">
<p>Hasil running dari kodingan for loop</p>
<img src="https://github.com/hafidzilkhairi/hello-golang/blob/task/Screenshots/conditional.png?raw=true">
<p>Hasil running dari kodingan conditional if/else</p>
</center>

<h4><li>For loop dan conditional if/else</li></h4>
<ul>
<li>array adalah keadan dimana suatu variable mempunyai index serta nilai dalam index tersebut</li>
<li>function merupakan baris baris kode yang dideklarasikan dengan syntax func sehingga dapat dipaggil berulang kali</li>
</ul>
<center>
<img src="https://github.com/hafidzilkhairi/hello-golang/blob/task/Screenshots/array.png?raw=true">
<p>Hasil running dari kodingan array</p>
<img src="https://github.com/hafidzilkhairi/hello-golang/blob/task/Screenshots/function.png?raw=true">
<p>Hasil running dari kodingan function</p>
</center>

<h4><li>Array dan Function</li></h4>
<ul>
<li>array adalah keadan dimana suatu variable mempunyai index serta nilai dalam index tersebut</li>
<li>function merupakan baris baris kode yang dideklarasikan dengan syntax func sehingga dapat dipaggil berulang kali</li>
</ul>
<center>
<img src="https://github.com/hafidzilkhairi/hello-golang/blob/task/Screenshots/array.png?raw=true">
<p>Hasil running dari kodingan array</p>
<img src="https://github.com/hafidzilkhairi/hello-golang/blob/task/Screenshots/function.png?raw=true">
<p>Hasil running dari kodingan function</p>
</center>

<h4><li>Struct dan Method</li></h4>
<ul>
<li>struct merupakan type data bentukan yang isi typenya dapat dikostumisasi</li>
<li>method merupakan suatu fungsi yang terdapat didalam struct atau kelas/class</li>
</ul>
<center>
<img src="https://github.com/hafidzilkhairi/hello-golang/blob/task/Screenshots/struct.png?raw=true">
<h4><li>Struct dan Metnod</li></h4>
<p>Hasil running dari kodingan array</p>
<img src="https://github.com/hafidzilkhairi/hello-golang/blob/task/Screenshots/method.png?raw=true">
<h4><li>Struct dan Method</li></h4>
<p>Hasil running dari kodingan function</p>
</center>

<h4><li>Multiple Return Value dan Command Line</li></h4>
<ul>
<li>multiple return merupakan suatu keadaan dimana fungsi dapat mengembalikan nilai lebih dari 1 nilai</li>
<li>command line merupakan keadaan dimana suatu nilai dapat di oper dari inputan user ke dalam kodingan ketika pemanggilan program</li>
</ul>
<center>
<img src="https://github.com/hafidzilkhairi/hello-golang/blob/task/Screenshots/multiple-return.png?raw=true">
<p>Hasil running dari kodingan Multiple Return</p>
<img src="https://github.com/hafidzilkhairi/hello-golang/blob/task/Screenshots/command-line.png?raw=true">
<p>Hasil running dari kodingan Command Line</p>
</center>

<h4><li>Simple Web Application</li></h4>
<ul>
<li>simple web application berfungsi untuk mengadakan layanan server sehingga user dapat mengaksesnya melalui protokol http
</li>
<li>command line merupakan keadaan dimana suatu nilai dapat di oper dari inputan user ke dalam kodingan ketika pemanggilan program</li>
</ul>
<center>
<img src="https://github.com/hafidzilkhairi/hello-golang/blob/task/Screenshots/simple-web-application.png?raw=true">
<p>Hasil running dari kodingan Simple Web Application</p>
</center>

<h4><li>Simple Web Application dengan config file</li></h4>
<ul>
<li>simple web application berfungsi untuk mengadakan layanan server sehingga user dapat mengaksesnya melalui protokol http
</li>
<li>command line merupakan keadaan dimana suatu nilai dapat di oper dari inputan user ke dalam kodingan ketika pemanggilan program</li>
</ul>
<center>
<img src="https://github.com/hafidzilkhairi/hello-golang/blob/task/Screenshots/simple-web-application-with-viper.png?raw=true">
<p>Direktori dari web app dengan config file</p>
<img src="https://github.com/hafidzilkhairi/hello-golang/blob/task/Screenshots/simple-web-application.png?raw=true">
<p>Hasil running dari kodingan Simple Web Application</p>
</center>


</ol>