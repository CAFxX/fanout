package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 uint16, p1 uint16) string {
    var v6 float64 = 3.14159
    v7 := GenS1{F0: -614}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return ""
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]string)
    var v3 uint16 = 255
    var v4 uint = 0
    var v5 int16 = 574
    v4 += 65535
    var v8 int = 9223372036854775807
    v9 := GenS1{F0: 384}
    for fi10 := 0; fi10 < 6; fi10++ {
        v5 = (v5 * v5)
        if (false || (42 != -1)) {
            v11 := GenS1{F0: 42}
            _ = v11
        } else {
            v12 := GenS1{F0: -982}
            if true {
                var v13 uint64 = 18446744073709551615
                _ = v13
            }
            _ = v12
        }
        _ = fi10
    }
    var v14 int = v8
    v1[(v8 - 1)] = 0
    var v15 int16 = (v5 - (v5 % (((v5 << (((v4 << ((v4) & 63))) & 63))) | 1)))
    if false {
        v16 := GenS1{F0: 42}
        for fi17 := 0; fi17 < 9; fi17++ {
            v8 -= v8
            _ = fi17
        }
        _ = v16
    }
    switch v14 {
    case 32767:
        var v18 int32 = -790
        _ = v18
    case 9223372036854775807:
        var v19 uint32 = 984
        _ = v19
    }
    var v20 uint32 = uint32(v3)
    v21 := GenS1{F0: 42}
    for fi22 := 0; fi22 < 4; fi22++ {
        for fi23 := 0; fi23 < 4; fi23++ {
            var v24 int16 = v15
            _ = fi23
            _ = v24
        }
        var v25 float64 = 0.0
        v3 += v3
        _ = fi22
        _ = v25
    }
    var v26 string = ""
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
