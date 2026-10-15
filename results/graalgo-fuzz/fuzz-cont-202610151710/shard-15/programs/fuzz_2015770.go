package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 uint32 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint64, p1 uint32) uint16 {
    p0 = p0
    v7 := []int{640}
    p1 += p1
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 int64 = 42
        p1 = p1
        if len(v7) > 0 {
            v7[0] = (-(fi8))
        }
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return 42
}
func genFn1() string {
    var v10 uint32 = 2147483647
    if false {
        for fi11 := 0; fi11 < 14; fi11++ {
            for fi12 := 0; fi12 < 13; fi12++ {
                v10 = v10
                v13 := []int{(fi12 | (fi12 | 127))}
                _ = fi12
                _ = v13
            }
            if false {
                v14 := []int{140, -1, fi11}
                _ = v14
            }
            switch (fi11 + (fi11 % (((fi11 * 2147483647)) | 1))) {
            case 127:
                v15 := []int{fi11, 0, -982, (-(fi11))}
                _ = v15
            case 772:
                fmt.Println(fi11, fi11)
            case -371:
                v16 := GenS1{F0: 4294967295, F1: 0, F2: v10}
                _ = v16
            default:
                switch 1 {
                case 9223372036854775807:
                    var v17 int32 = -251
                    _ = v17
                case 91:
                    v10 = (v10 | v10)
                }
            }
            _ = fi11
        }
        v10 = (v10 & 2)
    }
    var v18 int16 = -137
    if true {
        var v19 uint32 = (v10 - v10)
        v20 := GenS2{F0: "hello", F1: -323}
        for fi21 := 0; fi21 < 20; fi21++ {
            var v22 int = fi21
            _ = fi21
            _ = v22
        }
        _ = v19
        _ = v20
    }
    _ = v10
    _ = v18
    return "foo bar"
}

func main() {
    var v1 string = "foo bar"
    var v2 uint64 = 18446744073709551615
    v3 := make(map[uint32]uint32)
    var v4 float64 = 0.1
    v5 := make(map[string]uint32)
    v6 := []int{221, 42}
    v23 := v3[4294967295]
    v1 = v1
    v2 = (v2 - v2)
    switch 32767 {
    case 128:
        for fi24 := 0; fi24 < 3; fi24++ {
            v6 = append(v6, (-(fi24)))
            if true {
                v25 := GenS3{F0: 127}
                v26 := v5[v1]
                var v27 string = ("\n\t" + "\n\t")
                _ = v25
                _ = v26
                _ = v27
            }
            fmt.Println(v23, v1)
            _ = fi24
        }
    }
    var v28 uint8 = 2
    switch 127 {
    case 0:
        var v29 uint = 4294967295
        _ = v29
    }
    v4 += (1.5 - (v4 - v4))
    fmt.Println(v28, v28)
    var v30 bool = (0 >= 255)
    var v31 uint16 = 2
    v2 += (v2 | (v2 - (v2 & 2)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
