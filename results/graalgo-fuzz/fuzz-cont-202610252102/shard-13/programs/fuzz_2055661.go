package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 int32; F2 uint16 }
type GenS3 struct { F0 string; F1 int; F2 uint }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int32; F1 uint16 }

func main() {
    v1 := []int{-831}
    var v2 uint64 = 42
    v3 := []string{"", ("foo bar" + ("\n\t" + "")), (("foo bar" + "a") + ("hello" + ("a" + ((("\n\t" + "hello") + "a") + "foo bar"))))}
    var v4 int64 = 0
    v5 := []string{}
    v6 := GenS1{F0: 349}
    if len(v3) > 0 {
        v3[0] = "a"
    }
    var v7 bool = (4294967295 != (v2 / (((v2 ^ 550)) | 1)))
    fmt.Println(v2)
    v4 += v4
    switch (("" + "hello") + ("foo bar" + "hello")) {
    case "":
        var v8 uint16 = 321
        _ = v8
    case "hello":
        switch "foo bar" {
        case "a":
            if v7 {
                v4 = (v4 % (((v4 % (((v4 / (((v4 + v4)) | 1))) | 1))) | 1))
                _ = len(v5)
                if false {
                    var v9 int = 9223372036854775807
                    switch "a" {
                    case "foo bar":
                        if true {
                            var v10 int64 = v4
                            switch v9 {
                            case 414:
                                v3 = append(v3, ("foo bar" + "foo bar"))
                            case -9223372036854775808:
                                var v11 uint = 69
                                _ = v11
                            }
                            var v12 int64 = (^(v4))
                            _ = v10
                            _ = v12
                        }
                    case "a":
                        fmt.Println(v7)
                    case "":
                        if false {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        } else {
                            for fi13 := 0; fi13 < 17; fi13++ {
                                if (v7 && (v7 && v7)) {
                                    if v7 {
                                        fmt.Println(v4, v7)
                                        var v14 int64 = (v4 << ((4294967295) & 63))
                                        for fi15 := 0; fi15 < 8; fi15++ {
                                            if len(v1) > 0 {
                                                v1[0] = v9
                                            }
                                            _ = fi15
                                        }
                                        _ = v14
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    var v16 int32 = 2147483647
                                    _ = v16
                                }
                                if v7 {
                                    v17 := GenS1{F0: 32767}
                                    _ = len(v5)
                                    var v18 int16 = -32768
                                    _ = v17
                                    _ = v18
                                } else {
                                    v9 = (fi13 - (v9 * (v9 ^ 689)))
                                }
                                switch v9 {
                                case -40:
                                    if false {
                                        v19 := GenS1{F0: 0}
                                        _ = v19
                                    }
                                case 289:
                                    var v20 string = ("a" + ("foo bar" + "\n\t"))
                                    _ = v20
                                }
                                _ = fi13
                            }
                            v2 += (v2 % ((v2) | 1))
                            switch "\n\t" {
                            case "a":
                                for fi21 := 0; fi21 < 1; fi21++ {
                                    if (!v7) {
                                        for fi22 := 0; fi22 < 4; fi22++ {
                                            if len(v3) > 0 {
                                                v3[0] = (("a" + "\n\t") + "a")
                                            }
                                            _ = fi22
                                        }
                                    } else {
                                        switch (v9 / ((fi21) | 1)) {
                                        case 25:
                                            v23 := GenS1{F0: 21}
                                            _ = v23
                                        }
                                        v24 := GenS2{F0: int32(v2), F1: 0, F2: 107}
                                        _ = v24
                                    }
                                    if (!v7) {
                                        v1 = append(v1, -2147483648)
                                        var v25 int = (v9 / ((fi21) | 1))
                                        for fi26 := 0; fi26 < 13; fi26++ {
                                            _ = len(v1)
                                            fmt.Println(fi21, v25, v9)
                                            v27 := GenS3{F0: "\n\t", F1: (fi21 / ((v25) | 1)), F2: 18446744073709551615}
                                            _ = fi26
                                            _ = v27
                                        }
                                        _ = v25
                                    }
                                    fmt.Println(v9, v7)
                                    _ = fi21
                                }
                            case "hello":
                                for fi28 := 0; fi28 < 19; fi28++ {
                                    switch (^(v9)) {
                                    case 2147483647:
                                        v29 := GenS2{F0: -1, F1: 0, F2: 0}
                                        _ = v29
                                    }
                                    _ = fi28
                                }
                            default:
                                switch v9 {
                                case 63:
                                    var v30 uint32 = 751
                                    _ = v30
                                case 42:
                                    var v31 int8 = 0
                                    _ = v31
                                case 72:
                                    var v32 int8 = -53
                                    _ = v32
                                default:
                                    fmt.Println(v9)
                                }
                            }
                        }
                    }
                    v2 = v2
                    _ = v9
                }
            }
        case "foo bar":
            v33 := GenS3{F0: ("foo bar" + ""), F1: 1, F2: 894}
            _ = v33
        case "hello":
            switch "foo bar" {
            case "a":
                for fi34 := 0; fi34 < 2; fi34++ {
                    for fi35 := 0; fi35 < 3; fi35++ {
                        if (122 != v2) {
                            if (!(float64(-2.25) >= 0.1)) {
                                if v7 {
                                    fmt.Println(v2)
                                } else {
                                    v36 := GenS4{F0: -2147483648}
                                    _ = v36
                                }
                                v2 += (v2 | 18446744073709551615)
                                v37 := GenS2{F0: 9, F1: -2147483648, F2: 557}
                                _ = v37
                            } else {
                                var v38 string = "a"
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                switch v38 {
                                case "foo bar":
                                    v39 := GenS4{F0: 1}
                                    _ = v39
                                case "\n\t":
                                    v2 = v2
                                }
                                _ = v38
                            }
                        } else {
                            var v40 int16 = -898
                            _ = v40
                        }
                        var v41 int = fi34
                        _ = fi35
                        _ = v41
                    }
                    _ = fi34
                }
            case "":
                if false {
                    var v42 float64 = 1e10
                    _ = v42
                } else {
                    for fi43 := 0; fi43 < 13; fi43++ {
                        v4 += 42
                        if false {
                            switch ("" + ((("foo bar" + ("" + ("foo bar" + "hello"))) + "foo bar") + "")) {
                            case "foo bar":
                                v7 = (!v7)
                            case "\n\t":
                                var v44 float32 = 1.5
                                _ = v44
                            default:
                                var v45 int8 = 1
                                _ = v45
                            }
                            v46 := GenS1{F0: 1}
                            _ = v46
                        }
                        _ = fi43
                    }
                    fmt.Println(v2, v7)
                }
            case "hello":
                v47 := GenS5{F0: -623, F1: 42}
                _ = v47
            default:
                for fi48 := 0; fi48 < 8; fi48++ {
                    v7 = false
                    _ = fi48
                }
            }
        }
    }
    var v49 int64 = -472
    fmt.Println(v7, v7)
    fmt.Println(v49, v4, v2)
    var v50 uint8 = 2
    fmt.Println(v49, v7, v50)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v3 = append(v3, "hello")
    v50 = (v50 ^ 128)
    var v51 int32 = -16
    v52 := GenS3{F0: ("hello" + ""), F1: -638, F2: 255}
    if (0.0 <= 0.0) {
        var v53 bool = ((v4 / ((0) | 1)) < 42)
        var v54 bool = ((!((!(119 <= -6)) && (v53 || v7))) || v53)
        _ = v53
        _ = v54
    } else {
        if v7 {
            for fi55 := 0; fi55 < 2; fi55++ {
                v49 = (v49 % ((-9223372036854775808) | 1))
                v2 += (v2 << ((2) & 63))
                var v56 int16 = 0
                _ = fi55
                _ = v56
            }
            v4 = 0
        } else {
            var v57 int16 = 42
            fmt.Println(v50, v51, v49)
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = v57
        }
        var v58 uint = 255
        var v59 int8 = 1
        _ = v58
        _ = v59
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v6
    _ = v7
}
