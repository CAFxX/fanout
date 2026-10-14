package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint; F1 string; F2 uint32 }
func genFn0(p0 uint8) int8 {
    var v7 int32 = 42
    fmt.Println(v7)
    v8 := GenS1{F0: "a", F1: 1}
    _ = p0
    _ = v7
    _ = v8
    return -128
}

func main() {
    var v1 bool = (!true)
    var v2 bool = v1
    var v3 bool = v1
    var v4 bool = v1
    v5 := []int{-1, -747, 128, 32767}
    v6 := []string{("\n\t" + ""), "a", ("" + "foo bar"), ""}
    var v9 float32 = (float32((float64(1.5) * 1.5)) / (1e10 / float32((1.5 / float64(0.1)))))
    for fi10 := 0; fi10 < 20; fi10++ {
        _ = len(v6)
        v1 = v3
        var v11 int8 = -1
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        if ("\n\t" >= "") {
            v9 = (v9 * float32(1e10))
        } else {
            switch 1 {
            case 32768:
                for fi13 := 0; fi13 < 1; fi13++ {
                    fmt.Println(v4, v4, fi13)
                    _ = fi13
                }
            case 2147483647:
                var v14 float64 = (((0.0 / 0.1) * 1.5) - -2.25)
                _ = v14
            case -1:
                v2 = v1
            default:
                fmt.Println(v2, fi12)
            }
            switch "foo bar" {
            case "":
                if v3 {
                    v15 := GenS2{F0: 42}
                    if false {
                        if v2 {
                            var v16 int64 = 1
                            var v17 int8 = 25
                            _ = v16
                            _ = v17
                        }
                        for fi18 := 0; fi18 < 16; fi18++ {
                            v19 := GenS3{F0: 32767}
                            var v20 float64 = (1.5 / (-2.25 - float64(v9)))
                            v21 := GenS1{F0: "", F1: 18446744073709551615}
                            _ = fi18
                            _ = v19
                            _ = v20
                            _ = v21
                        }
                    }
                    _ = v15
                }
            }
        }
        fmt.Println(v2, v2, v2)
        v5 = append(v5, 1)
        _ = fi12
    }
    v22 := GenS3{F0: 1}
    v6 = append(v6, "foo bar")
    if (float64(-2.25) >= -2.25) {
        var v23 float32 = v9
        _ = v23
    }
    if (0 < 416) {
        v4 = (42 >= 0)
    }
    for fi24 := 0; fi24 < 5; fi24++ {
        var v25 bool = (false || ((fi24 + (^(fi24))) < fi24))
        _ = fi24
        _ = v25
    }
    if false {
        if ((245 > 1) && v4) {
            v2 = (!true)
            v4 = (v2 && (!false))
        }
        fmt.Println(v3, v1)
    }
    if (81 == 0) {
        v4 = false
        var v26 string = "\n\t"
        var v27 uint64 = 421
        _ = v26
        _ = v27
    } else {
        for fi28 := 0; fi28 < 11; fi28++ {
            var v29 uint8 = 2
            var v30 float64 = 3.14159
            _ = fi28
            _ = v29
            _ = v30
        }
        v5 = append(v5, 1)
    }
    var v31 int64 = 268
    v32 := GenS3{F0: 1}
    if len(v6) > 0 {
        v6[0] = "\n\t"
    }
    v33 := GenS4{F0: 18446744073709551615, F1: "a", F2: 813}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
