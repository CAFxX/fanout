package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 int8; F1 int16 }
func genFn0(p0 uint16) string {
    p0 += p0
    if true {
        var v4 uint16 = 32767
        _ = v4
    } else {
        fmt.Println(p0)
        v5 := []int{944}
        _ = v5
    }
    var v6 int16 = 42
    for fi7 := 0; fi7 < 1; fi7++ {
        for fi8 := 0; fi8 < 20; fi8++ {
            var v9 uint = 42
            v10 := GenS1{F0: fi8, F1: (^(v9))}
            _ = fi8
            _ = v10
            _ = v9
        }
        v11 := GenS1{F0: fi7, F1: 192}
        var v12 uint16 = (^(p0))
        _ = fi7
        _ = v11
        _ = v12
    }
    _ = p0
    _ = v6
    return "foo bar"
}
func genFn1(p0 int64, p1 int) int {
    v13 := GenS2{F0: 990, F1: 187}
    v14 := []int{p1}
    var v15 float64 = 3.14159
    for fi16 := 0; fi16 < 14; fi16++ {
        for fi17 := 0; fi17 < 5; fi17++ {
            v14 = append(v14, fi16)
            var v18 int8 = 127
            _ = fi17
            _ = v18
        }
        _ = len(v14)
        _ = fi16
    }
    if len(v14) > 0 {
        v14[0] = 128
    }
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    _ = v15
    return (^(p1))
}

func main() {
    var v1 int8 = 83
    var v2 int64 = 535
    var v3 string = (("\n\t" + "foo bar") + "foo bar")
    var v19 float64 = (1e10 * 1e10)
    var v20 uint16 = 2
    v21 := GenS3{F0: 42, F1: 32767}
    var v22 uint = 255
    switch -2147483648 {
    case 976:
        for fi23 := 0; fi23 < 12; fi23++ {
            var v24 uint64 = 0
            _ = fi23
            _ = v24
        }
    case 4294967295:
        var v25 uint16 = (v20 & 65535)
        _ = v25
    default:
        fmt.Println(v22, v20, v19)
    }
    var v26 float64 = v19
    v27 := GenS2{F0: (v22 >> ((v22) & 63)), F1: uint64(v22)}
    switch ("hello" + v3) {
    case "a":
        switch (v3 + v3) {
        case "a":
            v1 = (v1 % ((v1) | 1))
        case "foo bar":
            var v28 int8 = v1
            _ = v28
        default:
            v19 = v19
        }
    case "hello":
        fmt.Println(v26, v22, v1)
    }
    fmt.Println(v26, v2, v20)
    var v29 int8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v26
    _ = v27
    _ = v29
    _ = v3
}
