package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint16 }
func genFn0(p0 string, p1 uint32) int {
    v7 := GenS1{F0: (((p0 + "a") + (("" + "\n\t") + p0)) + (p0 + "")), F1: p1}
    v8 := []int{-9223372036854775808}
    switch (p0 + p0) {
    case "\n\t":
        fmt.Println(p0, p1)
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 uint16 = 65535
    var v2 bool = true
    var v3 int64 = 1
    var v4 bool = ((3.14159 / 1e10) == (0.0 - 0.0))
    var v5 float32 = 0.1
    var v6 uint8 = 149
    switch "a" {
    case "foo bar":
        v9 := GenS1{F0: ("foo bar" + "hello"), F1: 2147483647}
        _ = v9
    case "a":
        var v10 uint16 = (v1 % ((42) | 1))
        _ = v10
    default:
        for fi11 := 0; fi11 < 17; fi11++ {
            v12 := GenS1{F0: "", F1: 902}
            _ = fi11
            _ = v12
        }
    }
    v5 = v5
    v3 -= v3
    fmt.Println(v6, v2)
    switch -497 {
    case -1:
        v5 = v5
    case -256:
        v1 -= v1
    case 32768:
        v13 := GenS2{F0: 0}
        _ = v13
    }
    v14 := []int{genFn0(("foo bar" + "foo bar"), 2147483648), -1, 2147483647}
    var v15 uint16 = v1
    if false {
        for fi16 := 0; fi16 < 16; fi16++ {
            for fi17 := 0; fi17 < 12; fi17++ {
                var v18 uint64 = 1
                if len(v14) > 0 {
                    v14[0] = (fi17 | genFn0("hello", 357))
                }
                var v19 float64 = 1e10
                if true {
                    v20 := GenS2{F0: v18}
                    _ = v20
                } else {
                    if len(v14) > 0 {
                        v14[0] = fi16
                    }
                    switch ("a" + ("\n\t" + "foo bar")) {
                    case "foo bar":
                        v4 = v4
                    case "a":
                        v21 := GenS2{F0: uint64(v1)}
                        _ = v21
                    case "\n\t":
                        v14 = append(v14, genFn0("hello", 15))
                    default:
                        var v22 int = 776
                        _ = v22
                    }
                    v14 = append(v14, fi17)
                }
                switch "" {
                case "\n\t":
                    if (v5 < (v5 * (3.14159 / v5))) {
                        fmt.Println(fi17, v15, v1)
                        if ((668 >= (v1 * (v15 % ((42) | 1)))) && v2) {
                            v23 := GenS1{F0: "a", F1: 1}
                            _ = v23
                        }
                        var v24 uint8 = v6
                        _ = v24
                    } else {
                        var v25 int = fi17
                        _ = v25
                    }
                case "foo bar":
                    switch fi17 {
                    case 405:
                        v14 = append(v14, genFn0("", 2147483647))
                    case -9223372036854775808:
                        v19 = 0.1
                    case 42:
                        v26 := GenS3{F0: 228}
                        _ = v26
                    default:
                        var v27 int8 = 64
                        _ = v27
                    }
                }
                _ = fi17
                _ = v18
                _ = v19
            }
            v28 := GenS4{F0: (^(v15))}
            _ = fi16
            _ = v28
        }
        if (float64(0.0) <= (0.0 * 3.14159)) {
            fmt.Println(v6)
        } else {
            var v29 int16 = 950
            _ = v29
        }
    }
    v30 := GenS2{F0: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
