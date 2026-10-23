package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 int16 }

func main() {
    v1 := []uint32{2147483647, 1, 4294967295}
    var v2 string = ("" + (("\n\t" + "a") + ("" + "")))
    v3 := []uint32{}
    v4 := make(map[uint32]bool)
    var v5 int32 = 386
    var v6 uint16 = 895
    var v7 bool = true
    _ = len(v3)
    switch -822 {
    case 789:
        v1 = append(v1, 2)
    case -1:
        fmt.Println(v6, v7, v2)
    }
    fmt.Println(v5, v7, v5)
    var v8 float32 = 3.14159
    var v9 uint32 = 2
    fmt.Println(v6, v8)
    var v10 int64 = 0
    var v11 uint8 = 17
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v11 += (v11 + 42)
    if (1.5 >= (0.1 / (v8 / ((-2.25 / v8) + 0.1)))) {
        for fi12 := 0; fi12 < 9; fi12++ {
            v13 := GenS1{F0: 42}
            _ = fi12
            _ = v13
        }
    }
    v14 := GenS2{F0: 678, F1: (v9 % ((v9) | 1))}
    if v7 {
        switch (v2 + ((v2 + v2) + "")) {
        case "hello":
            var v15 int = 32768
            _ = v15
        default:
            v9 = v9
        }
    } else {
        v7 = true
        var v16 uint8 = v11
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v16
    }
    for fi17 := 0; fi17 < 4; fi17++ {
        if (false || v7) {
            for fi18 := 0; fi18 < 14; fi18++ {
                var v19 int16 = 42
                _ = fi18
                _ = v19
            }
            if (0 >= -65) {
                if ((^(v5)) >= (^(v5))) {
                    v5 -= (v5 % (((v5 >> ((1) & 63))) | 1))
                }
                v20 := GenS3{F0: 0, F1: 1}
                _ = v20
            }
        } else {
            v8 = v8
            if len(v3) > 0 {
                v3[0] = (^(v9))
            }
        }
        v2 = "foo bar"
        v9 += v9
        _ = fi17
    }
    for fi21 := 0; fi21 < 3; fi21++ {
        var v22 uint64 = 984
        v23 := GenS3{F0: 388, F1: 42}
        if v7 {
            fmt.Println(v7, v11)
            if (!v7) {
                v24 := GenS4{F0: 407}
                var v25 float64 = -2.25
                _ = len(v1)
                _ = v24
                _ = v25
            }
            var v26 int8 = 28
            _ = v26
        }
        _ = fi21
        _ = v22
        _ = v23
    }
    if true {
        for fi27 := 0; fi27 < 4; fi27++ {
            v28 := GenS3{F0: 65535, F1: -247}
            if true {
                var v29 uint8 = (v11 - (v11 * v11))
                _ = v29
            }
            _ = fi27
            _ = v28
        }
        v7 = false
    } else {
        if true {
            v2 = v2
            if len(v1) > 0 {
                v1[0] = (^(v9))
            }
            var v30 uint32 = v9
            _ = v30
        }
        v1 = append(v1, v9)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
