package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32 }
func genFn0() uint8 {
    v5 := GenS1{F0: 127}
    v6 := GenS1{F0: -1}
    if (("" + (("\n\t" + "foo bar") + "")) >= "") {
        v7 := []int{0}
        _ = v7
    }
    _ = v5
    _ = v6
    return 161
}
func genFn1(p0 uint64) uint16 {
    for fi8 := 0; fi8 < 1; fi8++ {
        var v9 int8 = 29
        _ = fi8
        _ = v9
    }
    if (65535 != 32767) {
        if (!((false || false) || ((3.14159 / 1.5) == 0.0))) {
            var v10 int16 = 1000
            if true {
                p0 -= (p0 << ((494) & 63))
            }
            v10 -= -1
            _ = v10
        } else {
            for fi11 := 0; fi11 < 19; fi11++ {
                if false {
                    var v12 uint = 255
                    _ = v12
                } else {
                    v13 := []int{}
                    _ = len(v13)
                    _ = v13
                }
                p0 += 42
                _ = fi11
            }
            var v14 uint16 = 285
            _ = v14
        }
        fmt.Println(p0)
        var v15 int = 9223372036854775807
        _ = v15
    } else {
        fmt.Println(p0)
    }
    var v16 int64 = 0
    _ = p0
    _ = v16
    return 2
}

func main() {
    v1 := make(map[string]int)
    var v2 int16 = 399
    var v3 bool = true
    var v4 int64 = 9223372036854775807
    switch 0 {
    case 127:
        var v17 int32 = 573
        _ = v17
    case 4294967295:
        v3 = true
    case 2147483647:
        v2 -= 502
    default:
        v18 := GenS2{F0: 0}
        _ = v18
    }
    var v19 uint = 65535
    v1["a"] = 32767
    delete(v1, ("" + ("\n\t" + ("\n\t" + (((("a" + ("" + "a")) + "foo bar") + "\n\t") + "foo bar")))))
    fmt.Println(v19, v19)
    v2 -= (v2 + v2)
    if (v4 <= 9223372036854775807) {
        v20 := GenS2{F0: int32(v2)}
        v1[((("foo bar" + "hello") + "hello") + "foo bar")] = 128
        fmt.Println(v19, v3, v19)
        _ = v20
    }
    fmt.Println(v4)
    v21 := []int{-1, 1, 4294967295}
    fmt.Println(v4)
    if true {
        v19 -= (v19 - (v19 - v19))
        v3 = (!false)
        _ = len(v1)
    }
    if len(v21) > 0 {
        v21[0] = 1
    }
    delete(v1, "")
    v21 = append(v21, -9223372036854775808)
    v4 = v4
    var v22 float32 = 1e10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
