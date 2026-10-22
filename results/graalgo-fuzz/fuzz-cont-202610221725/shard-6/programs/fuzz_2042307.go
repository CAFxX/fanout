package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64; F1 uint8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int8; F1 int }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 string; F1 uint16; F2 bool }
func genFn0(p0 uint32, p1 uint) uint32 {
    for fi7 := 0; fi7 < 10; fi7++ {
        for fi8 := 0; fi8 < 8; fi8++ {
            var v9 uint8 = 0
            p0 -= (^(p0))
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 14; fi10++ {
            p1 = p1
            p1 += 65535
            _ = fi10
        }
        var v11 uint32 = p0
        _ = fi7
        _ = v11
    }
    v12 := []int{4294967295, int(p1), 128}
    var v13 uint = p1
    var v14 float32 = 1e10
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    _ = v14
    return 1
}
func genFn1(p0 uint8, p1 uint64, p2 int16) int32 {
    var v15 int16 = (p2 >> ((0) & 63))
    p0 = (p0 - (p0 ^ (p0 & (p0 & p0))))
    if false {
        v16 := []int{}
        _ = v16
    }
    fmt.Println(p2, p1)
    v17 := GenS1{F0: (v15 - 32767)}
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v17
    return 42
}

func main() {
    var v1 uint16 = 32767
    v2 := make(map[int]string)
    v3 := []uint32{0}
    v4 := make(map[uint32]bool)
    v5 := []int{}
    var v6 int8 = 127
    for fi18 := 0; fi18 < 16; fi18++ {
        if true {
            if true {
                switch "" {
                case "\n\t":
                    for fi19 := 0; fi19 < 13; fi19++ {
                        v20 := GenS1{F0: 1}
                        var v21 float64 = 0.1
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                case "hello":
                    var v22 uint8 = 1
                    _ = v22
                case "":
                    if true {
                        switch "\n\t" {
                        case "a":
                            v23 := GenS1{F0: -1}
                            _ = v23
                        case "foo bar":
                            v6 = v6
                        }
                        for fi24 := 0; fi24 < 19; fi24++ {
                            if false {
                                _ = len(v4)
                            } else {
                                v25 := GenS1{F0: 0}
                                fmt.Println(fi18, fi18, v6)
                                _ = v25
                            }
                            if (!(true || true)) {
                                v6 -= (v6 / (((v6 / ((v6) | 1))) | 1))
                                fmt.Println(fi18)
                                v1 = 65535
                            }
                            v26 := GenS2{F0: "foo bar"}
                            _ = fi24
                            _ = v26
                        }
                    }
                }
            }
        }
        _ = fi18
    }
    v6 = v6
    v1 = 390
    _ = len(v5)
    for fi27 := 0; fi27 < 11; fi27++ {
        for fi28 := 0; fi28 < 11; fi28++ {
            if (1 < 9223372036854775807) {
                v29 := v4[genFn0(genFn0(4294967295, 1), 443)]
                if (v29 && (315 >= 1)) {
                    v3 = append(v3, 0)
                    if true {
                        var v30 int64 = 9223372036854775807
                        v3 = append(v3, 2147483647)
                        v2[fi28] = "foo bar"
                        v30 = v30
                        _ = v30
                    }
                    v4[1000] = (false || v29)
                } else {
                    switch "a" {
                    case "":
                        delete(v2, fi28)
                    case "\n\t":
                        v6 = -1
                    case "a":
                        if v29 {
                            if v29 {
                                if ((0 > 6) && false) {
                                    switch "" {
                                    case "hello":
                                        v1 -= 1
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = 4294967295
                                    }
                                    var v31 uint = 2
                                    _ = v31
                                }
                                v32 := GenS2{F0: "foo bar"}
                                _ = v32
                            }
                            v2[(fi27 / (((fi27 ^ fi28)) | 1))] = ("hello" + "\n\t")
                            if v29 {
                                if false {
                                    switch (("" + ("" + "")) + "") {
                                    case "":
                                        var v33 int8 = (v6 << ((uint(v1)) & 63))
                                        _ = v33
                                    case "foo bar":
                                        var v34 float64 = 0.0
                                        _ = v34
                                    case "a":
                                        v35 := GenS3{F0: 461, F1: 127}
                                        _ = v35
                                    }
                                    v36 := GenS4{F0: 0}
                                    _ = v36
                                }
                                delete(v2, (fi28 | 32768))
                            } else {
                                if (!v29) {
                                    v1 += 954
                                } else {
                                    v6 = v6
                                    var v37 int32 = 0
                                    var v38 int64 = 9223372036854775807
                                    _ = v37
                                    _ = v38
                                }
                                var v39 uint8 = 2
                                _ = v39
                            }
                            v6 = (v6 & v6)
                            var v40 uint64 = 1
                            var v41 string = "\n\t"
                            _ = v40
                            _ = v41
                        }
                    }
                    if (!((v29 || v29) || (("foo bar" + "\n\t") != "foo bar"))) {
                        v1 = (v1 >> ((42) & 63))
                        v29 = v29
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        if ((((("\n\t" + (("" + "") + "a")) + "a") + "a") + ("foo bar" + (("foo bar" + (("\n\t" + "") + ("hello" + "foo bar"))) + ((("hello" + "a") + "hello") + "hello")))) < "a") {
                            if (0 > 255) {
                                _ = len(v3)
                                v29 = v29
                            } else {
                                var v42 uint32 = 2147483647
                                var v43 int32 = genFn1(0, 2, -32768)
                                _ = v42
                                _ = v43
                            }
                            fmt.Println(fi27)
                        }
                        v29 = v29
                        var v44 int32 = 2147483647
                        var v45 int64 = -628
                        v46 := v2[(fi28 + 605)]
                        _ = v44
                        _ = v45
                        _ = v46
                    }
                }
                fmt.Println(fi27)
                _ = v29
            } else {
                var v47 uint64 = 2
                _ = v47
            }
            _ = fi28
        }
        for fi48 := 0; fi48 < 5; fi48++ {
            if false {
                fmt.Println(v1, fi48, fi27)
                var v49 int32 = 265
                if len(v5) > 0 {
                    v5[0] = (fi48 | (fi27 | -9223372036854775808))
                }
                _ = v49
            }
            fmt.Println(v6, fi48, v1)
            _ = fi48
        }
        if false {
            switch (fi27 | fi27) {
            case -655:
                if (32767 != 32767) {
                    var v50 float64 = 3.14159
                    _ = v50
                }
            case -169:
                v51 := GenS5{F0: (v6 * -31), F1: (-(fi27))}
                _ = v51
            default:
                v52 := GenS2{F0: "a"}
                _ = v52
            }
        }
        _ = fi27
    }
    for fi53 := 0; fi53 < 10; fi53++ {
        if false {
            var v54 int64 = 42
            v1 = (v1 & (v1 / ((v1) | 1)))
            for fi55 := 0; fi55 < 20; fi55++ {
                fmt.Println(fi53)
                if false {
                    v56 := GenS5{F0: (v6 - (v6 | (v6 >> ((65535) & 63)))), F1: (fi53 & 42)}
                    _ = v56
                }
                _ = fi55
            }
            _ = v54
        } else {
            v57 := GenS6{F0: v6}
            _ = v57
        }
        var v58 int64 = -9223372036854775808
        var v59 uint = 2
        _ = fi53
        _ = v58
        _ = v59
    }
    var v60 int64 = 9223372036854775807
    v61 := GenS7{F0: "a", F1: v1, F2: (9223372036854775807 == 127)}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v60
    _ = v61
}
