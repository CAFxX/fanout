package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 bool; F2 int }

func main() {
    var v1 uint16 = 2
    var v2 float32 = 1.5
    var v3 int32 = -1
    v4 := make(map[int]int)
    for fi5 := 0; fi5 < 18; fi5++ {
        for fi6 := 0; fi6 < 15; fi6++ {
            var v7 bool = ((fi5 | fi5) <= fi5)
            var v8 string = ("foo bar" + "a")
            var v9 float32 = -2.25
            switch (fi6 << ((18446744073709551615) & 63)) {
            case -182:
                if false {
                    v8 = "hello"
                    v10 := GenS1{F0: 907}
                    v4[-1] = (fi5 << ((42) & 63))
                    _ = v10
                }
                if v7 {
                    var v11 uint64 = 4294967295
                    if (2 == 2) {
                        fmt.Println(fi5)
                    } else {
                        if ((^(v11)) >= 42) {
                            v12 := []int{(fi5 % (((fi5 | 9223372036854775807)) | 1)), fi5, (fi6 * int(v11))}
                            v8 = ("\n\t" + v8)
                            _ = v12
                        } else {
                            v11 = (v11 / (((v11 - (v11 >> ((255) & 63)))) | 1))
                            var v13 uint32 = 4294967295
                            if ((^(fi6)) > (fi6 | fi5)) {
                                switch fi6 {
                                case 128:
                                    _ = len(v4)
                                }
                            } else {
                                if true {
                                    v4[-2147483648] = fi5
                                    v14 := GenS1{F0: (v13 | 2)}
                                    v1 = (v1 ^ 779)
                                    _ = v14
                                } else {
                                    fmt.Println(v2, fi6)
                                    fmt.Println(v8)
                                }
                                fmt.Println(v1)
                            }
                            _ = v13
                        }
                        var v15 int32 = (^(v3))
                        _ = v15
                    }
                    fmt.Println(v8)
                    _ = v11
                }
            }
            v1 += 32767
            if v7 {
                v16 := v4[fi5]
                var v17 uint32 = 597
                _ = v16
                _ = v17
            }
            v18 := GenS1{F0: 0}
            v19 := GenS2{F0: "foo bar", F1: (v7 && v7), F2: fi6}
            v1 = (v1 + v1)
            var v20 uint = 18446744073709551615
            fmt.Println(v8, fi5, v9)
            _ = fi6
            _ = v18
            _ = v19
            _ = v20
            _ = v7
            _ = v8
            _ = v9
        }
        _ = fi5
    }
    switch 4294967295 {
    case 32768:
        var v21 uint = 377
        _ = v21
    }
    var v22 uint16 = 123
    if false {
        var v23 int16 = -32768
        var v24 int8 = 100
        _ = v23
        _ = v24
    } else {
        if (65535 >= 2) {
            fmt.Println(v1, v1, v2)
        }
        v25 := []int{0, 32768}
        _ = v25
    }
    v26 := GenS1{F0: 2147483647}
    var v27 int8 = 0
    fmt.Println(v2)
    v3 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v22
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
