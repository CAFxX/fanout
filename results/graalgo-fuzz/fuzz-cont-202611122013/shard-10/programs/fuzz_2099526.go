package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 int64; F2 uint64 }
type GenS3 struct { F0 int64; F1 int; F2 int16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint32; F1 uint }
func genFn0() uint64 {
    v4 := GenS1{F0: "foo bar"}
    var v5 uint16 = 42
    var v6 bool = (!((307 > 127) || false))
    var v7 float32 = 1.5
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 1
}
func genFn1() int64 {
    var v8 int64 = -3
    v8 -= v8
    var v9 int16 = 197
    v9 -= 1
    v9 -= (v9 & v9)
    _ = v8
    _ = v9
    return 9223372036854775807
}

func main() {
    var v1 string = ("foo bar" + "a")
    var v2 int32 = 1
    v3 := make(map[string]uint32)
    v2 -= -244
    var v10 int64 = 1
    v11 := []int{-1}
    v2 = v2
    v12 := GenS1{F0: "hello"}
    for fi13 := 0; fi13 < 13; fi13++ {
        for fi14 := 0; fi14 < 7; fi14++ {
            if len(v11) > 0 {
                v11[0] = (-(fi14))
            }
            _ = fi14
        }
        _ = fi13
    }
    var v15 bool = ((!true) || true)
    for fi16 := 0; fi16 < 8; fi16++ {
        v11 = append(v11, 128)
        var v17 int32 = 42
        _ = fi16
        _ = v17
    }
    if v15 {
        var v18 string = ("" + (v1 + v1))
        var v19 uint16 = 87
        _ = len(v3)
        _ = v18
        _ = v19
    } else {
        v2 -= (v2 / (((v2 % (((v2 ^ (v2 | (^(v2))))) | 1))) | 1))
        v20 := GenS2{F0: 65535, F1: v10, F2: 1}
        _ = v20
    }
    for fi21 := 0; fi21 < 12; fi21++ {
        var v22 int8 = -105
        if false {
            if (!false) {
                if v15 {
                    fmt.Println(v22, v1)
                } else {
                    v2 += v2
                    var v23 uint16 = 597
                    _ = v23
                }
                v24 := GenS2{F0: 729, F1: genFn1(), F2: 336}
                if v15 {
                    v25 := GenS3{F0: -1, F1: fi21, F2: -1}
                    v22 = v22
                    _ = v25
                } else {
                    for fi26 := 0; fi26 < 2; fi26++ {
                        switch v1 {
                        case "foo bar":
                            if false {
                                v22 = v22
                                var v27 float64 = 0.0
                                _ = v27
                            } else {
                                switch (((("\n\t" + (v1 + v1)) + "") + v1) + "hello") {
                                case "hello":
                                    var v28 int16 = -747
                                    _ = v28
                                case "":
                                    v29 := GenS4{F0: (v2 + v2)}
                                    _ = v29
                                case "a":
                                    v15 = false
                                default:
                                    v30 := GenS3{F0: (v10 >> ((18446744073709551615) & 63)), F1: 32768, F2: -960}
                                    _ = v30
                                }
                                var v31 float32 = float32(-2.25)
                                var v32 int8 = -113
                                _ = v31
                                _ = v32
                            }
                        case "a":
                            fmt.Println(v1, v15, fi21)
                        }
                        if (fi26 > (^(fi21))) {
                            switch (v1 + (("hello" + "a") + ("a" + v1))) {
                            case "\n\t":
                                switch fi21 {
                                case -1:
                                    v11 = append(v11, 1)
                                case 32768:
                                    v22 = (-(v22))
                                }
                            case "hello":
                                delete(v3, v1)
                            default:
                                if (141 != -1) {
                                    v3[((v1 + v1) + ("a" + "a"))] = 2
                                }
                            }
                        } else {
                            delete(v3, v1)
                            if false {
                                v11 = append(v11, -2147483648)
                                v33 := GenS5{F0: 914, F1: 182}
                                v15 = true
                                _ = v33
                            }
                        }
                        var v34 uint16 = 218
                        _ = fi26
                        _ = v34
                    }
                    var v35 float64 = (((-2.25 / 3.14159) / (1.5 + (0.1 - 1e10))) - -2.25)
                    for fi36 := 0; fi36 < 2; fi36++ {
                        v2 += v2
                        var v37 uint8 = 1
                        var v38 uint = 1
                        _ = fi36
                        _ = v37
                        _ = v38
                    }
                    _ = v35
                }
                _ = v24
            }
            var v39 float64 = 1e10
            v15 = v15
            _ = v39
        }
        v10 = int64(v2)
        _ = fi21
        _ = v22
    }
    var v40 float32 = 3.14159
    if (!false) {
        if false {
            v2 -= -836
            for fi41 := 0; fi41 < 15; fi41++ {
                v1 = ""
                v10 += (-(v10))
                _ = fi41
            }
        } else {
            v1 = v1
            var v42 float64 = 0.0
            _ = v42
        }
        for fi43 := 0; fi43 < 5; fi43++ {
            switch v1 {
            case "foo bar":
                if v15 {
                    var v44 int8 = 0
                    var v45 int8 = v44
                    _ = v44
                    _ = v45
                }
            case "a":
                _ = len(v3)
            case "":
                if (v15 && (!true)) {
                    var v46 uint = 42
                    _ = v46
                }
            default:
                var v47 uint8 = 83
                _ = v47
            }
            _ = len(v11)
            _ = fi43
        }
    } else {
        var v48 uint8 = 255
        _ = v48
    }
    v15 = v15
    v40 = v40
    for fi49 := 0; fi49 < 5; fi49++ {
        for fi50 := 0; fi50 < 9; fi50++ {
            var v51 uint32 = 959
            var v52 int16 = 1
            v10 = (-(v10))
            _ = fi50
            _ = v51
            _ = v52
        }
        if len(v11) > 0 {
            v11[0] = (fi49 ^ (fi49 | 577))
        }
        _ = fi49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v40:", v40)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v40
}
