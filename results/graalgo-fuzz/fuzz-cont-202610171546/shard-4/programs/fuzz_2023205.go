package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint }
func genFn0(p0 uint, p1 bool, p2 uint32) bool {
    p1 = p1
    v7 := GenS1{F0: ("\n\t" + ""), F1: 0, F2: p0}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return false
}

func main() {
    v1 := make(map[int]string)
    var v2 int = 812
    var v3 uint = 778
    var v4 bool = true
    v5 := []string{("" + ("foo bar" + "hello")), "\n\t", "hello"}
    v6 := make(map[string]string)
    for fi8 := 0; fi8 < 13; fi8++ {
        _ = len(v5)
        _ = fi8
    }
    if (-32768 != 0) {
        if ((1 <= 42) || false) {
            switch (v2 | 0) {
            case -59:
                for fi9 := 0; fi9 < 13; fi9++ {
                    for fi10 := 0; fi10 < 3; fi10++ {
                        fmt.Println(fi9, v3)
                        v2 = 4294967295
                        fmt.Println(fi9, v2)
                        _ = fi10
                    }
                    _ = fi9
                }
            case 128:
                v11 := GenS1{F0: "foo bar", F1: 2, F2: (v3 & (v3 + 42))}
                _ = v11
            }
            switch "" {
            case "hello":
                var v12 float64 = 0.0
                _ = v12
            case "foo bar":
                var v13 string = "hello"
                _ = v13
            case "\n\t":
                v2 = (v2 / (((-(v2))) | 1))
            default:
                var v14 float64 = 1.5
                _ = v14
            }
        }
        for fi15 := 0; fi15 < 12; fi15++ {
            v4 = ((-1 != 32767) || v4)
            fmt.Println(v2, v2, v4)
            _ = fi15
        }
        for fi16 := 0; fi16 < 2; fi16++ {
            v3 += v3
            _ = fi16
        }
    } else {
        if (-1 < -19) {
            switch "hello" {
            case "hello":
                for fi17 := 0; fi17 < 19; fi17++ {
                    v5 = append(v5, (("" + ("foo bar" + ("\n\t" + "a"))) + "hello"))
                    fmt.Println(fi17)
                    _ = fi17
                }
            case "a":
                fmt.Println(v4, v4)
            }
            if v4 {
                fmt.Println(v4, v3, v4)
            }
            if len(v5) > 0 {
                v5[0] = "a"
            }
        }
        var v18 bool = true
        _ = v18
    }
    var v19 uint8 = 1
    var v20 bool = (42 != -568)
    var v21 uint32 = 42
    v5 = append(v5, "")
    var v22 int16 = -1
    v5 = append(v5, "a")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
