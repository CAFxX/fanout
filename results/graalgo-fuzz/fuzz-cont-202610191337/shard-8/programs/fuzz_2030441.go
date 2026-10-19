package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool; F1 uint32; F2 uint8 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint32 = 65535
    v3 := []int{-552, -948, 4294967295}
    var v4 int64 = 603
    v5 := GenS1{F0: false}
    for fi6 := 0; fi6 < 17; fi6++ {
        for fi7 := 0; fi7 < 11; fi7++ {
            switch "" {
            case "\n\t":
                v8 := GenS1{F0: false}
                _ = v8
            case "foo bar":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                switch ("a" + "hello") {
                case "foo bar":
                    v4 -= (v4 + (-(v4)))
                case "":
                    fmt.Println(v2)
                default:
                    var v9 uint16 = 2
                    _ = v9
                }
            case "a":
                var v10 int64 = -435
                _ = v10
            }
            if (v2 >= v2) {
                var v11 int16 = -62
                _ = v11
            } else {
                _ = len(v3)
            }
            v2 = v2
            _ = fi7
        }
        _ = fi6
    }
    var v12 int32 = -857
    _ = len(v3)
    var v13 float64 = 0.1
    switch "\n\t" {
    case "foo bar":
        if true {
            v14 := GenS2{F0: -1}
            v2 += 186
            _ = v14
        }
    case "a":
        fmt.Println(v12, v12)
    default:
        switch "a" {
        case "\n\t":
            _ = len(v3)
        case "hello":
            fmt.Println(v2, v13)
        case "a":
            v15 := GenS1{F0: (!(32767 != 42))}
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 7; fi16++ {
        if (-2.25 == 3.14159) {
            v17 := GenS3{F0: false, F1: (v2 | (v2 + v2)), F2: 1}
            _ = v17
        } else {
            switch 459 {
            case 914:
                v13 = 1.5
            case -668:
                for fi18 := 0; fi18 < 3; fi18++ {
                    var v19 uint32 = (v2 / ((4294967295) | 1))
                    switch "\n\t" {
                    case "hello":
                        if (42 != 2) {
                            v12 = v12
                            var v20 string = ("hello" + "foo bar")
                            if false {
                                fmt.Println(v19)
                                if (255 == 42) {
                                    var v21 uint16 = 32767
                                    _ = v21
                                }
                            } else {
                                var v22 uint = 18446744073709551615
                                _ = v22
                            }
                            _ = v20
                        }
                    case "":
                        v19 -= (v19 & (v19 / ((v2) | 1)))
                    }
                    _ = len(v3)
                    _ = fi18
                    _ = v19
                }
            case 32768:
                v23 := GenS1{F0: (false && false)}
                _ = v23
            default:
                var v24 uint16 = 32767
                _ = v24
            }
        }
        for fi25 := 0; fi25 < 9; fi25++ {
            v2 = (^(v2))
            if true {
                fmt.Println(fi16, fi25)
            }
            fmt.Println(v2, fi16)
            _ = fi25
        }
        _ = fi16
    }
    for fi26 := 0; fi26 < 1; fi26++ {
        fmt.Println(v4, v12, v2)
        _ = fi26
    }
    if len(v3) > 0 {
        v3[0] = 583
    }
    var v27 int = 32767
    var v28 int16 = -4
    var v29 float32 = (-2.25 + (1e10 / 1.5))
    if (0 != 1) {
        if true {
            v28 += (^(v28))
            var v30 string = ("" + "foo bar")
            _ = v30
        }
        v12 = int32(v28)
        v3 = append(v3, v27)
    } else {
        for fi31 := 0; fi31 < 17; fi31++ {
            v32 := v1[(^(v2))]
            _ = fi31
            _ = v32
        }
    }
    v12 -= 0
    _ = len(v1)
    if true {
        _ = len(v3)
        v12 = v12
    } else {
        for fi33 := 0; fi33 < 14; fi33++ {
            if true {
                if false {
                    var v34 int16 = 42
                    for fi35 := 0; fi35 < 7; fi35++ {
                        var v36 int16 = v28
                        _ = fi35
                        _ = v36
                    }
                    v37 := v1[(v2 | 4294967295)]
                    _ = v34
                    _ = v37
                }
            } else {
                for fi38 := 0; fi38 < 7; fi38++ {
                    v4 += v4
                    fmt.Println(v2, v28, v29)
                    _ = fi38
                }
                var v39 uint16 = 1
                _ = v39
            }
            var v40 uint8 = 255
            _ = fi33
            _ = v40
        }
        for fi41 := 0; fi41 < 20; fi41++ {
            var v42 int = -9223372036854775808
            _ = fi41
            _ = v42
        }
    }
    v28 = (v28 % ((168) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
