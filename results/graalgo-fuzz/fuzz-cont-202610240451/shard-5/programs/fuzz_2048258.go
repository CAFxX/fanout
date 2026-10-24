package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 uint8, p1 int64, p2 uint8) int {
    for fi5 := 0; fi5 < 16; fi5++ {
        var v6 int64 = (-(p1))
        fmt.Println(p0, p0, fi5)
        switch (fi5 ^ (fi5 + (fi5 + 32767))) {
        case 9223372036854775807:
            for fi7 := 0; fi7 < 4; fi7++ {
                v8 := GenS1{F0: 4294967295}
                p0 += p2
                _ = fi7
                _ = v8
            }
        case 4294967295:
            var v9 float64 = 1e10
            _ = v9
        case -2147483648:
            p1 = 717
        }
        _ = fi5
        _ = v6
    }
    var v10 uint32 = 456
    v11 := []int{427, 0}
    var v12 int8 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    return 336
}
func genFn1(p0 int32, p1 string) uint64 {
    switch -280 {
    case 9223372036854775807:
        if true {
            var v13 uint32 = 0
            _ = v13
        } else {
            var v14 uint = 2
            for fi15 := 0; fi15 < 16; fi15++ {
                if (!(!true)) {
                    var v16 int32 = p0
                    _ = v16
                }
                var v17 int16 = -240
                _ = fi15
                _ = v17
            }
            fmt.Println(p0, v14)
            _ = v14
        }
    }
    var v18 uint32 = 0
    var v19 uint = uint(v18)
    _ = p0
    _ = p1
    _ = v18
    _ = v19
    return 4294967295
}

func main() {
    var v1 int16 = 42
    v2 := []uint32{2147483647, 0}
    var v3 uint32 = 2
    v4 := make(map[uint32]int)
    if len(v2) > 0 {
        v2[0] = (v3 % (((v3 % (((v3 & v3)) | 1))) | 1))
    }
    var v20 int32 = -587
    var v21 uint16 = 263
    var v22 bool = false
    var v23 int = genFn0(113, -408, 127)
    v1 += 1
    for fi24 := 0; fi24 < 6; fi24++ {
        v25 := GenS2{F0: (v21 / ((v21) | 1))}
        var v26 float32 = ((-2.25 - 0.1) + 0.0)
        _ = fi24
        _ = v25
        _ = v26
    }
    if (false || ((-(v23)) == (v23 & genFn0(203, 0, 128)))) {
        switch "hello" {
        case "":
            v2 = append(v2, v3)
        case "hello":
            _ = len(v4)
        }
    }
    fmt.Println(v20, v21)
    v21 = (v21 | (v21 + v21))
    var v27 uint16 = (v21 + 0)
    switch ("" + "") {
    case "a":
        for fi28 := 0; fi28 < 8; fi28++ {
            if ("a" < "\n\t") {
                fmt.Println(v23, v1)
                _ = len(v4)
            } else {
                var v29 int32 = 2147483647
                if v22 {
                    v4[(v3 << ((353) & 63))] = v23
                    var v30 uint64 = 699
                    for fi31 := 0; fi31 < 20; fi31++ {
                        var v32 float32 = 1e10
                        _ = fi31
                        _ = v32
                    }
                    _ = v30
                }
                var v33 float32 = ((0.1 - ((3.14159 - 3.14159) * ((3.14159 * 0.0) / 0.1))) / float32(1.5))
                _ = v29
                _ = v33
            }
            fmt.Println(v3, v27)
            _ = fi28
        }
    case "":
        var v34 int8 = 0
        _ = v34
    case "\n\t":
        var v35 bool = v22
        _ = v35
    }
    switch genFn0(42, -546, 2) {
    case -2147483648:
        fmt.Println(v21, v23, v21)
    case 42:
        var v36 int16 = (-(v1))
        _ = v36
    }
    fmt.Println(v3, v20)
    var v37 uint8 = 1
    fmt.Println(v27, v23, v23)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v27
    _ = v3
    _ = v37
    _ = v4
}
