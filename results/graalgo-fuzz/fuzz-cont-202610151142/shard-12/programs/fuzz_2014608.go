package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32; F1 uint }
type GenS4 struct { F0 uint8; F1 int32; F2 int }
func genFn0() int32 {
    var v4 uint16 = 32767
    v5 := []int{}
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 uint32 = 863
        var v8 uint16 = v4
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 uint = uint(v4)
    v9 += v9
    _ = v4
    _ = v5
    _ = v9
    return 42
}
func genFn1(p0 string, p1 int16, p2 uint16) int32 {
    var v10 float64 = 1.5
    v11 := GenS1{F0: genFn0()}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return -177
}

func main() {
    var v1 int16 = -681
    v2 := make(map[string]int)
    var v3 int16 = (-(v1))
    if (!false) {
        v12 := []int{}
        v13 := GenS2{F0: 656}
        _ = v12
        _ = v13
    }
    switch 1 {
    case 0:
        v14 := []int{-852, -9223372036854775808, 0, 127}
        _ = v14
    default:
        switch 9223372036854775807 {
        case 42:
            if (true || true) {
                for fi15 := 0; fi15 < 12; fi15++ {
                    var v16 uint32 = 2147483647
                    switch "hello" {
                    case "\n\t":
                        switch "" {
                        case "hello":
                            if ((!true) && ((true || (false && (true && true))) || false)) {
                                v1 = v3
                            }
                        default:
                            for fi17 := 0; fi17 < 12; fi17++ {
                                v18 := []int{}
                                v1 -= 301
                                fmt.Println(fi15)
                                _ = fi17
                                _ = v18
                            }
                        }
                    default:
                        v19 := GenS2{F0: 887}
                        _ = v19
                    }
                    _ = fi15
                    _ = v16
                }
                delete(v2, "\n\t")
                var v20 float64 = 0.0
                _ = v20
            }
        case 955:
            fmt.Println(v1)
        case -104:
            var v21 uint16 = 65535
            _ = v21
        }
    }
    v22 := GenS1{F0: 866}
    for fi23 := 0; fi23 < 12; fi23++ {
        var v24 bool = false
        _ = fi23
        _ = v24
    }
    if (70 == 0) {
        for fi25 := 0; fi25 < 11; fi25++ {
            if true {
                var v26 int8 = 16
                fmt.Println(fi25, fi25)
                _ = v26
            } else {
                v1 = (v3 % ((119) | 1))
                if true {
                    fmt.Println(v1, v1, v3)
                    if (true && (!(646 <= 708))) {
                        for fi27 := 0; fi27 < 3; fi27++ {
                            v28 := GenS2{F0: 1}
                            v2["foo bar"] = (fi25 - 128)
                            v1 -= (-(v3))
                            _ = fi27
                            _ = v28
                        }
                    } else {
                        var v29 bool = ((!(4294967295 <= 4294967295)) || (float32((1e10 / 3.14159)) < float32(1.5)))
                        if ((592 != 0) && (!false)) {
                            v30 := []int{}
                            fmt.Println(v1, v29)
                            switch "" {
                            case "foo bar":
                                fmt.Println(v1, v1, v3)
                            case "":
                                var v31 int64 = -9223372036854775808
                                _ = v31
                            }
                            _ = v30
                        } else {
                            for fi32 := 0; fi32 < 13; fi32++ {
                                var v33 int16 = (^(v3))
                                _ = fi32
                                _ = v33
                            }
                        }
                        for fi34 := 0; fi34 < 3; fi34++ {
                            v29 = ((0 <= 685) || true)
                            var v35 int8 = 17
                            _ = fi34
                            _ = v35
                        }
                        _ = v29
                    }
                }
                v36 := GenS3{F0: 1, F1: 18446744073709551615}
                _ = v36
            }
            _ = fi25
        }
        v3 = (v3 << ((18446744073709551615) & 63))
        if (652 <= -932) {
            var v37 uint8 = 127
            if true {
                for fi38 := 0; fi38 < 7; fi38++ {
                    v1 = v3
                    _ = fi38
                }
            }
            fmt.Println(v37)
            _ = v37
        }
    } else {
        fmt.Println(v1, v3)
        var v39 uint64 = 687
        v40 := GenS1{F0: 1}
        _ = v39
        _ = v40
    }
    var v41 uint = 607
    switch 32768 {
    case 127:
        v1 = v1
    }
    if true {
        v42 := GenS4{F0: 128, F1: 42, F2: -265}
        for fi43 := 0; fi43 < 1; fi43++ {
            var v44 int8 = -40
            fmt.Println(v41, v3)
            _ = fi43
            _ = v44
        }
        _ = v42
    } else {
        if false {
            v3 = -711
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v41
}
