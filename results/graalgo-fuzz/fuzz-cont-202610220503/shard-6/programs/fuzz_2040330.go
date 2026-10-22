package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 int64 }

func main() {
    v1 := []uint32{}
    var v2 int64 = -828
    v3 := make(map[string]uint32)
    v4 := make(map[string]bool)
    var v5 int16 = 42
    v6 := make(map[string]uint32)
    v7 := GenS1{F0: -67}
    v5 -= v5
    switch "\n\t" {
    case "foo bar":
        delete(v3, "")
    default:
        v8 := GenS1{F0: 100}
        _ = v8
    }
    if true {
        switch "foo bar" {
        case "\n\t":
            v9 := GenS1{F0: 0}
            _ = v9
        }
    }
    var v10 int8 = 40
    var v11 string = ""
    if (true && (0.0 != 1.5)) {
        var v12 uint64 = 317
        _ = v12
    } else {
        v13 := GenS1{F0: (v10 - v10)}
        v14 := GenS2{F0: 32767, F1: 42}
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 16; fi15++ {
        fmt.Println(v11)
        _ = fi15
    }
    if true {
        var v16 string = v11
        var v17 uint = 4294967295
        fmt.Println(v11)
        _ = v16
        _ = v17
    } else {
        switch -513 {
        case 32768:
            _ = len(v1)
        case 127:
            for fi18 := 0; fi18 < 4; fi18++ {
                fmt.Println(v2, v10)
                for fi19 := 0; fi19 < 12; fi19++ {
                    v20 := GenS2{F0: 255, F1: 0}
                    _ = fi19
                    _ = v20
                }
                var v21 int16 = v5
                _ = fi18
                _ = v21
            }
        default:
            var v22 int8 = 100
            _ = v22
        }
    }
    var v23 int = 32767
    var v24 int32 = 0
    var v25 uint = 18446744073709551615
    v26 := GenS3{F0: (v2 % ((v2) | 1))}
    switch v23 {
    case -9223372036854775808:
        v23 -= v23
    }
    for fi27 := 0; fi27 < 5; fi27++ {
        if (238 <= 170) {
            switch int(v2) {
            case 4294967295:
                for fi28 := 0; fi28 < 6; fi28++ {
                    if (!false) {
                        switch ("a" + "a") {
                        case "hello":
                            delete(v4, v11)
                            v29 := GenS2{F0: 42, F1: 2147483647}
                            var v30 bool = false
                            _ = v29
                            _ = v30
                        }
                        var v31 int64 = (-(v2))
                        _ = v31
                    }
                    _ = fi28
                }
            case 716:
                for fi32 := 0; fi32 < 6; fi32++ {
                    _ = len(v6)
                    if len(v1) > 0 {
                        v1[0] = 2
                    }
                    _ = fi32
                }
            }
            v25 -= 660
            v2 -= (v2 % (((v2 * 1)) | 1))
        }
        _ = fi27
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
