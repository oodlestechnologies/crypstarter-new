<template>
	<div v-if="initialized">
		<span class="cnbcnb" @click="openNav()"><img :src="hamberger" /></span>
		<div id="mySidenav" class="sidenav">
			<div class="m-v-logo">
				<img :src="image" />
			</div>
			<div class="c-bond">
				<router-link to="/">
					<img :src="dashboard" />
					<p>Dashboard</p>
				</router-link>
			</div>
			<div class="c-bond">
				<router-link to="/staking">
					<img :src="staking" />
					<p>Staking</p>
				</router-link>
			</div>
			<div class="c-bond">
				<router-link to="/Unstaking">
					<img :src="unstaking" />
					<p>UnStaking</p>
				</router-link>
			</div>
			<div class="c-bond">
				<router-link to="/Withdraw">
					<img :src="withdraw" />
					<p>Withdraw Reward</p>
				</router-link>
			</div>
			<hr class="horizontal-row" />
			<br />
			<div class="c-bond">
				<a href="http://explorer.crypstarter.network/" target="_blank">
					<img :src="explorer" />
					<p>Explorer</p>
				</a>
			</div>
			<div class="c-bond">
				<a href="https://docs.crypstarter.network/" target="_blank">
					<img :src="documents" />
					<p>Documents</p>
				</a>
			</div>
		</div>

		

		<SpWallet ref="wallet" v-on:dropdown-opened="$refs.menu.closeDropdown()" />
		<SpLayout>
			<template v-slot:sidebar>
				<Sidebar />
			</template>
			<template v-slot:content>
				<router-view />
				<!-- <Footer /> -->
			</template>
		</SpLayout>
	</div>
</template>
<style>
body {
	margin: 0;
	padding: auto;
	color: white;
	background: #171622;
}

/*  */
.horizontal-row {
	margin-left: 30px;
	margin-right: 40px;
	height: 1px;
	border-width: 0;
	color: #171622;
	background-color: #171622;
}

.m-v-logo img{
	width: 360px;
    margin-left: 28px;
    height: 70px;
	margin-bottom: 5px;
}

.cnbcnb img {
	width: 50px;
	position: fixed;
	left: 10px;
	top: 10px;
	z-index: 9;
}

.sidenav {
	height: 100%;
	width: 0;
	position: fixed;
	z-index: 1;
	top: 20;
	left: 0;
	background-color: #fec007;
	overflow-x: hidden;
	transition: 0.5s;
	padding-top: 20px;
	margin-top: 80px;
	border-top-right-radius: 10px;
	border-bottom-right-radius: 10px;
}

.sidenav a {
	padding: 8px 8px 8px 32px;
	text-decoration: none;
	font-size: 25px;
	color: #818181;
	display: block;
	transition: 0.3s;
}

.sidenav a:hover,
.offcanvas a:focus {
	color: #f1f1f1;
}

.sidenav .closebtn {
	position: absolute;
	top: 0;
	right: 25px;
	font-size: 36px;
	margin-left: 50px;
}

@media screen and (max-height: 450px) {
	.sidenav {
		padding-top: 15px;
	}
	.sidenav a {
		font-size: 18px;
	}
}

.c-bond {
	padding-bottom: 20px;
}
.c-bond a {
	text-decoration: none;
}
.c-bond img {
	margin-bottom: 30px;
}
.c-bond p {
	font-size: 18px;
	margin-top: -60px;
	margin-left: 67px;
	color: #171622;
	font-weight: 500;
}

