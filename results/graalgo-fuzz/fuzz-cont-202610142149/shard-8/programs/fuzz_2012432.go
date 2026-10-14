package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 string; F1 int8; F2 uint }
type GenS3 struct { F0 uint; F1 uint32; F2 int8 }
type GenS4 struct { F0 uint16 }

func main() {
    var v1 int = -373
    var v2 float32 = (0.1 - float32(1e10))
    var v3 string = "a"
    var v4 int32 = 42
    v5 := make(map[int]int)
    v3 = (v3 + ("a" + (("foo bar" + "a") + v3)))
    for fi6 := 0; fi6 < 11; fi6++ {
        fmt.Println(v1, v1, fi6)
        v7 := []int{}
        _ = fi6
        _ = v7
    }
    if false {
        var v8 float32 = float32(0.0)
        v8 += v2
        v9 := GenS1{F0: 2, F1: -32768}
        _ = v8
        _ = v9
    }
    if ((true || (127 >= 0)) || (42 >= 0)) {
        var v10 uint64 = 4294967295
        var v11 string = ("hello" + ("foo bar" + "\n\t"))
        _ = v10
        _ = v11
    } else {
        var v12 uint = 659
        if (((v4 % (((v4 >> (((v12 * 2)) & 63))) | 1)) > 580) || false) {
            v4 += 0
        }
        _ = v12
    }
    switch v1 {
    case -210:
        for fi13 := 0; fi13 < 15; fi13++ {
            var v14 uint16 = 1
            switch v3 {
            case "a":
                v3 = "a"
            default:
                v4 += -1
            }
            _ = fi13
            _ = v14
        }
    case -2147483648:
        if false {
            v15 := v5[(v1 << ((2) & 63))]
            var v16 uint8 = 99
            v15 -= 1
            _ = v15
            _ = v16
        } else {
            switch v3 {
            case "hello":
                v17 := []int{(-(v1))}
                _ = v17
            case "foo bar":
                v18 := []int{v1}
                _ = v18
            }
            v2 = ((v2 + (((1e10 / -2.25) + v2) - 1.5)) * v2)
            if (!(32767 != 42)) {
                v19 := GenS1{F0: 0, F1: 1}
                _ = v19
            } else {
                for fi20 := 0; fi20 < 7; fi20++ {
                    var v21 uint32 = 247
                    _ = len(v5)
                    _ = fi20
                    _ = v21
                }
                v22 := GenS1{F0: 255, F1: -520}
                var v23 int16 = 598
                _ = v22
                _ = v23
            }
        }
    case 128:
        v1 = 0
    default:
        v1 -= -176
    }
    var v24 uint8 = 128
    v25 := GenS1{F0: 257, F1: 1000}
    v1 += v1
    if ((false && (42 != 42)) || (0 < 42)) {
        for fi26 := 0; fi26 < 10; fi26++ {
            var v27 uint32 = 4294967295
            var v28 uint = 1
            if false {
                var v29 int = v1
                var v30 int16 = 42
                var v31 string = v3
                _ = v29
                _ = v30
                _ = v31
            } else {
                v2 = float32(0.1)
            }
            _ = fi26
            _ = v27
            _ = v28
        }
    }
    for fi32 := 0; fi32 < 11; fi32++ {
        v3 = v3
        fmt.Println(v3, fi32)
        _ = fi32
    }
    v33 := GenS1{F0: 42, F1: 0}
    v24 += (v24 % ((v24) | 1))
    var v34 uint32 = uint32(v24)
    v3 = v3
    v3 = v3
    for fi35 := 0; fi35 < 9; fi35++ {
        if ((1 < 42) || false) {
            v4 -= (v4 + (v4 - -132))
            if ((int8(v1) < 42) || false) {
                delete(v5, (fi35 + fi35))
                if true {
                    var v36 uint64 = 2
                    fmt.Println(v2)
                    v2 = (((1e10 + 1.5) * (v2 / v2)) / (1e10 * (v2 + v2)))
                    _ = v36
                }
                var v37 uint16 = 255
                _ = v37
            }
        } else {
            switch (v3 + (v3 + v3)) {
            case "\n\t":
                var v38 uint64 = 18446744073709551615
                _ = v38
            }
            if (0 >= -32768) {
                var v39 uint8 = (^(v24))
                _ = v39
            } else {
                v40 := GenS1{F0: 42, F1: 1000}
                v41 := GenS2{F0: v3, F1: -106, F2: 4294967295}
                _ = v40
                _ = v41
            }
        }
        var v42 uint64 = 18446744073709551615
        for fi43 := 0; fi43 < 15; fi43++ {
            var v44 uint = 65535
            switch v3 {
            case "foo bar":
                v45 := GenS1{F0: 201, F1: 27}
                _ = v45
            case "hello":
                if false {
                    var v46 uint16 = 2
                    if true {
                        if true {
                            if false {
                                v2 = v2
                                var v47 int8 = -1
                                switch (v3 + (v3 + "hello")) {
                                case "\n\t":
                                    v34 = v34
                                }
                                _ = v47
                            }
                            var v48 float32 = -2.25
                            _ = v48
                        } else {
                            switch (v1 & fi35) {
                            case 807:
                                var v49 int64 = 891
                                _ = v49
                            case 305:
                                fmt.Println(v1, v3)
                            }
                            switch ((v3 + "") + (("a" + ("\n\t" + "hello")) + "foo bar")) {
                            case "foo bar":
                                var v50 uint16 = uint16(v34)
                                _ = v50
                            case "a":
                                v24 = v24
                            }
                            var v51 uint8 = 2
                            _ = v51
                        }
                        switch (fi35 / ((v1) | 1)) {
                        case 32767:
                            var v52 uint64 = (v42 % (((v42 - 861)) | 1))
                            _ = v52
                        case 825:
                            v53 := GenS3{F0: 4294967295, F1: uint32(v44), F2: -3}
                            v1 += -1
                            v54 := GenS2{F0: ((v3 + v3) + "\n\t"), F1: 127, F2: 255}
                            _ = v53
                            _ = v54
                        case 127:
                            v55 := GenS4{F0: 745}
                            _ = v55
                        }
                    } else {
                        switch v3 {
                        case "":
                            var v56 int64 = -89
                            _ = v56
                        }
                        var v57 uint = v44
                        if (v3 > v3) {
                            v3 = "foo bar"
                            if false {
                                v34 = 2
                                switch (fi35 - (v1 + 42)) {
                                case 2147483647:
                                    v34 -= (v34 << (((v57 / (((v57 / (((v57 / ((858) | 1))) | 1))) | 1))) & 63))
                                case 128:
                                    var v58 int16 = 42
                                    _ = v58
                                }
                                v59 := GenS3{F0: (v44 << ((4294967295) & 63)), F1: (v34 - (v34 ^ (^(v34)))), F2: 0}
                                v2 = v2
                                _ = v59
                            }
                            v5[449] = (v1 ^ 32767)
                            v2 += (0.1 - (v2 + v2))
                        } else {
                            var v60 float32 = v2
                            var v61 int8 = 0
                            var v62 uint8 = 31
                            _ = v60
                            _ = v61
                            _ = v62
                        }
                        _ = v57
                    }
                    _ = v46
                }
            case "":
                var v63 uint8 = 1
                _ = v63
            }
            v64 := GenS2{F0: (((v3 + "\n\t") + (v3 + (v3 + v3))) + v3), F1: 127, F2: v44}
            _ = fi43
            _ = v44
            _ = v64
        }
        _ = fi35
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
