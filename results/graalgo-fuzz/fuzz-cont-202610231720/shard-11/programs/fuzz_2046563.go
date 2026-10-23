package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 int; F1 int }
type GenS4 struct { F0 uint16 }
func genFn0(p0 uint, p1 uint) uint {
    p0 = p1
    v7 := GenS1{F0: 42}
    p0 = p1
    var v8 uint16 = 32767
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return p1
}
func genFn1(p0 int64) uint64 {
    for fi9 := 0; fi9 < 9; fi9++ {
        for fi10 := 0; fi10 < 3; fi10++ {
            v11 := []int{fi10, fi9}
            if false {
                var v12 uint = 457
                v13 := GenS2{F0: true, F1: -915}
                var v14 float64 = 1.5
                _ = v12
                _ = v13
                _ = v14
            } else {
                p0 -= (p0 | (p0 & (-(p0))))
                v11 = append(v11, (fi9 + (fi9 / (((fi9 % ((fi10) | 1))) | 1))))
                if true {
                    var v15 int16 = 42
                    v15 -= v15
                    v15 += -32768
                    _ = v15
                } else {
                    var v16 uint16 = 223
                    v17 := GenS1{F0: 2}
                    var v18 string = ""
                    _ = v16
                    _ = v17
                    _ = v18
                }
                _ = len(v11)
                if (false && (240 <= 42)) {
                    if (-32768 >= 0) {
                        var v19 uint32 = 2147483648
                        var v20 bool = false
                        switch "foo bar" {
                        case "hello":
                            if (!(!true)) {
                                v21 := GenS2{F0: v20, F1: -95}
                                _ = len(v11)
                                _ = v21
                            }
                        case "foo bar":
                            if ((!true) || (!true)) {
                                switch (fi9 * (fi9 >> ((18446744073709551615) & 63))) {
                                case -890:
                                    var v22 int32 = 1
                                    _ = v22
                                default:
                                    if true {
                                        v11 = append(v11, -766)
                                        var v23 uint64 = 821
                                        if len(v11) > 0 {
                                            v11[0] = (fi9 << ((0) & 63))
                                        }
                                        _ = v23
                                    } else {
                                        var v24 int32 = -383
                                        switch ((("a" + "a") + (("a" + "foo bar") + "")) + (("" + "a") + "hello")) {
                                        case "\n\t":
                                            if (false && v20) {
                                                v25 := GenS3{F0: (fi10 - -986), F1: 0}
                                                _ = v25
                                            }
                                        case "hello":
                                            v20 = (18446744073709551615 <= 1)
                                        case "":
                                            v19 = (v19 % (((^(v19))) | 1))
                                        default:
                                            var v26 bool = v20
                                            _ = v26
                                        }
                                        _ = v24
                                    }
                                }
                                v19 = (v19 + 2147483648)
                                v20 = (!(v20 && false))
                            } else {
                                var v27 int16 = -38
                                fmt.Println(v20, v27)
                                _ = v27
                            }
                        default:
                            var v28 int = (fi9 ^ fi9)
                            _ = v28
                        }
                        _ = v19
                        _ = v20
                    }
                } else {
                    if false {
                        var v29 int16 = 829
                        p0 -= (p0 - -1)
                        var v30 float64 = 3.14159
                        v30 = v30
                        v31 := GenS2{F0: true, F1: -1}
                        _ = v29
                        _ = v30
                        _ = v31
                    } else {
                        if (!((3.14159 / 1.5) >= float64(((1e10 - float32(3.14159)) + 1.5)))) {
                            switch fi10 {
                            case -288:
                                v32 := GenS2{F0: (true || (1 != 32767)), F1: -1}
                                _ = v32
                            case 1:
                                var v33 string = "hello"
                                _ = v33
                            default:
                                v34 := GenS1{F0: 203}
                                _ = v34
                            }
                            var v35 uint8 = 127
                            _ = v35
                        }
                        v36 := GenS3{F0: fi10, F1: 32767}
                        var v37 uint = 456
                        _ = v36
                        _ = v37
                    }
                    p0 -= p0
                }
                v11 = append(v11, (fi10 % ((-514) | 1)))
                p0 = (p0 + (p0 * (p0 | 1)))
            }
            _ = fi10
            _ = v11
        }
        v38 := GenS1{F0: 42}
        _ = fi9
        _ = v38
    }
    if (0.0 != 1.5) {
        var v39 uint64 = 907
        v39 += (v39 >> ((1) & 63))
        _ = v39
    } else {
        var v40 uint = 4294967295
        _ = v40
    }
    for fi41 := 0; fi41 < 7; fi41++ {
        fmt.Println(fi41, fi41)
        p0 -= (p0 << ((534) & 63))
        _ = fi41
    }
    p0 = 605
    _ = p0
    return 683
}

func main() {
    var v1 uint = 2
    var v2 float32 = 0.0
    v3 := []uint32{65535}
    var v4 int64 = 9223372036854775807
    var v5 uint32 = 950
    var v6 uint64 = 2
    v6 += 18446744073709551615
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch 9223372036854775807 {
    case 0:
        var v42 uint64 = genFn1((v4 | (^(v4))))
        _ = v42
    case -885:
        v6 = uint64(v5)
    case 576:
        if false {
            fmt.Println(v6)
        } else {
            if len(v3) > 0 {
                v3[0] = v5
            }
            v43 := GenS4{F0: 688}
            v6 -= (v6 * (v6 | (v6 % ((v6) | 1))))
            _ = v43
        }
    }
    switch ("hello" + "foo bar") {
    case "hello":
        if (42 != 1) {
            var v44 float32 = -2.25
            _ = v44
        }
    case "":
        v2 -= float32((3.14159 - (-2.25 - 3.14159)))
    default:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    for fi45 := 0; fi45 < 6; fi45++ {
        var v46 uint16 = 42
        _ = fi45
        _ = v46
    }
    v47 := GenS4{F0: 255}
    v4 -= -9223372036854775808
    v3 = append(v3, (v5 << (((v1 | 4294967295)) & 63)))
    switch ("foo bar" + "") {
    case "":
        var v48 uint64 = (^(v6))
        _ = v48
    case "hello":
        var v49 uint16 = 183
        _ = v49
    case "\n\t":
        var v50 int8 = 42
        _ = v50
    }
    v4 -= 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
}
