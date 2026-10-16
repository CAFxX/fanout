package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 int32; F1 uint64 }

func main() {
    var v1 uint16 = 255
    v2 := []int{780, 520, 32768}
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]string)
    var v5 int8 = -1
    v6 := make(map[uint32]int)
    switch "a" {
    case "\n\t":
        delete(v4, 2147483648)
    default:
        if true {
            if len(v2) > 0 {
                v2[0] = -2147483648
            }
            if true {
                if true {
                    fmt.Println(v5, v5)
                }
            }
            var v7 bool = true
            _ = v7
        }
    }
    _ = len(v4)
    v8 := v6[689]
    if true {
        for fi9 := 0; fi9 < 5; fi9++ {
            for fi10 := 0; fi10 < 19; fi10++ {
                if (-32768 >= 1000) {
                    delete(v6, 0)
                    v8 = v8
                    v8 -= 1
                } else {
                    if false {
                        v11 := v3[0]
                        _ = v11
                    }
                    switch "" {
                    case "\n\t":
                        var v12 uint8 = 255
                        _ = v12
                    }
                    v13 := GenS1{F0: -1}
                    _ = v13
                }
                var v14 uint16 = v1
                _ = fi10
                _ = v14
            }
            switch 2147483647 {
            case 644:
                if len(v2) > 0 {
                    v2[0] = fi9
                }
            case -229:
                v15 := GenS2{F0: uint8(v1), F1: 675}
                _ = v15
            case 127:
                v5 = v5
            }
            _ = fi9
        }
    }
    var v16 uint = 255
    var v17 int32 = 2147483647
    fmt.Println(v17)
    delete(v4, 154)
    v18 := GenS3{F0: 0, F1: 127}
    var v19 uint16 = v1
    var v20 float64 = (((float64(float32(1.5)) + -2.25) * -2.25) - 1.5)
    var v21 string = ""
    v5 = -128
    if ((v16 >> (((v16 << ((v16) & 63))) & 63)) == v16) {
        if false {
            var v22 int8 = (v5 | 30)
            v23 := GenS4{F0: v17, F1: 2}
            _ = v22
            _ = v23
        }
        var v24 int64 = 1
        for fi25 := 0; fi25 < 18; fi25++ {
            for fi26 := 0; fi26 < 18; fi26++ {
                var v27 int8 = 42
                _ = fi26
                _ = v27
            }
            _ = fi25
        }
        _ = v24
    }
    var v28 string = v21
    var v29 int64 = 1
    v30 := GenS4{F0: (-(v17)), F1: 18446744073709551615}
    fmt.Println(v1, v19, v19)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
