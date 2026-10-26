package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 uint64 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int64 = -9223372036854775808
    var v3 float32 = 0.1
    var v4 int8 = 127
    var v5 string = ("hello" + "a")
    var v6 int16 = 1
    _ = len(v1)
    var v7 int = 42
    v8 := v1[2147483647]
    for fi9 := 0; fi9 < 9; fi9++ {
        v10 := []int{9223372036854775807}
        _ = fi9
        _ = v10
    }
    v7 += v7
    var v11 uint32 = 65535
    switch ((v5 + ("a" + "a")) + v8) {
    case "a":
        var v12 int16 = v6
        _ = v12
    }
    v1[(v11 << ((65535) & 63))] = ((v5 + v5) + v8)
    v13 := []int{(-(v7))}
    if true {
        v3 = float32(float64(float32(0.0)))
    }
    switch (v7 % (((v7 >> ((65535) & 63))) | 1)) {
    case 128:
        switch (v7 / (((v7 * (v7 ^ -1))) | 1)) {
        case 398:
            var v14 uint64 = 605
            _ = v14
        case 32767:
            if false {
                var v15 bool = ((v11 % ((v11) | 1)) >= 827)
                if (float32(-2.25) >= float32(1.5)) {
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                } else {
                    var v16 int64 = (v2 + (v2 ^ v2))
                    var v17 string = ((v5 + (v5 + v5)) + ("foo bar" + v5))
                    for fi18 := 0; fi18 < 13; fi18++ {
                        var v19 uint32 = (v11 & v11)
                        fmt.Println(v7, v3)
                        for fi20 := 0; fi20 < 8; fi20++ {
                            fmt.Println(v15, fi20, fi20)
                            if len(v13) > 0 {
                                v13[0] = (v7 % (((fi18 >> ((18446744073709551615) & 63))) | 1))
                            }
                            v4 = (v4 + v4)
                            v21 := GenS1{F0: 2, F1: v2, F2: 1}
                            _ = fi20
                            _ = v21
                        }
                        _ = fi18
                        _ = v19
                    }
                    _ = v16
                    _ = v17
                }
                fmt.Println(v6, v5, v4)
                _ = v15
            }
        default:
            var v22 int8 = v4
            _ = v22
        }
    case 32768:
        v23 := GenS1{F0: 159, F1: 1, F2: 18446744073709551615}
        _ = v23
    case 0:
        if false {
            var v24 float64 = float64(3.14159)
            fmt.Println(v6)
            v11 = 0
            _ = v24
        }
    default:
        v3 = ((v3 / 1.0) / ((0.0 - 0.0) - float32(1e10)))
    }
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    var v25 float32 = (((v3 + 3.14159) - 0.1) - float32(1.5))
    switch 2147483647 {
    case -732:
        fmt.Println(v8)
    case -91:
        var v26 uint16 = 2
        _ = v26
    case 42:
        var v27 uint = 18446744073709551615
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
