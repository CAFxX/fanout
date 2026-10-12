package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 int8 }
func genFn0(p0 int8, p1 int32) int16 {
    for fi5 := 0; fi5 < 3; fi5++ {
        p0 = (-(p0))
        _ = fi5
    }
    var v6 float64 = 0.1
    _ = p0
    _ = p1
    _ = v6
    return 42
}

func main() {
    v1 := make(map[string]bool)
    v2 := []int{}
    var v3 bool = false
    var v4 uint64 = 4294967295
    switch "foo bar" {
    case "foo bar":
        v7 := GenS1{F0: 42, F1: v4, F2: 0}
        _ = v7
    case "a":
        switch "a" {
        case "foo bar":
            delete(v1, "foo bar")
        case "hello":
            switch 643 {
            case -590:
                v4 = (v4 << ((2) & 63))
            default:
                fmt.Println(v3)
            }
        case "a":
            if true {
                v1["\n\t"] = v3
                fmt.Println(v3, v3)
                var v8 uint = 18446744073709551615
                _ = v8
            }
        default:
            fmt.Println(v4, v3)
        }
    case "hello":
        if (1.5 >= 1.5) {
            var v9 float64 = ((0.1 + float64((float32(0.0) - 0.1))) + (((1e10 / 3.14159) - (0.0 * float64((0.0 - (-2.25 + ((3.14159 / (((((1.5 + 1.5) + 1e10) * 3.14159) / 1.5) - (0.0 / (0.0 + 1e10)))) * 1e10)))))) / 3.14159))
            var v10 string = "foo bar"
            fmt.Println(v9, v9, v9)
            _ = v10
            _ = v9
        } else {
            for fi11 := 0; fi11 < 10; fi11++ {
                fmt.Println(fi11, v4)
                var v12 uint8 = 0
                var v13 uint16 = 0
                _ = fi11
                _ = v12
                _ = v13
            }
            var v14 bool = (!((!v3) || (-1 <= 32767)))
            v1["foo bar"] = true
            _ = v14
        }
    default:
        if v3 {
            if v3 {
                v4 = 2
                v15 := GenS1{F0: 18446744073709551615, F1: v4, F2: 100}
                _ = v15
            }
            if false {
                for fi16 := 0; fi16 < 18; fi16++ {
                    v3 = v3
                    var v17 bool = (int16(v4) <= -916)
                    _ = fi16
                    _ = v17
                }
            } else {
                v3 = (v3 && (614 >= (v4 - v4)))
                var v18 int16 = -32768
                v18 -= v18
                _ = v18
            }
            v19 := GenS1{F0: 65535, F1: v4, F2: 42}
            _ = v19
        }
    }
    if false {
        var v20 float64 = float64(3.14159)
        v21 := GenS1{F0: 4294967295, F1: v4, F2: 1}
        _ = v20
        _ = v21
    } else {
        v4 = (v4 << ((18446744073709551615) & 63))
        v22 := GenS1{F0: 312, F1: v4, F2: -79}
        _ = v22
    }
    if true {
        var v23 uint16 = 204
        _ = v23
    }
    for fi24 := 0; fi24 < 3; fi24++ {
        switch ("\n\t" + "\n\t") {
        case "":
            var v25 float32 = -2.25
            _ = v25
        case "hello":
            switch "" {
            case "":
                fmt.Println(v3)
            default:
                var v26 int32 = 2147483647
                _ = v26
            }
        case "\n\t":
            var v27 int64 = 9223372036854775807
            _ = v27
        }
        fmt.Println(v3)
        var v28 int64 = 862
        _ = fi24
        _ = v28
    }
    var v29 uint8 = 2
    if v3 {
        for fi30 := 0; fi30 < 3; fi30++ {
            for fi31 := 0; fi31 < 6; fi31++ {
                v3 = (!((!(255 > 598)) && (("a" + (("a" + (("" + "") + ("\n\t" + "hello"))) + "a")) > (("foo bar" + "foo bar") + (("foo bar" + "a") + ("foo bar" + ("a" + ("" + ("hello" + "foo bar")))))))))
                _ = fi31
            }
            for fi32 := 0; fi32 < 1; fi32++ {
                delete(v1, "hello")
                var v33 uint64 = (^(v4))
                fmt.Println(v33, fi32, v33)
                _ = fi32
                _ = v33
            }
            fmt.Println(fi30, fi30, fi30)
            _ = fi30
        }
        _ = len(v1)
        v3 = (1 < -396)
    }
    if len(v2) > 0 {
        v2[0] = -287
    }
    v2 = append(v2, 490)
    v29 = (v29 % ((42) | 1))
    v3 = true
    var v34 int32 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v34
    _ = v4
}
