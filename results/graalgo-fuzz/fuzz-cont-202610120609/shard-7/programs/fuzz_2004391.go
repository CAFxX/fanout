package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 string }
type GenS2 struct { F0 uint; F1 uint32; F2 int32 }
type GenS3 struct { F0 bool; F1 int16; F2 uint32 }
type GenS4 struct { F0 uint8 }
func genFn0() bool {
    var v7 float32 = 3.14159
    for fi8 := 0; fi8 < 1; fi8++ {
        var v9 uint = 18446744073709551615
        _ = fi8
        _ = v9
    }
    v7 += (v7 - (-2.25 * v7))
    var v10 uint8 = 246
    _ = v10
    _ = v7
    return true
}
func genFn1(p0 uint64, p1 uint32, p2 string) uint64 {
    if (-731 != -417) {
        p1 -= p1
        var v11 uint16 = 32767
        var v12 string = ((("" + "hello") + ("hello" + "hello")) + p2)
        _ = v11
        _ = v12
    } else {
        for fi13 := 0; fi13 < 14; fi13++ {
            v14 := []int{(^(fi13)), (fi13 << ((81) & 63)), -1, -668}
            var v15 int64 = 9223372036854775807
            switch p2 {
            case "a":
                fmt.Println(p1)
            default:
                var v16 int8 = 127
                _ = v16
            }
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    p0 += (p0 >> ((2) & 63))
    _ = p0
    _ = p1
    _ = p2
    return (p0 * p0)
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 bool = true
    v3 := []string{}
    v4 := make(map[int]int)
    var v5 uint = 504
    var v6 int16 = 340
    v5 += v5
    v17 := GenS1{F0: 0, F1: 101, F2: "hello"}
    fmt.Println(v6)
    if v2 {
        v18 := GenS2{F0: 0, F1: 65535, F2: 211}
        if false {
            var v19 float32 = 0.0
            v20 := GenS2{F0: (v5 ^ v5), F1: 4294967295, F2: 2147483647}
            v5 = 306
            _ = v19
            _ = v20
        }
        if (!((v5 ^ v5) > v5)) {
            _ = len(v3)
            for fi21 := 0; fi21 < 16; fi21++ {
                _ = len(v4)
                switch (^(fi21)) {
                case 32768:
                    fmt.Println(v2)
                }
                var v22 uint32 = 4294967295
                _ = fi21
                _ = v22
            }
        }
        _ = v18
    } else {
        var v23 uint8 = 0
        _ = v23
    }
    v6 = (^(v6))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v5 += (^(v5))
    v24 := v1[42]
    for fi25 := 0; fi25 < 10; fi25++ {
        if ((v2 && (!(uint8(v24) > 173))) || v2) {
            v5 = v5
            var v26 string = ("a" + "foo bar")
            _ = v26
        } else {
            var v27 float32 = 3.14159
            var v28 int = -813
            _ = v27
            _ = v28
        }
        _ = fi25
    }
    v29 := GenS3{F0: false, F1: v6, F2: v24}
    var v30 string = "foo bar"
    var v31 int = 42
    v2 = (v2 && false)
    for fi32 := 0; fi32 < 13; fi32++ {
        if (true && true) {
            v33 := GenS4{F0: 0}
            for fi34 := 0; fi34 < 17; fi34++ {
                if true {
                    var v35 uint = 42
                    switch (v31 - (v31 * fi34)) {
                    case 32767:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case -103:
                        switch ("" + (v30 + "\n\t")) {
                        case "\n\t":
                            var v36 int64 = 432
                            _ = v36
                        case "a":
                            fmt.Println(v30)
                        case "":
                            fmt.Println(fi32, v5, v30)
                        }
                    case 1:
                        if (!v2) {
                            v35 = (v35 & v35)
                            v35 -= (v5 + 682)
                        } else {
                            v4[(v31 - (-(v31)))] = -941
                        }
                    }
                    _ = v35
                } else {
                    fmt.Println(v5, v6, v2)
                }
                if (!(((v5 % (((^(v5))) | 1)) >= v5) && (!true))) {
                    v37 := GenS3{F0: v2, F1: (v6 | v6), F2: (^(v24))}
                    v6 = 1000
                    _ = v37
                }
                _ = fi34
            }
            var v38 uint8 = 127
            _ = v33
            _ = v38
        }
        if (0 >= 42) {
            if v2 {
                v31 += -2147483648
                v30 = v30
                var v39 int16 = int16(v24)
                _ = v39
            }
            v40 := GenS1{F0: -624, F1: 0, F2: v30}
            _ = v40
        } else {
            v24 = 887
        }
        switch (v31 * (^(v31))) {
        case 885:
            if v2 {
                v41 := GenS3{F0: (1 <= int64(fi32)), F1: v6, F2: (v24 + (v24 & v24))}
                _ = v41
            } else {
                switch v30 {
                case "hello":
                    fmt.Println(fi32)
                case "a":
                    var v42 int32 = -648
                    _ = v42
                case "":
                    switch "" {
                    case "\n\t":
                        for fi43 := 0; fi43 < 20; fi43++ {
                            v30 = v30
                            _ = fi43
                        }
                    case "hello":
                        fmt.Println(fi32, v6)
                    default:
                        v3 = append(v3, (("foo bar" + (((v30 + (v30 + "hello")) + "hello") + "a")) + "a"))
                    }
                }
                if (255 != uint16(v5)) {
                    if len(v3) > 0 {
                        v3[0] = v30
                    }
                    v44 := GenS1{F0: 0, F1: 1, F2: v30}
                    _ = v44
                }
            }
        case 1:
            if (255 != 255) {
                for fi45 := 0; fi45 < 11; fi45++ {
                    v46 := GenS4{F0: 0}
                    v47 := v4[fi32]
                    _ = fi45
                    _ = v46
                    _ = v47
                }
            }
        case 309:
            switch 4294967295 {
            case -2147483648:
                switch fi32 {
                case 2147483647:
                    var v48 uint = 547
                    _ = v48
                case 975:
                    for fi49 := 0; fi49 < 20; fi49++ {
                        v50 := GenS4{F0: 2}
                        _ = fi49
                        _ = v50
                    }
                default:
                    v51 := GenS2{F0: v5, F1: v24, F2: 42}
                    _ = v51
                }
            }
        default:
            v52 := GenS3{F0: (((v30 + v30) >= ("a" + v30)) || v2), F1: (v6 << ((18446744073709551615) & 63)), F2: (v24 / ((65535) | 1))}
            _ = v52
        }
        _ = fi32
    }
    var v53 bool = (((0.0 + (-2.25 + (float32((3.14159 * 1e10)) * (float32(1e10) + (0.1 + (float32(((0.1 * 3.14159) - 0.0)) + (3.14159 * 3.14159))))))) / (((1.5 * 1.5) * (1e10 * (1.5 + float32(3.14159)))) * -2.25)) <= 0.0)
    v2 = v53
    var v54 uint16 = 546
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v24
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v6
}
