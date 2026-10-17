package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 int; F1 uint16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int16; F1 uint64 }

func main() {
    var v1 int32 = 42
    var v2 string = ("a" + "foo bar")
    var v3 uint64 = 2
    for fi4 := 0; fi4 < 5; fi4++ {
        var v5 int64 = 42
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 18; fi6++ {
        switch v2 {
        case "\n\t":
            var v7 uint64 = v3
            _ = v7
        case "a":
            for fi8 := 0; fi8 < 5; fi8++ {
                var v9 uint = 891
                if (1 > 1) {
                    fmt.Println(v1, v1, v1)
                } else {
                    var v10 int = fi8
                    v9 = (v9 | 556)
                    _ = v10
                }
                var v11 uint = 18446744073709551615
                if false {
                    var v12 uint8 = 209
                    _ = v12
                } else {
                    v2 = v2
                    switch "" {
                    case "hello":
                        if (!((v1 << ((v9) & 63)) <= (v1 * -2147483648))) {
                            var v13 uint8 = 237
                            _ = v13
                        }
                    default:
                        var v14 int64 = 0
                        _ = v14
                    }
                    v9 += 2
                    v9 -= v9
                }
                switch (v2 + v2) {
                case "hello":
                    fmt.Println(v2, v11)
                case "":
                    v1 += v1
                case "a":
                    v15 := []int{fi8, (fi8 + fi6), -574}
                    v16 := GenS1{F0: 4294967295, F1: ((v2 == (v2 + v2)) && (!(!((v11 << ((uint(v3)) & 63)) > 0))))}
                    v17 := GenS2{F0: (v11 << ((v11) & 63)), F1: ("" + (v2 + v2))}
                    _ = v15
                    _ = v16
                    _ = v17
                default:
                    switch fi6 {
                    case -863:
                        var v18 uint = 1
                        _ = v18
                    case -1:
                        v1 = (v1 >> ((v11) & 63))
                    case 32768:
                        var v19 bool = true
                        _ = v19
                    }
                }
                v20 := GenS3{F0: fi8, F1: 1}
                switch v2 {
                case "a":
                    var v21 uint = v9
                    _ = v21
                case "\n\t":
                    v1 = v1
                case "":
                    fmt.Println(v9, v11, v2)
                default:
                    if false {
                        v9 -= 141
                    }
                }
                _ = fi8
                _ = v11
                _ = v20
                _ = v9
            }
        case "foo bar":
            v22 := []int{-689, fi6}
            _ = v22
        default:
            v23 := []int{(fi6 % ((878) | 1)), (fi6 / (((fi6 ^ fi6)) | 1)), (fi6 << ((395) & 63)), fi6}
            _ = v23
        }
        v24 := []int{}
        for fi25 := 0; fi25 < 5; fi25++ {
            if false {
                v24 = append(v24, fi25)
            } else {
                fmt.Println(fi6, fi25, fi6)
            }
            fmt.Println(fi6)
            if (!(!(0 > 1))) {
                var v26 int16 = -1
                v27 := GenS1{F0: 894, F1: ((v3 != v3) && false)}
                _ = v26
                _ = v27
            } else {
                v3 = v3
                if true {
                    v24 = append(v24, 0)
                    var v28 int8 = 83
                    _ = v28
                }
                fmt.Println(v3, v1)
            }
            _ = fi25
        }
        _ = fi6
        _ = v24
    }
    v1 = v1
    switch -556 {
    case 128:
        var v29 float32 = (0.1 + float32(float64(((0.0 - (((3.14159 / 1.5) - (3.14159 + float32(3.14159))) / float32(1e10))) + (1e10 / 0.1)))))
        _ = v29
    default:
        for fi30 := 0; fi30 < 19; fi30++ {
            for fi31 := 0; fi31 < 4; fi31++ {
                v2 = ("foo bar" + ((v2 + "foo bar") + (v2 + ((v2 + "hello") + v2))))
                switch (fi31 + fi30) {
                case 672:
                    var v32 uint32 = 174
                    _ = v32
                case 2147483647:
                    switch 349 {
                    case 42:
                        var v33 int8 = 1
                        _ = v33
                    default:
                        v34 := []int{1, fi30, (fi31 % ((fi30) | 1))}
                        _ = v34
                    }
                case 281:
                    v35 := []int{fi30}
                    _ = v35
                }
                if (int16(fi30) < 0) {
                    v3 = v3
                    var v36 uint64 = (v3 % ((623) | 1))
                    _ = v36
                }
                _ = fi31
            }
            for fi37 := 0; fi37 < 18; fi37++ {
                var v38 float32 = (3.14159 * 1.5)
                _ = fi37
                _ = v38
            }
            var v39 string = v2
            _ = fi30
            _ = v39
        }
    }
    for fi40 := 0; fi40 < 2; fi40++ {
        for fi41 := 0; fi41 < 5; fi41++ {
            var v42 uint16 = 712
            v43 := GenS4{F0: 947}
            v42 += v42
            _ = fi41
            _ = v42
            _ = v43
        }
        fmt.Println(v2, v1, fi40)
        var v44 string = v2
        _ = fi40
        _ = v44
    }
    fmt.Println(v1, v2)
    for fi45 := 0; fi45 < 9; fi45++ {
        v46 := []int{fi45}
        var v47 uint8 = 128
        _ = fi45
        _ = v46
        _ = v47
    }
    if true {
        v48 := []int{}
        _ = v48
    }
    v49 := []int{2147483647, 9223372036854775807}
    switch 32768 {
    case 42:
        v49 = append(v49, -160)
    case 32768:
        v49 = append(v49, 128)
    case -1:
        v3 += 2
    default:
        if true {
            v3 = v3
            v50 := GenS1{F0: 255, F1: true}
            switch "foo bar" {
            case "\n\t":
                var v51 uint16 = 65535
                _ = v51
            case "hello":
                switch ("\n\t" + v2) {
                case "":
                    for fi52 := 0; fi52 < 17; fi52++ {
                        var v53 string = "\n\t"
                        v2 = "hello"
                        switch fi52 {
                        case 128:
                            v49 = append(v49, (fi52 % ((42) | 1)))
                        case 2147483647:
                            v49 = append(v49, fi52)
                        default:
                            var v54 bool = ((!false) && (1 <= 65535))
                            _ = v54
                        }
                        _ = fi52
                        _ = v53
                    }
                }
            case "":
                v55 := GenS5{F0: 0, F1: (v3 * (v3 / (((v3 & v3)) | 1)))}
                _ = v55
            default:
                if ((^(v3)) != 42) {
                    fmt.Println(v3, v3)
                    for fi56 := 0; fi56 < 11; fi56++ {
                        var v57 bool = (1 == 2147483648)
                        var v58 int = -617
                        _ = fi56
                        _ = v57
                        _ = v58
                    }
                    v3 -= v3
                }
            }
            _ = v50
        } else {
            _ = len(v49)
        }
    }
    if true {
        _ = len(v49)
        v1 -= v1
        v3 -= (v3 >> ((1) & 63))
    } else {
        if true {
            for fi59 := 0; fi59 < 1; fi59++ {
                var v60 float64 = 3.14159
                var v61 string = (v2 + v2)
                var v62 int16 = 941
                _ = fi59
                _ = v60
                _ = v61
                _ = v62
            }
        }
        for fi63 := 0; fi63 < 3; fi63++ {
            if true {
                var v64 int64 = -1
                _ = v64
            }
            _ = fi63
        }
    }
    fmt.Println(v2, v2, v2)
    fmt.Println(v3)
    v49 = append(v49, 127)
    v3 += v3
    var v65 int8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v49:", len(v49), v49)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v49
    _ = v65
}
