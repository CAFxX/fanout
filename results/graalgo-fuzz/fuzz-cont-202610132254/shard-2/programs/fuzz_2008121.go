package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 int32; F1 int }
func genFn0(p0 int16, p1 string, p2 int8) int {
    p0 += -32768
    p2 -= (p2 / (((p2 - p2)) | 1))
    _ = p0
    _ = p1
    _ = p2
    return 1
}
func genFn1(p0 int8) uint16 {
    var v5 int64 = 1
    for fi6 := 0; fi6 < 1; fi6++ {
        p0 = (-(p0))
        _ = fi6
    }
    _ = p0
    _ = v5
    return 255
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]int)
    var v3 int8 = 11
    var v4 string = ""
    var v7 uint8 = 2
    var v8 int8 = v3
    for fi9 := 0; fi9 < 14; fi9++ {
        v10 := []int{}
        _ = fi9
        _ = v10
    }
    switch v4 {
    case "foo bar":
        v11 := []int{521}
        _ = v11
    }
    var v12 uint = 65535
    v13 := GenS1{F0: 43, F1: 42}
    v4 = (v4 + v4)
    var v14 int32 = 0
    var v15 float32 = 0.1
    fmt.Println(v4)
    v14 += v14
    var v16 uint32 = 0
    v8 = v3
    fmt.Println(v15)
    if (!((v7 % ((127) | 1)) > (v7 + (v7 * 0)))) {
        switch (v4 + v4) {
        case "hello":
            v16 += v16
        case "foo bar":
            if false {
                var v17 string = v4
                _ = v17
            } else {
                var v18 int = 2147483647
                v15 = 1.5
                _ = v18
            }
        case "":
            for fi19 := 0; fi19 < 11; fi19++ {
                v3 = (v3 & 42)
                fmt.Println(v8, v7, v3)
                _ = fi19
            }
        }
    } else {
        fmt.Println(v16)
        var v20 uint = v12
        if false {
            if false {
                v21 := GenS2{F0: 1, F1: 159}
                if false {
                    fmt.Println(v14)
                } else {
                    switch ("foo bar" + v4) {
                    case "":
                        v22 := []int{}
                        _ = v22
                    case "\n\t":
                        v3 += v3
                    }
                    var v23 float32 = 0.1
                    for fi24 := 0; fi24 < 2; fi24++ {
                        if false {
                            var v25 int64 = 314
                            if (uint16(v7) != 255) {
                                v8 += (v8 ^ int8(v16))
                            }
                            _ = v25
                        } else {
                            v3 = (v3 ^ 127)
                            fmt.Println(v14, v4, v20)
                            v15 = v23
                        }
                        _ = fi24
                    }
                    _ = v23
                }
                _ = v21
            }
            var v26 uint = v12
            _ = v26
        }
        _ = v20
    }
    v8 = (^(v8))
    var v27 float32 = v15
    if true {
        v3 -= (-(v3))
        v27 += 3.14159
        v28 := []int{}
        _ = v28
    } else {
        v14 -= v14
        switch "" {
        case "hello":
            v29 := GenS1{F0: v14, F1: 42}
            _ = v29
        }
        v30 := v2[((v4 + "foo bar") + (v4 + v4))]
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
