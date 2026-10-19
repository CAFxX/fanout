package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 uint; F1 int8; F2 uint }
type GenS5 struct { F0 uint16; F1 uint; F2 string }
type GenS6 struct { F0 string; F1 string; F2 int8 }
func genFn0(p0 uint64, p1 bool) int {
    p0 -= 42
    if false {
        var v5 string = ("foo bar" + "")
        if (!p1) {
            v6 := GenS1{F0: 679, F1: 1, F2: -503}
            var v7 int8 = -128
            for fi8 := 0; fi8 < 17; fi8++ {
                v9 := GenS2{F0: -1}
                var v10 float32 = ((((1e10 / (1.5 - 0.1)) - 1e10) - -2.25) / -2.25)
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = v6
            _ = v7
        } else {
            for fi11 := 0; fi11 < 10; fi11++ {
                p1 = p1
                p1 = p1
                var v12 int16 = 1000
                _ = fi11
                _ = v12
            }
        }
        var v13 bool = ((("foo bar" <= (v5 + (((v5 + v5) + v5) + "foo bar"))) || (469 >= 65535)) || false)
        _ = v13
        _ = v5
    } else {
        var v14 uint16 = 1
        switch -2147483648 {
        case -9223372036854775808:
            for fi15 := 0; fi15 < 12; fi15++ {
                for fi16 := 0; fi16 < 13; fi16++ {
                    v17 := GenS3{F0: p0, F1: "\n\t"}
                    switch (((("hello" + "\n\t") + "foo bar") + ("\n\t" + (("a" + "\n\t") + ((("" + "\n\t") + "foo bar") + ("a" + ("a" + ("foo bar" + "a"))))))) + ("\n\t" + "hello")) {
                    case "":
                        fmt.Println(p0, p0)
                    case "hello":
                        var v18 int64 = -1
                        _ = v18
                    }
                    v19 := GenS3{F0: (p0 - p0), F1: ""}
                    _ = fi16
                    _ = v17
                    _ = v19
                }
                p1 = true
                _ = fi15
            }
        }
        _ = v14
    }
    if p1 {
        if (((float32((3.14159 + -2.25)) * 0.1) <= 1e10) || (0.0 > 0.0)) {
            var v20 int64 = 108
            _ = v20
        } else {
            v21 := []int{4294967295}
            v22 := GenS2{F0: -32768}
            if (false || p1) {
                var v23 string = "a"
                _ = v23
            }
            _ = v21
            _ = v22
        }
    } else {
        fmt.Println(p1, p1)
    }
    v24 := []int{32768, 2147483647, int(p0), 0}
    p0 += 214
    _ = p0
    _ = p1
    _ = v24
    return 32768
}

func main() {
    v1 := make(map[int]string)
    var v2 int = 0
    var v3 uint = 2
    v4 := make(map[uint32]string)
    var v25 uint32 = 4294967295
    if true {
        switch v2 {
        case 32767:
            fmt.Println(v2)
        case 2147483647:
            v26 := GenS4{F0: (v3 | 221), F1: -20, F2: v3}
            _ = v26
        default:
            var v27 uint8 = 122
            _ = v27
        }
        v28 := GenS5{F0: 255, F1: 42, F2: ("a" + ("\n\t" + "foo bar"))}
        var v29 float64 = ((-2.25 / 1.0) * 1e10)
        _ = v28
        _ = v29
    }
    v2 = v2
    v30 := []int{1, v2, v2, 32767}
    v31 := v4[v25]
    v2 = (v2 ^ -430)
    if false {
        for fi32 := 0; fi32 < 6; fi32++ {
            for fi33 := 0; fi33 < 20; fi33++ {
                var v34 bool = false
                _ = len(v30)
                _ = fi33
                _ = v34
            }
            _ = fi32
        }
    } else {
        v35 := GenS6{F0: (v31 + ("foo bar" + "a")), F1: (v31 + "foo bar"), F2: -1}
        _ = v35
    }
    var v36 uint8 = 128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v30
    _ = v31
    _ = v36
    _ = v4
}
