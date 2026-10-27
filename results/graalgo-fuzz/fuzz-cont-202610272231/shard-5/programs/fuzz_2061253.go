package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16; F1 uint64; F2 int64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint8; F1 int32 }
func genFn0(p0 int16, p1 bool, p2 int8) int16 {
    var v6 string = "hello"
    if ((0.1 / -2.25) > (1.5 - 0.0)) {
        p1 = (false || false)
    }
    var v7 string = v6
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -32768
}

func main() {
    v1 := []string{"\n\t"}
    var v2 uint64 = 2
    var v3 int8 = 42
    var v4 int32 = -237
    v5 := make(map[string]string)
    v8 := GenS1{F0: "\n\t", F1: 42}
    var v9 int16 = 1
    fmt.Println(v9, v9)
    var v10 uint8 = 1
    if (42 < 42) {
        var v11 int64 = 0
        _ = v11
    } else {
        v2 -= 207
        fmt.Println(v10, v4)
        v12 := GenS2{F0: v4}
        _ = v12
    }
    if true {
        fmt.Println(v2, v2, v9)
        fmt.Println(v2, v4, v10)
        v9 = (v9 >> ((2) & 63))
    } else {
        for fi13 := 0; fi13 < 7; fi13++ {
            fmt.Println(v2, v4, v2)
            if false {
                var v14 uint16 = uint16(v10)
                var v15 int16 = (v9 - genFn0(genFn0((v9 * v9), false, v3), false, (v3 | (v3 ^ 127))))
                _ = v14
                _ = v15
            }
            _ = fi13
        }
    }
    v2 = v2
    var v16 int = 2147483647
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if true {
        var v17 int = (v16 + v16)
        v18 := GenS3{F0: 255, F1: v2, F2: 42}
        v19 := GenS4{F0: v16}
        _ = v17
        _ = v18
        _ = v19
    } else {
        if (true || ("a" > "\n\t")) {
            for fi20 := 0; fi20 < 7; fi20++ {
                v4 += (-(v4))
                if false {
                    v21 := GenS1{F0: ("foo bar" + "\n\t"), F1: 2}
                    _ = v21
                } else {
                    v22 := GenS5{F0: v10, F1: (v4 >> ((1) & 63))}
                    for fi23 := 0; fi23 < 3; fi23++ {
                        var v24 string = ""
                        _ = fi23
                        _ = v24
                    }
                    _ = v22
                }
                v25 := GenS4{F0: (fi20 >> ((0) & 63))}
                _ = fi20
                _ = v25
            }
        }
        for fi26 := 0; fi26 < 9; fi26++ {
            fmt.Println(fi26)
            _ = fi26
        }
        for fi27 := 0; fi27 < 19; fi27++ {
            _ = len(v5)
            switch v16 {
            case -355:
                for fi28 := 0; fi28 < 18; fi28++ {
                    _ = len(v1)
                    delete(v5, "\n\t")
                    var v29 int = int(v4)
                    _ = fi28
                    _ = v29
                }
            }
            _ = fi27
        }
    }
    fmt.Println(v16)
    if false {
        v4 += 42
        for fi30 := 0; fi30 < 5; fi30++ {
            var v31 uint16 = 0
            _ = fi30
            _ = v31
        }
        _ = len(v5)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
