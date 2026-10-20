package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 uint; F1 uint16; F2 uint8 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]uint32)
    var v3 uint32 = 63
    var v4 uint16 = 1
    if true {
        fmt.Println(v4, v4)
    } else {
        for fi5 := 0; fi5 < 15; fi5++ {
            fmt.Println(fi5, fi5, v4)
            _ = fi5
        }
    }
    var v6 int8 = -1
    v1["hello"] = "hello"
    v6 += (v6 / ((v6) | 1))
    v3 += 65535
    if false {
        if false {
            v2[4294967295] = 393
            for fi7 := 0; fi7 < 10; fi7++ {
                for fi8 := 0; fi8 < 18; fi8++ {
                    switch ("" + ("a" + ("" + "foo bar"))) {
                    case "":
                        if false {
                            var v9 float64 = -2.25
                            fmt.Println(fi8, fi8, v6)
                            if (0.0 >= -2.25) {
                                v10 := []int{fi7, fi7, fi8, (fi7 * 42)}
                                _ = v10
                            }
                            v11 := GenS1{F0: v3, F1: v6}
                            _ = v11
                            _ = v9
                        } else {
                            v12 := GenS1{F0: 444, F1: v6}
                            var v13 uint64 = 42
                            _ = v12
                            _ = v13
                        }
                    default:
                        v14 := []int{0, fi7, (fi8 / ((int(v6)) | 1))}
                        _ = v14
                    }
                    _ = fi8
                }
                fmt.Println(fi7, v4)
                _ = fi7
            }
        }
        v15 := GenS1{F0: 1, F1: 1}
        _ = v15
    }
    if false {
        if false {
            v4 = 65535
            v16 := []int{}
            _ = v16
        } else {
            switch "" {
            case "\n\t":
                v17 := GenS2{F0: 65535, F1: 255, F2: 127}
                _ = v17
            }
            fmt.Println(v4)
        }
        v3 -= (v3 % ((65535) | 1))
        v2[-608] = v3
    }
    fmt.Println(v6, v6)
    var v18 uint16 = (v4 & (v4 * 0))
    delete(v2, int(v18))
    v19 := []int{}
    var v20 string = "a"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v6
}
