package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 uint64; F1 int8; F2 uint64 }
type GenS5 struct { F0 uint32; F1 int16; F2 uint64 }

func main() {
    var v1 int8 = -1
    var v2 uint32 = 96
    var v3 float32 = float32(-2.25)
    var v4 uint32 = 2
    var v5 bool = ((1.5 < -2.25) || (9223372036854775807 <= 4294967295))
    v6 := []string{"a", "\n\t"}
    switch 9223372036854775807 {
    case 128:
        var v7 string = "\n\t"
        _ = v7
    case -1:
        v8 := GenS1{F0: 1}
        _ = v8
    case 32767:
        var v9 float32 = (float32(float64((v3 * float32(3.14159)))) / (v3 / (3.14159 * -2.25)))
        _ = v9
    default:
        v1 = (v1 | (^(v1)))
    }
    v10 := GenS1{F0: 921}
    var v11 int = 32768
    if v5 {
        if (("a" > "hello") || false) {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            if ((!v5) && false) {
                var v12 uint64 = 205
                _ = v12
            }
        } else {
            fmt.Println(v1)
            for fi13 := 0; fi13 < 3; fi13++ {
                v14 := GenS2{F0: int16(v11), F1: 0}
                _ = fi13
                _ = v14
            }
            _ = len(v6)
        }
    }
    _ = len(v6)
    switch "" {
    case "\n\t":
        var v15 bool = false
        _ = v15
    }
    if true {
        fmt.Println(v3, v5, v1)
        v5 = (!v5)
        v16 := GenS3{F0: 1, F1: v1}
        _ = v16
    }
    v17 := GenS2{F0: 483, F1: 42}
    var v18 int16 = -32768
    switch 0 {
    case 127:
        for fi19 := 0; fi19 < 20; fi19++ {
            v2 -= 85
            var v20 bool = true
            v21 := GenS4{F0: 0, F1: -28, F2: 1}
            _ = fi19
            _ = v20
            _ = v21
        }
    case 2147483647:
        if v5 {
            var v22 bool = ((v11 + (v11 >> ((65535) & 63))) >= (v11 | v11))
            var v23 uint16 = uint16(v2)
            switch (^(v11)) {
            case -1:
                var v24 uint32 = v4
                _ = v24
            case 42:
                var v25 int = (v11 / ((v11) | 1))
                _ = v25
            case 2147483647:
                _ = len(v6)
            default:
                var v26 string = ("\n\t" + (("\n\t" + "foo bar") + "a"))
                _ = v26
            }
            _ = v22
            _ = v23
        } else {
            v4 -= (v2 | (v2 % ((139) | 1)))
            _ = len(v6)
            var v27 int32 = 42
            _ = v27
        }
    }
    v28 := GenS1{F0: 108}
    v29 := GenS4{F0: 4294967295, F1: (v1 * 1), F2: 1}
    var v30 uint = 42
    var v31 bool = ((2147483648 >= 42) && false)
    if v5 {
        var v32 int64 = 9223372036854775807
        _ = v32
    }
    for fi33 := 0; fi33 < 15; fi33++ {
        v30 -= (v30 >> (((v30 & 255)) & 63))
        fmt.Println(v5, v31)
        for fi34 := 0; fi34 < 9; fi34++ {
            var v35 bool = true
            v36 := GenS1{F0: v30}
            if v31 {
                v18 -= -328
                v37 := GenS5{F0: 730, F1: 32767, F2: 18446744073709551615}
                _ = v37
            } else {
                if (false || (9223372036854775807 >= 9223372036854775807)) {
                    var v38 uint8 = 128
                    _ = v38
                } else {
                    if ((18446744073709551615 < (v30 + v30)) || v35) {
                        v35 = v31
                        v39 := GenS1{F0: v30}
                        _ = v39
                    } else {
                        var v40 float32 = (v3 + v3)
                        if (((!v31) && ((-(v1)) >= (v1 ^ v1))) && v5) {
                            v30 += (v30 ^ (v30 + v30))
                        }
                        _ = v40
                    }
                    if v31 {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        v41 := GenS5{F0: v2, F1: int16(v30), F2: 4294967295}
                        _ = v41
                    } else {
                        fmt.Println(v30, v35, v2)
                        if v35 {
                            v5 = (fi34 >= v11)
                            var v42 float64 = (0.1 + 0.1)
                            var v43 int16 = -32768
                            _ = v42
                            _ = v43
                        }
                    }
                    var v44 uint32 = v2
                    _ = v44
                }
            }
            _ = fi34
            _ = v35
            _ = v36
        }
        _ = fi33
    }
    var v45 uint16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v45
    _ = v5
    _ = v6
}
