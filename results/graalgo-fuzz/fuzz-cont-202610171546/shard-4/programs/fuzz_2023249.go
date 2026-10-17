package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 bool }
type GenS2 struct { F0 uint16; F1 uint16; F2 int64 }
type GenS3 struct { F0 int; F1 int; F2 uint64 }
type GenS4 struct { F0 uint; F1 int8; F2 int16 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]string)
    var v3 uint16 = 32767
    var v4 int8 = 19
    v5 := make(map[string]int)
    var v6 uint64 = 0
    var v7 uint64 = (v6 >> ((286) & 63))
    if (!(false && false)) {
        _ = len(v2)
        var v8 uint = 0
        v3 -= 686
        _ = v8
    }
    v6 = (v6 % ((v6) | 1))
    if (-2.25 < (3.14159 + (-2.25 * 1.5))) {
        if false {
            if (false || true) {
                if (!((false && (30 >= (^(v4)))) && false)) {
                    v3 = (v3 / ((uint16(v6)) | 1))
                    for fi9 := 0; fi9 < 20; fi9++ {
                        var v10 uint16 = v3
                        v4 -= v4
                        _ = fi9
                        _ = v10
                    }
                } else {
                    switch -2147483648 {
                    case 32768:
                        v11 := GenS1{F0: -497, F1: 2, F2: (false || (2 <= v3))}
                        _ = v11
                    case -141:
                        var v12 int = 128
                        _ = v12
                    default:
                        v7 += 2
                    }
                    switch "foo bar" {
                    case "a":
                        v3 = v3
                    }
                    var v13 uint64 = (v6 << ((18446744073709551615) & 63))
                    _ = v13
                }
            } else {
                var v14 float32 = 0.1
                if (v3 < (^(v3))) {
                    switch "a" {
                    case "hello":
                        if true {
                            v3 += 369
                            v15 := GenS2{F0: (^(v3)), F1: v3, F2: 211}
                            v14 = ((0.1 + float32(float64((float32(((float64(((v14 - (float32(0.0) * 3.14159)) / ((v14 / v14) * v14))) + 0.1) - 1.5)) + v14)))) + 3.14159)
                            _ = v15
                        }
                    case "foo bar":
                        var v16 int16 = 32767
                        _ = v16
                    }
                }
                _ = v14
            }
        } else {
            var v17 float64 = 0.0
            for fi18 := 0; fi18 < 15; fi18++ {
                v17 = (((float64((1.5 + 1.5)) + (v17 / -2.25)) / v17) + v17)
                v19 := []int{(fi18 / ((32768) | 1)), (fi18 << ((1) & 63))}
                _ = fi18
                _ = v19
            }
            _ = v17
        }
        var v20 int16 = 32767
        _ = v20
    }
    switch 2147483647 {
    case -2147483648:
        var v21 bool = false
        _ = v21
    case 0:
        v7 = v6
    case 223:
        delete(v1, 32768)
    }
    for fi22 := 0; fi22 < 19; fi22++ {
        v6 = uint64(v3)
        switch (("hello" + "hello") + "foo bar") {
        case "hello":
            v6 = v6
        }
        fmt.Println(fi22, v7)
        _ = fi22
    }
    for fi23 := 0; fi23 < 20; fi23++ {
        var v24 uint16 = (v3 - 255)
        v3 += (^(v3))
        for fi25 := 0; fi25 < 4; fi25++ {
            v4 = v4
            v26 := GenS3{F0: fi23, F1: (fi25 << ((18446744073709551615) & 63)), F2: v7}
            _ = fi25
            _ = v26
        }
        _ = fi23
        _ = v24
    }
    var v27 uint8 = 1
    v6 = 0
    var v28 int = -9223372036854775808
    v29 := []int{}
    for fi30 := 0; fi30 < 4; fi30++ {
        for fi31 := 0; fi31 < 6; fi31++ {
            fmt.Println(v6)
            fmt.Println(fi30)
            _ = fi31
        }
        v32 := GenS1{F0: (^(fi30)), F1: (v27 & (^(v27))), F2: false}
        _ = fi30
        _ = v32
    }
    if true {
        if len(v29) >= 2 {
            v29 = v29[1:len(v29)]
        }
        for fi33 := 0; fi33 < 9; fi33++ {
            if false {
                _ = len(v29)
            } else {
                v34 := GenS1{F0: (fi33 ^ v28), F1: (v27 | (v27 >> ((1) & 63))), F2: true}
                if (false && ((v3 << ((232) & 63)) <= v3)) {
                    for fi35 := 0; fi35 < 3; fi35++ {
                        switch -9223372036854775808 {
                        case -104:
                            v36 := v1[-9223372036854775808]
                            _ = v36
                        case 0:
                            v3 = v3
                        }
                        _ = fi35
                    }
                    v37 := GenS4{F0: 532, F1: (v4 - (v4 << ((420) & 63))), F2: 491}
                    if false {
                        v4 = v4
                        var v38 float32 = float32(1e10)
                        if len(v29) >= 2 {
                            v29 = v29[1:len(v29)]
                        }
                        _ = v38
                    }
                    _ = v37
                }
                var v39 int32 = 815
                _ = v34
                _ = v39
            }
            if len(v29) > 0 {
                v29[0] = (v28 | v28)
            }
            _ = fi33
        }
        _ = len(v29)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
