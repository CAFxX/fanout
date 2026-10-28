package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 uint }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 uint; F1 uint; F2 uint16 }
func genFn0(p0 uint64) uint16 {
    var v7 int = 2147483647
    var v8 uint = 2
    var v9 uint64 = 0
    fmt.Println(p0, v8)
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 0
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int = 127
    var v3 float64 = -2.25
    var v4 float64 = ((v3 / v3) - v3)
    var v5 string = "a"
    v6 := make(map[int]int)
    if (32768 != (v2 | v2)) {
        v2 = 42
        for fi10 := 0; fi10 < 18; fi10++ {
            v5 = ((((v5 + v5) + ("\n\t" + "a")) + (v5 + (v5 + "foo bar"))) + v5)
            v11 := GenS1{F0: -32768, F1: -621, F2: 255}
            _ = fi10
            _ = v11
        }
    }
    var v12 uint16 = 2
    for fi13 := 0; fi13 < 13; fi13++ {
        for fi14 := 0; fi14 < 19; fi14++ {
            v1[(fi14 / ((v2) | 1))] = 2147483647
            v2 = fi14
            _ = len(v6)
            var v15 int16 = 42
            _ = fi14
            _ = v15
        }
        v16 := GenS2{F0: 1000, F1: 1000}
        switch 32767 {
        case 2147483647:
            v17 := GenS2{F0: 32767, F1: -1}
            _ = v17
        case 426:
            fmt.Println(v12, v12)
        case 127:
            switch ((v5 + ("foo bar" + v5)) + ("a" + "foo bar")) {
            case "hello":
                v12 += v12
            case "\n\t":
                var v18 uint = 18446744073709551615
                _ = v18
            case "a":
                fmt.Println(v4)
            default:
                var v19 string = v5
                _ = v19
            }
        default:
            v20 := v1[(v2 * 9223372036854775807)]
            _ = v20
        }
        _ = fi13
        _ = v16
    }
    if ((false || false) && ((139 < 706) && (!(!false)))) {
        var v21 int32 = 739
        v2 = -2147483648
        _ = v21
    } else {
        var v22 int32 = -1
        _ = v22
    }
    v23 := GenS3{F0: 0, F1: 2, F2: 0}
    if (65535 > 2147483647) {
        var v24 int64 = 42
        for fi25 := 0; fi25 < 4; fi25++ {
            v2 = 32767
            v26 := GenS1{F0: -48, F1: v24, F2: 2}
            v4 = v3
            _ = fi25
            _ = v26
        }
        _ = v24
    }
    switch v5 {
    case "\n\t":
        v2 = (v2 - v2)
    }
    var v27 uint = 65535
    v28 := GenS2{F0: 1000, F1: 32767}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v2
    _ = v23
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
