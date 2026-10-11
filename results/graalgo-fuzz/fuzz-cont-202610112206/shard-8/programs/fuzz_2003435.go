package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 int32; F1 bool }
func genFn0(p0 int) int16 {
    fmt.Println(p0)
    v5 := GenS1{F0: 1, F1: "\n\t"}
    var v6 int32 = 42
    v7 := GenS1{F0: 696, F1: "foo bar"}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 bool = (-505 > 762)
    v2 := make(map[int]int)
    v3 := make(map[uint32]string)
    var v4 uint16 = 570
    var v8 bool = ((89 <= 128) && false)
    var v9 int16 = 1000
    v9 = (v9 - v9)
    var v10 int16 = v9
    fmt.Println(v8)
    for fi11 := 0; fi11 < 12; fi11++ {
        for fi12 := 0; fi12 < 12; fi12++ {
            switch fi12 {
            case -9223372036854775808:
                if true {
                    switch fi11 {
                    case 42:
                        var v13 float32 = ((1.5 + 1.5) + (0.1 * 1e10))
                        _ = v13
                    case 1:
                        v14 := []int{(^(fi11)), (fi12 / ((32768) | 1))}
                        _ = v14
                    case -2147483648:
                        switch fi11 {
                        case -468:
                            v15 := []int{fi12, fi12}
                            _ = v15
                        case 587:
                            if (!((v10 / ((int16(fi11)) | 1)) == v9)) {
                                if v1 {
                                    v16 := GenS1{F0: 0, F1: ("" + "hello")}
                                    v17 := v2[(^(fi12))]
                                    _ = v16
                                    _ = v17
                                }
                            }
                        case 42:
                            v1 = v8
                        default:
                            if (42 >= -2147483648) {
                                v18 := GenS2{F0: 2147483647, F1: false}
                                _ = v18
                            }
                        }
                    }
                    fmt.Println(v8, v1, fi11)
                }
            case 42:
                var v19 int64 = 175
                if v1 {
                    var v20 int64 = v19
                    _ = v20
                }
                _ = v19
            case 9223372036854775807:
                v21 := GenS1{F0: 4294967295, F1: ("hello" + ("foo bar" + "foo bar"))}
                _ = v21
            default:
                v22 := GenS1{F0: 580, F1: "\n\t"}
                _ = v22
            }
            var v23 string = (("hello" + ("hello" + "")) + "")
            _ = fi12
            _ = v23
        }
        v8 = v8
        fmt.Println(fi11, v8)
        _ = fi11
    }
    for fi24 := 0; fi24 < 18; fi24++ {
        var v25 string = ("a" + "a")
        _ = fi24
        _ = v25
    }
    var v26 uint32 = 738
    v27 := v3[v26]
    if (2147483647 >= -550) {
        v3[2147483647] = (v27 + "foo bar")
        v8 = v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
