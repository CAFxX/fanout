package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint64 }
type GenS2 struct { F0 uint8; F1 uint32; F2 uint8 }
func genFn0(p0 int8, p1 uint16, p2 uint8) uint32 {
    var v5 uint64 = 18446744073709551615
    var v6 int32 = 248
    var v7 uint = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 int64, p1 string, p2 int) uint8 {
    if true {
        var v8 float32 = (-2.25 / 0.1)
        for fi9 := 0; fi9 < 5; fi9++ {
            v10 := []int{fi9, (fi9 | (^(fi9))), (^(fi9))}
            _ = fi9
            _ = v10
        }
        _ = v8
    }
    var v11 uint16 = 525
    p0 = 42
    var v12 string = (p1 + p1)
    p0 -= -392
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 1
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]uint32)
    var v3 bool = ((991 > 156) && true)
    var v4 float32 = ((1e10 * -2.25) - 0.1)
    var v13 int16 = 1
    fmt.Println(v13, v3, v3)
    fmt.Println(v4, v3, v4)
    var v14 uint8 = 1
    v14 -= genFn1(9223372036854775807, "", 128)
    v15 := v1[121]
    var v16 uint32 = 700
    v3 = (1 > 1)
    var v17 uint32 = genFn0(100, 1, (v14 % ((genFn1(-491, "a", v15)) | 1)))
    switch (^(v15)) {
    case 1:
        var v18 uint8 = v14
        _ = v18
    }
    var v19 bool = v3
    if (v19 || (-128 > -18)) {
        if ((-9223372036854775808 != 0) || (!false)) {
            var v20 int8 = 42
            fmt.Println(v3)
            v21 := GenS1{F0: 565, F1: v3, F2: 794}
            _ = v20
            _ = v21
        } else {
            if ((!true) && (1.5 != 1.5)) {
                v19 = true
            } else {
                v13 -= v13
                if v19 {
                    switch -1 {
                    case 2147483647:
                        fmt.Println(v4, v3)
                    case -1:
                        if true {
                            if (v4 < v4) {
                                v16 += 65535
                                var v22 int32 = 0
                                var v23 uint = 4294967295
                                _ = v22
                                _ = v23
                            } else {
                                var v24 uint32 = (v17 * 1)
                                _ = v24
                            }
                            v3 = (2 < 42)
                        } else {
                            switch (v15 * v15) {
                            case 2147483647:
                                var v25 int = -888
                                _ = v25
                            case 9223372036854775807:
                                var v26 uint16 = 111
                                _ = v26
                            }
                            for fi27 := 0; fi27 < 4; fi27++ {
                                fmt.Println(v13, v13, fi27)
                                _ = fi27
                            }
                        }
                    }
                    for fi28 := 0; fi28 < 11; fi28++ {
                        v29 := GenS2{F0: v14, F1: v17, F2: (v14 & 0)}
                        _ = fi28
                        _ = v29
                    }
                    delete(v1, (v17 << ((540) & 63)))
                } else {
                    v30 := []int{v15, 32768, v15}
                    v16 = (v17 % ((uint32(v14)) | 1))
                    v3 = v3
                    _ = v30
                }
            }
            switch "a" {
            case "":
                v13 -= v13
            case "hello":
                _ = len(v2)
            }
            delete(v1, genFn0(-1, 42, (^(v14))))
        }
        var v31 bool = ((!false) || (1 <= v15))
        _ = v31
    } else {
        switch "foo bar" {
        case "":
            v15 += (^(v15))
        }
        var v32 uint16 = 65535
        _ = v32
    }
    v33 := GenS2{F0: (v14 << ((971) & 63)), F1: 2147483648, F2: (^(v14))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v33
    _ = v4
}
