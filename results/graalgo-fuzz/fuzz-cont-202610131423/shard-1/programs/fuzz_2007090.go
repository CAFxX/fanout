package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 uint16; F1 uint32; F2 int16 }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint8) string {
    for fi5 := 0; fi5 < 12; fi5++ {
        v6 := []int{}
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 12; fi7++ {
        v8 := []int{(fi7 & fi7)}
        p0 += (p0 / ((127) | 1))
        _ = fi7
        _ = v8
    }
    var v9 uint16 = 882
    var v10 int64 = 0
    v10 += (v10 >> ((2) & 63))
    _ = p0
    _ = v10
    _ = v9
    return "hello"
}
func genFn1(p0 int8, p1 uint16) uint8 {
    v11 := []int{32768, -31, 32767}
    switch (("hello" + "foo bar") + (("\n\t" + "\n\t") + "a")) {
    case "hello":
        if (((4294967295 <= 737) || (!(false && false))) || (!true)) {
            v12 := GenS1{F0: 0, F1: -1}
            _ = v12
        } else {
            if true {
                p0 = (p0 / ((p0) | 1))
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                for fi13 := 0; fi13 < 2; fi13++ {
                    fmt.Println(fi13)
                    _ = fi13
                }
            }
            v14 := GenS2{F0: p1, F1: 2, F2: 1}
            v11 = append(v11, 128)
            _ = v14
        }
    }
    p0 -= (p0 & 127)
    switch 356 {
    case -307:
        var v15 int8 = p0
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v11
    return 177
}

func main() {
    var v1 int = 32768
    var v2 uint64 = 4294967295
    var v3 bool = true
    v4 := []uint32{}
    var v16 string = "a"
    var v17 int32 = -1
    v18 := GenS2{F0: 1, F1: 2, F2: 42}
    var v19 int8 = -1
    if len(v4) > 0 {
        v4[0] = 143
    }
    if (v19 < (v19 / (((v19 | (v19 & v19))) | 1))) {
        if (0.1 >= float64(-2.25)) {
            for fi20 := 0; fi20 < 15; fi20++ {
                v2 = (v2 | v2)
                for fi21 := 0; fi21 < 8; fi21++ {
                    var v22 uint64 = 0
                    var v23 int64 = -52
                    switch (fi21 + (fi21 | fi20)) {
                    case 4294967295:
                        fmt.Println(v3, v2, v22)
                    case 83:
                        switch (fi20 >> ((65535) & 63)) {
                        case 0:
                            v24 := GenS2{F0: 981, F1: 1, F2: 42}
                            _ = v24
                        case 1:
                            v25 := GenS1{F0: 858, F1: int32(v19)}
                            _ = v25
                        default:
                            v17 -= 0
                        }
                    case 42:
                        switch -444 {
                        case 32768:
                            var v26 float64 = 3.14159
                            if ((!v3) && ((0 != v2) && (float64(-2.25) <= v26))) {
                                fmt.Println(v26, v26, v22)
                                v4 = append(v4, 2147483647)
                            }
                            _ = v26
                        }
                    }
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                _ = fi20
            }
            fmt.Println(v1, v16, v16)
        } else {
            switch 32768 {
            case 32768:
                for fi27 := 0; fi27 < 1; fi27++ {
                    v28 := GenS1{F0: 1, F1: (v17 * v17)}
                    _ = fi27
                    _ = v28
                }
            case -975:
                var v29 int32 = int32(v2)
                _ = v29
            default:
                var v30 int = (v1 >> ((8) & 63))
                _ = v30
            }
        }
        v31 := GenS3{F0: 1}
        switch (v16 + "hello") {
        case "hello":
            var v32 uint8 = 127
            _ = v32
        default:
            for fi33 := 0; fi33 < 19; fi33++ {
                if v3 {
                    v19 = (v19 << ((356) & 63))
                    if v3 {
                        if (1 > 4294967295) {
                            for fi34 := 0; fi34 < 19; fi34++ {
                                if v3 {
                                    v16 = v16
                                    fmt.Println(fi33, fi33)
                                    var v35 bool = true
                                    v36 := GenS2{F0: 255, F1: 2147483648, F2: -32768}
                                    _ = v35
                                    _ = v36
                                }
                                _ = fi34
                            }
                        } else {
                            var v37 float32 = 1.5
                            _ = len(v4)
                            _ = v37
                        }
                        switch v16 {
                        case "hello":
                            if v3 {
                                v1 -= (fi33 + v1)
                            } else {
                                v1 -= v1
                                v17 -= (^(v17))
                            }
                        default:
                            if len(v4) > 0 {
                                v4[0] = 42
                            }
                        }
                    }
                } else {
                    v3 = (v2 != 0)
                }
                v1 += (fi33 & (v1 << ((2) & 63)))
                _ = fi33
            }
        }
        _ = v31
    } else {
        if v3 {
            var v38 int64 = -1
            fmt.Println(v38, v1, v2)
            _ = v38
        }
        var v39 uint16 = 255
        _ = v39
    }
    _ = len(v4)
    if ((v1 * -525) >= v1) {
        fmt.Println(v3)
    } else {
        switch v1 {
        case 128:
            if v3 {
                fmt.Println(v19, v19, v16)
                var v40 uint32 = 2147483647
                switch v1 {
                case 150:
                    v4 = append(v4, 2147483648)
                case -607:
                    if v3 {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        switch "foo bar" {
                        case "":
                            fmt.Println(v16, v1, v1)
                        default:
                            _ = len(v4)
                        }
                        for fi41 := 0; fi41 < 13; fi41++ {
                            switch 32767 {
                            case -2147483648:
                                v19 = (v19 % (((v19 % ((-12) | 1))) | 1))
                            }
                            _ = fi41
                        }
                    } else {
                        v42 := GenS1{F0: 29, F1: v17}
                        _ = v42
                    }
                case -9223372036854775808:
                    v16 = v16
                default:
                    v43 := GenS3{F0: 1}
                    _ = v43
                }
                _ = v40
            } else {
                var v44 bool = (float32(1.5) == -2.25)
                _ = v44
            }
        case 32768:
            v3 = v3
        }
        if v3 {
            if len(v4) > 0 {
                v4[0] = 2147483648
            }
        } else {
            var v45 uint = 218
            for fi46 := 0; fi46 < 12; fi46++ {
                var v47 uint8 = 192
                _ = fi46
                _ = v47
            }
            _ = v45
        }
    }
    if true {
        var v48 int64 = -1
        _ = v48
    } else {
        switch 987 {
        case -1:
            for fi49 := 0; fi49 < 12; fi49++ {
                fmt.Println(v16)
                _ = fi49
            }
        }
        v16 = "\n\t"
    }
    v50 := GenS2{F0: 0, F1: 42, F2: 497}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v16 = (v16 + (v16 + (v16 + "a")))
    v17 += (v17 << ((4294967295) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v50
}
