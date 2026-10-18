package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 int8 }
type GenS2 struct { F0 int32; F1 uint64; F2 string }
type GenS3 struct { F0 string; F1 uint32; F2 int32 }

func main() {
    v1 := []uint32{1, 183, 1, 4294967295}
    var v2 uint8 = 255
    var v3 float32 = (0.0 / (float32(float64(1.5)) - 0.1))
    v4 := make(map[uint32]string)
    v5 := []int{32767, -9223372036854775808}
    var v6 float64 = 1e10
    v6 = v6
    if false {
        switch "" {
        case "hello":
            switch ("a" + "") {
            case "hello":
                if false {
                    for fi7 := 0; fi7 < 15; fi7++ {
                        v2 = (v2 >> ((255) & 63))
                        _ = fi7
                    }
                }
            case "foo bar":
                for fi8 := 0; fi8 < 8; fi8++ {
                    if false {
                        if ((v2 ^ (v2 >> ((344) & 63))) <= v2) {
                            v9 := GenS1{F0: 65535, F1: -1, F2: 94}
                            if len(v1) > 0 {
                                v1[0] = 42
                            }
                            _ = v9
                        }
                        v2 += 2
                    }
                    var v10 float32 = 0.0
                    _ = fi8
                    _ = v10
                }
            case "":
                fmt.Println(v6, v3)
            default:
                var v11 uint8 = v2
                _ = v11
            }
        case "foo bar":
            switch 32768 {
            case 127:
                v6 = ((v6 - ((v6 / ((float64(-2.25) + v6) / (v6 / 1.0))) / (0.1 + (v6 - (v6 / ((v6 * (float64((v3 / -2.25)) - float64(v3))) * v6)))))) / (v6 * 1e10))
            }
        case "":
            fmt.Println(v2, v6, v3)
        }
        v12 := GenS1{F0: 255, F1: 2147483647, F2: -128}
        if ((v2 + (v2 / (((v2 / ((v2) | 1))) | 1))) >= 42) {
            v2 += (v2 & (v2 & 1))
            var v13 uint16 = 1
            var v14 int = -22
            _ = v13
            _ = v14
        } else {
            if (!false) {
                for fi15 := 0; fi15 < 18; fi15++ {
                    for fi16 := 0; fi16 < 11; fi16++ {
                        var v17 uint = 473
                        v18 := v4[833]
                        var v19 uint64 = 2
                        _ = fi16
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                    _ = fi15
                }
            }
        }
        _ = v12
    }
    for fi20 := 0; fi20 < 20; fi20++ {
        if true {
            switch (fi20 ^ (fi20 - (fi20 & (fi20 / (((fi20 >> ((1) & 63))) | 1))))) {
            case -2147483648:
                var v21 bool = true
                _ = v21
            case 4294967295:
                _ = len(v4)
            }
        }
        v3 += v3
        _ = fi20
    }
    v22 := GenS1{F0: 631, F1: 2147483647, F2: 15}
    switch 53 {
    case 128:
        v3 = (float32(-2.25) - -2.25)
    default:
        var v23 bool = (("hello" >= "a") || (!((42 >= 255) || (1e10 <= 1.5))))
        _ = v23
    }
    var v24 uint = 592
    for fi25 := 0; fi25 < 7; fi25++ {
        v6 = v6
        v2 += 255
        v2 += (v2 + 1)
        _ = fi25
    }
    v26 := GenS2{F0: 1, F1: 42, F2: ("a" + (("\n\t" + "") + ("a" + (("hello" + "") + "hello"))))}
    if ((818 <= 886) && false) {
        switch -9223372036854775808 {
        case 42:
            fmt.Println(v2, v3, v2)
        }
    } else {
        switch -268 {
        case 127:
            for fi27 := 0; fi27 < 2; fi27++ {
                fmt.Println(v6)
                _ = fi27
            }
        }
    }
    v28 := GenS2{F0: -2147483648, F1: 4294967295, F2: "foo bar"}
    if (((v24 / ((v24) | 1)) > v24) || false) {
        v29 := GenS3{F0: "foo bar", F1: 0, F2: 0}
        if len(v5) > 0 {
            v5[0] = -174
        }
        _ = v29
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v24
    _ = v26
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
