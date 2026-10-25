package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint16 }
func genFn0(p0 uint8, p1 int64) int {
    v5 := GenS1{F0: ("a" + ("\n\t" + "a"))}
    var v6 uint64 = 0
    if false {
        var v7 int32 = -1
        for fi8 := 0; fi8 < 19; fi8++ {
            for fi9 := 0; fi9 < 3; fi9++ {
                var v10 string = "hello"
                var v11 float64 = float64(0.0)
                var v12 uint16 = 32767
                var v13 uint64 = 1
                v14 := GenS1{F0: v10}
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
                _ = v13
                _ = v14
            }
            switch 9223372036854775807 {
            case -616:
                switch "foo bar" {
                case "a":
                    for fi15 := 0; fi15 < 6; fi15++ {
                        var v16 int64 = (p1 * 42)
                        var v17 int16 = 341
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                }
            case 32767:
                fmt.Println(v6)
            default:
                v18 := []int{fi8, (^(fi8))}
                _ = v18
            }
            for fi19 := 0; fi19 < 6; fi19++ {
                fmt.Println(v6, v6, fi8)
                _ = fi19
            }
            _ = fi8
        }
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 uint16, p1 string) int {
    v20 := GenS2{F0: "\n\t"}
    var v21 string = p1
    var v22 uint = 739
    if (!false) {
        switch int(v22) {
        case -9223372036854775808:
            v23 := []int{4294967295, 127, 2147483647, -2147483648}
            _ = v23
        case 32767:
            fmt.Println(p0, v21, p0)
        }
        var v24 float32 = (((((((0.1 - 0.1) / (((float32(-2.25) / (1e10 + (3.14159 * 3.14159))) * 0.0) - 1.5)) * (float32(0.1) / 3.14159)) / ((((3.14159 - -2.25) + 1.5) / (1e10 + 3.14159)) / 1e10)) * (float32(3.14159) / (3.14159 + 1.5))) + (1.5 / 0.1)) / float32(-2.25))
        var v25 string = (((p1 + "foo bar") + (p1 + v21)) + ("\n\t" + "\n\t"))
        _ = v24
        _ = v25
    }
    _ = p0
    _ = p1
    _ = v20
    _ = v21
    _ = v22
    return 0
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]uint32)
    v3 := []string{}
    var v4 uint16 = 0
    fmt.Println(v4)
    if false {
        v4 = v4
    } else {
        v26 := GenS3{F0: v4, F1: 4294967295}
        var v27 int8 = 92
        _ = v26
        _ = v27
    }
    var v28 int8 = -1
    v29 := GenS2{F0: "foo bar"}
    var v30 uint16 = v4
    fmt.Println(v4, v28, v30)
    for fi31 := 0; fi31 < 18; fi31++ {
        if true {
            for fi32 := 0; fi32 < 10; fi32++ {
                v28 -= 127
                v30 = 1
                v33 := GenS4{F0: -128}
                v28 -= (v28 ^ (v28 / ((-76) | 1)))
                var v34 int16 = -1
                _ = fi32
                _ = v33
                _ = v34
            }
        } else {
            v35 := GenS4{F0: 127}
            if ((fi31 ^ -2147483648) == 761) {
                var v36 int = fi31
                var v37 uint64 = 1
                _ = v36
                _ = v37
            } else {
                if ((!false) || false) {
                    var v38 uint32 = 661
                    v30 = 255
                    _ = v38
                } else {
                    var v39 int = int(v28)
                    if (2147483647 == -385) {
                        for fi40 := 0; fi40 < 7; fi40++ {
                            var v41 uint32 = 42
                            var v42 int64 = -9223372036854775808
                            _ = fi40
                            _ = v41
                            _ = v42
                        }
                        v43 := GenS2{F0: "a"}
                        switch "a" {
                        case "\n\t":
                            v44 := GenS5{F0: (v4 + 65535)}
                            _ = v44
                        case "":
                            v30 = v4
                        default:
                            var v45 int8 = v28
                            _ = v45
                        }
                        _ = v43
                    }
                    fmt.Println(v28)
                    _ = v39
                }
                if false {
                    _ = len(v3)
                    v28 -= (^(v28))
                } else {
                    for fi46 := 0; fi46 < 10; fi46++ {
                        if true {
                            v28 += int8(v30)
                            if (!(4294967295 >= 798)) {
                                if (false && true) {
                                    delete(v2, int(v4))
                                    switch (fi46 >> ((0) & 63)) {
                                    case -220:
                                        v3 = append(v3, "")
                                    case 9223372036854775807:
                                        var v47 int8 = v28
                                        _ = v47
                                    case 107:
                                        var v48 bool = true
                                        _ = v48
                                    default:
                                        if len(v3) > 0 {
                                            v3[0] = "\n\t"
                                        }
                                    }
                                } else {
                                    if (2 < 178) {
                                        _ = len(v3)
                                    }
                                    var v49 int32 = 1
                                    _ = v49
                                }
                                v50 := GenS1{F0: "\n\t"}
                                var v51 int32 = -35
                                _ = v50
                                _ = v51
                            } else {
                                if len(v3) > 0 {
                                    v3[0] = ("\n\t" + "\n\t")
                                }
                            }
                            v4 += v30
                        }
                        var v52 uint8 = 0
                        v4 -= v30
                        _ = fi46
                        _ = v52
                    }
                }
            }
            v53 := GenS4{F0: (v28 - (v28 + v28))}
            _ = v35
            _ = v53
        }
        var v54 uint16 = 100
        v30 -= 1
        _ = fi31
        _ = v54
    }
    v2[4294967295] = 42
    v3 = append(v3, ("hello" + "hello"))
    _ = len(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
