package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 string }
func genFn0(p0 int32, p1 int8, p2 int32) uint64 {
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 string = "hello"
        v9 := GenS1{F0: 42}
        if true {
            var v10 uint8 = 0
            var v11 bool = true
            _ = v10
            _ = v11
        } else {
            if (!(true || false)) {
                p0 = p2
                for fi12 := 0; fi12 < 3; fi12++ {
                    switch fi7 {
                    case 127:
                        v13 := []int{32768, -109, fi7}
                        _ = v13
                    }
                    if (1 > -1) {
                        v14 := []int{1, (fi7 / (((fi12 << ((2) & 63))) | 1))}
                        var v15 int32 = (p0 + -1)
                        _ = v14
                        _ = v15
                    } else {
                        fmt.Println(fi12, v8, p1)
                        switch fi12 {
                        case 32767:
                            v16 := []int{fi7, fi12, fi7, fi12}
                            _ = v16
                        }
                        fmt.Println(fi12, p2)
                    }
                    if ((1e10 <= -2.25) || (p1 >= (p1 - (-(p1))))) {
                        var v17 string = (v8 + ((v8 + v8) + ((v8 + ((("a" + "foo bar") + v8) + (v8 + ("a" + v8)))) + (v8 + "foo bar"))))
                        if ((0.1 + -2.25) != ((0.0 / float32(1e10)) - -2.25)) {
                            v8 = (v17 + (v17 + (v8 + "foo bar")))
                            p0 = p2
                        }
                        switch (v17 + v8) {
                        case "":
                            v18 := GenS2{F0: 0, F1: ""}
                            _ = v18
                        case "foo bar":
                            var v19 string = ((((v17 + v8) + v17) + "foo bar") + v17)
                            _ = v19
                        }
                        fmt.Println(p0, p1)
                        _ = v17
                    }
                    _ = fi12
                }
            } else {
                p2 = (p2 + p2)
            }
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    v20 := []int{32768, 127, 32768}
    _ = len(v20)
    if len(v20) > 0 {
        v20[0] = 32768
    }
    if len(v20) >= 2 {
        v20 = v20[1:len(v20)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    return 388
}

func main() {
    var v1 uint32 = 771
    var v2 int8 = 7
    v3 := []string{"hello", "\n\t", ""}
    var v4 string = ("foo bar" + "\n\t")
    var v5 string = (((v4 + v4) + "hello") + "foo bar")
    var v6 float64 = 0.1
    var v21 float32 = 3.14159
    for fi22 := 0; fi22 < 6; fi22++ {
        v23 := GenS2{F0: 9223372036854775807, F1: (("foo bar" + ((v4 + "a") + "hello")) + v4)}
        var v24 uint8 = 42
        _ = fi22
        _ = v23
        _ = v24
    }
    var v25 bool = (156 < 248)
    for fi26 := 0; fi26 < 11; fi26++ {
        fmt.Println(v1, v4, v6)
        _ = fi26
    }
    fmt.Println(v21)
    var v27 float64 = -2.25
    v2 -= v2
    var v28 uint16 = 1
    v25 = (v4 > ("foo bar" + v5))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v25
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
