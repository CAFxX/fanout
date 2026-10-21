package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 uint64 }

func main() {
    var v1 uint = 255
    var v2 string = "a"
    var v3 int8 = 94
    if (false && (true || (-9223372036854775808 >= 9223372036854775807))) {
        fmt.Println(v2)
        for fi4 := 0; fi4 < 20; fi4++ {
            if false {
                v1 = 18446744073709551615
                var v5 uint16 = 32767
                v3 += (v3 / ((v3) | 1))
                _ = v5
            }
            _ = fi4
        }
    } else {
        var v6 uint32 = 42
        _ = v6
    }
    var v7 float32 = 0.0
    fmt.Println(v1, v7, v3)
    var v8 string = ""
    v9 := []int{127, 4294967295}
    fmt.Println(v1, v2, v8)
    if false {
        if false {
            var v10 uint16 = 32767
            v11 := GenS1{F0: (v1 >> ((65535) & 63)), F1: 127}
            var v12 string = v8
            _ = v10
            _ = v11
            _ = v12
        } else {
            if true {
                var v13 uint64 = 4294967295
                v8 = v8
                _ = v13
            }
            if true {
                if true {
                    v8 = (v2 + v8)
                    v3 += (v3 ^ (v3 + (v3 ^ v3)))
                    v9 = append(v9, -9223372036854775808)
                } else {
                    var v14 uint16 = 506
                    var v15 int8 = 100
                    _ = v14
                    _ = v15
                }
                fmt.Println(v7, v8)
                v16 := GenS2{F0: -1, F1: 981}
                _ = v16
            } else {
                for fi17 := 0; fi17 < 2; fi17++ {
                    if false {
                        var v18 float32 = ((-2.25 + (float32((0.1 / -2.25)) * v7)) + float32(3.14159))
                        _ = v18
                    } else {
                        _ = len(v9)
                        fmt.Println(v7)
                    }
                    _ = fi17
                }
                v1 = (^(v1))
                switch 960 {
                case 128:
                    var v19 float32 = v7
                    _ = v19
                case 9223372036854775807:
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                }
            }
        }
    } else {
        fmt.Println(v2, v1, v2)
        if true {
            var v20 float32 = 3.14159
            fmt.Println(v8, v2)
            _ = v20
        } else {
            v1 = (^(v1))
            fmt.Println(v8)
            fmt.Println(v3, v2, v1)
        }
        v21 := GenS3{F0: 2, F1: 42}
        _ = v21
    }
    v1 -= v1
    if (((("a" + v8) + v8) + ("foo bar" + (v8 + "foo bar"))) != "a") {
        var v22 int16 = -1
        v2 = (("foo bar" + (v2 + v2)) + v2)
        var v23 int16 = (v22 >> ((42) & 63))
        _ = v22
        _ = v23
    }
    if (((2 <= 923) || true) && false) {
        var v24 int32 = 272
        _ = len(v9)
        _ = v24
    } else {
        fmt.Println(v8, v8)
        var v25 int32 = 19
        _ = v25
    }
    v2 = ""
    if (!true) {
        v26 := GenS4{F0: 2}
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
