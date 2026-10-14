package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }

func main() {
    var v1 int8 = 127
    var v2 int = 2147483647
    var v3 uint8 = 2
    var v4 uint8 = (v3 / (((v3 << ((371) & 63))) | 1))
    var v5 bool = true
    v6 := make(map[string]bool)
    v5 = true
    if (2 < 1) {
        var v7 uint = 126
        _ = v7
    }
    for fi8 := 0; fi8 < 8; fi8++ {
        _ = len(v6)
        v5 = false
        _ = fi8
    }
    if true {
        v5 = (!v5)
    } else {
        if ((1 <= 705) || (17 > 255)) {
            for fi9 := 0; fi9 < 10; fi9++ {
                fmt.Println(v4, v1)
                var v10 int64 = 42
                v4 -= v3
                _ = fi9
                _ = v10
            }
        }
        if v5 {
            v5 = v5
        }
        v6["foo bar"] = true
    }
    var v11 int32 = -1
    switch "" {
    case "hello":
        v11 -= -1
    case "":
        switch (v2 << ((594) & 63)) {
        case 42:
            fmt.Println(v4)
        case 32768:
            if (-342 >= -861) {
                fmt.Println(v2, v5, v11)
                v4 = v3
            }
        }
    }
    v12 := GenS1{F0: (v2 & 32767), F1: 32767}
    var v13 uint = 0
    var v14 int16 = int16(v11)
    fmt.Println(v13)
    var v15 int8 = (v1 & 0)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
