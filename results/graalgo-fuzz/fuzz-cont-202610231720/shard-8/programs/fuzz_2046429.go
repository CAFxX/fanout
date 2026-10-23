package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 uint32 }

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{("" + "a")}
    v3 := []string{"", "\n\t"}
    var v4 uint8 = 56
    v5 := []uint32{233}
    for fi6 := 0; fi6 < 16; fi6++ {
        for fi7 := 0; fi7 < 16; fi7++ {
            v8 := GenS1{F0: 65535, F1: (v4 - 42)}
            var v9 uint64 = 970
            fmt.Println(fi7, fi6)
            var v10 int16 = 42
            _ = len(v1)
            var v11 uint8 = (v4 + v4)
            _ = fi7
            _ = v10
            _ = v11
            _ = v8
            _ = v9
        }
        v4 -= (v4 + 255)
        _ = fi6
    }
    if len(v2) > 0 {
        v2[0] = ""
    }
    fmt.Println(v4)
    switch "" {
    case "\n\t":
        v4 = (v4 & (v4 * (v4 % (((v4 / ((v4) | 1))) | 1))))
    case "hello":
        fmt.Println(v4)
    case "a":
        var v12 int = 9223372036854775807
        _ = v12
    default:
        for fi13 := 0; fi13 < 19; fi13++ {
            v4 = v4
            var v14 uint16 = uint16(v4)
            var v15 bool = (702 >= 828)
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 8; fi16++ {
        v1[42] = 65535
        for fi17 := 0; fi17 < 8; fi17++ {
            if (int16(v4) == 850) {
                v18 := GenS1{F0: 32767, F1: (v4 * (v4 - (^(v4))))}
                _ = v18
            }
            _ = fi17
        }
        for fi19 := 0; fi19 < 7; fi19++ {
            v20 := GenS1{F0: 0, F1: v4}
            _ = fi19
            _ = v20
        }
        _ = fi16
    }
    _ = len(v1)
    v21 := GenS1{F0: 0, F1: v4}
    for fi22 := 0; fi22 < 5; fi22++ {
        for fi23 := 0; fi23 < 10; fi23++ {
            var v24 uint = 15
            var v25 int64 = 1
            _ = fi23
            _ = v24
            _ = v25
        }
        if (347 >= 764) {
            for fi26 := 0; fi26 < 8; fi26++ {
                v4 = v4
                _ = fi26
            }
        } else {
            fmt.Println(fi22, v4)
            var v27 string = "hello"
            fmt.Println(v4, v4, v4)
            _ = v27
        }
        _ = fi22
    }
    switch -325 {
    case -912:
        fmt.Println(v4)
    default:
        for fi28 := 0; fi28 < 6; fi28++ {
            switch "hello" {
            case "a":
                var v29 float64 = 0.1
                _ = v29
            default:
                for fi30 := 0; fi30 < 14; fi30++ {
                    fmt.Println(fi30, fi30)
                    var v31 int64 = -1
                    var v32 uint8 = v4
                    fmt.Println(v32, v31, v4)
                    _ = fi30
                    _ = v31
                    _ = v32
                }
            }
            _ = len(v1)
            var v33 string = ""
            _ = fi28
            _ = v33
        }
    }
    var v34 string = "foo bar"
    for fi35 := 0; fi35 < 12; fi35++ {
        var v36 uint32 = 282
        _ = fi35
        _ = v36
    }
    if (100 < 127) {
        var v37 uint = 2
        for fi38 := 0; fi38 < 7; fi38++ {
            if ((3.14159 <= 3.14159) && false) {
                if ((v34 < "") && (false || false)) {
                    v34 = "hello"
                }
                v37 -= uint(v4)
                var v39 int8 = -1
                _ = v39
            }
            _ = fi38
        }
        _ = v37
    } else {
        for fi40 := 0; fi40 < 18; fi40++ {
            v34 = ((v34 + "hello") + v34)
            switch "\n\t" {
            case "\n\t":
                v34 = ((v34 + "") + ((("a" + "hello") + v34) + v34))
            case "foo bar":
                v41 := GenS1{F0: 738, F1: (^(v4))}
                _ = v41
            case "a":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            default:
                fmt.Println(fi40, v4, fi40)
            }
            v42 := GenS1{F0: 16, F1: 0}
            _ = fi40
            _ = v42
        }
    }
    switch -530 {
    case 2147483647:
        fmt.Println(v4, v4)
    default:
        for fi43 := 0; fi43 < 12; fi43++ {
            v44 := GenS2{F0: 1}
            _ = fi43
            _ = v44
        }
    }
    var v45 uint32 = 0
    v34 = "a"
    v46 := GenS1{F0: 1, F1: v4}
    switch -1 {
    case 4294967295:
        var v47 int16 = -448
        _ = v47
    case -478:
        for fi48 := 0; fi48 < 1; fi48++ {
            v4 = 37
            v49 := GenS1{F0: 255, F1: (v4 & (v4 ^ 169))}
            v45 = v45
            _ = fi48
            _ = v49
        }
    case 128:
        for fi50 := 0; fi50 < 12; fi50++ {
            v51 := GenS3{F0: 2, F1: v45}
            v52 := v1[(fi50 - (fi50 / ((fi50) | 1)))]
            _ = fi50
            _ = v51
            _ = v52
        }
    default:
        v34 = v34
    }
    v4 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v34
    _ = v4
    _ = v45
    _ = v46
    _ = v5
}
