package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 int64; F2 int32 }
type GenS3 struct { F0 int32; F1 uint64; F2 uint64 }
type GenS4 struct { F0 int }
func genFn0() int16 {
    for fi5 := 0; fi5 < 18; fi5++ {
        for fi6 := 0; fi6 < 19; fi6++ {
            var v7 int16 = 32767
            _ = fi6
            _ = v7
        }
        fmt.Println(fi5)
        _ = fi5
    }
    v8 := GenS1{F0: 128}
    var v9 float32 = 0.1
    _ = v8
    _ = v9
    return -645
}

func main() {
    var v1 uint16 = 1
    var v2 uint16 = v1
    v3 := make(map[uint32]string)
    v4 := []string{"hello", ("" + ("a" + ("hello" + "\n\t")))}
    if true {
        fmt.Println(v2)
        v10 := GenS2{F0: (v1 / ((223) | 1)), F1: 196, F2: -1}
        _ = v10
    } else {
        v11 := GenS2{F0: (^(v1)), F1: 1, F2: -262}
        _ = v11
    }
    var v12 int8 = -128
    v13 := GenS1{F0: 40}
    v14 := GenS3{F0: 42, F1: 546, F2: 4294967295}
    if (4294967295 != 107) {
        if (-1 != 0) {
            var v15 float32 = (-2.25 + (1e10 + 3.14159))
            if ((^(v2)) > (v2 * v2)) {
                v15 = float32(0.1)
            } else {
                delete(v3, 1)
                if (true || false) {
                    for fi16 := 0; fi16 < 7; fi16++ {
                        v1 -= (v2 << ((65535) & 63))
                        _ = fi16
                    }
                    for fi17 := 0; fi17 < 1; fi17++ {
                        delete(v3, 1)
                        _ = fi17
                    }
                    switch 32768 {
                    case -1:
                        var v18 float32 = v15
                        _ = v18
                    default:
                        for fi19 := 0; fi19 < 8; fi19++ {
                            v12 = v12
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            var v20 uint32 = 675
                            _ = fi19
                            _ = v20
                        }
                    }
                }
                v12 = (v12 + (v12 & (v12 << ((258) & 63))))
            }
            _ = v15
        }
    }
    var v21 bool = false
    v22 := GenS2{F0: 2, F1: -1, F2: 0}
    v21 = false
    switch ((("foo bar" + ("" + "hello")) + (((("foo bar" + "a") + ("a" + "foo bar")) + "") + "\n\t")) + "\n\t") {
    case "foo bar":
        if (2147483648 >= 2147483648) {
            for fi23 := 0; fi23 < 8; fi23++ {
                if true {
                    v24 := GenS1{F0: 90}
                    _ = len(v3)
                    _ = v24
                }
                v3[984] = "a"
                for fi25 := 0; fi25 < 10; fi25++ {
                    v26 := GenS1{F0: 128}
                    _ = fi25
                    _ = v26
                }
                _ = fi23
            }
        }
    case "a":
        var v27 uint16 = 23
        _ = v27
    default:
        if v21 {
            switch 42 {
            case 403:
                var v28 uint32 = 65535
                _ = v28
            case 2147483647:
                if ((42 <= v12) && v21) {
                    for fi29 := 0; fi29 < 14; fi29++ {
                        var v30 uint16 = (v2 << ((18446744073709551615) & 63))
                        if v21 {
                            fmt.Println(v21, v2)
                            v30 += 32767
                            v30 = (v1 + v30)
                        }
                        _ = fi29
                        _ = v30
                    }
                    v4 = append(v4, (("foo bar" + "\n\t") + "\n\t"))
                    v31 := GenS2{F0: v1, F1: -818, F2: -1}
                    _ = v31
                } else {
                    switch "" {
                    case "foo bar":
                        if v21 {
                            var v32 uint32 = 65535
                            switch "\n\t" {
                            case "hello":
                                if true {
                                    v32 += v32
                                    v2 = v2
                                }
                            case "\n\t":
                                v4 = append(v4, "a")
                            case "":
                                if len(v4) > 0 {
                                    v4[0] = (((("\n\t" + "a") + "hello") + "hello") + "")
                                }
                            default:
                                switch ("\n\t" + ("\n\t" + "foo bar")) {
                                case "a":
                                    var v33 bool = (float64(1.5) <= 3.14159)
                                    _ = v33
                                default:
                                    var v34 uint8 = 127
                                    _ = v34
                                }
                            }
                            if v21 {
                                fmt.Println(v12)
                            } else {
                                v2 -= 32767
                                v1 = 255
                            }
                            _ = v32
                        }
                    case "":
                        v21 = v21
                    }
                    var v35 uint32 = 2147483647
                    _ = v35
                }
            default:
                if (128 > 127) {
                    _ = len(v4)
                } else {
                    var v36 uint16 = 2
                    if v21 {
                        v37 := GenS4{F0: 42}
                        if (1 != 18446744073709551615) {
                            v38 := GenS3{F0: -995, F1: 42, F2: 608}
                            fmt.Println(v21, v1, v2)
                            var v39 uint = 18446744073709551615
                            _ = v38
                            _ = v39
                        }
                        v36 -= (^(v36))
                        _ = v37
                    } else {
                        var v40 uint32 = 0
                        _ = v40
                    }
                    var v41 int64 = 0
                    _ = v36
                    _ = v41
                }
            }
            for fi42 := 0; fi42 < 11; fi42++ {
                for fi43 := 0; fi43 < 9; fi43++ {
                    if v21 {
                        var v44 int64 = -334
                        _ = v44
                    } else {
                        if true {
                            switch fi43 {
                            case 747:
                                var v45 float32 = 1.5
                                var v46 int32 = 773
                                var v47 int32 = int32(v2)
                                _ = v45
                                _ = v46
                                _ = v47
                            case -280:
                                v2 -= v2
                            default:
                                v48 := GenS2{F0: 1, F1: 1, F2: 1}
                                _ = v48
                            }
                            var v49 bool = (false || (v2 == (^(v2))))
                            _ = v49
                        } else {
                            var v50 uint16 = 1
                            if (!false) {
                                var v51 int16 = -1
                                var v52 uint = 787
                                _ = v51
                                _ = v52
                            }
                            _ = v50
                        }
                    }
                    _ = fi43
                }
                for fi53 := 0; fi53 < 13; fi53++ {
                    var v54 uint = 18446744073709551615
                    _ = fi53
                    _ = v54
                }
                _ = fi42
            }
        }
    }
    var v55 int = -947
    v21 = (-44 > v12)
    if (!v21) {
        v56 := v3[2147483647]
        v2 -= (^(v1))
        for fi57 := 0; fi57 < 6; fi57++ {
            var v58 uint16 = v2
            _ = len(v4)
            switch "\n\t" {
            case "":
                for fi59 := 0; fi59 < 11; fi59++ {
                    if (v58 == 65535) {
                        v60 := GenS2{F0: v2, F1: -1, F2: 0}
                        _ = v60
                    } else {
                        var v61 int64 = -1
                        var v62 int16 = -636
                        if v21 {
                            delete(v3, 4294967295)
                        } else {
                            v63 := GenS4{F0: v55}
                            _ = v63
                        }
                        _ = v61
                        _ = v62
                    }
                    fmt.Println(fi57)
                    v2 -= (v1 | 283)
                    _ = fi59
                }
            case "foo bar":
                fmt.Println(fi57, v58)
            }
            _ = fi57
            _ = v58
        }
        _ = v56
    }
    v2 -= 2
    v64 := GenS4{F0: (v55 & v55)}
    var v65 uint32 = 647
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v55:", v55)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v55
    _ = v64
    _ = v65
}
