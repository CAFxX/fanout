package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 bool }
type GenS2 struct { F0 int8; F1 int64; F2 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 bool; F1 int16; F2 uint64 }
func genFn0() uint8 {
    if true {
        v4 := []int{}
        if (!true) {
            var v5 uint = 65535
            v6 := GenS1{F0: -74, F1: 32768, F2: (!true)}
            for fi7 := 0; fi7 < 18; fi7++ {
                v8 := GenS1{F0: -16, F1: (fi7 ^ fi7), F2: true}
                _ = fi7
                _ = v8
            }
            _ = v5
            _ = v6
        } else {
            var v9 float64 = ((0.1 - float64(1e10)) * 1.5)
            for fi10 := 0; fi10 < 6; fi10++ {
                if (127 < -128) {
                    for fi11 := 0; fi11 < 6; fi11++ {
                        var v12 int64 = 0
                        var v13 float64 = 3.14159
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                }
                v9 -= (1.5 + v9)
                _ = fi10
            }
            _ = v9
        }
        for fi14 := 0; fi14 < 18; fi14++ {
            var v15 int64 = 81
            _ = fi14
            _ = v15
        }
        _ = v4
    } else {
        switch "hello" {
        case "hello":
            var v16 int64 = -853
            _ = v16
        }
        if true {
            var v17 uint8 = 172
            _ = v17
        } else {
            var v18 float32 = 1e10
            v19 := GenS2{F0: 127, F1: 1, F2: 1}
            _ = v18
            _ = v19
        }
    }
    var v20 float64 = 1e10
    var v21 int32 = 2147483647
    v21 = 42
    _ = v20
    _ = v21
    return 163
}

func main() {
    var v1 float32 = 3.14159
    var v2 bool = (32767 >= 42)
    var v3 int8 = -69
    v2 = ("hello" != "")
    v22 := GenS1{F0: (v3 / ((v3) | 1)), F1: 609, F2: (true || false)}
    var v23 uint = 124
    v24 := GenS3{F0: 42}
    if (-1 == 1) {
        for fi25 := 0; fi25 < 20; fi25++ {
            v2 = (!((v3 + -41) == (v3 << (((v23 & (^(v23)))) & 63))))
            _ = fi25
        }
        fmt.Println(v2, v2)
    } else {
        v26 := GenS3{F0: 0}
        if true {
            switch -403 {
            case -1:
                var v27 int64 = 215
                _ = v27
            case -44:
                switch ("hello" + "hello") {
                case "foo bar":
                    v3 = (v3 >> ((v23) & 63))
                case "hello":
                    v3 += v3
                case "a":
                    v23 += (^(v23))
                default:
                    v28 := GenS4{F0: 76}
                    _ = v28
                }
            default:
                if (118 != genFn0()) {
                    v2 = ((!((v1 + v1) >= v1)) || true)
                    var v29 uint8 = 0
                    v30 := []int{504, 9223372036854775807}
                    _ = v29
                    _ = v30
                }
            }
            v31 := GenS2{F0: (v3 | v3), F1: 0, F2: 790}
            var v32 string = ""
            _ = v31
            _ = v32
        } else {
            v23 = 383
            v1 -= (float32((0.1 * -2.25)) * v1)
        }
        _ = v26
    }
    v33 := []int{127, 40}
    if false {
        v34 := GenS5{F0: v2, F1: 1, F2: 18446744073709551615}
        v23 -= (v23 & v23)
        _ = v34
    }
    var v35 float32 = 0.1
    _ = len(v33)
    fmt.Println(v1, v2, v23)
    var v36 bool = true
    for fi37 := 0; fi37 < 11; fi37++ {
        switch (("" + "a") + (("\n\t" + "\n\t") + "foo bar")) {
        case "a":
            var v38 uint = (v23 >> (((v23 * (v23 << ((v23) & 63)))) & 63))
            _ = v38
        case "hello":
            if v36 {
                if len(v33) >= 2 {
                    v33 = v33[1:len(v33)]
                }
                v36 = v36
                v2 = (v36 && v36)
            }
        default:
            v35 += v1
        }
        v39 := GenS4{F0: 153}
        _ = fi37
        _ = v39
    }
    var v40 int32 = 0
    v41 := GenS2{F0: -70, F1: 0, F2: 65535}
    v42 := GenS1{F0: (v3 ^ v3), F1: -1, F2: v2}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v33
    _ = v35
    _ = v36
    _ = v40
    _ = v41
    _ = v42
}
