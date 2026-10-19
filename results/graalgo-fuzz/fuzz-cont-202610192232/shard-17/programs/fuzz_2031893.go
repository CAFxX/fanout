package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint; F1 uint16 }
type GenS5 struct { F0 bool }
func genFn0() int16 {
    v4 := []int{128, 128}
    v5 := GenS1{F0: 1, F1: 406}
    switch -997 {
    case 9223372036854775807:
        for fi6 := 0; fi6 < 9; fi6++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v7 := GenS2{F0: 255}
            var v8 float64 = (3.14159 + 0.0)
            _ = fi6
            _ = v7
            _ = v8
        }
    default:
        _ = len(v4)
    }
    if ((!(0.1 >= (-2.25 / (1e10 - float32(0.1))))) && false) {
        for fi9 := 0; fi9 < 20; fi9++ {
            _ = len(v4)
            _ = fi9
        }
        for fi10 := 0; fi10 < 10; fi10++ {
            var v11 string = ("foo bar" + "\n\t")
            _ = fi10
            _ = v11
        }
        for fi12 := 0; fi12 < 3; fi12++ {
            v13 := GenS2{F0: 619}
            if (-128 == 127) {
                var v14 uint64 = 2
                for fi15 := 0; fi15 < 15; fi15++ {
                    v14 = v14
                    if ((v14 ^ (^(v14))) < 811) {
                        var v16 int16 = -1
                        v14 = 42
                        v17 := GenS3{F0: 575}
                        _ = v16
                        _ = v17
                    } else {
                        var v18 int16 = 1
                        fmt.Println(fi12, fi12)
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        var v19 int32 = 396
                        v4 = append(v4, 37)
                        _ = v18
                        _ = v19
                    }
                    _ = fi15
                }
                fmt.Println(fi12)
                _ = v14
            }
            for fi20 := 0; fi20 < 18; fi20++ {
                if false {
                    fmt.Println(fi20)
                    fmt.Println(fi12, fi20)
                    var v21 string = "hello"
                    _ = v21
                }
                var v22 int64 = -659
                v22 = v22
                _ = fi20
                _ = v22
            }
            _ = fi12
            _ = v13
        }
    }
    _ = v4
    _ = v5
    return -558
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]bool)
    v3 := make(map[string]uint32)
    var v23 float32 = (0.0 / 0.1)
    var v24 float32 = (v23 / v23)
    var v25 uint32 = 2147483648
    for fi26 := 0; fi26 < 13; fi26++ {
        v23 += 1.5
        _ = fi26
    }
    delete(v1, 776)
    for fi27 := 0; fi27 < 12; fi27++ {
        var v28 uint64 = 18446744073709551615
        _ = fi27
        _ = v28
    }
    var v29 int64 = 9223372036854775807
    for fi30 := 0; fi30 < 10; fi30++ {
        fmt.Println(v23, v24)
        for fi31 := 0; fi31 < 7; fi31++ {
            var v32 int = (fi31 >> ((673) & 63))
            v33 := []int{(fi30 ^ v32), fi30, fi30}
            _ = fi31
            _ = v32
            _ = v33
        }
        _ = fi30
    }
    var v34 string = ""
    v35 := GenS4{F0: uint(v25), F1: 812}
    var v36 int = 1
    v36 = v36
    if ((int8(v25) > 13) && false) {
        v37 := GenS4{F0: 255, F1: 2}
        _ = v37
    } else {
        for fi38 := 0; fi38 < 14; fi38++ {
            fmt.Println(v24)
            _ = fi38
        }
        if true {
            v39 := GenS5{F0: ((!(0 <= v36)) || (502 <= (v25 & 2147483647)))}
            for fi40 := 0; fi40 < 8; fi40++ {
                fmt.Println(v34, v23, v29)
                v41 := GenS1{F0: -1, F1: 980}
                _ = fi40
                _ = v41
            }
            _ = v39
        }
        switch 322 {
        case -274:
            switch v34 {
            case "foo bar":
                for fi42 := 0; fi42 < 11; fi42++ {
                    fmt.Println(v36, v25)
                    _ = fi42
                }
            default:
                var v43 float64 = 0.1
                _ = v43
            }
        case 1:
            switch v34 {
            case "":
                v24 = v23
            }
        default:
            v25 = v25
        }
    }
    switch v36 {
    case 653:
        fmt.Println(v36, v36, v36)
    case 9223372036854775807:
        v25 = 2147483647
    case 127:
        v24 = v23
    default:
        v25 = (v25 >> ((42) & 63))
    }
    var v44 int32 = -457
    v45 := []int{(^(v36))}
    fmt.Println(v23)
    v24 = v24
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v45:", len(v45), v45)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v44
    _ = v45
}
