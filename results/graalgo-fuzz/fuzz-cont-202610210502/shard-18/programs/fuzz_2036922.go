package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 int16; F2 int64 }
func genFn0(p0 int64) bool {
    for fi5 := 0; fi5 < 17; fi5++ {
        for fi6 := 0; fi6 < 15; fi6++ {
            v7 := []int{}
            v8 := GenS1{F0: "a"}
            p0 = (^(p0))
            p0 -= p0
            var v9 int64 = p0
            fmt.Println(p0, fi6)
            var v10 bool = (!(42 > 928))
            _ = fi6
            _ = v10
            _ = v7
            _ = v8
            _ = v9
        }
        p0 += p0
        _ = fi5
    }
    p0 = 601
    for fi11 := 0; fi11 < 19; fi11++ {
        v12 := []int{fi11, (fi11 + (fi11 ^ (fi11 * fi11))), fi11, 42}
        _ = fi11
        _ = v12
    }
    v13 := []int{}
    _ = p0
    _ = v13
    return false
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int8 = 0
    var v3 uint16 = 427
    var v4 uint64 = 18446744073709551615
    var v14 uint16 = v3
    v2 = v2
    delete(v1, 2147483648)
    if false {
        delete(v1, 4294967295)
        var v15 string = "\n\t"
        _ = v15
    }
    for fi16 := 0; fi16 < 2; fi16++ {
        v3 = 65535
        _ = fi16
    }
    _ = len(v1)
    for fi17 := 0; fi17 < 3; fi17++ {
        v18 := GenS1{F0: "\n\t"}
        v3 -= v14
        v14 = (v14 % ((42) | 1))
        _ = fi17
        _ = v18
    }
    fmt.Println(v14)
    var v19 int = 1
    var v20 uint16 = v3
    _ = len(v1)
    v20 = 1
    for fi21 := 0; fi21 < 4; fi21++ {
        v22 := GenS2{F0: (v14 - (v3 / ((v3) | 1))), F1: 419, F2: 1}
        var v23 int64 = -1
        for fi24 := 0; fi24 < 15; fi24++ {
            var v25 int32 = 2147483647
            _ = fi24
            _ = v25
        }
        _ = fi21
        _ = v22
        _ = v23
    }
    var v26 uint16 = v14
    if (2 > 0) {
        if (32767 >= 931) {
            if false {
                var v27 int = (v19 ^ v19)
                v19 -= (v19 >> ((7) & 63))
                var v28 string = "hello"
                _ = v27
                _ = v28
            }
        }
        switch -2147483648 {
        case -2147483648:
            var v29 uint16 = v26
            _ = v29
        case 1:
            var v30 int8 = int8(v4)
            _ = v30
        case 2147483647:
            v2 -= v2
        default:
            v31 := []int{v19, v19}
            _ = v31
        }
        v1[0] = 748
    }
    v1[1] = 9223372036854775807
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v4
}
