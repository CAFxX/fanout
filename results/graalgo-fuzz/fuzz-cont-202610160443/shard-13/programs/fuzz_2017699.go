package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 uint8; F1 uint32; F2 uint }
type GenS4 struct { F0 uint8; F1 int64 }
func genFn0(p0 uint32, p1 uint8, p2 string) bool {
    var v6 uint = 784
    p0 = (p0 * (^(p0)))
    var v7 uint = v6
    v6 -= (v7 + v7)
    switch "\n\t" {
    case "hello":
        if (!(true || (-108 <= -128))) {
            fmt.Println(p1)
        } else {
            v8 := []int{32767, -1}
            _ = v8
        }
    case "a":
        var v9 uint8 = 67
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return false
}
func genFn1(p0 int) uint16 {
    v10 := GenS1{F0: "foo bar", F1: 1}
    v11 := []int{p0, (-(p0))}
    var v12 uint = 18446744073709551615
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    return 414
}

func main() {
    var v1 string = (((("" + "foo bar") + "hello") + "") + "")
    v2 := make(map[string]bool)
    var v3 uint8 = 28
    var v4 bool = false
    var v5 float64 = 1e10
    v5 += v5
    for fi13 := 0; fi13 < 20; fi13++ {
        v5 = 1.5
        _ = fi13
    }
    var v14 int32 = 72
    for fi15 := 0; fi15 < 20; fi15++ {
        var v16 int64 = 9223372036854775807
        for fi17 := 0; fi17 < 16; fi17++ {
            var v18 float64 = v5
            switch 9223372036854775807 {
            case -2147483648:
                v19 := GenS2{F0: v1, F1: 500}
                _ = v19
            case 32767:
                v20 := GenS3{F0: v3, F1: 806, F2: 289}
                v21 := GenS2{F0: v1, F1: 42}
                fmt.Println(v14, v18)
                _ = v20
                _ = v21
            case 9223372036854775807:
                v22 := GenS2{F0: v1, F1: 700}
                _ = v22
            }
            _ = fi17
            _ = v18
        }
        v23 := GenS3{F0: v3, F1: 2, F2: 18446744073709551615}
        _ = fi15
        _ = v16
        _ = v23
    }
    switch 1 {
    case 2147483647:
        fmt.Println(v1)
    default:
        v14 = (v14 ^ (v14 ^ v14))
    }
    var v24 uint8 = (v3 | v3)
    var v25 int64 = -327
    switch -1 {
    case 435:
        for fi26 := 0; fi26 < 19; fi26++ {
            v27 := GenS1{F0: v1, F1: 32767}
            _ = fi26
            _ = v27
        }
    case 127:
        v24 -= v24
    case 917:
        v14 = -2147483648
    default:
        v28 := GenS2{F0: (("a" + (v1 + v1)) + (v1 + v1)), F1: 1000}
        _ = v28
    }
    if v4 {
        _ = len(v2)
        fmt.Println(v14)
    }
    for fi29 := 0; fi29 < 10; fi29++ {
        v3 = v24
        var v30 int = (fi29 & fi29)
        _ = fi29
        _ = v30
    }
    v24 = (v24 << ((255) & 63))
    v2[v1] = (v4 || (!(347 >= 4294967295)))
    switch v1 {
    case "\n\t":
        for fi31 := 0; fi31 < 5; fi31++ {
            var v32 string = v1
            v33 := []int{}
            _ = fi31
            _ = v32
            _ = v33
        }
    default:
        var v34 int32 = v14
        _ = v34
    }
    v25 += 42
    var v35 int8 = -99
    if v4 {
        for fi36 := 0; fi36 < 7; fi36++ {
            v37 := GenS2{F0: ("a" + v1), F1: 1}
            v38 := GenS4{F0: 225, F1: v25}
            fmt.Println(v1)
            _ = fi36
            _ = v37
            _ = v38
        }
    }
    for fi39 := 0; fi39 < 12; fi39++ {
        switch fi39 {
        case 128:
            v40 := GenS3{F0: (v24 * 127), F1: 4294967295, F2: 1}
            _ = v40
        }
        for fi41 := 0; fi41 < 1; fi41++ {
            switch v1 {
            case "a":
                v25 = (v25 ^ v25)
            }
            var v42 int32 = (-(v14))
            _ = len(v2)
            v43 := GenS1{F0: ((("a" + "a") + "\n\t") + "\n\t"), F1: 893}
            _ = fi41
            _ = v42
            _ = v43
        }
        _ = fi39
    }
    var v44 string = v1
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v35
    _ = v4
    _ = v44
    _ = v5
}
