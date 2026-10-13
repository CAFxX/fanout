package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 uint32; F1 int16; F2 uint32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int }

func main() {
    var v1 bool = ((true && true) || ("foo bar" > ("\n\t" + "\n\t")))
    v2 := []uint32{491, 839}
    var v3 float32 = float32(0.1)
    fmt.Println(v1, v1)
    if v1 {
        switch -39 {
        case 32767:
            _ = len(v2)
        }
        if v1 {
            if v1 {
                for fi4 := 0; fi4 < 1; fi4++ {
                    var v5 uint8 = 1
                    _ = fi4
                    _ = v5
                }
            }
            var v6 uint64 = 4294967295
            _ = v6
        } else {
            v7 := GenS1{F0: 2, F1: 129}
            for fi8 := 0; fi8 < 17; fi8++ {
                v3 -= float32(((0.0 + 3.14159) / (1.5 + 3.14159)))
                v9 := GenS1{F0: 42, F1: 128}
                _ = fi8
                _ = v9
            }
            _ = v7
        }
        v3 = 3.14159
    }
    v3 += 3.14159
    v1 = v1
    switch "" {
    case "hello":
        switch 32768 {
        case 42:
            fmt.Println(v1, v3)
        default:
            var v10 float64 = 3.14159
            _ = v10
        }
    case "":
        if false {
            if (v1 || (((!v1) || (v1 && v1)) || (2147483648 > 2))) {
                for fi11 := 0; fi11 < 8; fi11++ {
                    if v1 {
                        var v12 float32 = v3
                        var v13 int64 = 42
                        _ = v12
                        _ = v13
                    }
                    v14 := GenS1{F0: 18446744073709551615, F1: 127}
                    _ = fi11
                    _ = v14
                }
                switch "a" {
                case "":
                    if (314 != -323) {
                        _ = len(v2)
                        v3 += (0.0 / 3.14159)
                    } else {
                        v15 := GenS2{F0: 652, F1: 1, F2: 258}
                        v3 = v3
                        _ = v15
                    }
                case "hello":
                    for fi16 := 0; fi16 < 15; fi16++ {
                        if v1 {
                            for fi17 := 0; fi17 < 17; fi17++ {
                                if v1 {
                                    if v1 {
                                        var v18 float32 = v3
                                        v18 = v18
                                        v19 := GenS1{F0: 65535, F1: 255}
                                        _ = v18
                                        _ = v19
                                    }
                                }
                                var v20 int16 = 753
                                var v21 uint = 65535
                                _ = fi17
                                _ = v20
                                _ = v21
                            }
                            if len(v2) > 0 {
                                v2[0] = 2147483647
                            }
                        }
                        _ = fi16
                    }
                }
            } else {
                var v22 int16 = 235
                v3 = 1.5
                for fi23 := 0; fi23 < 3; fi23++ {
                    var v24 float32 = v3
                    if (0 != 0) {
                        for fi25 := 0; fi25 < 3; fi25++ {
                            if v1 {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                if len(v2) > 0 {
                                    v2[0] = 65535
                                }
                            }
                            v24 = (v24 * -2.25)
                            _ = fi25
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v2 = append(v2, 2)
                    }
                    _ = fi23
                    _ = v24
                }
                _ = v22
            }
        }
    default:
        var v26 int64 = -1
        _ = v26
    }
    var v27 int32 = 42
    for fi28 := 0; fi28 < 15; fi28++ {
        for fi29 := 0; fi29 < 16; fi29++ {
            v3 += (((0.0 * v3) - (v3 - 0.1)) * ((((0.1 + v3) - (v3 - 3.14159)) - (v3 + v3)) - 3.14159))
            if v1 {
                var v30 uint32 = 65535
                var v31 uint16 = 1
                var v32 uint32 = v30
                _ = v30
                _ = v31
                _ = v32
            } else {
                var v33 uint64 = 18446744073709551615
                fmt.Println(fi29, v33, v33)
                _ = v33
            }
            if v1 {
                v1 = ((2147483647 == 42) && v1)
                switch ("hello" + "hello") {
                case "foo bar":
                    switch (^(fi29)) {
                    case -1:
                        v34 := GenS3{F0: 206}
                        _ = v34
                    case 127:
                        var v35 uint32 = 845
                        _ = v35
                    case 32767:
                        var v36 float64 = -2.25
                        _ = v36
                    }
                case "a":
                    fmt.Println(v3, v3)
                case "\n\t":
                    v27 = v27
                }
            }
            _ = fi29
        }
        v3 = float32(-2.25)
        _ = fi28
    }
    v2 = append(v2, 1)
    var v37 uint64 = 469
    v38 := GenS4{F0: -1}
    var v39 uint64 = (v37 - v37)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v37
    _ = v38
    _ = v39
}
