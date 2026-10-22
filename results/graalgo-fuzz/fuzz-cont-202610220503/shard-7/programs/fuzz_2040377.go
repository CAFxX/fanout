package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }

func main() {
    var v1 int16 = 1000
    v2 := make(map[uint32]int)
    v3 := []uint32{853, 42, 42}
    if false {
        v4 := v2[42]
        var v5 uint = 65535
        _ = v4
        _ = v5
    } else {
        fmt.Println(v1)
        var v6 int16 = 76
        _ = v6
    }
    var v7 int8 = 42
    switch -2147483648 {
    case 9223372036854775807:
        v7 -= (v7 % (((v7 << ((916) & 63))) | 1))
    case 127:
        v8 := GenS1{F0: "a", F1: ((-9223372036854775808 != 42) && (131 == 128))}
        _ = v8
    }
    var v9 uint = 1
    if true {
        if (!(-2.25 >= (float64(1.5) / 3.14159))) {
            var v10 int64 = 0
            switch (("\n\t" + "hello") + ("hello" + "a")) {
            case "":
                for fi11 := 0; fi11 < 10; fi11++ {
                    if false {
                        v7 += v7
                    }
                    _ = fi11
                }
            case "foo bar":
                v12 := GenS1{F0: ("" + "a"), F1: true}
                _ = v12
            case "a":
                var v13 float64 = float64(-2.25)
                _ = v13
            }
            _ = v10
        }
        fmt.Println(v7, v9)
        fmt.Println(v1, v7, v1)
    } else {
        switch 32768 {
        case 2147483647:
            v1 += (v1 / ((int16(v7)) | 1))
        case 0:
            fmt.Println(v9, v7)
        }
        var v14 uint16 = 202
        v1 = v1
        _ = v14
    }
    _ = len(v3)
    var v15 int32 = 1
    fmt.Println(v15, v7, v9)
    for fi16 := 0; fi16 < 17; fi16++ {
        var v17 uint64 = 2
        v9 -= (v9 >> ((0) & 63))
        v18 := v2[0]
        _ = fi16
        _ = v17
        _ = v18
    }
    _ = len(v2)
    var v19 bool = (2 == 4294967295)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v7
    _ = v9
}
