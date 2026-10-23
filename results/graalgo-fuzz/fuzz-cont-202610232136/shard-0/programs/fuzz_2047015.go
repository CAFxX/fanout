package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 int32 }
type GenS2 struct { F0 int; F1 int64; F2 int64 }

func main() {
    var v1 int8 = -2
    v2 := make(map[string]string)
    v3 := []string{}
    var v4 int32 = -2147483648
    v5 := []uint32{2147483647, 42, 2, 284}
    v6 := []uint32{0}
    if (!false) {
        for fi7 := 0; fi7 < 6; fi7++ {
            fmt.Println(v4, v1)
            v8 := GenS1{F0: (!(3.14159 == (0.1 + 0.1))), F1: 42, F2: (v4 / (((v4 | (v4 * (^(v4))))) | 1))}
            _ = fi7
            _ = v8
        }
        v2["foo bar"] = "\n\t"
    }
    var v9 string = "hello"
    var v10 uint64 = 42
    for fi11 := 0; fi11 < 10; fi11++ {
        if len(v6) > 0 {
            v6[0] = 42
        }
        switch v9 {
        case "\n\t":
            fmt.Println(v4, v10)
        default:
            fmt.Println(fi11, v4)
        }
        _ = fi11
    }
    var v12 string = ""
    if true {
        var v13 int64 = 42
        for fi14 := 0; fi14 < 18; fi14++ {
            v15 := v2["\n\t"]
            var v16 uint32 = 42
            _ = fi14
            _ = v15
            _ = v16
        }
        var v17 int64 = v13
        _ = v13
        _ = v17
    } else {
        v18 := v2["\n\t"]
        fmt.Println(v10)
        _ = v18
    }
    v9 = ("foo bar" + (v9 + "\n\t"))
    if (82 != 196) {
        v6 = append(v6, 0)
        v9 = "\n\t"
    } else {
        if true {
            v4 += v4
        } else {
            for fi19 := 0; fi19 < 2; fi19++ {
                for fi20 := 0; fi20 < 19; fi20++ {
                    v9 = v12
                    var v21 string = ("hello" + v9)
                    if true {
                        if (true && (false && false)) {
                            v1 += (v1 - v1)
                        } else {
                            fmt.Println(fi19)
                            v22 := GenS1{F0: true, F1: (v10 + 0), F2: v4}
                            var v23 bool = false
                            _ = v22
                            _ = v23
                        }
                        if true {
                            v21 = ((("" + (v21 + v9)) + v12) + v9)
                        }
                    } else {
                        var v24 float64 = ((float64((0.0 + 0.0)) + (1e10 / (1.5 * (((0.0 / 1e10) * 0.0) + (float64(float32(1e10)) - 3.14159))))) + ((float64(1e10) - 3.14159) + -2.25))
                        _ = v24
                    }
                    _ = fi20
                    _ = v21
                }
                switch v9 {
                case "a":
                    v25 := GenS2{F0: fi19, F1: 1, F2: 0}
                    _ = v25
                case "":
                    v26 := GenS1{F0: (!(42 != 1)), F1: (v10 / (((v10 * v10)) | 1)), F2: (v4 * v4)}
                    _ = v26
                case "foo bar":
                    var v27 uint32 = 65535
                    _ = v27
                }
                _ = fi19
            }
            v28 := GenS1{F0: (399 < 4294967295), F1: v10, F2: 2147483647}
            for fi29 := 0; fi29 < 2; fi29++ {
                var v30 int = fi29
                _ = len(v2)
                _ = fi29
                _ = v30
            }
            _ = v28
        }
        v31 := GenS1{F0: false, F1: 4294967295, F2: (^(v4))}
        switch "hello" {
        case "a":
            var v32 int32 = 1
            _ = v32
        case "foo bar":
            fmt.Println(v12, v4)
        }
        _ = v31
    }
    delete(v2, v12)
    var v33 int16 = 1000
    v10 = (v10 >> ((961) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
