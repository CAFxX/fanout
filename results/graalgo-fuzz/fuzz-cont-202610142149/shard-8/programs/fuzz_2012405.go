package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint64; F1 uint32 }
type GenS6 struct { F0 uint; F1 int32; F2 int64 }
type GenS7 struct { F0 uint64; F1 int64; F2 int16 }
func genFn0(p0 uint64, p1 uint16, p2 int) int32 {
    if (p2 > int(p0)) {
        switch p2 {
        case 4294967295:
            var v7 int64 = -820
            _ = v7
        }
    }
    v8 := []int{(p2 / ((p2) | 1)), (p2 ^ -1), (p2 << ((2) & 63))}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 534
}
func genFn1(p0 uint16, p1 int16, p2 uint8) bool {
    p2 = p2
    fmt.Println(p2, p1)
    if true {
        var v9 int = 2147483647
        if (("\n\t" + "foo bar") != "\n\t") {
            p2 -= p2
        } else {
            p2 = 42
        }
        switch v9 {
        case 32768:
            var v10 uint16 = p0
            _ = v10
        case -329:
            for fi11 := 0; fi11 < 6; fi11++ {
                var v12 float64 = (0.1 * float64(1e10))
                _ = fi11
                _ = v12
            }
        }
        _ = v9
    } else {
        for fi13 := 0; fi13 < 5; fi13++ {
            var v14 uint32 = 244
            _ = fi13
            _ = v14
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return (p0 != 42)
}

func main() {
    var v1 uint64 = 4294967295
    var v2 float64 = -2.25
    v3 := make(map[string]uint32)
    var v4 uint16 = 42
    var v5 uint = 65535
    v6 := []uint32{2147483647, 724, uint32(v1)}
    v6 = append(v6, 744)
    for fi15 := 0; fi15 < 4; fi15++ {
        if true {
            v16 := GenS1{F0: -1, F1: (v1 & (^(v1))), F2: 2}
            for fi17 := 0; fi17 < 16; fi17++ {
                var v18 uint = v5
                switch (fi17 >> ((v5) & 63)) {
                case 4294967295:
                    var v19 int8 = 76
                    _ = v19
                default:
                    var v20 float64 = 3.14159
                    _ = v20
                }
                v4 -= 0
                _ = fi17
                _ = v18
            }
            switch (fi15 % ((1) | 1)) {
            case -9223372036854775808:
                var v21 string = ("" + ("hello" + "a"))
                _ = v21
            case 32767:
                if false {
                    for fi22 := 0; fi22 < 17; fi22++ {
                        var v23 uint64 = v1
                        var v24 int32 = -399
                        switch "a" {
                        case "a":
                            _ = len(v3)
                        case "\n\t":
                            var v25 uint64 = v23
                            _ = v25
                        }
                        if true {
                            fmt.Println(v24)
                            var v26 int8 = 100
                            _ = v26
                        } else {
                            v1 = (v1 << (((v5 >> ((v5) & 63))) & 63))
                            var v27 bool = false
                            _ = v27
                        }
                        if (fi22 != int(v4)) {
                            v28 := GenS1{F0: -32768, F1: 4294967295, F2: (v1 & 674)}
                            fmt.Println(v24, v23)
                            v24 += (v24 | -957)
                            _ = v28
                        }
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    if (221 > 2147483647) {
                        fmt.Println(v4, v4, v2)
                        v29 := GenS2{F0: ("hello" + "")}
                        _ = v29
                    } else {
                        delete(v3, "")
                    }
                } else {
                    switch (("foo bar" + (("\n\t" + "foo bar") + ("foo bar" + ""))) + "\n\t") {
                    case "\n\t":
                        if true {
                            for fi30 := 0; fi30 < 16; fi30++ {
                                v31 := GenS3{F0: 2}
                                var v32 uint8 = 1
                                _ = fi30
                                _ = v31
                                _ = v32
                            }
                        } else {
                            switch "a" {
                            case "foo bar":
                                if true {
                                    for fi33 := 0; fi33 < 17; fi33++ {
                                        v34 := GenS3{F0: 216}
                                        var v35 int8 = 19
                                        v1 = v1
                                        var v36 float64 = ((v2 / 0.1) + v2)
                                        _ = fi33
                                        _ = v34
                                        _ = v35
                                        _ = v36
                                    }
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                } else {
                                    v4 = v4
                                    fmt.Println(v1)
                                    var v37 int32 = 0
                                    _ = v37
                                }
                            }
                            for fi38 := 0; fi38 < 11; fi38++ {
                                switch ("hello" + "") {
                                case "hello":
                                    v5 = (^(v5))
                                }
                                _ = fi38
                            }
                            v39 := GenS4{F0: (v5 << ((v5) & 63))}
                            _ = v39
                        }
                    }
                }
            case 42:
                v40 := GenS5{F0: 42, F1: 4294967295}
                _ = v40
            default:
                delete(v3, "foo bar")
            }
            _ = v16
        } else {
            v41 := GenS3{F0: 244}
            _ = v41
        }
        v4 = v4
        for fi42 := 0; fi42 < 16; fi42++ {
            var v43 int8 = 0
            if ((-2.25 + 1e10) >= 1.5) {
                var v44 float64 = 1.5
                v45 := v3[("hello" + "foo bar")]
                _ = len(v3)
                _ = v44
                _ = v45
            }
            _ = fi42
            _ = v43
        }
        _ = fi15
    }
    var v46 int8 = 63
    fmt.Println(v5, v5, v4)
    for fi47 := 0; fi47 < 8; fi47++ {
        var v48 bool = (!true)
        v1 = (v1 / ((2) | 1))
        _ = fi47
        _ = v48
    }
    if (!true) {
        v2 = 1e10
        v4 -= (v4 % ((v4) | 1))
        for fi49 := 0; fi49 < 13; fi49++ {
            v1 = (v1 ^ v1)
            v6 = append(v6, 247)
            v50 := v3[("hello" + ("a" + ""))]
            _ = fi49
            _ = v50
        }
    } else {
        var v51 string = ("hello" + "a")
        v4 = v4
        if false {
            v2 += float64(3.14159)
            v52 := GenS6{F0: uint(v4), F1: -697, F2: int64(v5)}
            _ = v52
        }
        _ = v51
    }
    var v53 uint32 = 1
    v54 := GenS7{F0: uint64(v53), F1: int64(v4), F2: 741}
    v1 = uint64(v53)
    v46 = (^(v46))
    v6 = append(v6, v53)
    fmt.Println(v4)
    v6 = append(v6, (v53 & (v53 & (v53 * 2147483648))))
    v55 := v3[("" + ("" + "a"))]
    fmt.Println(v55)
    v53 = (v55 << (((v5 & v5)) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v55:", v55)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v53
    _ = v54
    _ = v55
    _ = v6
}
