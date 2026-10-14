package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint8 }
type GenS2 struct { F0 int16; F1 string; F2 string }
type GenS3 struct { F0 int16; F1 uint32; F2 int64 }

func main() {
    var v1 uint = 4294967295
    v2 := make(map[uint32]bool)
    var v3 int64 = 575
    var v4 float32 = (1.5 + 3.14159)
    var v5 uint64 = 1
    if (606 == 325) {
        var v6 int16 = -1
        _ = v6
    } else {
        var v7 uint32 = 0
        if false {
            var v8 uint16 = 576
            delete(v2, v7)
            if (-821 < 1) {
                v3 = v3
                if true {
                    v9 := GenS1{F0: 2147483647, F1: ("" + "a"), F2: 42}
                    v10 := GenS1{F0: 1, F1: (("\n\t" + "foo bar") + ("hello" + ("\n\t" + "a"))), F2: 255}
                    _ = v10
                    _ = v9
                } else {
                    v7 -= v7
                    var v11 int32 = 1
                    v4 = v4
                    _ = v11
                }
            }
            _ = v8
        }
        fmt.Println(v7, v7, v4)
        _ = v7
    }
    for fi12 := 0; fi12 < 7; fi12++ {
        v13 := GenS2{F0: 1000, F1: "a", F2: ""}
        v5 += v5
        _ = fi12
        _ = v13
    }
    v14 := []int{-289, 9223372036854775807, 127}
    if true {
        var v15 int64 = (v3 * (v3 / (((v3 ^ (v3 & 30))) | 1)))
        _ = v15
    } else {
        if false {
            v4 += v4
        }
        fmt.Println(v4)
    }
    var v16 uint8 = 0
    for fi17 := 0; fi17 < 17; fi17++ {
        for fi18 := 0; fi18 < 1; fi18++ {
            if true {
                v16 = (^(v16))
                if len(v14) > 0 {
                    v14[0] = (fi17 * fi17)
                }
                v19 := GenS1{F0: int32(fi17), F1: "", F2: v16}
                _ = v19
            } else {
                v16 = 0
            }
            _ = len(v14)
            v3 = (v3 - v3)
            v20 := GenS1{F0: 42, F1: ("foo bar" + "hello"), F2: (v16 - (v16 >> ((255) & 63)))}
            _ = fi18
            _ = v20
        }
        _ = fi17
    }
    switch "" {
    case "\n\t":
        for fi21 := 0; fi21 < 2; fi21++ {
            v3 = (v3 ^ v3)
            v3 += 257
            v22 := GenS1{F0: 0, F1: (((("a" + ("foo bar" + ("a" + "\n\t"))) + (("\n\t" + "foo bar") + ("hello" + "a"))) + "hello") + "\n\t"), F2: 230}
            _ = fi21
            _ = v22
        }
    case "hello":
        v5 += v5
    default:
        v14 = append(v14, -869)
    }
    for fi23 := 0; fi23 < 4; fi23++ {
        _ = len(v14)
        _ = fi23
    }
    v24 := GenS3{F0: int16(v5), F1: 2, F2: v3}
    for fi25 := 0; fi25 < 17; fi25++ {
        v2[2] = true
        v26 := GenS3{F0: -1, F1: 862, F2: (v3 >> ((v1) & 63))}
        var v27 int64 = 795
        _ = fi25
        _ = v26
        _ = v27
    }
    var v28 uint = 4294967295
    switch 32767 {
    case 32768:
        switch ("\n\t" + "\n\t") {
        case "\n\t":
            for fi29 := 0; fi29 < 15; fi29++ {
                if true {
                    v30 := GenS2{F0: -1, F1: ("hello" + "foo bar"), F2: "a"}
                    v31 := GenS3{F0: 720, F1: 0, F2: (v3 / (((^(v3))) | 1))}
                    v14 = append(v14, 9223372036854775807)
                    _ = v30
                    _ = v31
                } else {
                    v5 -= v5
                    if true {
                        var v32 int16 = 583
                        for fi33 := 0; fi33 < 8; fi33++ {
                            if true {
                                if false {
                                    _ = len(v14)
                                }
                            }
                            if true {
                                var v34 uint8 = v16
                                v16 = v16
                                _ = v34
                            }
                            var v35 float64 = float64(1.5)
                            _ = fi33
                            _ = v35
                        }
                        _ = v32
                    }
                }
                v36 := GenS2{F0: -1, F1: ("foo bar" + ("foo bar" + "hello")), F2: "hello"}
                var v37 int64 = (v3 | -1)
                _ = fi29
                _ = v36
                _ = v37
            }
        case "foo bar":
            for fi38 := 0; fi38 < 18; fi38++ {
                v39 := GenS1{F0: int32(fi38), F1: "\n\t", F2: uint8(v5)}
                for fi40 := 0; fi40 < 16; fi40++ {
                    if false {
                        v4 = ((3.14159 + (v4 + 0.1)) - v4)
                        v2[2] = (((!true) && (true && false)) && (v4 <= 0.1))
                        var v41 int = (-(fi38))
                        v3 = 151
                        _ = v41
                    } else {
                        var v42 int64 = v3
                        _ = v42
                    }
                    if false {
                        fmt.Println(fi38, v28)
                        switch "foo bar" {
                        case "foo bar":
                            v4 -= (float32((((3.14159 / 0.1) * (0.0 + float64(v4))) + 1.5)) / 1.5)
                            var v43 int16 = -881
                            _ = v43
                        default:
                            var v44 uint32 = 2
                            _ = v44
                        }
                        if true {
                            var v45 int64 = 1
                            if (0 <= v16) {
                                v45 -= v45
                            } else {
                                v2[2147483647] = (84 != -79)
                            }
                            var v46 uint16 = 2
                            _ = v45
                            _ = v46
                        } else {
                            v1 = 277
                            fmt.Println(v1, v4)
                            var v47 int8 = -128
                            _ = v47
                        }
                    }
                    fmt.Println(fi38, v4, v4)
                    _ = fi40
                }
                fmt.Println(v4, v4, v1)
                _ = fi38
                _ = v39
            }
        default:
            v1 = (v1 >> ((65535) & 63))
        }
    case 42:
        v1 = (v28 ^ (v28 << ((v28) & 63)))
    case 4294967295:
        if (2147483648 >= 189) {
            var v48 uint32 = 175
            _ = v48
        }
    default:
        for fi49 := 0; fi49 < 14; fi49++ {
            for fi50 := 0; fi50 < 6; fi50++ {
                v28 = (v28 << ((289) & 63))
                v4 = v4
                if (0 >= 42) {
                    fmt.Println(v3, v16)
                    _ = len(v14)
                }
                _ = fi50
            }
            _ = fi49
        }
    }
    var v51 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v14
    _ = v16
    _ = v2
    _ = v24
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v51
}
