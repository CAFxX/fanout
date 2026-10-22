package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 string }
type GenS2 struct { F0 string; F1 bool; F2 uint8 }
type GenS3 struct { F0 int; F1 int16; F2 uint32 }
type GenS4 struct { F0 bool; F1 int16 }
type GenS5 struct { F0 uint64; F1 uint }
type GenS6 struct { F0 string; F1 int }
type GenS7 struct { F0 uint32; F1 uint32 }
type GenS8 struct { F0 uint32; F1 uint64; F2 int8 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int16; F1 int64 }
func genFn0(p0 uint16, p1 uint8, p2 int32) int8 {
    fmt.Println(p0, p0)
    if false {
        switch 2147483647 {
        case 375:
            v4 := GenS1{F0: (^(p0)), F1: 42, F2: "hello"}
            _ = v4
        default:
            p2 += (p2 << ((915) & 63))
        }
        p0 = 0
    }
    var v5 int64 = 899
    var v6 bool = true
    fmt.Println(p2, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return -128
}
func genFn1(p0 uint32, p1 int8, p2 bool) int32 {
    p0 += 668
    v7 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -504
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]int)
    v3 := make(map[uint32]string)
    v8 := []int{1, -1, 32768, -1}
    switch 9223372036854775807 {
    case 42:
        v9 := GenS1{F0: 255, F1: 18446744073709551615, F2: "\n\t"}
        _ = v9
    case 128:
        var v10 bool = true
        _ = v10
    case -1:
        switch "hello" {
        case "foo bar":
            for fi11 := 0; fi11 < 4; fi11++ {
                var v12 uint64 = 1
                switch (("\n\t" + ((("\n\t" + "a") + "") + ("\n\t" + "\n\t"))) + "\n\t") {
                case "foo bar":
                    for fi13 := 0; fi13 < 11; fi13++ {
                        delete(v1, ("foo bar" + ""))
                        _ = fi13
                    }
                case "hello":
                    for fi14 := 0; fi14 < 15; fi14++ {
                        v12 = 0
                        delete(v2, (^(fi14)))
                        _ = fi14
                    }
                case "a":
                    v15 := GenS2{F0: (("a" + ("hello" + ("\n\t" + "\n\t"))) + "foo bar"), F1: (true || true), F2: 128}
                    _ = v15
                }
                _ = fi11
                _ = v12
            }
        default:
            v16 := GenS3{F0: 32767, F1: -869, F2: 0}
            _ = v16
        }
    }
    delete(v3, 809)
    var v17 int32 = -897
    for fi18 := 0; fi18 < 8; fi18++ {
        for fi19 := 0; fi19 < 11; fi19++ {
            v2[fi19] = 32768
            var v20 int = fi18
            v20 = (fi19 & v20)
            fmt.Println(fi19, v17)
            _ = fi19
            _ = v20
        }
        _ = fi18
    }
    v21 := GenS2{F0: (("hello" + "\n\t") + "a"), F1: (89 < 42), F2: 1}
    switch "foo bar" {
    case "\n\t":
        v22 := GenS3{F0: -266, F1: 1000, F2: 1}
        _ = v22
    case "hello":
        v17 += genFn1(42, genFn0(241, 179, v17), true)
    case "foo bar":
        v23 := GenS1{F0: 1, F1: 564, F2: "a"}
        _ = v23
    default:
        var v24 uint64 = 0
        _ = v24
    }
    for fi25 := 0; fi25 < 5; fi25++ {
        v26 := v2[(fi25 / (((fi25 | fi25)) | 1))]
        _ = fi25
        _ = v26
    }
    if true {
        if false {
            if true {
                if false {
                    v17 -= genFn1(579, 127, false)
                } else {
                    v17 -= (-(v17))
                    if len(v8) > 0 {
                        v8[0] = 32768
                    }
                    for fi27 := 0; fi27 < 14; fi27++ {
                        v28 := GenS3{F0: (fi27 << ((211) & 63)), F1: 587, F2: 1}
                        if (((("a" + "\n\t") + "foo bar") + "hello") > "hello") {
                            v29 := GenS4{F0: false, F1: -32768}
                            if false {
                                if len(v8) >= 2 {
                                    v8 = v8[1:len(v8)]
                                }
                                for fi30 := 0; fi30 < 2; fi30++ {
                                    var v31 uint16 = 814
                                    switch 9223372036854775807 {
                                    case 950:
                                        v31 = (v31 * v31)
                                    }
                                    _ = fi30
                                    _ = v31
                                }
                                v32 := v3[2147483647]
                                _ = v32
                            } else {
                                var v33 int64 = 310
                                _ = v33
                            }
                            _ = v29
                        }
                        fmt.Println(v17, v17)
                        _ = fi27
                        _ = v28
                    }
                }
                fmt.Println(v17)
                switch "foo bar" {
                case "foo bar":
                    v17 = v17
                case "a":
                    var v34 string = ("\n\t" + ("a" + ("foo bar" + (("\n\t" + "hello") + "a"))))
                    _ = v34
                }
            }
        }
        if (!(-334 <= 1)) {
            if ((false && (!(!true))) || false) {
                v35 := GenS1{F0: 42, F1: 4294967295, F2: "hello"}
                _ = v35
            } else {
                v36 := v1["a"]
                v37 := GenS5{F0: 4294967295, F1: uint(v36)}
                _ = v36
                _ = v37
            }
            v17 = v17
            v38 := GenS6{F0: "foo bar", F1: 9223372036854775807}
            _ = v38
        } else {
            if (42 >= -1) {
                switch -9223372036854775808 {
                case -2147483648:
                    v39 := GenS5{F0: 2, F1: 255}
                    _ = v39
                default:
                    var v40 uint = 0
                    _ = v40
                }
                fmt.Println(v17)
            }
            var v41 string = "\n\t"
            v42 := v3[135]
            _ = v41
            _ = v42
        }
        fmt.Println(v17)
    } else {
        v17 += (v17 | (v17 * (v17 << ((18446744073709551615) & 63))))
        v43 := GenS6{F0: "foo bar", F1: 9223372036854775807}
        _ = v43
    }
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    var v44 int = 479
    var v45 uint32 = 28
    v17 = 1
    if (!(!((0 > 1) || (!false)))) {
        v44 = (-(v44))
        if len(v8) > 0 {
            v8[0] = v44
        }
        if (2 <= 0) {
            for fi46 := 0; fi46 < 9; fi46++ {
                if (1e10 <= (1e10 * -2.25)) {
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    for fi47 := 0; fi47 < 2; fi47++ {
                        var v48 float32 = (3.14159 / 1.0)
                        v49 := GenS7{F0: v45, F1: (v45 >> ((463) & 63))}
                        _ = fi47
                        _ = v48
                        _ = v49
                    }
                } else {
                    delete(v3, 42)
                    delete(v2, (fi46 / ((v44) | 1)))
                    fmt.Println(v44)
                }
                if false {
                    if false {
                        switch (v44 / (((^(fi46))) | 1)) {
                        case 92:
                            switch ((("hello" + "") + "\n\t") + "a") {
                            case "a":
                                v44 += (v44 - (v44 * (v44 >> ((1) & 63))))
                            }
                        case 1:
                            switch "hello" {
                            case "hello":
                                v45 += (v45 | 333)
                            case "":
                                var v50 float32 = 3.14159
                                _ = v50
                            case "\n\t":
                                v51 := GenS8{F0: (v45 / (((v45 % ((v45) | 1))) | 1)), F1: 148, F2: int8(v17)}
                                _ = v51
                            }
                        case 32767:
                            var v52 float32 = 0.1
                            _ = v52
                        default:
                            v44 += (v44 - v44)
                        }
                    }
                }
                var v53 int8 = 51
                _ = fi46
                _ = v53
            }
            v54 := GenS2{F0: "\n\t", F1: true, F2: 255}
            var v55 string = "a"
            _ = v54
            _ = v55
        }
    } else {
        v56 := GenS7{F0: 4294967295, F1: (v45 / (((^(v45))) | 1))}
        v17 = -1
        _ = v56
    }
    for fi57 := 0; fi57 < 8; fi57++ {
        v58 := GenS9{F0: 4294967295}
        if true {
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
        } else {
            for fi59 := 0; fi59 < 19; fi59++ {
                if ((0.0 - 1e10) <= -2.25) {
                    var v60 float32 = -2.25
                    if (1 <= -468) {
                        var v61 float32 = v60
                        v60 += (v60 / 3.14159)
                        _ = v61
                    }
                    _ = v60
                }
                v44 = fi57
                v62 := GenS10{F0: 42, F1: -1}
                fmt.Println(v44, v44, v17)
                _ = fi59
                _ = v62
            }
            switch fi57 {
            case 807:
                v45 = (^(v45))
            }
        }
        fmt.Println(fi57)
        _ = fi57
        _ = v58
    }
    var v63 uint16 = 605
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v63:", v63)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v44
    _ = v45
    _ = v63
    _ = v8
}
