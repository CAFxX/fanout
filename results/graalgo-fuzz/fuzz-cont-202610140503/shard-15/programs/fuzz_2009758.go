package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint }

func main() {
    v1 := make(map[string]string)
    var v2 bool = (1 >= 2147483647)
    var v3 float32 = (3.14159 + ((1.5 - -2.25) * 1e10))
    var v4 uint32 = 628
    v5 := []string{((("foo bar" + "\n\t") + "hello") + "a")}
    var v6 bool = (18446744073709551615 <= 1)
    for fi7 := 0; fi7 < 15; fi7++ {
        if v6 {
            for fi8 := 0; fi8 < 5; fi8++ {
                v9 := GenS1{F0: ((!false) && v2)}
                switch 2147483647 {
                case 32767:
                    fmt.Println(v2, v2, fi7)
                case 9223372036854775807:
                    var v10 uint32 = v4
                    _ = v10
                default:
                    _ = len(v1)
                }
                _ = fi8
                _ = v9
            }
            for fi11 := 0; fi11 < 19; fi11++ {
                var v12 int16 = -1
                if v6 {
                    var v13 int32 = 0
                    v3 -= v3
                    _ = v13
                }
                v12 = 203
                if v6 {
                    if ((1e10 - (0.0 + 0.0)) == (((-2.25 / ((1e10 * 0.0) * (((0.0 * 1e10) / (float64((1.5 + (((v3 + 1.5) + 0.0) / 3.14159))) - 0.1)) - (3.14159 - -2.25)))) * 0.0) - float64((float32(1.5) * 1e10)))) {
                        var v14 uint32 = v4
                        v4 = v4
                        if v6 {
                            var v15 float64 = (1e10 * (1.5 * 1.5))
                            v16 := GenS1{F0: (fi7 <= (fi7 | (fi7 >> ((323) & 63))))}
                            _ = v15
                            _ = v16
                        }
                        _ = v14
                    }
                    fmt.Println(fi11, fi11, fi7)
                    switch fi11 {
                    case 32767:
                        if (!false) {
                            v12 = v12
                            fmt.Println(v3, v4, fi7)
                            v17 := GenS2{F0: 83}
                            _ = v17
                        } else {
                            v3 = v3
                        }
                    default:
                        v2 = true
                    }
                }
                var v18 string = ("hello" + "")
                _ = fi11
                _ = v12
                _ = v18
            }
        } else {
            v4 = 2147483648
            _ = len(v5)
            v2 = v2
        }
        fmt.Println(v3, v4)
        if (!(-1 <= 42)) {
            v19 := GenS2{F0: 381}
            if false {
                var v20 uint64 = 42
                _ = v20
            }
            v1["hello"] = ("a" + "\n\t")
            _ = v19
        }
        _ = fi7
    }
    v1["foo bar"] = ("\n\t" + ("a" + ("foo bar" + "foo bar")))
    for fi21 := 0; fi21 < 2; fi21++ {
        for fi22 := 0; fi22 < 19; fi22++ {
            if (!v6) {
                switch "\n\t" {
                case "a":
                    switch (fi22 % (((fi22 >> ((1) & 63))) | 1)) {
                    case -442:
                        _ = len(v5)
                        if true {
                            v2 = v6
                            fmt.Println(v6)
                            v3 += 1e10
                        }
                        v6 = (fi21 <= -739)
                    }
                case "hello":
                    var v23 int = 4294967295
                    _ = v23
                case "foo bar":
                    if (2147483647 >= -279) {
                        if len(v5) > 0 {
                            v5[0] = ""
                        }
                        var v24 uint64 = 18446744073709551615
                        var v25 int32 = -1
                        _ = v24
                        _ = v25
                    } else {
                        switch (fi22 >> ((4294967295) & 63)) {
                        case 4294967295:
                            v3 = ((((float32(float64(v3)) / v3) + (float32(1.5) - v3)) + (v3 - v3)) / -2.25)
                        case -195:
                            if (v2 && ((1e10 * (v3 + v3)) <= 0.1)) {
                                var v26 float64 = 1e10
                                if (0.1 != v3) {
                                    var v27 float32 = v3
                                    var v28 uint = 619
                                    v3 = v3
                                    _ = v27
                                    _ = v28
                                } else {
                                    if v2 {
                                        switch ("hello" + "\n\t") {
                                        case "a":
                                            v3 = (v3 - v3)
                                            fmt.Println(v4)
                                        case "hello":
                                            switch (fi21 ^ fi21) {
                                            case -698:
                                                v26 = v26
                                            case -334:
                                                if v6 {
                                                    if v2 {
                                                        v29 := GenS3{F0: 255, F1: 0}
                                                        if (true && true) {
                                                            _ = len(v1)
                                                            var v30 bool = false
                                                            v6 = (0 >= 1)
                                                            _ = v30
                                                        } else {
                                                            v2 = (v4 != (v4 << ((42) & 63)))
                                                        }
                                                        var v31 uint8 = 192
                                                        _ = v29
                                                        _ = v31
                                                    }
                                                    v32 := GenS1{F0: v6}
                                                    v26 = ((v26 + (1.5 + ((v26 - (0.0 / 1e10)) / -2.25))) / ((float64((v3 * 0.0)) / v26) / ((0.1 - 0.1) + (v26 * 3.14159))))
                                                    _ = v32
                                                }
                                            }
                                        case "\n\t":
                                            fmt.Println(v3, v2, fi21)
                                        }
                                        v3 += 1.5
                                    }
                                    fmt.Println(fi21)
                                    var v33 string = (("a" + "\n\t") + ("" + "foo bar"))
                                    _ = v33
                                }
                                var v34 uint = 18446744073709551615
                                _ = v26
                                _ = v34
                            } else {
                                v2 = (((v4 ^ v4) >= v4) && (!true))
                                var v35 int64 = 1
                                if v6 {
                                    switch fi21 {
                                    case -1:
                                        switch "hello" {
                                        case "hello":
                                            v4 = v4
                                        default:
                                            var v36 int64 = v35
                                            _ = v36
                                        }
                                        v35 = (v35 % ((v35) | 1))
                                        if false {
                                            var v37 int16 = -1
                                            _ = v37
                                        } else {
                                            var v38 int8 = -128
                                            switch "\n\t" {
                                            case "\n\t":
                                                if len(v5) > 0 {
                                                    v5[0] = "foo bar"
                                                }
                                            }
                                            if (v2 && (!v6)) {
                                                var v39 float32 = -2.25
                                                if len(v5) > 0 {
                                                    v5[0] = ("hello" + "hello")
                                                }
                                                v4 += 1
                                                _ = v39
                                            }
                                            _ = v38
                                        }
                                    case 405:
                                        fmt.Println(fi21, fi21)
                                    }
                                    fmt.Println(fi21, v3, fi21)
                                    v40 := GenS1{F0: ((fi22 >> ((0) & 63)) < (fi22 - fi21))}
                                    _ = v40
                                } else {
                                    if (-1 >= 127) {
                                        delete(v1, ("a" + "foo bar"))
                                        switch fi22 {
                                        case 553:
                                            v6 = (v35 == v35)
                                        default:
                                            v6 = true
                                        }
                                        fmt.Println(v6, v35, v35)
                                    }
                                    fmt.Println(v2)
                                }
                                _ = v35
                            }
                        }
                        var v41 uint32 = (v4 << ((65535) & 63))
                        fmt.Println(v41, fi22)
                        _ = v41
                    }
                }
                var v42 string = ""
                _ = v42
            } else {
                v43 := GenS3{F0: 1, F1: 0}
                _ = v43
            }
            var v44 int32 = 0
            v2 = ("foo bar" != "a")
            _ = fi22
            _ = v44
        }
        _ = len(v1)
        v45 := GenS4{F0: fi21}
        _ = fi21
        _ = v45
    }
    switch "\n\t" {
    case "":
        var v46 string = "a"
        _ = v46
    }
    v4 += (v4 / (((v4 - v4)) | 1))
    v4 -= 65535
    fmt.Println(v6, v2, v2)
    v47 := GenS5{F0: 276}
    if ((v2 && ((-32768 == 1000) || v6)) || (128 == 250)) {
        if (v4 < 2) {
            if len(v5) > 0 {
                v5[0] = ""
            }
        } else {
            v3 -= 1.5
        }
    } else {
        v4 = v4
    }
    for fi48 := 0; fi48 < 14; fi48++ {
        if len(v5) > 0 {
            v5[0] = "\n\t"
        }
        for fi49 := 0; fi49 < 2; fi49++ {
            v50 := GenS4{F0: (fi48 >> ((148) & 63))}
            switch fi49 {
            case 824:
                v3 = v3
            case 127:
                if true {
                    fmt.Println(fi49)
                    v3 += 0.0
                } else {
                    if v6 {
                        var v51 int8 = 0
                        var v52 string = (("a" + ("foo bar" + "hello")) + ((("\n\t" + ("a" + "\n\t")) + "a") + "\n\t"))
                        _ = v51
                        _ = v52
                    } else {
                        var v53 int64 = 224
                        _ = v53
                    }
                }
            }
            switch (fi48 - (fi49 | (^(fi49)))) {
            case 32768:
                fmt.Println(v3, v2, fi49)
            case 4294967295:
                fmt.Println(fi49, v2, v3)
            }
            _ = fi49
            _ = v50
        }
        _ = fi48
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
}
