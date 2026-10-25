package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int }
func genFn0(p0 int64, p1 uint32, p2 uint32) int16 {
    p0 -= (p0 + p0)
    switch "" {
    case "a":
        if ((0.0 > 3.14159) && (!(!false))) {
            fmt.Println(p1)
            var v5 uint = 65535
            _ = v5
        } else {
            fmt.Println(p1, p1, p0)
            switch 32767 {
            case 862:
                for fi6 := 0; fi6 < 14; fi6++ {
                    v7 := []int{466, -1}
                    _ = fi6
                    _ = v7
                }
            case -2147483648:
                v8 := []int{1, 42, -980}
                _ = v8
            }
        }
    case "foo bar":
        for fi9 := 0; fi9 < 15; fi9++ {
            var v10 float32 = 0.0
            v11 := []int{4294967295, 127}
            if (!true) {
                v12 := GenS1{F0: 9223372036854775807, F1: (1e10 == -2.25), F2: -32768}
                v13 := GenS2{F0: 0}
                fmt.Println(v10)
                _ = v12
                _ = v13
            }
            _ = fi9
            _ = v10
            _ = v11
        }
    case "\n\t":
        if (!false) {
            p0 -= (p0 | 695)
        }
    }
    var v14 bool = false
    var v15 uint64 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    return 32767
}

func main() {
    var v1 string = ("a" + "hello")
    v2 := []string{v1, v1, (v1 + ""), "hello"}
    var v3 int32 = 42
    var v4 uint32 = 2
    var v16 int8 = -1
    if true {
        v17 := GenS1{F0: 9223372036854775807, F1: (255 != 255), F2: -874}
        var v18 int64 = 0
        _ = v17
        _ = v18
    } else {
        v16 -= (-(v16))
        v3 = (^(v3))
    }
    var v19 uint16 = 2
    for fi20 := 0; fi20 < 1; fi20++ {
        var v21 uint = 255
        _ = fi20
        _ = v21
    }
    v22 := GenS3{F0: -9223372036854775808}
    v1 = v1
    for fi23 := 0; fi23 < 12; fi23++ {
        v24 := GenS2{F0: 927}
        _ = fi23
        _ = v24
    }
    var v25 float32 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v22
    _ = v25
    _ = v3
    _ = v4
}
