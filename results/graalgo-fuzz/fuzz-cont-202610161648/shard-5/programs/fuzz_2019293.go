package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 int16 }
func genFn0() uint {
    var v5 uint8 = 1
    for fi6 := 0; fi6 < 20; fi6++ {
        if false {
            var v7 float32 = 1.5
            _ = v7
        }
        _ = fi6
    }
    v8 := []int{}
    var v9 int = 1
    fmt.Println(v9, v9)
    _ = v5
    _ = v8
    _ = v9
    return 2
}
func genFn1(p0 int32, p1 bool, p2 int8) bool {
    var v10 string = "hello"
    fmt.Println(v10, v10)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return p1
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[uint32]int)
    var v3 uint8 = 202
    v4 := make(map[int]string)
    v11 := []int{9223372036854775807}
    for fi12 := 0; fi12 < 3; fi12++ {
        v13 := GenS1{F0: 1}
        _ = fi12
        _ = v13
    }
    var v14 int32 = -563
    var v15 uint16 = 42
    delete(v2, 42)
    if true {
        var v16 int = -9223372036854775808
        switch v16 {
        case 127:
            var v17 float32 = 3.14159
            _ = v17
        case 42:
            var v18 uint8 = 127
            _ = v18
        }
        _ = v16
    } else {
        v19 := GenS1{F0: 1}
        v20 := GenS1{F0: -1}
        for fi21 := 0; fi21 < 9; fi21++ {
            if (((true && false) || true) && (18446744073709551615 < 4294967295)) {
                if (2 <= 0) {
                    for fi22 := 0; fi22 < 15; fi22++ {
                        v23 := GenS1{F0: 42}
                        v24 := GenS1{F0: 126}
                        v25 := GenS1{F0: -32}
                        _ = fi22
                        _ = v23
                        _ = v24
                        _ = v25
                    }
                }
            }
            _ = fi21
        }
        _ = v19
        _ = v20
    }
    v26 := GenS2{F0: true, F1: 1000}
    v27 := GenS1{F0: -1}
    var v28 int = 869
    v29 := GenS1{F0: -105}
    for fi30 := 0; fi30 < 8; fi30++ {
        v14 = 260
        for fi31 := 0; fi31 < 2; fi31++ {
            delete(v2, 1)
            _ = fi31
        }
        if (float32(1.5) >= (0.1 - -2.25)) {
            v14 = v14
        }
        _ = fi30
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
