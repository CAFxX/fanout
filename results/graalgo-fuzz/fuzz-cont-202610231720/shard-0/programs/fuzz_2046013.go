package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 uint }
func genFn0(p0 uint) string {
    if true {
        for fi4 := 0; fi4 < 7; fi4++ {
            v5 := GenS1{F0: 0, F1: 2147483647, F2: 1}
            _ = fi4
            _ = v5
        }
        if ((!false) && (true || false)) {
            fmt.Println(p0)
        }
        v6 := []int{127, 127, 1, -370}
        _ = v6
    } else {
        switch ("hello" + "hello") {
        case "":
            switch 128 {
            case -2147483648:
                var v7 float64 = -2.25
                _ = v7
            case 939:
                var v8 uint16 = 0
                _ = v8
            }
        case "hello":
            var v9 uint16 = 65535
            _ = v9
        case "a":
            for fi10 := 0; fi10 < 15; fi10++ {
                fmt.Println(p0)
                var v11 float32 = 1.5
                _ = fi10
                _ = v11
            }
        default:
            var v12 int64 = -9223372036854775808
            _ = v12
        }
    }
    switch "a" {
    case "\n\t":
        p0 = 614
    case "hello":
        if true {
            for fi13 := 0; fi13 < 5; fi13++ {
                if ((fi13 & int(p0)) < (-(fi13))) {
                    v14 := []int{}
                    _ = v14
                } else {
                    var v15 int32 = 208
                    _ = v15
                }
                _ = fi13
            }
            v16 := GenS1{F0: 0, F1: 1, F2: p0}
            _ = v16
        }
    default:
        p0 = 1
    }
    fmt.Println(p0)
    _ = p0
    return ("a" + "hello")
}
func genFn1() string {
    var v17 uint64 = 36
    var v18 int64 = -847
    _ = v17
    _ = v18
    return "foo bar"
}

func main() {
    var v1 string = "\n\t"
    var v2 float32 = 0.0
    var v3 int16 = 0
    for fi19 := 0; fi19 < 14; fi19++ {
        if true {
            v2 -= v2
            fmt.Println(v2, fi19)
            v1 = v1
        }
        _ = fi19
    }
    fmt.Println(v2)
    var v20 uint64 = 4294967295
    fmt.Println(v3, v20)
    for fi21 := 0; fi21 < 10; fi21++ {
        v2 = float32(-2.25)
        _ = fi21
    }
    if false {
        v22 := GenS1{F0: 936, F1: 0, F2: 4294967295}
        _ = v22
    }
    var v23 uint16 = 255
    var v24 int = 1
    var v25 int32 = -1
    for fi26 := 0; fi26 < 9; fi26++ {
        if (((v20 | (v20 >> ((65535) & 63))) > 0) && (v23 >= 383)) {
            v20 = v20
            v1 = v1
        }
        v27 := []int{(-(v24)), int(v20), fi26, int(v20)}
        _ = fi26
        _ = v27
    }
    v28 := []int{(v24 >> ((18446744073709551615) & 63)), (^(v24))}
    if (2 == uint64(v23)) {
        v28 = append(v28, v24)
    } else {
        v24 += v24
    }
    var v29 float64 = float64((1.5 - 3.14159))
    v30 := GenS1{F0: 0, F1: (v25 & v25), F2: 18446744073709551615}
    var v31 uint = 255
    if ((-242 >= 42) || true) {
        for fi32 := 0; fi32 < 18; fi32++ {
            v20 -= v20
            _ = fi32
        }
        var v33 int8 = 0
        _ = v33
    }
    v34 := GenS1{F0: -1, F1: v25, F2: 255}
    v24 -= v24
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v34
}
