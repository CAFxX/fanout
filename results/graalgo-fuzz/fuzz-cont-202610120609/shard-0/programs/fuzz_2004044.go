package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 string; F1 string; F2 int32 }
func genFn0(p0 uint8) string {
    v6 := GenS1{F0: 0}
    v7 := []int{-1, 127}
    if true {
        if (1000 <= 1000) {
            p0 -= 127
            v8 := GenS1{F0: -395}
            _ = v8
        } else {
            p0 -= (p0 - (p0 | (p0 * (p0 * (p0 << ((114) & 63))))))
            v7 = append(v7, 0)
            switch 42 {
            case 2147483647:
                var v9 string = ""
                _ = v9
            case 9223372036854775807:
                if len(v7) > 0 {
                    v7[0] = 9223372036854775807
                }
            case 753:
                var v10 float64 = 1e10
                _ = v10
            }
        }
    } else {
        var v11 float32 = (1.5 / float32((-2.25 / 3.14159)))
        _ = v11
    }
    v12 := GenS1{F0: 32767}
    _ = p0
    _ = v12
    _ = v6
    _ = v7
    return "\n\t"
}
func genFn1(p0 int8) uint8 {
    var v13 uint8 = 21
    fmt.Println(p0)
    if (42 == 246) {
        switch 42 {
        case -2147483648:
            v14 := []int{32767, 614}
            _ = v14
        default:
            if true {
                var v15 float64 = 0.0
                v16 := GenS2{F0: 567}
                _ = v15
                _ = v16
            }
        }
        v13 -= v13
        p0 = p0
    }
    var v17 bool = (-1 <= -1)
    _ = p0
    _ = v13
    _ = v17
    return 228
}

func main() {
    var v1 int = 778
    var v2 uint = 1
    var v3 int32 = -2147483648
    var v4 uint32 = 2147483647
    var v5 uint = v2
    switch v1 {
    case -1:
        v18 := GenS2{F0: (-(v1))}
        _ = v18
    case 42:
        var v19 float64 = (0.1 * -2.25)
        _ = v19
    }
    if (v4 != v4) {
        v20 := []int{-1, v1, v1}
        var v21 uint64 = 963
        switch (v1 * (v1 << ((v2) & 63))) {
        case 4294967295:
            var v22 float64 = 1e10
            _ = v22
        case 0:
            for fi23 := 0; fi23 < 10; fi23++ {
                for fi24 := 0; fi24 < 14; fi24++ {
                    v25 := GenS1{F0: v1}
                    var v26 uint32 = 477
                    v2 = (v2 + (v2 & (v5 >> ((0) & 63))))
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                _ = fi23
            }
        case -1:
            fmt.Println(v21)
        default:
            for fi27 := 0; fi27 < 3; fi27++ {
                var v28 int32 = (^(v3))
                var v29 uint8 = 95
                _ = fi27
                _ = v28
                _ = v29
            }
        }
        _ = v20
        _ = v21
    }
    var v30 int8 = 0
    v4 -= v4
    var v31 uint64 = 2
    for fi32 := 0; fi32 < 7; fi32++ {
        v33 := []int{4294967295, fi32}
        _ = fi32
        _ = v33
    }
    v34 := GenS2{F0: v1}
    var v35 float32 = 1.5
    var v36 uint8 = 255
    fmt.Println(v35)
    fmt.Println(v35, v4, v5)
    var v37 bool = false
    fmt.Println(v31, v37)
    switch v1 {
    case -391:
        if v37 {
            v38 := GenS1{F0: 42}
            _ = v38
        } else {
            v39 := GenS3{F0: 32768, F1: v31}
            if v37 {
                v40 := []int{int(v30), v1}
                _ = v40
            } else {
                v5 += (^(v5))
                v4 += (^(v4))
                if v37 {
                    if true {
                        var v41 uint32 = (^(v4))
                        v1 = -2147483648
                        _ = v41
                    } else {
                        v30 = (v30 >> ((723) & 63))
                    }
                    v4 -= 42
                } else {
                    v37 = ("a" < "")
                }
            }
            v42 := GenS3{F0: -2147483648, F1: 1}
            _ = v39
            _ = v42
        }
    case 128:
        var v43 int32 = v3
        _ = v43
    default:
        for fi44 := 0; fi44 < 18; fi44++ {
            if v37 {
                for fi45 := 0; fi45 < 2; fi45++ {
                    var v46 float32 = (0.1 - (v35 / float32(((3.14159 + 0.1) + (0.1 / (-2.25 - ((1.5 + float64(float32(0.0))) / 1.0)))))))
                    v47 := []int{fi45, (fi44 % ((9223372036854775807) | 1))}
                    v4 -= v4
                    _ = fi45
                    _ = v46
                    _ = v47
                }
                if (v2 <= v5) {
                    fmt.Println(v30, v4)
                    v48 := []int{9223372036854775807, int(v2)}
                    var v49 int32 = 858
                    _ = v48
                    _ = v49
                }
            }
            _ = fi44
        }
    }
    v50 := GenS4{F0: ("\n\t" + "hello"), F1: "\n\t", F2: 103}
    for fi51 := 0; fi51 < 17; fi51++ {
        v5 = v2
        _ = fi51
    }
    if ("hello" >= ("\n\t" + "\n\t")) {
        v36 = v36
    }
    v52 := GenS3{F0: v1, F1: uint64(v4)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v50
    _ = v52
}
