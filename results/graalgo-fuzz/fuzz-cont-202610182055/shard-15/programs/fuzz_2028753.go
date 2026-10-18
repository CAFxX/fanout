package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 int, p1 string) string {
    if (0 > 127) {
        p1 = (p1 + p1)
    }
    if ((867 < 0) || (18446744073709551615 == 18446744073709551615)) {
        var v6 int16 = 1000
        var v7 string = p1
        if (((0 != -9223372036854775808) || true) || (1 > 42)) {
            switch "" {
            case "":
                if (v6 > (v6 ^ v6)) {
                    v7 = v7
                    v8 := []int{(-(p0)), p0}
                    _ = v8
                }
            case "a":
                for fi9 := 0; fi9 < 10; fi9++ {
                    v7 = "a"
                    var v10 uint16 = 875
                    for fi11 := 0; fi11 < 1; fi11++ {
                        var v12 float32 = 1e10
                        fmt.Println(v7)
                        v13 := GenS1{F0: 754}
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                    _ = fi9
                    _ = v10
                }
            case "\n\t":
                v14 := GenS2{F0: 4294967295}
                _ = v14
            }
        }
        _ = v6
        _ = v7
    }
    v15 := []int{p0, (p0 & p0), p0, p0}
    var v16 int64 = -9223372036854775808
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    return p1
}
func genFn1(p0 int16, p1 int) int64 {
    for fi17 := 0; fi17 < 15; fi17++ {
        if false {
            for fi18 := 0; fi18 < 16; fi18++ {
                var v19 bool = true
                v20 := []int{}
                _ = fi18
                _ = v19
                _ = v20
            }
            switch 568 {
            case -2147483648:
                var v21 int64 = -9223372036854775808
                _ = v21
            }
            switch "a" {
            case "":
                var v22 int16 = (p0 ^ p0)
                _ = v22
            }
        }
        var v23 float64 = 0.1
        _ = fi17
        _ = v23
    }
    var v24 string = "\n\t"
    v24 = (("hello" + ("hello" + "\n\t")) + v24)
    var v25 int8 = 18
    _ = p0
    _ = p1
    _ = v24
    _ = v25
    return 42
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 uint = v1
    v3 := []string{"", ("a" + ""), "\n\t"}
    v4 := []string{"\n\t", ""}
    var v5 float64 = (-2.25 / (0.1 * 1e10))
    fmt.Println(v2)
    var v26 int16 = -648
    v5 = float64(0.1)
    var v27 bool = false
    v2 = 4294967295
    v28 := GenS1{F0: 780}
    var v29 uint16 = 2
    for fi30 := 0; fi30 < 1; fi30++ {
        var v31 uint8 = 245
        _ = fi30
        _ = v31
    }
    _ = len(v3)
    v27 = v27
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
