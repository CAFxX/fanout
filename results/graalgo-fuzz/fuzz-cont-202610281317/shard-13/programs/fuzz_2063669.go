package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 uint }
func genFn0(p0 int8, p1 int) uint16 {
    var v5 uint32 = 4294967295
    if false {
        p0 += p0
    }
    if (true && true) {
        var v6 int32 = 1
        var v7 uint64 = 4294967295
        var v8 uint = 65535
        _ = v6
        _ = v7
        _ = v8
    }
    switch "\n\t" {
    case "\n\t":
        for fi9 := 0; fi9 < 4; fi9++ {
            v10 := GenS1{F0: 42, F1: 1, F2: 255}
            var v11 uint = 839
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    _ = p0
    _ = p1
    _ = v5
    return 42
}
func genFn1() uint32 {
    var v12 uint = 65535
    var v13 float64 = -2.25
    var v14 uint16 = 98
    if true {
        var v15 uint32 = 2147483647
        v12 = v12
        _ = v15
    } else {
        for fi16 := 0; fi16 < 18; fi16++ {
            v12 = (v12 % ((v12) | 1))
            v12 = v12
            var v17 int64 = 1
            _ = fi16
            _ = v17
        }
        if true {
            var v18 int32 = -263
            var v19 uint16 = 65535
            _ = v18
            _ = v19
        } else {
            v20 := []int{1, 350, -516, 4294967295}
            var v21 float64 = float64(3.14159)
            var v22 bool = false
            _ = v20
            _ = v21
            _ = v22
        }
        for fi23 := 0; fi23 < 15; fi23++ {
            v14 += (v14 + v14)
            v13 = ((v13 / v13) * 0.0)
            _ = fi23
        }
    }
    fmt.Println(v12, v13, v14)
    _ = v12
    _ = v13
    _ = v14
    return 251
}

func main() {
    v1 := []uint32{2, 897, 10, 771}
    var v2 uint64 = 42
    var v3 int64 = 0
    var v4 string = ((("\n\t" + (("hello" + "foo bar") + "")) + "a") + ("hello" + "a"))
    v4 = "a"
    var v24 uint = 1
    if (true && false) {
        v25 := GenS1{F0: 1000, F1: 97, F2: 2}
        _ = v25
    }
    v2 = (v2 - (v2 % (((^(v2))) | 1)))
    for fi26 := 0; fi26 < 17; fi26++ {
        v3 -= v3
        _ = fi26
    }
    v4 = v4
    var v27 uint = (v24 >> ((42) & 63))
    var v28 uint = v27
    v3 -= (v3 * int64(v2))
    if false {
        if len(v1) > 0 {
            v1[0] = 65535
        }
    }
    var v29 uint64 = v2
    var v30 float64 = 1.5
    switch ("" + (((v4 + (v4 + ("foo bar" + "a"))) + v4) + (((v4 + "\n\t") + ("" + v4)) + ((((("" + "") + ("a" + "a")) + v4) + "") + v4)))) {
    case "\n\t":
        if false {
            var v31 int32 = -794
            _ = v31
        } else {
            v29 -= 1
            for fi32 := 0; fi32 < 14; fi32++ {
                var v33 uint8 = 2
                v30 = 3.14159
                var v34 string = v4
                _ = fi32
                _ = v33
                _ = v34
            }
            fmt.Println(v24, v30)
        }
    }
    v35 := GenS1{F0: 536, F1: 42, F2: 808}
    var v36 int64 = 1
    v30 = 0.1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v35
    _ = v36
    _ = v4
}
