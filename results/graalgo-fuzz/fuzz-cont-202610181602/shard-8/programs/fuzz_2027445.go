package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 string; F1 uint64; F2 bool }
func genFn0() bool {
    var v5 float32 = 1e10
    if false {
        var v6 int64 = 42
        _ = v6
    } else {
        var v7 int32 = 0
        var v8 string = "hello"
        var v9 uint16 = 87
        _ = v7
        _ = v8
        _ = v9
    }
    v5 -= (float32(-2.25) / 1.0)
    v5 = (((v5 + (3.14159 + v5)) / v5) - (v5 + 1e10))
    var v10 uint8 = 128
    _ = v10
    _ = v5
    return true
}
func genFn1(p0 string, p1 uint16) bool {
    var v11 string = (p0 + (p0 + (p0 + "")))
    if false {
        var v12 int = 2147483647
        fmt.Println(p1, p1, v12)
        if false {
            v12 += 0
            switch "hello" {
            case "hello":
                var v13 int32 = -1
                _ = v13
            case "\n\t":
                if true {
                    v14 := GenS1{F0: p0}
                    _ = v14
                }
            case "":
                v15 := []int{(v12 + (^(v12))), v12}
                _ = v15
            }
        } else {
            fmt.Println(v12, p0)
            switch 128 {
            case -1:
                fmt.Println(v11)
            case 127:
                for fi16 := 0; fi16 < 10; fi16++ {
                    switch v11 {
                    case "\n\t":
                        if (0.1 == 0.1) {
                            v11 = p0
                        }
                    case "a":
                        if true {
                            v17 := GenS1{F0: v11}
                            for fi18 := 0; fi18 < 7; fi18++ {
                                v12 += (fi16 - fi18)
                                v19 := GenS1{F0: ""}
                                v20 := GenS1{F0: (p0 + "foo bar")}
                                _ = fi18
                                _ = v19
                                _ = v20
                            }
                            if false {
                                for fi21 := 0; fi21 < 7; fi21++ {
                                    v22 := GenS1{F0: p0}
                                    if false {
                                        fmt.Println(fi21, fi16)
                                        v12 += (v12 - v12)
                                        switch fi21 {
                                        case -754:
                                            var v23 float32 = 3.14159
                                            _ = v23
                                        case 32767:
                                            v24 := GenS2{F0: 127, F1: 0, F2: 0}
                                            _ = v24
                                        }
                                        v25 := GenS1{F0: v11}
                                        _ = v25
                                    } else {
                                        var v26 float32 = -2.25
                                        _ = v26
                                    }
                                    var v27 uint64 = 0
                                    _ = fi21
                                    _ = v22
                                    _ = v27
                                }
                            } else {
                                for fi28 := 0; fi28 < 8; fi28++ {
                                    v29 := GenS3{F0: false}
                                    var v30 string = ("" + p0)
                                    _ = fi28
                                    _ = v29
                                    _ = v30
                                }
                                switch fi16 {
                                case -720:
                                    var v31 int64 = 9223372036854775807
                                    _ = v31
                                case 2147483647:
                                    if (728 > 4294967295) {
                                        var v32 bool = false
                                        _ = v32
                                    } else {
                                        v33 := GenS1{F0: (p0 + "foo bar")}
                                        v11 = "foo bar"
                                        var v34 int64 = 352
                                        _ = v33
                                        _ = v34
                                    }
                                case 297:
                                    switch fi16 {
                                    case 32768:
                                        v12 += (fi16 % (((fi16 * (fi16 % (((fi16 ^ (v12 - fi16))) | 1)))) | 1))
                                    case 32767:
                                        var v35 uint = 336
                                        _ = v35
                                    case 349:
                                        var v36 bool = (-2147483648 == -138)
                                        _ = v36
                                    default:
                                        for fi37 := 0; fi37 < 7; fi37++ {
                                            switch fi16 {
                                            case -258:
                                                var v38 float64 = 3.14159
                                                _ = v38
                                            case 1:
                                                v39 := GenS2{F0: 127, F1: 239, F2: 1}
                                                _ = v39
                                            case 41:
                                                fmt.Println(v12)
                                            }
                                            fmt.Println(p1, fi16)
                                            v40 := []int{}
                                            _ = fi37
                                            _ = v40
                                        }
                                    }
                                }
                                var v41 uint16 = (p1 / ((p1) | 1))
                                _ = v41
                            }
                            _ = v17
                        }
                    }
                    _ = fi16
                }
            case -2147483648:
                p0 = "\n\t"
            default:
                p0 = "a"
            }
        }
        _ = v12
    } else {
        v42 := []int{9223372036854775807, 2147483647}
        var v43 uint8 = 128
        _ = v42
        _ = v43
    }
    for fi44 := 0; fi44 < 1; fi44++ {
        var v45 int16 = 208
        p1 -= p1
        _ = fi44
        _ = v45
    }
    var v46 float64 = ((float64((float32(1e10) * ((-2.25 - 3.14159) * (0.0 - 1e10)))) / -2.25) / (float64((1e10 - 1.5)) * ((3.14159 + (((0.1 + ((0.1 / 3.14159) / 0.1)) / 1.0) + 0.0)) / 1.0)))
    _ = p0
    _ = p1
    _ = v11
    _ = v46
    return (4294967295 < 32767)
}

