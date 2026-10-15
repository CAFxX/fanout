package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0() uint16 {
    var v6 uint = 4294967295
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := GenS1{F0: (!(("hello" + ("\n\t" + (("hello" + "hello") + "hello"))) >= ("" + "foo bar")))}
        var v9 int32 = -674
        v10 := []int{}
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    v11 := []int{-1, 128, 9223372036854775807}
    _ = v11
    _ = v6
    return 1
}
func genFn1() int64 {
    if false {
        for fi12 := 0; fi12 < 18; fi12++ {
            fmt.Println(fi12)
            fmt.Println(fi12)
            _ = fi12
        }
    }
    v13 := []int{42, -929}
    _ = v13
    return 42
}

func main() {
    var v1 uint = 2
    var v2 bool = true
    var v3 string = "\n\t"
    var v4 bool = v2
    v5 := []int{-67, 9223372036854775807, -319}
    var v14 int64 = genFn1()
    var v15 float32 = 1.5
    for fi16 := 0; fi16 < 19; fi16++ {
        var v17 float64 = 1e10
        _ = fi16
        _ = v17
    }
    var v18 uint16 = 695
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    switch -2147483648 {
    case -1:
        var v19 uint16 = v18
        _ = v19
    case -133:
        var v20 uint8 = 1
        _ = v20
    }
    var v21 uint64 = 4294967295
    for fi22 := 0; fi22 < 6; fi22++ {
        switch int(v14) {
        case 312:
            var v23 uint32 = 632
            _ = v23
        case 32768:
            v3 = ""
        case 32767:
            if ((v14 > 1) || v2) {
                fmt.Println(fi22, v3)
                var v24 uint64 = (v21 / (((v21 << ((v1) & 63))) | 1))
                var v25 int8 = 127
                _ = v24
                _ = v25
            }
        default:
            v1 -= (^(v1))
        }
        _ = fi22
    }
    var v26 int32 = 1
    fmt.Println(v2, v3, v26)
    if ((v1 % ((42) | 1)) > (v1 - uint(v21))) {
        if true {
            var v27 int = -9223372036854775808
            fmt.Println(v15)
            if v2 {
                _ = len(v5)
            }
            _ = v27
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v21
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
