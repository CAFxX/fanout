package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 int8; F1 uint8; F2 uint64 }
type GenS3 struct { F0 uint32; F1 uint16 }
func genFn0(p0 int64) uint64 {
    var v5 string = "\n\t"
    fmt.Println(v5)
    v5 = v5
    var v6 int32 = 0
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 uint64 = 544
        v8 = (v8 << ((0) & 63))
        fmt.Println(p0, v8)
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = v5
    _ = v6
    return 1
}

func main() {
    var v1 int8 = -128
    var v2 int32 = 0
    v3 := []int{}
    v4 := []int{1}
    var v9 int16 = 734
    switch "\n\t" {
    case "foo bar":
        v9 += 1
    }
    var v10 uint32 = 2147483648
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v3 = append(v3, 786)
    v11 := GenS1{F0: (v9 - v9), F1: 498}
    if (97 < 127) {
        var v12 int64 = 905
        _ = v12
    }
    if len(v3) > 0 {
        v3[0] = -1
    }
    if false {
        for fi13 := 0; fi13 < 6; fi13++ {
            var v14 int16 = v9
            var v15 int8 = (-(v1))
            v14 = v14
            _ = fi13
            _ = v14
            _ = v15
        }
        var v16 int8 = 1
        var v17 int64 = 9223372036854775807
        _ = v16
        _ = v17
    }
    v9 = (-(v9))
    for fi18 := 0; fi18 < 17; fi18++ {
        for fi19 := 0; fi19 < 7; fi19++ {
            _ = len(v4)
            v20 := GenS1{F0: -1, F1: v10}
            _ = fi19
            _ = v20
        }
        if false {
            fmt.Println(v9)
            if true {
                if false {
                    for fi21 := 0; fi21 < 13; fi21++ {
                        if len(v4) > 0 {
                            v4[0] = fi18
                        }
                        v10 = v10
                        _ = fi21
                    }
                }
                var v22 int32 = (^(v2))
                _ = v22
            }
            v23 := GenS1{F0: 1000, F1: (v10 % (((v10 & v10)) | 1))}
            _ = v23
        } else {
            if true {
                for fi24 := 0; fi24 < 11; fi24++ {
                    v25 := GenS1{F0: (v9 / ((32767) | 1)), F1: 1}
                    if len(v3) > 0 {
                        v3[0] = int(v9)
                    }
                    _ = fi24
                    _ = v25
                }
                var v26 uint16 = 622
                if (0 == v9) {
                    var v27 string = ((("\n\t" + "") + "hello") + "hello")
                    var v28 uint64 = 0
                    var v29 int64 = 9223372036854775807
                    _ = v27
                    _ = v28
                    _ = v29
                }
                _ = v26
            }
            v30 := GenS2{F0: (v1 ^ (v1 >> ((177) & 63))), F1: 234, F2: 2}
            _ = v30
        }
        v31 := GenS2{F0: 64, F1: 127, F2: 1}
        _ = fi18
        _ = v31
    }
    var v32 int = 2147483647
    if (0 == 18446744073709551615) {
        var v33 int32 = 1
        fmt.Println(v2, v1)
        _ = v33
    }
    v1 -= (v1 ^ v1)
    fmt.Println(v1, v1)
    var v34 uint8 = 177
    v35 := GenS3{F0: v10, F1: 65535}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v32
    _ = v34
    _ = v35
    _ = v4
    _ = v9
}
