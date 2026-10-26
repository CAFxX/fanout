package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int; F1 uint64 }
type GenS5 struct { F0 uint8 }

func main() {
    var v1 int = 0
    v2 := make(map[uint32]int)
    var v3 int64 = 429
    var v4 uint16 = 65535
    var v5 int8 = -128
    v6 := []uint32{867, 4294967295, 209, 1}
    v7 := GenS1{F0: 230}
    switch v1 {
    case 0:
        v8 := GenS1{F0: 128}
        _ = v8
    case -537:
        var v9 uint32 = 4294967295
        _ = v9
    }
    v3 += v3
    fmt.Println(v1, v1)
    var v10 int = v1
    v1 = (v1 ^ -2147483648)
    fmt.Println(v4)
    v11 := GenS1{F0: 153}
    var v12 float32 = float32((0.1 - 3.14159))
    var v13 uint8 = 42
    for fi14 := 0; fi14 < 4; fi14++ {
        var v15 string = "foo bar"
        var v16 int8 = (v5 & (v5 << ((2) & 63)))
        v4 += (v4 | v4)
        _ = fi14
        _ = v15
        _ = v16
    }
    if (491 <= 1) {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        if false {
            _ = len(v2)
            var v17 uint64 = 42
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = v17
        }
        switch v10 {
        case 882:
            for fi18 := 0; fi18 < 12; fi18++ {
                for fi19 := 0; fi19 < 18; fi19++ {
                    _ = len(v2)
                    fmt.Println(v1, v1, v3)
                    _ = fi19
                }
                var v20 int32 = -791
                _ = fi18
                _ = v20
            }
        }
    } else {
        v21 := GenS2{F0: 988, F1: v4}
        _ = v21
    }
    v22 := v2[348]
    var v23 uint = 18446744073709551615
    for fi24 := 0; fi24 < 11; fi24++ {
        v25 := GenS3{F0: "hello"}
        if false {
            if ((42 != 834) && (!(!true))) {
                var v26 int8 = v5
                _ = v26
            }
        } else {
            v27 := GenS4{F0: 32767, F1: 1}
            _ = v27
        }
        v28 := GenS5{F0: (v13 * v13)}
        _ = fi24
        _ = v25
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
