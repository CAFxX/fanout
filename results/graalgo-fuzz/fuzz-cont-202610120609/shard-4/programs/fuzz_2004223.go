package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 int16 }
func genFn0(p0 int8) int64 {
    p0 += -39
    v7 := []int{}
    _ = p0
    _ = v7
    return 0
}
func genFn1() int {
    v8 := GenS1{F0: -1}
    var v9 uint = 18446744073709551615
    _ = v8
    _ = v9
    return -2147483648
}

func main() {
    v1 := []string{""}
    var v2 uint = 18446744073709551615
    v3 := []int{-574}
    v4 := make(map[int]bool)
    var v5 uint32 = 0
    v6 := []int{}
    v5 -= v5
    var v10 float64 = 1.5
    fmt.Println(v2)
    var v11 float64 = v10
    var v12 float64 = 1e10
    v13 := GenS2{F0: v5, F1: 42}
    v12 -= (v11 + v10)
    if (-9223372036854775808 != 635) {
        v14 := GenS1{F0: 733}
        fmt.Println(v2)
        var v15 int32 = 1
        _ = v14
        _ = v15
    } else {
        var v16 uint = v2
        _ = len(v3)
        _ = v16
    }
    var v17 int = 9223372036854775807
    if false {
        if ((!(false || ((^(v5)) != v5))) && (0 < -1)) {
            v12 -= v12
            var v18 int = v17
            if false {
                fmt.Println(v2, v11, v12)
                if (v11 >= v11) {
                    for fi19 := 0; fi19 < 6; fi19++ {
                        if false {
                            fmt.Println(v17, fi19)
                            if (!(!(0 <= -902))) {
                                var v20 uint8 = 1
                                if false {
                                    for fi21 := 0; fi21 < 11; fi21++ {
                                        v5 = (v5 ^ 4294967295)
                                        _ = fi21
                                    }
                                } else {
                                    var v22 string = "\n\t"
                                    _ = v22
                                }
                                _ = v20
                            }
                        }
                        _ = fi19
                    }
                }
            }
            _ = v18
        }
    }
    var v23 int64 = -660
    v24 := GenS1{F0: -32768}
    v17 = -563
    var v25 int8 = 1
    if true {
        for fi26 := 0; fi26 < 12; fi26++ {
            var v27 int16 = 429
            if true {
                for fi28 := 0; fi28 < 2; fi28++ {
                    var v29 int64 = (v23 ^ v23)
                    switch ("foo bar" + "\n\t") {
                    case "hello":
                        v4[fi26] = (!false)
                    }
                    var v30 int8 = int8(v27)
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                v31 := GenS2{F0: (v5 % (((v5 - 2)) | 1)), F1: (v27 & 1)}
                v5 += v5
                _ = v31
            }
            _ = fi26
            _ = v27
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
