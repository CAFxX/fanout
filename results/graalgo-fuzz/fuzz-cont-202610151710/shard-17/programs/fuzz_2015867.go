package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int8 }
type GenS2 struct { F0 int8; F1 uint16; F2 uint }
type GenS3 struct { F0 bool }
func genFn0(p0 uint32, p1 uint32, p2 int16) int32 {
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := []int{}
        if len(v8) > 0 {
            v8[0] = fi7
        }
        _ = fi7
        _ = v8
    }
    var v9 int64 = -1
    switch -2147483648 {
    case 128:
        fmt.Println(p1)
    default:
        var v10 string = "foo bar"
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return -1
}
func genFn1(p0 uint64) int {
    v11 := GenS1{F0: 4294967295, F1: 42, F2: 2}
    v12 := []int{259, 0, 366, 0}
    var v13 float32 = float32(3.14159)
    p0 = (p0 << ((4294967295) & 63))
    switch "" {
    case "foo bar":
        p0 -= (p0 % (((p0 & (p0 / (((p0 * p0)) | 1)))) | 1))
    case "a":
        var v14 int16 = -719
        _ = v14
    default:
        p0 = (p0 | p0)
    }
    _ = p0
    _ = v11
    _ = v12
    _ = v13
    return 32768
}

func main() {
    var v1 uint = 65535
    v2 := make(map[int]bool)
    v3 := make(map[int]int)
    v4 := make(map[int]bool)
    var v5 uint16 = 32767
    v6 := make(map[int]uint32)
    if (-923 > 0) {
        v15 := GenS1{F0: -2147483648, F1: 54, F2: 32}
        v16 := GenS2{F0: 100, F1: 1, F2: (v1 * (^(v1)))}
        _ = v15
        _ = v16
    }
    var v17 float64 = ((1.5 + (0.0 - ((-2.25 * ((3.14159 * (1.5 - 3.14159)) - 1.5)) * (0.1 + 1.5)))) + 0.1)
    for fi18 := 0; fi18 < 18; fi18++ {
        delete(v3, fi18)
        var v19 int16 = 32767
        _ = fi18
        _ = v19
    }
    v5 -= v5
    for fi20 := 0; fi20 < 8; fi20++ {
        if true {
            for fi21 := 0; fi21 < 16; fi21++ {
                v22 := GenS1{F0: (fi20 - fi20), F1: int8(v1), F2: -1}
                _ = fi21
                _ = v22
            }
            var v23 uint = 42
            _ = v23
        }
        _ = fi20
    }
    for fi24 := 0; fi24 < 8; fi24++ {
        v2[fi24] = true
        v17 += ((1e10 + v17) * (v17 * ((v17 - v17) / (v17 * (0.0 / (v17 / 1.0))))))
        v25 := GenS3{F0: false}
        _ = fi24
        _ = v25
    }
    v26 := GenS3{F0: (!(false || ((-495 < 32767) || true)))}
    for fi27 := 0; fi27 < 7; fi27++ {
        v1 += (v1 << (((^(v1))) & 63))
        var v28 int8 = 63
        if true {
            v17 = (0.0 - v17)
            for fi29 := 0; fi29 < 13; fi29++ {
                if true {
                    var v30 int64 = -41
                    v31 := []int{-2147483648}
                    _ = v30
                    _ = v31
                }
                v32 := GenS2{F0: 50, F1: (v5 + v5), F2: 18446744073709551615}
                _ = fi29
                _ = v32
            }
            if true {
                for fi33 := 0; fi33 < 7; fi33++ {
                    v34 := GenS1{F0: fi33, F1: (v28 >> ((v1) & 63)), F2: (v28 + v28)}
                    fmt.Println(v17, fi33, fi27)
                    var v35 uint = v1
                    _ = fi33
                    _ = v34
                    _ = v35
                }
                var v36 int16 = -39
                _ = v36
            } else {
                v37 := []int{fi27, (-(fi27)), (fi27 >> ((v1) & 63))}
                var v38 uint32 = 453
                v17 = v17
                _ = v37
                _ = v38
            }
        }
        _ = fi27
        _ = v28
    }
    v1 = uint(v5)
    var v39 uint32 = 2147483647
    v17 = ((1.5 * v17) - (1.5 / float64((((3.14159 + ((1e10 * 0.0) / float32(1.5))) - 0.0) - -2.25))))
    var v40 float32 = 1.5
    v40 = v40
    if (101 != 0) {
        var v41 int8 = 100
        fmt.Println(v17, v39, v1)
        var v42 int16 = 32767
        _ = v41
        _ = v42
    }
    v43 := []int{32768, -9223372036854775808}
    v44 := GenS1{F0: -663, F1: 0, F2: 0}
    v2[370] = false
    fmt.Println(v40, v17)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v2
    _ = v26
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
