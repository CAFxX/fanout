package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 int }
func genFn0() int16 {
    if false {
        for fi7 := 0; fi7 < 9; fi7++ {
            var v8 uint32 = 1
            _ = fi7
            _ = v8
        }
        var v9 float64 = (1e10 + -2.25)
        _ = v9
    }
    v10 := []int{-810}
    var v11 int = -914
    if (-911 == 73) {
        var v12 float64 = float64(1e10)
        _ = len(v10)
        _ = v12
    } else {
        v11 = (v11 | v11)
    }
    v11 = v11
    _ = v10
    _ = v11
    return 42
}
func genFn1(p0 int64, p1 uint16, p2 int32) uint16 {
    switch "\n\t" {
    case "a":
        var v13 string = ("" + "a")
        _ = v13
    case "hello":
        v14 := GenS1{F0: false, F1: -128}
        _ = v14
    case "\n\t":
        var v15 int8 = -128
        _ = v15
    default:
        switch 111 {
        case 0:
            var v16 bool = (2 >= 122)
            _ = v16
        }
    }
    var v17 float64 = 1.5
    var v18 int64 = (p0 >> ((18446744073709551615) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v18
    return p1
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 int64 = -416
    var v3 bool = false
    var v4 string = ("hello" + "")
    var v5 int64 = v2
    var v6 int32 = 0
    var v19 int32 = (v6 | v6)
    var v20 uint8 = 151
    var v21 int32 = v19
    var v22 string = (("\n\t" + ("foo bar" + "foo bar")) + v4)
    v3 = v3
    for fi23 := 0; fi23 < 3; fi23++ {
        for fi24 := 0; fi24 < 8; fi24++ {
            v25 := GenS2{F0: fi24}
            _ = fi24
            _ = v25
        }
        for fi26 := 0; fi26 < 5; fi26++ {
            v27 := []int{32767, (fi26 + (fi23 - fi23))}
            v20 = (^(v20))
            _ = fi26
            _ = v27
        }
        fmt.Println(v1, v5, v22)
        _ = fi23
    }
    if (genFn1(v2, 1, 996) == 255) {
        var v28 float32 = 0.0
        for fi29 := 0; fi29 < 7; fi29++ {
            v21 -= 938
            for fi30 := 0; fi30 < 11; fi30++ {
                v31 := []int{}
                _ = fi30
                _ = v31
            }
            _ = fi29
        }
        v2 += v5
        _ = v28
    } else {
        fmt.Println(v2, v22, v1)
        v2 -= -9223372036854775808
    }
    var v32 uint8 = v20
    if v3 {
        var v33 int = -9223372036854775808
        fmt.Println(v19, v20, v22)
        _ = v33
    } else {
        var v34 string = (v4 + (v4 + (v4 + ((("\n\t" + (((v4 + v4) + v4) + v22)) + "a") + "\n\t"))))
        for fi35 := 0; fi35 < 5; fi35++ {
            if true {
                var v36 string = v34
                fmt.Println(v5, v1)
                _ = v36
            }
            _ = fi35
        }
        _ = v34
    }
    fmt.Println(v3)
    if v3 {
        v4 = "\n\t"
    }
    v37 := GenS1{F0: v3, F1: 100}
    fmt.Println(v2, v22)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v32
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
