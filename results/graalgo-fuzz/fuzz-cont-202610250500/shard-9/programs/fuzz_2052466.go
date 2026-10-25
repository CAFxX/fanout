package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 int8 }
type GenS2 struct { F0 uint8; F1 string; F2 uint }
type GenS3 struct { F0 uint32; F1 int64 }

func main() {
    var v1 int16 = -1
    var v2 int = 9223372036854775807
    v3 := []uint32{347, 2, 42}
    v2 -= (v2 - 128)
    if len(v3) > 0 {
        v3[0] = 583
    }
    if len(v3) > 0 {
        v3[0] = 1
    }
    for fi4 := 0; fi4 < 2; fi4++ {
        var v5 string = ("foo bar" + "a")
        if len(v3) > 0 {
            v3[0] = 1
        }
        _ = fi4
        _ = v5
    }
    fmt.Println(v1, v1)
    _ = len(v3)
    fmt.Println(v2, v2)
    if false {
        fmt.Println(v2, v1)
        if false {
            var v6 int32 = int32(v2)
            if len(v3) > 0 {
                v3[0] = 2147483647
            }
            _ = v6
        }
        v7 := GenS1{F0: 929, F1: 191, F2: 1}
        _ = v7
    } else {
        switch (("hello" + ("" + "a")) + "foo bar") {
        case "\n\t":
            if false {
                fmt.Println(v2)
                v8 := GenS2{F0: 42, F1: (("hello" + "foo bar") + "a"), F2: 65535}
                _ = v8
            } else {
                var v9 int16 = (v1 - (v1 << ((2) & 63)))
                fmt.Println(v2)
                v2 -= v2
                _ = v9
            }
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    if true {
        _ = len(v3)
        if (false || true) {
            v1 = 473
            fmt.Println(v2)
        } else {
            v10 := GenS2{F0: 236, F1: ("" + "\n\t"), F2: 535}
            _ = v10
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v11 uint8 = 255
    v2 = 128
    var v12 string = "hello"
    if (true && false) {
        _ = len(v3)
    }
    v13 := GenS3{F0: 246, F1: -989}
    var v14 float64 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
}
