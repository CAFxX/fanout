package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 int }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool; F1 uint32; F2 uint16 }

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[uint32]bool)
    var v3 int8 = -1
    v3 = v3
    v4 := GenS1{F0: 0, F1: 79, F2: -394}
    v3 = 127
    fmt.Println(v3, v3)
    var v5 int = -1
    var v6 uint32 = 0
    if false {
        if (!false) {
            for fi7 := 0; fi7 < 19; fi7++ {
                if true {
                    v6 += (v6 - v6)
                } else {
                    fmt.Println(v3, v5, v5)
                }
                _ = fi7
            }
            switch (v5 / ((v5) | 1)) {
            case 0:
                v8 := GenS2{F0: 306, F1: 865}
                _ = v8
            default:
                for fi9 := 0; fi9 < 9; fi9++ {
                    switch (("hello" + "") + (("foo bar" + "\n\t") + ("" + (("foo bar" + ((("a" + ((("hello" + (("\n\t" + "") + ("\n\t" + ((("hello" + "") + "hello") + ("" + "foo bar"))))) + "hello") + "\n\t")) + ("\n\t" + "\n\t")) + "")) + "\n\t")))) {
                    case "":
                        v10 := GenS2{F0: 1, F1: 481}
                        _ = v10
                    case "a":
                        fmt.Println(v1)
                    }
                    v11 := []int{-9223372036854775808}
                    var v12 int16 = 0
                    _ = fi9
                    _ = v11
                    _ = v12
                }
            }
            fmt.Println(v3, v3, v1)
        }
        for fi13 := 0; fi13 < 15; fi13++ {
            fmt.Println(v1, v5)
            _ = fi13
        }
        if (!((v1 << ((65535) & 63)) == v1)) {
            fmt.Println(v1, v3)
            v14 := []int{(^(v5)), (-(v5)), (-(v5))}
            v5 = (v5 ^ -1)
            _ = v14
        } else {
            if ((v5 % (((v5 >> ((42) & 63))) | 1)) == (v5 & (v5 ^ (v5 % ((v5) | 1))))) {
                if true {
                    v5 = (v5 + v5)
                }
                v15 := GenS3{F0: -77}
                _ = v15
            }
            v2[0] = (!(v3 <= (v3 - v3)))
        }
    } else {
        if false {
            v16 := GenS1{F0: 1, F1: 255, F2: int(v6)}
            _ = v16
        } else {
            var v17 float32 = 0.0
            v6 += (v6 * v6)
            _ = v17
        }
        var v18 uint = 7
        _ = v18
    }
    var v19 int = 2147483647
    var v20 uint32 = v6
    v21 := GenS4{F0: true, F1: 65535, F2: 416}
    v22 := []int{9223372036854775807}
    var v23 int32 = v1
    v1 += (-(v1))
    fmt.Println(v23, v19)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
