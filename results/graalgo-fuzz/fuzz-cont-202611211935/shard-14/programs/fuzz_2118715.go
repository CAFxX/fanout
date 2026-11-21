package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 bool }
func genFn0(p0 uint) uint8 {
    var v7 int8 = -104
    for fi8 := 0; fi8 < 19; fi8++ {
        for fi9 := 0; fi9 < 16; fi9++ {
            var v10 string = (("" + "hello") + "a")
            _ = fi9
            _ = v10
        }
        if (1 < 1) {
            p0 -= p0
            if true {
                var v11 int8 = v7
                _ = v11
            } else {
                p0 = p0
            }
            v7 = v7
        }
        _ = fi8
    }
    v12 := []int{}
    v13 := GenS1{F0: 1}
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    _ = p0
    _ = v12
    _ = v13
    _ = v7
    return 128
}

func main() {
    var v1 int64 = -1
    var v2 uint32 = 448
    var v3 uint = 372
    v4 := make(map[uint32]int)
    v5 := []int{2147483647, -9223372036854775808}
    v6 := make(map[uint32]string)
    var v14 string = "foo bar"
    v15 := GenS1{F0: (v1 | (v1 - v1))}
    switch v14 {
    case "hello":
        if (!(v1 != -9223372036854775808)) {
            v16 := v4[v2]
            for fi17 := 0; fi17 < 14; fi17++ {
                fmt.Println(v16)
                for fi18 := 0; fi18 < 8; fi18++ {
                    if false {
                        if (!false) {
                            if (v2 != v2) {
                                fmt.Println(fi18, v3)
                                var v19 uint64 = 2
                                fmt.Println(fi17, v3, v14)
                                _ = v19
                            }
                            switch (v16 << (((v3 | v3)) & 63)) {
                            case 2147483647:
                                var v20 int32 = 42
                                _ = v20
                            case -226:
                                if true {
                                    v16 -= (fi18 ^ fi18)
                                    var v21 float32 = float32(3.14159)
                                    _ = v21
                                } else {
                                    if (!(true || (true || (true || false)))) {
                                        v22 := GenS2{F0: (v16 - (v16 + fi18))}
                                        v23 := GenS3{F0: 4294967295, F1: true}
                                        _ = v22
                                        _ = v23
                                    } else {
                                        var v24 float64 = ((1.5 - 0.0) * 1e10)
                                        _ = v24
                                    }
                                }
                            }
                            var v25 bool = (false || (0.1 >= 0.1))
                            _ = v25
                        }
                        var v26 uint8 = 42
                        v5 = append(v5, (fi17 + (fi18 + (v16 % ((fi17) | 1)))))
                        var v27 string = ""
                        _ = v26
                        _ = v27
                    }
                    v28 := GenS1{F0: (v1 - -1)}
                    var v29 int = (^(v16))
                    _ = fi18
                    _ = v28
                    _ = v29
                }
                var v30 uint64 = 1
                _ = fi17
                _ = v30
            }
            _ = v16
        }
    case "\n\t":
        _ = len(v5)
    default:
        v31 := GenS1{F0: 9223372036854775807}
        _ = v31
    }
    switch -2147483648 {
    case 32767:
        v2 = v2
    case -2147483648:
        v32 := GenS3{F0: 4294967295, F1: false}
        _ = v32
    }
    v14 = v14
    v33 := GenS1{F0: v1}
    fmt.Println(v14)
    v14 = (v14 + "a")
    if true {
        if len(v5) > 0 {
            v5[0] = 4294967295
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
