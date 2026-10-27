package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16 }
func genFn0() uint64 {
    var v5 bool = (1 <= 42)
    v5 = (42 <= 527)
    for fi6 := 0; fi6 < 7; fi6++ {
        v7 := GenS1{F0: 127, F1: 905, F2: 864}
        _ = fi6
        _ = v7
    }
    _ = v5
    return 4294967295
}

func main() {
    v1 := []int{0}
    var v2 bool = (2147483647 <= -106)
    v3 := make(map[int]uint32)
    var v4 string = (("foo bar" + "foo bar") + "\n\t")
    for fi8 := 0; fi8 < 5; fi8++ {
        for fi9 := 0; fi9 < 8; fi9++ {
            v10 := GenS1{F0: 0, F1: 42, F2: 65535}
            v2 = (3.14159 <= 0.1)
            var v11 string = v4
            _ = fi9
            _ = v10
            _ = v11
        }
        var v12 int32 = -333
        _ = fi8
        _ = v12
    }
    var v13 uint32 = 1
    if (!v2) {
        if true {
            for fi14 := 0; fi14 < 19; fi14++ {
                v15 := GenS2{F0: -378}
                for fi16 := 0; fi16 < 20; fi16++ {
                    var v17 uint = 385
                    _ = len(v1)
                    _ = fi16
                    _ = v17
                }
                _ = fi14
                _ = v15
            }
        }
        v2 = (0 != 2147483647)
    }
    for fi18 := 0; fi18 < 19; fi18++ {
        delete(v3, (^(fi18)))
        for fi19 := 0; fi19 < 4; fi19++ {
            var v20 uint16 = 837
            var v21 int64 = 42
            _ = fi19
            _ = v20
            _ = v21
        }
        var v22 int32 = -1
        _ = fi18
        _ = v22
    }
    var v23 int16 = -1
    v13 -= (v13 + v13)
    var v24 uint32 = (v13 + (^(v13)))
    if ((v24 * (^(v24))) > v24) {
        for fi25 := 0; fi25 < 19; fi25++ {
            v26 := GenS1{F0: 21, F1: (v13 - v24), F2: 255}
            v2 = (v2 && ((!(255 <= 42)) && v2))
            var v27 uint8 = 56
            _ = fi25
            _ = v26
            _ = v27
        }
    }
    fmt.Println(v4, v13, v4)
    v28 := GenS3{F0: v23}
    v29 := GenS2{F0: v23}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v23
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
