package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 uint8; F1 int32; F2 bool }
type GenS3 struct { F0 uint8 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 255
    v3 := []int{32768, 9223372036854775807, 128}
    var v4 uint8 = 153
    v5 := make(map[string]uint32)
    var v6 int64 = 9223372036854775807
    v2 += v2
    v6 = 0
    var v7 uint32 = 65535
    if (float32((float64(float32(-2.25)) * 1.5)) != -2.25) {
        fmt.Println(v2, v2, v6)
        v8 := GenS1{F0: ((v6 > (^(v6))) || true), F1: 720}
        _ = v8
    } else {
        var v9 float64 = ((1e10 * 1e10) + (3.14159 * 0.1))
        v10 := GenS1{F0: true, F1: -303}
        switch (("" + "a") + "a") {
        case "foo bar":
            for fi11 := 0; fi11 < 9; fi11++ {
                v12 := v5[""]
                v13 := GenS2{F0: 127, F1: 42, F2: false}
                var v14 uint16 = 0
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
            }
        default:
            v6 -= 941
        }
        _ = v10
        _ = v9
    }
    v2 -= v4
    var v15 int32 = 1
    for fi16 := 0; fi16 < 4; fi16++ {
        v17 := GenS1{F0: (!true), F1: v15}
        v2 -= v4
        _ = fi16
        _ = v17
    }
    v6 = v6
    if (274 >= 833) {
        var v18 int8 = -128
        _ = v18
    } else {
        v3 = append(v3, 532)
        switch 1 {
        case -963:
            fmt.Println(v2, v4)
        case 127:
            if ((!true) || true) {
                if len(v3) > 0 {
                    v3[0] = 4294967295
                }
            } else {
                v19 := GenS1{F0: false, F1: (v15 % ((v15) | 1))}
                _ = v19
            }
        case -726:
            v15 = (v15 % ((v15) | 1))
        default:
            v6 = (v6 + 9223372036854775807)
        }
    }
    v20 := v5["hello"]
    var v21 float64 = -2.25
    v7 = v20
    var v22 string = ""
    v23 := GenS3{F0: 0}
    switch -370 {
    case 2147483647:
        if (true && false) {
            v22 = v22
            if (true || false) {
                switch -2147483648 {
                case 109:
                    v24 := v1[(v20 / ((4294967295) | 1))]
                    _ = v24
                default:
                    v20 = 2147483647
                }
                var v25 int32 = 328
                if true {
                    switch (("hello" + ((v22 + v22) + v22)) + v22) {
                    case "":
                        if true {
                            v20 -= 177
                            var v26 int16 = 42
                            _ = v26
                        }
                    case "hello":
                        v15 -= -198
                    }
                    v2 += (^(v2))
                } else {
                    if false {
                        if false {
                            var v27 float64 = ((float64((float32(v21) + ((-2.25 + float32((v21 * v21))) / 1.0))) - (v21 + (v21 + v21))) / 3.14159)
                            _ = v27
                        }
                    } else {
                        var v28 float32 = -2.25
                        for fi29 := 0; fi29 < 10; fi29++ {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            var v30 uint8 = 0
                            _ = fi29
                            _ = v30
                        }
                        v25 += v25
                        _ = v28
                    }
                    if false {
                        for fi31 := 0; fi31 < 20; fi31++ {
                            v21 = float64(1e10)
                            _ = fi31
                        }
                    } else {
                        v7 += 65535
                        var v32 bool = (((^(v4)) >= v2) || false)
                        _ = v32
                    }
                }
                _ = v25
            } else {
                for fi33 := 0; fi33 < 2; fi33++ {
                    _ = len(v1)
                    fmt.Println(fi33)
                    _ = fi33
                }
            }
        } else {
            var v34 uint = 514
            _ = v34
        }
    case 1:
        v35 := v1[(v7 >> ((299) & 63))]
        _ = v35
    case 4294967295:
        v36 := GenS3{F0: 1}
        _ = v36
    default:
        for fi37 := 0; fi37 < 8; fi37++ {
            if false {
                v6 -= 9223372036854775807
                v6 = -1
            } else {
                for fi38 := 0; fi38 < 19; fi38++ {
                    delete(v1, v20)
                    _ = fi38
                }
            }
            _ = fi37
        }
    }
    v6 += (v6 >> ((611) & 63))
    v4 = v4
    for fi39 := 0; fi39 < 3; fi39++ {
        fmt.Println(v21, v20, v22)
        _ = fi39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
