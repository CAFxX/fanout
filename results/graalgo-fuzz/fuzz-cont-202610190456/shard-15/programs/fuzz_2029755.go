package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 uint16 }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 bool }
func genFn0(p0 uint64, p1 uint32) uint64 {
    v4 := []int{}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    _ = p0
    _ = p1
    _ = v4
    return (p0 % ((p0) | 1))
}
func genFn1() uint8 {
    var v5 uint = 264
    fmt.Println(v5)
    var v6 uint64 = 515
    fmt.Println(v6, v5)
    _ = v5
    _ = v6
    return 11
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]string)
    var v3 float32 = -2.25
    v7 := GenS1{F0: -9223372036854775808, F1: 0, F2: 1}
    v3 += v3
    if false {
        for fi8 := 0; fi8 < 15; fi8++ {
            fmt.Println(v3)
            v3 = (1.5 - float32(float64((v3 * v3))))
            switch fi8 {
            case 9223372036854775807:
                if true {
                    v9 := []int{42, fi8, (fi8 ^ fi8)}
                    for fi10 := 0; fi10 < 16; fi10++ {
                        var v11 int16 = -1
                        _ = fi10
                        _ = v11
                    }
                    v12 := GenS2{F0: 65535, F1: ((0.0 * 1e10) != v3)}
                    _ = v12
                    _ = v9
                } else {
                    v13 := []int{fi8, fi8}
                    v3 = 0.0
                    var v14 float32 = float32(float64((1.5 + v3)))
                    _ = v13
                    _ = v14
                }
            default:
                v3 = 0.1
            }
            _ = fi8
        }
    } else {
        fmt.Println(v3)
    }
    var v15 uint64 = 98
    var v16 int64 = -1
    v15 = genFn0((v15 | (v15 - (^(v15)))), 0)
    v17 := GenS2{F0: 32767, F1: (v3 > v3)}
    v16 -= v16
    if (4294967295 <= 167) {
        if true {
            if (true && (!(!(!(!((!(!true)) && false)))))) {
                var v18 float64 = -2.25
                v18 += (1e10 / (-2.25 + float64((v3 + (1.5 / v3)))))
                _ = v18
            } else {
                v19 := []int{}
                v20 := v2[((((("\n\t" + "\n\t") + "foo bar") + "") + "") + "hello")]
                _ = v19
                _ = v20
            }
            var v21 int8 = 127
            var v22 uint16 = 32767
            _ = v21
            _ = v22
        }
        for fi23 := 0; fi23 < 7; fi23++ {
            switch "" {
            case "foo bar":
                var v24 string = ""
                _ = v24
            }
            v16 -= 9223372036854775807
            _ = fi23
        }
    } else {
        for fi25 := 0; fi25 < 1; fi25++ {
            var v26 uint64 = genFn0((v15 / (((v15 ^ genFn0(v15, 0))) | 1)), 902)
            if (true && false) {
                if (120 >= 1000) {
                    var v27 int16 = 32767
                    v15 = 467
                    fmt.Println(fi25, v16)
                    _ = v27
                }
            } else {
                v15 = v26
                if (9223372036854775807 >= (^(fi25))) {
                    v28 := GenS3{F0: (1 > 127)}
                    fmt.Println(v16)
                    _ = v28
                } else {
                    v29 := []int{fi25, fi25, fi25}
                    _ = v29
                }
                for fi30 := 0; fi30 < 4; fi30++ {
                    if ((fi30 & fi30) >= fi30) {
                        fmt.Println(v3, v16, v3)
                        v16 = (v16 & (v16 * v16))
                    } else {
                        v16 = (v16 >> ((775) & 63))
                        v16 = v16
                    }
                    switch "a" {
                    case "\n\t":
                        if true {
                            var v31 int8 = 42
                            _ = v31
                        } else {
                            v32 := GenS2{F0: 0, F1: false}
                            _ = v32
                        }
                    case "foo bar":
                        fmt.Println(fi30)
                    case "":
                        v26 += (^(v15))
                    }
                    _ = fi30
                }
            }
            _ = fi25
            _ = v26
        }
    }
    var v33 int16 = -1
    var v34 uint64 = v15
    var v35 string = "\n\t"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v7
}
