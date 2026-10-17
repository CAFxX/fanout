package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int8; F1 uint32 }
func genFn0(p0 uint64, p1 uint8) string {
    if false {
        v5 := []int{396, 32768, 127, 125}
        _ = v5
    } else {
        for fi6 := 0; fi6 < 9; fi6++ {
            p1 = (p1 | (p1 ^ p1))
            v7 := GenS1{F0: 255, F1: fi6}
            _ = fi6
            _ = v7
        }
        switch "" {
        case "foo bar":
            v8 := []int{}
            _ = v8
        case "a":
            for fi9 := 0; fi9 < 8; fi9++ {
                for fi10 := 0; fi10 < 6; fi10++ {
                    v11 := []int{(^(fi10))}
                    _ = fi10
                    _ = v11
                }
                if (42 >= p1) {
                    var v12 bool = true
                    v12 = v12
                    _ = v12
                }
                _ = fi9
            }
        case "":
            p0 = 18446744073709551615
        default:
            p1 = 255
        }
        fmt.Println(p0)
    }
    v13 := GenS2{F0: 127, F1: 4294967295}
    v14 := []int{-638, 32767, 9223372036854775807, 127}
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return ""
}

func main() {
    v1 := make(map[string]string)
    var v2 bool = false
    v3 := []string{(("a" + "hello") + ("a" + ("" + "")))}
    v4 := []string{"\n\t", "a", ((("\n\t" + "hello") + "foo bar") + "a")}
    for fi15 := 0; fi15 < 16; fi15++ {
        switch ("\n\t" + ("a" + "hello")) {
        case "":
            var v16 string = "\n\t"
            _ = v16
        case "foo bar":
            for fi17 := 0; fi17 < 9; fi17++ {
                if (!(4294967295 == 96)) {
                    fmt.Println(fi15, fi17)
                    v18 := GenS2{F0: 25, F1: 2147483647}
                    var v19 int8 = -8
                    _ = v18
                    _ = v19
                }
                _ = fi17
            }
        case "\n\t":
            fmt.Println(v2, fi15)
        }
        v2 = v2
        for fi20 := 0; fi20 < 8; fi20++ {
            var v21 uint16 = 30
            v21 -= 1
            v2 = ((1.5 + 0.1) != 0.0)
            _ = len(v1)
            v2 = v2
            var v22 uint = 4294967295
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = fi15
    }
    v2 = (18 != 42)
    fmt.Println(v2)
    v4 = append(v4, "\n\t")
    v23 := v1["hello"]
    v23 = "hello"
    v2 = true
    if len(v3) > 0 {
        v3[0] = v23
    }
    fmt.Println(v23, v23)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
