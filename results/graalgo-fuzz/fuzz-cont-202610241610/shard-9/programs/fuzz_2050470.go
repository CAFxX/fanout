package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 int = 4294967295
    var v2 string = ("hello" + "a")
    var v3 int32 = 1
    v3 = 42
    for fi4 := 0; fi4 < 17; fi4++ {
        v1 += 9223372036854775807
        if (3.14159 != (((((3.14159 + -2.25) - 3.14159) / (float64(((float32(0.1) + (1e10 - float32(1e10))) + 3.14159)) + -2.25)) / 1e10) + 3.14159)) {
            v2 = v2
            for fi5 := 0; fi5 < 6; fi5++ {
                var v6 string = (v2 + "")
                v1 += (fi5 - fi4)
                v1 = fi5
                _ = fi5
                _ = v6
            }
            v7 := GenS1{F0: v3}
            _ = v7
        }
        _ = fi4
    }
    var v8 int16 = 1000
    v9 := []int{-2147483648}
    if (!true) {
        var v10 int32 = (v3 & (v3 / (((v3 / ((v3) | 1))) | 1)))
        _ = v10
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        if ((0.1 / ((0.1 - 1e10) - 0.1)) <= (3.14159 * (float64((1.5 * 1.5)) / 3.14159))) {
            var v12 float64 = 0.1
            _ = v12
        } else {
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
        }
        var v13 string = "hello"
        _ = fi11
        _ = v13
    }
    v8 = (-(v8))
    if (v1 <= v1) {
        v14 := GenS1{F0: 42}
        var v15 string = "foo bar"
        var v16 bool = true
        _ = v14
        _ = v15
        _ = v16
    } else {
        v8 = v8
        v17 := GenS1{F0: v3}
        v18 := GenS1{F0: (v3 / (((v3 + v3)) | 1))}
        _ = v17
        _ = v18
    }
    v19 := GenS1{F0: (v3 >> ((65535) & 63))}
    _ = len(v9)
    if len(v9) > 0 {
        v9[0] = 98
    }
    v3 = 0
    for fi20 := 0; fi20 < 7; fi20++ {
        v2 = v2
        fmt.Println(v8, v8)
        for fi21 := 0; fi21 < 5; fi21++ {
            v9 = append(v9, (^(v1)))
            var v22 uint32 = 4294967295
            var v23 int16 = v8
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = fi20
    }
    _ = len(v9)
    var v24 int16 = (v8 ^ (-(v8)))
    var v25 int = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v8
    _ = v9
}
