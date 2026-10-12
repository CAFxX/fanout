package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 uint8 }
func genFn0() int32 {
    var v5 int = 4294967295
    if (false && false) {
        v5 += v5
        var v6 bool = (true || false)
        _ = v6
    }
    if (!(-32768 == 1)) {
        if (128 != 42) {
            v5 = v5
            for fi7 := 0; fi7 < 19; fi7++ {
                var v8 int8 = 100
                if (false && true) {
                    for fi9 := 0; fi9 < 20; fi9++ {
                        v10 := GenS1{F0: 528, F1: 818}
                        fmt.Println(v8, v5)
                        _ = fi9
                        _ = v10
                    }
                }
                v5 += v5
                _ = fi7
                _ = v8
            }
        } else {
            for fi11 := 0; fi11 < 5; fi11++ {
                v5 = v5
                var v12 bool = (false && true)
                _ = fi11
                _ = v12
            }
            for fi13 := 0; fi13 < 11; fi13++ {
                var v14 float64 = 1.5
                fmt.Println(fi13, v5, v14)
                _ = fi13
                _ = v14
            }
            v15 := GenS1{F0: -267, F1: 0}
            _ = v15
        }
        v5 -= v5
        v5 = v5
    }
    v16 := GenS1{F0: 0, F1: 18446744073709551615}
    if ((("hello" + "") < "\n\t") || (!(-9223372036854775808 != -1))) {
        fmt.Println(v5)
    } else {
        var v17 int16 = 1000
        for fi18 := 0; fi18 < 7; fi18++ {
            var v19 uint8 = 234
            v19 = v19
            v20 := GenS2{F0: 640}
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = v17
    }
    _ = v16
    _ = v5
    return 42
}

func main() {
    v1 := make(map[int]bool)
    var v2 int8 = 83
    v3 := make(map[int]int)
    var v4 uint = 4294967295
    v21 := GenS3{F0: (true && false), F1: 2}
    var v22 float64 = 0.0
    delete(v3, 42)
    if false {
        v1[32768] = (!("\n\t" > ("hello" + "\n\t")))
        if (42 == -2147483648) {
            v22 -= 1.5
            v23 := []int{}
            _ = v23
        }
        fmt.Println(v2)
    }
    switch -2147483648 {
    case -761:
        v22 += v22
    case 1:
        v22 -= v22
    }
    var v24 uint64 = 488
    fmt.Println(v4)
    if (!false) {
        v2 += v2
    } else {
        var v25 uint64 = uint64(v4)
        for fi26 := 0; fi26 < 10; fi26++ {
            var v27 uint32 = 2147483648
            v24 = v25
            v22 += ((v22 + (v22 / v22)) / float64((float32(-2.25) / 1.0)))
            _ = fi26
            _ = v27
        }
        _ = v25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v24
    _ = v3
    _ = v4
}
