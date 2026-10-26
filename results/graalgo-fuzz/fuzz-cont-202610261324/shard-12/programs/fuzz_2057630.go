package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 int16; F1 string; F2 int }
func genFn0(p0 bool, p1 int8, p2 bool) uint16 {
    p0 = (p2 && p0)
    p2 = (!(p2 || p2))
    v6 := []int{-561}
    v6 = append(v6, -2147483648)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 111
}
func genFn1(p0 uint32) uint {
    var v7 uint16 = 32767
    var v8 uint32 = p0
    _ = p0
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 float32 = 1e10
    var v2 bool = true
    var v3 int16 = 312
    v4 := []uint32{4294967295, 0}
    var v5 string = "\n\t"
    if true {
        v5 = v5
    } else {
        for fi9 := 0; fi9 < 7; fi9++ {
            v5 = (v5 + ((v5 + (((v5 + "a") + v5) + (((("hello" + "foo bar") + v5) + v5) + v5))) + (v5 + v5)))
            v3 -= -32768
            _ = fi9
        }
        var v10 float64 = 0.1
        _ = v10
    }
    var v11 uint8 = 2
    fmt.Println(v5, v5, v1)
    if (0 == -1) {
        if v2 {
            var v12 float32 = ((v1 - 0.1) * -2.25)
            switch "a" {
            case "hello":
                if v2 {
                    if true {
                        var v13 string = v5
                        var v14 uint = 1
                        v13 = ("a" + (((v5 + v13) + v5) + (v5 + ((v5 + ((v13 + (v13 + v5)) + v13)) + (v13 + v13)))))
                        _ = v13
                        _ = v14
                    } else {
                        var v15 string = v5
                        fmt.Println(v3, v1, v3)
                        var v16 uint64 = 18446744073709551615
                        _ = v15
                        _ = v16
                    }
                }
            case "":
                if len(v4) > 0 {
                    v4[0] = 0
                }
            case "a":
                v12 += v1
            default:
                var v17 uint64 = 139
                _ = v17
            }
            for fi18 := 0; fi18 < 1; fi18++ {
                fmt.Println(fi18, v5)
                v19 := GenS1{F0: 18446744073709551615, F1: (v5 + v5)}
                _ = fi18
                _ = v19
            }
            _ = v12
        }
        if len(v4) > 0 {
            v4[0] = 1
        }
    }
    v3 -= 726
    v20 := GenS2{F0: 877, F1: 100}
    var v21 int = 190
    if len(v4) > 0 {
        v4[0] = 2
    }
    var v22 int64 = 0
    for fi23 := 0; fi23 < 6; fi23++ {
        v24 := GenS2{F0: 0, F1: -114}
        switch v5 {
        case "hello":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case "":
            v5 = v5
        }
        _ = fi23
        _ = v24
    }
    var v25 int = v21
    v5 = v5
    switch -2147483648 {
    case 128:
        v26 := GenS3{F0: v3, F1: "", F2: v25}
        _ = v26
    case 321:
        var v27 int = v21
        _ = v27
    }
    fmt.Println(v22, v3, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
