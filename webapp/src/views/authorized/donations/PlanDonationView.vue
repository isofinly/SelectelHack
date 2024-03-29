<script setup>
import InfoIcon from '@/assets/icons/InfoIcon.vue'

import Datepicker from '@vuepic/vue-datepicker';
import '@vuepic/vue-datepicker/dist/main.css';

import UISelectorButton from '@/components/ui/UISelectorButton.vue';
import UIDropdownWithSearch from '@/components/ui/UIDropdownWithSearch.vue';
import { MainButton } from 'vue-tg';
</script>
<template>
    <section class="bg-gray-50 dark:bg-gray-900">
        <div class="flex flex-col items-center justify-center px-6 py-8 mx-auto md:h-screen lg:py-0">
            <div class="w-full bg-white rounded-lg shadow dark:border md:mt-0 sm:max-w-md xl:p-0 dark:bg-gray-800 dark:border-gray-700">
                <div class="p-6 space-y-4 md:space-y-6 sm:p-8">
                    <h1 class="text-xl font-bold leading-tight tracking-tight text-gray-900 md:text-2xl dark:text-white text-center">
                        Напоминание о донации
                    </h1>
                    <div class="space-y-4 md:space-y-6" action="#">
                        <div>
                            <b>Выберите тип донации</b><br>
                            <span>После выбора типа донации автоматически отобразится ближайшая доступная дата с учётом интервалов между донациями</span>
                            <UIDropdownWithSearch :options="donations" v-model="blood_type" class="mt-1"/>
                        </div>
                        <div>
                            <b>Дата донации</b>
                            <Datepicker v-model="date"
                            locale="ru-RU"
                            class="mt-1"
                            :enable-time-picker="false"
                            :min-date="new Date()"
                            format="dd.MM.yyyy"/>
                        </div>
                        <div>
                            <b>Тип донации</b>
                            <div class="grid gap-2 mb-6 grid-cols-2">
                                <UISelectorButton
                                    name="Безвозмездно"
                                    @select="selectType(0)"
                                    :active="type == 0"
                                    class="mt-2"
                                    :features="[
                                        'Питание или компенсация питания.',
                                        '5% МРОТ порядка 700-1500 ₽.',
                                        'Учитывается при получении звания Почетного донора.'
                                    ]"/>
                                <UISelectorButton
                                    name="Платно"
                                    @select="selectType(1)"
                                    :active="type == 1"
                                    class="mt-2"
                                    :features="[
                                        'Деньги или социальная поддержка.'
                                    ]"
                                    :debuffs="[
                                        'Не учитывается при получении звания почетного донора.'
                                    ]"/>
                            </div>
                        </div>
                        <div>
                            <b>Место сдачи</b>
                            <div class="grid gap-2 mb-6 grid-cols-2">
                                <UISelectorButton
                                    name="Стационарный пункт"
                                    @select="selectPlace(0)"
                                    :active="place == 0"
                                    class="mt-2"
                                    :features="[
                                        'Центр крови.',
                                        'Станция переливания.'
                                    ]"/>
                                <UISelectorButton
                                    name="Выездная акция"
                                    @select="selectPlace(1)"
                                    :active="place == 1"
                                    class="mt-2"
                                    :features="[
                                        'День донора.',
                                        'Выезды в ВУЗы.',
                                        'Предвижные мобильные бригады.'
                                    ]"/>
                            </div>
                        </div>
                        <div>
                            <b>Город</b>
                            <UIDropdownWithSearch :options="cities" v-model="city" @changed="updateCenters" class="mt-1">
                                Выберите город
                            </UIDropdownWithSearch>
                        </div>
                        <div v-if="place == 0">
                            <b>Центр крови</b><br>
                            <span>Важно: если ваш центр крови принимает по записи, то нужно отдельно записаться на сайте центра крови или через Госуслуги. Планирование донации на сайте DonorSearch позволит нам за 3 дня до указанной даты напомнить о вашем намерении совершить донацию и подготовиться к ней.</span>
                            <UIDropdownWithSearch :options="centers" v-model="center" class="mt-1">
                                Выберите центр крови
                            </UIDropdownWithSearch>
                        </div>
                        <div class="flex items-center p-4 mb-4 text-sm text-red-800 border border-red-300 rounded-lg bg-red-50 dark:bg-gray-800 dark:text-red-400 dark:border-red-800" role="alert">
                            <InfoIcon/>
                            <span class="sr-only">Info</span>
                            <div>
                                <span class="font-medium">Планирование не означает запись на донацию в центр крови.</span>
                            </div>
                        </div>
                        <b></b>
                    </div>
                </div>
            </div>
        </div>
    </section>
    <MainButton @click="sendToBot" text="Создать напоминание"></MainButton>
</template>

<script>
import { RegionService } from '@/services/RegionService';
import { StationService } from '@/services/StationService';
import { useWebApp } from 'vue-tg';
export default {
    name: 'CreateDonationView',
    components: {
        MainButton
    },
    data() {
        return {
            date: new Date(),
            blood_type: "blood",
            donations: {
                "blood": 'Цельная кровь',
                "plasma": 'Плазма',
                "platelets": 'Тромбоциты',
                "erythrocytes": 'Эритроциты',
                "leukocytes": 'Гранулоциты'
            },
            type: 0,
            place: 0,

            cities: {},
            centers: {},
            city: 0,
            center: 0
        }
    },
    methods: {
        selectType(type) {
            this.type = type
        },
        selectPlace(place) {
            this.place = place
        },
        updateCenters() {
            if(this.city < 1)
                return;
            StationService.getStations(this.city, (data) => {
                this.centers = data.results.reduce((acc, cur) => {
                    acc[cur.id] = cur.title;
                    return acc;
                }, {});
                if(localStorage.getItem("station")) {
                    let json  = JSON.parse(localStorage.getItem("station"));
                    this.center = json.id;
                    localStorage.removeItem("station");
                }
            }, () => {
                this.$notify({text: "Не удалось получить доступные станции", type: "error"});
            })
        },
        sendToBot() {
            if(this.city < 1) {
                this.$notify({text: "Выберите город", type: "error"});
                return;
            }
            if(this.place == 0 && this.center < 1) {
                this.$notify({text: "Выберите центр крови", type: "error"});
                return;
            }
            useWebApp().sendData(JSON.stringify(
                {
                    type: "plan_donation",
                    data: {
                        date: this.date.toISOString().slice(0, 10),
                        blood_type: this.blood_type,
                        type: this.type,
                        place: this.place,
                        city_id: this.city,
                        center_id: this.center
                    },
                    hash: this.$cookies.get("hash"),
                    id: this.$cookies.get("id")
                }));
        }
    },
    mounted() {
        RegionService.getCities((data) => {
            this.cities = data.results.reduce((acc, cur) => {
                acc[cur.id] = cur.title;
                return acc;
            }, {});
            if(localStorage.getItem("station")) {
                let json  = JSON.parse(localStorage.getItem("station"));
                this.city = json.city_id;
            }
            this.updateCenters();
        }, () => {
            this.$notify({text: "Не удалось получить доступные города", type: "error"});
        })
    }
}
</script>