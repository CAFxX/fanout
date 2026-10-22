package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 uint }
func genFn0(p0 string, p1 int64, p2 bool) bool {
    p0 = ("" + "\n\t")
    p0 = p0
    p0 = ""
    var v6 float64 = float64(3.14159)
    v7 := GenS1{F0: -938, F1: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return p2
}
func genFn1(p0 int64, p1 int32, p2 int) int8 {
    var v8 uint32 = 663
    p1 += -1
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 126
}

func main() {
    var v1 string = (("" + "foo bar") + "foo bar")
    v2 := []string{}
    var v3 uint64 = 0
    var v4 int16 = 696
    v5 := make(map[string]string)
    v9 := v5[""]
    if false {
        v4 = v4
        if false {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if true {
                for fi10 := 0; fi10 < 8; fi10++ {
                    switch "" {
                    case "a":
                        v11 := GenS2{F0: 65535}
                        _ = v11
                    case "":
                        var v12 uint64 = (v3 ^ v3)
                        _ = v12
                    case "hello":
                        var v13 uint16 = 0
                        _ = v13
                    default:
                        v4 = 1
                    }
                    if (127 != 127) {
                        v4 = -32768
                        fmt.Println(v1, v9)
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    _ = fi10
                }
                v1 = v9
            }
        } else {
            v4 -= v4
        }
    } else {
        var v14 int8 = -119
        var v15 float64 = (1e10 + 0.0)
        _ = v14
        _ = v15
    }
    if (false && (127 < 0)) {
        v4 -= v4
    }
    var v16 uint64 = 1
    v1 = v9
    var v17 uint32 = 65535
    for fi18 := 0; fi18 < 2; fi18++ {
        var v19 int32 = 112
        var v20 int8 = 119
        v21 := GenS3{F0: 254, F1: v3}
        _ = fi18
        _ = v19
        _ = v20
        _ = v21
    }
    switch -2147483648 {
    case 32768:
        for fi22 := 0; fi22 < 5; fi22++ {
            v4 = v4
            _ = fi22
        }
    }
    fmt.Println(v4, v1, v16)
    v23 := GenS4{F0: 69}
    if false {
        fmt.Println(v4, v9)
        for fi24 := 0; fi24 < 20; fi24++ {
            v17 -= uint32(v16)
            if (42 < 1) {
                if len(v2) > 0 {
                    v2[0] = ("a" + ("\n\t" + "a"))
                }
                _ = len(v2)
                v25 := v5["hello"]
                _ = v25
            } else {
                if len(v2) > 0 {
                    v2[0] = v1
                }
            }
            _ = fi24
        }
    }
    if true {
        v26 := v5[v1]
        var v27 bool = (42 < 17)
        _ = v26
        _ = v27
    } else {
        var v28 uint = 771
        var v29 int16 = (v4 + 1)
        _ = v28
        _ = v29
    }
    v30 := GenS1{F0: v4, F1: v17}
    switch 229 {
    case -2147483648:
        v17 += v17
    default:
        if true {
            v1 = v9
            if true {
                var v31 string = ""
                _ = v31
            } else {
                v9 = ("" + "a")
                v32 := GenS4{F0: 2}
                _ = v32
            }
        }
    }
    v33 := GenS3{F0: 1, F1: 578}
    for fi34 := 0; fi34 < 2; fi34++ {
        var v35 int = fi34
        var v36 uint64 = (v3 | 0)
        var v37 int8 = -128
        _ = fi34
        _ = v35
        _ = v36
        _ = v37
    }
    switch v9 {
    case "":
        for fi38 := 0; fi38 < 3; fi38++ {
            var v39 int = 806
            _ = fi38
            _ = v39
        }
    case "a":
        v40 := GenS1{F0: (v4 % ((1) | 1)), F1: v17}
        _ = v40
    case "foo bar":
        var v41 uint16 = 42
        _ = v41
    }
    v1 = (("\n\t" + ((v1 + (v1 + (v9 + "a"))) + "hello")) + v9)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v30
    _ = v33
    _ = v4
    _ = v5
    _ = v9
}
