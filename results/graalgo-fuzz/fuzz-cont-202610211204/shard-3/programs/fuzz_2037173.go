package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 uint16; F1 uint16; F2 bool }
type GenS4 struct { F0 string; F1 uint64 }

func main() {
    v1 := []uint32{}
    v2 := []int{}
    var v3 bool = (-2147483648 > 0)
    v4 := []string{}
    v5 := make(map[int]int)
    var v6 uint16 = 1
    v7 := GenS1{F0: 0}
    for fi8 := 0; fi8 < 19; fi8++ {
        if (-1 != 0) {
            for fi9 := 0; fi9 < 2; fi9++ {
                var v10 uint32 = 2
                _ = len(v4)
                if (!v3) {
                    _ = len(v4)
                    v11 := GenS2{F0: -128, F1: 1}
                    if (1 != 149) {
                        v12 := GenS1{F0: 2}
                        _ = v12
                    }
                    _ = v11
                }
                _ = fi9
                _ = v10
            }
            for fi13 := 0; fi13 < 12; fi13++ {
                v6 = (v6 / (((v6 >> ((65535) & 63))) | 1))
                var v14 bool = ((!v3) || true)
                _ = fi13
                _ = v14
            }
            var v15 int64 = 1
            _ = v15
        } else {
            v16 := GenS1{F0: 42}
            _ = v16
        }
        for fi17 := 0; fi17 < 12; fi17++ {
            var v18 int64 = -9223372036854775808
            v19 := GenS3{F0: 65535, F1: v6, F2: (!v3)}
            if false {
                if v3 {
                    v20 := GenS1{F0: 446}
                    fmt.Println(v6, fi17)
                    switch fi17 {
                    case -555:
                        var v21 string = ((("a" + "a") + "hello") + ((("hello" + "\n\t") + "") + ""))
                        _ = v21
                    }
                    _ = v20
                }
                var v22 int16 = 963
                v23 := GenS3{F0: 2, F1: v6, F2: ((v6 * (v6 >> ((503) & 63))) > 65535)}
                _ = v22
                _ = v23
            } else {
                delete(v5, int(v6))
            }
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = fi8
    }
    fmt.Println(v6)
    if (-1 == 0) {
        var v24 int16 = -32768
        var v25 float32 = 3.14159
        if (!(1 == 108)) {
            fmt.Println(v24)
            var v26 uint32 = 2147483648
            v26 = v26
            _ = v26
        } else {
            if true {
                _ = len(v5)
            }
            v5[0] = -26
            var v27 string = ""
            _ = v27
        }
        _ = v24
        _ = v25
    } else {
        var v28 int = -1
        _ = v28
    }
    fmt.Println(v3, v6)
    if len(v4) > 0 {
        v4[0] = (("\n\t" + "\n\t") + "\n\t")
    }
    if v3 {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    v29 := GenS4{F0: "\n\t", F1: 18446744073709551615}
    var v30 float64 = -2.25
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
