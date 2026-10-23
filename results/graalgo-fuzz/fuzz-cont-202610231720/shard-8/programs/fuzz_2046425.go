package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32 }
func genFn0() int8 {
    var v7 bool = false
    v7 = (("foo bar" == "a") || (v7 && (1 != -32768)))
    _ = v7
    return 42
}
func genFn1(p0 int) int64 {
    var v8 uint = 1
    var v9 uint = v8
    for fi10 := 0; fi10 < 12; fi10++ {
        p0 = (^(p0))
        var v11 int8 = -128
        _ = fi10
        _ = v11
    }
    if (1 > 1) {
        v8 += v8
        if (!(genFn0() == genFn0())) {
            if (false && true) {
                if false {
                    var v12 string = "\n\t"
                    for fi13 := 0; fi13 < 15; fi13++ {
                        switch (v12 + "a") {
                        case "":
                            var v14 float32 = float32(0.1)
                            _ = v14
                        case "hello":
                            switch -2147483648 {
                            case 42:
                                var v15 bool = false
                                _ = v15
                            case 127:
                                p0 = 42
                            default:
                                var v16 int16 = -1
                                _ = v16
                            }
                        case "\n\t":
                            v17 := []int{-193}
                            _ = v17
                        }
                        v18 := GenS1{F0: 0}
                        _ = fi13
                        _ = v18
                    }
                    _ = v12
                } else {
                    var v19 uint = v9
                    var v20 uint = v8
                    _ = v19
                    _ = v20
                }
                v21 := GenS1{F0: 0}
                _ = v21
            }
            v9 = (v9 * v9)
        }
    } else {
        var v22 int = p0
        v22 = p0
        _ = v22
    }
    v23 := []int{}
    _ = p0
    _ = v23
    _ = v8
    _ = v9
    return -9223372036854775808
}

func main() {
    var v1 float32 = (0.1 / (1e10 * -2.25))
    v2 := make(map[string]bool)
    var v3 uint8 = 233
    var v4 int = 32767
    var v5 uint = 42
    v6 := make(map[int]bool)
    var v24 uint = (v5 >> ((449) & 63))
    v24 -= 425
    delete(v2, "")
    var v25 int32 = 539
    if true {
        v26 := GenS1{F0: (v25 * int32(v4))}
        if true {
            var v27 int16 = 0
            var v28 uint16 = 2
            _ = v27
            _ = v28
        }
        fmt.Println(v24)
        _ = v26
    }
    v29 := []int{4294967295}
    v30 := GenS2{F0: 2147483648}
    fmt.Println(v24, v25, v3)
    for fi31 := 0; fi31 < 5; fi31++ {
        v25 -= v25
        v5 = (v5 * v24)
        _ = fi31
    }
    var v32 float32 = (1.5 / (-2.25 * v1))
    var v33 float32 = ((v1 / 1.5) + ((v32 + float32(((0.1 / 3.14159) + 3.14159))) / -2.25))
    var v34 float64 = ((1.5 / float64(-2.25)) / 1e10)
    _ = len(v6)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
