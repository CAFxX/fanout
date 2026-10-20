package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 int64; F1 uint32; F2 string }
type GenS4 struct { F0 int8 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]string)
    var v3 int = 4294967295
    var v4 int8 = 42
    var v5 int64 = -1
    switch v3 {
    case 42:
        switch (v3 + int(v5)) {
        case -982:
            v5 = (v5 & v5)
        default:
            var v6 float64 = (0.1 * (0.1 * -2.25))
            _ = v6
        }
    case 40:
        var v7 int64 = v5
        _ = v7
    }
    for fi8 := 0; fi8 < 5; fi8++ {
        if (255 >= 138) {
            switch ("\n\t" + "\n\t") {
            case "\n\t":
                v4 = (^(v4))
            case "hello":
                for fi9 := 0; fi9 < 17; fi9++ {
                    v5 = 9223372036854775807
                    v10 := v1[(fi9 + v3)]
                    v5 = (v5 | (^(v5)))
                    _ = fi9
                    _ = v10
                }
            }
            v11 := GenS1{F0: 42}
            v12 := GenS2{F0: true, F1: 127}
            _ = v11
            _ = v12
        }
        fmt.Println(v3, v5, fi8)
        _ = fi8
    }
    for fi13 := 0; fi13 < 5; fi13++ {
        v4 -= (v4 + v4)
        v4 = -1
        v2[(("foo bar" + "a") + "")] = ""
        _ = fi13
    }
    v14 := GenS2{F0: (v4 >= (v4 & v4)), F1: (v4 * (v4 >> ((431) & 63)))}
    for fi15 := 0; fi15 < 4; fi15++ {
        v16 := GenS1{F0: 2}
        v17 := GenS1{F0: 1}
        _ = fi15
        _ = v16
        _ = v17
    }
    delete(v2, "\n\t")
    fmt.Println(v5, v3, v5)
    var v18 int64 = (v5 + v5)
    for fi19 := 0; fi19 < 3; fi19++ {
        v20 := GenS3{F0: v5, F1: 2147483647, F2: "foo bar"}
        var v21 float64 = 3.14159
        v5 -= (v5 - v18)
        _ = fi19
        _ = v20
        _ = v21
    }
    var v22 uint16 = 42
    v23 := GenS3{F0: (v5 | (v5 << ((4294967295) & 63))), F1: 42, F2: ""}
    var v24 int16 = -32768
    var v25 float32 = 3.14159
    if true {
        v24 = v24
        var v26 int8 = v4
        if ((-(v5)) <= v18) {
            for fi27 := 0; fi27 < 11; fi27++ {
                _ = len(v2)
                v28 := GenS4{F0: (v26 | 42)}
                _ = fi27
                _ = v28
            }
        }
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
