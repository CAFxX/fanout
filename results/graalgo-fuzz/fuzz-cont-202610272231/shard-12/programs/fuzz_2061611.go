package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string; F1 bool }

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 838
    v3 := make(map[uint32]string)
    var v4 uint8 = uint8(v2)
    var v5 int = -162
    var v6 float64 = 1.5
    v7 := GenS1{F0: 1, F1: 1}
    var v8 string = "\n\t"
    var v9 int8 = -23
    if ((-376 < 0) || (false && false)) {
        fmt.Println(v2, v5)
        var v10 uint16 = (v2 % (((v2 ^ v2)) | 1))
        v11 := GenS2{F0: 42}
        _ = v10
        _ = v11
    }
    if true {
        var v12 int64 = -1
        _ = v12
    }
    fmt.Println(v8, v2, v6)
    v2 = v2
    v13 := GenS3{F0: v8}
    var v14 float64 = -2.25
    fmt.Println(v8)
    var v15 int16 = 42
    v5 = (^(v5))
    switch v5 {
    case 32767:
        switch "foo bar" {
        case "":
            if false {
                var v16 uint = 65535
                var v17 float64 = float64(1.5)
                for fi18 := 0; fi18 < 3; fi18++ {
                    var v19 int16 = v15
                    var v20 int64 = -9223372036854775808
                    switch "\n\t" {
                    case "":
                        v9 -= v9
                    }
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                _ = v16
                _ = v17
            }
        default:
            v21 := GenS4{F0: (("foo bar" + (v8 + (("a" + (v8 + ((((("a" + v8) + v8) + (((v8 + "foo bar") + (v8 + v8)) + ((v8 + ((v8 + "foo bar") + (v8 + "a"))) + v8))) + v8) + "a"))) + "foo bar"))) + "foo bar"), F1: true}
            _ = v21
        }
    case -9223372036854775808:
        for fi22 := 0; fi22 < 1; fi22++ {
            for fi23 := 0; fi23 < 9; fi23++ {
                var v24 string = v8
                _ = len(v3)
                _ = fi23
                _ = v24
            }
            for fi25 := 0; fi25 < 15; fi25++ {
                v6 -= v14
                if true {
                    switch "hello" {
                    case "":
                        v26 := []int{fi25, fi25}
                        _ = v26
                    case "foo bar":
                        var v27 int32 = 1
                        _ = v27
                    case "hello":
                        v28 := []int{(v5 / ((4294967295) | 1)), fi25, fi25}
                        _ = v28
                    }
                    if (((v2 & (v2 + (v2 * v2))) >= (^(v2))) || false) {
                        var v29 float64 = v14
                        _ = v29
                    }
                } else {
                    if (!(v8 < v8)) {
                        v2 += v2
                        if false {
                            switch v8 {
                            case "\n\t":
                                if true {
                                    delete(v3, 322)
                                    var v30 float32 = -2.25
                                    _ = len(v3)
                                    _ = v30
                                } else {
                                    v2 += (v2 | (v2 & 1))
                                }
                            case "hello":
                                var v31 uint32 = 2147483648
                                _ = v31
                            default:
                                if false {
                                    var v32 uint16 = 0
                                    var v33 int32 = 42
                                    var v34 uint = 798
                                    _ = v32
                                    _ = v33
                                    _ = v34
                                }
                            }
                        }
                        fmt.Println(fi25, v15)
                    } else {
                        var v35 int8 = v9
                        switch v8 {
                        case "":
                            var v36 float64 = float64((1e10 / 1.5))
                            _ = v36
                        case "hello":
                            if (!true) {
                                switch fi25 {
                                case -531:
                                    fmt.Println(v15, v35)
                                case -153:
                                    v37 := GenS4{F0: v8, F1: true}
                                    _ = v37
                                default:
                                    var v38 int16 = v15
                                    _ = v38
                                }
                                switch v5 {
                                case 32767:
                                    fmt.Println(fi22)
                                default:
                                    var v39 int = v5
                                    _ = v39
                                }
                            } else {
                                fmt.Println(v9)
                            }
                        }
                        fmt.Println(v5)
                        _ = v35
                    }
                }
                var v40 int64 = -29
                _ = fi25
                _ = v40
            }
            _ = fi22
        }
    default:
        v14 = v6
    }
    for fi41 := 0; fi41 < 20; fi41++ {
        v2 -= (^(v2))
        _ = fi41
    }
    switch "foo bar" {
    case "a":
        var v42 int = (v5 / (((v5 + 364)) | 1))
        _ = v42
    case "":
        if false {
            v15 = (v15 >> ((212) & 63))
            var v43 uint32 = 42
            if false {
                var v44 float64 = v14
                var v45 int32 = 2147483647
                _ = v44
                _ = v45
            }
            _ = v43
        }
    case "hello":
        v46 := GenS3{F0: ("hello" + v8)}
        _ = v46
    default:
        v6 = (3.14159 + 1.5)
    }
    for fi47 := 0; fi47 < 7; fi47++ {
        for fi48 := 0; fi48 < 4; fi48++ {
            v49 := []int{1, (fi47 | (v5 * v5))}
            var v50 uint8 = (v4 << ((1) & 63))
            v5 += fi48
            _ = fi48
            _ = v49
            _ = v50
        }
        v51 := GenS2{F0: v15}
        v52 := v3[0]
        _ = fi47
        _ = v51
        _ = v52
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
