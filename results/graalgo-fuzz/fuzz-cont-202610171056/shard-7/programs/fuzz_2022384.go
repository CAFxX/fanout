package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64; F1 bool }
type GenS4 struct { F0 string; F1 int32; F2 int8 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int8 }

func main() {
    var v1 float64 = 0.1
    v2 := []uint32{496, 0, 314}
    v3 := []int{170}
    v4 := make(map[uint32]int)
    for fi5 := 0; fi5 < 13; fi5++ {
        v2 = append(v2, 4294967295)
        _ = fi5
    }
    switch -2147483648 {
    case -508:
        for fi6 := 0; fi6 < 2; fi6++ {
            for fi7 := 0; fi7 < 2; fi7++ {
                if len(v3) > 0 {
                    v3[0] = fi6
                }
                v8 := GenS1{F0: 84, F1: 65535, F2: fi7}
                switch (fi6 >> ((479) & 63)) {
                case -2147483648:
                    if true {
                        switch fi6 {
                        case -84:
                            fmt.Println(fi7, fi7)
                        }
                        _ = len(v4)
                        if (2 == 1) {
                            var v9 int16 = 1
                            _ = v9
                        }
                        switch ("hello" + "hello") {
                        case "a":
                            v10 := GenS1{F0: -1, F1: 1, F2: fi7}
                            _ = v10
                        case "hello":
                            switch "foo bar" {
                            case "hello":
                                fmt.Println(fi7, fi7)
                            default:
                                if false {
                                    switch fi6 {
                                    case 9223372036854775807:
                                        v11 := GenS1{F0: -1, F1: 2, F2: fi6}
                                        _ = len(v2)
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        _ = v11
                                    }
                                    v12 := v4[672]
                                    _ = v12
                                } else {
                                    var v13 int8 = 53
                                    if true {
                                        var v14 uint = 42
                                        _ = v14
                                    }
                                    _ = len(v4)
                                    _ = v13
                                }
                            }
                        }
                        v1 -= -2.25
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v1 = 0.0
                    }
                default:
                    switch fi6 {
                    case 200:
                        if (!(0 <= 1)) {
                            v1 += v1
                        }
                    }
                }
                v3 = append(v3, (fi7 / ((32767) | 1)))
                if true {
                    var v15 uint8 = 255
                    _ = v15
                }
                v16 := GenS2{F0: -32768}
                v17 := GenS3{F0: 42, F1: (true && (!(42 >= -1)))}
                _ = fi7
                _ = v16
                _ = v17
                _ = v8
            }
            for fi18 := 0; fi18 < 14; fi18++ {
                v1 -= (1e10 * float64(0.0))
                fmt.Println(fi18, fi6)
                _ = fi18
            }
            _ = fi6
        }
    case -2147483648:
        for fi19 := 0; fi19 < 20; fi19++ {
            v4[2] = 1
            _ = fi19
        }
    case 371:
        for fi20 := 0; fi20 < 17; fi20++ {
            if false {
                var v21 float32 = 3.14159
                _ = v21
            } else {
                v1 -= 0.0
            }
            switch fi20 {
            case -9223372036854775808:
                var v22 int64 = 0
                _ = v22
            case 42:
                for fi23 := 0; fi23 < 14; fi23++ {
                    if true {
                        var v24 uint64 = 21
                        _ = v24
                    } else {
                        v2 = append(v2, 0)
                    }
                    _ = fi23
                }
            }
            fmt.Println(fi20, v1)
            _ = fi20
        }
    default:
        _ = len(v4)
    }
    if ((1e10 - 3.14159) < 1e10) {
        for fi25 := 0; fi25 < 9; fi25++ {
            var v26 string = ""
            v27 := GenS4{F0: v26, F1: 886, F2: 127}
            v1 = (0.0 - float64((1.5 * ((0.0 * 0.1) - -2.25))))
            _ = fi25
            _ = v26
            _ = v27
        }
        var v28 int64 = -1
        _ = v28
    } else {
        if len(v2) > 0 {
            v2[0] = 1
        }
        switch -2147483648 {
        case -684:
            v1 = (3.14159 * v1)
        case 42:
            v29 := GenS5{F0: 1}
            _ = v29
        case 128:
            v30 := GenS6{F0: 100}
            _ = v30
        default:
            for fi31 := 0; fi31 < 20; fi31++ {
                var v32 int32 = 826
                if ((!(!((fi31 << ((0) & 63)) >= -1))) || (!false)) {
                    delete(v4, 2)
                    fmt.Println(v32, v32)
                }
                v32 -= (v32 >> ((0) & 63))
                _ = fi31
                _ = v32
            }
        }
    }
    var v33 int32 = 2147483647
    var v34 uint64 = 0
    v1 = 1.5
    for fi35 := 0; fi35 < 11; fi35++ {
        var v36 uint = 18446744073709551615
        _ = fi35
        _ = v36
    }
    v37 := GenS2{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v37
    _ = v4
}
