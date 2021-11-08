<template>
	<div v-if="initialized">
		<SpWallet ref="wallet" v-on:dropdown-opened="$refs.menu.closeDropdown()" />
		<SpLayout>
			<template v-slot:sidebar>
				<Sidebar />
			</template>
			<template v-slot:content>
				<router-view />
			</template>
		</SpLayout>
	</div>
</template>

<style>
body {
	 margin: 0;
	 padding: auto;
	 color: white;
}
.sp-token-send__holder .sp-token-send .sp-token-send__main__form .sp-token-send__main__rcpt__wrapper .sp-token-send__main__rcpt__memo__btn,.sp-token-send__holder .sp-token-send .sp-token-send__main__form .sp-amount-select .sp-form-group .sp-amount-select__denom .sp-amount-select__denom__selected
{
	color: #fff;
	opacity: 1;
}
 #app {
	 background: #171622;
}
 .sp-sidebar {
	 height: 100vh;
	 background:  #171622;
	 color: white;
	 font-size: 16px;
	 display: none;
}
 .sp-box {
	 background: #212130;
}
 .sp-button {
    background: #E649EE;
    border: none;
}
 .sp-fill {
	 padding-top: 11rem;
}
 .sp-wallet-create {
	 width: 35rem;
	 max-height: 95vh;
	 background:  #171622;
	 overflow: auto;
}
 .sp-card.sp-card-primary {
	 background: #E649EE;
    border: none;
}
 .sp-button.sp-button-primary {
	 background: #FEC007!important;
	 border: none!important;
}
 .sp-button.sp-button-secondary {
	 background: #eb5757;
	 color: white;
	 border: none;
}
 .sp-link-icon .sp-link-icon-wrapper {
	 color: white;
}
 .sp-transfer-list .sp-transfer-list__table .sp-transfer-list__status .sp-transfer-list__status__wrapper .sp-transfer-list__status__icon.sp-transfer-list__status__icon__empty {
	 color: white;
}
 .sp-component-title span {
	 display: block;
	 opacity: initial;
}
 .sp-transfer-list .sp-transfer-list__table .sp-transfer-list__status .sp-transfer-list__status__wrapper .sp-transfer-list__status__action .sp-transfer-list__status__action__text {
	 opacity: initial;
}
 .sp-transfer-list .sp-transfer-list__table .sp-transfer-list__status .sp-transfer-list__status__wrapper .sp-transfer-list__status__action .sp-transfer-list__status__action__date {
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
 .sp-transfer-list .sp-transfer-list__table .sp-transfer-list__status .sp-transfer-list__status__wrapper .sp-transfer-list__status__icon.sp-transfer-list__status__icon__empty {
	 background: #FEC007;
}
 .sp-button:disabled {
	 cursor: not-allowed;
	 color: white !important;
	 opacity: 0.5;
	
}
.sp-token-send__holder .sp-token-send .sp-token-send__main__form .sp-amount-select .sp-form-group .sp-amount-select__denom .sp-amount-select__denom__modal
{
	background: #1f203f;
}
.sp-token-send__holder .sp-token-send .sp-token-send__main__form .sp-amount-select .sp-form-group .sp-amount-select__denom .sp-amount-select__denom__modal .sp-amount-select__denom__modal__search .sp-amount-select__denom__modal__search__input,.sp-token-send__holder .sp-token-send .sp-token-send__main__form .sp-amount-select .sp-form-group .sp-amount-select__denom .sp-amount-select__denom__modal .sp-amount-select__denom__modal__search .sp-icon,.sp-token-send__holder .sp-token-send .sp-token-send__main__form .sp-token-send__main__footer.sp-token-send__main__footer__open .sp-token-send__main__btns .sp-token-send__main__btns__reset__fees,.sp-token-send__holder .sp-token-send .sp-token-send__main__form .sp-token-send__main__footer .sp-token-send__main__btns .sp-token-send__main__btns__tx .sp-token-send__main__btns__reset
 {
	 opacity: 1
 }
 .sp-wallet-create .sp-wallet-create__back .sp-icon {
	 opacity: initial;
}
 .sp-link-icon .sp-link-icon-wrapper {
	 color: white;
	 opacity: initial;
}
 .sp-card.sp-card-secondary {
	 background: #E649EE;
     border: none;
	 color:#fff;
}
 .sp-textarea {
	 background: #a1a6f5;
}
 .sp-input {
	 background: #34344B;
}
 .sp-token-send__holder .sp-token-send .sp-token-send__main__form .sp-token-send__main__rcpt__wrapper .sp-token-send__main__rcpt__icon {
	color: #181631;
    background: #FEC007;
}
 .sp-box .sp-box-header {
	 opacity: initial;
}
 .sp-token-send__holder .sp-token-send .sp-token-send__main__form .sp-token-send__main__footer .sp-box-message {
	 opacity: initial;
}
 .sp-token-send__holder .sp-assets__wrapper .sp-assets .sp-assets__main__header {
	 opacity: initial;
}
 .custom-header {
	 display: flex;
	 position: absolute;
	 top: 20px;
	 left: 20px;
	 z-index: 99;
    background: #181631;
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
	 height: 40px;
	 width: 200px;
	 margin-right: 10px;
	 margin-top: -5px
}
 @media (max-width: 767px) {
	 .custom-header {
		 top: 15px;
		 left: 5px;
	}
	.sp-transfer-list .sp-transfer-list__table td.sp-transfer-list__table__address
	{
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
		 top: 0;
		 right: 0;
	}
}
 @media (max-width: 500px) {
	 .custom-header {
		 display: none;
	}
	.sp-wallet,.sp-wallet-menu.sp-opened {
    width: 100%;
}
.sp-wallet-menu
{
	max-width: 100%;
}
}

 .sp-wallet-menu {
	 background:  #171622;
}
 .sp-header-text {
	 color: white;
}
 input {
	 color: black;
}
.sp-dummy-fill {
    background: rgb(17 17 17);
    border-radius: .4rem;
    height: .8rem;
    flex: 1;
}
.sp-token-send__holder .sp-token-send .sp-token-send__main__form .sp-amount-select .sp-form-group .sp-input:last-child{
color:white;
border:none !important;
}

.sp-input .sp-input-large{
color:#fff;
}

.sp-form-group .sp-input:first-child{
border:none !important;
}

.sp-token-send__holder{
     display: flex!important;
    flex-wrap: wrap!important;
    margin-right: -15px;
    margin-left: -15px;
    width: 100%;
	flex-direction:row !important;
}
.sp-token-send__holder .sp-token-send{
    flex: 0 0 50%!important;
    max-width: 50%!important;
    position: relative!important;
    padding-right: 15px;
    padding-left: 15px;
}
.sp-component.sp-assets__wrapper{
    flex: 0 0 50%!important;
    max-width: 50%!important;
    position: relative!important;
    padding-right: 15px;
    padding-left: 15px;

}

.sp-assets__main{
background: #17A2B8 !important;
} 

input{
color:#fff !important;
}

.sp-transfer-list .sp-transfer-list__table .sp-transfer-list__status .sp-transfer-list__status__wrapper .sp-transfer-list__status__icon.sp-transfer-list__status__icon__received{
    color: #181631;
    background: #FEC007;
}

.sp-shadow{
box-shadow:none !important;
}
</style>

<script>
import './scss/app.scss'
import '@starport/vue/lib/starport-vue.css'
import Sidebar from './components/Sidebar'

export default {
	components: {
		Sidebar
	},
	data() {
		return {
			initialized: false
		}
	},
	computed: {
		hasWallet() {
			return this.$store.hasModule(['common', 'wallet'])
		}
	},
	async created() {
		await this.$store.dispatch('common/env/init')
		this.initialized = true
		await this.$store.dispatch("cosmonaut.voter.voter/QueryPollAll",{options:{subscribe:true, all:true},params:{}})
        await this.$store.dispatch("cosmonaut.voter.voter/QueryVoteAll",{options:{subscribe:true, all:true},params:{}})
	},
	errorCaptured(err) {
		console.log(err)
		return false
	}
}
</script>