/*  */
.custom-footer {
	width: 100%;
	left: 0;
	position: absolute;
	bottom: 0;
	font-size: 14px;
	text-align: center;
	padding: 10px;
	background: #212130;
	/* margin-top: 10px; */
	/* margin-left: -200px; */
	z-index: 9999999;
}
* {
	transition: 0.3s all ease-in-out;
}
.sp-token-send__holder .sp-token-send .sp-token-send__main__form .sp-token-send__main__rcpt__wrapper .sp-token-send__main__rcpt__memo__btn,
.sp-token-send__holder .sp-token-send .sp-token-send__main__form .sp-amount-select .sp-form-group .sp-amount-select__denom .sp-amount-select__denom__selected {
	color: #fff;
	opacity: 1;
}
#app {
	background: #171622;
}
.sp-sidebar {
	background: #fec007;
	color: #171622;
	font-size: 16px;
	justify-content: start;
	border-radius: 20px;
	/* width: 275px; */
	width: 325px;
	margin: 15px;
	/* position: fixed; */
	height: 96vh;
	top:0;
	min-height: 400px;
	/* display: none; */
}

/* .sp-link-icon__text {
    color: black;
}
.sp-link-icon__icon {
    color: black;
} */
.sp-sidebar .sp-hamburger {
	display: block;
	background: transparent;
	position: relative;
	z-index: 999;
	/* left: 200px; */
	left: 250px;
	top: -15px;
}
.sp-sidebar.sp-opened .sp-sidebar__content,
.sp-sidebar.sp-opened .sp-sidebar__footer,
.sp-sidebar.sp-opened .sp-sidebar__header {
	width: 100%;
}
.sp-sidebar.sp-mob-opened {
	width: 102px;
	padding-left: 4rem;
	padding-right: 4rem;
}
.sp-sidebar.sp-mob-opened .sp-hamburger {
	left: 75px;
	color: white;
}
.sp-sidebar.sp-mob-opened .custom-header {
	/* left: 40px; */
	left: 25px;
	z-index: 99;

	/* width: 28px; */
	width: 40px;
	overflow: hidden;
}
.sp-box {
	background: #212130;
}
.sp-button {
	background: #fec007;
	border: none;
	height: 51px;
	color: #171622;
}
.sp-button.sp-button-primary {
	color: #000;
}
.sp-fill {
	padding-top: 11rem;
	padding-left: 35px;
	/* margin-left: 332px; */
}
.sp-transfer-list
{
	padding-bottom: 5rem;
}
.sp-sidebar__content
{
    max-height: 62vh;
    /* overflow: AUTO!important; */
}
.sp-sidebar.sp-mob-opened + .sp-fill
{
/* margin-left: 102px; */
}
.sp-wallet-create {
	width: 35rem;
	max-height: 95vh;
	background: #171622;
	overflow: auto;
}
.sp-card.sp-card-primary {
	/* background: #E649EE; */
	background: #fec007;
	border: none;
	color: #171622;
}
.sp-button.sp-button-primary {
	background: #fec007 !important;
	border: none !important;
}
.sp-button.sp-button-secondary {
	background: #eb5757;
	color: white;
	border: none;
}
.sp-link-icon .sp-link-icon-wrapper {
	color: white;
}
/* .sp-transfer-list .sp-transfer-list__table .sp-transfer-list__status .sp-transfer-list__status__wrapper .sp-transfer-list__status__icon.sp-transfer-list__status__icon__empty {
	 color: white;
} */
.sp-component-title span {
	display: block;
	opacity: initial;
}
.sp-transfer-list
	.sp-transfer-list__table
	.sp-transfer-list__status
	.sp-transfer-list__status__wrapper
	.sp-transfer-list__status__action
	.sp-transfer-list__status__action__text {
	opacity: initial;
}
.sp-transfer-list
	.sp-transfer-list__table
	.sp-transfer-list__status
	.sp-transfer-list__status__wrapper
	.sp-transfer-list__status__action
	.sp-transfer-list__status__action__date {
	opacity: initial;
}
.sp-status-api .sp-status-api__text {
	opacity: initial;
}
.sp-status-rpc .sp-status-rpc__text {
	opacity: initial;
}
.sp-status-ws .sp-status-ws__text {
	opacity: initial;
}
.sp-relayers__holder .sp-relayers__main .sp-relayers__main__message {
	opacity: initial;
}
.sp-relayers__holder .sp-relayers__add .sp-relayers__add__main__message {
	opacity: initial;
}
.sp-welcome .sp-welcome__main {
	z-index: auto;
}
.sp-transfer-list
	.sp-transfer-list__table
	.sp-transfer-list__status
	.sp-transfer-list__status__wrapper
	.sp-transfer-list__status__icon.sp-transfer-list__status__icon__empty {
	background: #fec007;
	color: black;
	height: 39px;
	width: 39px;
}
.sp-button:disabled {
	cursor: not-allowed;
	color: #171622 !important;
	opacity: 0.5;
}
.sp-token-send__holder .sp-token-send .sp-token-send__main__form .sp-amount-select .sp-form-group .sp-amount-select__denom .sp-amount-select__denom__modal {
	background: #1f203f;
}
.sp-token-send__holder
	.sp-token-send
	.sp-token-send__main__form
	.sp-amount-select
	.sp-form-group
	.sp-amount-select__denom
	.sp-amount-select__denom__modal
	.sp-amount-select__denom__modal__search
	.sp-amount-select__denom__modal__search__input,
