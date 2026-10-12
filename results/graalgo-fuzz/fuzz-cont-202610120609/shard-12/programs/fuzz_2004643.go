package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 uint8; F2 int32 }
type GenS3 struct { F0 int; F1 int8; F2 bool }
func genFn0(p0 int16, p1 int) uint64 {
    p1 = p1
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 uint = 255
    v2 := []int{1}
    var v3 float64 = (3.14159 - float64(1e10))
    v4 := []string{("\n\t" + ("" + "\n\t")), "", ("hello" + ("hello" + ((("hello" + "hello") + "foo bar") + "foo bar"))), "\n\t"}
    v5 := []int{}
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 int = fi6
        _ = fi6
        _ = v7
    }
    fmt.Println(v1, v1)
    if (!((1e10 / 0.1) <= -2.25)) {
        var v8 uint16 = 632
        fmt.Println(v8)
        _ = v8
    } else {
        if (!false) {
            if false {
                var v9 uint16 = 42
                v1 = v1
                _ = v9
            } else {
                v3 = ((v3 * 3.14159) + (1.5 + (v3 - ((((1e10 + ((float64(1e10) / v3) / 1.0)) / v3) + (3.14159 / v3)) + 1e10))))
                var v10 uint = (^(v1))
                v11 := GenS1{F0: 42}
                _ = v10
                _ = v11
            }
            v1 += v1
            v12 := GenS1{F0: 128}
            _ = v12
        } else {
            var v13 uint64 = 0
            for fi14 := 0; fi14 < 8; fi14++ {
                fmt.Println(v13, v1, v13)
                if (true || false) {
                    var v15 int = fi14
                    _ = v15
                }
                _ = fi14
            }
            switch 0 {
            case 32768:
                fmt.Println(v3, v3, v1)
            case 1:
                v1 = (v1 * v1)
            }
            _ = v13
        }
    }
    if false {
        var v16 string = (("a" + "a") + ("foo bar" + "foo bar"))
        _ = v16
    } else {
        v17 := GenS1{F0: 4294967295}
        if (2 <= 18446744073709551615) {
            v3 = float64(float32(v3))
            var v18 uint32 = 835
            _ = v18
        }
        _ = v17
    }
    v19 := GenS1{F0: 943}
    for fi20 := 0; fi20 < 6; fi20++ {
        var v21 uint = 65535
        if len(v5) > 0 {
            v5[0] = (fi20 & 4294967295)
        }
        _ = fi20
        _ = v21
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if (false || false) {
        _ = len(v5)
    }
    if (true || true) {
        var v22 int8 = 1
        if false {
            v23 := GenS2{F0: 1, F1: 127, F2: -2147483648}
            var v24 uint32 = 2
            fmt.Println(v1, v22, v3)
            _ = v23
            _ = v24
        }
        _ = v22
    } else {
        var v25 uint32 = 2147483647
        v25 = (^(v25))
        switch 361 {
        case 1:
            var v26 int16 = -377
            _ = v26
        case 249:
            fmt.Println(v25)
        case 0:
            v3 = 0.0
        default:
            var v27 uint8 = 1
            _ = v27
        }
        _ = v25
    }
    for fi28 := 0; fi28 < 15; fi28++ {
        switch (fi28 + (fi28 & (fi28 * (fi28 | 42)))) {
        case -12:
            if false {
                v29 := GenS3{F0: (fi28 % ((fi28) | 1)), F1: -49, F2: true}
                _ = v29
            } else {
                switch "foo bar" {
                case "a":
                    var v30 float64 = v3
                    _ = v30
                case "hello":
                    var v31 uint8 = 0
                    _ = v31
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
        case -173:
            for fi32 := 0; fi32 < 18; fi32++ {
                var v33 uint = (v1 & 4294967295)
                _ = fi32
                _ = v33
            }
        case -838:
            var v34 int32 = 2147483647
            _ = v34
        default:
            switch (fi28 ^ 4294967295) {
            case -1:
                v3 = v3
            case -9223372036854775808:
                if (fi28 == fi28) {
                    switch (-(fi28)) {
                    case 32768:
                        if (2147483647 == 1) {
                            v3 -= 3.14159
                        }
                    }
                    if len(v4) > 0 {
                        v4[0] = "a"
                    }
                } else {
                    var v35 uint64 = 18446744073709551615
                    _ = v35
                }
            }
        }
        var v36 float32 = 1e10
        _ = fi28
        _ = v36
    }
    var v37 int16 = -607
    switch 986 {
    case 1:
        v1 += v1
    }
    for fi38 := 0; fi38 < 2; fi38++ {
        fmt.Println(v1, fi38, v37)
        _ = fi38
    }
    v37 = v37
    var v39 int16 = (v37 * (v37 | v37))
    v40 := GenS2{F0: 9223372036854775807, F1: 2, F2: 33}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v37
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
