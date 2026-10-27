package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32; F1 int16 }
type GenS5 struct { F0 int64 }
func genFn0() int32 {
    v5 := GenS1{F0: 4294967295}
    var v6 string = "\n\t"
    var v7 int64 = -155
    v8 := GenS1{F0: 0}
    v9 := []int{433, 32768}
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return -2147483648
}
func genFn1(p0 int16) string {
    p0 -= 1000
    if false {
        var v10 string = "foo bar"
        _ = v10
    } else {
        v11 := GenS2{F0: false}
        var v12 uint8 = 216
        _ = v11
        _ = v12
    }
    if false {
        for fi13 := 0; fi13 < 2; fi13++ {
            if true {
                if (!(4294967295 <= 528)) {
                    v14 := GenS1{F0: 2147483647}
                    fmt.Println(fi13, p0)
                    _ = v14
                }
                v15 := GenS3{F0: false}
                _ = v15
            } else {
                v16 := GenS3{F0: false}
                _ = v16
            }
            if (2147483647 >= 318) {
                if true {
                    if false {
                        if true {
                            var v17 bool = (((("hello" + ("\n\t" + "\n\t")) + ("hello" + ("foo bar" + ""))) != (("foo bar" + "\n\t") + ("foo bar" + "hello"))) && false)
                            _ = v17
                        } else {
                            var v18 string = "a"
                            _ = v18
                        }
                    }
                }
            }
            v19 := GenS4{F0: 2147483648, F1: int16(fi13)}
            _ = fi13
            _ = v19
        }
        if (float64((-2.25 * 0.1)) >= (float64((0.0 - float32((1e10 / 1.0)))) / 1.5)) {
            v20 := []int{-9223372036854775808, 9223372036854775807}
            var v21 float64 = 0.0
            _ = v20
            _ = v21
        } else {
            p0 += (^(p0))
            v22 := []int{}
            var v23 int32 = -2147483648
            _ = v22
            _ = v23
        }
    }
    p0 += (^(p0))
    var v24 int32 = -2147483648
    _ = p0
    _ = v24
    return ("a" + "hello")
}

func main() {
    var v1 uint64 = 42
    v2 := []int{127}
    var v3 int8 = 0
    var v4 uint = 255
    var v25 int16 = 126
    v25 += (v25 ^ (v25 | v25))
    var v26 string = "a"
    if true {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v3 -= v3
    } else {
        if false {
            var v27 int16 = v25
            v3 -= v3
            _ = v27
        } else {
            v28 := GenS5{F0: 42}
            _ = len(v2)
            var v29 int8 = (v3 - v3)
            _ = v28
            _ = v29
        }
        v26 = "\n\t"
    }
    fmt.Println(v3, v26)
    if len(v2) > 0 {
        v2[0] = -9223372036854775808
    }
    v25 += v25
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
