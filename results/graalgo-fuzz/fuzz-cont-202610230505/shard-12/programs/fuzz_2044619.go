package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0() int32 {
    v5 := []int{}
    var v6 bool = true
    fmt.Println(v6)
    if len(v5) > 0 {
        v5[0] = -650
    }
    for fi7 := 0; fi7 < 3; fi7++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi7
    }
    _ = v5
    _ = v6
    return 1
}

func main() {
    v1 := make(map[int]bool)
    var v2 string = "a"
    v3 := []uint32{318, 2147483648, 575, 546}
    v4 := []uint32{}
    v8 := GenS1{F0: 0}
    v2 = (v2 + ("foo bar" + (v2 + ((v2 + v2) + v2))))
    switch v2 {
    case "hello":
        var v9 float64 = 1e10
        _ = v9
    case "\n\t":
        var v10 int8 = 100
        _ = v10
    case "a":
        if ((18446744073709551615 == 18446744073709551615) && true) {
            if (2 != 32767) {
                v2 = v2
                if ((0.0 - 3.14159) < 1.5) {
                    var v11 float32 = (0.1 + ((0.1 / 1.0) / 1.0))
                    for fi12 := 0; fi12 < 16; fi12++ {
                        v11 = float32((0.1 - 0.1))
                        _ = fi12
                    }
                    _ = v11
                } else {
                    if (((((((0.0 * (3.14159 / (((3.14159 - (3.14159 - 1e10)) * (1e10 * 0.1)) + (3.14159 * 1e10)))) - float32(-2.25)) / float32(float64(-2.25))) / 1.0) + (1.5 / 1e10)) > 0.0) || (true || false)) {
                        v2 = v2
                    } else {
                        v13 := GenS1{F0: 24}
                        var v14 string = "a"
                        fmt.Println(v14, v14)
                        _ = v13
                        _ = v14
                    }
                }
            } else {
                v2 = "\n\t"
                v2 = v2
            }
            var v15 uint64 = 215
            var v16 uint32 = 0
            _ = v15
            _ = v16
        } else {
            v17 := GenS1{F0: 1}
            _ = len(v1)
            _ = v17
        }
    }
    v1[2147483647] = true
    v18 := GenS1{F0: 1}
    if true {
        var v19 uint = 18446744073709551615
        for fi20 := 0; fi20 < 12; fi20++ {
            var v21 float64 = (1.5 - -2.25)
            _ = fi20
            _ = v21
        }
        _ = v19
    }
    switch -9223372036854775808 {
    case 32767:
        fmt.Println(v2)
    }
    if true {
        v22 := GenS1{F0: 42}
        _ = v22
    } else {
        switch 42 {
        case -35:
            if ((48 > 127) && false) {
                v2 = "hello"
                for fi23 := 0; fi23 < 19; fi23++ {
                    fmt.Println(fi23, fi23)
                    v2 = (("foo bar" + ((("a" + (("foo bar" + (v2 + ((v2 + v2) + "foo bar"))) + v2)) + v2) + (v2 + v2))) + v2)
                    _ = fi23
                }
                var v24 int = 128
                _ = v24
            } else {
                var v25 int8 = 127
                _ = v25
            }
        case 1:
            fmt.Println(v2)
        }
        var v26 float64 = (float64((3.14159 / 1.0)) / 1.0)
        _ = v26
    }
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
