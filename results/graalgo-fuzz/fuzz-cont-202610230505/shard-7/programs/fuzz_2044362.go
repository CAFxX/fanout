package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint8, p1 uint16, p2 uint32) uint64 {
    v4 := []int{}
    switch 127 {
    case -2147483648:
        var v5 int16 = -952
        _ = v5
    case -9223372036854775808:
        fmt.Println(p2, p1, p1)
    }
    p1 += 981
    v6 := GenS1{F0: -9223372036854775808}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v6
    return 2
}
func genFn1(p0 uint8) uint16 {
    p0 = p0
    for fi7 := 0; fi7 < 6; fi7++ {
        if true {
            v8 := GenS1{F0: 960}
            var v9 uint16 = 358
            _ = v8
            _ = v9
        }
        var v10 int = (fi7 | fi7)
        _ = fi7
        _ = v10
    }
    p0 += (p0 - p0)
    if ("\n\t" < ("foo bar" + "hello")) {
        v11 := GenS2{F0: 32767}
        fmt.Println(p0)
        _ = v11
    } else {
        v12 := []int{9223372036854775807, 0, 4294967295, -2147483648}
        switch "a" {
        case "":
            var v13 int32 = -399
            _ = v13
        }
        if len(v12) > 0 {
            v12[0] = 32768
        }
        _ = v12
    }
    p0 += (^(p0))
    _ = p0
    return 892
}

func main() {
    v1 := make(map[string]int)
    var v2 int16 = 32767
    v3 := make(map[string]int)
    v14 := GenS1{F0: 158}
    v15 := GenS1{F0: 0}
    v2 = (v2 | (v2 % (((v2 + (-(v2)))) | 1)))
    var v16 float32 = 0.0
    v16 -= v16
    switch "a" {
    case "foo bar":
        if true {
            if true {
                if false {
                    for fi17 := 0; fi17 < 8; fi17++ {
                        var v18 uint32 = 738
                        _ = fi17
                        _ = v18
                    }
                    delete(v3, "a")
                }
                switch 32767 {
                case 42:
                    fmt.Println(v2, v2)
                case -1:
                    switch (("\n\t" + "") + (("a" + "") + "hello")) {
                    case "":
                        v16 = v16
                    case "\n\t":
                        fmt.Println(v16, v2)
                    default:
                        var v19 string = ""
                        _ = v19
                    }
                case -9223372036854775808:
                    v16 += v16
                }
            } else {
                var v20 uint16 = 0
                fmt.Println(v2, v2, v20)
                v16 = float32(0.1)
                _ = v20
            }
        }
    case "":
        v16 += v16
    case "\n\t":
        var v21 string = ("a" + "\n\t")
        _ = v21
    default:
        var v22 uint16 = 0
        _ = v22
    }
    for fi23 := 0; fi23 < 19; fi23++ {
        for fi24 := 0; fi24 < 16; fi24++ {
            fmt.Println(v16)
            fmt.Println(fi24)
            v25 := []int{fi23, fi23, fi24}
            _ = fi24
            _ = v25
        }
        _ = fi23
    }
    v16 -= (((float32(-2.25) / 1.0) / (float32(0.0) / v16)) + 0.0)
    switch "\n\t" {
    case "foo bar":
        v16 -= (v16 * 1e10)
    case "":
        v2 -= (v2 + (v2 - 32767))
    }
    v26 := []int{}
    if ("\n\t" > "") {
        v27 := GenS1{F0: 9223372036854775807}
        var v28 int32 = -2147483648
        _ = v27
        _ = v28
    }
    for fi29 := 0; fi29 < 11; fi29++ {
        var v30 uint8 = 0
        v31 := v3[""]
        v32 := v3[""]
        _ = fi29
        _ = v30
        _ = v31
        _ = v32
    }
    for fi33 := 0; fi33 < 8; fi33++ {
        delete(v1, "")
        _ = fi33
    }
    if true {
        switch 128 {
        case -560:
            var v34 int64 = -727
            _ = v34
        default:
            if len(v26) >= 2 {
                v26 = v26[1:len(v26)]
            }
        }
        switch 127 {
        case 0:
            v35 := v1[""]
            _ = v35
        case -1:
            var v36 int8 = 0
            _ = v36
        }
        switch "foo bar" {
        case "a":
            v26 = append(v26, -686)
        case "\n\t":
            delete(v1, ("a" + ""))
        case "foo bar":
            fmt.Println(v16, v16)
        default:
            v26 = append(v26, 32768)
        }
    } else {
        v16 = (v16 / float32((float64(0.0) + (3.14159 + -2.25))))
        var v37 uint8 = 0
        if true {
            fmt.Println(v2, v37)
        }
        _ = v37
    }
    for fi38 := 0; fi38 < 17; fi38++ {
        for fi39 := 0; fi39 < 19; fi39++ {
            v16 = (0.0 / 1.5)
            v40 := GenS1{F0: -1}
            _ = fi39
            _ = v40
        }
        _ = fi38
    }
    v16 = v16
    var v41 int32 = 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v41:", v41)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v26
    _ = v3
    _ = v41
}
