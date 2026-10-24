package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 int64; F2 string }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 string; F1 bool }
func genFn0(p0 int16, p1 int32) int8 {
    switch 153 {
    case 1:
        p1 = 0
    case 32768:
        if ((-(p1)) <= p1) {
            switch ("a" + "") {
            case "\n\t":
                var v7 uint64 = 797
                _ = v7
            case "hello":
                var v8 uint32 = 463
                _ = v8
            case "a":
                p1 = (p1 | -2147483648)
            }
            if (0 != 520) {
                for fi9 := 0; fi9 < 15; fi9++ {
                    v10 := []int{fi9, (fi9 << ((262) & 63))}
                    _ = fi9
                    _ = v10
                }
                v11 := GenS1{F0: 0, F1: true, F2: (p1 | (p1 << ((18446744073709551615) & 63)))}
                p1 = (p1 - p1)
                _ = v11
            } else {
                v12 := []int{32768}
                fmt.Println(p1)
                v13 := GenS1{F0: 127, F1: (0 < 255), F2: (p1 & p1)}
                _ = v12
                _ = v13
            }
        } else {
            v14 := GenS1{F0: 100, F1: (1e10 >= float64(3.14159)), F2: 2147483647}
            switch (("foo bar" + "a") + "hello") {
            case "\n\t":
                if true {
                    var v15 uint8 = 177
                    fmt.Println(p0, v15, v15)
                    v15 -= (^(v15))
                    _ = v15
                } else {
                    v16 := GenS2{F0: false}
                    _ = v16
                }
            }
            _ = v14
        }
    default:
        v17 := []int{}
        _ = v17
    }
    v18 := GenS2{F0: (454 <= 0)}
    _ = p0
    _ = p1
    _ = v18
    return 100
}

func main() {
    v1 := make(map[string]int)
    v2 := []int{-566}
    v3 := []uint32{938, 2147483648}
    var v4 int64 = -9223372036854775808
    var v5 uint8 = 1
    v6 := []uint32{136, 42, 4294967295}
    v19 := GenS2{F0: false}
    switch 42 {
    case -9223372036854775808:
        if ((true || (2 > 0)) && false) {
            v2 = append(v2, 9223372036854775807)
        }
    }
    switch "" {
    case "a":
        v20 := GenS3{F0: 75, F1: (v4 + -9223372036854775808), F2: ""}
        _ = v20
    }
    var v21 float32 = 1e10
    fmt.Println(v4, v4, v4)
    var v22 bool = (false || false)
    v23 := GenS4{F0: (v5 % ((128) | 1))}
    v24 := GenS5{F0: "\n\t", F1: (v22 && false)}
    var v25 int64 = (-(v4))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
