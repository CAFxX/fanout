package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32; F1 uint; F2 int }
type GenS4 struct { F0 uint64; F1 uint16; F2 uint8 }
type GenS5 struct { F0 uint16 }
func genFn0(p0 int8, p1 int) bool {
    fmt.Println(p0, p0)
    v5 := GenS1{F0: 255, F1: 42}
    _ = p0
    _ = p1
    _ = v5
    return true
}

func main() {
    var v1 bool = true
    v2 := []uint32{4294967295, 2147483647, 786}
    var v3 int32 = -445
    v4 := make(map[uint32]bool)
    for fi6 := 0; fi6 < 15; fi6++ {
        for fi7 := 0; fi7 < 4; fi7++ {
            switch fi6 {
            case -1:
                fmt.Println(fi6, fi7)
            case -2147483648:
                v3 -= 1
            case 2147483647:
                var v8 string = "\n\t"
                _ = v8
            }
            fmt.Println(v3, fi6)
            switch "a" {
            case "a":
                var v9 uint64 = 42
                _ = v9
            }
            _ = fi7
        }
        _ = fi6
    }
    switch "a" {
    case "":
        fmt.Println(v3, v1)
    case "a":
        for fi10 := 0; fi10 < 4; fi10++ {
            v3 -= (-(v3))
            if (!((v1 || ((-9223372036854775808 <= 632) || ("hello" != ""))) && v1)) {
                for fi11 := 0; fi11 < 17; fi11++ {
                    if (128 >= 128) {
                        v12 := GenS1{F0: 869, F1: -9223372036854775808}
                        _ = v12
                    }
                    switch fi10 {
                    case 128:
                        v3 = 42
                    }
                    switch "\n\t" {
                    case "":
                        var v13 uint16 = 2
                        _ = v13
                    case "a":
                        switch 0 {
                        case 2147483647:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        case 0:
                            var v14 bool = v1
                            _ = v14
                        default:
                            switch "a" {
                            case "foo bar":
                                var v15 int32 = (v3 + (v3 + (v3 + -1)))
                                _ = v15
                            case "":
                                fmt.Println(v1, v1)
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = len(v2)
                            case "hello":
                                v3 -= (-(v3))
                            default:
                                var v16 string = ""
                                _ = v16
                            }
                        }
                    default:
                        v17 := GenS2{F0: 32767}
                        _ = v17
                    }
                    var v18 uint = 667
                    _ = fi11
                    _ = v18
                }
            }
            if true {
                switch (fi10 >> ((923) & 63)) {
                case -2147483648:
                    for fi19 := 0; fi19 < 9; fi19++ {
                        v3 = v3
                        var v20 uint = 911
                        var v21 bool = ((!false) && (int16(fi10) <= 292))
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                default:
                    v22 := GenS2{F0: -32768}
                    _ = v22
                }
                var v23 float32 = 0.1
                _ = v23
            }
            _ = fi10
        }
    case "foo bar":
        v24 := GenS1{F0: 42, F1: -880}
        _ = v24
    }
    v25 := GenS3{F0: 1, F1: 0, F2: 42}
    v3 = (v3 - v3)
    var v26 bool = (-254 > 0)
    for fi27 := 0; fi27 < 6; fi27++ {
        v3 = v3
        switch (fi27 | (fi27 & -9223372036854775808)) {
        case 128:
            v28 := GenS4{F0: 4294967295, F1: 42, F2: 80}
            _ = v28
        }
        switch (fi27 * (fi27 * -535)) {
        case 127:
            var v29 string = ("" + "hello")
            _ = v29
        case 4294967295:
            v30 := GenS5{F0: 2}
            _ = v30
        case 2147483647:
            v4[47] = (((float64(0.1) - 3.14159) / (3.14159 * 1e10)) >= ((3.14159 + -2.25) - 1.5))
        }
        _ = fi27
    }
    v26 = (v1 || v26)
    if true {
        _ = len(v2)
        for fi31 := 0; fi31 < 14; fi31++ {
            v32 := GenS5{F0: 255}
            _ = fi31
            _ = v32
        }
        var v33 uint = 1
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
