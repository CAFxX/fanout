package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 uint64 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 string; F1 bool; F2 int8 }
type GenS4 struct { F0 int8; F1 int16; F2 bool }
type GenS5 struct { F0 bool; F1 uint; F2 int8 }
type GenS6 struct { F0 uint; F1 uint; F2 int }
func genFn0(p0 uint8, p1 bool, p2 uint16) int {
    var v6 string = ""
    var v7 uint8 = p0
    v8 := GenS1{F0: 2147483648, F1: -388, F2: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 32768
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]int)
    v3 := make(map[string]int)
    v4 := []string{("hello" + ("a" + "a"))}
    var v5 string = "\n\t"
    v5 = v5
    switch -1 {
    case 127:
        v5 = "hello"
    default:
        switch "" {
        case "foo bar":
            if true {
                fmt.Println(v5)
                for fi9 := 0; fi9 < 16; fi9++ {
                    for fi10 := 0; fi10 < 4; fi10++ {
                        var v11 uint8 = 2
                        if true {
                            v3[v5] = (fi9 % (((fi9 * fi10)) | 1))
                            fmt.Println(v5, v11)
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        } else {
                            v11 = v11
                        }
                        var v12 bool = true
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                    _ = fi9
                }
                var v13 uint64 = 18446744073709551615
                _ = v13
            }
        case "hello":
            for fi14 := 0; fi14 < 1; fi14++ {
                v5 = "\n\t"
                fmt.Println(v5, fi14)
                delete(v2, 4294967295)
                _ = fi14
            }
        default:
            for fi15 := 0; fi15 < 16; fi15++ {
                var v16 uint32 = 175
                fmt.Println(v16)
                _ = fi15
                _ = v16
            }
        }
    }
    v17 := GenS1{F0: 2147483647, F1: 1, F2: 467}
    switch -1 {
    case 207:
        v18 := v1[1]
        _ = v18
    case 413:
        switch 4294967295 {
        case -1:
            v19 := GenS2{F0: 552, F1: 1}
            _ = v19
        case 872:
            if false {
                for fi20 := 0; fi20 < 9; fi20++ {
                    if (!false) {
                        var v21 float64 = 3.14159
                        v21 = v21
                        _ = len(v4)
                        _ = v21
                    } else {
                        v5 = v5
                        v5 = (v5 + "foo bar")
                        var v22 uint8 = 1
                        _ = v22
                    }
                    switch (fi20 / ((genFn0(139, (true || (912 <= 9223372036854775807)), 208)) | 1)) {
                    case -179:
                        var v23 uint32 = 2147483647
                        _ = v23
                    default:
                        v5 = (v5 + v5)
                    }
                    _ = len(v4)
                    _ = fi20
                }
                fmt.Println(v5)
                switch v5 {
                case "":
                    v24 := GenS1{F0: 516, F1: 1000, F2: 121}
                    _ = v24
                case "foo bar":
                    v25 := GenS3{F0: v5, F1: (3.14159 >= (1e10 * 1e10)), F2: 42}
                    _ = v25
                default:
                    if len(v4) > 0 {
                        v4[0] = v5
                    }
                }
            } else {
                var v26 int16 = 32767
                _ = v26
            }
        default:
            for fi27 := 0; fi27 < 20; fi27++ {
                switch "foo bar" {
                case "a":
                    v28 := GenS4{F0: 42, F1: -1, F2: (v5 > "\n\t")}
                    _ = v28
                case "hello":
                    _ = len(v4)
                case "foo bar":
                    v5 = (v5 + ((v5 + "foo bar") + "hello"))
                }
                if false {
                    var v29 bool = false
                    v30 := GenS3{F0: v5, F1: ((!v29) || v29), F2: 127}
                    _ = v29
                    _ = v30
                }
                if len(v4) > 0 {
                    v4[0] = (v5 + v5)
                }
                _ = fi27
            }
        }
    default:
        v31 := GenS5{F0: true, F1: 339, F2: 127}
        _ = v31
    }
    var v32 uint32 = 835
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v33 := GenS3{F0: v5, F1: false, F2: 100}
    v34 := GenS6{F0: 940, F1: 42, F2: 2147483647}
    if len(v4) > 0 {
        v4[0] = ((v5 + v5) + ("" + "hello"))
    }
    if ((65535 < 65535) && (1e10 > (3.14159 / -2.25))) {
        v1[v32] = -611
    } else {
        if false {
            delete(v1, 460)
            fmt.Println(v5)
            for fi35 := 0; fi35 < 15; fi35++ {
                var v36 int = -1
                var v37 uint64 = 2
                for fi38 := 0; fi38 < 7; fi38++ {
                    var v39 float32 = float32(3.14159)
                    v40 := GenS6{F0: 1, F1: 18446744073709551615, F2: v36}
                    v36 = fi38
                    var v41 uint8 = 37
                    _ = fi38
                    _ = v39
                    _ = v40
                    _ = v41
                }
                _ = fi35
                _ = v36
                _ = v37
            }
        }
        v3[(v5 + v5)] = -791
        var v42 string = (v5 + v5)
        _ = v42
    }
    var v43 int32 = 676
    v2[v32] = 0
    v4 = append(v4, (v5 + "foo bar"))
    v43 += v43
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v43
    _ = v5
}
