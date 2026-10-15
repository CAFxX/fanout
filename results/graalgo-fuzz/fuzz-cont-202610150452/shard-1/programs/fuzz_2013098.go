package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 uint32; F1 string; F2 int64 }
type GenS4 struct { F0 uint32; F1 int16; F2 uint64 }
type GenS5 struct { F0 uint16; F1 int8; F2 uint32 }
func genFn0() uint16 {
    var v5 int = 42
    v6 := []int{(-(v5)), v5, v5, (v5 + v5)}
    if true {
        if (-128 != 119) {
            v5 -= v5
        } else {
            v5 = v5
        }
    } else {
        var v7 uint64 = 2
        v7 = v7
        for fi8 := 0; fi8 < 16; fi8++ {
            var v9 bool = (2147483647 < 4294967295)
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            for fi10 := 0; fi10 < 9; fi10++ {
                if (!v9) {
                    v11 := GenS1{F0: (!(!(0.0 >= 3.14159)))}
                    v9 = (!((!true) && (!((!(!false)) || false))))
                    _ = v11
                } else {
                    v12 := GenS2{F0: 206, F1: (true || (fi8 >= 9223372036854775807))}
                    if len(v6) > 0 {
                        v6[0] = (v5 + fi8)
                    }
                    var v13 uint32 = 65535
                    _ = v12
                    _ = v13
                }
                _ = len(v6)
                switch "a" {
                case "":
                    var v14 uint8 = 0
                    switch (fi10 + fi10) {
                    case 1:
                        v7 = 42
                    case 4294967295:
                        var v15 float32 = 1e10
                        _ = v15
                    }
                    if v9 {
                        v16 := GenS3{F0: 154, F1: ("a" + "a"), F2: 767}
                        fmt.Println(fi10)
                        var v17 float32 = 0.1
                        _ = v16
                        _ = v17
                    } else {
                        switch fi10 {
                        case -2147483648:
                            switch "foo bar" {
                            case "\n\t":
                                switch "hello" {
                                case "foo bar":
                                    var v18 uint64 = v7
                                    _ = v18
                                case "hello":
                                    if ((fi10 + v5) <= v5) {
                                        if ((0 < 0) && false) {
                                            v14 -= 128
                                        } else {
                                            switch (fi10 / (((fi10 ^ fi8)) | 1)) {
                                            case 127:
                                                v19 := GenS1{F0: true}
                                                _ = v19
                                            default:
                                                if (!(v9 || v9)) {
                                                    v5 += (fi8 | v5)
                                                    switch 519 {
                                                    case -1:
                                                        v6 = append(v6, fi10)
                                                    }
                                                    v14 = 42
                                                } else {
                                                    var v20 uint8 = (v14 ^ (v14 + (^(v14))))
                                                    var v21 uint32 = 4294967295
                                                    _ = v20
                                                    _ = v21
                                                }
                                            }
                                        }
                                    } else {
                                        v6 = append(v6, (^(v5)))
                                    }
                                }
                            }
                        case 32767:
                            v7 += (^(v7))
                        default:
                            var v22 int = fi10
                            _ = v22
                        }
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        var v23 uint32 = 867
                        _ = v23
                    }
                    v24 := GenS4{F0: 302, F1: 1, F2: (v7 << ((uint(v7)) & 63))}
                    v25 := GenS3{F0: 2, F1: (("foo bar" + "hello") + ""), F2: 0}
                    _ = v14
                    _ = v24
                    _ = v25
                }
                if false {
                    var v26 uint8 = 176
                    _ = v26
                } else {
                    if false {
                        var v27 uint64 = v7
                        var v28 float32 = (0.0 / float32(-2.25))
                        fmt.Println(v7)
                        _ = v27
                        _ = v28
                    } else {
                        v5 += fi8
                    }
                    var v29 uint64 = (^(v7))
                    fmt.Println(v9, v7)
                    _ = v29
                }
                var v30 uint16 = 406
                _ = fi10
                _ = v30
            }
            _ = fi8
            _ = v9
        }
        _ = v7
    }
    _ = v5
    _ = v6
    return 840
}

func main() {
    v1 := []int{}
    v2 := make(map[uint32]uint32)
    var v3 int32 = 2147483647
    v4 := []int{32767}
    v31 := GenS2{F0: -598, F1: false}
    fmt.Println(v3)
    v32 := v2[2147483648]
    v33 := GenS4{F0: v32, F1: -1, F2: 977}
    var v34 bool = true
    v35 := GenS5{F0: 0, F1: 116, F2: (v32 ^ (^(v32)))}
    var v36 bool = v34
    var v37 float64 = -2.25
    v34 = true
    v34 = ("a" >= "hello")
    var v38 uint8 = 255
    v39 := GenS4{F0: (^(v32)), F1: -1, F2: 317}
    v38 -= (v38 + (^(v38)))
    for fi40 := 0; fi40 < 5; fi40++ {
        var v41 uint16 = genFn0()
        _ = fi40
        _ = v41
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
}
