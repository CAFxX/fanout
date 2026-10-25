package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint16 }
type GenS2 struct { F0 int8; F1 int }

func main() {
    var v1 uint8 = 128
    var v2 float32 = 1e10
    var v3 int16 = 566
    v4 := []int{}
    v5 := []int{0, 339, 128, 9223372036854775807}
    var v6 float32 = float32(1.5)
    var v7 uint = 42
    if false {
        var v8 int16 = int16(v1)
        switch 1 {
        case -9223372036854775808:
            v9 := GenS1{F0: 0, F1: v8, F2: 42}
            _ = v9
        }
        var v10 bool = false
        _ = v10
        _ = v8
    } else {
        if (42 >= -128) {
            if (true || false) {
                v11 := GenS2{F0: -128, F1: 127}
                var v12 uint64 = 4294967295
                var v13 int32 = 951
                _ = v11
                _ = v12
                _ = v13
            }
        }
    }
    v14 := GenS1{F0: -2147483648, F1: v3, F2: 65535}
    fmt.Println(v2, v3)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    switch 4294967295 {
    case 9223372036854775807:
        switch "\n\t" {
        case "a":
            fmt.Println(v3, v2, v3)
        case "foo bar":
            v1 += (v1 / (((v1 | uint8(v7))) | 1))
        case "hello":
            fmt.Println(v3, v7, v3)
        default:
            switch 1 {
            case -1:
                v3 -= v3
            case 9223372036854775807:
                for fi15 := 0; fi15 < 1; fi15++ {
                    v16 := GenS2{F0: -1, F1: (fi15 - (-(fi15)))}
                    _ = len(v4)
                    if ("" <= "foo bar") {
                        v17 := GenS2{F0: 100, F1: (fi15 ^ (-(fi15)))}
                        if (!(("\n\t" + ("a" + ("hello" + "hello"))) <= "")) {
                            fmt.Println(v3, fi15, v1)
                            var v18 int32 = 42
                            _ = v18
                        } else {
                            var v19 uint = v7
                            v20 := GenS1{F0: 42, F1: (v3 & v3), F2: 772}
                            _ = v19
                            _ = v20
                        }
                        _ = v17
                    } else {
                        var v21 float64 = 0.0
                        v6 += 0.1
                        fmt.Println(v6, v7, v7)
                        _ = v21
                    }
                    _ = fi15
                    _ = v16
                }
            }
        }
    }
    for fi22 := 0; fi22 < 7; fi22++ {
        switch (-(fi22)) {
        case 125:
            var v23 uint8 = uint8(v7)
            _ = v23
        }
        _ = fi22
    }
    for fi24 := 0; fi24 < 17; fi24++ {
        var v25 int = 127
        v26 := GenS2{F0: 127, F1: 0}
        var v27 float64 = 1.5
        _ = fi24
        _ = v25
        _ = v26
        _ = v27
    }
    _ = len(v4)
    var v28 uint64 = 0
    var v29 int16 = v3
    var v30 int16 = (v29 % ((v29) | 1))
    var v31 int64 = 1
    if (3.14159 < (float64((3.14159 + v2)) - 1e10)) {
        var v32 int8 = 127
        _ = v32
    } else {
        var v33 uint = 1
        v28 = v28
        _ = v33
    }
    if true {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        v31 += (-(v31))
        var v34 bool = (42 == -2147483648)
        if v34 {
            if len(v4) > 0 {
                v4[0] = 0
            }
            for fi35 := 0; fi35 < 1; fi35++ {
                for fi36 := 0; fi36 < 4; fi36++ {
                    var v37 uint16 = 1
                    var v38 bool = ((v1 ^ v1) <= (v1 * (v1 & v1)))
                    v37 += (v37 | 32767)
                    _ = fi36
                    _ = v37
                    _ = v38
                }
                _ = fi35
            }
            if len(v5) > 0 {
                v5[0] = 9223372036854775807
            }
        } else {
            v4 = append(v4, -2147483648)
        }
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
