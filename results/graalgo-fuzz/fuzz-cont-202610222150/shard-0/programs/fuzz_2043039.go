package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 int32; F1 uint64 }
func genFn0(p0 uint32, p1 int, p2 uint) bool {
    p2 = (p2 & 4294967295)
    var v7 bool = true
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return v7
}

func main() {
    var v1 int = -2147483648
    var v2 float32 = 1.5
    v3 := make(map[int]string)
    var v4 uint32 = 582
    v5 := make(map[uint32]int)
    var v6 uint32 = (v4 / ((4294967295) | 1))
    fmt.Println(v4)
    switch "a" {
    case "":
        v8 := GenS1{F0: -907}
        _ = v8
    case "foo bar":
        for fi9 := 0; fi9 < 16; fi9++ {
            var v10 uint32 = v4
            switch (-(fi9)) {
            case 128:
                v11 := GenS1{F0: 1000}
                _ = v11
            case -577:
                fmt.Println(fi9, v6, fi9)
            }
            for fi12 := 0; fi12 < 17; fi12++ {
                v4 -= (v10 + (v10 & (v4 & v4)))
                switch -9223372036854775808 {
                case 42:
                    var v13 bool = true
                    _ = v13
                case 32768:
                    var v14 int32 = -449
                    _ = v14
                }
                v15 := v3[v1]
                _ = fi12
                _ = v15
            }
            _ = fi9
            _ = v10
        }
    }
    if (!false) {
        fmt.Println(v4, v6)
        var v16 string = "a"
        _ = v16
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        var v18 uint = uint(v4)
        _ = fi17
        _ = v18
    }
    var v19 uint32 = (v6 << ((4294967295) & 63))
    v2 = (v2 / 3.14159)
    fmt.Println(v19)
    v20 := GenS2{F0: true, F1: 587}
    var v21 string = "\n\t"
    for fi22 := 0; fi22 < 13; fi22++ {
        v23 := GenS3{F0: -2147483648, F1: 2}
        _ = fi22
        _ = v23
    }
    var v24 bool = false
    v25 := GenS3{F0: 1, F1: 158}
    var v26 bool = (255 < 255)
    switch (^(v1)) {
    case -2147483648:
        fmt.Println(v4, v24, v26)
    case -9223372036854775808:
        var v27 int = v1
        _ = v27
    case 0:
        if true {
            v26 = ((!true) || (!(!(v4 < v6))))
            v28 := v3[947]
            _ = v28
        } else {
            v26 = v24
            var v29 uint = 832
            _ = v29
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
