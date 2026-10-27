package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 bool; F2 uint }
type GenS3 struct { F0 int32; F1 int; F2 uint8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint8; F1 uint64 }
func genFn0() string {
    for fi4 := 0; fi4 < 10; fi4++ {
        fmt.Println(fi4)
        var v5 uint32 = 633
        v6 := GenS1{F0: 255}
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 int = 32767
    v7 = -1
    _ = v7
    return "a"
}
func genFn1(p0 uint64, p1 int32, p2 int8) int {
    if true {
        if (!(42 < p2)) {
            for fi8 := 0; fi8 < 19; fi8++ {
                switch "hello" {
                case "hello":
                    switch "hello" {
                    case "foo bar":
                        p1 = (p1 ^ p1)
                    case "\n\t":
                        v9 := GenS2{F0: 42, F1: false, F2: uint(p0)}
                        _ = v9
                    }
                case "a":
                    v10 := []int{fi8}
                    _ = v10
                case "\n\t":
                    fmt.Println(fi8)
                }
                for fi11 := 0; fi11 < 13; fi11++ {
                    if (true && false) {
                        switch (fi8 ^ -861) {
                        case 2147483647:
                            p2 -= (-(p2))
                            var v12 uint8 = 2
                            v12 += v12
                            _ = v12
                        case 660:
                            v13 := []int{(fi11 << ((65535) & 63)), -1, int(p0), (^(fi8))}
                            _ = v13
                        case 32768:
                            v14 := GenS3{F0: (-(p1)), F1: (fi11 - fi11), F2: 104}
                            _ = v14
                        default:
                            v15 := []int{}
                            _ = v15
                        }
                    }
                    fmt.Println(fi8, p2)
                    v16 := []int{fi11, (-(fi11)), fi11, 0}
                    if len(v16) >= 2 {
                        v16 = v16[1:len(v16)]
                    }
                    _ = fi11
                    _ = v16
                }
                _ = fi8
            }
        } else {
            v17 := []int{-2147483648, 127, 32768}
            var v18 float64 = ((((-2.25 - (((1e10 * 0.1) * (-2.25 / float64(1e10))) / 1e10)) / 3.14159) / 3.14159) / 1.0)
            if ("a" >= "hello") {
                v19 := GenS2{F0: 2, F1: ("\n\t" == ("hello" + "\n\t")), F2: 738}
                _ = v19
            } else {
                if (128 <= 0) {
                    v17 = append(v17, -79)
                    fmt.Println(v18, p1, v18)
                    if ((p0 + (p0 * 18446744073709551615)) != p0) {
                        v18 -= ((v18 + (1.5 / ((v18 * 0.0) + 1e10))) + 1e10)
                    }
                }
                v20 := GenS1{F0: 128}
                fmt.Println(p2, p1)
                _ = v20
            }
            _ = v17
            _ = v18
        }
    }
    var v21 uint16 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    return -1
}

func main() {
    var v1 float64 = (-2.25 - 1e10)
    v2 := []int{}
    var v3 bool = (770 == 2147483648)
    v22 := GenS3{F0: -307, F1: 78, F2: 29}
    var v23 float32 = (1e10 * float32(1e10))
    _ = len(v2)
    var v24 int16 = 1000
    fmt.Println(v24)
    v25 := GenS4{F0: 1}
    switch "hello" {
    case "foo bar":
        var v26 uint8 = 150
        _ = v26
    case "\n\t":
        fmt.Println(v1, v24)
    case "a":
        var v27 uint64 = 1
        _ = v27
    default:
        fmt.Println(v3)
    }
    if len(v2) > 0 {
        v2[0] = 9223372036854775807
    }
    v2 = append(v2, 127)
    if v3 {
        v28 := GenS5{F0: 153, F1: 69}
        _ = v28
    } else {
        var v29 int32 = 42
        var v30 uint = 0
        v3 = v3
        _ = v29
        _ = v30
    }
    switch ("a" + "\n\t") {
    case "hello":
        switch ("a" + "\n\t") {
        case "\n\t":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            if len(v2) > 0 {
                v2[0] = -711
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
