package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 string }

func main() {
    v1 := []uint32{}
    var v2 uint16 = 32767
    v3 := []uint32{186, 2147483647}
    v4 := []uint32{2147483647}
    var v5 string = "\n\t"
    if false {
        _ = len(v3)
        var v6 int16 = 0
        _ = v6
    } else {
        var v7 int32 = 1
        _ = v7
    }
    if ((v2 >= v2) && false) {
        var v8 uint64 = 4294967295
        var v9 uint8 = 209
        if (false && (!false)) {
            var v10 uint16 = 80
            _ = v10
        }
        _ = v8
        _ = v9
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v2 = 2
    v2 = (v2 / (((v2 ^ (^(v2)))) | 1))
    var v11 float64 = 1.5
    for fi12 := 0; fi12 < 3; fi12++ {
        if (!((32767 == 9223372036854775807) && (false || false))) {
            v11 += ((v11 / 1.0) * v11)
            for fi13 := 0; fi13 < 12; fi13++ {
                v11 = (0.1 * v11)
                fmt.Println(v5)
                var v14 int64 = -536
                _ = fi13
                _ = v14
            }
        } else {
            fmt.Println(v5)
            for fi15 := 0; fi15 < 10; fi15++ {
                var v16 int32 = 42
                v1 = append(v1, uint32(v2))
                _ = fi15
                _ = v16
            }
        }
        switch (-(fi12)) {
        case 0:
            v2 -= (^(v2))
        case -642:
            switch (("\n\t" + ("" + "hello")) + "hello") {
            case "foo bar":
                if true {
                    var v17 int64 = 9223372036854775807
                    _ = v17
                } else {
                    var v18 uint = 42
                    var v19 string = v5
                    for fi20 := 0; fi20 < 17; fi20++ {
                        v21 := GenS1{F0: 42, F1: 353, F2: v19}
                        v22 := GenS1{F0: 42, F1: 1, F2: "foo bar"}
                        v19 = v19
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                    _ = v18
                    _ = v19
                }
            }
        case 128:
            var v23 int8 = 100
            _ = v23
        }
        _ = fi12
    }
    var v24 float32 = 1.5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
