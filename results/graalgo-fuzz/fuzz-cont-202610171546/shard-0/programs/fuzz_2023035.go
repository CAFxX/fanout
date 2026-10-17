package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 int16 }
func genFn0(p0 int32, p1 uint8, p2 uint32) uint32 {
    for fi5 := 0; fi5 < 3; fi5++ {
        p0 = p0
        if true {
            fmt.Println(fi5)
            for fi6 := 0; fi6 < 16; fi6++ {
                var v7 float32 = 1e10
                _ = fi6
                _ = v7
            }
        }
        p0 = 2147483647
        _ = fi5
    }
    var v8 uint16 = 2
    for fi9 := 0; fi9 < 3; fi9++ {
        var v10 bool = true
        p1 += (p1 / (((^(p1))) | 1))
        v11 := []int{(-(fi9))}
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 int16 = 32767
    v13 := []int{2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    _ = v8
    return p2
}
func genFn1(p0 uint16) uint32 {
    v14 := GenS1{F0: 0, F1: 894}
    p0 -= p0
    v15 := []int{}
    p0 -= (^(p0))
    _ = p0
    _ = v14
    _ = v15
    return genFn0(-2147483648, 0, 2147483648)
}

func main() {
    v1 := []int{}
    var v2 int = -943
    var v3 int64 = 42
    var v4 uint64 = 2
    var v16 uint8 = uint8(v4)
    if (((0.0 - -2.25) - (3.14159 * 3.14159)) == 0.0) {
        if true {
            for fi17 := 0; fi17 < 5; fi17++ {
                v2 += (^(fi17))
                fmt.Println(v4, v2)
                fmt.Println(v3, v4, fi17)
                _ = fi17
            }
            v16 = v16
            if (-9223372036854775808 == 0) {
                if len(v1) > 0 {
                    v1[0] = -790
                }
            } else {
                switch v2 {
                case -843:
                    v2 = 89
                default:
                    fmt.Println(v16)
                }
                if true {
                    if ((!false) || (2 < 1)) {
                        _ = len(v1)
                        var v18 int8 = 9
                        _ = v18
                    }
                    for fi19 := 0; fi19 < 7; fi19++ {
                        if false {
                            _ = len(v1)
                            var v20 uint64 = v4
                            _ = v20
                        }
                        if false {
                            var v21 float64 = 1.5
                            v2 = fi19
                            _ = v21
                        } else {
                            var v22 float64 = 0.1
                            v23 := GenS1{F0: 32767, F1: v4}
                            v4 -= (^(v4))
                            _ = v22
                            _ = v23
                        }
                        v2 = fi19
                        _ = fi19
                    }
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v24 uint = 2
                    v24 = (v24 * 2)
                    _ = v24
                }
            }
        } else {
            for fi25 := 0; fi25 < 11; fi25++ {
                var v26 int16 = 280
                _ = fi25
                _ = v26
            }
        }
        switch "hello" {
        case "\n\t":
            v27 := GenS1{F0: -32768, F1: v4}
            _ = v27
        case "foo bar":
            var v28 uint8 = 2
            _ = v28
        }
        var v29 float32 = ((0.0 - float32(1e10)) + ((float32(float64(float32(1.5))) / 1.0) * (-2.25 - 0.0)))
        _ = v29
    } else {
        _ = len(v1)
        if true {
            for fi30 := 0; fi30 < 4; fi30++ {
                fmt.Println(v2, v3)
                var v31 int8 = 28
                v3 = (v3 + (v3 & int64(fi30)))
                _ = fi30
                _ = v31
            }
            v2 = (v2 ^ v2)
            switch (v2 & v2) {
            case -9223372036854775808:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case 1:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case 2147483647:
                v2 = 390
            }
        } else {
            var v32 uint64 = v4
            _ = v32
        }
    }
    v33 := GenS1{F0: -306, F1: v4}
    v34 := GenS2{F0: -424}
    v2 = v2
    var v35 uint32 = 289
    v16 = v16
    switch 2147483647 {
    case 1:
        var v36 uint32 = genFn0(0, 87, 2147483647)
        _ = v36
    default:
        switch "\n\t" {
        case "":
            switch (v2 * v2) {
            case -2147483648:
                v1 = append(v1, v2)
            default:
                for fi37 := 0; fi37 < 1; fi37++ {
                    switch (fi37 * (v2 + v2)) {
                    case -1:
                        v38 := GenS2{F0: 1}
                        _ = v38
                    case 42:
                        v4 += (^(v4))
                    case 2147483647:
                        var v39 int64 = v3
                        _ = v39
                    }
                    _ = fi37
                }
            }
        }
    }
    var v40 int64 = 895
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v40
}
