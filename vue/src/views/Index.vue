<template>
	<div>
		<div class="custom-header">
		<img :src="image" />
		<p>Crypstarter</p>
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
   		image:require("/public/Mask Group.svg"),

   		}
   	},
	computed: {
		address() {
			return this.$store.getters['common/wallet/address']
		}
	}
}
</script>
