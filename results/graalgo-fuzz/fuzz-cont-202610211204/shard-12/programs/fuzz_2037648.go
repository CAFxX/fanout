package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 uint64; F1 int; F2 uint }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int64; F1 int8 }
type GenS7 struct { F0 int8; F1 int16; F2 int }
func genFn0(p0 uint, p1 int32) uint64 {
    v5 := GenS1{F0: ((true && true) && false)}
    p0 -= p0
    p0 = (p0 + 4294967295)
    p0 += p0
    v6 := GenS2{F0: 0}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 2
}

func main() {
    var v1 int8 = 42
    v2 := make(map[uint32]int)
    var v3 float64 = (1e10 - 1.5)
    v4 := make(map[int]string)
    v7 := GenS2{F0: -1}
    switch "hello" {
    case "a":
        var v8 uint32 = 42
        _ = v8
    case "hello":
        fmt.Println(v3)
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        fmt.Println(v3, v1)
        switch (^(fi9)) {
        case 128:
            v1 = 95
        case 2147483647:
            if (976 == 1) {
                for fi10 := 0; fi10 < 16; fi10++ {
                    v11 := []int{fi10}
                    v1 = v1
                    _ = len(v2)
                    _ = fi10
                    _ = v11
                }
                var v12 uint64 = 18446744073709551615
                _ = v12
            } else {
                for fi13 := 0; fi13 < 14; fi13++ {
                    var v14 uint64 = 4294967295
                    var v15 uint32 = 0
                    v16 := GenS3{F0: v15, F1: -1}
                    v17 := v2[92]
                    var v18 uint = 408
                    v14 += (^(v14))
                    v19 := v2[(v15 << ((v18) & 63))]
                    _ = fi13
                    _ = v14
                    _ = v15
                    _ = v16
                    _ = v17
                    _ = v18
                    _ = v19
                }
                v3 = float64(3.14159)
            }
        case -261:
            switch "foo bar" {
            case "foo bar":
                v20 := GenS4{F0: 137, F1: (-(fi9)), F2: 0}
                _ = v20
            default:
                delete(v2, 42)
            }
        default:
            for fi21 := 0; fi21 < 4; fi21++ {
                var v22 string = ("\n\t" + "\n\t")
                if false {
                    var v23 uint8 = 41
                    _ = v23
                }
                _ = fi21
                _ = v22
            }
        }
        _ = fi9
    }
    var v24 bool = true
    if true {
        switch "hello" {
        case "foo bar":
            var v25 uint64 = 42
            _ = v25
        case "":
            fmt.Println(v24, v1, v3)
        }
        var v26 int64 = 9223372036854775807
        for fi27 := 0; fi27 < 13; fi27++ {
            v28 := []int{2147483647, 32768}
            _ = fi27
            _ = v28
        }
        _ = v26
    }
    switch 179 {
    case -1:
        for fi29 := 0; fi29 < 7; fi29++ {
            var v30 int64 = 0
            var v31 uint = 529
            _ = fi29
            _ = v30
            _ = v31
        }
    default:
        v4[128] = ("" + ("" + ("a" + (("a" + "a") + "\n\t"))))
    }
    v1 -= v1
    for fi32 := 0; fi32 < 15; fi32++ {
        if v24 {
            fmt.Println(v24)
        } else {
            v1 -= 1
            fmt.Println(v1)
            v33 := GenS5{F0: (v1 << ((18446744073709551615) & 63))}
            _ = v33
        }
        _ = fi32
    }
    var v34 uint32 = 2147483647
    v24 = (false && (237 > v34))
    if v24 {
        for fi35 := 0; fi35 < 20; fi35++ {
            fmt.Println(v3, v34)
            fmt.Println(fi35, v34, v3)
            _ = fi35
        }
        v36 := GenS6{F0: 491, F1: 100}
        v1 -= -128
        _ = v36
    }
    v37 := GenS3{F0: (^(v34)), F1: 42}
    fmt.Println(v34, v3, v1)
    if true {
        for fi38 := 0; fi38 < 13; fi38++ {
            var v39 string = ""
            v39 = "hello"
            fmt.Println(v34, v3)
            _ = fi38
            _ = v39
        }
        var v40 uint8 = 193
        _ = v40
    }
    var v41 uint64 = 2
    switch ("\n\t" + "") {
    case "a":
        for fi42 := 0; fi42 < 9; fi42++ {
            v3 -= float64(0.1)
            _ = fi42
        }
    case "hello":
        v1 += v1
    case "\n\t":
        v43 := GenS7{F0: (v1 << ((4294967295) & 63)), F1: 1, F2: 128}
        _ = v43
    default:
        if v24 {
            v41 += (v41 ^ 18446744073709551615)
            var v44 uint = 42
            _ = v44
        }
    }
    switch ("a" + "hello") {
    case "foo bar":
        v45 := GenS4{F0: v41, F1: 42, F2: 1}
        _ = v45
    case "a":
        v46 := []int{}
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v34
    _ = v37
    _ = v4
    _ = v41
    _ = v7
}
