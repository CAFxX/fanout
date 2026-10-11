package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 uint64 }
type GenS2 struct { F0 int8; F1 uint32; F2 int8 }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 int; F1 int }
type GenS5 struct { F0 string; F1 int64; F2 int16 }
func genFn0(p0 uint64) uint16 {
    var v6 string = ("\n\t" + "a")
    for fi7 := 0; fi7 < 6; fi7++ {
        for fi8 := 0; fi8 < 12; fi8++ {
            var v9 int16 = -1
            v10 := GenS1{F0: (fi7 / ((-579) | 1)), F1: (fi8 | 4294967295), F2: (p0 ^ (p0 * (p0 ^ 42)))}
            _ = fi8
            _ = v10
            _ = v9
        }
        var v11 string = ((v6 + (v6 + ((v6 + v6) + ("hello" + "\n\t")))) + "hello")
        _ = fi7
        _ = v11
    }
    _ = p0
    _ = v6
    return 288
}

func main() {
    var v1 int = -9223372036854775808
    var v2 uint = 556
    var v3 int64 = 9223372036854775807
    v4 := []string{"\n\t", "\n\t"}
    var v5 int = (v1 >> ((v2) & 63))
    v3 = (v3 | -390)
    v5 += v1
    for fi12 := 0; fi12 < 19; fi12++ {
        if (!false) {
            for fi13 := 0; fi13 < 16; fi13++ {
                var v14 bool = false
                _ = fi13
                _ = v14
            }
            var v15 uint8 = 144
            _ = v15
        }
        if len(v4) > 0 {
            v4[0] = ("" + "hello")
        }
        _ = fi12
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        if (((("hello" + "foo bar") + (("a" + "foo bar") + ("a" + ("" + "\n\t")))) + (((("\n\t" + "") + "a") + "hello") + ("hello" + "\n\t"))) == "\n\t") {
            var v17 uint32 = 0
            if false {
                v18 := GenS1{F0: 0, F1: fi16, F2: 0}
                for fi19 := 0; fi19 < 11; fi19++ {
                    switch (("\n\t" + "") + "hello") {
                    case "foo bar":
                        v20 := GenS1{F0: fi19, F1: v1, F2: 371}
                        _ = v20
                    default:
                        var v21 float64 = 0.0
                        _ = v21
                    }
                    _ = fi19
                }
                _ = v18
            }
            if true {
                for fi22 := 0; fi22 < 4; fi22++ {
                    if ((true || (true && (261 < (fi16 * -506)))) && true) {
                        switch "foo bar" {
                        case "":
                            var v23 int32 = 703
                            _ = v23
                        }
                        var v24 string = "a"
                        switch fi16 {
                        case 42:
                            v25 := GenS2{F0: -49, F1: (v17 << ((175) & 63)), F2: 1}
                            _ = v25
                        case 0:
                            if true {
                                v24 = v24
                                var v26 int32 = 42
                                var v27 string = "hello"
                                if true {
                                    var v28 int16 = 193
                                    var v29 uint = v2
                                    v5 += v5
                                    _ = v28
                                    _ = v29
                                } else {
                                    fmt.Println(v26)
                                    v30 := GenS2{F0: 42, F1: v17, F2: -128}
                                    _ = v30
                                }
                                var v31 float32 = 3.14159
                                _ = v26
                                _ = v27
                                _ = v31
                            }
                        case 460:
                            switch v5 {
                            case 0:
                                fmt.Println(fi16)
                                switch -263 {
                                case -458:
                                    if (654 <= (v17 * v17)) {
                                        if (!(!(true && (!((((float32(0.0) * ((float32(-2.25) - (0.1 + 1.5)) / (3.14159 + 1.5))) + 0.1) / (float32(float64(float32(1e10))) / 1.5)) < float32(float64((((-2.25 - float32(float64(0.0))) + 3.14159) - ((0.0 - 0.1) - 0.1))))))))) {
                                            v32 := GenS3{F0: 484, F1: 127}
                                            _ = v32
                                        }
                                    } else {
                                        v33 := GenS3{F0: uint64(v17), F1: 1}
                                        _ = v33
                                    }
                                case -913:
                                    if true {
                                        v5 -= (fi22 ^ (fi22 & (^(v5))))
                                        switch v24 {
                                        case "":
                                            v2 += (v2 - 0)
                                        default:
                                            if true {
                                                var v34 float64 = -2.25
                                                v35 := GenS1{F0: (fi16 / (((fi16 | (fi22 << ((v2) & 63)))) | 1)), F1: v1, F2: 2}
                                                _ = v34
                                                _ = v35
                                            }
                                        }
                                    } else {
                                        if true {
                                            _ = len(v4)
                                            _ = len(v4)
                                        }
                                        v5 = fi16
                                        v2 += v2
                                        if (true && false) {
                                            v3 = 1
                                            fmt.Println(v24, v1, v3)
                                        }
                                    }
                                }
                            case 32767:
                                var v36 int = (v5 % (((fi22 * (v5 + (v5 >> ((v2) & 63))))) | 1))
                                _ = v36
                            }
                        }
                        _ = v24
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v17 += (v17 + v17)
                    var v37 uint8 = 185
                    if false {
                        var v38 int16 = -1
                        _ = v38
                    }
                    _ = fi22
                    _ = v37
                }
            }
            _ = v17
        } else {
            switch "foo bar" {
            case "":
                _ = len(v4)
            case "foo bar":
                var v39 int64 = v3
                _ = v39
            case "\n\t":
                v1 = v1
            default:
                v40 := GenS3{F0: 2, F1: 0}
                _ = v40
            }
            fmt.Println(v1, fi16, v2)
            if false {
                var v41 uint64 = 1
                var v42 string = "a"
                if (v3 <= 483) {
                    if true {
                        if (true && true) {
                            fmt.Println(v41)
                            if len(v4) > 0 {
                                v4[0] = ""
                            }
                            v43 := GenS1{F0: (^(fi16)), F1: fi16, F2: 42}
                            _ = v43
                        }
                        v44 := GenS4{F0: fi16, F1: v1}
                        _ = v44
                    }
                }
                _ = v41
                _ = v42
            } else {
                v45 := GenS5{F0: "a", F1: (v3 ^ 9223372036854775807), F2: -1}
                if (42 != uint64(v2)) {
                    for fi46 := 0; fi46 < 11; fi46++ {
                        fmt.Println(fi46, v3)
                        _ = fi46
                    }
                } else {
                    for fi47 := 0; fi47 < 20; fi47++ {
                        v3 -= (v3 | (v3 ^ v3))
                        var v48 uint64 = 42
                        v3 = (^(v3))
                        _ = fi47
                        _ = v48
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                fmt.Println(v3, fi16)
                _ = v45
            }
        }
        switch v1 {
        case 296:
            for fi49 := 0; fi49 < 20; fi49++ {
                v50 := GenS2{F0: 6, F1: 598, F2: -10}
                _ = fi49
                _ = v50
            }
        case 344:
            v51 := GenS2{F0: -1, F1: 42, F2: 74}
            _ = v51
        default:
            var v52 int8 = 0
            _ = v52
        }
        v1 += (v1 ^ v1)
        _ = fi16
    }
    var v53 uint16 = 42
    v5 = v5
    var v54 uint16 = v53
    var v55 bool = (1 == 128)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v55
}
