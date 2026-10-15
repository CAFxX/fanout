package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 int32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int32; F1 int32 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint32; F1 string }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]int)
    var v3 int32 = 42
    v4 := make(map[uint32]int)
    var v5 int16 = -1
    v6 := GenS1{F0: 592}
    delete(v2, 1)
    for fi7 := 0; fi7 < 3; fi7++ {
        if (18446744073709551615 > 65535) {
            v8 := GenS1{F0: (fi7 * 298)}
            _ = v8
        }
        _ = fi7
    }
    var v9 int64 = 0
    if false {
        v9 = 1
    } else {
        var v10 bool = ((290 != 0) || true)
        v9 = v9
        var v11 float64 = ((float64((-2.25 * -2.25)) + (float64(3.14159) * -2.25)) * 3.14159)
        _ = v10
        _ = v11
    }
    if ((2147483648 <= 2147483647) && false) {
        v12 := GenS2{F0: 42}
        if (true && false) {
            for fi13 := 0; fi13 < 15; fi13++ {
                if true {
                    if false {
                        var v14 float64 = (0.0 / 3.14159)
                        v15 := GenS3{F0: (v3 ^ (v3 << ((358) & 63))), F1: (^(v3))}
                        var v16 uint64 = 2
                        _ = v14
                        _ = v15
                        _ = v16
                    } else {
                        for fi17 := 0; fi17 < 4; fi17++ {
                            var v18 string = "a"
                            v18 = "hello"
                            var v19 uint32 = 436
                            fmt.Println(v18)
                            v20 := GenS4{F0: int16(v9)}
                            v21 := GenS1{F0: fi17}
                            _ = fi17
                            _ = v18
                            _ = v19
                            _ = v20
                            _ = v21
                        }
                        delete(v2, 4294967295)
                    }
                    for fi22 := 0; fi22 < 14; fi22++ {
                        switch "hello" {
                        case "":
                            var v23 int16 = (v5 ^ 42)
                            _ = v23
                        }
                        fmt.Println(fi22)
                        _ = fi22
                    }
                } else {
                    v24 := GenS1{F0: fi13}
                    var v25 uint64 = 1
                    var v26 bool = false
                    _ = v24
                    _ = v25
                    _ = v26
                }
                fmt.Println(v9)
                v9 = v9
                _ = fi13
            }
            v5 += v5
        }
        _ = v12
    } else {
        var v27 float32 = (0.0 + float32(0.0))
        if true {
            var v28 uint8 = 247
            _ = v28
        } else {
            fmt.Println(v3, v5, v9)
        }
        _ = v27
    }
    if (false || true) {
        var v29 uint16 = 1
        v3 = (v3 & v3)
        switch "foo bar" {
        case "":
            for fi30 := 0; fi30 < 5; fi30++ {
                fmt.Println(v5, v5, fi30)
                v31 := GenS2{F0: v29}
                _ = fi30
                _ = v31
            }
        case "a":
            fmt.Println(v9, v3, v29)
        default:
            if (32767 == 32767) {
                if true {
                    v9 = v9
                    if false {
                        for fi32 := 0; fi32 < 10; fi32++ {
                            var v33 uint32 = 0
                            _ = fi32
                            _ = v33
                        }
                    }
                } else {
                    if ((((true || (true && (0 <= 1))) && false) || false) && (18446744073709551615 > 42)) {
                        v34 := GenS5{F0: (v3 | 17), F1: v3}
                        v9 += (v9 * v9)
                        _ = v34
                    }
                    switch -1 {
                    case 561:
                        for fi35 := 0; fi35 < 12; fi35++ {
                            v36 := []int{fi35, fi35}
                            _ = fi35
                            _ = v36
                        }
                    }
                    v5 += v5
                }
                fmt.Println(v3)
                v5 += (v5 << ((502) & 63))
            }
        }
        _ = v29
    } else {
        v9 = 9223372036854775807
        if false {
            var v37 uint32 = 42
            v38 := []int{128, 127, 0}
            v39 := GenS6{F0: (-84 == 4294967295)}
            _ = v37
            _ = v38
            _ = v39
        } else {
            switch 42 {
            case -2147483648:
                if false {
                    v40 := GenS2{F0: 1}
                    switch "foo bar" {
                    case "a":
                        v3 = (-(v3))
                    }
                    _ = v40
                } else {
                    var v41 int = 878
                    v3 = v3
                    var v42 float32 = (1e10 / ((1.5 - -2.25) - 0.1))
                    _ = v41
                    _ = v42
                }
            case 2147483647:
                v5 += v5
            case 9223372036854775807:
                var v43 uint16 = 0
                _ = v43
            default:
                v9 += v9
            }
            var v44 int8 = 1
            _ = v44
        }
        v45 := GenS7{F0: 2, F1: "foo bar"}
        _ = v45
    }
    switch ("" + ("foo bar" + "a")) {
    case "hello":
        v9 = v9
    default:
        v3 = (v3 ^ (v3 % ((v3) | 1)))
    }
    v46 := []int{-772, 32768}
    for fi47 := 0; fi47 < 12; fi47++ {
        fmt.Println(v3, fi47)
        var v48 int = fi47
        _ = fi47
        _ = v48
    }
    fmt.Println(v5)
    v5 = (v5 / ((v5) | 1))
    _ = len(v2)
    var v49 int16 = (v5 << ((1) & 63))
    v5 = v5
    v3 = (v3 % (((v3 / ((238) | 1))) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v49
    _ = v5
    _ = v6
    _ = v9
}
