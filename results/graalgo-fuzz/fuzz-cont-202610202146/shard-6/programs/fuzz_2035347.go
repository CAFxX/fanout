package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint; F1 int32; F2 int64 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int8; F1 uint; F2 uint32 }
func genFn0() int8 {
    v4 := GenS1{F0: 30, F1: false}
    v5 := []int{-49, -2147483648, 823}
    v6 := GenS2{F0: 1, F1: "hello"}
    _ = v4
    _ = v5
    _ = v6
    return -1
}
func genFn1(p0 uint32, p1 uint8, p2 string) int32 {
    if (151 <= 32767) {
        for fi7 := 0; fi7 < 14; fi7++ {
            p1 += p1
            var v8 uint32 = 2147483647
            fmt.Println(p2)
            _ = fi7
            _ = v8
        }
        if true {
            var v9 float32 = 0.1
            _ = v9
        }
    }
    var v10 uint32 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return -209
}

func main() {
    var v1 uint16 = 42
    v2 := make(map[uint32]string)
    var v3 float64 = 1.5
    if false {
        var v11 int64 = -1
        fmt.Println(v3)
        _ = v11
    }
    v3 = 1.5
    v12 := GenS3{F0: 42, F1: 1}
    if (false && true) {
        fmt.Println(v3, v3)
        for fi13 := 0; fi13 < 10; fi13++ {
            fmt.Println(v1, v3)
            for fi14 := 0; fi14 < 2; fi14++ {
                if false {
                    var v15 uint64 = 65
                    v16 := []int{(fi14 << ((0) & 63)), fi14}
                    _ = v15
                    _ = v16
                }
                _ = fi14
            }
            if false {
                v3 += v3
                var v17 int64 = 9223372036854775807
                _ = v17
            }
            _ = fi13
        }
    } else {
        if (!false) {
            if false {
                v18 := []int{924, 520, 42, -999}
                if (1.5 == (((-2.25 * 0.1) / 1.0) + -2.25)) {
                    fmt.Println(v3, v1)
                    v1 = 65535
                    if (352 != 32767) {
                        switch "\n\t" {
                        case "hello":
                            v3 += v3
                        default:
                            switch "hello" {
                            case "\n\t":
                                for fi19 := 0; fi19 < 20; fi19++ {
                                    v20 := GenS2{F0: -128, F1: "hello"}
                                    _ = fi19
                                    _ = v20
                                }
                            }
                        }
                        v21 := GenS3{F0: 0, F1: 4294967295}
                        v1 = v1
                        _ = v21
                    }
                } else {
                    fmt.Println(v1)
                    v3 = 1e10
                }
                _ = v18
            }
            v22 := GenS4{F0: 0}
            _ = v22
        }
    }
    v23 := []int{936, -457, -2147483648, 128}
    switch 32767 {
    case 663:
        for fi24 := 0; fi24 < 10; fi24++ {
            for fi25 := 0; fi25 < 15; fi25++ {
                v1 -= (v1 / (((v1 * 1)) | 1))
                if len(v23) >= 2 {
                    v23 = v23[1:len(v23)]
                }
                var v26 int32 = 1
                _ = fi25
                _ = v26
            }
            v27 := GenS5{F0: 96, F1: -1, F2: -1}
            _ = fi24
            _ = v27
        }
    case 614:
        delete(v2, 42)
    }
    v28 := GenS3{F0: (^(v1)), F1: 2147483648}
    switch "foo bar" {
    case "hello":
        var v29 int16 = -32768
        _ = v29
    case "foo bar":
        fmt.Println(v1, v1)
    default:
        if false {
            fmt.Println(v1, v1)
        } else {
            v30 := GenS6{F0: 78}
            fmt.Println(v3, v1)
            v31 := GenS7{F0: 1, F1: 42, F2: 74}
            _ = v30
            _ = v31
        }
    }
    switch 9223372036854775807 {
    case 42:
        var v32 uint = 2
        _ = v32
    case 127:
        var v33 bool = true
        _ = v33
    case -9223372036854775808:
        if ((((3.14159 - v3) + v3) - (3.14159 + v3)) > -2.25) {
            v34 := v2[85]
            fmt.Println(v34, v34)
            var v35 uint64 = 4294967295
            _ = v34
            _ = v35
        }
    }
    var v36 uint64 = 280
    var v37 int64 = -35
    if true {
        for fi38 := 0; fi38 < 5; fi38++ {
            switch fi38 {
            case 464:
                fmt.Println(fi38, v1, fi38)
            case 9223372036854775807:
                delete(v2, 65535)
            default:
                var v39 bool = false
                _ = v39
            }
            var v40 int16 = 0
            if false {
                for fi41 := 0; fi41 < 18; fi41++ {
                    v42 := GenS1{F0: v1, F1: (!(v1 < 1))}
                    _ = fi41
                    _ = v42
                }
                for fi43 := 0; fi43 < 11; fi43++ {
                    switch (fi43 / (((fi43 * fi38)) | 1)) {
                    case 1:
                        if (!(((v40 % (((v40 ^ (^(v40)))) | 1)) >= 559) && true)) {
                            var v44 uint32 = 722
                            _ = v44
                        }
                    }
                    delete(v2, 2)
                    v2[42] = "a"
                    var v45 float32 = ((float32((float64(1e10) / v3)) + 3.14159) * float32(1e10))
                    v36 += v36
                    _ = fi43
                    _ = v45
                }
            } else {
                for fi46 := 0; fi46 < 7; fi46++ {
                    v47 := GenS2{F0: 42, F1: "a"}
                    _ = fi46
                    _ = v47
                }
                v37 = -9223372036854775808
            }
            _ = fi38
            _ = v40
        }
        var v48 bool = (-2147483648 > 42)
        v1 -= (v1 << ((255) & 63))
        _ = v48
    }
    if ("foo bar" < "\n\t") {
        switch "a" {
        case "foo bar":
            for fi49 := 0; fi49 < 20; fi49++ {
                v50 := GenS6{F0: 2}
                var v51 uint32 = 859
                _ = fi49
                _ = v50
                _ = v51
            }
        case "hello":
            fmt.Println(v36, v1, v36)
        }
    } else {
        v1 = (v1 << ((1) & 63))
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v12
    _ = v2
    _ = v23
    _ = v28
    _ = v3
    _ = v36
    _ = v37
}
