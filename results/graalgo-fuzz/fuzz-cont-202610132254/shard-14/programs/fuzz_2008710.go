package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 uint; F1 string; F2 uint8 }

func main() {
    v1 := make(map[string]int)
    var v2 uint32 = 1
    var v3 uint32 = v2
    var v4 uint16 = 32767
    var v5 int64 = 872
    fmt.Println(v3, v3, v2)
    var v6 int64 = (-(v5))
    if false {
        v5 = 548
        switch "hello" {
        case "":
            v7 := GenS1{F0: (v2 + (v3 << ((65535) & 63)))}
            _ = v7
        case "a":
            if (false || false) {
                v2 += v2
            }
        default:
            v5 = v5
        }
        v2 -= 2147483647
    } else {
        v8 := v1["foo bar"]
        v3 += 750
        _ = v8
    }
    v9 := GenS2{F0: 975, F1: 0}
    if (!(-128 >= 1)) {
        v10 := GenS2{F0: 0, F1: 124}
        if (true && false) {
            fmt.Println(v3, v3, v5)
            delete(v1, "\n\t")
        } else {
            delete(v1, "a")
        }
        _ = v10
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        if true {
            if false {
                var v12 uint16 = uint16(v3)
                _ = v12
            } else {
                if (42 > 1) {
                    if (4294967295 >= uint64(v3)) {
                        fmt.Println(v6)
                        v2 += v3
                    }
                    for fi13 := 0; fi13 < 17; fi13++ {
                        if true {
                            _ = len(v1)
                            var v14 uint64 = 476
                            _ = v14
                        }
                        v15 := []int{}
                        v1["\n\t"] = fi11
                        _ = fi13
                        _ = v15
                    }
                    v1["a"] = (fi11 << ((65535) & 63))
                } else {
                    for fi16 := 0; fi16 < 6; fi16++ {
                        v4 = 1
                        _ = fi16
                    }
                }
                v3 = 603
                fmt.Println(v5)
            }
            fmt.Println(fi11, v4, v5)
            v3 = v2
        }
        v1["hello"] = fi11
        if false {
            v2 = v2
            v6 = (-(v5))
            v1["\n\t"] = 4294967295
        }
        _ = fi11
    }
    v5 += v6
    var v17 float32 = -2.25
    fmt.Println(v3, v2, v4)
    if false {
        v2 -= (v3 ^ (v3 / ((v3) | 1)))
    } else {
        v18 := GenS3{F0: 4294967295, F1: "a", F2: 42}
        v17 = v17
        var v19 int16 = 42
        _ = v18
        _ = v19
    }
    v2 = v3
    for fi20 := 0; fi20 < 18; fi20++ {
        v17 = (v17 * v17)
        var v21 uint = 65535
        _ = fi20
        _ = v21
    }
    _ = len(v1)
    if false {
        fmt.Println(v5, v3)
        v4 += (v4 ^ (v4 & (^(v4))))
    }
    v4 -= 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
