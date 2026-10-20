package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 uint }
func genFn0(p0 int) uint16 {
    v4 := []int{(p0 + p0), (-(p0))}
    for fi5 := 0; fi5 < 3; fi5++ {
        if false {
            for fi6 := 0; fi6 < 1; fi6++ {
                var v7 uint32 = 2147483647
                _ = fi6
                _ = v7
            }
        } else {
            var v8 uint8 = 1
            fmt.Println(fi5, v8)
            p0 += -9223372036854775808
            _ = v8
        }
        p0 += p0
        for fi9 := 0; fi9 < 11; fi9++ {
            switch "a" {
            case "":
                var v10 string = ""
                _ = v10
            case "\n\t":
                var v11 float32 = 0.0
                _ = v11
            case "hello":
                var v12 uint = 255
                _ = v12
            default:
                p0 = fi9
            }
            _ = len(v4)
            _ = fi9
        }
        _ = fi5
    }
    var v13 int8 = -35
    if (true || (!true)) {
        fmt.Println(p0)
    }
    _ = p0
    _ = v13
    _ = v4
    return 255
}
func genFn1(p0 int8, p1 bool) int8 {
    v14 := GenS1{F0: "foo bar", F1: 9223372036854775807, F2: 721}
    fmt.Println(p0)
    p1 = true
    var v15 int64 = 1
    v15 = v15
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    return 1
}

func main() {
    var v1 bool = (true && (2147483647 > -724))
    var v2 int = 0
    var v3 uint16 = 411
    switch "hello" {
    case "hello":
        v3 -= 65535
    case "a":
        var v16 int32 = 0
        _ = v16
    case "foo bar":
        v17 := GenS1{F0: "hello", F1: -660, F2: 2}
        _ = v17
    }
    for fi18 := 0; fi18 < 10; fi18++ {
        if false {
            for fi19 := 0; fi19 < 2; fi19++ {
                v20 := GenS1{F0: "foo bar", F1: (fi18 << ((4294967295) & 63)), F2: 65535}
                fmt.Println(v1, v3, fi18)
                v21 := []int{32767, int(v3), fi19}
                _ = fi19
                _ = v20
                _ = v21
            }
        }
        v3 -= 676
        _ = fi18
    }
    var v22 int16 = 0
    for fi23 := 0; fi23 < 6; fi23++ {
        for fi24 := 0; fi24 < 10; fi24++ {
            var v25 bool = (v2 < 1)
            var v26 float64 = ((0.1 - 1e10) - (1e10 * (0.0 / 1.0)))
            _ = fi24
            _ = v25
            _ = v26
        }
        for fi27 := 0; fi27 < 6; fi27++ {
            v28 := []int{9223372036854775807}
            v1 = (false && v1)
            _ = fi27
            _ = v28
        }
        _ = fi23
    }
    v29 := []int{v2}
    var v30 string = "foo bar"
    for fi31 := 0; fi31 < 10; fi31++ {
        for fi32 := 0; fi32 < 7; fi32++ {
            v33 := GenS1{F0: "\n\t", F1: (^(v2)), F2: 0}
            _ = fi32
            _ = v33
        }
        switch fi31 {
        case 128:
            v30 = (v30 + ((v30 + v30) + v30))
        default:
            var v34 uint64 = 1
            _ = v34
        }
        _ = fi31
    }
    for fi35 := 0; fi35 < 12; fi35++ {
        v29 = append(v29, fi35)
        fmt.Println(v1, v30)
        v22 += v22
        _ = fi35
    }
    v22 = v22
    v30 = v30
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v22
    _ = v29
    _ = v3
    _ = v30
}
