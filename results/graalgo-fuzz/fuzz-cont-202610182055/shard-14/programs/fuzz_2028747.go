package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 bool; F1 bool; F2 bool }
func genFn0(p0 uint8) string {
    var v6 uint32 = 393
    v7 := []int{-14, 127}
    for fi8 := 0; fi8 < 20; fi8++ {
        for fi9 := 0; fi9 < 19; fi9++ {
            v6 += (v6 ^ (^(v6)))
            if false {
                if len(v7) > 0 {
                    v7[0] = fi9
                }
                if len(v7) > 0 {
                    v7[0] = (fi8 + int(v6))
                }
                if (!true) {
                    fmt.Println(fi9, fi9)
                    var v10 int32 = 0
                    if ((true && (!false)) && true) {
                        if len(v7) > 0 {
                            v7[0] = fi9
                        }
                    }
                    _ = v10
                } else {
                    v11 := GenS1{F0: 4294967295}
                    p0 -= p0
                    switch 9223372036854775807 {
                    case 0:
                        if false {
                            fmt.Println(fi8)
                            v6 = (v6 / ((v6) | 1))
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                        } else {
                            v12 := GenS1{F0: 1}
                            if (false && (0 > 1)) {
                                if (false && true) {
                                    var v13 int64 = -9223372036854775808
                                    v14 := GenS1{F0: v6}
                                    _ = v13
                                    _ = v14
                                } else {
                                    v15 := GenS2{F0: 797}
                                    if false {
                                        switch (("foo bar" + "hello") + ((("\n\t" + "a") + ("foo bar" + "foo bar")) + "foo bar")) {
                                        case "\n\t":
                                            var v16 int8 = 1
                                            _ = v16
                                        case "a":
                                            v6 -= (v6 >> ((834) & 63))
                                        default:
                                            if true {
                                                v7 = append(v7, fi8)
                                                if len(v7) >= 2 {
                                                    v7 = v7[1:len(v7)]
                                                }
                                                var v17 int32 = 0
                                                _ = v17
                                            } else {
                                                v18 := GenS1{F0: 481}
                                                if (0 >= 42) {
                                                    p0 -= 107
                                                    var v19 float32 = 3.14159
                                                    var v20 uint8 = p0
                                                    _ = v19
                                                    _ = v20
                                                } else {
                                                    switch (fi8 + fi8) {
                                                    case 9223372036854775807:
                                                        if (false || (474 == -387)) {
                                                            fmt.Println(fi8, v6, fi8)
                                                            fmt.Println(v6, fi8)
                                                        } else {
                                                            v7 = append(v7, fi9)
                                                        }
                                                    }
                                                    p0 += 42
                                                }
                                                _ = v18
                                            }
                                        }
                                        v7 = append(v7, (^(fi9)))
                                        v7 = append(v7, (fi9 | fi9))
                                    } else {
                                        p0 -= (p0 - 0)
                                    }
                                    var v21 uint = 255
                                    var v22 string = "\n\t"
                                    var v23 uint16 = 42
                                    _ = v15
                                    _ = v21
                                    _ = v22
                                    _ = v23
                                }
                                p0 -= 2
                                var v24 float32 = 3.14159
                                _ = v24
                            } else {
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                                _ = len(v7)
                            }
                            p0 = uint8(v6)
                            _ = v12
                        }
                    case 128:
                        switch (fi9 / ((int(v6)) | 1)) {
                        case 2147483647:
                            p0 += p0
                        }
                    default:
                        v6 = v6
                    }
                    _ = v11
                }
            }
            var v25 float64 = (1e10 * 3.14159)
            _ = fi9
            _ = v25
        }
        _ = fi8
    }
    _ = p0
    _ = v6
    _ = v7
    return "\n\t"
}
func genFn1() uint8 {
    for fi26 := 0; fi26 < 5; fi26++ {
        v27 := []int{}
        if false {
            for fi28 := 0; fi28 < 7; fi28++ {
                if len(v27) > 0 {
                    v27[0] = fi26
                }
                var v29 uint = 65535
                v29 = (v29 | (^(v29)))
                _ = fi28
                _ = v29
            }
            var v30 int = fi26
            for fi31 := 0; fi31 < 17; fi31++ {
                fmt.Println(fi31)
                v30 = -2147483648
                fmt.Println(fi26, fi26)
                _ = fi31
            }
            _ = v30
        } else {
            var v32 int = (fi26 << ((1) & 63))
            if len(v27) >= 2 {
                v27 = v27[1:len(v27)]
            }
            _ = v32
        }
        for fi33 := 0; fi33 < 15; fi33++ {
            var v34 uint8 = 244
            var v35 int16 = -32768
            if false {
                var v36 int64 = 9223372036854775807
                _ = v36
            }
            _ = fi33
            _ = v34
            _ = v35
        }
        _ = fi26
        _ = v27
    }
    var v37 int32 = -2147483648
    _ = v37
    return 127
}

func main() {
    var v1 uint64 = 771
    v2 := make(map[string]uint32)
    var v3 int16 = -1
    var v4 int8 = 119
    v5 := []uint32{1, 113, 2, 2}
    if false {
        switch "hello" {
        case "a":
            v38 := GenS2{F0: (v1 + (v1 | (v1 % ((v1) | 1))))}
            _ = v38
        case "foo bar":
            fmt.Println(v4, v1)
        default:
            v3 += (v3 ^ 42)
        }
        for fi39 := 0; fi39 < 10; fi39++ {
            v4 += v4
            if len(v5) > 0 {
                v5[0] = 0
            }
            _ = fi39
        }
    } else {
        v4 -= (v4 + -1)
    }
    v40 := GenS2{F0: v1}
    var v41 uint = 18446744073709551615
    v4 += 27
    switch 128 {
    case 127:
        for fi42 := 0; fi42 < 5; fi42++ {
            for fi43 := 0; fi43 < 2; fi43++ {
                switch "a" {
                case "\n\t":
                    if false {
                        v41 = v41
                        if false {
                            switch "foo bar" {
                            case "":
                                fmt.Println(v1)
                            case "foo bar":
                                delete(v2, "")
                            }
                            v44 := v2["foo bar"]
                            _ = v44
                        }
                        switch (fi42 - fi42) {
                        case 32767:
                            var v45 int32 = 2147483647
                            _ = v45
                        default:
                            v46 := GenS1{F0: 464}
                            _ = v46
                        }
                    }
                    fmt.Println(v3, fi43)
                case "hello":
                    delete(v2, ((("" + "a") + ("\n\t" + ("a" + ""))) + "hello"))
                }
                _ = fi43
            }
            var v47 int64 = 9223372036854775807
            _ = fi42
            _ = v47
        }
    case 9223372036854775807:
        v1 -= v1
    case -2147483648:
        fmt.Println(v3)
    default:
        var v48 uint32 = 313
        _ = v48
    }
    v3 += v3
    var v49 uint8 = 169
    v50 := GenS3{F0: (^(v41)), F1: (v4 * (v4 % ((v4) | 1)))}
    delete(v2, "a")
    var v51 int32 = -1
    v52 := GenS4{F0: false, F1: (float64(float32(((3.14159 / 1.0) * ((((1e10 - 3.14159) / 0.1) + 3.14159) / float64(0.1))))) == -2.25), F2: (((^(v41)) != (v41 >> (((^(v41))) & 63))) || (0 < 127))}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
}
