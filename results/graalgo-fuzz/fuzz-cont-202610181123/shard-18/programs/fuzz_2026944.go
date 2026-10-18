package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 int8 }
func genFn0(p0 uint8, p1 int8) int16 {
    fmt.Println(p0, p0)
    p0 = p0
    _ = p0
    _ = p1
    return 1000
}

func main() {
    var v1 int32 = -760
    var v2 int32 = (v1 / (((-(v1))) | 1))
    v3 := []int{314}
    var v4 float64 = 1e10
    v5 := []int{}
    var v6 int32 = -119
    var v7 uint64 = 616
    if ((v7 - v7) >= v7) {
        for fi8 := 0; fi8 < 10; fi8++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            fmt.Println(v4, fi8)
            _ = fi8
        }
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        if ((754 < 0) && true) {
            var v10 int64 = 0
            v11 := GenS1{F0: 42, F1: 225}
            _ = v10
            _ = v11
        } else {
            if true {
                v12 := GenS1{F0: 0, F1: 42}
                switch "" {
                case "":
                    v6 += (-(v1))
                }
                _ = v12
            }
            v7 = (v7 & 399)
        }
        for fi13 := 0; fi13 < 6; fi13++ {
            v4 = v4
            v14 := GenS1{F0: -587, F1: 20}
            _ = fi13
            _ = v14
        }
        for fi15 := 0; fi15 < 14; fi15++ {
            _ = len(v5)
            v1 -= v2
            _ = len(v3)
            _ = fi15
        }
        _ = fi9
    }
    var v16 bool = false
    if v16 {
        for fi17 := 0; fi17 < 10; fi17++ {
            for fi18 := 0; fi18 < 7; fi18++ {
                v19 := GenS1{F0: -634, F1: 42}
                _ = fi18
                _ = v19
            }
            var v20 int64 = 348
            if true {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v21 := GenS2{F0: v16}
                v4 = (((v4 * v4) * 1e10) / (-2.25 - (1.5 / 1.5)))
                _ = v21
            } else {
                if ("a" <= "hello") {
                    for fi22 := 0; fi22 < 14; fi22++ {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = fi22
                    }
                    if v16 {
                        var v23 int64 = v20
                        _ = v23
                    } else {
                        v1 = (v1 * v2)
                        for fi24 := 0; fi24 < 18; fi24++ {
                            if len(v5) > 0 {
                                v5[0] = -1
                            }
                            if len(v5) > 0 {
                                v5[0] = (fi24 + fi24)
                            }
                            _ = fi24
                        }
                    }
                }
                for fi25 := 0; fi25 < 8; fi25++ {
                    if (!((422 > v7) && true)) {
                        v26 := GenS3{F0: 1, F1: v16}
                        _ = v26
                    } else {
                        v1 += (^(v1))
                        fmt.Println(v16, fi17, v4)
                    }
                    _ = fi25
                }
            }
            _ = fi17
            _ = v20
        }
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if v16 {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi27 := 0; fi27 < 13; fi27++ {
            v16 = (!(false || v16))
            _ = fi27
        }
    }
    v3 = append(v3, 9223372036854775807)
    _ = len(v3)
    for fi28 := 0; fi28 < 10; fi28++ {
        v29 := GenS4{F0: 0}
        _ = fi28
        _ = v29
    }
    for fi30 := 0; fi30 < 19; fi30++ {
        fmt.Println(v16)
        _ = fi30
    }
    for fi31 := 0; fi31 < 1; fi31++ {
        var v32 bool = v16
        if false {
            v1 -= v6
            var v33 uint8 = 118
            _ = v33
        } else {
            _ = len(v5)
        }
        _ = fi31
        _ = v32
    }
    if (float32(1e10) < 0.1) {
        var v34 uint = 65535
        _ = v34
    }
    var v35 float64 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
