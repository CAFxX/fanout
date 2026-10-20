package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 string }
type GenS2 struct { F0 int64; F1 string }
func genFn0(p0 uint16) uint {
    if ((true || false) && (0 == 90)) {
        if true {
            for fi6 := 0; fi6 < 6; fi6++ {
                p0 = 42
                for fi7 := 0; fi7 < 9; fi7++ {
                    v8 := GenS1{F0: "a", F1: "hello", F2: "hello"}
                    v9 := GenS1{F0: ("foo bar" + "foo bar"), F1: "a", F2: "\n\t"}
                    _ = fi7
                    _ = v8
                    _ = v9
                }
                v10 := GenS1{F0: ("" + "a"), F1: "", F2: (("\n\t" + "\n\t") + ("" + ("\n\t" + "")))}
                _ = fi6
                _ = v10
            }
            var v11 uint32 = 2147483648
            _ = v11
        } else {
            v12 := GenS2{F0: 9223372036854775807, F1: "hello"}
            _ = v12
        }
    }
    if true {
        var v13 float64 = -2.25
        _ = v13
    }
    p0 += p0
    p0 -= p0
    p0 -= 182
    _ = p0
    return 154
}

func main() {
    v1 := make(map[int]bool)
    var v2 string = ("\n\t" + (("foo bar" + "foo bar") + ("\n\t" + ("a" + ""))))
    var v3 uint16 = 338
    var v4 float32 = 0.0
    v5 := []string{((((v2 + "") + v2) + ((v2 + v2) + v2)) + v2), "hello", v2, v2}
    delete(v1, -2147483648)
    var v14 uint16 = (v3 | 0)
    var v15 int64 = 1
    var v16 uint64 = 733
    if (!false) {
        if false {
            fmt.Println(v3, v15)
            var v17 string = "\n\t"
            var v18 int64 = (-(v15))
            _ = v17
            _ = v18
        } else {
            var v19 bool = (18446744073709551615 != v16)
            var v20 bool = v19
            if false {
                for fi21 := 0; fi21 < 8; fi21++ {
                    fmt.Println(fi21, v14, v3)
                    fmt.Println(v20)
                    _ = fi21
                }
            }
            _ = v19
            _ = v20
        }
        var v22 int8 = -128
        _ = v22
    } else {
        delete(v1, -1)
    }
    v4 -= (v4 + 0.1)
    for fi23 := 0; fi23 < 14; fi23++ {
        for fi24 := 0; fi24 < 11; fi24++ {
            v15 = 42
            _ = fi24
        }
        _ = fi23
    }
    if true {
        var v25 bool = (!(2 > 255))
        _ = v25
    }
    if ((v4 != 0.0) || true) {
        var v26 uint8 = 53
        var v27 uint32 = 814
        if (!(!((v27 <= (v27 - (v27 % ((v27) | 1)))) && true))) {
            if (991 == -537) {
                for fi28 := 0; fi28 < 3; fi28++ {
                    delete(v1, (fi28 >> ((4294967295) & 63)))
                    v29 := GenS1{F0: v2, F1: "", F2: v2}
                    _ = fi28
                    _ = v29
                }
                if false {
                    v15 = (v15 - (v15 - v15))
                    fmt.Println(v2, v15, v26)
                    v14 = (^(v14))
                }
            } else {
                var v30 int8 = -1
                v15 += -194
                if len(v5) > 0 {
                    v5[0] = (v2 + ((((("hello" + v2) + v2) + v2) + v2) + v2))
                }
                _ = v30
            }
        }
        _ = v26
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
