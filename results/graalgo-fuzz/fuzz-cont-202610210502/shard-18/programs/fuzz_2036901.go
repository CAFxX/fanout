package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 int16 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 int16 = 1
    var v2 uint64 = 2
    var v3 float32 = ((((0.0 - (1.5 * (-2.25 + float32(-2.25)))) + 3.14159) + -2.25) * 0.1)
    var v4 uint16 = 42
    v1 += -32768
    switch 32767 {
    case 789:
        v5 := []int{}
        _ = v5
    case -887:
        if (100 != -33) {
            for fi6 := 0; fi6 < 2; fi6++ {
                v7 := []int{fi6, (fi6 >> ((1) & 63)), fi6, fi6}
                _ = fi6
                _ = v7
            }
            var v8 uint32 = 65535
            _ = v8
        }
    }
    var v9 int64 = 0
    fmt.Println(v3, v2, v2)
    v10 := GenS1{F0: 4294967295, F1: false, F2: v1}
    var v11 uint = 847
    fmt.Println(v4, v4, v2)
    if false {
        if true {
            var v12 int8 = 1
            _ = v12
        }
        for fi13 := 0; fi13 < 2; fi13++ {
            v9 = (v9 ^ (v9 % (((v9 | -1)) | 1)))
            fmt.Println(v9, fi13)
            if true {
                if true {
                    if (fi13 > fi13) {
                        for fi14 := 0; fi14 < 2; fi14++ {
                            v11 = v11
                            v15 := GenS2{F0: 1, F1: 2147483647}
                            _ = fi14
                            _ = v15
                        }
                    }
                }
                v16 := []int{}
                _ = v16
            }
            _ = fi13
        }
        v17 := GenS3{F0: 127}
        _ = v17
    }
    if (true && true) {
        v18 := []int{-658, 258, 9223372036854775807, -2147483648}
        _ = v18
    }
    fmt.Println(v4, v3, v1)
    if true {
        for fi19 := 0; fi19 < 10; fi19++ {
            fmt.Println(fi19, v1)
            _ = fi19
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
