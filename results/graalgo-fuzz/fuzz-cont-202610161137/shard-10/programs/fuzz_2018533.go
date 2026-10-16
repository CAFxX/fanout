package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 int; F1 int16; F2 int64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint8; F1 int64; F2 string }
func genFn0() int8 {
    var v6 int16 = -32768
    if ((-546 >= 9223372036854775807) || ((834 == 21) || (2147483647 < 1))) {
        if true {
            var v7 uint32 = 2147483647
            fmt.Println(v7, v6)
            _ = v7
        }
        for fi8 := 0; fi8 < 17; fi8++ {
            var v9 uint16 = 163
            fmt.Println(v9, fi8, v6)
            v6 = 0
            _ = fi8
            _ = v9
        }
    }
    v10 := GenS1{F0: 1}
    fmt.Println(v6)
    switch "a" {
    case "\n\t":
        var v11 uint = 255
        _ = v11
    }
    _ = v10
    _ = v6
    return 127
}
func genFn1(p0 string) string {
    var v12 uint16 = 8
    switch "a" {
    case "\n\t":
        v13 := GenS1{F0: 2}
        _ = v13
    }
    _ = p0
    _ = v12
    return "hello"
}

func main() {
    v1 := []int{}
    var v2 int8 = 47
    v3 := make(map[int]uint32)
    var v4 int64 = 42
    var v5 int8 = (v2 / ((v2) | 1))
    fmt.Println(v4)
    var v14 int64 = (v4 / ((v4) | 1))
    var v15 float32 = 3.14159
    v4 = v4
    v16 := GenS2{F0: "\n\t"}
    fmt.Println(v2, v15)
    if (583 <= 4294967295) {
        v17 := GenS3{F0: true, F1: 0}
        for fi18 := 0; fi18 < 7; fi18++ {
            v19 := GenS3{F0: (("a" + "a") != "foo bar"), F1: 50}
            _ = fi18
            _ = v19
        }
        _ = v17
    }
    v20 := GenS1{F0: 127}
    v3[127] = 2147483648
    v3[42] = 65535
    switch "\n\t" {
    case "foo bar":
        var v21 uint32 = 1
        _ = v21
    case "a":
        for fi22 := 0; fi22 < 20; fi22++ {
            if (!(2147483648 != 42)) {
                v23 := GenS4{F0: (fi22 >> ((18446744073709551615) & 63)), F1: -976, F2: v14}
                var v24 uint8 = 242
                _ = v23
                _ = v24
            }
            if false {
                var v25 float32 = float32(((1e10 - 1e10) - -2.25))
                v26 := GenS5{F0: 18446744073709551615}
                _ = v25
                _ = v26
            } else {
                fmt.Println(v2, v15)
                if false {
                    switch (fi22 >> ((42) & 63)) {
                    case 9223372036854775807:
                        var v27 uint16 = 0
                        _ = v27
                    case 42:
                        fmt.Println(fi22, v15, v4)
                    case 4294967295:
                        var v28 int16 = 42
                        _ = v28
                    default:
                        for fi29 := 0; fi29 < 15; fi29++ {
                            v30 := GenS2{F0: "a"}
                            _ = fi29
                            _ = v30
                        }
                    }
                } else {
                    switch ("" + "foo bar") {
                    case "foo bar":
                        var v31 int32 = 228
                        _ = v31
                    default:
                        var v32 int8 = 42
                        _ = v32
                    }
                    delete(v3, (^(fi22)))
                    var v33 int8 = (v5 * v2)
                    _ = v33
                }
            }
            _ = fi22
        }
    }
    if (-2147483648 < 1) {
        for fi34 := 0; fi34 < 7; fi34++ {
            v4 -= (v14 ^ v4)
            v35 := GenS4{F0: (^(fi34)), F1: -32768, F2: v14}
            if true {
                v2 += genFn0()
                switch (fi34 * -9223372036854775808) {
                case 9223372036854775807:
                    for fi36 := 0; fi36 < 7; fi36++ {
                        v14 += (v14 & v14)
                        v5 -= 127
                        var v37 int64 = v4
                        _ = fi36
                        _ = v37
                    }
                case 1:
                    for fi38 := 0; fi38 < 10; fi38++ {
                        v39 := GenS6{F0: 42, F1: v4, F2: "hello"}
                        _ = fi38
                        _ = v39
                    }
                }
                v2 = v2
            }
            _ = fi34
            _ = v35
        }
        var v40 uint64 = 42
        _ = v40
    } else {
        v2 += v2
    }
    var v41 float64 = 3.14159
    v2 -= (v5 ^ 23)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v41
    _ = v5
}
