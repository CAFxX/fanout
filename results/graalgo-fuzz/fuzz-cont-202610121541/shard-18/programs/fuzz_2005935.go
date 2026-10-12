package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 uint32 }
func genFn0(p0 int16) uint64 {
    if true {
        var v4 bool = true
        p0 = 1000
        v5 := GenS1{F0: ((p0 != p0) && ((979 <= 0) && (1e10 <= -2.25)))}
        _ = v4
        _ = v5
    }
    var v6 int = -9223372036854775808
    _ = p0
    _ = v6
    return 700
}
func genFn1(p0 uint64, p1 uint32, p2 string) uint8 {
    for fi7 := 0; fi7 < 2; fi7++ {
        v8 := []int{}
        p2 = "hello"
        fmt.Println(fi7, p1, p1)
        _ = fi7
        _ = v8
    }
    v9 := []int{139, 127}
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 127
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]string)
    v3 := make(map[string]bool)
    var v10 uint8 = 41
    v10 += v10
    for fi11 := 0; fi11 < 1; fi11++ {
        for fi12 := 0; fi12 < 12; fi12++ {
            switch (fi12 & 32767) {
            case 0:
                if true {
                    v13 := GenS1{F0: false}
                    v14 := GenS1{F0: false}
                    fmt.Println(fi11, fi11, v10)
                    _ = v13
                    _ = v14
                } else {
                    var v15 int8 = 70
                    v10 = 211
                    var v16 uint64 = 42
                    fmt.Println(v16, fi12)
                    v17 := GenS1{F0: (!((!false) || (32767 >= 323)))}
                    _ = v15
                    _ = v16
                    _ = v17
                }
            case 4294967295:
                var v18 uint32 = 539
                _ = v18
            case 484:
                switch ("a" + "a") {
                case "\n\t":
                    if (true && (42 <= 228)) {
                        v19 := []int{(fi11 | 9223372036854775807), fi12}
                        var v20 uint16 = 242
                        _ = v19
                        _ = v20
                    } else {
                        var v21 string = "foo bar"
                        if true {
                            switch v21 {
                            case "hello":
                                v10 -= (v10 | v10)
                            }
                        }
                        _ = v21
                    }
                }
            default:
                _ = len(v2)
            }
            fmt.Println(v10, fi12, v10)
            _ = fi12
        }
        _ = fi11
    }
    v10 += 197
    if true {
        v10 = (v10 & (v10 + (v10 % (((v10 << ((1) & 63))) | 1))))
        if ((789 == 4294967295) || true) {
            v22 := GenS1{F0: (391 <= 42)}
            _ = v22
        } else {
            if false {
                v10 = v10
            } else {
                v3[(((("" + ("" + ((("\n\t" + "hello") + "hello") + "\n\t"))) + "a") + "a") + "")] = false
                switch 0 {
                case 0:
                    v23 := []int{-2147483648, 898, -674, 1}
                    _ = v23
                case 42:
                    if false {
                        if (0 != 848) {
                            var v24 string = "\n\t"
                            switch "hello" {
                            case "":
                                v25 := GenS2{F0: 255, F1: 32767}
                                _ = v25
                            case "\n\t":
                                v26 := GenS1{F0: (0.1 > 0.1)}
                                _ = v26
                            case "foo bar":
                                if (0.1 != (0.1 + 1.5)) {
                                    v27 := []int{1, 0, 127, 32767}
                                    var v28 int = -9223372036854775808
                                    _ = v27
                                    _ = v28
                                }
                            }
                            var v29 bool = true
                            _ = v24
                            _ = v29
                        }
                        var v30 int64 = 1
                        v30 -= v30
                        _ = v30
                    } else {
                        fmt.Println(v10)
                        fmt.Println(v10)
                        var v31 float32 = 3.14159
                        _ = v31
                    }
                default:
                    v10 = genFn1(4294967295, 42, ("" + ("\n\t" + "\n\t")))
                }
            }
            for fi32 := 0; fi32 < 10; fi32++ {
                v33 := GenS3{F0: 4294967295, F1: 9223372036854775807}
                v10 = (v10 - (v10 + v10))
                _ = fi32
                _ = v33
            }
            v34 := []int{-142, 128}
            _ = v34
        }
        var v35 uint64 = 0
        _ = v35
    } else {
        switch 9223372036854775807 {
        case -312:
            v2[289] = ("foo bar" + "\n\t")
        case 0:
            fmt.Println(v10)
        case -593:
            v36 := []int{-2147483648, 97, -2147483648}
            _ = v36
        }
        var v37 float32 = ((((0.0 * -2.25) * (3.14159 / 1.0)) * 0.0) - ((0.0 + 1.5) * -2.25))
        if true {
            var v38 uint64 = genFn0(0)
            _ = v38
        } else {
            switch 0 {
            case -9223372036854775808:
                fmt.Println(v10, v37)
            }
        }
        _ = v37
    }
    v39 := GenS4{F0: 4294967295}
    var v40 string = "\n\t"
    var v41 uint32 = 807
    var v42 int64 = 342
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v39
    _ = v40
    _ = v41
    _ = v42
}
