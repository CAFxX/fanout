package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint; F1 int32; F2 int16 }
func genFn0() uint16 {
    var v7 float32 = 1e10
    v7 += 3.14159
    fmt.Println(v7)
    v7 = 1e10
    _ = v7
    return 255
}
func genFn1() string {
    for fi8 := 0; fi8 < 9; fi8++ {
        for fi9 := 0; fi9 < 8; fi9++ {
            switch (fi8 * fi8) {
            case 4294967295:
                v10 := GenS1{F0: 2, F1: 171}
                _ = v10
            }
            v11 := []int{(^(fi9))}
            var v12 float64 = 3.14159
            _ = fi9
            _ = v11
            _ = v12
        }
        var v13 int8 = 1
        switch "hello" {
        case "foo bar":
            for fi14 := 0; fi14 < 10; fi14++ {
                fmt.Println(v13)
                v15 := []int{(fi14 + 610), fi14, 14}
                if true {
                    v13 = (v13 << ((611) & 63))
                    v13 += (v13 + 127)
                } else {
                    if (!(false || (1 <= 1))) {
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                        v13 = (v13 | v13)
                    }
                    v13 = v13
                    if ((3.14159 + 0.0) > float32(3.14159)) {
                        fmt.Println(fi8)
                    }
                }
                _ = fi14
                _ = v15
            }
        }
        _ = fi8
        _ = v13
    }
    for fi16 := 0; fi16 < 17; fi16++ {
        for fi17 := 0; fi17 < 19; fi17++ {
            if (true || false) {
                var v18 int64 = 790
                var v19 int32 = -1
                v19 = (^(v19))
                var v20 float64 = -2.25
                _ = v18
                _ = v19
                _ = v20
            }
            switch 128 {
            case 2147483647:
                v21 := GenS2{F0: 42}
                _ = v21
            }
            _ = fi17
        }
        v22 := GenS3{F0: 51, F1: 128}
        _ = fi16
        _ = v22
    }
    return ("" + (("" + ((("a" + "\n\t") + "") + "hello")) + "hello"))
}

func main() {
    var v1 uint64 = 2
    v2 := make(map[uint32]int)
    v3 := []uint32{uint32(v1), 83, 17, 0}
    var v4 float64 = (((-2.25 / 0.1) + -2.25) + (0.1 / 1.0))
    v5 := make(map[uint32]string)
    var v6 string = ("hello" + ("\n\t" + "a"))
    v3 = append(v3, 2147483648)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if true {
        for fi23 := 0; fi23 < 2; fi23++ {
            var v24 uint = 65535
            _ = len(v3)
            _ = fi23
            _ = v24
        }
        v1 = 18446744073709551615
        if (2147483647 != 289) {
            for fi25 := 0; fi25 < 9; fi25++ {
                v3 = append(v3, 2147483647)
                for fi26 := 0; fi26 < 12; fi26++ {
                    switch v6 {
                    case "foo bar":
                        if len(v3) > 0 {
                            v3[0] = 42
                        }
                    }
                    fmt.Println(v4, v4)
                    var v27 uint32 = 861
                    _ = fi26
                    _ = v27
                }
                v1 = 0
                _ = fi25
            }
        } else {
            v28 := GenS4{F0: "hello"}
            for fi29 := 0; fi29 < 5; fi29++ {
                var v30 int64 = -9223372036854775808
                v4 += v4
                var v31 bool = (fi29 == fi29)
                _ = fi29
                _ = v30
                _ = v31
            }
            v32 := GenS4{F0: (("a" + v6) + "\n\t")}
            _ = v28
            _ = v32
        }
    } else {
        for fi33 := 0; fi33 < 9; fi33++ {
            for fi34 := 0; fi34 < 10; fi34++ {
                v35 := GenS4{F0: "foo bar"}
                if true {
                    switch fi33 {
                    case 32767:
                        v36 := v5[95]
                        _ = v36
                    case 0:
                        v37 := GenS4{F0: v6}
                        _ = v37
                    case 128:
                        delete(v2, uint32(v1))
                    }
                    var v38 uint = 255
                    _ = v38
                } else {
                    if true {
                        v3 = append(v3, 2147483648)
                        v39 := GenS1{F0: 2, F1: 4294967295}
                        _ = v39
                    }
                    if true {
                        var v40 uint16 = genFn0()
                        _ = v40
                    } else {
                        _ = len(v3)
                        var v41 float32 = (1e10 / (((float32(v4) / 1.0) * (((float32(v4) + -2.25) * ((-2.25 + 0.1) / 0.1)) + ((-2.25 - 1.5) / ((0.1 / 1e10) - 1.5)))) - float32(v4)))
                        _ = v41
                    }
                }
                var v42 bool = true
                _ = fi34
                _ = v35
                _ = v42
            }
            _ = fi33
        }
    }
    fmt.Println(v6, v6)
    if (!true) {
        for fi43 := 0; fi43 < 8; fi43++ {
            if true {
                if (1 != 652) {
                    v1 += v1
                    if (-5 < 0) {
                        fmt.Println(v6, v1, v1)
                        var v44 float32 = ((-2.25 * (float32((float64((float32(0.0) + 3.14159)) - ((v4 * (v4 - 0.0)) * v4))) - ((0.1 - 1.5) + 3.14159))) / float32(v4))
                        _ = v44
                    } else {
                        for fi45 := 0; fi45 < 5; fi45++ {
                            switch 127 {
                            case 398:
                                var v46 uint8 = 2
                                _ = v46
                            case -1:
                                v6 = v6
                            case -9223372036854775808:
                                switch v6 {
                                case "a":
                                    v47 := GenS1{F0: genFn0(), F1: 2147483648}
                                    _ = v47
                                case "foo bar":
                                    v48 := GenS5{F0: 880}
                                    _ = v48
                                }
                            }
                            v49 := GenS5{F0: 550}
                            _ = fi45
                            _ = v49
                        }
                        switch 326 {
                        case 4294967295:
                            fmt.Println(fi43, v6, v6)
                        case -184:
                            var v50 uint = 891
                            _ = v50
                        default:
                            v51 := GenS6{F0: 1, F1: 0, F2: 600}
                            _ = v51
                        }
                    }
                }
                fmt.Println(v4)
                fmt.Println(v1)
            }
            switch ("hello" + v6) {
            case "foo bar":
                var v52 uint16 = 317
                _ = v52
            default:
                var v53 float32 = (float32(v4) / -2.25)
                _ = v53
            }
            fmt.Println(v4, v6, v4)
            _ = fi43
        }
        var v54 int8 = 1
        _ = v54
    } else {
        var v55 float32 = 1.5
        for fi56 := 0; fi56 < 10; fi56++ {
            fmt.Println(v4, v55, v1)
            var v57 uint16 = 1
            _ = fi56
            _ = v57
        }
        v55 -= (v55 - (float32(v4) - ((float32((3.14159 - float64(v55))) / ((0.1 - (0.1 + 1.5)) * v55)) / 1e10)))
        _ = v55
    }
    var v58 uint = 759
    v58 -= uint(v1)
    v1 = v1
    fmt.Println(v1, v4)
    for fi59 := 0; fi59 < 19; fi59++ {
        v58 = 255
        v4 = float64(0.0)
        var v60 int = -896
        _ = fi59
        _ = v60
    }
    fmt.Println(v58, v6)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v6
}
