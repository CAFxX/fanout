package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 int64; F1 uint32; F2 uint64 }
type GenS3 struct { F0 int16; F1 uint8; F2 bool }
func genFn0(p0 int) uint16 {
    p0 -= p0
    switch "a" {
    case "foo bar":
        if (false || (0 == 379)) {
            for fi5 := 0; fi5 < 7; fi5++ {
                var v6 int32 = -1
                _ = fi5
                _ = v6
            }
        }
    default:
        switch (p0 >> ((15) & 63)) {
        case 572:
            if false {
                p0 = (p0 ^ -9223372036854775808)
                for fi7 := 0; fi7 < 16; fi7++ {
                    v8 := []int{-429, (fi7 + (p0 * 1)), p0, -600}
                    var v9 float64 = 1.5
                    _ = fi7
                    _ = v8
                    _ = v9
                }
                p0 = (p0 - (p0 / (((p0 ^ p0)) | 1)))
            }
        }
    }
    if true {
        v10 := GenS1{F0: -810, F1: -993}
        v11 := []int{}
        fmt.Println(p0)
        _ = v10
        _ = v11
    } else {
        v12 := GenS2{F0: 0, F1: 827, F2: 42}
        p0 += (p0 | p0)
        v13 := []int{p0, p0, (p0 * (p0 | p0)), p0}
        _ = v12
        _ = v13
    }
    fmt.Println(p0)
    _ = p0
    return 65535
}
func genFn1(p0 int8, p1 uint16, p2 int64) uint8 {
    v14 := []int{127, -9223372036854775808, -197, int(p1)}
    p0 += (p0 + (p0 - (p0 - p0)))
    if len(v14) > 0 {
        v14[0] = 1
    }
    var v15 uint16 = (p1 % ((p1) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    return 2
}

func main() {
    var v1 string = "foo bar"
    v2 := []string{v1, "foo bar"}
    var v3 string = ("" + "hello")
    var v4 uint = 42
    v16 := GenS2{F0: 453, F1: 1, F2: 2}
    fmt.Println(v1)
    _ = len(v2)
    var v17 uint32 = 4294967295
    var v18 int16 = -521
    var v19 string = "hello"
    if true {
        var v20 string = "hello"
        _ = v20
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v21 int64 = -21
    var v22 int64 = (v21 & -701)
    v4 -= (v4 << ((v4) & 63))
    switch "hello" {
    case "":
        v19 = v19
    case "foo bar":
        if false {
            v3 = v1
            switch v1 {
            case "a":
                if true {
                    if (((176 <= genFn1(0, 255, v22)) || (!(v17 < v17))) && (!(v18 >= (v18 - 32767)))) {
                        var v23 string = v19
                        var v24 float64 = 3.14159
                        v25 := GenS3{F0: v18, F1: 127, F2: false}
                        _ = v23
                        _ = v24
                        _ = v25
                    }
                    v1 = (v19 + (v3 + (((v3 + ((v1 + "a") + "\n\t")) + v3) + ("foo bar" + "a"))))
                    switch -958 {
                    case 127:
                        v3 = "a"
                    default:
                        var v26 int16 = v18
                        _ = v26
                    }
                } else {
                    switch v1 {
                    case "hello":
                        v3 = ((((v3 + v19) + (v1 + v3)) + v19) + v3)
                    default:
                        v22 -= (^(v22))
                    }
                    var v27 uint64 = 2
                    var v28 uint32 = 2
                    _ = v27
                    _ = v28
                }
            case "\n\t":
                var v29 bool = false
                _ = v29
            case "hello":
                _ = len(v2)
            default:
                var v30 int8 = -104
                _ = v30
            }
            v3 = (((v1 + v19) + v19) + v19)
        } else {
            for fi31 := 0; fi31 < 3; fi31++ {
                var v32 uint32 = 2147483648
                var v33 uint64 = 1
                _ = fi31
                _ = v32
                _ = v33
            }
        }
    }
    fmt.Println(v18)
    fmt.Println(v4, v22, v1)
    fmt.Println(v18, v22, v22)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
