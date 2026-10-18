package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string }
func genFn0(p0 int8) uint8 {
    v5 := GenS1{F0: (-32768 != -32768)}
    p0 = p0
    v6 := []int{395, 9223372036854775807, 42, 0}
    _ = p0
    _ = v5
    _ = v6
    return 1
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]int)
    var v3 float64 = float64((0.0 / 0.1))
    var v4 int8 = 127
    var v7 int32 = 42
    if ((v7 * -967) != (^(v7))) {
        if false {
            var v8 string = ("hello" + "foo bar")
            switch v8 {
            case "\n\t":
                v9 := []int{}
                _ = v9
            default:
                if (32767 < 127) {
                    for fi10 := 0; fi10 < 8; fi10++ {
                        fmt.Println(fi10)
                        v11 := GenS1{F0: (v8 <= v8)}
                        v4 -= (^(v4))
                        _ = fi10
                        _ = v11
                    }
                    v12 := GenS1{F0: true}
                    _ = v12
                } else {
                    if (!(((0.0 / (-2.25 - float32(0.0))) - (-2.25 - 0.1)) >= -2.25)) {
                        var v13 uint8 = 42
                        _ = v13
                    }
                    for fi14 := 0; fi14 < 7; fi14++ {
                        switch v8 {
                        case "\n\t":
                            v4 = v4
                        }
                        if (-2.25 < float32((3.14159 + float64(float32(v3))))) {
                            v4 = v4
                            switch (fi14 >> ((879) & 63)) {
                            case 1:
                                if ((v3 * float64(3.14159)) > ((v3 + 0.0) - -2.25)) {
                                    var v15 int64 = 42
                                    var v16 int32 = v7
                                    _ = v15
                                    _ = v16
                                } else {
                                    var v17 uint16 = 726
                                    fmt.Println(fi14)
                                    _ = v17
                                }
                            case -2147483648:
                                v18 := GenS1{F0: true}
                                _ = v18
                            default:
                                if ((v8 + v8) <= v8) {
                                    var v19 int8 = v4
                                    _ = v19
                                }
                            }
                        } else {
                            v4 += (v4 / (((v4 >> ((0) & 63))) | 1))
                            var v20 uint64 = 18446744073709551615
                            _ = v20
                        }
                        v21 := GenS2{F0: "foo bar"}
                        _ = fi14
                        _ = v21
                    }
                }
            }
            var v22 int = -521
            _ = v22
            _ = v8
        }
        fmt.Println(v7)
        if ((false && (((v4 + v4) > v4) || (2 != 180))) && true) {
            var v23 bool = (!true)
            if v23 {
                v3 += v3
                fmt.Println(v7)
            }
            _ = v23
        }
    }
    v7 = (^(v7))
    v3 = v3
    v4 = -128
    for fi24 := 0; fi24 < 9; fi24++ {
        var v25 float32 = float32((3.14159 / 1.0))
        for fi26 := 0; fi26 < 7; fi26++ {
            v3 = -2.25
            _ = fi26
        }
        _ = fi24
        _ = v25
    }
    var v27 uint64 = 1
    for fi28 := 0; fi28 < 8; fi28++ {
        delete(v2, "a")
        fmt.Println(v7, v3)
        _ = fi28
    }
    switch ("foo bar" + ("foo bar" + ("foo bar" + ("hello" + "")))) {
    case "hello":
        delete(v1, 0)
    case "a":
        v27 = v27
    default:
        for fi29 := 0; fi29 < 8; fi29++ {
            for fi30 := 0; fi30 < 2; fi30++ {
                v31 := GenS2{F0: "\n\t"}
                _ = fi30
                _ = v31
            }
            for fi32 := 0; fi32 < 10; fi32++ {
                v27 -= v27
                _ = fi32
            }
            _ = fi29
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v7
}
