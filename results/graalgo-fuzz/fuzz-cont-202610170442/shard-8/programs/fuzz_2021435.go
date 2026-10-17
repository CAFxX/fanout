package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 uint16; F2 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint; F1 int16; F2 uint32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint8; F1 bool }
type GenS7 struct { F0 int64; F1 uint64; F2 int64 }
type GenS8 struct { F0 int }
func genFn0(p0 int) int64 {
    p0 = (p0 - p0)
    v4 := GenS1{F0: 70}
    v5 := []int{p0, (p0 & 2147483647), p0, p0}
    fmt.Println(p0)
    _ = p0
    _ = v4
    _ = v5
    return -1
}
func genFn1() int8 {
    if false {
        var v6 float64 = 0.1
        var v7 uint32 = 65535
        for fi8 := 0; fi8 < 13; fi8++ {
            v9 := []int{(fi8 % ((fi8) | 1))}
            _ = fi8
            _ = v9
        }
        _ = v6
        _ = v7
    }
    for fi10 := 0; fi10 < 8; fi10++ {
        var v11 float32 = (1e10 / -2.25)
        for fi12 := 0; fi12 < 9; fi12++ {
            fmt.Println(v11)
            v13 := GenS2{F0: (4294967295 != 813), F1: 65535, F2: "foo bar"}
            v11 -= 1e10
            _ = fi12
            _ = v13
        }
        switch fi10 {
        case 32768:
            var v14 uint16 = 32767
            _ = v14
        case 4294967295:
            if true {
                v11 -= v11
            }
        default:
            switch "\n\t" {
            case "\n\t":
                v15 := GenS1{F0: 22}
                _ = v15
            case "":
                v16 := GenS1{F0: 100}
                _ = v16
            case "foo bar":
                var v17 float64 = 1e10
                _ = v17
            default:
                switch 1 {
                case 0:
                    v11 = v11
                case -634:
                    var v18 uint = 18446744073709551615
                    _ = v18
                }
            }
        }
        _ = fi10
        _ = v11
    }
    return 127
}

