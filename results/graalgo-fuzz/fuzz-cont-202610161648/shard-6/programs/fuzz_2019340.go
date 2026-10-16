package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 int }
func genFn0() int32 {
    var v5 float32 = 0.1
    var v6 uint = 60
    var v7 int = -2147483648
    var v8 float64 = float64(v5)
    v7 += v7
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return -265
}

func main() {
    v1 := []uint32{65535, 2}
    var v2 int8 = 42
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]bool)
    if false {
        v2 -= v2
    } else {
        v9 := GenS1{F0: -591}
        fmt.Println(v2)
        if true {
            v2 = (^(v2))
            v2 = v2
            v10 := v3[297]
            _ = v10
        } else {
            fmt.Println(v2)
            var v11 string = ("a" + "\n\t")
            _ = v11
        }
        _ = v9
    }
    var v12 uint16 = 527
    var v13 int32 = genFn0()
    if (!(("hello" + "foo bar") > ("hello" + ""))) {
        v2 -= v2
    } else {
        _ = len(v3)
        if len(v1) > 0 {
            v1[0] = 0
        }
    }
    v14 := GenS1{F0: 42}
    for fi15 := 0; fi15 < 14; fi15++ {
        v16 := GenS1{F0: -719}
        _ = fi15
        _ = v16
    }
    var v17 bool = true
    var v18 uint16 = 110
    v1 = append(v1, 1)
    v17 = (true && (875 > 1))
    var v19 uint8 = 127
    var v20 uint8 = v19
    v17 = v17
    for fi21 := 0; fi21 < 13; fi21++ {
        v17 = ((v2 >> ((18446744073709551615) & 63)) >= v2)
        if true {
            v4[4294967295] = v17
            switch -634 {
            case 42:
                if v17 {
                    fmt.Println(fi21)
                    var v22 uint = 0
                    _ = v22
                }
            }
            v12 = 2
        }
        v23 := GenS2{F0: 42}
        _ = fi21
        _ = v23
    }
    switch "foo bar" {
    case "":
        v24 := GenS3{F0: 1, F1: 128}
        _ = v24
    case "a":
        fmt.Println(v18)
    case "hello":
        v25 := GenS3{F0: int32(v2), F1: -176}
        _ = v25
    }
    var v26 uint64 = 308
    var v27 uint32 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
