package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 uint8; F2 string }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 uint8; F1 bool; F2 int }
type GenS5 struct { F0 int32; F1 int32; F2 uint16 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint32; F1 bool; F2 int64 }

func main() {
    v1 := []string{}
    var v2 uint64 = 0
    var v3 int8 = 0
    var v4 int32 = 1
    var v5 uint = 0
    v5 = v5
    var v6 uint64 = 4294967295
    var v7 float32 = -2.25
    v8 := GenS1{F0: 0}
    for fi9 := 0; fi9 < 17; fi9++ {
        for fi10 := 0; fi10 < 3; fi10++ {
            _ = len(v1)
            _ = fi10
        }
        _ = fi9
    }
    if len(v1) > 0 {
        v1[0] = ""
    }
    fmt.Println(v5, v7)
    v3 = v3
    v11 := GenS1{F0: 0}
    for fi12 := 0; fi12 < 20; fi12++ {
        var v13 int16 = 42
        v14 := GenS2{F0: (v2 ^ (^(v2))), F1: 1, F2: "foo bar"}
        v1 = append(v1, "")
        _ = fi12
        _ = v13
        _ = v14
    }
    v15 := GenS3{F0: 42, F1: (v2 % ((v2) | 1))}
    fmt.Println(v4, v5)
    for fi16 := 0; fi16 < 2; fi16++ {
        if true {
            v5 -= (^(v5))
            if (false && (false && (true || true))) {
                switch "foo bar" {
                case "foo bar":
                    if len(v1) > 0 {
                        v1[0] = ""
                    }
                default:
                    for fi17 := 0; fi17 < 19; fi17++ {
                        var v18 bool = (!false)
                        if (v18 || (fi16 != fi17)) {
                            var v19 int64 = 582
                            if ((826 < 65535) || (!(v18 && (!v18)))) {
                                if false {
                                    v20 := GenS2{F0: 278, F1: 128, F2: "\n\t"}
                                    var v21 float64 = (1.5 * -2.25)
                                    if (v18 && v18) {
                                        var v22 bool = true
                                        switch (fi16 - -292) {
                                        case 9223372036854775807:
                                            switch 32768 {
                                            case 32767:
                                                var v23 int32 = (v4 | v4)
                                                _ = v23
                                            }
                                        }
                                        v2 = v2
                                        switch fi16 {
                                        case 1:
                                            v24 := GenS4{F0: 0, F1: true, F2: fi16}
                                            switch "\n\t" {
                                            case "\n\t":
                                                v4 = (v4 * (v4 << ((v5) & 63)))
                                                fmt.Println(v4, v2, v2)
                                            case "a":
                                                _ = len(v1)
                                            case "":
                                                var v25 int32 = v4
                                                _ = v25
                                            default:
                                                v26 := GenS5{F0: (v4 - (v4 | (-(v4)))), F1: int32(v19), F2: 896}
                                                _ = v26
                                            }
                                            var v27 float32 = v7
                                            v18 = v18
                                            _ = v24
                                            _ = v27
                                        case -2147483648:
                                            var v28 bool = true
                                            _ = v28
                                        case 42:
                                            v29 := GenS2{F0: (v2 - (^(v6))), F1: 42, F2: "hello"}
                                            _ = v29
                                        }
                                        _ = v22
                                    } else {
                                        var v30 string = "foo bar"
                                        v31 := GenS2{F0: (^(v6)), F1: 255, F2: (v30 + v30)}
                                        var v32 float64 = (((-2.25 + v21) + ((v21 - (v21 - (1.5 * v21))) - (0.0 - v21))) - 1.5)
                                        _ = v30
                                        _ = v31
                                        _ = v32
                                    }
                                    v33 := GenS2{F0: 185, F1: uint8(v2), F2: "a"}
                                    var v34 int16 = -32768
                                    _ = v20
                                    _ = v21
                                    _ = v33
                                    _ = v34
                                }
                            } else {
                                v2 = 0
                                v18 = v18
                            }
                            _ = v19
                        } else {
                            var v35 int64 = -9223372036854775808
                            v36 := GenS6{F0: 42}
                            _ = v35
                            _ = v36
                        }
                        switch (^(fi16)) {
                        case 32767:
                            fmt.Println(v3, v5, v18)
                        default:
                            if (((v6 + v2) < v6) && v18) {
                                if (int64(fi17) >= 315) {
                                    v37 := GenS4{F0: 255, F1: (true || (-1 > v3)), F2: (fi17 - fi16)}
                                    _ = v37
                                }
                            }
                        }
                        _ = fi17
                        _ = v18
                    }
                }
                v2 = v6
                v7 -= (((v7 / v7) + (0.1 - (v7 / (0.0 / (v7 + (((float32(1e10) + 0.0) * (-2.25 * v7)) - (float32(0.1) + (v7 * -2.25)))))))) * 0.0)
            } else {
                v38 := GenS7{F0: 937, F1: (!(0 == 65535)), F2: 0}
                v5 = v5
                v5 -= 0
                _ = v38
            }
            v39 := GenS2{F0: 0, F1: 25, F2: ("\n\t" + "hello")}
            _ = v39
        }
        var v40 bool = (!(!true))
        var v41 uint32 = 65535
        _ = fi16
        _ = v40
        _ = v41
    }
    if (!true) {
        var v42 uint8 = 127
        _ = v42
    } else {
        var v43 float64 = ((0.0 - float64(1.5)) * float64((1.5 * (v7 * (v7 + (float32(-2.25) + v7))))))
        var v44 string = (("\n\t" + "\n\t") + "hello")
        var v45 int16 = 1
        _ = v43
        _ = v44
        _ = v45
    }
    if false {
        v7 -= float32(((3.14159 * (-2.25 + (0.1 + 3.14159))) * 0.1))
    }
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
