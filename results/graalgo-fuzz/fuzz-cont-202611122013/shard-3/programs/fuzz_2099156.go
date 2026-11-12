package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16; F1 int; F2 uint16 }
type GenS5 struct { F0 uint; F1 int32; F2 bool }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int; F1 uint; F2 uint8 }
func genFn0() uint16 {
    if true {
        var v7 bool = true
        _ = v7
    } else {
        var v8 string = "hello"
        switch v8 {
        case "foo bar":
            v8 = ("a" + ((v8 + "foo bar") + ((v8 + "hello") + (v8 + v8))))
        case "\n\t":
            fmt.Println(v8)
        case "a":
            v9 := []int{9223372036854775807, 42, 42}
            _ = v9
        default:
            if (!(0 >= 29)) {
                if true {
                    var v10 uint8 = 2
                    var v11 float64 = -2.25
                    switch 42 {
                    case 502:
                        v8 = "hello"
                    case 4294967295:
                        var v12 uint64 = 18446744073709551615
                        _ = v12
                    }
                    _ = v10
                    _ = v11
                }
            } else {
                v8 = v8
            }
        }
        _ = v8
    }
    v13 := []int{-1, 32767, 4294967295}
    for fi14 := 0; fi14 < 11; fi14++ {
        var v15 int = 522
        _ = fi14
        _ = v15
    }
    _ = v13
    return 255
}
func genFn1(p0 uint16, p1 bool, p2 uint8) int64 {
    for fi16 := 0; fi16 < 1; fi16++ {
        var v17 int16 = -32768
        _ = fi16
        _ = v17
    }
    v18 := GenS1{F0: -764, F1: (("foo bar" + "a") + "hello")}
    var v19 uint64 = uint64(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v19
    return 42
}

func main() {
    var v1 uint8 = 1
    v2 := make(map[int]uint32)
    var v3 uint16 = 65535
    v4 := []int{42}
    var v5 int64 = -741
    v6 := make(map[int]int)
    fmt.Println(v1, v5, v3)
    v20 := v2[-9223372036854775808]
    var v21 uint16 = (v3 | genFn0())
    switch (("foo bar" + ("" + "\n\t")) + "\n\t") {
    case "":
        v22 := GenS1{F0: 127, F1: "a"}
        _ = v22
    }
    v5 += 849
    v21 += 9
    for fi23 := 0; fi23 < 18; fi23++ {
        for fi24 := 0; fi24 < 9; fi24++ {
            v25 := GenS1{F0: (-(fi23)), F1: ""}
            var v26 bool = false
            _ = fi24
            _ = v25
            _ = v26
        }
        if (false || false) {
            for fi27 := 0; fi27 < 17; fi27++ {
                fmt.Println(fi27)
                _ = fi27
            }
            fmt.Println(v5, v20, v5)
        } else {
            switch "" {
            case "foo bar":
                v5 = 1
            }
        }
        var v28 float32 = 3.14159
        _ = fi23
        _ = v28
    }
    if (!(1 == 1)) {
        fmt.Println(v5)
        var v29 string = ("a" + ("a" + "hello"))
        for fi30 := 0; fi30 < 16; fi30++ {
            v1 = 42
            v29 = v29
            switch fi30 {
            case -651:
                switch v29 {
                case "\n\t":
                    v31 := v6[(fi30 ^ -477)]
                    _ = v31
                case "hello":
                    if (v5 >= v5) {
                        v20 += v20
                        switch v29 {
                        case "foo bar":
                            var v32 int16 = -1
                            _ = v32
                        case "\n\t":
                            _ = len(v2)
                        }
                        v33 := GenS1{F0: (fi30 / ((146) | 1)), F1: "a"}
                        _ = v33
                    }
                default:
                    for fi34 := 0; fi34 < 2; fi34++ {
                        v29 = v29
                        if true {
                            fmt.Println(fi30, v29, v5)
                        } else {
                            v29 = v29
                            switch fi34 {
                            case 9223372036854775807:
                                v4 = append(v4, -292)
                            case 42:
                                v35 := GenS2{F0: (v20 * v20), F1: 42}
                                _ = v35
                            default:
                                v36 := GenS3{F0: (v29 + v29)}
                                _ = v36
                            }
                        }
                        if (true && true) {
                            var v37 uint8 = uint8(v20)
                            switch 32767 {
                            case 440:
                                var v38 int32 = -1
                                _ = v38
                            default:
                                if ("a" >= v29) {
                                    var v39 uint16 = (v21 ^ v3)
                                    _ = v39
                                } else {
                                    var v40 int16 = 610
                                    _ = v40
                                }
                            }
                            if (3.14159 < (1e10 * float64(1.5))) {
                                v5 = genFn1(289, (fi30 != (fi34 * -2147483648)), (v1 >> ((4294967295) & 63)))
                                v21 -= (v3 + genFn0())
                                if (true || false) {
                                    v21 = (v21 - v21)
                                    if (v1 > v37) {
                                        if true {
                                            var v41 int = fi34
                                            switch (fi30 - fi34) {
                                            case 32768:
                                                v42 := GenS1{F0: 32768, F1: v29}
                                                _ = v42
                                            case -594:
                                                var v43 int32 = 1
                                                _ = v43
                                            case -834:
                                                var v44 float64 = 1.5
                                                fmt.Println(v44, v3)
                                                var v45 int16 = -32768
                                                _ = v44
                                                _ = v45
                                            default:
                                                fmt.Println(v29, v3)
                                            }
                                            delete(v2, (^(fi34)))
                                            v3 = (v21 % (((v21 & (v3 ^ 734))) | 1))
                                            v37 = uint8(v20)
                                            v29 = "foo bar"
                                            v4 = append(v4, fi30)
                                            _ = v41
                                        }
                                        if true {
                                            var v46 float32 = 3.14159
                                            _ = v46
                                        } else {
                                            v20 += (v20 / (((v20 + 678)) | 1))
                                            v3 = 506
                                            if len(v4) > 0 {
                                                v4[0] = (fi30 & int(v1))
                                            }
                                        }
                                    } else {
                                        v37 = v1
                                        v2[fi34] = v20
                                        v37 += (v37 - v1)
                                        v47 := GenS4{F0: (^(v3)), F1: int(v20), F2: (^(v3))}
                                        _ = v47
                                    }
                                }
                                fmt.Println(v5, v21, v29)
                            }
                            var v48 float32 = -2.25
                            fmt.Println(v37, v3, v21)
                            _ = v37
                            _ = v48
                        }
                        _ = fi34
                    }
                }
            default:
                v49 := GenS3{F0: v29}
                _ = v49
            }
            _ = fi30
        }
        _ = v29
    } else {
        if ((!true) || ((1e10 <= float64(0.0)) || (4294967295 < 1))) {
            for fi50 := 0; fi50 < 15; fi50++ {
                v1 += (v1 | (v1 - (v1 >> ((996) & 63))))
                v51 := GenS5{F0: 4294967295, F1: 688, F2: (!(("\n\t" + "\n\t") == "a"))}
                _ = fi50
                _ = v51
            }
            if (!true) {
                if (32767 <= 127) {
                    switch 127 {
                    case 2147483647:
                        fmt.Println(v3, v1, v20)
                    }
                }
                fmt.Println(v1)
            } else {
                if false {
                    v5 += (v5 << ((2) & 63))
                    fmt.Println(v3, v5, v20)
                    v52 := GenS1{F0: 4294967295, F1: ""}
                    _ = v52
                }
                v1 -= 2
                var v53 uint8 = 49
                _ = v53
            }
        } else {
            fmt.Println(v1, v21)
        }
    }
    v5 -= int64(v20)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch "" {
    case "":
        v20 = 2
    case "foo bar":
        for fi54 := 0; fi54 < 3; fi54++ {
            v3 = (v21 ^ v21)
            _ = fi54
        }
    default:
        var v55 float64 = (0.0 / 1e10)
        _ = v55
    }
    v56 := GenS6{F0: uint32(v3)}
    v20 = (^(v20))
    v57 := GenS7{F0: 4294967295, F1: 0, F2: v1}
    switch 128 {
    case 127:
        fmt.Println(v1, v5)
    case 4294967295:
        delete(v2, 343)
    case -1:
        if false {
            v58 := GenS6{F0: 42}
            var v59 uint8 = 2
            _ = v58
            _ = v59
        }
    default:
        var v60 bool = (true || true)
        _ = v60
    }
    if true {
        if len(v4) > 0 {
            v4[0] = -678
        }
        for fi61 := 0; fi61 < 5; fi61++ {
            var v62 int = fi61
            _ = fi61
            _ = v62
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v57
    _ = v6
}
