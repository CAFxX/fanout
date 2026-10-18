package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 uint16 }
func genFn0(p0 string) bool {
    var v6 float32 = 0.0
    var v7 int = 742
    v8 := []int{}
    p0 = ""
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return true
}
func genFn1(p0 uint16, p1 int64) bool {
    switch 1 {
    case 209:
        v9 := GenS1{F0: 4294967295}
        _ = v9
    case 32767:
        var v10 string = "\n\t"
        _ = v10
    default:
        var v11 uint16 = p0
        _ = v11
    }
    v12 := GenS2{F0: 63}
    var v13 string = "hello"
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return false
}

func main() {
    var v1 bool = false
    var v2 int = 1
    var v3 uint64 = 18446744073709551615
    v4 := make(map[string]int)
    var v5 bool = v1
    var v14 int = v2
    v15 := []int{v14, 0, v2}
    if false {
        for fi16 := 0; fi16 < 3; fi16++ {
            if v5 {
                _ = len(v4)
                v17 := GenS2{F0: 2}
                _ = v17
            } else {
                var v18 string = ((("hello" + ("\n\t" + ("\n\t" + ("a" + "")))) + "a") + "\n\t")
                _ = v18
            }
            var v19 float64 = (-2.25 + 3.14159)
            _ = fi16
            _ = v19
        }
        v14 = (-(v2))
        var v20 float64 = 0.1
        _ = v20
    } else {
        for fi21 := 0; fi21 < 9; fi21++ {
            v3 += (^(v3))
            _ = fi21
        }
    }
    v22 := GenS1{F0: 2}
    if v1 {
        v23 := v4["foo bar"]
        if true {
            if (!(9223372036854775807 <= -9223372036854775808)) {
                var v24 uint8 = 42
                v25 := GenS3{F0: -128, F1: 32767}
                if len(v15) > 0 {
                    v15[0] = (v2 + (v23 & (v23 << ((2) & 63))))
                }
                _ = v24
                _ = v25
            }
            if true {
                v3 += v3
                var v26 string = ""
                if (v1 || v5) {
                    if false {
                        var v27 int64 = 42
                        _ = v27
                    } else {
                        _ = len(v15)
                        fmt.Println(v14, v23, v26)
                    }
                    var v28 int32 = int32(v3)
                    v23 = (v14 + v23)
                    _ = v28
                } else {
                    if v5 {
                        v3 = v3
                        if (808 > int32(v23)) {
                            v26 = v26
                            var v29 float64 = float64(1.5)
                            _ = v29
                        }
                    }
                    v30 := GenS1{F0: 42}
                    _ = v30
                }
                _ = v26
            }
        }
        _ = v23
    } else {
        delete(v4, "hello")
        v4[""] = v14
    }
    v2 = (v2 | (v14 >> ((323) & 63)))
    v31 := GenS2{F0: 128}
    for fi32 := 0; fi32 < 10; fi32++ {
        v2 = v2
        switch 4294967295 {
        case 176:
            v3 += v3
        case 127:
            var v33 int64 = 519
            _ = v33
        }
        _ = fi32
    }
    var v34 int64 = 1
    v35 := GenS2{F0: 0}
    v5 = (18446744073709551615 <= 4294967295)
    if (v14 != -889) {
        var v36 float64 = (((3.14159 + 3.14159) + -2.25) + 1.5)
        _ = v36
    }
    v3 = 18446744073709551615
    _ = len(v15)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v3
    _ = v31
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
