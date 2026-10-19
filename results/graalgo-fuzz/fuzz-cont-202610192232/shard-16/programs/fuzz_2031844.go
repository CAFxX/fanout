package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 uint; F2 uint16 }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 int8; F1 int32; F2 uint16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint64; F1 string; F2 string }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 string; F1 int16; F2 int8 }

func main() {
    v1 := []int{2147483647}
    var v2 uint = 18446744073709551615
    var v3 string = "a"
    var v4 uint64 = 45
    v5 := make(map[string]uint32)
    v3 = v3
    for fi6 := 0; fi6 < 6; fi6++ {
        if true {
            var v7 int64 = int64(fi6)
            var v8 float32 = (float32(1e10) * -2.25)
            _ = v7
            _ = v8
        } else {
            switch v3 {
            case "hello":
                if (1000 != -1) {
                    var v9 uint = (v2 + 4294967295)
                    for fi10 := 0; fi10 < 8; fi10++ {
                        var v11 int64 = -9223372036854775808
                        _ = fi10
                        _ = v11
                    }
                    v5[v3] = 1
                    _ = v9
                }
            default:
                v2 = v2
            }
        }
        _ = fi6
    }
    for fi12 := 0; fi12 < 1; fi12++ {
        switch (v3 + v3) {
        case "":
            fmt.Println(v4, v2, v4)
        }
        v2 = 18446744073709551615
        v2 -= (^(v2))
        _ = fi12
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        for fi14 := 0; fi14 < 13; fi14++ {
            v1 = append(v1, int(v2))
            switch v3 {
            case "a":
                var v15 uint8 = 255
                _ = v15
            case "foo bar":
                v2 += v2
            case "hello":
                if ((1 == -32768) || (127 == 127)) {
                    v16 := GenS1{F0: 42}
                    _ = v16
                }
            }
            v17 := GenS1{F0: 1}
            _ = fi14
            _ = v17
        }
        if ((fi13 << ((v2) & 63)) >= fi13) {
            switch (v3 + ("\n\t" + v3)) {
            case "\n\t":
                if len(v1) > 0 {
                    v1[0] = (fi13 << ((0) & 63))
                }
            case "foo bar":
                switch (fi13 / (((fi13 - fi13)) | 1)) {
                case -9223372036854775808:
                    if (v4 != v4) {
                        v3 = (v3 + (v3 + "a"))
                        var v18 uint32 = 2147483647
                        _ = v18
                    }
                }
            default:
                v2 = v2
            }
        }
        _ = fi13
    }
    var v19 uint = 876
    var v20 int8 = 1
    for fi21 := 0; fi21 < 13; fi21++ {
        if true {
            var v22 int32 = -1
            if true {
                v23 := GenS1{F0: 42}
                switch v3 {
                case "foo bar":
                    var v24 string = v3
                    _ = v24
                case "hello":
                    _ = len(v5)
                }
                _ = v23
            } else {
                var v25 float64 = (1.5 + (1.5 + 0.1))
                var v26 bool = ((128 <= 149) || (!false))
                _ = v25
                _ = v26
            }
            var v27 uint = v2
            _ = v22
            _ = v27
        } else {
            for fi28 := 0; fi28 < 1; fi28++ {
                var v29 float32 = 1.5
                v5["hello"] = 886
                _ = fi28
                _ = v29
            }
        }
        v4 += (^(v4))
        var v30 int = (fi21 << ((0) & 63))
        _ = fi21
        _ = v30
    }
    if (-1 > 0) {
        for fi31 := 0; fi31 < 9; fi31++ {
            v20 += v20
            v32 := GenS2{F0: -912, F1: v2, F2: 0}
            _ = fi31
            _ = v32
        }
        v2 += (v2 / (((v2 ^ (v2 >> (((^(v19))) & 63)))) | 1))
    } else {
        var v33 string = (v3 + v3)
        _ = v33
    }
    v34 := GenS3{F0: 127, F1: (((v3 + "") + "hello") + "\n\t")}
    var v35 uint = v2
    if len(v1) > 0 {
        v1[0] = 944
    }
    var v36 float32 = float32(float64(1e10))
    switch -9223372036854775808 {
    case 721:
        if false {
            v35 = v19
        } else {
            if (9223372036854775807 >= 1) {
                switch 1 {
                case 4294967295:
                    if ((v2 ^ v19) == v35) {
                        if (v3 <= v3) {
                            v35 = v19
                            v37 := GenS4{F0: v20, F1: -1, F2: 1}
                            _ = v37
                        } else {
                            for fi38 := 0; fi38 < 14; fi38++ {
                                v35 += (v35 + 0)
                                _ = fi38
                            }
                            var v39 float64 = ((3.14159 + 0.0) - 0.0)
                            v40 := v5[(v3 + (v3 + (v3 + ((("\n\t" + (v3 + v3)) + "") + v3))))]
                            _ = v39
                            _ = v40
                        }
                        fmt.Println(v4)
                    } else {
                        for fi41 := 0; fi41 < 2; fi41++ {
                            for fi42 := 0; fi42 < 17; fi42++ {
                                v43 := GenS5{F0: true}
                                v1 = append(v1, (fi42 << ((1) & 63)))
                                _ = fi42
                                _ = v43
                            }
                            _ = fi41
                        }
                        var v44 float32 = (float32(0.1) - 3.14159)
                        var v45 uint32 = 42
                        _ = v44
                        _ = v45
                    }
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            for fi46 := 0; fi46 < 9; fi46++ {
                for fi47 := 0; fi47 < 15; fi47++ {
                    var v48 int8 = v20
                    if true {
                        var v49 uint8 = uint8(v4)
                        var v50 string = (v3 + (v3 + v3))
                        _ = v49
                        _ = v50
                    } else {
                        v4 = 811
                        var v51 int8 = v20
                        v52 := GenS6{F0: (v3 >= "hello")}
                        var v53 int8 = v20
                        if false {
                            if false {
                                v48 -= 127
                                fmt.Println(v48)
                            } else {
                                switch "hello" {
                                case "":
                                    fmt.Println(v51)
                                case "foo bar":
                                    var v54 bool = (2 > 640)
                                    _ = v54
                                }
                                switch (("hello" + v3) + "a") {
                                case "\n\t":
                                    var v55 float64 = 3.14159
                                    _ = v55
                                }
                            }
                            var v56 float64 = ((float64(((v36 + v36) / (float32(3.14159) * v36))) * 3.14159) + float64(v36))
                            fmt.Println(v48, v51)
                            _ = v56
                        } else {
                            var v57 string = v3
                            _ = v57
                        }
                        _ = v51
                        _ = v52
                        _ = v53
                    }
                    _ = fi47
                    _ = v48
                }
                var v58 uint32 = 5
                fmt.Println(v19)
                _ = fi46
                _ = v58
            }
            for fi59 := 0; fi59 < 20; fi59++ {
                switch v3 {
                case "":
                    v36 -= v36
                default:
                    var v60 uint = 935
                    _ = v60
                }
                for fi61 := 0; fi61 < 12; fi61++ {
                    v2 = 640
                    _ = fi61
                }
                if false {
                    if (true && true) {
                        delete(v5, (((("foo bar" + v3) + v3) + (v3 + "\n\t")) + v3))
                        v2 += v2
                    }
                    v62 := GenS7{F0: v4, F1: ((((("hello" + "hello") + v3) + v3) + v3) + v3), F2: v3}
                    _ = v62
                }
                _ = fi59
            }
        }
    case 9223372036854775807:
        for fi63 := 0; fi63 < 4; fi63++ {
            switch int(v4) {
            case 42:
                v5[(v3 + v3)] = 4294967295
            }
            var v64 int16 = -1
            _ = fi63
            _ = v64
        }
    }
    v2 += (v35 & v2)
    if (787 > 1) {
        if (255 == 105) {
            fmt.Println(v36)
            v65 := GenS8{F0: 2}
            fmt.Println(v2)
            _ = v65
        } else {
            v20 -= v20
            v36 = v36
        }
        switch 877 {
        case -2147483648:
            for fi66 := 0; fi66 < 4; fi66++ {
                if (true && false) {
                    if false {
                        v67 := GenS8{F0: 32767}
                        if (false || false) {
                            if (!true) {
                                var v68 int64 = int64(v4)
                                switch v3 {
                                case "hello":
                                    v69 := GenS9{F0: ((("a" + "a") + ("\n\t" + v3)) + (((v3 + "foo bar") + v3) + "\n\t")), F1: int16(v2), F2: (v20 * v20)}
                                    _ = v69
                                case "\n\t":
                                    if len(v1) > 0 {
                                        v1[0] = fi66
                                    }
                                default:
                                    v20 = (v20 << (((v2 % (((^(v2))) | 1))) & 63))
                                }
                                _ = v68
                            } else {
                                v20 = (v20 * -125)
                            }
                        }
                        _ = v67
                    }
                    if len(v1) > 0 {
                        v1[0] = fi66
                    }
                } else {
                    v4 += 18446744073709551615
                }
                _ = fi66
            }
        }
        _ = len(v5)
    }
    var v70 int64 = 0
    v35 -= 4294967295
    _ = len(v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v70:", v70)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v70
}
