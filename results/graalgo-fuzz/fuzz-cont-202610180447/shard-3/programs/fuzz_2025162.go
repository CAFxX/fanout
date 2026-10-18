package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool; F1 uint8; F2 uint32 }
func genFn0(p0 uint) uint {
    var v7 uint8 = 42
    p0 = 255
    _ = p0
    _ = v7
    return 18446744073709551615
}

func main() {
    var v1 uint8 = 144
    v2 := []int{-47, 82, -9223372036854775808}
    var v3 int64 = 42
    v4 := make(map[uint32]string)
    v5 := make(map[string]uint32)
    var v6 float64 = (0.0 + (1e10 / ((3.14159 / 1.5) * 0.1)))
    if false {
        switch 1 {
        case -9223372036854775808:
            for fi8 := 0; fi8 < 18; fi8++ {
                if false {
                    var v9 string = "hello"
                    delete(v4, 65535)
                    v2 = append(v2, -2147483648)
                    _ = v9
                }
                v4[uint32(v1)] = "hello"
                var v10 float64 = v6
                _ = fi8
                _ = v10
            }
        case 2147483647:
            v3 = v3
        case 32768:
            v3 -= (v3 * 608)
        }
        fmt.Println(v6, v1, v3)
        fmt.Println(v1, v1)
    } else {
        switch -60 {
        case -9223372036854775808:
            var v11 uint8 = v1
            _ = v11
        case 1:
            v12 := GenS1{F0: 28}
            _ = v12
        default:
            _ = len(v4)
        }
        v6 -= 3.14159
        v3 = v3
    }
    v6 -= v6
    v13 := GenS1{F0: 182}
    switch "\n\t" {
    case "hello":
        v6 = float64(1.5)
    case "foo bar":
        for fi14 := 0; fi14 < 13; fi14++ {
            v15 := GenS1{F0: 765}
            if ((v1 & 208) > v1) {
                fmt.Println(v6, v3, fi14)
                var v16 float32 = (float32((3.14159 - 3.14159)) / (0.1 - 0.0))
                _ = v16
            } else {
                var v17 float64 = v6
                for fi18 := 0; fi18 < 11; fi18++ {
                    var v19 int32 = -2147483648
                    switch ("\n\t" + "a") {
                    case "a":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case "foo bar":
                        var v20 float32 = 0.1
                        _ = v20
                    default:
                        v6 -= float64((3.14159 * 1e10))
                    }
                    _ = fi18
                    _ = v19
                }
                v1 = v1
                _ = v17
            }
            _ = fi14
            _ = v15
        }
    case "a":
        for fi21 := 0; fi21 < 20; fi21++ {
            v22 := GenS1{F0: -32768}
            _ = fi21
            _ = v22
        }
    default:
        if (2147483647 < -35) {
            if (!true) {
                var v23 uint16 = 1
                v24 := GenS2{F0: 32768}
                var v25 uint = 636
                _ = v23
                _ = v24
                _ = v25
            }
        } else {
            switch 42 {
            case -2147483648:
                for fi26 := 0; fi26 < 13; fi26++ {
                    var v27 uint = 18
                    _ = fi26
                    _ = v27
                }
            case -458:
                fmt.Println(v1, v3)
            case 42:
                switch "hello" {
                case "foo bar":
                    var v28 string = "\n\t"
                    _ = v28
                case "a":
                    for fi29 := 0; fi29 < 9; fi29++ {
                        _ = len(v2)
                        var v30 int = fi29
                        _ = fi29
                        _ = v30
                    }
                case "hello":
                    v31 := GenS3{F0: (false || (0 >= -37)), F1: v1, F2: 2147483647}
                    _ = v31
                }
            }
            v1 = v1
            v1 = (^(v1))
        }
    }
    v32 := GenS1{F0: -1}
    v3 = v3
    var v33 uint8 = (^(v1))
    switch 128 {
    case 42:
        var v34 int = -424
        _ = v34
    case -9223372036854775808:
        if false {
            v6 -= 1.5
        }
    case -2147483648:
        if true {
            for fi35 := 0; fi35 < 6; fi35++ {
                for fi36 := 0; fi36 < 8; fi36++ {
                    if false {
                        var v37 bool = true
                        v1 = (v33 >> ((255) & 63))
                        _ = v37
                    }
                    _ = fi36
                }
                v2 = append(v2, -2147483648)
                _ = fi35
            }
            switch "" {
            case "\n\t":
                var v38 float64 = 0.0
                _ = v38
            case "":
                var v39 int16 = -1
                _ = v39
            case "a":
                var v40 uint64 = 610
                _ = v40
            default:
                for fi41 := 0; fi41 < 5; fi41++ {
                    v33 = v1
                    var v42 float32 = ((1.5 - 1e10) + -2.25)
                    _ = fi41
                    _ = v42
                }
            }
        } else {
            for fi43 := 0; fi43 < 9; fi43++ {
                _ = len(v4)
                switch "\n\t" {
                case "foo bar":
                    v33 += v33
                case "":
                    fmt.Println(fi43, v6)
                }
                v44 := GenS3{F0: (true && false), F1: v1, F2: 4294967295}
                _ = fi43
                _ = v44
            }
            for fi45 := 0; fi45 < 9; fi45++ {
                for fi46 := 0; fi46 < 13; fi46++ {
                    var v47 float32 = 1.5
                    var v48 int32 = 2147483647
                    _ = fi46
                    _ = v47
                    _ = v48
                }
                _ = fi45
            }
            var v49 int16 = 0
            _ = v49
        }
    }
    var v50 string = ("hello" + "a")
    v6 += v6
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v50
    _ = v6
}
