package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int64; F1 bool }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]bool)
    v3 := make(map[int]string)
    v4 := make(map[string]string)
    v5 := make(map[uint32]bool)
    var v6 uint16 = 32767
    if (("foo bar" + "") < ("foo bar" + "hello")) {
        var v7 bool = (42 == -1)
        var v8 uint = 1
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 42, F1: 564}
    v10 := []int{1, 4294967295}
    if (565 >= 587) {
        if (("foo bar" != (("\n\t" + "foo bar") + (("\n\t" + "foo bar") + ""))) && false) {
            if true {
                _ = len(v10)
                var v11 uint64 = 506
                v12 := GenS1{F0: (v11 & v11), F1: 127}
                _ = v11
                _ = v12
            }
        } else {
            for fi13 := 0; fi13 < 6; fi13++ {
                fmt.Println(fi13)
                v6 -= (v6 >> ((0) & 63))
                _ = len(v10)
                _ = fi13
            }
            var v14 uint16 = (v6 ^ v6)
            _ = v14
        }
        v6 -= v6
        v6 += v6
    }
    fmt.Println(v6)
    switch "hello" {
    case "\n\t":
        v6 -= 65535
    default:
        if ((!((v6 % ((32767) | 1)) == v6)) && true) {
            if false {
                _ = len(v10)
            }
            if (false || (!((32768 < 127) || false))) {
                v6 = (v6 / ((65535) | 1))
                if false {
                    v15 := GenS1{F0: 2, F1: -9223372036854775808}
                    v10 = append(v10, 128)
                    _ = v15
                }
            }
            switch -9223372036854775808 {
            case -9223372036854775808:
                v6 = 42
            case 42:
                if (false || true) {
                    if (!(42 != 0)) {
                        var v16 int64 = 974
                        fmt.Println(v16)
                        switch "" {
                        case "hello":
                            for fi17 := 0; fi17 < 15; fi17++ {
                                v16 += v16
                                v10 = append(v10, 0)
                                if true {
                                    fmt.Println(fi17, fi17, v16)
                                    for fi18 := 0; fi18 < 11; fi18++ {
                                        v19 := GenS2{F0: 1}
                                        v20 := GenS2{F0: 4294967295}
                                        _ = fi18
                                        _ = v19
                                        _ = v20
                                    }
                                }
                                _ = fi17
                            }
                        case "":
                            v21 := GenS3{F0: -128}
                            _ = v21
                        case "\n\t":
                            v16 += (^(v16))
                        default:
                            var v22 int8 = -1
                            _ = v22
                        }
                        _ = v16
                    }
                    fmt.Println(v6)
                }
            }
        } else {
            var v23 int = -69
            if true {
                for fi24 := 0; fi24 < 6; fi24++ {
                    if (460 < 0) {
                        switch fi24 {
                        case 1:
                            if false {
                                switch fi24 {
                                case 127:
                                    var v25 uint16 = v6
                                    _ = v25
                                }
                                delete(v5, 889)
                                for fi26 := 0; fi26 < 15; fi26++ {
                                    var v27 int64 = 277
                                    fmt.Println(v23, fi26)
                                    _ = fi26
                                    _ = v27
                                }
                            }
                        }
                        var v28 uint32 = 4294967295
                        v28 += v28
                        _ = v28
                    } else {
                        var v29 string = ""
                        if false {
                            for fi30 := 0; fi30 < 2; fi30++ {
                                var v31 int16 = 42
                                _ = fi30
                                _ = v31
                            }
                            v32 := GenS2{F0: 2}
                            switch v29 {
                            case "foo bar":
                                v33 := GenS3{F0: -9}
                                _ = v33
                            }
                            _ = v32
                        }
                        v34 := GenS4{F0: (^(v23))}
                        _ = v29
                        _ = v34
                    }
                    _ = fi24
                }
            }
            var v35 uint16 = v6
            _ = v23
            _ = v35
        }
    }
    fmt.Println(v6)
    var v36 uint8 = 2
    fmt.Println(v6)
    var v37 int8 = 0
    switch ("foo bar" + "hello") {
    case "hello":
        var v38 float32 = (3.14159 + float32(1.5))
        _ = v38
    default:
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
    }
    if len(v10) > 0 {
        v10[0] = 4294967295
    }
    v37 -= -128
    _ = len(v10)
    v39 := GenS5{F0: 966, F1: ((1 == -9223372036854775808) || (!false))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
