package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 int64; F1 string; F2 string }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 string; F1 string; F2 uint64 }
func genFn0(p0 string, p1 bool, p2 string) string {
    if p1 {
        fmt.Println(p2, p2, p1)
        for fi7 := 0; fi7 < 16; fi7++ {
            for fi8 := 0; fi8 < 9; fi8++ {
                var v9 int = -1
                _ = fi8
                _ = v9
            }
            for fi10 := 0; fi10 < 3; fi10++ {
                switch "" {
                case "":
                    v11 := []int{fi10, fi10}
                    p0 = p2
                    var v12 uint32 = 1
                    _ = v11
                    _ = v12
                case "\n\t":
                    p0 = p2
                }
                v13 := GenS1{F0: int64(fi7), F1: p2}
                _ = fi10
                _ = v13
            }
            _ = fi7
        }
        switch p2 {
        case "":
            fmt.Println(p1)
        case "\n\t":
            v14 := GenS2{F0: -1, F1: p2, F2: p2}
            _ = v14
        case "a":
            p2 = (p2 + "")
        }
    } else {
        fmt.Println(p1, p0, p1)
        var v15 uint = 749
        p0 = p0
        _ = v15
    }
    if false {
        for fi16 := 0; fi16 < 14; fi16++ {
            p1 = false
            _ = fi16
        }
        for fi17 := 0; fi17 < 1; fi17++ {
            var v18 float64 = -2.25
            _ = fi17
            _ = v18
        }
        v19 := []int{2147483647, 32768, 128}
        _ = v19
    } else {
        for fi20 := 0; fi20 < 10; fi20++ {
            v21 := []int{fi20}
            _ = fi20
            _ = v21
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return p0
}

func main() {
    var v1 float64 = float64(3.14159)
    var v2 int8 = 0
    var v3 string = "a"
    var v4 uint16 = 255
    var v5 uint8 = 1
    v6 := make(map[uint32]string)
    if ((v2 < v2) || true) {
        for fi22 := 0; fi22 < 2; fi22++ {
            for fi23 := 0; fi23 < 18; fi23++ {
                var v24 int16 = -1
                var v25 uint32 = 2147483648
                delete(v6, (v25 + (v25 & v25)))
                _ = fi23
                _ = v24
                _ = v25
            }
            v26 := []int{1, -2147483648, fi22, fi22}
            _ = fi22
            _ = v26
        }
        fmt.Println(v3)
    } else {
        var v27 uint32 = uint32(v5)
        _ = v27
    }
    var v28 int = 42
    fmt.Println(v3, v5)
    v29 := GenS1{F0: -1, F1: (((v3 + "foo bar") + v3) + v3)}
    v1 -= float64(-2.25)
    var v30 uint8 = 255
    var v31 float64 = (0.1 - (((-2.25 * 0.0) - 3.14159) / 1e10))
    if (uint64(v30) == 333) {
        var v32 uint64 = 0
        v5 += 113
        v32 -= 1
        _ = v32
    } else {
        var v33 string = v3
        switch v3 {
        case "\n\t":
            v34 := []int{(v28 | (v28 | v28)), (v28 | v28), v28, int(v2)}
            _ = v34
        case "hello":
            switch 128 {
            case 42:
                switch -987 {
                case 32767:
                    var v35 int8 = v2
                    _ = v35
                case 1:
                    for fi36 := 0; fi36 < 8; fi36++ {
                        for fi37 := 0; fi37 < 7; fi37++ {
                            v38 := GenS3{F0: (v3 + ((v33 + v3) + v33)), F1: 347}
                            _ = fi37
                            _ = v38
                        }
                        _ = fi36
                    }
                case 42:
                    var v39 uint32 = 868
                    _ = v39
                default:
                    for fi40 := 0; fi40 < 16; fi40++ {
                        v41 := []int{int(v2), (^(v28)), int(v4)}
                        if false {
                            var v42 float32 = 1e10
                            _ = v42
                        }
                        if len(v41) > 0 {
                            v41[0] = (fi40 - -350)
                        }
                        _ = fi40
                        _ = v41
                    }
                }
            }
        default:
            v43 := GenS4{F0: v3, F1: ((("foo bar" + ((v33 + (v33 + ((v3 + v33) + v3))) + "a")) + v3) + "hello"), F2: 2}
            _ = v43
        }
        var v44 uint32 = 435
        _ = v33
        _ = v44
    }
    v28 -= int(v30)
    v45 := []int{(v28 | 42), (v28 % (((v28 % ((-2147483648) | 1))) | 1))}
    v4 = (v4 ^ (v4 << ((65535) & 63)))
    v2 -= (v2 - v2)
    for fi46 := 0; fi46 < 5; fi46++ {
        var v47 bool = false
        _ = fi46
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v45
    _ = v5
    _ = v6
}
