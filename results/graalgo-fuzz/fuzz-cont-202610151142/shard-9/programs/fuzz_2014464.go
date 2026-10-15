package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int32) uint64 {
    if (964 != 2) {
        var v4 uint = 65535
        var v5 uint32 = 2147483648
        var v6 int64 = 42
        _ = v4
        _ = v5
        _ = v6
    } else {
        v7 := []int{-1}
        for fi8 := 0; fi8 < 8; fi8++ {
            p0 += (^(p0))
            _ = fi8
        }
        _ = v7
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        var v10 int = fi9
        v11 := GenS1{F0: 42, F1: ("foo bar" + "\n\t")}
        for fi12 := 0; fi12 < 4; fi12++ {
            var v13 uint64 = 238
            v14 := GenS1{F0: 4294967295, F1: "\n\t"}
            v10 += -631
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = p0
    return 2
}
func genFn1() uint8 {
    var v15 uint32 = 42
    var v16 uint32 = v15
    for fi17 := 0; fi17 < 4; fi17++ {
        if true {
            switch "a" {
            case "\n\t":
                v18 := []int{(fi17 % ((11) | 1)), fi17, fi17}
                _ = v18
            case "hello":
                v19 := GenS1{F0: v16, F1: "hello"}
                _ = v19
            default:
                for fi20 := 0; fi20 < 14; fi20++ {
                    v21 := []int{-893, (fi20 / ((615) | 1)), -9223372036854775808}
                    v22 := GenS2{F0: 0}
                    var v23 int32 = 684
                    v16 -= (^(v15))
                    _ = fi20
                    _ = v21
                    _ = v22
                    _ = v23
                }
            }
        }
        v24 := GenS1{F0: (^(v15)), F1: "\n\t"}
        _ = fi17
        _ = v24
    }
    var v25 uint64 = genFn0(485)
    switch 1 {
    case 1:
        v26 := []int{}
        _ = v26
    default:
        switch "\n\t" {
        case "\n\t":
            fmt.Println(v25, v16, v15)
        }
    }
    _ = v15
    _ = v16
    _ = v25
    return 1
}

func main() {
    var v1 uint8 = 0
    var v2 int64 = -1
    var v3 int32 = -85
    var v27 float32 = 1e10
    for fi28 := 0; fi28 < 18; fi28++ {
        var v29 int16 = 738
        v29 = v29
        _ = fi28
        _ = v29
    }
    fmt.Println(v27)
    v30 := GenS2{F0: uint16(v1)}
    var v31 uint8 = 174
    var v32 uint = 42
    var v33 int16 = 0
    v34 := GenS2{F0: uint16(v32)}
    if (!(v33 >= v33)) {
        var v35 bool = false
        v27 += 3.14159
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
}
