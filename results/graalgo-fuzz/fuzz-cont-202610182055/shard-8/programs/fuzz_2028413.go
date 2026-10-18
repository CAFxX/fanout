package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 bool }
func genFn0() uint16 {
    var v6 int16 = -279
    if false {
        v7 := GenS1{F0: 0}
        var v8 uint32 = 884
        v9 := []int{}
        _ = v7
        _ = v8
        _ = v9
    } else {
        if (-1 == 0) {
            v6 -= (v6 * v6)
            v10 := GenS2{F0: 9223372036854775807, F1: ("hello" >= (("hello" + "a") + "\n\t"))}
            _ = v10
        } else {
            var v11 uint64 = 1
            _ = v11
        }
    }
    v6 -= (v6 | v6)
    v12 := []int{-2147483648, 4294967295, 2147483647, 0}
    if false {
        if (true || (!((4294967295 < 363) && (0 < 1)))) {
            fmt.Println(v6)
        } else {
            for fi13 := 0; fi13 < 18; fi13++ {
                if (!(fi13 <= (fi13 % (((fi13 + fi13)) | 1)))) {
                    var v14 int = fi13
                    var v15 int16 = -1
                    _ = v14
                    _ = v15
                }
                v6 = int16(fi13)
                _ = fi13
            }
        }
        v6 -= -711
        if false {
            var v16 bool = ((34 <= 100) || true)
            _ = v16
        }
    }
    _ = v12
    _ = v6
    return 42
}
func genFn1() string {
    if (42 == 65535) {
        if false {
            var v17 uint64 = 1
            v18 := []int{1, 9223372036854775807}
            var v19 int16 = 1000
            _ = v17
            _ = v18
            _ = v19
        } else {
            for fi20 := 0; fi20 < 6; fi20++ {
                for fi21 := 0; fi21 < 9; fi21++ {
                    var v22 int64 = -444
                    switch "" {
                    case "hello":
                        v22 += 0
                    case "\n\t":
                        v23 := GenS1{F0: 42}
                        var v24 float64 = (1e10 * ((-2.25 / 3.14159) + ((3.14159 * float64((1e10 + 1.5))) * -2.25)))
                        v24 = v24
                        _ = v23
                        _ = v24
                    case "foo bar":
                        v22 -= v22
                    }
                    if (0 < 640) {
                        v22 = 9223372036854775807
                    }
                    _ = fi21
                    _ = v22
                }
                _ = fi20
            }
            v25 := GenS1{F0: 0}
            _ = v25
        }
    } else {
        var v26 int = 1
        if false {
            for fi27 := 0; fi27 < 2; fi27++ {
                v28 := []int{(fi27 - fi27)}
                fmt.Println(v26)
                v29 := GenS2{F0: 1, F1: ((1 < -1) && false)}
                _ = fi27
                _ = v28
                _ = v29
            }
            var v30 uint32 = 2
            var v31 int16 = -520
            _ = v30
            _ = v31
        }
        _ = v26
    }
    v32 := GenS2{F0: 9223372036854775807, F1: (0 != -1)}
    _ = v32
    return ("a" + ("\n\t" + "foo bar"))
}

func main() {
    var v1 uint16 = 32767
    var v2 uint64 = 18446744073709551615
    var v3 uint = 374
    var v4 uint = (v3 / ((uint(v1)) | 1))
    v5 := make(map[uint32]int)
    v2 -= (v2 ^ (v2 + (v2 >> ((65535) & 63))))
    fmt.Println(v2, v1)
    v33 := []int{32768, -586}
    for fi34 := 0; fi34 < 13; fi34++ {
        if true {
            v3 = v3
            fmt.Println(fi34, v4)
        }
        _ = fi34
    }
    v1 -= (v1 / ((v1) | 1))
    v3 = v4
    var v35 uint = 4294967295
    fmt.Println(v1)
    v2 += (v2 % (((v2 << ((42) & 63))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v35
    _ = v4
    _ = v5
}
