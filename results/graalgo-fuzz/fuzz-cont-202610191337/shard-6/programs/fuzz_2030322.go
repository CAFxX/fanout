package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 string; F2 uint8 }
func genFn0(p0 uint8, p1 uint) int {
    if false {
        for fi4 := 0; fi4 < 17; fi4++ {
            v5 := []int{fi4, fi4, (fi4 << (((p1 / ((p1) | 1))) & 63)), 32768}
            _ = fi4
            _ = v5
        }
    }
    p1 -= p1
    v6 := GenS1{F0: (("a" + (("foo bar" + "\n\t") + "\n\t")) + "foo bar")}
    if ((2 >= 65535) && ("\n\t" >= "\n\t")) {
        var v7 int8 = 1
        _ = v7
    }
    if true {
        p1 -= 42
        switch "" {
        case "\n\t":
            var v8 int = -2147483648
            _ = v8
        case "a":
            v9 := []int{4294967295}
            _ = v9
        default:
            switch "hello" {
            case "hello":
                if (p0 > p0) {
                    v10 := GenS2{F0: 1000}
                    var v11 bool = (!true)
                    var v12 int32 = 634
                    _ = v10
                    _ = v11
                    _ = v12
                }
            default:
                v13 := GenS1{F0: "foo bar"}
                _ = v13
            }
        }
    }
    _ = p0
    _ = p1
    _ = v6
    return 1
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []uint32{4294967295, 42, 4294967295}
    v3 := []int{}
    v14 := GenS3{F0: 42, F1: "foo bar", F2: 2}
    var v15 float64 = (0.1 / 1e10)
    var v16 float64 = (v15 - (-2.25 / float64(((0.1 - ((1.5 - 0.1) * -2.25)) - ((0.0 * 1e10) * (1e10 / (1e10 + 1.5)))))))
    var v17 int8 = -73
    fmt.Println(v15)
    var v18 int32 = -1
    v19 := GenS2{F0: 1}
    v15 += v15
    v17 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
