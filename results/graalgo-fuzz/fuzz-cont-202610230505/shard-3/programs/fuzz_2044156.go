package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 int64 }

func main() {
    var v1 uint64 = 0
    v2 := []string{}
    var v3 int8 = 42
    var v4 float64 = 0.0
    var v5 int8 = (v3 ^ v3)
    if (!true) {
        if true {
            v1 = v1
        }
        for fi6 := 0; fi6 < 9; fi6++ {
            v7 := GenS1{F0: 2147483647, F1: 2}
            _ = fi6
            _ = v7
        }
    } else {
        v4 += ((3.14159 + v4) - (v4 + v4))
    }
    if false {
        for fi8 := 0; fi8 < 4; fi8++ {
            var v9 int = fi8
            _ = len(v2)
            _ = fi8
            _ = v9
        }
    } else {
        if false {
            switch "" {
            case "a":
                _ = len(v2)
            case "hello":
                var v10 uint = 127
                _ = v10
            case "":
                v5 += v5
            }
        }
        v3 -= v3
        switch -1 {
        case 128:
            var v11 int64 = 9223372036854775807
            _ = v11
        case 338:
            var v12 bool = ((v1 * v1) > 18446744073709551615)
            _ = v12
        }
    }
    var v13 float32 = (1e10 * ((-2.25 / 3.14159) * (1e10 + (1e10 + -2.25))))
    v14 := GenS2{F0: 216}
    v4 = (v4 * -2.25)
    v13 -= (-2.25 - -2.25)
    switch 475 {
    case -647:
        for fi15 := 0; fi15 < 13; fi15++ {
            var v16 int = (fi15 << ((401) & 63))
            if false {
                switch fi15 {
                case -886:
                    switch "a" {
                    case "\n\t":
                        _ = len(v2)
                    case "hello":
                        _ = len(v2)
                    case "foo bar":
                        v5 = 49
                    }
                case 127:
                    for fi17 := 0; fi17 < 5; fi17++ {
                        var v18 int16 = 1
                        fmt.Println(v3)
                        _ = fi17
                        _ = v18
                    }
                case -877:
                    v4 -= v4
                }
            } else {
                for fi19 := 0; fi19 < 16; fi19++ {
                    fmt.Println(v13, fi19)
                    v2 = append(v2, ("\n\t" + ("\n\t" + "")))
                    var v20 uint32 = 4294967295
                    _ = fi19
                    _ = v20
                }
                v3 -= v3
            }
            if false {
                var v21 int32 = 980
                _ = v21
            }
            _ = fi15
            _ = v16
        }
    default:
        v3 += v5
    }
    var v22 bool = (("hello" >= "foo bar") && (-849 <= 42))
    fmt.Println(v22, v4, v5)
    if false {
        if len(v2) > 0 {
            v2[0] = ""
        }
        for fi23 := 0; fi23 < 19; fi23++ {
            v1 = v1
            v13 = v13
            _ = fi23
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
