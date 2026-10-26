package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 int16 }
type GenS2 struct { F0 int; F1 string; F2 int64 }
func genFn0() int8 {
    var v6 uint64 = 0
    v6 -= 18446744073709551615
    v7 := []int{32768}
    var v8 int8 = 100
    _ = v6
    _ = v7
    _ = v8
    return v8
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int = -9223372036854775808
    var v3 uint32 = 1
    v4 := []uint32{v3}
    v5 := make(map[string]string)
    if (9223372036854775807 <= 9223372036854775807) {
        v9 := GenS1{F0: -9223372036854775808, F1: 1, F2: 42}
        var v10 string = ("" + "foo bar")
        _ = v10
        _ = v9
    } else {
        delete(v1, 2147483648)
    }
    fmt.Println(v3, v2)
    for fi11 := 0; fi11 < 8; fi11++ {
        delete(v1, v3)
        switch ((("" + "\n\t") + ("foo bar" + "")) + ("hello" + (("foo bar" + "") + "hello"))) {
        case "":
            v3 += (^(v3))
        case "foo bar":
            v12 := GenS2{F0: -1, F1: (("hello" + "\n\t") + ("foo bar" + "")), F2: 9223372036854775807}
            _ = v12
        default:
            v3 -= v3
        }
        var v13 int16 = 32767
        _ = fi11
        _ = v13
    }
    switch v2 {
    case 4294967295:
        if false {
            for fi14 := 0; fi14 < 8; fi14++ {
                v2 = (v2 & (-(fi14)))
                for fi15 := 0; fi15 < 16; fi15++ {
                    if true {
                        var v16 int = fi14
                        _ = v16
                    }
                    var v17 uint8 = 127
                    _ = fi15
                    _ = v17
                }
                _ = fi14
            }
            v3 += 0
        }
    case -179:
        v18 := v1[(v3 >> ((2) & 63))]
        _ = v18
    }
    var v19 int32 = -1
    if true {
        fmt.Println(v19, v3, v2)
        var v20 int16 = -32768
        if (uint16(v3) >= 255) {
            var v21 float64 = (-2.25 * ((3.14159 - 1.5) + float64(float32(-2.25))))
            _ = v21
        }
        _ = v20
    }
    for fi22 := 0; fi22 < 15; fi22++ {
        var v23 float64 = (0.1 - 0.1)
        v3 = v3
        _ = fi22
        _ = v23
    }
    if len(v4) > 0 {
        v4[0] = (v3 / ((v3) | 1))
    }
    v2 = 1
    v19 += v19
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
