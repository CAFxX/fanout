package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 uint8; F1 bool; F2 uint16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint64; F1 uint16; F2 uint32 }
type GenS6 struct { F0 uint64; F1 bool; F2 int32 }
type GenS7 struct { F0 int16; F1 bool }
type GenS8 struct { F0 bool }
func genFn0(p0 bool) uint64 {
    fmt.Println(p0)
    v5 := []int{-742, -578, -2147483648, 42}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v6 int16 = 1
    switch 127 {
    case 42:
        v7 := GenS1{F0: true, F1: "foo bar"}
        _ = v7
    case -609:
        fmt.Println(p0, p0)
    case 543:
        switch 345 {
        case 765:
            for fi8 := 0; fi8 < 17; fi8++ {
                var v9 int8 = -1
                v9 = v9
                _ = fi8
                _ = v9
            }
        case 2147483647:
            v10 := GenS1{F0: false, F1: "a"}
            _ = v10
        default:
            switch "foo bar" {
            case "\n\t":
                var v11 uint32 = 65535
                _ = v11
            case "hello":
                for fi12 := 0; fi12 < 10; fi12++ {
                    v13 := GenS2{F0: -1, F1: -117}
                    v14 := GenS2{F0: -9223372036854775808, F1: -25}
                    var v15 int64 = 9223372036854775807
                    _ = fi12
                    _ = v13
                    _ = v14
                    _ = v15
                }
            default:
                for fi16 := 0; fi16 < 6; fi16++ {
                    switch "a" {
                    case "foo bar":
                        v17 := GenS2{F0: -9223372036854775808, F1: 1}
                        _ = v17
                    case "hello":
                        var v18 int8 = int8(v6)
                        _ = v18
                    default:
                        for fi19 := 0; fi19 < 17; fi19++ {
                            var v20 bool = (42 <= 858)
                            v21 := GenS2{F0: -1, F1: 1}
                            _ = fi19
                            _ = v20
                            _ = v21
                        }
                    }
                    v22 := GenS3{F0: 127, F1: true, F2: 266}
                    _ = fi16
                    _ = v22
                }
            }
        }
    }
    _ = p0
    _ = v5
    _ = v6
    return 18446744073709551615
}

