package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 string }
func genFn0(p0 uint16, p1 string, p2 uint64) uint {
    var v4 bool = true
    if true {
        p1 = "a"
        fmt.Println(p2, p2)
    }
    switch -920 {
    case -2147483648:
        for fi5 := 0; fi5 < 18; fi5++ {
            if v4 {
                var v6 uint64 = p2
                _ = v6
            } else {
                switch fi5 {
                case -1:
                    for fi7 := 0; fi7 < 6; fi7++ {
                        if (-1 != 0) {
                            v8 := []int{0}
                            fmt.Println(p1, fi7)
                            if v4 {
                                v8 = append(v8, (fi7 - (fi5 >> ((244) & 63))))
                            }
                            _ = v8
                        }
                        var v9 float32 = (0.0 + 1.5)
                        if (!(false && (p2 != p2))) {
                            var v10 uint32 = uint32(p0)
                            _ = v10
                        }
                        _ = fi7
                        _ = v9
                    }
                case -9223372036854775808:
                    var v11 uint32 = 322
                    _ = v11
                case 32767:
                    switch (fi5 * 32767) {
                    case 240:
                        var v12 uint16 = (p0 | p0)
                        _ = v12
                    case 127:
                        if v4 {
                            var v13 int64 = -1
                            _ = v13
                        }
                    case -32:
                        v14 := []int{fi5, (fi5 % ((fi5) | 1))}
                        _ = v14
                    }
                }
            }
            _ = fi5
        }
    default:
        var v15 string = p1
        _ = v15
    }
    var v16 uint32 = 2147483647
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v4
    return 794
}

func main() {
    v1 := make(map[uint32]string)
    var v2 float64 = (3.14159 * 3.14159)
    var v3 string = "a"
    v3 = v3
    var v17 uint32 = 4294967295
    v18 := GenS1{F0: 9223372036854775807, F1: 307, F2: "a"}
    v3 = "a"
    _ = len(v1)
    if (32768 > -848) {
        var v19 uint8 = 1
        fmt.Println(v2, v2, v17)
        for fi20 := 0; fi20 < 4; fi20++ {
            delete(v1, (v17 * 42))
            switch (v3 + ((v3 + v3) + v3)) {
            case "a":
                var v21 float32 = (1.5 - (1.5 - 3.14159))
                _ = v21
            }
            _ = fi20
        }
        _ = v19
    }
    var v22 uint8 = 128
    var v23 int32 = 1
    v3 = "hello"
    var v24 float32 = 0.1
    var v25 int64 = 9223372036854775807
    v2 = v2
    delete(v1, uint32(v22))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
