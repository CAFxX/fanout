package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 uint64; F2 int64 }
type GenS3 struct { F0 int8; F1 string }
func genFn0(p0 string) uint8 {
    v7 := []int{}
    var v8 uint8 = 255
    for fi9 := 0; fi9 < 16; fi9++ {
        var v10 float32 = (-2.25 * 0.1)
        p0 = ("" + p0)
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 int8 = 121
    var v2 uint = 42
    v3 := make(map[int]string)
    v4 := make(map[uint32]uint32)
    var v5 uint8 = 102
    v6 := make(map[uint32]string)
    if true {
        v5 -= 127
        v11 := []int{2147483647, 127, 32768, -9223372036854775808}
        _ = v11
    } else {
        for fi12 := 0; fi12 < 8; fi12++ {
            v2 -= v2
            v13 := []int{fi12}
            v5 += genFn0("foo bar")
            _ = fi12
            _ = v13
        }
        for fi14 := 0; fi14 < 5; fi14++ {
            v15 := []int{(fi14 + -1)}
            if len(v15) > 0 {
                v15[0] = (fi14 & (fi14 << ((v2) & 63)))
            }
            _ = fi14
            _ = v15
        }
        v5 -= 128
    }
    v2 = v2
    fmt.Println(v5, v2, v1)
    fmt.Println(v2)
    if false {
        v1 += 45
        v2 -= v2
        if (!false) {
            v1 += v1
            if (-789 == 128) {
                v16 := GenS1{F0: (v5 % ((127) | 1))}
                var v17 uint32 = 65535
                for fi18 := 0; fi18 < 3; fi18++ {
                    var v19 int32 = 979
                    _ = fi18
                    _ = v19
                }
                _ = v16
                _ = v17
            }
            var v20 int8 = (-(v1))
            _ = v20
        }
    } else {
        if true {
            v21 := GenS2{F0: -1, F1: 345, F2: -144}
            v22 := []int{-2147483648}
            if (!((!(v1 <= (-(v1)))) || (42 == -1))) {
                switch 32767 {
                case 336:
                    var v23 int32 = -1
                    _ = v23
                case -9223372036854775808:
                    if false {
                        var v24 int = 32768
                        v25 := GenS3{F0: (v1 << ((v2) & 63)), F1: (("foo bar" + (("a" + "") + "\n\t")) + "foo bar")}
                        if len(v22) >= 2 {
                            v22 = v22[1:len(v22)]
                        }
                        _ = v24
                        _ = v25
                    } else {
                        var v26 int64 = 9223372036854775807
                        _ = v26
                    }
                case 9223372036854775807:
                    delete(v4, 2147483648)
                default:
                    _ = len(v22)
                }
            }
            _ = v21
            _ = v22
        } else {
            fmt.Println(v1)
            v5 = (v5 + (v5 ^ (v5 >> (((v2 % ((v2) | 1))) & 63))))
        }
        v5 += 85
        if false {
            for fi27 := 0; fi27 < 11; fi27++ {
                fmt.Println(v5, v5, v1)
                var v28 int64 = -9223372036854775808
                _ = fi27
                _ = v28
            }
            v5 += 128
            delete(v6, 0)
        } else {
            var v29 uint8 = 5
            v6[614] = ("\n\t" + (("hello" + (("foo bar" + (("\n\t" + "hello") + "foo bar")) + "hello")) + ""))
            _ = v29
        }
    }
    v1 = (v1 & (v1 << ((255) & 63)))
    for fi30 := 0; fi30 < 17; fi30++ {
        var v31 uint = v2
        var v32 uint32 = 434
        fmt.Println(v5, v1, v2)
        _ = fi30
        _ = v31
        _ = v32
    }
    delete(v4, 2147483647)
    v33 := []int{42, -392, 9223372036854775807}
    v2 -= (^(v2))
    v1 += (v1 >> ((479) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
