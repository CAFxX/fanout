package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 uint8 }
func genFn0() uint8 {
    var v4 uint16 = 255
    var v5 int16 = 102
    v4 += (v4 & v4)
    v5 = -32768
    _ = v4
    _ = v5
    return uint8(v4)
}
func genFn1(p0 uint8) uint32 {
    v6 := []int{2147483647}
    p0 += (p0 * p0)
    fmt.Println(p0)
    _ = len(v6)
    _ = p0
    _ = v6
    return 42
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]bool)
    v3 := make(map[int]uint32)
    var v7 uint32 = 0
    if false {
        fmt.Println(v7)
        fmt.Println(v7)
    } else {
        fmt.Println(v7)
    }
    var v8 uint16 = 151
    switch "a" {
    case "\n\t":
        var v9 int16 = 1000
        _ = v9
    default:
        for fi10 := 0; fi10 < 8; fi10++ {
            v11 := v1["hello"]
            _ = fi10
            _ = v11
        }
    }
    v12 := []int{-7, -160}
    if true {
        v8 += 433
    }
    v7 += (^(v7))
    v8 -= (v8 * 715)
    if true {
        if ((v8 % ((65535) | 1)) >= 1) {
            for fi13 := 0; fi13 < 16; fi13++ {
                v7 = 2147483648
                v8 += v8
                var v14 int16 = -175
                _ = fi13
                _ = v14
            }
            var v15 uint8 = uint8(v7)
            var v16 uint64 = 42
            _ = v15
            _ = v16
        } else {
            if true {
                v7 = v7
                for fi17 := 0; fi17 < 18; fi17++ {
                    v8 = 32767
                    if (v8 > v8) {
                        if (!false) {
                            fmt.Println(fi17, v7, fi17)
                            if false {
                                fmt.Println(fi17)
                                var v18 float64 = -2.25
                                _ = v18
                            }
                            var v19 int = (fi17 & -597)
                            _ = v19
                        } else {
                            v8 = (v8 >> ((uint(v7)) & 63))
                        }
                    }
                    v8 += 715
                    _ = fi17
                }
                v20 := GenS1{F0: 4294967295}
                _ = v20
            } else {
                if false {
                    var v21 uint = 42
                    v22 := GenS2{F0: -1, F1: 2147483647}
                    _ = v21
                    _ = v22
                }
            }
            v7 += genFn1(2)
            v8 = 1
        }
        if (((-9223372036854775808 > 0) || false) || ((v8 & v8) <= v8)) {
            for fi23 := 0; fi23 < 19; fi23++ {
                if (!false) {
                    for fi24 := 0; fi24 < 12; fi24++ {
                        v8 = (v8 ^ (v8 * 119))
                        v2[(v7 * (v7 | 0))] = (!(v7 > (v7 & v7)))
                        _ = fi24
                    }
                }
                for fi25 := 0; fi25 < 5; fi25++ {
                    var v26 uint = uint(v8)
                    var v27 string = ("a" + "foo bar")
                    var v28 int32 = 0
                    _ = fi25
                    _ = v26
                    _ = v27
                    _ = v28
                }
                for fi29 := 0; fi29 < 12; fi29++ {
                    if len(v12) > 0 {
                        v12[0] = fi29
                    }
                    _ = fi29
                }
                _ = fi23
            }
        }
    }
    fmt.Println(v7)
    var v30 uint = 255
    v31 := GenS3{F0: 42}
    var v32 float64 = 0.0
    if true {
        if true {
            v8 = (v8 & (v8 % ((v8) | 1)))
            if true {
                if false {
                    var v33 int16 = -1
                    _ = v33
                } else {
                    v30 = (v30 / (((v30 ^ (v30 % ((v30) | 1)))) | 1))
                    v34 := GenS3{F0: 0}
                    delete(v2, genFn1(genFn0()))
                    _ = v34
                }
                fmt.Println(v32)
                v30 -= 1
            } else {
                if true {
                    if (-32768 == -32768) {
                        fmt.Println(v7, v8)
                    }
                }
                _ = len(v12)
                v30 += 955
            }
        }
    }
    v2[v7] = (!true)
    v35 := GenS2{F0: 127, F1: -838}
    var v36 string = "hello"
    for fi37 := 0; fi37 < 9; fi37++ {
        switch v36 {
        case "":
            if true {
                fmt.Println(v32)
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
            }
        default:
            if (!false) {
                var v38 float32 = 0.0
                _ = v38
            }
        }
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        v7 = (v7 & uint32(v30))
        _ = fi37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v35
    _ = v36
    _ = v7
    _ = v8
}
