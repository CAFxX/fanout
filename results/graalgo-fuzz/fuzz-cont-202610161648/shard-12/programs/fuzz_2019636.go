package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 int8 }
func genFn0(p0 int) uint8 {
    for fi7 := 0; fi7 < 1; fi7++ {
        p0 = (p0 ^ p0)
        _ = fi7
    }
    for fi8 := 0; fi8 < 17; fi8++ {
        if false {
            for fi9 := 0; fi9 < 15; fi9++ {
                v10 := GenS1{F0: 2147483647, F1: 933, F2: 0}
                _ = fi9
                _ = v10
            }
            var v11 int = (^(p0))
            v12 := []int{(v11 / (((v11 << ((65535) & 63))) | 1)), fi8}
            _ = v11
            _ = v12
        }
        v13 := GenS1{F0: 1, F1: 837, F2: -45}
        _ = fi8
        _ = v13
    }
    fmt.Println(p0)
    _ = p0
    return 230
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint64 = 0
    v3 := make(map[int]string)
    var v4 float32 = 1.5
    var v5 uint16 = 32767
    var v6 int64 = 1
    v14 := GenS1{F0: 455, F1: (v5 + uint16(v2)), F2: -1}
    var v15 string = "hello"
    delete(v1, 4294967295)
    v15 = v15
    var v16 uint8 = 42
    if (!false) {
        for fi17 := 0; fi17 < 19; fi17++ {
            var v18 uint16 = uint16(v16)
            _ = fi17
            _ = v18
        }
    } else {
        for fi19 := 0; fi19 < 17; fi19++ {
            if true {
                var v20 int32 = 1
                var v21 int8 = -1
                var v22 uint32 = uint32(v5)
                _ = v20
                _ = v21
                _ = v22
            } else {
                v5 += 583
            }
            if (true && ((fi19 % (((^(fi19))) | 1)) <= (fi19 & fi19))) {
                v5 += 255
            }
            _ = fi19
        }
    }
    var v23 bool = false
    for fi24 := 0; fi24 < 14; fi24++ {
        if (v6 >= -858) {
            v25 := []int{fi24, -1, 32767}
            v5 = v5
            _ = v25
        }
        var v26 int64 = (v6 >> ((2) & 63))
        _ = fi24
        _ = v26
    }
    v27 := GenS1{F0: 26, F1: v5, F2: -80}
    v15 = v15
    var v28 float32 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
