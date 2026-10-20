package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 int32 }

func main() {
    v1 := []uint32{2147483647, 79}
    v2 := make(map[int]uint32)
    var v3 string = ("\n\t" + "")
    v4 := make(map[string]uint32)
    var v5 int = 4294967295
    fmt.Println(v5, v3)
    v3 = v3
    fmt.Println(v5)
    v6 := GenS1{F0: -278, F1: -128}
    for fi7 := 0; fi7 < 2; fi7++ {
        if false {
            var v8 int64 = -9223372036854775808
            _ = v8
        } else {
            v5 = 2147483647
            v9 := v2[(v5 + (v5 >> ((255) & 63)))]
            _ = v9
        }
        if (true || false) {
            if len(v1) > 0 {
                v1[0] = 4294967295
            }
            _ = len(v2)
            if (127 > -1) {
                if false {
                    var v10 int32 = 0
                    _ = v10
                }
                v11 := GenS2{F0: -1}
                v12 := GenS2{F0: 0}
                _ = v11
                _ = v12
            }
        } else {
            var v13 uint8 = 42
            var v14 uint8 = 2
            _ = v13
            _ = v14
        }
        _ = fi7
    }
    if true {
        if len(v1) > 0 {
            v1[0] = 2
        }
    }
    switch v3 {
    case "\n\t":
        for fi15 := 0; fi15 < 10; fi15++ {
            var v16 uint64 = 42
            _ = len(v1)
            _ = fi15
            _ = v16
        }
    }
    if true {
        var v17 bool = false
        var v18 int64 = -476
        _ = v17
        _ = v18
    }
    switch "hello" {
    case "a":
        switch (v5 & (v5 >> ((0) & 63))) {
        case 640:
            v19 := GenS2{F0: 0}
            _ = v19
        case 32768:
            if (2 != 314) {
                switch 2147483647 {
                case 4294967295:
                    var v20 uint16 = 207
                    _ = v20
                }
            }
        }
    case "foo bar":
        v2[v5] = 98
    case "":
        switch v3 {
        case "hello":
            if (65535 > 4294967295) {
                var v21 uint16 = 497
                var v22 int8 = 127
                fmt.Println(v22, v21)
                _ = v21
                _ = v22
            }
        case "foo bar":
            v3 = (v3 + v3)
        case "\n\t":
            for fi23 := 0; fi23 < 8; fi23++ {
                if false {
                    v24 := GenS2{F0: 42}
                    _ = v24
                }
                var v25 float32 = (0.1 / 1.5)
                v1 = append(v1, 946)
                _ = fi23
                _ = v25
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