func main() {
    var v1 uint16 = 962
    var v2 bool = (!(!(949 == 128)))
    v3 := []string{"a", ("\n\t" + "\n\t"), (((("hello" + "") + ("a" + (("foo bar" + "hello") + "foo bar"))) + ("foo bar" + ("hello" + ("" + "foo bar")))) + "hello"), ("" + "\n\t")}
    var v4 uint16 = v1
    var v47 int64 = 42
    v4 -= 0
    if v2 {
        v48 := GenS4{F0: 594}
        switch (("a" + "a") + "hello") {
        case "\n\t":
            var v49 uint64 = 0
            _ = v49
        case "hello":
            fmt.Println(v47)
        case "a":
            fmt.Println(v47)
        }
        for fi50 := 0; fi50 < 4; fi50++ {
            v51 := GenS4{F0: 42}
            v2 = (v4 >= v1)
            v3 = append(v3, "hello")
            _ = fi50
            _ = v51
        }
        _ = v48
    } else {
        v1 -= v4
        var v52 uint8 = 255
        _ = v52
    }
    for fi53 := 0; fi53 < 4; fi53++ {
        if v2 {
            v54 := GenS2{F0: 1, F1: 0, F2: uint8(v1)}
            for fi55 := 0; fi55 < 16; fi55++ {
                var v56 int = -733
                _ = fi55
                _ = v56
            }
            _ = v54
        } else {
            v47 = v47
            v4 = 1
            if false {
                var v57 string = ("hello" + "foo bar")
                var v58 int = (fi53 * 32767)
                _ = v57
                _ = v58
            }
        }
        switch fi53 {
        case 496:
            v59 := GenS5{F0: "\n\t", F1: 2, F2: false}
            _ = v59
        case -477:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case 32767:
            if len(v3) > 0 {
                v3[0] = "\n\t"
            }
        }
        if (true && (fi53 != (fi53 % ((-9223372036854775808) | 1)))) {
            v60 := GenS4{F0: 110}
            var v61 float32 = -2.25
            var v62 uint16 = (v4 ^ (v4 ^ v4))
            _ = v60
            _ = v61
            _ = v62
        }
        _ = fi53
    }
    switch -942 {
    case 32767:
        var v63 float64 = 0.1
        _ = v63
    default:
        _ = len(v3)
    }
    for fi64 := 0; fi64 < 6; fi64++ {
        v4 += (v1 / ((2) | 1))
        var v65 string = "a"
        _ = fi64
        _ = v65
    }
    var v66 string = "\n\t"
    for fi67 := 0; fi67 < 4; fi67++ {
        if true {
            var v68 int = 1
            _ = v68
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v2 = false
        }
        var v69 string = v66
        v70 := GenS5{F0: ("hello" + v69), F1: 4294967295, F2: (false && true)}
        _ = fi67
        _ = v69
        _ = v70
    }
    var v71 int = 128
    fmt.Println(v66, v66, v66)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v66:", v66)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v66
    _ = v71
}
