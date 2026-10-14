package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 int; F1 int8 }
func genFn0(p0 uint16) uint8 {
    v7 := []int{485, 32768}
    var v8 string = ("hello" + "hello")
    if true {
        var v9 int16 = -1
        var v10 int32 = -593
        _ = v10
        _ = v9
    }
    _ = p0
    _ = v7
    _ = v8
    return 255
}

func main() {
    var v1 int64 = 237
    var v2 int = 42
    var v3 int32 = 42
    v4 := make(map[string]bool)
    v5 := make(map[int]int)
    var v6 bool = ((128 > 228) || false)
    for fi11 := 0; fi11 < 13; fi11++ {
        v2 += 2147483647
        if v6 {
            var v12 float32 = (float32((0.1 - (3.14159 * (0.0 * -2.25)))) + 0.1)
            v13 := []int{}
            fmt.Println(fi11, v3, fi11)
            _ = v12
            _ = v13
        } else {
            if v6 {
                v14 := []int{v2}
                _ = v14
            } else {
                if v6 {
                    if true {
                        v3 = (v3 * -554)
                        var v15 int16 = 42
                        _ = v15
                    }
                    var v16 int = fi11
                    for fi17 := 0; fi17 < 2; fi17++ {
                        var v18 int8 = -128
                        _ = fi17
                        _ = v18
                    }
                    _ = v16
                }
                v2 += v2
                var v19 string = "foo bar"
                _ = v19
            }
            var v20 int64 = (v1 - (v1 + (v1 >> ((564) & 63))))
            _ = v20
        }
        _ = fi11
    }
    v21 := GenS1{F0: true}
    if (false || true) {
        fmt.Println(v6, v1, v1)
        var v22 int64 = v1
        _ = v22
    }
    v2 += (v2 << ((864) & 63))
    var v23 uint16 = 42
    var v24 uint16 = (^(v23))
    switch "" {
    case "hello":
        v25 := GenS1{F0: (!(("hello" + "foo bar") != "foo bar"))}
        _ = v25
    case "\n\t":
        v26 := GenS2{F0: 70, F1: (v1 - -123)}
        _ = v26
    }
    v27 := []int{-9223372036854775808, (v2 % (((v2 >> ((174) & 63))) | 1)), (v2 & 128)}
    fmt.Println(v23)
    switch ("foo bar" + "foo bar") {
    case "":
        v28 := GenS1{F0: (!(320 <= 1))}
        _ = v28
    case "foo bar":
        if false {
            for fi29 := 0; fi29 < 18; fi29++ {
                for fi30 := 0; fi30 < 13; fi30++ {
                    v31 := GenS3{F0: (v2 * 1), F1: 27}
                    if v6 {
                        var v32 int16 = 0
                        _ = v32
                    } else {
                        if len(v27) >= 2 {
                            v27 = v27[1:len(v27)]
                        }
                        v33 := v5[2147483647]
                        _ = v33
                    }
                    _ = fi30
                    _ = v31
                }
                _ = fi29
            }
            v2 = (v2 << ((716) & 63))
        }
    }
    v34 := GenS1{F0: true}
    var v35 int64 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v23
    _ = v24
    _ = v27
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
