package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 int16 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 int; F1 uint64; F2 int16 }
type GenS4 struct { F0 string; F1 uint64; F2 string }
type GenS5 struct { F0 uint32; F1 string }
type GenS6 struct { F0 int; F1 uint64; F2 bool }
func genFn0(p0 uint, p1 uint) int16 {
    p0 = p0
    for fi7 := 0; fi7 < 14; fi7++ {
        switch (-(fi7)) {
        case 761:
            var v8 string = ("foo bar" + ("\n\t" + "foo bar"))
            _ = v8
        case 161:
            var v9 uint16 = 42
            _ = v9
        }
        if (!(true && false)) {
            for fi10 := 0; fi10 < 11; fi10++ {
                v11 := GenS1{F0: 1, F1: 0, F2: 1000}
                _ = fi10
                _ = v11
            }
        }
        _ = fi7
    }
    var v12 uint16 = 722
    if true {
        var v13 uint32 = 2147483647
        p0 -= p0
        fmt.Println(v12)
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v12
    return 42
}
func genFn1(p0 int, p1 int) uint32 {
    p0 = (p1 & (^(p1)))
    switch 373 {
    case 0:
        fmt.Println(p1)
    }
    v14 := GenS1{F0: 1, F1: 4294967295, F2: 1}
    switch "" {
    case "hello":
        v15 := GenS1{F0: 42, F1: 104, F2: genFn0(18446744073709551615, 2)}
        _ = v15
    case "a":
        v16 := []int{-2147483648, (p0 | (p0 * (p0 << ((65535) & 63)))), 0, (^(p1))}
        _ = v16
    case "\n\t":
        v17 := GenS2{F0: -2147483648, F1: 42}
        _ = v17
    default:
        v18 := []int{-9223372036854775808}
        _ = v18
    }
    _ = p0
    _ = p1
    _ = v14
    return 65535
}

func main() {
    var v1 bool = (!false)
    var v2 bool = true
    var v3 int32 = 0
    v4 := make(map[int]int)
    var v5 int64 = 0
    var v6 int64 = v5
    _ = len(v4)
    v3 = v3
    var v19 int32 = (v3 * (v3 & (-(v3))))
    v3 = v3
    v20 := GenS1{F0: 42, F1: 249, F2: 0}
    for fi21 := 0; fi21 < 14; fi21++ {
        switch -2147483648 {
        case -447:
            v22 := GenS2{F0: -1, F1: 255}
            _ = v22
        case 1:
            v23 := GenS3{F0: (fi21 >> ((1) & 63)), F1: 18446744073709551615, F2: 42}
            _ = v23
        case 128:
            v2 = (127 == 128)
        }
        for fi24 := 0; fi24 < 14; fi24++ {
            fmt.Println(v1, v6)
            _ = fi24
        }
        if (2147483647 <= (fi21 << ((0) & 63))) {
            fmt.Println(v6, v5)
            v25 := GenS4{F0: "a", F1: 2, F2: "\n\t"}
            var v26 uint8 = 0
            _ = v25
            _ = v26
        }
        _ = fi21
    }
    v4[-9223372036854775808] = 42
    for fi27 := 0; fi27 < 8; fi27++ {
        v28 := GenS2{F0: (v3 << ((675) & 63)), F1: 2}
        v29 := GenS5{F0: 21, F1: "foo bar"}
        _ = fi27
        _ = v28
        _ = v29
    }
    v30 := GenS6{F0: 2147483647, F1: 2, F2: false}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
