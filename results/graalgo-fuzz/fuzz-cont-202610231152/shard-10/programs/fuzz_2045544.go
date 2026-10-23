package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string; F1 int8 }

func main() {
    var v1 string = ("\n\t" + "a")
    var v2 uint8 = 1
    var v3 uint16 = 0
    for fi4 := 0; fi4 < 8; fi4++ {
        for fi5 := 0; fi5 < 19; fi5++ {
            if (3.14159 < float64(((float32(0.0) * ((-2.25 * float32(3.14159)) + float32((1.5 / ((3.14159 / (0.0 - 1.5)) + (1.5 + 0.0)))))) / 0.1))) {
                v6 := []int{}
                _ = v6
            } else {
                v7 := []int{(fi5 >> ((359) & 63)), 127}
                var v8 bool = (1.5 < 0.0)
                v9 := GenS1{F0: 744}
                if ((262 >= 147) || v8) {
                    v10 := GenS1{F0: 0}
                    _ = v10
                }
                _ = v7
                _ = v8
                _ = v9
            }
            v3 = (v3 * (v3 - 65535))
            _ = fi5
        }
        v3 = v3
        var v11 int64 = 0
        _ = fi4
        _ = v11
    }
    v12 := GenS2{F0: v2}
    fmt.Println(v1, v2)
    v13 := GenS2{F0: 61}
    for fi14 := 0; fi14 < 19; fi14++ {
        if true {
            v15 := []int{(fi14 / ((fi14) | 1)), fi14, fi14, (fi14 / ((fi14) | 1))}
            v1 = v1
            _ = v15
        }
        if false {
            for fi16 := 0; fi16 < 14; fi16++ {
                v17 := GenS1{F0: 9223372036854775807}
                var v18 uint16 = v3
                _ = fi16
                _ = v17
                _ = v18
            }
            var v19 uint8 = v2
            var v20 uint32 = 4294967295
            _ = v19
            _ = v20
        }
        _ = fi14
    }
    var v21 uint64 = 836
    var v22 float64 = float64(1e10)
    switch v1 {
    case "foo bar":
        if false {
            v3 = v3
            v3 += (v3 % ((v3) | 1))
            var v23 string = (v1 + v1)
            _ = v23
        }
    case "\n\t":
        for fi24 := 0; fi24 < 13; fi24++ {
            v21 += (v21 / ((v21) | 1))
            _ = fi24
        }
    case "":
        v22 += ((((((1.5 + ((float64(1e10) * 0.1) - (1e10 * float64(1e10)))) + -2.25) + v22) + v22) + float64(-2.25)) / (-2.25 / v22))
    default:
        switch int(v2) {
        case -2147483648:
            v1 = v1
        case 127:
            v25 := GenS1{F0: -197}
            _ = v25
        default:
            var v26 int64 = -1
            _ = v26
        }
    }
    v27 := []int{-711, -1}
    for fi28 := 0; fi28 < 15; fi28++ {
        if len(v27) > 0 {
            v27[0] = (^(fi28))
        }
        v2 += v2
        var v29 int32 = 0
        _ = fi28
        _ = v29
    }
    for fi30 := 0; fi30 < 12; fi30++ {
        var v31 float32 = 0.0
        for fi32 := 0; fi32 < 19; fi32++ {
            v33 := GenS2{F0: 127}
            v34 := GenS2{F0: v2}
            _ = fi32
            _ = v33
            _ = v34
        }
        _ = fi30
        _ = v31
    }
    switch 150 {
    case 1:
        for fi35 := 0; fi35 < 20; fi35++ {
            for fi36 := 0; fi36 < 8; fi36++ {
                v37 := GenS3{F0: v1, F1: -1}
                v38 := GenS3{F0: v1, F1: 42}
                _ = fi36
                _ = v37
                _ = v38
            }
            var v39 uint = 2
            var v40 int64 = 9223372036854775807
            _ = fi35
            _ = v39
            _ = v40
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v27
    _ = v3
}
