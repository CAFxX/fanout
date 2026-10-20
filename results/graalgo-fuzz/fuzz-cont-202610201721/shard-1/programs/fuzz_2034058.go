package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 uint }

func main() {
    var v1 float64 = 0.0
    v2 := make(map[int]uint32)
    var v3 int32 = -2147483648
    var v4 uint32 = 65535
    v5 := make(map[string]string)
    var v6 int8 = -88
    if true {
        v7 := GenS1{F0: 42}
        v8 := GenS2{F0: 255, F1: 1}
        _ = v7
        _ = v8
    } else {
        var v9 uint8 = 2
        _ = v9
    }
    v3 = v3
    switch 561 {
    case 2147483647:
        v1 -= v1
    case -313:
        for fi10 := 0; fi10 < 1; fi10++ {
            fmt.Println(v3)
            v11 := GenS1{F0: 1}
            for fi12 := 0; fi12 < 2; fi12++ {
                fmt.Println(v4, fi12)
                _ = fi12
            }
            _ = fi10
            _ = v11
        }
    default:
        v13 := GenS2{F0: 42, F1: 2}
        _ = v13
    }
    v14 := []int{-1}
    fmt.Println(v4)
    _ = len(v14)
    for fi15 := 0; fi15 < 3; fi15++ {
        var v16 uint64 = 1
        _ = fi15
        _ = v16
    }
    if (false || true) {
        v17 := GenS1{F0: 65535}
        v14 = append(v14, 32767)
        fmt.Println(v4, v1)
        _ = v17
    }
    v6 += 127
    v1 = v1
    if false {
        fmt.Println(v6, v6, v3)
        v14 = append(v14, 386)
    } else {
        if true {
            if (("\n\t" <= "") || true) {
                if (true || false) {
                    var v18 int = 0
                    _ = v18
                }
                if false {
                    var v19 int64 = -9223372036854775808
                    _ = v19
                } else {
                    var v20 uint = 1
                    for fi21 := 0; fi21 < 1; fi21++ {
                        v6 = (^(v6))
                        _ = fi21
                    }
                    var v22 uint32 = 462
                    _ = v20
                    _ = v22
                }
                var v23 float32 = -2.25
                _ = v23
            } else {
                for fi24 := 0; fi24 < 1; fi24++ {
                    v3 -= int32(fi24)
                    _ = fi24
                }
                if false {
                    v25 := v2[727]
                    var v26 uint8 = 42
                    if (!(!(v6 < (v6 * (^(v6)))))) {
                        if len(v14) > 0 {
                            v14[0] = 127
                        }
                        v1 = (v1 / 1.0)
                    }
                    _ = v25
                    _ = v26
                } else {
                    fmt.Println(v3, v1)
                    var v27 bool = (-18 == (-(v6)))
                    v28 := v2[762]
                    _ = v27
                    _ = v28
                }
                v3 -= v3
            }
        }
        for fi29 := 0; fi29 < 16; fi29++ {
            if ((v3 | v3) < v3) {
                if true {
                    if (!true) {
                        if true {
                            if len(v14) >= 2 {
                                v14 = v14[1:len(v14)]
                            }
                        } else {
                            for fi30 := 0; fi30 < 1; fi30++ {
                                v1 = v1
                                v3 += (v3 >> ((0) & 63))
                                _ = fi30
                            }
                            v1 = v1
                            if (42 > -210) {
                                var v31 int8 = -1
                                _ = v31
                            }
                        }
                    } else {
                        v5[""] = ""
                        for fi32 := 0; fi32 < 16; fi32++ {
                            if true {
                                v2[42] = (v4 ^ (^(v4)))
                                _ = len(v5)
                            } else {
                                v1 = v1
                                var v33 string = ""
                                v33 = v33
                                _ = v33
                            }
                            if (false || true) {
                                v6 = v6
                                var v34 int32 = 610
                                _ = v34
                            }
                            _ = fi32
                        }
                        for fi35 := 0; fi35 < 4; fi35++ {
                            var v36 int8 = v6
                            var v37 uint16 = 0
                            _ = fi35
                            _ = v36
                            _ = v37
                        }
                    }
                    v2[(-(fi29))] = (v4 - (v4 & v4))
                } else {
                    switch "a" {
                    case "hello":
                        if (2 <= 1) {
                            v38 := v5["a"]
                            var v39 uint32 = (^(v4))
                            if false {
                                v3 = 1
                                v38 = v38
                            }
                            _ = v38
                            _ = v39
                        } else {
                            v3 = 1
                            v6 = v6
                            v1 -= v1
                        }
                    case "a":
                        fmt.Println(v6, v3, fi29)
                    default:
                        var v40 int64 = 0
                        _ = v40
                    }
                    var v41 uint32 = (v4 >> ((255) & 63))
                    _ = v41
                }
                v42 := GenS2{F0: 920, F1: 127}
                v43 := GenS3{F0: 1}
                _ = v42
                _ = v43
            } else {
                if len(v14) > 0 {
                    v14[0] = (-(fi29))
                }
            }
            _ = fi29
        }
        var v44 uint = 4294967295
        _ = v44
    }
    v1 = ((v1 * v1) * (v1 + float64(float32(v1))))
    if (42 != 255) {
        v45 := GenS2{F0: 466, F1: 1}
        _ = v45
    } else {
        v6 -= (v6 + (v6 / ((v6) | 1)))
    }
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
