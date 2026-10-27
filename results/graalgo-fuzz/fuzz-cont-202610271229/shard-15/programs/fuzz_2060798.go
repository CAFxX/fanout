package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 bool }

func main() {
    var v1 bool = true
    var v2 uint64 = 42
    v3 := []uint32{}
    var v4 uint16 = 301
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    _ = len(v3)
    for fi5 := 0; fi5 < 4; fi5++ {
        var v6 bool = true
        _ = fi5
        _ = v6
    }
    var v7 uint64 = (v2 ^ 1)
    if (0.0 != 3.14159) {
        for fi8 := 0; fi8 < 14; fi8++ {
            for fi9 := 0; fi9 < 13; fi9++ {
                v2 = v2
                v3 = append(v3, 1)
                var v10 uint32 = 502
                _ = fi9
                _ = v10
            }
            v3 = append(v3, 4294967295)
            _ = fi8
        }
        var v11 int64 = 1
        _ = v11
    }
    if (!true) {
        v3 = append(v3, 2147483647)
        if v1 {
            switch "foo bar" {
            case "foo bar":
                var v12 float64 = (3.14159 * 0.0)
                _ = v12
            }
            v13 := GenS1{F0: -624}
            _ = v13
        } else {
            var v14 uint16 = 748
            if (v2 != 4294967295) {
                v15 := GenS1{F0: 42}
                _ = v15
            } else {
                for fi16 := 0; fi16 < 6; fi16++ {
                    var v17 uint = 255
                    v18 := GenS1{F0: 0}
                    v1 = ((!((1 <= 127) && (0 <= 353))) || (!(32 < 201)))
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                fmt.Println(v4, v1)
                var v19 int64 = -9223372036854775808
                _ = v19
            }
            for fi20 := 0; fi20 < 4; fi20++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                switch fi20 {
                case -1:
                    v2 = (v7 / ((v2) | 1))
                }
                v1 = (v1 && ("" > "hello"))
                _ = fi20
            }
            _ = v14
        }
        switch 127 {
        case 1:
            for fi21 := 0; fi21 < 20; fi21++ {
                for fi22 := 0; fi22 < 1; fi22++ {
                    v23 := GenS1{F0: 42}
                    v4 = (v4 | (v4 % (((v4 | 2)) | 1)))
                    _ = fi22
                    _ = v23
                }
                var v24 uint = 918
                fmt.Println(v4, fi21, v2)
                _ = fi21
                _ = v24
            }
        case 180:
            fmt.Println(v4, v1)
        case -572:
            for fi25 := 0; fi25 < 4; fi25++ {
                v7 += v7
                _ = len(v3)
                _ = fi25
            }
        default:
            var v26 uint32 = 4294967295
            _ = v26
        }
    }
    v4 = (v4 + 0)
    fmt.Println(v2, v7, v7)
    var v27 int64 = 735
    if ((v1 || (-2147483648 > -1)) || (42 != -1)) {
        var v28 float64 = 0.1
        v27 = v27
        _ = len(v3)
        _ = v28
    } else {
        v29 := GenS2{F0: -1, F1: true}
        _ = v29
    }
    if ("" != ("hello" + "a")) {
        var v30 int = 32767
        _ = v30
    }
    if (float32((0.1 + 1.5)) > 0.1) {
        if (42 <= 383) {
            var v31 float64 = 0.1
            for fi32 := 0; fi32 < 11; fi32++ {
                _ = len(v3)
                _ = fi32
            }
            _ = v31
        }
        v3 = append(v3, 65535)
    } else {
        v33 := GenS1{F0: v27}
        v7 += (v2 >> ((255) & 63))
        _ = v33
    }
    for fi34 := 0; fi34 < 15; fi34++ {
        var v35 uint16 = v4
        fmt.Println(fi34)
        v35 -= 858
        _ = fi34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v7
}
