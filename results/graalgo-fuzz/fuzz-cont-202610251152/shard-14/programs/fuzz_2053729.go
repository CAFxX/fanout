package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 string }
type GenS2 struct { F0 int8; F1 uint }
func genFn0(p0 uint32) int16 {
    var v6 string = ("hello" + "hello")
    var v7 uint16 = 32767
    var v8 int8 = 102
    switch 0 {
    case 553:
        var v9 float64 = 0.0
        _ = v9
    case 128:
        v10 := GenS1{F0: 127, F1: 0, F2: ("foo bar" + "\n\t")}
        _ = v10
    }
    for fi11 := 0; fi11 < 13; fi11++ {
        switch 4294967295 {
        case 127:
            var v12 uint16 = v7
            _ = v12
        }
        if false {
            v8 -= 0
            var v13 string = ((v6 + v6) + "a")
            _ = v13
        }
        _ = fi11
    }
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 uint16 = 520
    var v2 string = ""
    var v3 string = ((v2 + v2) + v2)
    v4 := make(map[uint32]bool)
    v5 := []int{32767, 661, 32767}
    switch -179 {
    case 32767:
        v2 = v2
    case 0:
        v14 := GenS2{F0: 52, F1: 4294967295}
        _ = v14
    case 128:
        v1 += (^(v1))
    }
    if (0.0 <= 0.1) {
        fmt.Println(v2, v2)
        if true {
            v1 = 695
            v2 = v3
        }
    }
    for fi15 := 0; fi15 < 12; fi15++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v16 := GenS1{F0: 0, F1: 2147483647, F2: v3}
        for fi17 := 0; fi17 < 7; fi17++ {
            v1 -= v1
            v1 -= (v1 / ((65535) | 1))
            _ = fi17
        }
        _ = fi15
        _ = v16
    }
    v2 = v2
    v3 = v3
    var v18 string = "hello"
    if false {
        for fi19 := 0; fi19 < 12; fi19++ {
            fmt.Println(v1, v3, v2)
            v3 = "\n\t"
            _ = fi19
        }
        var v20 int8 = 19
        var v21 uint16 = (v1 & v1)
        _ = v20
        _ = v21
    }
    v1 = (^(v1))
    v3 = "foo bar"
    var v22 int8 = 1
    v1 = v1
    if true {
        if false {
            fmt.Println(v22)
            var v23 uint16 = v1
            fmt.Println(v3, v1)
            _ = v23
        } else {
            var v24 float64 = (float64(3.14159) / 1e10)
            v25 := GenS1{F0: v22, F1: -919, F2: v2}
            if true {
                delete(v4, 88)
            }
            _ = v24
            _ = v25
        }
        var v26 uint8 = 2
        _ = v26
    }
    for fi27 := 0; fi27 < 14; fi27++ {
        v28 := GenS2{F0: (v22 * v22), F1: 626}
        _ = fi27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
