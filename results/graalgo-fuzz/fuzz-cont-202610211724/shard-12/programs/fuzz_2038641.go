package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 string }
type GenS2 struct { F0 bool }
func genFn0(p0 string) int {
    for fi7 := 0; fi7 < 10; fi7++ {
        p0 = p0
        v8 := []int{-530}
        var v9 uint64 = 0
        _ = fi7
        _ = v8
        _ = v9
    }
    fmt.Println(p0)
    p0 = p0
    v10 := GenS1{F0: 0, F1: 65535, F2: ((p0 + (p0 + "a")) + "a")}
    _ = p0
    _ = v10
    return 127
}
func genFn1() int32 {
    var v11 uint32 = 393
    for fi12 := 0; fi12 < 7; fi12++ {
        for fi13 := 0; fi13 < 8; fi13++ {
            v14 := []int{(fi13 % (((fi13 >> ((65535) & 63))) | 1)), (-(fi12)), -39, (fi12 & 9223372036854775807)}
            _ = fi13
            _ = v14
        }
        v15 := GenS2{F0: false}
        if (false && false) {
            v11 = v11
            v11 = (v11 | 4294967295)
        } else {
            if ((false || false) || (!(127 < 102))) {
                v16 := []int{fi12}
                _ = v16
            } else {
                var v17 uint8 = 128
                if false {
                    switch ("hello" + ("\n\t" + "\n\t")) {
                    case "hello":
                        var v18 float32 = 1.5
                        _ = v18
                    }
                    v19 := []int{-693, -9223372036854775808, -697, (fi12 << ((2) & 63))}
                    _ = v19
                } else {
                    if ((-466 < 0) && (v17 != (v17 - 9))) {
                        v20 := []int{(fi12 / (((-(fi12))) | 1)), -2147483648}
                        var v21 uint64 = 1
                        _ = v20
                        _ = v21
                    }
                }
                _ = v17
            }
            v11 = v11
        }
        _ = fi12
        _ = v15
    }
    var v22 uint32 = v11
    _ = v11
    _ = v22
    return -301
}

func main() {
    v1 := make(map[string]bool)
    var v2 int32 = -477
    var v3 uint16 = 42
    var v4 int = 2147483647
    v5 := make(map[string]bool)
    v6 := make(map[string]bool)
    v23 := GenS1{F0: v2, F1: 2, F2: ("a" + "\n\t")}
    var v24 uint32 = 2147483647
    v25 := []int{-9223372036854775808, (-(v4)), (v4 << ((65535) & 63))}
    v26 := GenS1{F0: (v2 | v2), F1: 785, F2: "a"}
    v1["foo bar"] = true
    if ((5 <= 0) || (0.1 == 1.5)) {
        for fi27 := 0; fi27 < 17; fi27++ {
            var v28 string = "foo bar"
            var v29 int64 = -9223372036854775808
            v3 -= 473
            _ = fi27
            _ = v28
            _ = v29
        }
        v24 += 579
        var v30 float32 = ((3.14159 + float32(3.14159)) / -2.25)
        _ = v30
    } else {
        v2 = -927
        fmt.Println(v3, v24, v3)
    }
    var v31 uint = 255
    v32 := GenS2{F0: true}
    fmt.Println(v31, v3, v4)
    v1["a"] = (1 > 42)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
