package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 int64; F1 string }
type GenS4 struct { F0 int32; F1 uint32; F2 int64 }
func genFn0(p0 string, p1 uint16) uint {
    if true {
        if false {
            p1 = (p1 * 32767)
            fmt.Println(p0, p0)
            v4 := GenS1{F0: 2}
            _ = v4
        } else {
            var v5 int64 = -1
            v5 = 1
            if false {
                var v6 uint8 = 54
                var v7 uint64 = 18446744073709551615
                fmt.Println(v5, p0)
                _ = v6
                _ = v7
            }
            _ = v5
        }
    } else {
        var v8 int8 = -61
        _ = v8
    }
    if true {
        v9 := []int{4294967295, 1, -544, -9223372036854775808}
        switch 0 {
        case 42:
            if true {
                p1 = p1
            }
        case -226:
            if true {
                var v10 int32 = 0
                _ = v10
            }
        }
        var v11 uint8 = 2
        _ = v11
        _ = v9
    } else {
        switch ((((p0 + ("hello" + "")) + (p0 + p0)) + p0) + (p0 + "hello")) {
        case "foo bar":
            v12 := GenS2{F0: (float64(float32(0.0)) >= float64(1.5)), F1: 0}
            _ = v12
        case "":
            switch 0 {
            case 128:
                var v13 int32 = 0
                _ = v13
            case 9223372036854775807:
                if (false || false) {
                    for fi14 := 0; fi14 < 14; fi14++ {
                        if false {
                            v15 := []int{(^(fi14)), fi14}
                            if false {
                                switch "\n\t" {
                                case "\n\t":
                                    if len(v15) > 0 {
                                        v15[0] = (fi14 | (fi14 / ((128) | 1)))
                                    }
                                case "a":
                                    p0 = (p0 + (((p0 + (p0 + p0)) + p0) + p0))
                                }
                                var v16 int8 = 42
                                _ = v16
                            }
                            _ = v15
                        }
                        p1 += 65535
                        _ = fi14
                    }
                    var v17 uint = 65535
                    _ = v17
                }
            case 42:
                if true {
                    if (4294967295 != 107) {
                        var v18 int8 = -40
                        _ = v18
                    }
                    fmt.Println(p1)
                } else {
                    var v19 uint64 = 42
                    p0 = "foo bar"
                    _ = v19
                }
            }
        }
    }
    _ = p0
    _ = p1
    return 65535
}
func genFn1() int32 {
    var v20 uint32 = 2147483647
    var v21 bool = true
    switch 389 {
    case 9223372036854775807:
        v22 := GenS3{F0: 640, F1: "a"}
        _ = v22
    case -2147483648:
        var v23 int16 = -32768
        _ = v23
    case -564:
        v20 = v20
    default:
        v24 := GenS1{F0: 1}
        _ = v24
    }
    if true {
        if v21 {
            v25 := GenS1{F0: 1}
            if (306 <= -47) {
                fmt.Println(v21, v20)
                var v26 string = "a"
                v20 += (^(v20))
                _ = v26
            } else {
                v21 = v21
                switch -2147483648 {
                case 2147483647:
                    if ((1e10 / 1.5) >= -2.25) {
                        switch 32767 {
                        case 1:
                            var v27 bool = (-1 > 579)
                            _ = v27
                        case -2147483648:
                            v28 := GenS3{F0: -1, F1: "a"}
                            _ = v28
                        default:
                            v29 := GenS3{F0: 42, F1: "foo bar"}
                            _ = v29
                        }
                        var v30 float32 = float32(1.5)
                        _ = v30
                    }
                case 42:
                    if v21 {
                        fmt.Println(v20, v21)
                    } else {
                        var v31 uint8 = 255
                        _ = v31
                    }
                case -362:
                    switch 42 {
                    case -9223372036854775808:
                        for fi32 := 0; fi32 < 4; fi32++ {
                            for fi33 := 0; fi33 < 3; fi33++ {
                                v20 -= (v20 & (^(v20)))
                                fmt.Println(v20)
                                _ = fi33
                            }
                            if (!v21) {
                                v20 = (v20 + (v20 >> ((2) & 63)))
                                v34 := GenS1{F0: 744}
                                v20 += v20
                                _ = v34
                            }
                            _ = fi32
                        }
                    }
                default:
                    v35 := []int{42, -521, 128}
                    _ = v35
                }
                switch ("hello" + "a") {
                case "a":
                    v36 := GenS1{F0: 42}
                    _ = v36
                case "":
                    for fi37 := 0; fi37 < 10; fi37++ {
                        switch (fi37 + fi37) {
                        case 647:
                            var v38 float32 = 0.1
                            _ = v38
                        default:
                            var v39 int16 = 1
                            _ = v39
                        }
                        v40 := GenS4{F0: -723, F1: (v20 * v20), F2: 444}
                        _ = fi37
                        _ = v40
                    }
                case "hello":
                    var v41 int16 = 193
                    _ = v41
                default:
                    v42 := []int{-1, -664}
                    _ = v42
                }
            }
            if true {
                v20 -= 802
            } else {
                var v43 bool = (v21 && false)
                if false {
                    var v44 int32 = 2147483647
                    _ = v44
                }
                _ = v43
            }
            _ = v25
        }
    } else {
        fmt.Println(v21)
        fmt.Println(v21, v20)
    }
    _ = v20
    _ = v21
    return 2147483647
}

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[string]int)
    var v3 int8 = 0
    for fi45 := 0; fi45 < 3; fi45++ {
        v1 += v1
        for fi46 := 0; fi46 < 14; fi46++ {
            var v47 float32 = v1
            _ = fi46
            _ = v47
        }
        var v48 uint64 = 0
        _ = fi45
        _ = v48
    }
    v49 := GenS2{F0: ((v1 / v1) >= float32(float64(3.14159))), F1: -416}
    v1 = ((v1 * v1) - (float32(0.0) / (3.14159 + 1.5)))
    for fi50 := 0; fi50 < 12; fi50++ {
        v51 := []int{(fi50 + (fi50 & (fi50 ^ fi50))), fi50}
        switch ("hello" + (("foo bar" + "a") + "a")) {
        case "foo bar":
            switch "" {
            case "\n\t":
                v3 += -1
            default:
                switch -298 {
                case -22:
                    var v52 int16 = 0
                    _ = v52
                case 42:
                    v1 += v1
                case 1:
                    v3 = (v3 * (v3 >> ((genFn0("hello", 42)) & 63)))
                default:
                    if true {
                        if false {
                            v2[("" + ("hello" + "a"))] = 127
                            v1 -= 3.14159
                        }
                    }
                }
            }
        case "hello":
            var v53 uint = 2
            _ = v53
        }
        _ = fi50
        _ = v51
    }
    for fi54 := 0; fi54 < 8; fi54++ {
        switch -2147483648 {
        case 646:
            for fi55 := 0; fi55 < 18; fi55++ {
                if true {
                    v1 += 3.14159
                } else {
                    var v56 int32 = 2147483647
                    _ = v56
                }
                _ = fi55
            }
        case 4294967295:
            v3 -= 100
        case 2147483647:
            v3 -= v3
        default:
            v1 = float32(-2.25)
        }
        _ = fi54
    }
    v57 := GenS1{F0: 0}
    v3 = 127
    for fi58 := 0; fi58 < 9; fi58++ {
        var v59 int = 128
        _ = fi58
        _ = v59
    }
    fmt.Println(v3)
    var v60 uint32 = 431
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v49
    _ = v57
    _ = v60
}
