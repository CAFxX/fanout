package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 int }

func main() {
    v1 := make(map[uint32]string)
    var v2 int32 = -2147483648
    var v3 int = 32767
    v3 = 4294967295
    if true {
        v4 := []int{(v3 * v3), (v3 << ((4294967295) & 63))}
        v2 = (v2 ^ (v2 / ((v2) | 1)))
        _ = v4
    } else {
        v5 := GenS1{F0: 0, F1: v3}
        var v6 uint8 = 124
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 float64 = 1e10
        _ = fi7
        _ = v8
    }
    v1[1] = ("a" + ("\n\t" + "\n\t"))
    if true {
        for fi9 := 0; fi9 < 2; fi9++ {
            v3 = (^(v3))
            _ = fi9
        }
        var v10 int = (v3 * (v3 | (v3 % ((-316) | 1))))
        _ = v10
    }
    for fi11 := 0; fi11 < 9; fi11++ {
        for fi12 := 0; fi12 < 13; fi12++ {
            v13 := []int{1, fi11, 0, (v3 | fi11)}
            v14 := GenS1{F0: 6, F1: (-(fi11))}
            fmt.Println(v2)
            _ = fi12
            _ = v13
            _ = v14
        }
        v15 := GenS2{F0: (-32768 < 42), F1: 4294967295}
        fmt.Println(fi11)
        _ = fi11
        _ = v15
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        for fi17 := 0; fi17 < 2; fi17++ {
            fmt.Println(fi17)
            var v18 uint = 65535
            fmt.Println(fi17, v2, fi16)
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    v19 := []int{v3, (v3 | (v3 >> ((4294967295) & 63)))}
    if true {
        if (!(!(false || ((0.0 + float32(float64((((((0.0 / 0.1) * -2.25) - -2.25) / 3.14159) / 1.0)))) < (1.5 + float32(-2.25)))))) {
            v2 += (v2 * v2)
            for fi20 := 0; fi20 < 12; fi20++ {
                v3 = fi20
                _ = fi20
            }
            fmt.Println(v3, v2)
        }
        var v21 float32 = 3.14159
        _ = v21
    } else {
        var v22 bool = false
        for fi23 := 0; fi23 < 13; fi23++ {
            fmt.Println(v3)
            if len(v19) >= 2 {
                v19 = v19[1:len(v19)]
            }
            v22 = (0 != 302)
            _ = fi23
        }
        _ = v22
    }
    _ = len(v1)
    var v24 float32 = (-2.25 * (3.14159 + 3.14159))
    var v25 uint16 = 42
    v26 := GenS3{F0: 128}
    var v27 uint8 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
}
