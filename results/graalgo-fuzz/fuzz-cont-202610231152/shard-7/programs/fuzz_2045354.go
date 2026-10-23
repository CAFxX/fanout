package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 uint; F1 int; F2 uint64 }
func genFn0(p0 int16, p1 bool, p2 int32) bool {
    var v4 uint64 = 1
    fmt.Println(v4, v4, p0)
    var v5 float64 = -2.25
    v6 := GenS1{F0: (^(v4))}
    v7 := []int{42, 127, 9223372036854775807, -9223372036854775808}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return false
}
func genFn1() bool {
    var v8 int32 = 0
    v9 := GenS1{F0: 945}
    for fi10 := 0; fi10 < 12; fi10++ {
        v8 = -600
        _ = fi10
    }
    for fi11 := 0; fi11 < 2; fi11++ {
        v8 = v8
        v8 = (v8 & (v8 & int32(fi11)))
        if false {
            v12 := GenS1{F0: 1}
            _ = v12
        }
        _ = fi11
    }
    var v13 int16 = 0
    _ = v13
    _ = v8
    _ = v9
    return false
}

func main() {
    var v1 uint32 = 4294967295
    v2 := make(map[string]bool)
    v3 := make(map[uint32]bool)
    var v14 float32 = 0.1
    switch 194 {
    case 0:
        var v15 uint32 = (v1 ^ 404)
        _ = v15
    default:
        if (("" + (("hello" + "a") + "a")) >= "\n\t") {
            switch "\n\t" {
            case "hello":
                fmt.Println(v1, v1)
            case "\n\t":
                for fi16 := 0; fi16 < 13; fi16++ {
                    var v17 float32 = 3.14159
                    v18 := GenS2{F0: 100, F1: 2147483647}
                    var v19 uint = 369
                    _ = fi16
                    _ = v17
                    _ = v18
                    _ = v19
                }
            }
            fmt.Println(v1)
        }
    }
    switch 32768 {
    case 27:
        fmt.Println(v14, v14)
    }
    if true {
        if (((v1 * (v1 % ((513) | 1))) <= v1) && false) {
            var v20 uint32 = (^(v1))
            v21 := v3[(v20 * 272)]
            _ = v20
            _ = v21
        } else {
            v1 = 42
            switch 981 {
            case -968:
                var v22 uint16 = 1
                _ = v22
            case 32768:
                if (180 < 0) {
                    v14 -= ((3.14159 - 1.5) / 1e10)
                    v1 = (^(v1))
                }
            case -282:
                fmt.Println(v14, v1)
            }
            v2["foo bar"] = ((v14 / (v14 / float32((-2.25 + (3.14159 * -2.25))))) <= float32(3.14159))
        }
        if false {
            v1 += (v1 % ((v1) | 1))
        } else {
            var v23 int64 = -1
            if false {
                fmt.Println(v1, v1)
                var v24 uint64 = 82
                _ = v24
            }
            v25 := GenS3{F0: (^(v1)), F1: v1}
            _ = v23
            _ = v25
        }
        _ = len(v3)
    }
    v14 = (v14 + 1e10)
    v26 := GenS3{F0: v1, F1: v1}
    for fi27 := 0; fi27 < 5; fi27++ {
        switch (fi27 - fi27) {
        case 0:
            v28 := GenS4{F0: uint(v1), F1: (^(fi27)), F2: 2}
            _ = v28
        case 127:
            v29 := v2[("hello" + "foo bar")]
            _ = v29
        }
        v1 = (^(v1))
        _ = fi27
    }
    for fi30 := 0; fi30 < 4; fi30++ {
        var v31 uint = 974
        var v32 uint = (v31 * v31)
        _ = fi30
        _ = v31
        _ = v32
    }
    var v33 int32 = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    _ = v1
    _ = v14
    _ = v2
    _ = v26
    _ = v3
    _ = v33
}
