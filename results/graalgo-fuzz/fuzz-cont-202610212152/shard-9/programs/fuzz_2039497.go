package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 uint32; F1 int; F2 uint }
type GenS4 struct { F0 uint; F1 int; F2 int32 }

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{1, 2, 0}
    var v3 uint32 = 65535
    v4 := GenS1{F0: 717}
    if true {
        for fi5 := 0; fi5 < 16; fi5++ {
            if false {
                if true {
                    var v6 uint16 = 0
                    _ = v6
                } else {
                    switch fi5 {
                    case 185:
                        var v7 int64 = 267
                        _ = v7
                    default:
                        v8 := GenS1{F0: 617}
                        _ = v8
                    }
                    var v9 string = "hello"
                    _ = v9
                }
                switch "a" {
                case "\n\t":
                    if false {
                        var v10 bool = false
                        _ = v10
                    } else {
                        v3 = 2147483647
                        v3 -= 394
                        switch fi5 {
                        case -1:
                            v11 := GenS2{F0: uint8(v3), F1: 42}
                            _ = v11
                        default:
                            v12 := GenS1{F0: 42}
                            _ = v12
                        }
                    }
                case "hello":
                    if false {
                        v3 -= (^(v3))
                        switch fi5 {
                        case 42:
                            v13 := GenS1{F0: 350}
                            _ = v13
                        }
                        v14 := GenS3{F0: (v3 % ((v3) | 1)), F1: 128, F2: 4}
                        _ = v14
                    }
                }
                v3 = v3
            } else {
                _ = len(v1)
                switch (fi5 & 162) {
                case 0:
                    var v15 float32 = 0.1
                    _ = v15
                case 4294967295:
                    for fi16 := 0; fi16 < 16; fi16++ {
                        v3 = v3
                        if false {
                            var v17 int32 = 2147483647
                            _ = v17
                        } else {
                            var v18 int = 562
                            _ = len(v1)
                            _ = v18
                        }
                        _ = fi16
                    }
                case 376:
                    var v19 int16 = 257
                    _ = v19
                default:
                    for fi20 := 0; fi20 < 6; fi20++ {
                        if (true || false) {
                            v21 := GenS2{F0: 241, F1: 1}
                            fmt.Println(fi5)
                            _ = v21
                        }
                        v22 := GenS2{F0: 127, F1: 4294967295}
                        v3 = 187
                        v3 -= v3
                        if (-9223372036854775808 <= -9223372036854775808) {
                            v3 = (v3 - (v3 / ((v3) | 1)))
                            v3 = (v3 % ((v3) | 1))
                            switch "\n\t" {
                            case "\n\t":
                                v3 -= v3
                            case "hello":
                                v23 := GenS1{F0: 646}
                                _ = v23
                            default:
                                if true {
                                    var v24 int8 = 35
                                    v25 := GenS2{F0: 174, F1: 0}
                                    fmt.Println(v3, v24)
                                    _ = v24
                                    _ = v25
                                }
                            }
                            fmt.Println(fi5, fi20)
                            fmt.Println(v3, fi20)
                            fmt.Println(fi5, fi5)
                        } else {
                            v26 := GenS3{F0: 2147483648, F1: fi20, F2: 4294967295}
                            _ = v26
                        }
                        _ = fi20
                        _ = v22
                    }
                }
                for fi27 := 0; fi27 < 4; fi27++ {
                    v3 -= 42
                    _ = len(v2)
                    _ = fi27
                }
            }
            for fi28 := 0; fi28 < 8; fi28++ {
                v29 := GenS4{F0: 255, F1: 65, F2: -9}
                switch (-(fi5)) {
                case -945:
                    fmt.Println(fi28, fi28, fi28)
                case -1:
                    fmt.Println(fi5, fi28)
                case 343:
                    var v30 bool = ((0.1 != 3.14159) || (!((v3 + v3) < (v3 >> ((0) & 63)))))
                    _ = v30
                default:
                    v3 += 2147483648
                }
                v1[v3] = -958
                _ = fi28
                _ = v29
            }
            var v31 int32 = -1
            _ = fi5
            _ = v31
        }
        fmt.Println(v3)
        if true {
            if false {
                var v32 int32 = -290
                v3 = (^(v3))
                if (42 != 255) {
                    if false {
                        v33 := GenS3{F0: v3, F1: 4294967295, F2: 2}
                        _ = v33
                    }
                }
                _ = v32
            } else {
                var v34 int32 = 1
                v35 := GenS1{F0: 951}
                _ = v34
                _ = v35
            }
            var v36 bool = (!true)
            _ = v36
        }
    }
    v3 = v3
    v3 -= (^(v3))
    var v37 uint8 = 115
    fmt.Println(v3)
    if false {
        for fi38 := 0; fi38 < 20; fi38++ {
            v37 -= 2
            _ = fi38
        }
        v37 += v37
    }
    if false {
        v37 = v37
        if (-9223372036854775808 <= 874) {
            var v39 int = 32767
            v40 := GenS3{F0: (^(v3)), F1: v39, F2: 2}
            _ = v39
            _ = v40
        }
        if false {
            v3 += (v3 * v3)
        }
    } else {
        v37 += (v37 % (((v37 ^ v37)) | 1))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
}
