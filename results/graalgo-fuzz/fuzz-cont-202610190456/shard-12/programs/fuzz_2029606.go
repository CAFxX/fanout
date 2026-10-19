package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 uint32; F2 uint64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint; F1 int8 }
func genFn0(p0 int, p1 int8) bool {
    v4 := GenS1{F0: 18446744073709551615}
    switch "hello" {
    case "hello":
        if (!false) {
            for fi5 := 0; fi5 < 12; fi5++ {
                v6 := GenS1{F0: 18446744073709551615}
                var v7 uint32 = 0
                _ = fi5
                _ = v6
                _ = v7
            }
        } else {
            v8 := GenS1{F0: 65535}
            var v9 bool = (28 > 254)
            _ = v8
            _ = v9
        }
    default:
        v10 := []int{128, p0, p0}
        _ = v10
    }
    p1 += p1
    var v11 bool = (!(true && false))
    _ = p0
    _ = p1
    _ = v11
    _ = v4
    return v11
}
func genFn1() int {
    var v12 int16 = 1
    fmt.Println(v12)
    v13 := []int{1, 719, -936, -1}
    v12 = (v12 % (((v12 << ((2) & 63))) | 1))
    _ = v12
    _ = v13
    return 128
}

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{521, 2, 2147483647, 0}
    v3 := make(map[string]uint32)
    switch -453 {
    case 0:
        for fi14 := 0; fi14 < 17; fi14++ {
            var v15 float32 = -2.25
            v16 := v1["a"]
            _ = fi14
            _ = v15
            _ = v16
        }
    case 1:
        v17 := GenS1{F0: 424}
        _ = v17
    case 32768:
        v18 := GenS2{F0: 1, F1: 2, F2: 864}
        _ = v18
    }
    if len(v2) > 0 {
        v2[0] = 61
    }
    v19 := v3["a"]
    v20 := GenS2{F0: 128, F1: 1, F2: 4294967295}
    if false {
        v21 := GenS2{F0: 127, F1: 2147483647, F2: 832}
        switch 127 {
        case -2147483648:
            fmt.Println(v19)
        case -1:
            var v22 float32 = 1.5
            _ = v22
        case 32768:
            switch "\n\t" {
            case "\n\t":
                v23 := GenS2{F0: 255, F1: v19, F2: 0}
                _ = v23
            }
        default:
            var v24 int64 = 1
            _ = v24
        }
        v19 += 1
        _ = v21
    } else {
        if false {
            for fi25 := 0; fi25 < 5; fi25++ {
                _ = len(v2)
                _ = fi25
            }
            if (!(!(!false))) {
                var v26 int = -2147483648
                _ = v26
            }
        } else {
            v19 = (v19 % (((v19 % ((2) | 1))) | 1))
            v19 = v19
            var v27 int = genFn1()
            _ = v27
        }
    }
    var v28 uint = 386
    v29 := GenS3{F0: (("" + ("foo bar" + "hello")) + "a")}
    for fi30 := 0; fi30 < 7; fi30++ {
        v28 = 42
        _ = fi30
    }
    v2 = append(v2, 2147483647)
    fmt.Println(v19)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v31 := GenS1{F0: 255}
    if len(v2) > 0 {
        v2[0] = uint32(v28)
    }
    var v32 int8 = 100
    var v33 bool = true
    switch 32767 {
    case 32768:
        if len(v2) > 0 {
            v2[0] = v19
        }
    case -544:
        var v34 string = ""
        _ = v34
    }
    v35 := GenS1{F0: (v28 - (v28 + (v28 - v28)))}
    v36 := GenS4{F0: (^(v28)), F1: (v32 - v32)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v35
    _ = v36
}
