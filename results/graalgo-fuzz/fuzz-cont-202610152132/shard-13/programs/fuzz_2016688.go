package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 uint16 }
type GenS2 struct { F0 uint8; F1 int8; F2 uint32 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 int64) int64 {
    v4 := GenS1{F0: "a", F1: ("" + "hello"), F2: 65535}
    p0 += (p0 ^ p0)
    switch "a" {
    case "hello":
        v5 := []int{9223372036854775807}
        _ = v5
    default:
        p0 = p0
    }
    v6 := []int{32767, -9223372036854775808, 9223372036854775807, 1}
    _ = p0
    _ = v4
    _ = v6
    return 0
}
func genFn1() int {
    var v7 float64 = float64(3.14159)
    var v8 uint = 65535
    var v9 float32 = -2.25
    _ = v7
    _ = v8
    _ = v9
    return 2147483647
}

func main() {
    var v1 int8 = 0
    var v2 uint64 = 18446744073709551615
    v3 := make(map[uint32]string)
    v2 = (^(v2))
    for fi10 := 0; fi10 < 1; fi10++ {
        v1 += v1
        var v11 string = ""
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 1; fi12++ {
        for fi13 := 0; fi13 < 4; fi13++ {
            var v14 int64 = 1
            var v15 int8 = (v1 - 42)
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi12
    }
    v1 -= (-(v1))
    if true {
        if false {
            var v16 int64 = 1
            for fi17 := 0; fi17 < 10; fi17++ {
                var v18 uint8 = 2
                v19 := GenS1{F0: ("\n\t" + ("a" + "")), F1: ("a" + ("foo bar" + "foo bar")), F2: 229}
                _ = len(v3)
                _ = fi17
                _ = v18
                _ = v19
            }
            v20 := GenS2{F0: 42, F1: (v1 * v1), F2: 2147483648}
            _ = v16
            _ = v20
        } else {
            v1 = 0
        }
        _ = len(v3)
        v21 := []int{-1, -9223372036854775808, 32767, -2147483648}
        _ = v21
    } else {
        v22 := v3[0]
        fmt.Println(v2)
        v1 -= (^(v1))
        _ = v22
    }
    fmt.Println(v2)
    v1 += v1
    if (!true) {
        v23 := GenS3{F0: 2147483648}
        v24 := []int{-1, -9223372036854775808, -9223372036854775808, genFn1()}
        _ = v23
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v3
}
