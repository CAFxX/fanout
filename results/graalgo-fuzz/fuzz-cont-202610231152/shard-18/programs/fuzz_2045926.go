package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 int8; F1 int8; F2 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64; F1 uint; F2 uint8 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint8 }
func genFn0() int8 {
    switch ("" + (("hello" + "a") + "")) {
    case "":
        switch 323 {
        case 9223372036854775807:
            v5 := GenS1{F0: 1, F1: 933}
            _ = v5
        case 198:
            for fi6 := 0; fi6 < 18; fi6++ {
                if (!true) {
                    var v7 int8 = 0
                    var v8 bool = (0 < 395)
                    _ = v7
                    _ = v8
                }
                v9 := GenS2{F0: -1, F1: 42, F2: 0}
                v10 := []int{fi6}
                _ = fi6
                _ = v10
                _ = v9
            }
        default:
            var v11 string = "a"
            _ = v11
        }
    case "a":
        v12 := GenS3{F0: -1}
        _ = v12
    default:
        var v13 int32 = -803
        _ = v13
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        var v15 uint64 = 4294967295
        v16 := GenS4{F0: 2, F1: 1, F2: 0}
        fmt.Println(fi14, fi14)
        _ = fi14
        _ = v15
        _ = v16
    }
    v17 := []int{82, 0, -2147483648, 127}
    var v18 uint64 = 4294967295
    _ = v17
    _ = v18
    return 127
}
func genFn1(p0 string) uint8 {
    switch 179 {
    case -2147483648:
        for fi19 := 0; fi19 < 12; fi19++ {
            if true {
                v20 := []int{400, fi19, -91, 128}
                v21 := GenS5{F0: 2}
                for fi22 := 0; fi22 < 16; fi22++ {
                    if (!(true || (((p0 + p0) >= ((("foo bar" + p0) + "a") + "hello")) && false))) {
                        var v23 uint64 = 4294967295
                        if len(v20) > 0 {
                            v20[0] = fi19
                        }
                        var v24 int = (fi22 | (fi19 + (fi19 << ((2) & 63))))
                        _ = v23
                        _ = v24
                    } else {
                        if false {
                            var v25 string = p0
                            fmt.Println(fi22)
                            var v26 int = 1
                            _ = v25
                            _ = v26
                        }
                    }
                    p0 = (p0 + "\n\t")
                    var v27 int8 = 74
                    _ = fi22
                    _ = v27
                }
                _ = v20
                _ = v21
            } else {
                p0 = ((p0 + ("a" + p0)) + ("\n\t" + (p0 + (p0 + "hello"))))
            }
            var v28 int8 = 1
            for fi29 := 0; fi29 < 19; fi29++ {
                fmt.Println(v28, v28, v28)
                var v30 uint = 0
                v30 -= v30
                _ = fi29
                _ = v30
            }
            _ = fi19
            _ = v28
        }
    case 42:
        var v31 int8 = genFn0()
        _ = v31
    default:
        if false {
            var v32 uint32 = 125
            for fi33 := 0; fi33 < 11; fi33++ {
                p0 = p0
                fmt.Println(fi33, v32, v32)
                _ = fi33
            }
            _ = v32
        } else {
            p0 = ("" + p0)
            for fi34 := 0; fi34 < 13; fi34++ {
                p0 = ((p0 + p0) + p0)
                fmt.Println(fi34, fi34)
                _ = fi34
            }
        }
    }
    v35 := []int{42, 1, 0, -876}
    for fi36 := 0; fi36 < 9; fi36++ {
        fmt.Println(fi36)
        v35 = append(v35, (fi36 + (fi36 / (((fi36 >> ((4294967295) & 63))) | 1))))
        var v37 uint64 = 584
        _ = fi36
        _ = v37
    }
    var v38 int32 = 59
    var v39 int16 = 42
    _ = p0
    _ = v35
    _ = v38
    _ = v39
    return 127
}

func main() {
    v1 := make(map[string]bool)
    v2 := []string{}
    var v3 float64 = 3.14159
    var v4 uint = 801
    fmt.Println(v3, v4)
    for fi40 := 0; fi40 < 13; fi40++ {
        v3 -= v3
        _ = fi40
    }
    v41 := GenS3{F0: 879}
    var v42 uint32 = 92
    if true {
        v4 = v4
        v1["hello"] = false
    }
    v43 := GenS3{F0: -193}
    v44 := GenS2{F0: 0, F1: 127, F2: int16(v42)}
    var v45 uint = v4
    switch 9223372036854775807 {
    case 42:
        if (501 <= v42) {
            for fi46 := 0; fi46 < 3; fi46++ {
                if true {
                    if false {
                        var v47 uint16 = 255
                        var v48 int = fi46
                        _ = v47
                        _ = v48
                    }
                    for fi49 := 0; fi49 < 4; fi49++ {
                        var v50 int8 = genFn0()
                        v51 := GenS5{F0: v45}
                        fmt.Println(fi49, v42, fi46)
                        _ = fi49
                        _ = v50
                        _ = v51
                    }
                    for fi52 := 0; fi52 < 1; fi52++ {
                        fmt.Println(v45)
                        _ = fi52
                    }
                }
                for fi53 := 0; fi53 < 7; fi53++ {
                    var v54 int8 = genFn0()
                    if len(v2) > 0 {
                        v2[0] = "hello"
                    }
                    _ = fi53
                    _ = v54
                }
                _ = fi46
            }
        } else {
            for fi55 := 0; fi55 < 18; fi55++ {
                if (662 <= 2) {
                    var v56 int32 = -661
                    for fi57 := 0; fi57 < 13; fi57++ {
                        var v58 uint8 = 2
                        v3 = (v3 * v3)
                        _ = fi57
                        _ = v58
                    }
                    _ = v56
                } else {
                    if (!true) {
                        for fi59 := 0; fi59 < 14; fi59++ {
                            v42 -= 2147483648
                            if true {
                                var v60 float32 = 3.14159
                                v45 = (v4 % (((^(v4))) | 1))
                                v60 = v60
                                delete(v1, "hello")
                                _ = v60
                            } else {
                                var v61 float32 = 0.0
                                var v62 uint64 = 358
                                v45 -= v45
                                _ = v61
                                _ = v62
                            }
                            _ = fi59
                        }
                    }
                    if false {
                        for fi63 := 0; fi63 < 1; fi63++ {
                            var v64 int16 = 0
                            v4 = (^(v4))
                            _ = fi63
                            _ = v64
                        }
                        v65 := GenS6{F0: 83}
                        _ = v65
                    } else {
                        var v66 bool = true
                        v67 := GenS2{F0: genFn0(), F1: 0, F2: 1000}
                        _ = v66
                        _ = v67
                    }
                    v1["foo bar"] = ("" > "hello")
                }
                _ = fi55
            }
            var v68 int8 = 42
            if true {
                v69 := GenS5{F0: (v45 | (v4 << ((v4) & 63)))}
                v45 = (v45 * v45)
                _ = v69
            } else {
                v1[""] = (v3 <= -2.25)
                v3 = v3
                var v70 uint32 = (v42 - v42)
                _ = v70
            }
            _ = v68
        }
    default:
        var v71 uint32 = v42
        _ = v71
    }
    v72 := v1[((((("\n\t" + "") + ("foo bar" + (("\n\t" + "") + "\n\t"))) + "hello") + "hello") + "foo bar")]
    v4 += (v4 + (^(v4)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v45:", v45)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v72
}
