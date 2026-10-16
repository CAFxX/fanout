package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 uint8; F1 uint64; F2 uint8 }
type GenS3 struct { F0 bool; F1 uint }

func main() {
    var v1 string = ("a" + "foo bar")
    var v2 int16 = -285
    var v3 int8 = 42
    v4 := make(map[uint32]uint32)
    var v5 uint32 = 335
    v6 := make(map[uint32]bool)
    var v7 uint64 = 0
    for fi8 := 0; fi8 < 1; fi8++ {
        for fi9 := 0; fi9 < 3; fi9++ {
            v10 := []int{(fi9 % ((-2147483648) | 1)), 4294967295, fi8}
            fmt.Println(v3, v5)
            v7 -= 797
            _ = fi9
            _ = v10
        }
        fmt.Println(v2, fi8)
        var v11 uint16 = 0
        _ = fi8
        _ = v11
    }
    v12 := []int{1, 4294967295, 2147483647}
    if false {
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
    }
    v13 := GenS1{F0: v7, F1: v3}
    fmt.Println(v7, v5)
    var v14 int = -734
    var v15 bool = (-1 >= -1)
    if v15 {
        var v16 uint8 = 42
        v16 += (v16 | 1)
        for fi17 := 0; fi17 < 6; fi17++ {
            if (!(!v15)) {
                if true {
                    v15 = true
                    v14 = (v14 - (fi17 >> ((280) & 63)))
                } else {
                    for fi18 := 0; fi18 < 11; fi18++ {
                        v19 := GenS2{F0: (v16 >> ((65535) & 63)), F1: v7, F2: (v16 >> ((0) & 63))}
                        if v15 {
                            var v20 uint32 = v5
                            _ = v20
                        } else {
                            v12 = append(v12, -429)
                            if (65535 != 0) {
                                var v21 float64 = 0.0
                                v3 -= v3
                                _ = v21
                            }
                            var v22 int64 = -9223372036854775808
                            switch v1 {
                            case "":
                                v14 = 128
                            }
                            _ = v22
                        }
                        _ = fi18
                        _ = v19
                    }
                    fmt.Println(v2)
                    v12 = append(v12, (v14 & (-(v14))))
                }
                var v23 int = -597
                _ = v23
            }
            v14 = v14
            var v24 int16 = (v2 >> ((4294967295) & 63))
            _ = fi17
            _ = v24
        }
        _ = v16
    }
    switch 164 {
    case -1:
        _ = len(v4)
    }
    var v25 uint = 4294967295
    var v26 int64 = 1
    if ((v25 % ((v25) | 1)) > v25) {
        fmt.Println(v7)
    }
    var v27 uint64 = v7
    v25 -= v25
    v12 = append(v12, v14)
    for fi28 := 0; fi28 < 5; fi28++ {
        v29 := GenS3{F0: v15, F1: v25}
        _ = fi28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
