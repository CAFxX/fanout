package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 int8; F1 uint16; F2 uint }
func genFn0(p0 string, p1 string) bool {
    if (2 <= 0) {
        var v4 int32 = 2147483647
        var v5 int16 = -816
        var v6 int64 = 327
        _ = v4
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: -95, F1: 18446744073709551615}
    var v8 int32 = -159
    if (true || (!false)) {
        v8 = (-(v8))
        var v9 float32 = 1e10
        p0 = ("\n\t" + p1)
        _ = v9
    }
    v8 -= (v8 | v8)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return false
}
func genFn1() uint64 {
    v10 := []int{214, 2147483647}
    for fi11 := 0; fi11 < 16; fi11++ {
        fmt.Println(fi11)
        _ = fi11
    }
    v12 := GenS1{F0: -9223372036854775808, F1: 18446744073709551615}
    _ = v10
    _ = v12
    return 4294967295
}

func main() {
    v1 := []int{-234, 2147483647}
    var v2 string = ("hello" + "a")
    v3 := []string{(((v2 + ((("hello" + "foo bar") + v2) + v2)) + v2) + v2)}
    v1 = append(v1, -1)
    if (v2 == v2) {
        v2 = v2
        v1 = append(v1, -4)
        var v13 uint32 = 130
        _ = v13
    } else {
        v2 = "foo bar"
    }
    var v14 uint = 65535
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v15 float64 = 0.0
    v16 := GenS1{F0: -549, F1: 4294967295}
    v15 -= v15
    if true {
        switch -142 {
        case 2147483647:
            v17 := GenS1{F0: -527, F1: 0}
            _ = v17
        case 411:
            v14 += (^(v14))
        case 1:
            v18 := GenS1{F0: -9223372036854775808, F1: 1}
            _ = v18
        }
        for fi19 := 0; fi19 < 16; fi19++ {
            v14 = (v14 - v14)
            var v20 uint8 = 42
            for fi21 := 0; fi21 < 2; fi21++ {
                switch fi19 {
                case 257:
                    v22 := GenS1{F0: -1, F1: 1}
                    _ = v22
                default:
                    v23 := GenS1{F0: 42, F1: genFn1()}
                    var v24 float64 = v15
                    _ = v23
                    _ = v24
                }
                var v25 uint8 = (v20 ^ 127)
                fmt.Println(v20, v2, fi19)
                if len(v1) > 0 {
                    v1[0] = (fi19 >> (((v14 ^ 0)) & 63))
                }
                v26 := GenS2{F0: 25, F1: 2, F2: 0}
                _ = fi21
                _ = v25
                _ = v26
            }
            _ = fi19
            _ = v20
        }
        _ = len(v3)
    }
    for fi27 := 0; fi27 < 17; fi27++ {
        switch (fi27 - fi27) {
        case -2147483648:
            if (258 >= 405) {
                for fi28 := 0; fi28 < 13; fi28++ {
                    var v29 int64 = 42
                    var v30 uint32 = 8
                    _ = fi28
                    _ = v29
                    _ = v30
                }
            } else {
                var v31 int16 = 1000
                _ = v31
            }
        case -105:
            v1 = append(v1, (fi27 % (((fi27 ^ (fi27 - (fi27 + (fi27 * (fi27 ^ (fi27 | -870))))))) | 1)))
        }
        v14 -= v14
        _ = fi27
    }
    v2 = (v2 + ((("hello" + (v2 + (v2 + v2))) + (v2 + v2)) + (v2 + v2)))
    var v32 int32 = -2147483648
    v33 := GenS2{F0: 127, F1: 42, F2: v14}
    var v34 int = 242
    fmt.Println(v34, v14)
    for fi35 := 0; fi35 < 3; fi35++ {
        if (42 > 0) {
            var v36 int = -2147483648
            _ = v36
        }
        _ = fi35
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
}
