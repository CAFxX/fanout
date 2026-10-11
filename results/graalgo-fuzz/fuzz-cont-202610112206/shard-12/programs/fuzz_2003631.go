package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 uint64 }

func main() {
    var v1 uint16 = 0
    v2 := make(map[uint32]bool)
    v3 := []string{("" + ""), "foo bar", "", ((("\n\t" + "a") + "a") + ("\n\t" + "hello"))}
    v4 := make(map[uint32]int)
    v5 := make(map[int]bool)
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 float64 = 1.5
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 32768, F1: ((true || ((128 > 128) || (!(32767 == -735)))) || false), F2: true}
    for fi9 := 0; fi9 < 18; fi9++ {
        var v10 float64 = 0.0
        fmt.Println(v1)
        v11 := GenS2{F0: fi9}
        _ = fi9
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 7; fi12++ {
        v1 = 0
        if (false || (!(943 != 42))) {
            if true {
                v1 -= v1
            }
            var v13 int64 = 1
            if false {
                for fi14 := 0; fi14 < 10; fi14++ {
                    if ((111 > 80) || ((1e10 - 1.5) < (float32(1.5) - ((float32(0.1) - 0.1) - (0.1 + (0.0 - 1e10)))))) {
                        _ = len(v5)
                        var v15 uint = 255
                        _ = v15
                    }
                    _ = len(v3)
                    _ = fi14
                }
            } else {
                fmt.Println(fi12)
                if true {
                    var v16 int = (fi12 >> ((1) & 63))
                    v17 := GenS1{F0: (fi12 & v16), F1: (!false), F2: false}
                    v18 := GenS1{F0: -2147483648, F1: (false || (false || false)), F2: false}
                    _ = v16
                    _ = v17
                    _ = v18
                } else {
                    switch "" {
                    case "hello":
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case "":
                        if ((1.5 / 0.1) > 3.14159) {
                            if false {
                                v19 := GenS2{F0: -469}
                                for fi20 := 0; fi20 < 7; fi20++ {
                                    switch "a" {
                                    case "":
                                        if true {
                                            v13 -= v13
                                        }
                                    case "a":
                                        v5[(fi20 % (((fi20 - (fi12 + 9223372036854775807))) | 1))] = (617 >= 489)
                                    case "\n\t":
                                        var v21 uint16 = 32767
                                        _ = v21
                                    }
                                    _ = fi20
                                }
                                _ = v19
                            }
                        } else {
                            v13 = v13
                            v22 := GenS3{F0: 417, F1: 575}
                            v1 = v1
                            _ = v22
                        }
                    case "\n\t":
                        if ((!true) || true) {
                            v13 += v13
                            if false {
                                delete(v4, 1)
                                v1 += v1
                                v23 := GenS4{F0: 1}
                                _ = v23
                            } else {
                                var v24 string = ("foo bar" + "\n\t")
                                if ("\n\t" > "foo bar") {
                                    v13 = (v13 + (v13 | (v13 | (v13 - -1))))
                                    _ = len(v5)
                                    var v25 uint = 42
                                    _ = v25
                                }
                                v24 = v24
                                _ = v24
                            }
                        } else {
                            switch (fi12 << ((255) & 63)) {
                            case -9223372036854775808:
                                var v26 int32 = -161
                                _ = v26
                            case 1:
                                v5[(fi12 & 2147483647)] = false
                            case 339:
                                fmt.Println(v13)
                            }
                        }
                    default:
                        var v27 uint = 2
                        _ = v27
                    }
                    for fi28 := 0; fi28 < 6; fi28++ {
                        if false {
                            var v29 uint = 710
                            var v30 float32 = (-2.25 - 0.1)
                            _ = v29
                            _ = v30
                        } else {
                            switch "foo bar" {
                            case "":
                                delete(v2, 16)
                            case "\n\t":
                                v1 += (v1 / ((v1) | 1))
                            }
                        }
                        v13 += (v13 + (v13 << ((4294967295) & 63)))
                        _ = fi28
                    }
                    v31 := GenS3{F0: -1, F1: 655}
                    _ = v31
                }
                v1 -= (^(v1))
            }
            _ = v13
        }
        var v32 string = ""
        _ = fi12
        _ = v32
    }
    var v33 float64 = 3.14159
    v33 = (1e10 * 0.1)
    v33 -= (-2.25 - (1e10 * (v33 + v33)))
    v33 += (v33 - float64((-2.25 / ((float32(v33) * 0.0) / 1e10))))
    if true {
        for fi34 := 0; fi34 < 18; fi34++ {
            for fi35 := 0; fi35 < 16; fi35++ {
                var v36 uint8 = 255
                _ = fi35
                _ = v36
            }
            fmt.Println(v33)
            _ = fi34
        }
    }
    v33 -= (1e10 + 0.1)
    var v37 float32 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v37
    _ = v4
    _ = v5
    _ = v8
}
