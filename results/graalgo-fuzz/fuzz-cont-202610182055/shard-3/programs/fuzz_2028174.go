package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64; F1 bool; F2 string }
func genFn0(p0 uint32, p1 int64, p2 uint) uint {
    var v6 string = "a"
    switch -2147483648 {
    case -1:
        if true {
            p1 = int64(p2)
        }
    case 128:
        var v7 bool = true
        _ = v7
    default:
        for fi8 := 0; fi8 < 7; fi8++ {
            v9 := []int{-2147483648}
            _ = fi8
            _ = v9
        }
    }
    v10 := GenS1{F0: (p0 & p0)}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v6
    return (^(p2))
}
func genFn1(p0 int32, p1 uint8, p2 int16) int64 {
    v11 := []int{}
    v12 := GenS1{F0: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 9223372036854775807
}

func main() {
    var v1 uint8 = 127
    v2 := make(map[uint32]int)
    v3 := []uint32{13, 65535, 2, 65535}
    v4 := make(map[uint32]int)
    v5 := make(map[string]uint32)
    switch 32767 {
    case 128:
        var v13 string = "hello"
        _ = v13
    case 4294967295:
        for fi14 := 0; fi14 < 4; fi14++ {
            var v15 int8 = 100
            fmt.Println(v15)
            _ = fi14
            _ = v15
        }
    default:
        for fi16 := 0; fi16 < 7; fi16++ {
            for fi17 := 0; fi17 < 12; fi17++ {
                var v18 int = fi16
                _ = fi17
                _ = v18
            }
            _ = fi16
        }
    }
    v19 := GenS1{F0: 65535}
    v1 += v1
    switch ("a" + "\n\t") {
    case "hello":
        fmt.Println(v1)
    }
    v20 := GenS1{F0: 42}
    v1 = 12
    for fi21 := 0; fi21 < 3; fi21++ {
        for fi22 := 0; fi22 < 16; fi22++ {
            switch "hello" {
            case "\n\t":
                switch "\n\t" {
                case "foo bar":
                    v23 := GenS1{F0: 13}
                    _ = v23
                case "\n\t":
                    fmt.Println(v1, v1)
                }
            case "hello":
                if (fi22 >= (fi21 * (^(fi22)))) {
                    var v24 float32 = float32(float64((-2.25 + 0.0)))
                    v25 := GenS1{F0: 42}
                    switch "foo bar" {
                    case "a":
                        v1 -= v1
                    case "foo bar":
                        v26 := GenS2{F0: 121}
                        _ = v26
                    default:
                        _ = len(v3)
                    }
                    _ = v24
                    _ = v25
                }
            case "a":
                v2[uint32(v1)] = fi21
            }
            _ = len(v3)
            _ = len(v3)
            _ = fi22
        }
        if (9223372036854775807 < 416) {
            v27 := GenS3{F0: uint64(v1), F1: true, F2: ""}
            _ = v27
        }
        _ = fi21
    }
    var v28 uint = 42
    if ((!(!(false || true))) && false) {
        if false {
            v1 -= v1
            var v29 int64 = -9223372036854775808
            _ = v29
        }
        for fi30 := 0; fi30 < 6; fi30++ {
            for fi31 := 0; fi31 < 6; fi31++ {
                if (genFn0(42, -9223372036854775808, (v28 % ((uint(v1)) | 1))) != (^(v28))) {
                    v32 := v4[2]
                    v33 := GenS2{F0: 100}
                    if (!true) {
                        if false {
                            v34 := GenS2{F0: 127}
                            _ = v34
                        }
                        var v35 float32 = 1e10
                        _ = v35
                    }
                    _ = v32
                    _ = v33
                }
                fmt.Println(fi30, fi31)
                _ = fi31
            }
            _ = fi30
        }
    } else {
        for fi36 := 0; fi36 < 16; fi36++ {
            var v37 uint = v28
            _ = fi36
            _ = v37
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
