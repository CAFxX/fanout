package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 int32; F1 int8; F2 uint8 }
type GenS4 struct { F0 int32; F1 string; F2 int16 }
func genFn0() uint8 {
    var v5 int32 = 464
    var v6 int32 = (v5 | (^(v5)))
    var v7 uint8 = 30
    _ = v5
    _ = v6
    _ = v7
    return (^(v7))
}
func genFn1() int16 {
    v8 := []int{9223372036854775807}
    var v9 float32 = 0.1
    _ = v8
    _ = v9
    return 32767
}

func main() {
    v1 := make(map[string]string)
    var v2 float64 = 0.1
    var v3 int = 127
    v4 := []uint32{65535, 65535, 292}
    v2 = 3.14159
    v10 := GenS1{F0: 116, F1: -463}
    v3 += (-(v3))
    if ((937 != 2147483648) || (((("a" + "foo bar") > ("hello" + "a")) || (255 > 1)) || false)) {
        _ = len(v1)
        var v11 string = (("a" + "") + "\n\t")
        var v12 uint64 = 2
        _ = v11
        _ = v12
    }
    v13 := GenS1{F0: 18446744073709551615, F1: 931}
    for fi14 := 0; fi14 < 17; fi14++ {
        v3 += (fi14 + (v3 | (v3 % (((-(v3))) | 1))))
        if true {
            for fi15 := 0; fi15 < 10; fi15++ {
                var v16 int = 4294967295
                v17 := GenS1{F0: 4294967295, F1: -179}
                if true {
                    v18 := GenS2{F0: 4294967295, F1: 42}
                    _ = v18
                }
                switch fi14 {
                case 32768:
                    if false {
                        fmt.Println(fi14)
                    } else {
                        v19 := v1["hello"]
                        _ = v19
                    }
                }
                var v20 int32 = 0
                _ = fi15
                _ = v16
                _ = v17
                _ = v20
            }
            v2 = float64(0.1)
            switch "hello" {
            case "":
                var v21 float32 = -2.25
                _ = v21
            case "\n\t":
                if (!((0 >= -2147483648) && (v3 <= v3))) {
                    var v22 uint = 1
                    v3 += (^(fi14))
                    fmt.Println(v22)
                    _ = v22
                }
            default:
                v3 = 42
            }
        }
        var v23 uint64 = 4294967295
        _ = fi14
        _ = v23
    }
    for fi24 := 0; fi24 < 14; fi24++ {
        if (!(86 == 16)) {
            var v25 uint16 = 65535
            if true {
                var v26 int = (-(fi24))
                v4 = append(v4, 631)
                _ = v26
            } else {
                fmt.Println(v25, fi24, fi24)
                if true {
                    var v27 int16 = -156
                    _ = v27
                }
                _ = len(v1)
            }
            _ = v25
        }
        var v28 uint8 = 160
        _ = fi24
        _ = v28
    }
    delete(v1, "hello")
    for fi29 := 0; fi29 < 10; fi29++ {
        v3 = 128
        _ = fi29
    }
    switch v3 {
    case -2147483648:
        v30 := GenS3{F0: -1, F1: 0, F2: 26}
        _ = v30
    case 32768:
        v3 += (^(v3))
    default:
        v31 := GenS4{F0: -2147483648, F1: "hello", F2: 42}
        _ = v31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
