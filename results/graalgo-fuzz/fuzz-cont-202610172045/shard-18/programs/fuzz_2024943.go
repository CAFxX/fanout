package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32; F1 int }
type GenS5 struct { F0 int16; F1 int64 }
type GenS6 struct { F0 int64; F1 int64 }
func genFn0() string {
    var v7 int = 127
    fmt.Println(v7)
    v7 -= v7
    v7 = (-(v7))
    v7 = 32768
    _ = v7
    return ""
}
func genFn1(p0 int, p1 uint16, p2 int8) int8 {
    p1 += p1
    var v8 uint16 = 65535
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return p2
}

func main() {
    v1 := make(map[int]string)
    var v2 float64 = (float64(float32(float64(0.1))) / 1.5)
    v3 := make(map[string]string)
    var v4 int64 = -1
    v5 := make(map[int]int)
    v6 := make(map[string]bool)
    var v9 int = 230
    v10 := GenS1{F0: (v9 / (((v9 & v9)) | 1))}
    v2 -= v2
    v4 += (v4 << ((326) & 63))
    var v11 int8 = genFn1(v9, 65535, 0)
    if true {
        v4 -= v4
        for fi12 := 0; fi12 < 4; fi12++ {
            if false {
                if false {
                    switch v9 {
                    case 0:
                        if (!false) {
                            fmt.Println(v4)
                            v4 = v4
                            fmt.Println(v4)
                        } else {
                            v13 := GenS2{F0: 42, F1: -877}
                            switch -273 {
                            case -9223372036854775808:
                                var v14 int16 = 32767
                                _ = v14
                            case 0:
                                switch "" {
                                case "\n\t":
                                    var v15 uint8 = 255
                                    _ = v15
                                case "hello":
                                    switch "a" {
                                    case "a":
                                        v9 -= (fi12 >> ((65535) & 63))
                                    case "foo bar":
                                        if true {
                                            var v16 bool = (255 <= 241)
                                            var v17 int16 = -1
                                            v18 := []int{v9, v9, (v9 + fi12)}
                                            _ = v16
                                            _ = v17
                                            _ = v18
                                        } else {
                                            var v19 float64 = float64(0.1)
                                            switch fi12 {
                                            case 32768:
                                                fmt.Println(v9, v11, v2)
                                            case 3:
                                                v4 = (^(v4))
                                            }
                                            switch "foo bar" {
                                            case "":
                                                for fi20 := 0; fi20 < 8; fi20++ {
                                                    if (((true && false) && true) || true) {
                                                        var v21 int = fi20
                                                        if true {
                                                            v22 := GenS1{F0: (^(v9))}
                                                            v23 := GenS1{F0: (v9 | (^(v9)))}
                                                            _ = v22
                                                            _ = v23
                                                        } else {
                                                            if ((42 <= 0) || ((25 < 0) || false)) {
                                                                switch (("foo bar" + "a") + "hello") {
                                                                case "":
                                                                    v24 := GenS2{F0: 73, F1: -1}
                                                                    _ = v24
                                                                case "\n\t":
                                                                    fmt.Println(v4, v2)
                                                                }
                                                                v25 := []int{(v9 | fi12), v9, (-(fi20))}
                                                                if true {
                                                                    var v26 int64 = (v4 ^ v4)
                                                                    _ = v26
                                                                }
                                                                _ = v25
                                                            } else {
                                                                if false {
                                                                    v27 := GenS2{F0: 128, F1: 1000}
                                                                    _ = v27
                                                                } else {
                                                                    var v28 int8 = 127
                                                                    v2 -= ((1.5 / v2) / (v2 * 1.5))
                                                                    var v29 uint32 = 42
                                                                    _ = v28
                                                                    _ = v29
                                                                }
                                                                if false {
                                                                    v21 = (fi20 & v21)
                                                                    var v30 float32 = ((0.1 * (((-2.25 / 3.14159) + 3.14159) / 1.5)) / (3.14159 + -2.25))
                                                                    _ = v30
                                                                }
                                                            }
                                                            v21 = 32768
                                                            var v31 int = 1
                                                            _ = v31
                                                        }
                                                        v4 = v4
                                                        _ = v21
                                                    } else {
                                                        fmt.Println(v9, v19, v9)
                                                        v4 = (v4 & (v4 ^ (v4 & (-(v4)))))
                                                        if (116 > 42) {
                                                            delete(v5, (fi12 - (-(fi12))))
                                                            var v32 string = "foo bar"
                                                            _ = v32
                                                        } else {
                                                            fmt.Println(v4)
                                                            v19 = float64(-2.25)
                                                            v19 = v2
                                                        }
                                                    }
                                                    v33 := GenS1{F0: v9}
                                                    v34 := GenS3{F0: int16(v11)}
                                                    var v35 float32 = 3.14159
                                                    v36 := GenS1{F0: int(v11)}
                                                    _ = fi20
                                                    _ = v33
                                                    _ = v34
                                                    _ = v35
                                                    _ = v36
                                                }
                                            case "\n\t":
                                                for fi37 := 0; fi37 < 8; fi37++ {
                                                    if (((0 != 925) && true) || (32767 > 255)) {
                                                        var v38 int = fi12
                                                        if false {
                                                            if (true && false) {
                                                                v2 -= v19
                                                                v39 := GenS4{F0: 0, F1: v9}
                                                                _ = v39
                                                            }
                                                            v40 := v5[fi12]
                                                            if (0.1 <= 1.5) {
                                                                switch fi37 {
                                                                case 0:
                                                                    v11 += (v11 | -1)
                                                                case -2147483648:
                                                                    if (0.1 < 0.0) {
                                                                        var v41 int64 = (v4 ^ (v4 ^ (-(v4))))
                                                                        _ = v41
                                                                    } else {
                                                                        v19 = (v19 + float64((((0.1 + ((-2.25 * 0.0) * -2.25)) * (0.1 * -2.25)) * -2.25)))
                                                                        v4 = 9223372036854775807
                                                                    }
                                                                }
                                                                var v42 int8 = v11
                                                                if (false && true) {
                                                                    var v43 int8 = v42
                                                                    _ = v43
                                                                } else {
                                                                    if (!(!true)) {
                                                                        switch (fi12 * 800) {
                                                                        case -933:
                                                                            switch int(v4) {
                                                                            case 127:
                                                                                var v44 uint8 = 42
                                                                                _ = v44
                                                                            default:
                                                                                if true {
                                                                                    v45 := []int{v9}
                                                                                    if true {
                                                                                        _ = len(v45)
                                                                                        if true {
                                                                                            var v46 uint8 = 255
                                                                                            _ = v46
                                                                                        } else {
                                                                                            if (v4 == (^(v4))) {
                                                                                                var v47 float64 = v2
                                                                                                delete(v6, "a")
                                                                                                _ = v47
                                                                                            } else {
                                                                                                if false {
                                                                                                    if (!(664 >= 2147483647)) {
                                                                                                        var v48 int8 = -96
                                                                                                        if (true || false) {
                                                                                                            v49 := GenS3{F0: 1}
                                                                                                            var v50 uint8 = 127
                                                                                                            v6[("a" + "foo bar")] = false
                                                                                                            _ = v49
                                                                                                            _ = v50
                                                                                                        }
                                                                                                        v42 -= 51
                                                                                                        fmt.Println(v48)
                                                                                                        _ = v48
                                                                                                    }
                                                                                                    if len(v45) >= 2 {
                                                                                                        v45 = v45[1:len(v45)]
                                                                                                    }
                                                                                                } else {
                                                                                                    v19 += v2
                                                                                                    var v51 uint8 = 127
                                                                                                    fmt.Println(v9, v42, v19)
                                                                                                    _ = v51
                                                                                                }
                                                                                            }
                                                                                            fmt.Println(v11, v40)
                                                                                        }
                                                                                    } else {
                                                                                        delete(v3, ((("a" + "a") + "\n\t") + "foo bar"))
                                                                                        var v52 uint = 151
                                                                                        _ = v52
                                                                                    }
                                                                                    v40 += (fi37 & v40)
                                                                                    _ = v45
                                                                                }
                                                                            }
                                                                        default:
                                                                            var v53 string = ("a" + "\n\t")
                                                                            _ = v53
                                                                        }
                                                                        v38 += -9223372036854775808
                                                                        switch "a" {
                                                                        case "hello":
                                                                            if true {
                                                                                var v54 bool = false
                                                                                var v55 uint8 = 42
                                                                                _ = v54
                                                                                _ = v55
                                                                            }
                                                                        case "foo bar":
                                                                            var v56 int = -2147483648
                                                                            _ = v56
                                                                        case "a":
                                                                            fmt.Println(v9)
                                                                        default:
                                                                            fmt.Println(v4, v11)
                                                                        }
                                                                    }
                                                                    v40 += (fi37 + (-(v38)))
                                                                    switch (fi37 << ((2) & 63)) {
                                                                    case 32767:
                                                                        var v57 int32 = 42
                                                                        _ = v57
                                                                    case 590:
                                                                        switch ("" + ("a" + "a")) {
                                                                        case "a":
                                                                            if (127 > 1) {
                                                                                if ((v2 / v2) == (float64((-2.25 - float32(0.0))) + v2)) {
                                                                                    v58 := GenS5{F0: 0, F1: v4}
                                                                                    var v59 uint64 = 172
                                                                                    var v60 uint = 42
                                                                                    _ = v58
                                                                                    _ = v59
                                                                                    _ = v60
                                                                                }
                                                                            }
                                                                        }
                                                                    case -839:
                                                                        v4 += (v4 - v4)
                                                                    default:
                                                                        var v61 int64 = (v4 + v4)
                                                                        _ = v61
                                                                    }
                                                                }
                                                                _ = v42
                                                            }
                                                            delete(v1, (v38 * v40))
                                                            switch v38 {
                                                            case -2147483648:
                                                                if (!true) {
                                                                    switch "foo bar" {
                                                                    case "hello":
                                                                        if true {
                                                                            v62 := v1[(v40 - 32767)]
                                                                            v11 = (-(v11))
                                                                            _ = v62
                                                                        }
                                                                    case "":
                                                                        if true {
                                                                            v63 := GenS4{F0: 280, F1: fi37}
                                                                            if false {
                                                                                var v64 float64 = (1e10 - v19)
                                                                                _ = v64
                                                                            } else {
                                                                                if ("foo bar" < "a") {
                                                                                    fmt.Println(fi37, v11)
                                                                                    v2 += v2
                                                                                }
                                                                                v19 = float64((-2.25 * float32(0.0)))
                                                                            }
                                                                            var v65 string = "a"
                                                                            _ = v63
                                                                            _ = v65
                                                                        } else {
                                                                            var v66 int8 = v11
                                                                            v2 += ((-2.25 / v19) * float64((1e10 + 1.5)))
                                                                            _ = v66
                                                                        }
                                                                    }
                                                                    v4 = v4
                                                                }
                                                            default:
                                                                var v67 bool = false
                                                                _ = v67
                                                            }
                                                            _ = v40
                                                        } else {
                                                            v4 -= (v4 | v4)
                                                        }
                                                        if ((2147483647 == -1) || true) {
                                                            var v68 bool = true
                                                            if v68 {
                                                                if (!v68) {
                                                                    v69 := GenS6{F0: (v4 ^ int64(v11)), F1: (v4 * (v4 | (v4 + (-(v4)))))}
                                                                    _ = v69
                                                                } else {
                                                                    v70 := []int{(fi37 ^ v38), fi12, (v9 * (fi12 ^ 9223372036854775807)), (^(v38))}
                                                                    _ = v70
                                                                }
                                                                fmt.Println(v4, v2, v68)
                                                            }
                                                            switch 169 {
                                                            case 32768:
                                                                var v71 uint8 = 255
                                                                _ = v71
                                                            case -9223372036854775808:
                                                                v72 := v5[v9]
                                                                _ = v72
                                                            case -733:
                                                                if v68 {
                                                                    fmt.Println(v2)
                                                                    switch "\n\t" {
                                                                    case "a":
                                                                        var v73 int32 = 61
                                                                        _ = v73
                                                                    case "\n\t":
                                                                        var v74 float32 = 0.0
                                                                        _ = v74
                                                                    case "foo bar":
                                                                        v4 = (v4 / (((v4 & (v4 << ((2) & 63)))) | 1))
                                                                    }
                                                                    var v75 uint = 49
                                                                    _ = v75
                                                                }
                                                            default:
                                                                if v68 {
                                                                    if (!false) {
                                                                        var v76 bool = (-1 > 2147483647)
                                                                        var v77 int16 = 379
                                                                        v6[("hello" + "")] = true
                                                                        _ = v76
                                                                        _ = v77
                                                                    } else {
                                                                        switch fi37 {
                                                                        case -1:
                                                                            v19 = (v2 - 0.1)
                                                                        case -983:
                                                                            v78 := GenS1{F0: fi12}
                                                                            _ = v78
                                                                        default:
                                                                            switch (fi37 / (((v38 * 32768)) | 1)) {
                                                                            case 257:
                                                                                var v79 uint = 255
                                                                                _ = v79
                                                                            case 32768:
                                                                                v2 -= ((((1e10 - 0.0) + v19) * v2) / v19)
                                                                            }
                                                                        }
                                                                    }
                                                                    v80 := GenS3{F0: 1000}
                                                                    v81 := GenS5{F0: 976, F1: 9223372036854775807}
                                                                    _ = v80
                                                                    _ = v81
                                                                } else {
                                                                    var v82 int32 = 42
                                                                    v82 = (^(v82))
                                                                    var v83 bool = v68
                                                                    var v84 float32 = (-2.25 * 1.5)
                                                                    _ = v82
                                                                    _ = v83
                                                                    _ = v84
                                                                }
                                                            }
                                                            _ = v68
                                                        }
                                                        _ = v38
                                                    }
                                                    _ = fi37
                                                }
                                            case "hello":
                                                for fi85 := 0; fi85 < 19; fi85++ {
                                                    v19 -= v2
                                                    v86 := v6[(("\n\t" + "foo bar") + "a")]
                                                    var v87 bool = ((^(v11)) > (v11 << ((0) & 63)))
                                                    var v88 int = (fi12 - 176)
                                                    v2 = (-2.25 + (v2 / float64(float32((v19 + (3.14159 - v19))))))
                                                    var v89 uint16 = 2
                                                    _ = fi85
                                                    _ = v86
                                                    _ = v87
                                                    _ = v88
                                                    _ = v89
                                                }
                                            default:
                                                switch "a" {
                                                case "a":
                                                    fmt.Println(v9, v11)
                                                }
                                            }
                                            _ = v19
                                        }
                                    case "hello":
                                        for fi90 := 0; fi90 < 8; fi90++ {
                                            if false {
                                                delete(v3, ("foo bar" + "\n\t"))
                                                if true {
                                                    v91 := GenS2{F0: 0, F1: -421}
                                                    fmt.Println(v4, v4)
                                                    v4 -= 0
                                                    _ = v91
                                                } else {
                                                    var v92 int16 = -1
                                                    _ = v92
                                                }
                                            } else {
                                                v4 += v4
                                            }
                                            _ = fi90
                                        }
                                    default:
                                        if true {
                                            v93 := GenS6{F0: (v4 + (v4 * 385)), F1: v4}
                                            v94 := []int{}
                                            _ = v93
                                            _ = v94
                                        } else {
                                            v11 = v11
                                        }
                                    }
                                }
                            case 4294967295:
                                switch (("foo bar" + "hello") + "a") {
                                case "hello":
                                    v95 := []int{(fi12 & -9223372036854775808)}
                                    _ = v95
                                default:
                                    v11 = v11
                                }
                            }
                            for fi96 := 0; fi96 < 7; fi96++ {
                                var v97 uint8 = 2
                                fmt.Println(fi12, fi12, v2)
                                _ = fi96
                                _ = v97
                            }
                            _ = v13
                        }
                    case 2147483647:
                        v98 := []int{v9}
                        _ = v98
                    case 4294967295:
                        if (false && false) {
                            var v99 float32 = ((0.0 * (((float32(v2) - 1.5) / 1.5) / (float32((float64(1e10) - (-2.25 + 1e10))) + 1.5))) * 1e10)
                            var v100 int32 = -597
                            fmt.Println(fi12, fi12, v9)
                            _ = v100
                            _ = v99
                        }
                    default:
                        for fi101 := 0; fi101 < 19; fi101++ {
                            _ = len(v3)
                            v9 = (fi12 << ((65535) & 63))
                            var v102 uint64 = 121
                            _ = fi101
                            _ = v102
                        }
                    }
                    var v103 int = v9
                    _ = v103
                } else {
                    for fi104 := 0; fi104 < 10; fi104++ {
                        v11 = genFn1((v9 & 160), 255, 42)
                        _ = fi104
                    }
                }
                v3[""] = ((("\n\t" + ("" + "\n\t")) + ("hello" + "")) + "\n\t")
                v105 := GenS3{F0: -1}
                _ = v105
            }
            var v106 float64 = v2
            var v107 uint16 = 2
            _ = fi12
            _ = v106
            _ = v107
        }
    }
    v2 = v2
    v2 -= float64(0.0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
