package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64; F1 uint }
type GenS5 struct { F0 uint16; F1 int8; F2 int }

func main() {
    var v1 float32 = 0.1
    var v2 uint = 2
    var v3 uint64 = 4294967295
    for fi4 := 0; fi4 < 5; fi4++ {
        v5 := GenS1{F0: (false && false), F1: 48, F2: (fi4 % ((fi4) | 1))}
        v6 := []int{-9223372036854775808, -2147483648, (fi4 << ((v2) & 63)), fi4}
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 string = "hello"
    if true {
        var v8 uint32 = 0
        _ = v8
    } else {
        var v9 int = 127
        _ = v9
    }
    v10 := []int{127, -2147483648, -568, 594}
    var v11 float32 = (-2.25 / v1)
    var v12 uint32 = 655
    for fi13 := 0; fi13 < 13; fi13++ {
        for fi14 := 0; fi14 < 3; fi14++ {
            v11 = float32(1e10)
            switch v7 {
            case "":
                if len(v10) > 0 {
                    v10[0] = -9223372036854775808
                }
            }
            var v15 uint32 = (v12 >> (((v2 << ((42) & 63))) & 63))
            _ = fi14
            _ = v15
        }
        if (!(true || (v11 == float32(3.14159)))) {
            if true {
                v16 := GenS1{F0: (0 < 2), F1: 127, F2: 32767}
                v17 := GenS2{F0: int8(fi13)}
                v10 = append(v10, (^(fi13)))
                _ = v16
                _ = v17
            }
            var v18 int64 = 0
            for fi19 := 0; fi19 < 4; fi19++ {
                switch (fi13 * fi13) {
                case 1:
                    v2 = uint(v3)
                case -731:
                    fmt.Println(v18, v11)
                default:
                    v12 = v12
                }
                var v20 int64 = 934
                v20 = v18
                _ = fi19
                _ = v20
            }
            _ = v18
        }
        var v21 string = ((v7 + "hello") + (v7 + "hello"))
        _ = fi13
        _ = v21
    }
    var v22 int8 = 86
    v3 += 1
    if (!false) {
        v7 = (v7 + (v7 + "foo bar"))
        v1 += (v1 / (-2.25 - v11))
    }
    var v23 uint8 = 128
    if (!(false && (true || ((1 <= -250) && (!(695 != 652)))))) {
        switch 124 {
        case 0:
            if (true && ((0.1 * v1) <= (-2.25 + v11))) {
                if (9223372036854775807 <= int64(v23)) {
                    var v24 uint16 = 65535
                    v23 -= (^(v23))
                    _ = v24
                }
                for fi25 := 0; fi25 < 1; fi25++ {
                    fmt.Println(v1, v7)
                    _ = fi25
                }
            }
        default:
            for fi26 := 0; fi26 < 11; fi26++ {
                if (("hello" + v7) >= "a") {
                    var v27 string = ""
                    var v28 uint32 = (v12 | 42)
                    for fi29 := 0; fi29 < 19; fi29++ {
                        switch (("hello" + v27) + v7) {
                        case "hello":
                            v30 := GenS1{F0: true, F1: v23, F2: 0}
                            _ = v30
                        default:
                            fmt.Println(v23, v22)
                        }
                        _ = len(v10)
                        _ = fi29
                    }
                    _ = v27
                    _ = v28
                }
                v12 = v12
                v3 = v3
                _ = fi26
            }
        }
        switch (v7 + v7) {
        case "foo bar":
            v31 := GenS3{F0: -1}
            _ = v31
        case "\n\t":
            _ = len(v10)
        }
        switch 127 {
        case -219:
            var v32 float32 = (-2.25 / (v1 * v1))
            _ = v32
        case 42:
            if (308 >= 0) {
                var v33 uint16 = 910
                var v34 float32 = ((v1 + float32(0.1)) + 0.1)
                var v35 uint8 = v23
                _ = v33
                _ = v34
                _ = v35
            }
        }
    } else {
        var v36 uint32 = uint32(v3)
        v23 += (v23 << ((uint(v3)) & 63))
        _ = v36
    }
    var v37 float32 = (v11 + ((v11 + float32((0.1 * (0.1 - (3.14159 - (3.14159 + 0.0)))))) / v1))
    var v38 uint8 = 42
    var v39 uint64 = (^(v3))
    if ((true || (((!((v38 * v23) <= (v23 - (v38 & 42)))) || true) || true)) || (v3 != 2)) {
        for fi40 := 0; fi40 < 16; fi40++ {
            var v41 uint64 = (^(v39))
            for fi42 := 0; fi42 < 12; fi42++ {
                _ = len(v10)
                var v43 int = 4294967295
                _ = fi42
                _ = v43
            }
            if true {
                var v44 float32 = (1.5 + 1.5)
                _ = v44
            } else {
                var v45 int64 = 466
                v23 += (v38 % ((v38) | 1))
                if true {
                    switch ("\n\t" + "a") {
                    case "\n\t":
                        if true {
                            var v46 int16 = -1
                            v47 := GenS4{F0: (-(v45)), F1: uint(v12)}
                            v22 = v22
                            _ = v46
                            _ = v47
                        } else {
                            if (-32768 < 1) {
                                var v48 bool = (4294967295 == v39)
                                _ = v48
                            } else {
                                for fi49 := 0; fi49 < 9; fi49++ {
                                    v22 -= v22
                                    var v50 uint32 = (v12 * v12)
                                    var v51 int = fi49
                                    _ = fi49
                                    _ = v50
                                    _ = v51
                                }
                                v52 := GenS3{F0: 0}
                                v53 := GenS2{F0: int8(v2)}
                                _ = v52
                                _ = v53
                            }
                        }
                    case "":
                        v54 := GenS1{F0: (false && false), F1: 42, F2: fi40}
                        _ = v54
                    case "hello":
                        if len(v10) > 0 {
                            v10[0] = fi40
                        }
                    default:
                        v22 += v22
                    }
                    var v55 uint16 = 32767
                    _ = v55
                } else {
                    var v56 float64 = 0.0
                    _ = v56
                }
                _ = v45
            }
            _ = fi40
            _ = v41
        }
        v10 = append(v10, -326)
        v3 -= (v39 >> ((v2) & 63))
    } else {
        if len(v10) > 0 {
            v10[0] = 0
        }
        var v57 uint8 = uint8(v3)
        _ = v57
    }
    switch v7 {
    case "a":
        _ = len(v10)
    case "":
        var v58 int = 566
        _ = v58
    default:
        v1 = (v37 + v11)
    }
    v59 := GenS5{F0: 42, F1: v22, F2: 9223372036854775807}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v59
    _ = v7
}
