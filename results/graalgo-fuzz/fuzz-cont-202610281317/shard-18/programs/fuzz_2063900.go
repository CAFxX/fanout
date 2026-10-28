package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int }
type GenS2 struct { F0 uint32 }
func genFn0() uint32 {
    if (1 != -9223372036854775808) {
        var v7 int16 = 32767
        _ = v7
    } else {
        var v8 string = ""
        _ = v8
    }
    switch 1 {
    case 32768:
        var v9 uint8 = 127
        _ = v9
    case 2147483647:
        var v10 bool = (-66 > 1)
        _ = v10
    }
    return 65535
}

func main() {
    var v1 uint64 = 1
    var v2 float64 = ((float64(-2.25) / 0.1) * (3.14159 * -2.25))
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]int)
    var v5 float64 = v2
    v6 := make(map[int]string)
    v2 = v2
    v11 := v4[893]
    v12 := v4[genFn0()]
    v13 := GenS1{F0: 255, F1: 32, F2: v12}
    v1 -= (v1 | v1)
    if false {
        var v14 int8 = 96
        v1 -= v1
        v11 = v11
        _ = v14
    } else {
        for fi15 := 0; fi15 < 6; fi15++ {
            switch (v11 & (v11 - (v11 + (fi15 | v12)))) {
            case 32767:
                var v16 bool = false
                _ = v16
            case -2147483648:
                v17 := GenS1{F0: 1, F1: 18446744073709551615, F2: fi15}
                _ = v17
            }
            fmt.Println(v2, v2)
            switch 2147483647 {
            case 128:
                v18 := []int{597, (^(v12)), v11}
                _ = v18
            case -619:
                fmt.Println(v1, v12)
            }
            _ = fi15
        }
        v19 := v6[v11]
        v1 -= 0
        _ = v19
    }
    if (18 < (-(v11))) {
        var v20 int16 = 197
        _ = v20
    } else {
        v21 := GenS2{F0: 351}
        switch 4294967295 {
        case 42:
            if true {
                v12 -= -958
            }
        case -785:
            if (0 >= 1) {
                switch (v12 - -2147483648) {
                case -736:
                    var v22 uint8 = 172
                    _ = v22
                case 315:
                    if (!(true && true)) {
                        v23 := []int{(v12 << ((18446744073709551615) & 63)), v12}
                        _ = v23
                    } else {
                        v24 := []int{v11, 128, -1, (-(v12))}
                        _ = v24
                    }
                default:
                    var v25 uint = 1
                    _ = v25
                }
            }
        default:
            v5 = v5
        }
        v26 := GenS2{F0: 1}
        _ = v21
        _ = v26
    }
    var v27 int = 819
    for fi28 := 0; fi28 < 8; fi28++ {
        v29 := []int{}
        _ = fi28
        _ = v29
    }
    switch ("hello" + "a") {
    case "\n\t":
        for fi30 := 0; fi30 < 2; fi30++ {
            _ = len(v4)
            _ = fi30
        }
    default:
        delete(v6, v11)
    }
    fmt.Println(v11)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
