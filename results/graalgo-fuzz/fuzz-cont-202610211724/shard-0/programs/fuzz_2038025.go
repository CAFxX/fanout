package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int64 }
func genFn0() uint {
    for fi4 := 0; fi4 < 20; fi4++ {
        switch fi4 {
        case 1:
            for fi5 := 0; fi5 < 12; fi5++ {
                if (!(3.14159 >= 0.0)) {
                    var v6 uint64 = 4294967295
                    _ = v6
                } else {
                    var v7 float32 = float32((0.1 / 3.14159))
                    v7 = 1.5
                    v8 := GenS1{F0: fi4, F1: 2, F2: 970}
                    var v9 uint64 = 42
                    v10 := []int{fi4, 1, fi5, (fi4 % (((fi5 ^ (fi4 % ((-1) | 1)))) | 1))}
                    v7 -= v7
                    switch fi5 {
                    case -643:
                        var v11 int = (fi4 ^ 128)
                        fmt.Println(v11, fi5)
                        _ = v11
                    default:
                        if ((100 < -1) || (!false)) {
                            var v12 string = ("hello" + "foo bar")
                            _ = v12
                        } else {
                            v7 += 1e10
                        }
                    }
                    _ = v10
                    _ = v7
                    _ = v8
                    _ = v9
                }
                _ = fi5
            }
        case -663:
            v13 := []int{fi4, -90, 4294967295, 0}
            _ = v13
        default:
            v14 := []int{(fi4 % ((fi4) | 1))}
            _ = v14
        }
        _ = fi4
    }
    var v15 int16 = 0
    v16 := []int{32768}
    _ = v15
    _ = v16
    return 65535
}
func genFn1(p0 uint16) uint64 {
    var v17 int32 = 703
    v18 := []int{127, -1}
    _ = p0
    _ = v17
    _ = v18
    return uint64(p0)
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int32 = 1
    var v3 int16 = 760
    fmt.Println(v2, v2)
    v2 = (v2 >> ((222) & 63))
    var v19 string = ""
    switch 32767 {
    case 32768:
        var v20 int8 = -4
        _ = v20
    case 0:
        for fi21 := 0; fi21 < 17; fi21++ {
            fmt.Println(v19, v19)
            _ = fi21
        }
    case 127:
        v2 -= v2
    default:
        v2 += int32(v3)
    }
    fmt.Println(v2, v19, v3)
    v19 = "\n\t"
    v19 = v19
    v22 := v1[2147483647]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v3
}
