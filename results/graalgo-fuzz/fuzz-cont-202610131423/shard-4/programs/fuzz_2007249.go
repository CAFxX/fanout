package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
func genFn0(p0 int) uint32 {
    for fi5 := 0; fi5 < 6; fi5++ {
        p0 += (fi5 - p0)
        v6 := []int{p0, p0, 2147483647, (-(p0))}
        _ = fi5
        _ = v6
    }
    var v7 uint32 = 925
    v8 := []int{}
    _ = p0
    _ = v7
    _ = v8
    return v7
}
func genFn1(p0 int, p1 bool) int {
    for fi9 := 0; fi9 < 13; fi9++ {
        v10 := []int{p0, -9223372036854775808, 42}
        if p1 {
            var v11 int8 = 127
            var v12 uint16 = 255
            var v13 uint64 = 18446744073709551615
            _ = v11
            _ = v12
            _ = v13
        } else {
            v14 := GenS1{F0: fi9, F1: 0}
            _ = v14
        }
        var v15 int8 = 49
        _ = fi9
        _ = v10
        _ = v15
    }
    if p1 {
        switch -2147483648 {
        case 4294967295:
            var v16 uint = 819
            _ = v16
        case 1:
            p0 = -2147483648
        case -2147483648:
            fmt.Println(p0)
        default:
            var v17 string = ("foo bar" + ("hello" + "\n\t"))
            _ = v17
        }
    } else {
        p1 = true
        if (2 <= 42) {
            v18 := []int{p0}
            _ = len(v18)
            _ = v18
        } else {
            p1 = p1
        }
    }
    var v19 int16 = 32767
    p1 = true
    _ = p0
    _ = p1
    _ = v19
    return -264
}

func main() {
    var v1 uint64 = 182
    var v2 uint32 = 2
    v3 := []uint32{0, v2, v2, (v2 + v2)}
    var v4 float32 = 1.5
    var v20 int8 = -43
    fmt.Println(v4, v2, v2)
    fmt.Println(v4, v4)
    switch ("hello" + ((("\n\t" + "hello") + "a") + "a")) {
    case "foo bar":
        if (true && (!true)) {
            if len(v3) > 0 {
                v3[0] = v2
            }
            if ((1 <= 362) || false) {
                var v21 string = "hello"
                v3 = append(v3, (v2 ^ (v2 >> ((18446744073709551615) & 63))))
                _ = v21
            } else {
                for fi22 := 0; fi22 < 13; fi22++ {
                    v4 = v4
                    _ = fi22
                }
                var v23 float32 = v4
                _ = v23
            }
        } else {
            if false {
                for fi24 := 0; fi24 < 13; fi24++ {
                    v1 -= (v1 & (^(v1)))
                    var v25 uint32 = (v2 * genFn0(fi24))
                    _ = fi24
                    _ = v25
                }
                for fi26 := 0; fi26 < 2; fi26++ {
                    v27 := GenS1{F0: genFn1(fi26, false), F1: 124}
                    v28 := GenS1{F0: -582, F1: 2147483647}
                    _ = fi26
                    _ = v27
                    _ = v28
                }
            } else {
                v1 += 4294967295
                v2 = v2
                v29 := GenS1{F0: 32767, F1: -98}
                _ = v29
            }
        }
    }
    if false {
        var v30 float64 = ((-2.25 - 3.14159) / 1.0)
        _ = v30
    }
    fmt.Println(v1)
    v31 := GenS1{F0: 127, F1: -1}
    fmt.Println(v20)
    var v32 bool = false
    v2 = (^(v2))
    v20 -= v20
    v3 = append(v3, 42)
    v4 -= v4
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v31
    _ = v32
    _ = v4
}
