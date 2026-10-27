package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 bool }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 int32 }
func genFn0() uint64 {
    v4 := []int{1, 4294967295, 2147483647}
    var v5 uint32 = 1
    var v6 uint8 = 128
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v6 = (v6 << ((18446744073709551615) & 63))
    _ = v4
    _ = v5
    _ = v6
    return 1
}

func main() {
    var v1 uint16 = 474
    v2 := make(map[int]string)
    v3 := make(map[string]int)
    v1 += (v1 % ((v1) | 1))
    v1 = v1
    var v7 uint32 = 0
    if false {
        v8 := GenS1{F0: 32767, F1: v1, F2: (true || true)}
        v9 := GenS1{F0: 726, F1: v1, F2: (true || (v1 < v1))}
        for fi10 := 0; fi10 < 10; fi10++ {
            for fi11 := 0; fi11 < 19; fi11++ {
                if (!true) {
                    var v12 int32 = -1
                    v2[fi11] = "foo bar"
                    _ = v12
                }
                v7 += (^(v7))
                v13 := GenS1{F0: -32768, F1: (^(v1)), F2: (false || (0 <= 4294967295))}
                v7 = v7
                _ = fi11
                _ = v13
            }
            v14 := GenS1{F0: 0, F1: 149, F2: false}
            v7 -= (v7 >> ((2) & 63))
            _ = fi10
            _ = v14
        }
        _ = v8
        _ = v9
    } else {
        v1 -= 32767
        switch 32767 {
        case -1:
            var v15 uint = 588
            _ = v15
        case 2147483647:
            var v16 uint = 65535
            _ = v16
        }
        var v17 uint = 65535
        _ = v17
    }
    fmt.Println(v1, v7)
    v1 -= v1
    for fi18 := 0; fi18 < 11; fi18++ {
        var v19 int16 = -1
        if true {
            var v20 int = 128
            v21 := v2[2147483647]
            _ = v20
            _ = v21
        } else {
            var v22 int = 32767
            v23 := GenS1{F0: v19, F1: v1, F2: false}
            for fi24 := 0; fi24 < 2; fi24++ {
                v25 := []int{(^(v22))}
                fmt.Println(fi24, fi18)
                _ = fi24
                _ = v25
            }
            _ = v22
            _ = v23
        }
        v26 := GenS2{F0: v7, F1: (v19 ^ -32768)}
        _ = fi18
        _ = v19
        _ = v26
    }
    var v27 int32 = 449
    if (true || false) {
        delete(v3, "a")
        for fi28 := 0; fi28 < 10; fi28++ {
            var v29 bool = (!(!true))
            v30 := GenS2{F0: v7, F1: -1}
            v31 := GenS2{F0: (v7 - 1), F1: -489}
            _ = fi28
            _ = v29
            _ = v30
            _ = v31
        }
    }
    var v32 float32 = ((float32(((0.1 / 1.0) - 0.0)) - 0.0) + 0.0)
    if false {
        var v33 uint16 = v1
        _ = v33
    }
    fmt.Println(v32)
    v34 := GenS3{F0: 2147483647}
    for fi35 := 0; fi35 < 17; fi35++ {
        var v36 int8 = 100
        _ = fi35
        _ = v36
    }
    var v37 uint16 = v1
    v38 := GenS1{F0: 32767, F1: v37, F2: true}
    var v39 int32 = v27
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v32
    _ = v34
    _ = v37
    _ = v38
    _ = v39
    _ = v7
}
