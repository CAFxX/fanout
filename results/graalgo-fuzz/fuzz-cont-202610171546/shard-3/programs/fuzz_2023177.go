package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 uint32 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool; F1 int16 }

func main() {
    var v1 bool = (!((65535 <= 65535) && false))
    var v2 float32 = -2.25
    v3 := []string{}
    var v4 float64 = float64(v2)
    v5 := GenS1{F0: 100, F1: 42, F2: 2}
    switch "" {
    case "":
        var v6 uint = 745
        _ = v6
    default:
        var v7 float32 = (v2 / 1.5)
        _ = v7
    }
    var v8 int16 = 1000
    v9 := GenS1{F0: 10, F1: 42, F2: 2}
    v10 := GenS2{F0: (("" + "foo bar") + "hello"), F1: 2}
    var v11 int32 = 42
    for fi12 := 0; fi12 < 2; fi12++ {
        for fi13 := 0; fi13 < 18; fi13++ {
            var v14 float32 = (float32(1e10) / (float32((float64(v2) + 3.14159)) * 3.14159))
            fmt.Println(v8)
            var v15 string = ("foo bar" + ("\n\t" + "\n\t"))
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi12
    }
    var v16 uint32 = 431
    v1 = v1
    v16 = v16
    if true {
        for fi17 := 0; fi17 < 12; fi17++ {
            for fi18 := 0; fi18 < 10; fi18++ {
                var v19 uint8 = 222
                fmt.Println(fi17, v16, v19)
                v20 := GenS3{F0: false}
                _ = fi18
                _ = v19
                _ = v20
            }
            for fi21 := 0; fi21 < 3; fi21++ {
                v22 := GenS1{F0: 1, F1: -523, F2: v16}
                _ = fi21
                _ = v22
            }
            _ = fi17
        }
        var v23 int = -238
        if (false || v1) {
            var v24 uint32 = v16
            var v25 uint8 = 128
            _ = v24
            _ = v25
        }
        _ = v23
    }
    for fi26 := 0; fi26 < 6; fi26++ {
        var v27 uint16 = 1
        if ((((v16 * 2147483647) > (v16 << ((143) & 63))) && v1) || false) {
            for fi28 := 0; fi28 < 9; fi28++ {
                v1 = (127 <= 222)
                v29 := GenS4{F0: (v4 > (v4 / v4)), F1: (v8 << ((42) & 63))}
                var v30 uint16 = 65535
                _ = fi28
                _ = v29
                _ = v30
            }
            switch fi26 {
            case 32767:
                v1 = v1
            case 4294967295:
                var v31 uint8 = 8
                _ = v31
            }
            if v1 {
                v2 -= (float32(v4) / ((float32(v4) + v2) / v2))
                var v32 uint16 = v27
                var v33 int = fi26
                _ = v32
                _ = v33
            }
        }
        _ = fi26
        _ = v27
    }
    fmt.Println(v16, v8)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
