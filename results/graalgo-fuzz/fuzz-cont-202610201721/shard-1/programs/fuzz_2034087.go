package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
func genFn0(p0 uint64, p1 string, p2 uint64) int {
    for fi5 := 0; fi5 < 3; fi5++ {
        if true {
            p1 = p1
            var v6 int16 = 1
            _ = v6
        }
        _ = fi5
    }
    v7 := GenS1{F0: 0, F1: 9223372036854775807}
    var v8 uint16 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 32768
}

func main() {
    v1 := []int{-9223372036854775808, 0}
    v2 := make(map[string]uint32)
    var v3 int8 = -63
    var v4 uint64 = 1
    _ = len(v2)
    v9 := GenS1{F0: 0, F1: 9223372036854775807}
    switch "hello" {
    case "":
        if ("foo bar" != ("\n\t" + ("foo bar" + "\n\t"))) {
            v4 = v4
        } else {
            fmt.Println(v3)
            v3 += (v3 - v3)
            var v10 uint32 = 2
            _ = v10
        }
    default:
        v3 -= 1
    }
    _ = len(v1)
    var v11 int64 = 42
    for fi12 := 0; fi12 < 8; fi12++ {
        v11 += v11
        _ = fi12
    }
    v13 := GenS1{F0: 42, F1: (v11 & (v11 % (((^(v11))) | 1)))}
    switch 2147483647 {
    case 9223372036854775807:
        v14 := GenS1{F0: uint(v4), F1: (v11 * v11)}
        _ = v14
    case 32768:
        switch int(v3) {
        case -1:
            switch (("hello" + (("a" + "a") + "foo bar")) + "foo bar") {
            case "\n\t":
                for fi15 := 0; fi15 < 8; fi15++ {
                    var v16 bool = false
                    for fi17 := 0; fi17 < 4; fi17++ {
                        v11 += (v11 ^ v11)
                        _ = fi17
                    }
                    v16 = (v16 && false)
                    _ = fi15
                    _ = v16
                }
            case "a":
                for fi18 := 0; fi18 < 11; fi18++ {
                    var v19 uint = 0
                    v19 = 47
                    v20 := GenS1{F0: 0, F1: (v11 << ((1) & 63))}
                    _ = fi18
                    _ = v19
                    _ = v20
                }
            case "":
                switch "foo bar" {
                case "hello":
                    fmt.Println(v4, v4, v3)
                case "a":
                    v4 = (v4 >> ((18446744073709551615) & 63))
                }
            }
        }
    case 1:
        if false {
            if false {
                if (v11 <= v11) {
                    fmt.Println(v11)
                }
            }
            fmt.Println(v3)
        } else {
            v21 := GenS1{F0: 852, F1: v11}
            _ = v21
        }
    }
    var v22 int = 32767
    v11 = v11
    if len(v1) > 0 {
        v1[0] = (v22 ^ (v22 * v22))
    }
    if (966 == 1000) {
        v23 := v2[("hello" + ("hello" + "hello"))]
        _ = v23
    }
    if (1 == 32767) {
        v1 = append(v1, v22)
    } else {
        v4 = (v4 << ((1) & 63))
        fmt.Println(v3)
    }
    var v24 uint32 = 1
    fmt.Println(v3, v11)
    switch "" {
    case "foo bar":
        var v25 float64 = float64(1.5)
        _ = v25
    case "\n\t":
        switch genFn0((v4 % (((v4 & 42)) | 1)), "\n\t", (v4 / ((0) | 1))) {
        case -113:
            var v26 int16 = 508
            _ = v26
        default:
            if (!(true || false)) {
                switch (v22 % ((v22) | 1)) {
                case -2147483648:
                    var v27 int = v22
                    _ = v27
                case -74:
                    var v28 int = (v22 | genFn0(v4, (("hello" + "") + "hello"), v4))
                    _ = v28
                default:
                    var v29 string = ""
                    _ = v29
                }
            }
        }
    case "hello":
        switch (v22 + v22) {
        case 32768:
            var v30 int64 = (-(v11))
            _ = v30
        default:
            var v31 uint16 = uint16(v24)
            _ = v31
        }
    }
    v24 = v24
    var v32 uint16 = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v22
    _ = v24
    _ = v3
    _ = v32
    _ = v4
    _ = v9
}
