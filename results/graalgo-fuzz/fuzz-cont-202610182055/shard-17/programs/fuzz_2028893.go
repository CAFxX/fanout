package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 int16; F2 bool }
type GenS3 struct { F0 int16 }
func genFn0(p0 int32, p1 uint8, p2 uint) int32 {
    for fi6 := 0; fi6 < 13; fi6++ {
        switch fi6 {
        case 127:
            fmt.Println(p0, p2, p2)
        }
        _ = fi6
    }
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 int = fi7
        _ = fi7
        _ = v8
    }
    if ((-(p0)) <= p0) {
        if (332 < (p2 | (p2 - (p2 & p2)))) {
            v9 := []int{-325, 0}
            _ = v9
        }
    } else {
        v10 := []int{2147483647, 32767, int(p1)}
        var v11 uint8 = (p1 / ((p1) | 1))
        for fi12 := 0; fi12 < 17; fi12++ {
            fmt.Println(p1)
            p2 = 18446744073709551615
            p1 = p1
            _ = fi12
        }
        _ = v10
        _ = v11
    }
    v13 := GenS1{F0: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return p0
}

func main() {
    var v1 int8 = 127
    v2 := []int{4294967295, -9223372036854775808}
    var v3 uint = 698
    v4 := make(map[string]bool)
    var v5 int8 = v1
    var v14 uint64 = 678
    fmt.Println(v3, v14, v3)
    v15 := GenS2{F0: (v14 >> ((v3) & 63)), F1: 32767, F2: (!true)}
    for fi16 := 0; fi16 < 2; fi16++ {
        v17 := GenS2{F0: v14, F1: 42, F2: (4294967295 <= 2147483647)}
        for fi18 := 0; fi18 < 18; fi18++ {
            var v19 string = ("" + "\n\t")
            fmt.Println(v19)
            var v20 uint = (v3 + 376)
            fmt.Println(fi16, fi18, v14)
            v2 = append(v2, fi18)
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = fi16
        _ = v17
    }
    v3 = v3
    v21 := GenS3{F0: 931}
    v22 := GenS3{F0: 0}
    var v23 bool = ((false || ("\n\t" == (("foo bar" + "\n\t") + "\n\t"))) || false)
    switch "a" {
    case "\n\t":
        var v24 uint32 = 0
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
