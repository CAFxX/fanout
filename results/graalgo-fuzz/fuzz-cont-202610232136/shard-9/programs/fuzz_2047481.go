package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 bool; F1 bool; F2 bool }
type GenS6 struct { F0 int; F1 uint16 }
func genFn0(p0 string, p1 int64) int {
    var v6 bool = (235 <= 65535)
    if true {
        switch 172 {
        case 224:
            v7 := GenS1{F0: "foo bar"}
            _ = v7
        case 2147483647:
            v8 := []int{9223372036854775807}
            _ = v8
        case -837:
            v6 = (313 <= 987)
        default:
            var v9 bool = (3.14159 > (float64((-2.25 + 1.5)) - ((0.1 - 1.5) * 3.14159)))
            _ = v9
        }
    } else {
        v10 := GenS1{F0: (p0 + ("\n\t" + (p0 + p0)))}
        var v11 int16 = 1
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v6
    return -268
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint16 = 0
    var v3 int64 = -1
    var v4 int8 = -1
    v5 := make(map[int]string)
    if (32767 >= 1) {
        var v12 int8 = int8(v2)
        v3 = (-(v3))
        v13 := []int{genFn0("a", v3), 127}
        _ = v12
        _ = v13
    } else {
        fmt.Println(v4, v2)
    }
    v3 = (^(v3))
    v14 := v5[-2147483648]
    fmt.Println(v2, v4)
    switch "\n\t" {
    case "a":
        fmt.Println(v14)
    case "foo bar":
        v15 := GenS2{F0: 745}
        _ = v15
    case "":
        switch v14 {
        case "a":
            var v16 string = "\n\t"
            _ = v16
        case "hello":
            _ = len(v1)
        case "":
            if false {
                v17 := GenS2{F0: 32767}
                _ = v17
            }
        }
    }
    fmt.Println(v3)
    for fi18 := 0; fi18 < 10; fi18++ {
        v14 = "a"
        if (525 <= 255) {
            v19 := GenS2{F0: 32767}
            var v20 uint = 4294967295
            v3 -= v3
            _ = v19
            _ = v20
        } else {
            v5[genFn0(v14, 9223372036854775807)] = v14
            var v21 uint16 = (v2 % (((v2 * v2)) | 1))
            fmt.Println(v14, v4, v2)
            _ = v21
        }
        v22 := GenS2{F0: int16(v4)}
        _ = fi18
        _ = v22
    }
    switch v14 {
    case "a":
        v23 := GenS3{F0: 992, F1: uint(v2)}
        _ = v23
    case "":
        switch v14 {
        case "a":
            for fi24 := 0; fi24 < 5; fi24++ {
                if ((false || (32767 == int16(v4))) && (119 == -67)) {
                    v25 := []int{(^(fi24)), (fi24 ^ int(v4)), 1}
                    _ = v25
                }
                v4 += 0
                v26 := []int{(fi24 << ((1) & 63)), 9223372036854775807, (-(fi24)), (fi24 & (-(fi24)))}
                _ = fi24
                _ = v26
            }
        case "foo bar":
            fmt.Println(v3)
        default:
            var v27 float32 = float32((3.14159 / float64(float32(-2.25))))
            _ = v27
        }
    }
    var v28 uint32 = 0
    var v29 string = (v14 + v14)
    if true {
        v30 := GenS2{F0: -755}
        var v31 float64 = 0.0
        _ = v30
        _ = v31
    } else {
        if true {
            for fi32 := 0; fi32 < 18; fi32++ {
                v4 -= v4
                _ = len(v1)
                switch (fi32 / (((fi32 % ((fi32) | 1))) | 1)) {
                case -1:
                    v4 = (v4 + v4)
                }
                _ = fi32
            }
            var v33 int32 = 2147483647
            v34 := []int{-2147483648, 677}
            _ = v33
            _ = v34
        }
    }
    switch v29 {
    case "\n\t":
        fmt.Println(v2)
    case "a":
        v35 := []int{-1, -9223372036854775808, 32768, 267}
        _ = v35
    case "foo bar":
        if true {
            for fi36 := 0; fi36 < 7; fi36++ {
                var v37 uint16 = (^(v2))
                for fi38 := 0; fi38 < 7; fi38++ {
                    v14 = ("" + v14)
                    if true {
                        _ = len(v1)
                        if false {
                            var v39 int16 = 1000
                            _ = v39
                        }
                    }
                    v40 := []int{genFn0(v14, (v3 % ((v3) | 1))), (fi38 ^ 42)}
                    _ = fi38
                    _ = v40
                }
                _ = fi36
                _ = v37
            }
            fmt.Println(v28, v29, v14)
            v2 += (v2 << ((0) & 63))
        }
    }
    for fi41 := 0; fi41 < 1; fi41++ {
        v4 = (v4 >> ((18446744073709551615) & 63))
        var v42 bool = ("hello" >= v29)
        v43 := GenS4{F0: (^(v3))}
        _ = fi41
        _ = v42
        _ = v43
    }
    v4 -= v4
    v44 := v1[795]
    for fi45 := 0; fi45 < 18; fi45++ {
        var v46 uint = 2
        _ = fi45
        _ = v46
    }
    for fi47 := 0; fi47 < 8; fi47++ {
        switch fi47 {
        case -1:
            v48 := GenS5{F0: true, F1: (!((v2 * v2) >= (v2 - v2))), F2: (1e10 != float64((1e10 + (3.14159 - 3.14159))))}
            _ = v48
        default:
            fmt.Println(v2, v14, fi47)
        }
        v49 := GenS6{F0: 4294967295, F1: v2}
        _ = fi47
        _ = v49
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v44
    _ = v5
}
