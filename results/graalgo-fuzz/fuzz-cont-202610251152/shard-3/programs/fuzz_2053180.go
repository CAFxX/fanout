package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 uint }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int; F1 uint8 }
type GenS5 struct { F0 int16; F1 bool }

func main() {
    v1 := make(map[int]string)
    var v2 bool = (float32(0.1) > 1.5)
    var v3 int16 = -1
    if (!(!v2)) {
        v3 += (v3 % (((-(v3))) | 1))
        var v4 bool = true
        if true {
            fmt.Println(v4, v2)
        } else {
            v3 = v3
        }
        _ = v4
    } else {
        v3 = (v3 % (((v3 * v3)) | 1))
    }
    if (153 <= 995) {
        v5 := GenS1{F0: 65535, F1: "", F2: 626}
        for fi6 := 0; fi6 < 6; fi6++ {
            switch (fi6 | fi6) {
            case 32767:
                switch fi6 {
                case 32767:
                    v7 := GenS2{F0: 127, F1: true}
                    _ = v7
                case 2147483647:
                    switch ((("foo bar" + "\n\t") + ((("\n\t" + "hello") + ("hello" + "\n\t")) + "hello")) + (("" + "\n\t") + (("" + "foo bar") + ("hello" + "foo bar")))) {
                    case "foo bar":
                        v8 := []int{-1}
                        _ = v8
                    }
                case -1:
                    var v9 bool = v2
                    _ = v9
                default:
                    if v2 {
                        v10 := GenS3{F0: 0}
                        _ = v10
                    }
                }
            default:
                var v11 float64 = (3.14159 + float64(0.0))
                _ = v11
            }
            _ = fi6
        }
        v12 := GenS3{F0: 135}
        _ = v12
        _ = v5
    } else {
        v13 := []int{918}
        _ = v13
    }
    var v14 uint64 = 0
    switch "hello" {
    case "a":
        var v15 int8 = 100
        _ = v15
    case "foo bar":
        if (210 <= 42) {
            if (("a" < "\n\t") && v2) {
                v3 -= (v3 % ((42) | 1))
                if ((v2 && v2) || v2) {
                    var v16 uint32 = 512
                    v17 := GenS4{F0: -955, F1: 164}
                    _ = v16
                    _ = v17
                }
                for fi18 := 0; fi18 < 20; fi18++ {
                    for fi19 := 0; fi19 < 16; fi19++ {
                        v20 := GenS3{F0: 210}
                        _ = fi19
                        _ = v20
                    }
                    v14 = v14
                    if v2 {
                        if true {
                            delete(v1, 2147483647)
                        } else {
                            switch fi18 {
                            case 32768:
                                fmt.Println(fi18, v14, v2)
                            case 9223372036854775807:
                                v1[(^(fi18))] = ("a" + "\n\t")
                            case -544:
                                var v21 uint8 = 128
                                _ = v21
                            default:
                                for fi22 := 0; fi22 < 18; fi22++ {
                                    if v2 {
                                        v2 = true
                                        v23 := GenS5{F0: (v3 & v3), F1: (v2 && v2)}
                                        _ = v23
                                    } else {
                                        v24 := GenS1{F0: 65535, F1: "foo bar", F2: 788}
                                        delete(v1, fi22)
                                        _ = v24
                                    }
                                    fmt.Println(v3, v2)
                                    if (2 == 1) {
                                        v25 := []int{-791, fi22, (fi18 * (fi22 * fi22)), -1}
                                        var v26 uint32 = 686
                                        _ = v25
                                        _ = v26
                                    }
                                    _ = fi22
                                }
                            }
                            if false {
                                _ = len(v1)
                            } else {
                                v14 += v14
                            }
                            if false {
                                v27 := GenS1{F0: 4294967295, F1: "a", F2: 2}
                                _ = v27
                            } else {
                                v28 := v1[fi18]
                                _ = v28
                            }
                        }
                        v14 += v14
                    } else {
                        v2 = ((3.14159 * (-2.25 + float32(0.0))) >= 1.5)
                    }
                    _ = fi18
                }
            }
            v29 := GenS5{F0: -32768, F1: ((!v2) || (18446744073709551615 != 18446744073709551615))}
            for fi30 := 0; fi30 < 18; fi30++ {
                v2 = v2
                var v31 float32 = 1.5
                v32 := GenS3{F0: 81}
                _ = fi30
                _ = v31
                _ = v32
            }
            _ = v29
        } else {
            v1[-610] = (("a" + "foo bar") + "foo bar")
        }
    case "hello":
        fmt.Println(v2)
    }
    fmt.Println(v3)
    v33 := []int{-1, -1, 617}
    fmt.Println(v14)
    if len(v33) >= 2 {
        v33 = v33[1:len(v33)]
    }
    var v34 uint32 = 692
    var v35 float64 = (1.5 - 3.14159)
    var v36 uint64 = uint64(v34)
    if v2 {
        switch ((("" + "hello") + "\n\t") + "a") {
        case "hello":
            if len(v33) >= 2 {
                v33 = v33[1:len(v33)]
            }
        case "\n\t":
            v14 += (v36 + 699)
        case "foo bar":
            if v2 {
                fmt.Println(v2, v2, v2)
            } else {
                var v37 float64 = v35
                if ((v34 << ((143) & 63)) < v34) {
                    var v38 int8 = 42
                    var v39 int8 = (v38 | (v38 + (v38 ^ v38)))
                    var v40 int32 = 42
                    _ = v38
                    _ = v39
                    _ = v40
                }
                _ = v37
            }
        }
    }
    delete(v1, 9223372036854775807)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
}
