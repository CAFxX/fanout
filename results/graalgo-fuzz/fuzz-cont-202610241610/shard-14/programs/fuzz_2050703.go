package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 uint64; F1 uint32; F2 int8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32; F1 uint8 }

func main() {
    var v1 int = -1
    v2 := make(map[int]string)
    var v3 string = (("\n\t" + "a") + ("a" + ""))
    var v4 bool = true
    v5 := make(map[int]int)
    var v6 int64 = -1
    var v7 int64 = v6
    var v8 float64 = ((((float64((float32((float64(1.5) / float64((((((0.0 + 0.0) / (((0.0 - -2.25) + ((0.1 * 3.14159) / 0.1)) / (((float32((3.14159 / 1.0)) + 0.0) + -2.25) - 0.0))) / 3.14159) - 1.5) - -2.25)))) * 0.1)) + 0.1) * 1.5) + (1.5 - 1e10)) / -2.25)
    if v4 {
        v7 = 1
        for fi9 := 0; fi9 < 8; fi9++ {
            for fi10 := 0; fi10 < 19; fi10++ {
                fmt.Println(v6, fi9)
                switch fi9 {
                case -480:
                    v7 = (v7 ^ (v6 ^ 1))
                case 0:
                    var v11 uint64 = 18446744073709551615
                    _ = v11
                default:
                    v8 = (v8 - 3.14159)
                }
                _ = fi10
            }
            _ = fi9
        }
        v12 := []int{}
        _ = v12
    }
    v3 = v3
    v8 += (v8 - 1e10)
    var v13 int8 = 42
    for fi14 := 0; fi14 < 18; fi14++ {
        if v4 {
            var v15 int8 = (v13 * (v13 >> ((85) & 63)))
            v16 := GenS1{F0: v3, F1: 2}
            if ((42 < 42) || false) {
                v1 += fi14
                switch (fi14 << ((889) & 63)) {
                case -751:
                    if (!(((255 != 127) && (2147483648 > 0)) && v4)) {
                        v17 := GenS1{F0: v3, F1: 2147483647}
                        v18 := []int{9223372036854775807, 478}
                        _ = v17
                        _ = v18
                    }
                case 32768:
                    switch (v1 - (fi14 / (((-(fi14))) | 1))) {
                    case -281:
                        if true {
                            var v19 uint16 = 2
                            fmt.Println(v19, v19)
                            _ = v19
                        } else {
                            fmt.Println(v6, fi14, v15)
                        }
                    case -9223372036854775808:
                        for fi20 := 0; fi20 < 1; fi20++ {
                            var v21 bool = v4
                            _ = fi20
                            _ = v21
                        }
                    case -239:
                        if false {
                            v22 := []int{(-(v1))}
                            fmt.Println(v7, v1, v7)
                            _ = v22
                        } else {
                            fmt.Println(fi14)
                        }
                    }
                default:
                    var v23 uint8 = 0
                    _ = v23
                }
                for fi24 := 0; fi24 < 9; fi24++ {
                    var v25 uint32 = 0
                    fmt.Println(v8, fi24, v1)
                    v26 := GenS1{F0: "hello", F1: (^(v25))}
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            } else {
                switch "\n\t" {
                case "":
                    v27 := GenS1{F0: ((v3 + v3) + "\n\t"), F1: 65535}
                    _ = v27
                case "foo bar":
                    for fi28 := 0; fi28 < 5; fi28++ {
                        v5[32767] = 127
                        v7 = 1
                        var v29 int8 = v13
                        _ = fi28
                        _ = v29
                    }
                case "\n\t":
                    var v30 int64 = v7
                    _ = v30
                }
                var v31 float32 = 1.5
                _ = v31
            }
            _ = v15
            _ = v16
        } else {
            if v4 {
                if (1 == 591) {
                    fmt.Println(fi14, v7, fi14)
                    var v32 uint8 = 2
                    _ = v32
                } else {
                    if v4 {
                        for fi33 := 0; fi33 < 6; fi33++ {
                            fmt.Println(v3, fi33, v3)
                            _ = fi33
                        }
                        var v34 int64 = 42
                        _ = v34
                    } else {
                        v7 -= (v6 % ((v7) | 1))
                        var v35 int8 = 42
                        var v36 int32 = -1
                        _ = v35
                        _ = v36
                    }
                    var v37 int32 = 2147483647
                    _ = v37
                }
            } else {
                switch "a" {
                case "a":
                    v38 := []int{fi14, fi14, (fi14 % ((2147483647) | 1))}
                    _ = v38
                default:
                    v39 := GenS2{F0: 4294967295, F1: 797, F2: v13}
                    _ = v39
                }
            }
            v7 -= 9223372036854775807
            for fi40 := 0; fi40 < 15; fi40++ {
                v41 := v5[(fi40 * v1)]
                var v42 uint8 = 42
                v43 := []int{v1, v41, v41}
                _ = fi40
                _ = v41
                _ = v42
                _ = v43
            }
        }
        delete(v2, (-(v1)))
        v4 = true
        _ = fi14
    }
    fmt.Println(v6, v6, v3)
    fmt.Println(v8, v8)
    switch v3 {
    case "":
        for fi44 := 0; fi44 < 14; fi44++ {
            for fi45 := 0; fi45 < 4; fi45++ {
                v46 := GenS3{F0: v13}
                switch v3 {
                case "hello":
                    var v47 int16 = 1000
                    _ = v47
                default:
                    var v48 bool = true
                    _ = v48
                }
                var v49 float32 = 3.14159
                _ = fi45
                _ = v46
                _ = v49
            }
            switch 4294967295 {
            case 0:
                for fi50 := 0; fi50 < 9; fi50++ {
                    if true {
                        fmt.Println(v7, v13, v4)
                        v51 := []int{}
                        _ = v51
                    }
                    fmt.Println(v3, v1, v1)
                    _ = fi50
                }
            case 1:
                for fi52 := 0; fi52 < 2; fi52++ {
                    fmt.Println(v8, v8, v6)
                    var v53 int64 = (v6 - -601)
                    _ = fi52
                    _ = v53
                }
            default:
                var v54 string = "a"
                _ = v54
            }
            _ = fi44
        }
    case "\n\t":
        if false {
            for fi55 := 0; fi55 < 2; fi55++ {
                var v56 int = (fi55 ^ v1)
                if v4 {
                    v7 += (v7 + v6)
                    var v57 int32 = 898
                    fmt.Println(v57)
                    _ = v57
                } else {
                    v58 := v5[fi55]
                    var v59 int64 = v7
                    var v60 int64 = (v7 << ((4294967295) & 63))
                    _ = v58
                    _ = v59
                    _ = v60
                }
                _ = fi55
                _ = v56
            }
        }
    default:
        fmt.Println(v3)
    }
    var v61 int16 = -32768
    v62 := GenS2{F0: 408, F1: 4294967295, F2: v13}
    var v63 uint32 = 1
    v8 = v8
    if false {
        switch 752 {
        case 158:
            v64 := []int{}
            _ = v64
        case 9223372036854775807:
            v65 := []int{(v1 * v1), v1}
            _ = v65
        case 2147483647:
            v66 := GenS4{F0: 783, F1: 111}
            _ = v66
        default:
            fmt.Println(v3, v7)
        }
        var v67 bool = (v4 || false)
        _ = v67
    }
    var v68 uint16 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v61:", v61)
    fmt.Println("v63:", v63)
    fmt.Println("v68:", v68)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v61
    _ = v62
    _ = v63
    _ = v68
    _ = v7
    _ = v8
}
