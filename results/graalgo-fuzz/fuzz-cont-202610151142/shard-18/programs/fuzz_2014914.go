package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int; F1 uint32; F2 string }
type GenS5 struct { F0 uint64 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint8 = 128
    v3 := []string{("" + "hello"), ("foo bar" + ("hello" + ("" + "a")))}
    var v4 uint32 = 2147483648
    var v5 bool = (-128 > 42)
    var v6 uint8 = 250
    delete(v1, v4)
    if v5 {
        if (-128 >= -67) {
            v7 := GenS1{F0: 4294967295, F1: (v6 - v6), F2: v5}
            switch -902 {
            case 32768:
                var v8 int16 = -1
                _ = v8
            case 380:
                if true {
                    var v9 int64 = 154
                    var v10 uint8 = v6
                    var v11 uint16 = 453
                    _ = v10
                    _ = v11
                    _ = v9
                }
            default:
                var v12 bool = false
                _ = v12
            }
            for fi13 := 0; fi13 < 17; fi13++ {
                v6 = v6
                var v14 float64 = float64(1.5)
                v1[v4] = 2147483648
                _ = fi13
                _ = v14
            }
            _ = v7
        } else {
            if false {
                for fi15 := 0; fi15 < 15; fi15++ {
                    v5 = v5
                    _ = fi15
                }
                var v16 int64 = 1
                _ = v16
            }
        }
        v5 = (0 != 32768)
    }
    v17 := GenS2{F0: 4294967295}
    fmt.Println(v4, v4)
    delete(v1, (v4 % ((v4) | 1)))
    v18 := GenS2{F0: 255}
    if v5 {
        if v5 {
            var v19 uint8 = (v6 * (v2 + (v6 & (^(v6)))))
            fmt.Println(v6, v4, v4)
            _ = v19
        } else {
            var v20 string = ""
            v21 := GenS1{F0: 0, F1: uint8(v4), F2: (!(v5 || (0 != -837)))}
            for fi22 := 0; fi22 < 19; fi22++ {
                switch v20 {
                case "a":
                    switch fi22 {
                    case 9223372036854775807:
                        v4 -= v4
                    default:
                        if v5 {
                            fmt.Println(v6, v20, v2)
                            var v23 float64 = 0.0
                            switch v20 {
                            case "foo bar":
                                v24 := v1[v4]
                                _ = v24
                            case "hello":
                                v4 = 1
                            case "\n\t":
                                for fi25 := 0; fi25 < 17; fi25++ {
                                    v1[v4] = v4
                                    _ = len(v3)
                                    _ = fi25
                                }
                            }
                            _ = v23
                        } else {
                            switch fi22 {
                            case 328:
                                if true {
                                    var v26 uint = 1
                                    _ = v26
                                }
                            default:
                                v27 := v1[863]
                                _ = v27
                            }
                        }
                    }
                case "foo bar":
                    switch fi22 {
                    case 128:
                        v6 += (v6 ^ v2)
                    }
                }
                var v28 int64 = -9223372036854775808
                v29 := GenS3{F0: 2}
                _ = fi22
                _ = v28
                _ = v29
            }
            _ = v20
            _ = v21
        }
        fmt.Println(v4, v5, v6)
    } else {
        var v30 int32 = -2147483648
        if v5 {
            var v31 string = "a"
            if len(v3) > 0 {
                v3[0] = v31
            }
            _ = v31
        }
        _ = v30
    }
    var v32 uint16 = 0
    var v33 int = 9223372036854775807
    for fi34 := 0; fi34 < 5; fi34++ {
        var v35 int16 = 804
        _ = fi34
        _ = v35
    }
    _ = len(v1)
    v36 := GenS4{F0: 42, F1: v4, F2: "hello"}
    if len(v3) > 0 {
        v3[0] = "a"
    }
    if (239 > 32767) {
        for fi37 := 0; fi37 < 14; fi37++ {
            fmt.Println(v33, v6, fi37)
            _ = fi37
        }
        v38 := GenS5{F0: 2}
        switch "" {
        case "hello":
            v39 := GenS5{F0: 1}
            _ = v39
        }
        _ = v38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
