package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 uint; F1 string }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 string; F1 int; F2 int32 }
type GenS6 struct { F0 int64; F1 uint8; F2 bool }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int16; F1 int32; F2 int }
type GenS9 struct { F0 uint64; F1 uint; F2 string }
type GenS10 struct { F0 uint64 }
func genFn0(p0 uint16, p1 uint64) bool {
    for fi4 := 0; fi4 < 10; fi4++ {
        var v5 int = 74
        var v6 uint16 = 65535
        _ = fi4
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: -1}
    p0 -= 603
    _ = p0
    _ = p1
    _ = v7
    return (!(-1 >= -9223372036854775808))
}
func genFn1() uint {
    if false {
        var v8 uint32 = 0
        for fi9 := 0; fi9 < 1; fi9++ {
            var v10 int16 = -356
            _ = fi9
            _ = v10
        }
        if false {
            v8 = (v8 / (((^(v8))) | 1))
            var v11 uint64 = 1
            v12 := GenS1{F0: 100}
            _ = v11
            _ = v12
        }
        _ = v8
    }
    v13 := GenS1{F0: 42}
    if (false || true) {
        v14 := GenS2{F0: (("" + "foo bar") + "\n\t"), F1: 2}
        for fi15 := 0; fi15 < 8; fi15++ {
            for fi16 := 0; fi16 < 7; fi16++ {
                switch (fi16 >> ((255) & 63)) {
                case 9223372036854775807:
                    var v17 uint32 = 210
                    _ = v17
                default:
                    v18 := []int{(^(fi15)), (fi15 | fi16)}
                    _ = v18
                }
                _ = fi16
            }
            if (!(false || false)) {
                var v19 uint16 = 32767
                if ((v19 & v19) >= (v19 + v19)) {
                    var v20 int = -709
                    for fi21 := 0; fi21 < 8; fi21++ {
                        if (!(42 < 1)) {
                            var v22 string = "hello"
                            v23 := []int{fi21, (^(fi15)), int(v19), -1}
                            switch "" {
                            case "a":
                                v24 := GenS3{F0: 4294967295, F1: ("" + "\n\t")}
                                _ = v24
                            case "foo bar":
                                if true {
                                    var v25 uint8 = 36
                                    var v26 int64 = 1
                                    v27 := GenS4{F0: 549}
                                    _ = v25
                                    _ = v26
                                    _ = v27
                                } else {
                                    var v28 float64 = 0.0
                                    v29 := GenS5{F0: v22, F1: 750, F2: 0}
                                    var v30 int = (v20 - 127)
                                    _ = v28
                                    _ = v29
                                    _ = v30
                                }
                            default:
                                fmt.Println(v19, fi21, v20)
                            }
                            _ = v22
                            _ = v23
                        } else {
                            v31 := []int{42, fi15, fi15, fi15}
                            _ = v31
                        }
                        var v32 uint8 = 1
                        _ = fi21
                        _ = v32
                    }
                    _ = v20
                } else {
                    v33 := GenS6{F0: 42, F1: 157, F2: (255 == 184)}
                    _ = v33
                }
                _ = v19
            } else {
                v34 := []int{}
                var v35 bool = false
                switch "foo bar" {
                case "\n\t":
                    if (-1 != 196) {
                        if len(v34) > 0 {
                            v34[0] = fi15
                        }
                    }
                case "foo bar":
                    switch fi15 {
                    case 2147483647:
                        var v36 int32 = 526
                        _ = v36
                    }
                case "":
                    var v37 float64 = 0.1
                    _ = v37
                }
                _ = v34
                _ = v35
            }
            v38 := []int{}
            _ = fi15
            _ = v38
        }
        _ = v14
    } else {
        var v39 int = -518
        v40 := []int{v39, (v39 * (v39 % ((v39) | 1)))}
        v39 += (v39 + 859)
        _ = v39
        _ = v40
    }
    var v41 uint64 = 108
    _ = v13
    _ = v41
    return 2
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]int)
    var v3 int8 = 0
    for fi42 := 0; fi42 < 3; fi42++ {
        switch (-(fi42)) {
        case 32767:
            v3 = (-(v3))
        case -2147483648:
            var v43 float64 = float64(1.5)
            _ = v43
        }
        _ = fi42
    }
    switch "a" {
    case "":
        v44 := []int{251, 42}
        _ = v44
    case "\n\t":
        switch 1 {
        case -63:
            var v45 string = "foo bar"
            _ = v45
        case 42:
            v46 := []int{-9223372036854775808, 32767}
            _ = v46
        case 1:
            var v47 int16 = -32768
            _ = v47
        }
    default:
        for fi48 := 0; fi48 < 19; fi48++ {
            for fi49 := 0; fi49 < 20; fi49++ {
                var v50 float32 = ((0.0 + 0.0) - 3.14159)
                if ((1.5 - 3.14159) != float64(-2.25)) {
                    v50 = (0.0 / 1e10)
                }
                v50 = (float32((3.14159 * (1e10 / float64(v50)))) - float32(float64(((float32(0.0) * 1.5) / v50))))
                _ = fi49
                _ = v50
            }
            fmt.Println(v3)
            _ = fi48
        }
    }
    var v51 int = 127
    var v52 uint32 = 2147483648
    v51 -= (v51 >> ((4294967295) & 63))
    switch v51 {
    case -837:
        if true {
            v52 = (v52 + v52)
            if true {
                for fi53 := 0; fi53 < 19; fi53++ {
                    var v54 int16 = -32768
                    for fi55 := 0; fi55 < 10; fi55++ {
                        v1[v51] = (2 == 0)
                        v52 = 1
                        _ = fi55
                    }
                    var v56 int = (v51 % ((fi53) | 1))
                    _ = fi53
                    _ = v54
                    _ = v56
                }
                v2[(v52 & 909)] = (v51 % ((128) | 1))
                var v57 int = (v51 - 4294967295)
                _ = v57
            }
        }
    default:
        v58 := []int{2147483647}
        _ = v58
    }
    if true {
        fmt.Println(v51, v3, v52)
        var v59 uint = 4294967295
        v52 += v52
        _ = v59
    }
    v60 := GenS4{F0: v52}
    switch (("foo bar" + "") + "hello") {
    case "foo bar":
        var v61 int8 = -1
        _ = v61
    }
    if false {
        v62 := GenS7{F0: genFn1()}
        for fi63 := 0; fi63 < 7; fi63++ {
            if ((3.14159 / (0.1 + 1.5)) <= (1e10 / 3.14159)) {
                v64 := GenS8{F0: 0, F1: 0, F2: v51}
                var v65 uint8 = 242
                v51 -= v51
                _ = v64
                _ = v65
            }
            v51 = int(v52)
            _ = fi63
        }
        if (!false) {
            for fi66 := 0; fi66 < 3; fi66++ {
                v52 = 989
                _ = fi66
            }
        } else {
            v67 := GenS4{F0: (v52 ^ v52)}
            _ = v67
        }
        _ = v62
    }
    for fi68 := 0; fi68 < 1; fi68++ {
        if true {
            for fi69 := 0; fi69 < 10; fi69++ {
                var v70 int32 = -772
                fmt.Println(fi68, v52, v3)
                _ = fi69
                _ = v70
            }
            if true {
                v3 -= v3
                fmt.Println(v51)
                v71 := []int{(v51 | v51), (fi68 >> ((0) & 63)), (-(v51)), (fi68 / ((fi68) | 1))}
                _ = v71
            }
        } else {
            var v72 int8 = v3
            v52 = (v52 - (v52 % (((^(v52))) | 1)))
            _ = v72
        }
        v73 := GenS9{F0: 1, F1: 42, F2: "foo bar"}
        switch fi68 {
        case 1:
            if false {
                fmt.Println(v51, fi68)
                v74 := GenS2{F0: ("\n\t" + (("\n\t" + "hello") + ("hello" + "a"))), F1: 2}
                _ = v74
            }
        case 2147483647:
            v75 := GenS10{F0: 632}
            _ = v75
        case -422:
            fmt.Println(v52)
        default:
            if false {
                var v76 int8 = (v3 * (v3 << ((42) & 63)))
                var v77 uint32 = (^(v52))
                _ = v76
                _ = v77
            }
        }
        _ = fi68
        _ = v73
    }
    for fi78 := 0; fi78 < 16; fi78++ {
        var v79 uint16 = 841
        _ = fi78
        _ = v79
    }
    delete(v1, v51)
    for fi80 := 0; fi80 < 16; fi80++ {
        v52 -= v52
        _ = fi80
    }
    var v81 uint32 = v52
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v51
    _ = v52
    _ = v60
    _ = v81
}
