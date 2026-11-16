package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 int32 }
func genFn0() bool {
    v7 := GenS1{F0: 0}
    if true {
        switch 2147483647 {
        case 42:
            v8 := GenS1{F0: 1}
            _ = v8
        case -2147483648:
            for fi9 := 0; fi9 < 12; fi9++ {
                var v10 uint = 65535
                for fi11 := 0; fi11 < 8; fi11++ {
                    if false {
                        v12 := []int{1, (fi11 | fi11), (fi11 ^ fi9)}
                        v10 += 744
                        fmt.Println(fi11, fi11, fi9)
                        if ("a" != "foo bar") {
                            fmt.Println(v10)
                        } else {
                            var v13 uint64 = 18446744073709551615
                            _ = v13
                        }
                        _ = v12
                    }
                    var v14 string = "hello"
                    _ = fi11
                    _ = v14
                }
                var v15 int = (fi9 & fi9)
                _ = fi9
                _ = v10
                _ = v15
            }
        }
    }
    v16 := GenS1{F0: 128}
    _ = v16
    _ = v7
    return false
}
func genFn1() int8 {
    if (((float64(3.14159) < 0.0) && (!false)) || false) {
        switch "a" {
        case "":
            var v17 uint32 = 592
            _ = v17
        }
        for fi18 := 0; fi18 < 19; fi18++ {
            for fi19 := 0; fi19 < 16; fi19++ {
                fmt.Println(fi19)
                _ = fi19
            }
            var v20 uint32 = 65535
            _ = fi18
            _ = v20
        }
    }
    var v21 int = -9223372036854775808
    v21 -= 826
    fmt.Println(v21)
    for fi22 := 0; fi22 < 4; fi22++ {
        var v23 uint = 18446744073709551615
        switch "foo bar" {
        case "\n\t":
            var v24 uint64 = 18446744073709551615
            _ = v24
        case "":
            v25 := []int{493}
            _ = v25
        default:
            var v26 uint = 255
            _ = v26
        }
        _ = fi22
        _ = v23
    }
    _ = v21
    return -90
}

func main() {
    v1 := []int{-261}
    v2 := make(map[string]int)
    var v3 uint32 = 2
    var v4 int64 = -9223372036854775808
    var v5 uint64 = 669
    v6 := make(map[int]uint32)
    var v27 uint64 = (v5 % (((v5 ^ (^(v5)))) | 1))
    var v28 int64 = (v4 / ((int64(v27)) | 1))
    switch 127 {
    case 2147483647:
        if ((-2.25 > 1e10) || (126 >= 0)) {
            fmt.Println(v27, v28, v4)
            if true {
                for fi29 := 0; fi29 < 11; fi29++ {
                    v1 = append(v1, fi29)
                    _ = fi29
                }
                switch ((("foo bar" + ("" + "a")) + "\n\t") + "a") {
                case "a":
                    var v30 string = ("a" + "\n\t")
                    _ = v30
                case "foo bar":
                    v27 += (v5 % (((v27 - (^(v5)))) | 1))
                default:
                    var v31 uint16 = 0
                    _ = v31
                }
            }
            if len(v1) > 0 {
                v1[0] = -1
            }
        }
    default:
        v32 := GenS1{F0: 42}
        _ = v32
    }
    for fi33 := 0; fi33 < 18; fi33++ {
        v34 := GenS1{F0: 144}
        if false {
            v4 += (v4 ^ 42)
        } else {
            var v35 uint = 18446744073709551615
            v27 += 118
            var v36 int8 = genFn1()
            _ = v35
            _ = v36
        }
        _ = fi33
        _ = v34
    }
    var v37 bool = (false && true)
    switch ("a" + "foo bar") {
    case "":
        var v38 uint64 = v27
        _ = v38
    case "hello":
        for fi39 := 0; fi39 < 7; fi39++ {
            v4 = v28
            var v40 uint16 = uint16(v5)
            _ = fi39
            _ = v40
        }
    default:
        v27 += (v27 * (^(v5)))
    }
    v4 += -1
    v37 = true
    v5 = 4294967295
    fmt.Println(v28, v4)
    v41 := GenS2{F0: (v3 & v3), F1: -2147483648}
    var v42 uint8 = 127
    for fi43 := 0; fi43 < 15; fi43++ {
        v28 += 42
        _ = fi43
    }
    switch 0 {
    case -2147483648:
        v27 += 0
    case 128:
        v5 += (v5 & v27)
    case -872:
        v44 := GenS1{F0: 0}
        _ = v44
    default:
        var v45 int = 671
        _ = v45
    }
    switch "a" {
    case "foo bar":
        var v46 uint64 = 2
        _ = v46
    case "\n\t":
        v5 += 2
    case "a":
        fmt.Println(v28, v28, v27)
    default:
        v2[""] = 42
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v37
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
