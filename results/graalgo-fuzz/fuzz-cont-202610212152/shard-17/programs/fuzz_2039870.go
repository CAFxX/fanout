package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 int, p1 uint8) string {
    p1 += p1
    v4 := GenS1{F0: -575, F1: 2147483647}
    v5 := GenS1{F0: 0, F1: -1}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return ""
}
func genFn1() int8 {
    for fi6 := 0; fi6 < 6; fi6++ {
        fmt.Println(fi6)
        _ = fi6
    }
    v7 := GenS1{F0: -1, F1: 0}
    v8 := []int{}
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]int)
    var v3 float32 = (3.14159 + 1e10)
    v3 += ((-2.25 + v3) / 1.5)
    var v9 int8 = 50
    var v10 int16 = -225
    switch -9223372036854775808 {
    case 9223372036854775807:
        for fi11 := 0; fi11 < 10; fi11++ {
            if true {
                for fi12 := 0; fi12 < 15; fi12++ {
                    switch fi12 {
                    case 4294967295:
                        v13 := []int{(^(fi11)), (fi11 | (fi12 * fi11)), (fi11 - fi12)}
                        _ = v13
                    default:
                        v10 = (^(v10))
                    }
                    _ = fi12
                }
            }
            var v14 bool = false
            _ = fi11
            _ = v14
        }
    case 712:
        if ((v9 * v9) <= 9) {
            if (v9 <= v9) {
                if (!("a" < "a")) {
                    if true {
                        var v15 string = "\n\t"
                        _ = v15
                    } else {
                        v10 = (v10 % (((^(v10))) | 1))
                        for fi16 := 0; fi16 < 19; fi16++ {
                            var v17 int32 = 42
                            v9 -= v9
                            _ = fi16
                            _ = v17
                        }
                    }
                }
                var v18 uint16 = 32767
                _ = v18
            }
            switch "a" {
            case "hello":
                var v19 int16 = 1000
                _ = v19
            case "foo bar":
                switch -465 {
                case 893:
                    for fi20 := 0; fi20 < 18; fi20++ {
                        delete(v2, "\n\t")
                        for fi21 := 0; fi21 < 14; fi21++ {
                            fmt.Println(fi21)
                            v22 := []int{}
                            _ = fi21
                            _ = v22
                        }
                        _ = fi20
                    }
                case 32768:
                    v23 := GenS2{F0: 2}
                    _ = v23
                }
            }
        }
    }
    for fi24 := 0; fi24 < 16; fi24++ {
        v3 += (v3 - (1.5 * float32(float64(v3))))
        _ = len(v2)
        fmt.Println(fi24, v3, v3)
        _ = fi24
    }
    v25 := GenS1{F0: -9223372036854775808, F1: -2147483648}
    for fi26 := 0; fi26 < 11; fi26++ {
        var v27 int8 = v9
        _ = fi26
        _ = v27
    }
    v9 -= genFn1()
    v28 := GenS3{F0: 0}
    v29 := []int{128, 2147483647, 32768}
    var v30 int64 = 42
    for fi31 := 0; fi31 < 19; fi31++ {
        fmt.Println(v3)
        v30 -= 42
        _ = fi31
    }
    delete(v2, ((("\n\t" + "") + ("\n\t" + "foo bar")) + ""))
    for fi32 := 0; fi32 < 2; fi32++ {
        var v33 uint64 = 207
        _ = fi32
        _ = v33
    }
    var v34 float32 = (v3 + v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v9
}
