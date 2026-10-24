package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0() int16 {
    var v4 float64 = 0.1
    v4 = (v4 + v4)
    v4 -= (((v4 / v4) + v4) - v4)
    _ = v4
    return 1
}
func genFn1(p0 uint64, p1 uint8) int8 {
    switch 42 {
    case -9223372036854775808:
        p1 = (p1 >> ((65535) & 63))
    case 2147483647:
        p1 += 1
    case -2147483648:
        switch "\n\t" {
        case "foo bar":
            v5 := GenS1{F0: genFn0()}
            _ = v5
        case "hello":
            v6 := []int{-149}
            _ = v6
        case "":
            var v7 uint32 = 0
            _ = v7
        default:
            fmt.Println(p0)
        }
    }
    switch -560 {
    case 128:
        p1 = p1
    case -952:
        var v8 int = 4294967295
        _ = v8
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    return 2
}

func main() {
    v1 := []string{("foo bar" + ""), "hello"}
    var v2 uint32 = 2
    var v3 int32 = -698
    v2 += v2
    fmt.Println(v2, v2)
    fmt.Println(v2, v3)
    v9 := GenS1{F0: -1}
    if ((((false && false) && (v3 >= v3)) || (!true)) || true) {
        for fi10 := 0; fi10 < 9; fi10++ {
            if (true || false) {
                var v11 float32 = float32((0.0 / 0.1))
                var v12 float64 = 1e10
                _ = v11
                _ = v12
            } else {
                v3 = (v3 + (v3 - 1))
            }
            v3 = (v3 + v3)
            fmt.Println(v2)
            _ = fi10
        }
    } else {
        var v13 int64 = -9223372036854775808
        v14 := GenS1{F0: 32767}
        _ = v13
        _ = v14
    }
    var v15 int64 = 0
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v9
}
