package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 uint32 }
type GenS2 struct { F0 int16; F1 uint16; F2 int }
func genFn0(p0 int64) uint64 {
    for fi5 := 0; fi5 < 8; fi5++ {
        var v6 int8 = 42
        if true {
            v6 -= v6
            for fi7 := 0; fi7 < 1; fi7++ {
                v8 := []int{}
                if (v6 >= (v6 / ((int8(fi7)) | 1))) {
                    switch "a" {
                    case "hello":
                        _ = len(v8)
                    case "a":
                        var v9 int64 = p0
                        _ = v9
                    case "\n\t":
                        fmt.Println(fi5, fi5)
                    }
                }
                if true {
                    var v10 int32 = 375
                    _ = v10
                } else {
                    v6 -= 1
                    v11 := GenS1{F0: 1, F1: "hello", F2: 4294967295}
                    _ = v11
                }
                _ = fi7
                _ = v8
            }
        }
        _ = fi5
        _ = v6
    }
    p0 = (p0 ^ (p0 << ((2) & 63)))
    var v12 uint64 = 0
    _ = p0
    _ = v12
    return v12
}
func genFn1() int8 {
    switch 32768 {
    case 0:
        var v13 uint8 = 42
        _ = v13
    case 42:
        var v14 uint32 = 0
        _ = v14
    case -650:
        var v15 uint = 18446744073709551615
        _ = v15
    default:
        switch "hello" {
        case "\n\t":
            var v16 uint = 1
            _ = v16
        }
    }
    var v17 int64 = 0
    for fi18 := 0; fi18 < 15; fi18++ {
        v17 = v17
        var v19 uint16 = 961
        _ = fi18
        _ = v19
    }
    v20 := GenS1{F0: 2, F1: "foo bar", F2: 2147483648}
    _ = v17
    _ = v20
    return 42
}

func main() {
    v1 := []int{2147483647, 42, 802, -805}
    v2 := make(map[string]string)
    v3 := []uint32{65535, 1, 4294967295}
    var v4 float32 = -2.25
    var v21 int = 130
    if (42 <= 4294967295) {
        v4 += 0.0
    }
    if false {
        if true {
            var v22 uint = 4294967295
            _ = v22
        }
    } else {
        if (1e10 >= ((float64(v4) * 3.14159) / (0.0 - 0.1))) {
            for fi23 := 0; fi23 < 7; fi23++ {
                var v24 uint8 = 119
                _ = fi23
                _ = v24
            }
        }
        switch (v21 >> ((255) & 63)) {
        case 0:
            if true {
                var v25 uint64 = 4294967295
                _ = v25
            } else {
                v4 = 0.0
                v26 := GenS1{F0: 18446744073709551615, F1: ("foo bar" + "foo bar"), F2: 4294967295}
                _ = v26
            }
        }
        v27 := GenS1{F0: 2, F1: "foo bar", F2: 127}
        _ = v27
    }
    v4 = v4
    v21 += v21
    if false {
        for fi28 := 0; fi28 < 20; fi28++ {
            switch v21 {
            case 2147483647:
                for fi29 := 0; fi29 < 20; fi29++ {
                    v21 += fi28
                    _ = fi29
                }
            default:
                fmt.Println(v4)
            }
            _ = fi28
        }
        var v30 float64 = (1.5 - (1.5 * 3.14159))
        v31 := GenS2{F0: 1, F1: 0, F2: 9223372036854775807}
        _ = v30
        _ = v31
    } else {
        var v32 int16 = 32767
        _ = len(v2)
        _ = v32
    }
    v4 -= v4
    if (genFn1() == 1) {
        if true {
            var v33 int = 127
            if (!(1 < 128)) {
                if true {
                    for fi34 := 0; fi34 < 11; fi34++ {
                        v33 = v21
                        var v35 int8 = -1
                        _ = fi34
                        _ = v35
                    }
                    v36 := GenS2{F0: 42, F1: 853, F2: v33}
                    _ = v36
                }
                v37 := v2[""]
                _ = v37
            } else {
                v21 += v33
            }
            var v38 uint = 42
            _ = v33
            _ = v38
        } else {
            v4 -= (v4 - v4)
        }
    } else {
        v21 = v21
        v21 += (^(v21))
        fmt.Println(v4, v21)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
