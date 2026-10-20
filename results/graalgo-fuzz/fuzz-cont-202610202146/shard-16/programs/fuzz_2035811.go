package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }

func main() {
    var v1 string = "hello"
    v2 := make(map[uint32]string)
    var v3 uint32 = 632
    v4 := make(map[int]int)
    var v5 int32 = 1
    var v6 bool = (-364 >= -6)
    v7 := GenS1{F0: uint64(v3), F1: 0}
    fmt.Println(v5, v3, v1)
    v8 := []int{-173, -192, 0}
    if false {
        v9 := GenS1{F0: uint64(v3), F1: 590}
        var v10 int = 32768
        if true {
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            if ((536 > 0) && (50 <= uint8(v3))) {
                var v11 bool = ((v1 <= (("hello" + "a") + ("hello" + (((((v1 + v1) + "hello") + "\n\t") + "\n\t") + "foo bar")))) && (95 != 16))
                var v12 uint32 = 0
                v13 := GenS1{F0: 18446744073709551615, F1: 65535}
                _ = v11
                _ = v12
                _ = v13
            }
        }
        _ = v10
        _ = v9
    } else {
        v14 := GenS1{F0: 1, F1: 2}
        _ = v14
    }
    for fi15 := 0; fi15 < 20; fi15++ {
        fmt.Println(v6)
        _ = fi15
    }
    var v16 float32 = float32((3.14159 / 1.0))
    v17 := GenS1{F0: 42, F1: uint16(v3)}
    for fi18 := 0; fi18 < 7; fi18++ {
        var v19 int32 = v5
        _ = fi18
        _ = v19
    }
    for fi20 := 0; fi20 < 2; fi20++ {
        var v21 string = "\n\t"
        _ = fi20
        _ = v21
    }
    fmt.Println(v3, v1, v6)
    if (1000 > 32767) {
        v3 -= v3
        v1 = ((((("\n\t" + "hello") + v1) + "a") + v1) + v1)
    }
    for fi22 := 0; fi22 < 11; fi22++ {
        v6 = (((float32(0.0) - (float32(0.1) * 0.0)) - 1.5) == v16)
        v23 := GenS1{F0: 1, F1: 32767}
        _ = fi22
        _ = v23
    }
    var v24 int32 = v5
    for fi25 := 0; fi25 < 9; fi25++ {
        v6 = false
        v5 = (v5 / ((2147483647) | 1))
        _ = fi25
    }
    _ = len(v8)
    v16 -= 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
