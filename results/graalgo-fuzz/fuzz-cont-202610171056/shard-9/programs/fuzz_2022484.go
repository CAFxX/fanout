package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 int = 129
    var v2 int32 = 0
    var v3 uint64 = 18446744073709551615
    fmt.Println(v3, v3, v3)
    var v4 int16 = -1
    v5 := GenS1{F0: 1}
    if false {
        if (false || (202 == 1)) {
            var v6 int8 = 3
            fmt.Println(v4, v2)
            _ = v6
        } else {
            switch (^(v1)) {
            case 0:
                for fi7 := 0; fi7 < 1; fi7++ {
                    v4 = 0
                    if true {
                        var v8 int8 = 102
                        _ = v8
                    } else {
                        for fi9 := 0; fi9 < 3; fi9++ {
                            fmt.Println(fi9, v2, fi9)
                            fmt.Println(fi7, v3, v1)
                            v10 := []int{(fi9 + v1)}
                            _ = fi9
                            _ = v10
                        }
                        var v11 float64 = ((1.5 + 0.0) + (float64((-2.25 + (3.14159 * 0.0))) - (0.0 + 1.5)))
                        _ = v11
                    }
                    _ = fi7
                }
            case -9223372036854775808:
                v12 := []int{}
                _ = v12
            case 1:
                v4 = v4
            }
            for fi13 := 0; fi13 < 12; fi13++ {
                switch "" {
                case "\n\t":
                    v14 := GenS1{F0: 42}
                    _ = v14
                default:
                    for fi15 := 0; fi15 < 1; fi15++ {
                        v3 = 4294967295
                        var v16 int16 = (v4 / (((v4 ^ v4)) | 1))
                        _ = fi15
                        _ = v16
                    }
                }
                switch v1 {
                case 32768:
                    v4 += v4
                default:
                    var v17 float32 = 3.14159
                    _ = v17
                }
                _ = fi13
            }
        }
        switch (("" + "hello") + "\n\t") {
        case "hello":
            var v18 uint8 = 2
            _ = v18
        case "":
            v19 := GenS1{F0: -89}
            _ = v19
        case "\n\t":
            for fi20 := 0; fi20 < 11; fi20++ {
                var v21 int16 = v4
                v21 = v21
                if false {
                    var v22 float32 = -2.25
                    v23 := GenS1{F0: 972}
                    _ = v22
                    _ = v23
                } else {
                    v1 = (-(fi20))
                    var v24 int8 = 1
                    _ = v24
                }
                _ = fi20
                _ = v21
            }
        default:
            var v25 string = "a"
            _ = v25
        }
        if false {
            var v26 int64 = 5
            _ = v26
        }
    }
    switch v1 {
    case 42:
        switch -566 {
        case 128:
            var v27 bool = ((!(100 <= 42)) || false)
            _ = v27
        default:
            var v28 uint = 18446744073709551615
            _ = v28
        }
    case 9223372036854775807:
        fmt.Println(v3, v4)
    default:
        v3 -= (^(v3))
    }
    v29 := []int{}
    v4 = (v4 % (((v4 << ((18446744073709551615) & 63))) | 1))
    v2 = v2
    var v30 int8 = -33
    if (v3 != v3) {
        var v31 string = ("foo bar" + "")
        fmt.Println(v2, v31, v4)
        _ = v31
    } else {
        v1 = (v1 | (^(v1)))
        fmt.Println(v1, v1)
        for fi32 := 0; fi32 < 16; fi32++ {
            v1 += 0
            v33 := GenS1{F0: 0}
            _ = fi32
            _ = v33
        }
    }
    v34 := GenS1{F0: 0}
    if true {
        var v35 int8 = (v30 - v30)
        _ = v35
    } else {
        switch v1 {
        case 42:
            for fi36 := 0; fi36 < 19; fi36++ {
                if len(v29) > 0 {
                    v29[0] = v1
                }
                _ = fi36
            }
        default:
            var v37 int8 = v30
            _ = v37
        }
    }
    for fi38 := 0; fi38 < 4; fi38++ {
        v3 = (^(v3))
        fmt.Println(v2)
        _ = fi38
    }
    v39 := GenS1{F0: 9223372036854775807}
    v3 -= v3
    fmt.Println(v3, v2)
    var v40 uint = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
