package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
func genFn0() int16 {
    if (42 > -2147483648) {
        v5 := []int{-141, 128, 4294967295, 32768}
        switch -415 {
        case -9223372036854775808:
            var v6 int = -960
            _ = v6
        case -165:
            switch ("\n\t" + "a") {
            case "foo bar":
                var v7 uint = 255
                _ = v7
            case "hello":
                var v8 uint8 = 169
                _ = v8
            }
        }
        _ = v5
    }
    var v9 uint8 = 24
    v10 := GenS1{F0: 0, F1: 100}
    _ = v10
    _ = v9
    return -533
}

func main() {
    var v1 float64 = (((1.5 - 1.5) / 0.1) + -2.25)
    var v2 int64 = 9223372036854775807
    var v3 uint = 955
    v4 := make(map[uint32]string)
    if (true && (42 == 42)) {
        v11 := []int{}
        _ = v11
    } else {
        v3 = v3
    }
    v3 = v3
    v1 = (1.5 * ((((v1 / float64(float32(v1))) * ((-2.25 - 1e10) - (3.14159 * float64(1e10)))) - float64(0.0)) / (0.0 - v1)))
    _ = len(v4)
    for fi12 := 0; fi12 < 7; fi12++ {
        v2 = v2
        var v13 uint64 = 18446744073709551615
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 12; fi14++ {
        fmt.Println(v2, v3, v2)
        var v15 int32 = -572
        fmt.Println(v1, v1)
        _ = fi14
        _ = v15
    }
    var v16 float64 = v1
    v17 := GenS1{F0: 556, F1: 0}
    v4[11] = "\n\t"
    var v18 int64 = v2
    var v19 uint16 = 143
    v20 := []int{32767, 0}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
