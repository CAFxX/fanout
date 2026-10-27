package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 uint32 }
type GenS4 struct { F0 uint }
func genFn0(p0 int8, p1 uint16, p2 uint32) int8 {
    for fi5 := 0; fi5 < 11; fi5++ {
        switch ("\n\t" + "hello") {
        case "a":
            v6 := GenS1{F0: 690, F1: -870, F2: 52}
            _ = v6
        case "\n\t":
            var v7 int16 = int16(fi5)
            _ = v7
        }
        v8 := GenS1{F0: 239, F1: 0, F2: (fi5 + (^(fi5)))}
        _ = fi5
        _ = v8
    }
    v9 := GenS1{F0: 0, F1: -32768, F2: 32}
    switch "\n\t" {
    case "foo bar":
        var v10 uint = 42
        _ = v10
    default:
        fmt.Println(p1)
    }
    var v11 float64 = float64(-2.25)
    if (18446744073709551615 > 269) {
        p0 = (-(p0))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v9
    return p0
}

func main() {
    var v1 uint8 = 1
    v2 := make(map[int]uint32)
    v3 := make(map[string]int)
    var v4 uint = 188
    if ((!(-2.25 == (0.1 + 1e10))) && (42 != 0)) {
        var v12 int = 617
        for fi13 := 0; fi13 < 8; fi13++ {
            if (!true) {
                v14 := []int{(v12 & 32767), (v12 - (fi13 * fi13)), v12}
                _ = v14
            } else {
                v12 += fi13
                v15 := GenS2{F0: 2}
                _ = v15
            }
            fmt.Println(v12, v1)
            _ = fi13
        }
        if false {
            if (4294967295 <= 407) {
                var v16 int8 = 100
                _ = v16
            }
        } else {
            fmt.Println(v12)
            v17 := GenS2{F0: uint8(v4)}
            fmt.Println(v4, v4, v1)
            _ = v17
        }
        _ = v12
    }
    var v18 int32 = int32(v4)
    v1 += (v1 - v1)
    if (!(1e10 >= 1e10)) {
        for fi19 := 0; fi19 < 13; fi19++ {
            v20 := GenS2{F0: v1}
            _ = fi19
            _ = v20
        }
        v4 = v4
    }
    fmt.Println(v18)
    switch -53 {
    case 127:
        if (!true) {
            var v21 int32 = v18
            var v22 int16 = 377
            v23 := []int{832, -1, 1, -9223372036854775808}
            _ = v21
            _ = v22
            _ = v23
        }
    }
    var v24 int8 = 42
    if (0.1 > (1.5 - float64(float32(((1.5 + 3.14159) - 3.14159))))) {
        v25 := v2[-9223372036854775808]
        _ = v25
    }
    if true {
        for fi26 := 0; fi26 < 15; fi26++ {
            var v27 string = ("\n\t" + (("" + "foo bar") + ("a" + ("" + ("a" + ("hello" + ("hello" + ("a" + "hello"))))))))
            _ = fi26
            _ = v27
        }
        for fi28 := 0; fi28 < 18; fi28++ {
            v29 := GenS3{F0: 945, F1: 4294967295}
            v30 := []int{fi28, (^(fi28)), (^(fi28)), 127}
            _ = fi28
            _ = v29
            _ = v30
        }
    }
    v31 := GenS4{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v24
    _ = v3
    _ = v31
    _ = v4
}