func main() {
    v1 := make(map[string]string)
    var v2 uint64 = 827
    var v3 int16 = 1
    v19 := GenS2{F0: (uint(v2) != 42), F1: 916, F2: ((("hello" + "hello") + "") + "")}
    if (0 == 65535) {
        var v20 int8 = 127
        v21 := []int{-776}
        _ = v20
        _ = v21
    } else {
        if (v3 < (v3 * v3)) {
            switch ("hello" + (("" + ("a" + "a")) + (((("hello" + (("" + "a") + "hello")) + ("foo bar" + "hello")) + ("\n\t" + "\n\t")) + "foo bar"))) {
            case "a":
                var v22 float64 = (((-2.25 - ((1e10 + 0.0) + ((-2.25 - (0.1 / float64(1e10))) * 0.0))) - (1.5 - 3.14159)) * 0.0)
                _ = v22
            case "":
                if false {
                    v3 += v3
                }
            case "\n\t":
                var v23 int8 = -1
                _ = v23
            default:
                fmt.Println(v3, v3)
            }
            v3 = 274
            var v24 float64 = (-2.25 + float64((((-2.25 / 1e10) + 0.0) * 1.5)))
            _ = v24
        } else {
            v2 += v2
        }
    }
    v1["\n\t"] = ("\n\t" + "\n\t")
    if (-91 <= 0) {
        fmt.Println(v3, v3)
    }
    v25 := v1["\n\t"]
    v25 = "foo bar"
    v2 = v2
    if (!(!true)) {
        v26 := GenS3{F0: 9223372036854775807}
        var v27 bool = false
        _ = v26
        _ = v27
    }
    if ((v2 * v2) != 4294967295) {
        v25 = ((((v25 + v25) + (v25 + (v25 + ("\n\t" + v25)))) + v25) + "foo bar")
        var v28 uint16 = 42
        v3 = v3
        _ = v28
    } else {
        if false {
            fmt.Println(v2, v2)
            var v29 bool = (((3.14159 != (-2.25 + float32((float64(0.1) + (0.1 * 0.1))))) || (false || true)) && ((v2 << ((18446744073709551615) & 63)) < (^(v2))))
            _ = v29
        }
        var v30 uint32 = 2147483648
        v31 := GenS4{F0: 2, F1: v3, F2: (v30 & (v30 % ((v30) | 1)))}
        _ = v30
        _ = v31
    }
    if (!true) {
        if (2147483647 > 65535) {
            v32 := []int{}
            if len(v32) > 0 {
                v32[0] = 104
            }
            _ = v32
        } else {
            v33 := GenS2{F0: true, F1: 255, F2: v25}
            if (true || (1000 >= v3)) {
                v1[((v25 + v25) + "\n\t")] = v25
                if (true && (v25 < v25)) {
                    fmt.Println(v2, v25)
                }
            }
            fmt.Println(v2, v25, v3)
            _ = v33
        }
        v34 := GenS5{F0: 42}
        _ = v34
    }
    for fi35 := 0; fi35 < 10; fi35++ {
        if (42 < 255) {
            switch v25 {
            case "hello":
                if (!(false || ((42 <= 42) && ((fi35 + 0) == 343)))) {
                    fmt.Println(fi35)
                    if (1.5 <= ((((((0.0 / 0.1) / 0.1) * 0.0) / (3.14159 - float32(1.5))) - 1.5) / (float32(1e10) + 3.14159))) {
                        switch fi35 {
                        case -2147483648:
                            v36 := []int{32767, 325, fi35, int(v2)}
                            _ = v36
                        case 9223372036854775807:
                            fmt.Println(fi35, v3, v3)
                        default:
                            for fi37 := 0; fi37 < 18; fi37++ {
                                if false {
                                    v38 := []int{(fi35 / ((fi35) | 1)), fi37, (fi37 >> ((255) & 63)), fi35}
                                    var v39 string = ((v25 + v25) + ((v25 + ((v25 + "foo bar") + (((((v25 + (v25 + v25)) + v25) + ("a" + "\n\t")) + ("" + v25)) + "\n\t"))) + v25))
                                    _ = v38
                                    _ = v39
                                }
                                _ = fi37
                            }
                        }
                    } else {
                        var v40 string = v25
                        _ = v40
                    }
                    switch v25 {
                    case "a":
                        var v41 uint8 = 255
                        _ = v41
                    case "":
                        _ = len(v1)
                    default:
                        fmt.Println(fi35, fi35)
                    }
                } else {
                    v42 := GenS6{F0: 128, F1: ((v3 >> ((527) & 63)) < v3)}
                    v43 := GenS5{F0: -9223372036854775808}
                    _ = v42
                    _ = v43
                }
            case "a":
                for fi44 := 0; fi44 < 17; fi44++ {
                    switch fi44 {
                    case 0:
                        delete(v1, v25)
                    case 4294967295:
                        v3 = int16(v2)
                    default:
                        v25 = v25
                    }
                    if false {
                        v3 += (v3 ^ v3)
                        v45 := []int{2147483647, (fi44 / ((fi44) | 1))}
                        _ = len(v45)
                        if (false && (!(((1.5 - 1e10) - float32((1e10 + 0.1))) > 3.14159))) {
                            v46 := GenS7{F0: 712, F1: 733, F2: 9223372036854775807}
                            if (42 >= -1) {
                                var v47 int16 = (v3 / ((v3) | 1))
                                _ = v47
                            }
                            if true {
                                if true {
                                    fmt.Println(fi44)
                                    var v48 int = 32767
                                    v25 = v25
                                    _ = v48
                                } else {
                                    switch (v25 + v25) {
                                    case "foo bar":
                                        var v49 float32 = 1.5
                                        _ = v49
                                    case "hello":
                                        switch (fi44 ^ (fi44 - (fi35 << ((2) & 63)))) {
                                        case 128:
                                            switch fi44 {
                                            case -480:
                                                v2 -= 2
                                            case 1:
                                                var v50 uint16 = 2
                                                _ = v50
                                            case 299:
                                                switch (fi35 * 806) {
                                                case -2147483648:
                                                    v2 = (v2 >> ((65535) & 63))
                                                case 127:
                                                    var v51 int64 = -9223372036854775808
                                                    _ = v51
                                                default:
                                                    switch 32768 {
                                                    case -30:
                                                        if true {
                                                            if (false || false) {
                                                                v3 = v3
                                                                v3 = (v3 * v3)
                                                            }
                                                            var v52 uint64 = 18446744073709551615
                                                            _ = v52
                                                        }
                                                    case -80:
                                                        v53 := GenS7{F0: -104, F1: v2, F2: genFn0(fi35)}
                                                        _ = v53
                                                    }
                                                }
                                            default:
                                                var v54 int = fi35
                                                _ = v54
                                            }
                                        case -2147483648:
                                            var v55 uint16 = 2
                                            _ = v55
                                        case -342:
                                            v3 -= v3
                                        }
                                    case "a":
                                        v56 := GenS8{F0: fi44}
                                        _ = v56
                                    }
                                }
                                v45 = append(v45, (fi35 + 9223372036854775807))
                            } else {
                                delete(v1, ("\n\t" + v25))
                            }
                            _ = v46
                        }
                        v2 = (v2 / (((v2 | 1)) | 1))
                        _ = v45
                    } else {
                        v25 = "foo bar"
                        v3 += (v3 + (v3 ^ v3))
                    }
                    fmt.Println(fi44, fi35)
                    v25 = v25
                    v57 := []int{fi44}
                    var v58 int32 = 2147483647
                    _ = fi44
                    _ = v57
                    _ = v58
                }
            }
        } else {
            var v59 int = 1
            v60 := []int{}
            for fi61 := 0; fi61 < 4; fi61++ {
                if (2147483647 != 648) {
                    switch 300 {
                    case -9223372036854775808:
                        var v62 int = 32767
                        _ = v62
                    case 32767:
                        delete(v1, v25)
                    default:
                        fmt.Println(v2, v59, v59)
                    }
                    v63 := GenS7{F0: 637, F1: (v2 >> ((1) & 63)), F2: 1}
                    _ = v63
                } else {
                    if len(v60) > 0 {
                        v60[0] = v59
                    }
                    v1["foo bar"] = v25
                    var v64 uint64 = (v2 % (((v2 & 263)) | 1))
                    _ = v64
                }
                fmt.Println(v3, v25)
                v25 = ("" + v25)
                _ = fi61
            }
            _ = v59
            _ = v60
        }
        v3 = -1
        var v65 int16 = (^(v3))
        _ = fi35
        _ = v65
    }
    fmt.Println(v3)
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v25
    _ = v3
}
