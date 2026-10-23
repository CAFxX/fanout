package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32 }

func main() {
    var v1 bool = false
    v2 := make(map[int]uint32)
    var v3 uint16 = 218
    v4 := make(map[uint32]uint32)
    var v5 uint8 = 0
    fmt.Println(v1)
    for fi6 := 0; fi6 < 3; fi6++ {
        switch -9223372036854775808 {
        case -583:
            v5 += v5
        }
        _ = fi6
    }
    delete(v4, 4294967295)
    v7 := GenS1{F0: -2147483648, F1: v3, F2: -9223372036854775808}
    v1 = (47 < -9223372036854775808)
    for fi8 := 0; fi8 < 16; fi8++ {
        fmt.Println(v1, v3, fi8)
        var v9 int = (-(fi8))
        _ = fi8
        _ = v9
    }
    if v1 {
        var v10 int = 9223372036854775807
        v3 = 255
        _ = v10
    }
    v11 := []int{int(v5), -2147483648}
    v11 = append(v11, -769)
    var v12 uint32 = 1
    var v13 float64 = 3.14159
    v14 := GenS2{F0: 127}
    v15 := GenS3{F0: int32(v3)}
    v3 -= v3
    if true {
        v16 := GenS3{F0: 0}
        for fi17 := 0; fi17 < 15; fi17++ {
            v18 := GenS2{F0: -92}
            switch fi17 {
            case 0:
                for fi19 := 0; fi19 < 18; fi19++ {
                    if false {
                        var v20 uint = 238
                        _ = v20
                    } else {
                        v21 := GenS3{F0: -2147483648}
                        v12 = 810
                        fmt.Println(v12, fi19)
                        _ = v21
                    }
                    _ = len(v2)
                    var v22 int16 = -32768
                    _ = fi19
                    _ = v22
                }
            case -510:
                if (2147483647 <= (^(v12))) {
                    v5 += (v5 >> ((968) & 63))
                }
            case -9223372036854775808:
                v23 := GenS1{F0: 0, F1: v3, F2: -1}
                _ = v23
            default:
                var v24 uint = 0
                _ = v24
            }
            _ = fi17
            _ = v18
        }
        _ = v16
    }
    var v25 float32 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
