package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 uint64; F1 uint; F2 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint; F1 uint32; F2 int64 }
type GenS5 struct { F0 int64; F1 uint16 }

func main() {
    var v1 int16 = 1000
    v2 := []string{"hello", "foo bar"}
    v3 := []uint32{2147483647, 2, 130, 4294967295}
    v4 := make(map[uint32]uint32)
    if ("a" == "hello") {
        var v5 uint64 = 345
        for fi6 := 0; fi6 < 13; fi6++ {
            v1 = v1
            fmt.Println(fi6, v1)
            _ = fi6
        }
        if (33 <= int8(v1)) {
            switch "" {
            case "a":
                var v7 uint = 845
                _ = v7
            case "\n\t":
                for fi8 := 0; fi8 < 10; fi8++ {
                    v1 = v1
                    _ = fi8
                }
            }
            if (-896 <= 42) {
                for fi9 := 0; fi9 < 9; fi9++ {
                    switch "foo bar" {
                    case "":
                        delete(v4, 2147483648)
                    case "hello":
                        var v10 uint8 = 145
                        _ = v10
                    default:
                        if ((-(v1)) < v1) {
                            _ = len(v4)
                            var v11 int16 = 1000
                            for fi12 := 0; fi12 < 7; fi12++ {
                                var v13 uint64 = (v5 % (((^(v5))) | 1))
                                var v14 uint = 255
                                switch (fi12 & int(v14)) {
                                case 128:
                                    fmt.Println(v1)
                                case -1:
                                    v5 -= (v5 - v13)
                                case 32768:
                                    if (1.5 > ((3.14159 * (0.1 + float32(0.1))) + 1.5)) {
                                        var v15 uint16 = 0
                                        _ = len(v4)
                                        _ = v15
                                    } else {
                                        v1 -= v1
                                    }
                                }
                                _ = fi12
                                _ = v13
                                _ = v14
                            }
                            _ = v11
                        } else {
                            for fi16 := 0; fi16 < 10; fi16++ {
                                v17 := GenS1{F0: 9223372036854775807, F1: -522}
                                var v18 uint = 18446744073709551615
                                delete(v4, 65535)
                                switch "a" {
                                case "":
                                    fmt.Println(v18, fi9, v5)
                                case "foo bar":
                                    v18 = 560
                                default:
                                    v2 = append(v2, "foo bar")
                                }
                                fmt.Println(v18)
                                _ = fi16
                                _ = v17
                                _ = v18
                            }
                            v19 := GenS2{F0: (v5 * 4294967295), F1: 2, F2: false}
                            for fi20 := 0; fi20 < 7; fi20++ {
                                v1 = -913
                                _ = fi20
                            }
                            _ = v19
                        }
                    }
                    _ = fi9
                }
            } else {
                v21 := GenS3{F0: 2}
                _ = v21
            }
            for fi22 := 0; fi22 < 14; fi22++ {
                for fi23 := 0; fi23 < 18; fi23++ {
                    var v24 float32 = ((float32(1.5) - ((1.5 / 1e10) * ((1.5 + 1.5) * 3.14159))) + 0.1)
                    v25 := GenS4{F0: 2, F1: 4294967295, F2: int64(fi23)}
                    if len(v2) > 0 {
                        v2[0] = "a"
                    }
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                var v26 int16 = v1
                _ = fi22
                _ = v26
            }
        }
        _ = v5
    }
    v1 += -1
    if (!(322 <= -1)) {
        var v27 string = (("\n\t" + ("a" + "a")) + ("" + ""))
        for fi28 := 0; fi28 < 18; fi28++ {
            for fi29 := 0; fi29 < 14; fi29++ {
                var v30 uint32 = 2147483648
                _ = len(v4)
                if (!(!((65535 < 262) && ((("a" + "hello") + v27) >= "")))) {
                    v1 -= v1
                    v30 = v30
                    v27 = v27
                    v1 = 906
                } else {
                    if (!true) {
                        var v31 int8 = 1
                        switch "" {
                        case "hello":
                            if false {
                                if (false || ((fi28 % ((fi28) | 1)) <= (fi28 ^ fi28))) {
                                    v32 := GenS5{F0: -423, F1: 1}
                                    _ = len(v4)
                                    _ = v32
                                }
                                switch "a" {
                                case "\n\t":
                                    v31 -= (v31 / (((v31 | 100)) | 1))
                                    v4[(^(v30))] = v30
                                    v30 -= (v30 << ((18446744073709551615) & 63))
                                case "hello":
                                    var v33 string = v27
                                    _ = v33
                                }
                            } else {
                                v27 = v27
                            }
                        }
                        _ = v31
                    }
                }
                switch (((v27 + v27) + "a") + "a") {
                case "":
                    if len(v3) > 0 {
                        v3[0] = (v30 << ((65535) & 63))
                    }
                case "hello":
                    v30 += (v30 << ((845) & 63))
                }
                _ = len(v4)
                _ = fi29
                _ = v30
            }
            v34 := GenS1{F0: 42, F1: 949}
            switch 0 {
            case -2147483648:
                var v35 bool = (false || ((true || true) && true))
                _ = v35
            }
            _ = fi28
            _ = v34
        }
        _ = v27
    } else {
        for fi36 := 0; fi36 < 13; fi36++ {
            var v37 float64 = 1.5
            var v38 uint = 255
            var v39 int8 = 127
            _ = fi36
            _ = v37
            _ = v38
            _ = v39
        }
        switch -970 {
        case 32767:
            fmt.Println(v1)
        case 32768:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case -1:
            var v40 bool = true
            _ = v40
        default:
            switch "hello" {
            case "a":
                var v41 int8 = 100
                _ = v41
            case "\n\t":
                v1 += v1
            case "foo bar":
                v3 = append(v3, 2)
            }
        }
    }
    var v42 uint8 = 127
    var v43 uint64 = 0
    var v44 float64 = 0.1
    var v45 int = -391
    v42 = (^(v42))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
}
