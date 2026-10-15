package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16; F1 uint32 }

func main() {
    var v1 int16 = -32768
    v2 := []string{"hello", ("hello" + "foo bar")}
    var v3 bool = false
    v4 := make(map[string]bool)
    for fi5 := 0; fi5 < 7; fi5++ {
        var v6 int64 = -1
        v2 = append(v2, "")
        var v7 int32 = -938
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 uint64 = 153
    if v3 {
        var v9 float64 = 0.0
        v3 = v3
        _ = v9
    } else {
        v1 = (v1 - (v1 << ((275) & 63)))
        for fi10 := 0; fi10 < 20; fi10++ {
            var v11 int16 = v1
            _ = fi10
            _ = v11
        }
    }
    if v3 {
        for fi12 := 0; fi12 < 11; fi12++ {
            v13 := v4[""]
            v3 = (0 < 496)
            fmt.Println(v1, v3, v13)
            _ = fi12
            _ = v13
        }
    } else {
        switch 32767 {
        case 0:
            var v14 int8 = 100
            _ = v14
        case 32768:
            for fi15 := 0; fi15 < 6; fi15++ {
                for fi16 := 0; fi16 < 18; fi16++ {
                    if len(v2) > 0 {
                        v2[0] = "\n\t"
                    }
                    if len(v2) > 0 {
                        v2[0] = "a"
                    }
                    v17 := GenS1{F0: 39}
                    _ = fi16
                    _ = v17
                }
                for fi18 := 0; fi18 < 16; fi18++ {
                    switch (fi18 << ((0) & 63)) {
                    case -9223372036854775808:
                        var v19 int32 = -1
                        _ = v19
                    }
                    _ = fi18
                }
                _ = fi15
            }
        case -898:
            if len(v2) > 0 {
                v2[0] = "hello"
            }
        default:
            v20 := GenS1{F0: 113}
            _ = v20
        }
        var v21 uint32 = 42
        var v22 uint64 = v8
        _ = v21
        _ = v22
    }
    v23 := GenS1{F0: 2}
    fmt.Println(v8, v3, v3)
    v24 := GenS2{F0: 621}
    v8 = 18446744073709551615
    v1 = v1
    switch -791 {
    case 32768:
        if v3 {
            if (!(v3 || v3)) {
                for fi25 := 0; fi25 < 10; fi25++ {
                    var v26 int16 = -380
                    if len(v2) > 0 {
                        v2[0] = (((("" + ("\n\t" + "a")) + "foo bar") + (("\n\t" + "\n\t") + "\n\t")) + ((("\n\t" + ("" + "hello")) + "a") + "hello"))
                    }
                    _ = fi25
                    _ = v26
                }
            }
            var v27 uint8 = 0
            _ = v27
        }
    default:
        v28 := GenS3{F0: v1, F1: 2147483647}
        _ = v28
    }
    var v29 int16 = (v1 + (v1 << ((18446744073709551615) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v29
    _ = v3
    _ = v4
    _ = v8
}
