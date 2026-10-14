package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }

func main() {
    var v1 float32 = (0.0 - (0.0 / ((0.1 / 1.0) - 0.0)))
    v2 := []string{"a"}
    v3 := []int{}
    var v4 int = 128
    v5 := make(map[int]uint32)
    var v6 int = 1
    var v7 string = ("\n\t" + "")
    fmt.Println(v1)
    v7 = ""
    v4 += 32768
    fmt.Println(v4, v1, v4)
    for fi8 := 0; fi8 < 16; fi8++ {
        switch ("hello" + (("" + v7) + v7)) {
        case "":
            if false {
                v5[9223372036854775807] = 2147483648
                var v9 uint8 = 42
                _ = v9
            }
        case "hello":
            var v10 uint8 = 127
            _ = v10
        }
        _ = len(v3)
        _ = fi8
    }
    switch v7 {
    case "a":
        v6 += -758
    }
    switch v6 {
    case -278:
        for fi11 := 0; fi11 < 5; fi11++ {
            var v12 string = v7
            _ = fi11
            _ = v12
        }
    case -752:
        fmt.Println(v4)
    case 523:
        v1 = v1
    default:
        var v13 int64 = -1
        _ = v13
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        v15 := GenS1{F0: 241, F1: -56}
        _ = fi14
        _ = v15
    }
    var v16 uint8 = 255
    for fi17 := 0; fi17 < 15; fi17++ {
        fmt.Println(v1, v6)
        _ = fi17
    }
    v1 = (v1 / v1)
    for fi18 := 0; fi18 < 2; fi18++ {
        for fi19 := 0; fi19 < 3; fi19++ {
            delete(v5, fi19)
            switch v7 {
            case "\n\t":
                v16 = v16
            case "a":
                var v20 int = fi18
                _ = v20
            default:
                v1 = 1e10
            }
            v21 := GenS1{F0: 0, F1: -933}
            _ = fi19
            _ = v21
        }
        for fi22 := 0; fi22 < 10; fi22++ {
            v3 = append(v3, v6)
            v1 += (v1 + v1)
            v23 := GenS1{F0: v16, F1: -223}
            _ = fi22
            _ = v23
        }
        _ = fi18
    }
    var v24 int = (^(v6))
    v25 := v5[(v6 / ((v24) | 1))]
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
