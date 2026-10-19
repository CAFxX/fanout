package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 int = 32767
    v2 := make(map[string]bool)
    v3 := make(map[int]string)
    switch ("\n\t" + "hello") {
    case "hello":
        v1 = -395
    case "foo bar":
        _ = len(v3)
    default:
        for fi4 := 0; fi4 < 10; fi4++ {
            if true {
                v1 = (fi4 + -433)
                if (65535 <= 0) {
                    var v5 int = fi4
                    var v6 float64 = 0.0
                    _ = v5
                    _ = v6
                } else {
                    v7 := GenS1{F0: 0}
                    v8 := []int{(fi4 | 606)}
                    _ = v7
                    _ = v8
                }
            }
            _ = fi4
        }
    }
    v9 := []int{(v1 % (((v1 ^ (v1 << ((1) & 63)))) | 1)), (v1 | v1)}
    for fi10 := 0; fi10 < 19; fi10++ {
        if false {
            if true {
                for fi11 := 0; fi11 < 20; fi11++ {
                    v1 -= 88
                    _ = fi11
                }
                if false {
                    v12 := GenS1{F0: -103}
                    _ = v12
                } else {
                    var v13 uint32 = 2
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    _ = v13
                }
                v1 = -530
            }
            v14 := GenS1{F0: 103}
            if (0 > 42) {
                var v15 int16 = 42
                _ = v15
            } else {
                var v16 uint8 = 72
                _ = v16
            }
            _ = v14
        } else {
            for fi17 := 0; fi17 < 3; fi17++ {
                if (-32768 != -32768) {
                    var v18 uint16 = 2
                    v1 += v1
                    v1 += fi10
                    fmt.Println(fi10)
                    _ = v18
                }
                var v19 int64 = 1
                _ = fi17
                _ = v19
            }
        }
        if true {
            var v20 bool = (0 > 1)
            _ = v20
        }
        var v21 int = -389
        _ = fi10
        _ = v21
    }
    if len(v9) > 0 {
        v9[0] = (v1 & 9223372036854775807)
    }
    var v22 int32 = 42
    delete(v2, "\n\t")
    var v23 int32 = (^(v22))
    var v24 uint8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v9
}
