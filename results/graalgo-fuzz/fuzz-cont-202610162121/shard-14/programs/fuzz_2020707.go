package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 uint16; F1 int; F2 uint8 }
type GenS3 struct { F0 int8; F1 int64; F2 string }
type GenS4 struct { F0 int16; F1 int16 }
func genFn0(p0 int, p1 string, p2 uint) int32 {
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 int16 = 961
        _ = fi7
        _ = v8
    }
    p2 += p2
    for fi9 := 0; fi9 < 1; fi9++ {
        fmt.Println(p1, fi9, p0)
        for fi10 := 0; fi10 < 14; fi10++ {
            var v11 int = (p0 / ((fi9) | 1))
            var v12 string = ("" + p1)
            p2 += (p2 * (p2 ^ p2))
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}
func genFn1() uint {
    var v13 float32 = 3.14159
    var v14 uint8 = 42
    _ = v13
    _ = v14
    return 255
}

func main() {
    var v1 bool = (!true)
    v2 := make(map[int]int)
    var v3 uint16 = 65535
    v4 := []uint32{2, 2147483648}
    var v5 int32 = -331
    v6 := make(map[int]string)
    for fi15 := 0; fi15 < 16; fi15++ {
        var v16 int16 = 1
        if v1 {
            v2[fi15] = 0
        }
        _ = fi15
        _ = v16
    }
    if false {
        if (18446744073709551615 > 2) {
            if v1 {
                v17 := GenS1{F0: -2147483648, F1: -1}
                _ = v17
            } else {
                if false {
                    var v18 int32 = genFn0(-1, "", 18446744073709551615)
                    _ = v18
                }
                switch "\n\t" {
                case "hello":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                case "":
                    for fi19 := 0; fi19 < 10; fi19++ {
                        for fi20 := 0; fi20 < 1; fi20++ {
                            v3 -= (^(v3))
                            if (0.0 != (1e10 * 1.5)) {
                                var v21 string = ((("\n\t" + "hello") + "\n\t") + ("hello" + "a"))
                                _ = v21
                            }
                            _ = fi20
                        }
                        _ = fi19
                    }
                }
            }
        }
        v5 = (v5 << ((255) & 63))
        var v22 int16 = -32768
        _ = v22
    }
    var v23 bool = v1
    for fi24 := 0; fi24 < 6; fi24++ {
        if v23 {
            if (2 < 0) {
                var v25 string = "\n\t"
                _ = v25
            } else {
                _ = len(v6)
                var v26 int8 = 0
                _ = v26
            }
        } else {
            v1 = v1
            for fi27 := 0; fi27 < 8; fi27++ {
                var v28 int16 = 32767
                v29 := GenS1{F0: (^(fi27)), F1: (^(v5))}
                _ = fi27
                _ = v28
                _ = v29
            }
        }
        var v30 uint32 = 805
        if (fi24 >= fi24) {
            var v31 int32 = -574
            _ = v31
        } else {
            v30 += v30
            v32 := GenS2{F0: 396, F1: (fi24 / (((fi24 & -1)) | 1)), F2: 253}
            _ = len(v2)
            _ = v32
        }
        _ = fi24
        _ = v30
    }
    for fi33 := 0; fi33 < 14; fi33++ {
        var v34 uint = 303
        _ = fi33
        _ = v34
    }
    v1 = (4294967295 < 2147483647)
    v23 = v1
    for fi35 := 0; fi35 < 5; fi35++ {
        fmt.Println(v23)
        var v36 string = ""
        v36 = (v36 + v36)
        _ = fi35
        _ = v36
    }
    var v37 int = 127
    if v23 {
        v38 := GenS3{F0: -128, F1: -1, F2: "foo bar"}
        v39 := GenS4{F0: 32767, F1: 1000}
        _ = v38
        _ = v39
    }
    fmt.Println(v1, v3)
    var v40 bool = false
    switch "" {
    case "a":
        if v23 {
            var v41 int32 = 0
            if v23 {
                v40 = v40
                if false {
                    v41 = v41
                    for fi42 := 0; fi42 < 19; fi42++ {
                        var v43 int64 = 1
                        _ = fi42
                        _ = v43
                    }
                    if v23 {
                        v2[(v37 % ((790) | 1))] = v37
                        fmt.Println(v41, v40, v5)
                        fmt.Println(v5)
                    } else {
                        var v44 int64 = 687
                        if v23 {
                            v5 = v41
                            var v45 int8 = -25
                            v45 = (v45 >> ((42) & 63))
                            _ = v45
                        } else {
                            v37 -= (^(v37))
                            var v46 string = "hello"
                            _ = v46
                        }
                        _ = v44
                    }
                } else {
                    _ = len(v2)
                }
                v47 := GenS1{F0: v37, F1: v5}
                _ = v47
            } else {
                var v48 int64 = 42
                v3 -= v3
                _ = v48
            }
            v2[v37] = 1
            _ = v41
        }
    case "foo bar":
        if (-382 == 0) {
            var v49 string = ""
            var v50 float32 = (0.0 + 0.1)
            _ = v49
            _ = v50
        }
    case "\n\t":
        v3 = (^(v3))
    default:
        v37 = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v37
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
