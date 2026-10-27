package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 int16 }
type GenS2 struct { F0 int8; F1 uint8; F2 uint16 }
type GenS3 struct { F0 uint64; F1 string; F2 bool }
func genFn0() uint8 {
    var v7 string = ("foo bar" + "\n\t")
    fmt.Println(v7)
    var v8 uint64 = 1
    _ = v7
    _ = v8
    return 127
}

func main() {
    v1 := []string{"", ("foo bar" + "\n\t")}
    v2 := make(map[uint32]string)
    var v3 float32 = ((-2.25 / 1e10) + (float32(float64(-2.25)) + float32(float64((0.0 * ((0.1 - 3.14159) / ((1e10 + 1.5) + -2.25)))))))
    v4 := make(map[uint32]uint32)
    var v5 int16 = 32767
    var v6 uint8 = 128
    v2[2] = "hello"
    v6 = (v6 / ((116) | 1))
    v3 += 3.14159
    var v9 int32 = 42
    if (!true) {
        v5 = (v5 << ((65535) & 63))
        delete(v4, 0)
    }
    var v10 uint32 = 984
    v3 -= v3
    v11 := GenS1{F0: true, F1: 1, F2: -32768}
    var v12 float32 = v3
    var v13 string = "hello"
    if true {
        for fi14 := 0; fi14 < 5; fi14++ {
            v15 := GenS2{F0: 1, F1: genFn0(), F2: 42}
            for fi16 := 0; fi16 < 14; fi16++ {
                fmt.Println(fi16)
                _ = fi16
            }
            _ = fi14
            _ = v15
        }
        switch -2147483648 {
        case -773:
            for fi17 := 0; fi17 < 3; fi17++ {
                v9 += 2147483647
                _ = fi17
            }
        case 128:
            v18 := GenS3{F0: 4294967295, F1: "", F2: true}
            _ = v18
        case -1:
            if true {
                var v19 float32 = 1.5
                var v20 bool = false
                _ = v19
                _ = v20
            }
        }
    }
    v10 = (v10 - 0)
    delete(v4, v10)
    v21 := GenS2{F0: 100, F1: 2, F2: 2}
    if true {
        v12 += (v12 + 0.0)
        fmt.Println(v13)
        if (!true) {
            v2[(^(v10))] = (v13 + "")
            if len(v1) > 0 {
                v1[0] = ((((v13 + ("" + (v13 + v13))) + "foo bar") + v13) + "hello")
            }
        }
    }
    v2[2147483648] = v13
    var v22 int64 = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
