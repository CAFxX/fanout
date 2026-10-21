package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 string }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int64, p1 uint16, p2 string) uint32 {
    v4 := []int{128, 2147483647, -380}
    v5 := GenS1{F0: false, F1: p2, F2: p2}
    fmt.Println(p1, p0, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 2147483647
}
func genFn1() int32 {
    switch 32768 {
    case -131:
        if false {
            if (false || ((float32(float64(3.14159)) * (3.14159 / 0.1)) < -2.25)) {
                var v6 uint32 = genFn0(-378, 2, "")
                v6 = (^(v6))
                _ = v6
            } else {
                v7 := GenS1{F0: (-2.25 == (1e10 - 0.0)), F1: "a", F2: ("foo bar" + "foo bar")}
                v8 := GenS1{F0: false, F1: "\n\t", F2: "hello"}
                _ = v7
                _ = v8
            }
            var v9 uint32 = genFn0(42, 580, (("\n\t" + "") + ""))
            v10 := GenS1{F0: true, F1: "hello", F2: "a"}
            _ = v10
            _ = v9
        }
    }
    var v11 uint = 42
    var v12 uint64 = 18446744073709551615
    v13 := GenS2{F0: (^(v12))}
    for fi14 := 0; fi14 < 15; fi14++ {
        v15 := []int{32767, 0}
        v11 += (v11 / (((v11 & (v11 + (v11 / (((v11 / (((v11 - v11)) | 1))) | 1))))) | 1))
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
        _ = fi14
        _ = v15
    }
    _ = v11
    _ = v12
    _ = v13
    return 42
}

func main() {
    var v1 int8 = 95
    v2 := []uint32{}
    var v3 uint64 = 18446744073709551615
    var v16 uint64 = 2
    var v17 int = 32768
    var v18 uint8 = 128
    if false {
        v1 -= -128
        v2 = append(v2, 555)
        var v19 int8 = v1
        _ = v19
    }
    if (464 > 32767) {
        for fi20 := 0; fi20 < 6; fi20++ {
            v17 += (v17 * v17)
            _ = fi20
        }
    }
    v17 = (v17 + v17)
    v1 -= (^(v1))
    for fi21 := 0; fi21 < 19; fi21++ {
        var v22 string = ""
        if false {
            fmt.Println(v22)
        }
        var v23 bool = (!false)
        _ = fi21
        _ = v22
        _ = v23
    }
    var v24 int64 = 42
    var v25 int8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v3
}
