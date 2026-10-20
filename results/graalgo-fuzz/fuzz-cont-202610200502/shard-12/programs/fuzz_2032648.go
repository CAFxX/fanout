package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 int }

func main() {
    var v1 bool = (100 >= -95)
    var v2 int64 = 0
    v3 := make(map[string]int)
    v4 := GenS1{F0: false, F1: 220, F2: 127}
    var v5 bool = v1
    v6 := GenS1{F0: (false || (!v5)), F1: 128, F2: 4294967295}
    var v7 int8 = -128
    if v5 {
        var v8 uint = 2
        _ = v8
    } else {
        if (-32768 > 32767) {
            var v9 uint64 = 18446744073709551615
            for fi10 := 0; fi10 < 19; fi10++ {
                var v11 uint = 487
                fmt.Println(v1, v9, v9)
                var v12 uint32 = 4294967295
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = v9
        } else {
            delete(v3, "hello")
            var v13 uint8 = 127
            _ = v13
        }
        v2 = (^(v2))
        v14 := GenS1{F0: v1, F1: 255, F2: 4294967295}
        _ = v14
    }
    for fi15 := 0; fi15 < 16; fi15++ {
        fmt.Println(v7)
        _ = fi15
    }
    var v16 bool = ((1 != 2) || true)
    fmt.Println(v7)
    for fi17 := 0; fi17 < 5; fi17++ {
        v5 = (false && false)
        switch ("\n\t" + (("hello" + ("a" + "")) + "hello")) {
        case "\n\t":
            fmt.Println(v7)
        case "a":
            v5 = v5
        case "hello":
            var v18 int64 = 9223372036854775807
            _ = v18
        }
        if (false && (v5 && (4294967295 == 65535))) {
            fmt.Println(v16, fi17, v7)
            v1 = v5
        } else {
            var v19 uint32 = 954
            var v20 int64 = (-(v2))
            v1 = ((false && (1000 == 85)) || (42 < 65535))
            _ = v19
            _ = v20
        }
        _ = fi17
    }
    var v21 uint8 = 42
    v1 = true
    var v22 int16 = 0
    v16 = (96 == (v21 % ((128) | 1)))
    v3["hello"] = -226
    var v23 uint64 = 414
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
