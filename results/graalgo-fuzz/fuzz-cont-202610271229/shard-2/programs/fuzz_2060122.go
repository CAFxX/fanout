package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 uint16; F1 int; F2 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 bool; F1 uint64; F2 string }

func main() {
    var v1 float64 = (0.1 / (0.1 + ((-2.25 / float64(3.14159)) + 0.1)))
    var v2 uint8 = 38
    var v3 int = int(v2)
    v4 := GenS1{F0: ("a" + "hello"), F1: 1}
    var v5 int = v3
    v6 := []int{32767, v5, 1, (v3 ^ (v5 - (v3 | (v3 | (v3 | v3)))))}
    if (!true) {
        var v7 uint16 = 912
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        v1 = float64(0.0)
        _ = v7
    }
    var v8 int32 = 0
    switch "a" {
    case "\n\t":
        v8 = (-(v8))
    case "foo bar":
        v5 = (v5 - (v3 + 9223372036854775807))
    default:
        fmt.Println(v3, v5)
    }
    var v9 uint = 4294967295
    if (v8 == (v8 ^ v8)) {
        v10 := GenS2{F0: 1, F1: int(v2), F2: 1}
        v11 := GenS2{F0: 250, F1: v3, F2: uint16(v2)}
        _ = v10
        _ = v11
    } else {
        switch (v5 + v3) {
        case 708:
            switch v3 {
            case -1:
                var v12 int8 = -1
                _ = v12
            }
        }
        var v13 bool = false
        for fi14 := 0; fi14 < 3; fi14++ {
            var v15 float32 = (0.1 + (1e10 / 1e10))
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = fi14
            _ = v15
        }
        _ = v13
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v16 int64 = 800
    v2 -= (v2 & (v2 & 235))
    for fi17 := 0; fi17 < 8; fi17++ {
        if false {
            var v18 int = v5
            for fi19 := 0; fi19 < 16; fi19++ {
                v20 := GenS3{F0: 42}
                var v21 int64 = (v16 >> ((v9) & 63))
                _ = fi19
                _ = v20
                _ = v21
            }
            var v22 int8 = -124
            _ = v18
            _ = v22
        } else {
            var v23 int64 = (v16 << ((65535) & 63))
            switch "a" {
            case "":
                switch "" {
                case "":
                    v16 = (v23 << (((v9 & (v9 / ((v9) | 1)))) & 63))
                case "\n\t":
                    v16 = 661
                case "a":
                    v9 = (v9 - v9)
                default:
                    if (true || ((!true) || false)) {
                        v6 = append(v6, 0)
                        if false {
                            var v24 uint16 = 32767
                            if (v8 < v8) {
                                var v25 uint8 = v2
                                _ = v25
                            }
                            v26 := GenS2{F0: (v24 | v24), F1: (v5 << ((4294967295) & 63)), F2: 42}
                            _ = v24
                            _ = v26
                        }
                        for fi27 := 0; fi27 < 9; fi27++ {
                            var v28 int8 = 1
                            _ = len(v6)
                            var v29 float32 = (float32(float64(3.14159)) * ((0.0 + -2.25) / float32(1.5)))
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                    }
                }
            case "hello":
                if true {
                    switch (v5 >> ((v9) & 63)) {
                    case -2147483648:
                        fmt.Println(v16, v2)
                    case 4294967295:
                        for fi30 := 0; fi30 < 1; fi30++ {
                            var v31 uint64 = 2
                            _ = fi30
                            _ = v31
                        }
                    case -17:
                        v8 += (v8 / ((608) | 1))
                    }
                    fmt.Println(v8)
                } else {
                    var v32 uint8 = v2
                    if (true || (42 == uint16(v2))) {
                        var v33 int = 243
                        _ = v33
                    } else {
                        v23 -= -438
                        v9 -= v9
                        var v34 int = v5
                        _ = v34
                    }
                    _ = v32
                }
            case "a":
                v2 -= v2
            }
            v8 = int32(fi17)
            _ = v23
        }
        _ = fi17
    }
    for fi35 := 0; fi35 < 8; fi35++ {
        v36 := GenS2{F0: 255, F1: (-(v3)), F2: 0}
        var v37 uint32 = 2147483648
        _ = fi35
        _ = v36
        _ = v37
    }
    v38 := GenS4{F0: false, F1: 0, F2: "\n\t"}
    v39 := GenS2{F0: 65535, F1: v5, F2: 255}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