func main() {
    var v1 float64 = (float64(float32(3.14159)) - 1e10)
    var v2 float32 = float32(v1)
    v3 := make(map[string]int)
    var v4 int16 = 686
    v1 = float64(0.0)
    v23 := []int{-564, 128}
    if false {
        var v24 int8 = 47
        switch (("hello" + "foo bar") + "hello") {
        case "a":
            v25 := GenS2{F0: -102, F1: (v24 * (v24 - (v24 * 127)))}
            _ = v25
        case "\n\t":
            if (-1 == -2) {
                if (true || (4294967295 < 0)) {
                    switch "hello" {
                    case "hello":
                        v1 = ((v1 - (1.5 - (v1 / v1))) / float64(1.5))
                    case "\n\t":
                        if true {
                            if (576 < 42) {
                                var v26 uint64 = 208
                                if true {
                                    v3["\n\t"] = -2147483648
                                } else {
                                    if false {
                                        v27 := v3[(((("\n\t" + "\n\t") + "foo bar") + "") + "foo bar")]
                                        _ = v27
                                    } else {
                                        var v28 float64 = v1
                                        _ = v28
                                    }
                                }
                                var v29 uint32 = 42
                                _ = v26
                                _ = v29
                            } else {
                                fmt.Println(v1)
                                v24 = (-(v24))
                            }
                            for fi30 := 0; fi30 < 8; fi30++ {
                                fmt.Println(v4, v4)
                                _ = fi30
                            }
                            fmt.Println(v4, v4, v1)
                        }
                    case "foo bar":
                        v31 := GenS1{F0: true, F1: ("hello" + "\n\t")}
                        _ = v31
                    default:
                        fmt.Println(v24)
                    }
                    var v32 int32 = 611
                    fmt.Println(v1)
                    _ = v32
                } else {
                    v33 := GenS2{F0: -780, F1: (v24 % (((v24 >> ((661) & 63))) | 1))}
                    v34 := GenS4{F0: (v4 & 0)}
                    if len(v23) > 0 {
                        v23[0] = 169
                    }
                    _ = v33
                    _ = v34
                }
                var v35 int8 = 1
                _ = v35
            } else {
                for fi36 := 0; fi36 < 4; fi36++ {
                    var v37 uint32 = 2
                    _ = fi36
                    _ = v37
                }
            }
        default:
            var v38 uint32 = 1
            _ = v38
        }
        _ = v24
    }
    v39 := GenS4{F0: v4}
    var v40 bool = true
    fmt.Println(v40)
    fmt.Println(v40)
    if v40 {
        for fi41 := 0; fi41 < 20; fi41++ {
            for fi42 := 0; fi42 < 1; fi42++ {
                if (!v40) {
                    var v43 int64 = 813
                    v44 := GenS5{F0: genFn0(v40), F1: 1, F2: 1}
                    var v45 uint32 = 636
                    v2 += (v2 - ((0.0 + (((v2 / -2.25) / v2) - float32(v1))) / 0.1))
                    _ = v43
                    _ = v44
                    _ = v45
                } else {
                    var v46 uint = 4294967295
                    _ = v46
                }
                v47 := GenS6{F0: 2, F1: (42 < 0), F2: 2147483647}
                if v40 {
                    var v48 float64 = 3.14159
                    var v49 int16 = (v4 >> ((18446744073709551615) & 63))
                    v50 := GenS3{F0: 93, F1: (-2147483648 > (fi42 << ((255) & 63))), F2: 42}
                    _ = v48
                    _ = v49
                    _ = v50
                } else {
                    v51 := GenS1{F0: (!(!("hello" >= "a"))), F1: ("a" + "hello")}
                    var v52 string = ("foo bar" + "foo bar")
                    _ = v51
                    _ = v52
                }
                v53 := GenS3{F0: 42, F1: (genFn0(v40) <= 2), F2: 0}
                fmt.Println(fi42, v4, v1)
                _ = fi42
                _ = v47
                _ = v53
            }
            if v40 {
                v1 = v1
                if false {
                    var v54 bool = (1 > 2147483647)
                    _ = v54
                } else {
                    v1 += float64(float32(v1))
                }
                if (1 == 42) {
                    if v40 {
                        v55 := GenS6{F0: 937, F1: v40, F2: 129}
                        v1 -= float64((1.5 - ((v2 / (1e10 * (float32(v1) / (v2 / v2)))) + (v2 + float32(1.5)))))
                        var v56 bool = (false || v40)
                        _ = v55
                        _ = v56
                    }
                    switch 32768 {
                    case -719:
                        v2 = v2
                    }
                    if v40 {
                        if len(v23) >= 2 {
                            v23 = v23[1:len(v23)]
                        }
                        fmt.Println(v4, v4)
                    }
                }
            } else {
                v1 = 0.0
                v57 := GenS7{F0: (v4 >> ((255) & 63)), F1: (!v40)}
                if len(v23) >= 2 {
                    v23 = v23[1:len(v23)]
                }
                _ = v57
            }
            v1 -= 3.14159
            _ = fi41
        }
    }
    v58 := GenS3{F0: 0, F1: (!true), F2: 2}
    if len(v23) >= 2 {
        v23 = v23[1:len(v23)]
    }
    var v59 uint8 = 131
    fmt.Println(v59)
    var v60 int64 = 1
    v40 = (885 < (v4 % ((v4) | 1)))
    v61 := GenS8{F0: true}
    v1 -= float64(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v58
    _ = v59
    _ = v60
    _ = v61
}
