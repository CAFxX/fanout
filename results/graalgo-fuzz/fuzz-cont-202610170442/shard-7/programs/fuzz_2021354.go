package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 bool; F2 int8 }
type GenS4 struct { F0 int32; F1 int32 }
type GenS5 struct { F0 int; F1 int64 }
type GenS6 struct { F0 uint32; F1 uint }
type GenS7 struct { F0 int8; F1 uint; F2 string }

func main() {
    var v1 uint8 = 255
    var v2 uint = 4294967295
    var v3 int16 = 32767
    v4 := make(map[int]string)
    v5 := make(map[string]string)
    switch "hello" {
    case "\n\t":
        v3 = 1
    }
    fmt.Println(v2, v1, v3)
    var v6 uint32 = uint32(v1)
    if (false || (!(true && (!(!(v6 <= 4294967295)))))) {
        switch -2147483648 {
        case -810:
            v7 := GenS1{F0: v6}
            _ = v7
        case 0:
            fmt.Println(v1, v2)
        case 9223372036854775807:
            for fi8 := 0; fi8 < 18; fi8++ {
                var v9 uint8 = (v1 + v1)
                _ = fi8
                _ = v9
            }
        }
        var v10 uint = 4294967295
        v11 := GenS1{F0: (v6 / ((v6) | 1))}
        _ = v10
        _ = v11
    } else {
        fmt.Println(v6, v3)
    }
    for fi12 := 0; fi12 < 7; fi12++ {
        var v13 uint8 = (^(v1))
        if true {
            var v14 uint32 = (^(v6))
            switch ("a" + ("" + "hello")) {
            case "hello":
                for fi15 := 0; fi15 < 16; fi15++ {
                    v6 -= v6
                    _ = fi15
                }
            }
            _ = v14
        }
        _ = fi12
        _ = v13
    }
    switch "a" {
    case "\n\t":
        var v16 int = 1
        _ = v16
    }
    var v17 int64 = 20
    v3 = (v3 + v3)
    v18 := GenS1{F0: (v6 >> ((v2) & 63))}
    for fi19 := 0; fi19 < 1; fi19++ {
        var v20 int64 = (v17 + 0)
        _ = fi19
        _ = v20
    }
    if (!false) {
        _ = len(v5)
        if true {
            v21 := v5[((("" + ("a" + "")) + "\n\t") + (("foo bar" + "foo bar") + "a"))]
            v22 := []int{4294967295}
            v6 = v6
            _ = v21
            _ = v22
        } else {
            v17 = (v17 / ((v17) | 1))
            fmt.Println(v1, v3)
        }
        if ((-(v17)) != -1) {
            if false {
                var v23 bool = true
                if ((42 > 42) || v23) {
                    v24 := GenS2{F0: 42}
                    var v25 int64 = (v17 % ((v17) | 1))
                    switch "\n\t" {
                    case "a":
                        var v26 uint64 = 997
                        _ = v26
                    default:
                        v27 := []int{9223372036854775807, 613, -517, 2147483647}
                        _ = v27
                    }
                    _ = v24
                    _ = v25
                }
                _ = v23
            }
            v28 := GenS3{F0: 42, F1: false, F2: 100}
            _ = v28
        } else {
            var v29 int16 = (v3 + v3)
            _ = v29
        }
    } else {
        _ = len(v4)
        for fi30 := 0; fi30 < 4; fi30++ {
            var v31 uint64 = 42
            v17 += v17
            if false {
                v1 = (v1 / ((232) | 1))
                switch ("" + "foo bar") {
                case "":
                    var v32 uint8 = v1
                    _ = v32
                }
                switch 32767 {
                case 32767:
                    if false {
                        v4[fi30] = "\n\t"
                        var v33 uint16 = 313
                        v34 := GenS3{F0: v31, F1: (v17 <= -441), F2: 127}
                        _ = v33
                        _ = v34
                    } else {
                        v35 := GenS4{F0: -1, F1: 1}
                        v36 := []int{fi30, (fi30 * fi30), fi30}
                        for fi37 := 0; fi37 < 17; fi37++ {
                            var v38 int8 = 42
                            var v39 int = fi37
                            switch (fi37 * 127) {
                            case -434:
                                v2 = v2
                            default:
                                _ = len(v36)
                                if true {
                                    v36 = append(v36, v39)
                                    switch ("hello" + "a") {
                                    case "foo bar":
                                        var v40 float64 = 1e10
                                        _ = v40
                                    case "a":
                                        if false {
                                            v41 := GenS5{F0: v39, F1: v17}
                                            var v42 uint8 = v1
                                            _ = v41
                                            _ = v42
                                        }
                                    case "\n\t":
                                        var v43 int32 = -1
                                        var v44 int32 = 2147483647
                                        _ = v43
                                        _ = v44
                                    default:
                                        var v45 uint = (v2 - 259)
                                        _ = v45
                                    }
                                    switch "" {
                                    case "\n\t":
                                        v6 -= v6
                                    case "":
                                        var v46 int32 = -1
                                        v47 := GenS3{F0: 4294967295, F1: true, F2: (v38 % ((-1) | 1))}
                                        if (((float32(-2.25) - float32(float64(0.0))) / 3.14159) < 3.14159) {
                                            var v48 int8 = (v38 - (v38 * 1))
                                            if true {
                                                if (uint32(v2) < 2147483648) {
                                                    delete(v5, "foo bar")
                                                    if true {
                                                        var v49 string = ("" + "a")
                                                        v31 = (v31 * v31)
                                                        _ = v49
                                                    }
                                                } else {
                                                    if len(v36) > 0 {
                                                        v36[0] = 1
                                                    }
                                                    switch ("\n\t" + "") {
                                                    case "foo bar":
                                                        var v50 uint8 = (v1 % ((v1) | 1))
                                                        _ = v50
                                                    case "":
                                                        switch "" {
                                                        case "hello":
                                                            v17 = (-(v17))
                                                        case "":
                                                            if (false || (128 <= 4294967295)) {
                                                                fmt.Println(v3)
                                                                switch ("a" + (("hello" + ("a" + "")) + "")) {
                                                                case "foo bar":
                                                                    var v51 string = "\n\t"
                                                                    _ = v51
                                                                }
                                                                switch fi37 {
                                                                case -506:
                                                                    _ = len(v5)
                                                                case 32768:
                                                                    var v52 float32 = -2.25
                                                                    _ = v52
                                                                }
                                                            }
                                                        }
                                                    case "a":
                                                        switch "a" {
                                                        case "\n\t":
                                                            if true {
                                                                v38 = v38
                                                                v53 := GenS2{F0: (v17 ^ int64(v31))}
                                                                _ = v53
                                                            }
                                                        default:
                                                            var v54 int = fi37
                                                            _ = v54
                                                        }
                                                    }
                                                }
                                                if len(v36) > 0 {
                                                    v36[0] = (fi37 << (((v2 << (((^(v2))) & 63))) & 63))
                                                }
                                                var v55 int = (fi37 - (v39 << (((v2 << (((v2 << ((v2) & 63))) & 63))) & 63)))
                                                _ = v55
                                            }
                                            if false {
                                                if (!(!(!true))) {
                                                    switch "hello" {
                                                    case "\n\t":
                                                        _ = len(v5)
                                                    case "a":
                                                        var v56 uint8 = (v1 - (v1 / ((v1) | 1)))
                                                        _ = v56
                                                    case "":
                                                        v3 = v3
                                                    }
                                                    var v57 int = (fi30 * fi30)
                                                    v31 = v31
                                                    var v58 int16 = v3
                                                    var v59 int64 = 181
                                                    var v60 string = "foo bar"
                                                    v3 -= v58
                                                    _ = v57
                                                    _ = v58
                                                    _ = v59
                                                    _ = v60
                                                }
                                                var v61 uint64 = 876
                                                var v62 float64 = 3.14159
                                                _ = v61
                                                _ = v62
                                            }
                                            var v63 string = ("a" + ("a" + "a"))
                                            _ = v48
                                            _ = v63
                                        }
                                        _ = v46
                                        _ = v47
                                    case "a":
                                        if (false || true) {
                                            fmt.Println(v39, v38)
                                        }
                                    default:
                                        var v64 float32 = float32(1.5)
                                        _ = v64
                                    }
                                } else {
                                    if ((v1 & v1) == 128) {
                                        fmt.Println(fi37, v31)
                                        fmt.Println(v39)
                                        v36 = append(v36, fi30)
                                    } else {
                                        var v65 uint = uint(v31)
                                        v66 := GenS6{F0: v6, F1: 442}
                                        v65 = v2
                                        _ = v65
                                        _ = v66
                                    }
                                }
                            }
                            _ = fi37
                            _ = v38
                            _ = v39
                        }
                        _ = v35
                        _ = v36
                    }
                default:
                    var v67 uint16 = 0
                    _ = v67
                }
            }
            _ = fi30
            _ = v31
        }
        var v68 int16 = v3
        _ = v68
    }
    v69 := GenS7{F0: 60, F1: (v2 ^ (v2 / ((v2) | 1))), F2: "foo bar"}
    var v70 int64 = v17
    v71 := v5[(("" + ("a" + "foo bar")) + "foo bar")]
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v69
    _ = v70
    _ = v71
}
