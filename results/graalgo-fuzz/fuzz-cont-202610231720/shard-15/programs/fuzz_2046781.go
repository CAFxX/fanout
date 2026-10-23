package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint64; F1 uint64 }
func genFn0(p0 int64, p1 bool) int16 {
    for fi6 := 0; fi6 < 2; fi6++ {
        v7 := []int{fi6, (^(fi6)), (fi6 ^ fi6)}
        _ = fi6
        _ = v7
    }
    var v8 string = "foo bar"
    var v9 uint8 = 67
    var v10 int8 = -1
    if (388 < 4294967295) {
        var v11 int8 = v10
        fmt.Println(p1, v11, v10)
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 uint32, p1 int8, p2 uint) uint32 {
    fmt.Println(p2, p1)
    switch ("a" + ("hello" + ((((((("\n\t" + "foo bar") + (("\n\t" + "hello") + ("foo bar" + "\n\t"))) + "hello") + "foo bar") + ("foo bar" + "hello")) + "\n\t") + "hello"))) {
    case "a":
        v12 := GenS1{F0: 65535}
        _ = v12
    case "foo bar":
        v13 := GenS1{F0: (p2 << ((p2) & 63))}
        _ = v13
    case "\n\t":
        v14 := GenS1{F0: (p2 * (p2 ^ p2))}
        _ = v14
    }
    v15 := GenS2{F0: 127}
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return (p0 % (((p0 / (((p0 + 902)) | 1))) | 1))
}

func main() {
    v1 := []string{(("foo bar" + "") + "foo bar")}
    v2 := make(map[uint32]uint32)
    var v3 int64 = 774
    v4 := make(map[uint32]bool)
    var v5 uint32 = 1
    var v16 float32 = 0.1
    v3 -= (v3 << ((2) & 63))
    v3 += (v3 * (^(v3)))
    for fi17 := 0; fi17 < 8; fi17++ {
        if (-21 > -32768) {
            fmt.Println(v16, v5)
            v3 = 42
        }
        v18 := GenS1{F0: 255}
        var v19 float32 = ((v16 - v16) * -2.25)
        _ = fi17
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 18; fi20++ {
        var v21 int16 = 279
        for fi22 := 0; fi22 < 12; fi22++ {
            switch "a" {
            case "hello":
                fmt.Println(v16, v5, v21)
            }
            var v23 uint32 = genFn1(v5, 1, 186)
            switch fi20 {
            case 32767:
                v21 -= (v21 + v21)
            case 549:
                var v24 int32 = 1
                _ = v24
            case 127:
                v23 = (v23 >> ((18446744073709551615) & 63))
            }
            var v25 float32 = ((1.5 + (v16 / (1.5 + ((v16 / (v16 / 1.0)) / ((v16 * float32((1e10 * (-2.25 * 0.0)))) - 0.0))))) - 1.5)
            fmt.Println(v25, v5, fi20)
            _ = fi22
            _ = v23
            _ = v25
        }
        _ = fi20
        _ = v21
    }
    v26 := GenS3{F0: genFn1((v5 / ((348) | 1)), -100, 255), F1: 667}
    v27 := GenS4{F0: 995}
    var v28 int64 = 0
    if (false && (-25 == 66)) {
        if (0.0 <= 0.0) {
            if (1 > -90) {
                var v29 int32 = 176
                var v30 int64 = -1
                if (!(-32768 < -1)) {
                    if ((("" + "hello") + "\n\t") == "") {
                        for fi31 := 0; fi31 < 20; fi31++ {
                            v32 := GenS1{F0: 876}
                            if (false && true) {
                                _ = len(v4)
                                for fi33 := 0; fi33 < 14; fi33++ {
                                    if true {
                                        var v34 uint8 = 42
                                        if ((!true) && false) {
                                            _ = len(v4)
                                        }
                                        v35 := GenS1{F0: 65535}
                                        _ = v34
                                        _ = v35
                                    } else {
                                        var v36 float64 = -2.25
                                        v37 := GenS2{F0: 2}
                                        _ = v36
                                        _ = v37
                                    }
                                    _ = fi33
                                }
                            }
                            var v38 uint64 = 650
                            _ = fi31
                            _ = v32
                            _ = v38
                        }
                        v29 -= v29
                    } else {
                        switch -266 {
                        case -2147483648:
                            var v39 uint32 = genFn1((v5 * (v5 + v5)), -1, uint(v5))
                            _ = v39
                        case 1:
                            var v40 bool = true
                            _ = v40
                        case 32768:
                            v29 = -2147483648
                        default:
                            v41 := GenS5{F0: 1}
                            _ = v41
                        }
                        v42 := GenS2{F0: 148}
                        _ = v42
                    }
                    var v43 int32 = int32(v5)
                    _ = v43
                }
                _ = v29
                _ = v30
            } else {
                v28 -= (v28 | (v3 * 9223372036854775807))
                for fi44 := 0; fi44 < 3; fi44++ {
                    var v45 int = (fi44 & (fi44 ^ (fi44 - (fi44 & (^(fi44))))))
                    v28 += 0
                    v46 := GenS6{F0: 62, F1: 0}
                    _ = fi44
                    _ = v45
                    _ = v46
                }
            }
        }
        fmt.Println(v28)
        delete(v2, (v5 | (v5 % ((v5) | 1))))
    }
    var v47 float64 = 1.5
    var v48 float64 = ((float64(((v16 / (((0.0 / float32(v47)) + 3.14159) - float32(v47))) - v16)) * ((v47 * ((1e10 + 1.5) / v47)) * v47)) - (v47 / float64(v16)))
    var v49 uint16 = 255
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v5
}