.sp-token-send__holder
	.sp-token-send
	.sp-token-send__main__form
	.sp-amount-select
	.sp-form-group
	.sp-amount-select__denom
	.sp-amount-select__denom__modal
	.sp-amount-select__denom__modal__search
	.sp-icon,
.sp-token-send__holder
	.sp-token-send
	.sp-token-send__main__form
	.sp-token-send__main__footer.sp-token-send__main__footer__open
	.sp-token-send__main__btns
	.sp-token-send__main__btns__reset__fees,
.sp-token-send__holder
	.sp-token-send
	.sp-token-send__main__form
	.sp-token-send__main__footer
	.sp-token-send__main__btns
	.sp-token-send__main__btns__tx
	.sp-token-send__main__btns__reset {
	opacity: 1;
}
.sp-wallet-create .sp-wallet-create__back .sp-icon {
	opacity: initial;
}
.sp-link-icon .sp-link-icon-wrapper {
	color: white;
	opacity: initial;
}
.sp-card.sp-card-secondary {
	background: #fec007;
	border: none;
	/* color:#fff; */
	color: #171622;
}
.sp-textarea {
	background: #a1a6f5;
}
.sp-input {
	background: #34344b;
}
.sp-token-send__holder .sp-token-send .sp-token-send__main__form .sp-token-send__main__rcpt__wrapper .sp-token-send__main__rcpt__icon {
	color: #181631;
	background: #fec007;
	height: 38px;
	width: 38px;
	margin-bottom: 28px;
}
.sp-box .sp-box-header {
	opacity: initial;
}
.sp-token-send__holder .sp-token-send .sp-token-send__main__form .sp-token-send__main__footer .sp-box-message {
	opacity: initial;
}
.sp-token-send__holder .sp-assets__wrapper .sp-assets .sp-assets__main__header {
	opacity: initial;
	/* font-size: 20px; */
}
.custom-header {
	display: flex;
	position: absolute;
	top: 20px;
	left: 20px;
	z-index: 99;
	width: 100%;
	overflow: hidden;
	/* background: #171622; */
}

.custom-header1 {
	display: flex;
	position: absolute;
	top: 20px;
	left: 20px;
	z-index: 99;
}
.sp-sidebar .sp-link-icon .sp-link-icon-wrapper .sp-link-icon__icon {
	font-size: 2rem;
}
.custom-header p {
	position: static;
	vertical-align: middle;
	font-size: 24px;
	font-weight: 600;
	margin: 0;
}

.custom-header img {
	vertical-align: middle;
	/* height: 35px;
	 width: 190px; */
	height: 38px;
	width: 240px;
	/* margin-right: 10px; */
	margin-top: -3px;
	margin-left: 10px;
}

