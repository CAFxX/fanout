package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 int8; F1 uint16; F2 int32 }
func genFn0(p0 int16, p1 int, p2 int64) uint32 {
    v5 := GenS1{F0: 1, F1: 1}
    var v6 int64 = (p2 * (p2 / ((p2) | 1)))
    switch ("foo bar" + "") {
    case "a":
        v7 := []int{p1, (p1 | (p1 & p1))}
        _ = v7
    default:
        p1 = 40
    }
    switch p1 {
    case -915:
        if ((!true) && (!false)) {
            for fi8 := 0; fi8 < 20; fi8++ {
                var v9 uint8 = 127
                _ = fi8
                _ = v9
            }
            p0 += (p0 / ((-1) | 1))
        } else {
            for fi10 := 0; fi10 < 15; fi10++ {
                for fi11 := 0; fi11 < 7; fi11++ {
                    fmt.Println(v6, v6)
                    var v12 uint32 = 2147483648
                    p2 = 1
                    var v13 string = ("a" + "\n\t")
                    if false {
                        var v14 int8 = -12
                        _ = v14
                    }
                    v15 := GenS1{F0: 114, F1: (v12 ^ (v12 / (((v12 + v12)) | 1)))}
                    _ = fi11
                    _ = v12
                    _ = v13
                    _ = v15
                }
                v16 := GenS1{F0: 1, F1: 610}
                v17 := []int{fi10, fi10, (-(p1))}
                _ = fi10
                _ = v16
                _ = v17
            }
        }
    }
    p1 += (^(p1))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 uint8 = 187
    v2 := make(map[int]int)
    var v3 float32 = -2.25
    v4 := make(map[string]bool)
    fmt.Println(v3)
    v1 += v1
    v18 := []int{}
    fmt.Println(v1)
    v3 = v3
    delete(v2, 4294967295)
    var v19 uint32 = 4294967295
    var v20 int16 = -229
    var v21 string = "foo bar"
    v4[(v21 + "")] = true
    _ = len(v2)
    v20 -= -1
    for fi22 := 0; fi22 < 2; fi22++ {
        fmt.Println(v21, v19)
        for fi23 := 0; fi23 < 11; fi23++ {
            if (!(true && (-1 != 0))) {
                v24 := GenS2{F0: 0}
                _ = v24
            } else {
                if (0 <= 4294967295) {
                    fmt.Println(v3)
                    v25 := GenS2{F0: 18446744073709551615}
                    _ = v25
                } else {
                    _ = len(v2)
                }
                v26 := GenS3{F0: -1, F1: true}
                _ = v26
            }
            fmt.Println(fi23, v3)
            fmt.Println(v3, v3, v19)
            _ = fi23
        }
        v27 := GenS4{F0: -1, F1: 2, F2: 1}
        _ = fi22
        _ = v27
    }
    switch 464 {
    case 202:
        if true {
            _ = len(v2)
        } else {
            if false {
                v28 := GenS2{F0: 18446744073709551615}
                var v29 uint64 = 983
                for fi30 := 0; fi30 < 5; fi30++ {
                    fmt.Println(v21)
                    _ = fi30
                }
                _ = v28
                _ = v29
            } else {
                var v31 uint8 = v1
                _ = v31
            }
            for fi32 := 0; fi32 < 3; fi32++ {
                var v33 uint64 = 2
                _ = fi32
                _ = v33
            }
        }
    case -2147483648:
        for fi34 := 0; fi34 < 6; fi34++ {
            var v35 string = ("a" + v21)
            _ = fi34
            _ = v35
        }
    case 2147483647:
        var v36 int64 = 902
        _ = v36
    }
    switch ("\n\t" + v21) {
    case "a":
        v20 -= 42
    case "foo bar":
        fmt.Println(v3, v20, v1)
    case "":
        var v37 float64 = 0.0
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
