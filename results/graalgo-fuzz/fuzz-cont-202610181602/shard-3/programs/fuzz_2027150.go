package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint; F1 string }
type GenS4 struct { F0 bool; F1 int8; F2 string }
func genFn0(p0 uint, p1 uint64, p2 uint16) uint {
    var v5 uint32 = 967
    switch -2147483648 {
    case -1:
        v5 = (v5 + v5)
    default:
        switch "a" {
        case "\n\t":
            v6 := GenS1{F0: -1, F1: (v5 + v5)}
            _ = v6
        case "":
            var v7 uint64 = 4294967295
            _ = v7
        }
    }
    v8 := GenS2{F0: (^(p2))}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v8
    return uint(v5)
}

func main() {
    var v1 int32 = 304
    var v2 uint = 64
    v3 := []int{32767}
    v4 := make(map[uint32]int)
    var v9 bool = (0 == 613)
    for fi10 := 0; fi10 < 12; fi10++ {
        switch "hello" {
        case "":
            fmt.Println(fi10, fi10, v2)
        case "foo bar":
            fmt.Println(v1, v2, fi10)
        default:
            v3 = append(v3, (fi10 % ((fi10) | 1)))
        }
        switch "\n\t" {
        case "a":
            if v9 {
                switch (fi10 / (((fi10 >> ((genFn0((v2 ^ v2), 1, 42)) & 63))) | 1)) {
                case -2147483648:
                    for fi11 := 0; fi11 < 9; fi11++ {
                        if (255 >= 42) {
                            var v12 uint64 = 4294967295
                            _ = v12
                        }
                        _ = fi11
                    }
                case 128:
                    fmt.Println(fi10, fi10)
                case 127:
                    var v13 float64 = 1e10
                    _ = v13
                default:
                    v14 := GenS1{F0: (v1 & v1), F1: 976}
                    _ = v14
                }
                v2 = v2
            }
        default:
            var v15 float32 = 1e10
            _ = v15
        }
        v16 := GenS3{F0: (^(v2)), F1: ("" + "hello")}
        _ = fi10
        _ = v16
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v9 = (true && false)
    _ = len(v4)
    v1 -= (v1 - v1)
    if v9 {
        fmt.Println(v9, v2, v1)
        for fi17 := 0; fi17 < 12; fi17++ {
            for fi18 := 0; fi18 < 4; fi18++ {
                fmt.Println(fi17, fi18)
                v19 := GenS4{F0: (159 == 128), F1: -41, F2: "a"}
                _ = fi18
                _ = v19
            }
            if (!(v2 < genFn0((v2 % (((v2 * (^(v2)))) | 1)), 42, 1))) {
                fmt.Println(v2, v2, fi17)
            }
            var v20 float32 = 1e10
            _ = fi17
            _ = v20
        }
    } else {
        switch (("a" + "foo bar") + "") {
        case "hello":
            v1 += 657
        case "foo bar":
            var v21 uint8 = 127
            _ = v21
        default:
            var v22 float64 = 0.0
            _ = v22
        }
        var v23 uint = (v2 * (^(v2)))
        _ = v23
    }
    for fi24 := 0; fi24 < 8; fi24++ {
        for fi25 := 0; fi25 < 6; fi25++ {
            v2 += (v2 / ((genFn0(780, 18446744073709551615, 65535)) | 1))
            var v26 float64 = 1.5
            _ = fi25
            _ = v26
        }
        fmt.Println(fi24, fi24, v2)
        _ = fi24
    }
    var v27 int32 = v1
    if len(v3) > 0 {
        v3[0] = 284
    }
    if v9 {
        fmt.Println(v9, v1)
        v27 += -2147483648
        switch "hello" {
        case "hello":
            for fi28 := 0; fi28 < 17; fi28++ {
                var v29 uint64 = 42
                switch "foo bar" {
                case "":
                    var v30 bool = v9
                    _ = v30
                case "hello":
                    if len(v3) > 0 {
                        v3[0] = (fi28 + fi28)
                    }
                case "a":
                    _ = len(v3)
                default:
                    v3 = append(v3, (^(fi28)))
                }
                fmt.Println(v1)
                _ = fi28
                _ = v29
            }
        }
    }
    var v31 bool = false
    if (v31 || (!(!(1 > 712)))) {
        v32 := v4[2147483647]
        var v33 bool = (594 == 1)
        for fi34 := 0; fi34 < 1; fi34++ {
            var v35 uint8 = 42
            var v36 uint = v2
            _ = fi34
            _ = v35
            _ = v36
        }
        _ = v32
        _ = v33
    }
    var v37 bool = (690 == (v2 % ((v2) | 1)))
    fmt.Println(v9, v1)
    var v38 float64 = (1.5 + 1.5)
    v9 = false
    fmt.Println(v27, v9, v31)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v31
    _ = v37
    _ = v38
    _ = v4
    _ = v9
}
