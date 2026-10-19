package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8; F1 string }
func genFn0() int {
    if false {
        var v6 uint8 = 42
        var v7 int8 = 127
        var v8 int16 = 1
        _ = v6
        _ = v7
        _ = v8
    }
    var v9 uint64 = 2
    for fi10 := 0; fi10 < 7; fi10++ {
        v11 := GenS1{F0: 171, F1: 0}
        for fi12 := 0; fi12 < 18; fi12++ {
            var v13 int32 = 2147483647
            v14 := []int{(-(fi10)), (fi10 + (fi10 >> ((4294967295) & 63)))}
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = fi10
        _ = v11
    }
    v15 := GenS2{F0: 2147483647}
    switch "" {
    case "a":
        fmt.Println(v9)
    case "foo bar":
        if true {
            v9 = 2
            v9 = v9
            v16 := []int{-442, 2147483647, 32768, 128}
            _ = v16
        }
    case "hello":
        for fi17 := 0; fi17 < 20; fi17++ {
            v9 -= 18446744073709551615
            v18 := GenS2{F0: -584}
            _ = fi17
            _ = v18
        }
    default:
        var v19 string = ""
        _ = v19
    }
    _ = v15
    _ = v9
    return 1
}
func genFn1(p0 uint16) int {
    for fi20 := 0; fi20 < 14; fi20++ {
        for fi21 := 0; fi21 < 4; fi21++ {
            v22 := []int{fi20, 32768, genFn0(), (fi20 << ((42) & 63))}
            _ = fi21
            _ = v22
        }
        if (255 < 1) {
            fmt.Println(fi20, p0)
            v23 := GenS1{F0: 18446744073709551615, F1: 9223372036854775807}
            _ = v23
        }
        _ = fi20
    }
    var v24 int = -9223372036854775808
    _ = p0
    _ = v24
    return (v24 & genFn0())
}

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 255
    var v3 bool = ((true && (404 > 203)) && true)
    v4 := []int{4294967295}
    var v5 bool = (835 > 853)
    v4 = append(v4, -2147483648)
    if v5 {
        v2 -= v2
        for fi25 := 0; fi25 < 12; fi25++ {
            fmt.Println(fi25, fi25)
            for fi26 := 0; fi26 < 18; fi26++ {
                var v27 int = genFn0()
                _ = fi26
                _ = v27
            }
            _ = fi25
        }
        v5 = v3
    }
    for fi28 := 0; fi28 < 1; fi28++ {
        if (0 <= (v2 ^ v2)) {
            for fi29 := 0; fi29 < 13; fi29++ {
                var v30 int32 = 432
                _ = fi29
                _ = v30
            }
            switch fi28 {
            case -2147483648:
                _ = len(v4)
            }
        }
        _ = fi28
    }
    v31 := GenS3{F0: 125, F1: ""}
    for fi32 := 0; fi32 < 9; fi32++ {
        if len(v4) > 0 {
            v4[0] = fi32
        }
        for fi33 := 0; fi33 < 6; fi33++ {
            v5 = (1 <= 4294967295)
            delete(v1, (fi33 * (fi33 / ((128) | 1))))
            fmt.Println(fi32)
            switch "" {
            case "\n\t":
                switch (fi33 & genFn1(2)) {
                case 127:
                    var v34 bool = v3
                    _ = v34
                }
            default:
                v35 := GenS2{F0: int32(v2)}
                fmt.Println(fi32, v5)
                var v36 string = "\n\t"
                _ = v35
                _ = v36
            }
            _ = fi33
        }
        _ = len(v4)
        _ = fi32
    }
    v37 := GenS1{F0: 47, F1: 1}
    if true {
        for fi38 := 0; fi38 < 9; fi38++ {
            switch "" {
            case "hello":
                if v5 {
                    var v39 int32 = 638
                    _ = v39
                }
            default:
                if ("a" == "\n\t") {
                    var v40 float32 = ((0.1 + float32(-2.25)) - -2.25)
                    _ = v40
                }
            }
            v2 = (v2 - v2)
            _ = fi38
        }
        for fi41 := 0; fi41 < 15; fi41++ {
            if (v3 && ((-2.25 - 1.5) == (0.1 - (1.5 * 0.1)))) {
                v5 = (42 <= 42)
                v42 := GenS3{F0: (v2 | v2), F1: "foo bar"}
                fmt.Println(v2, v5)
                _ = v42
            }
            v5 = v5
            var v43 uint64 = 2
            _ = fi41
            _ = v43
        }
    }
    v5 = false
    v5 = (-32768 == 42)
    var v44 float64 = 1.5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v37
    _ = v4
    _ = v44
    _ = v5
}
