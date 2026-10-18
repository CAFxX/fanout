package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 uint16 }
type GenS4 struct { F0 uint32; F1 uint64; F2 string }
type GenS5 struct { F0 int8; F1 string }
func genFn0(p0 uint8) uint16 {
    if (162 != 2) {
        if true {
            switch 42 {
            case 4294967295:
                fmt.Println(p0)
            }
        } else {
            fmt.Println(p0)
            var v4 int8 = 123
            _ = v4
        }
        p0 = (p0 * (^(p0)))
    }
    switch "foo bar" {
    case "hello":
        v5 := []int{-857, 276, 128, 0}
        _ = v5
    case "a":
        var v6 uint16 = uint16(p0)
        _ = v6
    default:
        var v7 uint16 = 42
        _ = v7
    }
    var v8 bool = (!true)
    _ = p0
    _ = v8
    return 255
}
func genFn1() bool {
    for fi9 := 0; fi9 < 12; fi9++ {
        var v10 float64 = ((1e10 - 0.1) - (1e10 + (1.5 - ((((0.0 * 0.0) + (float64(1e10) + 1e10)) - ((0.0 / float64(3.14159)) / ((float64(3.14159) * ((1e10 - float64(-2.25)) - float64(1.5))) / 1e10))) - 3.14159))))
        _ = fi9
        _ = v10
    }
    v11 := []int{0, -1, -13, -9223372036854775808}
    v12 := GenS1{F0: 2}
    var v13 uint32 = 472
    if len(v11) > 0 {
        v11[0] = -1
    }
    _ = v11
    _ = v12
    _ = v13
    return (!true)
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]uint32)
    var v3 bool = (true || (!(!(!false))))
    fmt.Println(v3)
    var v14 int32 = -2147483648
    for fi15 := 0; fi15 < 20; fi15++ {
        v16 := v2[2]
        _ = fi15
        _ = v16
    }
    switch -1 {
    case 127:
        switch "" {
        case "hello":
            var v17 int8 = -7
            _ = v17
        case "a":
            v14 -= v14
        case "\n\t":
            for fi18 := 0; fi18 < 9; fi18++ {
                v14 = (v14 << ((246) & 63))
                var v19 uint = 18446744073709551615
                v20 := GenS1{F0: 0}
                _ = fi18
                _ = v19
                _ = v20
            }
        default:
            switch "foo bar" {
            case "hello":
                fmt.Println(v14, v14)
            case "":
                if v3 {
                    var v21 uint8 = 0
                    for fi22 := 0; fi22 < 18; fi22++ {
                        if (((1.5 / 1.5) == (3.14159 / 0.1)) && (!v3)) {
                            var v23 string = (("" + (("hello" + ("\n\t" + ("" + "a"))) + "")) + "a")
                            var v24 uint64 = 4294967295
                            var v25 uint = 4294967295
                            _ = v23
                            _ = v24
                            _ = v25
                        }
                        if v3 {
                            var v26 int64 = -970
                            v26 = (v26 * (v26 >> ((4294967295) & 63)))
                            _ = v26
                        } else {
                            for fi27 := 0; fi27 < 3; fi27++ {
                                switch "\n\t" {
                                case "":
                                    var v28 uint16 = 42
                                    _ = v28
                                case "\n\t":
                                    v29 := GenS2{F0: (v21 / ((v21) | 1))}
                                    _ = v29
                                }
                                _ = fi27
                            }
                            v30 := GenS1{F0: genFn0(255)}
                            _ = v30
                        }
                        fmt.Println(v14, v14, fi22)
                        _ = fi22
                    }
                    _ = v21
                }
            default:
                v31 := GenS1{F0: 2}
                _ = v31
            }
        }
    default:
        fmt.Println(v14)
    }
    delete(v2, 736)
    v3 = v3
    for fi32 := 0; fi32 < 17; fi32++ {
        fmt.Println(fi32, v3)
        _ = fi32
    }
    v33 := GenS3{F0: 764, F1: 94}
    var v34 int16 = 0
    var v35 uint64 = 0
    var v36 int32 = v14
    v36 -= v36
    v37 := []int{2147483647, 32767}
    var v38 int64 = 42
    fmt.Println(v34)
    _ = len(v37)
    switch "foo bar" {
    case "":
        if ("\n\t" <= "a") {
            if v3 {
                var v39 float32 = 0.1
                if false {
                    switch -9223372036854775808 {
                    case 32768:
                        _ = len(v1)
                    case 127:
                        var v40 uint64 = (^(v35))
                        _ = v40
                    default:
                        fmt.Println(v38)
                    }
                    v34 = (-(v34))
                }
                v41 := GenS1{F0: genFn0(42)}
                _ = v39
                _ = v41
            }
        } else {
            v2[279] = 4294967295
            v42 := GenS4{F0: 42, F1: 2, F2: "\n\t"}
            _ = v42
        }
    case "a":
        for fi43 := 0; fi43 < 3; fi43++ {
            _ = len(v1)
            v1[127] = (!((1 != genFn0(197)) && (v3 || ((!(2147483647 >= (v14 << ((193) & 63)))) && v3))))
            v14 -= (v36 ^ (-(v36)))
            _ = fi43
        }
    case "\n\t":
        v44 := GenS5{F0: -1, F1: ("" + "\n\t")}
        _ = v44
    default:
        fmt.Println(v38, v14, v36)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
}
