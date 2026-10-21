package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 uint }
func genFn0() uint32 {
    var v6 int64 = 606
    var v7 float64 = (0.1 - (-2.25 - (1.5 - (0.0 + 0.1))))
    _ = v6
    _ = v7
    return 143
}

func main() {
    var v1 bool = (42 == -9)
    v2 := []uint32{42}
    v3 := make(map[string]int)
    v4 := make(map[uint32]bool)
    var v5 uint32 = 65535
    var v8 uint32 = v5
    var v9 uint = 0
    v10 := GenS1{F0: -32768}
    switch 1 {
    case 32768:
        fmt.Println(v8, v8)
    case 673:
        if (!(v1 && ((!(846 <= v9)) || true))) {
            var v11 int8 = -116
            switch "hello" {
            case "a":
                fmt.Println(v5, v9, v9)
            case "":
                v9 -= 42
            case "\n\t":
                fmt.Println(v5, v5, v5)
            }
            _ = v11
        } else {
            v12 := GenS1{F0: -540}
            _ = v12
        }
    case 737:
        v13 := GenS2{F0: 255, F1: (v9 - (v9 + 4294967295))}
        _ = v13
    default:
        v8 += genFn0()
    }
    v14 := GenS2{F0: 42, F1: v9}
    switch 1 {
    case 128:
        switch ("a" + "a") {
        case "a":
            if (!(-675 > -2147483648)) {
                var v15 uint32 = (^(v8))
                _ = v15
            }
        case "":
            fmt.Println(v9)
        case "\n\t":
            var v16 int = 1
            _ = v16
        }
    case -748:
        if (("a" + "foo bar") <= ("" + ("" + "hello"))) {
            var v17 int = 833
            _ = v17
        } else {
            for fi18 := 0; fi18 < 15; fi18++ {
                for fi19 := 0; fi19 < 8; fi19++ {
                    var v20 uint32 = 65535
                    v8 = v20
                    v5 += (^(v20))
                    _ = fi19
                    _ = v20
                }
                _ = fi18
            }
        }
    case -2147483648:
        var v21 float64 = 0.0
        _ = v21
    default:
        _ = len(v2)
    }
    v5 += v8
    fmt.Println(v9, v9, v8)
    v1 = v1
    switch 2147483647 {
    case 2147483647:
        var v22 int32 = 998
        _ = v22
    case -382:
        switch "a" {
        case "hello":
            v4[117] = (962 == 18446744073709551615)
        case "":
            var v23 bool = (9223372036854775807 != -259)
            _ = v23
        case "a":
            switch 32768 {
            case 757:
                var v24 bool = (v1 && (v9 > (v9 & uint(v8))))
                _ = v24
            default:
                v25 := GenS1{F0: 704}
                _ = v25
            }
        default:
            fmt.Println(v9, v1, v9)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
