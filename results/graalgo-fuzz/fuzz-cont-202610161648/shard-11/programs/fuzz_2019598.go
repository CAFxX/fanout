package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint64 }
type GenS2 struct { F0 uint; F1 int8; F2 bool }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 string; F1 uint16 }
func genFn0(p0 int, p1 bool) uint16 {
    for fi4 := 0; fi4 < 10; fi4++ {
        var v5 uint16 = 0
        switch (("" + "") + "") {
        case "foo bar":
            var v6 int = 127
            _ = v6
        case "\n\t":
            fmt.Println(v5, p1)
        }
        v7 := []int{(p0 * 2147483647)}
        _ = fi4
        _ = v5
        _ = v7
    }
    p1 = false
    p0 = (-(p0))
    v8 := []int{}
    p0 -= 128
    _ = p0
    _ = p1
    _ = v8
    return 121
}
func genFn1(p0 int) int8 {
    v9 := []int{(p0 << ((282) & 63)), 127}
    p0 = (^(p0))
    _ = len(v9)
    var v10 uint8 = 221
    _ = p0
    _ = v10
    _ = v9
    return 63
}

func main() {
    v1 := make(map[int]bool)
    var v2 int64 = 0
    var v3 int32 = 1
    for fi11 := 0; fi11 < 11; fi11++ {
        v12 := GenS1{F0: 2147483647, F1: 0, F2: 908}
        v13 := []int{(fi11 * -2147483648), fi11, (fi11 % ((295) | 1)), (fi11 * (-(fi11)))}
        v14 := GenS1{F0: 0, F1: 421, F2: 18446744073709551615}
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    if true {
        switch 460 {
        case 127:
            v15 := GenS2{F0: 763, F1: -42, F2: false}
            _ = v15
        default:
            v2 += v2
        }
        var v16 uint8 = 42
        v17 := GenS2{F0: 4294967295, F1: 100, F2: ("\n\t" == ("foo bar" + "\n\t"))}
        _ = v16
        _ = v17
    } else {
        v1[52] = true
    }
    v3 += v3
    v18 := []int{}
    v19 := GenS3{F0: 42}
    fmt.Println(v2)
    var v20 int = 1
    v3 = v3
    v21 := GenS4{F0: "\n\t", F1: 511}
    v20 += v20
    v2 += v2
    var v22 uint = 306
    var v23 int64 = 0
    for fi24 := 0; fi24 < 8; fi24++ {
        var v25 bool = false
        _ = fi24
        _ = v25
    }
    var v26 string = "a"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v26
    _ = v3
}
