package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 string; F2 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint8 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]int)
    v3 := make(map[string]uint32)
    var v4 uint = 1
    v5 := []int{32767, 394, 32767, 128}
    switch 4294967295 {
    case -333:
        v6 := GenS1{F0: 2}
        _ = v6
    case 42:
        var v7 uint32 = 2147483647
        _ = v7
    default:
        delete(v3, "")
    }
    v4 += v4
    if (true || true) {
        v8 := GenS1{F0: 18446744073709551615}
        if (!false) {
            v4 += (v4 | v4)
            if true {
                v4 = (v4 * (v4 ^ (v4 + v4)))
                fmt.Println(v4)
            } else {
                fmt.Println(v4)
            }
        } else {
            var v9 int32 = -1
            _ = v9
        }
        _ = v8
    } else {
        if (false && (42 == 0)) {
            if (!(0 >= -73)) {
                for fi10 := 0; fi10 < 11; fi10++ {
                    switch fi10 {
                    case 2147483647:
                        for fi11 := 0; fi11 < 17; fi11++ {
                            v4 -= (^(v4))
                            v12 := GenS1{F0: 18446744073709551615}
                            _ = fi11
                            _ = v12
                        }
                    case -1:
                        if (false && true) {
                            var v13 uint64 = 2
                            for fi14 := 0; fi14 < 20; fi14++ {
                                v15 := GenS1{F0: (^(v13))}
                                _ = fi14
                                _ = v15
                            }
                            if false {
                                var v16 int = (fi10 + fi10)
                                var v17 uint16 = 255
                                _ = v16
                                _ = v17
                            } else {
                                v18 := GenS1{F0: (v13 & (^(v13)))}
                                switch (fi10 - (fi10 + (fi10 / ((-9223372036854775808) | 1)))) {
                                case 4294967295:
                                    for fi19 := 0; fi19 < 4; fi19++ {
                                        delete(v3, "\n\t")
                                        var v20 string = "hello"
                                        switch v20 {
                                        case "a":
                                            v20 = ((v20 + (v20 + (("a" + v20) + (v20 + v20)))) + v20)
                                            if (v13 > (v13 & (v13 % ((v13) | 1)))) {
                                                v20 = v20
                                                var v21 uint8 = 65
                                                _ = v21
                                            } else {
                                                v22 := GenS2{F0: 65535, F1: v20, F2: 2147483647}
                                                _ = v22
                                            }
                                            v13 = v13
                                        case "hello":
                                            var v23 int64 = 424
                                            _ = v23
                                        default:
                                            v24 := GenS3{F0: (v4 * (v4 % (((v4 + 255)) | 1)))}
                                            _ = v24
                                        }
                                        _ = fi19
                                        _ = v20
                                    }
                                case -678:
                                    v13 += v13
                                case 711:
                                    for fi25 := 0; fi25 < 6; fi25++ {
                                        if true {
                                            fmt.Println(fi25, v4)
                                        } else {
                                            switch ("a" + (("hello" + "a") + "foo bar")) {
                                            case "foo bar":
                                                var v26 uint8 = 145
                                                delete(v3, "foo bar")
                                                var v27 uint = v4
                                                _ = v26
                                                _ = v27
                                            case "":
                                                v13 = 42
                                            }
                                            if true {
                                                v5 = append(v5, fi10)
                                            }
                                        }
                                        var v28 float64 = float64(0.1)
                                        _ = fi25
                                        _ = v28
                                    }
                                default:
                                    var v29 uint16 = 65535
                                    _ = v29
                                }
                                _ = v18
                            }
                            _ = v13
                        } else {
                            for fi30 := 0; fi30 < 9; fi30++ {
                                var v31 int = 127
                                var v32 int32 = -715
                                if false {
                                    if false {
                                        v33 := GenS4{F0: 111}
                                        _ = v33
                                    }
                                    v32 = 0
                                }
                                v2[2] = fi10
                                v31 = fi30
                                _ = fi30
                                _ = v31
                                _ = v32
                            }
                        }
                    case 42:
                        if true {
                            if false {
                                v4 += v4
                                v4 = (v4 / (((v4 << ((1) & 63))) | 1))
                            }
                            _ = len(v5)
                        }
                    default:
                        fmt.Println(v4, fi10)
                    }
                    fmt.Println(fi10, v4)
                    _ = fi10
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                var v34 uint = (v4 ^ v4)
                _ = v34
            } else {
                v4 -= (v4 & (^(v4)))
                v4 -= v4
                var v35 float64 = (0.1 / (-2.25 * ((float64(1e10) - 0.0) + 0.1)))
                _ = v35
            }
            var v36 float64 = (1.5 * 0.1)
            _ = v36
        }
    }
    fmt.Println(v4)
    delete(v1, 1)
    var v37 bool = true
    switch 127 {
    case 1:
        fmt.Println(v4)
    }
    v5 = append(v5, 0)
    v38 := GenS1{F0: 18446744073709551615}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
