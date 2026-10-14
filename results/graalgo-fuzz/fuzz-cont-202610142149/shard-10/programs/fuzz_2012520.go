package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 int8 }
type GenS2 struct { F0 int8; F1 uint32 }
func genFn0(p0 int64, p1 string, p2 int32) int32 {
    p1 = p1
    fmt.Println(p2, p1)
    p1 = (("a" + (("" + p1) + p1)) + p1)
    var v4 uint32 = 464
    if (!(42 <= 1)) {
        var v5 uint16 = 2
        _ = v5
    } else {
        for fi6 := 0; fi6 < 19; fi6++ {
            switch ((p1 + p1) + p1) {
            case "\n\t":
                fmt.Println(p0, p2, p2)
            case "":
                if ((!false) && (42 == 1)) {
                    v4 -= (v4 % (((v4 % (((v4 - (v4 >> ((2) & 63)))) | 1))) | 1))
                    var v7 string = (("hello" + ("hello" + (p1 + (p1 + (((("\n\t" + p1) + p1) + p1) + p1))))) + "\n\t")
                    v8 := GenS1{F0: 783, F1: "foo bar", F2: 127}
                    _ = v7
                    _ = v8
                }
            case "a":
                p1 = "a"
            default:
                v9 := GenS2{F0: 0, F1: 0}
                _ = v9
            }
            var v10 string = (p1 + (p1 + p1))
            v10 = "\n\t"
            _ = fi6
            _ = v10
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return p2
}

func main() {
    var v1 int16 = 1000
    v2 := make(map[string]string)
    v3 := []int{-9223372036854775808, 32768, 816}
    v1 = (v1 % ((v1) | 1))
    fmt.Println(v1)
    _ = len(v3)
    var v11 int32 = genFn0(-590, ("hello" + "\n\t"), genFn0(-1, ("" + "hello"), 368))
    v11 = genFn0(-287, ("\n\t" + ("hello" + "\n\t")), genFn0(42, "", (v11 ^ (v11 << ((18446744073709551615) & 63)))))
    for fi12 := 0; fi12 < 2; fi12++ {
        for fi13 := 0; fi13 < 11; fi13++ {
            v1 = (v1 | v1)
            v11 += (-(v11))
            _ = fi13
        }
        _ = fi12
    }
    v14 := v2[("\n\t" + ("" + ("\n\t" + "foo bar")))]
    fmt.Println(v1, v14, v11)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v3
}
