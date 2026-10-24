package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int }

func main() {
    var v1 uint8 = 127
    var v2 uint8 = (v1 - (v1 % (((v1 * v1)) | 1)))
    v3 := make(map[int]int)
    if true {
        var v4 int32 = -1
        for fi5 := 0; fi5 < 18; fi5++ {
            v4 += v4
            fmt.Println(v4)
            v6 := GenS1{F0: -322, F1: (fi5 % ((fi5) | 1)), F2: 1}
            _ = fi5
            _ = v6
        }
        fmt.Println(v2, v4, v2)
        _ = v4
    }
    v2 = v2
    var v7 float64 = 1e10
    fmt.Println(v2, v7)
    v8 := GenS1{F0: -2147483648, F1: -984, F2: 2}
    v9 := GenS1{F0: 2147483647, F1: 128, F2: 1}
    for fi10 := 0; fi10 < 17; fi10++ {
        for fi11 := 0; fi11 < 6; fi11++ {
            if false {
                v3[1] = fi10
                var v12 uint64 = 4294967295
                v7 -= -2.25
                _ = v12
            } else {
                v13 := GenS2{F0: -1}
                _ = v13
            }
            v14 := []int{fi11, (^(fi10)), -830}
            switch fi11 {
            case -542:
                v15 := GenS3{F0: fi11}
                _ = v15
            case 4294967295:
                v16 := GenS3{F0: 128}
                _ = v16
            default:
                if ((-544 == 1) && (true || (!(492 >= 2)))) {
                    v2 = v2
                    v7 += ((v7 - 1.5) / ((((1.5 * v7) + 3.14159) * v7) + ((((v7 + v7) + 3.14159) * (v7 / v7)) / (v7 - v7))))
                    fmt.Println(v7)
                } else {
                    v17 := GenS1{F0: -1, F1: (fi10 + 0), F2: 1}
                    _ = v17
                }
            }
            if true {
                delete(v3, fi10)
            } else {
                if false {
                    if (!true) {
                        v18 := GenS3{F0: (fi11 % ((fi10) | 1))}
                        _ = v18
                    } else {
                        v7 = float64(3.14159)
                        delete(v3, (fi11 ^ (fi10 / (((fi11 * (-(fi11)))) | 1))))
                    }
                    v19 := GenS3{F0: (fi11 + fi10)}
                    v1 = v1
                    _ = v19
                }
                var v20 bool = ((true || false) && (32767 != 150))
                _ = v20
            }
            _ = fi11
            _ = v14
        }
        _ = fi10
    }
    var v21 bool = true
    if (!v21) {
        v21 = (4294967295 != 1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
