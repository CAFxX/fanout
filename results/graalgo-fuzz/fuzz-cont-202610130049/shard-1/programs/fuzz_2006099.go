package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 string; F2 uint32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32; F1 uint8; F2 uint }
func genFn0(p0 uint32, p1 uint16) bool {
    p0 -= p0
    for fi7 := 0; fi7 < 16; fi7++ {
        v8 := []int{fi7}
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        fmt.Println(fi9, p0)
        var v10 int = fi9
        _ = fi9
        _ = v10
    }
    var v11 uint32 = p0
    _ = p0
    _ = p1
    _ = v11
    return true
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int64 = -9223372036854775808
    var v3 uint64 = 710
    var v4 int8 = 127
    v5 := make(map[uint32]uint32)
    v6 := []string{}
    for fi12 := 0; fi12 < 13; fi12++ {
        delete(v5, 805)
        var v13 int64 = (v2 << ((0) & 63))
        v2 -= v2
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        if (450 != 530) {
            for fi15 := 0; fi15 < 12; fi15++ {
                var v16 int8 = v4
                var v17 int32 = int32(fi15)
                v18 := GenS1{F0: v17}
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
            }
            _ = len(v5)
        }
        if (!true) {
            var v19 float64 = float64((float32(-2.25) * 0.0))
            _ = v19
        }
        v3 += (v3 * v3)
        _ = fi14
    }
    if true {
        if (-2147483648 > 1) {
            v4 -= 1
            v20 := GenS2{F0: 841, F1: ("foo bar" + "\n\t"), F2: 65535}
            _ = v20
        } else {
            v21 := GenS1{F0: 42}
            var v22 uint64 = (^(v3))
            _ = v21
            _ = v22
        }
        fmt.Println(v4, v4, v3)
        v23 := GenS2{F0: 4294967295, F1: ("foo bar" + "hello"), F2: 2147483648}
        _ = v23
    } else {
        if (918 <= 385) {
            v3 = v3
            var v24 int8 = (v4 + (v4 % ((122) | 1)))
            _ = v24
        }
        fmt.Println(v2, v2, v3)
    }
    var v25 uint8 = 128
    if false {
        v26 := GenS1{F0: 698}
        switch "\n\t" {
        case "\n\t":
            var v27 bool = false
            _ = v27
        case "a":
            v2 -= v2
        case "foo bar":
            for fi28 := 0; fi28 < 18; fi28++ {
                var v29 int32 = 0
                for fi30 := 0; fi30 < 20; fi30++ {
                    fmt.Println(v2, v2)
                    if (!(fi30 <= fi28)) {
                        v31 := GenS3{F0: (v4 >> ((606) & 63))}
                        v6 = append(v6, ("a" + ("hello" + (("" + "hello") + "hello"))))
                        _ = v31
                    } else {
                        fmt.Println(v4, v3, v25)
                        var v32 uint = 4294967295
                        fmt.Println(v2, v29)
                        var v33 uint64 = v3
                        v29 = v29
                        var v34 string = "hello"
                        _ = len(v6)
                        _ = v32
                        _ = v33
                        _ = v34
                    }
                    if true {
                        v35 := GenS2{F0: 1, F1: "\n\t", F2: 0}
                        fmt.Println(fi28, fi28)
                        var v36 int64 = 1
                        _ = v35
                        _ = v36
                    }
                    _ = fi30
                }
                _ = fi28
                _ = v29
            }
        }
        var v37 int8 = (v4 % (((v4 - (v4 & (v4 + -54)))) | 1))
        _ = v26
        _ = v37
    }
    if true {
        switch -1 {
        case 9223372036854775807:
            if (v25 != 127) {
                if (v2 <= 1) {
                    v25 = (v25 / (((v25 ^ v25)) | 1))
                    v38 := GenS3{F0: (v4 - 26)}
                    _ = v38
                }
                _ = len(v5)
            } else {
                var v39 uint32 = 2147483648
                var v40 int8 = (v4 + (v4 * -101))
                _ = v39
                _ = v40
            }
        case 128:
            switch -1 {
            case -137:
                var v41 uint16 = 42
                _ = v41
            case 9223372036854775807:
                v42 := v5[65535]
                _ = v42
            case -346:
                delete(v1, 528)
            }
        default:
            _ = len(v6)
        }
        v4 = (v4 << ((0) & 63))
        if (196 < 4294967295) {
            delete(v5, 2147483647)
            for fi43 := 0; fi43 < 9; fi43++ {
                fmt.Println(v2)
                v1[2147483648] = 65535
                fmt.Println(v4, v25)
                _ = fi43
            }
        } else {
            if (-2.25 > (0.1 - 1.5)) {
                fmt.Println(v3, v25, v3)
                if len(v6) > 0 {
                    v6[0] = ("hello" + "\n\t")
                }
                var v44 int16 = -316
                _ = v44
            }
        }
    } else {
        v45 := GenS4{F0: -2147483648, F1: 43, F2: 285}
        v46 := GenS1{F0: -559}
        var v47 uint = 627
        _ = v45
        _ = v46
        _ = v47
    }
    v48 := GenS2{F0: 996, F1: "foo bar", F2: 65535}
    v2 = (v2 & (-(v2)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v6
}
