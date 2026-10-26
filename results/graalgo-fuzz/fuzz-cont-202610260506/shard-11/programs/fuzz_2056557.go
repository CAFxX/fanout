package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 int64 }
func genFn0(p0 int8, p1 uint16, p2 uint64) uint {
    if false {
        switch -9223372036854775808 {
        case 32768:
            v5 := GenS1{F0: -86}
            _ = v5
        case 0:
            for fi6 := 0; fi6 < 17; fi6++ {
                p0 += (p0 / (((^(p0))) | 1))
                _ = fi6
            }
        }
        var v7 int16 = int16(p2)
        _ = v7
    }
    if true {
        var v8 int8 = (p0 >> ((4294967295) & 63))
        _ = v8
    }
    v9 := GenS2{F0: uint8(p1), F1: (p2 << ((255) & 63))}
    if false {
        var v10 uint8 = 127
        var v11 string = ("foo bar" + ("hello" + "foo bar"))
        _ = v10
        _ = v11
    }
    p0 += p0
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 4294967295
}

func main() {
    var v1 uint = 4294967295
    var v2 uint16 = 2
    v3 := make(map[int]int)
    v4 := make(map[string]string)
    for fi12 := 0; fi12 < 15; fi12++ {
        var v13 int = fi12
        var v14 int16 = 32767
        if ((0.0 / 3.14159) != 1e10) {
            if true {
                fmt.Println(v14, v14, v2)
                fmt.Println(v14)
            }
            v13 -= fi12
            fmt.Println(v13, v1, v2)
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    fmt.Println(v2, v2)
    if true {
        var v15 string = ("\n\t" + "a")
        _ = v15
    } else {
        var v16 int64 = -1
        _ = v16
    }
    v17 := GenS3{F0: -191}
    var v18 uint = (^(v1))
    for fi19 := 0; fi19 < 1; fi19++ {
        fmt.Println(v2)
        _ = fi19
    }
    v20 := GenS3{F0: 9223372036854775807}
    for fi21 := 0; fi21 < 13; fi21++ {
        switch "" {
        case "foo bar":
            v22 := GenS3{F0: -9223372036854775808}
            _ = v22
        case "a":
            if true {
                v1 = genFn0(100, v2, 18446744073709551615)
                if true {
                    var v23 bool = true
                    v23 = false
                    var v24 float64 = -2.25
                    _ = v23
                    _ = v24
                }
            } else {
                v1 = (v18 & genFn0(1, (v2 >> ((v1) & 63)), 0))
            }
        case "":
            v1 += v18
        }
        _ = fi21
    }
    _ = len(v3)
    var v25 float32 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v4
}
