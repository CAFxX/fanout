package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16; F1 string }
func genFn0(p0 int64) uint16 {
    var v7 uint = 0
    for fi8 := 0; fi8 < 18; fi8++ {
        if true {
            for fi9 := 0; fi9 < 20; fi9++ {
                p0 += -1
                var v10 float64 = -2.25
                fmt.Println(fi8, v10)
                _ = fi9
                _ = v10
            }
        } else {
            fmt.Println(p0, fi8, p0)
            var v11 bool = true
            v7 = (v7 ^ v7)
            _ = v11
        }
        _ = fi8
    }
    _ = p0
    _ = v7
    return 255
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint64 = 4294967295
    var v3 int16 = 32767
    var v4 float32 = (0.0 - (-2.25 + 1e10))
    var v5 int32 = 42
    var v6 int8 = 97
    v12 := []int{42, 2147483647, 42}
    for fi13 := 0; fi13 < 20; fi13++ {
        if false {
            for fi14 := 0; fi14 < 9; fi14++ {
                switch (fi14 % (((fi13 & fi13)) | 1)) {
                case -254:
                    switch 4294967295 {
                    case -1:
                        v2 += 0
                    case -214:
                        v15 := GenS1{F0: 0, F1: ((v2 + (v2 << ((4294967295) & 63))) >= v2)}
                        _ = v15
                    case -2147483648:
                        switch "\n\t" {
                        case "foo bar":
                            v16 := GenS2{F0: 0}
                            _ = v16
                        case "\n\t":
                            fmt.Println(v2, v2)
                        case "":
                            v1[42] = true
                            if ((v2 * 0) > (^(v2))) {
                                switch (fi13 * 32767) {
                                case 4294967295:
                                    v6 = (-(v6))
                                default:
                                    if true {
                                        switch "a" {
                                        case "foo bar":
                                            switch fi14 {
                                            case 4294967295:
                                                v4 += v4
                                            case 64:
                                                var v17 int16 = (-(v3))
                                                _ = v17
                                            default:
                                                var v18 uint32 = 252
                                                _ = v18
                                            }
                                            switch 220 {
                                            case -2147483648:
                                                var v19 int16 = (v3 % ((v3) | 1))
                                                _ = v19
                                            case 9223372036854775807:
                                                v6 += (v6 + (v6 >> ((65535) & 63)))
                                            case 0:
                                                var v20 int32 = 0
                                                _ = v20
                                            default:
                                                var v21 string = "hello"
                                                _ = v21
                                            }
                                        case "":
                                            var v22 int32 = (v5 % ((v5) | 1))
                                            _ = v22
                                        case "a":
                                            var v23 float64 = float64((v4 / 0.1))
                                            _ = v23
                                        default:
                                            v6 -= v6
                                        }
                                    }
                                }
                            }
                        }
                    default:
                        switch fi14 {
                        case 32767:
                            v24 := GenS2{F0: v5}
                            _ = v24
                        case 9223372036854775807:
                            fmt.Println(v6, fi13)
                        default:
                            if false {
                                switch 127 {
                                case 4294967295:
                                    var v25 uint = 51
                                    _ = v25
                                case 196:
                                    v6 = 40
                                default:
                                    v26 := GenS1{F0: 606, F1: false}
                                    _ = v26
                                }
                                fmt.Println(v5)
                            }
                        }
                    }
                }
                var v27 uint8 = 255
                var v28 int16 = v3
                var v29 int64 = -609
                var v30 uint32 = 1
                _ = fi14
                _ = v27
                _ = v28
                _ = v29
                _ = v30
            }
            for fi31 := 0; fi31 < 12; fi31++ {
                v32 := GenS2{F0: (^(v5))}
                if true {
                    var v33 uint8 = 255
                    v2 += v2
                    _ = v33
                } else {
                    v2 += 4294967295
                }
                delete(v1, 0)
                _ = fi31
                _ = v32
            }
            for fi34 := 0; fi34 < 9; fi34++ {
                var v35 string = "\n\t"
                var v36 int16 = v3
                _ = fi34
                _ = v35
                _ = v36
            }
        }
        _ = len(v12)
        switch (fi13 / (((fi13 ^ fi13)) | 1)) {
        case 17:
            for fi37 := 0; fi37 < 18; fi37++ {
                _ = len(v12)
                v38 := GenS1{F0: 862, F1: false}
                _ = fi37
                _ = v38
            }
        case -331:
            v39 := GenS1{F0: 18446744073709551615, F1: true}
            _ = v39
        case -58:
            fmt.Println(v5, fi13)
        default:
            if (32767 != 0) {
                var v40 uint16 = 65535
                var v41 uint16 = 2
                switch (((("\n\t" + ("hello" + ("\n\t" + "\n\t"))) + "\n\t") + "foo bar") + "hello") {
                case "hello":
                    switch "foo bar" {
                    case "a":
                        switch (fi13 * fi13) {
                        case -824:
                            if len(v12) >= 2 {
                                v12 = v12[1:len(v12)]
                            }
                        case -1:
                            fmt.Println(v41, v5, v4)
                        case 128:
                            v42 := GenS3{F0: (^(v40)), F1: "hello"}
                            _ = v42
                        }
                    }
                case "foo bar":
                    switch (fi13 << ((0) & 63)) {
                    case 32768:
                        fmt.Println(v5, v3)
                    default:
                        v5 = v5
                    }
                default:
                    v43 := GenS3{F0: 1, F1: ("\n\t" + ((("" + "") + ("\n\t" + (((("\n\t" + "\n\t") + ((("a" + ("foo bar" + "a")) + ("\n\t" + "a")) + ("a" + ("" + ("a" + ""))))) + "foo bar") + "foo bar"))) + "\n\t"))}
                    _ = v43
                }
                _ = v40
                _ = v41
            } else {
                v5 += v5
                fmt.Println(v2, v2, v3)
            }
        }
        _ = fi13
    }
    v1[65535] = true
    if (!(v3 > 0)) {
        var v44 float32 = v4
        var v45 int64 = -1
        _ = v44
        _ = v45
    } else {
        if false {
            switch ("\n\t" + "foo bar") {
            case "foo bar":
                delete(v1, 593)
            case "\n\t":
                var v46 int = 32767
                _ = v46
            default:
                v47 := GenS2{F0: (v5 + (-(v5)))}
                _ = v47
            }
        } else {
            for fi48 := 0; fi48 < 13; fi48++ {
                v4 = float32(1.5)
                fmt.Println(v5, v3)
                _ = fi48
            }
            var v49 float32 = v4
            _ = v49
        }
        v3 += v3
    }
    var v50 uint32 = 2
    v51 := GenS1{F0: 4294967295, F1: true}
    v4 = (v4 / ((-2.25 * 1.5) - ((v4 + (float32((1e10 + 3.14159)) * v4)) + (float32(((3.14159 - (float64(v4) - 0.1)) * -2.25)) * (v4 + v4)))))
    for fi52 := 0; fi52 < 4; fi52++ {
        var v53 int16 = (^(v3))
        v1[v50] = false
        _ = fi52
        _ = v53
    }
    v2 += 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v6
}
