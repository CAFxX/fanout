package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint; F1 int32; F2 string }
func genFn0(p0 int32, p1 uint) int {
    v7 := GenS1{F0: 918}
    p1 = 52
    for fi8 := 0; fi8 < 16; fi8++ {
        fmt.Println(p1)
        var v9 int64 = 1
        _ = fi8
        _ = v9
    }
    v10 := []int{0, 32768, 4294967295, -9223372036854775808}
    for fi11 := 0; fi11 < 7; fi11++ {
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        p1 = p1
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    return 42
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int = 9223372036854775807
    v3 := []string{"", "a"}
    var v4 uint32 = 1
    var v5 uint64 = 42
    v6 := make(map[int]uint32)
    var v12 int8 = -128
    var v13 uint16 = 32767
    delete(v6, v2)
    if (((!false) && false) || (255 == 209)) {
        v14 := GenS1{F0: 917}
        v5 = v5
        v15 := GenS1{F0: -750}
        _ = v14
        _ = v15
    } else {
        var v16 uint = 18446744073709551615
        var v17 float32 = (-2.25 / (0.1 / (0.1 + -2.25)))
        _ = v16
        _ = v17
    }
    v18 := GenS1{F0: 9223372036854775807}
    var v19 float32 = -2.25
    switch "hello" {
    case "":
        fmt.Println(v12)
    case "\n\t":
        fmt.Println(v19, v12, v4)
    case "a":
        for fi20 := 0; fi20 < 5; fi20++ {
            fmt.Println(v12, v5, v13)
            var v21 int = genFn0(int32(v12), 201)
            fmt.Println(v12, fi20, v2)
            _ = fi20
            _ = v21
        }
    default:
        v12 += 0
    }
    switch v2 {
    case 0:
        v13 = v13
    case -1:
        for fi22 := 0; fi22 < 2; fi22++ {
            var v23 int64 = 0
            if (int8(v13) == v12) {
                var v24 int64 = -1
                if true {
                    var v25 float32 = 3.14159
                    v24 = v23
                    _ = v25
                } else {
                    if (-2.25 >= (v19 + (float32(-2.25) / 1.5))) {
                        var v26 int64 = -141
                        if false {
                            switch fi22 {
                            case 128:
                                if len(v3) > 0 {
                                    v3[0] = ("hello" + "a")
                                }
                            default:
                                fmt.Println(fi22, v19, v19)
                            }
                            var v27 uint8 = 0
                            _ = v27
                        } else {
                            if (42 == (^(v13))) {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                            for fi28 := 0; fi28 < 5; fi28++ {
                                if true {
                                    if true {
                                        if (-128 != -128) {
                                            fmt.Println(fi28, v4)
                                            switch (("foo bar" + "foo bar") + "") {
                                            case "a":
                                                if true {
                                                    var v29 int8 = (v12 + v12)
                                                    fmt.Println(v29, v29, v29)
                                                    fmt.Println(v5)
                                                    _ = v29
                                                }
                                                v23 = (v24 % ((v23) | 1))
                                            case "\n\t":
                                                fmt.Println(v26, v23, v23)
                                            default:
                                                var v30 uint32 = 228
                                                _ = v30
                                            }
                                            var v31 uint = 18446744073709551615
                                            _ = v31
                                        } else {
                                            if false {
                                                var v32 uint8 = 127
                                                var v33 uint64 = (v5 + (^(v5)))
                                                var v34 int16 = -819
                                                _ = v32
                                                _ = v33
                                                _ = v34
                                            }
                                            if true {
                                                v35 := GenS1{F0: int64(v4)}
                                                switch ("\n\t" + "\n\t") {
                                                case "hello":
                                                    v5 = (^(v5))
                                                case "":
                                                    delete(v1, 2147483647)
                                                case "a":
                                                    fmt.Println(v13, v13, v26)
                                                default:
                                                    fmt.Println(v2, v19)
                                                }
                                                if ((3.14159 / (1e10 / (1e10 / 0.1))) < -2.25) {
                                                    fmt.Println(v26, fi22)
                                                    if true {
                                                        v2 = (fi28 ^ v2)
                                                    } else {
                                                        v6[(v2 / ((fi28) | 1))] = v4
                                                    }
                                                    var v36 bool = false
                                                    _ = v36
                                                } else {
                                                    if false {
                                                        if (v2 > (fi28 % (((v2 % (((fi22 / ((fi28) | 1))) | 1))) | 1))) {
                                                            v23 = int64(v4)
                                                            var v37 int64 = v24
                                                            v24 = v24
                                                            _ = v37
                                                        }
                                                        var v38 uint64 = 18446744073709551615
                                                        if false {
                                                            var v39 bool = true
                                                            if true {
                                                                _ = len(v3)
                                                                _ = len(v3)
                                                                if len(v3) > 0 {
                                                                    v3[0] = "\n\t"
                                                                }
                                                            } else {
                                                                v19 = v19
                                                                if v39 {
                                                                    delete(v6, genFn0(-544, uint(v4)))
                                                                    var v40 int32 = -2147483648
                                                                    if false {
                                                                        var v41 int = (fi28 >> ((562) & 63))
                                                                        _ = v41
                                                                    } else {
                                                                        v2 += fi22
                                                                    }
                                                                    var v42 float32 = v19
                                                                    _ = v40
                                                                    _ = v42
                                                                } else {
                                                                    _ = len(v6)
                                                                    fmt.Println(v2)
                                                                    switch "a" {
                                                                    case "foo bar":
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                    case "a":
                                                                        v3 = append(v3, "hello")
                                                                    case "":
                                                                        var v43 uint32 = (v4 * (v4 + 4294967295))
                                                                        _ = v43
                                                                    default:
                                                                        if (v39 || true) {
                                                                            v19 = (1.5 / float32(-2.25))
                                                                            var v44 int16 = -353
                                                                            _ = v44
                                                                        } else {
                                                                            v6[genFn0(2147483647, 255)] = 977
                                                                        }
                                                                    }
                                                                }
                                                                v45 := v1[4294967295]
                                                                _ = v45
                                                            }
                                                            v46 := GenS1{F0: (v23 ^ 1)}
                                                            _ = v39
                                                            _ = v46
                                                        } else {
                                                            v12 += (v12 * 1)
                                                        }
                                                        _ = v38
                                                    }
                                                    var v47 uint64 = 0
                                                    var v48 bool = true
                                                    v49 := GenS2{F0: v12}
                                                    _ = v47
                                                    _ = v48
                                                    _ = v49
                                                }
                                                v19 -= v19
                                                v5 = (v5 | (^(v5)))
                                                _ = v35
                                            }
                                        }
                                        if (!true) {
                                            switch (^(fi28)) {
                                            case 4294967295:
                                                v50 := GenS2{F0: (v12 / ((int8(v13)) | 1))}
                                                _ = v50
                                            case 32768:
                                                v6[v2] = v4
                                            case 32767:
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            }
                                        } else {
                                            fmt.Println(v5, v12)
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            v19 = (v19 / v19)
                                            v51 := GenS3{F0: -747}
                                            _ = v51
                                        }
                                        fmt.Println(fi22)
                                    }
                                    var v52 float64 = 1e10
                                    _ = v52
                                } else {
                                    v53 := GenS3{F0: 942}
                                    _ = v53
                                }
                                if len(v3) > 0 {
                                    v3[0] = (("\n\t" + "a") + "foo bar")
                                }
                                var v54 bool = (!(!("" != "a")))
                                _ = fi28
                                _ = v54
                            }
                        }
                        _ = v26
                    } else {
                        for fi55 := 0; fi55 < 6; fi55++ {
                            v4 = uint32(v13)
                            _ = fi55
                        }
                    }
                }
                if false {
                    if true {
                        var v56 int64 = v24
                        _ = v56
                    }
                    if ((-1 > -32768) || true) {
                        var v57 int32 = -2147483648
                        _ = v57
                    } else {
                        var v58 uint = 1
                        v5 += v5
                        _ = v58
                    }
                    fmt.Println(v19, v19)
                } else {
                    switch (fi22 - (fi22 * v2)) {
                    case 63:
                        v23 -= v24
                    }
                    for fi59 := 0; fi59 < 8; fi59++ {
                        v24 += v24
                        var v60 int64 = (v24 / ((v23) | 1))
                        switch ("hello" + "hello") {
                        case "\n\t":
                            switch (fi59 % ((fi22) | 1)) {
                            case 9223372036854775807:
                                if false {
                                    var v61 uint64 = (^(v5))
                                    v12 -= 0
                                    _ = v61
                                } else {
                                    switch (^(v2)) {
                                    case 128:
                                        _ = len(v3)
                                    case -319:
                                        v62 := GenS3{F0: 42}
                                        v60 -= v24
                                        if true {
                                            fmt.Println(v5)
                                            var v63 int64 = int64(v13)
                                            delete(v6, fi59)
                                            _ = v63
                                        } else {
                                            var v64 uint8 = 1
                                            var v65 int8 = (v12 >> ((2) & 63))
                                            _ = v64
                                            _ = v65
                                        }
                                        _ = v62
                                    case 2147483647:
                                        v6[fi22] = 65535
                                    default:
                                        var v66 int = (v2 - -85)
                                        _ = v66
                                    }
                                }
                            case 295:
                                var v67 bool = true
                                _ = v67
                            default:
                                v68 := GenS1{F0: (v60 % ((-357) | 1))}
                                _ = v68
                            }
                        default:
                            var v69 float32 = v19
                            _ = v69
                        }
                        _ = len(v3)
                        v13 += v13
                        _ = fi59
                        _ = v60
                    }
                    v6[fi22] = v4
                }
                _ = v24
            } else {
                for fi70 := 0; fi70 < 10; fi70++ {
                    var v71 uint = 0
                    _ = fi70
                    _ = v71
                }
            }
            _ = fi22
            _ = v23
        }
    default:
        switch "\n\t" {
        case "\n\t":
            for fi72 := 0; fi72 < 18; fi72++ {
                for fi73 := 0; fi73 < 16; fi73++ {
                    var v74 bool = ((18446744073709551615 != 593) || (v19 >= 0.1))
                    _ = fi73
                    _ = v74
                }
                var v75 int64 = 1
                v13 = 320
                _ = fi72
                _ = v75
            }
        }
    }
    var v76 uint32 = v4
    v13 -= (v13 * v13)
    var v77 float32 = (v19 + 3.14159)
    var v78 int = (^(v2))
    if (2 != 0) {
        v5 = (^(v5))
        v79 := GenS4{F0: 65535, F1: -1, F2: ("\n\t" + "")}
        _ = len(v3)
        _ = v79
    }
    for fi80 := 0; fi80 < 10; fi80++ {
        v5 -= v5
        v3 = append(v3, ("\n\t" + "hello"))
        v81 := GenS2{F0: v12}
        _ = fi80
        _ = v81
    }
    v13 += 155
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v76
    _ = v77
    _ = v78
}
