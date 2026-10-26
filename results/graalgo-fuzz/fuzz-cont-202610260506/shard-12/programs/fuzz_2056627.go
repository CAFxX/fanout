package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 bool }
type GenS2 struct { F0 int64; F1 int16 }

func main() {
    var v1 uint64 = 0
    var v2 string = "hello"
    var v3 bool = (v2 <= v2)
    var v4 uint16 = 255
    v5 := make(map[int]uint32)
    var v6 float32 = -2.25
    var v7 int32 = 804
    fmt.Println(v4, v4)
    var v8 int32 = (v7 - -1)
    for fi9 := 0; fi9 < 6; fi9++ {
        for fi10 := 0; fi10 < 1; fi10++ {
            v8 -= 2147483647
            v1 += (v1 >> ((876) & 63))
            v7 -= -2147483648
            var v11 bool = (v6 != float32(0.1))
            var v12 int8 = -128
            _ = fi10
            _ = v11
            _ = v12
        }
        v13 := []int{(fi9 % (((fi9 & fi9)) | 1)), (fi9 + (fi9 * fi9))}
        for fi14 := 0; fi14 < 8; fi14++ {
            if false {
                switch (("hello" + "foo bar") + (v2 + v2)) {
                case "":
                    v2 = v2
                default:
                    switch (fi14 >> ((uint(v4)) & 63)) {
                    case -2147483648:
                        switch fi9 {
                        case 32767:
                            fmt.Println(v6)
                            var v15 uint16 = (v4 - (v4 & (v4 / ((v4) | 1))))
                            _ = v15
                        }
                    case -509:
                        v6 = 1.5
                    }
                }
            } else {
                if (1 >= 0) {
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                    var v16 uint = 42
                    if (!(2 >= 2)) {
                        var v17 int = (fi14 >> ((v16) & 63))
                        fmt.Println(fi14, v4)
                        _ = v17
                    } else {
                        var v18 int16 = -1
                        v19 := GenS1{F0: (fi9 << (((v16 << (((v16 & v16)) & 63))) & 63)), F1: -1, F2: (v1 > v1)}
                        _ = len(v5)
                        _ = v18
                        _ = v19
                    }
                    var v20 uint = (v16 | v16)
                    _ = v16
                    _ = v20
                } else {
                    v1 += (v1 | v1)
                }
            }
            _ = fi14
        }
        _ = fi9
        _ = v13
    }
    for fi21 := 0; fi21 < 15; fi21++ {
        var v22 uint32 = 42
        _ = fi21
        _ = v22
    }
    fmt.Println(v8, v1, v1)
    fmt.Println(v3, v8, v2)
    var v23 bool = false
    v2 = (("\n\t" + "foo bar") + v2)
    if false {
        var v24 int8 = -1
        v25 := GenS1{F0: -513, F1: 42, F2: v23}
        if true {
            if (1000 > -32768) {
                var v26 float32 = (v6 / -2.25)
                var v27 int = 509
                fmt.Println(v4, v3)
                _ = v26
                _ = v27
            } else {
                v28 := GenS1{F0: 0, F1: 42, F2: v3}
                _ = v28
            }
        } else {
            v29 := GenS1{F0: int(v1), F1: 0, F2: (-2147483648 <= 0)}
            _ = v29
        }
        _ = v24
        _ = v25
    }
    fmt.Println(v1, v1)
    v30 := GenS2{F0: -72, F1: 237}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
