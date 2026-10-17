package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8; F1 uint8; F2 string }
func genFn0(p0 int) bool {
    var v5 uint = 566
    fmt.Println(p0)
    if false {
        var v6 string = ("hello" + ("\n\t" + "\n\t"))
        var v7 uint16 = 42
        _ = v6
        _ = v7
    }
    var v8 uint = (^(v5))
    switch (p0 + (p0 | (p0 - int(v5)))) {
    case -647:
        p0 = p0
    case 643:
        var v9 float32 = 1.5
        _ = v9
    }
    _ = p0
    _ = v5
    _ = v8
    return true
}
func genFn1() uint {
    var v10 float64 = 1.5
    v11 := []int{1}
    for fi12 := 0; fi12 < 9; fi12++ {
        v13 := GenS1{F0: -1}
        var v14 uint = 255
        _ = fi12
        _ = v13
        _ = v14
    }
    _ = len(v11)
    _ = v10
    _ = v11
    return 42
}

func main() {
    var v1 uint8 = 0
    var v2 float64 = -2.25
    var v3 uint8 = v1
    v4 := []uint32{42}
    var v15 uint16 = 255
    fmt.Println(v1, v2, v1)
    v1 += 74
    v1 = (v3 % (((v3 | v3)) | 1))
    fmt.Println(v15)
    v4 = append(v4, 58)
    var v16 uint64 = 18446744073709551615
    var v17 bool = (!(826 != 32767))
    v18 := GenS1{F0: int16(v1)}
    v19 := GenS2{F0: (v3 ^ v1)}
    for fi20 := 0; fi20 < 7; fi20++ {
        for fi21 := 0; fi21 < 10; fi21++ {
            var v22 string = "hello"
            v23 := GenS3{F0: v1, F1: (^(v1)), F2: ((((v22 + "foo bar") + v22) + v22) + (v22 + (v22 + "")))}
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = fi20
    }
    for fi24 := 0; fi24 < 20; fi24++ {
        fmt.Println(v1)
        _ = fi24
    }
    switch 42 {
    case -393:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
