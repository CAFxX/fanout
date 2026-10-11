package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16; F1 int32 }

func main() {
    v1 := []int{128}
    v2 := make(map[int]bool)
    var v3 string = "a"
    var v4 bool = true
    var v5 float32 = float32((float64(0.0) / 1.5))
    var v6 uint64 = 4294967295
    if len(v1) > 0 {
        v1[0] = -752
    }
    v4 = v4
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 uint32 = 65535
        _ = len(v2)
        v1 = append(v1, -9223372036854775808)
        _ = fi7
        _ = v8
    }
    var v9 float32 = v5
    v5 -= (3.14159 + ((1e10 / (((-2.25 + float32(((1e10 - (1.5 + (1e10 + 1e10))) + 1.5))) - v5) + (v5 / 1.5))) - 1e10))
    for fi10 := 0; fi10 < 2; fi10++ {
        v4 = v4
        var v11 int64 = 42
        _ = len(v2)
        _ = fi10
        _ = v11
    }
    if (!v4) {
        var v12 bool = false
        for fi13 := 0; fi13 < 5; fi13++ {
            v5 += v5
            if ((0 <= 1000) && (!v4)) {
                _ = len(v2)
                var v14 uint8 = 1
                _ = len(v1)
                _ = v14
            }
            _ = fi13
        }
        v1 = append(v1, 32767)
        _ = v12
    }
    for fi15 := 0; fi15 < 16; fi15++ {
        var v16 int = (-(fi15))
        var v17 int16 = -861
        _ = fi15
        _ = v16
        _ = v17
    }
    v18 := v2[4294967295]
    v19 := GenS1{F0: ((0 <= 9223372036854775807) || (-1 >= -35)), F1: 6, F2: 4294967295}
    if true {
        for fi20 := 0; fi20 < 15; fi20++ {
            v9 = float32(0.0)
            delete(v2, (-(fi20)))
            v3 = ""
            _ = fi20
        }
        var v21 float64 = 0.1
        _ = v21
    } else {
        v22 := GenS2{F0: 42}
        v18 = (true && (1 <= -1))
        v4 = true
        _ = v22
    }
    var v23 uint8 = 2
    for fi24 := 0; fi24 < 12; fi24++ {
        var v25 float64 = float64(0.0)
        var v26 bool = v4
        for fi27 := 0; fi27 < 11; fi27++ {
            v28 := GenS3{F0: 42, F1: 1}
            var v29 float32 = float32(1.5)
            switch 0 {
            case 34:
                if ((v9 - v5) >= 1e10) {
                    var v30 uint32 = 65535
                    _ = v30
                }
            }
            _ = fi27
            _ = v28
            _ = v29
        }
        _ = fi24
        _ = v25
        _ = v26
    }
    if (42 != 1) {
        var v31 uint16 = 32767
        v18 = true
        _ = v31
    } else {
        var v32 uint16 = 65535
        _ = v32
    }
    for fi33 := 0; fi33 < 2; fi33++ {
        if v18 {
            v4 = false
        }
        fmt.Println(v6, v5, fi33)
        v6 += (^(v6))
        _ = fi33
    }
    v34 := GenS3{F0: -1, F1: 0}
    fmt.Println(v18, v9)
    v35 := GenS1{F0: (-1 != int8(v6)), F1: 127, F2: 222}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
