package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 uint16; F2 uint8 }
type GenS3 struct { F0 uint16; F1 uint64; F2 int }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint32; F1 uint64; F2 uint16 }
func genFn0() int32 {
    v7 := GenS1{F0: 42}
    for fi8 := 0; fi8 < 13; fi8++ {
        if (float64(-2.25) != 0.0) {
            switch fi8 {
            case 32767:
                var v9 bool = (!(1 > 132))
                _ = v9
            case -162:
                for fi10 := 0; fi10 < 13; fi10++ {
                    fmt.Println(fi8)
                    _ = fi10
                }
            }
            if (359 >= -1) {
                v11 := GenS2{F0: 608, F1: 32767, F2: 136}
                _ = v11
            } else {
                fmt.Println(fi8)
                var v12 uint64 = 255
                _ = v12
            }
            var v13 int8 = 7
            _ = v13
        }
        var v14 int = 9223372036854775807
        for fi15 := 0; fi15 < 18; fi15++ {
            var v16 uint = 1
            var v17 int16 = 1
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = fi8
        _ = v14
    }
    _ = v7
    return 1
}

func main() {
    var v1 float64 = 1.5
    var v2 int8 = 0
    v3 := make(map[int]int)
    v4 := []uint32{2, 23, 4294967295}
    v5 := []string{}
    var v6 int32 = -41
    v18 := GenS2{F0: 68, F1: 255, F2: 1}
    v6 += (v6 & (^(v6)))
    for fi19 := 0; fi19 < 16; fi19++ {
        switch (((("" + "a") + ("" + "foo bar")) + "\n\t") + "foo bar") {
        case "foo bar":
            for fi20 := 0; fi20 < 14; fi20++ {
                if ((!((v2 + v2) == v2)) || false) {
                    v21 := GenS2{F0: 689, F1: 255, F2: 1}
                    _ = v21
                } else {
                    var v22 int8 = (-(v2))
                    v4 = append(v4, 2)
                    var v23 string = "hello"
                    _ = v22
                    _ = v23
                }
                v24 := GenS2{F0: 42, F1: 140, F2: 219}
                var v25 int32 = 42
                _ = fi20
                _ = v24
                _ = v25
            }
        case "hello":
            fmt.Println(v2, v1, v6)
        case "":
            for fi26 := 0; fi26 < 19; fi26++ {
                delete(v3, -1)
                _ = fi26
            }
        default:
            var v27 int32 = v6
            _ = v27
        }
        _ = fi19
    }
    v28 := GenS2{F0: 1, F1: 32767, F2: 187}
    v29 := GenS2{F0: 1000, F1: 833, F2: 1}
    switch "" {
    case "foo bar":
        v3[753] = 346
    case "hello":
        switch "\n\t" {
        case "hello":
            v30 := GenS1{F0: 909}
            _ = v30
        default:
            for fi31 := 0; fi31 < 11; fi31++ {
                for fi32 := 0; fi32 < 2; fi32++ {
                    var v33 int = fi31
                    var v34 float64 = v1
                    _ = fi32
                    _ = v33
                    _ = v34
                }
                var v35 uint16 = 65535
                _ = fi31
                _ = v35
            }
        }
    case "\n\t":
        v6 += (v6 % ((genFn0()) | 1))
    }
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v2)
    case "hello":
        fmt.Println(v1)
    case "foo bar":
        v6 -= (-(v6))
    default:
        v6 += (-(v6))
    }
    for fi36 := 0; fi36 < 2; fi36++ {
        v6 = 2147483647
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v37 bool = false
        _ = fi36
        _ = v37
    }
    _ = len(v3)
    v38 := GenS2{F0: 42, F1: 255, F2: 1}
    var v39 bool = (65535 >= 16)
    if false {
        var v40 float32 = -2.25
        v41 := GenS3{F0: 569, F1: 4294967295, F2: 1}
        if (!(42 < -1)) {
            switch 127 {
            case 128:
                switch 880 {
                case 9223372036854775807:
                    v3[353] = int(v6)
                case 32768:
                    _ = len(v3)
                }
            }
            var v42 uint32 = 2147483647
            v43 := GenS2{F0: 528, F1: 74, F2: 127}
            _ = v42
            _ = v43
        } else {
            var v44 int16 = 0
            _ = v44
        }
        _ = v40
        _ = v41
    } else {
        v45 := GenS4{F0: -1}
        var v46 uint16 = 490
        v47 := GenS5{F0: 2, F1: 3, F2: v46}
        _ = v45
        _ = v46
        _ = v47
    }
    v48 := GenS4{F0: 0}
    v49 := GenS2{F0: -32768, F1: 518, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
