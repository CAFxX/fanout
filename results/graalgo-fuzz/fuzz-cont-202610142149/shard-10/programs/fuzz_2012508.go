package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint16; F1 int; F2 uint64 }
type GenS5 struct { F0 int64; F1 int64 }
func genFn0(p0 bool, p1 uint8, p2 int64) uint32 {
    var v5 uint16 = 32767
    v6 := GenS1{F0: 42, F1: 1}
    p0 = p0
    if true {
        var v7 uint64 = 0
        _ = v7
    } else {
        v8 := []int{}
        v5 = (v5 / (((v5 - 899)) | 1))
        v9 := GenS2{F0: 4294967295, F1: (v5 * (v5 * v5))}
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 449
}
func genFn1() uint {
    var v10 string = "foo bar"
    if true {
        switch "foo bar" {
        case "foo bar":
            v11 := GenS1{F0: 0, F1: 0}
            _ = v11
        }
    } else {
        for fi12 := 0; fi12 < 5; fi12++ {
            v10 = v10
            _ = fi12
        }
        v10 = v10
        v10 = (v10 + v10)
    }
    v13 := GenS3{F0: 419}
    v14 := GenS2{F0: 485, F1: 255}
    fmt.Println(v10)
    _ = v10
    _ = v13
    _ = v14
    return 255
}

func main() {
    var v1 int64 = 114
    v2 := make(map[uint32]string)
    v3 := []uint32{}
    v4 := []int{-9223372036854775808, 1, 1, -1}
    var v15 int8 = 1
    var v16 int64 = (v1 >> ((65535) & 63))
    if len(v4) > 0 {
        v4[0] = 38
    }
    for fi17 := 0; fi17 < 15; fi17++ {
        for fi18 := 0; fi18 < 4; fi18++ {
            v1 -= v1
            v2[42] = ("" + "a")
            _ = fi18
        }
        v19 := GenS1{F0: 1000, F1: (fi17 | -142)}
        _ = fi17
        _ = v19
    }
    delete(v2, 65535)
    var v20 int16 = 0
    for fi21 := 0; fi21 < 1; fi21++ {
        for fi22 := 0; fi22 < 3; fi22++ {
            if len(v3) > 0 {
                v3[0] = 2
            }
            if false {
                v20 -= 143
            } else {
                v16 = 471
            }
            v23 := v2[0]
            _ = fi22
            _ = v23
        }
        _ = fi21
    }
    v2[2] = "a"
    v24 := GenS1{F0: v20, F1: 0}
    fmt.Println(v1, v15, v16)
    if true {
        for fi25 := 0; fi25 < 10; fi25++ {
            var v26 string = ("a" + "hello")
            _ = fi25
            _ = v26
        }
        if (!true) {
            fmt.Println(v1, v15, v15)
        }
    } else {
        for fi27 := 0; fi27 < 11; fi27++ {
            switch (-(fi27)) {
            case 977:
                v15 += 0
            case 826:
                if (320 <= 255) {
                    v1 = v1
                    var v28 int32 = -818
                    _ = v28
                }
            default:
                for fi29 := 0; fi29 < 17; fi29++ {
                    v15 += v15
                    var v30 int32 = 0
                    v31 := GenS4{F0: 65535, F1: 32768, F2: 42}
                    _ = fi29
                    _ = v30
                    _ = v31
                }
            }
            v1 += 0
            _ = fi27
        }
        var v32 int16 = 42
        for fi33 := 0; fi33 < 9; fi33++ {
            v20 = int16(v1)
            switch ("\n\t" + (("\n\t" + "\n\t") + ("" + "hello"))) {
            case "\n\t":
                if false {
                    v34 := GenS3{F0: genFn1()}
                    _ = v34
                } else {
                    for fi35 := 0; fi35 < 20; fi35++ {
                        if false {
                            v4 = append(v4, -2147483648)
                            switch -362 {
                            case 32768:
                                delete(v2, 800)
                            case 42:
                                var v36 uint16 = 167
                                _ = v36
                            case -735:
                                v16 = -176
                            }
                            v37 := GenS5{F0: 9223372036854775807, F1: -9223372036854775808}
                            _ = v37
                        } else {
                            v32 = v32
                        }
                        fmt.Println(v16, v16, v16)
                        var v38 bool = true
                        _ = fi35
                        _ = v38
                    }
                    var v39 uint8 = 127
                    _ = v39
                }
            case "a":
                var v40 int64 = (v16 | (v1 * v1))
                _ = v40
            default:
                switch "a" {
                case "hello":
                    if ((875 < 42) || false) {
                        v41 := v2[42]
                        _ = v41
                    }
                case "\n\t":
                    if (932 == 21) {
                        fmt.Println(v16, fi33)
                        fmt.Println(v20, v1)
                        var v42 string = ""
                        _ = v42
                    }
                }
            }
            _ = fi33
        }
        _ = v32
    }
    var v43 uint16 = 2
    _ = len(v2)
    for fi44 := 0; fi44 < 12; fi44++ {
        var v45 string = "hello"
        var v46 uint16 = (v43 >> ((735) & 63))
        _ = fi44
        _ = v45
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
    _ = v43
}
