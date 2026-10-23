package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 int64; F2 int }
type GenS3 struct { F0 string; F1 uint16; F2 uint16 }
type GenS4 struct { F0 int8; F1 uint8 }
type GenS5 struct { F0 string }
func genFn0(p0 uint16) int {
    if false {
        var v6 int16 = 780
        _ = v6
    }
    if false {
        var v7 int = 0
        p0 = 255
        v8 := []int{v7, (v7 << ((65535) & 63)), v7}
        _ = v7
        _ = v8
    }
    if (2 > 2) {
        var v9 float64 = (0.1 + (((float64(float32(0.0)) / 1e10) * -2.25) * 0.1))
        v10 := GenS1{F0: 2}
        _ = v10
        _ = v9
    }
    _ = p0
    return -1
}
func genFn1(p0 int16, p1 int32, p2 int) int16 {
    var v11 float32 = 3.14159
    var v12 float64 = float64(v11)
    var v13 uint32 = 863
    switch p2 {
    case -1:
        var v14 float64 = v12
        _ = v14
    case 9223372036854775807:
        if true {
            for fi15 := 0; fi15 < 2; fi15++ {
                v16 := GenS2{F0: 0, F1: -9223372036854775808, F2: genFn0(65535)}
                var v17 uint16 = 771
                _ = fi15
                _ = v16
                _ = v17
            }
            v12 = (v12 / 1.5)
            for fi18 := 0; fi18 < 18; fi18++ {
                for fi19 := 0; fi19 < 6; fi19++ {
                    var v20 string = ""
                    var v21 uint32 = v13
                    v22 := GenS1{F0: 16}
                    _ = fi19
                    _ = v20
                    _ = v21
                    _ = v22
                }
                _ = fi18
            }
        } else {
            if true {
                var v23 int8 = 1
                fmt.Println(p2, v12)
                switch "\n\t" {
                case "foo bar":
                    var v24 float64 = (v12 / (1e10 / ((v12 / v12) + ((0.0 - 1e10) / (v12 / v12)))))
                    _ = v24
                }
                _ = v23
            } else {
                v25 := GenS3{F0: (("hello" + ((("hello" + "foo bar") + "") + "foo bar")) + ("" + "\n\t")), F1: 593, F2: 1}
                _ = v25
            }
        }
    default:
        v26 := []int{genFn0(838), (p2 | p2)}
        _ = v26
    }
    v27 := []int{(p2 << ((20) & 63)), (^(p2))}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v13
    _ = v27
    return p0
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []int{}
    var v3 int16 = -32768
    v4 := make(map[int]uint32)
    var v5 uint8 = 104
    for fi28 := 0; fi28 < 17; fi28++ {
        var v29 uint = 2
        if true {
            v3 = (v3 - (-(v3)))
        } else {
            _ = len(v2)
        }
        _ = fi28
        _ = v29
    }
    v5 = (v5 & (v5 + v5))
    var v30 uint64 = 146
    for fi31 := 0; fi31 < 13; fi31++ {
        var v32 string = ""
        v5 = v5
        _ = fi31
        _ = v32
    }
    v3 = (v3 - genFn1(genFn1(int16(v30), 42, -552), 2147483647, -606))
    for fi33 := 0; fi33 < 4; fi33++ {
        if true {
            var v34 int64 = -9223372036854775808
            var v35 bool = false
            _ = v34
            _ = v35
        }
        var v36 uint16 = 32767
        v30 -= (v30 ^ 337)
        _ = fi33
        _ = v36
    }
    if (51 == 127) {
        var v37 string = "\n\t"
        _ = v37
    } else {
        v38 := GenS1{F0: (^(v5))}
        v30 = 4294967295
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v38
    }
    for fi39 := 0; fi39 < 1; fi39++ {
        v40 := GenS4{F0: int8(v5), F1: (v5 << ((116) & 63))}
        if false {
            if (!(((1 <= 0) || (-1 <= 0)) || (true && (!true)))) {
                if (true || true) {
                    v1[("hello" + (("" + "hello") + "hello"))] = 896
                }
            } else {
                if false {
                    if (true || (!false)) {
                        var v41 int16 = -844
                        if (3.14159 != (0.0 * (((1.5 / 1.5) * ((0.0 - 0.0) - ((1e10 * (1e10 - 3.14159)) + 1.5))) + 3.14159))) {
                            v42 := GenS1{F0: v5}
                            fmt.Println(fi39, fi39, fi39)
                            _ = v42
                        } else {
                            v41 += int16(fi39)
                        }
                        v30 = (v30 & (v30 ^ 622))
                        _ = v41
                    }
                }
            }
            switch "hello" {
            case "\n\t":
                v43 := GenS2{F0: -1, F1: 535, F2: (fi39 / ((genFn0(347)) | 1))}
                _ = v43
            case "foo bar":
                switch (^(fi39)) {
                case -665:
                    var v44 string = "hello"
                    _ = v44
                case -9223372036854775808:
                    v45 := GenS5{F0: "foo bar"}
                    _ = v45
                default:
                    switch "" {
                    case "a":
                        if (0.1 < 0.0) {
                            var v46 bool = (1 <= 4294967295)
                            _ = v46
                        } else {
                            v47 := GenS3{F0: "", F1: 65535, F2: uint16(v30)}
                            v48 := GenS5{F0: ""}
                            if (!false) {
                                if false {
                                    for fi49 := 0; fi49 < 3; fi49++ {
                                        if (float32(3.14159) < 1e10) {
                                            var v50 uint8 = (v5 ^ (^(v5)))
                                            v5 -= 2
                                            var v51 uint = 65535
                                            _ = v50
                                            _ = v51
                                        } else {
                                            switch genFn0(533) {
                                            case -394:
                                                fmt.Println(v3, v30)
                                            default:
                                                var v52 int8 = -113
                                                delete(v4, (fi49 + fi39))
                                                _ = v52
                                            }
                                        }
                                        var v53 int = fi39
                                        _ = fi49
                                        _ = v53
                                    }
                                    if (v30 > 42) {
                                        v30 = (v30 * (^(v30)))
                                    }
                                }
                            }
                            _ = v47
                            _ = v48
                        }
                    case "\n\t":
                        if len(v2) > 0 {
                            v2[0] = fi39
                        }
                    }
                }
            default:
                if true {
                    var v54 int = (fi39 >> ((255) & 63))
                    _ = v54
                }
            }
            var v55 string = "hello"
            _ = v55
        }
        var v56 int16 = 205
        _ = fi39
        _ = v40
        _ = v56
    }
    v5 -= (v5 + 2)
    v2 = append(v2, 128)
    var v57 int8 = 1
    v5 = (^(v5))
    var v58 uint16 = 1
    var v59 uint64 = 0
    v2 = append(v2, int(v3))
    var v60 float64 = (1e10 / -2.25)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v57
    _ = v58
    _ = v59
    _ = v60
}
