package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 int32 }
type GenS2 struct { F0 uint8; F1 int64; F2 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint; F1 uint16; F2 uint8 }
type GenS5 struct { F0 uint8; F1 uint32; F2 uint32 }
func genFn0(p0 uint64, p1 string, p2 int) int8 {
    fmt.Println(p2, p1)
    var v5 int = p2
    p0 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return int8(p0)
}
func genFn1(p0 uint32) int16 {
    v6 := []int{2147483647}
    if true {
        var v7 uint = 161
        _ = v7
    } else {
        var v8 uint = 42
        var v9 float32 = 1e10
        v10 := GenS1{F0: true, F1: 643, F2: -688}
        _ = v10
        _ = v8
        _ = v9
    }
    switch (("foo bar" + "a") + ("a" + "a")) {
    case "a":
        if true {
            if (("\n\t" + ("foo bar" + "foo bar")) <= "\n\t") {
                v6 = append(v6, -2147483648)
                if (-2.25 <= 0.1) {
                    var v11 uint16 = 65535
                    _ = v11
                }
            } else {
                var v12 int = 4294967295
                p0 += 42
                p0 = p0
                _ = v12
            }
        } else {
            for fi13 := 0; fi13 < 7; fi13++ {
                fmt.Println(fi13)
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                var v14 uint = 65535
                _ = fi13
                _ = v14
            }
            fmt.Println(p0)
            var v15 float32 = (float32(1e10) - (1e10 * 0.0))
            _ = v15
        }
    case "":
        p0 = (p0 & (p0 / ((p0) | 1)))
    default:
        p0 = p0
    }
    v16 := GenS2{F0: 1, F1: 368, F2: (1e10 <= 1e10)}
    var v17 int = 128
    _ = p0
    _ = v16
    _ = v17
    _ = v6
    return 42
}

func main() {
    var v1 float32 = (((1e10 - 1e10) + 1.5) * 1e10)
    v2 := make(map[int]string)
    var v3 uint8 = 128
    var v4 uint16 = 291
    switch 127 {
    case 32767:
        if false {
            fmt.Println(v1, v3, v4)
            var v18 bool = false
            if v18 {
                v4 -= v4
                v19 := GenS1{F0: (v18 && true), F1: -32768, F2: 1}
                v3 = v3
                _ = v19
            } else {
                v20 := GenS2{F0: (v3 - v3), F1: 42, F2: (!true)}
                _ = v20
            }
            _ = v18
        }
    }
    if false {
        _ = len(v2)
    } else {
        var v21 int64 = -1
        for fi22 := 0; fi22 < 12; fi22++ {
            v23 := GenS2{F0: 17, F1: v21, F2: (127 >= 127)}
            fmt.Println(fi22, fi22)
            v21 += v21
            _ = fi22
            _ = v23
        }
        var v24 bool = (false || false)
        _ = v21
        _ = v24
    }
    fmt.Println(v4, v1, v1)
    v25 := []int{4294967295}
    var v26 int8 = 42
    var v27 int64 = -681
    for fi28 := 0; fi28 < 3; fi28++ {
        v27 = v27
        _ = fi28
    }
    switch -61 {
    case 185:
        var v29 uint32 = 341
        _ = v29
    case 9223372036854775807:
        for fi30 := 0; fi30 < 8; fi30++ {
            switch (fi30 - fi30) {
            case 32767:
                v31 := GenS3{F0: -1}
                _ = v31
            case 219:
                _ = len(v25)
            }
            v32 := GenS4{F0: 0, F1: 42, F2: 0}
            if (fi30 < (^(fi30))) {
                if (v3 != (v3 | (^(v3)))) {
                    fmt.Println(v27)
                    fmt.Println(v1, v3)
                    for fi33 := 0; fi33 < 20; fi33++ {
                        fmt.Println(v26)
                        v34 := v2[fi33]
                        v35 := GenS3{F0: int16(v26)}
                        _ = fi33
                        _ = v34
                        _ = v35
                    }
                }
            } else {
                v3 += v3
                for fi36 := 0; fi36 < 9; fi36++ {
                    if (((3.14159 / (0.0 + (1.5 + 3.14159))) * (3.14159 / 3.14159)) >= (((0.1 + float64(1.5)) - (0.1 / float64(-2.25))) * (1e10 + 0.1))) {
                        var v37 int8 = (v26 ^ v26)
                        if true {
                            var v38 string = ("\n\t" + "foo bar")
                            _ = v38
                        }
                        _ = v37
                    } else {
                        if false {
                            if false {
                                var v39 int = (fi36 << ((2) & 63))
                                _ = v39
                            } else {
                                v40 := GenS2{F0: v3, F1: v27, F2: false}
                                var v41 uint32 = 42
                                var v42 uint64 = 42
                                _ = v40
                                _ = v41
                                _ = v42
                            }
                            v1 = v1
                            if false {
                                v4 -= (v4 ^ (v4 * v4))
                                var v43 uint32 = 4294967295
                                switch "foo bar" {
                                case "foo bar":
                                    switch (fi30 & fi30) {
                                    case 128:
                                        fmt.Println(v3)
                                    default:
                                        var v44 float32 = (0.1 / float32(((3.14159 * 3.14159) / (1.5 / (float64(v1) - (0.1 + 3.14159))))))
                                        _ = v44
                                    }
                                case "\n\t":
                                    fmt.Println(v27)
                                }
                                _ = v43
                            } else {
                                var v45 int16 = genFn1(2147483648)
                                _ = v45
                            }
                            var v46 uint32 = 0
                            _ = v46
                        } else {
                            var v47 float64 = 1.5
                            _ = v47
                        }
                        var v48 int64 = v27
                        v49 := GenS5{F0: 255, F1: 979, F2: 2147483648}
                        _ = v48
                        _ = v49
                    }
                    var v50 string = ("foo bar" + "\n\t")
                    _ = fi36
                    _ = v50
                }
            }
            _ = fi30
            _ = v32
        }
    case 32767:
        var v51 int32 = 1
        _ = v51
    default:
        v26 = (v26 ^ (v26 % ((v26) | 1)))
    }
    var v52 uint64 = 0
    fmt.Println(v52)
    delete(v2, 0)
    switch 42 {
    case 2147483647:
        var v53 bool = true
        _ = v53
    default:
        for fi54 := 0; fi54 < 11; fi54++ {
            for fi55 := 0; fi55 < 5; fi55++ {
                v56 := v2[(fi54 / ((fi54) | 1))]
                v56 = (v56 + v56)
                var v57 bool = (false && (0 >= 1))
                _ = fi55
                _ = v56
                _ = v57
            }
            _ = fi54
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v52
}
