package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 int; F1 int64; F2 int32 }

func main() {
    var v1 uint64 = 1
    v2 := make(map[uint32]string)
    var v3 uint32 = 0
    var v4 int16 = 0
    v5 := make(map[string]uint32)
    var v6 int32 = -1
    fmt.Println(v1, v4, v4)
    v7 := GenS1{F0: 1, F1: 280}
    v8 := GenS2{F0: -2147483648, F1: 1, F2: v6}
    v4 = v4
    v4 = -1
    for fi9 := 0; fi9 < 14; fi9++ {
        for fi10 := 0; fi10 < 1; fi10++ {
            v1 = v1
            v11 := GenS2{F0: (fi10 % ((int(v6)) | 1)), F1: 0, F2: 1}
            fmt.Println(v4, fi10)
            _ = fi10
            _ = v11
        }
        v2[4294967295] = ("a" + "\n\t")
        _ = fi9
    }
    for fi12 := 0; fi12 < 18; fi12++ {
        var v13 int64 = -1
        for fi14 := 0; fi14 < 10; fi14++ {
            v4 = v4
            var v15 int64 = (v13 * (v13 * 42))
            var v16 float32 = ((3.14159 * ((float32(3.14159) * (1e10 * 1.5)) * 1e10)) + (1e10 + (1e10 + 1.5)))
            var v17 int16 = v4
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
        v3 -= 1
        _ = fi12
        _ = v13
    }
    v6 = (v6 * v6)
    var v18 int32 = -2147483648
    var v19 int16 = (v4 << ((566) & 63))
    v1 = (v1 * 0)
    fmt.Println(v6)
    var v20 string = ""
    v2[(v3 << ((809) & 63))] = v20
    v1 -= (v1 & (v1 - 0))
    if (127 >= -616) {
        for fi21 := 0; fi21 < 20; fi21++ {
            v22 := v5[("hello" + v20)]
            if false {
                switch fi21 {
                case 32768:
                    v23 := GenS2{F0: 42, F1: 1, F2: v18}
                    _ = v23
                }
                for fi24 := 0; fi24 < 8; fi24++ {
                    var v25 int = fi24
                    v26 := []int{v25}
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            } else {
                if (1e10 <= -2.25) {
                    for fi27 := 0; fi27 < 12; fi27++ {
                        v20 = ((v20 + v20) + v20)
                        _ = fi27
                    }
                } else {
                    for fi28 := 0; fi28 < 1; fi28++ {
                        _ = len(v5)
                        _ = fi28
                    }
                    fmt.Println(v22, fi21)
                }
            }
            _ = fi21
            _ = v22
        }
    } else {
        fmt.Println(v18, v19)
        var v29 int = int(v3)
        v29 += 4294967295
        _ = v29
    }
    var v30 uint8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
