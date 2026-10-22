package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 uint64; F1 int64 }

func main() {
    var v1 bool = false
    v2 := []uint32{514, 207}
    v3 := []uint32{}
    var v4 int = 127
    fmt.Println(v1)
    fmt.Println(v4, v4)
    v1 = v1
    if v1 {
        _ = len(v3)
        switch -2147483648 {
        case -304:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case 0:
            if v1 {
                v5 := GenS1{F0: 159, F1: -1}
                for fi6 := 0; fi6 < 5; fi6++ {
                    fmt.Println(fi6, fi6)
                    v4 = v4
                    _ = fi6
                }
                v1 = v1
                _ = v5
            } else {
                v1 = v1
                switch ("a" + ("" + "hello")) {
                case "foo bar":
                    v7 := GenS1{F0: 65535, F1: 32767}
                    _ = v7
                case "\n\t":
                    v4 = v4
                case "a":
                    if false {
                        var v8 string = ("a" + "a")
                        var v9 uint32 = 0
                        for fi10 := 0; fi10 < 3; fi10++ {
                            _ = len(v3)
                            fmt.Println(v1)
                            _ = fi10
                        }
                        _ = v8
                        _ = v9
                    } else {
                        v4 = (v4 + (v4 << ((4294967295) & 63)))
                        var v11 uint64 = 339
                        if len(v2) > 0 {
                            v2[0] = 4294967295
                        }
                        _ = v11
                    }
                }
                for fi12 := 0; fi12 < 18; fi12++ {
                    if ("hello" >= "a") {
                        var v13 int32 = 42
                        var v14 uint = 42
                        fmt.Println(v14, v4, v1)
                        _ = v13
                        _ = v14
                    }
                    _ = fi12
                }
            }
        case -9223372036854775808:
            v15 := GenS1{F0: 389, F1: -341}
            _ = v15
        }
        var v16 int64 = 42
        _ = v16
    } else {
        v4 += 9223372036854775807
        v4 = v4
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v1, v1)
    var v17 float32 = float32(((((1e10 / (float64(0.1) + 0.1)) * -2.25) + 1e10) / float64((float32(((-2.25 - 0.0) + 1e10)) + 1.5))))
    _ = len(v3)
    for fi18 := 0; fi18 < 16; fi18++ {
        if v1 {
            for fi19 := 0; fi19 < 9; fi19++ {
                v4 = 128
                v20 := GenS2{F0: 0, F1: 1}
                v21 := GenS2{F0: 862, F1: 42}
                _ = fi19
                _ = v20
                _ = v21
            }
            for fi22 := 0; fi22 < 18; fi22++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi22
            }
        } else {
            fmt.Println(v4, v17, v1)
        }
        _ = fi18
    }
    v2 = append(v2, 2147483648)
    var v23 int32 = 2147483647
    if false {
        var v24 int8 = -128
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
