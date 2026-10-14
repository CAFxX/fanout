package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 bool }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 bool; F1 uint64; F2 uint }
type GenS4 struct { F0 int16; F1 int16 }
type GenS5 struct { F0 int64; F1 int16 }
type GenS6 struct { F0 uint64; F1 bool }
type GenS7 struct { F0 int8; F1 bool; F2 int32 }
func genFn0(p0 string, p1 int64) uint16 {
    var v6 int32 = -2147483648
    for fi7 := 0; fi7 < 11; fi7++ {
        fmt.Println(v6)
        p0 = ((("hello" + ("\n\t" + "foo bar")) + "") + p0)
        _ = fi7
    }
    var v8 uint16 = 32767
    v9 := GenS1{F0: -836, F1: 0, F2: (42 > 1)}
    v10 := GenS2{F0: 1, F1: false}
    _ = p0
    _ = p1
    _ = v10
    _ = v6
    _ = v8
    _ = v9
    return v8
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float64 = ((0.1 - ((1.5 * (((0.0 / (1e10 + (-2.25 * (float64((0.0 + 3.14159)) * 1e10)))) + 0.0) + 1.5)) + -2.25)) + 3.14159)
    v3 := make(map[uint32]int)
    v4 := []int{4294967295, -874}
    v5 := make(map[uint32]string)
    v11 := v3[2]
    fmt.Println(v2, v11)
    v2 = v2
    v11 = 1
    if true {
        v12 := GenS3{F0: (true || (38 != 255)), F1: 2, F2: 4294967295}
        switch v11 {
        case -9223372036854775808:
            v5[42] = ("" + "\n\t")
        case 9223372036854775807:
            if false {
                var v13 float64 = v2
                v13 = ((v2 + v13) * 0.0)
                _ = v13
            }
        case 42:
            v14 := GenS4{F0: 0, F1: 1000}
            _ = v14
        }
        _ = v12
    }
    for fi15 := 0; fi15 < 17; fi15++ {
        var v16 int8 = -99
        v17 := v5[4294967295]
        switch fi15 {
        case 432:
            for fi18 := 0; fi18 < 4; fi18++ {
                var v19 int16 = -32768
                if false {
                    v19 += v19
                    v1[(fi15 - 444)] = 420
                } else {
                    if len(v4) > 0 {
                        v4[0] = int(v16)
                    }
                }
                _ = fi18
                _ = v19
            }
        case -947:
            switch v17 {
            case "foo bar":
                if ((2 <= 927) || true) {
                    v20 := GenS5{F0: -922, F1: 0}
                    switch 32768 {
                    case 32768:
                        for fi21 := 0; fi21 < 9; fi21++ {
                            if len(v4) > 0 {
                                v4[0] = -1
                            }
                            if ((-2.25 + 0.1) <= (3.14159 - 0.0)) {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                if (42 < 32767) {
                                    if (!(v17 >= ((v17 + v17) + ("" + "\n\t")))) {
                                        var v22 int32 = -1
                                        if true {
                                            v11 = (fi21 >> ((466) & 63))
                                            if len(v4) > 0 {
                                                v4[0] = int(v22)
                                            }
                                            var v23 int32 = (v22 / (((v22 & (v22 | v22))) | 1))
                                            _ = v23
                                        } else {
                                            v2 = (-2.25 + v2)
                                            var v24 bool = false
                                            _ = v24
                                        }
                                        _ = v22
                                    }
                                    v25 := GenS6{F0: 18446744073709551615, F1: (true && true)}
                                    var v26 uint32 = 65535
                                    _ = v25
                                    _ = v26
                                } else {
                                    v17 = (v17 + "")
                                    var v27 bool = (fi21 == fi21)
                                    var v28 float64 = 1.5
                                    switch v17 {
                                    case "a":
                                        if len(v4) > 0 {
                                            v4[0] = fi15
                                        }
                                    case "foo bar":
                                        v29 := GenS2{F0: 2, F1: true}
                                        _ = v29
                                    }
                                    _ = v27
                                    _ = v28
                                }
                                v16 -= 87
                            }
                            v16 -= v16
                            _ = fi21
                        }
                    case -9223372036854775808:
                        fmt.Println(fi15, v17, v11)
                    case 127:
                        v30 := GenS7{F0: (v16 - (v16 * (v16 / ((v16) | 1)))), F1: true, F2: -163}
                        _ = v30
                    }
                    _ = v20
                } else {
                    var v31 int64 = 891
                    _ = v31
                }
            case "hello":
                switch v17 {
                case "":
                    var v32 int16 = 1000
                    _ = v32
                case "\n\t":
                    if false {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        fmt.Println(fi15, v11, fi15)
                        var v33 uint = 596
                        _ = v33
                    }
                default:
                    v17 = "hello"
                }
            }
        }
        _ = fi15
        _ = v16
        _ = v17
    }
    if false {
        fmt.Println(v2)
    } else {
        fmt.Println(v2)
        var v34 float32 = (1.5 / 1.5)
        _ = v34
    }
    if true {
        var v35 float64 = -2.25
        v2 += float64(float32(v35))
        var v36 string = ("foo bar" + "\n\t")
        _ = v35
        _ = v36
    }
    v2 = (v2 / 3.14159)
    fmt.Println(v2)
    var v37 string = "hello"
    v1[-9223372036854775808] = 2147483648
    fmt.Println(v2, v11)
    if (255 == 163) {
        v38 := GenS4{F0: 1000, F1: -1}
        var v39 uint16 = 65535
        if (!(205 >= 255)) {
            _ = len(v4)
            v39 = v39
        } else {
            if (v2 >= (1.5 - (v2 + (v2 + v2)))) {
                v39 = v39
                if (!false) {
                    var v40 int64 = 9223372036854775807
                    for fi41 := 0; fi41 < 6; fi41++ {
                        if true {
                            var v42 int32 = 639
                            _ = v42
                        }
                        fmt.Println(v39)
                        _ = fi41
                    }
                    _ = v40
                }
            }
            for fi43 := 0; fi43 < 18; fi43++ {
                v2 += v2
                _ = fi43
            }
            fmt.Println(v2)
        }
        _ = v38
        _ = v39
    } else {
        v37 = v37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
}
