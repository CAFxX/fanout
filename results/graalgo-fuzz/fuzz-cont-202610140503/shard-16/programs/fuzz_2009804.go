package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint32; F1 uint }
func genFn0() string {
    v6 := []int{-689}
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 int32 = -2147483648
        var v9 float32 = (((0.0 + 1.5) + 0.1) - 1.5)
        if true {
            switch "foo bar" {
            case "\n\t":
                v9 -= -2.25
            case "":
                for fi10 := 0; fi10 < 8; fi10++ {
                    fmt.Println(v8)
                    if false {
                        if true {
                            var v11 float32 = (0.1 + v9)
                            _ = len(v6)
                            _ = v11
                        } else {
                            v12 := GenS1{F0: -1, F1: 0}
                            fmt.Println(fi7, v9)
                            var v13 int64 = 57
                            _ = v12
                            _ = v13
                        }
                        fmt.Println(v9, fi10, fi10)
                        fmt.Println(v8, fi7, fi10)
                    } else {
                        var v14 uint16 = 0
                        _ = v14
                    }
                    v9 -= v9
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    _ = fi10
                }
            default:
                if false {
                    var v15 string = "\n\t"
                    v15 = ""
                    _ = v15
                } else {
                    fmt.Println(v9)
                }
            }
        } else {
            var v16 int8 = -1
            v6 = append(v6, fi7)
            _ = v16
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    var v17 int64 = -50
    _ = v17
    _ = v6
    return ""
}
func genFn1(p0 int32, p1 int16) int {
    if (0.0 >= 0.0) {
        for fi18 := 0; fi18 < 19; fi18++ {
            for fi19 := 0; fi19 < 9; fi19++ {
                var v20 int8 = 127
                p1 -= 42
                _ = fi19
                _ = v20
            }
            v21 := []int{fi18, fi18}
            _ = fi18
            _ = v21
        }
        if (42 < 255) {
            for fi22 := 0; fi22 < 1; fi22++ {
                for fi23 := 0; fi23 < 20; fi23++ {
                    p0 = (-(p0))
                    _ = fi23
                }
                var v24 uint8 = 0
                _ = fi22
                _ = v24
            }
        } else {
            var v25 string = (("\n\t" + "") + ("foo bar" + "hello"))
            for fi26 := 0; fi26 < 1; fi26++ {
                fmt.Println(v25, fi26)
                var v27 bool = (!true)
                fmt.Println(p0, p1, v25)
                _ = fi26
                _ = v27
            }
            if (v25 <= "hello") {
                p0 = -1
            } else {
                fmt.Println(p0, p0, v25)
                for fi28 := 0; fi28 < 19; fi28++ {
                    p0 += p0
                    _ = fi28
                }
                for fi29 := 0; fi29 < 11; fi29++ {
                    for fi30 := 0; fi30 < 8; fi30++ {
                        v31 := []int{722, fi30}
                        v25 = (("foo bar" + v25) + "a")
                        switch fi29 {
                        case 809:
                            switch v25 {
                            case "":
                                v25 = v25
                            case "foo bar":
                                var v32 int64 = 0
                                _ = v32
                            default:
                                if (2 <= 0) {
                                    if len(v31) > 0 {
                                        v31[0] = 210
                                    }
                                    v31 = append(v31, 863)
                                    fmt.Println(p1)
                                } else {
                                    p1 -= (p1 - (p1 | 42))
                                }
                            }
                            switch ((v25 + v25) + v25) {
                            case "foo bar":
                                if false {
                                    v25 = ""
                                    var v33 bool = true
                                    switch fi29 {
                                    case -2147483648:
                                        if len(v31) > 0 {
                                            v31[0] = 0
                                        }
                                    case -9223372036854775808:
                                        if v33 {
                                            p1 = (p1 ^ (p1 % ((362) | 1)))
                                            v25 = "a"
                                            v33 = false
                                        }
                                    }
                                    _ = v33
                                }
                            }
                        case -2147483648:
                            if (839 > 488) {
                                switch v25 {
                                case "foo bar":
                                    switch -9223372036854775808 {
                                    case 4294967295:
                                        fmt.Println(p1)
                                    default:
                                        var v34 uint = 1
                                        _ = v34
                                    }
                                case "a":
                                    var v35 float64 = (-2.25 + 0.1)
                                    _ = v35
                                case "hello":
                                    fmt.Println(fi29, p0, fi29)
                                }
                                v25 = ((("a" + v25) + ("" + ("\n\t" + v25))) + ("a" + v25))
                                p1 += (p1 / ((p1) | 1))
                            }
                            p1 = 1000
                        case 0:
                            v31 = append(v31, -9223372036854775808)
                        }
                        p1 += (p1 & -1)
                        v25 = v25
                        _ = fi30
                        _ = v31
                    }
                    v36 := GenS2{F0: v25}
                    v37 := GenS1{F0: 9223372036854775807, F1: 854}
                    _ = fi29
                    _ = v36
                    _ = v37
                }
            }
            _ = v25
        }
    }
    p0 = p0
    fmt.Println(p0, p1)
    var v38 float32 = (0.1 / 3.14159)
    _ = p0
    _ = p1
    _ = v38
    return 4294967295
}

func main() {
    var v1 uint32 = 2147483647
    v2 := make(map[uint32]uint32)
    var v3 int64 = 968
    var v4 uint32 = v1
    v5 := []string{}
    v1 = v1
    v39 := GenS3{F0: 42, F1: 162}
    if len(v5) > 0 {
        v5[0] = ""
    }
    _ = len(v2)
    for fi40 := 0; fi40 < 5; fi40++ {
        var v41 int8 = 0
        var v42 int64 = -254
        for fi43 := 0; fi43 < 10; fi43++ {
            switch 127 {
            case 9223372036854775807:
                v44 := GenS4{F0: ""}
                _ = v44
            case 4294967295:
                v2[(v4 ^ v4)] = (^(v4))
            default:
                var v45 float32 = 1.5
                v3 = v42
                var v46 int8 = (v41 & v41)
                _ = v45
                _ = v46
            }
            fmt.Println(v4, v41, v42)
            v1 += 1
            v47 := GenS5{F0: 175}
            v48 := GenS6{F0: v4}
            _ = fi43
            _ = v47
            _ = v48
        }
        _ = fi40
        _ = v41
        _ = v42
    }
    var v49 int16 = 1
    for fi50 := 0; fi50 < 4; fi50++ {
        var v51 uint = 65535
        _ = fi50
        _ = v51
    }
    v49 = v49
    v52 := GenS7{F0: 2147483647, F1: 0}
    var v53 float64 = 3.14159
    v1 -= 2147483648
    v3 = (^(v3))
    if true {
        if (false || (v1 > 2147483647)) {
            v54 := GenS3{F0: 1, F1: 1}
            _ = v54
        }
    }
    for fi55 := 0; fi55 < 10; fi55++ {
        v53 -= (v53 + (1.5 - (float64(1.5) - 1e10)))
        var v56 bool = false
        _ = fi55
        _ = v56
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v49
    _ = v5
    _ = v52
    _ = v53
}
