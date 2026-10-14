package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 uint }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool }
func genFn0(p0 uint, p1 int8) uint8 {
    if (p1 != (p1 >> ((18446744073709551615) & 63))) {
        switch "foo bar" {
        case "a":
            p0 = 0
        }
        v6 := []int{}
        if true {
            p1 -= p1
            for fi7 := 0; fi7 < 10; fi7++ {
                var v8 uint16 = 255
                v9 := GenS1{F0: 118, F1: 4294967295, F2: (^(p0))}
                _ = fi7
                _ = v8
                _ = v9
            }
        }
        _ = v6
    }
    if (true || (-670 != -675)) {
        for fi10 := 0; fi10 < 15; fi10++ {
            v11 := GenS2{F0: (p0 * p0), F1: 399}
            _ = fi10
            _ = v11
        }
    }
    v12 := []int{-601, -230}
    _ = len(v12)
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v12
    return 33
}

func main() {
    var v1 int = 2147483647
    var v2 int16 = 0
    v3 := []uint32{65535, 65535, 65535, 979}
    var v4 int64 = 9223372036854775807
    var v5 int64 = v4
    fmt.Println(v5, v4, v2)
    for fi13 := 0; fi13 < 20; fi13++ {
        fmt.Println(v5, v4, v2)
        v4 += v5
        _ = fi13
    }
    v3 = append(v3, 1)
    v4 = -391
    if (("hello" + ("hello" + ("\n\t" + ("a" + ("foo bar" + "foo bar"))))) >= "") {
        if false {
            fmt.Println(v1, v1)
        }
        if len(v3) > 0 {
            v3[0] = 0
        }
    }
    for fi14 := 0; fi14 < 20; fi14++ {
        var v15 float32 = 0.0
        _ = fi14
        _ = v15
    }
    var v16 uint32 = 994
    switch 2147483647 {
    case 777:
        v2 = 1000
    case 1:
        v4 -= v4
    case 62:
        if false {
            if true {
                v17 := GenS3{F0: 42}
                _ = v17
            }
            for fi18 := 0; fi18 < 18; fi18++ {
                v19 := GenS4{F0: true}
                switch 127 {
                case -9223372036854775808:
                    switch fi18 {
                    case 32768:
                        var v20 uint16 = 472
                        _ = v20
                    default:
                        var v21 int64 = v5
                        _ = v21
                    }
                }
                _ = fi18
                _ = v19
            }
            if false {
                var v22 string = ("\n\t" + "")
                _ = v22
            }
        } else {
            switch "\n\t" {
            case "foo bar":
                var v23 int8 = 100
                _ = v23
            case "hello":
                v16 = v16
            }
        }
    }
    v16 = (v16 >> ((255) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
