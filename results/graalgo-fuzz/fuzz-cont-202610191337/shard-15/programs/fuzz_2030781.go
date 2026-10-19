package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 bool }
type GenS2 struct { F0 uint32; F1 int16 }

func main() {
    v1 := make(map[int]string)
    var v2 int16 = 1000
    var v3 float32 = -2.25
    var v4 uint64 = 520
    var v5 int = -1
    v4 = 745
    var v6 string = "a"
    v6 = ""
    v7 := GenS1{F0: 247, F1: 394, F2: true}
    v8 := GenS2{F0: 2147483648, F1: v2}
    v3 = 3.14159
    v9 := v1[v5]
    for fi10 := 0; fi10 < 8; fi10++ {
        if (-80 > 1) {
            for fi11 := 0; fi11 < 1; fi11++ {
                var v12 int64 = 952
                _ = fi11
                _ = v12
            }
        } else {
            v1[fi10] = "hello"
            if true {
                v9 = v6
            }
        }
        _ = fi10
    }
    v6 = "foo bar"
    var v13 bool = true
    for fi14 := 0; fi14 < 18; fi14++ {
        for fi15 := 0; fi15 < 14; fi15++ {
            var v16 int16 = (v2 ^ (v2 / ((v2) | 1)))
            var v17 uint32 = 326
            _ = fi15
            _ = v16
            _ = v17
        }
        v4 = v4
        v3 -= 0.0
        _ = fi14
    }
    var v18 uint16 = 725
    v19 := []int{}
    fmt.Println(v9, v5, v2)
    switch (v5 * v5) {
    case -9223372036854775808:
        for fi20 := 0; fi20 < 9; fi20++ {
            for fi21 := 0; fi21 < 19; fi21++ {
                v18 = uint16(v4)
                var v22 uint = 777
                var v23 uint8 = 42
                v24 := GenS1{F0: (^(v23)), F1: (v4 + (v4 << (((v22 * v22)) & 63))), F2: ((!(9223372036854775807 == 178)) || (!false))}
                _ = fi21
                _ = v22
                _ = v23
                _ = v24
            }
            _ = fi20
        }
    case -600:
        v2 -= -425
    }
    if (779 >= 764) {
        for fi25 := 0; fi25 < 16; fi25++ {
            if (!(!true)) {
                fmt.Println(v9)
                var v26 uint16 = v18
                if len(v19) > 0 {
                    v19[0] = (fi25 ^ 0)
                }
                _ = v26
            } else {
                switch ("" + v9) {
                case "":
                    var v27 bool = (4294967295 < 707)
                    _ = v27
                case "hello":
                    v18 = (v18 & 32767)
                }
                v28 := v1[(fi25 / ((fi25) | 1))]
                _ = v28
            }
            var v29 int16 = v2
            v13 = false
            _ = fi25
            _ = v29
        }
        for fi30 := 0; fi30 < 4; fi30++ {
            v31 := GenS2{F0: 65535, F1: v2}
            v3 = v3
            v4 = 2
            _ = fi30
            _ = v31
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
