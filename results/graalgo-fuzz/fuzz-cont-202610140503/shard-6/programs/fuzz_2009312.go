package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint16 }
type GenS2 struct { F0 int; F1 int; F2 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int32 }

func main() {
    var v1 uint8 = 42
    var v2 int = 4294967295
    var v3 uint64 = 42
    var v4 uint16 = 255
    var v5 int16 = -32768
    if (-2.25 == float32((float64(0.1) * ((((float64(-2.25) - 0.1) + ((float64((float32(float64((1.5 - 1.5))) / 1e10)) / 1e10) * 1.5)) + 0.1) + 0.0)))) {
        if false {
            for fi6 := 0; fi6 < 17; fi6++ {
                if false {
                    v3 = (v3 - (v3 + (v3 >> ((2) & 63))))
                } else {
                    fmt.Println(v5, v1, v3)
                    v7 := []int{(^(fi6)), 0}
                    _ = v7
                }
                _ = fi6
            }
        }
        var v8 float32 = (((1.5 + ((1e10 + (float32(0.0) * ((((1.5 * float32((1e10 + 1.5))) + float32((1e10 * 1.5))) + 3.14159) - float32(1.5)))) - 0.1)) / 1e10) - 1e10)
        v5 = (v5 & 349)
        _ = v8
    } else {
        v9 := GenS1{F0: true, F1: "a", F2: v4}
        v2 -= (v2 >> ((237) & 63))
        v1 = (v1 ^ v1)
        _ = v9
    }
    v2 += v2
    var v10 uint8 = v1
    v11 := GenS2{F0: v2, F1: 948, F2: (v3 / (((v3 + v3)) | 1))}
    fmt.Println(v2, v1, v5)
    v12 := []int{}
    v10 = (v1 * (^(v10)))
    for fi13 := 0; fi13 < 16; fi13++ {
        for fi14 := 0; fi14 < 15; fi14++ {
            if true {
                v3 = 18446744073709551615
                v15 := GenS1{F0: true, F1: ("a" + ("a" + ("" + "\n\t"))), F2: v4}
                _ = v15
            }
            _ = fi14
        }
        _ = fi13
    }
    v2 += v2
    _ = len(v12)
    var v16 uint = 255
    v4 -= 65535
    switch 0 {
    case 844:
        var v17 float64 = 3.14159
        _ = v17
    case -43:
        for fi18 := 0; fi18 < 18; fi18++ {
            fmt.Println(v5)
            var v19 int = -1
            v20 := GenS2{F0: v2, F1: 334, F2: (v3 * 1)}
            _ = fi18
            _ = v19
            _ = v20
        }
    case 384:
        v16 = (v16 << ((v16) & 63))
    default:
        var v21 int64 = -610
        _ = v21
    }
    if ((v10 & v10) > (v10 / ((128) | 1))) {
        var v22 int32 = 1
        v23 := GenS2{F0: v2, F1: v2, F2: v3}
        v16 = (v16 / (((v16 / ((v16) | 1))) | 1))
        _ = v22
        _ = v23
    }
    var v24 int16 = (v5 ^ v5)
    for fi25 := 0; fi25 < 18; fi25++ {
        if (true && true) {
            if (((v1 * (v1 | v10)) > (v10 % ((v1) | 1))) && (42 == 42)) {
                var v26 int16 = -32768
                switch "hello" {
                case "foo bar":
                    if false {
                        fmt.Println(v2)
                    }
                case "a":
                    v26 -= (v26 / ((v24) | 1))
                }
                switch fi25 {
                case 847:
                    v26 += v24
                case 32767:
                    v24 -= (-(v24))
                case 2147483647:
                    if ((^(v26)) < (v26 | 42)) {
                        v27 := GenS2{F0: fi25, F1: 9223372036854775807, F2: (v3 & v3)}
                        fmt.Println(v24, v1)
                        if false {
                            for fi28 := 0; fi28 < 20; fi28++ {
                                if false {
                                    var v29 float32 = 0.0
                                    var v30 int32 = -1
                                    switch (fi28 ^ (fi28 / ((-770) | 1))) {
                                    case 0:
                                        var v31 int32 = -2147483648
                                        _ = v31
                                    case -557:
                                        if true {
                                            if true {
                                                v32 := GenS3{F0: false}
                                                var v33 bool = (!true)
                                                _ = v32
                                                _ = v33
                                            }
                                        } else {
                                            if len(v12) >= 2 {
                                                v12 = v12[1:len(v12)]
                                            }
                                            v16 += (v16 * (^(v16)))
                                        }
                                    default:
                                        if len(v12) > 0 {
                                            v12[0] = (v2 << ((449) & 63))
                                        }
                                    }
                                    _ = v29
                                    _ = v30
                                }
                                v34 := GenS1{F0: (v2 != 9223372036854775807), F1: "\n\t", F2: (v4 ^ (v4 * v4))}
                                _ = fi28
                                _ = v34
                            }
                        }
                        _ = v27
                    }
                }
                _ = v26
            }
        } else {
            v12 = append(v12, (-(fi25)))
            v35 := GenS4{F0: -1}
            v36 := GenS3{F0: (4294967295 < 42)}
            _ = v35
            _ = v36
        }
        v2 = fi25
        _ = fi25
    }
    for fi37 := 0; fi37 < 15; fi37++ {
        fmt.Println(v10)
        _ = fi37
    }
    v4 = (v4 % (((v4 << ((v16) & 63))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
