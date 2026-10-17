package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }

func main() {
    var v1 int16 = 0
    v2 := make(map[int]bool)
    v3 := make(map[int]int)
    delete(v2, 32768)
    v1 = (-(v1))
    fmt.Println(v1)
    v1 = v1
    var v4 int32 = 1
    fmt.Println(v1, v1)
    v1 = v1
    v4 += v4
    v4 = v4
    switch "a" {
    case "hello":
        if false {
            fmt.Println(v1, v4)
            for fi5 := 0; fi5 < 6; fi5++ {
                if (true && (!false)) {
                    fmt.Println(v1, v1)
                } else {
                    var v6 float32 = 1e10
                    for fi7 := 0; fi7 < 15; fi7++ {
                        switch (^(fi5)) {
                        case 934:
                            if (true && (!(765 >= 0))) {
                                _ = len(v2)
                            } else {
                                fmt.Println(v6, v4)
                                var v8 int8 = 121
                                v9 := []int{fi7, fi5, fi7, fi5}
                                v10 := GenS1{F0: 4294967295, F1: 4294967295}
                                var v11 float32 = -2.25
                                _ = v10
                                _ = v11
                                _ = v8
                                _ = v9
                            }
                        case 920:
                            var v12 uint8 = 255
                            fmt.Println(v1)
                            _ = v12
                        case 226:
                            _ = len(v2)
                        }
                        _ = fi7
                    }
                    v6 += ((1e10 * float32((float64((1.5 + float32((0.0 / 1.5)))) / 1.0))) + 0.0)
                    _ = v6
                }
                fmt.Println(fi5)
                _ = fi5
            }
            var v13 int32 = v4
            _ = v13
        }
    case "foo bar":
        v14 := GenS1{F0: 65535, F1: 2}
        _ = v14
    default:
        v1 = -1
    }
    fmt.Println(v1, v4)
    fmt.Println(v4, v1)
    var v15 float64 = 0.1
    var v16 int16 = (v1 + int16(v4))
    for fi17 := 0; fi17 < 15; fi17++ {
        for fi18 := 0; fi18 < 12; fi18++ {
            if true {
                v19 := v2[fi17]
                v20 := []int{}
                var v21 bool = true
                v16 = v16
                var v22 bool = (true && (0 < 843))
                _ = v19
                _ = v20
                _ = v21
                _ = v22
            }
            var v23 uint32 = 0
            v24 := GenS1{F0: 558, F1: 2}
            _ = fi18
            _ = v23
            _ = v24
        }
        _ = fi17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
