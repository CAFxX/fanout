package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 int; F1 string; F2 uint8 }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 uint16 = 42
    var v2 uint32 = 4294967295
    var v3 float32 = (3.14159 * 1e10)
    var v4 uint8 = 127
    v5 := []string{}
    var v6 uint64 = 18446744073709551615
    var v7 uint32 = 1
    if false {
        if (977 > 782) {
            v8 := GenS1{F0: v1, F1: (v6 % ((v6) | 1))}
            for fi9 := 0; fi9 < 17; fi9++ {
                var v10 uint32 = 452
                for fi11 := 0; fi11 < 9; fi11++ {
                    if false {
                        if (!false) {
                            if ((-926 >= 42) || (1 == 127)) {
                                v12 := GenS1{F0: (v1 / ((102) | 1)), F1: (v6 << ((65535) & 63))}
                                if (!(!(false || false))) {
                                    _ = len(v5)
                                } else {
                                    var v13 uint8 = v4
                                    _ = v13
                                }
                                fmt.Println(v7, v10, v3)
                                _ = v12
                            }
                            v14 := GenS1{F0: v1, F1: v6}
                            v4 -= (v4 & (^(v4)))
                            v10 += v2
                            _ = v14
                        }
                        var v15 float64 = 3.14159
                        _ = v15
                    }
                    _ = fi11
                }
                v2 = (v2 ^ (^(v2)))
                _ = fi9
                _ = v10
            }
            var v16 int = 480
            _ = v16
            _ = v8
        }
        var v17 int64 = 1
        _ = v17
    } else {
        v1 = 42
        var v18 int64 = 1
        for fi19 := 0; fi19 < 3; fi19++ {
            switch (-(fi19)) {
            case -2147483648:
                var v20 uint64 = 42
                _ = v20
            }
            _ = fi19
        }
        _ = v18
    }
    if ((1 < 190) && true) {
        var v21 int32 = 0
        _ = v21
    } else {
        v22 := GenS1{F0: v1, F1: (v6 & (v6 ^ v6))}
        v3 += ((0.1 * 3.14159) + 1e10)
        var v23 int32 = 2147483647
        _ = v22
        _ = v23
    }
    v3 = (1.5 + 0.0)
    if (!((v2 | (v7 + (^(v7)))) < 2)) {
        v24 := GenS2{F0: int(v4), F1: "\n\t", F2: (v4 - uint8(v1))}
        _ = v24
    }
    var v25 string = ""
    v26 := GenS3{F0: 18446744073709551615}
    if (-790 <= 42) {
        fmt.Println(v7)
        fmt.Println(v1)
        switch (v25 + ((v25 + v25) + ("a" + v25))) {
        case "":
            v7 += v7
        case "hello":
            for fi27 := 0; fi27 < 17; fi27++ {
                v28 := GenS3{F0: (^(v6))}
                var v29 uint16 = (v1 % (((v1 + v1)) | 1))
                _ = fi27
                _ = v28
                _ = v29
            }
        }
    }
    v30 := GenS1{F0: 335, F1: 2}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v3 += (v3 / 1.5)
    var v31 uint16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
