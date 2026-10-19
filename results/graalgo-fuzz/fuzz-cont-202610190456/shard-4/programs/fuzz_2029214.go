package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 bool; F1 int64 }
func genFn0(p0 uint64, p1 int32) int32 {
    v6 := []int{-617, 128, 32768, -1}
    for fi7 := 0; fi7 < 3; fi7++ {
        fmt.Println(p0, p1)
        switch "" {
        case "foo bar":
            if false {
                _ = len(v6)
            }
        case "\n\t":
            var v8 uint16 = 0
            _ = v8
        default:
            if true {
                for fi9 := 0; fi9 < 16; fi9++ {
                    v6 = append(v6, fi7)
                    var v10 string = ("\n\t" + ("hello" + "foo bar"))
                    _ = fi9
                    _ = v10
                }
                for fi11 := 0; fi11 < 18; fi11++ {
                    var v12 int16 = 32767
                    var v13 string = ("foo bar" + ("hello" + ("foo bar" + "foo bar")))
                    fmt.Println(fi7, p1)
                    if (fi7 == 128) {
                        v14 := GenS1{F0: 127}
                        _ = v14
                    } else {
                        var v15 float32 = 0.0
                        _ = v15
                    }
                    v16 := GenS1{F0: 2}
                    _ = fi11
                    _ = v12
                    _ = v13
                    _ = v16
                }
                if false {
                    fmt.Println(fi7, p0, p1)
                }
            }
        }
        v17 := GenS1{F0: 127}
        _ = fi7
        _ = v17
    }
    var v18 bool = ((p0 / ((p0) | 1)) != p0)
    _ = p0
    _ = p1
    _ = v18
    _ = v6
    return (p1 & p1)
}
func genFn1(p0 int) int8 {
    if false {
        v19 := GenS2{F0: 155, F1: 2147483647}
        if (0.1 < 0.1) {
            var v20 uint32 = 1
            v21 := []int{-914, 1}
            _ = v20
            _ = v21
        }
        _ = v19
    } else {
        v22 := []int{(p0 | p0)}
        _ = v22
    }
    for fi23 := 0; fi23 < 8; fi23++ {
        p0 = fi23
        _ = fi23
    }
    var v24 uint16 = 65535
    _ = p0
    _ = v24
    return 35
}

func main() {
    var v1 int8 = -128
    v2 := make(map[string]bool)
    var v3 uint64 = 541
    v4 := make(map[uint32]uint32)
    v5 := []int{}
    for fi25 := 0; fi25 < 11; fi25++ {
        for fi26 := 0; fi26 < 1; fi26++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            fmt.Println(fi25, v3)
            _ = fi26
        }
        v1 += genFn1(fi25)
        v1 += (v1 ^ (-(v1)))
        _ = fi25
    }
    var v27 string = "\n\t"
    for fi28 := 0; fi28 < 19; fi28++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v29 := GenS2{F0: 127, F1: 472}
        for fi30 := 0; fi30 < 16; fi30++ {
            fmt.Println(fi28, fi30, v27)
            if true {
                v31 := GenS3{F0: (0 > 65535), F1: -26}
                v27 = v27
                switch "foo bar" {
                case "a":
                    if (868 != 650) {
                        var v32 uint64 = (v3 >> ((4294967295) & 63))
                        delete(v4, 2147483647)
                        v27 = v27
                        _ = v32
                    } else {
                        if len(v5) > 0 {
                            v5[0] = fi28
                        }
                    }
                    var v33 int16 = 42
                    _ = len(v2)
                    _ = v33
                case "hello":
                    var v34 bool = true
                    _ = v34
                default:
                    var v35 float32 = (float32((float64(0.0) * 1.5)) / float32(((-2.25 + ((((float64(0.0) - float64(0.1)) + (3.14159 * (0.0 * 1.5))) / -2.25) - 3.14159)) - (0.1 + -2.25))))
                    _ = v35
                }
                var v36 int8 = genFn1(1)
                _ = v31
                _ = v36
            }
            _ = fi30
        }
        _ = fi28
        _ = v29
    }
    switch 9223372036854775807 {
    case 0:
        var v37 int32 = 0
        _ = v37
    default:
        delete(v4, 442)
    }
    v1 -= genFn1(32768)
    switch v27 {
    case "\n\t":
        var v38 float64 = 1.5
        _ = v38
    }
    _ = len(v2)
    var v39 uint8 = 145
    var v40 uint = 255
    v39 = v39
    for fi41 := 0; fi41 < 17; fi41++ {
        fmt.Println(v40)
        _ = fi41
    }
    switch (v27 + v27) {
    case "foo bar":
        v40 = 1
    case "\n\t":
        var v42 int32 = genFn0(v3, -2147483648)
        _ = v42
    case "":
        switch "foo bar" {
        case "hello":
            v27 = ((((v27 + v27) + "\n\t") + "a") + (v27 + ((v27 + (v27 + ("" + ("hello" + ("a" + v27))))) + v27)))
        case "a":
            v43 := GenS1{F0: v39}
            _ = v43
        case "":
            v44 := GenS3{F0: ((3.14159 / 3.14159) >= (((1e10 / 1.0) - (3.14159 / -2.25)) + (-2.25 - (0.1 / (1.5 / 1.5))))), F1: 9223372036854775807}
            _ = v44
        default:
            v5 = append(v5, -9223372036854775808)
        }
    default:
        var v45 int16 = 570
        _ = v45
    }
    fmt.Println(v40, v39, v39)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
