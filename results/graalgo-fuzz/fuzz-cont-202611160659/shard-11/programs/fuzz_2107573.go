package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16; F1 uint32; F2 bool }
type GenS4 struct { F0 uint32; F1 uint16 }
type GenS5 struct { F0 string }
func genFn0(p0 uint16, p1 int, p2 int32) int8 {
    var v5 bool = true
    fmt.Println(p1, p0, p2)
    p1 -= -665
    v6 := []int{(p1 % ((p1) | 1)), (-(p1))}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 float32 = 3.14159
    var v2 int16 = -32768
    v3 := make(map[uint32]int)
    v4 := []uint32{570, 2, 523}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v7 int = 1
    v7 = (v7 & v7)
    switch v7 {
    case -508:
        fmt.Println(v1, v1)
    case -702:
        if (-1 < -2147483648) {
            if false {
                fmt.Println(v1, v7, v1)
                v8 := GenS1{F0: "\n\t", F1: (v7 ^ v7)}
                _ = v8
            }
        } else {
            var v9 float64 = 3.14159
            var v10 uint64 = 42
            if true {
                v2 += -402
                delete(v3, 49)
            }
            _ = v10
            _ = v9
        }
    case 42:
        v11 := GenS1{F0: ("a" + "\n\t"), F1: 127}
        _ = v11
    default:
        v12 := GenS2{F0: 32767}
        _ = v12
    }
    var v13 int = 32768
    switch 32768 {
    case 127:
        if (!(15 < 2)) {
            for fi14 := 0; fi14 < 3; fi14++ {
                v15 := GenS2{F0: 65535}
                v16 := GenS3{F0: (v2 | -729), F1: 1, F2: false}
                _ = fi14
                _ = v15
                _ = v16
            }
        } else {
            var v17 int8 = 8
            switch "hello" {
            case "\n\t":
                var v18 uint16 = 32767
                _ = v18
            case "foo bar":
                if true {
                    var v19 float32 = (3.14159 - v1)
                    if false {
                        _ = len(v3)
                    }
                    _ = v19
                } else {
                    var v20 uint32 = 1
                    var v21 bool = false
                    v22 := GenS3{F0: v2, F1: v20, F2: false}
                    _ = v20
                    _ = v21
                    _ = v22
                }
            }
            if true {
                v23 := v3[2]
                if true {
                    if (!(false || true)) {
                        v24 := GenS1{F0: ("a" + "hello"), F1: v13}
                        for fi25 := 0; fi25 < 6; fi25++ {
                            switch "" {
                            case "foo bar":
                                v1 = float32((-2.25 - 0.1))
                            case "hello":
                                var v26 int8 = genFn0(42, (-(v23)), 0)
                                _ = v26
                            case "\n\t":
                                var v27 uint32 = 42
                                _ = v27
                            }
                            v28 := GenS4{F0: 0, F1: 42}
                            _ = fi25
                            _ = v28
                        }
                        _ = v24
                    } else {
                        var v29 int64 = -9223372036854775808
                        v17 -= (v17 + (v17 >> ((1) & 63)))
                        if (65535 == 65535) {
                            _ = len(v4)
                            v23 -= v23
                            if (318 > 2) {
                                delete(v3, 2147483647)
                                for fi30 := 0; fi30 < 17; fi30++ {
                                    v7 += -9223372036854775808
                                    var v31 int = int(v2)
                                    v29 = v29
                                    _ = fi30
                                    _ = v31
                                }
                            }
                        } else {
                            v32 := GenS1{F0: "", F1: v7}
                            if (true && false) {
                                v33 := GenS2{F0: 255}
                                v23 = (v7 / (((v13 - (-(v23)))) | 1))
                                _ = v33
                            } else {
                                v7 = v13
                            }
                            _ = v32
                        }
                        _ = v29
                    }
                    v2 -= 32767
                    switch ("foo bar" + ("" + (("foo bar" + "hello") + "\n\t"))) {
                    case "hello":
                        v34 := GenS5{F0: "a"}
                        _ = v34
                    case "\n\t":
                        fmt.Println(v23)
                    }
                } else {
                    fmt.Println(v17)
                    v1 += 1.5
                    _ = len(v3)
                }
                _ = v23
            }
            _ = v17
        }
    case -9223372036854775808:
        v13 = (v13 + (v7 / ((-1) | 1)))
    case -55:
        switch v13 {
        case 4294967295:
            v2 += v2
        case 127:
            switch "hello" {
            case "":
                if (v1 < 1.5) {
                    if len(v4) > 0 {
                        v4[0] = 2147483647
                    }
                } else {
                    var v35 string = ("hello" + "hello")
                    var v36 bool = false
                    _ = len(v3)
                    _ = v35
                    _ = v36
                }
            }
        }
    default:
        _ = len(v3)
    }
    var v37 uint16 = 581
    v2 = v2
    var v38 int32 = 497
    switch (v7 / (((v7 * v13)) | 1)) {
    case 460:
        switch v13 {
        case 1:
            fmt.Println(v7)
        case 9223372036854775807:
            var v39 int8 = 98
            _ = v39
        case 448:
            fmt.Println(v7, v37, v13)
        }
    case 1:
        v2 = int16(v7)
    }
    v13 += (^(v13))
    for fi40 := 0; fi40 < 12; fi40++ {
        for fi41 := 0; fi41 < 20; fi41++ {
            _ = len(v4)
            _ = fi41
        }
        _ = fi40
    }
    v37 = v37
    var v42 uint = 0
    var v43 int16 = (v2 >> ((0) & 63))
    var v44 int32 = v38
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v7
}
