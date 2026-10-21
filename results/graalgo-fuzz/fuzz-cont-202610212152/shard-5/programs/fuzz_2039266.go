package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 uint16; F1 uint8; F2 string }
func genFn0(p0 uint) bool {
    var v6 uint16 = 42
    var v7 uint = 510
    switch "a" {
    case "foo bar":
        v8 := GenS1{F0: 25, F1: (v7 % (((v7 & (p0 >> ((p0) & 63)))) | 1))}
        _ = v8
    }
    _ = p0
    _ = v6
    _ = v7
    return true
}
func genFn1() int64 {
    if false {
        v9 := []int{-2147483648, 507, 742, 9223372036854775807}
        var v10 int = -923
        switch ("foo bar" + "\n\t") {
        case "hello":
            v10 = v10
        default:
            var v11 string = ""
            _ = v11
        }
        _ = v10
        _ = v9
    } else {
        var v12 int32 = 42
        var v13 float32 = 1.5
        v14 := []int{1, 2147483647}
        _ = v12
        _ = v13
        _ = v14
    }
    var v15 uint64 = 18446744073709551615
    if true {
        v15 += v15
    }
    v15 = v15
    var v16 uint = 4294967295
    _ = v15
    _ = v16
    return 42
}

func main() {
    v1 := []int{-863, -943, 42, 927}
    var v2 bool = true
    var v3 string = ("" + "foo bar")
    var v4 uint8 = 44
    var v5 float64 = 1.5
    v2 = false
    var v17 string = ((v3 + (v3 + ("a" + ((("foo bar" + "a") + v3) + v3)))) + ("foo bar" + (v3 + (v3 + "\n\t"))))
    for fi18 := 0; fi18 < 16; fi18++ {
        v19 := GenS1{F0: 610, F1: 654}
        v4 += 204
        _ = fi18
        _ = v19
    }
    var v20 uint16 = 456
    v21 := GenS1{F0: 4294967295, F1: 4294967295}
    if ((0 >= 591) || (399 != 4294967295)) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v22 uint16 = v20
        if (!(1 >= 310)) {
            fmt.Println(v2)
            v23 := GenS1{F0: 4294967295, F1: 780}
            _ = v23
        }
        _ = v22
    } else {
        for fi24 := 0; fi24 < 1; fi24++ {
            var v25 int8 = -128
            var v26 int = fi24
            _ = fi24
            _ = v25
            _ = v26
        }
        for fi27 := 0; fi27 < 15; fi27++ {
            v17 = ("\n\t" + ((v3 + (v3 + v17)) + v17))
            v28 := GenS1{F0: 42, F1: 4294967295}
            _ = fi27
            _ = v28
        }
        if (494 < 4294967295) {
            v2 = v2
        } else {
            v3 = v3
        }
    }
    fmt.Println(v20, v4, v3)
    v29 := GenS2{F0: (^(v20)), F1: 0, F2: (v17 + ("\n\t" + ("\n\t" + ((v3 + (v3 + "")) + (v17 + (("\n\t" + "foo bar") + (v3 + v3)))))))}
    v3 = ((v3 + "foo bar") + v17)
    for fi30 := 0; fi30 < 19; fi30++ {
        switch (-(fi30)) {
        case -357:
            var v31 float32 = 1e10
            _ = v31
        case 42:
            v17 = "hello"
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi30
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
