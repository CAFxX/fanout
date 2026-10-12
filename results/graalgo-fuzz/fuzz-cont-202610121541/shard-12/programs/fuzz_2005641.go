package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64; F1 int16; F2 uint32 }
type GenS5 struct { F0 uint64; F1 bool; F2 uint }
type GenS6 struct { F0 int; F1 int8 }
func genFn0(p0 string, p1 int) uint {
    fmt.Println(p1, p1)
    p1 = (p1 | (p1 * (p1 ^ 769)))
    fmt.Println(p0)
    _ = p0
    _ = p1
    return 255
}

func main() {
    var v1 float64 = (1.5 + ((float64(-2.25) * 1e10) - (1e10 - 1.5)))
    v2 := make(map[int]uint32)
    var v3 uint32 = 560
    var v4 uint32 = v3
    var v5 int64 = 0
    v6 := []string{"", "a"}
    v1 -= v1
    var v7 float32 = 0.1
    v8 := GenS1{F0: 4294967295}
    for fi9 := 0; fi9 < 11; fi9++ {
        var v10 int16 = 168
        _ = fi9
        _ = v10
    }
    var v11 int16 = 818
    v4 += v4
    if true {
        switch 1 {
        case 32767:
            v1 -= (-2.25 * 3.14159)
        default:
            var v12 float64 = (float64(3.14159) - ((float64(-2.25) + v1) * v1))
            _ = v12
        }
    }
    v4 -= v3
    fmt.Println(v5, v3, v7)
    v13 := GenS2{F0: 127}
    fmt.Println(v5)
    for fi14 := 0; fi14 < 8; fi14++ {
        v6 = append(v6, "\n\t")
        var v15 int32 = -1
        _ = fi14
        _ = v15
    }
    var v16 int8 = 127
    for fi17 := 0; fi17 < 13; fi17++ {
        _ = len(v6)
        v2[fi17] = v3
        _ = fi17
    }
    for fi18 := 0; fi18 < 10; fi18++ {
        var v19 int = (fi18 | 2147483647)
        if (22 != 0) {
            for fi20 := 0; fi20 < 10; fi20++ {
                switch (fi18 >> ((4294967295) & 63)) {
                case 808:
                    v21 := GenS3{F0: 128}
                    _ = v21
                }
                if (true || (v11 >= v11)) {
                    if ((v16 / ((127) | 1)) == 0) {
                        v7 -= 1e10
                        if false {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            v22 := GenS1{F0: 42}
                            var v23 bool = ((int64(v4) < v5) || true)
                            _ = v22
                            _ = v23
                        } else {
                            v24 := GenS2{F0: v19}
                            if ((15 > 128) || true) {
                                v1 = (float64((-2.25 + (((v7 * v7) + float32(((v1 - ((0.1 + (((float64(((0.1 / float32((v1 * ((0.1 + (v1 - 3.14159)) * (v1 / v1))))) * v7)) * 3.14159) / 1.5) * -2.25)) + 3.14159)) / float64(v7)))) / v7))) * -2.25)
                            } else {
                                var v25 bool = (642 > 57)
                                _ = v25
                            }
                            v2[fi20] = 0
                            var v26 int32 = 0
                            v19 = v19
                            var v27 uint64 = 42
                            if true {
                                switch v19 {
                                case 127:
                                    v1 += v1
                                    fmt.Println(v5, v11, v16)
                                    v19 = 32767
                                case 32767:
                                    fmt.Println(fi18, v3)
                                case 1:
                                    v28 := v2[v19]
                                    _ = v28
                                }
                            } else {
                                var v29 uint16 = 32767
                                v27 = 2
                                if (((1e10 + (v1 + (v1 * v1))) > (v1 - v1)) || ((true && ((fi20 + (fi18 & int(v5))) <= fi18)) && false)) {
                                    var v30 uint64 = 0
                                    v19 = 32768
                                    var v31 uint64 = (v27 + (v30 + v27))
                                    switch ("\n\t" + "\n\t") {
                                    case "a":
                                        v6 = append(v6, "hello")
                                    case "hello":
                                        v31 = (v31 - (v30 ^ v30))
                                        v2[int(v27)] = v3
                                        if true {
                                            v29 = v29
                                            fmt.Println(v11, v27)
                                            var v32 uint32 = v3
                                            _ = v32
                                        } else {
                                            fmt.Println(v11)
                                        }
                                    }
                                    _ = v30
                                    _ = v31
                                } else {
                                    if false {
                                        v16 = (v16 & v16)
                                    } else {
                                        if (!(v16 != v16)) {
                                            v33 := GenS3{F0: fi18}
                                            _ = v33
                                        }
                                    }
                                }
                                var v34 string = ("" + "hello")
                                switch ("hello" + ("foo bar" + v34)) {
                                case "foo bar":
                                    if (true && false) {
                                        fmt.Println(v16, v29)
                                        if true {
                                            if (!(v29 > 255)) {
                                                if false {
                                                    v3 = v3
                                                    v35 := GenS4{F0: (v5 * 1), F1: 1000, F2: v4}
                                                    delete(v2, (fi18 | 177))
                                                    _ = v35
                                                } else {
                                                    delete(v2, 42)
                                                    if true {
                                                        v11 = 42
                                                    } else {
                                                        var v36 int32 = (v26 & v26)
                                                        fmt.Println(v5, v1, v19)
                                                        var v37 int16 = v11
                                                        _ = v36
                                                        _ = v37
                                                    }
                                                }
                                            }
                                            v27 -= 42
                                            if true {
                                                var v38 int8 = (-(v16))
                                                fmt.Println(v38, v19, v4)
                                                var v39 float64 = ((((v1 * 3.14159) + 0.0) - float64(v7)) / float64(3.14159))
                                                if false {
                                                    v5 = v5
                                                } else {
                                                    v26 -= v26
                                                    v40 := GenS5{F0: (v27 + (v27 >> ((672) & 63))), F1: (!false), F2: 65535}
                                                    _ = v40
                                                }
                                                _ = v38
                                                _ = v39
                                            }
                                        }
                                    } else {
                                        var v41 int32 = -1
                                        _ = v41
                                    }
                                case "\n\t":
                                    v6 = append(v6, ("foo bar" + v34))
                                }
                                _ = v29
                                _ = v34
                            }
                            _ = v24
                            _ = v26
                            _ = v27
                        }
                        v19 += fi18
                        fmt.Println(v3, v11, v16)
                    }
                    v11 = 678
                    _ = len(v6)
                } else {
                    var v42 int32 = 2147483647
                    var v43 uint64 = 2
                    switch v19 {
                    case -9223372036854775808:
                        v44 := GenS4{F0: v5, F1: (v11 * (v11 / (((v11 - -32768)) | 1))), F2: 795}
                        _ = v44
                    case -723:
                        if (0.0 > (((v1 / v1) * v1) + ((v1 / ((v1 - v1) * float64((float32((3.14159 * v1)) / 1e10)))) - float64(float32((v1 + 1.5)))))) {
                            var v45 uint8 = 125
                            v19 = fi18
                            switch (fi20 + fi18) {
                            case 128:
                                v46 := GenS6{F0: (fi20 >> ((255) & 63)), F1: (^(v16))}
                                _ = v46
                            default:
                                v6 = append(v6, (((("foo bar" + "a") + "a") + "hello") + "a"))
                                v45 -= (v45 - v45)
                            }
                            v45 = (v45 % ((255) | 1))
                            _ = v45
                        } else {
                            var v47 bool = true
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            _ = v47
                        }
                    default:
                        var v48 bool = (!true)
                        _ = v48
                    }
                    _ = v42
                    _ = v43
                }
                _ = fi20
            }
            _ = len(v6)
        }
        _ = fi18
        _ = v19
    }
    if true {
        v2[30] = (v3 + v4)
    } else {
        var v49 uint = 65535
        var v50 int32 = -1
        _ = v49
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
