package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32 }
func genFn0() uint8 {
    v4 := []int{-1, 127}
    var v5 int8 = 100
    v5 -= v5
    _ = v4
    _ = v5
    return 255
}

func main() {
    var v1 float32 = (0.1 - 0.0)
    v2 := make(map[uint32]uint32)
    var v3 int = 127
    var v6 uint = 0
    if (!(-2147483648 == (v3 * 1))) {
        if ((true || true) || false) {
            var v7 uint32 = 42
            delete(v2, v7)
            if false {
                v3 += int(v7)
                v8 := []int{}
                v3 = (v3 - (v3 % ((v3) | 1)))
                _ = v8
            } else {
                switch ("foo bar" + "a") {
                case "a":
                    var v9 float64 = -2.25
                    _ = v9
                case "foo bar":
                    if false {
                        v10 := GenS1{F0: 2147483647}
                        var v11 int32 = 2147483647
                        _ = v10
                        _ = v11
                    } else {
                        var v12 float64 = (1.5 - -2.25)
                        _ = v12
                    }
                case "\n\t":
                    switch ("hello" + "foo bar") {
                    case "\n\t":
                        if (((float64(float32(0.1)) < (float64((1.5 * v1)) * 1e10)) || false) && false) {
                            var v13 int = v3
                            _ = v13
                        } else {
                            v14 := GenS1{F0: 0}
                            v15 := []int{int(v7)}
                            _ = v14
                            _ = v15
                        }
                    case "":
                        fmt.Println(v6)
                    case "foo bar":
                        for fi16 := 0; fi16 < 7; fi16++ {
                            var v17 int64 = 1
                            _ = fi16
                            _ = v17
                        }
                    default:
                        v6 = 42
                    }
                default:
                    var v18 int = (v3 / (((v3 | v3)) | 1))
                    _ = v18
                }
                var v19 uint64 = 0
                _ = v19
            }
            _ = v7
        } else {
            v1 = -2.25
            for fi20 := 0; fi20 < 5; fi20++ {
                if false {
                    fmt.Println(v3)
                } else {
                    switch "a" {
                    case "":
                        v21 := []int{}
                        _ = v21
                    default:
                        switch ((("hello" + (("\n\t" + "") + "foo bar")) + "foo bar") + "a") {
                        case "a":
                            v22 := []int{}
                            _ = v22
                        default:
                            if (127 != 1) {
                                switch v3 {
                                case -1:
                                    fmt.Println(fi20)
                                case 2147483647:
                                    for fi23 := 0; fi23 < 15; fi23++ {
                                        v6 = 931
                                        _ = fi23
                                    }
                                }
                                var v24 string = "foo bar"
                                v25 := []int{1, fi20}
                                _ = v24
                                _ = v25
                            } else {
                                fmt.Println(v3)
                            }
                        }
                    }
                    v3 = 0
                }
                fmt.Println(v3)
                _ = fi20
            }
            delete(v2, 4294967295)
        }
        var v26 bool = false
        _ = v26
    }
    fmt.Println(v1, v6, v6)
    v2[42] = 4294967295
    v1 = v1
    _ = len(v2)
    fmt.Println(v6, v6)
    var v27 int16 = 42
    v28 := GenS2{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v6
}
