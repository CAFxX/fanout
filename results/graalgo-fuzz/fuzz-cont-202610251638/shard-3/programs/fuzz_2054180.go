package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8; F1 int64; F2 int16 }
func genFn0(p0 uint16, p1 string, p2 int8) int32 {
    v4 := []int{int(p2), 318, 4294967295}
    for fi5 := 0; fi5 < 9; fi5++ {
        for fi6 := 0; fi6 < 19; fi6++ {
            fmt.Println(p0, p1)
            _ = fi6
        }
        var v7 int = (fi5 >> ((2) & 63))
        _ = fi5
        _ = v7
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v8 uint = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v8
    return int32(v8)
}
func genFn1(p0 bool, p1 int64, p2 int8) int32 {
    v9 := GenS1{F0: 1}
    v10 := GenS1{F0: 65535}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return genFn0(795, "foo bar", (p2 | p2))
}

func main() {
    var v1 uint64 = 2
    v2 := make(map[string]bool)
    var v3 float64 = 1e10
    var v11 int64 = -1
    v1 = 769
    v1 += (v1 ^ (^(v1)))
    if (-2.25 <= v3) {
        if (true || true) {
            var v12 float32 = float32((float64((3.14159 - 0.1)) * ((v3 / ((v3 - 1e10) / ((3.14159 - float64(3.14159)) / float64(float32(v3))))) - 3.14159)))
            var v13 uint8 = 128
            for fi14 := 0; fi14 < 19; fi14++ {
                v2["a"] = (!true)
                _ = fi14
            }
            _ = v12
            _ = v13
        }
    }
    v11 = 0
    if false {
        for fi15 := 0; fi15 < 18; fi15++ {
            v16 := GenS1{F0: 969}
            _ = fi15
            _ = v16
        }
    }
    v17 := GenS1{F0: 633}
    for fi18 := 0; fi18 < 1; fi18++ {
        v3 += float64(1.5)
        _ = len(v2)
        _ = fi18
    }
    fmt.Println(v3, v3, v11)
    switch -2147483648 {
    case -307:
        v19 := []int{-227, 1}
        _ = v19
    case -9223372036854775808:
        if (false || (false && true)) {
            var v20 float32 = 1.5
            _ = v20
        } else {
            var v21 int16 = -927
            v1 += (v1 & (v1 << ((1) & 63)))
            _ = v21
        }
    case 32767:
        v22 := v2["a"]
        _ = v22
    }
    v23 := []int{-1}
    if true {
        v24 := GenS2{F0: -32768}
        delete(v2, (("foo bar" + "") + "a"))
        _ = v24
    } else {
        if (!(!(!true))) {
            if false {
                switch 0 {
                case 32768:
                    v23 = append(v23, 988)
                case 0:
                    for fi25 := 0; fi25 < 12; fi25++ {
                        var v26 bool = (65535 < 2147483647)
                        v26 = false
                        _ = fi25
                        _ = v26
                    }
                case 32767:
                    switch 1 {
                    case -716:
                        if true {
                            var v27 int64 = v11
                            switch "hello" {
                            case "a":
                                switch 127 {
                                case 1:
                                    var v28 float32 = ((0.1 + -2.25) * 0.1)
                                    _ = v28
                                case 4294967295:
                                    var v29 uint16 = 746
                                    _ = v29
                                case 24:
                                    for fi30 := 0; fi30 < 11; fi30++ {
                                        v11 -= v27
                                        var v31 uint32 = 1
                                        v27 = 118
                                        _ = fi30
                                        _ = v31
                                    }
                                default:
                                    v3 += float64(3.14159)
                                }
                            case "":
                                fmt.Println(v1)
                            case "hello":
                                for fi32 := 0; fi32 < 13; fi32++ {
                                    fmt.Println(v27, v27)
                                    var v33 uint32 = 65535
                                    fmt.Println(v3, fi32)
                                    _ = fi32
                                    _ = v33
                                }
                            }
                            _ = v27
                        } else {
                            v1 = v1
                            v2["\n\t"] = false
                        }
                    case 1:
                        for fi34 := 0; fi34 < 18; fi34++ {
                            v3 -= 1e10
                            _ = fi34
                        }
                    case 127:
                        v11 = -520
                    default:
                        v35 := GenS2{F0: -32768}
                        _ = v35
                    }
                default:
                    fmt.Println(v3)
                }
                if (-2.25 <= 1.5) {
                    fmt.Println(v3, v11)
                } else {
                    for fi36 := 0; fi36 < 13; fi36++ {
                        var v37 int64 = v11
                        v3 = 0.0
                        _ = fi36
                        _ = v37
                    }
                    if (!true) {
                        var v38 string = "hello"
                        delete(v2, "a")
                        _ = v38
                    } else {
                        var v39 uint16 = 2
                        for fi40 := 0; fi40 < 10; fi40++ {
                            v11 += v11
                            if (false || false) {
                                var v41 uint32 = 2147483647
                                _ = v41
                            } else {
                                v3 += v3
                                delete(v2, "a")
                                var v42 string = "foo bar"
                                _ = v42
                            }
                            if false {
                                v43 := GenS2{F0: 1}
                                v44 := GenS1{F0: (v39 - (^(v39)))}
                                fmt.Println(fi40, v11, v3)
                                _ = v43
                                _ = v44
                            }
                            _ = fi40
                        }
                        var v45 uint = 255
                        _ = v39
                        _ = v45
                    }
                    v11 = (v11 ^ 9223372036854775807)
                }
                if (!false) {
                    for fi46 := 0; fi46 < 15; fi46++ {
                        switch "" {
                        case "\n\t":
                            var v47 uint16 = 397
                            _ = v47
                        case "foo bar":
                            var v48 float32 = float32(((float64((float32(float64(-2.25)) + 0.1)) - (3.14159 * v3)) - 0.0))
                            _ = v48
                        default:
                            fmt.Println(v11, fi46)
                        }
                        _ = len(v23)
                        var v49 uint32 = 2
                        _ = fi46
                        _ = v49
                    }
                    var v50 uint = 0
                    _ = v50
                }
            } else {
                v51 := GenS2{F0: -274}
                if false {
                    var v52 uint16 = 32767
                    _ = v52
                }
                var v53 float32 = ((-2.25 / 0.1) / 0.1)
                _ = v51
                _ = v53
            }
        }
        v54 := GenS3{F0: 128, F1: -673, F2: 32767}
        _ = v54
    }
    _ = len(v23)
    var v55 float32 = -2.25
    if (!true) {
        v3 += 1.5
    }
    v56 := GenS1{F0: 32767}
    fmt.Println(v1, v55, v3)
    var v57 uint16 = 473
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v55:", v55)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v11
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v55
    _ = v56
    _ = v57
}
