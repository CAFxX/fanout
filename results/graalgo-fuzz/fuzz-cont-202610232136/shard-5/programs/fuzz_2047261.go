package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 uint64; F1 bool }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint8 }
func genFn0(p0 uint) int8 {
    var v5 uint8 = 255
    for fi6 := 0; fi6 < 3; fi6++ {
        v5 = (v5 + v5)
        var v7 float32 = 0.1
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 8; fi8++ {
        v9 := GenS1{F0: 10, F1: fi8}
        p0 -= 42
        v10 := []int{(fi8 - 753), 42}
        _ = fi8
        _ = v10
        _ = v9
    }
    fmt.Println(v5)
    var v11 uint16 = 0
    _ = p0
    _ = v11
    _ = v5
    return 127
}
func genFn1(p0 int16, p1 uint8) int32 {
    p0 -= (p0 / ((p0) | 1))
    if false {
        switch "hello" {
        case "foo bar":
            for fi12 := 0; fi12 < 6; fi12++ {
                p0 = p0
                var v13 uint16 = 65535
                _ = fi12
                _ = v13
            }
        case "hello":
            p0 = (p0 % ((p0) | 1))
        case "\n\t":
            if false {
                var v14 int64 = -1
                _ = v14
            }
        }
    }
    switch "\n\t" {
    case "":
        p0 -= (p0 * (^(p0)))
    case "foo bar":
        fmt.Println(p0, p0)
    case "a":
        switch 127 {
        case -826:
            p0 = p0
        case -1:
            if false {
                fmt.Println(p0, p0)
                var v15 float32 = (0.1 + (0.0 * 3.14159))
                _ = v15
            }
        }
    default:
        p0 = -175
    }
    _ = p0
    _ = p1
    return 2147483647
}

func main() {
    v1 := []string{"hello", (("\n\t" + "") + "foo bar"), ("foo bar" + ("foo bar" + "\n\t"))}
    v2 := make(map[uint32]uint32)
    var v3 uint64 = 438
    var v4 uint8 = 255
    var v16 float32 = float32((-2.25 / (3.14159 - 0.0)))
    for fi17 := 0; fi17 < 2; fi17++ {
        var v18 float32 = float32(0.0)
        fmt.Println(v4)
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 9; fi19++ {
        v4 += 0
        var v20 uint16 = 42
        var v21 uint16 = (v20 * v20)
        _ = fi19
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 18; fi22++ {
        v23 := GenS2{F0: -1, F1: 511}
        switch "\n\t" {
        case "\n\t":
            fmt.Println(v16, v16)
        case "":
            v16 -= 1e10
        }
        if false {
            for fi24 := 0; fi24 < 1; fi24++ {
                v4 = (v4 | v4)
                _ = fi24
            }
            v25 := v2[uint32(v3)]
            for fi26 := 0; fi26 < 2; fi26++ {
                if false {
                    _ = len(v2)
                    v27 := GenS2{F0: -2147483648, F1: (^(v3))}
                    switch fi22 {
                    case -467:
                        var v28 uint64 = 1
                        _ = v28
                    case -997:
                        if (727 != 713) {
                            var v29 string = ("" + "a")
                            v4 += uint8(v25)
                            _ = v29
                        } else {
                            var v30 int = fi26
                            _ = v30
                        }
                    case 32767:
                        if (false || (!false)) {
                            if false {
                                fmt.Println(fi22, v25)
                                v25 = v25
                            }
                            if true {
                                v16 += ((-2.25 * v16) - v16)
                            } else {
                                v31 := GenS3{F0: 42, F1: ("hello" > "a")}
                                switch "a" {
                                case "foo bar":
                                    var v32 int32 = -808
                                    _ = v32
                                case "hello":
                                    if false {
                                        v3 += 0
                                        var v33 float64 = (0.1 - 0.0)
                                        _ = v33
                                    }
                                case "a":
                                    v34 := GenS1{F0: (v4 * (v4 << ((340) & 63))), F1: fi26}
                                    _ = v34
                                }
                                _ = v31
                            }
                            v35 := GenS3{F0: (^(v3)), F1: false}
                            _ = v35
                        }
                    default:
                        v36 := v2[(v25 | (v25 ^ v25))]
                        _ = v36
                    }
                    var v37 uint = 18446744073709551615
                    _ = v27
                    _ = v37
                }
                _ = fi26
            }
            _ = v25
        } else {
            if ((3.14159 / 1.0) <= float64(v16)) {
                var v38 string = ""
                v39 := GenS1{F0: (v4 >> ((810) & 63)), F1: 9223372036854775807}
                _ = v38
                _ = v39
            } else {
                var v40 uint8 = v4
                _ = v40
            }
            switch -2147483648 {
            case 0:
                var v41 int8 = -1
                _ = v41
            }
        }
        _ = fi22
        _ = v23
    }
    v16 = v16
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v42 := GenS1{F0: (^(v4)), F1: -823}
    for fi43 := 0; fi43 < 11; fi43++ {
        v44 := GenS4{F0: 1}
        _ = fi43
        _ = v44
    }
    v16 -= ((((1.5 / v16) + v16) * 3.14159) + (float32(float64((v16 + (v16 - ((v16 / v16) / 1.5))))) / 1e10))
    v3 = v3
    v16 += 1e10
    v45 := GenS5{F0: 186}
    fmt.Println(v3, v3, v16)
    fmt.Println(v3)
    fmt.Println(v4)
    v1 = append(v1, ("\n\t" + ("\n\t" + (((("\n\t" + (((("foo bar" + "") + "") + (("a" + ("a" + "foo bar")) + "\n\t")) + "")) + "\n\t") + ("\n\t" + "")) + ("a" + ("\n\t" + ("hello" + ("hello" + ("a" + "a")))))))))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v45
}
