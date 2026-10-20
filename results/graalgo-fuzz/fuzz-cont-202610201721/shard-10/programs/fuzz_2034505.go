package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 int16 }

func main() {
    var v1 uint16 = 255
    v2 := make(map[string]int)
    v3 := make(map[uint32]int)
    v4 := []int{4294967295}
    var v5 int16 = -431
    var v6 int64 = 9223372036854775807
    var v7 float64 = (1e10 - (float64(0.1) + 0.0))
    var v8 int32 = -532
    var v9 uint16 = v1
    var v10 int = 9223372036854775807
    for fi11 := 0; fi11 < 7; fi11++ {
        switch fi11 {
        case -92:
            if false {
                var v12 uint32 = 42
                _ = v12
            } else {
                v3[1] = -2147483648
                if (2147483647 < 2) {
                    var v13 int32 = (v8 % (((^(v8))) | 1))
                    var v14 uint16 = 823
                    var v15 string = "a"
                    _ = v13
                    _ = v14
                    _ = v15
                }
                fmt.Println(v6, v8)
            }
        case 51:
            var v16 string = ("\n\t" + "foo bar")
            _ = v16
        case -639:
            for fi17 := 0; fi17 < 4; fi17++ {
                var v18 uint64 = 18446744073709551615
                _ = fi17
                _ = v18
            }
        default:
            var v19 uint8 = uint8(v1)
            _ = v19
        }
        _ = fi11
    }
    if true {
        var v20 string = ("a" + "\n\t")
        var v21 uint32 = 42
        _ = v20
        _ = v21
    } else {
        v4 = append(v4, v10)
        for fi22 := 0; fi22 < 16; fi22++ {
            switch (-(v10)) {
            case 128:
                v23 := GenS1{F0: true, F1: int64(v8), F2: (v5 - v5)}
                _ = v23
            case 4294967295:
                for fi24 := 0; fi24 < 19; fi24++ {
                    v25 := GenS1{F0: false, F1: v6, F2: v5}
                    _ = fi24
                    _ = v25
                }
            case -2147483648:
                var v26 int8 = 8
                _ = v26
            default:
                for fi27 := 0; fi27 < 18; fi27++ {
                    fmt.Println(v6, fi27)
                    _ = fi27
                }
            }
            v28 := GenS1{F0: false, F1: (v6 + int64(v8)), F2: v5}
            _ = fi22
            _ = v28
        }
    }
    fmt.Println(v8, v10)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
