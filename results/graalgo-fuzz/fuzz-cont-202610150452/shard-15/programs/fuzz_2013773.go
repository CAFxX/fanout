package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 int8; F2 uint8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int }

func main() {
    var v1 string = ""
    v2 := make(map[string]uint32)
    var v3 int = -1
    var v4 float32 = 0.0
    var v5 int8 = 100
    fmt.Println(v1)
    if (!false) {
        for fi6 := 0; fi6 < 16; fi6++ {
            fmt.Println(v5)
            _ = fi6
        }
    } else {
        if false {
            for fi7 := 0; fi7 < 20; fi7++ {
                v8 := []int{32767, v3, (fi7 << ((967) & 63))}
                v3 = (v3 & (-(fi7)))
                v5 -= (v5 + 100)
                _ = fi7
                _ = v8
            }
            var v9 uint = 0
            _ = v9
        }
        var v10 uint64 = 42
        v5 += (v5 % (((^(v5))) | 1))
        _ = v10
    }
    var v11 uint = 2
    var v12 float64 = 1e10
    for fi13 := 0; fi13 < 1; fi13++ {
        if false {
            v14 := GenS1{F0: -1}
            var v15 uint = (v11 >> (((v11 ^ 255)) & 63))
            _ = v14
            _ = v15
        } else {
            v16 := GenS2{F0: 18446744073709551615, F1: 120, F2: 0}
            v17 := []int{v3, v3}
            _ = v16
            _ = v17
        }
        switch (v1 + "\n\t") {
        case "foo bar":
            for fi18 := 0; fi18 < 11; fi18++ {
                v12 -= float64(v4)
                _ = fi18
            }
        case "hello":
            var v19 uint = 2
            _ = v19
        case "\n\t":
            var v20 int8 = 42
            _ = v20
        default:
            if true {
                v11 += v11
                if false {
                    for fi21 := 0; fi21 < 7; fi21++ {
                        v22 := GenS3{F0: (v3 << ((v11) & 63))}
                        var v23 int8 = v5
                        v24 := []int{}
                        switch "hello" {
                        case "a":
                            delete(v2, (v1 + "hello"))
                            if (2 >= 57) {
                                if len(v24) >= 2 {
                                    v24 = v24[1:len(v24)]
                                }
                            } else {
                                fmt.Println(fi13, v11, fi21)
                                var v25 float32 = (float32(float64((v4 - v4))) / v4)
                                var v26 uint = 65535
                                _ = v25
                                _ = v26
                            }
                        case "\n\t":
                            fmt.Println(v11, fi21)
                        case "":
                            if (true || true) {
                                v12 = (v12 + v12)
                            }
                        default:
                            v11 -= 1000
                        }
                        var v27 int16 = 0
                        var v28 int8 = 1
                        _ = fi21
                        _ = v22
                        _ = v23
                        _ = v24
                        _ = v27
                        _ = v28
                    }
                    fmt.Println(v1, v12)
                }
            } else {
                v11 -= 642
            }
        }
        _ = fi13
    }
    v4 += -2.25
    v29 := GenS3{F0: (v3 + v3)}
    v5 = 1
    for fi30 := 0; fi30 < 20; fi30++ {
        var v31 int = fi30
        for fi32 := 0; fi32 < 1; fi32++ {
            var v33 string = v1
            var v34 uint16 = 0
            switch (v3 ^ fi32) {
            case 2147483647:
                v2["\n\t"] = 97
            default:
                v35 := v2[v1]
                _ = v35
            }
            _ = fi32
            _ = v33
            _ = v34
        }
        fmt.Println(v1, v11, v5)
        _ = fi30
        _ = v31
    }
    v5 = (v5 | v5)
    v36 := []int{(v3 / ((v3) | 1)), (v3 % ((v3) | 1)), 129}
    for fi37 := 0; fi37 < 3; fi37++ {
        for fi38 := 0; fi38 < 9; fi38++ {
            if (!(42 < 4294967295)) {
                v11 -= (v11 / (((v11 | v11)) | 1))
            } else {
                _ = len(v36)
                if len(v36) >= 2 {
                    v36 = v36[1:len(v36)]
                }
            }
            _ = fi38
        }
        _ = fi37
    }
    v39 := GenS4{F0: (v3 + v3)}
    v12 = 1e10
    v5 = (v5 + (^(v5)))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v36
    _ = v39
    _ = v4
    _ = v5
}
