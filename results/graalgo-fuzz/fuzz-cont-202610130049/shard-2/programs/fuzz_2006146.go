package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 int8; F1 uint8; F2 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16; F1 int32; F2 int32 }
type GenS5 struct { F0 uint8; F1 string; F2 string }

func main() {
    var v1 int32 = -957
    v2 := []string{("hello" + "\n\t"), "hello"}
    var v3 bool = ((!(1.5 <= 3.14159)) && (true && false))
    if v3 {
        for fi4 := 0; fi4 < 4; fi4++ {
            v1 -= 42
            var v5 string = "hello"
            var v6 int16 = 1
            _ = fi4
            _ = v5
            _ = v6
        }
        if v3 {
            fmt.Println(v3)
            v7 := GenS1{F0: -128, F1: 916}
            _ = v7
        }
        if len(v2) > 0 {
            v2[0] = "\n\t"
        }
    } else {
        var v8 int = -991
        if (!false) {
            v3 = v3
            if (4294967295 < 1) {
                v9 := GenS1{F0: 108, F1: 590}
                _ = v9
            }
            for fi10 := 0; fi10 < 3; fi10++ {
                var v11 float32 = 0.0
                if v3 {
                    v8 = 932
                    v1 += -981
                } else {
                    v12 := GenS2{F0: 42, F1: 128, F2: -584}
                    v8 += (v8 / (((fi10 ^ (v8 << ((2) & 63)))) | 1))
                    _ = v12
                }
                v1 += (v1 + (v1 % ((v1) | 1)))
                _ = fi10
                _ = v11
            }
        }
        _ = v8
    }
    _ = len(v2)
    if (!true) {
        fmt.Println(v1)
        var v13 float32 = -2.25
        switch 2147483647 {
        case 42:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case -907:
            var v14 float64 = float64((((((v13 * (v13 * ((3.14159 - 1e10) - v13))) * ((v13 * 0.0) - 1.5)) * v13) * (v13 - -2.25)) + (v13 / 1e10)))
            _ = v14
        case 127:
            var v15 float64 = 1e10
            _ = v15
        default:
            fmt.Println(v3, v1, v3)
        }
        _ = v13
    }
    for fi16 := 0; fi16 < 17; fi16++ {
        v1 = 992
        _ = fi16
    }
    var v17 float64 = float64(0.0)
    if len(v2) > 0 {
        v2[0] = ("hello" + "\n\t")
    }
    fmt.Println(v17, v1, v17)
    v18 := GenS3{F0: (false || v3)}
    fmt.Println(v17)
    var v19 uint8 = 1
    v20 := GenS3{F0: v3}
    v3 = false
    if ("hello" > "") {
        if v3 {
            var v21 int16 = 1
            _ = v21
        } else {
            for fi22 := 0; fi22 < 2; fi22++ {
                var v23 bool = (v3 || (2 < 2147483647))
                _ = fi22
                _ = v23
            }
            for fi24 := 0; fi24 < 3; fi24++ {
                switch "" {
                case "a":
                    v25 := GenS4{F0: 42, F1: 830, F2: (v1 | v1)}
                    _ = v25
                case "":
                    if ((v17 - 1.5) == -2.25) {
                        switch fi24 {
                        case 0:
                            switch -238 {
                            case 2147483647:
                                var v26 int64 = 0
                                _ = v26
                            case -9223372036854775808:
                                v1 += (v1 | v1)
                            case -1:
                                for fi27 := 0; fi27 < 4; fi27++ {
                                    v28 := GenS5{F0: v19, F1: "foo bar", F2: ("" + "hello")}
                                    var v29 float32 = -2.25
                                    _ = fi27
                                    _ = v28
                                    _ = v29
                                }
                            default:
                                var v30 int32 = v1
                                _ = v30
                            }
                        }
                    }
                default:
                    for fi31 := 0; fi31 < 11; fi31++ {
                        v1 -= v1
                        if true {
                            v17 = v17
                            v32 := GenS3{F0: (!(!false))}
                            _ = v32
                        }
                        v33 := GenS3{F0: (uint(v19) > 62)}
                        _ = fi31
                        _ = v33
                    }
                }
                var v34 float32 = 0.1
                _ = fi24
                _ = v34
            }
        }
    }
    if v3 {
        var v35 bool = (!(!(!(32767 == 32767))))
        v19 += (v19 + v19)
        _ = v35
    } else {
        v19 = v19
        var v36 string = ((("hello" + ("foo bar" + "hello")) + "hello") + "")
        var v37 uint8 = 1
        _ = v36
        _ = v37
    }
    fmt.Println(v17)
    v38 := GenS4{F0: 2, F1: (v1 % (((^(v1))) | 1)), F2: v1}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v38
}
