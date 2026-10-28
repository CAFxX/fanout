package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 uint8; F2 uint }
func genFn0() int8 {
    var v4 uint32 = 2147483648
    if (!("a" != "foo bar")) {
        fmt.Println(v4)
        var v5 bool = true
        var v6 uint = 692
        _ = v5
        _ = v6
    } else {
        v4 += v4
        for fi7 := 0; fi7 < 5; fi7++ {
            var v8 float32 = -2.25
            v9 := []int{(-(fi7)), fi7, fi7}
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    if (!(false && (!(!((!false) && true))))) {
        v4 = 2147483648
        v10 := []int{0, 0, 802}
        switch 1 {
        case -762:
            v4 += (v4 ^ 2147483647)
        }
        _ = v10
    } else {
        var v11 float32 = 3.14159
        for fi12 := 0; fi12 < 14; fi12++ {
            v13 := GenS1{F0: 1}
            v14 := GenS1{F0: -563}
            v15 := GenS1{F0: 9223372036854775807}
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
        v16 := GenS2{F0: -1, F1: 49, F2: 923}
        _ = v11
        _ = v16
    }
    _ = v4
    return 100
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float64 = (1.5 - 1e10)
    v3 := []uint32{42, 194}
    v2 += (((1.5 * float64(1e10)) * v2) + ((0.1 + float64(0.0)) + (1e10 + v2)))
    switch -820 {
    case -891:
        for fi17 := 0; fi17 < 2; fi17++ {
            _ = len(v3)
            _ = len(v3)
            _ = fi17
        }
    case 9223372036854775807:
        if false {
            delete(v1, 9223372036854775807)
        } else {
            v18 := GenS2{F0: -40, F1: 1, F2: 628}
            var v19 int = 1
            _ = v18
            _ = v19
        }
    default:
        for fi20 := 0; fi20 < 5; fi20++ {
            v21 := GenS1{F0: 777}
            var v22 int32 = 3
            _ = fi20
            _ = v21
            _ = v22
        }
    }
    if false {
        v2 += ((0.1 - (v2 / v2)) - (v2 * v2))
        var v23 string = ""
        _ = v23
    }
    if false {
        var v24 int32 = -87
        _ = v24
    }
    switch 508 {
    case 127:
        _ = len(v3)
    case 4294967295:
        v2 += 1e10
    case -630:
        fmt.Println(v2)
    }
    fmt.Println(v2)
    if true {
        v2 = (1e10 * (v2 + (v2 * 0.1)))
        v2 -= (0.1 / -2.25)
    }
    fmt.Println(v2)
    v2 = 0.1
    var v25 int64 = -9223372036854775808
    if true {
        if (542 >= -1) {
            v2 = (v2 - v2)
            var v26 uint = 1
            v3 = append(v3, 463)
            _ = v26
        }
    }
    switch ("\n\t" + ("foo bar" + (((("" + ("\n\t" + "hello")) + "hello") + ("hello" + ("hello" + ("hello" + "\n\t")))) + "a"))) {
    case "":
        var v27 uint64 = 0
        _ = v27
    }
    if (4294967295 > 793) {
        for fi28 := 0; fi28 < 11; fi28++ {
            _ = len(v3)
            if true {
                fmt.Println(fi28, v2, v25)
            }
            _ = fi28
        }
    } else {
        var v29 int64 = v25
        v2 += 1e10
        v30 := v1[2147483647]
        _ = v29
        _ = v30
    }
    var v31 string = ("a" + "\n\t")
    v31 = ""
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v31
}
