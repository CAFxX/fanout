package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 int64; F1 int32; F2 uint16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint8; F1 uint8; F2 bool }
func genFn0(p0 uint8, p1 uint32) int64 {
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 string = "hello"
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 12; fi9++ {
        fmt.Println(p1, p1)
        _ = fi9
    }
    p0 = p0
    v10 := GenS1{F0: (p1 >> ((1) & 63))}
    v11 := GenS2{F0: 262, F1: -1}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 9223372036854775807
}
func genFn1() int32 {
    v12 := GenS3{F0: genFn0(255, 2), F1: 207, F2: 0}
    var v13 string = ("\n\t" + "")
    _ = v12
    _ = v13
    return 42
}

func main() {
    var v1 string = "a"
    var v2 int = -93
    var v3 uint32 = 4294967295
    var v4 int64 = 904
    var v5 uint16 = 42
    var v6 int = v2
    v2 -= -7
    v14 := []int{}
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    fmt.Println(v5, v1)
    v15 := GenS4{F0: 127}
    for fi16 := 0; fi16 < 20; fi16++ {
        var v17 bool = (1 <= 0)
        switch ((v1 + v1) + "a") {
        case "foo bar":
            if (-311 >= -32768) {
                fmt.Println(v5)
                v18 := GenS2{F0: (v5 * (v5 - v5)), F1: (v2 << ((2) & 63))}
                _ = v18
            }
        }
        for fi19 := 0; fi19 < 18; fi19++ {
            v20 := GenS5{F0: 42, F1: 156, F2: false}
            switch -9223372036854775808 {
            case 128:
                if len(v14) > 0 {
                    v14[0] = fi16
                }
            }
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
            _ = fi19
            _ = v20
        }
        _ = fi16
        _ = v17
    }
    var v21 string = ""
    for fi22 := 0; fi22 < 18; fi22++ {
        v2 = (-(v2))
        v14 = append(v14, (fi22 - int(v5)))
        for fi23 := 0; fi23 < 5; fi23++ {
            fmt.Println(fi23)
            fmt.Println(v6, fi22)
            _ = fi23
        }
        _ = fi22
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
