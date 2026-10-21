package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 int; F2 uint64 }
type GenS3 struct { F0 bool; F1 int32 }
func genFn0() int8 {
    var v7 uint64 = 42
    fmt.Println(v7)
    var v8 int8 = -1
    v9 := []int{1}
    v10 := GenS1{F0: 1}
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return v8
}
func genFn1(p0 int16, p1 int) int64 {
    p0 += p0
    p0 = (p0 ^ 0)
    _ = p0
    _ = p1
    return 627
}

func main() {
    v1 := []uint32{}
    var v2 int = 32768
    v3 := make(map[int]int)
    var v4 int16 = 1000
    v5 := make(map[string]uint32)
    v6 := make(map[int]uint32)
    for fi11 := 0; fi11 < 14; fi11++ {
        var v12 int64 = 9223372036854775807
        if (4294967295 >= 42) {
            v13 := GenS2{F0: "foo bar", F1: fi11, F2: 18446744073709551615}
            switch "hello" {
            case "":
                fmt.Println(fi11)
            }
            var v14 uint = 65535
            _ = v13
            _ = v14
        }
        _ = fi11
        _ = v12
    }
    for fi15 := 0; fi15 < 1; fi15++ {
        v16 := GenS2{F0: "", F1: 0, F2: 1}
        v17 := GenS1{F0: 65535}
        delete(v3, 42)
        _ = fi15
        _ = v16
        _ = v17
    }
    v18 := GenS3{F0: (false && (true && true)), F1: 42}
    switch v2 {
    case 4294967295:
        fmt.Println(v2, v4)
    case 1:
        v19 := GenS2{F0: "a", F1: v2, F2: 2}
        _ = v19
    case -383:
        v2 = (v2 ^ (v2 >> ((678) & 63)))
    }
    switch 1 {
    case -226:
        var v20 float32 = (0.0 + 1.5)
        _ = v20
    case 2147483647:
        for fi21 := 0; fi21 < 18; fi21++ {
            fmt.Println(v2, fi21)
            var v22 bool = (true && (false && (4294967295 > v2)))
            _ = fi21
            _ = v22
        }
    }
    _ = len(v3)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v23 bool = false
    for fi24 := 0; fi24 < 18; fi24++ {
        if v23 {
            v23 = (42 < 218)
        }
        _ = fi24
    }
    var v25 uint16 = 302
    var v26 int64 = -9223372036854775808
    fmt.Println(v4)
    _ = len(v3)
    fmt.Println(v26, v23, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
