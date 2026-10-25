package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }

func main() {
    v1 := []string{"foo bar", ("hello" + "hello")}
    var v2 string = ""
    var v3 string = "foo bar"
    var v4 int8 = 0
    var v5 string = "a"
    var v6 int = -157
    for fi7 := 0; fi7 < 8; fi7++ {
        for fi8 := 0; fi8 < 3; fi8++ {
            var v9 uint = 4294967295
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    v10 := GenS1{F0: 1, F1: (v2 + v5)}
    for fi11 := 0; fi11 < 20; fi11++ {
        var v12 float64 = (((0.1 - float64((0.1 + 0.0))) / (0.1 - -2.25)) + ((1e10 - 1e10) * (1e10 * 1e10)))
        if (42 < -422) {
            var v13 bool = ((3.14159 - (v12 - 1e10)) > (v12 * v12))
            _ = v13
        } else {
            v5 = (v3 + v2)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v14 float64 = 0.0
            _ = v14
        }
        if (false || (557 != 446)) {
            if true {
                if len(v1) > 0 {
                    v1[0] = "a"
                }
            } else {
                var v15 float32 = 1e10
                var v16 int64 = 149
                _ = v15
                _ = v16
            }
            var v17 float32 = 0.0
            _ = v17
        } else {
            fmt.Println(v6, v2, v5)
            for fi18 := 0; fi18 < 13; fi18++ {
                if true {
                    if true {
                        var v19 float32 = -2.25
                        _ = v19
                    } else {
                        var v20 uint16 = 32767
                        if (true && false) {
                            var v21 string = v3
                            _ = v21
                        } else {
                            var v22 int64 = 0
                            v2 = v5
                            _ = v22
                        }
                        _ = v20
                    }
                    switch fi18 {
                    case 32768:
                        if (((3.14159 + float32(1.5)) + float32((float64(3.14159) / (1.5 / (((0.1 / 1.0) * v12) + (1e10 * ((float64(1.5) - v12) + -2.25))))))) != 0.0) {
                            var v23 uint = 255
                            _ = v23
                        }
                    }
                } else {
                    fmt.Println(v2, v3, v4)
                }
                switch ("hello" + ((v3 + (v5 + v3)) + ("hello" + v5))) {
                case "foo bar":
                    var v24 float64 = v12
                    if len(v1) > 0 {
                        v1[0] = v5
                    }
                    v25 := GenS1{F0: 2147483647, F1: v2}
                    _ = v24
                    _ = v25
                default:
                    v2 = v2
                }
                _ = fi18
            }
        }
        _ = fi11
        _ = v12
    }
    var v26 int8 = v4
    v1 = append(v1, v3)
    var v27 uint64 = 616
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch (v6 + (v6 & (v6 >> ((18446744073709551615) & 63)))) {
    case 129:
        var v28 int = -205
        _ = v28
    default:
        for fi29 := 0; fi29 < 13; fi29++ {
            v30 := GenS1{F0: 484, F1: v3}
            fmt.Println(v5, v5)
            v1 = append(v1, "")
            _ = fi29
            _ = v30
        }
    }
    v3 = "\n\t"
    var v31 uint16 = 65535
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v32 float32 = (3.14159 * ((3.14159 * -2.25) - -2.25))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
