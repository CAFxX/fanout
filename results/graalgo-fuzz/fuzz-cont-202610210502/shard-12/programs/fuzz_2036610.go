package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint8 }
type GenS2 struct { F0 int32 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 236
    var v3 string = (("hello" + "hello") + ("\n\t" + ""))
    v4 := make(map[int]bool)
    var v5 int = 32767
    if false {
        var v6 float64 = float64(float32(1.5))
        _ = v6
    }
    var v7 float64 = 3.14159
    var v8 uint64 = 2
    var v9 uint32 = (v2 / ((2147483647) | 1))
    v5 -= 127
    fmt.Println(v5, v2)
    if (true || true) {
        v10 := GenS1{F0: -357, F1: v5, F2: 42}
        _ = v10
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        _ = len(v4)
        _ = fi11
    }
    var v12 float64 = 3.14159
    v13 := GenS1{F0: -2147483648, F1: -1, F2: 1}
    switch v5 {
    case 567:
        if true {
            v14 := []int{-1, v5}
            _ = v14
        }
    case -479:
        var v15 uint32 = v9
        _ = v15
    case -2147483648:
        var v16 uint = 42
        _ = v16
    default:
        for fi17 := 0; fi17 < 13; fi17++ {
            _ = len(v4)
            v9 += 2147483648
            _ = fi17
        }
    }
    v18 := []int{(v5 & 9223372036854775807), v5, -217}
    switch (v3 + "foo bar") {
    case "\n\t":
        v2 = v2
    case "foo bar":
        _ = len(v18)
    case "":
        for fi19 := 0; fi19 < 19; fi19++ {
            switch (v3 + v3) {
            case "a":
                v20 := GenS1{F0: -2147483648, F1: (-(fi19)), F2: 128}
                _ = v20
            case "":
                var v21 uint = 0
                _ = v21
            case "hello":
                switch "a" {
                case "a":
                    v22 := GenS1{F0: 507, F1: -9223372036854775808, F2: 116}
                    _ = v22
                case "hello":
                    for fi23 := 0; fi23 < 13; fi23++ {
                        if (-105 != -128) {
                            v9 += (^(v2))
                            fmt.Println(v8)
                            _ = len(v4)
                            switch fi19 {
                            case -212:
                                var v24 float64 = v7
                                _ = v24
                            case 106:
                                fmt.Println(v9, fi23, v7)
                            case 127:
                                v9 = (v2 ^ 2)
                            default:
                                var v25 float32 = 0.1
                                _ = v25
                            }
                            var v26 uint = 501
                            _ = v26
                        } else {
                            fmt.Println(v12)
                            switch 32768 {
                            case 844:
                                if false {
                                    v27 := GenS2{F0: 1}
                                    fmt.Println(v8, v9)
                                    _ = v27
                                } else {
                                    _ = len(v4)
                                    if len(v18) >= 2 {
                                        v18 = v18[1:len(v18)]
                                    }
                                    var v28 float64 = ((-2.25 * 0.0) * 1e10)
                                    _ = v28
                                }
                            case -222:
                                var v29 int64 = -1
                                _ = v29
                            case -163:
                                var v30 int64 = 42
                                _ = v30
                            }
                        }
                        _ = fi23
                    }
                case "foo bar":
                    v8 += v8
                }
            }
            switch (^(fi19)) {
            case -710:
                var v31 uint16 = 1
                _ = v31
            }
            for fi32 := 0; fi32 < 9; fi32++ {
                var v33 uint64 = v8
                _ = fi32
                _ = v33
            }
            _ = fi19
        }
    }
    var v34 string = "a"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
