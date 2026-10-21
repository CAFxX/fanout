package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint8; F1 uint16 }
type GenS5 struct { F0 string }
func genFn0(p0 bool, p1 uint) int64 {
    v5 := GenS1{F0: (p1 + p1), F1: 190, F2: "\n\t"}
    switch "" {
    case "foo bar":
        switch 42 {
        case 42:
            v6 := GenS1{F0: (^(p1)), F1: 201, F2: "\n\t"}
            _ = v6
        case 128:
            switch "hello" {
            case "":
                var v7 int8 = 98
                _ = v7
            case "\n\t":
                for fi8 := 0; fi8 < 15; fi8++ {
                    switch "hello" {
                    case "hello":
                        var v9 float32 = (1e10 / 0.1)
                        _ = v9
                    }
                    var v10 int32 = -1
                    var v11 int32 = (^(v10))
                    _ = fi8
                    _ = v10
                    _ = v11
                }
            }
        }
    case "hello":
        v12 := []int{}
        _ = v12
    case "":
        p0 = true
    }
    fmt.Println(p1, p1)
    p0 = (!((1 >= 1) && (!(true && true))))
    for fi13 := 0; fi13 < 20; fi13++ {
        for fi14 := 0; fi14 < 8; fi14++ {
            var v15 int32 = 936
            _ = fi14
            _ = v15
        }
        switch "\n\t" {
        case "":
            v16 := GenS2{F0: uint16(p1)}
            _ = v16
        case "foo bar":
            for fi17 := 0; fi17 < 13; fi17++ {
                p0 = (511 != 42)
                v18 := []int{(^(fi13)), 4294967295}
                var v19 bool = (-4 >= 1)
                _ = fi17
                _ = v18
                _ = v19
            }
        case "a":
            v20 := GenS3{F0: 42}
            _ = v20
        }
        v21 := GenS3{F0: 809}
        _ = fi13
        _ = v21
    }
    _ = p0
    _ = p1
    _ = v5
    return 460
}

func main() {
    v1 := make(map[int]int)
    var v2 int32 = -1
    v3 := []int{758, 2147483647}
    v4 := make(map[uint32]bool)
    for fi22 := 0; fi22 < 3; fi22++ {
        for fi23 := 0; fi23 < 9; fi23++ {
            switch ("hello" + ("\n\t" + "hello")) {
            case "\n\t":
                v24 := GenS4{F0: 127, F1: 0}
                _ = v24
            case "":
                v25 := GenS3{F0: int64(fi23)}
                _ = v25
            default:
                var v26 int8 = 127
                _ = v26
            }
            var v27 float64 = 3.14159
            _ = fi23
            _ = v27
        }
        fmt.Println(v2)
        if false {
            fmt.Println(v2, fi22)
        } else {
            _ = len(v3)
        }
        _ = fi22
    }
    v2 = v2
    v3 = append(v3, 128)
    delete(v4, 0)
    fmt.Println(v2)
    var v28 uint32 = 4294967295
    v29 := GenS5{F0: ("a" + "\n\t")}
    _ = len(v3)
    v2 -= v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
