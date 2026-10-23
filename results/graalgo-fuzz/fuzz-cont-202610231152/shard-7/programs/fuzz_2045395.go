package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8; F1 uint8; F2 uint32 }
type GenS5 struct { F0 string }

func main() {
    v1 := []uint32{0, 2147483647, 59}
    var v2 string = "hello"
    var v3 uint64 = 0
    var v4 bool = (2 != 42)
    var v5 bool = (1 >= 32767)
    v6 := GenS1{F0: 2, F1: -101}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v7 int8 = 127
    var v8 uint16 = 42
    if (418 < 117) {
        for fi9 := 0; fi9 < 13; fi9++ {
            switch v2 {
            case "hello":
                v10 := GenS1{F0: 255, F1: 40}
                _ = v10
            case "":
                var v11 string = (("a" + (v2 + v2)) + "a")
                _ = v11
            case "foo bar":
                if len(v1) > 0 {
                    v1[0] = 65535
                }
            }
            _ = fi9
        }
        for fi12 := 0; fi12 < 13; fi12++ {
            if (42 < 4294967295) {
                v8 = 324
            } else {
                for fi13 := 0; fi13 < 9; fi13++ {
                    fmt.Println(fi13)
                    v14 := GenS2{F0: 0, F1: uint(v3)}
                    v4 = v5
                    _ = fi13
                    _ = v14
                }
            }
            var v15 uint8 = 255
            switch v2 {
            case "a":
                var v16 bool = (-2147483648 >= 42)
                _ = v16
            case "":
                v1 = append(v1, 4294967295)
            case "foo bar":
                for fi17 := 0; fi17 < 17; fi17++ {
                    if true {
                        v18 := GenS3{F0: 65535}
                        var v19 int = 1
                        _ = v18
                        _ = v19
                    }
                    if v5 {
                        var v20 int64 = 0
                        _ = v20
                    } else {
                        v21 := GenS4{F0: v7, F1: (v15 << ((850) & 63)), F2: 108}
                        _ = v21
                    }
                    v1 = append(v1, 140)
                    _ = fi17
                }
            default:
                v22 := GenS4{F0: (v7 - (v7 + v7)), F1: (^(v15)), F2: uint32(v15)}
                _ = v22
            }
            _ = fi12
            _ = v15
        }
        v5 = (v4 || v4)
    }
    if (377 == v3) {
        for fi23 := 0; fi23 < 13; fi23++ {
            var v24 float32 = (0.0 * float32((1.5 * 0.1)))
            switch 9223372036854775807 {
            case -2147483648:
                v3 += (v3 << ((42) & 63))
            case 1:
                if len(v1) > 0 {
                    v1[0] = 65535
                }
            }
            _ = fi23
            _ = v24
        }
    }
    var v25 uint8 = 128
    fmt.Println(v5, v3)
    v26 := GenS2{F0: (v25 + v25), F1: 65535}
    if v4 {
        v27 := GenS5{F0: v2}
        for fi28 := 0; fi28 < 9; fi28++ {
            var v29 float32 = 1e10
            _ = fi28
            _ = v29
        }
        if false {
            switch v2 {
            case "hello":
                v5 = (42 != 0)
            case "\n\t":
                if (v4 && v5) {
                    for fi30 := 0; fi30 < 7; fi30++ {
                        fmt.Println(v7, v5, v3)
                        v31 := GenS3{F0: 42}
                        _ = fi30
                        _ = v31
                    }
                    v32 := GenS3{F0: 65535}
                    switch v2 {
                    case "\n\t":
                        var v33 int32 = -1
                        _ = v33
                    case "":
                        for fi34 := 0; fi34 < 19; fi34++ {
                            v35 := GenS4{F0: (v7 % ((0) | 1)), F1: 128, F2: 42}
                            v25 = 0
                            v36 := GenS2{F0: 42, F1: 42}
                            _ = fi34
                            _ = v35
                            _ = v36
                        }
                    }
                    _ = v32
                }
            default:
                v3 = 42
            }
            var v37 float64 = 3.14159
            _ = v37
        }
        _ = v27
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v38 bool = (!(1e10 < (3.14159 - float32(1.5))))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
