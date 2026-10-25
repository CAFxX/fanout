package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64; F1 uint64 }
func genFn0(p0 int16, p1 int, p2 uint) uint64 {
    var v7 int16 = (^(p0))
    switch p1 {
    case -2147483648:
        v8 := []int{}
        _ = v8
    case -382:
        p0 += p0
    case -105:
        p0 = p0
    }
    p1 -= (p1 + -941)
    var v9 uint16 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v9
    return 1
}

func main() {
    v1 := []string{"foo bar"}
    v2 := make(map[int]bool)
    v3 := make(map[int]string)
    var v4 int32 = -1
    var v5 bool = true
    var v6 int64 = 42
    v10 := GenS1{F0: int16(v6), F1: 127}
    var v11 float64 = float64(1e10)
    switch 9223372036854775807 {
    case 127:
        switch -122 {
        case 32767:
            v12 := GenS1{F0: 32767, F1: -9223372036854775808}
            _ = v12
        case 4294967295:
            var v13 float64 = v11
            _ = v13
        default:
            v5 = (true && (v5 || v5))
        }
    default:
        var v14 uint8 = 64
        _ = v14
    }
    v6 += (v6 ^ (-(v6)))
    for fi15 := 0; fi15 < 4; fi15++ {
        for fi16 := 0; fi16 < 1; fi16++ {
            var v17 int16 = 1000
            _ = fi16
            _ = v17
        }
        var v18 string = ("\n\t" + "a")
        switch fi15 {
        case 2147483647:
            switch fi15 {
            case -589:
                if (255 <= 985) {
                    v11 = v11
                    var v19 uint = 65535
                    _ = v19
                }
            case 1:
                for fi20 := 0; fi20 < 20; fi20++ {
                    switch (fi20 + (fi15 % ((fi20) | 1))) {
                    case 32768:
                        var v21 int64 = v6
                        _ = v21
                    case 9223372036854775807:
                        v22 := GenS2{F0: (v6 << ((209) & 63))}
                        _ = v22
                    }
                    v4 += v4
                    _ = fi20
                }
            default:
                v6 = v6
            }
        case 585:
            v23 := GenS1{F0: 1, F1: fi15}
            _ = v23
        case 561:
            var v24 int8 = 76
            _ = v24
        default:
            fmt.Println(v6, v6)
        }
        _ = fi15
        _ = v18
    }
    if (v6 <= v6) {
        for fi25 := 0; fi25 < 14; fi25++ {
            v5 = (1 <= 155)
            _ = fi25
        }
    }
    v26 := GenS1{F0: 32767, F1: -551}
    var v27 int64 = v6
    var v28 int64 = v6
    if v5 {
        if (-32768 != 0) {
            fmt.Println(v28, v28, v27)
        } else {
            var v29 float64 = v11
            v6 = v28
            fmt.Println(v6)
            _ = v29
        }
        v30 := GenS3{F0: 0}
        if ("\n\t" <= "\n\t") {
            if (((32767 > 1) || (470 > 554)) && (true || v5)) {
                if ((v5 && v5) && (v5 && v5)) {
                    v28 = (v27 | (v6 / ((1) | 1)))
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v31 := v3[4294967295]
                    v2[163] = v5
                    _ = v31
                }
                v6 -= 51
                var v32 bool = (v6 < (v27 + (v28 & (v28 & 42))))
                _ = v32
            } else {
                var v33 uint64 = genFn0(1, 32768, 42)
                var v34 uint = 0
                _ = v33
                _ = v34
            }
            fmt.Println(v5, v6, v4)
        } else {
            for fi35 := 0; fi35 < 1; fi35++ {
                for fi36 := 0; fi36 < 10; fi36++ {
                    switch 2147483647 {
                    case -9223372036854775808:
                        v37 := GenS4{F0: 42, F1: genFn0(-32768, (fi35 ^ 0), 4294967295)}
                        _ = v37
                    case 4294967295:
                        if v5 {
                            var v38 float32 = 3.14159
                            var v39 int32 = v4
                            _ = v38
                            _ = v39
                        }
                    }
                    _ = fi36
                }
                _ = fi35
            }
            var v40 int = int(v4)
            _ = v40
        }
        _ = v30
    }
    v41 := GenS2{F0: int64(v4)}
    delete(v3, -2147483648)
    var v42 int = 100
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
