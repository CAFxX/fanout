package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 int16 }

func main() {
    var v1 uint32 = 4294967295
    var v2 int8 = -1
    var v3 bool = false
    v4 := []string{"\n\t", ("hello" + ("hello" + "\n\t")), ("a" + "hello")}
    var v5 int32 = 728
    var v6 uint = 0
    for fi7 := 0; fi7 < 9; fi7++ {
        v5 -= (v5 + v5)
        if (65535 == 209) {
            if true {
                for fi8 := 0; fi8 < 2; fi8++ {
                    var v9 uint64 = 42
                    if v3 {
                        var v10 uint32 = (v1 % (((v1 / ((v1) | 1))) | 1))
                        if (false && false) {
                            var v11 uint32 = (v1 + (v1 << ((1) & 63)))
                            var v12 int = (^(fi8))
                            _ = len(v4)
                            _ = v11
                            _ = v12
                        } else {
                            fmt.Println(v5)
                            var v13 uint64 = (v9 * 128)
                            _ = v13
                        }
                        _ = v10
                    } else {
                        v3 = v3
                        v2 -= (v2 & 115)
                    }
                    _ = fi8
                    _ = v9
                }
                for fi14 := 0; fi14 < 19; fi14++ {
                    if (true && v3) {
                        fmt.Println(fi14, v1)
                        v15 := GenS1{F0: 128}
                        v1 = 42
                        _ = v15
                    }
                    _ = len(v4)
                    fmt.Println(fi7, v1, v6)
                    _ = fi14
                }
                var v16 int = (fi7 << ((v6) & 63))
                _ = v16
            }
            _ = len(v4)
            var v17 float32 = ((1.5 * float32(float64(1e10))) - 3.14159)
            _ = v17
        }
        _ = fi7
    }
    var v18 bool = v3
    if v3 {
        v3 = (!true)
        for fi19 := 0; fi19 < 8; fi19++ {
            var v20 int64 = 1
            for fi21 := 0; fi21 < 12; fi21++ {
                fmt.Println(v20, fi19, v6)
                v6 -= uint(v1)
                if v3 {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v6 += v6
                    var v22 int8 = v2
                    _ = v22
                }
                _ = fi21
            }
            _ = fi19
            _ = v20
        }
        for fi23 := 0; fi23 < 17; fi23++ {
            if v18 {
                v3 = (false && (v3 && (0.0 <= 0.1)))
                v24 := GenS1{F0: 127}
                _ = v24
            }
            fmt.Println(v5, v5, v2)
            _ = fi23
        }
    }
    var v25 float32 = -2.25
    var v26 float32 = 0.0
    var v27 bool = (255 != 2)
    if v27 {
        v28 := GenS2{F0: 42}
        _ = v28
    } else {
        if len(v4) > 0 {
            v4[0] = "a"
        }
        var v29 uint8 = 58
        for fi30 := 0; fi30 < 5; fi30++ {
            v31 := GenS3{F0: 811, F1: 42}
            if v3 {
                var v32 int8 = v2
                _ = v32
            }
            _ = fi30
            _ = v31
        }
        _ = v29
    }
    var v33 int8 = -1
    switch "" {
    case "hello":
        fmt.Println(v27, v2, v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
