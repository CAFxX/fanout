package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 int8 }

func main() {
    var v1 float32 = (-2.25 / 1e10)
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]bool)
    v4 := make(map[uint32]bool)
    var v5 int16 = -1
    var v6 uint16 = 0
    v7 := GenS1{F0: 1}
    if false {
        v8 := v4[516]
        var v9 int8 = 1
        _ = v8
        _ = v9
    } else {
        switch 32768 {
        case 9223372036854775807:
            if (v1 == v1) {
                fmt.Println(v5, v5, v5)
            } else {
                if true {
                    v10 := GenS1{F0: v5}
                    _ = v10
                } else {
                    if true {
                        var v11 string = ((("\n\t" + ((((("foo bar" + "\n\t") + "foo bar") + "a") + "hello") + "foo bar")) + ("" + (((("a" + "") + ("" + "")) + "\n\t") + "hello"))) + (("\n\t" + "") + ""))
                        v5 -= v5
                        v1 += v1
                        _ = v11
                    }
                    v12 := []int{0, -790, 342}
                    _ = v12
                }
                if false {
                    var v13 uint16 = (^(v6))
                    if (((-2.25 * 0.1) / -2.25) <= 0.0) {
                        v14 := []int{}
                        _ = v14
                    }
                    _ = v13
                }
            }
        case 523:
            fmt.Println(v6, v1, v1)
        }
        var v15 int16 = v5
        _ = v15
    }
    v16 := GenS1{F0: 459}
    switch -9223372036854775808 {
    case -1:
        var v17 int16 = v5
        _ = v17
    }
    delete(v3, 965)
    v6 += v6
    if true {
        v18 := GenS2{F0: 255, F1: 153}
        _ = v18
    }
    v19 := GenS3{F0: -1}
    v6 += (v6 | v6)
    v20 := GenS3{F0: 3}
    for fi21 := 0; fi21 < 10; fi21++ {
        for fi22 := 0; fi22 < 10; fi22++ {
            delete(v2, 2147483647)
            var v23 float32 = v1
            _ = fi22
            _ = v23
        }
        var v24 int16 = v5
        fmt.Println(fi21, v6, v6)
        _ = fi21
        _ = v24
    }
    fmt.Println(v6, v1, v5)
    for fi25 := 0; fi25 < 9; fi25++ {
        fmt.Println(fi25, v5)
        var v26 uint16 = v6
        _ = fi25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
