package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 int64; F2 int }
func genFn0(p0 uint32, p1 bool) uint64 {
    v5 := []int{}
    switch 2147483647 {
    case 42:
        if len(v5) > 0 {
            v5[0] = 989
        }
    default:
        p1 = (0.1 != ((float64(0.1) * -2.25) + 0.1))
    }
    fmt.Println(p1)
    for fi6 := 0; fi6 < 16; fi6++ {
        var v7 string = ("hello" + "a")
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v5
    return 71
}
func genFn1(p0 int8, p1 int64, p2 int16) int8 {
    switch 32767 {
    case 199:
        p0 = (p0 * (p0 << ((633) & 63)))
    case 32767:
        var v8 uint = 1
        _ = v8
    case 128:
        p0 = (^(p0))
    default:
        v9 := []int{-937, 4294967295, -1, -801}
        _ = v9
    }
    p1 += p1
    for fi10 := 0; fi10 < 14; fi10++ {
        if false {
            fmt.Println(p2)
            v11 := []int{613, fi10, fi10, 128}
            _ = v11
        } else {
            for fi12 := 0; fi12 < 4; fi12++ {
                v13 := []int{fi12}
                _ = fi12
                _ = v13
            }
            v14 := GenS1{F0: true, F1: -654}
            _ = v14
        }
        _ = fi10
    }
    var v15 uint16 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return -82
}

func main() {
    v1 := make(map[int]string)
    var v2 uint = 0
    var v3 uint8 = 127
    var v4 uint16 = 1
    switch ("hello" + "") {
    case "":
        if (true || true) {
            for fi16 := 0; fi16 < 12; fi16++ {
                var v17 int64 = -1
                var v18 uint32 = 363
                _ = fi16
                _ = v17
                _ = v18
            }
            var v19 int8 = 100
            v20 := GenS1{F0: true, F1: -1}
            _ = v19
            _ = v20
        }
    }
    if (true && false) {
        v21 := GenS2{F0: 2147483647}
        v22 := v1[-2147483648]
        if (!((v4 * v4) <= v4)) {
            v23 := []int{}
            _ = v23
        }
        _ = v21
        _ = v22
    } else {
        for fi24 := 0; fi24 < 5; fi24++ {
            v25 := GenS2{F0: 82}
            _ = fi24
            _ = v25
        }
        fmt.Println(v2)
    }
    var v26 string = "foo bar"
    for fi27 := 0; fi27 < 5; fi27++ {
        if (9223372036854775807 > -55) {
            _ = len(v1)
        }
        _ = fi27
    }
    if false {
        fmt.Println(v4)
        var v28 int32 = 42
        _ = v28
    }
    fmt.Println(v26)
    v1[9223372036854775807] = "a"
    v29 := []int{1, 32768, 1}
    var v30 int16 = 442
    if (!false) {
        fmt.Println(v2)
        v30 = (v30 - v30)
        v31 := GenS3{F0: -1, F1: -300, F2: 1}
        _ = v31
    } else {
        var v32 int64 = 156
        if true {
            v33 := GenS2{F0: -540}
            fmt.Println(v2, v3, v26)
            v4 = (v4 % (((v4 | (v4 / ((v4) | 1)))) | 1))
            _ = v33
        }
        _ = v32
    }
    v3 = 21
    v4 = (^(v4))
    var v34 bool = (!(3 < 42))
    switch v26 {
    case "hello":
        var v35 string = ((v26 + v26) + v26)
        _ = v35
    case "\n\t":
        var v36 string = (("hello" + v26) + "hello")
        _ = v36
    default:
        v37 := v1[-182]
        _ = v37
    }
    for fi38 := 0; fi38 < 1; fi38++ {
        v34 = (!v34)
        fmt.Println(v26, v30)
        if len(v29) > 0 {
            v29[0] = fi38
        }
        _ = fi38
    }
    switch (v26 + v26) {
    case "foo bar":
        fmt.Println(v30)
    case "hello":
        for fi39 := 0; fi39 < 19; fi39++ {
            var v40 float64 = 0.0
            var v41 bool = (!v34)
            _ = fi39
            _ = v40
            _ = v41
        }
    case "a":
        var v42 uint = (v2 | v2)
        _ = v42
    default:
        var v43 uint = (v2 << (((v2 * v2)) & 63))
        _ = v43
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v4
}
