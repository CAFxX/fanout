package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 uint32 }
type GenS2 struct { F0 uint32; F1 uint64; F2 uint32 }
func genFn0(p0 int, p1 uint32) uint8 {
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 float32 = ((1.5 * 3.14159) - float32(1.5))
        _ = fi7
        _ = v8
    }
    p0 += (p0 ^ 32767)
    for fi9 := 0; fi9 < 2; fi9++ {
        if ((-119 >= -14) && true) {
            fmt.Println(fi9, p1, fi9)
        }
        for fi10 := 0; fi10 < 19; fi10++ {
            p1 = 2147483647
            var v11 int64 = -9223372036854775808
            if (127 != 38) {
                v12 := []int{(fi10 / ((fi10) | 1)), p0, p0, 32768}
                p1 -= (p1 ^ 419)
                v13 := GenS1{F0: "a", F1: 255, F2: p1}
                if (2147483647 >= fi10) {
                    v14 := GenS1{F0: "foo bar", F1: 4294967295, F2: (^(p1))}
                    switch "\n\t" {
                    case "hello":
                        var v15 float32 = 1.5
                        _ = v15
                    case "foo bar":
                        if true {
                            p0 -= (fi9 / (((fi10 ^ (^(fi9)))) | 1))
                        }
                    default:
                        v16 := GenS2{F0: p1, F1: 0, F2: p1}
                        _ = v16
                    }
                    _ = v14
                }
                p1 = p1
                _ = v12
                _ = v13
            } else {
                fmt.Println(fi9, v11, p0)
                var v17 int16 = 718
                _ = v17
            }
            _ = fi10
            _ = v11
        }
        var v18 int32 = -2147483648
        _ = fi9
        _ = v18
    }
    _ = p0
    _ = p1
    return 137
}

func main() {
    v1 := []string{"a"}
    v2 := []string{(((("hello" + ("foo bar" + ("\n\t" + "foo bar"))) + (("a" + "hello") + "foo bar")) + "a") + "\n\t"), ("\n\t" + (("a" + "") + ("foo bar" + "a")))}
    v3 := make(map[string]bool)
    var v4 float64 = float64(1e10)
    var v5 int = -406
    var v6 float64 = v4
    switch (v5 >> ((947) & 63)) {
    case 32767:
        delete(v3, "\n\t")
    case 9223372036854775807:
        var v19 float64 = v6
        _ = v19
    }
    var v20 bool = (18446744073709551615 != 1)
    v1 = append(v1, ("hello" + ("" + "")))
    v21 := GenS1{F0: "a", F1: 18446744073709551615, F2: 2}
    var v22 uint = 255
    switch v5 {
    case 672:
        var v23 int = v5
        _ = v23
    case 4294967295:
        for fi24 := 0; fi24 < 2; fi24++ {
            var v25 int16 = -1
            _ = fi24
            _ = v25
        }
    default:
        v6 += v4
    }
    v5 = v5
    _ = len(v3)
    delete(v3, "")
    _ = len(v3)
    v2 = append(v2, ("a" + "\n\t"))
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    if (!v20) {
        for fi26 := 0; fi26 < 14; fi26++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi26
        }
        v5 -= 1
        var v27 bool = (1 > 42)
        _ = v27
    }
    v28 := GenS2{F0: 487, F1: 984, F2: 2147483647}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
