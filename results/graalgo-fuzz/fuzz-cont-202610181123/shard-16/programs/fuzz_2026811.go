package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 int32; F2 uint32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint; F1 string }
func genFn0(p0 int32, p1 uint, p2 int32) bool {
    var v4 int16 = 1
    v5 := []int{32768, int(p2), 128, 2147483647}
    fmt.Println(p2, v4, v4)
    switch 9223372036854775807 {
    case 659:
        v6 := GenS1{F0: (9223372036854775807 <= -9223372036854775808), F1: 4294967295, F2: 255}
        _ = v6
    case -9223372036854775808:
        switch -9223372036854775808 {
        case 4294967295:
            if (!((true && (1e10 == 3.14159)) || (0.1 == -2.25))) {
                for fi7 := 0; fi7 < 9; fi7++ {
                    v8 := GenS2{F0: 1}
                    for fi9 := 0; fi9 < 9; fi9++ {
                        v5 = append(v5, fi9)
                        switch (fi9 << ((27) & 63)) {
                        case 1:
                            v4 = v4
                        case 4294967295:
                            if false {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                switch (("foo bar" + ("\n\t" + "hello")) + "foo bar") {
                                case "a":
                                    v10 := GenS3{F0: ("\n\t" + ""), F1: (p0 / ((p2) | 1)), F2: 4294967295}
                                    _ = v10
                                case "":
                                    p2 += p2
                                }
                            }
                        default:
                            v11 := GenS4{F0: "foo bar"}
                            _ = v11
                        }
                        _ = fi9
                    }
                    _ = fi7
                    _ = v8
                }
                v12 := GenS5{F0: 4294967295}
                var v13 float64 = 1e10
                _ = v12
                _ = v13
            } else {
                var v14 uint8 = 2
                if (1.5 != 0.1) {
                    v15 := GenS6{F0: 1, F1: ""}
                    fmt.Println(p1)
                    _ = v15
                }
                _ = len(v5)
                _ = v14
            }
        case 32767:
            p2 += -714
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return false
}

func main() {
    var v1 uint8 = 234
    var v2 bool = false
    var v3 bool = true
    if (!(443 >= 255)) {
        v16 := []int{127, -1, 4294967295, 409}
        _ = v16
    }
    var v17 int16 = 42
    v17 = (^(v17))
    v1 -= v1
    v18 := GenS4{F0: ""}
    v19 := []int{-1}
    var v20 int16 = v17
    _ = len(v19)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
