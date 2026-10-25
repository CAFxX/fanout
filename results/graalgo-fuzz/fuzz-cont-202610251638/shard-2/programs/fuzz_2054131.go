package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 bool; F2 uint32 }
type GenS3 struct { F0 int16; F1 string; F2 int32 }
type GenS4 struct { F0 int8; F1 bool; F2 bool }
type GenS5 struct { F0 int; F1 string; F2 bool }
type GenS6 struct { F0 int16; F1 int }

func main() {
    v1 := []uint32{65535, 2147483648, 177, 42}
    var v2 float32 = (1.5 - 1.5)
    var v3 uint = 664
    var v4 bool = (true && (32767 < 988))
    v4 = v4
    v3 += 668
    _ = len(v1)
    _ = len(v1)
    if v4 {
        var v5 uint64 = 1
        _ = v5
    } else {
        fmt.Println(v2, v3, v2)
    }
    if (false || v4) {
        v6 := GenS1{F0: 9223372036854775807}
        for fi7 := 0; fi7 < 4; fi7++ {
            if (42 < 0) {
                for fi8 := 0; fi8 < 10; fi8++ {
                    if v4 {
                        if (true || ((42 < 2) || (!v4))) {
                            if v4 {
                                if v4 {
                                    switch (("\n\t" + ("\n\t" + "")) + "a") {
                                    case "foo bar":
                                        v9 := GenS2{F0: "hello", F1: (!((false || (-81 == 0)) && v4)), F2: 1}
                                        _ = v9
                                    case "hello":
                                        fmt.Println(v2, fi7)
                                    }
                                } else {
                                    var v10 int16 = -1
                                    _ = v10
                                }
                            } else {
                                if v4 {
                                    switch 579 {
                                    case -798:
                                        var v11 bool = v4
                                        _ = v11
                                    }
                                }
                                var v12 uint32 = 4294967295
                                if len(v1) > 0 {
                                    v1[0] = (^(v12))
                                }
                                _ = v12
                            }
                            v1 = append(v1, 2147483648)
                            if (true || (v4 && (3.14159 >= (1e10 * 1.5)))) {
                                v3 = 1
                            } else {
                                if v4 {
                                    switch fi7 {
                                    case 127:
                                        v3 = 4294967295
                                    case 134:
                                        var v13 int16 = 10
                                        if ((42 == 127) || false) {
                                            if true {
                                                var v14 bool = (false || v4)
                                                var v15 float64 = (1e10 - float64(3.14159))
                                                var v16 uint16 = 32767
                                                _ = v14
                                                _ = v15
                                                _ = v16
                                            } else {
                                                var v17 int32 = 2147483647
                                                var v18 int8 = -1
                                                var v19 uint16 = 2
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                v2 = ((v2 - v2) + 0.0)
                                                fmt.Println(v17, v13)
                                                switch ("" + ((("a" + "") + "") + ((("\n\t" + "foo bar") + ("" + "\n\t")) + "hello"))) {
                                                case "foo bar":
                                                    v3 += (v3 | v3)
                                                case "a":
                                                    v2 -= (float32((3.14159 * 0.0)) + 0.0)
                                                }
                                                _ = v17
                                                _ = v18
                                                _ = v19
                                            }
                                            v2 = float32(3.14159)
                                            v4 = (0 == 100)
                                            v13 = (-(v13))
                                            var v20 uint16 = 32767
                                            _ = v20
                                        }
                                        v13 -= -510
                                        _ = v13
                                    case 0:
                                        if true {
                                            v1 = append(v1, 2147483647)
                                        } else {
                                            var v21 uint16 = 1
                                            var v22 int = fi7
                                            _ = len(v1)
                                            v2 -= v2
                                            _ = v21
                                            _ = v22
                                        }
                                    default:
                                        if v4 {
                                            v2 = 0.0
                                            v23 := GenS3{F0: int16(v3), F1: "\n\t", F2: 692}
                                            _ = v23
                                        } else {
                                            var v24 uint32 = 717
                                            var v25 uint = 18446744073709551615
                                            _ = v24
                                            _ = v25
                                        }
                                    }
                                    v26 := GenS4{F0: -1, F1: (32767 >= -32768), F2: false}
                                    _ = v26
                                } else {
                                    _ = len(v1)
                                    if v4 {
                                        var v27 uint32 = 65535
                                        v1 = append(v1, v27)
                                        v28 := GenS5{F0: (^(fi7)), F1: ("foo bar" + ""), F2: v4}
                                        _ = v27
                                        _ = v28
                                    } else {
                                        if (!v4) {
                                            v29 := GenS2{F0: "", F1: false, F2: 0}
                                            var v30 bool = (!(935 < 2))
                                            _ = v29
                                            _ = v30
                                        }
                                    }
                                    if (1 == 2) {
                                        var v31 int64 = -1
                                        _ = v31
                                    } else {
                                        if len(v1) > 0 {
                                            v1[0] = uint32(v3)
                                        }
                                        v3 += 238
                                    }
                                }
                                if false {
                                    fmt.Println(v2)
                                    v1 = append(v1, 1)
                                    v32 := GenS1{F0: (fi7 ^ fi7)}
                                    _ = v32
                                }
                            }
                            fmt.Println(v3)
                        } else {
                            switch "foo bar" {
                            case "hello":
                                v4 = v4
                                v2 = v2
                                _ = len(v1)
                            default:
                                if v4 {
                                    var v33 int = fi7
                                    var v34 int64 = -1
                                    v35 := GenS5{F0: (fi8 << ((275) & 63)), F1: "a", F2: (1 != -63)}
                                    _ = v33
                                    _ = v34
                                    _ = v35
                                } else {
                                    if len(v1) > 0 {
                                        v1[0] = 2147483647
                                    }
                                    var v36 bool = (2 != 989)
                                    switch "foo bar" {
                                    case "a":
                                        if ((v4 && (true && v36)) && (v36 && (-1 < 42))) {
                                            fmt.Println(fi8, v3, fi8)
                                            switch (fi7 & (fi8 | (^(fi7)))) {
                                            case 4294967295:
                                                fmt.Println(fi7, fi8)
                                            }
                                            v4 = (!v4)
                                        } else {
                                            v36 = v36
                                            var v37 int = fi7
                                            fmt.Println(v3, fi8, v37)
                                            _ = v37
                                        }
                                    }
                                    _ = v36
                                }
                            }
                        }
                        _ = len(v1)
                        fmt.Println(fi7, fi8, v2)
                    }
                    v38 := GenS6{F0: -1, F1: fi7}
                    if (1e10 == 1.5) {
                        var v39 int32 = 44
                        var v40 float32 = v2
                        _ = v39
                        _ = v40
                    }
                    _ = fi8
                    _ = v38
                }
                var v41 int64 = -967
                v3 = (^(v3))
                _ = v41
            }
            for fi42 := 0; fi42 < 13; fi42++ {
                if len(v1) > 0 {
                    v1[0] = 616
                }
                _ = fi42
            }
            _ = fi7
        }
        _ = v6
    }
    v4 = true
    v2 = 1.5
    for fi43 := 0; fi43 < 9; fi43++ {
        v44 := GenS3{F0: -10, F1: "\n\t", F2: 0}
        _ = fi43
        _ = v44
    }
    if false {
        if (true && (!(("foo bar" + "a") != "a"))) {
            v3 = v3
            fmt.Println(v3, v2)
        }
    } else {
        for fi45 := 0; fi45 < 17; fi45++ {
            for fi46 := 0; fi46 < 4; fi46++ {
                v4 = true
                _ = fi46
            }
            var v47 float64 = 1.5
            _ = fi45
            _ = v47
        }
        v3 = v3
        var v48 float32 = ((float32(1.5) + v2) - (v2 / 1e10))
        _ = v48
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
