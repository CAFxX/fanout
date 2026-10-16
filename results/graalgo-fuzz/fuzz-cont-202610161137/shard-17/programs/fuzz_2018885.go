package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 uint }
func genFn0(p0 uint) int32 {
    var v4 uint32 = 0
    fmt.Println(p0)
    var v5 int32 = 42
    v6 := GenS1{F0: 255}
    switch "" {
    case "foo bar":
        v7 := []int{int(v4), 127, 127, 2147483647}
        _ = v7
    case "hello":
        v8 := GenS1{F0: 42}
        _ = v8
    default:
        v9 := []int{32768}
        _ = v9
    }
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return (v5 % ((v5) | 1))
}

func main() {
    var v1 float32 = 3.14159
    var v2 int32 = -1
    var v3 float32 = -2.25
    v2 += genFn0(410)
    for fi10 := 0; fi10 < 2; fi10++ {
        var v11 string = "\n\t"
        v12 := []int{}
        for fi13 := 0; fi13 < 20; fi13++ {
            switch v11 {
            case "hello":
                var v14 bool = false
                _ = v14
            }
            v11 = "a"
            _ = fi13
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    switch 128 {
    case -648:
        fmt.Println(v1, v3)
    case -1:
        switch "hello" {
        case "\n\t":
            v2 -= (v2 & genFn0(1))
        default:
            for fi15 := 0; fi15 < 2; fi15++ {
                var v16 uint8 = 127
                if true {
                    if true {
                        var v17 float64 = (1e10 - (3.14159 * 0.1))
                        if (fi15 > int(v2)) {
                            var v18 uint16 = 32767
                            for fi19 := 0; fi19 < 3; fi19++ {
                                fmt.Println(v16, v2)
                                v16 += v16
                                _ = fi19
                            }
                            v1 = float32(v17)
                            _ = v18
                        } else {
                            v17 += v17
                            v16 -= 42
                        }
                        _ = v17
                    }
                    switch (fi15 / (((-(fi15))) | 1)) {
                    case 42:
                        var v20 float32 = 0.1
                        _ = v20
                    case -9223372036854775808:
                        var v21 int16 = 42
                        _ = v21
                    case 9223372036854775807:
                        var v22 float32 = float32(-2.25)
                        _ = v22
                    default:
                        var v23 uint = 4294967295
                        _ = v23
                    }
                }
                for fi24 := 0; fi24 < 4; fi24++ {
                    if (1000 <= -32768) {
                        var v25 string = ("\n\t" + "")
                        v26 := GenS2{F0: 278, F1: 935}
                        _ = v25
                        _ = v26
                    } else {
                        if true {
                            v27 := GenS2{F0: 255, F1: 718}
                            var v28 int = 843
                            _ = v27
                            _ = v28
                        }
                    }
                    fmt.Println(fi15, v2)
                    fmt.Println(v3, v3, v16)
                    _ = fi24
                }
                _ = fi15
                _ = v16
            }
        }
    }
    fmt.Println(v1, v2, v2)
    v2 = (v2 | v2)
    for fi29 := 0; fi29 < 15; fi29++ {
        switch fi29 {
        case -753:
            var v30 uint16 = 0
            _ = v30
        case 9223372036854775807:
            v1 = (v3 - v3)
        }
        _ = fi29
    }
    var v31 float32 = (3.14159 / v3)
    switch -9223372036854775808 {
    case 32767:
        v32 := []int{345, -72, 32768}
        _ = v32
    case 310:
        v1 -= float32(((-2.25 / 1.0) * ((((float64(v3) - 3.14159) / 0.1) * 0.0) - (0.0 - -2.25))))
    default:
        v33 := []int{}
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
}
