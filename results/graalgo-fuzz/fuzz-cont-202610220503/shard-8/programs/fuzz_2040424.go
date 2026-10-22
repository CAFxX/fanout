package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 uint16; F2 int }
type GenS3 struct { F0 uint32; F1 int32; F2 int32 }
func genFn0(p0 int16, p1 string, p2 uint16) bool {
    p0 -= (-(p0))
    v7 := []int{-2147483648}
    var v8 uint16 = 1
    var v9 float32 = 0.1
    if false {
        if false {
            if true {
                for fi10 := 0; fi10 < 5; fi10++ {
                    if false {
                        if (v8 < p2) {
                            var v11 string = (("hello" + ("\n\t" + p1)) + (p1 + p1))
                            var v12 float32 = (1.5 * -2.25)
                            _ = v11
                            _ = v12
                        }
                        if (485 >= 159) {
                            var v13 string = "foo bar"
                            if (!false) {
                                _ = len(v7)
                                v14 := GenS1{F0: p0}
                                _ = v14
                            } else {
                                p2 = 341
                                if len(v7) > 0 {
                                    v7[0] = -1
                                }
                                if false {
                                    v15 := GenS1{F0: -462}
                                    v13 = "a"
                                    var v16 float64 = float64(((1.5 * float32(float64(v9))) * ((v9 / 0.1) + float32(1e10))))
                                    _ = v15
                                    _ = v16
                                }
                            }
                            var v17 bool = (-1 <= 42)
                            _ = v13
                            _ = v17
                        }
                    }
                    var v18 string = (p1 + p1)
                    _ = fi10
                    _ = v18
                }
                if true {
                    var v19 float64 = 3.14159
                    v20 := GenS2{F0: 0, F1: v8, F2: -368}
                    _ = v19
                    _ = v20
                }
            }
            v7 = append(v7, 4294967295)
            v7 = append(v7, 4294967295)
        }
    } else {
        _ = len(v7)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return (-9223372036854775808 >= 257)
}

func main() {
    var v1 float64 = (-2.25 * (-2.25 + 1.5))
    var v2 int64 = -24
    var v3 uint32 = 903
    v4 := make(map[string]string)
    var v5 bool = (!(1 == 18446744073709551615))
    v6 := make(map[uint32]uint32)
    fmt.Println(v3, v5)
    var v21 float64 = (v1 / (-2.25 + (v1 / v1)))
    v22 := v4["\n\t"]
    v23 := GenS3{F0: (v3 / (((^(v3))) | 1)), F1: 0, F2: 0}
    fmt.Println(v1)
    v21 = v1
    var v24 bool = v5
    if v24 {
        if v5 {
            if true {
                v3 -= v3
                fmt.Println(v1, v1)
            }
            for fi25 := 0; fi25 < 7; fi25++ {
                var v26 uint8 = 254
                for fi27 := 0; fi27 < 14; fi27++ {
                    var v28 int = (fi27 + (fi27 % ((-2147483648) | 1)))
                    var v29 int16 = 42
                    var v30 float32 = 1e10
                    _ = fi27
                    _ = v28
                    _ = v29
                    _ = v30
                }
                _ = fi25
                _ = v26
            }
            v31 := []int{-1, 32767, -19}
            _ = v31
        }
        fmt.Println(v21, v3)
        var v32 float64 = v21
        _ = v32
    }
    fmt.Println(v22)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
