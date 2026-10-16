package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint16; F1 uint }
func genFn0(p0 int8) uint64 {
    if false {
        if true {
            p0 += (p0 - -128)
        }
        for fi7 := 0; fi7 < 2; fi7++ {
            v8 := []int{fi7, (fi7 / ((fi7) | 1))}
            fmt.Println(fi7, p0)
            var v9 uint16 = 0
            _ = fi7
            _ = v8
            _ = v9
        }
        v10 := []int{}
        _ = v10
    }
    var v11 int8 = p0
    _ = p0
    _ = v11
    return 42
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{}
    var v3 uint8 = 127
    v4 := make(map[string]int)
    var v5 uint = 42
    v6 := []string{}
    for fi12 := 0; fi12 < 3; fi12++ {
        fmt.Println(v3, fi12)
        var v13 float64 = ((float64(0.1) - float64(3.14159)) + (3.14159 + (((1.5 / float64(1e10)) / (0.1 / float64(-2.25))) * (0.1 / (0.1 * 1.5)))))
        _ = fi12
        _ = v13
    }
    v3 = 128
    if (true && ((true && false) && (42 > 245))) {
        fmt.Println(v3, v3)
        for fi14 := 0; fi14 < 17; fi14++ {
            fmt.Println(fi14)
            v5 -= v5
            var v15 int64 = 9223372036854775807
            _ = fi14
            _ = v15
        }
    } else {
        v16 := GenS1{F0: 9223372036854775807, F1: 65535, F2: "\n\t"}
        for fi17 := 0; fi17 < 18; fi17++ {
            v3 = 1
            var v18 uint16 = 76
            if true {
                var v19 uint64 = 896
                _ = v19
            } else {
                v20 := GenS1{F0: 930, F1: uint32(v3), F2: "hello"}
                for fi21 := 0; fi21 < 14; fi21++ {
                    var v22 uint16 = 32767
                    v23 := GenS1{F0: 0, F1: 2, F2: ""}
                    v3 = 2
                    fmt.Println(v22)
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                if true {
                    var v24 int = fi17
                    var v25 uint16 = (v18 << (((v5 | (^(v5)))) & 63))
                    _ = v24
                    _ = v25
                } else {
                    v18 += (v18 + 757)
                }
                _ = v20
            }
            _ = fi17
            _ = v18
        }
        v2 = append(v2, 65535)
        _ = v16
    }
    if true {
        for fi26 := 0; fi26 < 5; fi26++ {
            _ = len(v4)
            _ = fi26
        }
        for fi27 := 0; fi27 < 5; fi27++ {
            v3 += v3
            var v28 int64 = 811
            var v29 bool = false
            _ = fi27
            _ = v28
            _ = v29
        }
        var v30 int = -372
        _ = v30
    } else {
        if (((v3 < (v3 >> (((v5 & v5)) & 63))) || (uint16(v3) >= uint16(v5))) && (false || true)) {
            v5 = 42
            v31 := GenS2{F0: 2147483647}
            if (v5 < v5) {
                for fi32 := 0; fi32 < 10; fi32++ {
                    fmt.Println(v5)
                    _ = fi32
                }
                var v33 int8 = -1
                v33 = (v33 / ((v33) | 1))
                _ = v33
            } else {
                for fi34 := 0; fi34 < 4; fi34++ {
                    var v35 uint64 = 42
                    fmt.Println(v3)
                    _ = fi34
                    _ = v35
                }
            }
            _ = v31
        } else {
            if (0 > 0) {
                var v36 int8 = -30
                _ = v36
            }
            fmt.Println(v5, v5)
        }
        v37 := GenS1{F0: 0, F1: 2147483647, F2: "a"}
        _ = v37
    }
    if true {
        v3 -= v3
    }
    v38 := GenS2{F0: 4294967295}
    v39 := GenS2{F0: 2147483647}
    switch "a" {
    case "":
        var v40 uint8 = (^(v3))
        _ = v40
    case "\n\t":
        var v41 float32 = (0.1 / -2.25)
        _ = v41
    case "a":
        if (-9223372036854775808 <= 32767) {
            v5 += 295
        }
    default:
        var v42 bool = (674 == 734)
        _ = v42
    }
    if (!(138 <= 2147483648)) {
        var v43 int8 = -1
        _ = v43
    } else {
        fmt.Println(v5, v3)
        v44 := GenS3{F0: false, F1: int32(v3)}
        switch 0 {
        case 32768:
            v45 := GenS1{F0: 9223372036854775807, F1: 65535, F2: "hello"}
            _ = v45
        case 32767:
            switch "a" {
            case "hello":
                var v46 int32 = 96
                _ = v46
            case "a":
                for fi47 := 0; fi47 < 19; fi47++ {
                    for fi48 := 0; fi48 < 16; fi48++ {
                        v49 := GenS1{F0: 42, F1: 930, F2: ("foo bar" + "\n\t")}
                        _ = fi48
                        _ = v49
                    }
                    v1[246] = 268
                    v3 += (^(v3))
                    _ = fi47
                }
            }
        }
        _ = v44
    }
    v50 := GenS4{F0: 127}
    switch "foo bar" {
    case "foo bar":
        v51 := GenS1{F0: 9223372036854775807, F1: 793, F2: "hello"}
        _ = v51
    case "hello":
        fmt.Println(v5, v3)
    case "a":
        if (42 >= 65535) {
            switch 42 {
            case -181:
                var v52 int64 = -1
                _ = v52
            case 42:
                v3 += 136
            case 128:
                var v53 uint8 = (v3 & uint8(v5))
                _ = v53
            default:
                v54 := GenS3{F0: true, F1: 1}
                _ = v54
            }
            v1[65535] = 2147483648
        }
    default:
        v5 += (^(v5))
    }
    for fi55 := 0; fi55 < 10; fi55++ {
        v3 -= (v3 * 255)
        var v56 string = ""
        var v57 int8 = 1
        _ = fi55
        _ = v56
        _ = v57
    }
    for fi58 := 0; fi58 < 11; fi58++ {
        fmt.Println(v5)
        fmt.Println(fi58, v5, v5)
        _ = fi58
    }
    if (true || (-926 != 2147483647)) {
        switch 9223372036854775807 {
        case 42:
            if false {
                v5 = 18446744073709551615
                var v59 uint32 = uint32(v3)
                _ = len(v4)
                _ = v59
            } else {
                for fi60 := 0; fi60 < 18; fi60++ {
                    if ((308 > 1000) || false) {
                        v61 := GenS5{F0: 717, F1: v5}
                        _ = len(v2)
                        v62 := GenS3{F0: ((-128 == 73) && true), F1: -1}
                        _ = v61
                        _ = v62
                    }
                    _ = fi60
                }
            }
        case 128:
            v3 = 0
        }
    } else {
        if (true || (!(v3 == (v3 ^ uint8(v5))))) {
            var v63 float32 = 3.14159
            if true {
                if len(v2) > 0 {
                    v2[0] = 650
                }
                switch "" {
                case "a":
                    var v64 uint8 = 42
                    _ = v64
                default:
                    v65 := v4[("foo bar" + ("hello" + ("\n\t" + "a")))]
                    _ = v65
                }
            }
            _ = v63
        }
    }
    switch ("foo bar" + "a") {
    case "hello":
        v3 = 255
    }
    if ((true || false) || (("\n\t" + "foo bar") <= "\n\t")) {
        var v66 int32 = 541
        v3 -= v3
        for fi67 := 0; fi67 < 19; fi67++ {
            if ((^(v5)) >= v5) {
                fmt.Println(fi67, v3)
            }
            _ = fi67
        }
        _ = v66
    }
    v68 := GenS2{F0: -600}
    delete(v1, 2147483647)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v50
    _ = v6
    _ = v68
}
