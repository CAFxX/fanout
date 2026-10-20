package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 int8 }

func main() {
    var v1 int32 = -2147483648
    var v2 int64 = 1
    v3 := make(map[string]bool)
    var v4 uint = 390
    v5 := make(map[uint32]string)
    v6 := []string{"\n\t"}
    v4 = v4
    v5[569] = (("" + "\n\t") + "\n\t")
    v7 := GenS1{F0: 897}
    for fi8 := 0; fi8 < 15; fi8++ {
        if (65535 <= 1) {
            v9 := GenS2{F0: 703, F1: -1}
            var v10 uint16 = 32767
            var v11 uint8 = 128
            _ = v10
            _ = v11
            _ = v9
        }
        fmt.Println(fi8, v4)
        for fi12 := 0; fi12 < 10; fi12++ {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            switch "\n\t" {
            case "hello":
                v13 := GenS1{F0: 0}
                _ = v13
            default:
                _ = len(v5)
            }
            var v14 uint64 = 1
            _ = fi12
            _ = v14
        }
        _ = fi8
    }
    if len(v6) > 0 {
        v6[0] = "hello"
    }
    for fi15 := 0; fi15 < 1; fi15++ {
        v1 = v1
        fmt.Println(v2, v1, v4)
        _ = fi15
    }
    if (42 == 21) {
        v6 = append(v6, "foo bar")
        if true {
            var v16 float64 = float64((1.5 - float32((0.0 - 1e10))))
            _ = v16
        } else {
            var v17 int8 = 0
            switch "a" {
            case "":
                if (1000 > -32768) {
                    if (!(4294967295 <= 407)) {
                        v18 := GenS2{F0: 192, F1: (v17 & (-(v17)))}
                        _ = v18
                    }
                    switch -2147483648 {
                    case -1:
                        v4 += v4
                    case 428:
                        delete(v3, ("hello" + "\n\t"))
                    default:
                        v1 += (v1 * v1)
                    }
                }
            case "a":
                switch -675 {
                case -484:
                    var v19 float64 = 1.5
                    _ = v19
                default:
                    _ = len(v6)
                }
            }
            var v20 int8 = v17
            _ = v17
            _ = v20
        }
    } else {
        var v21 uint = v4
        _ = v21
    }
    if false {
        _ = len(v5)
        fmt.Println(v4)
    }
    var v22 float64 = float64(3.14159)
    fmt.Println(v22, v4)
    v5[42] = (("hello" + "foo bar") + ("hello" + ("foo bar" + "\n\t")))
    switch "\n\t" {
    case "hello":
        v22 = (float64(1.5) * (0.1 * v22))
    case "a":
        var v23 float64 = 1.5
        _ = v23
    case "foo bar":
        v4 += 686
    }
    v4 = 42
    var v24 string = ("\n\t" + "")
    _ = len(v3)
    if (1 <= 167) {
        v6 = append(v6, v24)
        v4 = 937
        fmt.Println(v4, v24)
    } else {
        v25 := GenS1{F0: 99}
        var v26 float32 = (0.0 - (1.5 / 1e10))
        _ = v25
        _ = v26
    }
    var v27 uint = (^(v4))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v24
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
