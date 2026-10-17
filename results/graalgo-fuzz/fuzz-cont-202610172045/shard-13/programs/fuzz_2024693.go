package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 int }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 bool }

func main() {
    var v1 int8 = 42
    var v2 uint64 = 18446744073709551615
    v3 := make(map[string]uint32)
    var v4 int32 = 0
    v5 := make(map[uint32]int)
    var v6 int64 = 1
    v2 = v2
    v6 += v6
    for fi7 := 0; fi7 < 2; fi7++ {
        v8 := []int{}
        _ = fi7
        _ = v8
    }
    var v9 uint8 = 255
    if true {
        var v10 uint32 = 2147483647
        v4 = 1
        switch "a" {
        case "foo bar":
            var v11 int8 = (v1 & v1)
            _ = v11
        case "":
            v4 -= 1
        case "hello":
            fmt.Println(v1, v4)
        }
        _ = v10
    }
    fmt.Println(v4, v6)
    v2 += v2
    fmt.Println(v1)
    switch "hello" {
    case "\n\t":
        v2 = (^(v2))
    case "a":
        switch "foo bar" {
        case "":
            var v12 int16 = 1000
            _ = v12
        case "foo bar":
            switch 128 {
            case -938:
                fmt.Println(v9, v2)
            case 0:
                v6 += (v6 % ((-1) | 1))
            case 65:
                fmt.Println(v6)
            }
        case "a":
            if true {
                if false {
                    for fi13 := 0; fi13 < 13; fi13++ {
                        for fi14 := 0; fi14 < 1; fi14++ {
                            v15 := GenS1{F0: (-(v4)), F1: v9, F2: fi13}
                            v6 += 0
                            if (!false) {
                                fmt.Println(v2)
                            }
                            var v16 uint8 = (v9 << ((354) & 63))
                            v17 := []int{-9223372036854775808, fi14, fi13, fi14}
                            v18 := GenS1{F0: (v4 / ((890) | 1)), F1: (v9 | (v16 - v16)), F2: fi14}
                            _ = fi14
                            _ = v15
                            _ = v16
                            _ = v17
                            _ = v18
                        }
                        delete(v3, "")
                        _ = fi13
                    }
                } else {
                    v19 := GenS1{F0: -283, F1: 0, F2: 1}
                    _ = v19
                }
                v9 += (v9 >> ((65535) & 63))
            } else {
                v20 := []int{-654, int(v6), 508}
                var v21 int8 = v1
                for fi22 := 0; fi22 < 5; fi22++ {
                    for fi23 := 0; fi23 < 6; fi23++ {
                        if true {
                            switch "hello" {
                            case "a":
                                v24 := GenS2{F0: -569, F1: 42}
                                _ = v24
                            case "hello":
                                v6 = (v6 + v6)
                            }
                        } else {
                            v21 -= v1
                            var v25 uint = 65535
                            fmt.Println(v6, v21, fi22)
                            v21 -= v21
                            _ = v25
                        }
                        if (!true) {
                            v9 = (v9 | 255)
                        }
                        fmt.Println(v1, v4, fi22)
                        _ = fi23
                    }
                    _ = fi22
                }
                _ = v20
                _ = v21
            }
        }
    case "hello":
        v26 := GenS3{F0: false}
        _ = v26
    }
    if (false && ((v4 ^ (v4 << ((0) & 63))) == v4)) {
        for fi27 := 0; fi27 < 7; fi27++ {
            v4 = 1
            _ = fi27
        }
        for fi28 := 0; fi28 < 10; fi28++ {
            v6 = (v6 / (((^(v6))) | 1))
            _ = fi28
        }
        v29 := []int{9223372036854775807, 128, -193, 9223372036854775807}
        _ = v29
    }
    v30 := v3["hello"]
    for fi31 := 0; fi31 < 4; fi31++ {
        fmt.Println(v1, v4, v2)
        var v32 int = 2147483647
        v33 := []int{(v32 >> ((4294967295) & 63)), (v32 % ((int(v6)) | 1))}
        _ = fi31
        _ = v32
        _ = v33
    }
    var v34 int8 = v1
    var v35 int16 = -1
    switch "\n\t" {
    case "a":
        v36 := []int{}
        _ = v36
    case "":
        v37 := GenS3{F0: (float64(1.5) != 1e10)}
        _ = v37
    default:
        if (!(((0.0 * 0.1) == 0.1) || (!(34 >= 4294967295)))) {
            for fi38 := 0; fi38 < 4; fi38++ {
                v6 = (v6 * (v6 ^ v6))
                _ = fi38
            }
            fmt.Println(v4, v9, v35)
            if true {
                v9 = (v9 / ((uint8(v30)) | 1))
                fmt.Println(v6, v2)
                var v39 int16 = v35
                _ = v39
            }
        }
    }
    var v40 uint16 = 828
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v34
    _ = v35
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v9
}
