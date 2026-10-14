package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 uint16, p1 string) int64 {
    v4 := GenS1{F0: 2, F1: 18446744073709551615}
    for fi5 := 0; fi5 < 17; fi5++ {
        var v6 int = fi5
        _ = fi5
        _ = v6
    }
    var v7 bool = true
    var v8 int16 = -32768
    _ = p0
    _ = p1
    _ = v4
    _ = v7
    _ = v8
    return -9223372036854775808
}
func genFn1(p0 uint32, p1 uint16) int32 {
    p1 = 321
    for fi9 := 0; fi9 < 14; fi9++ {
        var v10 int8 = 0
        v11 := []int{(fi9 | 762)}
        var v12 string = "\n\t"
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 bool = (!(2147483647 <= 1))
    var v2 int32 = -109
    var v3 float32 = 0.1
    v13 := GenS2{F0: 100}
    for fi14 := 0; fi14 < 3; fi14++ {
        if v1 {
            switch "" {
            case "hello":
                var v15 int = fi14
                _ = v15
            case "foo bar":
                for fi16 := 0; fi16 < 11; fi16++ {
                    var v17 float32 = (float32(-2.25) * (v3 * 1e10))
                    _ = fi16
                    _ = v17
                }
            case "a":
                v3 = ((v3 + v3) * 3.14159)
            }
            v2 += v2
            var v18 float32 = 1e10
            _ = v18
        } else {
            v2 -= 1
            switch "\n\t" {
            case "\n\t":
                v2 = (v2 << ((506) & 63))
            }
            v3 -= v3
        }
        _ = fi14
    }
    var v19 uint = 463
    if (v1 || (0 == 0)) {
        v19 -= (v19 & (^(v19)))
        v1 = v1
    }
    switch ("\n\t" + "\n\t") {
    case "hello":
        v19 += v19
    default:
        var v20 int16 = 32767
        _ = v20
    }
    var v21 uint8 = 2
    if false {
        fmt.Println(v1)
        v22 := GenS1{F0: 2, F1: 18446744073709551615}
        _ = v22
    }
    var v23 int16 = -32768
    for fi24 := 0; fi24 < 2; fi24++ {
        fmt.Println(v3)
        _ = fi24
    }
    var v25 string = "foo bar"
    switch v25 {
    case "":
        v26 := []int{4294967295, 678, -210}
        _ = v26
    case "foo bar":
        var v27 bool = (!v1)
        _ = v27
    default:
        v28 := GenS3{F0: 65535}
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v21
    _ = v23
    _ = v25
    _ = v3
}
