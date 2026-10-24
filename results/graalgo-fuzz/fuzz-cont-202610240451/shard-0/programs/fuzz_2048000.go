package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
func genFn0(p0 bool) uint16 {
    p0 = p0
    fmt.Println(p0)
    _ = p0
    return 0
}

func main() {
    var v1 int = 72
    var v2 int32 = -1
    var v3 bool = false
    var v4 float32 = (1.5 + 0.0)
    for fi5 := 0; fi5 < 9; fi5++ {
        for fi6 := 0; fi6 < 8; fi6++ {
            if v3 {
                v7 := []int{fi5, v1, (fi5 >> ((255) & 63)), v1}
                var v8 uint16 = 2
                var v9 uint32 = 178
                _ = v7
                _ = v8
                _ = v9
            } else {
                var v10 uint8 = 128
                v3 = (v3 && (1 == -10))
                _ = v10
            }
            switch ("a" + "") {
            case "a":
                if (v4 >= ((1e10 - v4) - 3.14159)) {
                    fmt.Println(v1, v3, fi6)
                    var v11 int64 = 0
                    _ = v11
                }
            case "\n\t":
                if ((fi6 >> ((65535) & 63)) == fi6) {
                    v12 := []int{}
                    v2 -= 189
                    v13 := GenS1{F0: (("" + "hello") != ((("a" + "hello") + "a") + ("" + ("a" + ("\n\t" + ("" + "\n\t")))))), F1: v3}
                    var v14 int64 = 42
                    fmt.Println(v3)
                    if len(v12) >= 2 {
                        v12 = v12[1:len(v12)]
                    }
                    fmt.Println(v4, v2, v2)
                    _ = v12
                    _ = v13
                    _ = v14
                }
            case "hello":
                v15 := []int{}
                _ = v15
            default:
                v16 := GenS1{F0: (585 >= 1), F1: v3}
                _ = v16
            }
            var v17 float64 = (3.14159 - (((3.14159 / 1.5) / 0.1) - float64(v4)))
            _ = fi6
            _ = v17
        }
        _ = fi5
    }
    v1 = v1
    if true {
        for fi18 := 0; fi18 < 3; fi18++ {
            if v3 {
                var v19 uint8 = 127
                if true {
                    v20 := GenS1{F0: v3, F1: ((!((v19 & (v19 / (((^(v19))) | 1))) > v19)) && false)}
                    v2 = (v2 | (-(v2)))
                    _ = v20
                } else {
                    var v21 uint64 = 4294967295
                    _ = v21
                }
                _ = v19
            } else {
                fmt.Println(v2, v3)
            }
            if false {
                var v22 uint32 = 65535
                v2 = v2
                _ = v22
            }
            if true {
                v23 := GenS1{F0: (("foo bar" + "") != "foo bar"), F1: v3}
                _ = v23
            } else {
                for fi24 := 0; fi24 < 18; fi24++ {
                    v3 = true
                    _ = fi24
                }
            }
            _ = fi18
        }
    } else {
        v25 := GenS1{F0: ((18446744073709551615 >= 2) && v3), F1: v3}
        _ = v25
    }
    fmt.Println(v3, v4)
    v2 += v2
    fmt.Println(v4, v1, v2)
    v26 := GenS1{F0: v3, F1: (!((519 >= 128) || v3))}
    v4 -= v4
    v27 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
