package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 string; F1 uint; F2 uint16 }
type GenS3 struct { F0 uint; F1 string }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]uint32)
    var v3 int8 = 1
    v4 := []uint32{493, 65535, 2147483647}
    var v5 float32 = (0.1 + 1.5)
    v6 := []uint32{}
    v7 := GenS1{F0: 1000, F1: 317}
    v8 := GenS1{F0: 1000, F1: 0}
    v5 = (v5 * v5)
    var v9 int = -1
    v5 += (-2.25 / v5)
    fmt.Println(v3, v3)
    v3 = v3
    var v10 uint32 = 843
    if false {
        if false {
            if ((true || true) || false) {
                if false {
                    var v11 uint8 = 165
                    v5 = v5
                    v12 := GenS1{F0: -808, F1: (^(v10))}
                    _ = v11
                    _ = v12
                } else {
                    var v13 int8 = -1
                    var v14 int64 = 0
                    _ = v13
                    _ = v14
                }
                var v15 uint64 = 42
                _ = v15
            }
        } else {
            v5 += ((v5 * v5) / (v5 + v5))
            for fi16 := 0; fi16 < 1; fi16++ {
                var v17 string = "hello"
                v18 := GenS1{F0: 0, F1: 797}
                _ = fi16
                _ = v17
                _ = v18
            }
        }
        v9 = (-(v9))
    } else {
        switch "" {
        case "foo bar":
            fmt.Println(v3, v9, v10)
        }
    }
    var v19 string = ""
    for fi20 := 0; fi20 < 7; fi20++ {
        for fi21 := 0; fi21 < 14; fi21++ {
            var v22 int16 = 968
            v23 := GenS1{F0: (v22 + -32768), F1: 2147483647}
            v24 := v2[1]
            _ = fi21
            _ = v22
            _ = v23
            _ = v24
        }
        if (!false) {
            var v25 int64 = -1
            if (!(0.1 == float64(v5))) {
                switch "a" {
                case "hello":
                    if (!true) {
                        var v26 int64 = v25
                        var v27 float32 = 0.0
                        _ = v26
                        _ = v27
                    }
                case "\n\t":
                    if (false && true) {
                        v1[v10] = (v10 - 2147483647)
                    }
                case "foo bar":
                    v25 += 42
                default:
                    if false {
                        if (0 < 0) {
                            var v28 bool = true
                            _ = v28
                        }
                        v29 := GenS2{F0: v19, F1: 42, F2: 65535}
                        _ = v29
                    } else {
                        if true {
                            var v30 float64 = 1e10
                            v31 := GenS2{F0: "hello", F1: 18446744073709551615, F2: 255}
                            if (false || false) {
                                fmt.Println(v9, v19)
                            } else {
                                for fi32 := 0; fi32 < 20; fi32++ {
                                    var v33 uint32 = v10
                                    _ = fi32
                                    _ = v33
                                }
                                for fi34 := 0; fi34 < 18; fi34++ {
                                    if true {
                                        if false {
                                            if (!(!(v30 <= (v30 * (1e10 + float64(-2.25)))))) {
                                                var v35 int = (fi34 + (fi34 % ((v9) | 1)))
                                                _ = v35
                                            } else {
                                                delete(v2, (^(v10)))
                                                var v36 uint = 18446744073709551615
                                                var v37 bool = (v25 != (v25 & 0))
                                                _ = v36
                                                _ = v37
                                            }
                                            v3 += v3
                                        }
                                        v1[(^(v10))] = (v10 / (((v10 ^ v10)) | 1))
                                        var v38 bool = ((!(false || (521 == 2))) && (true || false))
                                        _ = v38
                                    } else {
                                        _ = len(v1)
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        _ = len(v2)
                                    }
                                    _ = fi34
                                }
                                if true {
                                    var v39 bool = false
                                    var v40 int = 128
                                    var v41 uint8 = 243
                                    _ = v39
                                    _ = v40
                                    _ = v41
                                } else {
                                    v25 = (v25 & v25)
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                            }
                            _ = v30
                            _ = v31
                        }
                        for fi42 := 0; fi42 < 7; fi42++ {
                            var v43 int8 = v3
                            _ = fi42
                            _ = v43
                        }
                    }
                }
            }
            _ = v25
        } else {
            for fi44 := 0; fi44 < 7; fi44++ {
                v10 += 1
                _ = fi44
            }
            v45 := GenS2{F0: "hello", F1: 526, F2: 2}
            _ = v45
        }
        v3 = v3
        _ = fi20
    }
    v46 := GenS3{F0: 4294967295, F1: v19}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
