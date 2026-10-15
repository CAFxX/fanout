package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 int32; F2 uint32 }
type GenS3 struct { F0 int; F1 uint16 }
type GenS4 struct { F0 int64; F1 bool; F2 int64 }
type GenS5 struct { F0 int64; F1 uint8 }
type GenS6 struct { F0 int64; F1 uint16 }
func genFn0(p0 int8) int16 {
    switch 127 {
    case 9223372036854775807:
        switch (("hello" + "foo bar") + "foo bar") {
        case "":
            var v6 string = "\n\t"
            _ = v6
        }
    case -1:
        fmt.Println(p0)
    }
    if (!((-(p0)) >= (p0 ^ p0))) {
        var v7 int = -2147483648
        _ = v7
    }
    var v8 float32 = float32(((((-2.25 * 1e10) / 3.14159) + 0.0) - (1e10 - 0.0)))
    if (true && ((("hello" + (("" + ((("foo bar" + ("" + ("foo bar" + "hello"))) + ("a" + "foo bar")) + (("hello" + ("hello" + "\n\t")) + "hello"))) + ("\n\t" + "hello"))) + "foo bar") <= "hello")) {
        v9 := []int{0}
        v10 := GenS1{F0: "hello"}
        if true {
            v11 := GenS2{F0: 42, F1: 1, F2: 399}
            _ = v11
        }
        _ = v10
        _ = v9
    } else {
        for fi12 := 0; fi12 < 17; fi12++ {
            for fi13 := 0; fi13 < 19; fi13++ {
                v14 := []int{-524, fi13, (^(fi13)), (^(fi12))}
                var v15 int32 = -595
                _ = fi13
                _ = v14
                _ = v15
            }
            _ = fi12
        }
        var v16 int8 = p0
        var v17 bool = false
        _ = v16
        _ = v17
    }
    switch "" {
    case "":
        v18 := []int{}
        _ = v18
    default:
        switch 32768 {
        case -1:
            var v19 bool = true
            _ = v19
        default:
            v20 := GenS3{F0: 32767, F1: 255}
            _ = v20
        }
    }
    _ = p0
    _ = v8
    return -32768
}
func genFn1(p0 int8) bool {
    v21 := []int{731, 662, 0, 1}
    switch "a" {
    case "hello":
        for fi22 := 0; fi22 < 18; fi22++ {
            var v23 uint = 517
            var v24 int64 = -9223372036854775808
            _ = fi22
            _ = v23
            _ = v24
        }
    default:
        var v25 uint64 = 42
        _ = v25
    }
    var v26 uint8 = 0
    switch 128 {
    case 9223372036854775807:
        switch "foo bar" {
        case "\n\t":
            p0 = p0
        case "foo bar":
            if false {
                fmt.Println(p0, v26)
                v27 := GenS1{F0: "hello"}
                _ = v27
            }
        }
    }
    var v28 string = ("foo bar" + "a")
    _ = p0
    _ = v21
    _ = v26
    _ = v28
    return true
}

func main() {
    var v1 int = 32768
    var v2 int8 = 1
    v3 := []uint32{818}
    var v4 string = "a"
    var v5 uint8 = 128
    if (true || false) {
        v29 := GenS1{F0: v4}
        for fi30 := 0; fi30 < 16; fi30++ {
            switch v1 {
            case 0:
                v1 += v1
            default:
                v2 += (v2 | v2)
            }
            v4 = "a"
            _ = fi30
        }
        _ = v29
    } else {
        var v31 uint8 = 84
        v32 := GenS4{F0: -718, F1: ((0.1 == (0.1 * 3.14159)) && ((true || false) || (v31 <= v5))), F2: -1}
        _ = v31
        _ = v32
    }
    var v33 int = (v1 ^ v1)
    var v34 uint8 = 173
    fmt.Println(v33, v5, v34)
    fmt.Println(v33, v5, v1)
    v1 = v1
    v35 := GenS2{F0: 880, F1: 0, F2: 513}
    switch v33 {
    case -2147483648:
        for fi36 := 0; fi36 < 16; fi36++ {
            if true {
                v33 = v1
                fmt.Println(v2, v2, v4)
                v34 = (v34 % ((127) | 1))
            } else {
                v37 := GenS1{F0: v4}
                _ = v37
            }
            _ = fi36
        }
    case -9223372036854775808:
        for fi38 := 0; fi38 < 4; fi38++ {
            for fi39 := 0; fi39 < 1; fi39++ {
                if true {
                    v34 = (v5 + 128)
                } else {
                    v33 = int(v2)
                }
                v40 := GenS5{F0: 1, F1: v34}
                v4 = v4
                _ = fi39
                _ = v40
            }
            var v41 int = fi38
            _ = fi38
            _ = v41
        }
    case 336:
        for fi42 := 0; fi42 < 20; fi42++ {
            var v43 int32 = -209
            _ = fi42
            _ = v43
        }
    }
    v44 := GenS6{F0: -1, F1: uint16(v5)}
    if (139 < 976) {
        var v45 uint64 = 728
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v44
    _ = v5
}
