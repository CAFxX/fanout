package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32 }

func main() {
    var v1 uint32 = 2147483647
    v2 := []uint32{v1, v1, v1, v1}
    v3 := []int{851, -1, -217}
    var v4 uint16 = 0
    var v5 float32 = float32(1.5)
    var v6 int16 = 1
    switch "a" {
    case "a":
        for fi7 := 0; fi7 < 9; fi7++ {
            if ((v1 | 0) < (^(v1))) {
                var v8 float32 = (v5 + v5)
                var v9 int16 = v6
                _ = v8
                _ = v9
            }
            v10 := GenS1{F0: 597, F1: ("" < (("hello" + "") + ("a" + ("foo bar" + "")))), F2: (^(v1))}
            var v11 uint64 = 907
            _ = fi7
            _ = v10
            _ = v11
        }
    case "foo bar":
        var v12 uint8 = 2
        _ = v12
    case "":
        var v13 bool = ((v5 != v5) && false)
        _ = v13
    }
    for fi14 := 0; fi14 < 17; fi14++ {
        v5 = (float32((3.14159 - float64((v5 * 1e10)))) / 1.0)
        v6 -= (v6 & -588)
        if false {
            switch "foo bar" {
            case "\n\t":
                if len(v3) > 0 {
                    v3[0] = (fi14 << ((1) & 63))
                }
            }
            fmt.Println(v5, fi14)
        }
        _ = fi14
    }
    switch "a" {
    case "":
        v4 = (^(v4))
    case "a":
        if (-1 < 113) {
            var v15 bool = true
            for fi16 := 0; fi16 < 1; fi16++ {
                v1 -= (v1 + uint32(v4))
                v15 = (0 == -584)
                _ = fi16
            }
            _ = v15
        }
    }
    for fi17 := 0; fi17 < 11; fi17++ {
        for fi18 := 0; fi18 < 8; fi18++ {
            v1 = v1
            _ = fi18
        }
        _ = fi17
    }
    v19 := GenS2{F0: 107}
    if true {
        var v20 float32 = v5
        var v21 float32 = v5
        _ = v20
        _ = v21
    }
    if false {
        for fi22 := 0; fi22 < 17; fi22++ {
            v23 := GenS3{F0: v1}
            _ = fi22
            _ = v23
        }
        v3 = append(v3, 4294967295)
    }
    _ = len(v2)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
