package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16; F1 int32; F2 int64 }
type GenS4 struct { F0 uint64; F1 int8; F2 uint16 }
type GenS5 struct { F0 int16; F1 int }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int64; F1 uint32; F2 string }
func genFn0(p0 uint16) uint8 {
    var v6 uint64 = 321
    v7 := []int{32767}
    v8 := GenS1{F0: -929, F1: (p0 | (p0 << ((308) & 63)))}
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 127
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]int)
    var v3 int64 = -9223372036854775808
    v4 := make(map[string]uint32)
    v5 := []string{("hello" + "foo bar")}
    for fi9 := 0; fi9 < 9; fi9++ {
        for fi10 := 0; fi10 < 16; fi10++ {
            v3 = v3
            var v11 int = -2147483648
            var v12 uint64 = 974
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi9
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v5 = append(v5, "hello")
    v3 = (^(v3))
    fmt.Println(v3)
    switch (("a" + "\n\t") + "hello") {
    case "foo bar":
        if (false || (255 != 255)) {
            v3 = (^(v3))
            v3 = v3
            for fi13 := 0; fi13 < 8; fi13++ {
                if (742 > 0) {
                    v3 = v3
                }
                _ = fi13
            }
        } else {
            if true {
                fmt.Println(v3)
            } else {
                v14 := GenS2{F0: -615}
                v15 := v4["hello"]
                _ = v14
                _ = v15
            }
            v3 = (v3 * (v3 * (v3 + 9223372036854775807)))
            for fi16 := 0; fi16 < 1; fi16++ {
                switch (fi16 * 4294967295) {
                case -9223372036854775808:
                    v3 = 1
                }
                switch "hello" {
                case "foo bar":
                    v3 -= v3
                case "\n\t":
                    var v17 int16 = -1
                    _ = v17
                case "a":
                    fmt.Println(fi16, fi16)
                }
                fmt.Println(fi16)
                _ = fi16
            }
        }
    case "":
        if ((!(357 < 1)) && true) {
            v18 := GenS2{F0: 42}
            v3 = v3
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = v18
        } else {
            v3 += (v3 | -1)
        }
    default:
        switch "foo bar" {
        case "foo bar":
            var v19 int8 = 106
            _ = v19
        default:
            if false {
                if true {
                    for fi20 := 0; fi20 < 19; fi20++ {
                        for fi21 := 0; fi21 < 13; fi21++ {
                            v3 -= (v3 / ((-1) | 1))
                            if false {
                                v22 := GenS3{F0: -924, F1: 0, F2: v3}
                                _ = len(v2)
                                if (int16(v3) >= 1) {
                                    if ((false && true) && true) {
                                        fmt.Println(fi20, fi20, v3)
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    }
                                    fmt.Println(fi21)
                                    if (true && true) {
                                        if ((!false) || false) {
                                            var v23 string = ""
                                            _ = v23
                                        }
                                        if ("\n\t" != "\n\t") {
                                            v3 += v3
                                            var v24 uint16 = 32767
                                            _ = v24
                                        }
                                    } else {
                                        switch (fi21 & fi20) {
                                        case 42:
                                            var v25 int16 = -908
                                            _ = v25
                                        case 127:
                                            fmt.Println(fi20, fi21)
                                        default:
                                            var v26 uint16 = 457
                                            var v27 uint = 2
                                            if (v26 <= 0) {
                                                v26 = 255
                                            } else {
                                                v26 = (^(v26))
                                                fmt.Println(fi21, v27)
                                                if (true || (((^(v3)) < -535) && true)) {
                                                    v3 -= v3
                                                }
                                            }
                                            _ = v26
                                            _ = v27
                                        }
                                        v3 -= (v3 & (v3 >> ((241) & 63)))
                                    }
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                } else {
                                    if false {
                                        var v28 uint32 = 42
                                        _ = v28
                                    } else {
                                        v29 := GenS4{F0: 553, F1: -98, F2: 32767}
                                        fmt.Println(fi21, fi21)
                                        _ = v29
                                    }
                                    v3 += 9223372036854775807
                                    switch fi20 {
                                    case -1:
                                        fmt.Println(v3)
                                    case 1:
                                        v3 = v3
                                    }
                                }
                                v3 -= (v3 % (((v3 >> ((1) & 63))) | 1))
                                switch -1 {
                                case -9223372036854775808:
                                    var v30 float64 = 1.5
                                    _ = v30
                                case -2147483648:
                                    v5 = append(v5, ("\n\t" + "\n\t"))
                                default:
                                    var v31 uint64 = 360
                                    _ = v31
                                }
                                var v32 string = "hello"
                                v33 := GenS5{F0: 1000, F1: fi20}
                                if (!false) {
                                    var v34 int = (fi20 << ((764) & 63))
                                    fmt.Println(fi21)
                                    _ = v34
                                } else {
                                    var v35 int16 = -32768
                                    var v36 uint32 = 2147483648
                                    v37 := GenS4{F0: 962, F1: 100, F2: 2}
                                    _ = v35
                                    _ = v36
                                    _ = v37
                                }
                                _ = v22
                                _ = v32
                                _ = v33
                            }
                            _ = fi21
                        }
                        _ = len(v4)
                        _ = fi20
                    }
                    var v38 bool = true
                    _ = v38
                } else {
                    v3 = v3
                    if (65535 <= 42) {
                        if (0.1 >= float64((float32(0.0) + 0.0))) {
                            fmt.Println(v3)
                        } else {
                            v39 := GenS5{F0: -32768, F1: 4294967295}
                            _ = v39
                        }
                    }
                    v40 := GenS6{F0: false}
                    _ = v40
                }
                var v41 int64 = v3
                v42 := GenS5{F0: int16(v3), F1: 127}
                _ = v41
                _ = v42
            } else {
                v43 := GenS1{F0: -2147483648, F1: 0}
                v44 := v4["hello"]
                v44 = (v44 << ((213) & 63))
                _ = v43
                _ = v44
            }
        }
    }
    v45 := GenS7{F0: v3, F1: 490, F2: "a"}
    fmt.Println(v3)
    var v46 int64 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
}
