package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 string; F1 bool; F2 int16 }
type GenS3 struct { F0 int; F1 uint; F2 uint64 }
func genFn0(p0 int, p1 int, p2 uint16) uint32 {
    v5 := GenS1{F0: 2147483647, F1: 1}
    switch "\n\t" {
    case "hello":
        v6 := GenS1{F0: 130, F1: 32767}
        _ = v6
    }
    v7 := []int{(p0 / ((p0) | 1))}
    p1 = (p1 & p1)
    var v8 int64 = 9223372036854775807
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 bool = false
    v2 := make(map[int]uint32)
    v3 := []int{9223372036854775807}
    v4 := make(map[string]string)
    v1 = (323 < 0)
    v9 := GenS2{F0: ("a" + "\n\t"), F1: (!(!v1)), F2: -428}
    var v10 int = 32768
    fmt.Println(v10, v10)
    v1 = (554 != 4294967295)
    v10 = (-(v10))
    var v11 int = 32768
    for fi12 := 0; fi12 < 6; fi12++ {
        fmt.Println(v10, fi12, v11)
        v1 = false
        fmt.Println(fi12, v10)
        _ = fi12
    }
    if (("foo bar" <= "foo bar") && v1) {
        v11 = v11
        fmt.Println(v10, v11)
        var v13 int8 = 0
        _ = v13
    } else {
        var v14 string = ("hello" + ("a" + ""))
        v15 := GenS3{F0: v11, F1: 4294967295, F2: 18446744073709551615}
        var v16 float64 = 1e10
        _ = v14
        _ = v15
        _ = v16
    }
    fmt.Println(v1, v1, v11)
    for fi17 := 0; fi17 < 20; fi17++ {
        v10 = v11
        v1 = ((18446744073709551615 < 508) || v1)
        _ = fi17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
