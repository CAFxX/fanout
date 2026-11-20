package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 int32; F1 uint64; F2 uint }
func genFn0() uint8 {
    var v6 uint16 = 2
    var v7 uint16 = (v6 | v6)
    _ = v6
    _ = v7
    return 128
}

func main() {
    var v1 string = "a"
    v2 := []int{-9223372036854775808, 128, 32768, -9223372036854775808}
    v3 := []uint32{744}
    var v4 int8 = -1
    var v5 int16 = 1000
    v8 := GenS1{F0: -9223372036854775808, F1: 967}
    v4 -= (v4 | v4)
    if len(v3) > 0 {
        v3[0] = 42
    }
    switch v1 {
    case "a":
        var v9 uint32 = 280
        _ = v9
    }
    fmt.Println(v1)
    for fi10 := 0; fi10 < 9; fi10++ {
        v11 := GenS1{F0: -702, F1: 42}
        _ = fi10
        _ = v11
    }
    v4 -= 42
    for fi12 := 0; fi12 < 4; fi12++ {
        var v13 int = 135
        var v14 bool = false
        _ = fi12
        _ = v13
        _ = v14
    }
    switch 1 {
    case 127:
        fmt.Println(v1, v5)
    case 32768:
        v15 := GenS2{F0: 0, F1: 18446744073709551615, F2: 65535}
        _ = v15
    case 0:
        if (!((true || (2147483647 >= 1)) && (v5 <= (v5 % ((v5) | 1))))) {
            v4 += v4
            if len(v3) > 0 {
                v3[0] = 2
            }
            if false {
                var v16 float64 = -2.25
                v1 = "hello"
                _ = v16
            }
        }
    default:
        for fi17 := 0; fi17 < 19; fi17++ {
            _ = len(v3)
            if false {
                var v18 int32 = 519
                fmt.Println(v4)
                _ = v18
            }
            _ = fi17
        }
    }
    for fi19 := 0; fi19 < 4; fi19++ {
        v20 := GenS1{F0: -1, F1: 32767}
        for fi21 := 0; fi21 < 4; fi21++ {
            var v22 int8 = v4
            switch ("hello" + "hello") {
            case "\n\t":
                var v23 float32 = (3.14159 * float32(1e10))
                _ = v23
            }
            fmt.Println(v1, v1, v4)
            _ = fi21
            _ = v22
        }
        v4 = -1
        _ = fi19
        _ = v20
    }
    v24 := GenS1{F0: 42, F1: 255}
    var v25 float64 = 0.1
    _ = len(v3)
    var v26 bool = false
    for fi27 := 0; fi27 < 2; fi27++ {
        switch v1 {
        case "foo bar":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        default:
            var v28 int8 = -69
            _ = v28
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi27
    }
    v26 = (!(!(0 != 421)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
