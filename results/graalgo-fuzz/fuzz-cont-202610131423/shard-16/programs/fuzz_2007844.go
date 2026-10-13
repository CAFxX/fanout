package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8; F1 string; F2 int }

func main() {
    var v1 uint64 = 0
    v2 := make(map[uint32]int)
    var v3 float32 = (1e10 / ((-2.25 + ((3.14159 * 0.0) * 1.5)) / 1.0))
    var v4 uint8 = 42
    v4 -= (^(v4))
    for fi5 := 0; fi5 < 10; fi5++ {
        switch "\n\t" {
        case "hello":
            for fi6 := 0; fi6 < 2; fi6++ {
                if (v4 == (v4 * v4)) {
                    v3 = (0.1 / 1.0)
                    if false {
                        v7 := []int{}
                        v1 -= (^(v1))
                        v2[4294967295] = (fi5 - (fi6 ^ (-(fi6))))
                        _ = v7
                    } else {
                        _ = len(v2)
                        if (32767 != -1) {
                            fmt.Println(v1, fi5)
                        } else {
                            v3 = 3.14159
                            var v8 uint64 = 4294967295
                            var v9 uint16 = 42
                            _ = v8
                            _ = v9
                        }
                        var v10 float64 = 0.1
                        _ = v10
                    }
                    switch (^(fi6)) {
                    case 2147483647:
                        var v11 float64 = 0.1
                        _ = v11
                    default:
                        v1 -= 18446744073709551615
                    }
                    fmt.Println(v1, v4)
                    var v12 uint64 = (^(v1))
                    _ = v12
                } else {
                    delete(v2, 2147483647)
                }
                if ((9 == -128) && (v1 >= (v1 & (^(v1))))) {
                    v1 = 2
                } else {
                    fmt.Println(v4, fi6, v4)
                    switch "\n\t" {
                    case "foo bar":
                        if true {
                            v13 := GenS1{F0: 32767, F1: 2, F2: true}
                            var v14 int8 = 0
                            var v15 uint = 18446744073709551615
                            if false {
                                v16 := []int{}
                                _ = len(v16)
                                v4 += v4
                                switch "\n\t" {
                                case "a":
                                    v17 := GenS2{F0: 65535}
                                    _ = v17
                                case "foo bar":
                                    v1 = (v1 % ((v1) | 1))
                                }
                                v1 = 2
                                _ = v16
                            }
                            _ = v13
                            _ = v14
                            _ = v15
                        } else {
                            if (("foo bar" != "") && true) {
                                v18 := GenS1{F0: 1, F1: 42, F2: true}
                                if (v1 <= 2) {
                                    fmt.Println(v3, fi5)
                                } else {
                                    if (42 >= 429) {
                                        v1 = (v1 << ((255) & 63))
                                    }
                                    var v19 uint = 0
                                    v3 += v3
                                    _ = v19
                                }
                                if true {
                                    var v20 int16 = 32767
                                    fmt.Println(fi5)
                                    v3 = 1e10
                                    _ = v20
                                }
                                _ = v18
                            } else {
                                if (665 != 903) {
                                    v21 := v2[2147483647]
                                    var v22 int8 = 24
                                    _ = v21
                                    _ = v22
                                }
                            }
                            var v23 string = "foo bar"
                            _ = v23
                        }
                    }
                }
                _ = fi6
            }
        }
        _ = fi5
    }
    var v24 bool = true
    v25 := GenS3{F0: 0, F1: "a", F2: 127}
    var v26 uint8 = 76
    v1 = v1
    switch "hello" {
    case "":
        var v27 int32 = -1
        _ = v27
    case "a":
        var v28 uint = 42
        _ = v28
    default:
        fmt.Println(v24, v26, v4)
    }
    v4 = (^(v4))
    v29 := []int{-1, 21, 325, 0}
    var v30 uint8 = (v4 | v26)
    var v31 float32 = ((v3 * -2.25) * 3.14159)
    if (-32768 > -992) {
        v30 = (v4 | 42)
    } else {
        var v32 uint8 = (v30 % ((v26) | 1))
        for fi33 := 0; fi33 < 18; fi33++ {
            var v34 uint = 42
            for fi35 := 0; fi35 < 18; fi35++ {
                if (271 != 546) {
                    if v24 {
                        if v24 {
                            v2[1] = (fi35 << (((v34 + (^(v34)))) & 63))
                            fmt.Println(v3)
                        }
                    }
                    var v36 uint32 = 2147483648
                    if len(v29) >= 2 {
                        v29 = v29[1:len(v29)]
                    }
                    v1 -= v1
                    _ = v36
                } else {
                    if v24 {
                        v37 := GenS3{F0: -8, F1: "foo bar", F2: (fi35 + 32767)}
                        if (!(-2147483648 > 1)) {
                            v24 = (1 == 255)
                        }
                        v38 := GenS1{F0: 32767, F1: 940, F2: (v24 && ((("\n\t" + (("foo bar" + ("hello" + "\n\t")) + "hello")) <= "hello") || false))}
                        _ = v37
                        _ = v38
                    }
                    var v39 int16 = 992
                    _ = v39
                }
                var v40 uint8 = v30
                var v41 bool = (32767 < uint16(v1))
                _ = fi35
                _ = v40
                _ = v41
            }
            var v42 uint16 = uint16(v26)
            _ = fi33
            _ = v34
            _ = v42
        }
        _ = v32
    }
    v24 = false
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
