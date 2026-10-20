package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 uint16 }
func genFn0() int64 {
    switch ("\n\t" + ("foo bar" + "foo bar")) {
    case "hello":
        var v7 bool = false
        _ = v7
    case "":
        var v8 bool = (18446744073709551615 != 0)
        _ = v8
    }
    var v9 bool = (32767 != 0)
    fmt.Println(v9)
    switch "a" {
    case "foo bar":
        for fi10 := 0; fi10 < 3; fi10++ {
            switch "hello" {
            case "a":
                for fi11 := 0; fi11 < 12; fi11++ {
                    var v12 int64 = 375
                    _ = fi11
                    _ = v12
                }
            default:
                v13 := []int{(fi10 / ((4294967295) | 1)), fi10, -744}
                _ = v13
            }
            fmt.Println(fi10)
            for fi14 := 0; fi14 < 20; fi14++ {
                fmt.Println(fi10, fi14)
                v15 := GenS1{F0: 0, F1: 1, F2: 32767}
                _ = fi14
                _ = v15
            }
            _ = fi10
        }
    default:
        var v16 uint32 = 2147483648
        _ = v16
    }
    for fi17 := 0; fi17 < 6; fi17++ {
        var v18 uint64 = 0
        if false {
            v9 = false
        }
        _ = fi17
        _ = v18
    }
    _ = v9
    return 62
}

func main() {
    var v1 string = "a"
    var v2 bool = false
    var v3 string = v1
    v4 := []uint32{42, 2147483648, 42}
    var v5 int32 = -209
    v6 := make(map[uint32]string)
    var v19 uint = 42
    for fi20 := 0; fi20 < 1; fi20++ {
        var v21 uint8 = 1
        switch int(v21) {
        case 107:
            var v22 int16 = 42
            _ = v22
        case 617:
            _ = len(v4)
        case 127:
            for fi23 := 0; fi23 < 19; fi23++ {
                var v24 int8 = -1
                v5 -= (v5 + 2147483647)
                var v25 string = (((v1 + ((("a" + "foo bar") + ("hello" + ("" + "a"))) + (v1 + v1))) + v3) + v1)
                _ = fi23
                _ = v24
                _ = v25
            }
        default:
            var v26 int = (fi20 + 32767)
            _ = v26
        }
        _ = fi20
        _ = v21
    }
    var v27 uint32 = 65535
    v28 := GenS1{F0: -9223372036854775808, F1: 42, F2: 42}
    var v29 float32 = 0.0
    if v2 {
        v30 := GenS1{F0: genFn0(), F1: -2, F2: 2}
        _ = v30
    }
    if true {
        var v31 bool = false
        _ = v31
    } else {
        var v32 uint64 = 18446744073709551615
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v32
    }
    fmt.Println(v1, v19)
    fmt.Println(v19)
    v3 = v1
    var v33 uint64 = 0
    var v34 float64 = 3.14159
    for fi35 := 0; fi35 < 5; fi35++ {
        fmt.Println(fi35)
        v19 = v19
        _ = fi35
    }
    switch "" {
    case "foo bar":
        v5 = 2147483647
    default:
        v1 = "hello"
    }
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
