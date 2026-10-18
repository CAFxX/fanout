package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 uint }
type GenS2 struct { F0 uint8; F1 int32; F2 uint64 }

func main() {
    v1 := []string{}
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]string)
    v5 := GenS1{F0: 100, F1: 1, F2: 1}
    var v6 uint16 = 910
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi7 := 0; fi7 < 17; fi7++ {
        for fi8 := 0; fi8 < 15; fi8++ {
            var v9 int8 = 100
            if (!true) {
                v10 := GenS1{F0: 2, F1: v9, F2: 1}
                v11 := GenS1{F0: 91, F1: 0, F2: 516}
                v12 := GenS1{F0: 0, F1: v9, F2: 639}
                _ = v10
                _ = v11
                _ = v12
            }
            v13 := GenS1{F0: 128, F1: v9, F2: 2}
            _ = fi8
            _ = v13
            _ = v9
        }
        _ = fi7
    }
    if true {
        fmt.Println(v6)
        if len(v1) > 0 {
            v1[0] = ("a" + "")
        }
        v6 -= 42
    } else {
        var v14 uint64 = 18446744073709551615
        _ = v14
    }
    switch "a" {
    case "":
        v6 = 65535
    }
    var v15 string = ""
    var v16 uint = 760
    v17 := GenS1{F0: 100, F1: 0, F2: (v16 - v16)}
    fmt.Println(v16, v15)
    switch (v15 + ((v15 + v15) + v15)) {
    case "":
        if false {
            if (18446744073709551615 > 1) {
                var v18 float32 = (1.5 / 1.0)
                v19 := GenS1{F0: 1, F1: 42, F2: v16}
                v20 := GenS1{F0: 1, F1: -12, F2: (^(v16))}
                _ = v18
                _ = v19
                _ = v20
            } else {
                v21 := GenS1{F0: 255, F1: 1, F2: v16}
                _ = v21
            }
        } else {
            if false {
                for fi22 := 0; fi22 < 16; fi22++ {
                    var v23 string = ("foo bar" + ((v15 + v15) + v15))
                    switch (fi22 << ((v16) & 63)) {
                    case 249:
                        var v24 uint16 = (v6 - (^(v6)))
                        _ = v24
                    default:
                        if (v23 <= ("hello" + (v23 + ("\n\t" + (v15 + (v23 + v23)))))) {
                            for fi25 := 0; fi25 < 4; fi25++ {
                                v6 -= v6
                                var v26 int64 = -756
                                _ = fi25
                                _ = v26
                            }
                            var v27 int = fi22
                            _ = v27
                        } else {
                            v2["\n\t"] = 42
                            v28 := GenS2{F0: 1, F1: -509, F2: 171}
                            if false {
                                var v29 uint32 = 2147483647
                                v16 = v16
                                _ = v29
                            } else {
                                for fi30 := 0; fi30 < 7; fi30++ {
                                    var v31 uint8 = 128
                                    switch (-(fi22)) {
                                    case 4294967295:
                                        var v32 uint64 = 42
                                        _ = v32
                                    case 9223372036854775807:
                                        v33 := GenS1{F0: (^(v31)), F1: -128, F2: uint(v31)}
                                        _ = v33
                                    default:
                                        v34 := v2["a"]
                                        _ = v34
                                    }
                                    switch ("" + v15) {
                                    case "\n\t":
                                        _ = len(v1)
                                    case "":
                                        v31 = v31
                                    }
                                    var v35 uint16 = v6
                                    if len(v1) > 0 {
                                        v1[0] = ""
                                    }
                                    var v36 int = 9223372036854775807
                                    _ = fi30
                                    _ = v31
                                    _ = v35
                                    _ = v36
                                }
                            }
                            _ = v28
                        }
                    }
                    _ = fi22
                    _ = v23
                }
            }
            var v37 int16 = -705
            v37 += 1000
            _ = v37
        }
    default:
        if false {
            var v38 int = -1
            var v39 uint16 = v6
            _ = v38
            _ = v39
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
