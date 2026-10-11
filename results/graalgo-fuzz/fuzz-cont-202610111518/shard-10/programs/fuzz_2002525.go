package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 string }
type GenS2 struct { F0 uint16; F1 int16; F2 bool }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 int32; F1 int32; F2 int8 }
type GenS5 struct { F0 uint32 }

func main() {
    var v1 uint64 = 42
    var v2 float64 = (0.0 - (1e10 - (1.5 + float64(0.1))))
    var v3 int8 = 127
    v4 := []int{4294967295}
    v5 := []uint32{2147483648, 42}
    var v6 uint32 = 42
    var v7 uint32 = v6
    _ = len(v4)
    v8 := GenS1{F0: 228, F1: 65535, F2: ((("\n\t" + "foo bar") + "") + ("hello" + ((((("foo bar" + "\n\t") + "foo bar") + "foo bar") + ("foo bar" + ("a" + ("a" + "a")))) + "a")))}
    if (("foo bar" + ("" + "foo bar")) > "a") {
        var v9 int64 = -75
        switch 128 {
        case -1:
            if false {
                fmt.Println(v2, v2, v2)
                for fi10 := 0; fi10 < 15; fi10++ {
                    v1 = v1
                    var v11 uint = 2
                    _ = fi10
                    _ = v11
                }
                v12 := GenS1{F0: 859, F1: 255, F2: ("" + "\n\t")}
                _ = v12
            } else {
                v4 = append(v4, 32768)
            }
        }
        if ((float32(v2) * -2.25) != 3.14159) {
            v2 -= 0.1
        } else {
            var v13 uint32 = v6
            _ = v13
        }
        _ = v9
    } else {
        var v14 int64 = 0
        _ = v14
    }
    if false {
        if (1000 >= int16(v1)) {
            for fi15 := 0; fi15 < 10; fi15++ {
                fmt.Println(v7, v2)
                _ = fi15
            }
        } else {
            for fi16 := 0; fi16 < 19; fi16++ {
                if false {
                    var v17 uint32 = (v6 / ((uint32(v1)) | 1))
                    var v18 int16 = 42
                    _ = v17
                    _ = v18
                }
                _ = fi16
            }
        }
        v19 := GenS1{F0: 42, F1: 1, F2: ("hello" + ("foo bar" + ("foo bar" + "")))}
        for fi20 := 0; fi20 < 2; fi20++ {
            v3 = v3
            if false {
                for fi21 := 0; fi21 < 9; fi21++ {
                    v22 := GenS1{F0: 65535, F1: 1, F2: ("foo bar" + (("\n\t" + "\n\t") + "hello"))}
                    switch "" {
                    case "":
                        if len(v5) > 0 {
                            v5[0] = v7
                        }
                    case "foo bar":
                        var v23 uint64 = 859
                        _ = v23
                    case "a":
                        v3 += -128
                    default:
                        v2 -= (v2 + (v2 + -2.25))
                    }
                    var v24 float64 = float64(0.0)
                    v3 = (v3 ^ -64)
                    fmt.Println(v3, v1, v6)
                    _ = fi21
                    _ = v22
                    _ = v24
                }
                var v25 string = "\n\t"
                v7 = v6
                _ = v25
            }
            _ = len(v4)
            _ = fi20
        }
        _ = v19
    } else {
        switch ("a" + "hello") {
        case "hello":
            if true {
                _ = len(v4)
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
        case "":
            var v26 bool = (42 >= 1)
            _ = v26
        case "foo bar":
            for fi27 := 0; fi27 < 11; fi27++ {
                for fi28 := 0; fi28 < 6; fi28++ {
                    v29 := GenS2{F0: 0, F1: -1, F2: (int64(v6) == -1)}
                    _ = fi28
                    _ = v29
                }
                _ = fi27
            }
        }
        v2 += 0.0
    }
    var v30 int32 = -1
    for fi31 := 0; fi31 < 19; fi31++ {
        fmt.Println(v1, v6)
        v7 -= v7
        _ = fi31
    }
    for fi32 := 0; fi32 < 13; fi32++ {
        var v33 int64 = -510
        for fi34 := 0; fi34 < 16; fi34++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v35 := GenS2{F0: 244, F1: -32768, F2: false}
            if true {
                var v36 string = ("a" + ("hello" + "a"))
                var v37 uint8 = uint8(v6)
                if ((v30 != v30) || (v3 > (v3 | (v3 & v3)))) {
                    v38 := GenS3{F0: 18446744073709551615, F1: 0}
                    fmt.Println(fi34, v2)
                    if (!(true || ((fi34 & 1) <= fi34))) {
                        _ = len(v5)
                        if true {
                            v5 = append(v5, v7)
                        } else {
                            var v39 uint = 0
                            _ = len(v4)
                            _ = v39
                        }
                        if (255 != 0) {
                            var v40 float64 = 0.1
                            var v41 uint64 = 42
                            _ = v40
                            _ = v41
                        }
                        if true {
                            var v42 uint = 4294967295
                            var v43 uint32 = v6
                            v37 -= 128
                            _ = v42
                            _ = v43
                        }
                    } else {
                        v44 := GenS1{F0: 644, F1: 527, F2: v36}
                        if (!true) {
                            v7 += 1
                            v33 = (v33 - v33)
                            var v45 int64 = (v33 | (-(v33)))
                            _ = v45
                        } else {
                            fmt.Println(v33, fi32)
                            switch ((v36 + "hello") + ("a" + v36)) {
                            case "":
                                v7 += (v7 % ((v7) | 1))
                            }
                            fmt.Println(v6, v33)
                        }
                        _ = v44
                    }
                    _ = v38
                } else {
                    switch ("a" + ("foo bar" + v36)) {
                    case "\n\t":
                        v46 := GenS3{F0: 18446744073709551615, F1: (^(v33))}
                        _ = v46
                    default:
                        if true {
                            if false {
                                fmt.Println(v3)
                                v47 := GenS4{F0: (v30 - v30), F1: (v30 >> ((2) & 63)), F2: 42}
                                _ = v47
                            } else {
                                v1 = v1
                                if true {
                                    if (v1 < 0) {
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    } else {
                                        var v48 uint64 = v1
                                        _ = v48
                                    }
                                } else {
                                    var v49 uint16 = 266
                                    _ = v49
                                }
                                var v50 uint = 65535
                                v51 := GenS3{F0: uint(v37), F1: v33}
                                _ = v50
                                _ = v51
                            }
                            v52 := GenS3{F0: 2, F1: v33}
                            v5 = append(v5, uint32(v37))
                            switch (fi34 - (fi32 / ((-1) | 1))) {
                            case 42:
                                fmt.Println(v3, v6)
                            case 0:
                                var v53 uint8 = (v37 >> ((18446744073709551615) & 63))
                                _ = v53
                            }
                            v54 := GenS4{F0: (-(v30)), F1: 42, F2: -62}
                            _ = v52
                            _ = v54
                        } else {
                            if true {
                                var v55 float32 = 0.1
                                fmt.Println(v2, fi34)
                                v56 := GenS5{F0: v6}
                                _ = v55
                                _ = v56
                            }
                            fmt.Println(v1, v37, v6)
                            _ = len(v5)
                        }
                    }
                    var v57 uint = 18446744073709551615
                    v58 := GenS3{F0: 4294967295, F1: int64(v1)}
                    _ = v57
                    _ = v58
                }
                _ = v36
                _ = v37
            }
            _ = fi34
            _ = v35
        }
        fmt.Println(v3)
        _ = fi32
        _ = v33
    }
    var v59 int16 = -1
    v30 -= v30
    if len(v5) > 0 {
        v5[0] = v6
    }
    var v60 float64 = 1.5
    v60 = (float64(1.5) + 0.1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v59
    _ = v6
    _ = v60
    _ = v7
    _ = v8
}
