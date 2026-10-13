package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 uint32 }
func genFn0(p0 int8, p1 uint32, p2 uint32) int64 {
    v7 := GenS1{F0: 127}
    fmt.Println(p0, p2, p1)
    v8 := []int{}
    var v9 bool = false
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 int8 = -88
    var v2 uint32 = 65535
    v3 := make(map[string]string)
    v4 := []int{128, 42}
    v5 := make(map[string]string)
    var v6 int32 = 42
    switch -9223372036854775808 {
    case -9223372036854775808:
        var v10 uint = 1
        _ = v10
    case 4294967295:
        var v11 string = (("foo bar" + (("" + "a") + "a")) + "hello")
        _ = v11
    default:
        v1 += (v1 >> ((1) & 63))
    }
    var v12 int = 32768
    v13 := GenS2{F0: false, F1: (^(v2))}
    var v14 int32 = (v6 ^ (v6 % ((-906) | 1)))
    v2 = (v2 - v2)
    if (!((((1.5 - ((1e10 * 1e10) - -2.25)) + (0.0 + (0.1 * ((0.1 / 1.0) - (-2.25 / float64(1e10)))))) * 0.1) >= 0.1)) {
        for fi15 := 0; fi15 < 19; fi15++ {
            var v16 int64 = 1
            _ = fi15
            _ = v16
        }
        var v17 int16 = -32768
        var v18 int16 = (v17 | 0)
        _ = v17
        _ = v18
    } else {
        v12 = v12
        v19 := GenS1{F0: 127}
        _ = v19
    }
    v20 := GenS1{F0: 42}
    switch v12 {
    case 769:
        var v21 int16 = 1000
        _ = v21
    case -2147483648:
        v1 += v1
    }
    v22 := GenS1{F0: 128}
    delete(v3, ("foo bar" + (((("hello" + "foo bar") + "hello") + ("foo bar" + "a")) + "")))
    var v23 int32 = int32(v12)
    switch 9223372036854775807 {
    case -219:
        for fi24 := 0; fi24 < 14; fi24++ {
            v14 += -2147483648
            var v25 bool = true
            _ = fi24
            _ = v25
        }
    case 42:
        v23 = v6
    case 128:
        v26 := GenS2{F0: false, F1: v2}
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
