package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool; F1 string; F2 uint32 }
type GenS5 struct { F0 bool; F1 uint64; F2 uint8 }
type GenS6 struct { F0 int }

func main() {
    v1 := make(map[string]int)
    var v2 float64 = 1.5
    v3 := make(map[string]bool)
    var v4 string = ("hello" + "")
    v5 := []int{}
    var v6 uint16 = 1
    switch 32767 {
    case 2147483647:
        v6 = v6
    case 127:
        v7 := GenS1{F0: 2}
        _ = v7
    }
    v8 := v1[v4]
    var v9 float64 = v2
    if true {
        var v10 int8 = -128
        if (false && true) {
            var v11 int8 = (v10 + (-(v10)))
            _ = v11
        }
        v6 -= v6
        _ = v10
    }
    for fi12 := 0; fi12 < 2; fi12++ {
        var v13 uint8 = uint8(v6)
        _ = fi12
        _ = v13
    }
    fmt.Println(v8, v6, v9)
    switch v8 {
    case -9223372036854775808:
        switch ("a" + "foo bar") {
        case "\n\t":
            if (1 > 2147483647) {
                v4 = v4
                v14 := GenS1{F0: 2147483647}
                _ = v14
            }
        }
    case 479:
        fmt.Println(v2, v9)
    case 255:
        if (true || (!false)) {
            switch v4 {
            case "foo bar":
                var v15 int16 = -1
                _ = v15
            case "":
                var v16 int16 = -32768
                _ = v16
            case "\n\t":
                var v17 int = 0
                _ = v17
            }
        } else {
            if (1e10 != 1e10) {
                var v18 uint8 = 67
                _ = v18
            }
            var v19 string = ""
            _ = len(v1)
            _ = v19
        }
    default:
        for fi20 := 0; fi20 < 8; fi20++ {
            fmt.Println(v2)
            fmt.Println(v8, v8)
            v21 := GenS2{F0: (0 > 0), F1: true}
            _ = fi20
            _ = v21
        }
    }
    var v22 uint16 = 40
    if (0.0 == 3.14159) {
        for fi23 := 0; fi23 < 2; fi23++ {
            for fi24 := 0; fi24 < 17; fi24++ {
                if false {
                    v3[v4] = false
                    switch fi24 {
                    case 1:
                        v9 += (((v9 - v9) - v2) * (v9 * 3.14159))
                    case 2147483647:
                        if (false && true) {
                            var v25 uint8 = 42
                            if false {
                                v9 = v2
                            }
                            v6 = v6
                            var v26 int = v8
                            v27 := GenS1{F0: 4294967295}
                            v25 -= (v25 * (v25 / ((v25) | 1)))
                            var v28 uint8 = 2
                            _ = v25
                            _ = v26
                            _ = v27
                            _ = v28
                        } else {
                            v29 := GenS2{F0: (4294967295 <= 4294967295), F1: true}
                            v30 := GenS3{F0: 127}
                            var v31 uint32 = 0
                            _ = v29
                            _ = v30
                            _ = v31
                        }
                    }
                    if false {
                        switch v4 {
                        case "hello":
                            var v32 int32 = -1
                            _ = v32
                        case "a":
                            if (float32(v9) == ((-2.25 / -2.25) / float32(float64(-2.25)))) {
                                v3[(v4 + (("foo bar" + v4) + "\n\t"))] = ((^(v6)) >= (v22 * 43))
                                v6 = (v22 / (((v6 + (v22 + (v22 - v6)))) | 1))
                            }
                            var v33 float64 = (((0.0 / 1e10) + float64((-2.25 + ((3.14159 + (0.0 + float32(v9))) * 1e10)))) - float64(1.5))
                            _ = v33
                        default:
                            _ = len(v5)
                        }
                        if true {
                            v8 = fi24
                            var v34 int64 = 13
                            _ = v34
                        }
                        fmt.Println(v2, v6, v2)
                        if false {
                            var v35 float32 = -2.25
                            var v36 uint16 = 0
                            v37 := GenS1{F0: 1}
                            _ = v35
                            _ = v36
                            _ = v37
                        } else {
                            v8 += (v8 - (-(v8)))
                            v9 = (v9 / 1.0)
                            v9 = (v2 / 0.1)
                        }
                    } else {
                        v38 := GenS4{F0: false, F1: ((("a" + ("foo bar" + "hello")) + ("a" + v4)) + v4), F2: 956}
                        var v39 int64 = 0
                        v40 := GenS1{F0: 4294967295}
                        _ = v38
                        _ = v39
                        _ = v40
                    }
                }
                _ = fi24
            }
            v9 = v2
            _ = fi23
        }
        var v41 float64 = ((float64(1e10) / (((v2 - v9) * 1.5) - (float64(3.14159) + (v9 / 1.0)))) - float64(float32((v2 / v9))))
        if (!(-104 > 100)) {
            if (382 == 0) {
                v6 = (^(v22))
            }
            for fi42 := 0; fi42 < 7; fi42++ {
                _ = len(v5)
                var v43 int = (-(fi42))
                v44 := GenS4{F0: (!((!((v6 * v6) >= (v22 % (((v22 - 65535)) | 1)))) || (v4 != "\n\t"))), F1: v4, F2: 295}
                _ = fi42
                _ = v43
                _ = v44
            }
        } else {
            _ = len(v3)
        }
        _ = v41
    } else {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    v9 -= 1e10
    fmt.Println(v9)
    if true {
        for fi45 := 0; fi45 < 6; fi45++ {
            v9 += v2
            _ = fi45
        }
        v46 := GenS5{F0: true, F1: 758, F2: 95}
        _ = len(v3)
        _ = v46
    }
    v47 := GenS6{F0: v8}
    if (!false) {
        var v48 int8 = 120
        var v49 uint32 = 759
        var v50 bool = (int8(v22) > v48)
        _ = v48
        _ = v49
        _ = v50
    }
    var v51 uint64 = 1
    if (-91 == 42) {
        fmt.Println(v9, v8, v6)
    } else {
        if (42 <= 1) {
            if true {
                for fi52 := 0; fi52 < 18; fi52++ {
                    v53 := GenS6{F0: -981}
                    if false {
                        var v54 float32 = 0.0
                        _ = v54
                    } else {
                        var v55 string = "\n\t"
                        _ = v55
                    }
                    v56 := GenS3{F0: 255}
                    _ = fi52
                    _ = v53
                    _ = v56
                }
                v57 := GenS6{F0: v8}
                v8 = (v8 & 42)
                _ = v57
            }
            var v58 uint8 = 105
            _ = v58
        }
        switch (v8 * (v8 / ((-1) | 1))) {
        case 32768:
            fmt.Println(v2, v6, v51)
        }
        if len(v5) > 0 {
            v5[0] = v8
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v51
    _ = v6
    _ = v8
    _ = v9
}
