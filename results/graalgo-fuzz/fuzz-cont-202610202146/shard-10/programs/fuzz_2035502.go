package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 int64 }
type GenS2 struct { F0 int64; F1 int64 }
func genFn0() int16 {
    v5 := []int{32768, 529}
    v6 := GenS1{F0: 42, F1: 65535, F2: -668}
    switch "" {
    case "hello":
        var v7 int16 = 1
        _ = v7
    default:
        v8 := GenS1{F0: -914, F1: 247, F2: -1}
        _ = v8
    }
    if false {
        switch 1 {
        case 1:
            var v9 uint64 = 42
            _ = v9
        case 128:
            _ = len(v5)
        case -2147483648:
            var v10 float32 = (float32(1.5) - 0.1)
            _ = v10
        default:
            var v11 float64 = 1.5
            _ = v11
        }
        var v12 string = "hello"
        if len(v5) > 0 {
            v5[0] = 296
        }
        _ = v12
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    _ = v5
    _ = v6
    return 32767
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint32 = 4294967295
    var v3 uint16 = 32767
    var v4 int32 = 432
    if true {
        var v13 uint32 = (^(v2))
        var v14 int16 = 42
        v4 -= v4
        _ = v13
        _ = v14
    }
    if (!true) {
        if ((0.1 > (0.0 * 1.5)) && true) {
            var v15 uint64 = 4294967295
            _ = v15
        } else {
            for fi16 := 0; fi16 < 4; fi16++ {
                v17 := GenS2{F0: 1, F1: -9223372036854775808}
                fmt.Println(v4, v2, v2)
                var v18 float32 = ((((1.5 - (0.0 * ((0.1 + float32(3.14159)) + -2.25))) * 0.1) * (1e10 + (((0.0 * (((1.5 / 1.5) / 1e10) + 0.0)) * 0.1) * 0.1))) + (1e10 + 3.14159))
                _ = fi16
                _ = v17
                _ = v18
            }
            for fi19 := 0; fi19 < 7; fi19++ {
                switch ("a" + "") {
                case "foo bar":
                    v20 := v1[508]
                    _ = v20
                }
                _ = fi19
            }
        }
        v21 := v1[0]
        v3 = v3
        _ = v21
    }
    v2 -= v2
    v2 += v2
    for fi22 := 0; fi22 < 15; fi22++ {
        for fi23 := 0; fi23 < 18; fi23++ {
            v3 = v3
            _ = fi23
        }
        switch "\n\t" {
        case "a":
            switch fi22 {
            case 128:
                var v24 uint32 = v2
                _ = v24
            }
        default:
            var v25 float64 = (1e10 - 0.1)
            _ = v25
        }
        v4 -= v4
        _ = fi22
    }
    if ((0.1 <= -2.25) || ((v3 << ((2) & 63)) < (v3 >> ((255) & 63)))) {
        if (!false) {
            v4 = -1
            var v26 uint16 = 714
            _ = v26
        }
        for fi27 := 0; fi27 < 13; fi27++ {
            var v28 float64 = (0.1 - -2.25)
            for fi29 := 0; fi29 < 12; fi29++ {
                v28 = v28
                _ = fi29
            }
            v3 = (v3 ^ v3)
            _ = fi27
            _ = v28
        }
        v4 = v4
    } else {
        if ((false && true) && true) {
            v30 := GenS1{F0: v4, F1: 234, F2: -9223372036854775808}
            v4 += v4
            if false {
                var v31 uint = 0
                if ((v4 < int32(v2)) && (false && (false && (!true)))) {
                    switch -914 {
                    case 127:
                        fmt.Println(v4, v4, v4)
                    }
                }
                if true {
                    fmt.Println(v4)
                    var v32 uint8 = 228
                    var v33 int64 = 586
                    _ = v32
                    _ = v33
                }
                _ = v31
            }
            _ = v30
        } else {
            var v34 uint = 1
            v2 -= 2
            _ = v34
        }
    }
    var v35 int = 127
    var v36 uint32 = (v2 % (((v2 * 4294967295)) | 1))
    v36 = 31
    v3 -= (v3 + v3)
    for fi37 := 0; fi37 < 2; fi37++ {
        v36 += v2
        switch fi37 {
        case 0:
            delete(v1, (v36 * (^(v36))))
        case 32767:
            var v38 int32 = (v4 / ((v4) | 1))
            _ = v38
        }
        _ = fi37
    }
    fmt.Println(v36)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
}
