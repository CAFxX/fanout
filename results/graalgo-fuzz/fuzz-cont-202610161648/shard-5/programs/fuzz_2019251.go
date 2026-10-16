package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8 }
func genFn0(p0 string) int8 {
    v7 := []int{-9223372036854775808, 128}
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 int32 = 42
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: ("" + (p0 + "\n\t"))}
    var v11 float64 = (0.1 / 3.14159)
    _ = p0
    _ = v10
    _ = v11
    _ = v7
    return 42
}

func main() {
    var v1 uint8 = 23
    var v2 float64 = 3.14159
    v3 := []string{"\n\t", "foo bar"}
    var v4 int32 = 2147483647
    v5 := make(map[int]int)
    v6 := []string{"a"}
    if ((!false) && true) {
        if (1 >= 2147483647) {
            if (false || true) {
                v4 -= 915
                if (1 < 0) {
                    for fi12 := 0; fi12 < 8; fi12++ {
                        v1 += (v1 ^ 255)
                        _ = fi12
                    }
                    fmt.Println(v1, v4, v4)
                    if (v2 >= v2) {
                        v4 = v4
                        v1 -= (v1 >> ((2) & 63))
                    }
                }
            }
            v4 -= v4
        }
    } else {
        for fi13 := 0; fi13 < 1; fi13++ {
            v6 = append(v6, ((("" + ("" + ("\n\t" + "foo bar"))) + ("a" + ("foo bar" + "hello"))) + ""))
            _ = fi13
        }
    }
    var v14 string = "\n\t"
    _ = len(v6)
    fmt.Println(v2, v14, v14)
    if true {
        if len(v6) > 0 {
            v6[0] = v14
        }
        v2 -= v2
    }
    v15 := GenS1{F0: v14}
    for fi16 := 0; fi16 < 5; fi16++ {
        switch v14 {
        case "foo bar":
            var v17 int = -9223372036854775808
            _ = v17
        }
        _ = fi16
    }
    v1 = v1
    v14 = "foo bar"
    v1 += (v1 + (v1 * 128))
    v18 := GenS2{F0: genFn0(("hello" + (v14 + (v14 + v14))))}
    for fi19 := 0; fi19 < 18; fi19++ {
        v5[-9223372036854775808] = (^(fi19))
        for fi20 := 0; fi20 < 11; fi20++ {
            v1 -= (v1 ^ (^(v1)))
            _ = fi20
        }
        fmt.Println(v1, v4, fi19)
        _ = fi19
    }
    for fi21 := 0; fi21 < 6; fi21++ {
        fmt.Println(v4, fi21, v2)
        var v22 string = v14
        _ = fi21
        _ = v22
    }
    v2 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
