package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
func genFn0(p0 uint64) int32 {
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 bool = ((fi7 ^ (fi7 + (fi7 >> ((42) & 63)))) == fi7)
        v9 := []int{}
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := []int{}
    _ = p0
    _ = v10
    return 42
}
func genFn1(p0 int8) int64 {
    for fi11 := 0; fi11 < 19; fi11++ {
        switch 127 {
        case -9223372036854775808:
            var v12 int32 = -1
            _ = v12
        case -2147483648:
            if true {
                p0 = 127
                switch "hello" {
                case "\n\t":
                    if true {
                        p0 -= -128
                    } else {
                        var v13 uint8 = 0
                        _ = v13
                    }
                default:
                    p0 -= p0
                }
                var v14 uint = 18446744073709551615
                _ = v14
            } else {
                v15 := []int{fi11, 1, 42}
                if false {
                    fmt.Println(fi11)
                    for fi16 := 0; fi16 < 3; fi16++ {
                        _ = len(v15)
                        var v17 int8 = -128
                        switch (^(fi11)) {
                        case 1:
                            p0 = 1
                        case 353:
                            var v18 bool = ((3.14159 / 0.1) >= (1.5 - ((-2.25 - 3.14159) * float64(3.14159))))
                            _ = v18
                        }
                        _ = fi16
                        _ = v17
                    }
                }
                _ = v15
            }
        }
        fmt.Println(fi11)
        _ = fi11
    }
    p0 = p0
    if true {
        var v19 float64 = 1.5
        for fi20 := 0; fi20 < 7; fi20++ {
            v21 := GenS1{F0: 2147483648, F1: 4294967295}
            var v22 uint32 = 65535
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = v19
    }
    v23 := GenS1{F0: 4294967295, F1: 777}
    _ = p0
    _ = v23
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = 127
    var v3 float64 = float64(3.14159)
    var v4 string = ""
    v5 := make(map[int]int)
    v6 := []string{"hello", "hello", v4, v4}
    fmt.Println(v4, v2, v2)
    v4 = ("" + "foo bar")
    v4 = (v4 + v4)
    v2 += v2
    v24 := GenS1{F0: 1, F1: 42}
    switch "" {
    case "\n\t":
        v3 = v3
    case "":
        v3 = ((((v3 * -2.25) - (v3 * (-2.25 / 1.5))) + 1.5) - -2.25)
    case "a":
        v2 += v2
    }
    for fi25 := 0; fi25 < 3; fi25++ {
        v26 := GenS1{F0: 0, F1: 597}
        switch fi25 {
        case -9223372036854775808:
            v27 := GenS1{F0: 526, F1: 254}
            _ = v27
        case 128:
            v2 = (v2 << ((1) & 63))
        }
        _ = fi25
        _ = v26
    }
    v28 := GenS1{F0: 981, F1: 736}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
