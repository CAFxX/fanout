package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 uint32 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 uint8 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]int)
    var v4 uint32 = 42
    switch "a" {
    case "":
        v4 = 2
    case "\n\t":
        var v5 float32 = (0.1 * -2.25)
        _ = v5
    case "foo bar":
        fmt.Println(v4)
    default:
        v6 := GenS1{F0: 0, F1: 18446744073709551615, F2: (v4 / ((v4) | 1))}
        _ = v6
    }
    v7 := GenS1{F0: 32, F1: 380, F2: v4}
    if false {
        v8 := []int{0, 660, -268}
        var v9 int = -9223372036854775808
        if true {
            if false {
                var v10 int32 = int32(v9)
                _ = v10
            } else {
                for fi11 := 0; fi11 < 17; fi11++ {
                    v8 = append(v8, fi11)
                    switch (v9 | v9) {
                    case 127:
                        v8 = append(v8, v9)
                    case 4294967295:
                        v12 := GenS2{F0: "hello", F1: v4}
                        _ = v12
                    case 0:
                        v4 -= v4
                    }
                    _ = fi11
                }
                v13 := GenS1{F0: 2, F1: 255, F2: 4294967295}
                switch "\n\t" {
                case "foo bar":
                    v14 := GenS3{F0: -32768}
                    _ = v14
                case "a":
                    var v15 int = (v9 << ((0) & 63))
                    _ = v15
                case "\n\t":
                    if false {
                        for fi16 := 0; fi16 < 8; fi16++ {
                            var v17 uint = 691
                            _ = fi16
                            _ = v17
                        }
                    } else {
                        if (393 == v4) {
                            if false {
                                var v18 uint = 1
                                _ = v18
                            } else {
                                var v19 uint64 = 548
                                var v20 bool = ((2 > 42) || true)
                                _ = v19
                                _ = v20
                            }
                            v9 = v9
                        } else {
                            if false {
                                v9 = 0
                                var v21 float32 = -2.25
                                _ = v21
                            }
                            if true {
                                if false {
                                    var v22 bool = true
                                    _ = len(v8)
                                    _ = v22
                                } else {
                                    if true {
                                        v9 -= 2147483647
                                        var v23 uint8 = 0
                                        v4 = (v4 & v4)
                                        _ = v23
                                    } else {
                                        v24 := v3[350]
                                        var v25 int16 = 42
                                        v26 := GenS2{F0: "hello", F1: 2147483647}
                                        _ = v24
                                        _ = v25
                                        _ = v26
                                    }
                                    var v27 int8 = -48
                                    v9 = v9
                                    _ = v27
                                }
                                if len(v8) >= 2 {
                                    v8 = v8[1:len(v8)]
                                }
                                if len(v8) >= 2 {
                                    v8 = v8[1:len(v8)]
                                }
                            } else {
                                v28 := GenS3{F0: 32767}
                                v29 := GenS1{F0: 1, F1: 1, F2: (v4 >> ((2) & 63))}
                                if (165 <= -905) {
                                    var v30 int = (v9 | (-(v9)))
                                    _ = v30
                                }
                                _ = v28
                                _ = v29
                            }
                        }
                        var v31 bool = (3.14159 <= 1e10)
                        var v32 int32 = -887
                        _ = v31
                        _ = v32
                    }
                }
                _ = v13
            }
            if false {
                if (((((-2.25 + 0.1) + (1.5 / 1.0)) / (1.5 * (-2.25 - 1e10))) + (-2.25 - 1.5)) != 1e10) {
                    var v33 int = v9
                    _ = v33
                } else {
                    var v34 int = (v9 ^ v9)
                    var v35 uint32 = (v4 >> ((65535) & 63))
                    v4 = (v35 & 65535)
                    _ = v34
                    _ = v35
                }
                v36 := GenS4{F0: 42, F1: 41}
                switch v9 {
                case 32767:
                    fmt.Println(v9, v4)
                }
                _ = v36
            }
            if (!true) {
                _ = len(v8)
                v4 = 2
            } else {
                if true {
                    if (!false) {
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                        v9 -= (v9 - 80)
                        var v37 int = (^(v9))
                        _ = v37
                    }
                    var v38 uint32 = (v4 / ((v4) | 1))
                    _ = v38
                } else {
                    var v39 uint16 = 546
                    _ = v39
                }
                if true {
                    var v40 int8 = 100
                    _ = v40
                }
            }
        }
        _ = v8
        _ = v9
    }
    var v41 uint64 = 637
    v41 += 792
    v4 = (^(v4))
    var v42 int32 = 2147483647
    if false {
        v4 += (v4 * (v4 >> ((42) & 63)))
        var v43 bool = true
        var v44 int32 = (^(v42))
        _ = v43
        _ = v44
    } else {
        v45 := GenS5{F0: 42}
        v42 = v42
        v42 = (v42 / (((v42 & 1)) | 1))
        _ = v45
    }
    v42 = v42
    _ = len(v1)
    v4 = 2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v7
}
