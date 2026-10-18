package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 int; F2 int32 }
func genFn0(p0 uint16, p1 int, p2 bool) int64 {
    for fi5 := 0; fi5 < 12; fi5++ {
        if ((0.0 * ((-2.25 - 1e10) - float64(float32(0.1)))) == 0.1) {
            for fi6 := 0; fi6 < 4; fi6++ {
                p2 = (!(p2 && true))
                p0 += (p0 - 66)
                if (p1 < -974) {
                    v7 := []int{(p1 * (p1 + (p1 * fi5))), fi5, (^(fi6)), (p1 / (((fi6 + 32767)) | 1))}
                    if false {
                        var v8 uint64 = 2
                        v8 = (^(v8))
                        p0 += p0
                        _ = v8
                    } else {
                        v9 := GenS1{F0: -85}
                        v10 := GenS1{F0: 100}
                        fmt.Println(fi6, p0, fi6)
                        var v11 int = 38
                        v12 := GenS1{F0: 127}
                        _ = v10
                        _ = v11
                        _ = v12
                        _ = v9
                    }
                    _ = v7
                } else {
                    v13 := []int{}
                    _ = v13
                }
                v14 := []int{(fi6 / ((fi6) | 1)), (p1 ^ fi5)}
                var v15 int8 = 1
                _ = fi6
                _ = v14
                _ = v15
            }
            for fi16 := 0; fi16 < 9; fi16++ {
                switch "\n\t" {
                case "a":
                    p2 = (p1 < p1)
                case "hello":
                    fmt.Println(p0)
                default:
                    fmt.Println(p0, fi16, fi5)
                }
                _ = fi16
            }
        } else {
            switch p1 {
            case -1:
                fmt.Println(p0, p1)
            case 964:
                switch (fi5 / ((-788) | 1)) {
                case -959:
                    fmt.Println(fi5)
                case 0:
                    var v17 int16 = -1
                    _ = v17
                }
            }
            p0 = (p0 | p0)
            if (p2 && false) {
                v18 := []int{}
                _ = v18
            } else {
                v19 := GenS1{F0: 126}
                var v20 string = (("foo bar" + "hello") + "hello")
                _ = v19
                _ = v20
            }
        }
        _ = fi5
    }
    var v21 uint64 = 18446744073709551615
    var v22 bool = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    _ = v22
    return 0
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := []string{("" + "hello"), "\n\t", "\n\t", ("hello" + "foo bar")}
    var v3 float64 = (1.5 - 1.5)
    var v4 bool = true
    var v23 bool = (!(2147483647 >= -1))
    for fi24 := 0; fi24 < 13; fi24++ {
        v25 := GenS1{F0: 42}
        if len(v2) > 0 {
            v2[0] = "foo bar"
        }
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 16; fi26++ {
        var v27 uint32 = 65535
        v1 -= (-(v1))
        _ = fi26
        _ = v27
    }
    if (!v23) {
        var v28 string = ("a" + ("foo bar" + "\n\t"))
        _ = len(v2)
        switch "hello" {
        case "":
            var v29 int16 = -1
            _ = v29
        case "foo bar":
            if len(v2) > 0 {
                v2[0] = "a"
            }
        }
        _ = v28
    } else {
        for fi30 := 0; fi30 < 6; fi30++ {
            var v31 bool = (v23 || (((((237 > 0) || (((v23 && v23) && false) && v4)) && (1e10 <= v3)) && (!(42 > 42))) && false))
            _ = fi30
            _ = v31
        }
        var v32 int = 491
        _ = v32
    }
    if v4 {
        if (128 <= 1) {
            if len(v2) > 0 {
                v2[0] = ((("hello" + "a") + ("" + ("foo bar" + ("" + ("hello" + "a"))))) + ("a" + ((("\n\t" + "") + ("\n\t" + "\n\t")) + "\n\t")))
            }
        } else {
            _ = len(v2)
        }
        v3 = 0.0
    }
    v3 += ((v3 / 3.14159) + 1e10)
    switch -183 {
    case 0:
        if v4 {
            fmt.Println(v23, v23, v1)
            var v33 bool = v23
            v3 = -2.25
            _ = v33
        }
    case 42:
        v1 += (v1 & 42)
    }
    v34 := GenS2{F0: 18446744073709551615, F1: -1, F2: 42}
    for fi35 := 0; fi35 < 4; fi35++ {
        var v36 int8 = -125
        v2 = append(v2, ("\n\t" + (("" + "a") + (("" + ("foo bar" + ("hello" + "a"))) + "foo bar"))))
        v2 = append(v2, ("\n\t" + "\n\t"))
        _ = fi35
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v34
    _ = v4
}
