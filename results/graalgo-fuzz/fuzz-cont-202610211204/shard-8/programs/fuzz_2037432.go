package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 int32 }
type GenS2 struct { F0 string; F1 bool; F2 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int; F1 int8 }
type GenS6 struct { F0 uint64 }

func main() {
    var v1 uint8 = 42
    var v2 float32 = (0.0 * -2.25)
    v3 := make(map[int]bool)
    v4 := []uint32{1}
    switch "a" {
    case "foo bar":
        var v5 uint64 = 574
        _ = v5
    case "a":
        var v6 uint64 = 1
        _ = v6
    default:
        v7 := GenS1{F0: 42, F1: 753, F2: 189}
        _ = v7
    }
    if true {
        var v8 uint = 65535
        v9 := GenS1{F0: -624, F1: uint32(v1), F2: 1}
        v1 += (v1 >> (((^(v8))) & 63))
        _ = v8
        _ = v9
    }
    if (true && false) {
        v2 -= float32((1.5 * 1.5))
        v4 = append(v4, 707)
        for fi10 := 0; fi10 < 11; fi10++ {
            v11 := GenS2{F0: (("hello" + ("\n\t" + "foo bar")) + "a"), F1: ((3.14159 + ((1e10 - (0.0 * float64(1.5))) + 0.1)) == 0.0), F2: 4294967295}
            var v12 int64 = 9223372036854775807
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    v13 := GenS2{F0: ((("hello" + "") + "hello") + "foo bar"), F1: true, F2: 530}
    if (-580 != 32767) {
        var v14 bool = true
        _ = v14
    }
    v15 := GenS1{F0: 2147483647, F1: 237, F2: -1}
    v2 += -2.25
    v2 = 0.0
    if false {
        v16 := GenS2{F0: "foo bar", F1: true, F2: 18446744073709551615}
        v1 = 128
        _ = v16
    } else {
        for fi17 := 0; fi17 < 5; fi17++ {
            for fi18 := 0; fi18 < 16; fi18++ {
                if ((!(255 != 42)) && ("\n\t" >= "\n\t")) {
                    fmt.Println(v1)
                    var v19 uint16 = 32767
                    if (-1 < 1) {
                        v2 -= (1e10 + (((v2 / float32(float64((-2.25 - v2)))) / v2) + v2))
                        if (433 < 0) {
                            if true {
                                v4 = append(v4, 955)
                                var v20 int32 = -90
                                v21 := GenS1{F0: v20, F1: 2, F2: 0}
                                _ = v20
                                _ = v21
                            }
                            switch 328 {
                            case 42:
                                v22 := GenS2{F0: "a", F1: false, F2: 42}
                                _ = v22
                            case 9223372036854775807:
                                if true {
                                    if true {
                                        v2 -= 3.14159
                                        v23 := GenS1{F0: 2147483647, F1: 42, F2: 42}
                                        _ = v23
                                    } else {
                                        v24 := GenS3{F0: (v1 % ((v1) | 1))}
                                        _ = v24
                                    }
                                    var v25 int64 = 9223372036854775807
                                    var v26 uint16 = 32767
                                    _ = v25
                                    _ = v26
                                }
                            case -9223372036854775808:
                                fmt.Println(v2, fi18)
                            default:
                                v27 := GenS1{F0: 2147483647, F1: 762, F2: 1}
                                _ = v27
                            }
                            if (((0.0 < (v2 * float32(-2.25))) && (v1 <= v1)) || true) {
                                var v28 uint64 = 42
                                _ = len(v4)
                                _ = v28
                            }
                        } else {
                            v2 -= (3.14159 * (v2 / (float32(3.14159) * v2)))
                            if false {
                                var v29 int8 = -52
                                _ = v29
                            } else {
                                v3[fi17] = ((v1 % ((v1) | 1)) <= v1)
                            }
                            v30 := GenS1{F0: 1, F1: 4294967295, F2: -1}
                            _ = v30
                        }
                    }
                    fmt.Println(fi17)
                    var v31 float64 = float64((v2 / v2))
                    switch 128 {
                    case 128:
                        var v32 float64 = float64((v2 * (1e10 / ((float32(float64(v2)) * (v2 / (1e10 - v2))) / v2))))
                        _ = v32
                    case 32767:
                        if len(v4) > 0 {
                            v4[0] = 0
                        }
                    default:
                        var v33 int8 = 0
                        _ = v33
                    }
                    var v34 int = (fi17 - fi17)
                    v35 := GenS2{F0: "\n\t", F1: (0 <= 412), F2: 42}
                    _ = v19
                    _ = v31
                    _ = v34
                    _ = v35
                } else {
                    v1 -= (v1 % ((v1) | 1))
                    var v36 uint = 4294967295
                    v37 := GenS4{F0: 9223372036854775807}
                    _ = v36
                    _ = v37
                }
                fmt.Println(v2, v1)
                fmt.Println(fi17, v2)
                _ = fi18
            }
            _ = fi17
        }
    }
    if true {
        v38 := v3[127]
        _ = v38
    } else {
        var v39 uint8 = (v1 % (((v1 - v1)) | 1))
        _ = v39
    }
    switch -385 {
    case 310:
        v2 += 0.1
    case -1:
        fmt.Println(v1, v1)
    case 32768:
        var v40 int16 = -27
        _ = v40
    }
    if (65535 == 858) {
        v4 = append(v4, 825)
        v41 := GenS1{F0: 42, F1: 2147483647, F2: -1}
        _ = v41
    } else {
        var v42 int16 = 1
        for fi43 := 0; fi43 < 8; fi43++ {
            var v44 int16 = (^(v42))
            _ = fi43
            _ = v44
        }
        _ = v42
    }
    var v45 bool = (42 > 106)
    if true {
        for fi46 := 0; fi46 < 5; fi46++ {
            switch 32768 {
            case 128:
                v47 := GenS4{F0: 42}
                _ = v47
            case -1:
                fmt.Println(v45, v1)
            case -9223372036854775808:
                switch "\n\t" {
                case "\n\t":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
            }
            _ = fi46
        }
        for fi48 := 0; fi48 < 8; fi48++ {
            fmt.Println(v1)
            var v49 float64 = 0.1
            var v50 int8 = -1
            _ = fi48
            _ = v49
            _ = v50
        }
        var v51 uint16 = 1
        _ = v51
    } else {
        switch "foo bar" {
        case "\n\t":
            switch ((("a" + "hello") + ("\n\t" + "a")) + "a") {
            case "":
                var v52 uint16 = 692
                _ = v52
            case "a":
                var v53 float32 = (float32(((1e10 + 1e10) / float64(v2))) - float32(-2.25))
                _ = v53
            case "foo bar":
                var v54 uint = 255
                _ = v54
            default:
                v55 := GenS4{F0: 9223372036854775807}
                _ = v55
            }
        default:
            var v56 int8 = -112
            _ = v56
        }
        v57 := GenS5{F0: 379, F1: 1}
        v58 := GenS1{F0: 2147483647, F1: 2147483647, F2: -1}
        _ = v57
        _ = v58
    }
    v59 := GenS6{F0: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v59
}
