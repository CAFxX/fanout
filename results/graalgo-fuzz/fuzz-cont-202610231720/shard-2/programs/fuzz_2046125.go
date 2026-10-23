package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint16 }
type GenS2 struct { F0 uint64 }
func genFn0() int8 {
    var v5 int = 2147483647
    var v6 uint32 = 65535
    _ = v5
    _ = v6
    return 35
}

func main() {
    var v1 int = 128
    v2 := []int{42, (v1 << ((4294967295) & 63))}
    var v3 uint8 = 255
    var v4 int32 = 1
    for fi7 := 0; fi7 < 13; fi7++ {
        if false {
            _ = len(v2)
            switch int(v3) {
            case 2147483647:
                var v8 int8 = genFn0()
                _ = v8
            }
            var v9 uint = 42
            _ = v9
        }
        v4 -= 1
        _ = fi7
    }
    v3 -= (v3 * v3)
    switch "hello" {
    case "hello":
        var v10 uint16 = 32767
        _ = v10
    case "foo bar":
        for fi11 := 0; fi11 < 18; fi11++ {
            fmt.Println(v4)
            v12 := GenS1{F0: 113, F1: 264, F2: 557}
            var v13 int8 = 88
            _ = fi11
            _ = v12
            _ = v13
        }
    case "":
        v4 = (^(v4))
    default:
        if (577 <= 384) {
            fmt.Println(v3, v4, v3)
            _ = len(v2)
        }
    }
    for fi14 := 0; fi14 < 7; fi14++ {
        var v15 uint8 = (v3 | (^(v3)))
        v3 = v15
        _ = fi14
        _ = v15
    }
    var v16 float32 = 1.5
    v17 := GenS1{F0: 32767, F1: 65535, F2: 437}
    v18 := GenS1{F0: 0, F1: 255, F2: 42}
    var v19 float32 = v16
    if false {
        fmt.Println(v16, v4, v1)
        v2 = append(v2, 4294967295)
    }
    if (true || (true && (-9223372036854775808 == -9223372036854775808))) {
        var v20 uint64 = 42
        var v21 int32 = -2147483648
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 18; fi22++ {
        for fi23 := 0; fi23 < 5; fi23++ {
            var v24 int32 = (^(v4))
            if (104 == (v3 & v3)) {
                fmt.Println(v1, v24, v4)
                v1 = -418
                fmt.Println(v24, v4)
            }
            _ = fi23
            _ = v24
        }
        var v25 uint = 1
        _ = fi22
        _ = v25
    }
    v26 := GenS2{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v26
    _ = v3
    _ = v4
}
