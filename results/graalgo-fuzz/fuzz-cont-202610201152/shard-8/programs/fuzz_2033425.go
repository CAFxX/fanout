package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint32; F1 int64 }
type GenS5 struct { F0 int32; F1 uint64 }
func genFn0(p0 bool, p1 int, p2 string) uint32 {
    fmt.Println(p0, p0)
    var v5 int = p1
    for fi6 := 0; fi6 < 5; fi6++ {
        if (1 != 0) {
            var v7 uint16 = 2
            v8 := []int{(v5 & (fi6 | int(v7))), p1, (v5 ^ v5)}
            _ = v7
            _ = v8
        } else {
            var v9 int32 = int32(p1)
            v10 := []int{fi6}
            var v11 int32 = v9
            _ = v10
            _ = v11
            _ = v9
        }
        var v12 uint16 = 42
        v5 -= p1
        _ = fi6
        _ = v12
    }
    if ((2147483647 == 2147483647) && (2 == 255)) {
        if p0 {
            var v13 int32 = 1
            if (2147483648 < 267) {
                p0 = ((!p0) || (622 < 535))
            } else {
                v14 := GenS1{F0: 42, F1: p2}
                var v15 float32 = 3.14159
                var v16 float32 = v15
                _ = v14
                _ = v15
                _ = v16
            }
            for fi17 := 0; fi17 < 9; fi17++ {
                switch v5 {
                case 1:
                    v13 += 507
                case 32768:
                    var v18 uint8 = 42
                    _ = v18
                }
                p2 = p2
                for fi19 := 0; fi19 < 11; fi19++ {
                    if p0 {
                        var v20 string = (((p2 + (((p2 + (((((p2 + "hello") + ("hello" + "")) + "") + p2) + "foo bar")) + p2) + "")) + p2) + "foo bar")
                        v21 := GenS2{F0: ("hello" < v20), F1: v13}
                        _ = v20
                        _ = v21
                    }
                    _ = fi19
                }
                _ = fi17
            }
            _ = v13
        }
        var v22 int32 = 2147483647
        v23 := GenS2{F0: (v5 > 4294967295), F1: v22}
        _ = v22
        _ = v23
    }
    p1 -= v5
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 2147483647
}
func genFn1() int64 {
    var v24 int = 9223372036854775807
    if false {
        v25 := []int{v24, (v24 | 42), v24, (v24 | 4294967295)}
        var v26 bool = false
        _ = v25
        _ = v26
    } else {
        v27 := []int{v24, v24, 9223372036854775807}
        var v28 int = (v24 << ((496) & 63))
        _ = v27
        _ = v28
    }
    if (42 <= 255) {
        v29 := []int{v24}
        _ = v29
    }
    if (313 != 9223372036854775807) {
        v24 = v24
        for fi30 := 0; fi30 < 11; fi30++ {
            v31 := []int{fi30, v24}
            var v32 string = "a"
            v33 := GenS1{F0: 42, F1: "\n\t"}
            _ = fi30
            _ = v31
            _ = v32
            _ = v33
        }
    } else {
        v34 := []int{v24, v24}
        _ = v34
    }
    for fi35 := 0; fi35 < 20; fi35++ {
        for fi36 := 0; fi36 < 10; fi36++ {
            if true {
                fmt.Println(fi36, fi36)
                if ((4294967295 >= 4294967295) && (false || (!(-1 > 1)))) {
                    var v37 int = (fi35 | fi35)
                    switch fi35 {
                    case 128:
                        v24 = 128
                    case -373:
                        v38 := GenS2{F0: (-2.25 <= -2.25), F1: 662}
                        _ = v38
                    case 32768:
                        v24 -= (fi35 & 4294967295)
                    default:
                        var v39 uint16 = 214
                        v39 -= (v39 * (v39 - v39))
                        if false {
                            v40 := []int{v37, 781, v24}
                            _ = v40
                        }
                        if false {
                            v39 = 42
                            var v41 int = v37
                            var v42 uint8 = 128
                            _ = v41
                            _ = v42
                        }
                        if false {
                            v43 := []int{543, (v24 - -998), fi36, (v37 * -9223372036854775808)}
                            _ = v43
                        } else {
                            var v44 bool = (((9223372036854775807 == 526) && false) || true)
                            v45 := GenS1{F0: 42, F1: ""}
                            _ = v44
                            _ = v45
                        }
                        v46 := []int{}
                        var v47 uint8 = 2
                        _ = v39
                        _ = v46
                        _ = v47
                    }
                    v37 = fi35
                    _ = v37
                } else {
                    var v48 int = (fi36 / (((fi35 ^ fi35)) | 1))
                    _ = v48
                }
                v24 = -9223372036854775808
                var v49 float32 = ((1e10 - 1e10) + (float32(1e10) / 0.1))
                if false {
                    fmt.Println(v24, fi36)
                }
                v50 := GenS3{F0: 1}
                if false {
                    v51 := []int{}
                    var v52 uint32 = 2147483648
                    _ = v51
                    _ = v52
                } else {
                    v24 = (fi36 << ((42) & 63))
                    v53 := GenS1{F0: 2147483647, F1: "\n\t"}
                    _ = v53
                }
                fmt.Println(v49, v49)
                _ = v49
                _ = v50
            }
            var v54 int16 = 1000
            _ = fi36
            _ = v54
        }
        v55 := []int{(fi35 * fi35), (v24 >> ((42) & 63))}
        fmt.Println(v24, v24)
        _ = fi35
        _ = v55
    }
    _ = v24
    return 788
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]uint32)
    var v3 bool = true
    var v4 float32 = 0.1
    var v56 int64 = genFn1()
    var v57 uint = 2
    switch "" {
    case "a":
        v58 := GenS4{F0: 362, F1: -1}
        _ = v58
    default:
        if false {
            v59 := v2[0]
            var v60 int32 = 2147483647
            _ = v59
            _ = v60
        } else {
            var v61 uint64 = 18446744073709551615
            var v62 int64 = (^(v56))
            var v63 float64 = ((((1e10 * ((1e10 * 1.5) - float64(0.1))) * float64((v4 + 1e10))) - 1e10) / 1.0)
            _ = v61
            _ = v62
            _ = v63
        }
    }
    var v64 int = 128
    var v65 uint32 = 2
    switch "" {
    case "\n\t":
        v66 := GenS5{F0: 730, F1: 593}
        _ = v66
    case "":
        v65 = (v65 + genFn0((!v3), -785, (("hello" + "hello") + "")))
    case "foo bar":
        switch (^(v64)) {
        case 590:
            var v67 int32 = -573
            _ = v67
        }
    default:
        var v68 bool = (490 < -2147483648)
        _ = v68
    }
    v57 = v57
    for fi69 := 0; fi69 < 10; fi69++ {
        var v70 int16 = -96
        var v71 float64 = (1.5 + 3.14159)
        switch "a" {
        case "":
            v65 = (v65 - (v65 & v65))
        case "a":
            v65 += (v65 - (v65 | genFn0(v3, (v64 % (((-(v64))) | 1)), "hello")))
        }
        _ = fi69
        _ = v70
        _ = v71
    }
    for fi72 := 0; fi72 < 18; fi72++ {
        if (1 != 4294967295) {
            v73 := []int{fi72, v64, (v64 * v64), v64}
            v73 = append(v73, v64)
            _ = v73
        } else {
            for fi74 := 0; fi74 < 6; fi74++ {
                v3 = false
                fmt.Println(v56)
                if (!true) {
                    switch (fi72 * 9223372036854775807) {
                    case -9223372036854775808:
                        if (v57 <= v57) {
                            v57 = 0
                        }
                    default:
                        v75 := GenS1{F0: v65, F1: "\n\t"}
                        _ = v75
                    }
                }
                v65 += 0
                _ = fi74
            }
            for fi76 := 0; fi76 < 12; fi76++ {
                var v77 string = ("" + ((("" + "a") + "a") + "hello"))
                _ = fi76
                _ = v77
            }
            for fi78 := 0; fi78 < 4; fi78++ {
                v4 += v4
                v4 += v4
                if v3 {
                    v56 += v56
                }
                _ = fi78
            }
        }
        v79 := GenS3{F0: 4294967295}
        _ = fi72
        _ = v79
    }
    if false {
        for fi80 := 0; fi80 < 14; fi80++ {
            for fi81 := 0; fi81 < 15; fi81++ {
                if v3 {
                    v56 -= genFn1()
                    v82 := []int{fi80, fi80, fi80}
                    v83 := GenS5{F0: 516, F1: 2}
                    _ = v82
                    _ = v83
                } else {
                    v84 := GenS4{F0: (v65 >> (((v57 % ((0) | 1))) & 63)), F1: v56}
                    _ = v84
                }
                _ = fi81
            }
            switch ("a" + "a") {
            case "\n\t":
                var v85 int64 = 78
                _ = v85
            case "a":
                var v86 float64 = 3.14159
                _ = v86
            case "foo bar":
                if ((true || false) || (0 > -1)) {
                    v57 = 2
                } else {
                    for fi87 := 0; fi87 < 7; fi87++ {
                        v65 -= genFn0(v3, (fi80 & v64), ("\n\t" + ""))
                        _ = fi87
                    }
                    fmt.Println(v57, v4, v3)
                }
            }
            _ = fi80
        }
        v3 = v3
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v56
    _ = v57
    _ = v64
    _ = v65
}
