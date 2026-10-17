package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint; F1 bool }
type GenS4 struct { F0 uint8; F1 int16; F2 int16 }
func genFn0() uint {
    v6 := GenS1{F0: 0, F1: 1, F2: 127}
    var v7 float32 = -2.25
    if (32767 < 111) {
        var v8 uint = 65535
        _ = v8
    } else {
        v9 := []int{0, -282}
        v7 = v7
        if (-109 != 816) {
            var v10 int64 = -60
            _ = v10
        } else {
            v7 = v7
            if (-109 != 100) {
                if len(v9) > 0 {
                    v9[0] = 42
                }
                v11 := GenS1{F0: 1, F1: 255, F2: 1}
                var v12 uint = 217
                _ = v11
                _ = v12
            } else {
                switch "a" {
                case "":
                    for fi13 := 0; fi13 < 15; fi13++ {
                        for fi14 := 0; fi14 < 2; fi14++ {
                            if true {
                                var v15 uint = 18446744073709551615
                                _ = v15
                            } else {
                                var v16 int16 = 1
                                v17 := GenS2{F0: -2147483648}
                                v18 := GenS3{F0: 255, F1: (!false)}
                                _ = v16
                                _ = v17
                                _ = v18
                            }
                            var v19 int16 = -32768
                            var v20 int64 = 1
                            _ = fi14
                            _ = v19
                            _ = v20
                        }
                        var v21 uint8 = 127
                        _ = fi13
                        _ = v21
                    }
                default:
                    switch "a" {
                    case "hello":
                        if false {
                            for fi22 := 0; fi22 < 7; fi22++ {
                                var v23 int32 = -972
                                v7 = v7
                                _ = fi22
                                _ = v23
                            }
                            _ = len(v9)
                        }
                    case "":
                        var v24 int8 = 25
                        _ = v24
                    case "foo bar":
                        fmt.Println(v7)
                    }
                }
                fmt.Println(v7)
            }
            for fi25 := 0; fi25 < 3; fi25++ {
                switch fi25 {
                case 1:
                    var v26 bool = ((2 > 743) && true)
                    _ = v26
                case 4294967295:
                    v27 := GenS3{F0: 4294967295, F1: (false && true)}
                    _ = v27
                }
                _ = fi25
            }
        }
        _ = v9
    }
    var v28 uint8 = 163
    var v29 float64 = float64((v7 + v7))
    _ = v28
    _ = v29
    _ = v6
    _ = v7
    return 18446744073709551615
}

