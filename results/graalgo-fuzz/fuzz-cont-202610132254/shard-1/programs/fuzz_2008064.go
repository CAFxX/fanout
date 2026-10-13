package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 int16 }
type GenS2 struct { F0 int64; F1 uint64; F2 int32 }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 uint; F1 uint; F2 uint32 }
type GenS5 struct { F0 uint16; F1 int }
type GenS6 struct { F0 bool; F1 bool; F2 uint8 }
func genFn0() uint16 {
    switch "foo bar" {
    case "\n\t":
        v5 := []int{32767, 32768, -2147483648}
        _ = v5
    }
    var v6 int = 127
    v7 := GenS1{F0: -32768, F1: 1, F2: -534}
    var v8 int16 = 863
    _ = v6
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 int16 = 32767
    var v2 int32 = -205
    var v3 uint16 = 346
    var v4 float32 = (1.5 / (0.0 - 0.1))
    var v9 string = (("a" + "a") + ("" + ((("hello" + "\n\t") + "hello") + "hello")))
    v10 := []int{128, 32767, 134, 256}
    if (v1 == v1) {
        v4 += float32(0.0)
        switch 0 {
        case 32768:
            if (155 == 859) {
                var v11 int16 = v1
                _ = len(v10)
                _ = v11
            }
        case 32767:
            for fi12 := 0; fi12 < 1; fi12++ {
                var v13 bool = false
                v4 = v4
                v14 := GenS1{F0: (v1 * 0), F1: v3, F2: (-(v1))}
                _ = fi12
                _ = v13
                _ = v14
            }
        case 4294967295:
            var v15 uint64 = 1
            _ = v15
        default:
            for fi16 := 0; fi16 < 12; fi16++ {
                for fi17 := 0; fi17 < 15; fi17++ {
                    v1 += (v1 / ((v1) | 1))
                    var v18 uint64 = 577
                    fmt.Println(v1, v1)
                    if ((v18 - v18) != (^(v18))) {
                        v1 -= -1
                        var v19 float64 = 3.14159
                        if false {
                            if (fi16 >= 32767) {
                                switch v9 {
                                case "a":
                                    fmt.Println(v19)
                                }
                                v20 := GenS2{F0: -1, F1: (v18 / ((830) | 1)), F2: v2}
                                _ = v20
                            }
                            var v21 float64 = (v19 + v19)
                            _ = v21
                        }
                        _ = v19
                    } else {
                        v9 = ("a" + v9)
                        if (!false) {
                            fmt.Println(v2, v2)
                        }
                        v9 = ("hello" + (v9 + v9))
                        if true {
                            if false {
                                if len(v10) > 0 {
                                    v10[0] = fi16
                                }
                            } else {
                                v4 -= ((0.0 - v4) + v4)
                                var v22 int64 = 1
                                var v23 int16 = v1
                                _ = v22
                                _ = v23
                            }
                        }
                        v9 = ""
                        v4 -= 0.0
                        _ = len(v10)
                    }
                    _ = fi17
                    _ = v18
                }
                _ = fi16
            }
        }
    } else {
        if ((v2 | v2) != v2) {
            var v24 float64 = 3.14159
            v25 := GenS1{F0: (v1 & (v1 & v1)), F1: v3, F2: -1}
            var v26 bool = (!true)
            _ = v24
            _ = v25
            _ = v26
        } else {
            if (v2 <= v2) {
                var v27 uint32 = 1
                v3 -= (v3 << ((38) & 63))
                v4 += ((1e10 / 1.0) - (v4 + float32(-2.25)))
                _ = v27
            }
            v2 -= (v2 << ((42) & 63))
            for fi28 := 0; fi28 < 5; fi28++ {
                var v29 uint = 2
                _ = fi28
                _ = v29
            }
        }
        var v30 uint64 = 0
        _ = v30
    }
    v1 += v1
    v31 := GenS1{F0: (v1 & (v1 / (((-(v1))) | 1))), F1: (^(v3)), F2: (v1 ^ 66)}
    var v32 uint8 = 128
    _ = len(v10)
    _ = len(v10)
    fmt.Println(v3, v9)
    v32 -= (^(v32))
    if true {
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
    } else {
        _ = len(v10)
        if true {
            var v33 uint16 = v3
            switch (((v9 + "a") + v9) + "hello") {
            case "foo bar":
                _ = len(v10)
            case "hello":
                var v34 int = 42
                _ = v34
            case "\n\t":
                var v35 int32 = (v2 & v2)
                _ = v35
            }
            var v36 int32 = (v2 | v2)
            _ = v33
            _ = v36
        }
    }
    if (!((((1.5 * 0.0) - 1.5) + 3.14159) < (3.14159 - (1.5 + 1e10)))) {
        v37 := GenS3{F0: (v1 * -1), F1: true}
        fmt.Println(v1, v4, v1)
        _ = v37
    } else {
        v38 := GenS4{F0: 4294967295, F1: 1, F2: 42}
        v1 -= 728
        _ = v38
    }
    if true {
        for fi39 := 0; fi39 < 19; fi39++ {
            switch v9 {
            case "a":
                for fi40 := 0; fi40 < 15; fi40++ {
                    v1 -= (-(v1))
                    _ = fi40
                }
            default:
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
            }
            var v41 int16 = (^(v1))
            if false {
                v41 = (^(v41))
                if len(v10) > 0 {
                    v10[0] = (fi39 / ((fi39) | 1))
                }
                var v42 uint = 882
                _ = v42
            } else {
                for fi43 := 0; fi43 < 19; fi43++ {
                    var v44 uint8 = v32
                    if (!false) {
                        v45 := GenS5{F0: uint16(v44), F1: (fi39 << ((65535) & 63))}
                        _ = v45
                    }
                    v4 += ((((0.0 / v4) / (v4 / v4)) / 1.0) + (v4 + float32(float64(((0.1 / ((3.14159 + ((v4 / ((v4 * ((0.1 / 3.14159) + float32((1.5 + ((-2.25 + -2.25) / 3.14159))))) * -2.25)) + 0.0)) * (0.0 + 0.0))) * 0.0)))))
                    _ = fi43
                    _ = v44
                }
                switch v9 {
                case "foo bar":
                    var v46 int8 = -128
                    _ = v46
                case "hello":
                    var v47 uint64 = 42
                    _ = v47
                default:
                    if (v32 == (v32 + (v32 - v32))) {
                        var v48 int32 = v2
                        switch fi39 {
                        case -393:
                            _ = len(v10)
                        case 1:
                            if true {
                                var v49 int8 = 127
                                if ((v3 >> ((42) & 63)) <= (v3 | (^(v3)))) {
                                    switch (fi39 + fi39) {
                                    case 723:
                                        v4 -= 1e10
                                    }
                                }
                                _ = v49
                            }
                        }
                        var v50 uint32 = 665
                        _ = v48
                        _ = v50
                    } else {
                        for fi51 := 0; fi51 < 9; fi51++ {
                            var v52 uint8 = 1
                            _ = fi51
                            _ = v52
                        }
                    }
                }
            }
            _ = fi39
            _ = v41
        }
        switch ("hello" + "\n\t") {
        case "a":
            if (v4 <= v4) {
                if false {
                    var v53 uint8 = 255
                    _ = v53
                } else {
                    v2 = v2
                    v54 := GenS6{F0: (3.14159 != float32((((1.5 * -2.25) / 0.1) - ((-2.25 / (0.1 / 1e10)) + 1e10)))), F1: (1.5 == float32(0.0)), F2: 255}
                    for fi55 := 0; fi55 < 18; fi55++ {
                        switch "" {
                        case "\n\t":
                            for fi56 := 0; fi56 < 7; fi56++ {
                                if (!(1 > 1)) {
                                    v3 -= (v3 ^ v3)
                                } else {
                                    var v57 uint32 = 65535
                                    v3 += 32767
                                    switch 9223372036854775807 {
                                    case 4294967295:
                                        v32 = v32
                                    case 128:
                                        var v58 float64 = (0.1 / (3.14159 * (1.5 + (1e10 + 1.5))))
                                        _ = v58
                                    default:
                                        fmt.Println(v4, v4, v3)
                                    }
                                    _ = v57
                                }
                                v9 = v9
                                fmt.Println(fi56, fi55)
                                _ = fi56
                            }
                        case "":
                            if len(v10) > 0 {
                                v10[0] = (fi55 / (((fi55 | (fi55 ^ 42))) | 1))
                            }
                        case "a":
                            v59 := GenS5{F0: v3, F1: fi55}
                            _ = v59
                        }
                        switch "foo bar" {
                        case "foo bar":
                            fmt.Println(v3, v32, v32)
                        case "a":
                            v2 -= v2
                        case "":
                            v4 = v4
                        }
                        _ = fi55
                    }
                    _ = v54
                }
                v9 = ("\n\t" + "hello")
            } else {
                if len(v10) > 0 {
                    v10[0] = 0
                }
                _ = len(v10)
                var v60 float64 = (float64(v4) - (((0.0 - (0.1 * (0.0 * float64((float32(0.1) - (v4 * v4)))))) * 3.14159) + 0.1))
                _ = v60
            }
        }
        var v61 uint64 = 624
        _ = v61
    } else {
        fmt.Println(v1, v32, v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v9
}
