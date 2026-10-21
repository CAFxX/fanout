package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 int32; F2 uint16 }
type GenS3 struct { F0 int8; F1 uint64; F2 int32 }
type GenS4 struct { F0 uint32; F1 uint16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint16; F1 string }
func genFn0() uint8 {
    var v7 uint16 = 32767
    v8 := GenS1{F0: int(v7)}
    v7 = v7
    v9 := GenS2{F0: 42, F1: 3, F2: (v7 / ((v7) | 1))}
    _ = v7
    _ = v8
    _ = v9
    return 1
}
func genFn1(p0 int8) int8 {
    p0 = p0
    var v10 int = 987
    var v11 float64 = 0.0
    _ = p0
    _ = v10
    _ = v11
    return (p0 << ((42) & 63))
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float64 = 3.14159
    v3 := make(map[string]int)
    var v4 uint32 = 4294967295
    v5 := make(map[int]int)
    var v6 uint64 = 547
    if (true || (1 == 154)) {
        v12 := GenS3{F0: 5, F1: (^(v6)), F2: 42}
        _ = v12
    } else {
        var v13 float64 = float64(-2.25)
        v14 := []int{2147483647, 32767}
        if (true && false) {
            for fi15 := 0; fi15 < 20; fi15++ {
                v13 += float64(0.0)
                var v16 uint = uint(v4)
                switch "" {
                case "foo bar":
                    for fi17 := 0; fi17 < 18; fi17++ {
                        v18 := GenS3{F0: genFn1(0), F1: 4294967295, F2: -1}
                        v14 = append(v14, (fi17 >> ((343) & 63)))
                        _ = fi17
                        _ = v18
                    }
                case "":
                    switch fi15 {
                    case 32767:
                        if true {
                            v4 += v4
                            if false {
                                var v19 uint64 = 1
                                fmt.Println(fi15, fi15)
                                var v20 float64 = (v13 - ((((v2 * v2) - v13) * ((v2 * v13) - ((((((float64(0.1) * (v2 * float64(3.14159))) - -2.25) * ((float64(1.5) + float64(3.14159)) - ((v13 * 1e10) / float64(float32(-2.25))))) * ((0.0 / v2) / (float64(0.0) / v2))) / v2) - ((v2 + float64(((1e10 + (1e10 / 3.14159)) - 1.5))) * 0.0)))) * v13))
                                _ = v19
                                _ = v20
                            }
                        }
                    case -1:
                        if true {
                            var v21 float32 = float32(v13)
                            for fi22 := 0; fi22 < 17; fi22++ {
                                v3["a"] = 128
                                _ = fi22
                            }
                            _ = v21
                        }
                    case 187:
                        var v23 int16 = 188
                        _ = v23
                    default:
                        if true {
                            switch (^(fi15)) {
                            case -1:
                                var v24 int32 = -1
                                _ = v24
                            case 1:
                                var v25 int = 4294967295
                                _ = v25
                            default:
                                var v26 uint64 = (v6 << ((730) & 63))
                                _ = v26
                            }
                        }
                    }
                }
                _ = fi15
                _ = v16
            }
            var v27 int64 = 42
            _ = v27
        } else {
            if true {
                v13 += -2.25
            }
            v28 := GenS3{F0: 42, F1: 943, F2: 1}
            _ = v28
        }
        _ = v13
        _ = v14
    }
    var v29 int8 = 0
    fmt.Println(v2, v6, v2)
    var v30 int16 = -32768
    v31 := GenS2{F0: 2147483647, F1: 42, F2: uint16(v4)}
    var v32 int8 = v29
    v32 = (^(v32))
    v33 := []int{2147483647, 292, 42, -640}
    switch 493 {
    case 9223372036854775807:
        for fi34 := 0; fi34 < 1; fi34++ {
            if true {
                var v35 uint32 = v4
                _ = v35
            } else {
                fmt.Println(v29)
                if false {
                    var v36 uint8 = 0
                    switch "foo bar" {
                    case "hello":
                        if (160 <= 1) {
                            v37 := v3["hello"]
                            v33 = append(v33, v37)
                            var v38 int32 = -871
                            _ = v37
                            _ = v38
                        } else {
                            v32 -= v32
                            var v39 string = "\n\t"
                            _ = v39
                        }
                    case "":
                        switch fi34 {
                        case 32768:
                            v40 := GenS4{F0: (v4 << ((42) & 63)), F1: 0}
                            _ = v40
                        default:
                            v2 = v2
                        }
                    case "a":
                        v29 -= (v29 | v32)
                    default:
                        fmt.Println(v30, v4)
                    }
                    _ = v36
                } else {
                    var v41 float32 = 3.14159
                    delete(v3, ("" + "a"))
                    _ = v41
                }
            }
            var v42 string = ("a" + (("\n\t" + "hello") + "a"))
            switch (("\n\t" + v42) + v42) {
            case "":
                switch int(v32) {
                case 0:
                    v43 := GenS5{F0: (v29 - v32)}
                    _ = v43
                case 4294967295:
                    v44 := GenS1{F0: (fi34 << ((311) & 63))}
                    _ = v44
                default:
                    for fi45 := 0; fi45 < 4; fi45++ {
                        fmt.Println(v32, v42)
                        _ = fi45
                    }
                }
            case "foo bar":
                v5[fi34] = 32768
            default:
                v4 = v4
            }
            _ = fi34
            _ = v42
        }
    case -89:
        if len(v33) >= 2 {
            v33 = v33[1:len(v33)]
        }
    case 127:
        v46 := GenS6{F0: 1, F1: ("a" + "hello")}
        _ = v46
    }
    var v47 uint8 = 1
    for fi48 := 0; fi48 < 6; fi48++ {
        if false {
            switch -1 {
            case -1:
                var v49 int8 = genFn1(int8(v6))
                _ = v49
            case -971:
                v47 += (v47 | genFn0())
            }
            switch (fi48 >> ((45) & 63)) {
            case -2147483648:
                v4 += v4
            case 9223372036854775807:
                v33 = append(v33, -1)
            }
        } else {
            delete(v3, ("hello" + "hello"))
            for fi50 := 0; fi50 < 15; fi50++ {
                fmt.Println(v6, v2)
                switch "hello" {
                case "a":
                    var v51 uint64 = 0
                    _ = v51
                }
                var v52 float64 = (v2 - v2)
                _ = fi50
                _ = v52
            }
            if (4294967295 != (v6 | uint64(v4))) {
                v53 := GenS2{F0: 42, F1: -2147483648, F2: uint16(v6)}
                var v54 uint8 = (v47 - v47)
                delete(v3, "\n\t")
                _ = v53
                _ = v54
            } else {
                var v55 string = "\n\t"
                _ = v55
            }
        }
        _ = fi48
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v47
    _ = v5
    _ = v6
}
