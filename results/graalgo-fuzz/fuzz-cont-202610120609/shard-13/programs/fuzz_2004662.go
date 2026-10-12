package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 int64 }

func main() {
    var v1 uint8 = 254
    v2 := []uint32{}
    v3 := make(map[uint32]string)
    var v4 uint16 = 949
    var v5 int64 = 42
    _ = len(v2)
    v6 := GenS1{F0: 0, F1: 4294967295, F2: v5}
    fmt.Println(v1, v5, v1)
    if (4294967295 >= 42) {
        for fi7 := 0; fi7 < 14; fi7++ {
            for fi8 := 0; fi8 < 6; fi8++ {
                var v9 bool = (-1 < 32767)
                _ = fi8
                _ = v9
            }
            if false {
                var v10 float64 = 0.1
                _ = len(v3)
                _ = v10
            }
            _ = fi7
        }
    }
    v11 := GenS1{F0: 42, F1: 65535, F2: v5}
    v2 = append(v2, 42)
    v4 = uint16(v1)
    for fi12 := 0; fi12 < 6; fi12++ {
        for fi13 := 0; fi13 < 1; fi13++ {
            var v14 string = "foo bar"
            _ = fi13
            _ = v14
        }
        var v15 uint = 42
        var v16 uint32 = 667
        _ = fi12
        _ = v15
        _ = v16
    }
    switch -1 {
    case -1:
        if ((!true) || false) {
            for fi17 := 0; fi17 < 10; fi17++ {
                delete(v3, 65535)
                v18 := v3[654]
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi17
                _ = v18
            }
            var v19 int32 = -620
            if true {
                var v20 int = 752
                switch (v20 | v20) {
                case 748:
                    v19 -= v19
                case -1:
                    fmt.Println(v1, v5)
                default:
                    _ = len(v2)
                }
                _ = v20
            }
            _ = v19
        }
    case -197:
        var v21 uint8 = v1
        _ = v21
    case 32767:
        var v22 int = -69
        _ = v22
    }
    switch "hello" {
    case "\n\t":
        var v23 float32 = 1.5
        _ = v23
    default:
        v24 := GenS1{F0: 966, F1: 417, F2: -819}
        _ = v24
    }
    _ = len(v3)
    var v25 uint32 = 65535
    fmt.Println(v4, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
