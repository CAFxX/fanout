package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 int16 }
type GenS2 struct { F0 int16; F1 uint; F2 string }
type GenS3 struct { F0 uint32 }
func genFn0(p0 uint32) int32 {
    v7 := []int{1, 32768}
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 uint = 311
        _ = fi8
        _ = v9
    }
    var v10 uint8 = 2
    v10 -= (^(v10))
    _ = p0
    _ = v10
    _ = v7
    return 42
}

func main() {
    var v1 bool = (1 >= 2)
    var v2 uint32 = 2
    var v3 int32 = -2147483648
    v4 := []int{127}
    var v5 uint = 734
    var v6 int = -66
    v1 = ((v6 << ((65535) & 63)) <= v6)
    var v11 int32 = int32(v6)
    fmt.Println(v5)
    _ = len(v4)
    for fi12 := 0; fi12 < 19; fi12++ {
        v6 = (fi12 << (((v5 & v5)) & 63))
        _ = fi12
    }
    v3 -= v3
    for fi13 := 0; fi13 < 16; fi13++ {
        for fi14 := 0; fi14 < 20; fi14++ {
            var v15 float64 = ((0.0 - 1e10) / -2.25)
            if len(v4) > 0 {
                v4[0] = v6
            }
            v16 := GenS1{F0: 127, F1: -32, F2: -32768}
            fmt.Println(v6)
            _ = fi14
            _ = v15
            _ = v16
        }
        for fi17 := 0; fi17 < 10; fi17++ {
            fmt.Println(v2)
            switch -2147483648 {
            case 32768:
                if v1 {
                    if (!v1) {
                        v11 = int32(v2)
                        if v1 {
                            switch fi17 {
                            case 4294967295:
                                if len(v4) > 0 {
                                    v4[0] = (fi13 % (((v6 % (((-(v6))) | 1))) | 1))
                                }
                            case 32768:
                                v2 = v2
                            case -2147483648:
                                v5 -= v5
                            }
                            var v18 uint8 = 196
                            _ = v18
                        } else {
                            fmt.Println(v2, v5, fi17)
                        }
                    }
                } else {
                    var v19 bool = false
                    var v20 int32 = v11
                    _ = v19
                    _ = v20
                }
            case 208:
                v6 = (v6 * (fi13 ^ (fi17 * (fi17 / ((fi13) | 1)))))
            }
            fmt.Println(v11)
            _ = fi17
        }
        _ = fi13
    }
    v11 = v11
    var v21 uint32 = (v2 | (v2 % ((v2) | 1)))
    var v22 float32 = 1e10
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v23 := GenS2{F0: 1, F1: v5, F2: ("a" + ("foo bar" + "hello"))}
    if (v11 >= -533) {
        v4 = append(v4, 4294967295)
    } else {
        fmt.Println(v11)
    }
    for fi24 := 0; fi24 < 3; fi24++ {
        var v25 int = -439
        if v1 {
            fmt.Println(v1, v1, v21)
            v26 := GenS3{F0: v2}
            _ = v26
        } else {
            v1 = v1
            if len(v4) > 0 {
                v4[0] = v6
            }
        }
        var v27 bool = true
        _ = fi24
        _ = v25
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
