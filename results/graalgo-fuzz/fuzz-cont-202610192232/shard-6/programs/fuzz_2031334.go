package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 int, p1 uint8, p2 int8) bool {
    var v5 int32 = -991
    v6 := GenS1{F0: 4294967295}
    var v7 uint8 = (p1 << ((388) & 63))
    var v8 float64 = 1.5
    v9 := GenS1{F0: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return (1 >= -1)
}

func main() {
    v1 := []int{-9223372036854775808}
    v2 := make(map[int]int)
    v3 := make(map[int]bool)
    var v4 uint64 = 1
    v10 := GenS1{F0: 2147483648}
    v4 -= (v4 % (((^(v4))) | 1))
    v4 = 1
    var v11 string = "\n\t"
    var v12 uint64 = 991
    var v13 int64 = -9223372036854775808
    var v14 uint64 = v4
    switch -9223372036854775808 {
    case -2147483648:
        fmt.Println(v4, v12, v13)
    case -9223372036854775808:
        var v15 string = (v11 + "\n\t")
        _ = v15
    }
    v4 = 4294967295
    for fi16 := 0; fi16 < 4; fi16++ {
        switch (v11 + (v11 + "")) {
        case "":
            fmt.Println(v4, v14, v4)
        case "a":
            for fi17 := 0; fi17 < 14; fi17++ {
                v14 = v14
                v12 = v4
                _ = fi17
            }
        }
        v18 := GenS1{F0: 4294967295}
        _ = fi16
        _ = v18
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v4 -= v4
    fmt.Println(v4)
    var v19 int8 = -95
    for fi20 := 0; fi20 < 9; fi20++ {
        switch -57 {
        case 0:
            var v21 int = 32768
            _ = v21
        case 32768:
            _ = len(v1)
        }
        if false {
            if false {
                for fi22 := 0; fi22 < 14; fi22++ {
                    v13 -= (v13 << ((0) & 63))
                    var v23 bool = false
                    _ = fi22
                    _ = v23
                }
                if len(v1) > 0 {
                    v1[0] = fi20
                }
            }
            v11 = (v11 + (v11 + "a"))
            if ((!(v4 == 825)) && false) {
                if false {
                    var v24 int = (fi20 % ((fi20) | 1))
                    _ = v24
                }
            }
        } else {
            switch 9223372036854775807 {
            case -9223372036854775808:
                v13 = (v13 & v13)
            case 9223372036854775807:
                for fi25 := 0; fi25 < 14; fi25++ {
                    if false {
                        switch -190 {
                        case -2147483648:
                            var v26 uint16 = 668
                            _ = v26
                        case 108:
                            if true {
                                v27 := GenS1{F0: 2}
                                v12 += v12
                                _ = v27
                            }
                        case 0:
                            var v28 float32 = -2.25
                            _ = v28
                        }
                        fmt.Println(v14, v13, v12)
                        v2[(fi25 ^ fi25)] = (fi25 / ((-39) | 1))
                        if true {
                            v29 := GenS2{F0: 250}
                            v13 -= (v13 - v13)
                            _ = v29
                        } else {
                            var v30 float32 = float32((-2.25 - 0.1))
                            var v31 int16 = 549
                            _ = v30
                            _ = v31
                        }
                        switch v11 {
                        case "hello":
                            if len(v1) > 0 {
                                v1[0] = fi25
                            }
                        case "a":
                            fmt.Println(v19, v11, fi20)
                        default:
                            v19 = v19
                        }
                        if true {
                            v13 += -9223372036854775808
                        }
                        fmt.Println(v4, v13, v19)
                    }
                    var v32 uint16 = 255
                    _ = fi25
                    _ = v32
                }
            default:
                v33 := v2[(fi20 + (^(fi20)))]
                _ = v33
            }
        }
        _ = fi20
    }
    fmt.Println(v13)
    var v34 int64 = (v13 | (v13 / ((0) | 1)))
    var v35 int8 = (v19 % ((int8(v34)) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
}
