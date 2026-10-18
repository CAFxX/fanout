package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 uint64; F1 int8; F2 int }
func genFn0(p0 uint32, p1 uint64, p2 int64) int {
    var v4 uint16 = 2
    v4 = v4
    v5 := GenS1{F0: "foo bar"}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return int(p1)
}

func main() {
    var v1 int16 = 1
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]uint32)
    switch genFn0(2147483648, 2, 1) {
    case 0:
        var v6 bool = true
        _ = v6
    case 2147483647:
        switch ("" + "foo bar") {
        case "foo bar":
            var v7 int32 = 0
            _ = v7
        case "\n\t":
            v1 = (v1 & (v1 % ((v1) | 1)))
        case "":
            v8 := GenS1{F0: "foo bar"}
            _ = v8
        default:
            v9 := GenS1{F0: "a"}
            _ = v9
        }
    }
    if (-446 < 1) {
        switch "\n\t" {
        case "foo bar":
            v10 := GenS1{F0: ("hello" + (("" + "") + "a"))}
            _ = v10
        case "a":
            v1 = 1
        case "\n\t":
            for fi11 := 0; fi11 < 3; fi11++ {
                v12 := v2[627]
                _ = fi11
                _ = v12
            }
        default:
            for fi13 := 0; fi13 < 6; fi13++ {
                switch "\n\t" {
                case "foo bar":
                    for fi14 := 0; fi14 < 17; fi14++ {
                        fmt.Println(fi13, fi14, fi13)
                        v15 := []int{(fi14 / (((fi14 - 189)) | 1)), genFn0(42, 839, 1)}
                        _ = fi14
                        _ = v15
                    }
                case "a":
                    var v16 int8 = -10
                    _ = v16
                case "hello":
                    var v17 int = -468
                    _ = v17
                }
                _ = fi13
            }
        }
    }
    delete(v3, 18)
    v3[0] = 2147483648
    for fi18 := 0; fi18 < 5; fi18++ {
        fmt.Println(v1, fi18)
        v19 := GenS2{F0: (float32(3.14159) <= 1.5), F1: 127}
        _ = fi18
        _ = v19
    }
    fmt.Println(v1)
    for fi20 := 0; fi20 < 10; fi20++ {
        if false {
            v21 := GenS2{F0: true, F1: 147}
            if false {
                v22 := []int{fi20}
                fmt.Println(v1, fi20)
                _ = v22
            }
            _ = v21
        } else {
            switch fi20 {
            case 4294967295:
                var v23 int16 = (v1 - v1)
                _ = v23
            case 9223372036854775807:
                fmt.Println(fi20, fi20)
            case 377:
                var v24 string = "\n\t"
                _ = v24
            }
            if (-1 >= -128) {
                v25 := []int{(fi20 ^ fi20), 415, (-(fi20)), fi20}
                for fi26 := 0; fi26 < 14; fi26++ {
                    v1 += -897
                    var v27 int32 = -606
                    var v28 uint64 = 969
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                _ = v25
            }
            var v29 float64 = (float64(((-2.25 + (3.14159 / 1.0)) / 1e10)) - float64((0.1 / 1.5)))
            _ = v29
        }
        for fi30 := 0; fi30 < 16; fi30++ {
            var v31 int32 = 42
            v32 := GenS3{F0: 580, F1: 100, F2: fi30}
            _ = fi30
            _ = v31
            _ = v32
        }
        _ = fi20
    }
    fmt.Println(v1)
    fmt.Println(v1)
    v33 := []int{}
    v34 := v2[4294967295]
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
}
