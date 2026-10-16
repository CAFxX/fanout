package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int; F1 int; F2 int16 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint64 = 640
    var v3 string = (((("" + ("" + "a")) + ("\n\t" + "")) + ("foo bar" + (("a" + "hello") + "\n\t"))) + "foo bar")
    var v4 uint32 = 1
    v5 := make(map[string]uint32)
    switch 796 {
    case -9:
        v2 += (v2 >> ((2) & 63))
    case 2147483647:
        v6 := GenS1{F0: 1}
        _ = v6
    case -1:
        var v7 int64 = 131
        _ = v7
    }
    v8 := GenS2{F0: 102, F1: 65535}
    var v9 int32 = 42
    v10 := GenS2{F0: 100, F1: 7}
    var v11 bool = (false && true)
    switch "a" {
    case "a":
        var v12 int8 = 127
        _ = v12
    case "\n\t":
        v2 += v2
    }
    v13 := []int{-205}
    var v14 string = (v3 + v3)
    if v11 {
        if true {
            if v11 {
                for fi15 := 0; fi15 < 1; fi15++ {
                    v16 := GenS3{F0: v9}
                    _ = fi15
                    _ = v16
                }
            } else {
                v9 -= v9
            }
        } else {
            if (-1 == -86) {
                switch -1 {
                case -740:
                    if v11 {
                        v3 = v14
                    }
                case 128:
                    if v11 {
                        for fi17 := 0; fi17 < 2; fi17++ {
                            for fi18 := 0; fi18 < 3; fi18++ {
                                fmt.Println(fi17, v3, v3)
                                var v19 int16 = 0
                                _ = fi18
                                _ = v19
                            }
                            switch fi17 {
                            case 504:
                                v20 := GenS1{F0: fi17}
                                _ = v20
                            case 0:
                                var v21 float64 = float64((3.14159 + -2.25))
                                _ = v21
                            case 1:
                                if len(v13) > 0 {
                                    v13[0] = (fi17 - fi17)
                                }
                            default:
                                v11 = true
                            }
                            _ = fi17
                        }
                    } else {
                        v9 += v9
                    }
                default:
                    v2 += (v2 % ((v2) | 1))
                }
                for fi22 := 0; fi22 < 15; fi22++ {
                    if (fi22 == (fi22 + (fi22 / ((-102) | 1)))) {
                        fmt.Println(v4, v4)
                    } else {
                        for fi23 := 0; fi23 < 18; fi23++ {
                            var v24 int8 = 42
                            var v25 int64 = 42
                            _ = fi23
                            _ = v24
                            _ = v25
                        }
                        for fi26 := 0; fi26 < 3; fi26++ {
                            v27 := GenS4{F0: 255}
                            _ = fi26
                            _ = v27
                        }
                        for fi28 := 0; fi28 < 2; fi28++ {
                            if v11 {
                                delete(v5, v3)
                                switch fi22 {
                                case 32768:
                                    if (431 != 964) {
                                        if (v11 && v11) {
                                            v3 = ("hello" + v3)
                                            var v29 float32 = float32(float64(0.0))
                                            var v30 uint32 = (^(v4))
                                            _ = v29
                                            _ = v30
                                        } else {
                                            _ = len(v1)
                                        }
                                    }
                                case -119:
                                    v4 = 1
                                }
                            }
                            _ = fi28
                        }
                    }
                    var v31 float32 = (((0.1 / 1.0) * 1e10) / (1.5 * -2.25))
                    _ = fi22
                    _ = v31
                }
            }
            v32 := GenS3{F0: v9}
            var v33 uint16 = 714
            _ = v32
            _ = v33
        }
        if true {
            var v34 float32 = -2.25
            var v35 uint16 = 2
            if (0 > 0) {
                if v11 {
                    if ((v14 + "") > ("a" + ((v3 + ("" + ("\n\t" + "a"))) + "\n\t"))) {
                        var v36 int8 = 100
                        v14 = ("foo bar" + "a")
                        var v37 uint16 = v35
                        _ = v36
                        _ = v37
                    }
                } else {
                    delete(v1, (v4 % (((v4 + 2147483648)) | 1)))
                    if v11 {
                        for fi38 := 0; fi38 < 14; fi38++ {
                            _ = len(v13)
                            _ = fi38
                        }
                        v39 := GenS2{F0: 32, F1: v35}
                        var v40 uint8 = 0
                        _ = v39
                        _ = v40
                    }
                    var v41 int = 591
                    _ = v41
                }
                v14 = v14
                var v42 uint16 = 42
                _ = v42
            } else {
                for fi43 := 0; fi43 < 14; fi43++ {
                    v11 = false
                    _ = fi43
                }
                var v44 float64 = (1.5 - 0.0)
                var v45 uint32 = 2
                _ = v44
                _ = v45
            }
            _ = v34
            _ = v35
        } else {
            v46 := GenS4{F0: 2}
            var v47 int64 = 508
            var v48 uint16 = 42
            _ = v46
            _ = v47
            _ = v48
        }
        fmt.Println(v2, v2)
    }
    fmt.Println(v2, v4)
    switch 127 {
    case 678:
        switch ("a" + "foo bar") {
        case "a":
            v49 := GenS1{F0: 1}
            _ = v49
        case "\n\t":
            var v50 float32 = float32(1.5)
            _ = v50
        default:
            var v51 int8 = 42
            _ = v51
        }
    case 42:
        if (65535 >= (v4 * 2147483647)) {
            if (537 < 980) {
                v13 = append(v13, 0)
            } else {
                switch ("foo bar" + (("\n\t" + ((v3 + v3) + (v3 + v3))) + v3)) {
                case "hello":
                    if (0 == -95) {
                        fmt.Println(v9, v4)
                    }
                }
            }
            var v52 int64 = 42
            v53 := GenS5{F0: -2147483648, F1: -1, F2: 42}
            _ = v52
            _ = v53
        }
    default:
        switch -2147483648 {
        case -2147483648:
            if ((-1 != 676) || (0 >= 42)) {
                var v54 int = 1
                var v55 float64 = 1e10
                _ = v54
                _ = v55
            } else {
                var v56 int16 = 252
                fmt.Println(v4)
                switch 4294967295 {
                case -9223372036854775808:
                    var v57 int16 = v56
                    _ = v57
                case 85:
                    _ = len(v13)
                case 127:
                    fmt.Println(v56, v14, v4)
                default:
                    v2 = v2
                }
                _ = v56
            }
        case 4294967295:
            var v58 string = ("\n\t" + "hello")
            _ = v58
        }
    }
    v59 := v5["hello"]
    v13 = append(v13, 686)
    var v60 uint8 = 255
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v60
    _ = v8
    _ = v9
}
