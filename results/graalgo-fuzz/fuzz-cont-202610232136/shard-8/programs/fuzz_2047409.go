package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 uint32 }
func genFn0() uint16 {
    var v7 bool = (4294967295 != 233)
    if true {
        v7 = (2147483647 == -2147483648)
        v8 := []int{2147483647, 9223372036854775807, 42, 2147483647}
        _ = v8
    }
    v9 := GenS1{F0: 18446744073709551615, F1: 65535}
    var v10 uint = 2
    _ = v10
    _ = v7
    _ = v9
    return 32767
}
func genFn1(p0 int16, p1 uint16) string {
    for fi11 := 0; fi11 < 3; fi11++ {
        v12 := []int{(fi11 * (fi11 & (fi11 >> ((514) & 63)))), (^(fi11)), -213, (fi11 * fi11)}
        _ = fi11
        _ = v12
    }
    fmt.Println(p0, p1)
    v13 := []int{711, -882, 861, -117}
    var v14 int16 = -32768
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return ("hello" + "")
}

func main() {
    var v1 uint = 206
    v2 := make(map[int]uint32)
    var v3 int16 = 1000
    var v4 uint64 = 18446744073709551615
    v5 := make(map[int]int)
    var v6 bool = true
    v3 -= 0
    if (-1 > 42) {
        for fi15 := 0; fi15 < 17; fi15++ {
            v2[(fi15 >> (((v1 | v1)) & 63))] = 1
            v16 := GenS1{F0: 18446744073709551615, F1: 2}
            _ = fi15
            _ = v16
        }
        var v17 int64 = -738
        var v18 uint32 = 2147483648
        _ = v17
        _ = v18
    }
    var v19 int64 = 9223372036854775807
    fmt.Println(v3)
    v19 = (v19 / ((-153) | 1))
    var v20 uint64 = (v4 * 2)
    var v21 int8 = -103
    if (v20 >= 18446744073709551615) {
        var v22 float64 = 0.1
        _ = v22
    } else {
        var v23 int32 = 5
        var v24 uint16 = 596
        if (v20 <= (v20 >> ((18446744073709551615) & 63))) {
            for fi25 := 0; fi25 < 15; fi25++ {
                var v26 int32 = (v23 / (((v23 | (v23 << ((1) & 63)))) | 1))
                var v27 string = "\n\t"
                _ = fi25
                _ = v26
                _ = v27
            }
            v21 -= (v21 | (v21 / ((v21) | 1)))
            var v28 bool = (v3 != v3)
            _ = v28
        }
        _ = v23
        _ = v24
    }
    v6 = v6
    v21 = (v21 | (v21 | v21))
    switch -875 {
    case 32767:
        v6 = (2147483648 >= 4294967295)
    case 9223372036854775807:
        for fi29 := 0; fi29 < 15; fi29++ {
            v4 = (v20 % ((239) | 1))
            if (!v6) {
                for fi30 := 0; fi30 < 19; fi30++ {
                    v31 := []int{fi29, int(v20)}
                    var v32 uint16 = 65535
                    _ = fi30
                    _ = v31
                    _ = v32
                }
            } else {
                v33 := GenS2{F0: 2147483647}
                _ = v33
            }
            _ = fi29
        }
    case 2147483647:
        for fi34 := 0; fi34 < 2; fi34++ {
            v35 := []int{(fi34 - 127), (^(fi34)), fi34, fi34}
            _ = fi34
            _ = v35
        }
    }
    v19 -= v19
    switch 32767 {
    case 4294967295:
        delete(v5, 2147483647)
    default:
        v36 := GenS2{F0: 42}
        _ = v36
    }
    v37 := []int{}
    v6 = v6
    var v38 int = 0
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
