package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8; F1 int16; F2 int64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 bool; F1 bool }
type GenS6 struct { F0 bool; F1 bool; F2 int32 }
func genFn0(p0 int16) int16 {
    if true {
        fmt.Println(p0)
    }
    p0 = p0
    for fi7 := 0; fi7 < 19; fi7++ {
        v8 := []int{fi7, fi7, (fi7 + (fi7 >> ((1) & 63)))}
        if ((!true) || false) {
            switch (fi7 & fi7) {
            case 128:
                p0 -= (p0 >> ((65535) & 63))
            default:
                _ = len(v8)
            }
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            var v9 uint = 4294967295
            _ = v9
        } else {
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
        }
        _ = fi7
        _ = v8
    }
    _ = p0
    return p0
}
func genFn1() bool {
    switch "" {
    case "":
        var v10 int = -1
        _ = v10
    case "hello":
        v11 := []int{-363, 32768, 758, 32768}
        _ = v11
    case "a":
        for fi12 := 0; fi12 < 20; fi12++ {
            switch "foo bar" {
            case "a":
                fmt.Println(fi12)
            }
            for fi13 := 0; fi13 < 4; fi13++ {
                var v14 uint64 = 822
                if (true || ((3.14159 == float32(float64(((1e10 - 0.1) * 1.5)))) || true)) {
                    if (((false && ((!(true && ((true || (2147483647 < (fi13 >> ((321) & 63)))) && (308 >= 246)))) || false)) || false) || true) {
                        if true {
                            var v15 float32 = (0.1 / 1.0)
                            _ = v15
                        } else {
                            v14 = v14
                            v14 -= 0
                        }
                        fmt.Println(fi13, fi13, fi13)
                    } else {
                        if (255 == 65535) {
                            switch fi13 {
                            case 140:
                                var v16 string = "a"
                                _ = v16
                            default:
                                fmt.Println(fi13)
                            }
                        } else {
                            var v17 float32 = 0.0
                            v18 := GenS1{F0: 10, F1: 249}
                            _ = v17
                            _ = v18
                        }
                        if ((true || ((fi13 ^ (fi13 >> ((1) & 63))) != (fi12 - -9223372036854775808))) || false) {
                            v14 = v14
                            fmt.Println(v14, fi13, fi12)
                        } else {
                            v14 = v14
                            v19 := GenS1{F0: 9223372036854775807, F1: 195}
                            v14 -= 884
                            _ = v19
                        }
                    }
                } else {
                    if (true || (1 <= -1)) {
                        var v20 uint64 = 42
                        _ = v20
                    } else {
                        switch (fi13 + -629) {
                        case -9223372036854775808:
                            v21 := GenS2{F0: 9223372036854775807}
                            _ = v21
                        case 128:
                            v14 -= (v14 % ((v14) | 1))
                        case 42:
                            fmt.Println(fi13, fi13, fi13)
                        }
                        v22 := []int{}
                        fmt.Println(fi13)
                        _ = v22
                    }
                    v23 := []int{(fi13 - (fi13 << ((2) & 63))), int(v14), (fi12 % ((fi12) | 1)), 2147483647}
                    _ = v23
                }
                if ((9223372036854775807 >= -1) && true) {
                    if (false || (!false)) {
                        switch 4294967295 {
                        case 42:
                            v24 := []int{(-(fi12))}
                            _ = v24
                        default:
                            v14 += v14
                            v25 := GenS3{F0: 127, F1: 1000, F2: 1}
                            _ = v25
                        }
                        var v26 string = ""
                        v27 := []int{(-(fi13))}
                        _ = v26
                        _ = v27
                    } else {
                        v28 := GenS2{F0: int64(v14)}
                        var v29 uint16 = 255
                        v29 -= (v29 % (((^(v29))) | 1))
                        _ = v28
                        _ = v29
                    }
                } else {
                    if (!false) {
                        v14 -= 0
                        var v30 int = 32767
                        var v31 uint8 = 2
                        _ = v30
                        _ = v31
                    }
                    v32 := GenS3{F0: -128, F1: 596, F2: -814}
                    fmt.Println(v14, fi13, fi12)
                    _ = v32
                }
                _ = fi13
                _ = v14
            }
            _ = fi12
        }
    }
    var v33 float32 = (float32((-2.25 / 1.0)) + -2.25)
    if (!(false && false)) {
        v34 := GenS2{F0: 0}
        _ = v34
    }
    _ = v33
    return (!(42 != 1))
}

func main() {
    var v1 int64 = -425
    var v2 uint32 = 255
    var v3 uint64 = 73
    var v4 int16 = 1
    v5 := []string{("" + ""), (("foo bar" + "hello") + (("foo bar" + "foo bar") + "hello"))}
    var v6 int32 = -557
    var v35 int64 = v1
    v36 := GenS1{F0: v35, F1: 109}
    for fi37 := 0; fi37 < 19; fi37++ {
        if false {
            if true {
                v6 -= v6
            } else {
                var v38 uint16 = 351
                _ = v38
            }
            _ = len(v5)
            v4 = (v4 % ((v4) | 1))
        } else {
            v39 := GenS4{F0: 0}
            _ = v39
        }
        _ = fi37
    }
    var v40 uint8 = 2
    v40 += v40
    for fi41 := 0; fi41 < 3; fi41++ {
        fmt.Println(v35, fi41)
        for fi42 := 0; fi42 < 15; fi42++ {
            var v43 uint64 = 0
            fmt.Println(v35, v1, fi42)
            _ = fi42
            _ = v43
        }
        _ = fi41
    }
    switch "foo bar" {
    case "\n\t":
        var v44 uint = 42
        _ = v44
    case "foo bar":
        for fi45 := 0; fi45 < 16; fi45++ {
            v46 := GenS5{F0: true, F1: true}
            var v47 int16 = (v4 + genFn0(v4))
            _ = fi45
            _ = v46
            _ = v47
        }
    case "":
        _ = len(v5)
    default:
        v48 := GenS6{F0: ((!true) && (678 == 255)), F1: (712 > 0), F2: v6}
        _ = v48
    }
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
