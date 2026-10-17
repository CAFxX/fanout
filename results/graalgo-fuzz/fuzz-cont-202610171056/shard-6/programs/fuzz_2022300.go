package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 int8 }
func genFn0(p0 int8, p1 uint32, p2 int16) int16 {
    v6 := GenS1{F0: 42}
    var v7 uint64 = 18446744073709551615
    var v8 float32 = ((3.14159 * (float32(0.0) / 3.14159)) * -2.25)
    switch ("foo bar" + (("hello" + "hello") + "")) {
    case "a":
        v9 := GenS1{F0: 255}
        _ = v9
    case "\n\t":
        var v10 int64 = int64(p0)
        _ = v10
    case "hello":
        var v11 float64 = float64(0.0)
        _ = v11
    default:
        fmt.Println(v7)
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        var v13 int16 = p2
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return (p2 | 124)
}

func main() {
    var v1 int32 = -1
    var v2 int32 = (v1 - v1)
    v3 := []int{768, 426}
    v4 := make(map[uint32]string)
    var v5 int32 = (v1 | (v1 | 1))
    switch 32767 {
    case 128:
        fmt.Println(v5, v2)
    }
    v2 = v1
    var v14 uint8 = 1
    if len(v3) > 0 {
        v3[0] = 9223372036854775807
    }
    v1 -= v5
    fmt.Println(v2, v5)
    _ = len(v4)
    switch "\n\t" {
    case "foo bar":
        if false {
            var v15 int32 = 2147483647
            v3 = append(v3, 49)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v15
        }
    case "a":
        fmt.Println(v14)
    }
    for fi16 := 0; fi16 < 7; fi16++ {
        fmt.Println(fi16, v14)
        _ = fi16
    }
    v4[381] = "foo bar"
    if false {
        for fi17 := 0; fi17 < 15; fi17++ {
            _ = len(v4)
            _ = fi17
        }
        v18 := GenS2{F0: "", F1: 100}
        if (true && (!true)) {
            v2 = v1
        } else {
            var v19 uint32 = 511
            _ = v19
        }
        _ = v18
    }
    if false {
        switch 9223372036854775807 {
        case 4294967295:
            v2 += v2
        case 0:
            switch "" {
            case "foo bar":
                var v20 float32 = float32((1e10 + ((-2.25 / 3.14159) * float64(-2.25))))
                _ = v20
            default:
                v5 = (v1 << ((35) & 63))
            }
        }
        var v21 int = 32767
        _ = v21
    }
    v1 += v1
    var v22 bool = false
    fmt.Println(v5)
    v4[65535] = "foo bar"
    var v23 int32 = (v2 * (v5 << ((42) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
