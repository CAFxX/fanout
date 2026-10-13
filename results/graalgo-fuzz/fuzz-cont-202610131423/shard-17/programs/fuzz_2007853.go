package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 uint; F1 uint32 }
func genFn0(p0 uint32, p1 int, p2 int32) uint8 {
    switch "foo bar" {
    case "a":
        var v6 uint16 = 32767
        _ = v6
    }
    p1 += p1
    _ = p0
    _ = p1
    _ = p2
    return 255
}

func main() {
    var v1 int32 = 2147483647
    var v2 uint32 = 1
    var v3 uint32 = v2
    var v4 uint = 42
    var v5 uint8 = 255
    v4 = (^(v4))
    v4 = v4
    for fi7 := 0; fi7 < 4; fi7++ {
        v4 = (v4 % ((65535) | 1))
        v8 := GenS1{F0: ("foo bar" + ("foo bar" + "hello")), F1: 32767}
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: "foo bar", F1: 493}
    v4 = (v4 >> ((v4) & 63))
    var v10 int16 = 1
    if (v1 < (v1 / (((v1 * (v1 % (((^(v1))) | 1)))) | 1))) {
        v2 = (v2 >> ((v4) & 63))
    } else {
        v5 = (^(v5))
        v11 := []int{32767, -963}
        _ = v11
    }
    fmt.Println(v1)
    if false {
        v12 := []int{}
        _ = v12
    }
    if true {
        v2 = (^(v2))
        v13 := GenS2{F0: v4, F1: (v3 * v3)}
        var v14 int8 = 0
        _ = v13
        _ = v14
    } else {
        v1 = v1
    }
    v10 -= -817
    v3 += v3
    if false {
        v15 := []int{0}
        _ = v15
    } else {
        v10 -= v10
        fmt.Println(v4, v1, v5)
        v16 := GenS1{F0: (("\n\t" + "a") + "\n\t"), F1: 255}
        _ = v16
    }
    if ("\n\t" <= ("a" + "hello")) {
        fmt.Println(v1)
    }
    v17 := []int{9223372036854775807, 32768}
    if len(v17) > 0 {
        v17[0] = -2147483648
    }
    var v18 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
