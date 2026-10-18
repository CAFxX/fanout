package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string; F1 bool }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 bool }
func genFn0() uint8 {
    for fi6 := 0; fi6 < 17; fi6++ {
        switch (-(fi6)) {
        case 42:
            v7 := []int{(-(fi6)), fi6}
            _ = v7
        }
        var v8 int32 = -1
        _ = fi6
        _ = v8
    }
    v9 := GenS1{F0: 1, F1: 0, F2: -887}
    for fi10 := 0; fi10 < 8; fi10++ {
        var v11 int64 = 1
        v11 += 9223372036854775807
        v12 := []int{(^(fi10)), fi10, -339}
        _ = fi10
        _ = v11
        _ = v12
    }
    switch "hello" {
    case "a":
        var v13 int32 = -1
        _ = v13
    case "\n\t":
        var v14 int32 = -459
        _ = v14
    }
    _ = v9
    return 2
}
func genFn1(p0 string) string {
    v15 := []int{-9223372036854775808, -127, 4294967295, 9}
    p0 = "hello"
    _ = p0
    _ = v15
    return ("\n\t" + "\n\t")
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[string]string)
    var v3 uint = 1
    v4 := []string{v1}
    var v5 float32 = 0.0
    v16 := GenS1{F0: 0, F1: -183, F2: 1}
    if (128 >= 1) {
        v17 := GenS1{F0: 2147483647, F1: 166, F2: 7}
        v18 := GenS2{F0: 42}
        for fi19 := 0; fi19 < 20; fi19++ {
            fmt.Println(v1)
            v2[((v1 + (("a" + (v1 + v1)) + (("hello" + v1) + ((("\n\t" + "hello") + ((v1 + "hello") + "hello")) + "\n\t")))) + ("hello" + (("" + v1) + "")))] = v1
            if ((1 != (v3 * 594)) && (!true)) {
                var v20 uint = (v3 << ((v3) & 63))
                if true {
                    v5 += 1e10
                    v2[(v1 + (v1 + "a"))] = "a"
                    var v21 int = 42
                    _ = v21
                }
                _ = v20
            } else {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v3 -= 255
            }
            _ = fi19
        }
        _ = v17
        _ = v18
    } else {
        v1 = v1
        switch ((v1 + (((v1 + v1) + v1) + "")) + "a") {
        case "hello":
            v22 := v2[(v1 + "")]
            _ = v22
        case "\n\t":
            for fi23 := 0; fi23 < 13; fi23++ {
                for fi24 := 0; fi24 < 4; fi24++ {
                    v1 = v1
                    v5 = -2.25
                    _ = fi24
                }
                _ = fi23
            }
        default:
            if (true || (int16(v3) < 674)) {
                if (v3 <= v3) {
                    v25 := GenS2{F0: 42}
                    v26 := GenS3{F0: false}
                    _ = v25
                    _ = v26
                }
                for fi27 := 0; fi27 < 5; fi27++ {
                    v5 += -2.25
                    var v28 int = (^(fi27))
                    switch v1 {
                    case "\n\t":
                        if len(v4) > 0 {
                            v4[0] = ""
                        }
                    }
                    _ = fi27
                    _ = v28
                }
                fmt.Println(v3)
            } else {
                for fi29 := 0; fi29 < 11; fi29++ {
                    var v30 uint64 = 485
                    _ = fi29
                    _ = v30
                }
                v3 = v3
                var v31 float32 = -2.25
                _ = v31
            }
        }
        v32 := GenS3{F0: false}
        _ = v32
    }
    if (202 != 268) {
        switch "" {
        case "":
            v5 = v5
        case "\n\t":
            v33 := GenS4{F0: v1, F1: false}
            _ = v33
        default:
            for fi34 := 0; fi34 < 13; fi34++ {
                switch "\n\t" {
                case "a":
                    v35 := GenS2{F0: 42}
                    _ = v35
                case "":
                    v5 -= v5
                case "\n\t":
                    var v36 uint = (v3 ^ 0)
                    _ = v36
                default:
                    v5 -= (0.0 + (v5 / v5))
                }
                _ = len(v4)
                _ = fi34
            }
        }
        v1 = v1
    } else {
        if true {
            var v37 int16 = 753
            if (0 == 2147483648) {
                for fi38 := 0; fi38 < 9; fi38++ {
                    var v39 float32 = 1e10
                    var v40 uint32 = 4294967295
                    _ = fi38
                    _ = v39
                    _ = v40
                }
            } else {
                fmt.Println(v3, v1)
                v3 -= 1
                switch 128 {
                case 32768:
                    v4 = append(v4, v1)
                case -9223372036854775808:
                    v5 = v5
                default:
                    v3 += (v3 >> (((^(v3))) & 63))
                }
            }
            v41 := GenS5{F0: (^(v3))}
            _ = v37
            _ = v41
        }
        fmt.Println(v5, v5)
        switch v1 {
        case "a":
            var v42 int64 = 292
            _ = v42
        case "hello":
            fmt.Println(v3, v1, v3)
        case "foo bar":
            fmt.Println(v5, v5, v5)
        }
    }
    fmt.Println(v5, v1, v1)
    if false {
        fmt.Println(v5)
    } else {
        switch 0 {
        case 32768:
            switch "" {
            case "":
                fmt.Println(v3, v5, v5)
            case "hello":
                if true {
                    switch -1 {
                    case 262:
                        if false {
                            for fi43 := 0; fi43 < 18; fi43++ {
                                if (23 >= 1) {
                                    var v44 int16 = 42
                                    _ = v44
                                } else {
                                    switch v1 {
                                    case "":
                                        var v45 int = fi43
                                        _ = v45
                                    }
                                    if true {
                                        v5 = (v5 / v5)
                                    } else {
                                        switch ((v1 + v1) + "\n\t") {
                                        case "\n\t":
                                            v5 = (((v5 / v5) * float32(-2.25)) + (v5 + ((1.5 - float32(float64(-2.25))) / float32(float64(v5)))))
                                        default:
                                            for fi46 := 0; fi46 < 7; fi46++ {
                                                var v47 int16 = -277
                                                _ = fi46
                                                _ = v47
                                            }
                                        }
                                        for fi48 := 0; fi48 < 11; fi48++ {
                                            var v49 int32 = 1
                                            _ = fi48
                                            _ = v49
                                        }
                                        var v50 uint64 = 0
                                        _ = v50
                                    }
                                    v5 -= 1e10
                                }
                                v2[v1] = v1
                                _ = fi43
                            }
                        }
                    case 2147483647:
                        for fi51 := 0; fi51 < 19; fi51++ {
                            var v52 uint = v3
                            _ = fi51
                            _ = v52
                        }
                    case 9223372036854775807:
                        v53 := GenS6{F0: false}
                        _ = v53
                    }
                    if false {
                        var v54 int = 9223372036854775807
                        for fi55 := 0; fi55 < 12; fi55++ {
                            v56 := v2[v1]
                            _ = fi55
                            _ = v56
                        }
                        _ = v54
                    } else {
                        switch v1 {
                        case "":
                            var v57 string = (("" + (((v1 + (v1 + "foo bar")) + v1) + ("" + v1))) + v1)
                            _ = v57
                        case "hello":
                            fmt.Println(v5, v3, v5)
                        default:
                            if false {
                                v3 -= (v3 % (((v3 | v3)) | 1))
                                v58 := GenS6{F0: (1 == 18446744073709551615)}
                                _ = v58
                            } else {
                                switch ("a" + v1) {
                                case "foo bar":
                                    var v59 uint64 = 2
                                    _ = v59
                                case "a":
                                    var v60 float32 = (1e10 * (v5 + 3.14159))
                                    _ = v60
                                }
                                fmt.Println(v3, v1)
                            }
                        }
                        var v61 int8 = 78
                        v62 := v2[v1]
                        _ = v61
                        _ = v62
                    }
                    v3 -= v3
                }
            case "a":
                var v63 float64 = float64((1.5 / v5))
                _ = v63
            default:
                v64 := GenS4{F0: "a", F1: (!(!(true && false)))}
                _ = v64
            }
        case 32767:
            v3 -= v3
        default:
            fmt.Println(v1, v3, v5)
        }
        v1 = ""
    }
    var v65 float32 = v5
    _ = len(v4)
    var v66 int8 = -127
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v65
    _ = v66
}
