package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 bool }
func genFn0(p0 int, p1 int32, p2 uint32) uint64 {
    var v4 int64 = int64(p1)
    var v5 uint16 = 42
    v6 := []int{(p0 / ((p0) | 1)), p0, (p0 & (-(p0)))}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 0
}
func genFn1() bool {
    v7 := []int{-439}
    v8 := GenS1{F0: 1, F1: 42}
    if len(v7) > 0 {
        v7[0] = -341
    }
    var v9 int8 = 42
    for fi10 := 0; fi10 < 18; fi10++ {
        var v11 uint8 = 255
        _ = fi10
        _ = v11
    }
    _ = v7
    _ = v8
    _ = v9
    return true
}

func main() {
    v1 := []uint32{}
    v2 := make(map[int]int)
    v3 := make(map[int]bool)
    var v12 int8 = -14
    v12 = (v12 | (v12 & (-(v12))))
    if true {
        for fi13 := 0; fi13 < 17; fi13++ {
            if (true || false) {
                v3[(fi13 << ((2) & 63))] = false
                if len(v1) > 0 {
                    v1[0] = 1
                }
                v12 = v12
            } else {
                var v14 uint64 = genFn0(fi13, 42, 202)
                _ = v14
            }
            var v15 bool = (1 != -1)
            delete(v2, (fi13 + (-(fi13))))
            _ = fi13
            _ = v15
        }
    } else {
        if (true || false) {
            var v16 string = "\n\t"
            _ = v16
        } else {
            switch "" {
            case "a":
                v12 += (v12 + (v12 | (v12 % (((v12 % (((v12 * (v12 ^ v12))) | 1))) | 1))))
            case "":
                switch ((("hello" + "hello") + "foo bar") + "hello") {
                case "\n\t":
                    v12 += (v12 >> ((1) & 63))
                case "":
                    delete(v3, -9223372036854775808)
                default:
                    var v17 int64 = -1
                    _ = v17
                }
            }
            var v18 float64 = -2.25
            _ = v18
        }
        if len(v1) > 0 {
            v1[0] = 620
        }
    }
    v19 := GenS1{F0: -1, F1: 0}
    for fi20 := 0; fi20 < 6; fi20++ {
        for fi21 := 0; fi21 < 1; fi21++ {
            switch (fi21 - -2147483648) {
            case 2147483647:
                v12 = (v12 / (((v12 + (v12 ^ 0))) | 1))
            default:
                v12 = 127
                var v22 float32 = 0.0
                _ = v22
            }
            _ = fi21
        }
        v12 = -76
        _ = fi20
    }
    var v23 uint32 = 30
    if false {
        if true {
            if true {
                v23 += 2
                fmt.Println(v12, v23)
            } else {
                for fi24 := 0; fi24 < 20; fi24++ {
                    switch fi24 {
                    case -1:
                        if true {
                            v12 += (-(v12))
                            _ = len(v1)
                        } else {
                            var v25 uint16 = 0
                            fmt.Println(v23)
                            _ = v25
                        }
                    default:
                        fmt.Println(v23, v12, v12)
                    }
                    for fi26 := 0; fi26 < 10; fi26++ {
                        switch (fi24 << ((871) & 63)) {
                        case 0:
                            fmt.Println(v12)
                        default:
                            var v27 uint64 = 875
                            _ = v27
                        }
                        var v28 float32 = float32(1.5)
                        _ = fi26
                        _ = v28
                    }
                    _ = fi24
                }
                var v29 uint32 = (^(v23))
                _ = len(v1)
                _ = v29
            }
            v23 += v23
            for fi30 := 0; fi30 < 5; fi30++ {
                v31 := GenS1{F0: -1, F1: 926}
                v23 -= 475
                _ = fi30
                _ = v31
            }
        } else {
            v32 := GenS2{F0: 1, F1: 547}
            _ = v32
        }
        for fi33 := 0; fi33 < 13; fi33++ {
            if false {
                if false {
                    if (true || (255 >= uint8(v23))) {
                        var v34 uint = 4294967295
                        _ = v34
                    }
                }
                switch fi33 {
                case 0:
                    v35 := GenS3{F0: ((!true) && (1 != 42))}
                    _ = v35
                case 590:
                    var v36 bool = (false && true)
                    _ = v36
                }
                var v37 int64 = 0
                _ = v37
            }
            _ = fi33
        }
        v38 := v3[128]
        _ = v38
    }
    v2[128] = int(v12)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v23
    _ = v3
}
