package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 uint16 }
func genFn0() uint8 {
    if true {
        var v6 uint = 2
        fmt.Println(v6)
        switch "\n\t" {
        case "a":
            v6 += (v6 / (((^(v6))) | 1))
        }
        _ = v6
    }
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 int16 = 506
        _ = fi7
        _ = v8
    }
    return 0
}
func genFn1(p0 int64) string {
    p0 += p0
    p0 += (p0 ^ p0)
    if false {
        v9 := []int{}
        _ = v9
    }
    v10 := []int{1, 788, -9223372036854775808}
    _ = p0
    _ = v10
    return ("" + ("a" + "hello"))
}

func main() {
    var v1 uint16 = 42
    v2 := make(map[int]uint32)
    var v3 uint16 = v1
    var v4 bool = false
    var v5 uint8 = 0
    for fi11 := 0; fi11 < 13; fi11++ {
        if v4 {
            v1 += (v1 * v3)
        }
        var v12 string = ""
        for fi13 := 0; fi13 < 13; fi13++ {
            if false {
                v14 := GenS1{F0: 112}
                fmt.Println(v1, v4, v1)
                var v15 int8 = 42
                _ = v14
                _ = v15
            }
            fmt.Println(v3, v1, v4)
            _ = fi13
        }
        _ = fi11
        _ = v12
    }
    var v16 int64 = 852
    v16 += (v16 + -120)
    fmt.Println(v5, v5, v3)
    var v17 uint = 2
    if true {
        v17 += v17
        v18 := []int{128, -1, 568}
        _ = v18
    }
    v3 = v3
    v1 += (v3 / (((^(v1))) | 1))
    v19 := []int{249, 42}
    var v20 uint64 = 2
    fmt.Println(v3)
    switch 1 {
    case 32767:
        v21 := GenS2{F0: (v5 | genFn0()), F1: -402}
        _ = v21
    case -675:
        var v22 uint64 = (v20 & (v20 & (^(v20))))
        _ = v22
    }
    for fi23 := 0; fi23 < 3; fi23++ {
        if v4 {
            fmt.Println(v5, v1)
            v24 := GenS2{F0: v5, F1: 990}
            _ = v24
        } else {
            for fi25 := 0; fi25 < 8; fi25++ {
                v26 := GenS3{F0: (v3 / (((v1 << (((v17 % (((v17 / ((v17) | 1))) | 1))) & 63))) | 1))}
                var v27 int8 = 1
                if v4 {
                    _ = len(v2)
                    var v28 uint = uint(v5)
                    if (!(v16 > v16)) {
                        var v29 uint64 = 314
                        v28 += (v17 % ((v28) | 1))
                        delete(v2, fi23)
                        _ = v29
                    } else {
                        if false {
                            v30 := v2[-9223372036854775808]
                            _ = v30
                        } else {
                            v4 = (((^(v27)) >= -1) && (-2147483648 >= 1))
                            var v31 string = (("\n\t" + ("a" + "foo bar")) + "hello")
                            switch -2147483648 {
                            case 1:
                                v32 := GenS2{F0: 24, F1: int16(v17)}
                                _ = v32
                            case 128:
                                var v33 uint16 = v1
                                _ = v33
                            }
                            _ = v31
                        }
                    }
                    _ = v28
                }
                _ = fi25
                _ = v26
                _ = v27
            }
        }
        delete(v2, 1)
        _ = fi23
    }
    v5 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
