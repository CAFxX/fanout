package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 string; F1 string; F2 int8 }
type GenS4 struct { F0 int8 }
func genFn0(p0 int8, p1 int8, p2 bool) int32 {
    var v6 uint16 = 255
    fmt.Println(p0, p1, v6)
    if (float32((-2.25 - ((0.1 / -2.25) - 0.1))) != (3.14159 / 1.5)) {
        var v7 float64 = 0.1
        var v8 uint32 = 4294967295
        var v9 int64 = 0
        _ = v7
        _ = v8
        _ = v9
    } else {
        fmt.Println(p0, p1, p2)
    }
    switch "foo bar" {
    case "a":
        fmt.Println(p0)
    case "hello":
        for fi10 := 0; fi10 < 19; fi10++ {
            for fi11 := 0; fi11 < 6; fi11++ {
                switch fi11 {
                case -593:
                    v12 := GenS1{F0: (true || p2), F1: p2}
                    _ = v12
                default:
                    p2 = p2
                }
                var v13 int64 = 0
                _ = fi11
                _ = v13
            }
            _ = fi10
        }
    default:
        p1 += p0
    }
    var v14 int8 = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v6
    return 57
}
func genFn1(p0 int16, p1 int8, p2 uint) int16 {
    v15 := []int{32768, 4294967295, -9223372036854775808, 1}
    if true {
        for fi16 := 0; fi16 < 5; fi16++ {
            var v17 int16 = 1
            _ = fi16
            _ = v17
        }
        var v18 string = "foo bar"
        if len(v15) > 0 {
            v15[0] = 4294967295
        }
        _ = v18
    } else {
        switch -577 {
        case 42:
            p0 += p0
        case 128:
            v19 := GenS1{F0: ((false || ((3.14159 + float32((0.0 * 0.1))) != (float32((0.0 + (1e10 + (1e10 + 0.0)))) - (0.0 - (float32(1e10) * 0.1))))) && (p1 == -32)), F1: true}
            _ = v19
        }
        if true {
            v20 := GenS2{F0: 255, F1: 0}
            _ = v20
        } else {
            if false {
                p1 += p1
                v21 := GenS2{F0: 65535, F1: 0}
                _ = v21
            } else {
                var v22 int64 = 9223372036854775807
                _ = v22
            }
            if ("a" < "a") {
                switch 559 {
                case 9223372036854775807:
                    var v23 int8 = p1
                    _ = v23
                default:
                    for fi24 := 0; fi24 < 15; fi24++ {
                        if (true && (!true)) {
                            var v25 int8 = (p1 % (((p1 * p1)) | 1))
                            for fi26 := 0; fi26 < 17; fi26++ {
                                fmt.Println(p1, fi26)
                                p2 = p2
                                p0 -= 1000
                                _ = fi26
                            }
                            if (false || true) {
                                for fi27 := 0; fi27 < 11; fi27++ {
                                    fmt.Println(p0, v25, p1)
                                    _ = fi27
                                }
                                v28 := GenS1{F0: false, F1: false}
                                var v29 uint16 = 65535
                                _ = v28
                                _ = v29
                            }
                            _ = v25
                        } else {
                            var v30 string = ((("hello" + "\n\t") + "foo bar") + "a")
                            for fi31 := 0; fi31 < 6; fi31++ {
                                if false {
                                    if true {
                                        var v32 float64 = 1.5
                                        var v33 float64 = float64(3.14159)
                                        p0 += p0
                                        var v34 float32 = (1e10 + float32(((v32 + v33) * (v33 / 3.14159))))
                                        if (true || true) {
                                            p2 = 65535
                                            if len(v15) > 0 {
                                                v15[0] = (fi24 % ((fi24) | 1))
                                            }
                                            v34 = (float32(float64(v34)) * (v34 / (-2.25 + -2.25)))
                                        }
                                        _ = v32
                                        _ = v33
                                        _ = v34
                                    } else {
                                        v35 := GenS1{F0: ((fi31 + fi24) >= (fi31 >> ((p2) & 63))), F1: false}
                                        if len(v15) >= 2 {
                                            v15 = v15[1:len(v15)]
                                        }
                                        p0 -= p0
                                        _ = v35
                                    }
                                    p1 = p1
                                    var v36 int = (fi24 + (-(fi24)))
                                    v37 := GenS3{F0: (((v30 + v30) + v30) + (((v30 + "a") + v30) + ("hello" + (v30 + (v30 + v30))))), F1: ((v30 + "\n\t") + (v30 + ("hello" + "hello"))), F2: p1}
                                    _ = v36
                                    _ = v37
                                } else {
                                    fmt.Println(fi24)
                                }
                                var v38 float64 = float64(float32(((1.5 * (0.1 - 3.14159)) * 1e10)))
                                _ = fi31
                                _ = v38
                            }
                            v30 = ("" + v30)
                            _ = v30
                        }
                        v39 := GenS4{F0: 100}
                        if len(v15) > 0 {
                            v15[0] = fi24
                        }
                        _ = fi24
                        _ = v39
                    }
                }
                p0 += (p0 + p0)
                v40 := GenS4{F0: -97}
                _ = v40
            }
            var v41 int32 = 0
            _ = v41
        }
        p2 += p2
    }
    var v42 uint16 = 652
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v42
    return (p0 / (((p0 / ((32767) | 1))) | 1))
}

