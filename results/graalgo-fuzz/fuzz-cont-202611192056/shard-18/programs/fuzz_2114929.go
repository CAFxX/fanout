package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 uint64 }
type GenS2 struct { F0 uint8; F1 uint; F2 string }
type GenS3 struct { F0 uint32 }
func genFn0() int64 {
    for fi4 := 0; fi4 < 8; fi4++ {
        v5 := []int{}
        if true {
            switch fi4 {
            case -252:
                v6 := GenS1{F0: 32768, F1: -1, F2: 2}
                _ = v6
            case 32768:
                var v7 int16 = 32767
                _ = v7
            case -2147483648:
                var v8 int32 = 946
                _ = v8
            }
            if len(v5) > 0 {
                v5[0] = -659
            }
            if false {
                var v9 bool = (128 > 255)
                _ = v9
            }
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi4
        _ = v5
    }
    var v10 int32 = 0
    switch -9223372036854775808 {
    case 32768:
        for fi11 := 0; fi11 < 2; fi11++ {
            var v12 string = ""
            v10 += (v10 % ((639) | 1))
            v10 = (v10 % ((v10) | 1))
            _ = fi11
            _ = v12
        }
    case 32767:
        var v13 uint64 = 42
        _ = v13
    case 227:
        v10 -= v10
    default:
        v14 := []int{33}
        _ = v14
    }
    _ = v10
    return -472
}

func main() {
    var v1 uint = 463
    v2 := make(map[int]string)
    var v3 float64 = float64(-2.25)
    fmt.Println(v1, v3)
    var v15 int16 = 92
    if true {
        switch "\n\t" {
        case "\n\t":
            if (0 > -1) {
                v3 -= ((3.14159 / 1.0) - v3)
            } else {
                var v16 uint32 = 20
                _ = v16
            }
        case "foo bar":
            v15 = (^(v15))
        case "hello":
            var v17 int64 = -528
            _ = v17
        default:
            for fi18 := 0; fi18 < 6; fi18++ {
                if true {
                    v19 := GenS1{F0: 42, F1: -1, F2: 18446744073709551615}
                    _ = v19
                } else {
                    var v20 uint8 = 1
                    delete(v2, fi18)
                    _ = v20
                }
                v3 = v3
                _ = fi18
            }
        }
        var v21 float32 = 1.5
        v1 -= (^(v1))
        _ = v21
    } else {
        v22 := GenS2{F0: 2, F1: 42, F2: "hello"}
        v3 += float64(1.5)
        _ = v22
    }
    for fi23 := 0; fi23 < 12; fi23++ {
        v15 = v15
        _ = fi23
    }
    v3 = (v3 - float64(1e10))
    switch 42 {
    case -641:
        if true {
            if true {
                if true {
                    v15 = (v15 * (v15 ^ (v15 * (-(v15)))))
                } else {
                    _ = len(v2)
                    v24 := []int{}
                    fmt.Println(v3, v1)
                    _ = v24
                }
            }
        } else {
            v3 = 3.14159
            v25 := GenS1{F0: -1, F1: -128, F2: 42}
            for fi26 := 0; fi26 < 12; fi26++ {
                switch 485 {
                case 32767:
                    for fi27 := 0; fi27 < 17; fi27++ {
                        fmt.Println(v1)
                        _ = fi27
                    }
                default:
                    for fi28 := 0; fi28 < 18; fi28++ {
                        var v29 uint = 1
                        _ = fi28
                        _ = v29
                    }
                }
                for fi30 := 0; fi30 < 7; fi30++ {
                    var v31 uint = 455
                    v32 := GenS3{F0: 42}
                    _ = fi30
                    _ = v31
                    _ = v32
                }
                v33 := v2[fi26]
                _ = fi26
                _ = v33
            }
            _ = v25
        }
    case 652:
        v34 := []int{4294967295, -86, 295, 42}
        _ = v34
    case 128:
        delete(v2, 2147483647)
    }
    v35 := []int{4294967295}
    v1 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v35
}
