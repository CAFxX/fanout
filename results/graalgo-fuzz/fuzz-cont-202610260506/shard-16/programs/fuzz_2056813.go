package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 uint8; F2 uint64 }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 uint16; F1 int; F2 int64 }
type GenS5 struct { F0 int64 }

func main() {
    v1 := make(map[string]string)
    v2 := []uint32{}
    var v3 float32 = ((0.1 - 3.14159) - 0.1)
    var v4 uint16 = 2
    v5 := GenS1{F0: uint32(v4)}
    var v6 uint32 = uint32(v4)
    if false {
        v6 += v6
        var v7 int8 = -15
        _ = v7
    } else {
        v8 := GenS2{F0: (v4 * v4), F1: 42, F2: 2}
        _ = v8
    }
    fmt.Println(v3, v3, v4)
    switch 2147483647 {
    case 2147483647:
        var v9 float64 = (1e10 * 0.1)
        _ = v9
    case -1:
        var v10 int32 = 42
        _ = v10
    }
    switch ("foo bar" + ("" + "hello")) {
    case "a":
        v11 := GenS2{F0: v4, F1: 221, F2: 981}
        _ = v11
    case "foo bar":
        var v12 bool = ((v4 >= 42) || true)
        _ = v12
    }
    if (!true) {
        if ((9223372036854775807 != -497) || (1 >= 1000)) {
            if (65535 == v6) {
                if (true && (0 >= 1000)) {
                    v13 := GenS1{F0: 628}
                    v6 += v6
                    _ = v13
                } else {
                    v14 := GenS2{F0: v4, F1: 136, F2: 2}
                    _ = v14
                }
                for fi15 := 0; fi15 < 5; fi15++ {
                    var v16 uint16 = v4
                    v17 := GenS2{F0: v4, F1: 42, F2: 4294967295}
                    if true {
                        fmt.Println(v3)
                        for fi18 := 0; fi18 < 7; fi18++ {
                            v19 := GenS2{F0: v4, F1: 131, F2: 944}
                            var v20 int = 32768
                            if (-1 > 84) {
                                if (!(!(107 >= 2))) {
                                    var v21 uint8 = 119
                                    v3 -= float32((1e10 - -2.25))
                                    var v22 string = "hello"
                                    if (!(!true)) {
                                        var v23 bool = (238 > (v21 + (v21 >> ((2) & 63))))
                                        v20 -= 1
                                        _ = v23
                                    }
                                    var v24 bool = (v22 >= v22)
                                    _ = v21
                                    _ = v22
                                    _ = v24
                                } else {
                                    var v25 float64 = 1e10
                                    var v26 uint32 = v6
                                    switch -517 {
                                    case 0:
                                        if (157 != -9223372036854775808) {
                                            if false {
                                                var v27 int64 = 42
                                                v28 := GenS3{F0: v27, F1: (-1 == 42)}
                                                if (!(((^(v6)) != 50) && (("a" + "foo bar") < ("" + "")))) {
                                                    v4 -= v16
                                                    v4 = 32767
                                                }
                                                _ = v27
                                                _ = v28
                                            } else {
                                                delete(v1, (("hello" + "\n\t") + ("\n\t" + ((("a" + ("\n\t" + "a")) + "") + ("a" + "\n\t")))))
                                                var v29 int16 = 0
                                                v30 := GenS1{F0: 850}
                                                if len(v2) > 0 {
                                                    v2[0] = (^(v26))
                                                }
                                                switch "hello" {
                                                case "foo bar":
                                                    v31 := v1["foo bar"]
                                                    _ = v31
                                                case "a":
                                                    var v32 float64 = float64(v3)
                                                    v1[("foo bar" + "\n\t")] = "hello"
                                                    _ = v32
                                                }
                                                _ = v29
                                                _ = v30
                                            }
                                            v33 := GenS3{F0: -9223372036854775808, F1: ((v3 == 0.0) || false)}
                                            fmt.Println(fi15)
                                            _ = v33
                                        }
                                    case 501:
                                        var v34 uint64 = 152
                                        v26 += (^(v6))
                                        _ = v34
                                    case 32768:
                                        v6 -= (v26 & (v6 & (v26 | (v6 * 42))))
                                    default:
                                        var v35 int16 = -1
                                        _ = v35
                                    }
                                    _ = v25
                                    _ = v26
                                }
                                var v36 string = "foo bar"
                                switch 9223372036854775807 {
                                case 4294967295:
                                    switch (-(fi15)) {
                                    case 450:
                                        v36 = v36
                                    case -9223372036854775808:
                                        v20 = (v20 * (v20 >> ((42) & 63)))
                                    case -1:
                                        var v37 uint64 = 4294967295
                                        _ = v37
                                    default:
                                        var v38 int64 = -1
                                        _ = v38
                                    }
                                case -2147483648:
                                    var v39 uint64 = 601
                                    _ = v39
                                case 9223372036854775807:
                                    delete(v1, "")
                                }
                                _ = v36
                            }
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                    }
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                if (42 > 281) {
                    var v40 uint8 = 42
                    var v41 float32 = v3
                    var v42 int32 = 42
                    _ = v40
                    _ = v41
                    _ = v42
                } else {
                    for fi43 := 0; fi43 < 20; fi43++ {
                        v3 = (v3 - -2.25)
                        _ = fi43
                    }
                    v44 := GenS1{F0: v6}
                    _ = v44
                }
            }
            fmt.Println(v6, v4, v4)
        }
        var v45 int16 = -1
        for fi46 := 0; fi46 < 6; fi46++ {
            v4 = v4
            v45 += 32767
            for fi47 := 0; fi47 < 17; fi47++ {
                if ((!true) && (!(float64(float32(((((1.5 - 1.5) + (1.5 - float64(v3))) / -2.25) + ((-2.25 + 0.0) * ((float64(((v3 / v3) / (1e10 - ((0.0 + (1.5 + v3)) * (v3 / v3))))) + (3.14159 - 0.0)) + (((0.1 * -2.25) * ((0.0 - 3.14159) / ((0.1 * 3.14159) * (3.14159 * 3.14159)))) - 1.5)))))) <= 0.1))) {
                    v3 -= 0.0
                } else {
                    v48 := GenS4{F0: 42, F1: (fi47 | fi47), F2: 0}
                    v2 = append(v2, 633)
                    fmt.Println(v6)
                    _ = v48
                }
                if true {
                    v6 += v6
                    var v49 int64 = 9223372036854775807
                    var v50 int = (fi46 << ((369) & 63))
                    _ = v49
                    _ = v50
                } else {
                    _ = len(v1)
                    if (-1 == -493) {
                        v51 := GenS3{F0: -550, F1: (0 != 16)}
                        v4 = v4
                        var v52 uint16 = (v4 >> ((65535) & 63))
                        switch (fi46 ^ (fi47 * 4294967295)) {
                        case -2147483648:
                            var v53 int64 = -1
                            _ = v53
                        case 32:
                            var v54 float64 = float64(float32(float64(v3)))
                            _ = v54
                        }
                        _ = v51
                        _ = v52
                    }
                    if (v3 > v3) {
                        v55 := GenS4{F0: 65535, F1: 127, F2: -9223372036854775808}
                        var v56 uint32 = v6
                        _ = v55
                        _ = v56
                    }
                    if false {
                        v6 += 2147483647
                        var v57 int16 = v45
                        _ = v57
                    } else {
                        v6 -= (v6 / (((v6 | v6)) | 1))
                        v3 -= 1.5
                        switch fi47 {
                        case -2147483648:
                            if (!false) {
                                switch (fi46 & (fi46 - fi47)) {
                                case 237:
                                    fmt.Println(fi46, v4)
                                default:
                                    v58 := v1[("\n\t" + "\n\t")]
                                    _ = v58
                                }
                                var v59 float32 = 1.5
                                v60 := v1["hello"]
                                _ = v59
                                _ = v60
                            } else {
                                v3 = v3
                                if false {
                                    v61 := GenS3{F0: 1, F1: false}
                                    _ = len(v2)
                                    _ = v61
                                } else {
                                    v4 -= 2
                                    var v62 int8 = -107
                                    switch -39 {
                                    case 127:
                                        var v63 int32 = 860
                                        var v64 bool = true
                                        _ = v63
                                        _ = v64
                                    }
                                    v65 := v1["hello"]
                                    _ = v62
                                    _ = v65
                                }
                                if (false && (true && true)) {
                                    if (("hello" + "hello") != "") {
                                        var v66 int16 = (v45 % ((-1) | 1))
                                        var v67 float32 = float32(1e10)
                                        fmt.Println(v45, v45)
                                        _ = v66
                                        _ = v67
                                    }
                                    var v68 uint16 = 2
                                    var v69 float64 = (1.5 + 1e10)
                                    _ = v68
                                    _ = v69
                                }
                            }
                        case 13:
                            v6 = v6
                            fmt.Println(fi47, fi47)
                            v1[("" + "")] = "foo bar"
                            if (-2147483648 > 805) {
                                switch ((("\n\t" + "\n\t") + "a") + "") {
                                case "\n\t":
                                    v3 -= v3
                                case "a":
                                    v45 -= v45
                                default:
                                    var v70 int16 = v45
                                    _ = v70
                                }
                            }
                            fmt.Println(v4, v6)
                            if false {
                                v4 = v4
                            } else {
                                var v71 string = "foo bar"
                                fmt.Println(v4)
                                v72 := v1[v71]
                                _ = v71
                                _ = v72
                            }
                            if false {
                                var v73 uint8 = 2
                                v74 := GenS1{F0: uint32(v73)}
                                _ = v73
                                _ = v74
                            } else {
                                v4 = (v4 << ((18446744073709551615) & 63))
                                var v75 uint64 = 633
                                _ = v75
                            }
                            v45 += 635
                        case -9223372036854775808:
                            var v76 uint = 42
                            _ = v76
                        }
                    }
                    v3 -= ((0.0 + v3) + (1e10 / v3))
                }
                var v77 uint8 = 128
                _ = fi47
                _ = v77
            }
            _ = fi46
        }
        _ = v45
    }
    switch "foo bar" {
    case "a":
        fmt.Println(v6, v4, v6)
    case "":
        v78 := v1[""]
        _ = v78
    default:
        v3 += (float32(0.1) + 0.0)
    }
    v79 := GenS1{F0: v6}
    fmt.Println(v6, v6)
    if (((("foo bar" + ((((("" + "hello") + "a") + ("" + "hello")) + (("a" + (((("" + "a") + (("a" + "foo bar") + "a")) + "hello") + "foo bar")) + "a")) + ("hello" + "foo bar"))) == ("\n\t" + ("a" + ("a" + "foo bar")))) && (((^(v6)) <= 2147483648) || (0 == 255))) || (!false)) {
        var v80 bool = (((v3 > v3) || (((("foo bar" + "hello") + ("" + "")) + "") < "\n\t")) && (65535 == 594))
        var v81 int16 = 32767
        _ = v80
        _ = v81
    } else {
        v82 := GenS4{F0: v4, F1: int(v4), F2: 42}
        _ = v82
    }
    fmt.Println(v6, v4, v3)
    var v83 int16 = -32768
    var v84 float32 = 1.5
    switch ("" + "\n\t") {
    case "\n\t":
        if true {
            var v85 int16 = v83
            v84 += v84
            _ = v85
        }
    case "foo bar":
        v86 := GenS5{F0: 0}
        _ = v86
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v83:", v83)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v79
    _ = v83
    _ = v84
}
