package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 int8; F1 int8; F2 int16 }
type GenS3 struct { F0 int }
func genFn0(p0 uint64, p1 uint32, p2 bool) int64 {
    var v7 uint32 = p1
    fmt.Println(p2, p0, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 91
}
func genFn1(p0 uint64, p1 int) uint32 {
    var v8 int = (^(p1))
    if (127 > 0) {
        switch "" {
        case "\n\t":
            if false {
                v9 := []int{-299, p1, (p1 / ((p1) | 1))}
                p0 -= 84
                _ = v9
            } else {
                var v10 int16 = 1000
                _ = v10
            }
        case "foo bar":
            if (!false) {
                switch ("\n\t" + "\n\t") {
                case "hello":
                    if false {
                        v11 := []int{}
                        var v12 int32 = 504
                        p1 += p1
                        _ = v11
                        _ = v12
                    } else {
                        if (42 == 4294967295) {
                            fmt.Println(v8)
                        } else {
                            p0 -= 1
                        }
                    }
                default:
                    for fi13 := 0; fi13 < 17; fi13++ {
                        p0 = 1
                        var v14 int8 = 127
                        var v15 uint64 = 939
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                }
            } else {
                v16 := GenS1{F0: 226, F1: 0}
                _ = v16
            }
        default:
            v17 := []int{-1}
            _ = v17
        }
    }
    for fi18 := 0; fi18 < 16; fi18++ {
        var v19 int32 = 1
        v8 = int(p0)
        for fi20 := 0; fi20 < 10; fi20++ {
            p1 += (^(fi18))
            _ = fi20
        }
        _ = fi18
        _ = v19
    }
    _ = p0
    _ = p1
    _ = v8
    return 65535
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = 100
    var v3 bool = true
    var v4 float32 = 0.0
    var v5 uint64 = 2
    v6 := make(map[int]int)
    v21 := GenS1{F0: (^(v5)), F1: v2}
    v3 = false
    v22 := v1[2147483648]
    var v23 uint64 = v5
    v24 := GenS2{F0: -1, F1: 1, F2: 661}
    fmt.Println(v2)
    fmt.Println(v22, v4, v23)
    v25 := GenS3{F0: -943}
    for fi26 := 0; fi26 < 12; fi26++ {
        for fi27 := 0; fi27 < 18; fi27++ {
            v5 -= v5
            _ = len(v1)
            _ = fi27
        }
        if false {
            for fi28 := 0; fi28 < 18; fi28++ {
                v29 := []int{(fi28 | fi26), fi28, fi28}
                var v30 string = ""
                if true {
                    fmt.Println(v4)
                    v31 := GenS1{F0: (v5 + v23), F1: 122}
                    _ = v31
                }
                _ = fi28
                _ = v29
                _ = v30
            }
            var v32 int16 = 1
            var v33 int32 = 467
            _ = v32
            _ = v33
        }
        _ = fi26
    }
    var v34 uint = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
