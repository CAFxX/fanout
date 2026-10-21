package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 int16, p1 int32) int32 {
    for fi5 := 0; fi5 < 5; fi5++ {
        var v6 uint32 = 42
        fmt.Println(fi5, fi5, fi5)
        _ = fi5
        _ = v6
    }
    if (!true) {
        v7 := []int{32768, 128}
        var v8 int64 = 9223372036854775807
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    return (p1 << ((4294967295) & 63))
}

func main() {
    v1 := []int{32768, 752, 32768, 32768}
    var v2 float64 = (1e10 * -2.25)
    var v3 float32 = -2.25
    v4 := make(map[uint32]uint32)
    fmt.Println(v2, v3)
    v2 += 0.0
    v3 -= (1.5 - 1e10)
    if true {
        if false {
            _ = len(v1)
            v3 = v3
        }
        var v9 int64 = 1
        _ = v9
    } else {
        v2 -= (v2 + (v2 + (1e10 + (-2.25 + 0.1))))
    }
    if false {
        v2 = 1e10
    } else {
        switch "\n\t" {
        case "\n\t":
            v3 = v3
        case "a":
            v10 := GenS1{F0: 32}
            _ = v10
        case "":
            v11 := GenS1{F0: 100}
            _ = v11
        default:
            switch ("hello" + "foo bar") {
            case "a":
                delete(v4, 111)
            case "":
                if (true && (!(v3 <= (v3 + 1e10)))) {
                    v2 = v2
                    fmt.Println(v2, v3)
                    v3 = float32(-2.25)
                } else {
                    var v12 uint = 0
                    var v13 float64 = -2.25
                    for fi14 := 0; fi14 < 15; fi14++ {
                        fmt.Println(v2, v2)
                        if true {
                            fmt.Println(v3, v3, fi14)
                            var v15 int32 = 770
                            v16 := GenS1{F0: 50}
                            _ = v15
                            _ = v16
                        }
                        _ = fi14
                    }
                    _ = v12
                    _ = v13
                }
            case "foo bar":
                var v17 int64 = -9223372036854775808
                _ = v17
            default:
                v2 = 0.1
            }
        }
    }
    v3 = v3
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v18 bool = true
    v2 += v2
    var v19 uint32 = 2
    v1 = append(v1, -2147483648)
    fmt.Println(v3, v3)
    v3 += 0.0
    v19 -= v19
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
