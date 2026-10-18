package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 int }
func genFn0(p0 uint32) int64 {
    p0 += (p0 << ((65535) & 63))
    if (32767 < 1000) {
        p0 -= p0
        var v5 int8 = 1
        v6 := []int{}
        _ = v5
        _ = v6
    } else {
        switch ("a" + ("a" + ("\n\t" + (("a" + ("foo bar" + "foo bar")) + "foo bar")))) {
        case "":
            p0 -= p0
        case "hello":
            v7 := []int{919}
            _ = v7
        case "\n\t":
            fmt.Println(p0)
        default:
            if false {
                p0 = (^(p0))
            }
        }
        v8 := GenS1{F0: 127}
        var v9 string = ("foo bar" + "hello")
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 16; fi10++ {
        v11 := []int{9223372036854775807, fi10, -2147483648}
        fmt.Println(p0, p0)
        _ = fi10
        _ = v11
    }
    var v12 uint64 = 2
    _ = p0
    _ = v12
    return -189
}

func main() {
    var v1 int64 = 1
    var v2 string = ""
    var v3 bool = false
    var v4 uint32 = 42
    v3 = v3
    v4 = v4
    var v13 uint = uint(v4)
    v14 := GenS1{F0: -9223372036854775808}
    if v3 {
        for fi15 := 0; fi15 < 6; fi15++ {
            v1 += v1
            for fi16 := 0; fi16 < 16; fi16++ {
                v17 := GenS2{F0: "", F1: 688}
                v18 := GenS1{F0: (fi15 - (fi15 & (^(fi16))))}
                var v19 int = fi16
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
            }
            _ = fi15
        }
        for fi20 := 0; fi20 < 18; fi20++ {
            fmt.Println(v3, v1)
            _ = fi20
        }
        switch v2 {
        case "":
            var v21 float32 = float32(0.1)
            _ = v21
        case "hello":
            for fi22 := 0; fi22 < 1; fi22++ {
                v2 = "a"
                fmt.Println(v1)
                _ = fi22
            }
        case "foo bar":
            v13 += (v13 + v13)
        default:
            for fi23 := 0; fi23 < 20; fi23++ {
                fmt.Println(fi23, v1)
                fmt.Println(fi23, fi23, v3)
                _ = fi23
            }
        }
    } else {
        v3 = (v3 && v3)
    }
    if v3 {
        switch "\n\t" {
        case "\n\t":
            v24 := []int{1, 1, 42, 399}
            _ = v24
        default:
            v13 = v13
        }
        var v25 uint16 = 740
        var v26 string = v2
        _ = v25
        _ = v26
    }
    var v27 float32 = ((3.14159 / 3.14159) / -2.25)
    var v28 int16 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
