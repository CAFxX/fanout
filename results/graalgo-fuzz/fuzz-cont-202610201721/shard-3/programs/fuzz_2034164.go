package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 float32 = float32(0.0)
    var v2 int32 = -571
    v3 := make(map[int]string)
    switch ("\n\t" + "") {
    case "\n\t":
        switch "a" {
        case "hello":
            v2 += (v2 << ((1) & 63))
        case "\n\t":
            delete(v3, 2147483647)
        case "":
            switch (("hello" + ("" + ("foo bar" + ("\n\t" + "")))) + "a") {
            case "a":
                fmt.Println(v2, v1)
            case "":
                v4 := GenS1{F0: v2}
                _ = v4
            case "hello":
                fmt.Println(v1)
            default:
                var v5 float32 = v1
                _ = v5
            }
        }
    }
    v2 = v2
    var v6 int = 1
    v7 := []int{}
    var v8 float32 = 1e10
    v6 = v6
    v1 = 1.5
    var v9 int8 = 1
    var v10 uint8 = 63
    if false {
        if (true && false) {
            v11 := GenS1{F0: (v2 - (v2 % (((v2 << ((1) & 63))) | 1)))}
            _ = v11
        } else {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            for fi12 := 0; fi12 < 5; fi12++ {
                switch fi12 {
                case 32768:
                    var v13 bool = false
                    _ = v13
                case -942:
                    switch (v6 / ((v6) | 1)) {
                    case 32767:
                        v1 += v8
                    case 102:
                        for fi14 := 0; fi14 < 12; fi14++ {
                            v10 -= v10
                            v10 = v10
                            _ = fi14
                        }
                    default:
                        v15 := GenS1{F0: v2}
                        _ = v15
                    }
                case 127:
                    var v16 int16 = -357
                    _ = v16
                default:
                    v3[(v6 >> ((0) & 63))] = "foo bar"
                }
                if (0 < v9) {
                    var v17 uint8 = v10
                    _ = v17
                } else {
                    switch (("\n\t" + "a") + "a") {
                    case "foo bar":
                        var v18 int16 = -32768
                        _ = v18
                    }
                    fmt.Println(v9)
                    if ((v2 + v2) == v2) {
                        fmt.Println(v6, v9, v9)
                        switch "\n\t" {
                        case "foo bar":
                            v2 += v2
                        case "":
                            var v19 float64 = 0.0
                            _ = v19
                        default:
                            v9 -= 42
                        }
                        var v20 int16 = -644
                        _ = v20
                    } else {
                        switch v6 {
                        case 668:
                            v9 += (v9 ^ v9)
                        case -2147483648:
                            for fi21 := 0; fi21 < 14; fi21++ {
                                var v22 int8 = v9
                                _ = fi21
                                _ = v22
                            }
                        case -9223372036854775808:
                            for fi23 := 0; fi23 < 9; fi23++ {
                                v10 = (v10 * 127)
                                v2 -= -415
                                _ = fi23
                            }
                        }
                        v24 := GenS1{F0: v2}
                        _ = v24
                    }
                }
                fmt.Println(v2, v2)
                _ = fi12
            }
        }
    }
    v10 = (v10 + v10)
    var v25 int32 = (v2 - v2)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v25
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
