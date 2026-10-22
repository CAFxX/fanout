package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 int; F2 int }
type GenS4 struct { F0 int32; F1 int32; F2 int16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint16; F1 uint16; F2 int64 }
type GenS7 struct { F0 uint32; F1 uint }
func genFn0() uint16 {
    v4 := []int{-1}
    if true {
        v5 := GenS1{F0: -982, F1: -1, F2: 65535}
        _ = v5
    } else {
        var v6 float32 = (1.5 * 1e10)
        v7 := GenS2{F0: 32767}
        _ = v6
        _ = v7
    }
    v8 := GenS3{F0: 2147483647, F1: 4294967295, F2: 128}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    _ = v4
    _ = v8
    return 1
}
func genFn1(p0 string) uint32 {
    fmt.Println(p0)
    v9 := GenS4{F0: -916, F1: -1, F2: -46}
    _ = p0
    _ = v9
    return 2147483647
}

func main() {
    var v1 bool = (true || (-2.25 <= float64((1e10 - 0.0))))
    v2 := make(map[int]bool)
    var v3 int64 = -867
    fmt.Println(v1)
    for fi10 := 0; fi10 < 13; fi10++ {
        _ = len(v2)
        for fi11 := 0; fi11 < 9; fi11++ {
            if (v3 != (^(v3))) {
                fmt.Println(v3)
            } else {
                var v12 uint32 = genFn1("a")
                switch fi11 {
                case 4294967295:
                    v13 := []int{(fi11 * fi11)}
                    _ = v13
                case 32768:
                    v14 := []int{fi11, fi10, fi11, -17}
                    delete(v2, fi11)
                    v12 -= v12
                    _ = v14
                }
                v1 = true
                _ = v12
            }
            switch "\n\t" {
            case "a":
                var v15 bool = v1
                _ = v15
            case "foo bar":
                v1 = false
            case "hello":
                v16 := GenS5{F0: 32767}
                _ = v16
            default:
                v17 := GenS6{F0: 492, F1: 1, F2: v3}
                _ = v17
            }
            _ = fi11
        }
        _ = fi10
    }
    fmt.Println(v1, v1)
    v3 += 1
    v18 := GenS2{F0: 2}
    switch "\n\t" {
    case "hello":
        var v19 int8 = 0
        _ = v19
    case "":
        for fi20 := 0; fi20 < 20; fi20++ {
            if v1 {
                v21 := GenS7{F0: 65535, F1: 42}
                for fi22 := 0; fi22 < 9; fi22++ {
                    var v23 uint64 = 1
                    var v24 int16 = 1000
                    fmt.Println(fi22, v1, v24)
                    switch -651 {
                    case -81:
                        var v25 uint8 = 201
                        _ = v25
                    case 4294967295:
                        v26 := GenS5{F0: int(v24)}
                        _ = v26
                    case 9223372036854775807:
                        var v27 bool = ((-1 < (v3 >> ((1) & 63))) && (!true))
                        _ = v27
                    default:
                        v24 = 600
                    }
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                delete(v2, (fi20 + 510))
                _ = v21
            }
            switch "" {
            case "\n\t":
                fmt.Println(v1)
            case "hello":
                delete(v2, (fi20 | (^(fi20))))
            default:
                v1 = v1
            }
            var v28 int64 = v3
            _ = fi20
            _ = v28
        }
    case "foo bar":
        v29 := []int{1, 0, 9223372036854775807}
        _ = v29
    }
    var v30 int8 = 42
    var v31 uint16 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v30
    _ = v31
}