.sp-wallet-menu {
	background: #171622;
}
.sp-header-text {
	color: white;
}
input {
	color: black;
}
.sp-dummy-fill {
	display: none;
}
.sp-denom-marker {
	/* display: none; */
	background: transparent !important;
	position: relative;
	width: 0;
}
.mystylezaxa .sp-assets__main__item {
	margin-bottom: 30px;
}
.sp-assets__main__item {
	position: relative;
	margin-bottom: 80px;
	width: 70%;
}
.mystylezaxa .sp-assets__main__item {
	width: 100%;
	position: relative;
}
.mystylezaxa .sp-assets__main__item .sp-assets__main__denom__name:after {
	content: 'Token';
	position: absolute;
	left: 0;
	font-weight: 600;
	font-size: 24px;
	top: -50px;
	width: 100%;
}

.mystylezaxa .sp-assets__main__item .sp-assets__main__denom__balance:after {
	content: 'Amount';
	position: absolute;
	left: 0;
	font-weight: 600;
	font-size: 24px;
	top: -50px;
	width: 100%;
}
.mystylezaxa .sp-assets__main__header {
	padding-top: 50px;
	margin-bottom: 75px;
	position: relative;
}

.mystylezaxa .sp-assets__main__header:after {
	content: 'Total Amount';
	position: absolute;
	left: 0;
	font-weight: 600;
	font-size: 24px;
	top: 0px;
}
.sp-wallet
{
	top: 2rem;
    right: 2rem;
    z-index: 90;
  
  	position: fixed;
}
/* .mystylezaxa .sp-assets__main__header:before{
content: "CST";
	position: absolute;
	left: 0;
	font-weight: 600;
	font-size: 24px;
	top: 0px;
} */

.sp-assets__main__header .sp-assets__main__denom__balance {
	position: relative;
	font-size: 32px;
}
.sp-assets__main__header .sp-assets__main__denom__name {
	position: absolute;
	display: none;
	background: transparent;
}
.sp-assets__main > .sp-assets__main__item ~ .sp-assets__main__item ~ .sp-assets__main__item .sp-assets__main__denom__balance:after {
	display: none;
}

.sp-assets__main > .sp-assets__main__item ~ .sp-assets__main__item ~ .sp-assets__main__item .sp-assets__main__denom__name:after {
	display: none;
}

.sp-assets__main > .sp-assets__main__item ~ .sp-assets__main__item .sp-assets__main__denom__balance:after {
	display: none;
}
.sp-assets__main > .sp-assets__main__item ~ .sp-assets__main__item .sp-assets__main__denom__name:after {
	display: none;
}
.sp-assets__main__header .sp-assets__main__denom__balance::after {
	content: 'CST';
	position: absolute;
	margin-left: 10px;
	font-weight: 600;
	font-size: 18px;
	top: 0;
}

.sp-assets__main > .sp-assets__main__item:after {
	content: 'N/A';
	position: absolute;
	left: 0;
	font-weight: 600;
	font-size: 40px;
	top: 0;
}

.mystylezaxa .sp-assets__main > .sp-assets__main__item ~ .sp-assets__main__item:after {
	display: none;
}
.mystylezaxa .sp-assets__main > .sp-assets__main__item:after {
	display: none;
}

.sp-assets__main > .sp-assets__main__item {
	display: block;
}
.sp-assets__main > .sp-assets__main__item ~ .sp-assets__main__item:after {
	content: 'Please access your Wallet';
	position: absolute;
	left: 0;
	font-weight: 400;
	font-size: 16px;
	top: 0;
}
.sp-assets__main > .sp-assets__main__item ~ .sp-assets__main__item ~ .sp-assets__main__item {
	display: none !important;
}
/* .sp-assets__main__item:last-child
{
	display: none
} */
.sp-token-send__holder .sp-token-send .sp-token-send__main__form .sp-amount-select .sp-form-group .sp-input:last-child {
	color: white;
	border: none !important;
}

.sp-input .sp-input-large {
	color: #fff;
}

.sp-form-group .sp-input:first-child {
	border: none !important;
}
.sp-token-send__holder {
	display: flex !important;
	flex-wrap: wrap !important;
	margin-right: -15px;
	margin-left: -15px;
	width: 100%;
	flex-direction: row !important;
}

.sp-token-send__holder .sp-token-send {
	/* flex: 0 0 50%!important; */
	/* max-width: 50%; */
	width: 65%;
	position: relative !important;
	padding-right: 15px;
	padding-left: 15px;
}

