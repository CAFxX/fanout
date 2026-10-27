package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 string }
func genFn0(p0 int8) bool {
    var v6 uint8 = 42
    for fi7 := 0; fi7 < 3; fi7++ {
        if false {
            for fi8 := 0; fi8 < 20; fi8++ {
                var v9 bool = (v6 >= v6)
                v10 := GenS1{F0: (false || (v6 != 237))}
                if ((fi7 & (fi7 >> ((65535) & 63))) < fi8) {
                    fmt.Println(fi7, fi7)
                } else {
                    v11 := GenS2{F0: 2, F1: (^(v6))}
                    if (-2.25 >= (1e10 * 0.0)) {
                        var v12 int8 = (p0 / ((0) | 1))
                        _ = v12
                    }
                    _ = v11
                }
                _ = fi8
                _ = v10
                _ = v9
            }
            switch (fi7 % ((4294967295) | 1)) {
            case 42:
                var v13 uint8 = v6
                _ = v13
            case 4294967295:
                v14 := GenS2{F0: 704, F1: (v6 + (v6 / ((v6) | 1)))}
                _ = v14
            case -1:
                switch fi7 {
                case 127:
                    var v15 uint16 = 0
                    _ = v15
                }
            default:
                v16 := GenS2{F0: 738, F1: 255}
                _ = v16
            }
            v17 := []int{}
            _ = v17
        }
        if true {
            switch "" {
            case "a":
                v18 := []int{}
                _ = v18
            case "\n\t":
                v19 := GenS1{F0: false}
                _ = v19
            default:
                v20 := GenS3{F0: 1, F1: -35}
                _ = v20
            }
            if (false && true) {
                v21 := GenS3{F0: -220, F1: 36}
                _ = v21
            } else {
                p0 = p0
            }
            switch fi7 {
            case 2147483647:
                var v22 int64 = 410
                _ = v22
            }
        } else {
            switch ("hello" + "\n\t") {
            case "foo bar":
                v6 += (v6 % ((v6) | 1))
            default:
                fmt.Println(p0)
            }
            fmt.Println(fi7, fi7, fi7)
        }
        var v23 bool = (!true)
        _ = fi7
        _ = v23
    }
    v6 = v6
    var v24 uint32 = 2147483647
    _ = p0
    _ = v24
    _ = v6
    return (true && (false || (0 == 986)))
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]string)
    v3 := []string{(("a" + "") + "foo bar"), ("\n\t" + "foo bar")}
    var v4 uint64 = 508
    var v5 int8 = -1
    delete(v1, 2147483647)
    v25 := v1[128]
    for fi26 := 0; fi26 < 19; fi26++ {
        v25 -= v25
        v27 := GenS3{F0: 0, F1: v5}
        _ = fi26
        _ = v27
    }
    fmt.Println(v4, v5, v4)
    switch (v25 * (^(v25))) {
    case 810:
        if true {
            v28 := GenS3{F0: 0, F1: 100}
            _ = v28
        }
    }
    for fi29 := 0; fi29 < 17; fi29++ {
        switch v25 {
        case 42:
            v30 := v2[248]
            _ = v30
        case 1:
            v1[(v25 * 497)] = v25
        default:
            switch "foo bar" {
            case "a":
                fmt.Println(v5, fi29)
            case "foo bar":
                fmt.Println(fi29)
            }
        }
        if true {
            for fi31 := 0; fi31 < 7; fi31++ {
                var v32 uint64 = 2
                if false {
                    switch v25 {
                    case 127:
                        v4 = (v4 >> ((42) & 63))
                    default:
                        v33 := GenS4{F0: ("foo bar" + ("\n\t" + ""))}
                        _ = v33
                    }
                } else {
                    var v34 int32 = int32(fi31)
                    if false {
                        fmt.Println(v32, fi29, v5)
                        v25 += (fi31 << ((1) & 63))
                        var v35 int = (v25 / ((v25) | 1))
                        _ = v35
                    }
                    v36 := GenS4{F0: ("" + "hello")}
                    fmt.Println(fi29, v25)
                    fmt.Println(fi31, v34, fi31)
                    _ = v34
                    _ = v36
                }
                _ = fi31
                _ = v32
            }
            if len(v3) > 0 {
                v3[0] = ("\n\t" + "foo bar")
            }
        }
        _ = fi29
    }
    v37 := GenS1{F0: (!(-128 > v5))}
    if len(v3) > 0 {
        v3[0] = "a"
    }
    var v38 string = "a"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
