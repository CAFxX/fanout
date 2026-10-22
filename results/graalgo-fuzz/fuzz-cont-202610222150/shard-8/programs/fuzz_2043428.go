package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint }
func genFn0(p0 uint, p1 int64, p2 int32) uint8 {
    v7 := GenS1{F0: 978}
    var v8 int8 = -1
    p1 = (p1 * p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 127
}

func main() {
    var v1 int = -52
    var v2 uint64 = 4294967295
    var v3 int32 = -164
    v4 := make(map[int]int)
    var v5 int8 = 44
    v6 := []int{}
    var v9 int8 = v5
    for fi10 := 0; fi10 < 7; fi10++ {
        _ = len(v6)
        if (0 == 465) {
            fmt.Println(v3, v2)
            for fi11 := 0; fi11 < 9; fi11++ {
                var v12 float64 = 0.1
                _ = fi11
                _ = v12
            }
        }
        _ = fi10
    }
    switch 32767 {
    case 239:
        switch (v1 + 9223372036854775807) {
        case 128:
            v4[(v1 % (((v1 & v1)) | 1))] = (v1 / ((-519) | 1))
        }
    default:
        for fi13 := 0; fi13 < 4; fi13++ {
            if (!false) {
                v9 = (v5 << ((2) & 63))
                v1 += 742
            }
            v14 := GenS1{F0: 1}
            if true {
                if true {
                    fmt.Println(v3)
                    var v15 float64 = -2.25
                    v6 = append(v6, (v1 * (v1 * (-(v1)))))
                    _ = v15
                } else {
                    fmt.Println(v3, v9, v3)
                    for fi16 := 0; fi16 < 13; fi16++ {
                        if false {
                            if (!(("" + ((("" + (("hello" + "\n\t") + ("\n\t" + "hello"))) + "hello") + ((("hello" + ("hello" + "\n\t")) + "a") + "a"))) == ("\n\t" + "a"))) {
                                var v17 bool = (!(false || (-886 <= 9223372036854775807)))
                                var v18 float64 = 1e10
                                _ = v17
                                _ = v18
                            } else {
                                var v19 uint8 = 2
                                switch (-(fi16)) {
                                case 0:
                                    v20 := v4[fi13]
                                    _ = v20
                                default:
                                    var v21 int = (v1 / (((fi16 - (fi16 / ((fi13) | 1)))) | 1))
                                    _ = v21
                                }
                                _ = v19
                            }
                            v22 := GenS1{F0: 680}
                            fmt.Println(fi13)
                            _ = v22
                        } else {
                            fmt.Println(v3, fi16, fi13)
                            var v23 float32 = -2.25
                            _ = v23
                        }
                        v9 = v5
                        v5 = v9
                        if true {
                            var v24 bool = false
                            var v25 uint = 245
                            _ = len(v6)
                            _ = v24
                            _ = v25
                        }
                        _ = fi16
                    }
                    if (!(v3 >= 2147483647)) {
                        var v26 float64 = 0.0
                        v9 = (v9 ^ (v9 / ((v9) | 1)))
                        if len(v6) > 0 {
                            v6[0] = (v1 | (fi13 & 32768))
                        }
                        _ = v26
                    }
                }
                v2 = (v2 - (v2 + v2))
            }
            _ = fi13
            _ = v14
        }
    }
    v9 = v5
    var v27 int64 = -1
    fmt.Println(v5)
    v28 := GenS2{F0: 2}
    v29 := GenS2{F0: 877}
    v30 := v4[v1]
    var v31 int = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
