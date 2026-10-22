package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 bool }
func genFn0(p0 uint16, p1 int16) uint64 {
    switch (("" + "") + "foo bar") {
    case "foo bar":
        p0 += (p0 % ((42) | 1))
    case "":
        var v6 int32 = 42
        _ = v6
    }
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := []int{fi7, 0, fi7}
        switch (("foo bar" + "\n\t") + ("a" + ("hello" + ("\n\t" + "\n\t")))) {
        case "foo bar":
            for fi9 := 0; fi9 < 1; fi9++ {
                switch (^(fi7)) {
                case 8:
                    var v10 uint = 255
                    _ = v10
                case 0:
                    _ = len(v8)
                case 2147483647:
                    var v11 int16 = p1
                    fmt.Println(fi9, v11)
                    _ = v11
                }
                v8 = append(v8, (fi9 & fi7))
                switch (fi7 + (fi9 ^ (^(fi9)))) {
                case 127:
                    fmt.Println(p1, fi7)
                case 4294967295:
                    if (752 > 603) {
                        _ = len(v8)
                        var v12 int = fi9
                        var v13 string = "foo bar"
                        _ = v12
                        _ = v13
                    }
                case -320:
                    var v14 uint = 380
                    _ = v14
                default:
                    fmt.Println(fi9, p0)
                }
                _ = fi9
            }
        case "":
            p1 -= p1
        }
        p1 -= (p1 & (p1 >> ((0) & 63)))
        _ = fi7
        _ = v8
    }
    fmt.Println(p1, p1)
    for fi15 := 0; fi15 < 2; fi15++ {
        for fi16 := 0; fi16 < 13; fi16++ {
            p0 = (^(p0))
            var v17 uint16 = (p0 | 42)
            if true {
                var v18 bool = true
                var v19 float64 = 1e10
                var v20 int16 = (p1 + 1)
                var v21 uint = 3
                _ = v18
                _ = v19
                _ = v20
                _ = v21
            } else {
                fmt.Println(v17, fi16)
                fmt.Println(fi15, p1, p1)
            }
            p0 += 32767
            _ = fi16
            _ = v17
        }
        switch "" {
        case "a":
            var v22 string = "\n\t"
            _ = v22
        case "":
            p0 -= p0
        case "hello":
            for fi23 := 0; fi23 < 9; fi23++ {
                p1 -= (p1 & p1)
                _ = fi23
            }
        }
        _ = fi15
    }
    _ = p0
    _ = p1
    return 171
}
func genFn1(p0 uint32, p1 string, p2 int32) uint64 {
    p0 -= (^(p0))
    p0 -= (p0 + p0)
    var v24 bool = true
    if false {
        fmt.Println(p0, p0)
        if v24 {
            v25 := []int{}
            _ = v25
        }
        var v26 string = ("a" + ("a" + p1))
        _ = v26
    }
    v24 = false
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    return uint64(p0)
}

func main() {
    v1 := []uint32{2, 0, 530, 9}
    var v2 uint32 = 2147483647
    v3 := make(map[string]uint32)
    v4 := make(map[int]string)
    var v5 int64 = -275
    switch 32767 {
    case -2147483648:
        if ((false || (((127 != 0) || true) && (1 < 0))) && false) {
            v5 += v5
            v5 = v5
        } else {
            v27 := GenS1{F0: 0}
            v28 := GenS2{F0: 18446744073709551615, F1: 0}
            v5 += 0
            _ = v27
            _ = v28
        }
    }
    var v29 bool = (!((true || (false || true)) && ((3.14159 == (((3.14159 + float32(1.5)) - 1.5) - 0.0)) || true)))
    v29 = v29
    v30 := GenS1{F0: 1}
    fmt.Println(v29, v5)
    v5 = 244
    fmt.Println(v29)
    v29 = ((v29 && false) && (true || (42 <= 2)))
    switch "" {
    case "":
        if (!v29) {
            v29 = (1 <= 1)
        } else {
            v2 = 283
            v5 = 9223372036854775807
            v2 = v2
        }
    default:
        v2 = (v2 & (v2 ^ v2))
    }
    v31 := GenS3{F0: ((1.5 == 1e10) && v29)}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
