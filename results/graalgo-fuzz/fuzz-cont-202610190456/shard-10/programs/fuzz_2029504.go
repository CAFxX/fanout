package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }

func main() {
    v1 := make(map[string]bool)
    var v2 int64 = -9223372036854775808
    v3 := []string{"hello", "\n\t", "foo bar"}
    var v4 int = 1
    var v5 uint16 = 42
    var v6 int32 = -2147483648
    var v7 string = ""
    var v8 uint64 = 4294967295
    switch v7 {
    case "foo bar":
        for fi9 := 0; fi9 < 5; fi9++ {
            v10 := GenS1{F0: 128, F1: uint32(v8)}
            _ = fi9
            _ = v10
        }
    default:
        var v11 int64 = v2
        _ = v11
    }
    for fi12 := 0; fi12 < 16; fi12++ {
        var v13 uint = 4294967295
        if (-9 != -128) {
            var v14 uint64 = v8
            _ = v14
        }
        _ = fi12
        _ = v13
    }
    for fi15 := 0; fi15 < 1; fi15++ {
        var v16 bool = (57 < 127)
        v17 := GenS1{F0: 128, F1: 65535}
        var v18 uint = 0
        _ = fi15
        _ = v16
        _ = v17
        _ = v18
    }
    fmt.Println(v6, v5, v8)
    switch (v7 + ((v7 + v7) + v7)) {
    case "a":
        switch v7 {
        case "hello":
            delete(v1, v7)
        case "":
            switch -800 {
            case 128:
                v8 = 18446744073709551615
            case -110:
                var v19 uint8 = 90
                _ = v19
            case -774:
                var v20 uint = 2
                _ = v20
            }
        }
    case "hello":
        var v21 uint8 = 42
        _ = v21
    default:
        var v22 int64 = v2
        _ = v22
    }
    var v23 int64 = 952
    if (18446744073709551615 == uint(v8)) {
        fmt.Println(v23, v4)
    }
    for fi24 := 0; fi24 < 11; fi24++ {
        _ = len(v1)
        _ = fi24
    }
    switch -1 {
    case 42:
        var v25 uint32 = 72
        _ = v25
    case 560:
        _ = len(v3)
    }
    switch v4 {
    case -9223372036854775808:
        var v26 int8 = 90
        _ = v26
    case 0:
        switch int(v6) {
        case -2147483648:
            v4 = (v4 ^ v4)
        }
    case 380:
        v2 -= v23
    }
    v6 += (v6 / ((v6) | 1))
    fmt.Println(v6, v4, v8)
    var v27 int16 = int16(v23)
    for fi28 := 0; fi28 < 18; fi28++ {
        var v29 uint = 0
        v30 := GenS1{F0: 42, F1: 2147483647}
        v29 = (v29 - v29)
        _ = fi28
        _ = v29
        _ = v30
    }
    switch v7 {
    case "foo bar":
        if true {
            switch int(v27) {
            case 672:
                v27 += -32768
            case -873:
                v2 = 1
            case 4294967295:
                switch (v4 + (v4 + v4)) {
                case 480:
                    var v31 int8 = int8(v8)
                    _ = v31
                }
            }
            v7 = (v7 + v7)
            v3 = append(v3, v7)
        } else {
            var v32 int8 = -54
            v33 := GenS1{F0: 73, F1: 2}
            var v34 bool = true
            _ = v32
            _ = v33
            _ = v34
        }
    default:
        var v35 int8 = 100
        _ = v35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v23
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
