package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 int64 }
func genFn0() uint16 {
    var v6 int16 = 100
    v6 = (-(v6))
    v7 := GenS1{F0: "a"}
    _ = v6
    _ = v7
    return 32767
}
func genFn1(p0 uint8) uint16 {
    var v8 int16 = -32768
    if (false || true) {
        v9 := []int{-494, 897}
        _ = v9
    } else {
        for fi10 := 0; fi10 < 12; fi10++ {
            var v11 int64 = 0
            if (false && (3.14159 >= 0.0)) {
                var v12 int32 = 0
                v12 = int32(fi10)
                for fi13 := 0; fi13 < 6; fi13++ {
                    p0 = 103
                    v14 := GenS1{F0: (("foo bar" + "\n\t") + "a")}
                    switch fi10 {
                    case 32767:
                        switch "foo bar" {
                        case "a":
                            v15 := GenS2{F0: -2147483648, F1: -1}
                            _ = v15
                        case "":
                            fmt.Println(fi13, fi13)
                        case "foo bar":
                            var v16 uint32 = 953
                            _ = v16
                        }
                    case 9223372036854775807:
                        v17 := []int{fi10}
                        _ = v17
                    case 1:
                        v8 = (v8 / ((v8) | 1))
                    }
                    _ = fi13
                    _ = v14
                }
                _ = v12
            }
            _ = fi10
            _ = v11
        }
    }
    var v18 int64 = 0
    _ = p0
    _ = v18
    _ = v8
    return 347
}

func main() {
    var v1 int16 = 243
    var v2 float64 = 1.5
    var v3 float32 = 1.5
    var v4 int32 = 2147483647
    v5 := make(map[string]bool)
    switch 0 {
    case -650:
        for fi19 := 0; fi19 < 6; fi19++ {
            var v20 int32 = (-(v4))
            _ = fi19
            _ = v20
        }
    case 32768:
        var v21 uint = 42
        _ = v21
    }
    if true {
        var v22 uint = 65535
        _ = v22
    }
    delete(v5, ("hello" + ("\n\t" + "hello")))
    for fi23 := 0; fi23 < 2; fi23++ {
        delete(v5, ("foo bar" + "\n\t"))
        v4 = -841
        var v24 float32 = v3
        _ = fi23
        _ = v24
    }
    v3 = float32((v2 / 1.0))
    v1 -= -566
    v25 := GenS2{F0: (v4 ^ (v4 / ((v4) | 1))), F1: 9223372036854775807}
    var v26 int32 = v4
    v27 := GenS2{F0: -1, F1: -1}
    v2 = ((0.0 + float64(-2.25)) / ((1e10 / (v2 / (float64(1.5) + float64((float32(1e10) - (float32(v2) / (v3 / 1.0))))))) - (v2 / v2)))
    v28 := []int{1, 32768}
    var v29 int16 = (v1 | (v1 >> ((42) & 63)))
    v5["a"] = ((v29 >= v1) && (1 <= 65535))
    var v30 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