.sp-component.sp-assets__wrapper {
	flex: 0 0 50% !important;
	max-width: 35%;
	position: relative !important;
	padding-right: 15px;
	padding-left: 15px;
}

.sp-assets__main {
	/* background: #ffbf09 !important; */
	background-image: url(../public/New_card_Resolutionsss.svg);
	color: #171622;
	background-repeat: no-repeat;
	background-size: cover;
	/* height: 180px; */
	height: 100%;
	font-weight: 600;
	font-size: 22px;
	padding: 35px;
	/* width: 100%; */
}

input {
	color: #fff !important;
}

.sp-transfer-list
	.sp-transfer-list__table
	.sp-transfer-list__status
	.sp-transfer-list__status__wrapper
	.sp-transfer-list__status__icon.sp-transfer-list__status__icon__received {
	color: #181631;
	background: #fec007;
}
.sp-mob-opened .orugnsasd{
	width: 0;
	overflow: hidden;
}
.sp-shadow {
	box-shadow: none !important;
}

.sp-amount-select__denom {
	display: none !important;
}

.sp-text.sp-bold.sp-active {
	color: #fec007;
}

.sp-icon.sp-icon-Lock {
	color: #fec007;
}

.sp-icon.sp-icon-Unlock {
	color: #fec007;
}

.custom-header1 a {
	color: #fff;
	text-decoration: none;
}

span.sp-icon.sp-icon-AddNew {
	color: white !important;
}

/* .sp-wallet .sp-wallet-menu-action .sp-link-icon .sp-link-icon__text{
	color: red !important;
} */

.sp-sidebar__content .sp-link-icon__icon {
	color: #171622;
}

.sp-sidebar__content .sp-link-icon__text {
	color: #171622;
}

.sp-assets__main__header__message:after {
	content: 'TOTAL AMOUNT';
	color: #000;
	background: #ffbf09;
	position: absolute;
	left: 0;
	top: 0;
	font-size: 16px;
	/* color: #b89f16; */
}
.sp-assets__main__header__message {
	position: relative;
	color: #ffbf09;
}
.sp-token-send__holder .sp-assets__wrapper .sp-assets .sp-assets__main__item .sp-assets__main__denom__name {
	/* margin-left: -17px; */
	font-size: 16px;
	font-weight: 600;
}

.sp-transfer-list__table th,
.sp-transfer-list__table td {
	border-bottom: 1px solid #5d5d82 !important;
	opacity: 1;
	color: #dcdcdc;
}

table {
	border-spacing: 0px !important;
}

.sp-transfer-list .sp-transfer-list__table .sp-transfer-list__status {
	text-align: left;
	padding-top: 20px;
	padding-bottom: 20px;
}

.sp-transfer-list .sp-transfer-list__table {
	width: 97%;
}

.sp-transfer-list .sp-transfer-list__table th {
	opacity: 1;
}

.sp-token-send__holder .sp-assets__wrapper .sp-assets .sp-assets__main__item .sp-assets__main__denom__balance {
	font-size: 16px;
	font-weight: 600;
}

.sp-assets__main__header__token,
.sp-assets__main__header__amount {
	font-size: 20px;
	font-weight: bold;
}

.sp-assets__main__header {
	/* border-bottom: 1px solid white; */
	padding-bottom: 10px;
}

.sp-token-send__holder .sp-token-send .sp-token-send__main__form .sp-token-send__main__rcpt__wrapper .sp-token-send__main__rcpt__memo__btn {
	margin-bottom: 25px !important;
}
@media (min-width: 900px) {
	.cnbcnb {
		display: none;
		width: 50px;
	}
}

@media (max-width: 900px) {
	.sp-sidebar {
		display: none;
		padding-top: 1rem;
		width: 50px;
		padding-left: 10px;
	}
}

