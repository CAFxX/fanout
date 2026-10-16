package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 int }
type GenS4 struct { F0 string; F1 int64 }
type GenS5 struct { F0 uint64 }
func genFn0(p0 uint64) uint {
    v5 := GenS1{F0: 255, F1: 42, F2: 0}
    p0 = (p0 | p0)
    for fi6 := 0; fi6 < 12; fi6++ {
        var v7 string = "\n\t"
        _ = fi6
        _ = v7
    }
    p0 = (^(p0))
    _ = p0
    _ = v5
    return uint(p0)
}

func main() {
    var v1 uint32 = 42
    var v2 float64 = ((0.0 / ((3.14159 + -2.25) * -2.25)) + (1.5 * float64(0.0)))
    var v3 bool = (42 != 1)
    var v4 int8 = 1
    v1 = 65535
    if (42 >= 42) {
        switch (("foo bar" + ("a" + "a")) + ("foo bar" + (("hello" + "") + ("" + ("\n\t" + ("hello" + ("hello" + ""))))))) {
        case "a":
            v2 -= (v2 - (float64(1e10) * float64(0.1)))
        case "":
            switch "hello" {
            case "hello":
                for fi8 := 0; fi8 < 12; fi8++ {
                    switch "" {
                    case "":
                        v9 := []int{127, fi8, (fi8 | fi8)}
                        _ = v9
                    case "foo bar":
                        v3 = (v3 || v3)
                    case "hello":
                        for fi10 := 0; fi10 < 12; fi10++ {
                            fmt.Println(v4, fi8)
                            _ = fi10
                        }
                    }
                    for fi11 := 0; fi11 < 8; fi11++ {
                        fmt.Println(fi11, v3)
                        _ = fi11
                    }
                    _ = fi8
                }
            }
        default:
            var v12 int32 = 821
            _ = v12
        }
        v13 := GenS1{F0: 2, F1: 32767, F2: 128}
        _ = v13
    } else {
        var v14 uint64 = 0
        for fi15 := 0; fi15 < 4; fi15++ {
            for fi16 := 0; fi16 < 8; fi16++ {
                v17 := GenS2{F0: "a"}
                v1 -= v1
                var v18 string = "foo bar"
                _ = fi16
                _ = v17
                _ = v18
            }
            _ = fi15
        }
        _ = v14
    }
    fmt.Println(v4)
    v4 -= v4
    v2 = 0.0
    for fi19 := 0; fi19 < 8; fi19++ {
        v20 := GenS3{F0: 4294967295, F1: 127}
        v21 := GenS3{F0: 79, F1: 0}
        _ = fi19
        _ = v20
        _ = v21
    }
    v4 -= 1
    v22 := GenS4{F0: "", F1: 42}
    fmt.Println(v2, v2)
    fmt.Println(v1)
    var v23 int16 = -280
    switch 32768 {
    case 177:
        v24 := GenS5{F0: 4294967295}
        _ = v24
    }
    switch "\n\t" {
    case "\n\t":
        v2 -= v2
    default:
        v25 := GenS3{F0: 1, F1: 376}
        _ = v25
    }
    fmt.Println(v23, v23)
    var v26 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v3
    _ = v4
}
