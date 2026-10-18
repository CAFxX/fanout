package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 uint8; F1 int8; F2 int64 }
type GenS5 struct { F0 int16; F1 string }
type GenS6 struct { F0 int64; F1 int; F2 uint }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int8 = 0
    v3 := []string{}
    v4 := make(map[int]uint32)
    v5 := make(map[string]bool)
    v6 := GenS1{F0: v2}
    v2 += (-(v2))
    var v7 uint32 = 151
    switch "" {
    case "foo bar":
        switch 9223372036854775807 {
        case -1:
            var v8 uint64 = 197
            _ = v8
        case 9223372036854775807:
            if true {
                var v9 int64 = 42
                _ = v9
            } else {
                var v10 int16 = 1000
                if true {
                    v11 := GenS1{F0: (v2 >> ((65535) & 63))}
                    _ = v11
                } else {
                    delete(v4, 702)
                    v12 := v4[702]
                    _ = v12
                }
                if true {
                    for fi13 := 0; fi13 < 11; fi13++ {
                        v5["\n\t"] = (713 <= 4294967295)
                        var v14 uint8 = 128
                        _ = fi13
                        _ = v14
                    }
                    v2 = 31
                } else {
                    _ = len(v3)
                }
                _ = v10
            }
        default:
            if false {
                v2 = (v2 | 1)
            }
        }
    case "":
        v7 -= (v7 ^ (^(v7)))
    case "hello":
        var v15 float64 = 0.1
        _ = v15
    default:
        var v16 int = -2147483648
        _ = v16
    }
    var v17 int16 = 0
    _ = len(v4)
    switch "" {
    case "foo bar":
        var v18 uint64 = 968
        _ = v18
    case "hello":
        if false {
            fmt.Println(v17, v17, v7)
        }
    case "\n\t":
        v19 := GenS1{F0: (-(v2))}
        _ = v19
    default:
        if true {
            v20 := GenS1{F0: (v2 - (v2 & (v2 * -1)))}
            v5["foo bar"] = (true || true)
            _ = v20
        }
    }
    var v21 int8 = v2
    v21 = v2
    v17 -= 42
    for fi22 := 0; fi22 < 14; fi22++ {
        v23 := GenS1{F0: 0}
        v7 += v7
        if ((188 < 32767) || (18446744073709551615 != 295)) {
            for fi24 := 0; fi24 < 17; fi24++ {
                if ((-2.25 + -2.25) > ((1.5 - (3.14159 + 1e10)) / (1.5 / (0.0 - (0.0 + -2.25))))) {
                    if false {
                        if ((-812 > 358) || false) {
                            v25 := GenS1{F0: (v21 << ((1) & 63))}
                            v2 = v2
                            if (!true) {
                                v2 += v21
                                v26 := GenS2{F0: 127}
                                _ = v26
                            }
                            _ = v25
                        }
                        v17 += v17
                        v17 = 28
                    }
                    var v27 uint = 1
                    switch "\n\t" {
                    case "hello":
                        v2 -= (v21 - -50)
                    case "":
                        if (906 == 42) {
                            switch "" {
                            case "hello":
                                v28 := GenS3{F0: fi24, F1: (v7 / (((v7 * (^(v7)))) | 1))}
                                _ = v28
                            default:
                                var v29 int64 = -1
                                _ = v29
                            }
                            delete(v1, (v7 - 2))
                        }
                    case "a":
                        v3 = append(v3, "")
                    }
                    _ = v27
                } else {
                    switch "a" {
                    case "\n\t":
                        fmt.Println(fi22)
                    case "":
                        if (!("" > "")) {
                            fmt.Println(v21, v17, v7)
                            v21 = (v21 << ((4294967295) & 63))
                            fmt.Println(v17)
                        }
                    case "foo bar":
                        var v30 int16 = -32768
                        _ = v30
                    default:
                        v21 -= 1
                    }
                }
                if ("foo bar" >= "") {
                    var v31 bool = false
                    v17 += v17
                    switch "\n\t" {
                    case "foo bar":
                        if v31 {
                            var v32 int = fi22
                            _ = v32
                        }
                    }
                    var v33 float32 = 1.5
                    _ = v31
                    _ = v33
                } else {
                    v2 -= v2
                }
                _ = fi24
            }
            var v34 int8 = v21
            _ = v34
        } else {
            var v35 float64 = -2.25
            fmt.Println(v21, v35)
            _ = v35
        }
        _ = fi22
        _ = v23
    }
    for fi36 := 0; fi36 < 18; fi36++ {
        if false {
            v37 := GenS4{F0: 2, F1: 30, F2: -1}
            switch "" {
            case "":
                v38 := GenS2{F0: 2}
                _ = v38
            case "hello":
                var v39 uint16 = 784
                _ = v39
            default:
                for fi40 := 0; fi40 < 8; fi40++ {
                    v21 -= 1
                    v41 := GenS4{F0: 42, F1: v2, F2: 1}
                    switch (fi40 / (((fi36 + -1)) | 1)) {
                    case 128:
                        fmt.Println(fi40)
                    case -1:
                        if false {
                            if false {
                                fmt.Println(v17, v21)
                                v42 := GenS4{F0: 0, F1: (v21 << ((uint(v7)) & 63)), F2: -328}
                                _ = v42
                            }
                            var v43 int64 = -817
                            v43 -= (-(v43))
                            _ = v43
                        } else {
                            if len(v3) > 0 {
                                v3[0] = "a"
                            }
                        }
                    default:
                        var v44 float64 = (1.5 + 0.0)
                        _ = v44
                    }
                    _ = fi40
                    _ = v41
                }
            }
            fmt.Println(fi36)
            _ = v37
        } else {
            fmt.Println(v2, v17)
            v45 := GenS2{F0: 255}
            _ = v45
        }
        _ = fi36
    }
    fmt.Println(v21, v17, v2)
    switch 1 {
    case -2147483648:
        if true {
            for fi46 := 0; fi46 < 1; fi46++ {
                var v47 uint = 65535
                v2 = (v21 >> ((v47) & 63))
                v48 := GenS5{F0: (v17 - (v17 | 986)), F1: "\n\t"}
                _ = fi46
                _ = v47
                _ = v48
            }
            v1[2147483647] = v7
            if ("a" != "") {
                var v49 uint8 = 143
                _ = v49
            } else {
                for fi50 := 0; fi50 < 2; fi50++ {
                    fmt.Println(v2)
                    _ = fi50
                }
                if (309 != 42) {
                    v7 = v7
                    v51 := GenS6{F0: -866, F1: 0, F2: 255}
                    _ = v51
                }
                v7 = v7
            }
        } else {
            var v52 int = 938
            _ = v52
        }
    case 737:
        _ = len(v1)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
