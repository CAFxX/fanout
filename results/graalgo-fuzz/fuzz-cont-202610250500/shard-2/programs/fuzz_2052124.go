package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 uint; F1 uint64; F2 uint16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int16; F1 int64; F2 uint8 }
type GenS6 struct { F0 int8 }
func genFn0(p0 uint32, p1 int16) int16 {
    v5 := GenS1{F0: 9223372036854775807}
    fmt.Println(p0, p1)
    v6 := GenS2{F0: 2, F1: 42}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (p1 | -1)
}

func main() {
    var v1 int = -2147483648
    var v2 bool = (("" >= ((("foo bar" + "\n\t") + ("hello" + ("a" + ("" + "hello")))) + "a")) && (-1 <= 1))
    v3 := []int{(-(v1))}
    var v4 int16 = 42
    if v2 {
        for fi7 := 0; fi7 < 1; fi7++ {
            if len(v3) > 0 {
                v3[0] = 315
            }
            if false {
                fmt.Println(fi7, v2)
            } else {
                v8 := GenS2{F0: 807, F1: -9223372036854775808}
                _ = v8
            }
            _ = fi7
        }
        var v9 uint32 = 668
        for fi10 := 0; fi10 < 5; fi10++ {
            for fi11 := 0; fi11 < 18; fi11++ {
                if len(v3) > 0 {
                    v3[0] = (fi11 * fi11)
                }
                _ = fi11
            }
            v1 = (^(v1))
            _ = fi10
        }
        _ = v9
    } else {
        if (!v2) {
            switch "hello" {
            case "hello":
                var v12 int64 = 9223372036854775807
                _ = v12
            case "":
                fmt.Println(v4)
            case "foo bar":
                v13 := GenS3{F0: 1, F1: 0, F2: 262}
                _ = v13
            }
            v14 := GenS4{F0: 312}
            _ = v14
        }
        v1 = v1
    }
    v1 = v1
    if v2 {
        var v15 float64 = float64(-2.25)
        fmt.Println(v1)
        _ = v15
    } else {
        if len(v3) > 0 {
            v3[0] = int(v4)
        }
        v16 := GenS5{F0: (v4 << ((255) & 63)), F1: 522, F2: 42}
        fmt.Println(v1, v1)
        _ = v16
    }
    var v17 string = ("foo bar" + ("" + ""))
    fmt.Println(v2, v4, v17)
    v2 = (v4 != genFn0(42, (-(v4))))
    v4 = -32768
    v17 = "a"
    v18 := GenS6{F0: 100}
    var v19 float32 = 1.5
    for fi20 := 0; fi20 < 8; fi20++ {
        var v21 string = ("a" + "\n\t")
        switch v21 {
        case "hello":
            var v22 uint16 = 151
            _ = v22
        case "\n\t":
            v2 = (v2 || (42 == 128))
        default:
            var v23 uint32 = 2
            _ = v23
        }
        _ = fi20
        _ = v21
    }
    for fi24 := 0; fi24 < 9; fi24++ {
        var v25 uint64 = 2
        fmt.Println(v19, v25, v4)
        _ = fi24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
