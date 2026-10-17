package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 string; F1 int64; F2 int8 }
type GenS5 struct { F0 uint16; F1 uint16 }
type GenS6 struct { F0 int16; F1 uint16 }
type GenS7 struct { F0 int32 }
func genFn0() string {
    var v5 string = "hello"
    var v6 int32 = 1
    v7 := []int{}
    v6 -= (v6 / ((v6) | 1))
    var v8 uint16 = 255
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return v5
}

func main() {
    var v1 string = "foo bar"
    v2 := []string{(("\n\t" + (v1 + "")) + "foo bar"), (v1 + v1), "a"}
    v3 := []string{"a", "foo bar"}
    v4 := []uint32{65535}
    v3 = append(v3, (v1 + v1))
    v9 := GenS1{F0: 2147483648, F1: 1}
    v10 := GenS2{F0: 1}
    v11 := GenS2{F0: 154}
    var v12 string = v1
    for fi13 := 0; fi13 < 15; fi13++ {
        if (!(367 < 505)) {
            v12 = v1
        } else {
            if true {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            } else {
                v14 := GenS3{F0: 2, F1: 42}
                v12 = "foo bar"
                v12 = ("hello" + v1)
                _ = v14
            }
        }
        fmt.Println(fi13, fi13)
        switch v1 {
        case "hello":
            if true {
                for fi15 := 0; fi15 < 8; fi15++ {
                    if (!false) {
                        if (0 >= -1) {
                            var v16 float32 = (3.14159 - (1e10 + (0.1 / 1.5)))
                            if (42 <= 22) {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                            v17 := GenS3{F0: 65535, F1: 42}
                            _ = v16
                            _ = v17
                        } else {
                            fmt.Println(fi13, fi15)
                            if false {
                                var v18 string = ("a" + ("a" + ""))
                                switch ("\n\t" + (("foo bar" + "foo bar") + "\n\t")) {
                                case "":
                                    fmt.Println(fi15, fi13, v12)
                                default:
                                    v19 := GenS4{F0: (v12 + "a"), F1: -104, F2: -27}
                                    _ = v19
                                }
                                _ = v18
                            } else {
                                if true {
                                    var v20 int = fi15
                                    _ = v20
                                }
                                v1 = v1
                                v12 = v1
                            }
                            if true {
                                switch v1 {
                                case "a":
                                    switch ("hello" + v1) {
                                    case "hello":
                                        _ = len(v3)
                                    case "a":
                                        if (false || true) {
                                            v12 = v1
                                        } else {
                                            v4 = append(v4, 602)
                                            v1 = (v1 + "a")
                                            var v21 float32 = 0.1
                                            _ = v21
                                        }
                                    case "":
                                        v22 := GenS5{F0: 376, F1: 65535}
                                        _ = v22
                                    }
                                case "\n\t":
                                    if (-2147483648 != 1) {
                                        var v23 int64 = 1
                                        _ = v23
                                    } else {
                                        fmt.Println(v1, v12)
                                    }
                                }
                                v1 = ("\n\t" + "hello")
                                fmt.Println(v12, fi13, v1)
                            }
                            var v24 uint32 = 2147483647
                            if true {
                                var v25 int64 = 1
                                switch v1 {
                                case "\n\t":
                                    if true {
                                        v24 = (^(v24))
                                    } else {
                                        fmt.Println(v1, fi13)
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                default:
                                    fmt.Println(v1, v12, v25)
                                }
                                if len(v4) > 0 {
                                    v4[0] = (v24 ^ v24)
                                }
                                _ = v25
                            }
                            _ = v24
                        }
                    } else {
                        if len(v4) > 0 {
                            v4[0] = 190
                        }
                        v26 := GenS2{F0: 127}
                        v1 = v1
                        if true {
                            v27 := GenS4{F0: v12, F1: -793, F2: -1}
                            switch "\n\t" {
                            case "\n\t":
                                if (42 >= 1) {
                                    v28 := GenS6{F0: 858, F1: 1}
                                    var v29 int64 = -229
                                    _ = v28
                                    _ = v29
                                }
                            }
                            _ = v27
                        }
                        var v30 string = "a"
                        _ = v26
                        _ = v30
                    }
                    _ = fi15
                }
                var v31 uint8 = 128
                fmt.Println(v1, v1, v1)
                _ = v31
            }
        case "":
            v2 = append(v2, "foo bar")
        default:
            var v32 int16 = 0
            _ = v32
        }
        _ = fi13
    }
    if (false || false) {
        v12 = v1
        for fi33 := 0; fi33 < 3; fi33++ {
            v34 := GenS4{F0: "", F1: -227, F2: -1}
            for fi35 := 0; fi35 < 20; fi35++ {
                v36 := GenS3{F0: 2, F1: 390}
                _ = fi35
                _ = v36
            }
            _ = fi33
            _ = v34
        }
    } else {
        for fi37 := 0; fi37 < 20; fi37++ {
            fmt.Println(fi37, v1)
            _ = fi37
        }
        for fi38 := 0; fi38 < 10; fi38++ {
            var v39 int32 = -69
            _ = fi38
            _ = v39
        }
        fmt.Println(v12, v1)
    }
    v1 = v12
    v12 = ("" + ("" + "foo bar"))
    v12 = ""
    var v40 int64 = 808
    var v41 int16 = 733
    _ = len(v3)
    for fi42 := 0; fi42 < 9; fi42++ {
        for fi43 := 0; fi43 < 10; fi43++ {
            v1 = v1
            _ = fi43
        }
        if (127 != 128) {
            v41 = v41
            v2 = append(v2, (v1 + ("hello" + (((v1 + ("hello" + "\n\t")) + "foo bar") + v1))))
            v44 := GenS1{F0: 0, F1: 42}
            _ = v44
        } else {
            for fi45 := 0; fi45 < 15; fi45++ {
                v40 += (v40 - v40)
                _ = fi45
            }
            v46 := GenS3{F0: 2, F1: (^(v40))}
            _ = v46
        }
        for fi47 := 0; fi47 < 17; fi47++ {
            v48 := GenS7{F0: 1}
            var v49 string = ""
            if (!(true && ((float64((float32(1.5) + 0.0)) <= float64(((0.0 * float32((-2.25 - 3.14159))) * 1e10))) || (v41 != v41)))) {
                var v50 int16 = v41
                _ = v50
            } else {
                if false {
                    if (-560 <= 42) {
                        if (18446744073709551615 < 2) {
                            v41 -= (v41 >> ((0) & 63))
                            v12 = "hello"
                            var v51 float32 = 0.0
                            fmt.Println(v41, v1, v12)
                            _ = v51
                        }
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    } else {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    v12 = ""
                } else {
                    if len(v2) > 0 {
                        v2[0] = "hello"
                    }
                    var v52 uint16 = 32767
                    v41 -= (-(v41))
                    _ = v52
                }
            }
            var v53 int64 = (v40 % ((v40) | 1))
            var v54 int64 = (v40 ^ (v53 * (v53 / ((v40) | 1))))
            _ = fi47
            _ = v48
            _ = v49
            _ = v53
            _ = v54
        }
        _ = fi42
    }
    fmt.Println(v12)
    for fi55 := 0; fi55 < 14; fi55++ {
        fmt.Println(v41, v1, v12)
        v56 := GenS7{F0: 42}
        _ = fi55
        _ = v56
    }
    v57 := GenS3{F0: 535, F1: v40}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v57
    _ = v9
}
