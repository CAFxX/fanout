package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint16) int8 {
    for fi4 := 0; fi4 < 13; fi4++ {
        var v5 uint = 18446744073709551615
        v5 += v5
        _ = fi4
        _ = v5
    }
    switch 32767 {
    case 439:
        v6 := GenS1{F0: 255}
        _ = v6
    default:
        v7 := []int{}
        _ = v7
    }
    p0 = 65535
    var v8 uint32 = 551
    _ = p0
    _ = v8
    return 42
}
func genFn1(p0 string) uint16 {
    for fi9 := 0; fi9 < 16; fi9++ {
        v10 := []int{}
        p0 = (p0 + (p0 + ("" + p0)))
        _ = fi9
        _ = v10
    }
    var v11 float64 = 0.0
    p0 = (p0 + (p0 + (("foo bar" + p0) + ("" + p0))))
    v12 := []int{28, 32768}
    _ = p0
    _ = v11
    _ = v12
    return 2
}

func main() {
    var v1 int64 = 464
    v2 := []int{}
    var v3 uint = 433
    if false {
        v13 := GenS2{F0: 2}
        _ = v13
    }
    switch 841 {
    case -2147483648:
        switch 2147483647 {
        case 2147483647:
            fmt.Println(v3, v1)
        }
    default:
        var v14 int8 = int8(v3)
        _ = v14
    }
    switch ("foo bar" + ("a" + "foo bar")) {
    case "foo bar":
        if (((1000 <= 32767) && true) || (-2.25 <= 0.0)) {
            var v15 float32 = 1e10
            var v16 uint64 = 4294967295
            fmt.Println(v15, v15)
            _ = v15
            _ = v16
        }
    case "a":
        v3 = 881
    default:
        var v17 bool = (("hello" == "foo bar") && (0 < 855))
        _ = v17
    }
    fmt.Println(v1)
    if (-1 <= -313) {
        if len(v2) > 0 {
            v2[0] = 9223372036854775807
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v1)
    v1 += v1
    v1 = (v1 ^ -717)
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
}