func main() {
    var v1 uint8 = 128
    v2 := []string{"", ((("hello" + "a") + "a") + ""), ("a" + "")}
    var v3 uint32 = 2147483648
    v4 := make(map[int]uint32)
    var v5 string = ("foo bar" + "\n\t")
    var v30 string = (v5 + "")
    for fi31 := 0; fi31 < 14; fi31++ {
        var v32 float64 = 1e10
        for fi33 := 0; fi33 < 19; fi33++ {
            v34 := GenS2{F0: 0}
            if true {
                switch (^(fi31)) {
                case 128:
                    var v35 uint = 65535
                    _ = len(v2)
                    switch 4294967295 {
                    case 32768:
                        if ((^(v1)) <= v1) {
                            v30 = v30
                        } else {
                            fmt.Println(v35, v5)
                            v32 = float64(0.0)
                            if (!(2 < 18446744073709551615)) {
                                v35 = 255
                            } else {
                                v36 := GenS4{F0: (v1 % ((v1) | 1)), F1: -396, F2: int16(v1)}
                                _ = v36
                            }
                        }
                    case -2147483648:
                        if false {
                            switch fi33 {
                            case 1:
                                if false {
                                    var v37 uint32 = v3
                                    _ = v37
                                } else {
                                    switch fi33 {
                                    case 127:
                                        switch v5 {
                                        case "hello":
                                            fmt.Println(v30)
                                        }
                                    }
                                    if (!(!(!((^(v1)) >= (^(v1)))))) {
                                        var v38 int32 = -2147483648
                                        _ = len(v2)
                                        v4[-1] = v3
                                        _ = v38
                                    }
                                    var v39 uint8 = 182
                                    _ = v39
                                }
                            case 4294967295:
                                v40 := GenS2{F0: 536}
                                _ = v40
                            default:
                                v5 = v30
                                fmt.Println(v32)
                            }
                            var v41 float32 = float32((1.5 - v32))
                            _ = v41
                        }
                    }
                    _ = v35
                case 127:
                    switch fi31 {
                    case 0:
                        if (573 <= -1) {
                            v30 = (("hello" + v30) + (((v5 + ("hello" + ((v5 + "a") + (v30 + (v5 + (v30 + (((v5 + v30) + "") + "foo bar"))))))) + v30) + "a"))
                        }
                    default:
                        var v42 uint16 = 255
                        _ = v42
                    }
                default:
                    switch (v30 + v5) {
                    case "\n\t":
                        v43 := GenS1{F0: -37, F1: 18446744073709551615, F2: 127}
                        _ = v43
                    default:
                        _ = len(v4)
                    }
                }
                v44 := GenS2{F0: 0}
                var v45 int64 = 263
                _ = v44
                _ = v45
            }
            _ = fi33
            _ = v34
        }
        _ = fi31
        _ = v32
    }
    switch (v5 + (v30 + ("hello" + "hello"))) {
    case "hello":
        var v46 uint64 = 4294967295
        _ = v46
    }
    var v47 string = ("\n\t" + "hello")
    v48 := GenS1{F0: 0, F1: 2, F2: 42}
    fmt.Println(v1)
    if (true && false) {
        var v49 uint32 = v3
        v1 = (v1 * (v1 & v1))
        for fi50 := 0; fi50 < 13; fi50++ {
            fmt.Println(fi50, v47, v5)
            v30 = "\n\t"
            _ = fi50
        }
        _ = v49
    } else {
        v47 = ("hello" + v5)
        for fi51 := 0; fi51 < 19; fi51++ {
            var v52 string = ""
            _ = fi51
            _ = v52
        }
        var v53 uint32 = (^(v3))
        _ = v53
    }
    var v54 uint16 = 569
    v30 = "hello"
    _ = len(v2)
    var v55 float32 = (1.5 + (0.1 / 1e10))
    fmt.Println(v5, v55)
    var v56 uint64 = 2
    var v57 int64 = -9223372036854775808
    switch 128 {
    case 32768:
        if (!false) {
            fmt.Println(v55)
            var v58 string = "a"
            _ = v58
        }
    case 359:
        v57 += -1
    case 2147483647:
        var v59 float64 = (-2.25 / 0.1)
        _ = v59
    default:
        if ((!true) && false) {
            v5 = ("" + v5)
        }
    }
    switch ("" + "\n\t") {
    case "\n\t":
        v30 = ""
    case "hello":
        if true {
            v2 = append(v2, v47)
            _ = len(v2)
            switch ("a" + (v47 + v47)) {
            case "a":
                var v60 int = 9223372036854775807
                _ = v60
            default:
                if false {
                    var v61 int64 = (v57 << ((65535) & 63))
                    _ = len(v2)
                    if false {
                        if true {
                            switch -9223372036854775808 {
                            case -2147483648:
                                for fi62 := 0; fi62 < 1; fi62++ {
                                    for fi63 := 0; fi63 < 5; fi63++ {
                                        delete(v4, (fi62 - (fi63 + fi62)))
                                        var v64 uint = uint(v56)
                                        _ = fi63
                                        _ = v64
                                    }
                                    if false {
                                        if len(v2) > 0 {
                                            v2[0] = v47
                                        }
                                    }
                                    _ = fi62
                                }
                            case 32768:
                                var v65 int16 = -884
                                _ = v65
                            case -9223372036854775808:
                                if len(v2) > 0 {
                                    v2[0] = v5
                                }
                            default:
                                var v66 int32 = -154
                                _ = v66
                            }
                            v56 = (v56 & v56)
                        }
                        var v67 int = 32768
                        _ = v67
                    } else {
                        var v68 bool = (226 < -32768)
                        v57 -= v57
                        _ = v68
                    }
                    _ = v61
                }
            }
        }
    default:
        var v69 uint64 = (v56 / ((v56) | 1))
        _ = v69
    }
    v70 := GenS2{F0: 1}
    v56 = (v56 | v56)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v54
    _ = v55
    _ = v56
    _ = v57
    _ = v70
}
