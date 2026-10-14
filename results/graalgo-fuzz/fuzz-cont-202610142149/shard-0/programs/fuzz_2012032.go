package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 uint16 }
type GenS2 struct { F0 int16; F1 uint16; F2 uint8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64; F1 int }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint }
func genFn0() int32 {
    var v6 float64 = -2.25
    for fi7 := 0; fi7 < 2; fi7++ {
        v6 -= v6
        _ = fi7
    }
    var v8 int8 = 100
    switch 9223372036854775807 {
    case -1:
        for fi9 := 0; fi9 < 8; fi9++ {
            for fi10 := 0; fi10 < 15; fi10++ {
                if true {
                    var v11 float32 = (0.1 * float32(float64((1e10 - 3.14159))))
                    _ = v11
                } else {
                    v8 = (v8 | v8)
                    v12 := GenS1{F0: 1, F1: true, F2: 255}
                    v13 := []int{}
                    _ = v12
                    _ = v13
                }
                v14 := []int{(fi10 / (((-(fi10))) | 1)), 9223372036854775807, fi9}
                _ = fi10
                _ = v14
            }
            _ = fi9
        }
    default:
        v6 -= float64((float32(-2.25) / 1e10))
    }
    if false {
        for fi15 := 0; fi15 < 13; fi15++ {
            v16 := GenS2{F0: -32768, F1: 42, F2: 42}
            v17 := GenS1{F0: 18446744073709551615, F1: (!false), F2: 301}
            _ = fi15
            _ = v16
            _ = v17
        }
    }
    _ = v6
    _ = v8
    return 449
}

func main() {
    v1 := []uint32{1, 2, 585}
    var v2 int16 = -32768
    v3 := make(map[string]uint32)
    var v4 int8 = -6
    v5 := []string{("hello" + "foo bar"), (("a" + "a") + ""), "foo bar"}
    var v18 float32 = 0.0
    v4 += (v4 << ((1) & 63))
    if (255 >= 42) {
        for fi19 := 0; fi19 < 9; fi19++ {
            if true {
                var v20 int64 = 9223372036854775807
                if (!false) {
                    v20 = v20
                    if true {
                        for fi21 := 0; fi21 < 6; fi21++ {
                            var v22 float64 = 1.5
                            var v23 int = fi19
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = fi21
                            _ = v22
                            _ = v23
                        }
                        var v24 uint32 = 245
                        for fi25 := 0; fi25 < 19; fi25++ {
                            v4 -= 126
                            switch ("a" + "\n\t") {
                            case "foo bar":
                                switch fi25 {
                                case 32767:
                                    var v26 int8 = (v4 >> ((244) & 63))
                                    _ = v26
                                case 127:
                                    var v27 float32 = -2.25
                                    _ = v27
                                case -37:
                                    var v28 float64 = (float64((0.1 - v18)) / 1.5)
                                    _ = v28
                                }
                            case "hello":
                                var v29 int16 = v2
                                _ = v29
                            }
                            v5 = append(v5, "a")
                            _ = fi25
                        }
                        _ = v24
                    }
                }
                _ = v20
            } else {
                for fi30 := 0; fi30 < 9; fi30++ {
                    if (!(((v4 | v4) >= -128) || (!(!("hello" == ""))))) {
                        fmt.Println(v2)
                        fmt.Println(v18)
                    }
                    if false {
                        if (42 < 128) {
                            v31 := GenS2{F0: -32768, F1: 752, F2: 2}
                            v32 := v3[""]
                            _ = v31
                            _ = v32
                        } else {
                            if true {
                                v33 := GenS1{F0: 0, F1: (42 == -676), F2: 32767}
                                v34 := GenS3{F0: (fi30 ^ (fi30 + 4294967295))}
                                v18 += v18
                                _ = v33
                                _ = v34
                            }
                            v18 += (0.0 - ((v18 * (v18 + (v18 / (((-2.25 * 1.5) / 1e10) - (v18 / (0.1 * (0.1 / ((((-2.25 / 0.1) + (v18 + v18)) + ((((v18 / v18) - v18) + v18) * float32(((-2.25 / (((3.14159 / 1.5) / ((1e10 - 1e10) + (1e10 - 0.1))) - -2.25)) + 1.5)))) / float32(3.14159))))))))) + 0.1))
                            fmt.Println(v4, v4, fi19)
                            v35 := GenS4{F0: -919, F1: fi30}
                            v4 -= (v4 - v4)
                            _ = v35
                        }
                        var v36 uint8 = 0
                        v37 := GenS5{F0: 2}
                        _ = v36
                        _ = v37
                    }
                    _ = fi30
                }
                switch ("\n\t" + ("\n\t" + "\n\t")) {
                case "a":
                    var v38 uint32 = 322
                    _ = v38
                case "\n\t":
                    var v39 uint8 = 127
                    _ = v39
                default:
                    v4 += v4
                }
            }
            _ = fi19
        }
    } else {
        if len(v1) > 0 {
            v1[0] = 65535
        }
        var v40 bool = (66 == 255)
        _ = v40
    }
    for fi41 := 0; fi41 < 8; fi41++ {
        var v42 uint64 = 625
        _ = fi41
        _ = v42
    }
    var v43 int32 = 0
    v2 = v2
    v44 := GenS6{F0: 319}
    fmt.Println(v43, v4, v18)
    switch "\n\t" {
    case "":
        delete(v3, (("hello" + "") + "hello"))
    case "foo bar":
        v43 = (v43 & 42)
    case "\n\t":
        v18 = v18
    default:
        var v45 uint64 = 4294967295
        _ = v45
    }
    fmt.Println(v4)
    switch 42 {
    case 1:
        if false {
            switch "a" {
            case "a":
                fmt.Println(v18, v18, v43)
            default:
                var v46 float64 = float64(v18)
                _ = v46
            }
            if true {
                if (424 >= 65535) {
                    v47 := v3[("foo bar" + "foo bar")]
                    _ = v47
                }
                var v48 int = -9223372036854775808
                _ = v48
            }
        }
    default:
        delete(v3, ("hello" + ("" + "foo bar")))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
}
