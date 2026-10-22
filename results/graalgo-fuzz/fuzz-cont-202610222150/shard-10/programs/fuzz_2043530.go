package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 int8; F2 uint16 }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 bool; F1 string }
func genFn0(p0 int16) uint16 {
    switch "foo bar" {
    case "\n\t":
        if true {
            var v5 uint = 18446744073709551615
            _ = v5
        }
    case "":
        v6 := []int{2147483647, 42, -1, 2147483647}
        _ = v6
    case "a":
        if true {
            if true {
                v7 := []int{969, 2147483647, 2147483647}
                p0 += -2
                _ = v7
            }
            var v8 uint8 = 76
            _ = v8
        }
    }
    v9 := []int{127, 9223372036854775807, 42, 42}
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    v10 := GenS1{F0: -2}
    if len(v9) > 0 {
        v9[0] = 2147483647
    }
    _ = p0
    _ = v10
    _ = v9
    return 42
}

func main() {
    var v1 bool = false
    v2 := []uint32{528, 65535, 2147483647, 920}
    v3 := make(map[string]int)
    v4 := []int{}
    v11 := GenS2{F0: 775, F1: -107, F2: 32767}
    if (v1 || ((!(1.5 < 0.0)) && (v1 || v1))) {
        for fi12 := 0; fi12 < 1; fi12++ {
            for fi13 := 0; fi13 < 14; fi13++ {
                v14 := GenS3{F0: false, F1: 0}
                if len(v4) > 0 {
                    v4[0] = -2147483648
                }
                _ = fi13
                _ = v14
            }
            v1 = (1 > 510)
            _ = fi12
        }
        switch -2147483648 {
        case -9223372036854775808:
            delete(v3, ("a" + "hello"))
        }
        fmt.Println(v1)
    }
    switch "a" {
    case "\n\t":
        fmt.Println(v1)
    case "hello":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    default:
        v4 = append(v4, 127)
    }
    v15 := GenS4{F0: (170 < 1000), F1: ("\n\t" + "hello")}
    fmt.Println(v1)
    if (("\n\t" != "hello") || true) {
        v1 = (false && (-406 < 0))
    }
    var v16 uint = 947
    var v17 uint16 = 42
    if ((32 < -67) || (42 == 1000)) {
        var v18 uint16 = v17
        var v19 int = 9223372036854775807
        _ = v18
        _ = v19
    } else {
        if (-128 == 100) {
            v1 = ((0 > 42) || (!((!true) && (-1 < 532))))
            v17 -= v17
        }
    }
    v1 = (0 <= 0)
    var v20 int16 = 1000
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
