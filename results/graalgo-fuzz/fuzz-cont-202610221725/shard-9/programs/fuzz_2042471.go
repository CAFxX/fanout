package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 uint32; F1 uint8 }
func genFn0(p0 uint16) int32 {
    switch (("foo bar" + "\n\t") + ("hello" + "foo bar")) {
    case "a":
        var v4 float32 = 1e10
        _ = v4
    case "hello":
        v5 := []int{42}
        _ = v5
    }
    v6 := GenS1{F0: 9223372036854775807}
    var v7 int = -2147483648
    if false {
        var v8 int64 = -295
        v9 := GenS1{F0: v8}
        var v10 int8 = 127
        _ = v10
        _ = v8
        _ = v9
    }
    for fi11 := 0; fi11 < 12; fi11++ {
        switch ((("" + "") + ("a" + "foo bar")) + "a") {
        case "\n\t":
            v12 := []int{(^(v7)), fi11}
            _ = v12
        }
        if (100 >= 42) {
            for fi13 := 0; fi13 < 6; fi13++ {
                v14 := []int{fi13}
                _ = fi13
                _ = v14
            }
            for fi15 := 0; fi15 < 19; fi15++ {
                v16 := []int{v7, (fi11 | v7), (fi15 & fi15)}
                v7 -= fi15
                v7 += 671
                var v17 uint8 = 174
                var v18 uint16 = p0
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
            }
            v7 = v7
        } else {
            var v19 bool = (((698 <= -32768) || false) || false)
            for fi20 := 0; fi20 < 6; fi20++ {
                v7 = 2147483647
                _ = fi20
            }
            p0 -= p0
            _ = v19
        }
        v21 := []int{v7, fi11, (fi11 / ((1) | 1))}
        _ = fi11
        _ = v21
    }
    _ = p0
    _ = v6
    _ = v7
    return 2147483647
}
func genFn1(p0 int16, p1 int) uint {
    var v22 uint32 = 4294967295
    p0 -= p0
    for fi23 := 0; fi23 < 16; fi23++ {
        fmt.Println(p1, fi23)
        var v24 uint8 = 2
        v25 := []int{(fi23 - (fi23 / ((-1) | 1)))}
        _ = fi23
        _ = v24
        _ = v25
    }
    fmt.Println(p1, v22, v22)
    _ = p0
    _ = p1
    _ = v22
    return 255
}

func main() {
    v1 := []uint32{2147483647, 858, 142}
    var v2 uint64 = 759
    v3 := make(map[uint32]bool)
    v2 = (v2 >> ((1) & 63))
    fmt.Println(v2)
    v3[2147483648] = (852 >= genFn1(1, 0))
    v26 := GenS2{F0: (^(v2)), F1: -1}
    var v27 uint = 65535
    var v28 uint8 = 1
    v28 = 127
    for fi29 := 0; fi29 < 8; fi29++ {
        if true {
            for fi30 := 0; fi30 < 12; fi30++ {
                v27 = (^(v27))
                var v31 uint8 = 42
                _ = fi30
                _ = v31
            }
            for fi32 := 0; fi32 < 15; fi32++ {
                var v33 uint64 = v2
                v2 = 18446744073709551615
                _ = fi32
                _ = v33
            }
            v28 = v28
        } else {
            for fi34 := 0; fi34 < 8; fi34++ {
                v28 = v28
                v2 -= 140
                if false {
                    var v35 uint8 = v28
                    _ = v35
                }
                _ = fi34
            }
            if false {
                if ((112 >= int8(v2)) || (2 > 65535)) {
                    switch (fi29 - fi29) {
                    case 1:
                        v36 := GenS1{F0: -9223372036854775808}
                        _ = v36
                    case 2147483647:
                        fmt.Println(fi29)
                    case -9223372036854775808:
                        v37 := GenS3{F0: 1, F1: 255}
                        _ = v37
                    }
                    fmt.Println(v2, v2, fi29)
                }
                if false {
                    var v38 uint32 = 196
                    _ = v38
                }
            } else {
                v39 := GenS3{F0: 42, F1: v28}
                var v40 int8 = -7
                _ = v39
                _ = v40
            }
            var v41 uint64 = v2
            _ = v41
        }
        _ = fi29
    }
    switch "a" {
    case "hello":
        fmt.Println(v2, v2)
    }
    var v42 int64 = 42
    v28 = (v28 & (v28 + 255))
    switch 811 {
    case 934:
        delete(v3, 129)
    case -9223372036854775808:
        v42 += v42
    case -223:
        if (v2 <= (v2 * 0)) {
            v43 := GenS3{F0: 65535, F1: (v28 & 255)}
            fmt.Println(v28)
            var v44 int8 = -128
            _ = v43
            _ = v44
        } else {
            _ = len(v1)
        }
    }
    for fi45 := 0; fi45 < 15; fi45++ {
        v2 -= 787
        for fi46 := 0; fi46 < 17; fi46++ {
            var v47 float64 = 3.14159
            v47 += float64(0.1)
            v27 = genFn1(-356, fi45)
            v27 -= (v27 | (v27 & 847))
            _ = fi46
            _ = v47
        }
        _ = fi45
    }
    for fi48 := 0; fi48 < 1; fi48++ {
        v49 := GenS3{F0: 102, F1: (v28 / ((v28) | 1))}
        v27 += 315
        _ = fi48
        _ = v49
    }
    var v50 bool = false
    v42 -= (v42 >> (((v27 * 65535)) & 63))
    var v51 int8 = -84
    var v52 uint16 = 255
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v42:", v42)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v42
    _ = v50
    _ = v51
    _ = v52
}
