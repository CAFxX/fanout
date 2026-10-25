package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 string; F1 uint16; F2 uint16 }
type GenS3 struct { F0 uint8; F1 uint16; F2 uint16 }
func genFn0(p0 string, p1 uint32) string {
    v4 := GenS1{F0: "foo bar", F1: 32767}
    var v5 uint32 = (p1 % ((4294967295) | 1))
    p0 = "\n\t"
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return p0
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]string)
    var v3 float64 = (0.0 * (0.0 * (1e10 * (-2.25 - -2.25))))
    v6 := v1[65535]
    fmt.Println(v6, v6)
    v7 := []int{}
    v3 += v3
    if v6 {
        fmt.Println(v6)
    } else {
        if v6 {
            v8 := GenS2{F0: "a", F1: 2, F2: 207}
            _ = v8
        } else {
            v7 = append(v7, -285)
            if v6 {
                v3 -= v3
                var v9 uint8 = 111
                v3 = (v3 * v3)
                _ = v9
            }
            if (4294967295 <= 950) {
                if false {
                    if v6 {
                        for fi10 := 0; fi10 < 18; fi10++ {
                            v11 := GenS1{F0: "a", F1: fi10}
                            var v12 float64 = v3
                            for fi13 := 0; fi13 < 1; fi13++ {
                                if true {
                                    var v14 int8 = 1
                                    var v15 int64 = 9223372036854775807
                                    if (0 <= -32768) {
                                        v16 := GenS1{F0: "a", F1: fi10}
                                        v12 = (v12 - 1e10)
                                        _ = v16
                                    }
                                    _ = v14
                                    _ = v15
                                }
                                v3 = 1.5
                                _ = fi13
                            }
                            _ = fi10
                            _ = v11
                            _ = v12
                        }
                    }
                    v17 := GenS2{F0: "", F1: 65535, F2: 2}
                    var v18 float64 = v3
                    _ = v17
                    _ = v18
                }
                var v19 float64 = v3
                fmt.Println(v6)
                _ = v19
            }
        }
        if ((!(2 != 255)) || v6) {
            v6 = (-2.25 <= (-2.25 - (0.0 / (3.14159 * 1e10))))
        }
    }
    v3 = v3
    v6 = false
    fmt.Println(v6, v3)
    v3 += v3
    if (!true) {
        var v20 uint32 = 400
        v21 := GenS2{F0: ("foo bar" + ("a" + "\n\t")), F1: 65535, F2: 253}
        v22 := GenS3{F0: 160, F1: 0, F2: 65535}
        _ = v20
        _ = v21
        _ = v22
    }
    var v23 int8 = -111
    var v24 int = 1
    for fi25 := 0; fi25 < 3; fi25++ {
        fmt.Println(v23, v6, fi25)
        _ = fi25
    }
    var v26 uint = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v26
    _ = v3
    _ = v6
    _ = v7
}
