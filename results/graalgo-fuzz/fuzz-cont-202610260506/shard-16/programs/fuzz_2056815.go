package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 bool; F1 uint16 }
type GenS5 struct { F0 uint32; F1 int }
type GenS6 struct { F0 uint16; F1 bool; F2 uint }
type GenS7 struct { F0 int64; F1 int64; F2 int32 }
func genFn0(p0 uint) uint64 {
    fmt.Println(p0)
    var v7 uint32 = 2147483647
    for fi8 := 0; fi8 < 17; fi8++ {
        v9 := GenS1{F0: -2147483648, F1: 1}
        v10 := GenS2{F0: 4294967295, F1: 42}
        v7 = (v7 & v7)
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = v7
    return 1
}
func genFn1(p0 uint, p1 bool) uint32 {
    var v11 uint32 = 271
    var v12 uint64 = 2
    if true {
        fmt.Println(p0)
    } else {
        for fi13 := 0; fi13 < 8; fi13++ {
            switch (fi13 ^ (fi13 % ((fi13) | 1))) {
            case 128:
                var v14 string = "\n\t"
                _ = v14
            case -730:
                for fi15 := 0; fi15 < 5; fi15++ {
                    var v16 uint8 = 224
                    fmt.Println(fi15)
                    var v17 string = "foo bar"
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            default:
                var v18 float64 = 3.14159
                _ = v18
            }
            _ = fi13
        }
        fmt.Println(v11, p0, p0)
    }
    if p1 {
        var v19 uint64 = 991
        _ = v19
    }
    var v20 int8 = 0
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v20
    return (^(v11))
}

func main() {
    v1 := []int{-1, -9223372036854775808, 42, 2147483647}
    v2 := make(map[int]uint32)
    var v3 int = -2147483648
    v4 := make(map[int]uint32)
    v5 := []int{}
    v6 := make(map[uint32]string)
    v21 := GenS3{F0: genFn0(339), F1: 0}
    delete(v2, v3)
    var v22 int32 = 2147483647
    v23 := GenS2{F0: 0, F1: -32768}
    fmt.Println(v3, v3)
    var v24 uint8 = 0
    for fi25 := 0; fi25 < 12; fi25++ {
        v24 = 1
        var v26 int = v3
        for fi27 := 0; fi27 < 10; fi27++ {
            var v28 string = ""
            v29 := GenS2{F0: v3, F1: int16(v22)}
            _ = fi27
            _ = v28
            _ = v29
        }
        _ = fi25
        _ = v26
    }
    _ = len(v5)
    v30 := GenS3{F0: 848, F1: 167}
    fmt.Println(v3, v22)
    switch ("\n\t" + (((("" + "\n\t") + ("a" + "a")) + "") + ("foo bar" + ""))) {
    case "hello":
        v31 := GenS4{F0: true, F1: 255}
        _ = v31
    default:
        switch "" {
        case "":
            switch (^(v3)) {
            case 2147483647:
                var v32 int32 = -2147483648
                _ = v32
            case 621:
                switch v3 {
                case 593:
                    var v33 int16 = -1
                    _ = v33
                case 32767:
                    v34 := GenS5{F0: 241, F1: (v3 / ((int(v22)) | 1))}
                    _ = v34
                case 744:
                    v22 += -1
                }
            }
        case "a":
            if false {
                v3 -= 9223372036854775807
                var v35 int64 = 0
                _ = v35
            } else {
                v36 := GenS6{F0: 444, F1: ((1 == 1) || false), F2: 42}
                _ = v36
            }
        case "foo bar":
            if true {
                v37 := GenS4{F0: true, F1: 487}
                for fi38 := 0; fi38 < 15; fi38++ {
                    var v39 int32 = v22
                    _ = fi38
                    _ = v39
                }
                v40 := GenS1{F0: (-(v22)), F1: 127}
                _ = v37
                _ = v40
            } else {
                v41 := GenS7{F0: 0, F1: -969, F2: 876}
                v3 = (v3 * (v3 >> ((uint(v24)) & 63)))
                v6[983] = "foo bar"
                _ = v41
            }
        }
    }
    var v42 float32 = (0.1 + 3.14159)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v30
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
