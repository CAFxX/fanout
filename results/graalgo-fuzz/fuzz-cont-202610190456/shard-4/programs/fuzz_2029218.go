package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 uint64; F1 uint; F2 int16 }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint8; F1 string }
func genFn0() int64 {
    var v6 int16 = -44
    var v7 uint8 = 163
    fmt.Println(v6, v7)
    v7 = (v7 ^ 42)
    _ = v6
    _ = v7
    return 1
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint64 = 0
    var v3 int8 = -79
    var v4 int32 = 42
    v5 := make(map[int]uint32)
    if false {
        v8 := GenS1{F0: 930, F1: "a"}
        _ = v8
    }
    var v9 float64 = (1e10 + 1e10)
    switch "hello" {
    case "":
        for fi10 := 0; fi10 < 2; fi10++ {
            v3 += (v3 >> ((255) & 63))
            _ = fi10
        }
    case "a":
        v2 -= v2
    default:
        v11 := GenS1{F0: 2, F1: ("a" + "hello")}
        _ = v11
    }
    if false {
        if (float32(float64(1.5)) != float32((v9 - 1.5))) {
            var v12 uint = 0
            _ = v12
        }
        fmt.Println(v9)
        v13 := []int{0, 232, 42, 32767}
        _ = v13
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        v15 := GenS1{F0: 785, F1: ""}
        if (!(2 >= 0)) {
            v4 = int32(v2)
        } else {
            for fi16 := 0; fi16 < 16; fi16++ {
                var v17 bool = false
                v18 := []int{fi16}
                v2 += (v2 * v2)
                _ = fi16
                _ = v17
                _ = v18
            }
        }
        _ = fi14
        _ = v15
    }
    fmt.Println(v9)
    v3 += (v3 + 62)
    v19 := GenS1{F0: 32767, F1: ("hello" + (("" + "a") + "hello"))}
    v20 := []int{776}
    if false {
        v2 = 572
    } else {
        fmt.Println(v4, v9, v3)
        var v21 int16 = 32767
        if false {
            var v22 int32 = v4
            if (846 < 0) {
                fmt.Println(v4, v9)
            } else {
                v2 -= v2
                var v23 bool = (339 > uint16(v2))
                _ = v23
            }
            _ = v22
        } else {
            v24 := GenS2{F0: 2147483648, F1: 1}
            var v25 int = 127
            for fi26 := 0; fi26 < 19; fi26++ {
                v21 += (v21 << ((136) & 63))
                var v27 bool = false
                switch 127 {
                case 127:
                    if (2 == 0) {
                        v25 += (v25 * (fi26 / ((v25) | 1)))
                        v4 = v4
                    } else {
                        v5[fi26] = 2
                    }
                case 128:
                    v28 := GenS3{F0: v2, F1: 0, F2: (v21 & 1000)}
                    _ = v28
                case -992:
                    if len(v20) >= 2 {
                        v20 = v20[1:len(v20)]
                    }
                default:
                    if false {
                        var v29 int64 = -1
                        var v30 float32 = 0.1
                        _ = v29
                        _ = v30
                    }
                }
                _ = fi26
                _ = v27
            }
            _ = v24
            _ = v25
        }
        _ = v21
    }
    v31 := GenS4{F0: "foo bar", F1: 127}
    v32 := GenS5{F0: 3}
    v33 := GenS6{F0: 42, F1: "\n\t"}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v9
}
