package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 bool }
type GenS2 struct { F0 uint32; F1 int64; F2 int8 }

func main() {
    v1 := []int{127}
    v2 := []uint32{4294967295, 2147483648, 446}
    v3 := []uint32{2, 65535, 1}
    v4 := make(map[string]string)
    var v5 uint8 = 2
    switch ("foo bar" + ("" + "hello")) {
    case "a":
        v6 := v4["hello"]
        _ = v6
    case "":
        v7 := GenS1{F0: 128, F1: 1, F2: (1.5 > 0.0)}
        _ = v7
    case "hello":
        if ((359 == 2) || false) {
            v8 := GenS1{F0: v5, F1: -78, F2: (127 != -586)}
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 1; fi9++ {
        switch -2147483648 {
        case 149:
            v5 = v5
        default:
            v5 += 184
        }
        _ = fi9
    }
    v5 = 2
    v5 = v5
    v10 := GenS1{F0: 67, F1: 1, F2: false}
    v5 = v5
    if (("foo bar" + "hello") <= "") {
        var v11 float32 = (3.14159 / (((0.0 * -2.25) + -2.25) - 1e10))
        _ = v11
    } else {
        v12 := GenS1{F0: (v5 & v5), F1: 0, F2: (-1 == -328)}
        _ = v12
    }
    var v13 uint32 = 42
    var v14 uint = 65535
    fmt.Println(v14)
    var v15 float64 = 1e10
    var v16 bool = true
    var v17 uint = v14
    var v18 int32 = 226
    v14 += (^(v14))
    switch "" {
    case "hello":
        if v16 {
            var v19 int64 = 0
            if v16 {
                _ = len(v3)
                for fi20 := 0; fi20 < 17; fi20++ {
                    for fi21 := 0; fi21 < 11; fi21++ {
                        if len(v2) > 0 {
                            v2[0] = v13
                        }
                        _ = len(v2)
                        _ = fi21
                    }
                    switch (("hello" + "a") + "\n\t") {
                    case "hello":
                        var v22 uint32 = (v13 ^ v13)
                        _ = v22
                    }
                    var v23 int8 = -1
                    _ = fi20
                    _ = v23
                }
            }
            fmt.Println(v13)
            _ = v19
        }
    case "":
        for fi24 := 0; fi24 < 1; fi24++ {
            for fi25 := 0; fi25 < 18; fi25++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v14 -= (v14 / ((v17) | 1))
                var v26 int8 = 1
                _ = fi25
                _ = v26
            }
            for fi27 := 0; fi27 < 8; fi27++ {
                if v16 {
                    v28 := GenS1{F0: (v5 % ((115) | 1)), F1: 719, F2: (65535 >= v17)}
                    _ = v28
                }
                var v29 int32 = (v18 * (v18 % ((v18) | 1)))
                v30 := GenS1{F0: (v5 | (v5 ^ (^(v5)))), F1: 978, F2: v16}
                _ = fi27
                _ = v29
                _ = v30
            }
            _ = fi24
        }
    case "foo bar":
        for fi31 := 0; fi31 < 18; fi31++ {
            if ("a" < "foo bar") {
                var v32 uint = uint(v5)
                v16 = (!v16)
                var v33 int8 = -1
                _ = v32
                _ = v33
            } else {
                if ((v5 << ((v17) & 63)) < 1) {
                    var v34 float64 = (1e10 * (v15 + float64(float32(float64(float32((1e10 * -2.25)))))))
                    var v35 uint32 = v13
                    _ = v34
                    _ = v35
                } else {
                    var v36 int8 = -128
                    _ = v36
                }
                if v16 {
                    var v37 float32 = 3.14159
                    _ = v37
                } else {
                    v38 := GenS1{F0: (v5 | (^(v5))), F1: (-(v18)), F2: (((v15 + float64(float32(float64(0.1)))) - (float64(0.1) / v15)) == 3.14159)}
                    v15 = v15
                    if len(v3) > 0 {
                        v3[0] = (v13 - 4294967295)
                    }
                    _ = v38
                }
            }
            _ = fi31
        }
    }
    if (v17 != (v17 & (^(v17)))) {
        var v39 int64 = -1
        _ = v39
    } else {
        v40 := GenS2{F0: v13, F1: 742, F2: -6}
        var v41 int16 = 42
        _ = v40
        _ = v41
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
