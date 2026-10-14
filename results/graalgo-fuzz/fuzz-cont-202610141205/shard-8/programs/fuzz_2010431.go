package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint; F1 int32; F2 int16 }

func main() {
    var v1 float64 = (1.5 - float64(3.14159))
    v2 := make(map[int]string)
    var v3 int8 = -128
    var v4 int64 = -9223372036854775808
    if true {
        v3 += (v3 / ((v3) | 1))
        for fi5 := 0; fi5 < 20; fi5++ {
            for fi6 := 0; fi6 < 4; fi6++ {
                v1 += ((((v1 * v1) - v1) / 1.0) / (v1 - 3.14159))
                fmt.Println(v4, fi6)
                _ = fi6
            }
            v7 := GenS1{F0: 32767, F1: 191}
            _ = fi5
            _ = v7
        }
    }
    fmt.Println(v3, v1, v3)
    v1 = float64(((0.0 * 3.14159) - 1.5))
    if false {
        var v8 uint32 = 1
        v9 := GenS1{F0: 347, F1: 127}
        v10 := []int{-624, -53}
        _ = v10
        _ = v8
        _ = v9
    }
    var v11 string = ("a" + "foo bar")
    for fi12 := 0; fi12 < 7; fi12++ {
        v13 := GenS1{F0: 0, F1: 42}
        var v14 int8 = (v3 | (v3 / ((v3) | 1)))
        switch "" {
        case "hello":
            var v15 int8 = -1
            _ = v15
        case "\n\t":
            v16 := GenS1{F0: 65535, F1: 127}
            _ = v16
        case "a":
            var v17 uint16 = 0
            _ = v17
        default:
            for fi18 := 0; fi18 < 18; fi18++ {
                if (false || false) {
                    var v19 uint64 = 4294967295
                    v14 += 100
                    _ = v19
                }
                delete(v2, fi12)
                v3 -= v3
                _ = fi18
            }
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    for fi20 := 0; fi20 < 9; fi20++ {
        v21 := GenS1{F0: 1, F1: 127}
        _ = fi20
        _ = v21
    }
    if (!(!true)) {
        var v22 int64 = v4
        _ = v22
    } else {
        if ((-32768 >= 32767) || (false && (v3 >= (v3 % ((v3) | 1))))) {
            v23 := GenS2{F0: 65535, F1: 1, F2: 762}
            _ = v23
        } else {
            var v24 int8 = v3
            v2[0] = "\n\t"
            v24 = (v24 & 84)
            _ = v24
        }
        v11 = ""
        fmt.Println(v11)
    }
    fmt.Println(v3, v3, v11)
    fmt.Println(v11, v3, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
}
