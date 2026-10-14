package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 string }
func genFn0() int {
    v5 := []int{}
    _ = len(v5)
    var v6 int8 = -77
    var v7 uint16 = 65535
    _ = v5
    _ = v6
    _ = v7
    return 9223372036854775807
}

func main() {
    var v1 int16 = 32767
    var v2 uint64 = 75
    var v3 uint32 = 2
    var v4 float64 = 1.5
    v4 = float64(0.1)
    if (((false && (!(((1.5 / float64(1.5)) > v4) && (!((0.0 * (1.5 / (0.0 + float32(v4)))) < float32((-2.25 - v4))))))) && false) && (1 != 9223372036854775807)) {
        v2 += (^(v2))
        v3 -= (v3 | (v3 << ((0) & 63)))
    } else {
        v8 := GenS1{F0: false}
        _ = v8
    }
    var v9 float32 = 0.0
    var v10 int16 = 263
    v11 := GenS1{F0: (!(false && ((v10 * v10) > (v10 % ((32767) | 1)))))}
    v3 -= (v3 / ((156) | 1))
    for fi12 := 0; fi12 < 11; fi12++ {
        v1 = (v1 - (^(v10)))
        for fi13 := 0; fi13 < 11; fi13++ {
            v14 := GenS1{F0: false}
            if true {
                v15 := GenS1{F0: true}
                switch "a" {
                case "foo bar":
                    var v16 string = ("hello" + "")
                    v4 = 0.1
                    v16 = (v16 + (v16 + v16))
                    _ = v16
                case "a":
                    var v17 float64 = (v4 / v4)
                    if false {
                        fmt.Println(v9, v3)
                    }
                    v1 += (v10 - v1)
                    _ = v17
                }
                fmt.Println(v9, v4)
                _ = v15
            } else {
                if true {
                    var v18 int16 = (v1 << ((0) & 63))
                    _ = v18
                } else {
                    if false {
                        var v19 bool = ((499 < 65535) || false)
                        _ = v19
                    } else {
                        v3 += 65535
                        v20 := []int{(fi13 + 2147483647)}
                        _ = v20
                    }
                }
                v21 := GenS2{F0: v10, F1: 128}
                if false {
                    v22 := []int{(fi13 ^ genFn0())}
                    if false {
                        fmt.Println(v2)
                        var v23 int8 = 100
                        fmt.Println(v23, v23, v3)
                        _ = v23
                    }
                    _ = v22
                } else {
                    v24 := GenS1{F0: true}
                    var v25 int8 = 1
                    v26 := []int{genFn0()}
                    _ = v24
                    _ = v25
                    _ = v26
                }
                var v27 uint8 = 255
                fmt.Println(v9, v3)
                _ = v21
                _ = v27
            }
            v28 := []int{42, fi12, 669}
            _ = fi13
            _ = v14
            _ = v28
        }
        v29 := GenS3{F0: "a"}
        _ = fi12
        _ = v29
    }
    var v30 uint8 = uint8(v2)
    var v31 int8 = -1
    if true {
        v32 := GenS1{F0: (false && (genFn0() <= -961))}
        v10 += v1
        var v33 int = int(v2)
        _ = v32
        _ = v33
    } else {
        v34 := []int{-470, 1, 4294967295, genFn0()}
        _ = v34
    }
    v2 = (v2 % (((v2 - v2)) | 1))
    switch -943 {
    case -1:
        if true {
            v35 := []int{}
            _ = len(v35)
            for fi36 := 0; fi36 < 9; fi36++ {
                v30 += v30
                _ = fi36
            }
            _ = v35
        } else {
            fmt.Println(v30, v31, v4)
        }
    default:
        v30 -= 255
    }
    v37 := GenS3{F0: ("a" + "a")}
    v38 := GenS2{F0: 0, F1: (v30 << ((18446744073709551615) & 63))}
    v39 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v9
}
