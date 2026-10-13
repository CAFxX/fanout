package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 int64; F1 uint8 }
type GenS4 struct { F0 int32; F1 int64; F2 uint }
type GenS5 struct { F0 uint; F1 int64 }
func genFn0(p0 uint64, p1 uint16) int {
    p0 -= 1
    p0 = (p0 % (((p0 ^ 18446744073709551615)) | 1))
    _ = p0
    _ = p1
    return -1
}
func genFn1(p0 int32, p1 int16) uint8 {
    for fi4 := 0; fi4 < 18; fi4++ {
        var v5 int64 = -584
        p0 += p0
        if true {
            for fi6 := 0; fi6 < 14; fi6++ {
                p1 = (p1 & (p1 * p1))
                v7 := GenS1{F0: p1, F1: (v5 ^ v5)}
                _ = fi6
                _ = v7
            }
            if (true && ("foo bar" <= ("" + ""))) {
                v8 := []int{}
                _ = v8
            } else {
                for fi9 := 0; fi9 < 6; fi9++ {
                    var v10 uint32 = 159
                    switch "" {
                    case "":
                        v11 := []int{fi9}
                        _ = v11
                    case "hello":
                        var v12 uint16 = 424
                        _ = v12
                    }
                    var v13 uint16 = 269
                    if true {
                        v10 = 4294967295
                    }
                    var v14 float32 = 0.0
                    _ = fi9
                    _ = v10
                    _ = v13
                    _ = v14
                }
            }
            if false {
                var v15 uint16 = 848
                v16 := []int{genFn0(42, (v15 << ((18446744073709551615) & 63))), fi4, genFn0(42, v15)}
                _ = v15
                _ = v16
            } else {
                p1 = (p1 & p1)
                var v17 int8 = -128
                for fi18 := 0; fi18 < 20; fi18++ {
                    var v19 string = (("\n\t" + "") + "\n\t")
                    _ = fi18
                    _ = v19
                }
                _ = v17
            }
        } else {
            if true {
                p0 = 1
                p1 -= int16(p0)
            } else {
                var v20 uint16 = 0
                for fi21 := 0; fi21 < 5; fi21++ {
                    var v22 int8 = 127
                    if false {
                        if true {
                            v23 := []int{4294967295, (fi21 + (fi21 & (fi21 ^ (-(fi21))))), fi4}
                            v22 = 1
                            fmt.Println(v22, fi21)
                            _ = v23
                        } else {
                            v24 := GenS2{F0: 439, F1: 1}
                            var v25 uint16 = v20
                            _ = v24
                            _ = v25
                        }
                        switch (^(fi4)) {
                        case 127:
                            var v26 float64 = (1e10 * 1.5)
                            _ = v26
                        case 42:
                            if (!(-2.25 != (3.14159 / 1.0))) {
                                v27 := []int{}
                                v22 += (v22 << ((820) & 63))
                                v28 := GenS3{F0: (v5 & -1), F1: 192}
                                var v29 string = ("foo bar" + "a")
                                p0 = p0
                                p1 = (p1 << ((537) & 63))
                                _ = v27
                                _ = v28
                                _ = v29
                            } else {
                                fmt.Println(v5, fi4, fi4)
                            }
                        case -9223372036854775808:
                            var v30 int32 = 736
                            _ = v30
                        default:
                            v5 = v5
                        }
                    }
                    _ = fi21
                    _ = v22
                }
                _ = v20
            }
            p1 += p1
        }
        _ = fi4
        _ = v5
    }
    var v31 int8 = 0
    v32 := []int{-371, 2147483647}
    v33 := GenS2{F0: 42, F1: 128}
    if false {
        p1 = (p1 >> ((892) & 63))
    } else {
        if (!(false && (!(!false)))) {
            p1 = p1
        }
    }
    _ = p0
    _ = p1
    _ = v31
    _ = v32
    _ = v33
    return 101
}

func main() {
    var v1 uint16 = 32767
    v2 := []string{"foo bar", "\n\t"}
    v3 := make(map[string]string)
    if len(v2) > 0 {
        v2[0] = ""
    }
    var v34 uint8 = 17
    for fi35 := 0; fi35 < 8; fi35++ {
        switch fi35 {
        case 32767:
            var v36 uint64 = 516
            _ = v36
        case 128:
            var v37 string = "foo bar"
            _ = v37
        case -1:
            v38 := GenS4{F0: 1, F1: 31, F2: 1}
            _ = v38
        default:
            fmt.Println(fi35, v1, v1)
        }
        _ = fi35
    }
    v34 = v34
    v1 -= v1
    if (true && (!((v1 | v1) > (v1 / ((v1) | 1))))) {
        v39 := GenS3{F0: -201, F1: genFn1(0, int16(v1))}
        _ = v39
    } else {
        v34 = (v34 << ((4294967295) & 63))
        v34 -= v34
    }
    if true {
        switch ("a" + "hello") {
        case "a":
            fmt.Println(v1)
        case "":
            if true {
                fmt.Println(v1)
                fmt.Println(v34, v1)
                fmt.Println(v1)
            }
        default:
            var v40 int8 = -118
            _ = v40
        }
        if (2 == 0) {
            v34 -= genFn1(-2147483648, -32768)
        } else {
            var v41 bool = (!(800 <= 2))
            _ = v41
        }
        if (uint(v34) >= 0) {
            fmt.Println(v34)
            for fi42 := 0; fi42 < 4; fi42++ {
                var v43 uint8 = (^(v34))
                var v44 uint16 = (v1 >> ((1) & 63))
                _ = fi42
                _ = v43
                _ = v44
            }
            if false {
                var v45 uint8 = (v34 ^ (v34 & v34))
                var v46 bool = false
                var v47 float64 = 3.14159
                _ = v45
                _ = v46
                _ = v47
            } else {
                v48 := GenS3{F0: 42, F1: 255}
                _ = v48
            }
        }
    } else {
        v34 += v34
    }
    var v49 uint8 = v34
    var v50 bool = (!false)
    fmt.Println(v1, v34, v49)
    v51 := GenS5{F0: 1, F1: 9223372036854775807}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v49
    _ = v50
    _ = v51
}
