package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 uint; F1 uint32 }
func genFn0(p0 int) string {
    fmt.Println(p0)
    var v5 int8 = -1
    for fi6 := 0; fi6 < 14; fi6++ {
        v7 := GenS1{F0: fi6, F1: 609, F2: 1}
        var v8 int = p0
        for fi9 := 0; fi9 < 7; fi9++ {
            v10 := GenS1{F0: (v8 & v8), F1: 693, F2: 65535}
            v11 := GenS1{F0: p0, F1: 79, F2: 816}
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi12 := 0; fi12 < 4; fi12++ {
        v13 := []int{}
        _ = len(v13)
        if (!((-565 == 9223372036854775807) && ("\n\t" != "a"))) {
            fmt.Println(fi12, v5, v5)
        } else {
            for fi14 := 0; fi14 < 2; fi14++ {
                var v15 uint = 1
                _ = fi14
                _ = v15
            }
            v5 = v5
        }
        _ = fi12
        _ = v13
    }
    switch p0 {
    case 128:
        p0 = p0
    case 0:
        v16 := GenS1{F0: p0, F1: 255, F2: 42}
        _ = v16
    case 32767:
        v5 = (v5 & v5)
    default:
        switch "" {
        case "":
            var v17 uint16 = 337
            _ = v17
        case "hello":
            fmt.Println(v5, p0)
        case "a":
            v18 := []int{830, (p0 ^ p0), p0, 2147483647}
            _ = v18
        default:
            var v19 uint16 = 503
            _ = v19
        }
    }
    _ = p0
    _ = v5
    return "foo bar"
}

func main() {
    var v1 uint32 = 42
    var v2 float32 = 0.1
    var v3 uint8 = 128
    var v4 int = 2147483647
    v4 -= (v4 % (((v4 & 180)) | 1))
    v2 += float32(1.5)
    v20 := GenS2{F0: 42}
    v21 := GenS3{F0: 42, F1: -100}
    v22 := GenS2{F0: 42}
    v2 -= v2
    switch "a" {
    case "hello":
        var v23 int32 = 0
        _ = v23
    case "foo bar":
        switch -395 {
        case 1:
            switch ("foo bar" + "hello") {
            case "\n\t":
                if ("" != "\n\t") {
                    switch "" {
                    case "":
                        var v24 uint16 = 255
                        _ = v24
                    case "hello":
                        v25 := []int{v4, (v4 % ((v4) | 1))}
                        _ = v25
                    }
                    v1 -= v1
                    v26 := GenS3{F0: 1, F1: -1}
                    _ = v26
                } else {
                    var v27 uint = 2
                    _ = v27
                }
            case "foo bar":
                for fi28 := 0; fi28 < 9; fi28++ {
                    v4 += (v4 / (((^(v4))) | 1))
                    _ = fi28
                }
            default:
                var v29 string = "\n\t"
                _ = v29
            }
        }
    }
    for fi30 := 0; fi30 < 11; fi30++ {
        if (42 > 80) {
            var v31 int8 = 100
            v1 = 945
            _ = v31
        } else {
            if true {
                v32 := GenS4{F0: 42, F1: (^(v1))}
                v33 := []int{(^(fi30)), (fi30 % ((-61) | 1)), 127}
                _ = v32
                _ = v33
            } else {
                switch "a" {
                case "":
                    if true {
                        for fi34 := 0; fi34 < 3; fi34++ {
                            var v35 bool = (32767 >= 42)
                            fmt.Println(v1)
                            _ = fi34
                            _ = v35
                        }
                        if (!(914 >= 87)) {
                            var v36 float32 = v2
                            _ = v36
                        }
                    } else {
                        fmt.Println(v4)
                        var v37 int64 = 9223372036854775807
                        _ = v37
                    }
                default:
                    var v38 uint = 255
                    _ = v38
                }
                v39 := []int{fi30}
                _ = v39
            }
        }
        var v40 uint64 = 18446744073709551615
        v2 = (((-2.25 + -2.25) + (v2 / (v2 - 0.1))) / (((1.5 - -2.25) - (1e10 * 3.14159)) / (v2 * 1e10)))
        _ = fi30
        _ = v40
    }
    v41 := []int{v4, (v4 << ((65535) & 63)), v4, v4}
    v42 := GenS1{F0: int(v3), F1: 18446744073709551615, F2: 65535}
    switch "foo bar" {
    case "a":
        if ((((v4 % ((32768) | 1)) != (-(v4))) && false) || (1000 <= 938)) {
            switch v4 {
            case -413:
                for fi43 := 0; fi43 < 17; fi43++ {
                    var v44 uint = 1
                    _ = fi43
                    _ = v44
                }
            case 2147483647:
                v3 += (^(v3))
            case 945:
                _ = len(v41)
            default:
                var v45 int8 = 0
                _ = v45
            }
            for fi46 := 0; fi46 < 13; fi46++ {
                v41 = append(v41, (v4 >> ((2) & 63)))
                _ = fi46
            }
            v3 += (^(v3))
        }
    case "":
        for fi47 := 0; fi47 < 17; fi47++ {
            fmt.Println(v1, v4)
            fmt.Println(v4)
            _ = fi47
        }
    default:
        var v48 int32 = 2147483647
        _ = v48
    }
    var v49 uint = 2
    v50 := GenS1{F0: v4, F1: (v49 - v49), F2: v49}
    fmt.Println(v3, v1)
    v4 = v4
    var v51 uint8 = (v3 >> ((v49) & 63))
    v49 = v49
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v49:", v49)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v49
    _ = v50
    _ = v51
}
