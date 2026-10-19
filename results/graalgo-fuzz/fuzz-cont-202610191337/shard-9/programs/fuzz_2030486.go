package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64; F1 int; F2 bool }

func main() {
    var v1 int8 = 58
    var v2 bool = true
    var v3 uint = 787
    for fi4 := 0; fi4 < 6; fi4++ {
        fmt.Println(fi4)
        var v5 uint8 = 127
        _ = fi4
        _ = v5
    }
    var v6 string = ("foo bar" + "")
    v7 := []int{-1, -1, 128}
    v8 := GenS1{F0: 297, F1: -128}
    var v9 float32 = (1.5 - (0.0 - 1.5))
    v2 = v2
    if v2 {
        var v10 uint = v3
        _ = v10
    } else {
        var v11 int64 = -9223372036854775808
        _ = v11
    }
    for fi12 := 0; fi12 < 3; fi12++ {
        var v13 int16 = -8
        v9 += 0.1
        v14 := GenS1{F0: v13, F1: (fi12 * 128)}
        _ = fi12
        _ = v13
        _ = v14
    }
    v1 = (v1 & v1)
    var v15 int8 = -15
    v16 := GenS1{F0: 323, F1: 1}
    if (!v2) {
        switch 32768 {
        case 127:
            if v2 {
                if false {
                    v17 := GenS1{F0: -1, F1: 4294967295}
                    var v18 uint64 = 0
                    _ = v17
                    _ = v18
                } else {
                    v6 = v6
                    for fi19 := 0; fi19 < 2; fi19++ {
                        fmt.Println(v1, fi19)
                        v7 = append(v7, fi19)
                        _ = fi19
                    }
                }
                for fi20 := 0; fi20 < 13; fi20++ {
                    v2 = true
                    var v21 bool = ((128 >= 1) && (2 >= 1))
                    switch 434 {
                    case 0:
                        v21 = true
                    case 4294967295:
                        v1 += (v1 * (v15 >> ((65535) & 63)))
                    case 128:
                        var v22 int8 = int8(fi20)
                        _ = v22
                    }
                    _ = fi20
                    _ = v21
                }
            } else {
                if (!v2) {
                    if v2 {
                        var v23 uint = (v3 % ((v3) | 1))
                        v24 := GenS2{F0: (-931 == 32767)}
                        if v2 {
                            v1 -= v1
                            v3 = 65535
                            var v25 uint32 = 824
                            _ = v25
                        } else {
                            if (42 > 63) {
                                fmt.Println(v6, v9, v3)
                            }
                        }
                        _ = v23
                        _ = v24
                    }
                }
                var v26 uint8 = 128
                var v27 float64 = 1e10
                _ = v26
                _ = v27
            }
        case 32768:
            switch "foo bar" {
            case "hello":
                v28 := GenS1{F0: 42, F1: -1}
                _ = v28
            default:
                v29 := GenS3{F0: v2}
                _ = v29
            }
        case -272:
            var v30 int64 = 0
            _ = v30
        }
    }
    switch "a" {
    case "foo bar":
        for fi31 := 0; fi31 < 17; fi31++ {
            v32 := GenS4{F0: 18446744073709551615, F1: fi31, F2: true}
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = fi31
            _ = v32
        }
    case "a":
        if (!(2 != 2)) {
            for fi33 := 0; fi33 < 17; fi33++ {
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                fmt.Println(v9, v2, v6)
                _ = fi33
            }
        } else {
            var v34 int = -304
            _ = v34
        }
    }
    v35 := GenS2{F0: (!(52 != 127))}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v35
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
