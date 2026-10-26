package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 string }
func genFn0(p0 int16, p1 uint8) uint {
    if true {
        p0 -= (p0 / ((p0) | 1))
        if false {
            var v5 int = 0
            v6 := []int{v5}
            _ = v5
            _ = v6
        }
        for fi7 := 0; fi7 < 4; fi7++ {
            switch (("\n\t" + "foo bar") + "") {
            case "":
                v8 := []int{}
                _ = v8
            case "\n\t":
                if true {
                    if false {
                        p1 += p1
                    }
                } else {
                    v9 := []int{(fi7 | 9223372036854775807), (fi7 % ((fi7) | 1)), fi7}
                    _ = v9
                }
            case "foo bar":
                p0 += 696
            }
            if (0.0 != ((float32(((3.14159 * 0.1) + 0.1)) + 1.5) + 1e10)) {
                if false {
                    v10 := GenS1{F0: -1, F1: 42, F2: ("\n\t" + "a")}
                    v11 := GenS1{F0: 0, F1: 2, F2: "hello"}
                    _ = v10
                    _ = v11
                } else {
                    for fi12 := 0; fi12 < 12; fi12++ {
                        p1 = p1
                        _ = fi12
                    }
                    if (true && (-1 > 43)) {
                        fmt.Println(p0, p1)
                    } else {
                        v13 := []int{(fi7 / ((int(p0)) | 1)), (fi7 / ((fi7) | 1))}
                        _ = v13
                    }
                    v14 := []int{}
                    _ = v14
                }
            }
            var v15 int64 = -1
            _ = fi7
            _ = v15
        }
    } else {
        var v16 int = 0
        _ = v16
    }
    fmt.Println(p0, p0)
    v17 := []int{4294967295, 9223372036854775807, 32767}
    _ = p0
    _ = p1
    _ = v17
    return 194
}

func main() {
    var v1 float64 = (1e10 / 1e10)
    v2 := make(map[int]uint32)
    v3 := []uint32{4294967295, 237, 2147483648}
    var v4 uint8 = 227
    for fi18 := 0; fi18 < 3; fi18++ {
        var v19 string = ("" + ("\n\t" + "a"))
        if true {
            _ = len(v2)
        }
        _ = fi18
        _ = v19
    }
    fmt.Println(v1, v1)
    var v20 bool = (!(1 != 881))
    v4 += 127
    v21 := v2[528]
    v22 := v2[42]
    for fi23 := 0; fi23 < 4; fi23++ {
        var v24 float64 = v1
        fmt.Println(v1)
        v1 = float64((1e10 / (3.14159 - 0.0)))
        _ = fi23
        _ = v24
    }
    var v25 uint = 514
    fmt.Println(v4)
    for fi26 := 0; fi26 < 19; fi26++ {
        var v27 int64 = 9223372036854775807
        var v28 uint8 = 1
        var v29 uint8 = (v28 << ((genFn0(-32768, (v4 ^ v4))) & 63))
        _ = fi26
        _ = v27
        _ = v28
        _ = v29
    }
    if (!v20) {
        delete(v2, 1)
        v25 = 156
    }
    var v30 int16 = 42
    for fi31 := 0; fi31 < 8; fi31++ {
        if v20 {
            fmt.Println(v1, v25, v21)
            var v32 uint16 = 42
            var v33 int32 = -1
            _ = v32
            _ = v33
        }
        _ = fi31
    }
    var v34 int64 = -147
    fmt.Println(v21, v4)
    v35 := GenS1{F0: 0, F1: uint32(v25), F2: "\n\t"}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v25
    _ = v3
    _ = v30
    _ = v34
    _ = v35
    _ = v4
}
