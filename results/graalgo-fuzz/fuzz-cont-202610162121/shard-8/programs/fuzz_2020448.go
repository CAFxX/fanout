package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 uint32; F2 int }
type GenS3 struct { F0 int32; F1 bool; F2 bool }
type GenS4 struct { F0 int32; F1 uint16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint16; F1 int32; F2 uint16 }
type GenS7 struct { F0 int64 }
func genFn0() bool {
    switch 32767 {
    case -475:
        var v5 float32 = (-2.25 * 3.14159)
        _ = v5
    case -399:
        var v6 float64 = 3.14159
        _ = v6
    case 9223372036854775807:
        v7 := []int{1, 1, 4294967295}
        _ = v7
    default:
        var v8 uint32 = 2147483647
        _ = v8
    }
    var v9 uint64 = 18446744073709551615
    v10 := GenS1{F0: -1}
    var v11 uint32 = 9
    var v12 float64 = (float64((3.14159 - 1e10)) * 1e10)
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return false
}

func main() {
    var v1 uint32 = 249
    var v2 uint32 = 42
    var v3 uint16 = 255
    var v4 bool = true
    for fi13 := 0; fi13 < 13; fi13++ {
        switch 235 {
        case -9223372036854775808:
            switch "a" {
            case "\n\t":
                v14 := []int{fi13, 32768, fi13}
                _ = v14
            case "foo bar":
                if v4 {
                    for fi15 := 0; fi15 < 6; fi15++ {
                        v16 := GenS1{F0: 42}
                        v17 := GenS1{F0: int16(v1)}
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                } else {
                    fmt.Println(fi13, v4, v3)
                    v3 += (^(v3))
                    v4 = (2 >= 0)
                }
            case "hello":
                switch (fi13 & 32767) {
                case 2147483647:
                    fmt.Println(v2)
                case 1:
                    for fi18 := 0; fi18 < 7; fi18++ {
                        v19 := GenS2{F0: -32768, F1: (^(v2)), F2: 1}
                        v3 += 255
                        v20 := []int{fi18, (fi18 / (((fi18 - -923)) | 1)), -49, (fi13 ^ (fi18 % ((fi18) | 1)))}
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                }
            default:
                for fi21 := 0; fi21 < 2; fi21++ {
                    var v22 string = ("\n\t" + "\n\t")
                    fmt.Println(v4, v1)
                    _ = fi21
                    _ = v22
                }
            }
        }
        _ = fi13
    }
    v1 -= v2
    for fi23 := 0; fi23 < 6; fi23++ {
        v24 := []int{}
        fmt.Println(v3, v2, fi23)
        v2 += v2
        _ = fi23
        _ = v24
    }
    v25 := GenS3{F0: 2147483647, F1: v4, F2: ((128 > 1) && ("hello" > "hello"))}
    for fi26 := 0; fi26 < 3; fi26++ {
        for fi27 := 0; fi27 < 20; fi27++ {
            v28 := GenS1{F0: 42}
            _ = fi27
            _ = v28
        }
        var v29 int16 = -32768
        v2 = (v1 << ((65535) & 63))
        _ = fi26
        _ = v29
    }
    v30 := []int{0, int(v3), 128, 177}
    var v31 int64 = 9223372036854775807
    var v32 string = "a"
    var v33 string = "foo bar"
    fmt.Println(v32)
    v31 = (v31 % ((42) | 1))
    if (255 <= 65535) {
        var v34 uint = 0
        _ = len(v30)
        _ = v34
    }
    var v35 float32 = float32(0.1)
    v36 := GenS3{F0: 2147483647, F1: true, F2: (42 != int8(v31))}
    var v37 uint8 = 255
    switch (v32 + "") {
    case "a":
        for fi38 := 0; fi38 < 6; fi38++ {
            v4 = (v35 == v35)
            v4 = v4
            v33 = (v32 + ((((v32 + "foo bar") + v32) + (v33 + (v32 + v32))) + v32))
            _ = fi38
        }
    case "foo bar":
        if true {
            for fi39 := 0; fi39 < 20; fi39++ {
                var v40 int8 = 51
                if len(v30) > 0 {
                    v30[0] = fi39
                }
                var v41 uint16 = v3
                _ = fi39
                _ = v40
                _ = v41
            }
            var v42 int = 54
            _ = v42
        } else {
            v2 -= 42
            fmt.Println(v37, v31, v35)
            v43 := GenS4{F0: -2147483648, F1: 0}
            _ = v43
        }
    default:
        v30 = append(v30, -9223372036854775808)
    }
    for fi44 := 0; fi44 < 13; fi44++ {
        v45 := GenS5{F0: v1}
        _ = fi44
        _ = v45
    }
    if v4 {
        for fi46 := 0; fi46 < 5; fi46++ {
            v32 = "\n\t"
            _ = fi46
        }
        if ((2147483647 > 4294967295) || (true || (!false))) {
            switch "hello" {
            case "foo bar":
                fmt.Println(v32, v35, v35)
            case "\n\t":
                if len(v30) >= 2 {
                    v30 = v30[1:len(v30)]
                }
            }
        }
        _ = len(v30)
    } else {
        for fi47 := 0; fi47 < 5; fi47++ {
            if len(v30) > 0 {
                v30[0] = 553
            }
            _ = len(v30)
            if false {
                switch int(v31) {
                case -152:
                    if ((2 == 2) && true) {
                        for fi48 := 0; fi48 < 16; fi48++ {
                            v49 := GenS6{F0: 32767, F1: -1, F2: v3}
                            _ = fi48
                            _ = v49
                        }
                        var v50 uint8 = 255
                        _ = v50
                    }
                case 0:
                    var v51 string = ""
                    _ = v51
                case 4294967295:
                    var v52 uint64 = 2
                    _ = v52
                default:
                    var v53 bool = v4
                    _ = v53
                }
                var v54 int64 = int64(v3)
                _ = v54
            } else {
                var v55 int32 = 0
                if (!false) {
                    if (42 == int8(v3)) {
                        if (((((float32(3.14159) * (v35 - 0.1)) + v35) * 3.14159) * ((((v35 - ((((float32(((((3.14159 - ((0.1 / float64(-2.25)) * -2.25)) / (float64((float32((0.1 - (-2.25 / 1e10))) + 0.0)) - 3.14159)) * (1.5 + 0.0)) - float64(v35))) + (v35 * v35)) / v35) - 1e10) + (v35 - v35))) - (v35 + 0.1)) * ((v35 * v35) * v35)) * 0.1)) == (-2.25 / v35)) {
                            v55 = v55
                            var v56 int8 = 78
                            fmt.Println(fi47)
                            _ = v56
                        } else {
                            if v4 {
                                for fi57 := 0; fi57 < 12; fi57++ {
                                    v31 = v31
                                    _ = fi57
                                }
                                v35 = ((-2.25 + float32(0.0)) * 0.0)
                            } else {
                                var v58 string = ""
                                _ = v58
                            }
                            v35 = 0.1
                            var v59 string = v32
                            _ = v59
                        }
                        var v60 float64 = float64(float32((0.1 - float64(v35))))
                        if (v31 > int64(v2)) {
                            var v61 int16 = -32768
                            for fi62 := 0; fi62 < 17; fi62++ {
                                if len(v30) >= 2 {
                                    v30 = v30[1:len(v30)]
                                }
                                _ = fi62
                            }
                            fmt.Println(v4, v37)
                            _ = v61
                        }
                        _ = v60
                    }
                } else {
                    v31 = (v31 << ((2) & 63))
                    v63 := GenS7{F0: v31}
                    v35 = float32((((1.5 - ((0.0 + 1.5) / ((1e10 * -2.25) + 0.0))) / float64((3.14159 / v35))) / float64(((((v35 - float32(float64((v35 - float32(0.1))))) * 3.14159) / (v35 + v35)) + ((((0.1 + ((((v35 + 0.1) * v35) + (3.14159 * 1.5)) - ((v35 / v35) + (v35 + (-2.25 + v35))))) + (v35 + v35)) + v35) + v35)))))
                    _ = v63
                }
                _ = v55
            }
            _ = fi47
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v35
    _ = v36
    _ = v37
    _ = v4
}
