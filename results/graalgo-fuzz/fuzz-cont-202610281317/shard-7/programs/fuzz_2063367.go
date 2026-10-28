package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 uint; F2 uint }

func main() {
    var v1 float64 = 1e10
    var v2 uint32 = 1
    var v3 int = 399
    v4 := []uint32{2}
    var v5 uint = 42
    var v6 int16 = -32768
    if len(v4) > 0 {
        v4[0] = (v2 + v2)
    }
    switch 32767 {
    case 45:
        _ = len(v4)
    case -452:
        v3 = v3
    case -143:
        switch "\n\t" {
        case "":
            var v7 int16 = v6
            _ = v7
        default:
            switch (v3 & (v3 << (((v5 | (^(v5)))) & 63))) {
            case -1:
                v2 += v2
            case 343:
                var v8 int8 = 42
                _ = v8
            case 9223372036854775807:
                switch ((((("a" + "foo bar") + "foo bar") + "hello") + "\n\t") + "\n\t") {
                case "\n\t":
                    v2 += v2
                case "hello":
                    v9 := GenS1{F0: int32(v3)}
                    _ = v9
                case "":
                    v10 := GenS2{F0: v3}
                    _ = v10
                }
            default:
                fmt.Println(v6, v2)
            }
        }
    }
    var v11 int8 = 1
    for fi12 := 0; fi12 < 14; fi12++ {
        fmt.Println(v5, fi12, v11)
        _ = fi12
    }
    switch "a" {
    case "\n\t":
        var v13 float64 = v1
        _ = v13
    default:
        if false {
            v14 := GenS3{F0: v3, F1: v5, F2: (v5 * v5)}
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            switch "foo bar" {
            case "a":
                v4 = append(v4, (v2 | 2147483648))
            case "\n\t":
                for fi15 := 0; fi15 < 3; fi15++ {
                    v5 += (v5 >> (((^(v5))) & 63))
                    var v16 bool = false
                    _ = fi15
                    _ = v16
                }
            case "foo bar":
                v5 = 1
            default:
                v5 += (v5 >> (((v5 / ((1) | 1))) & 63))
            }
            _ = v14
        }
    }
    v1 += v1
    v6 = 0
    fmt.Println(v3, v5)
    if (v5 > (v5 - (v5 * 255))) {
        v4 = append(v4, v2)
        v5 = (v5 & (v5 | v5))
    } else {
        for fi17 := 0; fi17 < 6; fi17++ {
            v1 -= 1e10
            v18 := GenS2{F0: (fi17 | 4294967295)}
            if (false || (!(2147483647 > 42))) {
                var v19 uint16 = 32767
                _ = v19
            }
            _ = fi17
            _ = v18
        }
        var v20 string = "hello"
        _ = v20
    }
    for fi21 := 0; fi21 < 1; fi21++ {
        for fi22 := 0; fi22 < 6; fi22++ {
            fmt.Println(v5, v5, v5)
            v23 := GenS3{F0: fi22, F1: v5, F2: 2}
            var v24 uint16 = 2
            _ = fi22
            _ = v23
            _ = v24
        }
        if true {
            if (1 != 0) {
                v1 = v1
                var v25 int32 = -61
                fmt.Println(v5)
                _ = v25
            }
        }
        var v26 float64 = v1
        _ = fi21
        _ = v26
    }
    fmt.Println(v11, v1, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
