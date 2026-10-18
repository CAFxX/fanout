package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string }
func genFn0(p0 uint64) uint {
    v6 := []int{32768}
    switch -9223372036854775808 {
    case -938:
        var v7 int64 = 9223372036854775807
        _ = v7
    case 1:
        p0 -= p0
    case 32767:
        if (32768 > 24) {
            _ = len(v6)
        } else {
            p0 += p0
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        }
    default:
        fmt.Println(p0)
    }
    _ = len(v6)
    if len(v6) > 0 {
        v6[0] = -230
    }
    fmt.Println(p0)
    _ = p0
    _ = v6
    return 255
}
func genFn1(p0 uint32, p1 bool, p2 string) int8 {
    v8 := []int{127}
    v9 := GenS1{F0: -1}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return -1
}

func main() {
    var v1 uint32 = 4294967295
    v2 := []uint32{v1}
    var v3 float32 = 1e10
    var v4 uint8 = 134
    var v5 int32 = 1
    var v10 uint64 = 0
    v1 = (^(v1))
    fmt.Println(v5, v10, v5)
    v11 := GenS2{F0: "a"}
    v4 += (v4 * v4)
    if (42 > 1) {
        if ((0 > 1) && true) {
            v12 := GenS2{F0: ("hello" + ("hello" + ""))}
            if ((2 > 32767) || (346 != 65535)) {
                var v13 int64 = -46
                v4 = v4
                if (9223372036854775807 < 1) {
                    v3 -= v3
                    for fi14 := 0; fi14 < 8; fi14++ {
                        v4 = (v4 * v4)
                        _ = fi14
                    }
                    fmt.Println(v10, v10, v1)
                }
                _ = v13
            } else {
                fmt.Println(v10)
            }
            _ = v12
        }
        v10 += 42
    } else {
        var v15 uint8 = (v4 & uint8(v1))
        _ = v15
    }
    switch 128 {
    case -742:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    if len(v2) > 0 {
        v2[0] = 787
    }
    if (("foo bar" + (("a" + (("" + ("a" + "a")) + "a")) + "")) >= "foo bar") {
        var v16 int64 = 0
        v10 = (v10 * (^(v10)))
        var v17 bool = false
        _ = v16
        _ = v17
    } else {
        _ = len(v2)
        v18 := GenS1{F0: -531}
        v5 += (v5 | (v5 % (((v5 + (-(v5)))) | 1)))
        _ = v18
    }
    var v19 int32 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
