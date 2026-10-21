package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64; F1 int64 }
type GenS5 struct { F0 uint8; F1 string }
type GenS6 struct { F0 int8; F1 int64; F2 uint16 }

func main() {
    var v1 float32 = 3.14159
    var v2 uint64 = 18446744073709551615
    var v3 int16 = 836
    v4 := make(map[int]bool)
    v3 = v3
    for fi5 := 0; fi5 < 8; fi5++ {
        v2 = (v2 >> ((898) & 63))
        if true {
            var v6 uint64 = v2
            _ = v6
        }
        _ = fi5
    }
    for fi7 := 0; fi7 < 4; fi7++ {
        if true {
            fmt.Println(v2)
            v1 = float32(3.14159)
        }
        _ = fi7
    }
    var v8 uint64 = (v2 - (^(v2)))
    fmt.Println(v2, v8)
    var v9 int8 = 1
    for fi10 := 0; fi10 < 14; fi10++ {
        if (45 >= 128) {
            if true {
                if ((-(fi10)) == (fi10 / ((fi10) | 1))) {
                    var v11 int64 = -505
                    v12 := v4[32767]
                    var v13 uint16 = 1
                    _ = v11
                    _ = v12
                    _ = v13
                } else {
                    if (255 >= 42) {
                        for fi14 := 0; fi14 < 12; fi14++ {
                            v15 := []int{fi10, fi10, 128}
                            if false {
                                var v16 string = "foo bar"
                                fmt.Println(fi10)
                                fmt.Println(v3, fi10)
                                var v17 string = "foo bar"
                                v4[(fi10 | fi14)] = (!false)
                                _ = v16
                                _ = v17
                            } else {
                                switch 42 {
                                case -1:
                                    var v18 uint64 = (v8 & v8)
                                    _ = v18
                                case -2147483648:
                                    var v19 int32 = 0
                                    _ = v19
                                case 128:
                                    if (197 > 127) {
                                        if (v8 <= v8) {
                                            v20 := GenS1{F0: int32(v3), F1: 8}
                                            v21 := GenS1{F0: 0, F1: 42}
                                            var v22 string = "\n\t"
                                            _ = v20
                                            _ = v21
                                            _ = v22
                                        } else {
                                            v23 := GenS2{F0: (v9 >> ((65535) & 63)), F1: 33}
                                            v24 := GenS3{F0: v9}
                                            _ = v23
                                            _ = v24
                                        }
                                    }
                                default:
                                    var v25 uint16 = 2
                                    _ = v25
                                }
                            }
                            _ = fi14
                            _ = v15
                        }
                        var v26 uint = 155
                        _ = v26
                    }
                }
                var v27 int8 = (v9 * (v9 << ((2) & 63)))
                for fi28 := 0; fi28 < 8; fi28++ {
                    if false {
                        var v29 float64 = (0.0 * 1.5)
                        _ = v29
                    } else {
                        v27 += (v27 >> ((255) & 63))
                        v2 = (v2 / (((v2 % (((v8 - (v2 + v2))) | 1))) | 1))
                        switch "" {
                        case "\n\t":
                            v30 := GenS4{F0: (v8 | 42), F1: -1}
                            _ = v30
                        case "hello":
                            var v31 float32 = -2.25
                            _ = v31
                        case "foo bar":
                            v8 += v2
                        default:
                            v8 = (v8 & (^(v2)))
                        }
                    }
                    var v32 int32 = -2147483648
                    v33 := []int{(fi28 / (((fi10 % ((int(v2)) | 1))) | 1)), int(v32), 128}
                    _ = fi28
                    _ = v32
                    _ = v33
                }
                _ = v27
            }
            v2 -= 2
            var v34 int8 = 100
            _ = v34
        } else {
            var v35 uint64 = 1
            var v36 int16 = (v3 << ((65535) & 63))
            _ = v35
            _ = v36
        }
        _ = fi10
    }
    var v37 float32 = v1
    if false {
        v3 = 32767
        var v38 int8 = v9
        switch 4294967295 {
        case 127:
            var v39 string = ""
            _ = v39
        case 32768:
            v40 := []int{4294967295, -615, 1, -9223372036854775808}
            _ = v40
        case 32767:
            v41 := GenS4{F0: 1, F1: 9223372036854775807}
            _ = v41
        }
        _ = v38
    }
    for fi42 := 0; fi42 < 17; fi42++ {
        var v43 uint = 18446744073709551615
        _ = fi42
        _ = v43
    }
    if (false && (!((!(!(true && true))) || (false || (-2147483648 <= 32767))))) {
        fmt.Println(v2, v37, v37)
        v3 += 32767
        for fi44 := 0; fi44 < 17; fi44++ {
            if (!((fi44 % (((fi44 | fi44)) | 1)) != fi44)) {
                if (3.14159 == 1.5) {
                    if false {
                        v45 := GenS5{F0: 127, F1: ""}
                        var v46 uint16 = 352
                        v1 -= float32(1.5)
                        _ = v45
                        _ = v46
                    }
                    for fi47 := 0; fi47 < 14; fi47++ {
                        var v48 uint32 = 42
                        _ = fi47
                        _ = v48
                    }
                }
                fmt.Println(fi44)
            } else {
                v49 := []int{(-(fi44)), 4294967295}
                _ = v49
            }
            _ = fi44
        }
    } else {
        v1 = v37
        v50 := GenS6{F0: v9, F1: int64(v3), F2: 0}
        _ = v50
    }
    switch "a" {
    case "":
        fmt.Println(v2, v3, v1)
    case "foo bar":
        var v51 uint64 = (^(v2))
        _ = v51
    }
    var v52 float32 = 1.5
    fmt.Println(v37, v52, v1)
    v1 += float32((1.5 - 3.14159))
    var v53 uint8 = 1
    v54 := GenS6{F0: v9, F1: 42, F2: 255}
    switch "foo bar" {
    case "a":
        delete(v4, 931)
    default:
        if false {
            v37 += (v37 * -2.25)
            var v55 uint32 = 0
            _ = v55
        } else {
            var v56 int8 = (v9 | (v9 % ((42) | 1)))
            _ = v56
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v52
    _ = v53
    _ = v54
    _ = v8
    _ = v9
}
