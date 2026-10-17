package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 bool }
func genFn0() int16 {
    v6 := GenS1{F0: 32767, F1: 637, F2: false}
    var v7 uint16 = 255
    _ = v6
    _ = v7
    return 42
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]string)
    var v3 string = "\n\t"
    var v4 uint64 = 42
    var v5 uint = 4294967295
    switch v3 {
    case "":
        var v8 uint32 = 162
        _ = v8
    default:
        switch "foo bar" {
        case "":
            v4 -= 1
        default:
            var v9 uint64 = (v4 - 0)
            _ = v9
        }
    }
    fmt.Println(v5, v3, v4)
    var v10 int64 = -9223372036854775808
    if (26 > 1) {
        var v11 uint64 = 2
        _ = len(v2)
        v12 := []int{}
        _ = v11
        _ = v12
    }
    var v13 float64 = ((-2.25 * ((((-2.25 + -2.25) - (float64((1.5 + (float32(-2.25) / ((3.14159 + 0.0) - 0.0)))) + 1.5)) + 1.5) - -2.25)) * 1e10)
    switch 471 {
    case -1:
        v14 := GenS1{F0: 42, F1: 255, F2: false}
        _ = v14
    }
    for fi15 := 0; fi15 < 20; fi15++ {
        v3 = "a"
        _ = fi15
    }
    switch 510 {
    case 127:
        v16 := GenS1{F0: 581, F1: 111, F2: false}
        _ = v16
    case 9223372036854775807:
        if false {
            var v17 string = "a"
            v17 = (v17 + (((v3 + "") + ((v3 + (v17 + v17)) + v3)) + v17))
            _ = v17
        }
    }
    switch 42 {
    case -185:
        for fi18 := 0; fi18 < 17; fi18++ {
            v10 -= -1
            v19 := GenS1{F0: 364, F1: 42, F2: ("foo bar" <= "hello")}
            _ = fi18
            _ = v19
        }
    case 32767:
        v4 = v4
    }
    fmt.Println(v3, v5, v13)
    v20 := GenS1{F0: genFn0(), F1: 65535, F2: false}
    _ = len(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
