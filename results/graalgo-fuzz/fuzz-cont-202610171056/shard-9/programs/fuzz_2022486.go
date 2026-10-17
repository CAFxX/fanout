package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 bool; F1 int32; F2 uint32 }

func main() {
    var v1 int = -2147483648
    var v2 int16 = 32767
    var v3 int = (v1 / (((v1 << ((65535) & 63))) | 1))
    v4 := make(map[string]int)
    v5 := GenS1{F0: 4294967295, F1: false}
    var v6 int64 = -554
    switch (v3 % ((127) | 1)) {
    case 1:
        if true {
            var v7 float32 = 1e10
            v8 := GenS1{F0: 806, F1: (!(((!true) || (18446744073709551615 < 18446744073709551615)) && true))}
            for fi9 := 0; fi9 < 2; fi9++ {
                var v10 int = 32768
                if (0.1 > (0.1 * ((-2.25 / 0.1) / float32((3.14159 * 3.14159))))) {
                    var v11 int = fi9
                    v12 := GenS2{F0: true, F1: 602, F2: 2147483648}
                    v3 += (v3 & v1)
                    _ = v11
                    _ = v12
                } else {
                    fmt.Println(v1)
                    var v13 float64 = 1e10
                    _ = v13
                }
                v14 := GenS1{F0: 65535, F1: (647 == 1)}
                _ = fi9
                _ = v10
                _ = v14
            }
            _ = v7
            _ = v8
        } else {
            v1 -= v3
            v2 -= v2
            for fi15 := 0; fi15 < 19; fi15++ {
                var v16 int8 = 27
                v6 = v6
                v2 -= v2
                _ = fi15
                _ = v16
            }
        }
    default:
        v2 = v2
    }
    if true {
        for fi17 := 0; fi17 < 11; fi17++ {
            var v18 int8 = 50
            var v19 int16 = v2
            _ = fi17
            _ = v18
            _ = v19
        }
        var v20 int16 = -32768
        _ = v20
    } else {
        var v21 uint = 4294967295
        v6 = (v6 / (((v6 * (v6 * v6))) | 1))
        v22 := []int{int(v2), (-(v1)), (-(v1))}
        _ = v21
        _ = v22
    }
    v2 = v2
    fmt.Println(v1, v6, v3)
    for fi23 := 0; fi23 < 6; fi23++ {
        v24 := GenS2{F0: false, F1: 326, F2: 927}
        fmt.Println(fi23, v3, fi23)
        _ = fi23
        _ = v24
    }
    var v25 float64 = ((float64(0.1) - (3.14159 / (3.14159 * (float64(0.1) / ((1e10 - (float64(1.5) * (((3.14159 + 0.0) + 3.14159) / 1.5))) / 3.14159))))) / 0.1)
    v26 := GenS2{F0: (!true), F1: 0, F2: 1}
    v27 := []int{(^(v1)), v3, v3}
    for fi28 := 0; fi28 < 7; fi28++ {
        if ((((v25 + (3.14159 + v25)) + 1.5) + 3.14159) <= (v25 / v25)) {
            var v29 string = ""
            if (!false) {
                fmt.Println(v25, fi28, v6)
                for fi30 := 0; fi30 < 15; fi30++ {
                    v31 := GenS2{F0: ((-(v6)) == v6), F1: 845, F2: 2147483648}
                    v6 += v6
                    _ = fi30
                    _ = v31
                }
                v1 = (fi28 >> ((1) & 63))
            }
            _ = v29
        } else {
            var v32 int = (v1 / ((4294967295) | 1))
            _ = v32
        }
        v3 = (v3 << ((59) & 63))
        _ = fi28
    }
    switch (v3 | 32767) {
    case -290:
        v6 += (v6 + v6)
    case 4294967295:
        var v33 int64 = -1
        _ = v33
    case 32768:
        v6 = -922
    default:
        var v34 int64 = 622
        _ = v34
    }
    _ = len(v4)
    v1 -= v3
    var v35 uint64 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
