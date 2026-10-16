package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 string }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 int64; F1 int; F2 int32 }
type GenS4 struct { F0 uint64; F1 int8 }
func genFn0() uint32 {
    v5 := []int{42, -818}
    var v6 bool = (81 >= -5)
    v7 := GenS1{F0: 42, F1: (("a" + "\n\t") + "hello"), F2: "hello"}
    var v8 float64 = (0.1 / 1.5)
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 int8 = 0
    v2 := make(map[string]int)
    var v3 int16 = -91
    var v4 int64 = 42
    var v9 string = "\n\t"
    v10 := GenS1{F0: 4294967295, F1: v9, F2: v9}
    v11 := GenS1{F0: 4294967295, F1: ("\n\t" + "foo bar"), F2: "\n\t"}
    v1 -= 100
    switch 945 {
    case 0:
        v3 -= (v3 | (v3 - (v3 ^ (v3 ^ v3))))
    case -1:
        var v12 uint = 1
        _ = v12
    case 1:
        if false {
            var v13 int16 = v3
            _ = v13
        }
    }
    delete(v2, "")
    v3 += v3
    switch ("foo bar" + v9) {
    case "a":
        v3 += (v3 - 32767)
    case "hello":
        fmt.Println(v4, v1, v4)
    case "\n\t":
        v14 := v2[(v9 + v9)]
        _ = v14
    }
    v3 = 32767
    v15 := []int{-2147483648, 4294967295}
    if true {
        switch ("\n\t" + (("" + ("a" + v9)) + v9)) {
        case "a":
            for fi16 := 0; fi16 < 8; fi16++ {
                for fi17 := 0; fi17 < 11; fi17++ {
                    v9 = ("foo bar" + "hello")
                    v18 := GenS2{F0: 5, F1: -1}
                    _ = fi17
                    _ = v18
                }
                switch fi16 {
                case -129:
                    if (42 <= -928) {
                        for fi19 := 0; fi19 < 5; fi19++ {
                            var v20 float64 = 0.0
                            _ = fi19
                            _ = v20
                        }
                    }
                case 1:
                    v1 = -38
                }
                _ = fi16
            }
        case "hello":
            for fi21 := 0; fi21 < 13; fi21++ {
                fmt.Println(v4, v1, v9)
                if true {
                    v22 := GenS1{F0: 2, F1: v9, F2: (v9 + v9)}
                    _ = v22
                }
                v2[v9] = (fi21 / ((fi21) | 1))
                _ = fi21
            }
        }
        for fi23 := 0; fi23 < 9; fi23++ {
            if len(v15) > 0 {
                v15[0] = fi23
            }
            v3 += 1000
            if (false && false) {
                fmt.Println(v1, fi23)
            } else {
                fmt.Println(v1)
                v1 -= v1
                v24 := v2[(((v9 + (v9 + "")) + "hello") + v9)]
                _ = v24
            }
            _ = fi23
        }
    } else {
        if (2147483648 != 50) {
            var v25 bool = (!((v4 ^ (v4 | v4)) > (v4 << ((967) & 63))))
            _ = v25
        } else {
            v26 := GenS3{F0: (v4 << ((888) & 63)), F1: 816, F2: 0}
            v27 := GenS1{F0: 945, F1: "", F2: v9}
            _ = v26
            _ = v27
        }
        v4 = (-(v4))
    }
    for fi28 := 0; fi28 < 18; fi28++ {
        v9 = v9
        _ = fi28
    }
    v1 += (v1 | v1)
    v29 := GenS4{F0: 42, F1: v1}
    v4 -= v4
    var v30 int32 = 452
    if (!true) {
        var v31 int32 = 1
        v31 -= v30
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v9
}
