package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 string; F1 uint8; F2 int64 }
type GenS3 struct { F0 int16; F1 int32; F2 uint64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint16; F1 int64 }
func genFn0(p0 uint64, p1 uint, p2 uint) uint {
    v7 := GenS1{F0: 1, F1: "foo bar"}
    if false {
        var v8 float32 = 1.5
        v9 := []int{32768, 0, 1, 9223372036854775807}
        _ = v8
        _ = v9
    }
    p2 += (^(p1))
    fmt.Println(p2)
    v10 := GenS1{F0: 327, F1: ""}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    return (p2 * 0)
}
func genFn1() int16 {
    var v11 string = (("foo bar" + "a") + "foo bar")
    v11 = v11
    var v12 int16 = 32767
    var v13 int = -485
    v11 = (((v11 + v11) + v11) + v11)
    _ = v11
    _ = v12
    _ = v13
    return (v12 & (v12 | (v12 >> ((1) & 63))))
}

func main() {
    v1 := make(map[int]string)
    var v2 uint64 = 1
    v3 := make(map[int]uint32)
    var v4 bool = (0 < 65535)
    var v5 int32 = 22
    var v6 float32 = float32(1.5)
    if (v4 && (genFn0(42, 999, 290) > 4294967295)) {
        v4 = (65535 <= 0)
        v14 := GenS2{F0: "", F1: 70, F2: 42}
        _ = v14
    } else {
        switch -2147483648 {
        case 42:
            for fi15 := 0; fi15 < 19; fi15++ {
                v16 := GenS1{F0: -1, F1: ("\n\t" + ("" + ("a" + ((("" + "hello") + ("hello" + "a")) + "foo bar"))))}
                var v17 int16 = 322
                v5 = 42
                _ = fi15
                _ = v16
                _ = v17
            }
        case 127:
            _ = len(v3)
        case 0:
            switch "\n\t" {
            case "":
                v2 -= v2
            default:
                var v18 uint8 = 0
                _ = v18
            }
        }
        var v19 uint8 = 255
        _ = v19
    }
    v5 += (v5 % ((0) | 1))
    _ = len(v1)
    v2 -= (v2 & 1)
    var v20 float64 = 1e10
    var v21 uint16 = 256
    switch "" {
    case "\n\t":
        for fi22 := 0; fi22 < 17; fi22++ {
            if true {
                for fi23 := 0; fi23 < 12; fi23++ {
                    if ((127 > 2) && v4) {
                        switch (^(fi23)) {
                        case -208:
                            v6 += v6
                        }
                        if (false && (!(v20 != 3.14159))) {
                            var v24 uint64 = (v2 | (v2 - (^(v2))))
                            _ = v24
                        }
                    } else {
                        v2 -= 142
                    }
                    v20 += 0.1
                    _ = fi23
                }
                var v25 uint64 = 547
                v25 += (v2 & (v2 * 595))
                _ = v25
            }
            _ = fi22
        }
    default:
        var v26 int = -152
        _ = v26
    }
    v27 := GenS3{F0: -294, F1: v5, F2: v2}
    if true {
        v4 = (v6 != 1e10)
    }
    if (0.1 >= float64((v6 - 0.1))) {
        v1[2147483647] = "\n\t"
        var v28 uint = 1
        _ = v28
    }
    for fi29 := 0; fi29 < 10; fi29++ {
        switch fi29 {
        case 127:
            switch 4294967295 {
            case 128:
                for fi30 := 0; fi30 < 13; fi30++ {
                    var v31 uint = 42
                    _ = fi30
                    _ = v31
                }
            default:
                v5 -= 0
            }
        case 0:
            if (!v4) {
                if true {
                    for fi32 := 0; fi32 < 12; fi32++ {
                        var v33 uint64 = (v2 ^ v2)
                        _ = fi32
                        _ = v33
                    }
                    var v34 uint16 = (^(v21))
                    _ = v34
                }
                v20 += (v20 - -2.25)
                if (v4 || v4) {
                    v35 := GenS2{F0: ("foo bar" + "foo bar"), F1: 42, F2: 543}
                    _ = v35
                } else {
                    var v36 int8 = -1
                    v36 -= v36
                    _ = v36
                }
            } else {
                v4 = v4
                v5 = (v5 << ((65535) & 63))
            }
        case 567:
            if (!(v20 <= 1.5)) {
                fmt.Println(v21, v21, v2)
                switch (fi29 - (^(fi29))) {
                case 32768:
                    if v4 {
                        v6 = float32(v20)
                    }
                }
                v5 = v5
            } else {
                switch fi29 {
                case -9223372036854775808:
                    for fi37 := 0; fi37 < 9; fi37++ {
                        if (((v6 / v6) * float32(float64(((v6 + v6) + float32(0.1))))) == 3.14159) {
                            v38 := GenS1{F0: 32767, F1: "hello"}
                            if v4 {
                                v2 = v2
                                _ = len(v1)
                                var v39 int8 = 42
                                _ = v39
                            }
                            switch ("foo bar" + "") {
                            case "":
                                if ((v5 / (((v5 + 0)) | 1)) == 42) {
                                    _ = len(v3)
                                    v6 = v6
                                    v5 = v5
                                }
                            default:
                                var v40 string = ""
                                _ = v40
                            }
                            fmt.Println(v5, v6, v4)
                            _ = v38
                        }
                        v6 = (v6 / v6)
                        fmt.Println(v2)
                        _ = fi37
                    }
                case -1:
                    switch fi29 {
                    case 42:
                        var v41 string = "a"
                        _ = v41
                    case 0:
                        for fi42 := 0; fi42 < 9; fi42++ {
                            v43 := []int{fi42}
                            v44 := GenS2{F0: "a", F1: 255, F2: 42}
                            _ = fi42
                            _ = v43
                            _ = v44
                        }
                    case -2147483648:
                        var v45 int64 = 42
                        _ = v45
                    }
                case 32767:
                    v46 := []int{}
                    _ = v46
                }
                v47 := v1[(-(fi29))]
                _ = v47
            }
        }
        switch ("foo bar" + "hello") {
        case "":
            v48 := GenS3{F0: int16(v21), F1: 522, F2: v2}
            _ = v48
        case "foo bar":
            for fi49 := 0; fi49 < 6; fi49++ {
                fmt.Println(fi49)
                var v50 int32 = v5
                var v51 float32 = v6
                _ = fi49
                _ = v50
                _ = v51
            }
        case "hello":
            var v52 int16 = -96
            _ = v52
        default:
            var v53 float64 = v20
            _ = v53
        }
        _ = fi29
    }
    v54 := GenS4{F0: 1}
    for fi55 := 0; fi55 < 6; fi55++ {
        _ = len(v3)
        v56 := []int{}
        v6 += float32(3.14159)
        _ = fi55
        _ = v56
    }
    for fi57 := 0; fi57 < 9; fi57++ {
        for fi58 := 0; fi58 < 20; fi58++ {
            fmt.Println(v5, v21, v20)
            var v59 int64 = -305
            var v60 float64 = 1.5
            v61 := GenS5{F0: 65535, F1: v59}
            v59 -= (v59 >> ((2) & 63))
            _ = fi58
            _ = v59
            _ = v60
            _ = v61
        }
        switch "a" {
        case "a":
            fmt.Println(v6, v5)
        case "foo bar":
            v62 := []int{}
            _ = v62
        default:
            var v63 int = fi57
            _ = v63
        }
        var v64 uint = 307
        _ = fi57
        _ = v64
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v6
}
