package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint; F1 bool; F2 int64 }
type GenS3 struct { F0 bool; F1 uint32; F2 int64 }
func genFn0(p0 uint64, p1 uint64, p2 uint64) int {
    v7 := []int{788, 42, -9223372036854775808, -907}
    p0 = (p1 >> ((65535) & 63))
    var v8 uint64 = (p1 << ((2) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 835
}

func main() {
    var v1 float32 = 1.5
    v2 := []int{127, 128, 128}
    var v3 uint16 = 379
    var v4 uint16 = (v3 ^ v3)
    v5 := []string{("a" + "hello"), "hello", ("a" + "hello")}
    v6 := make(map[string]string)
    switch 590 {
    case 9223372036854775807:
        v9 := v6[(("\n\t" + (("hello" + ("\n\t" + "a")) + "")) + "a")]
        _ = v9
    }
    var v10 float64 = (1e10 * 1.5)
    v1 = (((v1 + 3.14159) + 0.0) / v1)
    v3 += (v3 * 32767)
    v3 += v3
    switch "\n\t" {
    case "foo bar":
        v11 := GenS1{F0: 220, F1: 42}
        _ = v11
    default:
        if false {
            var v12 uint64 = 1
            v13 := GenS2{F0: 4294967295, F1: (718 < 2147483647), F2: 0}
            _ = v12
            _ = v13
        } else {
            for fi14 := 0; fi14 < 10; fi14++ {
                fmt.Println(v3, v4)
                _ = fi14
            }
            for fi15 := 0; fi15 < 6; fi15++ {
                fmt.Println(v3, v1)
                _ = fi15
            }
        }
    }
    if (0.1 >= v1) {
        v16 := GenS1{F0: -932, F1: 96}
        _ = v16
    }
    v6[("hello" + ((("foo bar" + ("a" + "a")) + "a") + ""))] = "a"
    v17 := GenS3{F0: (4294967295 >= 18446744073709551615), F1: 2, F2: -820}
    fmt.Println(v10, v3, v4)
    fmt.Println(v10, v3, v1)
    var v18 uint8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
