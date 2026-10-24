package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 int; F1 int; F2 uint32 }

func main() {
    v1 := []uint32{42, 903}
    var v2 uint = 65535
    var v3 string = ""
    var v4 uint8 = 0
    v5 := []uint32{}
    _ = len(v5)
    var v6 bool = false
    var v7 float64 = 0.1
    for fi8 := 0; fi8 < 17; fi8++ {
        for fi9 := 0; fi9 < 17; fi9++ {
            var v10 float32 = 1e10
            if len(v5) > 0 {
                v5[0] = 0
            }
            var v11 int = (fi9 & fi8)
            _ = fi9
            _ = v10
            _ = v11
        }
        var v12 int8 = 127
        var v13 float64 = v7
        _ = fi8
        _ = v12
        _ = v13
    }
    switch v3 {
    case "":
        for fi14 := 0; fi14 < 18; fi14++ {
            var v15 uint16 = 761
            var v16 uint32 = 121
            _ = fi14
            _ = v15
            _ = v16
        }
    case "foo bar":
        v2 -= 255
    case "hello":
        v7 -= 0.1
    }
    switch 308 {
    case -175:
        fmt.Println(v2, v4)
    default:
        v17 := GenS1{F0: 1}
        _ = v17
    }
    for fi18 := 0; fi18 < 12; fi18++ {
        v3 = (("\n\t" + v3) + (v3 + (v3 + v3)))
        _ = fi18
    }
    v19 := GenS1{F0: 984}
    for fi20 := 0; fi20 < 17; fi20++ {
        if len(v1) > 0 {
            v1[0] = 668
        }
        if (-2147483648 < -935) {
            for fi21 := 0; fi21 < 19; fi21++ {
                v2 -= (v2 & v2)
                var v22 int32 = 990
                _ = fi21
                _ = v22
            }
        }
        _ = fi20
    }
    if false {
        if len(v5) > 0 {
            v5[0] = 2147483648
        }
        for fi23 := 0; fi23 < 14; fi23++ {
            var v24 float64 = (v7 + 1.5)
            fmt.Println(v6)
            v6 = ((!false) && (128 != 42))
            _ = fi23
            _ = v24
        }
    }
    v5 = append(v5, 2147483647)
    for fi25 := 0; fi25 < 15; fi25++ {
        if true {
            var v26 float64 = ((v7 + (float64((-2.25 - 0.1)) * (v7 / -2.25))) - float64((-2.25 / (1e10 + (0.1 - 0.0)))))
            _ = v26
        } else {
            var v27 int8 = 23
            if v6 {
                var v28 uint16 = 878
                fmt.Println(v7, v6, v27)
                v3 = (((v3 + "foo bar") + v3) + ("hello" + v3))
                _ = v28
            } else {
                var v29 string = v3
                _ = v29
            }
            _ = v27
        }
        if (!(false && (32767 > 32767))) {
            for fi30 := 0; fi30 < 2; fi30++ {
                var v31 int32 = -619
                v4 = (v4 ^ v4)
                switch (("foo bar" + "hello") + v3) {
                case "\n\t":
                    fmt.Println(v6, fi25, v2)
                case "":
                    if v6 {
                        var v32 uint32 = 795
                        fmt.Println(v6)
                        if v6 {
                            if v6 {
                                var v33 int8 = 82
                                var v34 uint16 = 32767
                                var v35 float64 = v7
                                _ = v33
                                _ = v34
                                _ = v35
                            } else {
                                switch (fi25 + fi25) {
                                case -9223372036854775808:
                                    var v36 string = v3
                                    _ = v36
                                case 2147483647:
                                    v4 -= v4
                                    var v37 uint8 = (v4 >> ((161) & 63))
                                    _ = v37
                                case 9223372036854775807:
                                    if len(v1) > 0 {
                                        v1[0] = v32
                                    }
                                default:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                _ = len(v5)
                                var v38 int = (fi25 & (fi30 << (((v2 | 252)) & 63)))
                                _ = v38
                            }
                            if true {
                                switch fi25 {
                                case 127:
                                    if v6 {
                                        var v39 string = (v3 + ("" + v3))
                                        var v40 float32 = (1e10 - (float32(v7) - (float32((v7 + float64(1e10))) - (0.1 / -2.25))))
                                        _ = v39
                                        _ = v40
                                    }
                                    var v41 float32 = 0.1
                                    var v42 bool = true
                                    var v43 int64 = 0
                                    _ = v41
                                    _ = v42
                                    _ = v43
                                }
                                if (true && (!v6)) {
                                    switch v3 {
                                    case "foo bar":
                                        fmt.Println(v32, v7, v3)
                                    case "a":
                                        var v44 uint32 = v32
                                        _ = v44
                                    default:
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                    if (42 <= 9223372036854775807) {
                                        if v6 {
                                            var v45 bool = v6
                                            fmt.Println(v6, v32, v45)
                                            _ = v45
                                        }
                                        v31 -= (v31 >> ((158) & 63))
                                        var v46 int8 = 42
                                        _ = v46
                                    } else {
                                        v47 := GenS2{F0: 626, F1: -9223372036854775808}
                                        _ = v47
                                    }
                                } else {
                                    fmt.Println(v7, fi25, v2)
                                    v7 -= 3.14159
                                }
                            } else {
                                v1 = append(v1, (v32 << ((18446744073709551615) & 63)))
                                switch v3 {
                                case "foo bar":
                                    var v48 string = "foo bar"
                                    _ = v48
                                default:
                                    switch v3 {
                                    case "foo bar":
                                        if (1000 == 358) {
                                            v4 -= v4
                                            fmt.Println(v6)
                                        }
                                    case "":
                                        v3 = ("" + "\n\t")
                                        var v49 string = ((("\n\t" + v3) + "a") + v3)
                                        var v50 uint32 = 0
                                        _ = v49
                                        _ = v50
                                    }
                                }
                            }
                            fmt.Println(fi25, v4, v32)
                            var v51 uint32 = v32
                            _ = v51
                        }
                        _ = v32
                    } else {
                        fmt.Println(v3)
                    }
                case "foo bar":
                    var v52 bool = true
                    _ = v52
                default:
                    if v6 {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v6 = v6
                        var v53 int8 = -128
                        _ = v53
                    } else {
                        var v54 uint32 = 2147483648
                        if true {
                            v55 := GenS3{F0: (fi25 << ((1) & 63)), F1: -610, F2: (v54 << ((885) & 63))}
                            _ = v55
                        }
                        _ = v54
                    }
                }
                _ = fi30
                _ = v31
            }
        } else {
            v56 := GenS1{F0: -659}
            var v57 bool = v6
            _ = v56
            _ = v57
        }
        v2 = 4294967295
        _ = fi25
    }
    v4 = v4
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v6, v4, v6)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
