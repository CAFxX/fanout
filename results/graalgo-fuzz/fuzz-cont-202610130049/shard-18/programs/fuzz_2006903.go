package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 string; F1 uint32 }

func main() {
    var v1 uint = 355
    v2 := make(map[uint32]bool)
    v3 := make(map[int]int)
    var v4 int64 = 9223372036854775807
    var v5 uint64 = 148
    if false {
        if false {
            for fi6 := 0; fi6 < 16; fi6++ {
                v4 -= 42
                v1 -= 255
                _ = fi6
            }
            v5 = 18446744073709551615
        } else {
            if ("hello" == "") {
                for fi7 := 0; fi7 < 14; fi7++ {
                    v4 = -9223372036854775808
                    _ = fi7
                }
                var v8 int64 = v4
                switch 4294967295 {
                case -218:
                    var v9 uint8 = 177
                    _ = v9
                case 128:
                    v1 -= v1
                case 4294967295:
                    v4 += (v8 ^ 42)
                }
                _ = v8
            }
        }
        for fi10 := 0; fi10 < 7; fi10++ {
            fmt.Println(v1, v1)
            _ = fi10
        }
        for fi11 := 0; fi11 < 18; fi11++ {
            v5 -= v5
            switch ("foo bar" + ((("\n\t" + "\n\t") + ("" + "\n\t")) + "foo bar")) {
            case "\n\t":
                var v12 float64 = (0.0 / 1e10)
                _ = v12
            case "a":
                var v13 uint = v1
                _ = v13
            case "":
                var v14 uint16 = 32767
                _ = v14
            }
            _ = fi11
        }
    }
    v15 := []int{-1, 2147483647, 0, 2147483647}
    v5 -= v5
    var v16 int32 = -543
    var v17 int = -950
    for fi18 := 0; fi18 < 11; fi18++ {
        if false {
            v5 += (^(v5))
            switch "hello" {
            case "":
                var v19 string = ""
                _ = v19
            case "hello":
                v1 -= v1
            default:
                for fi20 := 0; fi20 < 20; fi20++ {
                    var v21 string = "foo bar"
                    _ = fi20
                    _ = v21
                }
            }
            if true {
                var v22 uint64 = 4294967295
                v23 := GenS1{F0: 100, F1: 1}
                _ = v22
                _ = v23
            } else {
                v24 := GenS1{F0: 42, F1: 127}
                if (!true) {
                    fmt.Println(v16, v16)
                } else {
                    _ = len(v3)
                    var v25 uint64 = v5
                    _ = v25
                }
                for fi26 := 0; fi26 < 16; fi26++ {
                    if true {
                        fmt.Println(fi18, v17)
                        v27 := GenS1{F0: 100, F1: 1}
                        v16 += v16
                        _ = v27
                    }
                    v1 = 134
                    _ = fi26
                }
                _ = v24
            }
        } else {
            for fi28 := 0; fi28 < 2; fi28++ {
                v5 += 726
                _ = fi28
            }
        }
        v17 -= (fi18 ^ fi18)
        _ = fi18
    }
    v17 -= v17
    v17 = v17
    var v29 uint16 = 644
    switch "\n\t" {
    case "a":
        v30 := v2[927]
        _ = v30
    default:
        v31 := GenS2{F0: "\n\t", F1: 2147483648}
        _ = v31
    }
    if len(v15) > 0 {
        v15[0] = int(v16)
    }
    v1 -= (v1 | (v1 + v1))
    v5 = (v5 << ((837) & 63))
    fmt.Println(v4, v29)
    if false {
        v16 -= (v16 & int32(v17))
    } else {
        if false {
            if ((0 < -907) && (122 != (^(v5)))) {
                if false {
                    v17 = (v17 & (v17 | v17))
                    var v32 uint8 = 128
                    _ = v32
                }
            }
            var v33 int16 = 85
            _ = v33
        } else {
            v1 = (v1 << (((v1 - v1)) & 63))
            var v34 int64 = (v4 / ((v4) | 1))
            var v35 int8 = 100
            _ = v34
            _ = v35
        }
        if (!false) {
            var v36 int32 = v16
            _ = v36
        } else {
            v1 = v1
            v1 -= 78
        }
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
