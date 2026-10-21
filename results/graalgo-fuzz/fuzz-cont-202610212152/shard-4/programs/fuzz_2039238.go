package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 uint16; F2 int64 }
func genFn0(p0 uint, p1 int8) int8 {
    p1 += (p1 * (p1 << ((p0) & 63)))
    for fi5 := 0; fi5 < 12; fi5++ {
        var v6 uint16 = 32767
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    return p1
}
func genFn1() int64 {
    var v7 uint16 = 1
    v7 = v7
    for fi8 := 0; fi8 < 4; fi8++ {
        v9 := []int{(fi8 & 32767), (fi8 >> ((18446744073709551615) & 63))}
        for fi10 := 0; fi10 < 12; fi10++ {
            if false {
                var v11 uint32 = 1
                _ = v11
            }
            v9 = append(v9, (fi10 & fi10))
            _ = fi10
        }
        for fi12 := 0; fi12 < 17; fi12++ {
            var v13 bool = true
            if false {
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
                v14 := GenS1{F0: 449, F1: "", F2: 0}
                var v15 uint = 1
                _ = v14
                _ = v15
            }
            v7 = v7
            _ = fi12
            _ = v13
        }
        _ = fi8
        _ = v9
    }
    _ = v7
    return 9223372036854775807
}

func main() {
    v1 := make(map[string]int)
    var v2 bool = true
    var v3 float32 = -2.25
    v4 := make(map[int]bool)
    v16 := GenS1{F0: 1000, F1: "hello", F2: -1}
    if false {
        switch ("\n\t" + "\n\t") {
        case "a":
            v3 -= -2.25
        case "\n\t":
            fmt.Println(v3, v2)
        case "foo bar":
            v17 := []int{}
            _ = v17
        default:
            var v18 uint32 = 351
            _ = v18
        }
        v3 = (v3 / v3)
        v19 := v4[720]
        _ = v19
    }
    v20 := []int{}
    if v2 {
        v2 = (2 <= 1)
        v2 = (3.14159 < 0.0)
        var v21 int = -507
        _ = v21
    }
    v3 = v3
    for fi22 := 0; fi22 < 17; fi22++ {
        if len(v20) > 0 {
            v20[0] = -2147483648
        }
        _ = fi22
    }
    if (!v2) {
        v23 := v1["\n\t"]
        var v24 uint32 = 1
        switch 730 {
        case 32767:
            var v25 string = "hello"
            _ = v25
        case -2147483648:
            for fi26 := 0; fi26 < 19; fi26++ {
                var v27 string = (((("" + "a") + ("hello" + ((("hello" + "\n\t") + ((("hello" + "a") + "hello") + "hello")) + ((("hello" + "") + "a") + "hello")))) + (("hello" + "foo bar") + "a")) + ("\n\t" + ("\n\t" + "foo bar")))
                v28 := v1["hello"]
                _ = fi26
                _ = v27
                _ = v28
            }
        }
        _ = v23
        _ = v24
    }
    var v29 int16 = 508
    var v30 uint16 = 42
    v31 := GenS2{F0: 115}
    v32 := GenS3{F0: v2, F1: (v30 ^ v30), F2: 791}
    var v33 float64 = 1e10
    var v34 bool = false
    v35 := GenS2{F0: -1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
}
