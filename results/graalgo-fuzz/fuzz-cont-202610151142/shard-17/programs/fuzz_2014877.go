package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 int16; F1 string }
func genFn0(p0 int, p1 int16, p2 int) int64 {
    fmt.Println(p2, p0, p1)
    v6 := []int{(p2 % ((p2) | 1)), p2, (p0 / (((p2 * p2)) | 1)), p2}
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := GenS1{F0: 42, F1: 30}
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        var v10 int16 = 32767
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -9223372036854775808
}

func main() {
    v1 := make(map[string]string)
    var v2 uint16 = 58
    var v3 float64 = -2.25
    var v4 int64 = 42
    v5 := []int{}
    var v11 uint32 = 2147483648
    var v12 uint = 65535
    if (true || (v12 <= (^(v12)))) {
        v12 = 42
        if ((v12 + 1) != v12) {
            if (true && ((1e10 + (0.0 * 0.0)) == 0.1)) {
                switch "" {
                case "":
                    if true {
                        var v13 uint8 = 23
                        _ = v13
                    }
                case "\n\t":
                    var v14 int16 = 844
                    _ = v14
                case "a":
                    var v15 int = -1
                    _ = v15
                }
            }
        }
        var v16 float32 = -2.25
        _ = v16
    } else {
        v11 -= v11
    }
    var v17 uint64 = 18446744073709551615
    delete(v1, "")
    if (((v11 - v11) == (v11 >> ((4294967295) & 63))) || true) {
        switch "\n\t" {
        case "":
            v1[""] = "\n\t"
        case "\n\t":
            v2 -= v2
        default:
            v17 += (v17 << ((v12) & 63))
        }
        if true {
            var v18 uint64 = (v17 & 720)
            switch 32767 {
            case -9223372036854775808:
                v4 = (v4 % ((v4) | 1))
            default:
                var v19 uint32 = uint32(v2)
                _ = v19
            }
            _ = len(v5)
            _ = v18
        } else {
            var v20 int64 = 1
            _ = v20
        }
        for fi21 := 0; fi21 < 7; fi21++ {
            v2 = 32767
            _ = fi21
        }
    } else {
        v2 += (v2 << (((v12 & uint(v2))) & 63))
    }
    v1[((("foo bar" + ((("" + ((("foo bar" + ("\n\t" + "\n\t")) + "") + "hello")) + "hello") + (("foo bar" + "\n\t") + ""))) + "") + (((("\n\t" + (("\n\t" + ("a" + "a")) + ("hello" + "hello"))) + ("a" + ("a" + "a"))) + "a") + "\n\t"))] = ("\n\t" + "")
    for fi22 := 0; fi22 < 20; fi22++ {
        v23 := GenS2{F0: 42, F1: ("foo bar" + "\n\t")}
        v11 -= uint32(v17)
        if true {
            fmt.Println(v3)
            v4 = 42
            if true {
                var v24 float32 = -2.25
                _ = v24
            }
        } else {
            var v25 int32 = 0
            v26 := GenS1{F0: 43, F1: 56}
            for fi27 := 0; fi27 < 16; fi27++ {
                var v28 uint = v12
                var v29 float32 = 1.5
                v17 += v17
                v25 = v25
                _ = fi27
                _ = v28
                _ = v29
            }
            _ = v25
            _ = v26
        }
        _ = fi22
        _ = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
