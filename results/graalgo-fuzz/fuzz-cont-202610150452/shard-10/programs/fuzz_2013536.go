package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int16 }
type GenS2 struct { F0 bool; F1 uint; F2 string }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 uint64 }
func genFn0() uint32 {
    var v5 float32 = (((3.14159 / (((float32(1.5) / 1.5) - 0.0) + -2.25)) / -2.25) + -2.25)
    var v6 bool = true
    if v6 {
        var v7 int = 32768
        fmt.Println(v7, v6)
        _ = v7
    } else {
        v8 := GenS1{F0: 2, F1: 42, F2: 1000}
        v6 = (v6 || v6)
        _ = v8
    }
    v9 := GenS1{F0: 255, F1: 0, F2: -797}
    v10 := GenS2{F0: v6, F1: 4294967295, F2: ("" + (((("a" + ("hello" + "")) + ("\n\t" + "foo bar")) + "foo bar") + "\n\t"))}
    _ = v10
    _ = v5
    _ = v6
    _ = v9
    return 2
}
func genFn1(p0 bool, p1 int8) uint8 {
    for fi11 := 0; fi11 < 12; fi11++ {
        v12 := []int{(fi11 & (fi11 / ((-330) | 1)))}
        _ = fi11
        _ = v12
    }
    p0 = (-428 < 0)
    p0 = true
    fmt.Println(p1, p0)
    var v13 uint16 = 0
    _ = p0
    _ = p1
    _ = v13
    return 128
}

func main() {
    var v1 float64 = 0.1
    var v2 int = -1
    var v3 int16 = -1
    v4 := []int{int(v3), v2}
    if len(v4) > 0 {
        v4[0] = v2
    }
    var v14 bool = (100 < 42)
    v15 := GenS3{F0: v2, F1: 235}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v1 = v1
    v4 = append(v4, v2)
    if v14 {
        switch (v2 ^ 128) {
        case 42:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case 343:
            fmt.Println(v2)
        }
        for fi16 := 0; fi16 < 2; fi16++ {
            for fi17 := 0; fi17 < 18; fi17++ {
                if v14 {
                    var v18 uint32 = 1
                    var v19 uint16 = 0
                    v20 := GenS4{F0: 1}
                    _ = v18
                    _ = v19
                    _ = v20
                }
                _ = fi17
            }
            var v21 float32 = -2.25
            _ = fi16
            _ = v21
        }
    }
    var v22 int32 = 862
    v14 = v14
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
