package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 int8 }

func main() {
    var v1 bool = false
    var v2 string = "\n\t"
    var v3 int8 = 100
    v4 := []int{-1}
    v5 := make(map[int]string)
    v6 := make(map[string]bool)
    v7 := GenS1{F0: 0, F1: (v3 | v3), F2: (v3 / ((0) | 1))}
    var v8 uint64 = 2
    var v9 uint64 = v8
    v3 = v3
    var v10 int = 42
    var v11 int = 32767
    switch (v11 * v10) {
    case 129:
        fmt.Println(v10)
    }
    if (4294967295 < 255) {
        v12 := GenS1{F0: -9223372036854775808, F1: v3, F2: int8(v11)}
        v13 := v5[v11]
        _ = v12
        _ = v13
    } else {
        if ((1.5 - (1.5 * (((1.5 / ((((1e10 + 0.0) - -2.25) - 1.5) - 1e10)) + -2.25) / (3.14159 / 1.0)))) == 1e10) {
            v1 = false
        }
        v11 = (v11 | 0)
        var v14 int16 = -32768
        _ = v14
    }
    switch (((v2 + ((v2 + v2) + "\n\t")) + "hello") + "foo bar") {
    case "\n\t":
        v1 = v1
    case "foo bar":
        v2 = v2
    default:
        v3 += v3
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch (v10 + (v11 % ((32767) | 1))) {
    case -823:
        v10 = 1
    case 32768:
        fmt.Println(v2, v1, v2)
    }
    switch v2 {
    case "a":
        v15 := v6["a"]
        _ = v15
    }
    if ((v2 + v2) < v2) {
        var v16 int8 = v3
        if v1 {
            _ = len(v4)
            for fi17 := 0; fi17 < 20; fi17++ {
                switch v11 {
                case 42:
                    switch v10 {
                    case 32768:
                        var v18 uint8 = 255
                        _ = v18
                    case 354:
                        if len(v4) > 0 {
                            v4[0] = 4294967295
                        }
                    case -258:
                        var v19 uint16 = 0
                        _ = v19
                    default:
                        if (!true) {
                            var v20 uint16 = 255
                            if v1 {
                                switch int(v3) {
                                case -987:
                                    switch fi17 {
                                    case 9223372036854775807:
                                        var v21 uint16 = (v20 - 887)
                                        _ = v21
                                    }
                                case -268:
                                    v1 = (true || v1)
                                case 9223372036854775807:
                                    var v22 uint8 = 42
                                    _ = v22
                                }
                                for fi23 := 0; fi23 < 20; fi23++ {
                                    var v24 float64 = 0.1
                                    if v1 {
                                        if (1 == 1) {
                                            var v25 uint16 = v20
                                            if (v3 != v3) {
                                                if ((v20 << ((1) & 63)) >= v25) {
                                                    v9 = v9
                                                    var v26 uint8 = 0
                                                    _ = v26
                                                } else {
                                                    switch (v11 << ((2) & 63)) {
                                                    case 127:
                                                        var v27 string = v2
                                                        _ = v27
                                                    default:
                                                        v25 = v25
                                                    }
                                                }
                                                v28 := GenS1{F0: 679, F1: 0, F2: v3}
                                                var v29 uint32 = 810
                                                _ = v28
                                                _ = v29
                                            } else {
                                                var v30 int32 = -618
                                                v9 -= (^(v8))
                                                _ = v30
                                            }
                                            _ = v25
                                        }
                                    } else {
                                        var v31 float64 = v24
                                        v10 = 1
                                        _ = v31
                                    }
                                    var v32 uint8 = 127
                                    _ = fi23
                                    _ = v24
                                    _ = v32
                                }
                            } else {
                                var v33 bool = (272 >= (^(v20)))
                                _ = v33
                            }
                            _ = v20
                        }
                    }
                default:
                    v34 := GenS1{F0: 1, F1: v16, F2: (v16 * 0)}
                    _ = v34
                }
                v4 = append(v4, v10)
                for fi35 := 0; fi35 < 17; fi35++ {
                    if (((3.14159 + 0.0) / (0.1 - (0.1 / (float64(-2.25) * 1e10)))) > (0.1 * 0.1)) {
                        v3 += v16
                        v9 = (v9 << ((0) & 63))
                        v36 := GenS1{F0: 0, F1: 98, F2: v16}
                        _ = v36
                    } else {
                        _ = len(v4)
                        var v37 string = (v2 + v2)
                        delete(v6, "a")
                        v38 := GenS1{F0: 9223372036854775807, F1: (v16 * (v3 + -93)), F2: v3}
                        _ = v37
                        _ = v38
                    }
                    v10 += fi17
                    _ = fi35
                }
                _ = fi17
            }
        }
        for fi39 := 0; fi39 < 16; fi39++ {
            if v1 {
                var v40 uint32 = 2
                for fi41 := 0; fi41 < 20; fi41++ {
                    switch v10 {
                    case 2147483647:
                        var v42 int8 = 127
                        _ = v42
                    }
                    _ = fi41
                }
                switch "" {
                case "hello":
                    fmt.Println(v16)
                default:
                    for fi43 := 0; fi43 < 4; fi43++ {
                        fmt.Println(v40)
                        _ = fi43
                    }
                }
                _ = v40
            }
            _ = fi39
        }
        _ = v16
    } else {
        for fi44 := 0; fi44 < 19; fi44++ {
            for fi45 := 0; fi45 < 5; fi45++ {
                fmt.Println(fi44, v11)
                _ = len(v6)
                _ = fi45
            }
            fmt.Println(v1)
            v8 = v9
            _ = fi44
        }
    }
    var v46 float32 = (0.1 + 1e10)
    for fi47 := 0; fi47 < 17; fi47++ {
        var v48 string = "a"
        _ = fi47
        _ = v48
    }
    var v49 float32 = v46
    v1 = false
    for fi50 := 0; fi50 < 4; fi50++ {
        var v51 uint32 = 2147483647
        v49 = v46
        _ = fi50
        _ = v51
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v46:", v46)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v49
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
