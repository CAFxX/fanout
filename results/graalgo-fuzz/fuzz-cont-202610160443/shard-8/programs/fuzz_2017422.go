package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 int8 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]bool)
    var v3 uint = 255
    var v4 uint64 = 18446744073709551615
    var v5 int = -70
    v4 += 401
    if (((v5 - v5) != v5) || (-2.25 >= 3.14159)) {
        for fi6 := 0; fi6 < 19; fi6++ {
            v5 -= v5
            _ = fi6
        }
    } else {
        var v7 uint64 = (v4 - (v4 * 953))
        if (442 == 32767) {
            fmt.Println(v4, v4, v3)
            if true {
                if (!(3.14159 == -2.25)) {
                    v3 = (v3 / ((65535) | 1))
                    _ = len(v1)
                } else {
                    v7 += 1
                    _ = len(v2)
                    for fi8 := 0; fi8 < 18; fi8++ {
                        v9 := []int{}
                        v3 += v3
                        _ = fi8
                        _ = v9
                    }
                }
                v10 := GenS1{F0: uint(v4), F1: v4}
                _ = v10
            } else {
                v11 := []int{32768, (v5 - (v5 % (((v5 % (((v5 & 2147483647)) | 1))) | 1))), v5}
                if true {
                    var v12 uint8 = 42
                    _ = v12
                } else {
                    v13 := GenS1{F0: (v3 * v3), F1: v4}
                    var v14 int = -2147483648
                    _ = v13
                    _ = v14
                }
                if ((2 > 0) && (0.1 < 1e10)) {
                    for fi15 := 0; fi15 < 19; fi15++ {
                        v4 -= (^(v4))
                        if len(v11) > 0 {
                            v11[0] = (fi15 % ((v5) | 1))
                        }
                        _ = fi15
                    }
                    fmt.Println(v7, v7, v4)
                }
                _ = v11
            }
            var v16 uint8 = 136
            _ = v16
        } else {
            v7 -= (v7 & v4)
            for fi17 := 0; fi17 < 20; fi17++ {
                fmt.Println(fi17)
                _ = fi17
            }
        }
        var v18 int16 = -453
        _ = v18
        _ = v7
    }
    v5 += (v5 - -1)
    for fi19 := 0; fi19 < 4; fi19++ {
        v20 := GenS2{F0: 42}
        v2[(("hello" + "hello") + "a")] = false
        _ = fi19
        _ = v20
    }
    fmt.Println(v4)
    v5 = 32768
    v3 += 68
    v21 := GenS1{F0: 0, F1: (v4 % ((v4) | 1))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
