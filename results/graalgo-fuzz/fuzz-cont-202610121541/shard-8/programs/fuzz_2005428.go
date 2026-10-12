package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8 }
func genFn0(p0 int) string {
    var v6 int8 = 42
    if (0 >= 18446744073709551615) {
        var v7 float64 = -2.25
        v8 := []int{p0, p0, (p0 * (-(p0)))}
        _ = v7
        _ = v8
    }
    _ = p0
    _ = v6
    return ("foo bar" + ((("a" + "hello") + "hello") + "\n\t"))
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int64 = 344
    var v3 int8 = 42
    var v4 float32 = (float32(-2.25) * -2.25)
    v5 := make(map[int]int)
    v2 = 9223372036854775807
    v9 := GenS1{F0: 684}
    v4 = v4
    if true {
        var v10 float32 = v4
        _ = v10
    }
    var v11 uint64 = 4294967295
    if false {
        v3 += v3
        var v12 float64 = 0.0
        _ = v12
    } else {
        var v13 int = 127
        _ = v13
    }
    v14 := v1[0]
    if true {
        for fi15 := 0; fi15 < 8; fi15++ {
            if false {
                fmt.Println(v2)
                if false {
                    for fi16 := 0; fi16 < 15; fi16++ {
                        switch "a" {
                        case "\n\t":
                            switch ("" + (("a" + ("foo bar" + "")) + "a")) {
                            case "hello":
                                var v17 uint8 = 228
                                _ = v17
                            }
                        default:
                            if true {
                                var v18 int32 = int32(v3)
                                var v19 string = ""
                                _ = v18
                                _ = v19
                            }
                        }
                        switch ("foo bar" + "foo bar") {
                        case "hello":
                            var v20 uint8 = 42
                            _ = v20
                        case "":
                            v21 := GenS2{F0: v3}
                            _ = v21
                        }
                        _ = fi16
                    }
                } else {
                    var v22 bool = false
                    v23 := []int{(^(fi15)), fi15, fi15, 1}
                    if len(v23) > 0 {
                        v23[0] = (fi15 - fi15)
                    }
                    _ = v22
                    _ = v23
                }
            }
            _ = fi15
        }
        v11 += 42
    } else {
        var v24 float64 = ((0.0 - 3.14159) + (1.5 - 1e10))
        fmt.Println(v24, v4, v4)
        switch 9223372036854775807 {
        case 128:
            var v25 int8 = v3
            _ = v25
        }
        _ = v24
    }
    v26 := GenS1{F0: (v14 - v14)}
    v27 := []int{-993, -1, -498}
    if true {
        v3 = -1
        fmt.Println(v11)
        fmt.Println(v4, v2, v14)
    } else {
        fmt.Println(v14)
    }
    switch ((("" + "a") + "\n\t") + "hello") {
    case "foo bar":
        fmt.Println(v2, v2)
    case "hello":
        var v28 uint = 107
        _ = v28
    case "\n\t":
        var v29 int32 = 2147483647
        _ = v29
    }
    var v30 int32 = -2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v9
}
