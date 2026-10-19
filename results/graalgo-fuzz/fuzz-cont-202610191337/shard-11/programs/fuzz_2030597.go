package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 uint16 }
func genFn0() string {
    for fi5 := 0; fi5 < 15; fi5++ {
        switch "foo bar" {
        case "hello":
            if (42 == 344) {
                for fi6 := 0; fi6 < 13; fi6++ {
                    var v7 uint8 = 128
                    _ = fi6
                    _ = v7
                }
            } else {
                var v8 float64 = (1.5 * 1.5)
                _ = v8
            }
        case "\n\t":
            var v9 int16 = 306
            _ = v9
        }
        if (false || (true || false)) {
            if (2147483647 != -359) {
                var v10 bool = (fi5 < (fi5 << ((213) & 63)))
                fmt.Println(v10)
                if (42 >= -856) {
                    var v11 uint32 = 4294967295
                    _ = v11
                } else {
                    for fi12 := 0; fi12 < 13; fi12++ {
                        fmt.Println(fi12, v10)
                        _ = fi12
                    }
                }
                _ = v10
            }
        }
        _ = fi5
    }
    var v13 int16 = -1
    if false {
        v14 := []int{-747}
        v15 := GenS1{F0: "hello", F1: 997}
        var v16 uint8 = 174
        _ = v14
        _ = v15
        _ = v16
    } else {
        if false {
            for fi17 := 0; fi17 < 17; fi17++ {
                v18 := GenS1{F0: ("hello" + "a"), F1: 1}
                fmt.Println(fi17)
                _ = fi17
                _ = v18
            }
            v19 := GenS1{F0: "hello", F1: 65535}
            _ = v19
        } else {
            for fi20 := 0; fi20 < 9; fi20++ {
                for fi21 := 0; fi21 < 20; fi21++ {
                    if (true || (fi21 >= 32768)) {
                        var v22 bool = (!true)
                        _ = v22
                    }
                    _ = fi21
                }
                v23 := GenS1{F0: "hello", F1: 1}
                _ = fi20
                _ = v23
            }
        }
        for fi24 := 0; fi24 < 20; fi24++ {
            var v25 uint = 0
            switch (fi24 / ((fi24) | 1)) {
            case 814:
                v25 += v25
            default:
                if (((1.5 / 1.0) * 0.0) != 1.5) {
                    var v26 bool = (2 <= 42)
                    _ = v26
                }
            }
            _ = fi24
            _ = v25
        }
        var v27 float64 = 3.14159
        _ = v27
    }
    _ = v13
    return "\n\t"
}
func genFn1(p0 int32, p1 int64) string {
    p1 -= 42
    v28 := GenS1{F0: "\n\t", F1: 196}
    _ = p0
    _ = p1
    _ = v28
    return ""
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]int)
    v3 := make(map[int]int)
    var v4 float32 = 0.1
    if (-109 > 1000) {
        v4 += (0.1 + v4)
        var v29 uint32 = 2147483647
        _ = v29
    } else {
        fmt.Println(v4)
    }
    if (true && (!true)) {
        v30 := GenS1{F0: ("hello" + "\n\t"), F1: 42}
        v31 := []int{9223372036854775807, 9223372036854775807, 0}
        _ = v30
        _ = v31
    } else {
        fmt.Println(v4)
        var v32 int = 4294967295
        _ = v32
    }
    var v33 int64 = 312
    _ = len(v1)
    var v34 bool = true
    fmt.Println(v33, v33)
    var v35 int = 2147483647
    for fi36 := 0; fi36 < 1; fi36++ {
        v37 := GenS1{F0: "foo bar", F1: 0}
        v38 := GenS2{F0: 65535}
        _ = fi36
        _ = v37
        _ = v38
    }
    v39 := []int{32767, (-(v35)), v35, v35}
    fmt.Println(v4, v33, v34)
    if (v4 <= v4) {
        v33 += (v33 * (v33 >> ((0) & 63)))
        v40 := GenS1{F0: ("\n\t" + "foo bar"), F1: 543}
        _ = v40
    }
    v4 += float32(1.5)
    var v41 float64 = float64((-2.25 + v4))
    fmt.Println(v4, v4)
    v41 += 1e10
    v41 -= v41
    v34 = v34
    v33 = v33
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v39
    _ = v4
    _ = v41
}
