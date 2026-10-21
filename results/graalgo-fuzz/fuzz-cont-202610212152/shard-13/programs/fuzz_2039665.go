package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 uint32 = 42
    v2 := []uint32{v1, (v1 >> ((0) & 63))}
    v3 := make(map[string]bool)
    fmt.Println(v1)
    fmt.Println(v1)
    fmt.Println(v1)
    if (true || false) {
        var v4 uint8 = 255
        _ = v4
    }
    for fi5 := 0; fi5 < 20; fi5++ {
        switch (fi5 + 32767) {
        case -9223372036854775808:
            fmt.Println(v1, fi5)
        default:
            if true {
                var v6 uint64 = 42
                _ = v6
            } else {
                var v7 int16 = 1000
                var v8 bool = false
                switch "" {
                case "":
                    if v8 {
                        var v9 uint8 = 92
                        _ = v9
                    } else {
                        switch ("\n\t" + "") {
                        case "foo bar":
                            if false {
                                fmt.Println(v1, fi5, fi5)
                                fmt.Println(fi5, v1)
                            }
                        case "\n\t":
                            v10 := GenS1{F0: 653, F1: (^(v7))}
                            _ = v10
                        case "":
                            for fi11 := 0; fi11 < 2; fi11++ {
                                if v8 {
                                    var v12 uint64 = 42
                                    v7 = (v7 + v7)
                                    _ = v12
                                }
                                _ = fi11
                            }
                        default:
                            switch "" {
                            case "hello":
                                fmt.Println(v1, fi5, fi5)
                            case "\n\t":
                                var v13 int16 = v7
                                _ = v13
                            default:
                                if ((true || (-2147483648 < 2147483647)) && v8) {
                                    if v8 {
                                        v14 := v3["\n\t"]
                                        v14 = false
                                        _ = v14
                                    } else {
                                        var v15 float64 = (1e10 / 1.0)
                                        v1 += v1
                                        switch (("" + "hello") + "hello") {
                                        case "foo bar":
                                            var v16 float64 = 1e10
                                            _ = v16
                                        default:
                                            v17 := GenS1{F0: 42, F1: int16(fi5)}
                                            _ = v17
                                        }
                                        _ = v15
                                    }
                                    fmt.Println(fi5)
                                }
                            }
                        }
                        for fi18 := 0; fi18 < 5; fi18++ {
                            v7 = v7
                            v7 += v7
                            var v19 float32 = -2.25
                            _ = fi18
                            _ = v19
                        }
                    }
                case "hello":
                    v20 := GenS2{F0: 0}
                    _ = v20
                }
                _ = v7
                _ = v8
            }
        }
        var v21 uint = 253
        switch (fi5 % (((^(fi5))) | 1)) {
        case -9223372036854775808:
            v22 := GenS1{F0: uint64(v1), F1: 42}
            _ = v22
        default:
            if false {
                v1 -= (^(v1))
            } else {
                if (false && true) {
                    if false {
                        v23 := GenS2{F0: 2}
                        _ = v23
                    } else {
                        var v24 uint8 = 255
                        for fi25 := 0; fi25 < 8; fi25++ {
                            v26 := GenS3{F0: 878}
                            _ = fi25
                            _ = v26
                        }
                        _ = v24
                    }
                    var v27 uint16 = 42
                    _ = v27
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v1 = (v1 / ((2147483648) | 1))
            }
        }
        _ = fi5
        _ = v21
    }
    v28 := v3["a"]
    fmt.Println(v28, v28)
    if (false && (float64(3.14159) > 1e10)) {
        var v29 int16 = -453
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v28 = v28
        _ = v29
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v30 uint64 = 0
    v31 := GenS2{F0: 142}
    if false {
        switch "hello" {
        case "":
            v32 := GenS1{F0: 4294967295, F1: 741}
            _ = v32
        case "foo bar":
            v30 += v30
        case "\n\t":
            v33 := v3[("" + "")]
            _ = v33
        default:
            fmt.Println(v28)
        }
        _ = len(v2)
        v28 = v28
    }
    v30 = v30
    if false {
        v28 = v28
        fmt.Println(v28, v1, v30)
        var v34 int8 = 81
        _ = v34
    }
    for fi35 := 0; fi35 < 9; fi35++ {
        delete(v3, "\n\t")
        _ = fi35
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v31
}
