package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 uint64; F1 string }

func main() {
    v1 := make(map[uint32]string)
    v2 := []uint32{712, 865, 138, 42}
    var v3 string = "\n\t"
    var v4 int16 = -1
    v5 := []uint32{219}
    var v6 float64 = (float64(0.0) + (((-2.25 + 0.1) / (1e10 * (float64((0.1 / float32(0.1))) + 1e10))) + 1.5))
    switch 42 {
    case 32767:
        if len(v5) > 0 {
            v5[0] = 42
        }
    case 127:
        v3 = (("a" + v3) + ("hello" + v3))
    case -809:
        if false {
            var v7 uint64 = 614
            v3 = v3
            _ = v7
        }
    }
    for fi8 := 0; fi8 < 6; fi8++ {
        if ((42 == 4294967295) && (!true)) {
            switch ((("a" + v3) + v3) + ("foo bar" + ("hello" + v3))) {
            case "foo bar":
                var v9 int16 = v4
                _ = v9
            case "\n\t":
                var v10 float32 = 0.1
                _ = v10
            case "":
                if (394 == 2) {
                    fmt.Println(v6, v4, fi8)
                    v11 := GenS1{F0: 255, F1: 42}
                    _ = v11
                }
            default:
                fmt.Println(v3)
            }
            delete(v1, 947)
            if (((v4 >> ((0) & 63)) <= (v4 | v4)) || false) {
                v12 := GenS2{F0: 0}
                _ = v12
            }
        }
        if false {
            var v13 uint8 = 128
            delete(v1, 283)
            _ = v13
        } else {
            v14 := GenS3{F0: 127, F1: "foo bar"}
            fmt.Println(fi8, v3, v3)
            _ = v14
        }
        for fi15 := 0; fi15 < 16; fi15++ {
            var v16 int16 = v4
            switch (-(fi8)) {
            case -2147483648:
                v6 += v6
            default:
                fmt.Println(fi8, fi8)
            }
            _ = fi15
            _ = v16
        }
        _ = fi8
    }
    var v17 int16 = -32768
    for fi18 := 0; fi18 < 7; fi18++ {
        for fi19 := 0; fi19 < 19; fi19++ {
            v4 = v17
            if (577 == 32767) {
                var v20 uint = 1
                _ = len(v1)
                _ = v20
            }
            v3 = "hello"
            var v21 float64 = v6
            v22 := GenS4{F0: 460, F1: ""}
            _ = fi19
            _ = v21
            _ = v22
        }
        _ = fi18
    }
    var v23 int8 = 127
    v23 = (v23 * int8(v4))
    v4 = -32768
    if true {
        for fi24 := 0; fi24 < 7; fi24++ {
            if true {
                if (128 < 255) {
                    var v25 uint32 = 4294967295
                    v23 = v23
                    _ = v25
                } else {
                    if true {
                        v3 = ((("" + v3) + v3) + v3)
                        if false {
                            for fi26 := 0; fi26 < 8; fi26++ {
                                v6 += v6
                                var v27 uint = 4294967295
                                _ = fi26
                                _ = v27
                            }
                            _ = len(v5)
                        } else {
                            v17 += v4
                            fmt.Println(fi24)
                        }
                    } else {
                        if (!(((929 >= 363) && (-2.25 != (v6 / v6))) || true)) {
                            for fi28 := 0; fi28 < 9; fi28++ {
                                v23 = v23
                                _ = len(v1)
                                _ = fi28
                            }
                            for fi29 := 0; fi29 < 18; fi29++ {
                                v30 := GenS4{F0: 0, F1: ""}
                                _ = len(v5)
                                var v31 uint8 = 42
                                if (float64(((3.14159 + 1.5) / 1.0)) > 1.5) {
                                    v31 = 238
                                    v31 -= v31
                                    var v32 int32 = -1
                                    v17 += 0
                                    _ = v32
                                }
                                _ = fi29
                                _ = v30
                                _ = v31
                            }
                            _ = len(v2)
                        } else {
                            for fi33 := 0; fi33 < 7; fi33++ {
                                var v34 bool = false
                                var v35 int = 186
                                _ = fi33
                                _ = v34
                                _ = v35
                            }
                            fmt.Println(v3)
                            var v36 uint32 = 2147483648
                            _ = v36
                        }
                    }
                    for fi37 := 0; fi37 < 14; fi37++ {
                        v3 = v3
                        _ = fi37
                    }
                    var v38 float32 = ((-2.25 / (3.14159 + 0.1)) - (1.5 / (0.1 + 1e10)))
                    _ = v38
                }
                v39 := GenS4{F0: 1, F1: ("foo bar" + v3)}
                _ = v39
            } else {
                if len(v5) > 0 {
                    v5[0] = 681
                }
                if (!((!((true || (1 >= 4294967295)) && false)) && true)) {
                    var v40 uint64 = 4294967295
                    _ = v40
                }
            }
            _ = fi24
        }
        switch -9223372036854775808 {
        case -106:
            fmt.Println(v3, v4)
        case 0:
            v41 := GenS4{F0: 42, F1: "\n\t"}
            _ = v41
        }
        v6 += (v6 + (1e10 / (-2.25 - (1e10 / v6))))
    } else {
        fmt.Println(v3, v4, v3)
        fmt.Println(v23)
        v4 = (v4 >> ((1) & 63))
    }
    v42 := GenS4{F0: 540, F1: ("a" + v3)}
    v43 := GenS1{F0: 255, F1: 933}
    var v44 float32 = 3.14159
    fmt.Println(v6, v4, v6)
    v45 := GenS2{F0: 99}
    switch -9223372036854775808 {
    case 241:
        fmt.Println(v4, v6, v17)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
