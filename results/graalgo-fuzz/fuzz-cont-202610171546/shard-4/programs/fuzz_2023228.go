package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 uint8; F1 uint8; F2 uint32 }
func genFn0(p0 string, p1 int, p2 uint8) int64 {
    var v6 int = (p1 & (p1 - 500))
    switch 77 {
    case 42:
        if false {
            v6 += 399
            fmt.Println(p2)
            var v7 uint32 = 561
            _ = v7
        }
    case 2147483647:
        var v8 int = (v6 % (((p1 >> ((uint(p2)) & 63))) | 1))
        _ = v8
    default:
        p0 = "\n\t"
    }
    v9 := []int{p1}
    v10 := GenS1{F0: p1, F1: p1, F2: ("foo bar" + "hello")}
    for fi11 := 0; fi11 < 20; fi11++ {
        v12 := GenS1{F0: -2147483648, F1: (^(v6)), F2: p0}
        switch -90 {
        case -2147483648:
            fmt.Println(p2, fi11)
        }
        var v13 bool = true
        _ = fi11
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v6
    _ = v9
    return 0
}

func main() {
    var v1 uint64 = 42
    var v2 string = ""
    var v3 uint8 = 193
    var v4 uint64 = 104
    var v5 uint8 = (^(v3))
    var v14 float64 = 1e10
    for fi15 := 0; fi15 < 6; fi15++ {
        var v16 int8 = 100
        v17 := []int{}
        var v18 float64 = 3.14159
        _ = fi15
        _ = v16
        _ = v17
        _ = v18
    }
    fmt.Println(v3, v4)
    var v19 float32 = 0.1
    var v20 uint32 = 58
    v21 := []int{4294967295, -248, -241, -722}
    fmt.Println(v20)
    v1 = (v4 - (v4 * 42))
    v22 := GenS2{F0: 202}
    var v23 int = 128
    v21 = append(v21, (v23 & (v23 * (v23 * (v23 / (((v23 - (v23 >> ((uint(v3)) & 63)))) | 1))))))
    fmt.Println(v3, v14, v23)
    var v24 string = ""
    switch "a" {
    case "a":
        for fi25 := 0; fi25 < 7; fi25++ {
            fmt.Println(v3, v5)
            v20 = v20
            _ = fi25
        }
    case "hello":
        v26 := GenS1{F0: (v23 | -1), F1: 9223372036854775807, F2: "\n\t"}
        _ = v26
    default:
        v4 += (^(v1))
    }
    if true {
        var v27 int64 = 9223372036854775807
        if (!false) {
            switch v2 {
            case "hello":
                v20 = (^(v20))
            case "foo bar":
                if true {
                    v28 := GenS2{F0: 47}
                    _ = v28
                }
            case "":
                v29 := GenS3{F0: v3, F1: 42}
                _ = v29
            default:
                v30 := GenS1{F0: 42, F1: -777, F2: "\n\t"}
                _ = v30
            }
            var v31 string = v24
            if true {
                v31 = (v24 + v24)
                var v32 uint32 = v20
                _ = v32
            }
            _ = v31
        } else {
            if true {
                var v33 float32 = v19
                if (!true) {
                    v34 := GenS4{F0: (^(v5)), F1: v5, F2: v20}
                    if true {
                        for fi35 := 0; fi35 < 10; fi35++ {
                            fmt.Println(v2, v24, v24)
                            _ = fi35
                        }
                        var v36 int64 = v27
                        v21 = append(v21, (v23 >> ((255) & 63)))
                        _ = v36
                    }
                    _ = v34
                }
                if false {
                    v23 += v23
                    fmt.Println(v24)
                    v37 := GenS1{F0: 9223372036854775807, F1: v23, F2: v2}
                    _ = v37
                } else {
                    if len(v21) >= 2 {
                        v21 = v21[1:len(v21)]
                    }
                }
                _ = v33
            }
            var v38 float32 = (v19 / 0.1)
            _ = v38
        }
        _ = v27
    } else {
        if (true || true) {
            if len(v21) > 0 {
                v21[0] = -2147483648
            }
            v24 = v24
            if false {
                for fi39 := 0; fi39 < 5; fi39++ {
                    switch v2 {
                    case "\n\t":
                        var v40 float64 = (v14 / float64(0.1))
                        _ = v40
                    case "hello":
                        for fi41 := 0; fi41 < 8; fi41++ {
                            var v42 int = (fi41 + 0)
                            v43 := GenS3{F0: v3, F1: -465}
                            _ = fi41
                            _ = v42
                            _ = v43
                        }
                    case "a":
                        v1 += (v4 | v1)
                    }
                    _ = fi39
                }
            }
        } else {
            var v44 float32 = (0.1 * v19)
            fmt.Println(v4, v20)
            _ = v44
        }
        if false {
            if (true && (uint32(v4) == (v20 / ((v20) | 1)))) {
                if true {
                    var v45 uint = uint(v20)
                    v24 = "\n\t"
                    _ = v45
                }
                if ((!false) && (2147483647 != 1)) {
                    if false {
                        _ = len(v21)
                    }
                    v23 = (v23 + (v23 | int(v20)))
                }
            }
            switch (v23 + (-(v23))) {
            case 128:
                v2 = ""
            case 2147483647:
                v20 = v20
            default:
                var v46 int16 = 1
                _ = v46
            }
        }
        if false {
            fmt.Println(v1, v2, v5)
            fmt.Println(v4, v2)
            v5 = v5
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
