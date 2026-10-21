package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 uint8; F1 int64; F2 bool }
type GenS3 struct { F0 int16; F1 int64 }

func main() {
    var v1 uint16 = 1
    var v2 float32 = 0.0
    v3 := make(map[uint32]bool)
    fmt.Println(v2, v2)
    for fi4 := 0; fi4 < 18; fi4++ {
        var v5 int64 = -9223372036854775808
        for fi6 := 0; fi6 < 2; fi6++ {
            var v7 string = "foo bar"
            if (568 != 363) {
                v2 += v2
            } else {
                var v8 int64 = v5
                v3[951] = false
                _ = v8
            }
            if false {
                switch v7 {
                case "a":
                    v9 := GenS1{F0: (200 <= 0), F1: -128}
                    _ = v9
                default:
                    switch "hello" {
                    case "hello":
                        var v10 float64 = 0.0
                        _ = v10
                    }
                }
            } else {
                var v11 string = v7
                if true {
                    switch ((("foo bar" + (v7 + v11)) + v11) + v7) {
                    case "foo bar":
                        if true {
                            var v12 int32 = 0
                            if false {
                                v1 -= (^(v1))
                                v12 = (v12 >> ((268) & 63))
                                v13 := GenS1{F0: (228 <= 65535), F1: -58}
                                _ = v13
                            }
                            _ = v12
                        } else {
                            var v14 uint64 = 2
                            v15 := GenS1{F0: (32767 == int16(v14)), F1: 127}
                            var v16 int = fi6
                            var v17 int32 = int32(v5)
                            _ = v14
                            _ = v15
                            _ = v16
                            _ = v17
                        }
                    case "\n\t":
                        v2 += 3.14159
                    case "a":
                        v5 = (-(v5))
                    default:
                        if false {
                            v11 = v11
                        } else {
                            switch 9223372036854775807 {
                            case -2147483648:
                                if false {
                                    var v18 string = v7
                                    var v19 int64 = 0
                                    v20 := []int{fi6, (fi6 ^ fi6)}
                                    _ = v18
                                    _ = v19
                                    _ = v20
                                } else {
                                    if (0 == 65535) {
                                        v1 += v1
                                    }
                                    v1 -= (v1 >> ((18446744073709551615) & 63))
                                }
                            default:
                                v21 := []int{fi6, fi6, 42, (fi6 + (fi4 ^ 32768))}
                                _ = v21
                            }
                            switch 32768 {
                            case 8:
                                switch v7 {
                                case "hello":
                                    v7 = ("" + v7)
                                }
                            }
                        }
                    }
                    fmt.Println(v11, v5)
                    switch v7 {
                    case "foo bar":
                        fmt.Println(fi4, fi4)
                    }
                    v22 := GenS2{F0: 255, F1: (v5 | v5), F2: false}
                    _ = v22
                }
                v3[602] = ((4294967295 > 693) || (!(((v5 / (((v5 << ((0) & 63))) | 1)) <= int64(v1)) || (!false))))
                _ = v11
            }
            _ = fi6
            _ = v7
        }
        for fi23 := 0; fi23 < 15; fi23++ {
            v1 = v1
            _ = fi23
        }
        _ = fi4
        _ = v5
    }
    delete(v3, 572)
    v2 = v2
    v3[284] = true
    var v24 uint8 = 25
    var v25 string = "foo bar"
    if false {
        for fi26 := 0; fi26 < 20; fi26++ {
            fmt.Println(v1)
            if (!(!false)) {
                for fi27 := 0; fi27 < 6; fi27++ {
                    v28 := GenS2{F0: v24, F1: 0, F2: false}
                    fmt.Println(v1)
                    _ = fi27
                    _ = v28
                }
            }
            var v29 int64 = 9223372036854775807
            _ = fi26
            _ = v29
        }
        v2 = 1.5
        v25 = v25
    }
    delete(v3, 2147483648)
    var v30 int64 = -1
    for fi31 := 0; fi31 < 3; fi31++ {
        if (v1 > 1) {
            v32 := GenS3{F0: -1, F1: (v30 & (v30 * v30))}
            _ = v32
        } else {
            v33 := GenS3{F0: -32768, F1: v30}
            fmt.Println(v30)
            v34 := GenS2{F0: (v24 + 239), F1: v30, F2: false}
            _ = v33
            _ = v34
        }
        if true {
            v24 = v24
            var v35 float32 = v2
            var v36 int32 = 42
            _ = v35
            _ = v36
        }
        _ = fi31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v30
}
