package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 bool; F2 uint8 }
func genFn0(p0 uint, p1 uint, p2 bool) int32 {
    if p2 {
        v4 := []int{127, 127}
        p1 += p0
        _ = v4
    } else {
        p1 = p0
    }
    v5 := []int{}
    if p2 {
        v5 = append(v5, 127)
        if p2 {
            v6 := GenS1{F0: 448}
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = v6
        }
        var v7 uint8 = 255
        _ = v7
    }
    p2 = (2 > 95)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 42
}

func main() {
    v1 := []string{"hello", "foo bar", ((("foo bar" + "a") + "a") + ("\n\t" + ""))}
    var v2 string = "a"
    var v3 string = (v2 + v2)
    v2 = v2
    if len(v1) > 0 {
        v1[0] = v3
    }
    fmt.Println(v2, v3)
    var v8 string = (v2 + ("\n\t" + ""))
    _ = len(v1)
    var v9 uint16 = 2
    switch v2 {
    case "":
        switch 42 {
        case -468:
            v3 = ((v3 + v2) + ("" + ("\n\t" + v3)))
        case 863:
            if true {
                if false {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    if false {
                        v2 = "hello"
                        v3 = (v3 + "a")
                    } else {
                        switch "\n\t" {
                        case "":
                            if false {
                                var v10 float64 = 3.14159
                                for fi11 := 0; fi11 < 11; fi11++ {
                                    var v12 uint16 = (v9 % (((v9 >> ((651) & 63))) | 1))
                                    _ = fi11
                                    _ = v12
                                }
                                v13 := GenS1{F0: (v9 + (v9 / ((v9) | 1)))}
                                _ = v10
                                _ = v13
                            }
                        }
                        v3 = (v3 + (v8 + v3))
                        var v14 uint64 = 429
                        _ = v14
                    }
                }
                for fi15 := 0; fi15 < 9; fi15++ {
                    _ = len(v1)
                    _ = fi15
                }
                if (442 < 0) {
                    v16 := GenS2{F0: 204}
                    var v17 int64 = -1
                    v18 := GenS2{F0: uint64(v9)}
                    _ = v16
                    _ = v17
                    _ = v18
                } else {
                    var v19 uint64 = 0
                    var v20 int32 = 0
                    _ = v19
                    _ = v20
                }
            }
        case 127:
            for fi21 := 0; fi21 < 10; fi21++ {
                var v22 uint32 = 42
                var v23 uint64 = 1
                var v24 uint32 = (v22 / (((^(v22))) | 1))
                _ = fi21
                _ = v22
                _ = v23
                _ = v24
            }
        default:
            v8 = ""
        }
    default:
        v8 = v3
    }
    v25 := GenS3{F0: (float64((1e10 + 3.14159)) == (((-2.25 + 0.1) * (float64(-2.25) + 1e10)) - (1.5 / (3.14159 * 1e10)))), F1: false, F2: 1}
    fmt.Println(v2, v8)
    var v26 int16 = -238
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v27 float32 = 3.14159
    v27 = ((float32(1e10) / v27) - v27)
    for fi28 := 0; fi28 < 2; fi28++ {
        var v29 string = (v2 + v8)
        fmt.Println(v27)
        switch "" {
        case "\n\t":
            v30 := GenS3{F0: true, F1: (127 >= 127), F2: 1}
            _ = v30
        }
        _ = fi28
        _ = v29
    }
    v3 = "hello"
    v2 = (v2 + (v2 + ("" + "foo bar")))
    v27 += 0.0
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v8
    _ = v9
}
