package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 int16 }
func genFn0(p0 uint8, p1 int64) int64 {
    v5 := GenS1{F0: 42, F1: ("\n\t" + "hello")}
    if (true && (-2.25 >= -2.25)) {
        var v6 uint64 = 337
        if false {
            var v7 uint = 1
            v8 := GenS1{F0: 2147483647, F1: ("hello" + "")}
            _ = v7
            _ = v8
        }
        _ = v6
    }
    var v9 float64 = 1.5
    if (((0.1 - 0.0) != 3.14159) && false) {
        for fi10 := 0; fi10 < 16; fi10++ {
            p1 += -9223372036854775808
            _ = fi10
        }
    }
    p1 = (^(p1))
    _ = p0
    _ = p1
    _ = v5
    _ = v9
    return (p1 >> ((2) & 63))
}

func main() {
    var v1 uint16 = 800
    v2 := make(map[int]string)
    var v3 uint16 = v1
    v4 := make(map[int]bool)
    for fi11 := 0; fi11 < 8; fi11++ {
        if true {
            v12 := GenS1{F0: -1, F1: "\n\t"}
            var v13 int8 = 1
            _ = v12
            _ = v13
        } else {
            var v14 uint = 608
            _ = v14
        }
        _ = fi11
    }
    fmt.Println(v3)
    fmt.Println(v1, v3)
    var v15 uint32 = 65535
    fmt.Println(v1, v3)
    var v16 int = 32768
    var v17 bool = true
    v18 := []int{(v16 / (((v16 % ((v16) | 1))) | 1)), (v16 - (v16 % ((v16) | 1))), v16, (v16 * v16)}
    v15 = (v15 << ((458) & 63))
    v19 := v2[(^(v16))]
    v17 = false
    if v17 {
        switch (v16 * (v16 | v16)) {
        case 2147483647:
            if len(v18) >= 2 {
                v18 = v18[1:len(v18)]
            }
        case -1:
            for fi20 := 0; fi20 < 3; fi20++ {
                v17 = (v17 || (!v17))
                switch v19 {
                case "\n\t":
                    v21 := GenS2{F0: 947, F1: 42}
                    _ = v21
                case "a":
                    v18 = append(v18, (-(fi20)))
                default:
                    v1 = 2
                }
                switch "hello" {
                case "":
                    var v22 uint64 = 1
                    _ = v22
                case "foo bar":
                    fmt.Println(fi20, fi20, v17)
                case "a":
                    if v17 {
                        for fi23 := 0; fi23 < 4; fi23++ {
                            if len(v18) >= 2 {
                                v18 = v18[1:len(v18)]
                            }
                            v18 = append(v18, fi20)
                            _ = fi23
                        }
                    } else {
                        fmt.Println(v17, v3)
                        var v24 uint16 = 42
                        if len(v18) >= 2 {
                            v18 = v18[1:len(v18)]
                        }
                        _ = v24
                    }
                default:
                    switch (v19 + v19) {
                    case "a":
                        var v25 bool = v17
                        _ = v25
                    case "hello":
                        if true {
                            if ((v15 - v15) == v15) {
                                if false {
                                    var v26 float32 = (3.14159 * float32(1.5))
                                    if (!true) {
                                        v4[fi20] = true
                                        if v17 {
                                            v18 = append(v18, (v16 % (((v16 >> ((42) & 63))) | 1)))
                                            if false {
                                                _ = len(v4)
                                            }
                                            v27 := GenS3{F0: (v15 / ((v15) | 1)), F1: v19}
                                            _ = v27
                                        }
                                    }
                                    var v28 int = int(v15)
                                    _ = v26
                                    _ = v28
                                } else {
                                    v29 := GenS4{F0: 0}
                                    _ = v29
                                }
                            }
                            switch 4294967295 {
                            case 128:
                                for fi30 := 0; fi30 < 10; fi30++ {
                                    delete(v2, (fi20 * 9223372036854775807))
                                    v1 = (v1 - 0)
                                    _ = fi30
                                }
                            case -134:
                                var v31 int32 = 235
                                _ = v31
                            case 32767:
                                var v32 uint16 = uint16(v15)
                                _ = v32
                            }
                            if ((78 <= 1) && (65535 != 255)) {
                                switch (v16 % ((fi20) | 1)) {
                                case 32767:
                                    for fi33 := 0; fi33 < 17; fi33++ {
                                        fmt.Println(v3, v15, v16)
                                        v3 -= (v3 >> ((1) & 63))
                                        _ = fi33
                                    }
                                }
                                var v34 uint8 = 241
                                for fi35 := 0; fi35 < 7; fi35++ {
                                    delete(v4, -1)
                                    var v36 float64 = 1e10
                                    var v37 uint = 255
                                    v34 = 128
                                    _ = fi35
                                    _ = v36
                                    _ = v37
                                }
                                _ = v34
                            }
                        }
                    case "\n\t":
                        for fi38 := 0; fi38 < 17; fi38++ {
                            var v39 float64 = 0.0
                            _ = fi38
                            _ = v39
                        }
                    default:
                        var v40 uint8 = 2
                        _ = v40
                    }
                }
                _ = fi20
            }
        default:
            if (("foo bar" + "") != ("a" + v19)) {
                if len(v18) >= 2 {
                    v18 = v18[1:len(v18)]
                }
                delete(v4, (-(v16)))
            } else {
                v17 = v17
                v17 = false
                if len(v18) >= 2 {
                    v18 = v18[1:len(v18)]
                }
            }
        }
        var v41 float32 = ((0.0 / -2.25) * (-2.25 * 0.0))
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
