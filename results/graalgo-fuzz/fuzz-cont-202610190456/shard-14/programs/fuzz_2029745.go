package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 int8; F1 int8 }

func main() {
    var v1 uint64 = 1
    var v2 uint16 = 42
    v3 := make(map[string]string)
    var v4 int32 = -150
    var v5 uint8 = 2
    v6 := []string{}
    var v7 int64 = 0
    _ = len(v6)
    if false {
        if true {
            switch 73 {
            case 396:
                v7 -= (v7 * (v7 * -9223372036854775808))
            case -2147483648:
                v4 = -216
            default:
                var v8 int32 = (v4 ^ 42)
                _ = v8
            }
            fmt.Println(v1, v7, v5)
        }
        v6 = append(v6, "")
        for fi9 := 0; fi9 < 1; fi9++ {
            var v10 float32 = float32(-2.25)
            _ = fi9
            _ = v10
        }
    }
    v5 += (v5 << ((1) & 63))
    if (42 >= 42) {
        var v11 uint64 = (v1 - v1)
        v12 := GenS1{F0: (v5 | (v5 * v5)), F1: v11, F2: (v5 + v5)}
        v13 := GenS1{F0: (v5 | 2), F1: v1, F2: uint8(v1)}
        _ = v11
        _ = v12
        _ = v13
    } else {
        var v14 uint8 = 67
        _ = v14
    }
    var v15 uint8 = (v5 - 2)
    v16 := GenS2{F0: 255, F1: 1000}
    for fi17 := 0; fi17 < 19; fi17++ {
        var v18 int16 = 1000
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 20; fi19++ {
        if false {
            var v20 int8 = -66
            var v21 string = (("\n\t" + "") + ("a" + ("hello" + "foo bar")))
            if false {
                fmt.Println(v20)
                v2 -= (v2 << ((18446744073709551615) & 63))
                var v22 uint32 = 2147483648
                _ = v22
            }
            _ = v20
            _ = v21
        } else {
            switch ("hello" + "foo bar") {
            case "\n\t":
                fmt.Println(fi19, v1)
            case "foo bar":
                var v23 uint = 510
                _ = v23
            }
            var v24 uint = 868
            _ = v24
        }
        var v25 uint64 = 2
        v26 := GenS2{F0: 255, F1: -32768}
        _ = fi19
        _ = v25
        _ = v26
    }
    v1 -= (v1 * 18446744073709551615)
    v27 := GenS3{F0: -128, F1: 83}
    var v28 float32 = (1.5 - 1.5)
    var v29 uint8 = (v5 / (((^(v5))) | 1))
    v1 = (v1 << ((uint(v15)) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
