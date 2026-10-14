package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int64 }
func genFn0() uint16 {
    var v6 int16 = 723
    v6 = v6
    _ = v6
    return 65535
}
func genFn1(p0 int32, p1 int16) int8 {
    p0 = (p0 / ((p0) | 1))
    switch "a" {
    case "a":
        for fi7 := 0; fi7 < 3; fi7++ {
            v8 := GenS1{F0: -1}
            for fi9 := 0; fi9 < 15; fi9++ {
                p1 += (p1 / ((p1) | 1))
                v10 := []int{(fi7 >> ((4294967295) & 63)), fi9, fi7}
                _ = fi9
                _ = v10
            }
            switch "" {
            case "\n\t":
                p1 = (p1 - (p1 ^ (p1 % ((361) | 1))))
            }
            _ = fi7
            _ = v8
        }
    case "foo bar":
        for fi11 := 0; fi11 < 15; fi11++ {
            var v12 string = ("hello" + "\n\t")
            var v13 float64 = ((1.5 - 0.1) * 0.1)
            v14 := []int{(fi11 + (^(fi11))), (fi11 & 32767), fi11}
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
    case "":
        switch 9223372036854775807 {
        case 1:
            v15 := GenS2{F0: 2}
            _ = v15
        case -9223372036854775808:
            var v16 int = 2147483647
            _ = v16
        case -1:
            if (1 > 127) {
                fmt.Println(p1)
                p1 = 119
                v17 := GenS3{F0: 42, F1: 4294967295}
                _ = v17
            } else {
                for fi18 := 0; fi18 < 7; fi18++ {
                    for fi19 := 0; fi19 < 10; fi19++ {
                        v20 := []int{(-(fi19)), fi18, (fi19 % ((fi18) | 1))}
                        _ = fi19
                        _ = v20
                    }
                    p0 -= p0
                    var v21 int8 = 1
                    _ = fi18
                    _ = v21
                }
                v22 := GenS3{F0: 42, F1: 255}
                for fi23 := 0; fi23 < 18; fi23++ {
                    fmt.Println(fi23, fi23, fi23)
                    _ = fi23
                }
                _ = v22
            }
        }
    default:
        if true {
            for fi24 := 0; fi24 < 4; fi24++ {
                v25 := GenS4{F0: 0}
                p1 += 32767
                switch "" {
                case "":
                    var v26 int8 = -109
                    _ = v26
                case "hello":
                    var v27 int64 = -9223372036854775808
                    _ = v27
                case "a":
                    var v28 int = (fi24 | fi24)
                    _ = v28
                default:
                    v29 := []int{89, (fi24 ^ 42), (fi24 & fi24)}
                    _ = v29
                }
                _ = fi24
                _ = v25
            }
            if (128 >= 128) {
                v30 := GenS5{F0: 0}
                v31 := GenS3{F0: 65535, F1: 2}
                v32 := GenS3{F0: 914, F1: 249}
                _ = v30
                _ = v31
                _ = v32
            } else {
                p0 += p0
                v33 := []int{}
                _ = v33
            }
            fmt.Println(p0, p0)
        } else {
            var v34 int8 = 42
            _ = v34
        }
    }
    for fi35 := 0; fi35 < 2; fi35++ {
        fmt.Println(p1)
        for fi36 := 0; fi36 < 7; fi36++ {
            var v37 int16 = (p1 - p1)
            if ("a" >= ("a" + "hello")) {
                var v38 uint16 = 324
                var v39 int16 = int16(v38)
                v40 := GenS5{F0: 9223372036854775807}
                _ = v38
                _ = v39
                _ = v40
            } else {
                var v41 uint16 = 255
                var v42 string = "foo bar"
                v43 := []int{fi35, fi36, 0}
                _ = v41
                _ = v42
                _ = v43
            }
            _ = fi36
            _ = v37
        }
        _ = fi35
    }
    _ = p0
    _ = p1
    return -1
}

func main() {
    v1 := make(map[int]string)
    var v2 uint = 4294967295
    v3 := []uint32{138}
    v4 := []uint32{867}
    v5 := make(map[string]int)
    fmt.Println(v2)
    switch 4294967295 {
    case 887:
        var v44 int = 98
        _ = v44
    case 9223372036854775807:
        v2 += (v2 + v2)
    }
    if false {
        v2 -= v2
        v45 := GenS5{F0: -460}
        _ = v45
    }
    var v46 uint16 = 2
    switch ("foo bar" + ("a" + ("\n\t" + ((("\n\t" + "a") + "\n\t") + ("" + "hello"))))) {
    case "a":
        for fi47 := 0; fi47 < 3; fi47++ {
            var v48 uint16 = 65535
            var v49 int = fi47
            _ = fi47
            _ = v48
            _ = v49
        }
    case "foo bar":
        fmt.Println(v46, v46)
    }
    var v50 uint64 = 0
    var v51 uint = (v2 - 255)
    v50 -= (^(v50))
    v51 -= (v2 >> (((v51 & v2)) & 63))
    v2 = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v50
    _ = v51
}
