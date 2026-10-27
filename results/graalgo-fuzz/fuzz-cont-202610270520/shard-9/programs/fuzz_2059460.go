package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 uint16; F1 int8; F2 int32 }
type GenS3 struct { F0 int8; F1 int8 }
type GenS4 struct { F0 bool; F1 int16 }
func genFn0(p0 uint64, p1 int32) int {
    var v7 float32 = (3.14159 + ((1.5 + (float32(float64((1e10 * 1.5))) * (0.1 + 3.14159))) / (3.14159 / (float32(1e10) + 0.1))))
    v8 := GenS1{F0: "hello", F1: 1}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 bool = (0.1 == -2.25)
    v2 := make(map[string]bool)
    var v3 int32 = 1
    v4 := make(map[int]string)
    var v5 bool = v1
    var v6 float32 = (1e10 / (-2.25 + (-2.25 + (float32(((float64((-2.25 / 1.0)) - (3.14159 * -2.25)) / 1e10)) + (0.0 - 3.14159)))))
    v5 = (!(v3 > 42))
    v9 := GenS1{F0: "foo bar", F1: 463}
    for fi10 := 0; fi10 < 8; fi10++ {
        var v11 string = ""
        if v1 {
            var v12 float32 = (float32(float64(1e10)) + (float32(3.14159) + v6))
            _ = v12
        } else {
            v5 = ((1 != 973) && false)
            v13 := []int{}
            v2[("\n\t" + "\n\t")] = false
            _ = v13
        }
        v5 = v1
        _ = fi10
        _ = v11
    }
    switch 1 {
    case 42:
        if v5 {
            var v14 float64 = 1.5
            var v15 uint = 4294967295
            _ = v14
            _ = v15
        }
    default:
        if v1 {
            for fi16 := 0; fi16 < 1; fi16++ {
                fmt.Println(v5, v1)
                _ = fi16
            }
            if (!true) {
                v6 = v6
            } else {
                var v17 uint32 = 151
                _ = v17
            }
        } else {
            if ("\n\t" == "a") {
                switch "a" {
                case "foo bar":
                    v1 = true
                }
                switch "hello" {
                case "foo bar":
                    fmt.Println(v3)
                case "":
                    switch ((("hello" + "") + "foo bar") + "a") {
                    case "hello":
                        _ = len(v2)
                    case "a":
                        v18 := []int{}
                        _ = v18
                    }
                }
                if true {
                    fmt.Println(v6, v5, v6)
                    for fi19 := 0; fi19 < 10; fi19++ {
                        if (206 < 128) {
                            v5 = (862 == 0)
                            var v20 uint8 = 141
                            var v21 float64 = 1e10
                            _ = v20
                            _ = v21
                        }
                        if (-20 >= -128) {
                            v22 := GenS1{F0: ("foo bar" + ("\n\t" + ("a" + "\n\t"))), F1: -30}
                            var v23 float64 = 3.14159
                            _ = v22
                            _ = v23
                        }
                        _ = fi19
                    }
                }
            }
        }
    }
    v24 := GenS1{F0: "hello", F1: 496}
    var v25 uint64 = 1
    var v26 uint64 = v25
    v27 := []int{-674, -975, 736, genFn0((v25 + (v25 ^ 42)), (v3 << ((2) & 63)))}
    fmt.Println(v26, v5, v3)
    if v1 {
        var v28 bool = v1
        var v29 uint64 = (v26 - (v26 + v25))
        _ = v28
        _ = v29
    } else {
        v30 := v2[("hello" + ((("" + "") + ("\n\t" + "")) + "hello"))]
        switch 127 {
        case -1:
            v31 := v2["\n\t"]
            _ = v31
        case 128:
            for fi32 := 0; fi32 < 7; fi32++ {
                v26 = v26
                var v33 int16 = -32768
                _ = fi32
                _ = v33
            }
        case 852:
            var v34 uint64 = v25
            _ = v34
        default:
            fmt.Println(v25, v3)
        }
        switch "" {
        case "foo bar":
            v2[("\n\t" + (((("\n\t" + (("foo bar" + "") + (("hello" + "") + "foo bar"))) + "\n\t") + "foo bar") + ("\n\t" + "hello")))] = (290 > 2)
        case "hello":
            for fi35 := 0; fi35 < 10; fi35++ {
                v2[(("foo bar" + "") + "a")] = v1
                _ = fi35
            }
        case "":
            v6 += ((v6 + -2.25) * (v6 + 3.14159))
        default:
            v26 -= v25
        }
        _ = v30
    }
    v26 += v26
    switch 128 {
    case -9223372036854775808:
        v36 := GenS2{F0: 65535, F1: 100, F2: v3}
        _ = v36
    default:
        if len(v27) > 0 {
            v27[0] = -366
        }
    }
    switch genFn0((v26 + (v25 | v25)), 42) {
    case 4294967295:
        v37 := GenS1{F0: ("" + "foo bar"), F1: 1}
        _ = v37
    default:
        v38 := GenS3{F0: -95, F1: 0}
        _ = v38
    }
    switch 9223372036854775807 {
    case 32767:
        var v39 bool = (633 <= -564)
        _ = v39
    default:
        v5 = (!(((^(v25)) != v26) || (1 < 100)))
    }
    var v40 uint64 = 2
    var v41 bool = v1
    if (genFn0(v40, (v3 % ((-1) | 1))) >= 127) {
        v26 = (v26 - (v40 & v26))
        if (false && false) {
            if (!((v25 * v25) >= v25)) {
                if false {
                    v25 = (v25 << ((80) & 63))
                    fmt.Println(v5)
                    v25 = (v40 - v26)
                } else {
                    var v42 uint32 = 0
                    for fi43 := 0; fi43 < 12; fi43++ {
                        var v44 int = (fi43 & (-(fi43)))
                        if len(v27) >= 2 {
                            v27 = v27[1:len(v27)]
                        }
                        if (!("a" <= ((("\n\t" + "hello") + "foo bar") + ("foo bar" + ("a" + ((("" + "hello") + "") + ("\n\t" + ("hello" + "hello")))))))) {
                            v42 = (v42 | 766)
                            v5 = true
                        } else {
                            _ = len(v2)
                        }
                        _ = fi43
                        _ = v44
                    }
                    _ = v42
                }
                var v45 float32 = (1e10 + v6)
                if (923 != 1) {
                    var v46 uint8 = 127
                    _ = v46
                } else {
                    var v47 string = ("foo bar" + "foo bar")
                    switch v47 {
                    case "a":
                        if (-2147483648 < (v3 >> ((18446744073709551615) & 63))) {
                            v48 := GenS2{F0: 1, F1: -25, F2: 42}
                            _ = v48
                        }
                    case "foo bar":
                        if (!((413 > 813) || (!v41))) {
                            v1 = (false || (!((v5 || v41) || v1)))
                            for fi49 := 0; fi49 < 20; fi49++ {
                                var v50 uint8 = 128
                                var v51 float64 = (float64(v6) * 0.1)
                                _ = fi49
                                _ = v50
                                _ = v51
                            }
                            v52 := GenS4{F0: v5, F1: 0}
                            _ = v52
                        } else {
                            delete(v4, 833)
                            v6 = (v45 * -2.25)
                        }
                    default:
                        v53 := GenS2{F0: 598, F1: 122, F2: v3}
                        _ = v53
                    }
                    _ = len(v27)
                    _ = v47
                }
                _ = v45
            } else {
                fmt.Println(v41, v5)
            }
            for fi54 := 0; fi54 < 19; fi54++ {
                for fi55 := 0; fi55 < 10; fi55++ {
                    v26 += 836
                    _ = fi55
                }
                if len(v27) >= 2 {
                    v27 = v27[1:len(v27)]
                }
                var v56 uint64 = 579
                _ = fi54
                _ = v56
            }
            fmt.Println(v25, v26)
        } else {
            delete(v4, 9223372036854775807)
        }
    } else {
        var v57 uint = 0
        v26 = (v40 | (v26 * (^(v25))))
        var v58 float32 = (3.14159 + (0.0 + v6))
        _ = v57
        _ = v58
    }
    v3 = (v3 % (((v3 / ((v3) | 1))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
    _ = v9
}
