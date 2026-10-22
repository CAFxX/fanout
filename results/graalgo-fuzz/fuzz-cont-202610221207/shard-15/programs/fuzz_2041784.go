package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 int8 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 uint8; F1 uint64; F2 uint64 }
func genFn0() uint {
    var v7 uint8 = 59
    v8 := []int{-1}
    v7 = (v7 >> ((721) & 63))
    _ = v7
    _ = v8
    return 707
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]bool)
    v3 := []string{"", "hello"}
    v4 := []uint32{65535, 222, 2}
    var v5 int8 = 15
    var v6 float64 = float64(1.5)
    var v9 uint32 = 65535
    v5 -= (v5 >> ((586) & 63))
    switch 1 {
    case 128:
        var v10 uint64 = 0
        _ = v10
    case -2147483648:
        switch 185 {
        case 292:
            v11 := GenS1{F0: 42, F1: 32768, F2: v5}
            _ = v11
        case 32767:
            v6 -= 1e10
        }
    case 32767:
        v6 += v6
    }
    var v12 float32 = (((0.0 + (3.14159 / 1.0)) - (0.0 + -2.25)) - 3.14159)
    v13 := GenS1{F0: -679, F1: -207, F2: v5}
    var v14 uint64 = 4294967295
    var v15 int = 32768
    var v16 int8 = 100
    if (!true) {
        v6 += float64(v12)
        if true {
            v17 := GenS2{F0: 255, F1: 2147483647}
            for fi18 := 0; fi18 < 18; fi18++ {
                v15 -= (v15 - (fi18 | 738))
                if ((false && ((v14 | v14) <= (^(v14)))) && false) {
                    switch ("hello" + (((("foo bar" + ("hello" + "")) + ("foo bar" + "a")) + "hello") + "")) {
                    case "":
                        v5 = v5
                    case "\n\t":
                        var v19 uint64 = (v14 | (^(v14)))
                        _ = v19
                    case "foo bar":
                        v14 = (^(v14))
                    }
                } else {
                    for fi20 := 0; fi20 < 3; fi20++ {
                        v1["foo bar"] = ("" + "\n\t")
                        var v21 uint16 = 437
                        delete(v2, "\n\t")
                        _ = fi20
                        _ = v21
                    }
                    v22 := GenS3{F0: 0, F1: (^(v14)), F2: v14}
                    _ = v22
                }
                delete(v2, ((("" + "foo bar") + "hello") + ""))
                _ = fi18
            }
            v23 := GenS3{F0: 2, F1: 2, F2: v14}
            _ = v17
            _ = v23
        }
    } else {
        if (!true) {
            v15 -= v15
        }
        var v24 float64 = v6
        _ = v24
    }
    v12 = (3.14159 + 0.1)
    var v25 uint = 2
    fmt.Println(v15, v6)
    fmt.Println(v12, v5, v15)
    if false {
        v5 = (v16 * v5)
        for fi26 := 0; fi26 < 7; fi26++ {
            if true {
                var v27 int8 = (v16 % (((v16 & (v5 - 42))) | 1))
                if (v14 >= uint64(v25)) {
                    switch (("hello" + "foo bar") + "\n\t") {
                    case "a":
                        _ = len(v2)
                    case "foo bar":
                        var v28 uint32 = v9
                        _ = v28
                    }
                    fmt.Println(v12, v9)
                    if (((-478 < 519) && true) && false) {
                        if (!(v14 <= 2)) {
                            var v29 int64 = 42
                            v25 -= v25
                            if true {
                                switch ("hello" + "a") {
                                case "a":
                                    for fi30 := 0; fi30 < 19; fi30++ {
                                        var v31 int16 = -973
                                        _ = fi30
                                        _ = v31
                                    }
                                default:
                                    _ = len(v1)
                                }
                                v25 = (v25 << ((v25) & 63))
                                fmt.Println(fi26, v25, v15)
                            }
                            _ = v29
                        }
                        if false {
                            v32 := GenS3{F0: 0, F1: (^(v14)), F2: 1}
                            _ = v32
                        } else {
                            if (!false) {
                                fmt.Println(v12, fi26, v27)
                                for fi33 := 0; fi33 < 2; fi33++ {
                                    v14 = 0
                                    var v34 float32 = (float32(v6) - float32((1.5 + ((((1e10 / 1.0) + 0.1) * v6) * 0.0))))
                                    _ = fi33
                                    _ = v34
                                }
                            }
                            for fi35 := 0; fi35 < 19; fi35++ {
                                var v36 int64 = 1
                                fmt.Println(v36, v25)
                                delete(v1, "")
                                _ = fi35
                                _ = v36
                            }
                            var v37 uint16 = 2
                            _ = v37
                        }
                    } else {
                        v6 = v6
                        var v38 uint8 = 2
                        fmt.Println(v12)
                        _ = v38
                    }
                } else {
                    v39 := GenS1{F0: -1, F1: fi26, F2: (v5 * (v27 / ((-32) | 1)))}
                    _ = v39
                }
                v25 = v25
                _ = v27
            } else {
                switch (("hello" + ((((("\n\t" + "hello") + "foo bar") + ("\n\t" + ((("\n\t" + "a") + ("\n\t" + "foo bar")) + ""))) + "hello") + "a")) + ("" + ("" + ("foo bar" + "\n\t")))) {
                case "":
                    var v40 float64 = v6
                    _ = v40
                }
                v5 = (^(v16))
            }
            v5 += (v16 >> (((v25 - (v25 - (v25 + 978)))) & 63))
            _ = fi26
        }
    } else {
        var v41 uint64 = 18446744073709551615
        var v42 uint8 = 84
        _ = v41
        _ = v42
    }
    for fi43 := 0; fi43 < 3; fi43++ {
        v5 = (-(v5))
        _ = fi43
    }
    var v44 string = ("\n\t" + ("" + ("" + "")))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v6
    _ = v9
}
