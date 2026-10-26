package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint32 }
type GenS2 struct { F0 bool; F1 string; F2 uint8 }
type GenS3 struct { F0 string }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]int)
    v3 := []int{2147483647, 127, -1}
    _ = len(v2)
    v4 := GenS1{F0: "foo bar", F1: (true && true), F2: 65535}
    _ = len(v2)
    var v5 uint16 = 65535
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 bool = true
        var v8 uint64 = 1
        if ("a" == "") {
            for fi9 := 0; fi9 < 5; fi9++ {
                v1[("a" + "\n\t")] = v7
                _ = fi9
            }
        } else {
            switch fi6 {
            case -332:
                var v10 int8 = -128
                _ = v10
            case -2147483648:
                for fi11 := 0; fi11 < 11; fi11++ {
                    v5 = v5
                    _ = fi11
                }
            case 0:
                v3 = append(v3, fi6)
            }
            if v7 {
                v5 = v5
                v5 = (v5 % ((v5) | 1))
            } else {
                if (0.0 >= 1.5) {
                    v7 = v7
                }
            }
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    fmt.Println(v5)
    var v12 int8 = 42
    v5 = (v5 - v5)
    switch 9223372036854775807 {
    case 128:
        delete(v1, "a")
    case 79:
        v13 := GenS2{F0: (false && (1.5 == float32(1e10))), F1: "", F2: 110}
        _ = v13
    default:
        for fi14 := 0; fi14 < 16; fi14++ {
            var v15 bool = (127 == 1)
            for fi16 := 0; fi16 < 20; fi16++ {
                var v17 float64 = 1.5
                fmt.Println(v12, v17, v15)
                if (-671 == 1) {
                    fmt.Println(fi16, v17, v12)
                }
                _ = fi16
                _ = v17
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi14
            _ = v15
        }
    }
    if true {
        v5 = v5
        fmt.Println(v12, v12)
        var v18 bool = false
        _ = v18
    } else {
        v5 -= (v5 & (v5 << ((18446744073709551615) & 63)))
        var v19 uint8 = 42
        _ = v19
    }
    v20 := GenS1{F0: "\n\t", F1: (("\n\t" + "") < "hello"), F2: 929}
    fmt.Println(v12, v5)
    v12 = (v12 - v12)
    for fi21 := 0; fi21 < 15; fi21++ {
        v22 := GenS3{F0: "a"}
        _ = fi21
        _ = v22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
