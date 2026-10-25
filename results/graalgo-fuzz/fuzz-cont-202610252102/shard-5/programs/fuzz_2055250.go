package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 uint8 }
func genFn0(p0 uint16, p1 int, p2 int8) int16 {
    switch "hello" {
    case "hello":
        switch p1 {
        case 907:
            if false {
                for fi5 := 0; fi5 < 9; fi5++ {
                    v6 := []int{32767, (fi5 * p1)}
                    v6 = append(v6, (fi5 * (fi5 * -2147483648)))
                    p0 = (p0 & p0)
                    _ = fi5
                    _ = v6
                }
                var v7 int16 = 42
                v7 = v7
                _ = v7
            }
        }
    case "":
        p1 = p1
    case "\n\t":
        v8 := []int{(p1 - (p1 >> ((743) & 63))), p1, int(p0)}
        _ = v8
    default:
        if false {
            if (562 > 42) {
                p2 = (p2 + (p2 % (((p2 | p2)) | 1)))
                v9 := []int{p1, p1, (p1 ^ (p1 - (p1 | 344)))}
                _ = v9
            }
        }
    }
    fmt.Println(p2, p0, p0)
    _ = p0
    _ = p1
    _ = p2
    return 1000
}

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 32767
    v3 := make(map[int]string)
    var v4 string = "foo bar"
    if true {
        switch 32767 {
        case 278:
            for fi10 := 0; fi10 < 14; fi10++ {
                v4 = ((v4 + "\n\t") + ((("" + v4) + ((((((v4 + ("" + v4)) + (v4 + ((v4 + v4) + "a"))) + v4) + (v4 + (((v4 + v4) + v4) + v4))) + "hello") + "")) + "foo bar"))
                _ = fi10
            }
        case 572:
            for fi11 := 0; fi11 < 4; fi11++ {
                v1[(fi11 * (-(fi11)))] = (("foo bar" + ("foo bar" + v4)) + "foo bar")
                fmt.Println(fi11, v4, fi11)
                fmt.Println(v2)
                _ = fi11
            }
        default:
            for fi12 := 0; fi12 < 1; fi12++ {
                v13 := v3[(^(fi12))]
                _ = fi12
                _ = v13
            }
        }
    } else {
        var v14 int = 2147483647
        _ = len(v3)
        v2 = 413
        _ = v14
    }
    _ = len(v1)
    if (true || (false || (42 != 937))) {
        var v15 float32 = float32((1e10 - (float64(1e10) * 3.14159)))
        v16 := GenS1{F0: 1}
        for fi17 := 0; fi17 < 5; fi17++ {
            v18 := []int{(fi17 / ((32768) | 1)), fi17, (fi17 * 42)}
            _ = fi17
            _ = v18
        }
        _ = v15
        _ = v16
    }
    v19 := []int{-419}
    v4 = v4
    var v20 bool = false
    v4 = v4
    fmt.Println(v4, v2)
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    v21 := GenS2{F0: 42, F1: 232}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
