package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
func genFn0() uint8 {
    if true {
        var v6 string = ("a" + "foo bar")
        _ = v6
    }
    if true {
        var v7 string = ""
        _ = v7
    }
    return 1
}

func main() {
    var v1 uint16 = 32767
    v2 := make(map[uint32]int)
    var v3 int32 = -2147483648
    var v4 bool = ((2 < 328) || false)
    var v5 uint16 = (v1 ^ (v1 / (((v1 * v1)) | 1)))
    if (1 == 127) {
        var v8 uint = 2
        var v9 uint64 = 18446744073709551615
        v9 = 1
        _ = v8
        _ = v9
    }
    if (-414 <= 127) {
        var v10 uint8 = 133
        if (v4 || ((1 <= 127) && false)) {
            v11 := GenS1{F0: -1, F1: 82}
            v12 := v2[128]
            _ = v11
            _ = v12
        } else {
            v13 := v2[uint32(v5)]
            var v14 int64 = int64(v13)
            switch ("hello" + "hello") {
            case "foo bar":
                v15 := GenS1{F0: 0, F1: 127}
                _ = v15
            case "hello":
                for fi16 := 0; fi16 < 17; fi16++ {
                    v17 := GenS1{F0: 28, F1: 42}
                    _ = fi16
                    _ = v17
                }
            default:
                for fi18 := 0; fi18 < 8; fi18++ {
                    var v19 float32 = (0.0 / float32(3.14159))
                    if (9223372036854775807 >= v14) {
                        v2[2147483647] = fi18
                    }
                    _ = fi18
                    _ = v19
                }
            }
            _ = v13
            _ = v14
        }
        _ = v10
    } else {
        v5 -= v5
    }
    v20 := GenS1{F0: -128, F1: 42}
    switch "hello" {
    case "hello":
        var v21 uint16 = 42
        _ = v21
    case "a":
        v22 := GenS1{F0: 42, F1: -1}
        _ = v22
    case "":
        fmt.Println(v1)
    }
    if false {
        v5 -= (v5 * 1)
        fmt.Println(v4, v3, v1)
        var v23 int16 = 685
        _ = v23
    } else {
        var v24 uint8 = genFn0()
        var v25 uint = 178
        _ = v24
        _ = v25
    }
    v26 := []int{}
    if false {
        for fi27 := 0; fi27 < 5; fi27++ {
            for fi28 := 0; fi28 < 19; fi28++ {
                if v4 {
                    var v29 uint32 = 0
                    var v30 uint32 = (v29 / ((v29) | 1))
                    _ = v29
                    _ = v30
                } else {
                    if len(v26) > 0 {
                        v26[0] = (fi28 >> ((4) & 63))
                    }
                    v31 := v2[2147483647]
                    _ = len(v26)
                    _ = v31
                }
                _ = fi28
            }
            var v32 uint8 = 42
            v5 += 362
            _ = fi27
            _ = v32
        }
        v3 -= 0
    }
    v5 -= v1
    var v33 uint8 = 108
    v4 = v4
    var v34 string = "a"
    if v4 {
        var v35 uint16 = (v5 - (v1 + v5))
        v1 += v1
        v2[1] = 9223372036854775807
        _ = v35
    }
    switch v34 {
    case "foo bar":
        v1 = (v1 & v5)
    }
    var v36 uint16 = (^(v1))
    var v37 int64 = 470
    v38 := GenS1{F0: -51, F1: 0}
    v3 = (v3 ^ -1)
    var v39 uint32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v33
    _ = v34
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
