package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 int32; F1 int64 }

func main() {
    v1 := []uint32{0, 4294967295, 4294967295}
    v2 := []string{("\n\t" + "a"), "\n\t", "foo bar", (("\n\t" + "\n\t") + "\n\t")}
    var v3 bool = true
    v4 := []string{"hello", ((((("foo bar" + "hello") + (("\n\t" + "") + ("a" + ("foo bar" + (("" + "foo bar") + ("a" + "a")))))) + ("" + "")) + "foo bar") + "\n\t"), ("a" + ("" + "")), (((("a" + ("hello" + "")) + "\n\t") + "") + "")}
    var v5 int16 = 1000
    v6 := make(map[string]string)
    v7 := GenS1{F0: ("foo bar" + "hello"), F1: 2}
    fmt.Println(v5, v5)
    if (!(42 == 0)) {
        v3 = true
        var v8 int8 = 1
        _ = v8
    }
    var v9 int64 = -833
    fmt.Println(v5)
    v6["foo bar"] = "hello"
    var v10 uint32 = 2147483648
    if v3 {
        for fi11 := 0; fi11 < 6; fi11++ {
            v12 := GenS2{F0: -2147483648, F1: 505}
            v9 = 481
            _ = len(v4)
            _ = fi11
            _ = v12
        }
        for fi13 := 0; fi13 < 14; fi13++ {
            if (462 >= 0) {
                v14 := GenS3{F0: 0, F1: (v9 >> ((677) & 63))}
                var v15 int8 = 100
                _ = v14
                _ = v15
            } else {
                var v16 int8 = 0
                v9 = (-(v9))
                var v17 float32 = 1.5
                _ = v16
                _ = v17
            }
            v18 := GenS3{F0: -2147483648, F1: (v9 & v9)}
            v10 += v10
            _ = fi13
            _ = v18
        }
        for fi19 := 0; fi19 < 5; fi19++ {
            var v20 int8 = 1
            for fi21 := 0; fi21 < 6; fi21++ {
                var v22 uint16 = 501
                switch ("foo bar" + "\n\t") {
                case "hello":
                    if false {
                        v20 += (v20 / ((v20) | 1))
                        v6[""] = "a"
                    }
                default:
                    var v23 string = "\n\t"
                    _ = v23
                }
                v24 := GenS2{F0: -625, F1: 439}
                _ = fi21
                _ = v22
                _ = v24
            }
            var v25 uint64 = 4294967295
            _ = fi19
            _ = v20
            _ = v25
        }
    }
    if v3 {
        v3 = (v3 || v3)
    } else {
        v10 += v10
        var v26 int16 = v5
        v5 += 412
        _ = v26
    }
    for fi27 := 0; fi27 < 1; fi27++ {
        v5 = v5
        if v3 {
            var v28 int8 = -1
            if (!(!v3)) {
                switch -2147483648 {
                case 2147483647:
                    v9 -= v9
                }
                if (330 <= 42) {
                    var v29 int8 = v28
                    _ = v29
                }
                v30 := GenS1{F0: ("a" + ("" + ("\n\t" + "a"))), F1: 18446744073709551615}
                _ = v30
            }
            _ = v28
        }
        _ = fi27
    }
    if len(v2) > 0 {
        v2[0] = ("a" + "foo bar")
    }
    v5 = (^(v5))
    v5 -= (v5 >> ((18446744073709551615) & 63))
    var v31 uint32 = 800
    var v32 bool = ("hello" < (("foo bar" + "a") + ("hello" + "foo bar")))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
