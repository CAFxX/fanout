package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 int16, p1 string) uint8 {
    p1 = (p1 + p1)
    p1 = "hello"
    _ = p0
    _ = p1
    return 255
}
func genFn1(p0 uint8, p1 uint16, p2 int64) uint64 {
    var v5 string = "\n\t"
    if (-2.25 != (float64(0.1) - 3.14159)) {
        v5 = ""
        if false {
            if false {
                if (true && true) {
                    fmt.Println(v5, v5, p0)
                    for fi6 := 0; fi6 < 18; fi6++ {
                        switch (v5 + "a") {
                        case "foo bar":
                            var v7 int64 = 1
                            _ = v7
                        case "":
                            v8 := []int{}
                            _ = v8
                        case "\n\t":
                            var v9 int = (fi6 - (fi6 | fi6))
                            _ = v9
                        default:
                            for fi10 := 0; fi10 < 20; fi10++ {
                                v11 := GenS1{F0: 464}
                                _ = fi10
                                _ = v11
                            }
                        }
                        _ = fi6
                    }
                } else {
                    switch -2147483648 {
                    case -990:
                        v12 := GenS1{F0: 0}
                        _ = v12
                    }
                }
                p0 += 2
                if (358 > 1) {
                    var v13 bool = false
                    p0 += (^(p0))
                    _ = v13
                }
            }
        }
    } else {
        var v14 bool = true
        if v14 {
            p1 = (^(p1))
        }
        _ = v14
    }
    var v15 uint16 = 255
    v16 := GenS1{F0: 2}
    v17 := []int{-2147483648, 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    _ = v17
    _ = v5
    return 561
}

func main() {
    v1 := []string{"hello"}
    var v2 int = 128
    v3 := make(map[uint32]uint32)
    var v4 string = ("" + "a")
    v4 = "hello"
    var v18 float64 = (1e10 - 1e10)
    var v19 int8 = 6
    switch ("hello" + (("\n\t" + (v4 + ("a" + ""))) + ("" + "\n\t"))) {
    case "\n\t":
        if (!((1 >= -1) || false)) {
            var v20 int16 = -32768
            var v21 uint8 = 42
            _ = v20
            _ = v21
        }
    default:
        if false {
            v19 = (v19 - (v19 % ((0) | 1)))
            v22 := GenS1{F0: 255}
            fmt.Println(v19, v18, v19)
            _ = v22
        }
    }
    for fi23 := 0; fi23 < 1; fi23++ {
        fmt.Println(v19)
        _ = fi23
    }
    switch v4 {
    case "":
        if true {
            v2 = (v2 >> ((1) & 63))
            var v24 float64 = 1.5
            _ = v24
        } else {
            switch (v4 + "hello") {
            case "a":
                var v25 int8 = (v19 << ((255) & 63))
                _ = v25
            case "hello":
                if (-9223372036854775808 != 497) {
                    var v26 uint16 = 1
                    _ = len(v3)
                    v18 = v18
                    _ = v26
                } else {
                    var v27 string = v4
                    if (((((1e10 * 0.1) + 1.5) * 1.5) / float32(float64(3.14159))) != (-2.25 - float32((v18 - 0.0)))) {
                        var v28 uint = 78
                        if ((0.1 <= -2.25) || (65535 <= 2)) {
                            for fi29 := 0; fi29 < 9; fi29++ {
                                _ = len(v3)
                                _ = fi29
                            }
                            _ = len(v1)
                        }
                        var v30 int = 128
                        _ = v28
                        _ = v30
                    } else {
                        var v31 uint64 = 4294967295
                        _ = v31
                    }
                    v18 = 1.5
                    _ = v27
                }
            default:
                v19 = (^(v19))
            }
        }
    }
    fmt.Println(v18, v19)
    if true {
        fmt.Println(v4, v2)
    }
    fmt.Println(v19, v18, v18)
    v32 := GenS1{F0: 49}
    if (2147483647 > 1) {
        fmt.Println(v2, v2, v19)
    } else {
        var v33 float64 = ((0.1 * v18) + float64(-2.25))
        _ = v33
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v32
    _ = v4
}
