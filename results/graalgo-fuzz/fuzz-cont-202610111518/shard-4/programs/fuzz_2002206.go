package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 uint }
type GenS2 struct { F0 int; F1 string; F2 uint64 }
type GenS3 struct { F0 int16; F1 int64; F2 bool }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint32; F1 string; F2 uint16 }
func genFn0(p0 int32, p1 uint8, p2 int8) uint64 {
    var v6 uint16 = 1
    v7 := GenS1{F0: 32768, F1: 100, F2: 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 uint64 = 223
    v2 := make(map[int]int)
    var v3 string = "a"
    var v4 int = -1
    var v5 string = "foo bar"
    if false {
        v8 := []int{v4}
        v9 := GenS2{F0: v4, F1: v5, F2: 0}
        var v10 uint8 = 1
        _ = v10
        _ = v8
        _ = v9
    }
    var v11 int8 = 1
    fmt.Println(v11, v3, v4)
    v12 := GenS3{F0: 1, F1: -186, F2: true}
    var v13 int = 32768
    v5 = v5
    v14 := v2[(v4 + 1)]
    var v15 bool = (v5 == v3)
    if ((v1 % ((v1) | 1)) == v1) {
        if v15 {
            v16 := GenS1{F0: (v14 * v13), F1: v11, F2: 4294967295}
            if ((v11 - 127) <= v11) {
                if false {
                    v11 += 82
                    for fi17 := 0; fi17 < 5; fi17++ {
                        v18 := GenS2{F0: v4, F1: ("\n\t" + v5), F2: genFn0(42, 1, 42)}
                        _ = fi17
                        _ = v18
                    }
                } else {
                    v19 := v2[v13]
                    if (!v15) {
                        v15 = false
                        var v20 uint8 = 127
                        v1 = v1
                        _ = v20
                    }
                    v21 := GenS4{F0: int16(v13)}
                    _ = v19
                    _ = v21
                }
                v15 = ((v1 << ((0) & 63)) > (v1 << ((42) & 63)))
            }
            v22 := GenS2{F0: v14, F1: ("\n\t" + v3), F2: v1}
            _ = v16
            _ = v22
        } else {
            for fi23 := 0; fi23 < 4; fi23++ {
                for fi24 := 0; fi24 < 2; fi24++ {
                    v25 := []int{(^(v4)), 625, (v13 ^ (v13 * fi24))}
                    _ = fi24
                    _ = v25
                }
                for fi26 := 0; fi26 < 8; fi26++ {
                    v5 = "\n\t"
                    _ = fi26
                }
                for fi27 := 0; fi27 < 1; fi27++ {
                    v28 := []int{(fi27 & (fi23 ^ -96)), fi23, (fi23 | v13)}
                    _ = fi27
                    _ = v28
                }
                _ = fi23
            }
            for fi29 := 0; fi29 < 13; fi29++ {
                var v30 float64 = -2.25
                _ = fi29
                _ = v30
            }
            var v31 float32 = float32(-2.25)
            _ = v31
        }
        if v15 {
            v14 += (v13 - (v4 << ((2) & 63)))
            switch v13 {
            case -9223372036854775808:
                v14 = 128
            }
        }
    } else {
        if (!(-1 == 1)) {
            if v15 {
                for fi32 := 0; fi32 < 18; fi32++ {
                    var v33 string = (v5 + v5)
                    v34 := GenS5{F0: 82, F1: (v33 + "foo bar"), F2: 42}
                    _ = fi32
                    _ = v33
                    _ = v34
                }
                switch v5 {
                case "\n\t":
                    v5 = ("hello" + v3)
                }
            } else {
                v35 := GenS5{F0: 2, F1: v3, F2: 255}
                var v36 int16 = -1
                _ = v35
                _ = v36
            }
            var v37 int8 = v11
            fmt.Println(v5)
            _ = v37
        } else {
            fmt.Println(v4, v15, v14)
        }
        v38 := []int{}
        _ = v38
    }
    v39 := GenS1{F0: 0, F1: v11, F2: 892}
    v13 = (v13 ^ v4)
    var v40 int16 = 32767
    v3 = (v5 + v5)
    v41 := v2[32767]
    var v42 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
}
