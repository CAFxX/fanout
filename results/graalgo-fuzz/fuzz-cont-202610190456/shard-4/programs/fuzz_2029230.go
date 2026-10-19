package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 int8 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 uint8 = 107
    var v2 int64 = 652
    v3 := []int{-776}
    var v4 float64 = 3.14159
    var v5 string = "\n\t"
    v6 := []string{"\n\t", (v5 + (v5 + v5))}
    _ = len(v6)
    v2 = -613
    if len(v6) > 0 {
        v6[0] = (v5 + "hello")
    }
    v6 = append(v6, (v5 + "a"))
    if true {
        var v7 int64 = (v2 << ((2) & 63))
        v8 := GenS1{F0: 4294967295, F1: "hello", F2: 1}
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v7
        _ = v8
    } else {
        switch -1 {
        case 32767:
            fmt.Println(v2)
        case -2147483648:
            v1 -= v1
        case 1:
            if false {
                if ((v4 < v4) || false) {
                    if (3.14159 <= 1.5) {
                        var v9 uint8 = 42
                        var v10 int64 = (v2 % ((9223372036854775807) | 1))
                        _ = v10
                        _ = v9
                    }
                } else {
                    fmt.Println(v2, v5, v1)
                    if false {
                        v3 = append(v3, 128)
                        var v11 uint8 = v1
                        _ = v11
                    } else {
                        v5 = v5
                        v2 -= v2
                        if (!false) {
                            if ((float64((3.14159 + 3.14159)) - v4) != float64(0.1)) {
                                v12 := GenS1{F0: 42, F1: "hello", F2: 42}
                                v13 := GenS2{F0: -1, F1: -671}
                                _ = v12
                                _ = v13
                            }
                            var v14 string = (v5 + (v5 + ("" + (("a" + (v5 + "a")) + ("hello" + v5)))))
                            _ = v14
                        }
                    }
                    for fi15 := 0; fi15 < 17; fi15++ {
                        switch (fi15 + fi15) {
                        case 0:
                            var v16 float32 = 0.0
                            _ = v16
                        }
                        var v17 int8 = -100
                        var v18 int8 = (v17 * (v17 >> ((4294967295) & 63)))
                        _ = fi15
                        _ = v17
                        _ = v18
                    }
                }
                v5 = (("hello" + "") + (v5 + v5))
                if ("" < (((("a" + v5) + v5) + "") + v5)) {
                    var v19 uint = 0
                    _ = v19
                } else {
                    for fi20 := 0; fi20 < 17; fi20++ {
                        var v21 int = fi20
                        var v22 bool = (-1 == int8(v21))
                        v23 := GenS3{F0: 160}
                        _ = fi20
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                    var v24 bool = false
                    var v25 int16 = -1
                    _ = v24
                    _ = v25
                }
            }
        }
        var v26 uint32 = 82
        _ = v26
    }
    var v27 uint16 = 0
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v28 uint64 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
