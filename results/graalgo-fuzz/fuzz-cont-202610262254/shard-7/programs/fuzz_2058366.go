package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32; F1 uint }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 int; F1 uint16; F2 uint64 }
func genFn0(p0 bool, p1 uint64, p2 uint8) uint64 {
    var v6 float32 = 3.14159
    v7 := GenS1{F0: "a", F1: p0}
    p2 = (p2 - p2)
    v8 := GenS2{F0: -2147483648}
    p1 = (^(p1))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return (p1 | 0)
}

func main() {
    var v1 int64 = 42
    v2 := make(map[int]string)
    v3 := []uint32{2147483648, 2147483648, 0}
    var v4 string = ((("foo bar" + "") + ("hello" + ("" + "foo bar"))) + (("a" + "") + "\n\t"))
    var v5 bool = (v4 >= v4)
    if v5 {
        delete(v2, 2147483647)
        v4 = v4
        if (!true) {
            switch -191 {
            case 9223372036854775807:
                var v9 string = ((v4 + v4) + v4)
                _ = v9
            case 32767:
                var v10 float64 = 1.5
                _ = v10
            default:
                var v11 int32 = -513
                _ = v11
            }
            var v12 int32 = -1
            _ = v12
        } else {
            for fi13 := 0; fi13 < 2; fi13++ {
                switch (-(fi13)) {
                case 4:
                    fmt.Println(fi13, fi13)
                }
                switch (fi13 / ((fi13) | 1)) {
                case 1:
                    for fi14 := 0; fi14 < 9; fi14++ {
                        v2[fi14] = (v4 + ((v4 + "foo bar") + "\n\t"))
                        var v15 float64 = 1e10
                        _ = fi14
                        _ = v15
                    }
                case 4294967295:
                    var v16 int = (fi13 ^ (fi13 % ((fi13) | 1)))
                    _ = v16
                default:
                    if (-2.25 == (0.1 * -2.25)) {
                        for fi17 := 0; fi17 < 3; fi17++ {
                            var v18 int16 = 1000
                            _ = fi17
                            _ = v18
                        }
                        switch fi13 {
                        case -650:
                            var v19 int8 = 57
                            _ = v19
                        case 42:
                            for fi20 := 0; fi20 < 12; fi20++ {
                                if (123 == 777) {
                                    v4 = (v4 + v4)
                                    v5 = true
                                }
                                switch fi20 {
                                case -9223372036854775808:
                                    fmt.Println(v5, v1)
                                case 9223372036854775807:
                                    if ((((0.1 - (((1e10 - 3.14159) + -2.25) - 0.0)) - (-2.25 / 1.5)) * ((0.1 + float64(float32(0.0))) * (1e10 * (3.14159 / 1e10)))) < 1.5) {
                                        v2[1] = v4
                                        if false {
                                            switch fi20 {
                                            case -43:
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            case 0:
                                                var v21 uint = 18446744073709551615
                                                _ = v21
                                            case 2147483647:
                                                fmt.Println(fi20)
                                                var v22 uint16 = 32767
                                                _ = v22
                                            default:
                                                var v23 bool = v5
                                                _ = v23
                                            }
                                            switch 32768 {
                                            case -497:
                                                switch v4 {
                                                case "a":
                                                    var v24 uint16 = 32767
                                                    _ = v24
                                                default:
                                                    v1 = (v1 + (-(v1)))
                                                }
                                            case -9223372036854775808:
                                                fmt.Println(v1)
                                            case -239:
                                                var v25 bool = false
                                                _ = v25
                                            }
                                            v26 := GenS3{F0: 0, F1: 0}
                                            _ = v26
                                        }
                                        if (801 <= 0) {
                                            var v27 bool = v5
                                            var v28 int32 = 517
                                            var v29 uint = 255
                                            delete(v2, fi20)
                                            _ = v27
                                            _ = v28
                                            _ = v29
                                        }
                                    } else {
                                        var v30 int8 = -1
                                        _ = v30
                                    }
                                case 32767:
                                    switch (fi20 | (fi20 * (^(fi20)))) {
                                    case -9223372036854775808:
                                        if true {
                                            if (1.5 < (float32((0.1 * 1e10)) + 1e10)) {
                                                v4 = v4
                                                var v31 uint64 = 0
                                                fmt.Println(v4, v1)
                                                _ = v31
                                            }
                                        }
                                    case 32768:
                                        var v32 uint64 = 0
                                        _ = v32
                                    default:
                                        fmt.Println(v1, v5, fi20)
                                    }
                                }
                                if (42 >= 1) {
                                    var v33 string = ((("a" + "a") + ((v4 + v4) + v4)) + v4)
                                    v5 = (255 != 42)
                                    delete(v2, fi20)
                                    _ = v33
                                }
                                v1 = (v1 >> ((848) & 63))
                                _ = fi20
                            }
                        }
                        v34 := GenS1{F0: "hello", F1: v5}
                        _ = v34
                    }
                }
                fmt.Println(v1, v4, v4)
                _ = fi13
            }
        }
    }
    v3 = append(v3, 1)
    v4 = (v4 + v4)
    fmt.Println(v1)
    var v35 bool = false
    v4 = ("" + v4)
    for fi36 := 0; fi36 < 15; fi36++ {
        v37 := GenS3{F0: 2, F1: 65535}
        fmt.Println(v1, v4, v4)
        _ = fi36
        _ = v37
    }
    var v38 uint64 = 143
    v35 = v5
    v39 := GenS4{F0: (v38 % ((v38) | 1)), F1: 128}
    var v40 int8 = 1
    v41 := GenS1{F0: (v4 + v4), F1: (!v5)}
    switch (v4 + "\n\t") {
    case "\n\t":
        if (72 <= 281) {
            switch v4 {
            case "a":
                if v5 {
                    if false {
                        for fi42 := 0; fi42 < 13; fi42++ {
                            v1 = v1
                            v5 = ((v38 * 1) < 1)
                            var v43 int8 = (v40 + (v40 ^ (v40 % ((v40) | 1))))
                            _ = fi42
                            _ = v43
                        }
                        for fi44 := 0; fi44 < 18; fi44++ {
                            var v45 float64 = 0.0
                            v2[-1] = v4
                            v45 = v45
                            _ = fi44
                            _ = v45
                        }
                        v40 = v40
                    } else {
                        v2[4294967295] = v4
                        v5 = v35
                        v46 := GenS3{F0: 2147483648, F1: 4294967295}
                        _ = v46
                    }
                    for fi47 := 0; fi47 < 20; fi47++ {
                        for fi48 := 0; fi48 < 19; fi48++ {
                            if len(v3) > 0 {
                                v3[0] = 65535
                            }
                            v38 -= 4294967295
                            _ = fi48
                        }
                        v49 := GenS4{F0: v38, F1: 94}
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi47
                        _ = v49
                    }
                    for fi50 := 0; fi50 < 8; fi50++ {
                        var v51 float32 = 1e10
                        _ = fi50
                        _ = v51
                    }
                } else {
                    v1 -= 581
                }
            case "\n\t":
                switch ("a" + ("\n\t" + ("" + v4))) {
                case "foo bar":
                    for fi52 := 0; fi52 < 16; fi52++ {
                        v53 := GenS2{F0: fi52}
                        var v54 float64 = (0.1 * 0.1)
                        _ = fi52
                        _ = v53
                        _ = v54
                    }
                case "a":
                    v55 := GenS5{F0: 32768, F1: 1, F2: 298}
                    _ = v55
                }
            }
        }
    case "hello":
        if (-616 != 127) {
            for fi56 := 0; fi56 < 16; fi56++ {
                for fi57 := 0; fi57 < 19; fi57++ {
                    var v58 uint = uint(v38)
                    _ = fi57
                    _ = v58
                }
                _ = fi56
            }
        } else {
            fmt.Println(v1)
            var v59 uint8 = 0
            var v60 int8 = v40
            _ = v59
            _ = v60
        }
    case "a":
        if false {
            if v35 {
                v3 = append(v3, 42)
            } else {
                var v61 float64 = 1.5
                _ = v61
            }
            v4 = ((v4 + v4) + v4)
            if v5 {
                if v5 {
                    fmt.Println(v1, v35)
                    switch -1 {
                    case -9223372036854775808:
                        var v62 uint16 = 42
                        _ = v62
                    default:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                }
            } else {
                v40 -= (v40 % ((v40) | 1))
            }
        }
    default:
        v38 = (v38 | 4294967295)
    }
    var v63 string = "a"
    for fi64 := 0; fi64 < 14; fi64++ {
        var v65 int64 = (v1 / ((v1) | 1))
        v35 = (int64(v38) == v1)
        _ = fi64
        _ = v65
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v63
}
