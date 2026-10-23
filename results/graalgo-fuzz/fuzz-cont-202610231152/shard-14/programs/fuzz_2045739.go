package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 string }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8; F1 uint32 }
type GenS5 struct { F0 uint64; F1 int8; F2 bool }
type GenS6 struct { F0 int }

func main() {
    v1 := make(map[int]string)
    var v2 string = ((("a" + "hello") + "foo bar") + ((("foo bar" + "hello") + ("" + "a")) + "foo bar"))
    v3 := []string{}
    v4 := []uint32{}
    v2 = ((v2 + (v2 + ("" + "foo bar"))) + "a")
    fmt.Println(v2)
    v5 := GenS1{F0: 1, F1: 64, F2: v2}
    switch v2 {
    case "a":
        var v6 float32 = (((1.5 - 1.5) / -2.25) - float32((1.5 / 0.1)))
        _ = v6
    case "":
        _ = len(v4)
    default:
        for fi7 := 0; fi7 < 14; fi7++ {
            _ = len(v4)
            _ = fi7
        }
    }
    v8 := GenS1{F0: 127, F1: 30, F2: v2}
    switch ("" + ("hello" + v2)) {
    case "\n\t":
        if false {
            var v9 uint32 = 841
            if (true || false) {
                var v10 int32 = 2147483647
                switch 42 {
                case -9223372036854775808:
                    v1[-1] = v2
                case 42:
                    if ("hello" > v2) {
                        v2 = (("foo bar" + "foo bar") + v2)
                        delete(v1, -9223372036854775808)
                        var v11 uint16 = 1
                        _ = v11
                    } else {
                        var v12 int32 = int32(v9)
                        v13 := GenS2{F0: (v9 <= v9), F1: 0}
                        v14 := GenS3{F0: "foo bar"}
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                case 193:
                    for fi15 := 0; fi15 < 7; fi15++ {
                        v1[2147483647] = ("" + "\n\t")
                        var v16 int64 = 9223372036854775807
                        _ = fi15
                        _ = v16
                    }
                default:
                    v10 = v10
                }
                _ = v10
            } else {
                v9 += (v9 * 4294967295)
                v17 := GenS4{F0: 128, F1: v9}
                _ = v17
            }
            _ = v9
        } else {
            for fi18 := 0; fi18 < 10; fi18++ {
                _ = len(v1)
                _ = fi18
            }
            for fi19 := 0; fi19 < 5; fi19++ {
                var v20 int64 = 0
                _ = fi19
                _ = v20
            }
        }
    case "a":
        var v21 uint8 = 0
        _ = v21
    }
    if (4294967295 <= 800) {
        v2 = v2
    } else {
        fmt.Println(v2)
    }
    v2 = v2
    for fi22 := 0; fi22 < 11; fi22++ {
        v23 := GenS2{F0: true, F1: -45}
        switch (fi22 + fi22) {
        case -699:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case 9223372036854775807:
            var v24 uint = 2
            _ = v24
        }
        _ = fi22
        _ = v23
    }
    switch v2 {
    case "foo bar":
        if (1 != 42) {
            for fi25 := 0; fi25 < 15; fi25++ {
                for fi26 := 0; fi26 < 11; fi26++ {
                    switch 127 {
                    case 940:
                        if true {
                            v1[fi25] = (v2 + (v2 + v2))
                        }
                    case 1:
                        if false {
                            if false {
                                v27 := GenS5{F0: 4294967295, F1: 2, F2: (-959 >= 0)}
                                fmt.Println(v2)
                                _ = v27
                            } else {
                                delete(v1, fi25)
                                fmt.Println(fi26, fi25)
                            }
                            if true {
                                var v28 int32 = 1
                                v2 = "a"
                                v29 := GenS6{F0: (fi26 % (((fi26 | fi25)) | 1))}
                                _ = v28
                                _ = v29
                            }
                        }
                    case -332:
                        switch v2 {
                        case "\n\t":
                            fmt.Println(fi25)
                        case "foo bar":
                            var v30 uint = 1
                            _ = v30
                        }
                    default:
                        var v31 uint32 = 42
                        _ = v31
                    }
                    v2 = v2
                    _ = fi26
                }
                _ = fi25
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if false {
                v2 = v2
            } else {
                var v32 uint = 42
                fmt.Println(v2, v2)
                _ = v32
            }
        } else {
            if (false || (255 < 919)) {
                v2 = "a"
                v33 := GenS4{F0: 255, F1: 0}
                v2 = ("\n\t" + (v2 + v2))
                _ = v33
            } else {
                v2 = "a"
            }
            var v34 uint64 = 2
            _ = v34
        }
    case "a":
        if (true && (!false)) {
            if true {
                if ((!true) || ((-2.25 > (1e10 - ((((float64((3.14159 + 3.14159)) - (0.0 / -2.25)) / 1.0) / 3.14159) + 0.1))) && false)) {
                    _ = len(v1)
                    _ = len(v3)
                }
                if true {
                    var v35 int = 32768
                    _ = v35
                }
                if (0 == 105) {
                    v4 = append(v4, 2)
                    _ = len(v4)
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                } else {
                    v36 := GenS4{F0: 42, F1: 2}
                    _ = v36
                }
            } else {
                if ((42 > 2) || true) {
                    var v37 int = 42
                    delete(v1, v37)
                    v37 = v37
                    _ = v37
                } else {
                    v2 = v2
                    for fi38 := 0; fi38 < 4; fi38++ {
                        var v39 int64 = -9223372036854775808
                        _ = fi38
                        _ = v39
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
            }
        } else {
            if (-1 >= 42) {
                if true {
                    if true {
                        fmt.Println(v2)
                    }
                } else {
                    v4 = append(v4, 2)
                    for fi40 := 0; fi40 < 18; fi40++ {
                        fmt.Println(fi40, v2)
                        _ = fi40
                    }
                }
                v2 = v2
            } else {
                var v41 int8 = -128
                v42 := v1[206]
                var v43 uint64 = 878
                _ = v41
                _ = v42
                _ = v43
            }
        }
    }
    switch 127 {
    case 458:
        switch "a" {
        case "a":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case "hello":
            if true {
                var v44 int8 = 100
                v45 := GenS6{F0: -586}
                var v46 int = -175
                _ = v44
                _ = v45
                _ = v46
            } else {
                if ((!(9223372036854775807 <= -257)) || false) {
                    for fi47 := 0; fi47 < 7; fi47++ {
                        v2 = v2
                        for fi48 := 0; fi48 < 3; fi48++ {
                            _ = len(v4)
                            v2 = v2
                            var v49 string = v2
                            _ = fi48
                            _ = v49
                        }
                        var v50 int32 = -2147483648
                        _ = fi47
                        _ = v50
                    }
                } else {
                    var v51 int = -249
                    fmt.Println(v2, v2)
                    var v52 int64 = 0
                    _ = v51
                    _ = v52
                }
                for fi53 := 0; fi53 < 9; fi53++ {
                    v2 = ("hello" + ("a" + "hello"))
                    switch fi53 {
                    case 4294967295:
                        switch fi53 {
                        case -2147483648:
                            if (166 >= 4294967295) {
                                var v54 float32 = 0.0
                                switch v2 {
                                case "":
                                    v2 = ("hello" + v2)
                                }
                                _ = v54
                            }
                        case 32768:
                            v2 = v2
                        }
                    }
                    for fi55 := 0; fi55 < 10; fi55++ {
                        var v56 int32 = 2147483647
                        switch fi55 {
                        case -2147483648:
                            v3 = append(v3, "foo bar")
                        case -9223372036854775808:
                            if (4294967295 == 0) {
                                var v57 bool = true
                                _ = v57
                            }
                        case 32767:
                            var v58 int = (fi53 & -1)
                            _ = v58
                        default:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        _ = fi55
                        _ = v56
                    }
                    _ = fi53
                }
            }
        }
    }
    v1[32768] = ("\n\t" + v2)
    var v59 float32 = float32(float64(1e10))
    var v60 bool = true
    v59 += ((((-2.25 - (float32(0.0) - ((v59 + v59) / (float32(1e10) + v59)))) + (v59 * 1.5)) + float32(3.14159)) / ((v59 - (1e10 - v59)) * v59))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v60
    _ = v8
}
