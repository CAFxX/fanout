package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 bool) uint {
    v7 := GenS1{F0: 2}
    p0 = (p0 || (!p0))
    var v8 float64 = 1.5
    _ = p0
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]int)
    var v3 int16 = -32768
    var v4 bool = (2147483648 == 2147483648)
    var v5 float32 = -2.25
    var v6 uint32 = 0
    var v9 int8 = 127
    var v10 int64 = 1
    fmt.Println(v9, v4)
    if v4 {
        v3 = (v3 % (((v3 - 32767)) | 1))
    }
    v1["hello"] = 775
    for fi11 := 0; fi11 < 9; fi11++ {
        v5 = (v5 + (1e10 + v5))
        v9 += 1
        var v12 uint32 = v6
        _ = fi11
        _ = v12
    }
    v6 = (v6 * (^(v6)))
    v13 := GenS2{F0: 4294967295}
    for fi14 := 0; fi14 < 10; fi14++ {
        for fi15 := 0; fi15 < 8; fi15++ {
            var v16 uint = 4294967295
            if true {
                v4 = v4
                fmt.Println(v6, v9, fi14)
            } else {
                v17 := GenS2{F0: 1}
                _ = v17
            }
            switch fi15 {
            case -9223372036854775808:
                fmt.Println(v3)
            default:
                v2[v6] = fi14
            }
            _ = fi15
            _ = v16
        }
        for fi18 := 0; fi18 < 18; fi18++ {
            if (!v4) {
                fmt.Println(v10, v9)
                var v19 uint = 18446744073709551615
                _ = v19
            }
            var v20 uint8 = 2
            _ = fi18
            _ = v20
        }
        v5 = v5
        _ = fi14
    }
    for fi21 := 0; fi21 < 14; fi21++ {
        for fi22 := 0; fi22 < 4; fi22++ {
            if false {
                v10 += 400
            } else {
                if v4 {
                    var v23 uint = 876
                    v6 += v6
                    fmt.Println(v10, v5, v9)
                    _ = v23
                }
                v24 := GenS2{F0: 42}
                _ = v24
            }
            v25 := GenS2{F0: 4294967295}
            v4 = (!v4)
            _ = fi22
            _ = v25
        }
        _ = fi21
    }
    var v26 int = -2147483648
    var v27 float32 = 1e10
    fmt.Println(v3, v3, v5)
    v26 = (^(v26))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
