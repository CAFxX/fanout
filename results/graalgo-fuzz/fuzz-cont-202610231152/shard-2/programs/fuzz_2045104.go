package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8 }
func genFn0(p0 uint16) int8 {
    v7 := GenS1{F0: (("" != "a") && ((p0 < (p0 & p0)) && true)), F1: 1}
    var v8 float32 = (-2.25 - float32((-2.25 + float64(1.5))))
    fmt.Println(p0, v8)
    var v9 float64 = (float64((v8 - v8)) + 0.1)
    v9 = ((v9 * 1e10) * v9)
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return -1
}
func genFn1() uint16 {
    var v10 int64 = -1
    v11 := GenS2{F0: -137}
    for fi12 := 0; fi12 < 10; fi12++ {
        fmt.Println(fi12, v10)
        _ = fi12
    }
    v10 += 1
    fmt.Println(v10)
    _ = v10
    _ = v11
    return 992
}

func main() {
    var v1 int = 32767
    v2 := []int{}
    v3 := []string{"a", ("foo bar" + "foo bar")}
    v4 := []uint32{}
    v5 := make(map[int]bool)
    v6 := make(map[uint32]string)
    v1 = v1
    v1 = 2147483647
    for fi13 := 0; fi13 < 3; fi13++ {
        for fi14 := 0; fi14 < 7; fi14++ {
            if true {
                switch "foo bar" {
                case "foo bar":
                    v2 = append(v2, 128)
                default:
                    if (553 > 32767) {
                        v1 = v1
                        v1 = 178
                        var v15 uint16 = 42
                        _ = v15
                    } else {
                        switch (fi14 % (((v1 % ((32768) | 1))) | 1)) {
                        case 128:
                            v1 = fi14
                        case 42:
                            v1 += -521
                        case 1:
                            v16 := v6[0]
                            _ = v16
                        }
                        v1 = 4294967295
                        fmt.Println(fi14, v1, fi14)
                    }
                }
                switch "\n\t" {
                case "":
                    _ = len(v3)
                case "foo bar":
                    v6[651] = ""
                case "a":
                    fmt.Println(fi13, fi13, v1)
                }
                v17 := GenS1{F0: false, F1: 0}
                _ = v17
            }
            v6[2147483648] = "hello"
            var v18 uint8 = 127
            switch "" {
            case "\n\t":
                var v19 int64 = 42
                _ = v19
            default:
                v20 := GenS1{F0: ((1.5 / (((0.1 / float64((-2.25 + (0.1 * (0.1 / ((float32(3.14159) - 0.0) * 0.1)))))) * 1.5) * float64(0.1))) <= -2.25), F1: 2147483648}
                _ = v20
            }
            _ = fi14
            _ = v18
        }
        _ = fi13
    }
    v5[v1] = true
    if false {
        var v21 uint32 = 532
        fmt.Println(v21)
        var v22 bool = (v1 <= (v1 ^ 9223372036854775807))
        _ = v21
        _ = v22
    } else {
        var v23 int16 = -930
        _ = len(v3)
        v24 := GenS1{F0: true, F1: 65535}
        _ = v23
        _ = v24
    }
    var v25 int16 = 1
    for fi26 := 0; fi26 < 2; fi26++ {
        switch 736 {
        case -9223372036854775808:
            if len(v2) > 0 {
                v2[0] = v1
            }
        case 32767:
            var v27 string = "a"
            _ = v27
        default:
            switch (("a" + ("hello" + "")) + "") {
            case "hello":
                var v28 uint = 1
                _ = v28
            case "a":
                var v29 uint8 = 127
                _ = v29
            }
        }
        v6[1] = "hello"
        if false {
            fmt.Println(v25)
            var v30 int8 = 0
            var v31 float32 = 1.5
            _ = v30
            _ = v31
        }
        _ = fi26
    }
    fmt.Println(v25)
    v25 = (v25 << ((476) & 63))
    switch "foo bar" {
    case "foo bar":
        var v32 uint64 = 18446744073709551615
        _ = v32
    case "":
        v33 := GenS3{F0: 1}
        _ = v33
    }
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
