package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 string }
type GenS2 struct { F0 int8; F1 uint64; F2 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint16; F1 uint16 }
func genFn0(p0 uint32, p1 uint, p2 int16) int16 {
    if ((!true) || (!true)) {
        v6 := []int{127, -806, -2147483648, -861}
        _ = v6
    }
    v7 := []int{-9223372036854775808, 4294967295, -9223372036854775808}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (p2 + (p2 % ((1000) | 1)))
}
func genFn1(p0 bool) uint32 {
    var v8 string = "hello"
    fmt.Println(v8, v8)
    var v9 int8 = -13
    if p0 {
        var v10 bool = (((0.1 + 0.0) - -2.25) > 1e10)
        _ = v10
    } else {
        for fi11 := 0; fi11 < 17; fi11++ {
            switch fi11 {
            case 127:
                v8 = ((v8 + v8) + "\n\t")
            case 32768:
                if (!(23 == 931)) {
                    v12 := GenS1{F0: fi11, F1: true, F2: "hello"}
                    _ = v12
                } else {
                    p0 = p0
                    v9 = v9
                }
            case -9223372036854775808:
                v13 := GenS2{F0: -16, F1: 0, F2: false}
                _ = v13
            }
            v14 := GenS2{F0: v9, F1: 595, F2: p0}
            _ = fi11
            _ = v14
        }
        switch v8 {
        case "\n\t":
            v15 := GenS3{F0: genFn0(0, 65535, int16(v9))}
            _ = v15
        case "":
            if (2 == 1) {
                var v16 int = 1
                _ = v16
            }
        case "foo bar":
            v8 = (v8 + "foo bar")
        }
        var v17 int8 = -128
        _ = v17
    }
    _ = p0
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 uint32 = 0
    var v2 float32 = 0.0
    v3 := make(map[int]int)
    v4 := make(map[int]string)
    v5 := make(map[int]string)
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v1)
    case "":
        if false {
            if false {
                var v18 int32 = -2147483648
                _ = v18
            } else {
                v4[1] = (("foo bar" + "\n\t") + "\n\t")
                v19 := []int{410, 32767, -147}
                _ = v19
            }
        }
    case "foo bar":
        var v20 int = 32768
        _ = v20
    }
    v21 := GenS1{F0: 9223372036854775807, F1: true, F2: "\n\t"}
    fmt.Println(v2)
    v4[-2147483648] = "a"
    v22 := GenS4{F0: (v1 << ((42) & 63))}
    v23 := GenS5{F0: 2, F1: 1}
    var v24 bool = false
    for fi25 := 0; fi25 < 8; fi25++ {
        var v26 uint32 = genFn1(v24)
        _ = fi25
        _ = v26
    }
    var v27 int = -9223372036854775808
    v24 = (!(!v24))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
