package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 int, p1 uint64, p2 uint) uint16 {
    fmt.Println(p1)
    var v6 uint8 = 0
    if true {
        p1 -= p1
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 548
}

func main() {
    var v1 uint64 = 429
    v2 := make(map[int]string)
    v3 := make(map[int]bool)
    v4 := make(map[string]uint32)
    v5 := make(map[int]bool)
    v1 = (v1 * 4294967295)
    v7 := GenS1{F0: 1}
    v8 := []int{32767}
    fmt.Println(v1)
    if (v1 < 893) {
        var v9 int8 = 27
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = v9
    }
    var v10 float32 = ((3.14159 * 0.0) + 0.1)
    for fi11 := 0; fi11 < 8; fi11++ {
        if (!true) {
            if true {
                fmt.Println(v1, v1, v1)
                fmt.Println(fi11, v10, v10)
                var v12 string = (((("\n\t" + ("foo bar" + "")) + "hello") + ("a" + "hello")) + ("foo bar" + "foo bar"))
                _ = v12
            } else {
                if true {
                    fmt.Println(v10)
                    if (!(1 > 255)) {
                        fmt.Println(fi11)
                        if false {
                            for fi13 := 0; fi13 < 12; fi13++ {
                                fmt.Println(fi13, fi11)
                                var v14 int16 = -100
                                _ = fi13
                                _ = v14
                            }
                            var v15 uint32 = 2147483647
                            _ = v15
                        } else {
                            for fi16 := 0; fi16 < 5; fi16++ {
                                v17 := GenS1{F0: 2}
                                v18 := GenS1{F0: 127}
                                v10 -= v10
                                _ = fi16
                                _ = v17
                                _ = v18
                            }
                            v10 = 0.1
                        }
                        for fi19 := 0; fi19 < 19; fi19++ {
                            fmt.Println(v1, fi19)
                            var v20 uint64 = (v1 >> ((0) & 63))
                            _ = fi19
                            _ = v20
                        }
                    }
                }
            }
        } else {
            if false {
                var v21 int16 = 42
                _ = v21
            } else {
                v10 -= v10
                var v22 int32 = -2147483648
                var v23 uint8 = 127
                _ = v22
                _ = v23
            }
        }
        v10 -= -2.25
        v10 += float32(float64(1e10))
        _ = fi11
    }
    var v24 int8 = 42
    fmt.Println(v10, v24, v24)
    var v25 uint = 4294967295
    v5[4294967295] = (float64(v10) == -2.25)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
