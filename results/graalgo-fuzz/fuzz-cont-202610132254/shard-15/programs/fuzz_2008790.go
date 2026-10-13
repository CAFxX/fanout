package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 int64; F1 uint16; F2 uint32 }
type GenS3 struct { F0 int8; F1 uint16; F2 int8 }
func genFn0(p0 uint, p1 bool) int64 {
    p0 = (p0 * (^(p0)))
    var v5 int8 = 123
    p0 -= p0
    _ = p0
    _ = p1
    _ = v5
    return -1
}

func main() {
    v1 := make(map[string]string)
    var v2 int32 = 42
    var v3 string = "hello"
    var v4 int = -1
    switch 128 {
    case 32768:
        for fi6 := 0; fi6 < 13; fi6++ {
            v7 := GenS1{F0: 1, F1: 0}
            _ = fi6
            _ = v7
        }
    case 2147483647:
        v4 += (v4 >> ((844) & 63))
    case -2147483648:
        var v8 uint64 = 388
        _ = v8
    }
    v9 := []int{(v4 % ((456) | 1)), v4, (v4 >> ((42) & 63)), v4}
    var v10 uint32 = 915
    if (false || false) {
        var v11 uint16 = 364
        _ = v11
    } else {
        delete(v1, v3)
        switch v3 {
        case "hello":
            for fi12 := 0; fi12 < 6; fi12++ {
                v1[("\n\t" + v3)] = ("" + v3)
                v13 := GenS2{F0: 570, F1: 0, F2: (v10 ^ v10)}
                _ = fi12
                _ = v13
            }
        default:
            var v14 bool = (false && false)
            _ = v14
        }
    }
    var v15 float32 = 3.14159
    v1[((((v3 + (v3 + v3)) + v3) + ("foo bar" + v3)) + (v3 + (v3 + v3)))] = v3
    if (v3 >= "a") {
        v15 = 1.5
        var v16 int32 = 91
        _ = v16
    } else {
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        fmt.Println(v2, v3, v2)
    }
    var v17 int64 = 612
    v4 += v4
    if true {
        v18 := GenS3{F0: 0, F1: 504, F2: 30}
        for fi19 := 0; fi19 < 6; fi19++ {
            if false {
                var v20 uint8 = 0
                v17 -= v17
                v15 = v15
                _ = v20
            }
            v4 += (fi19 << ((525) & 63))
            _ = fi19
        }
        for fi21 := 0; fi21 < 8; fi21++ {
            var v22 uint64 = 42
            v2 += 272
            _ = fi21
            _ = v22
        }
        _ = v18
    } else {
        v23 := GenS1{F0: 1000, F1: 42}
        v10 = v10
        _ = v23
    }
    var v24 uint64 = 0
    fmt.Println(v24, v4, v2)
    if true {
        var v25 bool = (!false)
        _ = v25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v17
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v9
}
