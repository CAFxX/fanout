package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 uint16; F1 int; F2 string }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int32; F1 bool; F2 string }
type GenS6 struct { F0 bool; F1 int32; F2 uint32 }

func main() {
    var v1 bool = ((!(false && ((1e10 / (float32(((((3.14159 - 0.0) / float64(((1.5 / 1e10) * 1.5))) * -2.25) / 3.14159)) + 3.14159)) != ((float32(1.5) * (0.1 - (0.1 + -2.25))) * ((-2.25 - (float32(0.1) + 3.14159)) + (float32(-2.25) / 1e10)))))) && ((752 >= 0) && false))
    var v2 int = 128
    v3 := []string{"", "hello", "foo bar"}
    var v4 int8 = 0
    v5 := make(map[string]string)
    v6 := make(map[int]uint32)
    v6[(v2 ^ v2)] = 0
    fmt.Println(v2)
    switch v2 {
    case 32768:
        var v7 int8 = (v4 + 31)
        _ = v7
    case 128:
        fmt.Println(v4)
    }
    if (!(!(852 != 827))) {
        fmt.Println(v2, v1)
        var v8 int32 = 42
        v2 = v2
        _ = v8
    }
    var v9 bool = true
    v4 = -1
    for fi10 := 0; fi10 < 11; fi10++ {
        var v11 uint = 18446744073709551615
        v2 += v2
        v9 = true
        _ = fi10
        _ = v11
    }
    if v9 {
        for fi12 := 0; fi12 < 15; fi12++ {
            v2 = (v2 | v2)
            _ = fi12
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v4 = 1
    }
    for fi13 := 0; fi13 < 17; fi13++ {
        for fi14 := 0; fi14 < 20; fi14++ {
            var v15 int = (^(v2))
            _ = fi14
            _ = v15
        }
        _ = fi13
    }
    if (((!(!(float32(-2.25) != ((float32(float64(-2.25)) / 1e10) * 3.14159)))) && (!(-9223372036854775808 <= 0))) || (-1 < v4)) {
        if (831 == 0) {
            fmt.Println(v9, v1, v2)
            if (!(!false)) {
                for fi16 := 0; fi16 < 5; fi16++ {
                    fmt.Println(fi16, v9, fi16)
                    for fi17 := 0; fi17 < 11; fi17++ {
                        v3 = append(v3, "")
                        if false {
                            v9 = v9
                        } else {
                            if false {
                                if len(v3) > 0 {
                                    v3[0] = ("" + "\n\t")
                                }
                            }
                            v18 := GenS1{F0: v2}
                            _ = v18
                        }
                        v2 -= v2
                        _ = fi17
                    }
                    v19 := v5[""]
                    _ = fi16
                    _ = v19
                }
                _ = len(v3)
                for fi20 := 0; fi20 < 13; fi20++ {
                    v4 = int8(fi20)
                    if (4294967295 <= 0) {
                        delete(v6, v2)
                        if v1 {
                            delete(v5, "foo bar")
                        } else {
                            delete(v6, 32768)
                            v21 := GenS1{F0: fi20}
                            _ = v21
                        }
                        fmt.Println(v1, fi20)
                    } else {
                        if false {
                            var v22 string = "hello"
                            if v1 {
                                var v23 int32 = 636
                                switch (v22 + "") {
                                case "a":
                                    v24 := GenS2{F0: 65535, F1: 111}
                                    _ = v24
                                case "\n\t":
                                    var v25 int16 = 32767
                                    _ = v25
                                case "":
                                    switch ("\n\t" + (v22 + ("" + "a"))) {
                                    case "a":
                                        if false {
                                            v26 := GenS3{F0: 0, F1: 32768, F2: "hello"}
                                            _ = v26
                                        } else {
                                            v2 = (fi20 % ((fi20) | 1))
                                            var v27 float32 = 1e10
                                            v28 := GenS1{F0: -1}
                                            _ = v27
                                            _ = v28
                                        }
                                    case "foo bar":
                                        v22 = v22
                                    default:
                                        switch (v22 + v22) {
                                        case "hello":
                                            if true {
                                                if (2 == 1) {
                                                    v22 = "hello"
                                                } else {
                                                    for fi29 := 0; fi29 < 10; fi29++ {
                                                        var v30 float32 = 3.14159
                                                        var v31 int64 = -1
                                                        v9 = (-313 < (v23 ^ 2147483647))
                                                        _ = fi29
                                                        _ = v30
                                                        _ = v31
                                                    }
                                                }
                                                v32 := GenS2{F0: 65535, F1: 1}
                                                _ = v32
                                            } else {
                                                switch "\n\t" {
                                                case "a":
                                                    for fi33 := 0; fi33 < 7; fi33++ {
                                                        switch v22 {
                                                        case "\n\t":
                                                            fmt.Println(v23, v4)
                                                        case "hello":
                                                            switch fi20 {
                                                            case 32767:
                                                                var v34 int16 = -229
                                                                _ = v34
                                                            default:
                                                                var v35 int = 32768
                                                                _ = v35
                                                            }
                                                        }
                                                        v36 := GenS4{F0: 161}
                                                        v22 = ("hello" + v22)
                                                        _ = fi33
                                                        _ = v36
                                                    }
                                                }
                                                v9 = v1
                                                fmt.Println(v22, v4, v22)
                                            }
                                        case "":
                                            v4 -= v4
                                        }
                                    }
                                default:
                                    for fi37 := 0; fi37 < 9; fi37++ {
                                        if v1 {
                                            fmt.Println(v1, fi37)
                                            v38 := GenS1{F0: (-(fi37))}
                                            v3 = append(v3, v22)
                                            _ = v38
                                        } else {
                                            v2 += (v2 % ((-90) | 1))
                                        }
                                        v39 := GenS2{F0: 65535, F1: 0}
                                        if ("hello" <= v22) {
                                            fmt.Println(v23)
                                            v1 = v9
                                            if (!(v22 >= v22)) {
                                                fmt.Println(fi20, v22)
                                                v23 = (-(v23))
                                            }
                                        } else {
                                            delete(v6, (fi37 | v2))
                                            v2 += 2147483647
                                        }
                                        _ = fi37
                                        _ = v39
                                    }
                                }
                                var v40 int64 = -952
                                _ = v23
                                _ = v40
                            }
                            _ = v22
                        } else {
                            v9 = v9
                        }
                    }
                    v1 = v9
                    _ = fi20
                }
            }
            var v41 string = ""
            _ = v41
        } else {
            v4 += -2
            if (true || true) {
                switch (v2 + 931) {
                case -9223372036854775808:
                    _ = len(v5)
                case 9223372036854775807:
                    _ = len(v6)
                case -796:
                    var v42 uint16 = 32767
                    _ = v42
                }
            }
        }
        v43 := GenS3{F0: 65535, F1: (v2 - (v2 << ((2) & 63))), F2: "foo bar"}
        _ = v43
    } else {
        if (!((((3.14159 + float64((-2.25 - (-2.25 / 3.14159)))) - 0.1) * -2.25) <= float64(0.1))) {
            v44 := GenS2{F0: 1, F1: 127}
            if v9 {
                v3 = append(v3, "hello")
                v45 := GenS4{F0: 2}
                v2 -= (v2 ^ (v2 / ((-1) | 1)))
                _ = v45
            } else {
                for fi46 := 0; fi46 < 1; fi46++ {
                    if v1 {
                        var v47 uint8 = 210
                        _ = v47
                    }
                    v4 = (v4 % ((v4) | 1))
                    _ = fi46
                }
                if v9 {
                    v48 := GenS5{F0: -559, F1: (42 == 1), F2: "hello"}
                    v5[""] = ("hello" + (("" + ("" + "a")) + ""))
                    v4 = v4
                    _ = v48
                } else {
                    v4 += (v4 ^ int8(v2))
                }
            }
            var v49 uint8 = 121
            _ = v44
            _ = v49
        } else {
            if (84 != v4) {
                fmt.Println(v2, v1, v9)
                switch "hello" {
                case "\n\t":
                    v50 := GenS4{F0: 42}
                    _ = v50
                default:
                    if len(v3) > 0 {
                        v3[0] = (("foo bar" + (("hello" + "foo bar") + "")) + "foo bar")
                    }
                }
            }
        }
        var v51 float64 = 1e10
        v52 := GenS5{F0: 240, F1: (v1 || true), F2: ""}
        _ = v51
        _ = v52
    }
    fmt.Println(v4)
    var v53 int8 = v4
    fmt.Println(v2)
    v3 = append(v3, "hello")
    fmt.Println(v2, v53, v1)
    v1 = v1
    var v54 uint8 = 240
    for fi55 := 0; fi55 < 3; fi55++ {
        var v56 int64 = -9223372036854775808
        if (false && ((("foo bar" + "a") < "foo bar") && (v4 <= (v4 * (^(v4)))))) {
            fmt.Println(v56, v53)
            v57 := GenS6{F0: (v54 > v54), F1: -812, F2: 42}
            delete(v6, 127)
            _ = v57
        }
        v54 = (v54 & (v54 >> ((709) & 63)))
        _ = fi55
        _ = v56
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v6
    _ = v9
}
