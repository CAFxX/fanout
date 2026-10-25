package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 uint32; F1 string; F2 uint64 }
func genFn0() uint32 {
    var v7 int64 = -245
    var v8 int8 = 33
    fmt.Println(v7)
    _ = v7
    _ = v8
    return 247
}
func genFn1(p0 uint, p1 string) bool {
    v9 := []int{-15, 128, 127, 32768}
    _ = len(v9)
    for fi10 := 0; fi10 < 5; fi10++ {
        v11 := GenS1{F0: 2}
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v9
    return (!true)
}

func main() {
    var v1 int8 = -128
    v2 := make(map[uint32]bool)
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]bool)
    var v5 uint16 = 256
    v6 := []uint32{}
    for fi12 := 0; fi12 < 11; fi12++ {
        var v13 int8 = (v1 - v1)
        if false {
            v14 := GenS2{F0: 0, F1: 2}
            switch "foo bar" {
            case "foo bar":
                switch (fi12 / (((fi12 | fi12)) | 1)) {
                case 2147483647:
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                case 127:
                    v15 := v3["a"]
                    _ = v15
                case -9223372036854775808:
                    v13 = v1
                }
            case "\n\t":
                _ = len(v3)
            }
            var v16 string = "a"
            _ = v14
            _ = v16
        }
        var v17 bool = (0.0 > 0.0)
        _ = fi12
        _ = v13
        _ = v17
    }
    v18 := v3[((("hello" + "foo bar") + "a") + ("a" + "\n\t"))]
    if ((v1 >> ((559) & 63)) < v1) {
        if false {
            if false {
                fmt.Println(v18, v1)
            }
        }
        switch 1 {
        case -216:
            var v19 int16 = 32767
            _ = v19
        case -186:
            for fi20 := 0; fi20 < 12; fi20++ {
                var v21 uint16 = (v5 << ((454) & 63))
                var v22 int8 = v1
                _ = fi20
                _ = v21
                _ = v22
            }
        }
    } else {
        if false {
            v5 += (^(v5))
            switch int(v18) {
            case 32768:
                if ((1 != -9223372036854775808) || (2 < 255)) {
                    switch -2147483648 {
                    case 42:
                        v23 := GenS1{F0: 381}
                        _ = v23
                    case 32767:
                        var v24 float32 = (1.5 - 0.0)
                        _ = v24
                    case -1:
                        fmt.Println(v18, v5)
                    }
                }
            default:
                var v25 int = -348
                _ = v25
            }
        }
        v18 = 100
        v26 := GenS3{F0: (^(v18)), F1: "\n\t", F2: 1}
        _ = v26
    }
    if (((false && false) && (42 <= 1)) && false) {
        v6 = append(v6, v18)
    } else {
        if len(v6) > 0 {
            v6[0] = genFn0()
        }
        v27 := GenS3{F0: (v18 * 701), F1: "a", F2: 195}
        for fi28 := 0; fi28 < 13; fi28++ {
            for fi29 := 0; fi29 < 8; fi29++ {
                fmt.Println(v5)
                v6 = append(v6, v18)
                var v30 int16 = 229
                _ = fi29
                _ = v30
            }
            fmt.Println(v1, v1)
            fmt.Println(v5, v1, v1)
            _ = fi28
        }
        _ = v27
    }
    fmt.Println(v18, v1, v5)
    v6 = append(v6, v18)
    var v31 string = "a"
    for fi32 := 0; fi32 < 15; fi32++ {
        v18 = (v18 | v18)
        var v33 int64 = 518
        var v34 uint16 = (^(v5))
        _ = fi32
        _ = v33
        _ = v34
    }
    var v35 int64 = 79
    v6 = append(v6, 65535)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v31
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
