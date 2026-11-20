package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 uint64 }
func genFn0(p0 uint) uint32 {
    var v4 string = ("hello" + "\n\t")
    var v5 int16 = -551
    _ = p0
    _ = v4
    _ = v5
    return 964
}
func genFn1(p0 int8, p1 bool, p2 bool) int {
    var v6 float32 = ((-2.25 - (1e10 * float32(1.5))) / 1.0)
    var v7 int = 980
    switch "\n\t" {
    case "\n\t":
        if true {
            v6 -= v6
        } else {
            if p1 {
                fmt.Println(v7)
            } else {
                v8 := GenS1{F0: 1000}
                if (!(p2 && (true && true))) {
                    var v9 int = v7
                    switch v9 {
                    case 32767:
                        var v10 int = (-(v7))
                        _ = v10
                    case 913:
                        var v11 bool = (1 >= 4294967295)
                        _ = v11
                    case 4294967295:
                        switch "a" {
                        case "a":
                            v12 := GenS1{F0: 1000}
                            _ = v12
                        case "hello":
                            v7 -= -9223372036854775808
                        default:
                            var v13 uint64 = 588
                            _ = v13
                        }
                    default:
                        v14 := []int{}
                        _ = v14
                    }
                    v15 := []int{-1}
                    _ = v15
                    _ = v9
                }
                v7 -= -585
                _ = v8
            }
            switch v7 {
            case 4294967295:
                if (p2 || false) {
                    var v16 int8 = p0
                    v17 := GenS1{F0: 32767}
                    v16 -= 100
                    _ = v16
                    _ = v17
                }
            case -975:
                v18 := []int{v7, (v7 ^ (v7 & (v7 & 1))), v7, -706}
                _ = v18
            }
        }
    default:
        v19 := GenS2{F0: 4294967295, F1: "foo bar"}
        _ = v19
    }
    p2 = p1
    var v20 int16 = 237
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v6
    _ = v7
    return -466
}

func main() {
    var v1 bool = (1 <= 653)
    var v2 uint32 = 0
    v3 := []string{("" + ("a" + "hello")), "hello", "a"}
    v21 := GenS1{F0: 287}
    if ((0 < 0) || (1 > -42)) {
        fmt.Println(v1, v1)
        v1 = (!v1)
        _ = len(v3)
    } else {
        v1 = false
        v2 = 65535
        if (4294967295 >= 4294967295) {
            v22 := GenS1{F0: 1}
            v1 = (!v1)
            _ = v22
        } else {
            for fi23 := 0; fi23 < 2; fi23++ {
                switch "" {
                case "\n\t":
                    if false {
                        var v24 uint = 250
                        _ = v24
                    }
                }
                if (42 >= 0) {
                    switch "" {
                    case "foo bar":
                        var v25 int64 = 9223372036854775807
                        _ = v25
                    }
                    switch (fi23 % ((fi23) | 1)) {
                    case -1:
                        for fi26 := 0; fi26 < 2; fi26++ {
                            v2 += (v2 - genFn0(65535))
                            fmt.Println(v1, v1)
                            var v27 float32 = 0.0
                            _ = fi26
                            _ = v27
                        }
                    default:
                        var v28 bool = v1
                        _ = v28
                    }
                    var v29 int = -9223372036854775808
                    _ = v29
                }
                fmt.Println(v1, fi23)
                _ = fi23
            }
        }
    }
    v30 := GenS1{F0: 926}
    var v31 uint = uint(v2)
    if v1 {
        fmt.Println(v1, v1, v1)
        if (!(2 < 42)) {
            v32 := GenS2{F0: 2, F1: "\n\t"}
            for fi33 := 0; fi33 < 14; fi33++ {
                v34 := GenS3{F0: 1}
                _ = len(v3)
                _ = fi33
                _ = v34
            }
            var v35 uint = (^(v31))
            _ = v32
            _ = v35
        } else {
            v2 = v2
        }
    } else {
        var v36 uint8 = 255
        _ = v36
    }
    if v1 {
        v1 = (!((1.5 * float64((3.14159 / 1.0))) >= (-2.25 / 1.5)))
    }
    fmt.Println(v2)
    if len(v3) > 0 {
        v3[0] = "foo bar"
    }
    for fi37 := 0; fi37 < 14; fi37++ {
        var v38 int = 9223372036854775807
        var v39 uint = 881
        var v40 int = fi37
        _ = fi37
        _ = v38
        _ = v39
        _ = v40
    }
    fmt.Println(v2, v31, v31)
    fmt.Println(v2, v31)
    if ("" == (("a" + "\n\t") + ((("" + "hello") + "hello") + ("a" + "")))) {
        v1 = true
        var v41 bool = ((genFn0(v31) == (v2 ^ v2)) && true)
        _ = v41
    }
    v42 := GenS2{F0: 42, F1: "a"}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v30
    _ = v31
    _ = v42
}
