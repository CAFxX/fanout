package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 int32 }

func main() {
    var v1 uint64 = 0
    v2 := []uint32{2, 42, 65535, 2147483647}
    v3 := make(map[uint32]string)
    var v4 int16 = -445
    v5 := []int{32768, -545, 32767}
    var v6 int = -1
    v2 = append(v2, 65535)
    if false {
        for fi7 := 0; fi7 < 11; fi7++ {
            if len(v2) > 0 {
                v2[0] = 266
            }
            if len(v2) > 0 {
                v2[0] = 2
            }
            var v8 bool = false
            _ = fi7
            _ = v8
        }
        delete(v3, 319)
        _ = len(v2)
    } else {
        for fi9 := 0; fi9 < 6; fi9++ {
            switch 128 {
            case 888:
                for fi10 := 0; fi10 < 4; fi10++ {
                    fmt.Println(fi10, v1, v4)
                    var v11 uint16 = 65535
                    if len(v2) > 0 {
                        v2[0] = 2
                    }
                    _ = fi10
                    _ = v11
                }
            case -410:
                if true {
                    var v12 int8 = 42
                    v13 := GenS1{F0: 42}
                    var v14 float64 = 1e10
                    _ = v12
                    _ = v13
                    _ = v14
                } else {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
            case 128:
                var v15 uint16 = 32767
                _ = v15
            }
            v16 := GenS2{F0: ""}
            _ = fi9
            _ = v16
        }
        var v17 int8 = 127
        _ = v17
    }
    if false {
        v4 = (^(v4))
    }
    if false {
        _ = len(v5)
        var v18 int16 = v4
        _ = v18
    }
    var v19 uint8 = 135
    fmt.Println(v4)
    delete(v3, 487)
    v20 := GenS3{F0: (v4 / ((0) | 1)), F1: 970}
    var v21 bool = ((!true) || (!true))
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
