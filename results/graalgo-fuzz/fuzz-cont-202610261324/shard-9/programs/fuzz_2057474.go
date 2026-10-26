package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 int }
func genFn0(p0 uint16, p1 string, p2 uint) int {
    v4 := GenS1{F0: 4294967295, F1: 349, F2: 32768}
    if true {
        var v5 string = ("hello" + "foo bar")
        if false {
            fmt.Println(p0, p0)
            v6 := []int{4294967295, 1, -2147483648, -41}
            var v7 int32 = 528
            _ = v6
            _ = v7
        } else {
            v8 := []int{32768, 9223372036854775807}
            var v9 uint = (p2 - (p2 % ((p2) | 1)))
            fmt.Println(p2, p1)
            _ = v8
            _ = v9
        }
        switch 128 {
        case -195:
            v10 := GenS1{F0: 348, F1: 2147483648, F2: 42}
            _ = v10
        case -142:
            v11 := []int{}
            _ = v11
        default:
            v12 := []int{-841}
            _ = v12
        }
        _ = v5
    } else {
        for fi13 := 0; fi13 < 1; fi13++ {
            p1 = ((p1 + p1) + "\n\t")
            v14 := GenS1{F0: 4294967295, F1: 1, F2: (fi13 % ((fi13) | 1))}
            v15 := GenS1{F0: 42, F1: 65535, F2: 390}
            _ = fi13
            _ = v14
            _ = v15
        }
        p1 = p1
    }
    v16 := []int{519, 635}
    for fi17 := 0; fi17 < 20; fi17++ {
        p1 = p1
        if len(v16) > 0 {
            v16[0] = fi17
        }
        _ = fi17
    }
    fmt.Println(p2, p0, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v4
    return -2147483648
}

func main() {
    var v1 int8 = 127
    var v2 int64 = 42
    var v3 int8 = 100
    v18 := GenS1{F0: 2, F1: 152, F2: 214}
    var v19 float32 = -2.25
    if (false && false) {
        for fi20 := 0; fi20 < 14; fi20++ {
            switch int(v1) {
            case 32768:
                v1 += (v1 << ((4294967295) & 63))
            }
            v1 += v1
            v21 := []int{-1, (^(fi20))}
            _ = fi20
            _ = v21
        }
        v2 += (v2 / ((v2) | 1))
        switch "hello" {
        case "foo bar":
            var v22 float64 = 0.1
            _ = v22
        }
    }
    var v23 int64 = 143
    if (false || ((((("a" + "\n\t") + "a") + "") + "\n\t") < "\n\t")) {
        for fi24 := 0; fi24 < 7; fi24++ {
            v25 := []int{fi24, 128}
            v3 -= (^(v1))
            _ = fi24
            _ = v25
        }
        v26 := GenS1{F0: 4294967295, F1: 0, F2: 42}
        v27 := []int{0, genFn0(454, "foo bar", 1), 4294967295}
        _ = v26
        _ = v27
    }
    v19 = v19
    v1 += v1
    v28 := GenS1{F0: 42, F1: 2147483648, F2: 0}
    var v29 float64 = (float64((v19 - -2.25)) * 0.1)
    var v30 uint8 = 128
    v19 = (-2.25 + v19)
    v23 = (v23 ^ -451)
    if (1 > 1) {
        var v31 uint8 = (v30 % (((v30 ^ v30)) | 1))
        v32 := []int{-9223372036854775808, genFn0(42, "\n\t", 42), 127}
        _ = v31
        _ = v32
    }
    var v33 uint64 = 2
    v34 := []int{32767, 127, -9223372036854775808, -9223372036854775808}
    v19 -= (v19 + (v19 * (v19 - v19)))
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v33
    _ = v34
}
