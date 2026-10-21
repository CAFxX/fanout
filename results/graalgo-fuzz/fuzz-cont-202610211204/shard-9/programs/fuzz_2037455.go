package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32; F1 int8 }
func genFn0(p0 bool, p1 uint32) int64 {
    p1 = (p1 + (p1 & p1))
    if (p0 || (-2.25 != 3.14159)) {
        v5 := GenS1{F0: 0}
        fmt.Println(p0)
        _ = v5
    }
    v6 := []int{127, 4294967295, 0}
    _ = p0
    _ = p1
    _ = v6
    return 42
}

func main() {
    var v1 float64 = 0.1
    v2 := make(map[uint32]bool)
    var v3 bool = true
    v4 := make(map[string]int)
    var v7 float32 = (1e10 + 1.5)
    for fi8 := 0; fi8 < 7; fi8++ {
        switch ("hello" + "a") {
        case "hello":
            v9 := []int{(fi8 / (((-(fi8))) | 1)), fi8, (fi8 | (fi8 / (((fi8 % ((fi8) | 1))) | 1)))}
            _ = v9
        case "foo bar":
            delete(v2, 891)
        }
        var v10 int8 = 0
        _ = fi8
        _ = v10
    }
    v11 := GenS2{F0: 0}
    v7 = (((1e10 + (1.5 * float32((((1.5 * v1) + v1) / v1)))) * 0.1) - 0.0)
    v3 = (127 >= 42)
    v12 := GenS1{F0: 731}
    if (!(!false)) {
        v1 = v1
        for fi13 := 0; fi13 < 16; fi13++ {
            if ((-32768 >= -32768) || (v3 || true)) {
                v3 = v3
                if ((361 >= 42) || ((!v3) || v3)) {
                    var v14 int8 = 127
                    switch (fi13 & (-(fi13))) {
                    case 550:
                        v7 = 1.5
                    case -2147483648:
                        if v3 {
                            v7 -= (v7 / ((((v7 * v7) - -2.25) * v7) * 3.14159))
                            var v15 bool = (!((-32768 == 1) || true))
                            _ = v15
                        } else {
                            var v16 string = (("\n\t" + "hello") + "foo bar")
                            switch (fi13 | (fi13 - fi13)) {
                            case -1:
                                v3 = v3
                            }
                            v3 = ((v14 >= v14) || ((!(!v3)) || true))
                            _ = v16
                        }
                    case 1:
                        var v17 int64 = -1
                        _ = v17
                    }
                    _ = v14
                } else {
                    if v3 {
                        v4["a"] = fi13
                    } else {
                        v18 := GenS2{F0: 89}
                        var v19 string = ("foo bar" + "hello")
                        _ = v18
                        _ = v19
                    }
                }
            }
            if false {
                switch ((("" + "hello") + "") + "\n\t") {
                case "foo bar":
                    v1 = float64(v7)
                case "a":
                    v20 := GenS1{F0: 746}
                    _ = v20
                }
                for fi21 := 0; fi21 < 19; fi21++ {
                    v22 := GenS3{F0: 65535, F1: -16}
                    _ = fi21
                    _ = v22
                }
            }
            var v23 int64 = -884
            _ = fi13
            _ = v23
        }
        switch ("a" + ("a" + "hello")) {
        case "foo bar":
            var v24 uint64 = 18446744073709551615
            _ = v24
        case "\n\t":
            v3 = v3
        default:
            for fi25 := 0; fi25 < 10; fi25++ {
                if (!v3) {
                    v7 = 1e10
                } else {
                    fmt.Println(v7, v3, v7)
                    fmt.Println(v7)
                    v26 := v2[0]
                    _ = v26
                }
                _ = fi25
            }
        }
    } else {
        var v27 float32 = 1.5
        v28 := v4["hello"]
        _ = v27
        _ = v28
    }
    switch 42 {
    case 4294967295:
        var v29 uint64 = 811
        _ = v29
    default:
        v3 = true
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
