package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 int8; F2 uint8 }
func genFn0() uint {
    var v5 uint16 = 65535
    var v6 string = "a"
    _ = v5
    _ = v6
    return 65535
}
func genFn1() uint16 {
    var v7 float64 = 1e10
    var v8 int16 = 1000
    fmt.Println(v8, v8)
    v7 = v7
    for fi9 := 0; fi9 < 3; fi9++ {
        v7 = v7
        v7 = 1.5
        v8 += v8
        _ = fi9
    }
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 int = 22
    var v2 bool = false
    var v3 string = ""
    var v4 uint32 = 4294967295
    fmt.Println(v2, v4, v3)
    fmt.Println(v2, v1, v2)
    var v10 uint = 0
    v4 = 1
    v11 := GenS1{F0: -51}
    var v12 uint64 = 1
    v13 := GenS1{F0: -128}
    for fi14 := 0; fi14 < 17; fi14++ {
        fmt.Println(v4)
        v15 := GenS1{F0: -43}
        switch -9223372036854775808 {
        case 9223372036854775807:
            v16 := GenS1{F0: -54}
            _ = v16
        case -436:
            v2 = v2
        }
        _ = fi14
        _ = v15
    }
    switch v3 {
    case "\n\t":
        v17 := []int{}
        _ = v17
    case "a":
        for fi18 := 0; fi18 < 20; fi18++ {
            v2 = v2
            for fi19 := 0; fi19 < 3; fi19++ {
                switch "a" {
                case "foo bar":
                    if (0.0 >= (0.1 * (3.14159 * 3.14159))) {
                        var v20 int64 = 0
                        if (!(v2 || v2)) {
                            var v21 uint = v10
                            v4 = (^(v4))
                            v22 := GenS1{F0: -110}
                            _ = v21
                            _ = v22
                        } else {
                            v23 := GenS1{F0: -1}
                            fmt.Println(v1, v10, v4)
                            v3 = (v3 + v3)
                            _ = v23
                        }
                        var v24 uint64 = v12
                        switch v3 {
                        case "":
                            var v25 int64 = v20
                            _ = v25
                        default:
                            v2 = false
                        }
                        _ = v20
                        _ = v24
                    }
                default:
                    var v26 float32 = (float32((0.0 + 1e10)) / 3.14159)
                    _ = v26
                }
                v27 := []int{fi19, (fi19 & (fi19 + v1)), v1, fi18}
                _ = fi19
                _ = v27
            }
            _ = fi18
        }
    case "hello":
        if v2 {
            var v28 int = (^(v1))
            fmt.Println(v4, v4)
            _ = v28
        } else {
            var v29 uint64 = (^(v12))
            _ = v29
        }
    default:
        var v30 float32 = 1.5
        _ = v30
    }
    var v31 uint64 = v12
    switch v1 {
    case -1:
        v2 = true
    case 4294967295:
        if v2 {
            var v32 int = (v1 % (((^(v1))) | 1))
            _ = v32
        } else {
            var v33 string = v3
            _ = v33
        }
    }
    switch (v3 + v3) {
    case "a":
        fmt.Println(v2, v31, v4)
    case "hello":
        v34 := GenS2{F0: 127, F1: 1, F2: 0}
        _ = v34
    case "":
        var v35 uint = (v10 >> ((v10) & 63))
        _ = v35
    default:
        v36 := []int{}
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v31
    _ = v4
}
