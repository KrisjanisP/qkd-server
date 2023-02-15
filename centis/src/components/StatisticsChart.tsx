import Chart, {ChartItem} from 'chart.js/auto';
import {useContext, useEffect, useRef} from "react";
import {ConfigContext} from "../utils/config-context";
import {wsConnect, wsSendRequest} from "../utils/promise-ws";

export default function StatisticsChart() {
    const config = useContext(ConfigContext)

    let chartCanvas = useRef();

    useEffect(() => {
        let chart = new Chart(chartCanvas.current as ChartItem, {
            type: 'line',
            data: {
                labels: ['1', '2', '3'],
                datasets: [{
                    label: 'Aija reservable Keys',
                    data: [12, 19, 3],
                    tension: 0.3,
                },
                    {
                        label: 'Brencis reservable Keys',
                        data: [2, 9, 13],
                        tension: 0.3,

                    }
                ]
            }
        });
        let interval = setInterval(async () => {
            let aijaWs = await wsConnect(config.aijaEndpoint)
            let brencisWS = await wsConnect(config.brencisEndpoint);

            chart.data.labels.push((chart.data.labels.length + 1).toString())
            chart.data.datasets[0].data.push(Math.floor(Math.random() * 100))
            chart.update()
        }, 1000)
        return () => {
            clearInterval(interval)
            chart.destroy();
        }
    }, [])

    return (
        <div className="my-4">
            <canvas ref={chartCanvas}></canvas>
        </div>
    )
}