package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 uint16 }
type GenS2 struct { F0 uint64 }

func main() {
    v1 := []int{32767, 2147483647, 1}
    var v2 int8 = -105
    var v3 uint32 = 251
    var v4 int16 = 1
    if (!true) {
        if len(v1) > 0 {
            v1[0] = 42
        }
    }
    if len(v1) > 0 {
        v1[0] = 4294967295
    }
    for fi5 := 0; fi5 < 5; fi5++ {
        v6 := GenS1{F0: 214, F1: 77, F2: 255}
        _ = fi5
        _ = v6
    }
    v7 := GenS2{F0: 2}
    if true {
        for fi8 := 0; fi8 < 20; fi8++ {
            fmt.Println(v4, v3, v3)
            for fi9 := 0; fi9 < 5; fi9++ {
                switch (("foo bar" + "a") + "hello") {
                case "a":
                    v10 := GenS1{F0: 187, F1: (v4 ^ int16(v2)), F2: 2}
                    if true {
                        var v11 uint8 = 128
                        fmt.Println(v11, v2)
                        fmt.Println(fi8, v2, v2)
                        v4 = 32767
                        switch "foo bar" {
                        case "hello":
                            fmt.Println(v11, v3, v3)
                        case "foo bar":
                            if len(v1) > 0 {
                                v1[0] = fi9
                            }
                        case "":
                            fmt.Println(fi9)
                        }
                        if false {
                            var v12 int8 = v2
                            v13 := GenS2{F0: 991}
                            _ = v12
                            _ = v13
                        }
                        _ = v11
                    }
                    _ = v10
                case "":
                    v2 -= v2
                default:
                    fmt.Println(fi8, v3, fi8)
                }
                v14 := GenS1{F0: 42, F1: (^(v4)), F2: 42}
                _ = fi9
                _ = v14
            }
            _ = fi8
        }
        v3 = (^(v3))
    }
    if len(v1) > 0 {
        v1[0] = 32767
    }
    v1 = append(v1, int(v3))
    for fi15 := 0; fi15 < 11; fi15++ {
        if (1 >= 1) {
            v3 -= 143
            v4 = v4
        }
        _ = fi15
    }
    if false {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    if len(v1) > 0 {
        v1[0] = 128
    }
    v3 -= (v3 | v3)
    var v16 float32 = (float32(1e10) - ((-2.25 / 1.0) / ((1.5 + 0.1) - float32(3.14159))))
    v1 = append(v1, 9223372036854775807)
    _ = len(v1)
    for fi17 := 0; fi17 < 9; fi17++ {
        fmt.Println(fi17, v4, v3)
        _ = fi17
    }
    fmt.Println(v16, v4)
    var v18 string = ""
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
