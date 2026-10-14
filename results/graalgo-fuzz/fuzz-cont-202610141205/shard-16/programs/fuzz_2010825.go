package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 uint64; F2 uint8 }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 uint64 }

func main() {
    v1 := []int{2147483647}
    v2 := make(map[string]bool)
    var v3 int = 42
    v4 := []uint32{846}
    v5 := make(map[string]string)
    v3 = 4294967295
    fmt.Println(v3)
    var v6 int = (^(v3))
    if false {
        if true {
            var v7 uint64 = 564
            if false {
                v6 -= (v3 % ((4294967295) | 1))
            }
            _ = v7
        } else {
            v8 := GenS1{F0: 2}
            _ = v8
        }
        v9 := GenS1{F0: 255}
        _ = v9
    } else {
        var v10 int64 = 1
        v3 = 32768
        switch int(v10) {
        case 2147483647:
            var v11 uint16 = 32767
            _ = v11
        default:
            if false {
                var v12 int16 = -1
                fmt.Println(v3, v10, v3)
                if (4294967295 > 42) {
                    var v13 uint32 = 0
                    for fi14 := 0; fi14 < 4; fi14++ {
                        v6 = 329
                        var v15 uint64 = 0
                        _ = fi14
                        _ = v15
                    }
                    v6 = (v3 << ((18446744073709551615) & 63))
                    _ = v13
                } else {
                    var v16 uint8 = 42
                    for fi17 := 0; fi17 < 16; fi17++ {
                        v16 = v16
                        _ = fi17
                    }
                    delete(v2, "hello")
                    _ = v16
                }
                _ = v12
            } else {
                switch (("" + "a") + ("a" + "foo bar")) {
                case "a":
                    for fi18 := 0; fi18 < 7; fi18++ {
                        v6 += fi18
                        _ = fi18
                    }
                case "hello":
                    v19 := GenS2{F0: 42, F1: 18446744073709551615, F2: 1}
                    _ = v19
                case "foo bar":
                    switch (^(v6)) {
                    case 2147483647:
                        var v20 uint16 = 42
                        _ = v20
                    case 0:
                        v21 := GenS3{F0: 224, F1: 65535}
                        _ = v21
                    case 32768:
                        v22 := GenS1{F0: 173}
                        _ = v22
                    }
                default:
                    switch v6 {
                    case 0:
                        var v23 uint32 = 0
                        _ = v23
                    case 769:
                        var v24 int64 = 9223372036854775807
                        _ = v24
                    }
                }
                if (49 < 113) {
                    if true {
                        fmt.Println(v3)
                        v10 += (v10 / ((1) | 1))
                        if false {
                            switch "\n\t" {
                            case "hello":
                                var v25 float32 = 1.5
                                _ = v25
                            case "foo bar":
                                v3 -= (v3 << ((0) & 63))
                            case "a":
                                switch (v3 ^ (v6 * v6)) {
                                case 32768:
                                    _ = len(v2)
                                case 4294967295:
                                    v26 := GenS3{F0: 246, F1: 0}
                                    _ = v26
                                case 32767:
                                    v3 += (v3 + v3)
                                default:
                                    for fi27 := 0; fi27 < 13; fi27++ {
                                        v1 = append(v1, v6)
                                        var v28 int8 = 0
                                        _ = fi27
                                        _ = v28
                                    }
                                }
                            }
                            fmt.Println(v3, v3)
                            if true {
                                for fi29 := 0; fi29 < 16; fi29++ {
                                    v1 = append(v1, (fi29 % ((fi29) | 1)))
                                    var v30 uint = 645
                                    _ = fi29
                                    _ = v30
                                }
                                v10 += 312
                                _ = len(v1)
                            }
                        }
                    } else {
                        if (1 >= -1) {
                            switch (v6 / (((v6 << ((42) & 63))) | 1)) {
                            case 924:
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            default:
                                v31 := GenS3{F0: 45, F1: 2147483648}
                                _ = v31
                            }
                        } else {
                            var v32 int16 = -32768
                            v33 := GenS4{F0: 1}
                            if len(v1) > 0 {
                                v1[0] = -2147483648
                            }
                            _ = v32
                            _ = v33
                        }
                        if len(v4) > 0 {
                            v4[0] = 42
                        }
                        v10 = v10
                    }
                }
            }
        }
        _ = v10
    }
    switch ("\n\t" + "a") {
    case "\n\t":
        var v34 uint32 = 67
        _ = v34
    case "hello":
        fmt.Println(v6, v6)
    default:
        for fi35 := 0; fi35 < 7; fi35++ {
            v6 -= v3
            _ = len(v4)
            var v36 uint16 = 581
            _ = fi35
            _ = v36
        }
    }
    v37 := GenS1{F0: 0}
    v6 = (v6 | v3)
    fmt.Println(v6, v6)
    var v38 uint8 = 127
    v6 = v6
    var v39 int16 = 419
    switch "a" {
    case "foo bar":
        delete(v2, "hello")
    case "a":
        v40 := GenS1{F0: v38}
        _ = v40
    case "":
        v41 := GenS4{F0: 2}
        _ = v41
    default:
        fmt.Println(v3)
    }
    v6 += v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
