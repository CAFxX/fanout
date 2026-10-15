package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 int32; F2 uint32 }

func main() {
    var v1 float64 = 1.5
    v2 := make(map[string]uint32)
    var v3 uint8 = 255
    var v4 bool = (3.14159 < (v1 * (0.1 + float64(3.14159))))
    var v5 int8 = 17
    for fi6 := 0; fi6 < 5; fi6++ {
        for fi7 := 0; fi7 < 1; fi7++ {
            _ = len(v2)
            switch fi6 {
            case 32767:
                v4 = (335 < -798)
            default:
                var v8 int8 = (-(v5))
                _ = v8
            }
            v9 := v2["a"]
            _ = fi7
            _ = v9
        }
        _ = fi6
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        v11 := GenS1{F0: 386}
        for fi12 := 0; fi12 < 12; fi12++ {
            _ = len(v2)
            v1 -= -2.25
            _ = len(v2)
            _ = fi12
        }
        _ = fi10
        _ = v11
    }
    delete(v2, ("" + "hello"))
    var v13 uint = 4294967295
    v14 := GenS1{F0: 414}
    for fi15 := 0; fi15 < 5; fi15++ {
        var v16 int16 = -733
        for fi17 := 0; fi17 < 4; fi17++ {
            fmt.Println(v4, v1)
            var v18 int64 = 42
            _ = fi17
            _ = v18
        }
        fmt.Println(v5, v16)
        _ = fi15
        _ = v16
    }
    var v19 string = "\n\t"
    switch "" {
    case "\n\t":
        if v4 {
            fmt.Println(v3, v4)
            v3 = 255
        }
    default:
        var v20 uint = 2
        _ = v20
    }
    fmt.Println(v5, v3)
    v3 = v3
    v4 = v4
    v21 := GenS2{F0: 1, F1: int32(v13), F2: 649}
    switch "\n\t" {
    case "a":
        for fi22 := 0; fi22 < 20; fi22++ {
            switch (fi22 * -9223372036854775808) {
            case 42:
                if (1 <= int16(v3)) {
                    v13 = 42
                    var v23 uint32 = 4294967295
                    _ = v23
                }
            case 461:
                v2[v19] = 4294967295
            case 128:
                if ((v5 ^ (^(v5))) < -101) {
                    var v24 int64 = 283
                    fmt.Println(v13, v4)
                    _ = v24
                } else {
                    var v25 float32 = (((3.14159 + 1.5) * 1e10) / (-2.25 * 1.5))
                    v26 := []int{(fi22 & 9223372036854775807), fi22}
                    _ = v25
                    _ = v26
                }
            default:
                var v27 float64 = (v1 - 0.0)
                _ = v27
            }
            for fi28 := 0; fi28 < 18; fi28++ {
                fmt.Println(fi22)
                fmt.Println(v19, v1)
                v29 := []int{(fi22 << ((v13) & 63)), fi28, (fi22 ^ (fi28 * -2147483648)), -1}
                _ = fi28
                _ = v29
            }
            _ = fi22
        }
    }
    v4 = (!(!(42 == 269)))
    var v30 int32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
