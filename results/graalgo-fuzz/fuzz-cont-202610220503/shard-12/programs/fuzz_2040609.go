package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 int32 }
func genFn0() int {
    switch 128 {
    case 1:
        if ((true && true) || ("a" > "")) {
            var v5 float32 = float32((0.1 + 0.1))
            fmt.Println(v5)
            var v6 uint64 = 2
            _ = v5
            _ = v6
        }
    case -2147483648:
        switch -1 {
        case 476:
            v7 := GenS1{F0: -803, F1: 42, F2: 42}
            _ = v7
        case 525:
            var v8 int = 2147483647
            _ = v8
        default:
            var v9 bool = false
            _ = v9
        }
    case 32767:
        var v10 int8 = -1
        _ = v10
    }
    var v11 string = ("foo bar" + "")
    switch -829 {
    case 32768:
        v11 = v11
    case -291:
        v12 := []int{1}
        _ = v12
    }
    _ = v11
    return 32767
}

func main() {
    var v1 string = "\n\t"
    v2 := make(map[int]string)
    v3 := make(map[string]int)
    var v4 int32 = 242
    v4 += v4
    v13 := GenS2{F0: 100}
    var v14 int64 = -9223372036854775808
    if false {
        v14 = v14
        var v15 int8 = 100
        _ = v15
    }
    v14 += (v14 << ((688) & 63))
    switch 127 {
    case 657:
        v16 := GenS3{F0: 548, F1: v4}
        _ = v16
    }
    var v17 uint64 = 1
    if (!true) {
        v4 = -266
        v18 := GenS3{F0: 881, F1: -637}
        v19 := GenS1{F0: -1, F1: (v14 | (^(v14))), F2: 1}
        _ = v18
        _ = v19
    }
    v20 := GenS2{F0: 52}
    switch -1 {
    case 9223372036854775807:
        if (v4 != v4) {
            v14 += -9223372036854775808
            var v21 int = 1
            v17 = v17
            _ = v21
        } else {
            v22 := GenS3{F0: 406, F1: (v4 / ((v4) | 1))}
            v23 := []int{-91}
            _ = v22
            _ = v23
        }
    case 128:
        if false {
            switch "foo bar" {
            case "a":
                v24 := GenS3{F0: 970, F1: v4}
                _ = v24
            case "hello":
                delete(v2, 32767)
            case "foo bar":
                v4 = v4
            default:
                var v25 float32 = ((-2.25 * 0.1) + 1e10)
                _ = v25
            }
            fmt.Println(v17)
            v4 = 0
        }
    default:
        if (0 != 952) {
            if true {
                fmt.Println(v14)
            }
            v26 := []int{551, 487}
            _ = v26
        } else {
            _ = len(v3)
        }
    }
    _ = len(v3)
    v27 := []int{32768, 9223372036854775807, 128}
    v28 := GenS3{F0: 1, F1: v4}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v20
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
