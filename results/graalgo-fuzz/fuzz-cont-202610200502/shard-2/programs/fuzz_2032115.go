package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 uint32 }
type GenS2 struct { F0 uint8; F1 uint8; F2 int64 }
type GenS3 struct { F0 uint64; F1 int; F2 int32 }
type GenS4 struct { F0 int16 }

func main() {
    var v1 uint32 = 2147483648
    var v2 int32 = 975
    v3 := make(map[string]bool)
    var v4 uint64 = 4294967295
    var v5 uint64 = 204
    v6 := []int{}
    if true {
        var v7 uint8 = uint8(v1)
        _ = v7
    }
    for fi8 := 0; fi8 < 2; fi8++ {
        v2 += -1
        if true {
            v9 := GenS1{F0: -182, F1: 2, F2: (^(v1))}
            _ = v9
        } else {
            if false {
                for fi10 := 0; fi10 < 7; fi10++ {
                    v11 := GenS2{F0: 0, F1: 145, F2: 9223372036854775807}
                    fmt.Println(v1, v5)
                    _ = fi10
                    _ = v11
                }
                if ((((1e10 / 1.0) / (-2.25 + (1e10 * 1e10))) >= (3.14159 / 0.1)) || (true && true)) {
                    v12 := GenS3{F0: v4, F1: (fi8 % ((-173) | 1)), F2: -1}
                    if true {
                        fmt.Println(v2, v5)
                        v13 := GenS1{F0: (v2 / ((2147483647) | 1)), F1: 0, F2: (v1 & v1)}
                        var v14 bool = false
                        _ = v13
                        _ = v14
                    } else {
                        switch (^(fi8)) {
                        case 9223372036854775807:
                            var v15 int8 = 61
                            _ = v15
                        case 819:
                            v16 := GenS2{F0: 0, F1: 128, F2: -9223372036854775808}
                            _ = v16
                        case -9223372036854775808:
                            v1 -= (v1 % ((v1) | 1))
                        }
                    }
                    v3["a"] = (false && true)
                    _ = v12
                } else {
                    v2 -= v2
                    var v17 uint = 18446744073709551615
                    _ = v17
                }
            }
            delete(v3, "")
        }
        if (false || false) {
            for fi18 := 0; fi18 < 7; fi18++ {
                if ((-81 != 100) && (-392 != 0)) {
                    v19 := GenS2{F0: 1, F1: 128, F2: -1}
                    var v20 int64 = -993
                    var v21 bool = (int(v5) == (-(fi8)))
                    _ = v19
                    _ = v20
                    _ = v21
                } else {
                    v4 = v4
                }
                _ = fi18
            }
        } else {
            var v22 string = ("\n\t" + (("" + "") + ("foo bar" + "hello")))
            for fi23 := 0; fi23 < 19; fi23++ {
                v24 := GenS4{F0: 42}
                fmt.Println(v1)
                _ = fi23
                _ = v24
            }
            _ = v22
        }
        _ = fi8
    }
    _ = len(v3)
    for fi25 := 0; fi25 < 4; fi25++ {
        fmt.Println(v5, v4)
        _ = fi25
    }
    if len(v6) > 0 {
        v6[0] = -9223372036854775808
    }
    var v26 int16 = 1000
    v3[""] = (false || ("foo bar" == ""))
    v5 -= v4
    var v27 int = -9223372036854775808
    v28 := GenS3{F0: v4, F1: 416, F2: v2}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
