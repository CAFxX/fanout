package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 int16; F1 bool }

func main() {
    var v1 bool = false
    v2 := make(map[uint32]bool)
    var v3 int32 = 280
    var v4 int8 = 42
    var v5 bool = ((-235 < -1) && (int64(v4) == 1))
    var v6 int64 = 1
    var v7 float64 = (1.5 + -2.25)
    v6 -= v6
    v8 := GenS1{F0: 4294967295}
    var v9 bool = (0 != 964)
    var v10 int64 = (v6 % ((-540) | 1))
    var v11 uint = 765
    switch "a" {
    case "\n\t":
        v10 = (v10 << ((v11) & 63))
    default:
        v12 := GenS1{F0: 2}
        _ = v12
    }
    if false {
        v13 := v2[65535]
        var v14 uint8 = 39
        _ = v13
        _ = v14
    } else {
        v15 := []int{-551, 253}
        _ = v15
    }
    if true {
        v3 += v3
    } else {
        for fi16 := 0; fi16 < 10; fi16++ {
            for fi17 := 0; fi17 < 11; fi17++ {
                if ((((3.14159 / 0.1) - (0.0 + 0.0)) - -2.25) <= 0.0) {
                    if v5 {
                        v5 = v5
                    }
                } else {
                    var v18 float64 = v7
                    switch (-(fi17)) {
                    case 127:
                        fmt.Println(v18, v11, v10)
                    case 512:
                        v10 -= v6
                    case 32768:
                        v2[42] = (v9 || (v1 && v1))
                    }
                    v19 := GenS2{F0: v11, F1: 2147483648}
                    _ = v18
                    _ = v19
                }
                v20 := []int{(fi16 + fi16)}
                _ = fi17
                _ = v20
            }
            var v21 uint = (^(v11))
            _ = fi16
            _ = v21
        }
    }
    v11 -= (^(v11))
    var v22 int16 = 1000
    if (100 > 0) {
        var v23 float32 = 1e10
        var v24 int32 = v3
        v6 = v6
        _ = v23
        _ = v24
    } else {
        for fi25 := 0; fi25 < 16; fi25++ {
            if v9 {
                v7 += v7
                if (0.0 < 0.1) {
                    if (!((v4 + 100) > v4)) {
                        for fi26 := 0; fi26 < 11; fi26++ {
                            v27 := GenS2{F0: (v11 << ((v11) & 63)), F1: 1}
                            _ = fi26
                            _ = v27
                        }
                        v28 := []int{-249, fi25, (-(fi25)), (fi25 - fi25)}
                        _ = v28
                    }
                } else {
                    v29 := GenS3{F0: int16(v3), F1: ("a" > ("foo bar" + "hello"))}
                    _ = v29
                }
                for fi30 := 0; fi30 < 6; fi30++ {
                    var v31 uint = (v11 >> ((v11) & 63))
                    v32 := []int{(^(fi25)), fi30}
                    _ = fi30
                    _ = v31
                    _ = v32
                }
            } else {
                v33 := []int{fi25, fi25, (fi25 * (fi25 * fi25))}
                v5 = ((v11 ^ (v11 - (v11 ^ (v11 >> ((v11) & 63))))) <= v11)
                _ = v33
            }
            _ = fi25
        }
    }
    var v34 int32 = 241
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v22
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