@media (max-width: 900px) {
	.sp-token-send__holder .sp-token-send {
		width: 100% !important;
	}
	.sp-token-send__holder {
		flex-direction: column-reverse !important;
	}
	.sp-fill {
		padding-top: 90px;
		    margin-left: 0px;
	}
	.custom-header {
		top: 10px;
		left: 16px;
		width: 29px;
	}

	.sp-sidebar .sp-hamburger {
		opacity: 0;
	}
	.sp-sidebar {
		display: none;
		padding-top: 1rem;
		width: 50px;
		padding-left: 10px;
	}
	.sp-sidebar .sp-link-icon .sp-link-icon-wrapper .sp-link-icon__text {
		display: none;
	}
	.sp-transfer-list .sp-transfer-list__table td.sp-transfer-list__table__address {
		word-break: break-all;
	}
	.custom-header p {
		width: 40px;
		overflow: hidden;
		white-space: nowrap;
	}
	.sp-wallet-create {
		width: 100%;
		max-height: 100vh;
	}
	.sp-wallet {
		top: 7px;
		right: 6px;
		padding-left: 50px;
	}
	.container {
		margin: auto;
		margin-left: 70px;
		padding: 0;
	}
	.sp-token-send__holder .sp-component.sp-assets__wrapper {
		flex: 0 0 100% !important;
		max-width: 100% !important;
	}

	.sp-token-send__holder .sp-component.sp-token-send {
		flex: 0 0 100% !important;
		max-width: 100% !important;
	}
	.custom-header {
		flex: 0 0 100% !important;
		max-width: 100% !important;
	}
	.custom-header1 {
		top: 5px !important;
	}

	.copyright p {
		display: none;
	}
}
</style>

<script>
import './scss/app.scss'
import '@starport/vue/lib/starport-vue.css'
import Sidebar from './components/Sidebar'
// import Footer from './components/Footer'

function numberWithCommas(x) {
	if (isNaN(x)) {
		return 0
	}

	return x.toLocaleString(undefined, { minimumFractionDigits: 2, maximumFractionDigits: 6 })

	//   return x.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ',');
	// return x.toLocaleString()

	// if(x===NaN){
	// 	console.log(x+"second")
	// 	return 0

	// }else{
	// 	console.log(x+"third")
	// 		return x.toLocaleString(undefined, { minimumFractionDigits: 2, maximumFractionDigits: 6 })

	// }
}
	var gridcell0 = document.getElementsByClassName('sp-text')
			var gridcell2 = document.getElementsByClassName('sp-assets__main__denom__balance')
			var gridcell1 = document.getElementsByClassName('sp-assets__main__item')

function myFUnciton(){
		// setTimeout(() => {
			
		let gridcell0temp = gridcell0
		let gridcell2temp = gridcell2
		let gridcell1temp = gridcell1
			// console.log( gridcell0[1].getAttribute('title') );
			let myAddress;
			if(gridcell0temp[1]){
			 myAddress = gridcell0temp[1].getAttribute('title')
			}
			//   console.log(myAddress);
			if (myAddress == null) {
				document.getElementById('app').classList.remove('mystylezaxa')

				if (document.getElementById('bts2')) {
					document.getElementById('bts2').classList.add('mystyle')
					document.getElementById('bts1').classList.remove('mystyle')
				}
				if(document.getElementById('validatoraddress2')){
					document.getElementById('validatoraddress2').classList.add('mystyle')
					document.getElementById('validatoraddress1').classList.remove('mystyle')

				}
				if(document.getElementById('amount2')){
					document.getElementById('amount2').classList.add('mystyle')
					document.getElementById('amount1').classList.remove('mystyle')

				}

				if(document.getElementById('withdraw')){
					document.getElementById('withdraw').classList.remove('mystyle')

				}

			} else {
					if (document.getElementById('bts2')) {
					document.getElementById('bts1').classList.add('mystyle')
					document.getElementById('bts2').classList.remove('mystyle')
				}

				if(document.getElementById('validatoraddress2')){
					document.getElementById('validatoraddress1').classList.add('mystyle')
					document.getElementById('validatoraddress2').classList.remove('mystyle')

				}

				if(document.getElementById('amount2')){
					document.getElementById('amount1').classList.add('mystyle')
					document.getElementById('amount2').classList.remove('mystyle')

				}

				if(document.getElementById('withdraw')){
					document.getElementById('withdraw').classList.add('mystyle')

				}


				document.getElementById('app').classList.add('mystylezaxa')
					var sta,stas,stass
				if (gridcell2temp.length > 0) {
				
				
					
					// for(var i=0; i<gridcell2temp.length; i++){
					
					// 			sta = gridcell2temp[i].innerHTML	
					// 			localStorage.setItem('',sta)
					// 			stas = numberWithDivision(sta)
					// 			stass = numberWithCommas(stas)
					// 			//console.log(stass)
					// 			gridcell2temp[i].innerHTML = stass
								
					//}
					
					var divs = gridcell1temp[0].innerHTML
					document.getElementsByClassName('sp-assets__main__header')[0].innerHTML = divs
					
				} 
			}
				
		// },1000)

	}

