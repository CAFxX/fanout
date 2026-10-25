package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint; F1 int64; F2 int16 }
type GenS3 struct { F0 uint16; F1 uint16; F2 string }
type GenS4 struct { F0 int16; F1 uint32; F2 int16 }
type GenS5 struct { F0 uint8 }

func main() {
    var v1 int = -1
    var v2 uint32 = 42
    var v3 float64 = (0.0 * float64(1e10))
    var v4 uint32 = 581
    v5 := make(map[string]int)
    var v6 int = int(v2)
    v7 := []int{2147483647, v1, v1, 636}
    fmt.Println(v2, v4)
    v5[""] = -2147483648
    v3 -= v3
    var v8 uint64 = 42
    if false {
        v9 := GenS1{F0: 115, F1: 18446744073709551615}
        if (-2.25 >= 1.5) {
            fmt.Println(v4, v2)
            v5["foo bar"] = (v6 % (((^(v1))) | 1))
            switch "" {
            case "":
                var v10 uint32 = (v2 + 4294967295)
                _ = v10
            case "\n\t":
                v11 := GenS1{F0: -32768, F1: 2}
                _ = v11
            case "hello":
                var v12 int64 = 0
                _ = v12
            default:
                fmt.Println(v2, v1)
            }
        }
        v2 -= (v4 & 42)
        _ = v9
    } else {
        v1 = v1
        switch "hello" {
        case "":
            if (255 > 172) {
                if (!(0.0 != 0.0)) {
                    if (false && (124 >= 128)) {
                        for fi13 := 0; fi13 < 18; fi13++ {
                            var v14 uint16 = 42
                            v1 -= v1
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                            _ = fi13
                            _ = v14
                        }
                        var v15 uint = 255
                        v3 = 1e10
                        _ = v15
                    }
                    v1 += (v6 >> ((1) & 63))
                    for fi16 := 0; fi16 < 10; fi16++ {
                        v1 = fi16
                        _ = fi16
                    }
                }
                var v17 uint16 = 32767
                _ = v17
            } else {
                v18 := GenS2{F0: 281, F1: -9223372036854775808, F2: 1000}
                var v19 string = "\n\t"
                v20 := GenS3{F0: 151, F1: 0, F2: v19}
                _ = v18
                _ = v19
                _ = v20
            }
        }
    }
    v21 := GenS4{F0: 727, F1: v4, F2: 32767}
    v22 := GenS1{F0: -32768, F1: 18446744073709551615}
    fmt.Println(v8, v4)
    var v23 int32 = -793
    fmt.Println(v3)
    if (false || (false && false)) {
        v3 = (v3 * v3)
    } else {
        var v24 uint = 65535
        if false {
            var v25 uint64 = v8
            v26 := GenS3{F0: 32767, F1: 255, F2: ("foo bar" + "a")}
            v1 -= (-(v1))
            _ = v25
            _ = v26
        } else {
            v27 := v5["a"]
            var v28 uint64 = 42
            v23 = v23
            _ = v27
            _ = v28
        }
        for fi29 := 0; fi29 < 10; fi29++ {
            v30 := GenS5{F0: 233}
            _ = fi29
            _ = v30
        }
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
