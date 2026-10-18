package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 uint }
type GenS2 struct { F0 string }
func genFn0(p0 int8, p1 int8) int16 {
    for fi5 := 0; fi5 < 7; fi5++ {
        var v6 bool = true
        _ = fi5
        _ = v6
    }
    var v7 uint32 = 2147483647
    var v8 int64 = -404
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -32768
}

func main() {
    var v1 uint64 = 0
    v2 := []uint32{2147483648, 601, 2}
    v3 := []uint32{}
    v4 := make(map[uint32]string)
    v1 = (v1 >> ((22) & 63))
    var v9 int16 = genFn0(-128, -61)
    v10 := GenS1{F0: 2, F1: (v9 * v9), F2: 36}
    var v11 uint8 = 2
    v9 = (v9 | genFn0(42, -26))
    fmt.Println(v1, v1)
    v12 := GenS1{F0: v11, F1: (v9 ^ v9), F2: 18446744073709551615}
    v11 += (v11 * v11)
    if (0.1 >= (0.1 / 1.0)) {
        delete(v4, 65535)
    }
    var v13 string = "a"
    v9 = (-(v9))
    var v14 int8 = -1
    if true {
        switch v13 {
        case "":
            if false {
                for fi15 := 0; fi15 < 16; fi15++ {
                    v16 := GenS2{F0: v13}
                    fmt.Println(v11)
                    _ = fi15
                    _ = v16
                }
            } else {
                v17 := GenS1{F0: (v11 ^ (v11 * (v11 & (v11 + 255)))), F1: (v9 + v9), F2: 42}
                switch (v13 + v13) {
                case "a":
                    v4[257] = "foo bar"
                }
                var v18 string = v13
                _ = v17
                _ = v18
            }
        }
        var v19 int64 = 9223372036854775807
        v20 := GenS2{F0: (v13 + (v13 + ("a" + "a")))}
        _ = v19
        _ = v20
    }
    switch v13 {
    case "":
        for fi21 := 0; fi21 < 16; fi21++ {
            if (-607 <= -40) {
                v1 = v1
                delete(v4, uint32(v1))
                v13 = "hello"
            } else {
                var v22 string = "\n\t"
                v11 = 255
                _ = v22
            }
            if true {
                fmt.Println(v13, fi21, fi21)
                for fi23 := 0; fi23 < 11; fi23++ {
                    var v24 uint64 = (v1 & (v1 & v1))
                    fmt.Println(v13, v9)
                    var v25 uint64 = uint64(v11)
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            }
            for fi26 := 0; fi26 < 5; fi26++ {
                var v27 float64 = float64(0.1)
                var v28 uint = 255
                _ = fi26
                _ = v27
                _ = v28
            }
            _ = fi21
        }
    case "a":
        v29 := GenS1{F0: 255, F1: 42, F2: 902}
        _ = v29
    }
    v30 := v4[2]
    var v31 uint8 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v9
}
