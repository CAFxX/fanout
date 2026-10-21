package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool; F1 uint; F2 int }
func genFn0(p0 uint16, p1 int8, p2 string) uint32 {
    fmt.Println(p1)
    v7 := []int{920, -392, -9223372036854775808, 268}
    v7 = append(v7, -2147483648)
    var v8 int16 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 125
}
func genFn1(p0 int32) int32 {
    v9 := GenS1{F0: -32, F1: 65535}
    p0 += -881
    if true {
        p0 = 42
    } else {
        v10 := GenS1{F0: 0, F1: 2}
        _ = v10
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        if false {
            if true {
                if (false && (fi11 < (fi11 ^ 9223372036854775807))) {
                    var v12 uint = 18446744073709551615
                    p0 += (p0 | -2147483648)
                    v13 := GenS2{F0: 32767}
                    _ = v12
                    _ = v13
                }
                v14 := GenS2{F0: 42}
                for fi15 := 0; fi15 < 12; fi15++ {
                    if true {
                        var v16 uint8 = 0
                        v17 := []int{(fi11 * fi11), (fi15 / (((fi15 >> ((815) & 63))) | 1)), (fi11 - fi11)}
                        var v18 uint32 = 1
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                    p0 -= 0
                    fmt.Println(fi11)
                    _ = fi15
                }
                _ = v14
            }
            switch "hello" {
            case "\n\t":
                var v19 string = ("foo bar" + "")
                _ = v19
            case "":
                v20 := []int{fi11, -787, 670, -747}
                _ = v20
            }
            for fi21 := 0; fi21 < 6; fi21++ {
                p0 += p0
                fmt.Println(p0, fi21)
                v22 := []int{9223372036854775807}
                _ = fi21
                _ = v22
            }
        }
        p0 = -83
        v23 := []int{fi11, fi11, (fi11 >> ((296) & 63))}
        _ = fi11
        _ = v23
    }
    fmt.Println(p0)
    _ = p0
    _ = v9
    return (p0 % (((p0 + -727)) | 1))
}

func main() {
    var v1 bool = ((((-251 <= -91) && true) || true) && true)
    v2 := []uint32{2147483648, 274, 328, 615}
    v3 := make(map[string]uint32)
    v4 := []string{}
    var v5 int8 = 127
    var v6 int32 = 963
    fmt.Println(v6, v5, v1)
    if false {
        var v24 float64 = 1.5
        if v1 {
            _ = len(v4)
            var v25 uint16 = 42
            _ = v25
        } else {
            var v26 uint = 42
            _ = v26
        }
        var v27 uint64 = 0
        _ = v24
        _ = v27
    } else {
        v28 := GenS3{F0: (125 <= 1), F1: 42, F2: 0}
        var v29 bool = (1 >= 817)
        _ = v28
        _ = v29
    }
    v6 = (v6 % (((v6 << ((42) & 63))) | 1))
    v30 := GenS3{F0: true, F1: 138, F2: 453}
    for fi31 := 0; fi31 < 15; fi31++ {
        delete(v3, "")
        _ = fi31
    }
    for fi32 := 0; fi32 < 6; fi32++ {
        for fi33 := 0; fi33 < 2; fi33++ {
            if v1 {
                var v34 string = "a"
                _ = v34
            }
            switch fi32 {
            case 42:
                var v35 int8 = (v5 / ((v5) | 1))
                _ = v35
            case 2147483647:
                if false {
                    if v1 {
                        if true {
                            v6 = v6
                        }
                    } else {
                        v1 = true
                    }
                    v6 = -1
                }
            case -906:
                var v36 float32 = ((((1e10 * -2.25) + 1.5) + 3.14159) - 3.14159)
                _ = v36
            }
            if v1 {
                var v37 float32 = 0.1
                _ = v37
            } else {
                v5 += 100
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v38 := GenS1{F0: v5, F1: 4294967295}
                _ = v38
            }
            _ = fi33
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v39 := GenS3{F0: v1, F1: 104, F2: fi32}
        _ = fi32
        _ = v39
    }
    if true {
        v40 := GenS3{F0: true, F1: 673, F2: -9223372036854775808}
        v5 = v5
        _ = v40
    }
    var v41 int = 1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
