package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 uint32 }
type GenS2 struct { F0 uint32; F1 uint64; F2 bool }
type GenS3 struct { F0 int; F1 int32 }
func genFn0(p0 uint8, p1 int64) uint {
    for fi5 := 0; fi5 < 16; fi5++ {
        var v6 int8 = 1
        v7 := GenS1{F0: 42, F1: 42, F2: 2147483648}
        _ = fi5
        _ = v6
        _ = v7
    }
    switch "\n\t" {
    case "a":
        p0 -= 2
    case "":
        var v8 uint32 = 691
        _ = v8
    }
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 int = 32767
    v2 := make(map[uint32]uint32)
    var v3 uint32 = 0
    v4 := make(map[string]bool)
    fmt.Println(v1, v3)
    v9 := GenS1{F0: -32768, F1: v3, F2: (v3 & v3)}
    delete(v4, "hello")
    v1 += 465
    var v10 uint8 = 255
    var v11 int16 = 32767
    fmt.Println(v11, v11)
    for fi12 := 0; fi12 < 8; fi12++ {
        v3 = v3
        fmt.Println(v11, v1)
        var v13 int64 = 9223372036854775807
        _ = fi12
        _ = v13
    }
    var v14 uint16 = 335
    v15 := GenS2{F0: (^(v3)), F1: 0, F2: true}
    if false {
        v3 = (v3 | 1)
        var v16 uint = 818
        delete(v4, ("foo bar" + "foo bar"))
        _ = v16
    }
    v17 := []int{v1, 2147483647, v1}
    var v18 uint8 = v10
    if (!("foo bar" < ("a" + "hello"))) {
        v19 := GenS2{F0: 629, F1: 421, F2: true}
        v18 = 114
        switch (v1 << ((genFn0(2, -745)) & 63)) {
        case 286:
            fmt.Println(v1, v18)
        case 1:
            var v20 string = ("hello" + "a")
            _ = v20
        case 584:
            v14 -= (^(v14))
        }
        _ = v19
    }
    for fi21 := 0; fi21 < 16; fi21++ {
        if (v11 != 0) {
            switch "foo bar" {
            case "a":
                for fi22 := 0; fi22 < 16; fi22++ {
                    switch 2147483647 {
                    case 32767:
                        v1 -= v1
                    default:
                        fmt.Println(v1, fi22, fi22)
                    }
                    v11 = v11
                    v3 = 0
                    _ = fi22
                }
            }
        }
        var v23 int32 = -2147483648
        if ("a" == "a") {
            v24 := GenS3{F0: v1, F1: (v23 | (v23 / ((v23) | 1)))}
            fmt.Println(v23, v18)
            _ = v24
        }
        _ = fi21
        _ = v23
    }
    v25 := v2[v3]
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v9
}
