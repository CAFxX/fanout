package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 int8 = 100
    v2 := make(map[int]int)
    var v3 int64 = 868
    var v4 uint = 65535
    var v5 int32 = 636
    var v6 float32 = 3.14159
    switch -93 {
    case 692:
        for fi7 := 0; fi7 < 2; fi7++ {
            fmt.Println(v6)
            var v8 float64 = 0.0
            for fi9 := 0; fi9 < 3; fi9++ {
                v4 += 255
                var v10 uint64 = 4294967295
                _ = fi9
                _ = v10
            }
            _ = fi7
            _ = v8
        }
    default:
        fmt.Println(v6, v5)
    }
    v2[9223372036854775807] = 32768
    if (!true) {
        for fi11 := 0; fi11 < 5; fi11++ {
            var v12 int = -1
            _ = fi11
            _ = v12
        }
        v13 := []int{909, 42}
        _ = v13
    }
    for fi14 := 0; fi14 < 11; fi14++ {
        if true {
            for fi15 := 0; fi15 < 7; fi15++ {
                v3 += (v3 ^ -627)
                _ = fi15
            }
            fmt.Println(v1, v4)
            var v16 string = (("a" + ("hello" + "hello")) + "a")
            _ = v16
        }
        fmt.Println(fi14, v1, fi14)
        v17 := GenS1{F0: 2147483648}
        _ = fi14
        _ = v17
    }
    v18 := GenS1{F0: 586}
    v5 = (v5 >> ((42) & 63))
    for fi19 := 0; fi19 < 5; fi19++ {
        for fi20 := 0; fi20 < 2; fi20++ {
            var v21 uint = (v4 / (((v4 % ((v4) | 1))) | 1))
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
