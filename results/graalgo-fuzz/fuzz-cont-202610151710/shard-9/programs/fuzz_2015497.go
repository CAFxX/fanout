package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 uint32; F1 int64; F2 string }
func genFn0(p0 int64) string {
    var v7 uint = 908
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 int64 = (p0 & 0)
        v9 -= (v9 + int64(v7))
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = v7
    return ""
}
func genFn1(p0 string, p1 uint32, p2 uint8) int {
    if true {
        p1 += (p1 | (p1 * (p1 >> ((42) & 63))))
        var v10 uint = 18446744073709551615
        for fi11 := 0; fi11 < 12; fi11++ {
            switch fi11 {
            case 2147483647:
                switch p0 {
                case "hello":
                    v12 := []int{(^(fi11)), (fi11 << ((746) & 63)), (fi11 % ((-841) | 1))}
                    _ = v12
                }
            default:
                v13 := GenS1{F0: p2, F1: p0}
                _ = v13
            }
            _ = fi11
        }
        _ = v10
    } else {
        if false {
            if false {
                var v14 int64 = -742
                _ = v14
            }
            fmt.Println(p2)
        }
        p2 = (p2 | (p2 & 128))
    }
    p2 = (p2 % (((p2 + p2)) | 1))
    var v15 int16 = 0
    for fi16 := 0; fi16 < 5; fi16++ {
        if true {
            v17 := GenS1{F0: (^(p2)), F1: p0}
            p2 += (p2 * 133)
            _ = v17
        } else {
            v18 := []int{}
            fmt.Println(p1, p1)
            _ = v18
        }
        for fi19 := 0; fi19 < 11; fi19++ {
            if (false || (!(!false))) {
                p1 = (p1 - p1)
                v20 := GenS1{F0: 2, F1: p0}
                if (-1 >= 42) {
                    fmt.Println(fi16, p0, p2)
                } else {
                    switch p0 {
                    case "":
                        v21 := []int{0, (-(fi16)), fi16}
                        _ = v21
                    case "a":
                        switch (-(fi16)) {
                        case 42:
                            var v22 int = 729
                            _ = v22
                        }
                    case "\n\t":
                        if (!false) {
                            v23 := []int{(fi19 & int(v15)), (fi16 >> ((2) & 63)), fi19, (fi19 % (((fi16 << ((9) & 63))) | 1))}
                            v24 := GenS1{F0: (^(p2)), F1: p0}
                            fmt.Println(p1, fi16)
                            _ = v23
                            _ = v24
                        } else {
                            var v25 float64 = (1.5 - (0.0 + 0.0))
                            var v26 int32 = -2147483648
                            v27 := []int{}
                            _ = v25
                            _ = v26
                            _ = v27
                        }
                    }
                    switch 0 {
                    case 0:
                        v28 := GenS2{F0: (v15 / ((v15) | 1)), F1: (-(fi16))}
                        _ = v28
                    case 32767:
                        switch (p0 + p0) {
                        case "foo bar":
                            v29 := GenS1{F0: p2, F1: ((p0 + p0) + (p0 + p0))}
                            _ = v29
                        case "":
                            var v30 uint64 = 18446744073709551615
                            _ = v30
                        default:
                            var v31 string = (p0 + "a")
                            _ = v31
                        }
                    }
                    v32 := []int{(^(fi19))}
                    _ = v32
                }
                p2 -= p2
                p0 = p0
                p1 -= (^(p1))
                _ = v20
            }
            var v33 int64 = -1
            var v34 float64 = (3.14159 + 1.5)
            _ = fi19
            _ = v33
            _ = v34
        }
        fmt.Println(v15, p1)
        _ = fi16
    }
    v35 := []int{32768, 4294967295, 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v35
    return -586
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 uint8 = 42
    var v3 int64 = -618
    var v4 uint8 = (v2 - 2)
    var v5 uint = 255
    v6 := []uint32{1}
    v2 = v4
    var v36 int = 535
    var v37 string = ("foo bar" + "a")
    fmt.Println(v4)
    fmt.Println(v36, v2, v1)
    v38 := GenS2{F0: 42, F1: (v36 >> ((v5) & 63))}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v37 = "hello"
    var v39 int16 = 0
    v40 := GenS3{F0: 1, F1: (v1 + (v3 / ((v3) | 1))), F2: ("\n\t" + v37)}
    var v41 int8 = 1
    var v42 uint8 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
