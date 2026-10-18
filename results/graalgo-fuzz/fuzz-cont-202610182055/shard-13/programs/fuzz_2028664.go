package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 string }
func genFn0() uint64 {
    v6 := []int{}
    var v7 string = "foo bar"
    v8 := GenS1{F0: 18446744073709551615, F1: 42, F2: (v7 + "a")}
    _ = v6
    _ = v7
    _ = v8
    return 18446744073709551615
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]bool)
    v3 := []int{-2147483648, -2147483648, 128, 42}
    v4 := make(map[string]int)
    var v5 int32 = 0
    v9 := GenS1{F0: 0, F1: 42, F2: "foo bar"}
    var v10 uint16 = 0
    fmt.Println(v10, v10)
    delete(v2, 751)
    fmt.Println(v5, v5)
    if false {
        v4["hello"] = 1
        v10 += 861
        var v11 int64 = -1
        _ = v11
    } else {
        for fi12 := 0; fi12 < 11; fi12++ {
            if (true && false) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            _ = fi12
        }
        switch -354 {
        case 42:
            delete(v2, 42)
        case 128:
            v5 -= v5
        default:
            fmt.Println(v5, v10)
        }
    }
    for fi13 := 0; fi13 < 6; fi13++ {
        v5 = v5
        _ = fi13
    }
    var v14 float32 = 1.5
    switch 127 {
    case 4294967295:
        var v15 int16 = 0
        _ = v15
    case 530:
        if (!true) {
            v5 += (v5 | v5)
            var v16 int32 = v5
            _ = v16
        } else {
            var v17 uint16 = (^(v10))
            _ = v17
        }
    default:
        for fi18 := 0; fi18 < 19; fi18++ {
            v14 -= v14
            var v19 bool = ((32767 < (fi18 % (((-(fi18))) | 1))) || false)
            v10 -= 475
            _ = fi18
            _ = v19
        }
    }
    v20 := GenS1{F0: 4294967295, F1: genFn0(), F2: ("hello" + "\n\t")}
    for fi21 := 0; fi21 < 16; fi21++ {
        v5 = 0
        if (!("\n\t" < "hello")) {
            var v22 int32 = (v5 | (v5 & -980))
            _ = v22
        }
        _ = fi21
    }
    for fi23 := 0; fi23 < 17; fi23++ {
        var v24 uint64 = 0
        _ = fi23
        _ = v24
    }
    var v25 string = ("\n\t" + "hello")
    if true {
        v5 += v5
    } else {
        var v26 uint = 4294967295
        _ = v26
    }
    var v27 uint64 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v20
    _ = v25
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
