package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 int8; F1 bool; F2 int32 }
type GenS4 struct { F0 uint32 }
func genFn0() uint8 {
    v4 := GenS1{F0: -1}
    var v5 int32 = 1
    var v6 bool = false
    _ = v4
    _ = v5
    _ = v6
    return 128
}

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = ((627 < 0) && false)
    v3 := make(map[string]string)
    v7 := GenS1{F0: -23}
    v8 := GenS2{F0: -870, F1: v2}
    v2 = v2
    if (2 > 344) {
        v9 := []int{-607}
        fmt.Println(v2)
        v10 := GenS2{F0: 1, F1: (722 <= 0)}
        _ = v10
        _ = v9
    }
    var v11 float64 = (0.1 - float64(1e10))
    var v12 int64 = 0
    v12 += v12
    v2 = ((v12 >> ((2) & 63)) != 0)
    v12 = (^(v12))
    if v2 {
        for fi13 := 0; fi13 < 14; fi13++ {
            v14 := GenS3{F0: -128, F1: (1.5 < float32(v11)), F2: -733}
            _ = fi13
            _ = v14
        }
        var v15 uint = 255
        for fi16 := 0; fi16 < 3; fi16++ {
            for fi17 := 0; fi17 < 13; fi17++ {
                if (1 != genFn0()) {
                    v15 = v15
                    var v18 int16 = -1
                    v19 := []int{fi16}
                    _ = v18
                    _ = v19
                } else {
                    if (true && v2) {
                        if true {
                            v2 = ((v11 == (v11 - 1.5)) && (((v11 * (v11 + v11)) >= v11) && (-128 != -128)))
                            var v20 bool = false
                            _ = v20
                        }
                    }
                    var v21 int16 = 42
                    _ = v21
                }
                if (v2 && ((v15 <= 450) && (4294967295 == (v15 + 42)))) {
                    v15 += (^(v15))
                    v2 = (0 >= 1)
                    switch "\n\t" {
                    case "a":
                        var v22 int = (fi16 << ((v15) & 63))
                        _ = v22
                    }
                } else {
                    switch fi17 {
                    case -748:
                        var v23 uint32 = 234
                        _ = v23
                    default:
                        v24 := []int{0, fi16, int(v15), (fi16 / ((2147483647) | 1))}
                        _ = v24
                    }
                }
                _ = fi17
            }
            fmt.Println(v15)
            switch (fi16 ^ (^(fi16))) {
            case 2147483647:
                var v25 int32 = 1
                _ = v25
            default:
                delete(v3, (("a" + "\n\t") + "a"))
            }
            _ = fi16
        }
        _ = v15
    } else {
        v26 := GenS3{F0: 0, F1: (v2 && false), F2: 42}
        _ = v26
    }
    for fi27 := 0; fi27 < 14; fi27++ {
        fmt.Println(fi27)
        v2 = (!true)
        _ = fi27
    }
    if v2 {
        v28 := []int{32767, 9223372036854775807, 647}
        _ = v28
    }
    delete(v1, 65535)
    switch ("foo bar" + "\n\t") {
    case "hello":
        v29 := []int{}
        _ = v29
    case "\n\t":
        v30 := GenS4{F0: 1}
        _ = v30
    }
    v11 -= 0.1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
