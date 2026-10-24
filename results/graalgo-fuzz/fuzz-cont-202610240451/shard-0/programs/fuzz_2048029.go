package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 int8 }
func genFn0(p0 int) int64 {
    v6 := GenS1{F0: 1, F1: 1}
    var v7 string = ("foo bar" + "")
    v7 = ("" + (((v7 + (v7 + (v7 + "a"))) + v7) + v7))
    v8 := []int{p0}
    fmt.Println(p0)
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 uint16, p1 bool) int {
    p1 = true
    switch ("foo bar" + "\n\t") {
    case "hello":
        switch "" {
        case "a":
            for fi9 := 0; fi9 < 2; fi9++ {
                if (!((42 > 2) || (-1 != 902))) {
                    v10 := GenS1{F0: -1, F1: 1}
                    _ = v10
                }
                if p1 {
                    v11 := GenS1{F0: -586, F1: 1000}
                    _ = v11
                } else {
                    v12 := []int{(fi9 | (fi9 - (fi9 / (((^(fi9))) | 1)))), 127, fi9}
                    _ = v12
                }
                _ = fi9
            }
        case "":
            for fi13 := 0; fi13 < 14; fi13++ {
                if p1 {
                    v14 := GenS1{F0: -1, F1: 0}
                    p0 = 42
                    _ = v14
                }
                _ = fi13
            }
        case "\n\t":
            if p1 {
                for fi15 := 0; fi15 < 13; fi15++ {
                    fmt.Println(fi15, p1)
                    fmt.Println(p0, p0)
                    var v16 string = "hello"
                    _ = fi15
                    _ = v16
                }
                switch "foo bar" {
                case "a":
                    for fi17 := 0; fi17 < 18; fi17++ {
                        var v18 uint64 = 42
                        p1 = (p1 && p1)
                        switch ("\n\t" + (("\n\t" + "foo bar") + "\n\t")) {
                        case "hello":
                            var v19 int16 = 0
                            _ = v19
                        case "a":
                            var v20 int32 = 942
                            _ = v20
                        case "":
                            var v21 int32 = -1
                            _ = v21
                        }
                        _ = fi17
                        _ = v18
                    }
                case "":
                    var v22 int16 = 1000
                    _ = v22
                }
            }
        }
    case "foo bar":
        var v23 bool = (true || ("\n\t" == (("hello" + "a") + "a")))
        _ = v23
    default:
        for fi24 := 0; fi24 < 9; fi24++ {
            var v25 uint32 = 504
            var v26 uint16 = 1
            var v27 int = fi24
            _ = fi24
            _ = v25
            _ = v26
            _ = v27
        }
    }
    var v28 int64 = genFn0(597)
    v29 := GenS1{F0: (^(v28)), F1: 1000}
    p0 = 32767
    _ = p0
    _ = p1
    _ = v28
    _ = v29
    return -1
}

func main() {
    var v1 string = ("foo bar" + "a")
    var v2 int32 = 2147483647
    var v3 uint64 = 533
    var v4 int8 = 1
    v5 := make(map[int]int)
    var v30 bool = false
    var v31 uint = 488
    var v32 int16 = 0
    var v33 uint = 255
    fmt.Println(v30, v2)
    var v34 int8 = v4
    for fi35 := 0; fi35 < 20; fi35++ {
        for fi36 := 0; fi36 < 11; fi36++ {
            if v30 {
                v34 -= (v34 + (v34 ^ int8(v32)))
            }
            var v37 uint8 = 128
            v38 := GenS1{F0: 0, F1: v32}
            _ = fi36
            _ = v37
            _ = v38
        }
        _ = fi35
    }
    v4 = (v4 ^ (v4 ^ v4))
    switch v1 {
    case "\n\t":
        for fi39 := 0; fi39 < 16; fi39++ {
            v40 := []int{fi39, fi39}
            v5[(fi39 - (fi39 % ((fi39) | 1)))] = (-(fi39))
            _ = fi39
            _ = v40
        }
    case "a":
        for fi41 := 0; fi41 < 4; fi41++ {
            var v42 int = fi41
            v4 -= v4
            var v43 uint8 = 2
            _ = fi41
            _ = v42
            _ = v43
        }
    case "hello":
        v5[2147483647] = 9223372036854775807
    default:
        v33 = v31
    }
    for fi44 := 0; fi44 < 15; fi44++ {
        v45 := GenS2{F0: (v4 ^ 122)}
        var v46 float64 = (((1e10 + 3.14159) + float64(((0.1 + 0.1) - -2.25))) - -2.25)
        for fi47 := 0; fi47 < 10; fi47++ {
            if (false || (v2 > (v2 >> ((2) & 63)))) {
                v4 -= (v34 * (v34 << (((v31 & 2)) & 63)))
                var v48 float32 = ((1.5 / (0.0 - 3.14159)) / (3.14159 - 1e10))
                v1 = v1
                _ = v48
            }
            _ = fi47
        }
        _ = fi44
        _ = v45
        _ = v46
    }
    fmt.Println(v3, v4)
    if (v1 >= "") {
        v34 -= -1
    }
    for fi49 := 0; fi49 < 20; fi49++ {
        fmt.Println(v31, v33, v3)
        v30 = v30
        _ = fi49
    }
    delete(v5, -204)
    v50 := []int{2147483647, 128}
    fmt.Println(v31, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v50:", len(v50), v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v50
}
