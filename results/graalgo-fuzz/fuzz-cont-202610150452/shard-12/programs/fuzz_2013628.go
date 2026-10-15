package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 int64; F1 uint16 }
func genFn0(p0 int32, p1 string) int16 {
    var v6 int = 0
    var v7 int64 = 26
    v8 := GenS1{F0: 42, F1: -676}
    v9 := []int{(v6 << ((4294967295) & 63)), 32767}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 286
}
func genFn1(p0 int16, p1 uint, p2 uint64) string {
    for fi10 := 0; fi10 < 2; fi10++ {
        v11 := GenS2{F0: 0, F1: 32767}
        var v12 bool = (42 > 329)
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 int8 = 127
    if ((p0 << ((p1) & 63)) < 41) {
        var v14 int32 = 1
        _ = v14
    } else {
        var v15 uint16 = 921
        v16 := []int{}
        if len(v16) > 0 {
            v16[0] = 1
        }
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 6; fi17++ {
        var v18 int16 = p0
        p1 -= p1
        _ = fi17
        _ = v18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return (("foo bar" + ("foo bar" + ("foo bar" + "a"))) + "\n\t")
}

func main() {
    var v1 uint64 = 2
    var v2 int32 = 0
    var v3 string = "foo bar"
    var v4 uint32 = 236
    var v5 int = 0
    fmt.Println(v2, v5, v3)
    fmt.Println(v4, v5)
    v19 := GenS1{F0: (v4 - v4), F1: -9223372036854775808}
    for fi20 := 0; fi20 < 1; fi20++ {
        var v21 bool = true
        _ = fi20
        _ = v21
    }
    v22 := GenS2{F0: -568, F1: 42}
    v3 = "\n\t"
    v4 += (v4 << ((uint(v4)) & 63))
    var v23 uint = 65535
    v23 = (v23 >> ((uint(v1)) & 63))
    for fi24 := 0; fi24 < 10; fi24++ {
        v25 := GenS1{F0: 170, F1: fi24}
        var v26 uint16 = 65535
        _ = fi24
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 1; fi27++ {
        v28 := []int{v5, (v5 ^ (v5 - v5))}
        _ = fi27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
