package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 uint32; F2 int16 }
func genFn0(p0 int) int16 {
    for fi7 := 0; fi7 < 6; fi7++ {
        if (false || false) {
            fmt.Println(p0, fi7)
        } else {
            v8 := []int{(fi7 - (-(fi7))), -2147483648, (p0 + (^(fi7)))}
            var v9 uint64 = 18446744073709551615
            _ = v8
            _ = v9
        }
        var v10 float32 = (1.5 * 1e10)
        _ = fi7
        _ = v10
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        if true {
            switch "hello" {
            case "a":
                v12 := GenS1{F0: 628, F1: 42, F2: ((fi11 >> ((42) & 63)) > (^(fi11)))}
                _ = v12
            case "hello":
                var v13 int32 = -1
                _ = v13
            default:
                for fi14 := 0; fi14 < 20; fi14++ {
                    var v15 int32 = 1
                    var v16 float32 = (0.0 - (0.0 / 3.14159))
                    fmt.Println(v15, p0, v15)
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            }
            p0 = (p0 - 9223372036854775807)
        }
        for fi17 := 0; fi17 < 13; fi17++ {
            if (true || false) {
                if true {
                    switch "" {
                    case "":
                        p0 = (p0 % ((p0) | 1))
                    }
                    if true {
                        fmt.Println(p0, fi17, fi11)
                    }
                } else {
                    fmt.Println(fi17, fi11)
                    v18 := []int{(fi17 + fi17), fi17}
                    var v19 uint = 0
                    _ = v18
                    _ = v19
                }
            }
            fmt.Println(fi11, p0, fi17)
            v20 := []int{-438, (^(fi11))}
            v21 := GenS2{F0: 18446744073709551615}
            var v22 string = "\n\t"
            var v23 int32 = -1
            _ = fi17
            _ = v20
            _ = v21
            _ = v22
            _ = v23
        }
        v24 := []int{fi11, 32768}
        _ = fi11
        _ = v24
    }
    v25 := []int{(-(p0)), (^(p0))}
    p0 += (p0 - p0)
    p0 = (p0 + -83)
    _ = p0
    _ = v25
    return 1
}

func main() {
    var v1 uint8 = 1
    v2 := make(map[int]uint32)
    var v3 int64 = 42
    var v4 int16 = 37
    var v5 int32 = 1
    v6 := []string{("a" + ((("" + ("" + ("hello" + "a"))) + "hello") + "a")), ""}
    var v26 uint64 = 2
    v3 = -1
    v27 := GenS2{F0: (v26 + v26)}
    for fi28 := 0; fi28 < 6; fi28++ {
        if true {
            var v29 uint = 255
            _ = v29
        }
        _ = fi28
    }
    var v30 bool = (57 == 4294967295)
    v5 += 0
    for fi31 := 0; fi31 < 7; fi31++ {
        var v32 int8 = 100
        var v33 uint = 65535
        _ = len(v2)
        _ = fi31
        _ = v32
        _ = v33
    }
    if v30 {
        v34 := GenS3{F0: v30, F1: uint32(v1), F2: (v4 / ((v4) | 1))}
        switch "a" {
        case "":
            fmt.Println(v4)
        }
        for fi35 := 0; fi35 < 8; fi35++ {
            for fi36 := 0; fi36 < 18; fi36++ {
                var v37 uint = 68
                _ = fi36
                _ = v37
            }
            for fi38 := 0; fi38 < 19; fi38++ {
                var v39 int = fi35
                v1 += v1
                if (v30 || (genFn0(fi35) < (v4 >> ((65535) & 63)))) {
                    v40 := v2[(-(fi35))]
                    if (-2147483648 <= int32(v4)) {
                        fmt.Println(v1, v26, v26)
                        var v41 uint32 = (v40 ^ 928)
                        v42 := v2[fi35]
                        _ = v41
                        _ = v42
                    } else {
                        _ = len(v2)
                    }
                    v2[v39] = v40
                    var v43 uint16 = 2
                    fmt.Println(v26, v43, v43)
                    _ = v40
                    _ = v43
                } else {
                    var v44 uint64 = v26
                    _ = v44
                }
                _ = fi38
                _ = v39
            }
            _ = fi35
        }
        _ = v34
    } else {
        v6 = append(v6, "")
        v2[-1] = 42
    }
    var v45 int8 = -1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v4
    _ = v45
    _ = v5
    _ = v6
}
