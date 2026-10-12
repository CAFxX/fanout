package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int16; F1 int16; F2 uint64 }
type GenS5 struct { F0 uint32; F1 uint64 }
type GenS6 struct { F0 uint; F1 uint32; F2 uint32 }
type GenS7 struct { F0 int8 }
func genFn0() bool {
    v4 := GenS1{F0: 152}
    v5 := []int{-18, 32767, 32767, 9223372036854775807}
    var v6 uint64 = 42
    _ = v4
    _ = v5
    _ = v6
    return false
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint = 255
    var v3 string = "\n\t"
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 uint = v2
        var v9 float32 = 1e10
        switch fi7 {
        case -2147483648:
            if true {
                var v10 bool = (false && true)
                _ = v10
            }
        case -9223372036854775808:
            v11 := GenS2{F0: (v3 + ((v3 + v3) + v3)), F1: 127}
            _ = v11
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi12 := 0; fi12 < 18; fi12++ {
        v13 := GenS3{F0: "hello"}
        v14 := []int{(fi12 >> (((v2 >> (((v2 >> ((v2) & 63))) & 63))) & 63)), fi12}
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 string = v3
    if ("" != "") {
        if (!false) {
            for fi16 := 0; fi16 < 13; fi16++ {
                v17 := v1[(v15 + v3)]
                switch v3 {
                case "hello":
                    v18 := GenS3{F0: "foo bar"}
                    _ = v18
                case "foo bar":
                    v1["a"] = false
                }
                _ = fi16
                _ = v17
            }
            var v19 float32 = ((0.0 + (0.0 / ((float32(-2.25) * float32((1.5 / (float64(-2.25) / (((3.14159 / 1.5) * (1e10 - 1e10)) - (3.14159 * 3.14159)))))) / 1e10))) / 1.5)
            _ = v19
        } else {
            if (3.14159 < ((3.14159 / 1e10) / 1.5)) {
                for fi20 := 0; fi20 < 18; fi20++ {
                    _ = len(v1)
                    v15 = v3
                    v2 -= 1
                    _ = fi20
                }
                switch ("\n\t" + "") {
                case "":
                    var v21 int8 = -82
                    _ = v21
                default:
                    switch ((((v15 + ("\n\t" + ("" + v15))) + v15) + v15) + v15) {
                    case "a":
                        if true {
                            v3 = v3
                            v1["\n\t"] = true
                        } else {
                            v2 -= (v2 & 65535)
                            switch v3 {
                            case "hello":
                                if true {
                                    switch "hello" {
                                    case "":
                                        for fi22 := 0; fi22 < 5; fi22++ {
                                            v2 = 0
                                            v23 := GenS2{F0: (v3 + v3), F1: -118}
                                            _ = fi22
                                            _ = v23
                                        }
                                    }
                                } else {
                                    var v24 bool = (false || true)
                                    _ = v24
                                }
                            case "a":
                                v3 = v3
                            case "":
                                var v25 int64 = 869
                                _ = v25
                            default:
                                v15 = (v3 + v15)
                            }
                            if (2147483647 == -1) {
                                v26 := GenS3{F0: ((v15 + "foo bar") + ("a" + (((((v3 + "hello") + v15) + (v15 + "a")) + "\n\t") + (v3 + v3))))}
                                v27 := []int{-244, 352}
                                _ = v26
                                _ = v27
                            } else {
                                v3 = ""
                            }
                        }
                    case "foo bar":
                        v2 = 42
                    case "hello":
                        var v28 uint16 = 615
                        _ = v28
                    default:
                        _ = len(v1)
                    }
                }
            } else {
                v15 = v15
                var v29 uint = v2
                _ = v29
            }
        }
    }
    switch 128 {
    case 9223372036854775807:
        var v30 float32 = 0.0
        _ = v30
    case 2147483647:
        if (42 != 42) {
            v31 := GenS2{F0: v15, F1: -88}
            _ = v31
        }
    case 748:
        v32 := v1[v3]
        _ = v32
    }
    if false {
        v2 = (v2 - (v2 & v2))
        switch -9223372036854775808 {
        case 0:
            fmt.Println(v2)
        case 32767:
            var v33 int16 = -32768
            _ = v33
        case 4294967295:
            fmt.Println(v15, v3)
        default:
            fmt.Println(v3)
        }
        v3 = v15
    } else {
        v15 = v15
    }
    var v34 uint32 = 2147483647
    v1[(("foo bar" + (v3 + v3)) + "")] = true
    var v35 uint8 = 127
    for fi36 := 0; fi36 < 14; fi36++ {
        v37 := GenS1{F0: 42}
        v34 += v34
        _ = fi36
        _ = v37
    }
    v38 := []int{}
    if (true && true) {
        v2 = (^(v2))
        var v39 uint16 = 311
        var v40 float64 = (3.14159 + 1.5)
        _ = v39
        _ = v40
    }
    switch 2147483647 {
    case 0:
        if true {
            for fi41 := 0; fi41 < 6; fi41++ {
                v3 = (("\n\t" + (v3 + (v3 + (v15 + ((v15 + "foo bar") + v15))))) + ("hello" + v3))
                if len(v38) > 0 {
                    v38[0] = fi41
                }
                _ = fi41
            }
        } else {
            for fi42 := 0; fi42 < 17; fi42++ {
                if false {
                    var v43 uint16 = 42
                    if true {
                        var v44 int64 = 268
                        _ = v44
                    }
                    var v45 int32 = 0
                    _ = v43
                    _ = v45
                } else {
                    v2 += 65535
                    fmt.Println(v34)
                }
                v3 = "foo bar"
                _ = fi42
            }
            v38 = append(v38, 42)
            var v46 string = (v3 + (v3 + "hello"))
            _ = v46
        }
    case 127:
        v47 := v1[v15]
        _ = v47
    case -667:
        var v48 int32 = 2147483647
        _ = v48
    }
    for fi49 := 0; fi49 < 8; fi49++ {
        if (true || true) {
            v34 -= (^(v34))
        }
        fmt.Println(fi49, v34, fi49)
        switch -772 {
        case 32767:
            for fi50 := 0; fi50 < 12; fi50++ {
                v2 = (v2 | v2)
                if false {
                    v1[((v3 + "") + ("a" + ""))] = (1 <= 4294967295)
                    switch (fi49 & fi50) {
                    case 1:
                        var v51 uint16 = 255
                        _ = v51
                    case -884:
                        v3 = ("\n\t" + v15)
                        v52 := GenS1{F0: -1}
                        var v53 bool = (true || false)
                        v53 = ("foo bar" == (v3 + ("" + v15)))
                        v3 = v15
                        _ = v52
                        _ = v53
                    default:
                        v54 := GenS4{F0: 32767, F1: 931, F2: uint64(v2)}
                        _ = v54
                    }
                    var v55 uint16 = 32767
                    _ = v55
                } else {
                    var v56 float32 = -2.25
                    fmt.Println(v2)
                    switch (^(fi50)) {
                    case 2147483647:
                        v57 := GenS3{F0: ""}
                        _ = v57
                    case -622:
                        v1["a"] = false
                        if true {
                            if (365 <= 42) {
                                switch (fi49 | fi49) {
                                case -612:
                                    v58 := GenS2{F0: v15, F1: -128}
                                    _ = v58
                                case 32767:
                                    v59 := GenS3{F0: v15}
                                    _ = v59
                                case 944:
                                    v15 = ("" + "\n\t")
                                }
                            } else {
                                v60 := v1["a"]
                                switch v15 {
                                case "foo bar":
                                    fmt.Println(v15)
                                default:
                                    fmt.Println(v35, v60)
                                }
                                v35 = (v35 * (^(v35)))
                                _ = v60
                            }
                            var v61 float32 = v56
                            var v62 uint64 = 4294967295
                            _ = v61
                            _ = v62
                        } else {
                            v63 := GenS5{F0: v34, F1: 4294967295}
                            v38 = append(v38, int(v34))
                            var v64 uint32 = v34
                            v65 := GenS4{F0: 32767, F1: -222, F2: 4294967295}
                            v56 = 0.1
                            var v66 uint32 = 2147483648
                            _ = v63
                            _ = v64
                            _ = v65
                            _ = v66
                        }
                        if false {
                            if true {
                                v35 = (v35 + 2)
                                v34 = v34
                                switch (fi49 | (-(fi49))) {
                                case 32767:
                                    var v67 bool = (true || true)
                                    _ = v67
                                case 42:
                                    v68 := GenS6{F0: (v2 % (((v2 / (((v2 - uint(v34))) | 1))) | 1)), F1: v34, F2: (v34 | (v34 | (v34 & (v34 & v34))))}
                                    _ = v68
                                }
                                var v69 uint8 = 38
                                _ = v69
                            }
                            if len(v38) >= 2 {
                                v38 = v38[1:len(v38)]
                            }
                        } else {
                            fmt.Println(v2, v35)
                        }
                    case 4294967295:
                        if (42 <= 0) {
                            var v70 uint32 = v34
                            v71 := GenS6{F0: (v2 - 750), F1: (v34 - uint32(v2)), F2: v34}
                            _ = v70
                            _ = v71
                        }
                    }
                    _ = v56
                }
                _ = fi50
            }
        case 65:
            if (false && true) {
                v34 = (v34 << ((0) & 63))
            }
        }
        _ = fi49
    }
    var v72 bool = (2 < v35)
    for fi73 := 0; fi73 < 2; fi73++ {
        if v72 {
            _ = len(v38)
            v74 := GenS2{F0: ("" + (v15 + (((v15 + v3) + (v15 + v15)) + "foo bar"))), F1: 1}
            var v75 int64 = 1
            _ = v74
            _ = v75
        }
        _ = fi73
    }
    var v76 int16 = -1
    if (v72 && ((^(v35)) > (v35 + (^(v35))))) {
        var v77 uint32 = (v34 - (v34 / (((v34 * (v34 ^ (^(v34))))) | 1)))
        _ = v77
    } else {
        v78 := GenS7{F0: 126}
        _ = v78
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v72:", v72)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v38
    _ = v72
    _ = v76
}
