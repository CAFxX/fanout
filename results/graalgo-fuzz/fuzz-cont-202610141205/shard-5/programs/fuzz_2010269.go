package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 bool }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 uint64; F1 uint }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint16; F1 string; F2 int8 }
type GenS7 struct { F0 uint16; F1 int64 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 string; F1 uint8; F2 bool }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]int)
    var v3 bool = (false || ((3.14159 - 3.14159) > -2.25))
    v4 := make(map[string]uint32)
    fmt.Println(v3)
    switch "hello" {
    case "hello":
        var v5 uint64 = 1
        _ = v5
    case "foo bar":
        v3 = ((1 < 2) || (float32(-2.25) >= -2.25))
    default:
        var v6 int16 = 160
        _ = v6
    }
    v7 := GenS1{F0: -1, F1: 1, F2: (100 < 1)}
    var v8 bool = (1.5 < 1e10)
    var v9 string = ((("" + ("\n\t" + "hello")) + "\n\t") + "a")
    v9 = "hello"
    v10 := GenS2{F0: 1, F1: 421}
    v11 := GenS3{F0: 100, F1: 2}
    for fi12 := 0; fi12 < 11; fi12++ {
        var v13 uint16 = 0
        v8 = (0 >= 42)
        _ = fi12
        _ = v13
    }
    fmt.Println(v9, v8)
    switch -509 {
    case 32767:
        v9 = v9
    case -964:
        if v3 {
            var v14 bool = v8
            if v14 {
                _ = len(v1)
            } else {
                if (65535 > 2) {
                    v15 := GenS4{F0: 1, F1: 255}
                    for fi16 := 0; fi16 < 7; fi16++ {
                        fmt.Println(v14, v14)
                        _ = fi16
                    }
                    _ = v15
                }
            }
            _ = v14
        }
    }
    v17 := GenS2{F0: 1, F1: 0}
    switch 2147483647 {
    case 128:
        for fi18 := 0; fi18 < 17; fi18++ {
            v19 := GenS2{F0: -303, F1: 586}
            _ = fi18
            _ = v19
        }
    case -1:
        if (((!true) && (((v3 || true) || v3) || (false && (1 < 1)))) || ((v9 < v9) && v8)) {
            switch 4294967295 {
            case 128:
                if true {
                    v20 := []int{}
                    _ = v20
                } else {
                    fmt.Println(v8)
                    switch 342 {
                    case 128:
                        for fi21 := 0; fi21 < 20; fi21++ {
                            v3 = v3
                            for fi22 := 0; fi22 < 10; fi22++ {
                                v23 := GenS3{F0: 1, F1: 117}
                                if v8 {
                                    var v24 uint8 = 102
                                    _ = v24
                                } else {
                                    v25 := []int{}
                                    var v26 int64 = -9223372036854775808
                                    if v8 {
                                        switch v9 {
                                        case "a":
                                            var v27 int64 = 9223372036854775807
                                            var v28 int32 = -2147483648
                                            v28 = (v28 >> ((4294967295) & 63))
                                            var v29 uint64 = 42
                                            delete(v4, "foo bar")
                                            fmt.Println(fi22)
                                            var v30 uint16 = 32767
                                            _ = v27
                                            _ = v28
                                            _ = v29
                                            _ = v30
                                        }
                                        v4[(v9 + ("" + v9))] = 99
                                        switch 127 {
                                        case -9223372036854775808:
                                            v31 := GenS3{F0: 1, F1: 169}
                                            _ = v31
                                        default:
                                            if v8 {
                                                fmt.Println(v9, v26, v8)
                                                var v32 uint8 = 42
                                                if len(v25) >= 2 {
                                                    v25 = v25[1:len(v25)]
                                                }
                                                var v33 int16 = 866
                                                fmt.Println(v33, v8, fi22)
                                                _ = v32
                                                _ = v33
                                            }
                                        }
                                        if v3 {
                                            var v34 int32 = 0
                                            var v35 uint8 = 128
                                            _ = v34
                                            _ = v35
                                        }
                                        fmt.Println(v3, v9)
                                        v9 = (("a" + "a") + ((v9 + v9) + v9))
                                    } else {
                                        var v36 uint64 = 98
                                        _ = v36
                                    }
                                    var v37 float64 = ((0.1 / 1.0) - float64(0.0))
                                    if (150 == 541) {
                                        v38 := v1[65535]
                                        var v39 uint = 4294967295
                                        _ = v38
                                        _ = v39
                                    }
                                    var v40 float64 = v37
                                    _ = v25
                                    _ = v26
                                    _ = v37
                                    _ = v40
                                }
                                _ = fi22
                                _ = v23
                            }
                            _ = fi21
                        }
                    }
                }
            case 514:
                v1[2147483647] = 2
            case 32767:
                v41 := GenS1{F0: -56, F1: 100, F2: (220 == 32767)}
                _ = v41
            }
            v3 = true
            delete(v2, v9)
        } else {
            switch 536 {
            case 0:
                var v42 int8 = 1
                _ = v42
            case 806:
                v9 = "a"
            case 42:
                v3 = v3
            }
        }
    case 245:
        if (v8 && (!true)) {
            switch 1 {
            case 484:
                for fi43 := 0; fi43 < 3; fi43++ {
                    switch "hello" {
                    case "foo bar":
                        for fi44 := 0; fi44 < 3; fi44++ {
                            if (!(v3 && (false || true))) {
                                v45 := []int{fi43, (fi43 + (fi44 ^ (fi43 * 128)))}
                                fmt.Println(fi44)
                                v9 = ("hello" + ((v9 + ((v9 + "\n\t") + ((v9 + "a") + (((v9 + v9) + v9) + v9)))) + (("hello" + ("a" + (v9 + v9))) + ((("" + v9) + v9) + (v9 + v9)))))
                                if len(v45) >= 2 {
                                    v45 = v45[1:len(v45)]
                                }
                                _ = v45
                            } else {
                                if (-128 <= 127) {
                                    var v46 int = 435
                                    _ = v46
                                } else {
                                    var v47 uint8 = 105
                                    var v48 int = fi44
                                    if v8 {
                                        var v49 int32 = -912
                                        fmt.Println(v47, v3, v8)
                                        v49 = (v49 ^ -2147483648)
                                        _ = v49
                                    }
                                    _ = v47
                                    _ = v48
                                }
                                v8 = v8
                                if v8 {
                                    switch (fi43 - fi43) {
                                    case 1:
                                        if v3 {
                                            v9 = ""
                                            fmt.Println(v8, fi43, v3)
                                            var v50 int = (fi43 & -138)
                                            _ = v50
                                        }
                                    case 4294967295:
                                        v51 := GenS5{F0: 26}
                                        v52 := GenS6{F0: 2, F1: ("" + v9), F2: -62}
                                        fmt.Println(v9, fi43)
                                        _ = v51
                                        _ = v52
                                    case 0:
                                        v8 = false
                                    }
                                    var v53 int16 = 556
                                    _ = v53
                                }
                            }
                            _ = fi44
                        }
                    case "a":
                        v54 := GenS4{F0: 1, F1: 0}
                        _ = v54
                    case "hello":
                        if (2 >= 18446744073709551615) {
                            if v3 {
                                v55 := v4[v9]
                                v56 := []int{}
                                v3 = (false || (121 == 1))
                                _ = v55
                                _ = v56
                            } else {
                                fmt.Println(v3)
                                if v3 {
                                    v8 = v3
                                }
                            }
                        }
                    }
                    _ = fi43
                }
            case -1:
                if ((32767 != -1) && v8) {
                    switch 9223372036854775807 {
                    case 32768:
                        if ((v9 + "hello") == "") {
                            fmt.Println(v9)
                            v57 := GenS5{F0: 127}
                            var v58 uint16 = 255
                            _ = v57
                            _ = v58
                        } else {
                            v8 = (((!true) && v8) || v8)
                            v9 = ((((v9 + (v9 + (("a" + ("" + "a")) + v9))) + v9) + (v9 + (v9 + "a"))) + (("hello" + (v9 + v9)) + (v9 + v9)))
                        }
                    case 127:
                        v59 := GenS7{F0: 42, F1: 9223372036854775807}
                        _ = v59
                    }
                    v60 := GenS1{F0: 2, F1: -99, F2: false}
                    if false {
                        v61 := GenS6{F0: 207, F1: "foo bar", F2: 0}
                        var v62 uint16 = 1
                        var v63 int64 = 1
                        _ = v61
                        _ = v62
                        _ = v63
                    }
                    _ = v60
                } else {
                    fmt.Println(v9, v8, v9)
                    switch "" {
                    case "foo bar":
                        for fi64 := 0; fi64 < 5; fi64++ {
                            for fi65 := 0; fi65 < 12; fi65++ {
                                if (193 != 227) {
                                    if v8 {
                                        var v66 float32 = (3.14159 / 1.0)
                                        delete(v2, v9)
                                        _ = v66
                                    } else {
                                        v67 := []int{-1, fi64}
                                        if false {
                                            v68 := GenS5{F0: 109}
                                            _ = v68
                                        }
                                        _ = v67
                                    }
                                    v69 := GenS8{F0: 42}
                                    switch fi65 {
                                    case 32768:
                                        v70 := GenS5{F0: 128}
                                        _ = v70
                                    case 890:
                                        v1[2] = 2
                                    case 127:
                                        if v8 {
                                            v71 := GenS8{F0: 350}
                                            _ = v71
                                        }
                                    }
                                    _ = v69
                                }
                                v72 := []int{}
                                _ = fi65
                                _ = v72
                            }
                            fmt.Println(fi64)
                            _ = fi64
                        }
                    case "a":
                        for fi73 := 0; fi73 < 9; fi73++ {
                            v74 := GenS9{F0: (("\n\t" + v9) + ("foo bar" + v9)), F1: 42, F2: v8}
                            _ = fi73
                            _ = v74
                        }
                    }
                }
            default:
                if true {
                    var v75 uint32 = 54
                    _ = v75
                }
            }
        } else {
            var v76 uint64 = 965
            v76 -= v76
            _ = v76
        }
    }
    if v3 {
        var v77 bool = false
        v3 = (-128 <= 127)
        v1[865] = 0
        _ = v77
    } else {
        var v78 uint = 42
        v79 := []int{-212, -452}
        var v80 uint64 = 42
        _ = v78
        _ = v79
        _ = v80
    }
    fmt.Println(v3, v8, v8)
    v9 = v9
    v9 = v9
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
