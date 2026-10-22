package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 bool }
type GenS2 struct { F0 uint32; F1 int64 }

func main() {
    var v1 float64 = 0.1
    var v2 int64 = 42
    var v3 bool = (2 == 42)
    var v4 uint = 42
    var v5 int16 = -225
    var v6 string = "\n\t"
    v1 = (1.5 - float64(3.14159))
    if (-573 >= (-(v5))) {
        if v3 {
            v7 := GenS1{F0: 42, F1: 918, F2: (v5 < (v5 * (v5 / (((-(v5))) | 1))))}
            var v8 int = -1
            switch ("a" + v6) {
            case "foo bar":
                var v9 float32 = float32(v1)
                _ = v9
            default:
                var v10 string = ("\n\t" + v6)
                _ = v10
            }
            _ = v7
            _ = v8
        } else {
            for fi11 := 0; fi11 < 9; fi11++ {
                if v3 {
                    v12 := []int{-878, int(v2)}
                    _ = len(v12)
                    v1 = (v1 / v1)
                    _ = v12
                }
                _ = fi11
            }
            switch 4294967295 {
            case 644:
                var v13 int = 2147483647
                _ = v13
            case -232:
                v3 = (2 == 623)
            default:
                v14 := []int{0, 358, 32767}
                _ = v14
            }
        }
    }
    if ("" != "") {
        fmt.Println(v3, v3, v5)
        var v15 bool = (1000 == v5)
        fmt.Println(v4)
        _ = v15
    } else {
        fmt.Println(v4, v1)
    }
    var v16 uint8 = 255
    v17 := []int{}
    var v18 uint8 = v16
    var v19 uint8 = v16
    var v20 int8 = 44
    v21 := GenS1{F0: 1, F1: 42, F2: true}
    switch (v6 + v6) {
    case "a":
        v4 = v4
    default:
        for fi22 := 0; fi22 < 13; fi22++ {
            if (((v20 << (((v4 >> (((^(v4))) & 63))) & 63)) < v20) || (int32(v20) >= 1)) {
                if v3 {
                    switch (("\n\t" + v6) + (("hello" + (((v6 + "") + "\n\t") + v6)) + v6)) {
                    case "hello":
                        v1 = (v1 * (3.14159 * v1))
                    case "foo bar":
                        for fi23 := 0; fi23 < 20; fi23++ {
                            switch fi22 {
                            case 128:
                                v24 := GenS1{F0: 0, F1: 364, F2: false}
                                v25 := GenS2{F0: 0, F1: 1}
                                _ = v24
                                _ = v25
                            case -1:
                                var v26 int16 = (v5 / ((v5) | 1))
                                _ = v26
                            default:
                                var v27 float32 = (((1e10 - 3.14159) - 0.1) + float32(float64(((-2.25 + float32(v1)) / (0.1 - float32((float64((0.0 - (3.14159 - (1e10 / 0.1)))) + v1)))))))
                                _ = v27
                            }
                            _ = fi23
                        }
                    case "\n\t":
                        if len(v17) > 0 {
                            v17[0] = (fi22 ^ (^(fi22)))
                        }
                    }
                    for fi28 := 0; fi28 < 10; fi28++ {
                        var v29 int8 = (v20 & v20)
                        _ = fi28
                        _ = v29
                    }
                } else {
                    v16 += (v19 & (^(v16)))
                }
                if len(v17) >= 2 {
                    v17 = v17[1:len(v17)]
                }
            }
            _ = fi22
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
