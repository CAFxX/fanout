package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 string; F1 int; F2 int8 }
type GenS3 struct { F0 uint32; F1 uint16; F2 uint8 }
func genFn0(p0 int32, p1 uint64) bool {
    var v7 bool = false
    if false {
        switch -9223372036854775808 {
        case 32767:
            var v8 uint8 = 160
            _ = v8
        }
        v9 := []int{int(p1)}
        p1 += (p1 % (((^(p1))) | 1))
        _ = v9
    } else {
        if (!v7) {
            for fi10 := 0; fi10 < 1; fi10++ {
                fmt.Println(p1)
                v11 := []int{}
                _ = fi10
                _ = v11
            }
            var v12 string = "hello"
            _ = v12
        }
    }
    v13 := []int{-816, 32767, 127}
    v13 = append(v13, 879)
    _ = p0
    _ = p1
    _ = v13
    _ = v7
    return v7
}

func main() {
    v1 := []string{"a"}
    var v2 bool = true
    var v3 float64 = 3.14159
    var v4 uint16 = 32767
    var v5 int64 = -9223372036854775808
    v6 := make(map[uint32]bool)
    for fi14 := 0; fi14 < 20; fi14++ {
        switch (-(fi14)) {
        case 32768:
            var v15 uint8 = 128
            _ = v15
        }
        v16 := GenS1{F0: (2 >= 255), F1: v2}
        _ = fi14
        _ = v16
    }
    v1 = append(v1, "a")
    v4 = 989
    for fi17 := 0; fi17 < 11; fi17++ {
        v1 = append(v1, "\n\t")
        switch 32767 {
        case -2147483648:
            switch (-(fi17)) {
            case 32767:
                v2 = false
            case 4294967295:
                v3 += 1e10
            case -9223372036854775808:
                for fi18 := 0; fi18 < 11; fi18++ {
                    if v2 {
                        _ = len(v6)
                    } else {
                        switch (fi18 >> ((4294967295) & 63)) {
                        case 9223372036854775807:
                            var v19 uint16 = 271
                            _ = v19
                        case 4294967295:
                            v1 = append(v1, ("\n\t" + ""))
                        case 127:
                            v5 = (-(v5))
                        default:
                            var v20 uint = 1
                            var v21 int8 = 0
                            v1 = append(v1, "hello")
                            _ = v20
                            _ = v21
                        }
                        var v22 string = "\n\t"
                        if v2 {
                            var v23 bool = (1 >= 100)
                            fmt.Println(v3)
                            var v24 uint16 = v4
                            _ = v23
                            _ = v24
                        }
                        v25 := GenS2{F0: v22, F1: (^(fi17)), F2: -128}
                        fmt.Println(v5, v3)
                        _ = v22
                        _ = v25
                    }
                    _ = len(v6)
                    v4 -= v4
                    _ = fi18
                }
            default:
                v5 = v5
            }
        case 42:
            fmt.Println(fi17, v3, fi17)
        }
        _ = fi17
    }
    if ((v5 - (-(v5))) != (-(v5))) {
        for fi26 := 0; fi26 < 4; fi26++ {
            switch (fi26 % ((128) | 1)) {
            case 0:
                fmt.Println(v5, fi26)
            default:
                v4 = 65535
            }
            _ = fi26
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v2 = (42 != 222)
    }
    v5 = v5
    var v27 int = 638
    fmt.Println(v2, v3)
    v28 := GenS1{F0: true, F1: (true && v2)}
    v29 := GenS3{F0: 827, F1: 601, F2: 1}
    v2 = (!false)
    var v30 uint = 18446744073709551615
    v30 -= v30
    if v2 {
        var v31 float64 = v3
        fmt.Println(v31, v31)
        _ = v31
    }
    var v32 bool = true
    var v33 uint16 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
