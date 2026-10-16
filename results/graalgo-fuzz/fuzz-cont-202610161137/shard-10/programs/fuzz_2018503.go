package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 uint8; F2 string }
type GenS3 struct { F0 int32; F1 uint8; F2 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint8; F1 uint8; F2 uint }
func genFn0(p0 uint8, p1 int, p2 string) int {
    p0 = (p0 << ((343) & 63))
    if (395 == 4) {
        var v5 int8 = 96
        if (!(-382 > p1)) {
            v6 := GenS1{F0: 1}
            v7 := GenS2{F0: 0, F1: p0, F2: (p2 + "hello")}
            p1 += -2147483648
            _ = v6
            _ = v7
        } else {
            var v8 uint = 4294967295
            var v9 int32 = 2147483647
            _ = v8
            _ = v9
        }
        _ = v5
    }
    _ = p0
    _ = p1
    _ = p2
    return p1
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int16 = -32768
    v3 := []uint32{978, 0, 42}
    var v4 int = 42
    v2 += (v2 | (v2 % ((-1) | 1)))
    for fi10 := 0; fi10 < 1; fi10++ {
        var v11 string = ""
        var v12 float64 = 1e10
        _ = fi10
        _ = v11
        _ = v12
    }
    switch ("\n\t" + "\n\t") {
    case "foo bar":
        for fi13 := 0; fi13 < 1; fi13++ {
            v2 = 715
            v14 := v1[-1]
            if true {
                if ((true && true) && false) {
                    var v15 uint32 = v14
                    for fi16 := 0; fi16 < 11; fi16++ {
                        v2 += (v2 << ((1) & 63))
                        _ = fi16
                    }
                    var v17 int32 = -32
                    _ = v15
                    _ = v17
                } else {
                    var v18 int8 = -43
                    v18 = v18
                    if false {
                        for fi19 := 0; fi19 < 9; fi19++ {
                            v1[(fi19 + v4)] = 4294967295
                            switch ("hello" + ((("\n\t" + "hello") + "hello") + "\n\t")) {
                            case "foo bar":
                                v18 += (v18 % ((v18) | 1))
                            }
                            if false {
                                v20 := GenS1{F0: v14}
                                _ = v20
                            } else {
                                v2 += 234
                                delete(v1, fi13)
                                fmt.Println(v2)
                            }
                            v21 := GenS3{F0: 2147483647, F1: 2, F2: 909}
                            if false {
                                v22 := v1[genFn0(127, -9223372036854775808, "hello")]
                                v4 = fi19
                                v23 := GenS3{F0: -1, F1: 84, F2: 255}
                                _ = v22
                                _ = v23
                            }
                            _ = fi19
                            _ = v21
                        }
                    } else {
                        fmt.Println(v18, v14, v4)
                        switch (("foo bar" + "hello") + "a") {
                        case "\n\t":
                            for fi24 := 0; fi24 < 5; fi24++ {
                                v4 = (fi24 % ((45) | 1))
                                _ = fi24
                            }
                        case "a":
                            var v25 int16 = v2
                            _ = v25
                        case "":
                            fmt.Println(v14, fi13)
                        default:
                            var v26 int64 = 622
                            _ = v26
                        }
                        if false {
                            var v27 int64 = -460
                            _ = v27
                        }
                    }
                    _ = v18
                }
                v4 = (v4 % (((fi13 + v4)) | 1))
            }
            _ = fi13
            _ = v14
        }
    case "\n\t":
        for fi28 := 0; fi28 < 14; fi28++ {
            v29 := GenS1{F0: 2147483648}
            v1[fi28] = 213
            _ = fi28
            _ = v29
        }
    default:
        var v30 bool = true
        _ = v30
    }
    _ = len(v3)
    v4 -= (v4 >> ((4294967295) & 63))
    for fi31 := 0; fi31 < 1; fi31++ {
        var v32 int16 = v2
        _ = fi31
        _ = v32
    }
    if false {
        for fi33 := 0; fi33 < 4; fi33++ {
            if (!(false && true)) {
                var v34 int16 = -536
                _ = v34
            }
            if true {
                v4 -= 145
                v35 := GenS3{F0: 1, F1: 1, F2: 255}
                _ = v35
            }
            _ = fi33
        }
    }
    if true {
        v4 += 1
        var v36 int32 = 0
        _ = len(v3)
        _ = v36
    }
    v2 += (v2 % ((v2) | 1))
    v4 = (v4 >> ((65535) & 63))
    if (float32(3.14159) < 0.0) {
        var v37 string = "hello"
        if false {
            v38 := v1[genFn0(1, v4, "foo bar")]
            if (!(v37 < v37)) {
                var v39 float32 = (1e10 / (((float32(0.0) * (float32((0.0 + 0.1)) / 1.5)) - (1e10 + 3.14159)) / 1.0))
                _ = v39
            }
            var v40 uint64 = 0
            _ = v38
            _ = v40
        }
        for fi41 := 0; fi41 < 8; fi41++ {
            for fi42 := 0; fi42 < 13; fi42++ {
                v4 += fi42
                v43 := GenS3{F0: -29, F1: 133, F2: 1}
                v4 = 1
                _ = fi42
                _ = v43
            }
            _ = fi41
        }
        _ = v37
    } else {
        if len(v3) > 0 {
            v3[0] = 4294967295
        }
    }
    v2 = -1
    v2 += (v2 >> ((952) & 63))
    if (!("foo bar" <= ("hello" + "\n\t"))) {
        v44 := GenS1{F0: 565}
        _ = v44
    }
    var v45 bool = (3.14159 > ((1e10 - float64(0.0)) + 1e10))
    v46 := v1[(v4 << ((870) & 63))]
    switch ("a" + "\n\t") {
    case "\n\t":
        var v47 uint8 = 180
        _ = v47
    case "":
        if (996 != 2147483647) {
            var v48 uint8 = 0
            v2 += v2
            _ = v48
        } else {
            if v45 {
                if (-381 <= -200) {
                    var v49 uint8 = 255
                    if v45 {
                        for fi50 := 0; fi50 < 3; fi50++ {
                            for fi51 := 0; fi51 < 7; fi51++ {
                                v52 := v1[-2147483648]
                                if (-128 != 97) {
                                    v53 := GenS4{F0: 42}
                                    var v54 uint = 255
                                    _ = v53
                                    _ = v54
                                } else {
                                    v52 += (^(v52))
                                    var v55 float64 = ((0.1 + 1e10) + 3.14159)
                                    if v45 {
                                        _ = len(v1)
                                        v56 := GenS2{F0: 896, F1: v49, F2: "foo bar"}
                                        switch 128 {
                                        case -192:
                                            var v57 float64 = (3.14159 - v55)
                                            _ = v57
                                        case 370:
                                            v58 := GenS3{F0: 2147483647, F1: 1, F2: 548}
                                            _ = v58
                                        case 2147483647:
                                            var v59 float32 = 1.5
                                            _ = v59
                                        default:
                                            v4 += genFn0(v49, 9223372036854775807, "")
                                        }
                                        _ = v56
                                    } else {
                                        v60 := GenS5{F0: v49, F1: 0, F2: 0}
                                        _ = v60
                                    }
                                    _ = v55
                                }
                                _ = fi51
                                _ = v52
                            }
                            _ = fi50
                        }
                        if v45 {
                            var v61 bool = ((1e10 > -2.25) && false)
                            fmt.Println(v61)
                            _ = v61
                        }
                        for fi62 := 0; fi62 < 6; fi62++ {
                            if v45 {
                                if true {
                                    v46 = v46
                                    var v63 int32 = -2147483648
                                    var v64 int8 = 1
                                    _ = v63
                                    _ = v64
                                }
                                v3 = append(v3, v46)
                            } else {
                                for fi65 := 0; fi65 < 11; fi65++ {
                                    var v66 uint32 = (v46 * (v46 - uint32(v49)))
                                    v49 += 1
                                    _ = fi65
                                    _ = v66
                                }
                                for fi67 := 0; fi67 < 19; fi67++ {
                                    v2 = v2
                                    if v45 {
                                        fmt.Println(v45)
                                    }
                                    var v68 uint16 = 808
                                    _ = fi67
                                    _ = v68
                                }
                                var v69 bool = (1 < 100)
                                _ = v69
                            }
                            switch (-(fi62)) {
                            case 128:
                                var v70 int16 = (v2 & (v2 * v2))
                                _ = v70
                            }
                            _ = fi62
                        }
                    }
                    var v71 int64 = 0
                    _ = v49
                    _ = v71
                }
                var v72 int = -750
                _ = v72
            }
            if v45 {
                v73 := v1[(v4 * genFn0(35, genFn0(42, v4, "hello"), "foo bar"))]
                _ = v73
            }
            switch v4 {
            case 2147483647:
                v2 = 0
            case -9223372036854775808:
                v4 = (v4 - (-(v4)))
            case 389:
                if len(v3) > 0 {
                    v3[0] = v46
                }
            }
        }
    case "foo bar":
        _ = len(v3)
    }
    v46 -= (v46 * v46)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
}
