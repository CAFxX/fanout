package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 string; F1 uint64; F2 int }
type GenS5 struct { F0 int16; F1 string; F2 uint }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint; F1 int; F2 int16 }
func genFn0(p0 uint32, p1 uint, p2 bool) bool {
    v6 := []int{32767, 4294967295, 32767, 1}
    var v7 string = "hello"
    if len(v6) > 0 {
        v6[0] = 2147483647
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (!p2)
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[int]string)
    var v3 int16 = 1000
    var v4 int16 = (v3 & 1000)
    v5 := make(map[string]string)
    v4 -= v4
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 float64 = 0.1
        _ = fi8
        _ = v9
    }
    v3 = 31
    var v10 string = "\n\t"
    if (!false) {
        if true {
            for fi11 := 0; fi11 < 19; fi11++ {
                v3 = 0
                if true {
                    v12 := GenS1{F0: v1}
                    _ = v12
                }
                v10 = "foo bar"
                _ = fi11
            }
            fmt.Println(v10)
        }
    } else {
        var v13 uint32 = 1
        var v14 int32 = -980
        _ = v13
        _ = v14
    }
    var v15 float32 = (float32((0.0 / 1.0)) * 0.1)
    v16 := []int{-9223372036854775808, -1, 127, 42}
    v10 = ((v10 + "hello") + v10)
    if true {
        for fi17 := 0; fi17 < 17; fi17++ {
            switch v10 {
            case "a":
                var v18 int64 = (v1 | (v1 ^ (v1 - v1)))
                _ = v18
            case "\n\t":
                if false {
                    for fi19 := 0; fi19 < 3; fi19++ {
                        delete(v5, v10)
                        _ = fi19
                    }
                    for fi20 := 0; fi20 < 5; fi20++ {
                        var v21 int64 = (v1 << ((65535) & 63))
                        _ = fi20
                        _ = v21
                    }
                    var v22 int = (fi17 >> ((2) & 63))
                    _ = v22
                } else {
                    var v23 bool = false
                    switch -767 {
                    case 127:
                        if ((42 > 1) || (127 > 100)) {
                            var v24 string = v10
                            switch v24 {
                            case "\n\t":
                                for fi25 := 0; fi25 < 1; fi25++ {
                                    if v23 {
                                        v26 := GenS1{F0: 42}
                                        if len(v16) >= 2 {
                                            v16 = v16[1:len(v16)]
                                        }
                                        _ = v26
                                    } else {
                                        var v27 int16 = (v4 * (v3 % ((v4) | 1)))
                                        v3 += v4
                                        _ = len(v16)
                                        _ = v27
                                    }
                                    if false {
                                        if (0 > 65535) {
                                            v4 += 1
                                            if v23 {
                                                if ("foo bar" == (v24 + "")) {
                                                    var v28 float64 = -2.25
                                                    _ = v28
                                                } else {
                                                    switch fi25 {
                                                    case -9223372036854775808:
                                                        var v29 uint32 = 1
                                                        _ = v29
                                                    case -2147483648:
                                                        if v23 {
                                                            if v23 {
                                                                v30 := GenS2{F0: 929, F1: (v24 + "")}
                                                                var v31 uint32 = 319
                                                                var v32 int8 = 40
                                                                if (1 > 0) {
                                                                    var v33 uint64 = 281
                                                                    _ = v33
                                                                }
                                                                switch "foo bar" {
                                                                case "foo bar":
                                                                    if (1 <= 65535) {
                                                                        switch (^(fi17)) {
                                                                        case 4294967295:
                                                                            if len(v16) >= 2 {
                                                                                v16 = v16[1:len(v16)]
                                                                            }
                                                                        case 42:
                                                                            var v34 uint16 = 303
                                                                            _ = v34
                                                                        default:
                                                                            var v35 uint64 = 2
                                                                            _ = v35
                                                                        }
                                                                        var v36 int = (fi25 >> ((uint(v31)) & 63))
                                                                        _ = v36
                                                                    }
                                                                case "hello":
                                                                    v23 = v23
                                                                }
                                                                _ = v30
                                                                _ = v31
                                                                _ = v32
                                                            }
                                                        } else {
                                                            fmt.Println(v23, v15, v4)
                                                            _ = len(v5)
                                                            var v37 int16 = v4
                                                            _ = v37
                                                        }
                                                    case 32767:
                                                        v4 = 1000
                                                    }
                                                    v23 = v23
                                                    v38 := GenS3{F0: 2}
                                                    v39 := GenS4{F0: (v10 + v10), F1: 0, F2: fi25}
                                                    _ = v38
                                                    _ = v39
                                                }
                                            } else {
                                                var v40 float32 = v15
                                                switch ("\n\t" + ("\n\t" + ("" + ("a" + "foo bar")))) {
                                                case "hello":
                                                    _ = len(v2)
                                                case "a":
                                                    v4 += (v4 - v4)
                                                case "\n\t":
                                                    v41 := GenS5{F0: v3, F1: (v10 + v24), F2: 4294967295}
                                                    _ = v41
                                                }
                                                _ = v40
                                            }
                                        }
                                        if v23 {
                                            switch fi17 {
                                            case 32767:
                                                v10 = v24
                                            case 1:
                                                var v42 float32 = -2.25
                                                _ = v42
                                            default:
                                                var v43 string = v24
                                                _ = v43
                                            }
                                            if true {
                                                v23 = v23
                                                var v44 int32 = 0
                                                v45 := GenS6{F0: 0}
                                                _ = v44
                                                _ = v45
                                            } else {
                                                var v46 uint16 = 255
                                                _ = v46
                                            }
                                            v47 := GenS7{F0: 639, F1: fi17, F2: (-(v3))}
                                            _ = v47
                                        } else {
                                            if ("a" >= "") {
                                                var v48 uint64 = 18446744073709551615
                                                _ = v48
                                            }
                                        }
                                    }
                                    v16 = append(v16, (fi17 << ((762) & 63)))
                                    _ = fi25
                                }
                            case "hello":
                                _ = len(v16)
                            default:
                                var v49 float64 = 1e10
                                _ = v49
                            }
                            _ = v24
                        } else {
                            var v50 uint64 = 64
                            _ = v50
                        }
                    case 9223372036854775807:
                        fmt.Println(v1)
                    default:
                        if true {
                            for fi51 := 0; fi51 < 18; fi51++ {
                                v3 += (v4 ^ v4)
                                _ = fi51
                            }
                            var v52 float64 = 1.5
                            _ = v52
                        }
                    }
                    v10 = (v10 + v10)
                    _ = v23
                }
            case "hello":
                v2[32767] = ("a" + ("foo bar" + ("a" + v10)))
            }
            for fi53 := 0; fi53 < 10; fi53++ {
                if ((v1 << ((42) & 63)) <= (v1 % (((^(v1))) | 1))) {
                    v15 += float32(3.14159)
                    v54 := v2[fi53]
                    if (0.1 != v15) {
                        var v55 float32 = ((v15 * (v15 + (((((v15 + v15) * v15) * 1e10) * float32(3.14159)) + 0.0))) * v15)
                        v54 = "foo bar"
                        _ = v55
                    } else {
                        fmt.Println(v15)
                    }
                    _ = v54
                }
                v1 -= v1
                _ = fi53
            }
            var v56 int16 = 1
            _ = fi17
            _ = v56
        }
        switch 4294967295 {
        case -2147483648:
            _ = len(v2)
        case -1:
            var v57 float32 = float32(3.14159)
            _ = v57
        case -9223372036854775808:
            fmt.Println(v1, v3, v1)
        }
        delete(v2, 32768)
    }
    fmt.Println(v15, v3)
    for fi58 := 0; fi58 < 13; fi58++ {
        for fi59 := 0; fi59 < 17; fi59++ {
            var v60 uint16 = 65535
            v61 := GenS5{F0: (v3 / ((v3) | 1)), F1: "", F2: 873}
            v16 = append(v16, 4294967295)
            _ = fi59
            _ = v60
            _ = v61
        }
        v4 = int16(v1)
        _ = fi58
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
