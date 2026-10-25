package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 int8 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 bool }

func main() {
    var v1 uint64 = 0
    var v2 uint16 = 65535
    v3 := make(map[int]uint32)
    var v4 uint16 = 251
    var v5 uint16 = 0
    v5 += 42
    for fi6 := 0; fi6 < 14; fi6++ {
        var v7 float32 = (0.1 / float32((0.1 * 0.1)))
        var v8 bool = (!false)
        if v8 {
            v9 := []int{(^(fi6)), -9223372036854775808}
            _ = v9
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        var v11 uint = 4294967295
        v11 = v11
        v4 = (v2 * v4)
        _ = fi10
        _ = v11
    }
    var v12 int64 = -9223372036854775808
    var v13 int = 42
    v2 = v2
    v4 += (v5 ^ 65535)
    v1 += v1
    for fi14 := 0; fi14 < 11; fi14++ {
        var v15 string = "\n\t"
        for fi16 := 0; fi16 < 6; fi16++ {
            v12 = v12
            var v17 uint8 = 42
            if false {
                if (v17 > 139) {
                    v18 := v3[(fi16 | (v13 / (((fi16 - (v13 * (v13 + 9223372036854775807)))) | 1)))]
                    _ = v18
                }
            }
            _ = fi16
            _ = v17
        }
        for fi19 := 0; fi19 < 18; fi19++ {
            if (-128 < 0) {
                v20 := GenS1{F0: 0, F1: 348, F2: 127}
                v21 := []int{}
                if true {
                    v2 = 32767
                    v5 -= 42
                    var v22 uint64 = v1
                    _ = v22
                } else {
                    if true {
                        var v23 uint64 = 42
                        fmt.Println(v2)
                        _ = v23
                    }
                    v15 = (v15 + v15)
                    switch ((v15 + (((v15 + ("\n\t" + "hello")) + ("a" + ((v15 + v15) + v15))) + v15)) + (v15 + v15)) {
                    case "a":
                        var v24 int = fi19
                        _ = v24
                    }
                }
                _ = v20
                _ = v21
            }
            _ = fi19
        }
        _ = fi14
        _ = v15
    }
    for fi25 := 0; fi25 < 3; fi25++ {
        switch (fi25 >> ((831) & 63)) {
        case 32767:
            if true {
                if false {
                    v26 := GenS2{F0: 935, F1: 237}
                    v27 := GenS2{F0: uint(v1), F1: 42}
                    v2 -= 0
                    _ = v26
                    _ = v27
                }
                v28 := GenS3{F0: false}
                v29 := []int{272, v13}
                _ = v28
                _ = v29
            } else {
                v30 := []int{}
                switch fi25 {
                case 297:
                    var v31 uint32 = 42
                    _ = v31
                }
                for fi32 := 0; fi32 < 10; fi32++ {
                    var v33 uint8 = 1
                    v2 = (v4 | (v4 + v5))
                    _ = fi32
                    _ = v33
                }
                _ = v30
            }
        default:
            v5 += (v2 - v4)
        }
        _ = len(v3)
        _ = fi25
    }
    var v34 uint32 = 324
    v2 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
