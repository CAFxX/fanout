package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 int = -2147483648
    var v2 uint64 = 18446744073709551615
    v3 := []uint32{2, uint32(v2), uint32(v2)}
    var v4 int32 = 1
    for fi5 := 0; fi5 < 20; fi5++ {
        var v6 int = (v1 / ((v1) | 1))
        var v7 bool = true
        _ = fi5
        _ = v6
        _ = v7
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v8 int = (v1 + (v1 & (v1 & 2147483647)))
    if len(v3) > 0 {
        v3[0] = uint32(v2)
    }
    var v9 uint16 = 988
    for fi10 := 0; fi10 < 1; fi10++ {
        for fi11 := 0; fi11 < 1; fi11++ {
            switch "foo bar" {
            case "hello":
                v2 -= v2
            default:
                var v12 bool = (255 < uint8(v2))
                _ = v12
            }
            var v13 int16 = -192
            _ = fi11
            _ = v13
        }
        _ = fi10
    }
    if false {
        switch v1 {
        case 1:
            v8 += v1
        case 265:
            switch ("hello" + (("hello" + ("hello" + "hello")) + ("a" + "a"))) {
            case "foo bar":
                var v14 string = "\n\t"
                _ = v14
            case "hello":
                if len(v3) > 0 {
                    v3[0] = 515
                }
            case "\n\t":
                switch ("foo bar" + ("hello" + (("a" + "hello") + "\n\t"))) {
                case "foo bar":
                    v3 = append(v3, 600)
                case "hello":
                    var v15 float32 = -2.25
                    _ = v15
                case "":
                    fmt.Println(v9, v8, v1)
                }
            }
        case 0:
            for fi16 := 0; fi16 < 6; fi16++ {
                var v17 int16 = 1000
                v17 -= v17
                _ = fi16
                _ = v17
            }
        default:
            if (!(v4 > v4)) {
                var v18 uint8 = 128
                _ = v18
            } else {
                var v19 int16 = 42
                var v20 int64 = -1
                _ = v19
                _ = v20
            }
        }
        if len(v3) > 0 {
            v3[0] = 806
        }
        var v21 int8 = 42
        _ = v21
    } else {
        v22 := GenS1{F0: 550, F1: 0, F2: 255}
        v23 := GenS1{F0: 1000, F1: 1, F2: 42}
        v24 := GenS1{F0: 1000, F1: 2, F2: 4294967295}
        _ = v22
        _ = v23
        _ = v24
    }
    switch (^(v1)) {
    case 988:
        v25 := GenS2{F0: 891}
        _ = v25
    case 423:
        var v26 int = 2147483647
        _ = v26
    }
    var v27 int8 = -128
    if len(v3) > 0 {
        v3[0] = 0
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