func main() {
    var v1 string = "\n\t"
    var v2 uint32 = 42
    var v3 int32 = -473
    v4 := make(map[int]int)
    var v5 uint64 = 1
    fmt.Println(v2, v3)
    var v43 int32 = (^(v3))
    v44 := GenS4{F0: -1}
    var v45 int8 = -14
    for fi46 := 0; fi46 < 8; fi46++ {
        if (!(571 >= 672)) {
            v47 := GenS1{F0: false, F1: (1 <= 127)}
            if (!(!false)) {
                var v48 float32 = 1e10
                switch -9223372036854775808 {
                case 128:
                    if (!true) {
                        _ = len(v4)
                        v43 += v43
                        var v49 int64 = 0
                        _ = v49
                    }
                }
                v48 = (v48 * float32((1.5 / ((1.5 + ((0.1 - (0.0 + 1e10)) * 1e10)) / (3.14159 * (0.1 - 1e10))))))
                _ = v48
            }
            _ = v47
        }
        v4[-588] = (fi46 + fi46)
        _ = fi46
    }
    var v50 float32 = 1.5
    v51 := []int{32768, 834}
    fmt.Println(v5, v3, v45)
    if (true || (!(((true && false) || true) && (!true)))) {
        var v52 int8 = v45
        for fi53 := 0; fi53 < 15; fi53++ {
            v4[fi53] = fi53
            v1 = (v1 + "foo bar")
            _ = fi53
        }
        _ = v52
    }
    v2 += (v2 / ((v2) | 1))
    var v54 bool = false
    if v54 {
        for fi55 := 0; fi55 < 1; fi55++ {
            var v56 int = (fi55 / (((-(fi55))) | 1))
            _ = fi55
            _ = v56
        }
        delete(v4, 9223372036854775807)
    }
    _ = len(v4)
    fmt.Println(v2, v5, v2)
    var v57 int = 1
    if v54 {
        var v58 string = v1
        if len(v51) > 0 {
            v51[0] = v57
        }
        for fi59 := 0; fi59 < 17; fi59++ {
            if (0 == 255) {
                var v60 int8 = -128
                switch v1 {
                case "foo bar":
                    if len(v51) >= 2 {
                        v51 = v51[1:len(v51)]
                    }
                default:
                    v61 := v4[(fi59 & (v57 * -279))]
                    _ = v61
                }
                _ = v60
            }
            delete(v4, (^(v57)))
            _ = fi59
        }
        _ = v58
    }
    for fi62 := 0; fi62 < 7; fi62++ {
        fmt.Println(v45)
        var v63 uint64 = v5
        for fi64 := 0; fi64 < 4; fi64++ {
            if v54 {
                var v65 bool = ("\n\t" == v1)
                if (v54 && false) {
                    switch "\n\t" {
                    case "a":
                        v45 = (^(v45))
                    case "\n\t":
                        switch 2147483647 {
                        case 0:
                            var v66 int64 = 1
                            _ = v66
                        default:
                            var v67 string = v1
                            _ = v67
                        }
                    case "foo bar":
                        if len(v51) >= 2 {
                            v51 = v51[1:len(v51)]
                        }
                    default:
                        fmt.Println(v1)
                    }
                }
                v68 := GenS1{F0: v65, F1: (!(v54 || v65))}
                _ = v65
                _ = v68
            }
            v69 := v4[v57]
            fmt.Println(v63, v5)
            _ = fi64
            _ = v69
        }
        _ = fi62
        _ = v63
    }
    v70 := v4[127]
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("slicelen v51:", len(v51), v51)
    fmt.Println("v54:", v54)
    fmt.Println("v57:", v57)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v50
    _ = v51
    _ = v54
    _ = v57
    _ = v70
}
