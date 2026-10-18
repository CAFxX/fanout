package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 uint; F1 int8; F2 uint }
type GenS3 struct { F0 uint }
func genFn0() int64 {
    var v6 uint16 = 65535
    switch 9223372036854775807 {
    case -9223372036854775808:
        v7 := GenS1{F0: "hello", F1: 2147483648}
        _ = v7
    case -1:
        if (42 == 0) {
            switch "\n\t" {
            case "\n\t":
                switch ("hello" + "foo bar") {
                case "a":
                    if (false && (!("hello" == "\n\t"))) {
                        v6 = v6
                        if false {
                            v8 := GenS1{F0: ("\n\t" + "a"), F1: 2147483648}
                            _ = v8
                        }
                    } else {
                        var v9 uint64 = 594
                        v6 = (v6 % (((v6 | 255)) | 1))
                        v10 := []int{42}
                        _ = v10
                        _ = v9
                    }
                default:
                    if (false || (0 < 56)) {
                        for fi11 := 0; fi11 < 2; fi11++ {
                            v12 := []int{fi11, fi11, 42, fi11}
                            _ = fi11
                            _ = v12
                        }
                    } else {
                        if (true || false) {
                            v6 += 255
                            var v13 int8 = 0
                            _ = v13
                        } else {
                            var v14 int64 = -116
                            v15 := GenS2{F0: 65535, F1: 1, F2: 0}
                            _ = v14
                            _ = v15
                        }
                        var v16 int = 0
                        v17 := []int{v16, (v16 - -279), (v16 >> ((4294967295) & 63)), v16}
                        _ = v16
                        _ = v17
                    }
                }
            case "a":
                if true {
                    v18 := []int{9223372036854775807, 0, 1}
                    _ = v18
                } else {
                    var v19 uint64 = 4294967295
                    _ = v19
                }
            default:
                var v20 uint = 447
                _ = v20
            }
        } else {
            v21 := []int{128, -329, 1, 9223372036854775807}
            var v22 uint8 = 0
            _ = v21
            _ = v22
        }
    default:
        v6 += v6
    }
    fmt.Println(v6)
    v6 = v6
    _ = v6
    return 9223372036854775807
}
func genFn1() string {
    v23 := GenS1{F0: "hello", F1: 4294967295}
    var v24 uint64 = 250
    v25 := GenS3{F0: uint(v24)}
    var v26 bool = (255 == 42)
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    return ("" + ("" + "\n\t"))
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]uint32)
    v3 := make(map[int]int)
    v4 := make(map[string]string)
    var v5 uint8 = 0
    fmt.Println(v5)
    var v27 int16 = -1
    v27 -= (^(v27))
    fmt.Println(v27, v27)
    v27 = -32768
    v28 := GenS3{F0: 18446744073709551615}
    v27 = v27
    v1[75] = 65535
    v29 := []int{}
    v27 = -32768
    if true {
        v5 = v5
        var v30 int = 9223372036854775807
        fmt.Println(v27, v27, v27)
        _ = v30
    } else {
        var v31 float64 = 0.0
        for fi32 := 0; fi32 < 4; fi32++ {
            var v33 uint64 = 329
            v34 := GenS1{F0: "\n\t", F1: 42}
            _ = len(v1)
            _ = fi32
            _ = v33
            _ = v34
        }
        fmt.Println(v27)
        _ = v31
    }
    var v35 int64 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
