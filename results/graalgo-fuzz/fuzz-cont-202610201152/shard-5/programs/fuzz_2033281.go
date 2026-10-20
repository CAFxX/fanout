package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 int16; F2 uint8 }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 uint8; F1 uint64 }
type GenS5 struct { F0 uint8; F1 uint16 }
type GenS6 struct { F0 int8; F1 int8; F2 uint }
type GenS7 struct { F0 string; F1 uint8 }
type GenS8 struct { F0 string; F1 uint16 }
func genFn0(p0 uint, p1 uint32) int {
    var v4 int64 = 1
    var v5 uint = p0
    p1 += 1
    if ("\n\t" != "a") {
        p1 = (p1 ^ (p1 >> ((4294967295) & 63)))
        for fi6 := 0; fi6 < 15; fi6++ {
            if true {
                if false {
                    p1 = (p1 - p1)
                    if false {
                        var v7 int8 = 0
                        for fi8 := 0; fi8 < 20; fi8++ {
                            v9 := []int{(fi8 / ((-1) | 1)), (fi6 / ((fi8) | 1))}
                            _ = fi8
                            _ = v9
                        }
                        v4 += 1
                        _ = v7
                    }
                }
                for fi10 := 0; fi10 < 16; fi10++ {
                    if ("a" > (("hello" + "foo bar") + "\n\t")) {
                        v11 := GenS1{F0: v4}
                        switch -9223372036854775808 {
                        case -1:
                            v5 = v5
                        case 1:
                            switch "hello" {
                            case "a":
                                var v12 uint8 = uint8(p1)
                                var v13 float32 = 0.1
                                var v14 float32 = (v13 + v13)
                                _ = v12
                                _ = v13
                                _ = v14
                            case "hello":
                                if false {
                                    var v15 int32 = 856
                                    fmt.Println(p0)
                                    _ = v15
                                } else {
                                    v16 := GenS2{F0: 255, F1: 329, F2: 42}
                                    _ = v16
                                }
                            case "foo bar":
                                v17 := GenS3{F0: 42, F1: 822}
                                _ = v17
                            default:
                                v18 := []int{fi6, fi6, (^(fi10)), fi10}
                                _ = v18
                            }
                        }
                        if true {
                            if (42 == -1) {
                                var v19 float64 = 1.5
                                fmt.Println(fi6)
                                _ = v19
                            } else {
                                v20 := []int{}
                                v5 -= (^(p0))
                                if len(v20) > 0 {
                                    v20[0] = fi6
                                }
                                _ = v20
                            }
                            var v21 string = ""
                            var v22 uint8 = 42
                            _ = v21
                            _ = v22
                        }
                        _ = v11
                    } else {
                        if ((fi6 | (fi6 - 42)) < 42) {
                            var v23 int = (fi6 + fi6)
                            var v24 int16 = -28
                            _ = v23
                            _ = v24
                        } else {
                            var v25 int64 = 373
                            v26 := []int{-9223372036854775808, (fi6 + (fi6 * 2147483647))}
                            _ = v25
                            _ = v26
                        }
                        v4 -= 42
                    }
                    _ = fi10
                }
            }
            _ = fi6
        }
    }
    var v27 int8 = int8(v4)
    _ = p0
    _ = p1
    _ = v27
    _ = v4
    _ = v5
    return 128
}
func genFn1(p0 int64, p1 int32) uint32 {
    for fi28 := 0; fi28 < 14; fi28++ {
        var v29 bool = (false && true)
        _ = fi28
        _ = v29
    }
    var v30 string = (("" + "foo bar") + "hello")
    _ = p0
    _ = p1
    _ = v30
    return 468
}

func main() {
    v1 := []string{"hello", (("hello" + ("\n\t" + "a")) + ("hello" + "foo bar")), "a"}
    var v2 int = 32767
    v3 := make(map[string]string)
    switch v2 {
    case 9223372036854775807:
        var v31 bool = (545 == 838)
        _ = v31
    case 4294967295:
        v32 := GenS4{F0: 43, F1: 4294967295}
        _ = v32
    case 210:
        switch ("" + ("foo bar" + ("hello" + "\n\t"))) {
        case "":
            if len(v1) > 0 {
                v1[0] = ("foo bar" + ("" + "hello"))
            }
        default:
            for fi33 := 0; fi33 < 7; fi33++ {
                v34 := GenS5{F0: 63, F1: 2}
                var v35 string = "\n\t"
                _ = fi33
                _ = v34
                _ = v35
            }
        }
    default:
        var v36 float32 = (3.14159 / 0.1)
        _ = v36
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v37 := GenS5{F0: 42, F1: 533}
    v38 := GenS4{F0: 1, F1: 1}
    v39 := GenS1{F0: -9223372036854775808}
    v40 := v3["\n\t"]
    fmt.Println(v2, v40)
    if false {
        var v41 bool = true
        if (!v41) {
            v41 = (!(false && (v2 > genFn0(2, 1))))
            fmt.Println(v41)
        } else {
            var v42 uint64 = 1
            _ = v42
        }
        _ = v41
    }
    v40 = "a"
    switch v2 {
    case -198:
        for fi43 := 0; fi43 < 11; fi43++ {
            var v44 bool = true
            _ = fi43
            _ = v44
        }
    default:
        if (!(100 < 127)) {
            var v45 uint8 = 127
            v2 -= (v2 ^ v2)
            _ = v45
        } else {
            for fi46 := 0; fi46 < 15; fi46++ {
                if false {
                    v47 := GenS6{F0: -128, F1: 1, F2: 1}
                    v48 := GenS7{F0: v40, F1: 138}
                    _ = v47
                    _ = v48
                } else {
                    fmt.Println(fi46, v2)
                    v1 = append(v1, "a")
                    v49 := GenS8{F0: "", F1: 2}
                    _ = v49
                }
                _ = fi46
            }
            var v50 int64 = 9223372036854775807
            if (!true) {
                v40 = "a"
                _ = len(v1)
                v51 := GenS2{F0: 65535, F1: 558, F2: 138}
                _ = v51
            }
            _ = v50
        }
    }
    v40 = ("hello" + v40)
    v52 := GenS5{F0: 2, F1: 34}
    fmt.Println(v40, v2)
    for fi53 := 0; fi53 < 2; fi53++ {
        var v54 int16 = -896
        _ = len(v1)
        for fi55 := 0; fi55 < 20; fi55++ {
            v40 = "hello"
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi55
        }
        _ = fi53
        _ = v54
    }
    v56 := GenS8{F0: "hello", F1: 576}
    v57 := GenS4{F0: 226, F1: 18446744073709551615}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v40
    _ = v52
    _ = v56
    _ = v57
}
