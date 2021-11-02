<template>
	<div>
		<div class="custom-header">
		<img :src="image" />
		</div>
		<div class="container">

			<SpTokenSend :address="address" />
			<SpTransferList :address="address" />
		</div>
	</div>
</template>

<script>
export default {
	name: 'Index',

	   data() {
   		return {
   		image:require("/public/Crypstarter Logo in SVG.svg"),

   		}
   	},
	computed: {
		address() {
			return this.$store.getters['common/wallet/address']
		}
	}
}
</script>
