package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 uint16 }
func genFn0(p0 int8) uint16 {
    var v7 uint32 = 940
    v8 := []int{}
    if len(v8) > 0 {
        v8[0] = -2147483648
    }
    _ = p0
    _ = v7
    _ = v8
    return 65535
}

func main() {
    var v1 int16 = -32768
    var v2 string = "a"
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]bool)
    v5 := make(map[uint32]bool)
    v6 := make(map[string]bool)
    var v9 int = -9223372036854775808
    for fi10 := 0; fi10 < 13; fi10++ {
        var v11 uint64 = 1
        var v12 uint8 = 42
        _ = fi10
        _ = v11
        _ = v12
    }
    if true {
        _ = len(v6)
        v9 += (v9 * v9)
    } else {
        for fi13 := 0; fi13 < 2; fi13++ {
            _ = len(v5)
            fmt.Println(v9, fi13)
            if true {
                for fi14 := 0; fi14 < 4; fi14++ {
                    v9 -= v9
                    _ = fi14
                }
                var v15 uint64 = 0
                _ = v15
            }
            _ = fi13
        }
        v16 := v5[836]
        _ = v16
    }
    v2 = v2
    delete(v4, 35)
    v17 := GenS1{F0: false, F1: 42, F2: 345}
    fmt.Println(v9)
    if (1e10 != -2.25) {
        var v18 int32 = 1
        v2 = v2
        v19 := GenS1{F0: (982 >= 65535), F1: 18446744073709551615, F2: 0}
        _ = v18
        _ = v19
    } else {
        var v20 int = v9
        _ = v20
    }
    for fi21 := 0; fi21 < 9; fi21++ {
        for fi22 := 0; fi22 < 12; fi22++ {
            v23 := GenS1{F0: false, F1: 2, F2: 2}
            switch 4294967295 {
            case 127:
                var v24 float32 = (0.1 + (1.5 + ((0.1 / ((-2.25 - 1e10) / 0.1)) * 1e10)))
                _ = v24
            case -456:
                var v25 uint8 = 127
                _ = v25
            default:
                if (false || (int16(v9) != -809)) {
                    var v26 string = "foo bar"
                    if false {
                        if (1 >= 1) {
                            v26 = "hello"
                        } else {
                            switch "foo bar" {
                            case "hello":
                                if (169 >= 4294967295) {
                                    switch 206 {
                                    case 0:
                                        v27 := GenS1{F0: ((v1 == (-(v1))) && true), F1: 2, F2: 521}
                                        _ = v27
                                    case 128:
                                        var v28 uint = 926
                                        _ = v28
                                    case 1:
                                        if (!((v2 + ((("foo bar" + "") + v26) + v26)) == (v2 + ((v26 + "") + "hello")))) {
                                            if (v26 < v2) {
                                                switch (fi21 - -1) {
                                                case -369:
                                                    v29 := GenS1{F0: false, F1: 311, F2: 2}
                                                    _ = v29
                                                }
                                            } else {
                                                v26 = "a"
                                            }
                                        }
                                    default:
                                        fmt.Println(fi21, fi21, fi22)
                                    }
                                } else {
                                    var v30 uint32 = 509
                                    var v31 int32 = 42
                                    v32 := []int{(^(v9))}
                                    _ = v30
                                    _ = v31
                                    _ = v32
                                }
                            default:
                                fmt.Println(fi22, v2)
                            }
                            v3[703] = fi21
                            v33 := v3[2]
                            _ = v33
                        }
                        v34 := []int{}
                        _ = v34
                    }
                    if true {
                        var v35 int16 = 910
                        v26 = v2
                        _ = v35
                    }
                    _ = v26
                }
            }
            fmt.Println(v2, fi21)
            _ = fi22
            _ = v23
        }
        v9 += 460
        _ = fi21
    }
    switch (v9 * (v9 << ((472) & 63))) {
    case -9223372036854775808:
        var v36 float32 = -2.25
        _ = v36
    case 0:
        v37 := v3[267]
        _ = v37
    default:
        if (!true) {
            for fi38 := 0; fi38 < 1; fi38++ {
                _ = len(v5)
                var v39 uint = 0
                var v40 uint8 = 127
                _ = fi38
                _ = v39
                _ = v40
            }
            if true {
                v1 = (^(v1))
            }
            var v41 uint64 = 4294967295
            _ = v41
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
