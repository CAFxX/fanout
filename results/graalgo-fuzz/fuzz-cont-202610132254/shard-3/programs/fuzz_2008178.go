package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 int8; F2 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint; F1 string; F2 uint }

func main() {
    v1 := make(map[string]int)
    var v2 bool = false
    var v3 float64 = 1e10
    var v4 int8 = 42
    var v5 bool = (0 < 897)
    var v6 uint64 = 2
    if (1 >= 1) {
        var v7 uint = 1
        v8 := GenS1{F0: (v6 / ((v6) | 1))}
        if (42 >= 0) {
            for fi9 := 0; fi9 < 13; fi9++ {
                v2 = (!v2)
                _ = fi9
            }
            var v10 uint64 = 910
            v11 := []int{42, 4294967295, 337}
            _ = v10
            _ = v11
        } else {
            for fi12 := 0; fi12 < 1; fi12++ {
                for fi13 := 0; fi13 < 2; fi13++ {
                    fmt.Println(v2)
                    switch "a" {
                    case "":
                        var v14 float64 = v3
                        _ = v14
                    case "\n\t":
                        v5 = (!(!(!v5)))
                    }
                    v7 += (^(v7))
                    _ = fi13
                }
                if ((v4 < 0) && false) {
                    switch (("a" + ("a" + (("foo bar" + "foo bar") + ""))) + ("a" + "a")) {
                    case "hello":
                        v15 := GenS2{F0: 0, F1: v4, F2: v6}
                        _ = v15
                    case "":
                        fmt.Println(v2, fi12, v6)
                    default:
                        v16 := GenS3{F0: 935}
                        _ = v16
                    }
                } else {
                    v17 := []int{}
                    for fi18 := 0; fi18 < 7; fi18++ {
                        v7 = v7
                        v19 := GenS2{F0: (fi12 ^ -1), F1: v4, F2: (^(v6))}
                        var v20 uint16 = 32767
                        v21 := GenS1{F0: v6}
                        _ = fi18
                        _ = v19
                        _ = v20
                        _ = v21
                    }
                    var v22 int8 = 100
                    _ = v17
                    _ = v22
                }
                _ = fi12
            }
            v3 += (v3 * 0.0)
            if true {
                v4 = -1
            }
        }
        _ = v7
        _ = v8
    } else {
        fmt.Println(v2, v2)
        v4 += v4
        switch -943 {
        case -454:
            for fi23 := 0; fi23 < 6; fi23++ {
                v24 := []int{(-(fi23)), (fi23 >> ((4294967295) & 63)), 2147483647}
                v25 := GenS4{F0: fi23}
                _ = fi23
                _ = v24
                _ = v25
            }
        case 42:
            v3 += -2.25
        case -2147483648:
            var v26 int = 879
            _ = v26
        default:
            v27 := GenS2{F0: 32768, F1: 127, F2: v6}
            _ = v27
        }
    }
    if (!true) {
        v3 -= float64(float32(v3))
    }
    if true {
        for fi28 := 0; fi28 < 1; fi28++ {
            var v29 int32 = 0
            v30 := GenS5{F0: 65535, F1: "\n\t", F2: 0}
            _ = fi28
            _ = v29
            _ = v30
        }
        var v31 uint8 = 111
        fmt.Println(v3, v5, v5)
        _ = v31
    } else {
        for fi32 := 0; fi32 < 16; fi32++ {
            v33 := GenS2{F0: fi32, F1: v4, F2: (v6 * (v6 - (v6 - (v6 | v6))))}
            _ = fi32
            _ = v33
        }
        v34 := []int{32767, 9223372036854775807, 4294967295, -515}
        for fi35 := 0; fi35 < 17; fi35++ {
            var v36 uint64 = v6
            _ = fi35
            _ = v36
        }
        _ = v34
    }
    fmt.Println(v4, v3, v4)
    var v37 uint64 = v6
    var v38 float64 = v3
    if true {
        fmt.Println(v38, v4, v38)
    } else {
        v39 := GenS1{F0: (^(v37))}
        v40 := v1["a"]
        v41 := GenS4{F0: v40}
        _ = v39
        _ = v40
        _ = v41
    }
    var v42 bool = (v37 >= v6)
    for fi43 := 0; fi43 < 16; fi43++ {
        fmt.Println(fi43)
        _ = fi43
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
