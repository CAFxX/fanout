package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 int16; F2 uint16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int32; F1 int16 }
type GenS6 struct { F0 uint8; F1 int8 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint16; F1 bool; F2 bool }
type GenS9 struct { F0 uint8; F1 int16 }
type GenS10 struct { F0 int64; F1 int8 }

func main() {
    var v1 string = ("" + "")
    v2 := make(map[uint32]string)
    var v3 bool = true
    if (v3 && (65535 <= 42)) {
        fmt.Println(v3, v1)
        v3 = v3
    }
    for fi4 := 0; fi4 < 3; fi4++ {
        if true {
            v1 = v1
            v5 := GenS1{F0: 595, F1: 128}
            switch fi4 {
            case 0:
                for fi6 := 0; fi6 < 20; fi6++ {
                    v7 := GenS1{F0: 883, F1: 42}
                    _ = fi6
                    _ = v7
                }
            }
            _ = v5
        } else {
            if v3 {
                var v8 float32 = -2.25
                var v9 int16 = 290
                _ = v8
                _ = v9
            } else {
                fmt.Println(v3, v3, fi4)
                fmt.Println(v3, v1, fi4)
                for fi10 := 0; fi10 < 9; fi10++ {
                    v11 := []int{-2147483648, fi4, (fi4 / ((fi4) | 1))}
                    var v12 uint8 = 1
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            }
            v3 = (32767 != 1000)
        }
        var v13 int8 = 50
        _ = fi4
        _ = v13
    }
    var v14 int = 1
    v2[2147483647] = (v1 + "hello")
    if v3 {
        var v15 float32 = ((((((0.0 - float32((((1.5 + (-2.25 - 0.0)) - (1e10 / float64(1.5))) - 3.14159))) * 0.0) / (0.0 + 3.14159)) / -2.25) + -2.25) + 3.14159)
        if true {
            switch (v14 << ((0) & 63)) {
            case 32768:
                v16 := []int{v14, 0}
                _ = v16
            default:
                v17 := GenS2{F0: v1}
                _ = v17
            }
        } else {
            v15 -= v15
        }
        switch (-(v14)) {
        case 396:
            v1 = "a"
        case 1:
            if (1 > 141) {
                if (1 < -9223372036854775808) {
                    for fi18 := 0; fi18 < 1; fi18++ {
                        if true {
                            for fi19 := 0; fi19 < 2; fi19++ {
                                v20 := GenS2{F0: (v1 + v1)}
                                _ = fi19
                                _ = v20
                            }
                            var v21 int64 = 9223372036854775807
                            var v22 int8 = -128
                            _ = v21
                            _ = v22
                        } else {
                            if (0 != -996) {
                                fmt.Println(fi18)
                                fmt.Println(v3, v3)
                                v23 := []int{v14, 42, 4294967295, v14}
                                _ = v23
                            }
                            v14 = fi18
                        }
                        v15 -= float32(float64(v15))
                        v24 := GenS3{F0: 114, F1: -1, F2: 766}
                        _ = fi18
                        _ = v24
                    }
                    var v25 int8 = 1
                    if v3 {
                        v25 = v25
                        if v3 {
                            fmt.Println(v25)
                        } else {
                            v26 := GenS3{F0: 2, F1: -957, F2: 32767}
                            if v3 {
                                fmt.Println(v25, v15, v14)
                                v27 := GenS4{F0: (v14 << ((2) & 63))}
                                v25 = (v25 / ((v25) | 1))
                                _ = v27
                            }
                            if (!(!(0 > 4294967295))) {
                                var v28 int64 = -1
                                for fi29 := 0; fi29 < 14; fi29++ {
                                    v30 := GenS5{F0: 924, F1: 1000}
                                    _ = fi29
                                    _ = v30
                                }
                                if ((v14 < (v14 / (((-(v14))) | 1))) && v3) {
                                    v3 = v3
                                }
                                _ = v28
                            }
                            _ = v26
                        }
                        var v31 int = -950
                        _ = v31
                    }
                    _ = v25
                } else {
                    delete(v2, 934)
                    v32 := []int{-326}
                    switch (("" + (("hello" + v1) + ("" + v1))) + "\n\t") {
                    case "hello":
                        switch "foo bar" {
                        case "foo bar":
                            v33 := v2[2147483648]
                            _ = v33
                        case "a":
                            var v34 uint8 = 135
                            _ = v34
                        }
                    }
                    _ = v32
                }
            } else {
                if true {
                    v15 -= (v15 * float32(float64(0.0)))
                } else {
                    if v3 {
                        v15 = 1.5
                    }
                }
                var v35 string = "foo bar"
                if v3 {
                    v2[400] = v1
                    v36 := []int{v14, -9223372036854775808}
                    v36 = append(v36, (v14 % (((^(v14))) | 1)))
                    _ = v36
                }
                _ = v35
            }
        }
        _ = v15
    }
    v37 := []int{(v14 ^ (v14 - (v14 >> ((4294967295) & 63))))}
    var v38 int8 = 1
    switch "foo bar" {
    case "\n\t":
        var v39 int8 = v38
        _ = v39
    case "a":
        var v40 uint16 = 22
        _ = v40
    case "hello":
        if true {
            fmt.Println(v1)
            if v3 {
                if true {
                    var v41 uint8 = 0
                    _ = v41
                }
            }
            switch v14 {
            case 9:
                var v42 bool = (0 == 42)
                _ = v42
            case 243:
                if (0 < -1) {
                    var v43 int64 = 42
                    _ = v43
                } else {
                    fmt.Println(v3, v1, v3)
                    var v44 int64 = 42
                    if (42 <= 42) {
                        var v45 int16 = -144
                        _ = v45
                    } else {
                        v46 := GenS6{F0: 127, F1: 127}
                        switch ("\n\t" + ("a" + (v1 + "hello"))) {
                        case "foo bar":
                            switch (v14 & 42) {
                            case 4294967295:
                                v47 := GenS7{F0: (v38 & (v38 >> ((140) & 63)))}
                                _ = v47
                            case 42:
                                v48 := GenS5{F0: -265, F1: -1}
                                _ = v48
                            default:
                                if len(v37) >= 2 {
                                    v37 = v37[1:len(v37)]
                                }
                            }
                        case "hello":
                            v1 = (v1 + v1)
                        default:
                            var v49 float64 = (3.14159 / ((-2.25 / float64(1e10)) / 1.0))
                            _ = v49
                        }
                        fmt.Println(v3)
                        _ = v46
                    }
                    _ = v44
                }
            }
        }
    }
    var v50 uint64 = 42
    if v3 {
        if v3 {
            switch "" {
            case "hello":
                var v51 uint32 = 4294967295
                _ = v51
            case "foo bar":
                if false {
                    var v52 uint32 = 42
                    _ = v52
                }
            }
        }
        switch (v14 + (v14 ^ v14)) {
        case -233:
            var v53 bool = true
            _ = v53
        case -658:
            for fi54 := 0; fi54 < 6; fi54++ {
                var v55 float32 = (1.5 + 0.0)
                var v56 uint8 = 0
                _ = fi54
                _ = v55
                _ = v56
            }
        }
        if (((255 != 42) && (4294967295 <= (v50 % ((0) | 1)))) && v3) {
            v14 += int(v50)
            delete(v2, 972)
            if (!(v50 < 693)) {
                v3 = (0 != 42)
            } else {
                for fi57 := 0; fi57 < 13; fi57++ {
                    for fi58 := 0; fi58 < 16; fi58++ {
                        var v59 uint = 0
                        fmt.Println(v59, fi57, v59)
                        if (263 < v50) {
                            v3 = (v3 && v3)
                            var v60 uint = v59
                            if (-2147483648 != 42) {
                                var v61 int32 = -1
                                fmt.Println(v59, v1)
                                var v62 int64 = -297
                                v60 = v59
                                var v63 uint16 = 615
                                if len(v37) > 0 {
                                    v37[0] = (^(v14))
                                }
                                _ = v61
                                _ = v62
                                _ = v63
                            }
                            _ = v60
                        }
                        fmt.Println(v59, v59)
                        _ = fi58
                        _ = v59
                    }
                    if v3 {
                        fmt.Println(v50, v38, v3)
                    }
                    for fi64 := 0; fi64 < 13; fi64++ {
                        v65 := GenS6{F0: 2, F1: v38}
                        delete(v2, 949)
                        v50 = (v50 - (v50 ^ v50))
                        _ = fi64
                        _ = v65
                    }
                    _ = fi57
                }
                v37 = append(v37, (v14 >> ((18446744073709551615) & 63)))
            }
        } else {
            var v66 uint16 = 576
            v67 := GenS5{F0: -980, F1: 361}
            _ = v66
            _ = v67
        }
    }
    for fi68 := 0; fi68 < 6; fi68++ {
        v38 = 42
        if true {
            v69 := GenS8{F0: 447, F1: v3, F2: true}
            _ = v69
        } else {
            _ = len(v37)
            v2[uint32(v50)] = "\n\t"
        }
        var v70 string = v1
        _ = fi68
        _ = v70
    }
    if v3 {
        v71 := GenS9{F0: 50, F1: -32768}
        fmt.Println(v1)
        _ = v71
    } else {
        v50 -= v50
        var v72 string = ("\n\t" + "foo bar")
        _ = v72
    }
    if (0 <= 68) {
        var v73 string = ("foo bar" + v1)
        _ = v73
    }
    switch (v14 >> ((39) & 63)) {
    case -2147483648:
        v74 := GenS10{F0: -884, F1: 0}
        _ = v74
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v38:", v38)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v50
}
