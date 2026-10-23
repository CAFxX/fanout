package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 string }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 uint; F1 uint8; F2 int8 }
func genFn0(p0 uint64, p1 int32, p2 bool) bool {
    fmt.Println(p2, p2, p2)
    var v4 int = -2147483648
    switch "hello" {
    case "a":
        v5 := GenS1{F0: -1, F1: 1, F2: "hello"}
        _ = v5
    case "foo bar":
        var v6 int64 = 0
        _ = v6
    case "hello":
        v4 = v4
    }
    for fi7 := 0; fi7 < 17; fi7++ {
        fmt.Println(fi7)
        v8 := []int{}
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return p2
}

func main() {
    var v1 uint64 = 0
    var v2 uint = 0
    v3 := make(map[uint32]uint32)
    for fi9 := 0; fi9 < 10; fi9++ {
        v2 = v2
        fmt.Println(v2)
        if false {
            v1 = (v1 - 731)
        } else {
            if true {
                var v10 uint64 = (v1 - (v1 + 389))
                _ = v10
            } else {
                var v11 float32 = 0.0
                _ = v11
            }
        }
        _ = fi9
    }
    v12 := v3[4294967295]
    for fi13 := 0; fi13 < 4; fi13++ {
        if false {
            var v14 bool = false
            _ = v14
        } else {
            v15 := v3[v12]
            fmt.Println(v12)
            _ = v15
        }
        for fi16 := 0; fi16 < 8; fi16++ {
            v17 := GenS2{F0: 255, F1: (fi13 | fi16)}
            _ = fi16
            _ = v17
        }
        _ = fi13
    }
    v18 := GenS2{F0: 2, F1: 4294967295}
    switch "\n\t" {
    case "a":
        switch 554 {
        case -648:
            v19 := []int{-9223372036854775808, 32767}
            _ = v19
        case 300:
            v2 -= v2
        }
    case "\n\t":
        var v20 bool = false
        _ = v20
    }
    if true {
        v1 = (v1 - v1)
    } else {
        v1 = (v1 ^ v1)
        for fi21 := 0; fi21 < 1; fi21++ {
            var v22 int = 32768
            _ = fi21
            _ = v22
        }
    }
    var v23 string = ("foo bar" + "")
    switch v23 {
    case "":
        if false {
            if false {
                if true {
                    v24 := []int{-9223372036854775808, -670, int(v12)}
                    if len(v24) >= 2 {
                        v24 = v24[1:len(v24)]
                    }
                    _ = v24
                } else {
                    v25 := GenS1{F0: 73, F1: v12, F2: v23}
                    _ = v25
                }
                v1 = v1
            } else {
                var v26 float64 = ((-2.25 / -2.25) + (0.1 / (float64((0.1 - float32((float64(-2.25) - 1.5)))) * 1.5)))
                _ = v26
            }
            v27 := GenS3{F0: v2, F1: 255, F2: 127}
            delete(v3, v12)
            _ = v27
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v18
    _ = v2
    _ = v23
    _ = v3
}
