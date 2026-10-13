package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 int16; F2 uint8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint8; F1 string; F2 bool }
func genFn0() bool {
    var v5 float64 = 1e10
    var v6 uint64 = 1
    if true {
        fmt.Println(v5)
        fmt.Println(v5, v6)
        for fi7 := 0; fi7 < 2; fi7++ {
            var v8 uint = 0
            v9 := GenS1{F0: false}
            for fi10 := 0; fi10 < 6; fi10++ {
                switch (fi7 ^ (fi7 & fi7)) {
                case 669:
                    if (((1e10 / float32(((((v5 / v5) - v5) - -2.25) + 1.5))) + ((float32(0.1) - 0.0) + (float32(v5) / 1.0))) != 0.1) {
                        v6 = 1
                        fmt.Println(fi10)
                        var v11 float32 = (1.5 + 0.1)
                        _ = v11
                    } else {
                        if false {
                            v12 := GenS1{F0: true}
                            _ = v12
                        }
                        v5 = float64(1e10)
                        var v13 int16 = -1
                        _ = v13
                    }
                case 489:
                    v14 := GenS2{F0: (true && false)}
                    _ = v14
                case 808:
                    var v15 float32 = float32(3.14159)
                    _ = v15
                default:
                    v16 := GenS1{F0: ((128 <= 1) && (257 > 0))}
                    _ = v16
                }
                switch (^(fi7)) {
                case -840:
                    v6 = (v6 * (v6 / (((^(v6))) | 1)))
                    v5 = v5
                case 606:
                    var v17 uint8 = 255
                    _ = v17
                case 0:
                    v6 = v6
                default:
                    if false {
                        v18 := GenS3{F0: -17, F1: 0, F2: 42}
                        v6 -= (^(v6))
                        _ = v18
                    }
                }
                _ = fi10
            }
            _ = fi7
            _ = v8
            _ = v9
        }
    } else {
        if true {
            fmt.Println(v6)
        } else {
            var v19 int16 = -1
            _ = v19
        }
        var v20 uint64 = 18446744073709551615
        if (true && false) {
            switch ("hello" + "") {
            case "hello":
                v21 := GenS3{F0: 0, F1: 1, F2: 251}
                _ = v21
            case "\n\t":
                v20 = v20
            case "a":
                v6 -= (v6 >> ((255) & 63))
            }
            for fi22 := 0; fi22 < 17; fi22++ {
                v6 += v6
                v23 := GenS4{F0: (v20 | (^(v6)))}
                v20 -= (v6 >> ((42) & 63))
                _ = fi22
                _ = v23
            }
        }
        _ = v20
    }
    _ = v5
    _ = v6
    return (!(((-12 < 32767) || false) && (634 > 561)))
}
func genFn1(p0 int8) int64 {
    fmt.Println(p0)
    if (1.5 != (float32((0.0 / float64(-2.25))) - (0.0 * 0.0))) {
        v24 := []int{-2147483648, 42, -36}
        _ = v24
    }
    p0 = 0
    v25 := []int{-9223372036854775808, 128, 32767}
    for fi26 := 0; fi26 < 12; fi26++ {
        if true {
            v27 := GenS1{F0: true}
            p0 -= (p0 - -128)
            if true {
                if len(v25) > 0 {
                    v25[0] = -701
                }
                p0 += 42
            } else {
                p0 += 83
                fmt.Println(fi26, fi26)
            }
            _ = v27
        }
        _ = fi26
    }
    _ = p0
    _ = v25
    return 9223372036854775807
}

func main() {
    var v1 uint16 = 2
    v2 := make(map[string]int)
    var v3 int16 = -32768
    var v4 string = "foo bar"
    for fi28 := 0; fi28 < 12; fi28++ {
        switch "a" {
        case "foo bar":
            v29 := GenS1{F0: false}
            _ = v29
        case "":
            v3 -= v3
        }
        _ = fi28
    }
    v4 = v4
    if true {
        fmt.Println(v3, v3)
        for fi30 := 0; fi30 < 10; fi30++ {
            var v31 string = ("hello" + v4)
            switch fi30 {
            case 9223372036854775807:
                for fi32 := 0; fi32 < 19; fi32++ {
                    v31 = ((((((v4 + ((v31 + ((v4 + v31) + "hello")) + v4)) + v4) + v31) + (v31 + v31)) + v4) + "")
                    fmt.Println(v1, fi30, v4)
                    _ = fi32
                }
            case -1:
                v31 = v4
            }
            for fi33 := 0; fi33 < 17; fi33++ {
                if (!false) {
                    v3 -= (v3 ^ v3)
                    v31 = ((v31 + v31) + (v4 + "\n\t"))
                } else {
                    var v34 uint16 = 950
                    var v35 int8 = 0
                    var v36 int = (fi30 - (fi30 * (fi33 - fi33)))
                    _ = v34
                    _ = v35
                    _ = v36
                }
                fmt.Println(fi33, fi33, v31)
                _ = fi33
            }
            _ = fi30
            _ = v31
        }
    }
    v1 = v1
    for fi37 := 0; fi37 < 17; fi37++ {
        var v38 uint64 = 458
        switch "" {
        case "a":
            var v39 string = "\n\t"
            _ = v39
        default:
            v40 := []int{-742}
            _ = v40
        }
        _ = fi37
        _ = v38
    }
    var v41 int8 = 42
    if false {
        var v42 int16 = -858
        _ = v42
    }
    if (1.5 != 0.1) {
        v4 = "a"
        var v43 int64 = -9223372036854775808
        if true {
            v43 = (v43 * (v43 & genFn1(v41)))
            var v44 uint8 = 0
            _ = v44
        }
        _ = v43
    } else {
        v45 := GenS5{F0: 3, F1: (((v4 + v4) + v4) + (v4 + "hello")), F2: (true || false)}
        _ = v45
    }
    for fi46 := 0; fi46 < 12; fi46++ {
        v47 := GenS4{F0: 806}
        _ = fi46
        _ = v47
    }
    for fi48 := 0; fi48 < 2; fi48++ {
        v4 = ("hello" + "a")
        _ = fi48
    }
    var v49 float32 = (0.1 + ((((float32(1e10) * 1.5) + 1.5) + float32(0.1)) - (-2.25 / -2.25)))
    v4 = ""
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v49
}
