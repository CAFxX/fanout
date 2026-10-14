package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 int8; F1 uint64 }
func genFn0(p0 bool, p1 int) int16 {
    for fi4 := 0; fi4 < 8; fi4++ {
        if p0 {
            fmt.Println(p0)
            fmt.Println(p0, p1)
            p0 = (2 == 285)
        } else {
            var v5 int64 = -1
            _ = v5
        }
        v6 := GenS1{F0: 0, F1: 2, F2: -1}
        if p0 {
            v7 := GenS1{F0: 65535, F1: 0, F2: -32768}
            _ = v7
        }
        _ = fi4
        _ = v6
    }
    var v8 int = (p1 & p1)
    _ = p0
    _ = p1
    _ = v8
    return 0
}
func genFn1(p0 uint16, p1 int64) bool {
    p0 = 42
    v9 := GenS2{F0: 4294967295}
    v10 := GenS2{F0: 4294967295}
    switch 1 {
    case -718:
        v11 := GenS3{F0: 477, F1: 410}
        _ = v11
    case 4294967295:
        var v12 string = "foo bar"
        _ = v12
    case 0:
        for fi13 := 0; fi13 < 3; fi13++ {
            fmt.Println(p0)
            fmt.Println(p0)
            _ = fi13
        }
    default:
        for fi14 := 0; fi14 < 16; fi14++ {
            var v15 string = ("a" + "a")
            _ = fi14
            _ = v15
        }
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return (genFn0(false, 702) < 0)
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint16 = 2
    var v3 int32 = 2147483647
    v2 = 32767
    for fi16 := 0; fi16 < 6; fi16++ {
        v17 := []int{32768, -1, fi16, fi16}
        v3 = 47
        _ = fi16
        _ = v17
    }
    v18 := GenS2{F0: 278}
    v3 = v3
    v2 += 110
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v2)
    case "hello":
        for fi19 := 0; fi19 < 8; fi19++ {
            v3 = (^(v3))
            v3 += 2147483647
            _ = fi19
        }
    }
    switch "hello" {
    case "foo bar":
        for fi20 := 0; fi20 < 9; fi20++ {
            var v21 int8 = -1
            var v22 int16 = genFn0(false, fi20)
            _ = fi20
            _ = v21
            _ = v22
        }
    case "":
        if false {
            fmt.Println(v3, v3)
            if (-669 < -32768) {
                fmt.Println(v3)
                for fi23 := 0; fi23 < 18; fi23++ {
                    var v24 string = ""
                    _ = fi23
                    _ = v24
                }
            } else {
                var v25 int = 1
                switch "" {
                case "foo bar":
                    v26 := GenS4{F0: 42, F1: 0}
                    _ = v26
                default:
                    var v27 bool = true
                    _ = v27
                }
                _ = v25
            }
        }
    case "\n\t":
        delete(v1, 594)
    }
    fmt.Println(v3, v2)
    var v28 bool = (4294967295 < 419)
    v29 := []int{}
    v2 = (v2 | (v2 - 452))
    for fi30 := 0; fi30 < 20; fi30++ {
        v31 := v1[182]
        var v32 uint64 = 798
        if (0 < 903) {
            var v33 uint8 = 128
            var v34 uint32 = (v31 & (^(v31)))
            var v35 int64 = -9223372036854775808
            _ = v33
            _ = v34
            _ = v35
        } else {
            v32 = (v32 % ((4294967295) | 1))
            var v36 uint32 = (v31 & v31)
            v36 -= v36
            _ = v36
        }
        _ = fi30
        _ = v31
        _ = v32
    }
    delete(v1, 2147483647)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v2
    _ = v28
    _ = v29
    _ = v3
}
