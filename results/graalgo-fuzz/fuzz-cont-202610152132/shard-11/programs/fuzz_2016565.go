package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint32 }
type GenS2 struct { F0 uint64; F1 bool }
func genFn0() bool {
    v7 := GenS1{F0: 653, F1: 0, F2: 42}
    var v8 uint8 = 255
    _ = v7
    _ = v8
    return (1e10 == 1.5)
}

func main() {
    var v1 uint = 42
    var v2 uint32 = 283
    v3 := make(map[int]uint32)
    var v4 int16 = -32768
    v5 := []string{((("" + ("a" + "a")) + ("a" + "foo bar")) + "\n\t"), ("hello" + (("foo bar" + ("hello" + "\n\t")) + ("a" + ("\n\t" + "\n\t"))))}
    var v6 int8 = -128
    if (734 >= 1) {
        _ = len(v3)
    } else {
        for fi9 := 0; fi9 < 13; fi9++ {
            for fi10 := 0; fi10 < 10; fi10++ {
                fmt.Println(v4, fi10, fi9)
                var v11 float32 = 3.14159
                _ = fi10
                _ = v11
            }
            _ = fi9
        }
    }
    if (float32(1.5) > 0.1) {
        if true {
            var v12 uint8 = 107
            fmt.Println(v12)
            fmt.Println(v4, v12)
            _ = v12
        }
        fmt.Println(v2, v2, v1)
    }
    for fi13 := 0; fi13 < 1; fi13++ {
        v6 += (v6 << (((v1 % ((v1) | 1))) & 63))
        for fi14 := 0; fi14 < 12; fi14++ {
            fmt.Println(fi13, v4)
            var v15 int = 0
            _ = fi14
            _ = v15
        }
        v1 += (v1 | (v1 & (v1 >> ((579) & 63))))
        _ = fi13
    }
    v3[4294967295] = (v2 ^ 65535)
    for fi16 := 0; fi16 < 18; fi16++ {
        _ = len(v3)
        for fi17 := 0; fi17 < 4; fi17++ {
            switch "a" {
            case "hello":
                if true {
                    var v18 bool = true
                    v5 = append(v5, "hello")
                    _ = v18
                } else {
                    v1 -= (v1 >> (((v1 - (v1 * v1))) & 63))
                    var v19 uint64 = 2
                    v20 := GenS2{F0: (^(v19)), F1: (true && (15 >= 1))}
                    _ = v19
                    _ = v20
                }
            case "foo bar":
                var v21 int16 = 744
                _ = v21
            }
            _ = fi17
        }
        _ = fi16
    }
    if (-413 != -1) {
        for fi22 := 0; fi22 < 11; fi22++ {
            for fi23 := 0; fi23 < 20; fi23++ {
                if true {
                    switch "hello" {
                    case "hello":
                        switch (fi22 - (^(fi23))) {
                        case -206:
                            v24 := GenS1{F0: 4294967295, F1: 9223372036854775807, F2: v2}
                            _ = v24
                        }
                    }
                }
                var v25 uint8 = 255
                var v26 int64 = -1
                _ = fi23
                _ = v25
                _ = v26
            }
            _ = fi22
        }
        v6 -= (v6 >> ((18446744073709551615) & 63))
        var v27 int64 = 42
        _ = v27
    }
    if true {
        v3[-591] = (v2 >> ((v1) & 63))
    } else {
        var v28 uint64 = 42
        _ = v28
    }
    switch 2147483647 {
    case 4294967295:
        var v29 int32 = 0
        _ = v29
    case -1:
        var v30 float32 = 1e10
        _ = v30
    case 475:
        var v31 int32 = 2147483647
        _ = v31
    }
    v32 := GenS2{F0: 2, F1: (!(1 <= 42))}
    var v33 uint32 = (^(v2))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