function numberWithDivision(p1) {
	const p2 = 1000000
	return p1 / p2
}

document.addEventListener('click', function () {
	if (parseInt(document.querySelector('#mySidenav').offsetWidth) > 0) {
		document.getElementById('mySidenav').style.width = '0'
	}
})

// document.getElementsByClassName(".c-bond").click(function(){
//             document.getElementsByClassName(".c-bond").collapse('hide');
//         });

// console.log(numberWithCommas(123456789987456321));

// function openNav() {
//     console.log("open");
//     document.getElementById("check").style.width = "250px";
// }

// function closeNav() {
//     document.getElementById("check").style.width = "0";
// }

// document.addEventListener("click", function(){
//   if( parseInt( document.getElementById().css('width') ) > 0 ){
//     document.getElementById("mySidenav").style.width = "0";
//   }
// });

export default {
	components: {
		Sidebar
		// Footer
	},

	methods: {
		openNav: function () {
			document.getElementById('mySidenav').style.width = '400px'
		},
		closeNav: function () {
			document.getElementById('mySidenav').style.width = '0'
		}
	},

	data() {
		setInterval(()=>{
			myFUnciton()
		// var elements = document.getElementsByClassName("sp-button-primary");
		// 	var elements2 = document.getElementsByClassName("sp-wallet-menu-item__status");
		// for(let i=0 ; i< elements.length -1;i++)
	 	// elements[i].addEventListener('click', myFUnciton, true); 
		// //  document.getElementsByClassName("sp-button-primary");
		//  if(elements2.length > 0){
		//   for(let i=0 ; i< elements2.length -1;i++)
	 	// 		elements2[i].addEventListener('click', myFUnciton, true); 
		 
		 },1000)

		
		//  document.getElementsByClassName("sp-button-primary");
		 
		//  },1000)
	// 
	//  myFUnciton()
		//
		//
		return {
			initialized: false,
			image: require('/public/Wide_logo-black.png'),
			dashboard: require('/public/dashboard.svg'),
			bond: require('/public/bond.svg'),
			staking: require('/public/staking.svg'),
			unstaking: require('/public/Unstaking.svg'),
			withdraw: require('/public/withdraw.svg'),
			explorer: require('/public/explorer.svg'),
			forum: require('/public/forum.svg'),
			governance: require('/public/governance.svg'),
			documents: require('/public/documents.svg'),
			hamberger: require('/public/menu_white_24dp.svg')
		}
	},
	computed: {
		hasWallet() {
			// myFUnciton()
			return this.$store.hasModule(['common', 'wallet'])
		}
	},
	async created() {
		// myFUnciton()
		await this.$store.dispatch('common/env/init')
		this.initialized = true
		await this.$store.dispatch('cosmonaut.voter.voter/QueryPollAll', { options: { subscribe: true, all: true }, params: {} })
		await this.$store.dispatch('cosmonaut.voter.voter/QueryVoteAll', { options: { subscribe: true, all: true }, params: {} })
	},
	errorCaptured(err) {
		console.log(err)
		return false
	}
}
</script>
