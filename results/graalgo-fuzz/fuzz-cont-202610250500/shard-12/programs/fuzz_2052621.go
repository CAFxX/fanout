package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 uint16 }
type GenS2 struct { F0 int32; F1 int8; F2 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool; F1 int8 }
type GenS5 struct { F0 int32; F1 uint32 }
func genFn0(p0 uint8, p1 uint) string {
    var v6 int64 = -1
    v7 := []int{}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return "\n\t"
}
func genFn1(p0 uint64, p1 int64) string {
    var v8 uint16 = 42
    if true {
        for fi9 := 0; fi9 < 8; fi9++ {
            for fi10 := 0; fi10 < 10; fi10++ {
                var v11 float64 = 0.1
                if (1 >= 309) {
                    var v12 int16 = 108
                    v13 := GenS1{F0: 74, F1: 127, F2: 255}
                    _ = v12
                    _ = v13
                } else {
                    var v14 int32 = 797
                    var v15 int8 = 0
                    _ = v14
                    _ = v15
                }
                _ = fi10
                _ = v11
            }
            var v16 uint8 = 42
            var v17 float32 = (0.1 / ((1.5 * (3.14159 * 0.1)) / ((0.1 - -2.25) - 0.0)))
            _ = fi9
            _ = v16
            _ = v17
        }
        fmt.Println(v8, v8)
    }
    v18 := GenS1{F0: 1, F1: 103, F2: v8}
    _ = p0
    _ = p1
    _ = v18
    _ = v8
    return "foo bar"
}

func main() {
    var v1 uint32 = 4294967295
    v2 := make(map[string]int)
    var v3 int16 = -1
    var v4 int64 = 399
    v5 := []uint32{(^(v1))}
    v1 = (v1 % ((2147483647) | 1))
    v3 += v3
    for fi19 := 0; fi19 < 16; fi19++ {
        var v20 uint16 = 0
        var v21 uint16 = (^(v20))
        _ = fi19
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 3; fi22++ {
        _ = len(v2)
        fmt.Println(v1, fi22, v1)
        for fi23 := 0; fi23 < 6; fi23++ {
            var v24 float32 = 1e10
            fmt.Println(fi23, fi23, v4)
            _ = fi23
            _ = v24
        }
        _ = fi22
    }
    v1 += 753
    if (!(true && true)) {
        switch "foo bar" {
        case "hello":
            if (true && (1 != (v4 + int64(v1)))) {
                if (42 == -532) {
                    if false {
                        v25 := GenS1{F0: 42, F1: -97, F2: 889}
                        fmt.Println(v1, v3)
                        var v26 uint16 = 998
                        _ = v25
                        _ = v26
                    }
                }
            }
        case "foo bar":
            for fi27 := 0; fi27 < 3; fi27++ {
                v28 := GenS1{F0: 0, F1: 0, F2: 243}
                var v29 int64 = v4
                if false {
                    if ((v29 >> ((787) & 63)) != (v29 / ((9223372036854775807) | 1))) {
                        for fi30 := 0; fi30 < 20; fi30++ {
                            v31 := GenS2{F0: -136, F1: 100, F2: -1}
                            if (false || true) {
                                if (0.1 < (1e10 / float32((((-2.25 / 1e10) / 1.0) / (1.5 - -2.25))))) {
                                    v32 := GenS3{F0: fi30}
                                    v4 -= v4
                                    var v33 uint32 = v1
                                    fmt.Println(v33, v33, v29)
                                    v34 := v2[""]
                                    _ = v32
                                    _ = v33
                                    _ = v34
                                } else {
                                    var v35 float32 = 1e10
                                    var v36 uint = 1
                                    switch fi27 {
                                    case -589:
                                        var v37 uint = (^(v36))
                                        _ = v37
                                    case 254:
                                        fmt.Println(fi27, v1, v4)
                                    }
                                    _ = v35
                                    _ = v36
                                }
                                var v38 int64 = (v4 / (((v29 * (v29 % ((1) | 1)))) | 1))
                                v39 := GenS3{F0: (-(fi27))}
                                _ = v38
                                _ = v39
                            } else {
                                switch fi30 {
                                case 0:
                                    switch fi30 {
                                    case 875:
                                        var v40 string = (((("" + "a") + "") + "a") + "hello")
                                        _ = v40
                                    case 4294967295:
                                        if true {
                                            if (1.5 < (1e10 - float64((0.0 * (3.14159 * -2.25))))) {
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                v1 += v1
                                                var v41 bool = false
                                                _ = v41
                                            }
                                            var v42 float32 = 0.1
                                            var v43 float32 = -2.25
                                            _ = v42
                                            _ = v43
                                        }
                                    default:
                                        v3 -= (v3 & v3)
                                    }
                                case 4294967295:
                                    var v44 int = fi30
                                    _ = v44
                                }
                                v29 = (v4 % ((v4) | 1))
                                v29 -= 1
                            }
                            v29 += 729
                            _ = fi30
                            _ = v31
                        }
                        var v45 int32 = 1
                        _ = v45
                    }
                } else {
                    if len(v5) > 0 {
                        v5[0] = v1
                    }
                }
                _ = fi27
                _ = v28
                _ = v29
            }
        }
        v3 = v3
    } else {
        v3 -= -181
    }
    fmt.Println(v1, v1, v1)
    var v46 float32 = 3.14159
    var v47 bool = false
    var v48 int = 127
    var v49 uint64 = 240
    switch ("\n\t" + "a") {
    case "\n\t":
        for fi50 := 0; fi50 < 19; fi50++ {
            switch ("foo bar" + "foo bar") {
            case "a":
                var v51 int8 = -43
                _ = v51
            default:
                if (1 < 37) {
                    var v52 uint32 = v1
                    v52 -= v1
                    _ = v52
                } else {
                    v1 = (v1 & v1)
                    var v53 float64 = 3.14159
                    switch int(v3) {
                    case 4294967295:
                        v54 := v2["\n\t"]
                        _ = v54
                    case 975:
                        v55 := GenS1{F0: 1, F1: -96, F2: 47}
                        _ = v55
                    default:
                        var v56 int8 = -58
                        _ = v56
                    }
                    _ = v53
                }
            }
            _ = fi50
        }
    case "":
        v57 := GenS4{F0: v47, F1: -1}
        _ = v57
    case "foo bar":
        v4 = -631
    default:
        delete(v2, "foo bar")
    }
    fmt.Println(v4)
    fmt.Println(v47, v49, v47)
    for fi58 := 0; fi58 < 8; fi58++ {
        var v59 uint16 = 32767
        _ = len(v5)
        _ = fi58
        _ = v59
    }
    for fi60 := 0; fi60 < 14; fi60++ {
        v48 = (fi60 - (v48 & fi60))
        v61 := GenS5{F0: 1, F1: (v1 / ((v1) | 1))}
        _ = fi60
        _ = v61
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
}
