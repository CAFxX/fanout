package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 int32 }
type GenS2 struct { F0 int; F1 uint8; F2 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16; F1 uint16 }
type GenS5 struct { F0 uint8; F1 uint64; F2 uint16 }
func genFn0(p0 bool, p1 uint) uint {
    if (!(!false)) {
        v6 := GenS1{F0: 1, F1: 2, F2: -1}
        _ = v6
    }
    var v7 uint = 65535
    fmt.Println(p1, p1)
    for fi8 := 0; fi8 < 20; fi8++ {
        switch 1000 {
        case 9223372036854775807:
            v9 := []int{}
            _ = v9
        case 0:
            switch (fi8 - (fi8 + fi8)) {
            case 2147483647:
                var v10 int8 = 78
                _ = v10
            case 9223372036854775807:
                var v11 int8 = 100
                _ = v11
            case 32767:
                p1 = 47
            }
        }
        v12 := GenS1{F0: 795, F1: 127, F2: 709}
        if true {
            var v13 int16 = 0
            switch fi8 {
            case 32767:
                for fi14 := 0; fi14 < 1; fi14++ {
                    v15 := []int{(fi8 - fi14), (fi14 | fi14), fi8, (fi8 * (fi8 % ((fi14) | 1)))}
                    _ = fi14
                    _ = v15
                }
            case 32768:
                v16 := GenS2{F0: fi8, F1: 55, F2: (-(fi8))}
                _ = v16
            case 94:
                v17 := GenS3{F0: 0}
                _ = v17
            default:
                for fi18 := 0; fi18 < 14; fi18++ {
                    var v19 string = ("a" + "hello")
                    v19 = ""
                    v20 := []int{127, 42}
                    _ = fi18
                    _ = v19
                    _ = v20
                }
            }
            v21 := GenS3{F0: 746}
            _ = v13
            _ = v21
        }
        _ = fi8
        _ = v12
    }
    if p0 {
        p0 = false
        var v22 uint8 = 128
        p1 = 1
        _ = v22
    } else {
        fmt.Println(p1, p0, p1)
    }
    _ = p0
    _ = p1
    _ = v7
    return 525
}
func genFn1(p0 uint8, p1 bool) uint32 {
    p0 = p0
    var v23 uint16 = 701
    v23 = v23
    _ = p0
    _ = p1
    _ = v23
    return 42
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]bool)
    var v3 int16 = -32768
    v4 := make(map[uint32]string)
    var v5 bool = (!(4294967295 != 827))
    fmt.Println(v3, v3)
    var v24 uint8 = 255
    v4[genFn1(1, (-2147483648 <= 969))] = ("foo bar" + ((("" + "a") + "\n\t") + "a"))
    for fi25 := 0; fi25 < 10; fi25++ {
        for fi26 := 0; fi26 < 7; fi26++ {
            v27 := GenS1{F0: 2147483647, F1: v24, F2: 2147483647}
            _ = fi26
            _ = v27
        }
        v28 := []int{-671, -9223372036854775808, fi25, 2147483647}
        _ = fi25
        _ = v28
    }
    var v29 int16 = (v3 ^ v3)
    v5 = (0 == 2)
    v30 := []int{32768, 417}
    var v31 string = "foo bar"
    fmt.Println(v24, v3, v24)
    _ = len(v30)
    fmt.Println(v31, v5, v5)
    var v32 uint16 = 255
    switch "\n\t" {
    case "":
        if v5 {
            switch -9223372036854775808 {
            case 9223372036854775807:
                switch (v31 + "a") {
                case "a":
                    v29 += (v29 ^ (v3 + -1))
                case "foo bar":
                    v33 := GenS4{F0: (v32 * v32), F1: (v32 + (v32 << ((255) & 63)))}
                    _ = v33
                default:
                    v34 := GenS4{F0: 65535, F1: v32}
                    _ = v34
                }
            case -548:
                if false {
                    v3 -= -1
                    fmt.Println(v32, v29)
                }
            case -1:
                v35 := GenS4{F0: (v32 & 32767), F1: v32}
                _ = v35
            }
            v3 = 32767
        } else {
            _ = len(v4)
            var v36 int32 = 871
            v29 += 42
            _ = v36
        }
    case "\n\t":
        if true {
            v31 = "foo bar"
            var v37 uint = 0
            v3 -= (v3 >> ((v37) & 63))
            _ = v37
        }
    }
    if len(v30) > 0 {
        v30[0] = 9223372036854775807
    }
    if true {
        fmt.Println(v32)
    } else {
        var v38 int8 = 100
        _ = len(v1)
        fmt.Println(v5, v24, v5)
        _ = v38
    }
    if (1 == 42) {
        v31 = v31
    }
    v39 := GenS5{F0: v24, F1: 2, F2: (^(v32))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v39
    _ = v4
    _ = v5
}
