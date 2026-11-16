package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 uint32; F1 int16; F2 string }
type GenS4 struct { F0 uint8; F1 string }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint16; F1 int8 }
type GenS7 struct { F0 int64 }

func main() {
    v1 := make(map[string]bool)
    v2 := []uint32{383, 132}
    var v3 string = "a"
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v4 uint = 18446744073709551615
        if true {
            if (true || (1 >= 74)) {
                var v5 int32 = 2147483647
                for fi6 := 0; fi6 < 20; fi6++ {
                    var v7 uint64 = 686
                    v8 := GenS1{F0: (v7 * (v7 | (v7 / ((v7) | 1))))}
                    _ = fi6
                    _ = v7
                    _ = v8
                }
                v5 = v5
                _ = v5
            } else {
                fmt.Println(v3, v3)
                v1["foo bar"] = (-249 < 9223372036854775807)
                var v9 uint32 = 2147483647
                _ = v9
            }
        } else {
            v3 = v3
            if true {
                if false {
                    if false {
                        for fi10 := 0; fi10 < 3; fi10++ {
                            var v11 float32 = 0.1
                            _ = fi10
                            _ = v11
                        }
                    }
                }
            }
        }
        _ = v4
    }
    v3 = v3
    v12 := GenS2{F0: 32767, F1: 1}
    var v13 int64 = -1
    if true {
        var v14 uint32 = 735
        v14 -= 65535
        for fi15 := 0; fi15 < 20; fi15++ {
            fmt.Println(v13, v14)
            if (-2.25 != -2.25) {
                for fi16 := 0; fi16 < 16; fi16++ {
                    v17 := GenS1{F0: 228}
                    _ = fi16
                    _ = v17
                }
            }
            _ = fi15
        }
        _ = v14
    }
    var v18 uint16 = 1
    v19 := GenS2{F0: 565, F1: 127}
    for fi20 := 0; fi20 < 12; fi20++ {
        for fi21 := 0; fi21 < 16; fi21++ {
            v3 = v3
            if ((((-2.25 * ((3.14159 + (1e10 + (3.14159 - 1e10))) * 0.1)) / 1.0) / 0.1) < 0.1) {
                v22 := GenS2{F0: -627, F1: -1}
                _ = v22
            } else {
                if (v3 > v3) {
                    if (1 < 955) {
                        v2 = append(v2, 227)
                        v3 = v3
                        if false {
                            fmt.Println(v13, v13, fi21)
                            var v23 int8 = 94
                            switch v3 {
                            case "":
                                switch (v3 + (v3 + "hello")) {
                                case "\n\t":
                                    switch 0 {
                                    case 128:
                                        if false {
                                            v18 = v18
                                        } else {
                                            v3 = v3
                                            v24 := GenS3{F0: 2147483647, F1: -459, F2: "\n\t"}
                                            var v25 uint8 = 255
                                            _ = v24
                                            _ = v25
                                        }
                                    default:
                                        if false {
                                            v13 -= v13
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            var v26 bool = (v23 < (v23 * 1))
                                            v26 = false
                                            _ = v26
                                        }
                                    }
                                case "hello":
                                    var v27 uint8 = 255
                                    _ = v27
                                case "":
                                    var v28 float64 = -2.25
                                    _ = v28
                                }
                            }
                            if len(v2) > 0 {
                                v2[0] = uint32(v18)
                            }
                            _ = v23
                        } else {
                            if true {
                                if true {
                                    fmt.Println(v3)
                                    var v29 int64 = (-(v13))
                                    if ((-1 < 32767) && ((0 != 2) && true)) {
                                        v1[v3] = true
                                        v30 := GenS4{F0: 1, F1: (v3 + (v3 + v3))}
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        _ = v30
                                    }
                                    _ = v29
                                } else {
                                    var v31 int8 = 100
                                    v31 = (v31 & v31)
                                    delete(v1, ("foo bar" + ""))
                                    _ = v31
                                }
                            }
                        }
                        var v32 bool = false
                        fmt.Println(v32, v18)
                        _ = v32
                    } else {
                        fmt.Println(v3, fi20)
                        var v33 uint64 = 42
                        _ = v33
                    }
                    if ((42 != 42) && true) {
                        v13 += -1
                        if (false || true) {
                            v34 := GenS1{F0: 221}
                            fmt.Println(v3)
                            v35 := GenS5{F0: 957}
                            fmt.Println(v18)
                            _ = v34
                            _ = v35
                        } else {
                            v36 := GenS6{F0: 0, F1: -1}
                            var v37 int = 32767
                            _ = v36
                            _ = v37
                        }
                        delete(v1, "")
                    }
                }
                v13 += (-(v13))
                var v38 int32 = -2147483648
                _ = v38
            }
            _ = fi21
        }
        v18 = v18
        _ = fi20
    }
    _ = len(v2)
    _ = len(v2)
    var v39 float64 = (((1.5 - -2.25) * 1.5) / 3.14159)
    v1[v3] = (!false)
    var v40 bool = true
    v40 = v40
    switch v3 {
    case "foo bar":
        var v41 uint = 2
        _ = v41
    case "hello":
        v42 := GenS7{F0: v13}
        _ = v42
    case "":
        var v43 int = 42
        _ = v43
    default:
        v39 -= 0.0
    }
    v44 := GenS6{F0: 32767, F1: 110}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v39
    _ = v40
    _ = v44
}
