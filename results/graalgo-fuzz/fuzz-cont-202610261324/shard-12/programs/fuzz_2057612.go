package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 int8; F2 uint16 }
func genFn0(p0 uint32, p1 string, p2 uint16) string {
    switch "a" {
    case "foo bar":
        var v4 string = p1
        _ = v4
    }
    for fi5 := 0; fi5 < 5; fi5++ {
        p2 = (p2 * (p2 << ((uint(p2)) & 63)))
        _ = fi5
    }
    v6 := GenS1{F0: 18446744073709551615}
    switch 2147483647 {
    case 9223372036854775807:
        p2 = 65535
    case -9223372036854775808:
        v7 := GenS1{F0: 18446744073709551615}
        _ = v7
    case -924:
        p2 += p2
    }
    var v8 int64 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v8
    return "a"
}
func genFn1() int8 {
    var v9 uint = 255
    fmt.Println(v9)
    fmt.Println(v9)
    var v10 uint32 = 42
    fmt.Println(v9, v10)
    _ = v10
    _ = v9
    return 1
}

func main() {
    var v1 uint8 = 128
    var v2 uint = 1
    v3 := []int{364, 1, -216}
    for fi11 := 0; fi11 < 17; fi11++ {
        v1 += (v1 % (((v1 - v1)) | 1))
        _ = fi11
    }
    switch 632 {
    case -871:
        var v12 uint8 = (v1 / ((255) | 1))
        _ = v12
    case -676:
        if len(v3) > 0 {
            v3[0] = 128
        }
    default:
        _ = len(v3)
    }
    v13 := GenS2{F0: genFn1(), F1: genFn1(), F2: 65535}
    var v14 uint32 = 0
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v14 = v14
    v15 := GenS1{F0: 18446744073709551615}
    var v16 float32 = float32((1e10 - 1e10))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
