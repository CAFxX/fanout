package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 string; F1 int16 }

func main() {
    var v1 uint = 18446744073709551615
    var v2 uint64 = 671
    v3 := []string{}
    var v4 float64 = (((0.0 / (3.14159 - (1e10 * float64(((3.14159 - (1.5 * (0.0 * ((0.1 + 3.14159) - 1.5)))) * (-2.25 / -2.25)))))) - -2.25) - (3.14159 * 1.5))
    if (false || false) {
        v5 := GenS1{F0: ((!(42 >= 0)) && ((true && (127 > 127)) || false)), F1: (v1 & v1)}
        _ = len(v3)
        _ = v5
    } else {
        var v6 string = "foo bar"
        switch v6 {
        case "a":
            v4 = 1.5
        case "\n\t":
            if len(v3) > 0 {
                v3[0] = "foo bar"
            }
        }
        var v7 uint = 255
        _ = v6
        _ = v7
    }
    if false {
        v1 = v1
    }
    for fi8 := 0; fi8 < 3; fi8++ {
        switch "foo bar" {
        case "a":
            for fi9 := 0; fi9 < 3; fi9++ {
                v1 -= v1
                var v10 uint32 = 932
                var v11 float64 = 0.1
                _ = fi9
                _ = v10
                _ = v11
            }
        case "foo bar":
            for fi12 := 0; fi12 < 10; fi12++ {
                v13 := GenS1{F0: false, F1: v1}
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi12
                _ = v13
            }
        case "":
            if ((!true) || (!((((float32(-2.25) - 0.0) >= float32(0.1)) || (uint(v2) <= v1)) && (63 == 100)))) {
                if true {
                    fmt.Println(v4, v4)
                } else {
                    var v14 int32 = 671
                    _ = v14
                }
            } else {
                for fi15 := 0; fi15 < 11; fi15++ {
                    switch fi8 {
                    case 127:
                        var v16 float32 = float32(v4)
                        _ = v16
                    case 32767:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    v4 -= 0.1
                    v4 -= -2.25
                    if ("" > "a") {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v17 float32 = 0.0
                        v17 = v17
                        _ = v17
                    }
                    v1 = (v1 / ((v1) | 1))
                    v1 = (v1 / ((v1) | 1))
                    if len(v3) > 0 {
                        v3[0] = ("\n\t" + "foo bar")
                    }
                    if ((v2 / ((0) | 1)) < 2) {
                        v2 += (^(v2))
                    }
                    _ = fi15
                }
                v1 = 0
                switch (-(fi8)) {
                case 42:
                    var v18 float64 = v4
                    _ = v18
                }
            }
        default:
            v2 -= v2
        }
        if (false || (42 < -2147483648)) {
            v2 = 2
        } else {
            var v19 int = fi8
            for fi20 := 0; fi20 < 6; fi20++ {
                v2 -= v2
                v21 := GenS1{F0: false, F1: v1}
                v1 -= (v1 % (((v1 + 1)) | 1))
                _ = fi20
                _ = v21
            }
            _ = v19
        }
        _ = fi8
    }
    _ = len(v3)
    v2 = v2
    var v22 uint = (v1 ^ 65535)
    v22 = 759
    for fi23 := 0; fi23 < 4; fi23++ {
        var v24 uint16 = 0
        if true {
            v1 = (v1 / ((v22) | 1))
            v24 -= v24
            if false {
                var v25 int64 = -244
                _ = v25
            }
        }
        if true {
            if true {
                var v26 uint64 = 42
                for fi27 := 0; fi27 < 19; fi27++ {
                    fmt.Println(v1, v2, v2)
                    if false {
                        v26 += v2
                    } else {
                        v3 = append(v3, "")
                        if (2147483648 != 2147483647) {
                            v4 += float64(0.0)
                            v28 := GenS1{F0: false, F1: 18446744073709551615}
                            _ = v28
                        } else {
                            v29 := GenS1{F0: true, F1: (v22 >> ((0) & 63))}
                            v30 := GenS2{F0: "\n\t", F1: 42}
                            _ = v29
                            _ = v30
                        }
                        switch "hello" {
                        case "foo bar":
                            fmt.Println(fi23)
                        }
                    }
                    fmt.Println(fi23, fi23)
                    var v31 int32 = -173
                    _ = fi27
                    _ = v31
                }
                var v32 int = 77
                _ = v26
                _ = v32
            }
        } else {
            v4 = -2.25
        }
        _ = fi23
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
