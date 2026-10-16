package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 uint32 }

func main() {
    var v1 float32 = 1e10
    var v2 int64 = 1
    v3 := []string{("foo bar" + "a"), "a", "", "hello"}
    var v4 uint8 = 0
    v5 := []int{638, 2147483647, -978, 331}
    var v6 int8 = 127
    if len(v5) > 0 {
        v5[0] = -328
    }
    v7 := GenS1{F0: (v4 + v4), F1: 32768}
    _ = len(v5)
    v8 := GenS2{F0: 42}
    if true {
        _ = len(v3)
        for fi9 := 0; fi9 < 2; fi9++ {
            if (false || false) {
                v1 -= v1
                v6 = v6
            } else {
                var v10 float64 = (float64(v1) + 1e10)
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = v10
            }
            for fi11 := 0; fi11 < 19; fi11++ {
                var v12 uint64 = 2
                v1 += (float32(-2.25) - (float32(1.5) / -2.25))
                _ = fi11
                _ = v12
            }
            for fi13 := 0; fi13 < 9; fi13++ {
                var v14 string = ""
                if (true && (v1 < (v1 + (-2.25 * 0.0)))) {
                    var v15 uint32 = 2
                    v15 -= v15
                    _ = v15
                } else {
                    if len(v3) > 0 {
                        v3[0] = (v14 + ("\n\t" + (v14 + v14)))
                    }
                }
                fmt.Println(v4, v2)
                _ = fi13
                _ = v14
            }
            _ = fi9
        }
        fmt.Println(v1)
    } else {
        if false {
            v2 = 604
            v6 += (v6 | 42)
        } else {
            _ = len(v3)
            if true {
                if true {
                    var v16 int8 = 0
                    _ = v16
                } else {
                    var v17 float32 = 0.0
                    _ = v17
                }
            } else {
                fmt.Println(v6, v2, v6)
            }
        }
        if (!(false || (-2147483648 < -662))) {
            v3 = append(v3, "foo bar")
            switch ("hello" + "hello") {
            case "foo bar":
                for fi18 := 0; fi18 < 12; fi18++ {
                    v2 = 0
                    _ = fi18
                }
            case "\n\t":
                v19 := GenS1{F0: (v4 + (v4 << ((0) & 63))), F1: 9223372036854775807}
                _ = v19
            case "hello":
                var v20 int = -2147483648
                _ = v20
            default:
                fmt.Println(v6)
            }
        } else {
            var v21 int8 = v6
            _ = v21
        }
    }
    v22 := GenS3{F0: v6, F1: 2147483648}
    for fi23 := 0; fi23 < 14; fi23++ {
        if true {
            v24 := GenS2{F0: -1}
            _ = v24
        } else {
            v6 += -39
            var v25 bool = (687 == 397)
            _ = v25
        }
        if len(v3) > 0 {
            v3[0] = "hello"
        }
        if len(v3) > 0 {
            v3[0] = (((("hello" + (("hello" + ("\n\t" + ("\n\t" + "a"))) + "\n\t")) + "hello") + (("hello" + "") + "")) + "")
        }
        _ = fi23
    }
    for fi26 := 0; fi26 < 20; fi26++ {
        v27 := GenS3{F0: v6, F1: 0}
        _ = fi26
        _ = v27
    }
    v28 := GenS3{F0: (v6 | v6), F1: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
