package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 string }
func genFn0(p0 uint32, p1 uint32) uint8 {
    var v7 float32 = float32(1e10)
    switch ((("\n\t" + (("hello" + "") + "foo bar")) + (("a" + "foo bar") + "")) + "\n\t") {
    case "\n\t":
        p0 -= p1
    case "a":
        switch -9223372036854775808 {
        case 127:
            var v8 uint32 = p1
            _ = v8
        case -1:
            var v9 float32 = v7
            _ = v9
        default:
            if (false || false) {
                for fi10 := 0; fi10 < 5; fi10++ {
                    v11 := []int{4294967295}
                    _ = len(v11)
                    _ = fi10
                    _ = v11
                }
                if true {
                    for fi12 := 0; fi12 < 15; fi12++ {
                        for fi13 := 0; fi13 < 8; fi13++ {
                            p0 = 786
                            if (((3.14159 + 3.14159) + 0.1) <= 0.0) {
                                v14 := []int{fi13, fi12}
                                _ = v14
                            } else {
                                v15 := GenS1{F0: 42, F1: -608}
                                fmt.Println(p1, v7)
                                _ = v15
                            }
                            _ = fi13
                        }
                        _ = fi12
                    }
                }
                var v16 string = ("" + "hello")
                _ = v16
            } else {
                var v17 float64 = -2.25
                fmt.Println(p0)
                p0 = p1
                _ = v17
            }
        }
    default:
        if (true || true) {
            fmt.Println(p0, v7)
        }
    }
    v18 := []int{-2147483648}
    _ = p0
    _ = p1
    _ = v18
    _ = v7
    return 0
}
func genFn1() int32 {
    if (-574 > -709) {
        var v19 int32 = 42
        switch "a" {
        case "":
            var v20 int16 = 1
            _ = v20
        }
        _ = v19
    }
    for fi21 := 0; fi21 < 5; fi21++ {
        var v22 int16 = -1
        switch "foo bar" {
        case "a":
            for fi23 := 0; fi23 < 5; fi23++ {
                v22 = v22
                _ = fi23
            }
        case "\n\t":
            var v24 float32 = 0.1
            _ = v24
        case "":
            var v25 uint64 = 559
            _ = v25
        }
        _ = fi21
        _ = v22
    }
    return 330
}

func main() {
    var v1 uint16 = 1
    var v2 int = -2147483648
    var v3 int32 = 2147483647
    v4 := make(map[string]bool)
    var v5 float64 = 1e10
    var v6 float32 = 0.1
    v26 := GenS2{F0: ("\n\t" + "foo bar")}
    if true {
        v1 += (v1 / ((v1) | 1))
        var v27 float64 = v5
        fmt.Println(v2)
        _ = v27
    }
    var v28 int64 = -233
    var v29 float32 = ((((v6 - float32(0.1)) / 1e10) * 0.1) * 1.5)
    v1 += (v1 / ((42) | 1))
    var v30 uint16 = (v1 + (v1 * 100))
    var v31 int32 = v3
    for fi32 := 0; fi32 < 3; fi32++ {
        var v33 int8 = -76
        if false {
            v34 := GenS1{F0: 1000, F1: 32767}
            for fi35 := 0; fi35 < 2; fi35++ {
                v36 := []int{(fi35 << ((0) & 63))}
                v4["a"] = true
                _ = fi35
                _ = v36
            }
            _ = v34
        }
        switch (v2 << ((1) & 63)) {
        case 2147483647:
            var v37 uint32 = 2147483648
            _ = v37
        case 448:
            var v38 float32 = 0.0
            _ = v38
        default:
            for fi39 := 0; fi39 < 16; fi39++ {
                v2 += fi32
                v40 := []int{}
                var v41 int32 = (v3 & 0)
                _ = fi39
                _ = v40
                _ = v41
            }
        }
        _ = fi32
        _ = v33
    }
    delete(v4, "hello")
    fmt.Println(v6, v31, v5)
    var v42 int64 = (v28 - v28)
    v43 := GenS2{F0: ""}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
