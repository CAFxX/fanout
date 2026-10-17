package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 string }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]int)
    v3 := make(map[uint32]bool)
    var v4 int32 = 2147483647
    var v5 float32 = (((3.14159 * 1.5) * 0.0) - 0.1)
    switch -2147483648 {
    case 404:
        for fi6 := 0; fi6 < 6; fi6++ {
            fmt.Println(v4, v4, fi6)
            for fi7 := 0; fi7 < 1; fi7++ {
                fmt.Println(fi7, v5, v4)
                v4 -= (-(v4))
                _ = fi7
            }
            for fi8 := 0; fi8 < 8; fi8++ {
                fmt.Println(v5)
                var v9 int8 = 42
                if false {
                    v10 := GenS1{F0: v9}
                    _ = v10
                } else {
                    var v11 float64 = 1e10
                    _ = v11
                }
                _ = fi8
                _ = v9
            }
            _ = fi6
        }
    case -9223372036854775808:
        var v12 int16 = 1000
        _ = v12
    }
    v13 := GenS2{F0: 2, F1: "hello"}
    for fi14 := 0; fi14 < 7; fi14++ {
        var v15 int = fi14
        var v16 uint = 632
        var v17 uint16 = 65535
        _ = fi14
        _ = v15
        _ = v16
        _ = v17
    }
    if true {
        var v18 int16 = -599
        _ = v18
    } else {
        switch -129 {
        case 32767:
            v4 += v4
        default:
            var v19 bool = false
            _ = v19
        }
        var v20 int64 = 0
        switch ("" + (("hello" + ("foo bar" + "")) + (("a" + "hello") + "a"))) {
        case "foo bar":
            v21 := []int{}
            _ = v21
        case "hello":
            for fi22 := 0; fi22 < 20; fi22++ {
                switch "" {
                case "hello":
                    var v23 int32 = v4
                    _ = v23
                default:
                    var v24 int32 = -397
                    _ = v24
                }
                if false {
                    var v25 uint64 = 0
                    v4 = (v4 ^ (v4 % ((v4) | 1)))
                    v26 := []int{fi22, -2147483648, fi22}
                    _ = v25
                    _ = v26
                } else {
                    var v27 int8 = 42
                    var v28 int8 = v27
                    _ = v27
                    _ = v28
                }
                _ = fi22
            }
        }
        _ = v20
    }
    if false {
        var v29 uint8 = 128
        _ = v29
    } else {
        v4 = v4
        v30 := GenS1{F0: 0}
        var v31 uint64 = 0
        _ = v30
        _ = v31
    }
    v4 -= v4
    for fi32 := 0; fi32 < 4; fi32++ {
        v33 := []int{(fi32 - fi32), fi32}
        v4 += v4
        _ = fi32
        _ = v33
    }
    if false {
        var v34 int = 42
        for fi35 := 0; fi35 < 11; fi35++ {
            v3[946] = (v4 != -1)
            if false {
                switch -365 {
                case -531:
                    delete(v2, v34)
                }
                fmt.Println(v4, v4, v34)
            } else {
                fmt.Println(v5, v5, v34)
                var v36 uint8 = 127
                _ = v36
            }
            _ = fi35
        }
        fmt.Println(v34)
        _ = v34
    } else {
        var v37 float64 = float64(0.0)
        _ = v37
    }
    var v38 string = "a"
    var v39 int64 = 530
    if ((v39 | v39) < (v39 | (v39 / (((^(v39))) | 1)))) {
        var v40 string = "foo bar"
        for fi41 := 0; fi41 < 15; fi41++ {
            v5 = (0.1 + (v5 / (3.14159 / (float32(((((3.14159 * -2.25) - 0.1) / -2.25) + 1.5)) - 3.14159))))
            v39 -= (v39 >> ((1) & 63))
            var v42 string = ("\n\t" + "a")
            _ = fi41
            _ = v42
        }
        for fi43 := 0; fi43 < 6; fi43++ {
            v4 = (v4 * v4)
            _ = fi43
        }
        _ = v40
    }
    for fi44 := 0; fi44 < 13; fi44++ {
        var v45 uint8 = 2
        _ = fi44
        _ = v45
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
