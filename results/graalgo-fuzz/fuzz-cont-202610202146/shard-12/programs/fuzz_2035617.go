package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int16; F1 int8 }
func genFn0() bool {
    var v7 int = -1
    v8 := GenS1{F0: (v7 % ((v7) | 1)), F1: 0}
    var v9 uint16 = 2
    _ = v7
    _ = v8
    _ = v9
    return (!(!(true || (!true))))
}

func main() {
    v1 := []int{}
    v2 := []uint32{321}
    var v3 int32 = 42
    v4 := []uint32{1, 2147483648, 1}
    v5 := make(map[uint32]int)
    v6 := make(map[int]string)
    v3 += (v3 / ((-764) | 1))
    for fi10 := 0; fi10 < 13; fi10++ {
        var v11 float64 = 0.0
        v11 -= (((v11 * (0.1 / ((v11 / (1.5 / float64(1e10))) * (v11 - ((((0.1 * float64(-2.25)) * 1e10) + 0.0) + (v11 + ((v11 * v11) * v11))))))) - v11) + v11)
        for fi12 := 0; fi12 < 20; fi12++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if (4294967295 > 42) {
                fmt.Println(v11, fi10)
                if (622 != 2147483647) {
                    v11 = (0.1 + 1e10)
                }
                var v13 string = "a"
                _ = v13
            }
            _ = fi12
        }
        _ = fi10
        _ = v11
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        fmt.Println(fi14, fi14)
        _ = fi14
    }
    var v15 int = 128
    for fi16 := 0; fi16 < 12; fi16++ {
        v3 = (v3 + (v3 >> ((65535) & 63)))
        for fi17 := 0; fi17 < 14; fi17++ {
            var v18 uint64 = 0
            v19 := GenS2{F0: 2147483648}
            var v20 string = "a"
            if false {
                if (327 >= 641) {
                    v18 += (v18 % ((v18) | 1))
                    v20 = ("" + "hello")
                    var v21 string = v20
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = v21
                } else {
                    fmt.Println(v3, v15)
                }
                v15 = 2147483647
                var v22 bool = ((v18 ^ 42) <= v18)
                if false {
                    if v22 {
                        v23 := GenS3{F0: 32767}
                        v22 = (v20 >= "hello")
                        _ = v23
                    }
                } else {
                    switch int(v3) {
                    case 2147483647:
                        fmt.Println(v20, fi16, v15)
                    }
                    var v24 uint32 = 636
                    if true {
                        v15 = fi17
                    }
                    _ = v24
                }
                v25 := GenS1{F0: (^(v15)), F1: 571}
                _ = v22
                _ = v25
            }
            v26 := GenS4{F0: -868, F1: 127}
            v27 := GenS3{F0: 32767}
            _ = fi17
            _ = v18
            _ = v19
            _ = v20
            _ = v26
            _ = v27
        }
        v28 := GenS4{F0: 1000, F1: 1}
        _ = fi16
        _ = v28
    }
    delete(v6, (v15 * (v15 / ((572) | 1))))
    v2 = append(v2, 2)
    v29 := GenS4{F0: 471, F1: 42}
    var v30 float32 = float32((1.5 / 0.1))
    var v31 uint16 = 1
    var v32 float64 = 0.1
    v31 = v31
    for fi33 := 0; fi33 < 11; fi33++ {
        v34 := GenS4{F0: 32767, F1: 127}
        v35 := GenS2{F0: 42}
        _ = fi33
        _ = v34
        _ = v35
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
