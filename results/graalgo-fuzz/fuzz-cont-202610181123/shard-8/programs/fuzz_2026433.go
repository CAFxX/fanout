package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 int32; F1 uint8; F2 int64 }
type GenS5 struct { F0 uint8; F1 string }
func genFn0() uint16 {
    var v4 int16 = 697
    if (!false) {
        v5 := []int{2147483647}
        v4 = v4
        var v6 string = (("hello" + "foo bar") + "a")
        _ = v5
        _ = v6
    }
    var v7 uint16 = 42
    v8 := GenS1{F0: true}
    for fi9 := 0; fi9 < 19; fi9++ {
        v7 += v7
        _ = fi9
    }
    _ = v4
    _ = v7
    _ = v8
    return (^(v7))
}
func genFn1(p0 string) uint64 {
    var v10 uint64 = 18446744073709551615
    var v11 int16 = int16(v10)
    for fi12 := 0; fi12 < 7; fi12++ {
        for fi13 := 0; fi13 < 5; fi13++ {
            p0 = p0
            _ = fi13
        }
        switch fi12 {
        case 128:
            if (-2147483648 == 572) {
                if true {
                    if ((v10 % (((v10 * v10)) | 1)) < v10) {
                        p0 = p0
                        v14 := GenS2{F0: 0, F1: 4}
                        _ = v14
                    } else {
                        if false {
                            switch (fi12 << ((1) & 63)) {
                            case 128:
                                switch fi12 {
                                case -840:
                                    var v15 uint64 = v10
                                    _ = v15
                                case -755:
                                    if false {
                                        var v16 uint8 = 1
                                        for fi17 := 0; fi17 < 3; fi17++ {
                                            v18 := GenS1{F0: false}
                                            var v19 uint8 = (v16 * (^(v16)))
                                            _ = fi17
                                            _ = v18
                                            _ = v19
                                        }
                                        _ = v16
                                    } else {
                                        var v20 uint16 = 65535
                                        _ = v20
                                    }
                                case 32768:
                                    for fi21 := 0; fi21 < 15; fi21++ {
                                        var v22 int64 = 42
                                        _ = fi21
                                        _ = v22
                                    }
                                default:
                                    v23 := []int{fi12, 0}
                                    _ = v23
                                }
                            case 4294967295:
                                p0 = ("hello" + "")
                            }
                            v24 := []int{(fi12 * (fi12 + (fi12 + fi12))), fi12}
                            _ = v24
                        }
                    }
                    p0 = p0
                    for fi25 := 0; fi25 < 16; fi25++ {
                        v26 := []int{(fi12 << ((1) & 63)), fi25, fi25, fi25}
                        var v27 bool = (1e10 >= (float32(float64(1.5)) / ((0.1 / 1.5) + (0.1 + 1e10))))
                        fmt.Println(v11)
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                } else {
                    fmt.Println(v11, fi12, v10)
                }
            } else {
                v28 := GenS3{F0: 28, F1: (v11 - v11)}
                var v29 float32 = -2.25
                var v30 uint32 = 2147483648
                _ = v28
                _ = v29
                _ = v30
            }
        case 1:
            v31 := []int{fi12, fi12, 128, (^(fi12))}
            _ = v31
        case 165:
            var v32 int64 = 42
            _ = v32
        default:
            p0 = ((p0 + "") + ("a" + ((("hello" + "foo bar") + ((p0 + ("" + "a")) + ("" + "\n\t"))) + p0)))
        }
        _ = fi12
    }
    _ = p0
    _ = v10
    _ = v11
    return 110
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]string)
    var v3 int8 = -83
    v3 += v3
    v3 -= -102
    fmt.Println(v3)
    v33 := v2["\n\t"]
    if (255 != 674) {
        var v34 uint8 = 180
        for fi35 := 0; fi35 < 3; fi35++ {
            v36 := GenS4{F0: 700, F1: (v34 * 255), F2: 969}
            _ = fi35
            _ = v36
        }
        for fi37 := 0; fi37 < 1; fi37++ {
            v3 += v3
            if false {
                v3 = (v3 << ((125) & 63))
            }
            switch v33 {
            case "foo bar":
                v34 -= (v34 % (((v34 >> ((18446744073709551615) & 63))) | 1))
            case "a":
                for fi38 := 0; fi38 < 7; fi38++ {
                    v39 := GenS2{F0: (v3 ^ 100), F1: v3}
                    _ = fi38
                    _ = v39
                }
            case "":
                v40 := GenS5{F0: (v34 ^ (v34 / (((v34 >> ((0) & 63))) | 1))), F1: "foo bar"}
                _ = v40
            default:
                switch fi37 {
                case -2147483648:
                    for fi41 := 0; fi41 < 10; fi41++ {
                        v42 := GenS4{F0: 727, F1: 1, F2: -1}
                        if true {
                            v3 = (v3 - 42)
                        } else {
                            v34 += v34
                            v43 := GenS2{F0: (v3 >> ((0) & 63)), F1: 38}
                            _ = v43
                        }
                        v33 = v33
                        _ = fi41
                        _ = v42
                    }
                case 32768:
                    v44 := GenS1{F0: (186 < 9223372036854775807)}
                    _ = v44
                case 42:
                    v34 = v34
                }
            }
            _ = fi37
        }
        _ = v34
    } else {
        var v45 uint8 = 71
        _ = v45
    }
    for fi46 := 0; fi46 < 1; fi46++ {
        v33 = v33
        for fi47 := 0; fi47 < 7; fi47++ {
            fmt.Println(fi46, fi47, v33)
            _ = fi47
        }
        v33 = v33
        _ = fi46
    }
    if (v3 > (v3 % (((v3 | (v3 * v3))) | 1))) {
        if true {
            v48 := GenS2{F0: (-(v3)), F1: (v3 << ((4294967295) & 63))}
            _ = v48
        } else {
            for fi49 := 0; fi49 < 1; fi49++ {
                switch (v33 + "\n\t") {
                case "":
                    for fi50 := 0; fi50 < 8; fi50++ {
                        if true {
                            v51 := GenS5{F0: 1, F1: v33}
                            v52 := []int{(fi50 ^ 32768)}
                            _ = v51
                            _ = v52
                        } else {
                            switch 127 {
                            case -440:
                                fmt.Println(v33, v33, v3)
                            case -2147483648:
                                v3 += v3
                            case -420:
                                var v53 uint64 = 18446744073709551615
                                _ = v53
                            }
                        }
                        _ = fi50
                    }
                case "a":
                    var v54 string = v33
                    _ = v54
                }
                var v55 uint = 4294967295
                _ = fi49
                _ = v55
            }
            v56 := []int{-2147483648, -1, -2147483648, 42}
            for fi57 := 0; fi57 < 11; fi57++ {
                switch (fi57 & fi57) {
                case 220:
                    var v58 bool = ((true || false) && (v33 < v33))
                    _ = v58
                default:
                    fmt.Println(v33, v3)
                }
                _ = fi57
            }
            _ = v56
        }
        v59 := []int{}
        fmt.Println(v33)
        _ = v59
    } else {
        v60 := GenS2{F0: v3, F1: v3}
        v3 = v3
        var v61 int64 = 9223372036854775807
        _ = v60
        _ = v61
    }
    for fi62 := 0; fi62 < 8; fi62++ {
        v63 := GenS3{F0: v3, F1: -728}
        _ = fi62
        _ = v63
    }
    for fi64 := 0; fi64 < 12; fi64++ {
        var v65 int64 = -686
        _ = fi64
        _ = v65
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
}
