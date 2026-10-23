package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 uint32; F1 int32; F2 uint8 }
type GenS4 struct { F0 uint8; F1 int16; F2 int16 }
type GenS5 struct { F0 int16 }
func genFn0(p0 uint32) bool {
    var v7 float32 = ((1.5 / -2.25) / 1e10)
    v8 := GenS1{F0: 1}
    var v9 int8 = 42
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return (2147483647 >= (p0 >> ((324) & 63)))
}

func main() {
    var v1 uint32 = 1
    var v2 string = ("hello" + "\n\t")
    v3 := []int{0, -9, -9223372036854775808}
    v4 := make(map[int]uint32)
    var v5 int32 = 926
    var v6 uint32 = 42
    for fi10 := 0; fi10 < 8; fi10++ {
        v5 = (v5 & v5)
        _ = fi10
    }
    v11 := GenS2{F0: 65535, F1: 0}
    v5 -= (v5 | (v5 - (v5 ^ (v5 ^ v5))))
    v6 = v6
    _ = len(v3)
    v12 := GenS3{F0: 2, F1: 144, F2: 234}
    if (1.5 >= (0.1 * 0.0)) {
        if false {
            v6 = (v6 % (((v6 * v6)) | 1))
            v13 := GenS3{F0: v1, F1: v5, F2: 193}
            var v14 uint = 634
            _ = v13
            _ = v14
        } else {
            for fi15 := 0; fi15 < 11; fi15++ {
                for fi16 := 0; fi16 < 3; fi16++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v17 uint8 = 128
                    _ = fi16
                    _ = v17
                }
                _ = fi15
            }
        }
        v1 += 65535
    } else {
        fmt.Println(v2, v5)
        fmt.Println(v1, v1)
        for fi18 := 0; fi18 < 13; fi18++ {
            v19 := v4[fi18]
            _ = fi18
            _ = v19
        }
    }
    v20 := GenS4{F0: 0, F1: -32768, F2: 0}
    if true {
        v4[-990] = (v1 << ((255) & 63))
        for fi21 := 0; fi21 < 1; fi21++ {
            v22 := GenS5{F0: int16(v1)}
            if ((!(!(0 <= 1000))) && true) {
                fmt.Println(fi21, v6, v2)
                var v23 uint = 1
                v23 += v23
                _ = v23
            }
            _ = fi21
            _ = v22
        }
        v5 -= 0
    } else {
        v1 -= v6
        for fi24 := 0; fi24 < 18; fi24++ {
            for fi25 := 0; fi25 < 5; fi25++ {
                v26 := GenS4{F0: 2, F1: -676, F2: 42}
                v27 := GenS3{F0: (v6 >> ((2) & 63)), F1: v5, F2: 0}
                var v28 uint16 = 1
                var v29 uint = 4294967295
                _ = fi25
                _ = v26
                _ = v27
                _ = v28
                _ = v29
            }
            for fi30 := 0; fi30 < 6; fi30++ {
                var v31 int32 = v5
                _ = fi30
                _ = v31
            }
            var v32 bool = (255 < 255)
            _ = fi24
            _ = v32
        }
        v6 = (v6 * v1)
    }
    var v33 uint16 = 0
    v34 := GenS4{F0: 2, F1: -690, F2: -32768}
    v35 := GenS3{F0: (v6 + 0), F1: v5, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
