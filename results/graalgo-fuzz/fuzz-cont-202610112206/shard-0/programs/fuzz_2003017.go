package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 uint32; F1 int32; F2 uint64 }
type GenS3 struct { F0 int16; F1 uint32; F2 uint }
type GenS4 struct { F0 uint64 }

func main() {
    var v1 float64 = 3.14159
    v2 := []string{"\n\t", ("hello" + "a")}
    var v3 int32 = -271
    if true {
        switch -2147483648 {
        case 4294967295:
            fmt.Println(v1, v1)
        case -2147483648:
            fmt.Println(v1, v3)
        default:
            fmt.Println(v3)
        }
        v4 := GenS1{F0: 127, F1: 65535}
        _ = v4
    }
    if true {
        var v5 int32 = -927
        for fi6 := 0; fi6 < 6; fi6++ {
            if true {
                var v7 uint64 = 42
                _ = v7
            } else {
                v3 = v5
            }
            var v8 string = "hello"
            _ = fi6
            _ = v8
        }
        _ = v5
    } else {
        v1 -= v1
    }
    v3 -= (v3 + (v3 | (-(v3))))
    v9 := GenS2{F0: 129, F1: 1, F2: 1}
    var v10 int32 = (-(v3))
    _ = len(v2)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi11 := 0; fi11 < 13; fi11++ {
        v12 := GenS3{F0: -32768, F1: 65535, F2: 303}
        var v13 int16 = -32768
        _ = fi11
        _ = v12
        _ = v13
    }
    v14 := GenS4{F0: 18446744073709551615}
    v10 = 1
    if false {
        _ = len(v2)
        if false {
            v1 -= v1
            fmt.Println(v3, v1, v10)
            var v15 uint = 4294967295
            _ = v15
        }
    } else {
        if (true && (-32768 >= -1)) {
            var v16 uint = 1
            var v17 uint16 = 1
            for fi18 := 0; fi18 < 3; fi18++ {
                v19 := GenS4{F0: 1}
                _ = fi18
                _ = v19
            }
            _ = v16
            _ = v17
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v9
}
