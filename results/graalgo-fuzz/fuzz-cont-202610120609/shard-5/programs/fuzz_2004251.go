package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0() int16 {
    var v7 int32 = -1
    v7 = -278
    v7 += v7
    var v8 string = ""
    _ = v7
    _ = v8
    return 32767
}
func genFn1(p0 int8) bool {
    if (1 <= 18446744073709551615) {
        for fi9 := 0; fi9 < 13; fi9++ {
            p0 -= 100
            v10 := GenS1{F0: 724}
            if false {
                fmt.Println(fi9, fi9)
            } else {
                for fi11 := 0; fi11 < 5; fi11++ {
                    var v12 float32 = (1.5 / 1.5)
                    v13 := []int{fi11, fi11, fi9, fi11}
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                v14 := []int{(^(fi9))}
                for fi15 := 0; fi15 < 6; fi15++ {
                    fmt.Println(p0, p0, p0)
                    _ = fi15
                }
                _ = v14
            }
            _ = fi9
            _ = v10
        }
        var v16 string = ("a" + "foo bar")
        var v17 int8 = -36
        _ = v16
        _ = v17
    }
    p0 -= 127
    _ = p0
    return ((-2.25 / -2.25) > 3.14159)
}

func main() {
    var v1 int32 = 217
    var v2 int8 = 33
    var v3 uint = 0
    var v4 int32 = 0
    v5 := []string{"\n\t"}
    var v6 int = int(v3)
    switch "hello" {
    case "a":
        if false {
            v2 += v2
            var v18 uint32 = 1
            _ = v18
        } else {
            var v19 uint8 = 128
            var v20 int16 = -654
            _ = v19
            _ = v20
        }
    case "foo bar":
        v21 := GenS1{F0: 1}
        _ = v21
    default:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    if ("" > "\n\t") {
        var v22 float32 = ((((1.5 / -2.25) - 1e10) / 1.0) - 3.14159)
        v5 = append(v5, "")
        var v23 bool = true
        _ = v22
        _ = v23
    }
    v4 -= (v1 / ((-524) | 1))
    if ((!(!true)) || true) {
        var v24 float32 = 0.1
        v5 = append(v5, "foo bar")
        v1 = (v1 / ((int32(v6)) | 1))
        _ = v24
    } else {
        var v25 int64 = -9223372036854775808
        v25 -= (-(v25))
        _ = v25
    }
    for fi26 := 0; fi26 < 15; fi26++ {
        v27 := GenS1{F0: 402}
        _ = fi26
        _ = v27
    }
    switch ("\n\t" + "foo bar") {
    case "":
        var v28 float32 = 0.0
        _ = v28
    }
    v3 -= (^(v3))
    var v29 float32 = 1e10
    v30 := GenS1{F0: 18446744073709551615}
    fmt.Println(v3)
    var v31 int32 = v4
    v1 -= v4
    _ = len(v5)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi32 := 0; fi32 < 2; fi32++ {
        _ = len(v5)
        _ = fi32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
