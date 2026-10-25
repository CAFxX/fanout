package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 int16; F1 int16; F2 int32 }
func genFn0(p0 int, p1 uint32) uint8 {
    switch p0 {
    case 162:
        v6 := GenS1{F0: 2}
        _ = v6
    }
    if (2 > 863) {
        if true {
            p0 = (^(p0))
            if (2 > 2) {
                p1 = p1
            } else {
                switch "hello" {
                case "\n\t":
                    p1 -= 269
                }
                var v7 uint8 = 2
                var v8 uint8 = (v7 + (v7 ^ (v7 & 24)))
                _ = v7
                _ = v8
            }
        }
        fmt.Println(p0)
    } else {
        v9 := []int{1, (p0 | (p0 % (((p0 - 127)) | 1))), -208}
        p0 = p0
        var v10 float64 = float64(((float32(-2.25) + 1.5) - float32(1.5)))
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    return 29
}
func genFn1() string {
    var v11 uint64 = 42
    v11 -= (^(v11))
    v11 = (v11 >> ((452) & 63))
    if false {
        if false {
            v11 -= 1
            v11 += v11
            for fi12 := 0; fi12 < 11; fi12++ {
                if true {
                    v13 := GenS1{F0: 1}
                    for fi14 := 0; fi14 < 8; fi14++ {
                        fmt.Println(fi12, fi12, fi12)
                        _ = fi14
                    }
                    _ = v13
                } else {
                    v11 -= (v11 * (v11 << ((65535) & 63)))
                }
                for fi15 := 0; fi15 < 19; fi15++ {
                    v16 := GenS1{F0: 65535}
                    v17 := []int{}
                    _ = len(v17)
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                _ = fi12
            }
        } else {
            v18 := []int{871, int(v11), -527}
            var v19 uint64 = v11
            _ = v18
            _ = v19
        }
    } else {
        v11 = (v11 % ((42) | 1))
        var v20 uint64 = v11
        fmt.Println(v11, v20)
        _ = v20
    }
    _ = v11
    return "a"
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint64 = 4294967295
    var v3 uint8 = 127
    var v4 int32 = -723
    v5 := make(map[int]uint32)
    v21 := []int{}
    var v22 uint16 = 61
    var v23 uint = 565
    v24 := GenS2{F0: 42, F1: 807}
    if len(v21) > 0 {
        v21[0] = -507
    }
    var v25 int32 = 1
    fmt.Println(v23)
    for fi26 := 0; fi26 < 2; fi26++ {
        fmt.Println(v25, v3)
        _ = fi26
    }
    v22 += v22
    for fi27 := 0; fi27 < 2; fi27++ {
        for fi28 := 0; fi28 < 20; fi28++ {
            v4 = (v25 / (((v4 + v25)) | 1))
            v29 := GenS3{F0: 32767, F1: 32767, F2: (v25 / (((v4 * (v25 | (v25 % ((v4) | 1))))) | 1))}
            v22 -= v22
            if false {
                if false {
                    v25 = int32(fi27)
                    v25 += (v4 / (((v4 | v25)) | 1))
                } else {
                    v25 = -2147483648
                    fmt.Println(v2)
                }
            } else {
                v21 = append(v21, (fi28 + fi27))
            }
            _ = fi28
            _ = v29
        }
        _ = fi27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
