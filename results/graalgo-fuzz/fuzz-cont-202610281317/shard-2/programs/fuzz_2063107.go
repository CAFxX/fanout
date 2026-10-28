package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 int8; F2 int16 }
type GenS3 struct { F0 int; F1 bool; F2 uint32 }

func main() {
    var v1 int64 = 9223372036854775807
    var v2 uint32 = 65535
    var v3 uint16 = 1
    var v4 int32 = -1
    v2 += (^(v2))
    v2 = 2147483648
    var v5 float64 = -2.25
    v6 := []int{int(v3), 127}
    if false {
        var v7 string = "a"
        if (true && false) {
            for fi8 := 0; fi8 < 1; fi8++ {
                if len(v6) > 0 {
                    v6[0] = fi8
                }
                _ = fi8
            }
            var v9 int8 = 42
            _ = v9
        } else {
            v4 += v4
            if false {
                _ = len(v6)
                fmt.Println(v3, v2, v4)
            } else {
                var v10 float64 = (float64((((0.1 / 1.5) - ((((-2.25 * 0.1) + 0.1) / 1.0) * float32(v5))) + 3.14159)) / 0.1)
                v11 := GenS1{F0: 503}
                _ = v10
                _ = v11
            }
            v12 := GenS1{F0: -383}
            _ = v12
        }
        v4 += (v4 * v4)
        _ = v7
    } else {
        if false {
            var v13 string = "foo bar"
            _ = v13
        }
        var v14 string = ""
        v3 = (v3 * v3)
        _ = v14
    }
    for fi15 := 0; fi15 < 12; fi15++ {
        if (1 < -1) {
            v3 -= 69
        }
        _ = len(v6)
        for fi16 := 0; fi16 < 15; fi16++ {
            fmt.Println(v1, v4, v4)
            _ = fi16
        }
        _ = fi15
    }
    v3 += (^(v3))
    for fi17 := 0; fi17 < 14; fi17++ {
        v18 := GenS2{F0: 278, F1: 81, F2: -34}
        _ = fi17
        _ = v18
    }
    if (1 > 127) {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    for fi19 := 0; fi19 < 4; fi19++ {
        var v20 float64 = (float64((0.0 * -2.25)) - 1e10)
        var v21 uint16 = 286
        fmt.Println(v4, v4, v4)
        _ = fi19
        _ = v20
        _ = v21
    }
    var v22 bool = true
    if false {
        v3 = v3
        if true {
            fmt.Println(v3)
        } else {
            v6 = append(v6, 32768)
            for fi23 := 0; fi23 < 9; fi23++ {
                if ("" != "foo bar") {
                    for fi24 := 0; fi24 < 10; fi24++ {
                        var v25 uint8 = 128
                        fmt.Println(fi24, v22, v5)
                        v2 = v2
                        _ = fi24
                        _ = v25
                    }
                }
                _ = fi23
            }
            var v26 uint16 = v3
            _ = v26
        }
        for fi27 := 0; fi27 < 4; fi27++ {
            var v28 int32 = v4
            _ = fi27
            _ = v28
        }
    } else {
        switch -282 {
        case 348:
            v29 := GenS3{F0: 128, F1: true, F2: v2}
            _ = v29
        default:
            _ = len(v6)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
