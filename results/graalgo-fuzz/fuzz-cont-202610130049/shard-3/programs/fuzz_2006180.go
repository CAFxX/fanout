package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 uint8; F2 uint }
type GenS3 struct { F0 int16; F1 int32; F2 uint64 }
func genFn0() int {
    if false {
        var v6 int8 = 1
        v6 -= 42
        for fi7 := 0; fi7 < 19; fi7++ {
            var v8 int64 = 42
            _ = fi7
            _ = v8
        }
        _ = v6
    } else {
        var v9 uint = 4294967295
        var v10 float64 = 0.0
        _ = v10
        _ = v9
    }
    var v11 int8 = 127
    if (-1 < -1) {
        var v12 int8 = 45
        for fi13 := 0; fi13 < 13; fi13++ {
            var v14 uint64 = 4294967295
            v15 := []int{(fi13 >> ((18446744073709551615) & 63)), (fi13 - 32767)}
            _ = fi13
            _ = v14
            _ = v15
        }
        if false {
            fmt.Println(v11, v12)
            fmt.Println(v11, v12)
            var v16 int16 = 32767
            _ = v16
        } else {
            var v17 int = -665
            v18 := GenS1{F0: "\n\t"}
            _ = v17
            _ = v18
        }
        _ = v12
    }
    fmt.Println(v11)
    var v19 uint64 = 18446744073709551615
    _ = v11
    _ = v19
    return 1
}
func genFn1(p0 int8, p1 int8, p2 uint) int {
    for fi20 := 0; fi20 < 9; fi20++ {
        p1 = p1
        v21 := []int{42, fi20, genFn0()}
        _ = fi20
        _ = v21
    }
    if false {
        var v22 float32 = 0.1
        for fi23 := 0; fi23 < 6; fi23++ {
            p0 -= (-(p1))
            _ = fi23
        }
        _ = v22
    } else {
        p1 = p1
        v24 := GenS2{F0: 4294967295, F1: 0, F2: p2}
        _ = v24
    }
    _ = p0
    _ = p1
    _ = p2
    return 2147483647
}

func main() {
    var v1 int64 = 223
    var v2 uint8 = 42
    v3 := []int{-2147483648}
    var v4 uint32 = 4294967295
    v5 := []string{((("\n\t" + "") + "a") + ("\n\t" + ("a" + "hello"))), "hello", ("" + "a")}
    v2 = (^(v2))
    for fi25 := 0; fi25 < 10; fi25++ {
        var v26 float64 = 1e10
        v26 = v26
        for fi27 := 0; fi27 < 12; fi27++ {
            v28 := GenS1{F0: ""}
            _ = fi27
            _ = v28
        }
        _ = fi25
        _ = v26
    }
    var v29 int16 = 0
    var v30 bool = false
    v31 := GenS1{F0: "hello"}
    var v32 uint64 = 1
    if (971 != -2147483648) {
        if v30 {
            fmt.Println(v29, v1, v2)
        } else {
            v33 := GenS3{F0: 1, F1: 2147483647, F2: (^(v32))}
            var v34 uint8 = 118
            _ = v33
            _ = v34
        }
        if len(v5) > 0 {
            v5[0] = "\n\t"
        }
    } else {
        if ((float64(float32(-2.25)) + 0.1) == (0.0 * -2.25)) {
            for fi35 := 0; fi35 < 5; fi35++ {
                v5 = append(v5, "\n\t")
                v2 = 35
                v36 := GenS3{F0: 32767, F1: 1, F2: (v32 | uint64(v4))}
                _ = fi35
                _ = v36
            }
            var v37 int64 = (v1 & (v1 * -1))
            _ = v37
        }
    }
    for fi38 := 0; fi38 < 10; fi38++ {
        var v39 string = ""
        _ = fi38
        _ = v39
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v2)
    if v30 {
        v1 += v1
        v4 = (^(v4))
        var v40 int = -22
        _ = v40
    } else {
        v30 = v30
        v41 := GenS3{F0: v29, F1: -2147483648, F2: (v32 << ((993) & 63))}
        _ = v41
    }
    var v42 uint = 255
    var v43 int32 = 391
    switch "\n\t" {
    case "":
        if ("" < "hello") {
            v2 += v2
            var v44 bool = v30
            var v45 uint16 = 973
            _ = v44
            _ = v45
        }
    case "a":
        v2 = 0
    case "\n\t":
        v46 := GenS1{F0: "foo bar"}
        _ = v46
    }
    switch ("" + "\n\t") {
    case "\n\t":
        fmt.Println(v4)
    case "":
        v47 := GenS1{F0: "hello"}
        _ = v47
    }
    var v48 uint64 = (v32 % ((v32) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v42
    _ = v43
    _ = v48
    _ = v5
}
