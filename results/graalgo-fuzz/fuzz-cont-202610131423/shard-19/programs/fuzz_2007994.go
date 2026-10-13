package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 uint64 }
type GenS2 struct { F0 uint; F1 int64; F2 int16 }
func genFn0(p0 int8, p1 uint64) int32 {
    var v6 uint16 = 963
    if false {
        v6 = (v6 % (((v6 >> ((4294967295) & 63))) | 1))
        for fi7 := 0; fi7 < 13; fi7++ {
            v8 := []int{fi7, (fi7 + (fi7 - (fi7 ^ 4294967295))), fi7, fi7}
            p0 = (-(p0))
            for fi9 := 0; fi9 < 16; fi9++ {
                var v10 uint8 = 1
                _ = fi9
                _ = v10
            }
            _ = fi7
            _ = v8
        }
        p1 += p1
    }
    var v11 float32 = 1.5
    if false {
        var v12 uint64 = p1
        _ = v12
    } else {
        var v13 int = 2147483647
        var v14 uint16 = v6
        _ = v13
        _ = v14
    }
    switch "a" {
    case "":
        v6 = v6
    case "a":
        var v15 uint32 = 1
        _ = v15
    default:
        for fi16 := 0; fi16 < 1; fi16++ {
            v17 := GenS1{F0: uint32(p1), F1: fi16, F2: 769}
            switch 1 {
            case 9223372036854775807:
                for fi18 := 0; fi18 < 16; fi18++ {
                    v11 += float32((((-2.25 + (3.14159 / 1.5)) * (1.5 / 3.14159)) + (1e10 / 1.5)))
                    _ = fi18
                }
            }
            fmt.Println(fi16, fi16)
            _ = fi16
            _ = v17
        }
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v6
    return -148
}
func genFn1(p0 uint16) int8 {
    p0 = p0
    var v19 int64 = -9223372036854775808
    p0 -= 1
    v19 = v19
    switch (("foo bar" + "hello") + "a") {
    case "hello":
        v20 := GenS2{F0: 50, F1: (v19 & v19), F2: 1000}
        _ = v20
    default:
        v21 := []int{234, 4294967295, -9223372036854775808, 32768}
        _ = v21
    }
    _ = p0
    _ = v19
    return 89
}

func main() {
    var v1 int = 1
    var v2 float64 = 0.1
    var v3 float32 = -2.25
    v4 := []int{v1}
    v5 := make(map[string]int)
    var v22 uint = 231
    fmt.Println(v1)
    v3 = 1.5
    switch -671 {
    case 836:
        if len(v4) > 0 {
            v4[0] = v1
        }
    case -309:
        v22 += (v22 - (^(v22)))
    default:
        var v23 int16 = -32768
        _ = v23
    }
    var v24 uint16 = 65535
    var v25 uint8 = 255
    var v26 uint8 = (v25 | (v25 + (v25 - v25)))
    for fi27 := 0; fi27 < 11; fi27++ {
        var v28 uint16 = v24
        _ = fi27
        _ = v28
    }
    fmt.Println(v1)
    for fi29 := 0; fi29 < 16; fi29++ {
        var v30 string = "\n\t"
        var v31 uint = v22
        if (v25 == (^(v25))) {
            for fi32 := 0; fi32 < 2; fi32++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v33 float64 = v2
                switch 128 {
                case 9223372036854775807:
                    v30 = v30
                default:
                    v24 -= (v24 | v24)
                }
                v22 = v31
                var v34 uint8 = 0
                _ = fi32
                _ = v33
                _ = v34
            }
            _ = len(v4)
        }
        _ = fi29
        _ = v30
        _ = v31
    }
    v24 -= v24
    var v35 uint16 = (v24 | v24)
    var v36 int8 = 88
    var v37 uint64 = 1
    var v38 string = "\n\t"
    v39 := GenS1{F0: 1, F1: v1, F2: (v37 | (^(v37)))}
    fmt.Println(v1)
    for fi40 := 0; fi40 < 5; fi40++ {
        for fi41 := 0; fi41 < 19; fi41++ {
            if false {
                v2 += (float64(v3) + float64(float32(v2)))
                fmt.Println(v1)
            } else {
                switch fi40 {
                case 613:
                    v4 = append(v4, (fi40 << ((4294967295) & 63)))
                }
            }
            _ = fi41
        }
        _ = fi40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v22
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
