package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 int; F1 int8; F2 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32; F1 int64; F2 string }
type GenS5 struct { F0 int32; F1 uint }
func genFn0(p0 uint32) int8 {
    for fi5 := 0; fi5 < 3; fi5++ {
        var v6 int64 = 1
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := GenS1{F0: 42, F1: 122}
        if (false && ((fi7 ^ fi7) != (fi7 | (fi7 ^ (-(fi7)))))) {
            switch "hello" {
            case "":
                for fi9 := 0; fi9 < 16; fi9++ {
                    p0 -= 65535
                    var v10 float32 = (0.1 + 1.5)
                    v10 += ((float32(1e10) + float32(0.0)) * ((v10 + v10) / (3.14159 + v10)))
                    var v11 float32 = float32(1.5)
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            case "a":
                if false {
                    for fi12 := 0; fi12 < 7; fi12++ {
                        var v13 float32 = 1e10
                        _ = fi12
                        _ = v13
                    }
                    p0 -= (p0 & 0)
                    fmt.Println(p0)
                }
            case "hello":
                v14 := []int{(^(fi7)), -37, (^(fi7))}
                _ = v14
            }
            v15 := []int{}
            _ = v15
        }
        if ((0.0 - -2.25) == ((0.0 + (1.5 - ((0.0 / 1.0) * 1.5))) / -2.25)) {
            if (42 > 1) {
                if true {
                    var v16 float32 = 0.0
                    _ = v16
                }
            }
            for fi17 := 0; fi17 < 9; fi17++ {
                p0 = p0
                _ = fi17
            }
        }
        _ = fi7
        _ = v8
    }
    _ = p0
    return -3
}
func genFn1() string {
    if (!(9223372036854775807 != 1)) {
        v18 := GenS1{F0: 100, F1: 255}
        v19 := GenS1{F0: -128, F1: 127}
        var v20 int8 = -128
        _ = v18
        _ = v19
        _ = v20
    }
    switch "foo bar" {
    case "hello":
        for fi21 := 0; fi21 < 15; fi21++ {
            v22 := GenS2{F0: fi21, F1: 100, F2: 255}
            var v23 uint8 = 1
            _ = fi21
            _ = v22
            _ = v23
        }
    default:
        for fi24 := 0; fi24 < 16; fi24++ {
            v25 := []int{fi24}
            _ = fi24
            _ = v25
        }
    }
    var v26 float64 = 0.1
    for fi27 := 0; fi27 < 10; fi27++ {
        for fi28 := 0; fi28 < 11; fi28++ {
            fmt.Println(v26, v26, fi28)
            v26 += 0.1
            switch "\n\t" {
            case "hello":
                if false {
                    v29 := GenS2{F0: fi28, F1: 36, F2: 42}
                    _ = v29
                } else {
                    if false {
                        switch (fi28 << ((1) & 63)) {
                        case 104:
                            var v30 float64 = v26
                            _ = v30
                        case -83:
                            switch "hello" {
                            case "a":
                                v31 := []int{(fi27 + (-(fi28)))}
                                _ = v31
                            }
                        case 32768:
                            var v32 int32 = 1
                            _ = v32
                        }
                        v26 = (3.14159 + (3.14159 / float64(1e10)))
                        var v33 uint8 = 255
                        _ = v33
                    }
                    var v34 int = (fi28 >> ((42) & 63))
                    _ = v34
                }
            case "":
                var v35 bool = false
                _ = v35
            case "\n\t":
                v26 = 0.1
            default:
                v36 := []int{(fi28 * (^(fi28))), 127, (fi28 ^ (fi27 / ((fi28) | 1)))}
                _ = v36
            }
            _ = fi28
        }
        var v37 int64 = 128
        _ = fi27
        _ = v37
    }
    _ = v26
    return "\n\t"
}

func main() {
    v1 := []uint32{2147483647, 224, 1}
    var v2 uint32 = 4294967295
    var v3 float64 = 1.5
    var v4 int = 0
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v38 int64 = 9223372036854775807
    if len(v1) > 0 {
        v1[0] = v2
    }
    v38 = (^(v38))
    if false {
        v4 -= 128
        switch (v4 ^ (^(v4))) {
        case 2147483647:
            v3 = v3
        case 9223372036854775807:
            for fi39 := 0; fi39 < 6; fi39++ {
                for fi40 := 0; fi40 < 2; fi40++ {
                    if false {
                        _ = len(v1)
                        switch "foo bar" {
                        case "":
                            var v41 uint32 = (v2 >> ((418) & 63))
                            _ = v41
                        case "foo bar":
                            v42 := GenS3{F0: "hello"}
                            _ = v42
                        }
                    } else {
                        var v43 uint = 0
                        v44 := GenS4{F0: 0, F1: -1, F2: "foo bar"}
                        if len(v1) > 0 {
                            v1[0] = v2
                        }
                        _ = v43
                        _ = v44
                    }
                    v4 = fi40
                    if false {
                        v45 := GenS3{F0: ("" + "")}
                        var v46 float64 = (((1.5 + (v3 * 0.0)) * v3) * ((((0.1 * (v3 - float64((1.5 * 1e10)))) / v3) + (v3 * v3)) * (1e10 * 0.1)))
                        v47 := GenS5{F0: 462, F1: 0}
                        var v48 float64 = v46
                        _ = v45
                        _ = v46
                        _ = v47
                        _ = v48
                    } else {
                        var v49 uint32 = 11
                        _ = v49
                    }
                    _ = len(v1)
                    _ = fi40
                }
                var v50 int = fi39
                _ = fi39
                _ = v50
            }
        }
        v3 = v3
    } else {
        var v51 uint = 0
        _ = v51
    }
    fmt.Println(v4, v38)
    switch "" {
    case "foo bar":
        if (131 > 127) {
            var v52 uint16 = 65535
            _ = v52
        } else {
            v53 := GenS2{F0: 2147483647, F1: 38, F2: 127}
            _ = v53
        }
    case "a":
        var v54 int32 = -2147483648
        _ = v54
    default:
        var v55 string = "foo bar"
        _ = v55
    }
    if (0 != 4294967295) {
        if ((2 < 618) && true) {
            var v56 uint32 = (v2 | v2)
            var v57 uint16 = 2
            v57 = (v57 | (v57 * (v57 ^ uint16(v56))))
            _ = v56
            _ = v57
        } else {
            if (-2.25 > (0.1 / 3.14159)) {
                v58 := GenS5{F0: 442, F1: 177}
                for fi59 := 0; fi59 < 3; fi59++ {
                    switch v4 {
                    case -33:
                        fmt.Println(v3)
                    case -1:
                        var v60 float32 = ((0.0 + float32(v3)) + 0.1)
                        _ = v60
                    case -821:
                        if false {
                            var v61 float64 = v3
                            fmt.Println(v2, v4)
                            v61 = (v3 * -2.25)
                            _ = v61
                        }
                    }
                    if (638 > 1000) {
                        fmt.Println(v2)
                    }
                    if len(v1) > 0 {
                        v1[0] = (v2 % (((v2 + v2)) | 1))
                    }
                    _ = fi59
                }
                _ = v58
            } else {
                var v62 int64 = v38
                _ = v62
            }
            v3 = 1e10
        }
        var v63 string = ("foo bar" + "hello")
        var v64 int = v4
        _ = v63
        _ = v64
    }
    var v65 int64 = (v38 / ((-1) | 1))
    v66 := GenS5{F0: 0, F1: 65535}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v65
    _ = v66
}
