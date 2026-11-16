package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 uint16 }

func main() {
    v1 := []string{("" + ((("hello" + "") + "a") + "\n\t")), "", "a"}
    v2 := []string{"foo bar"}
    v3 := make(map[string]string)
    var v4 float32 = 1e10
    for fi5 := 0; fi5 < 12; fi5++ {
        fmt.Println(v4)
        for fi6 := 0; fi6 < 9; fi6++ {
            fmt.Println(v4, v4, fi6)
            var v7 float64 = float64(v4)
            var v8 int16 = -442
            _ = fi6
            _ = v7
            _ = v8
        }
        v2 = append(v2, ("hello" + "a"))
        _ = fi5
    }
    fmt.Println(v4)
    fmt.Println(v4)
    v9 := GenS1{F0: -316, F1: 460, F2: 32767}
    switch (("hello" + ("foo bar" + "foo bar")) + "foo bar") {
    case "\n\t":
        var v10 int32 = -1
        _ = v10
    default:
        if (v4 <= v4) {
            for fi11 := 0; fi11 < 17; fi11++ {
                for fi12 := 0; fi12 < 20; fi12++ {
                    fmt.Println(fi12, fi11)
                    v13 := GenS1{F0: -2147483648, F1: 678, F2: 42}
                    v14 := GenS1{F0: 0, F1: 0, F2: 65535}
                    var v15 int16 = 804
                    v2 = append(v2, "")
                    _ = fi12
                    _ = v13
                    _ = v14
                    _ = v15
                }
                for fi16 := 0; fi16 < 14; fi16++ {
                    v4 += (v4 - (v4 + v4))
                    _ = fi16
                }
                if (false && true) {
                    v17 := v3[("foo bar" + "")]
                    fmt.Println(v4)
                    for fi18 := 0; fi18 < 19; fi18++ {
                        v3["hello"] = ("a" + v17)
                        _ = fi18
                    }
                    _ = v17
                }
                _ = fi11
            }
            if true {
                for fi19 := 0; fi19 < 15; fi19++ {
                    var v20 int32 = -2147483648
                    _ = fi19
                    _ = v20
                }
            } else {
                v4 = (v4 * 3.14159)
            }
            delete(v3, ((("" + "hello") + "foo bar") + "\n\t"))
        }
    }
    var v21 uint64 = 2
    if false {
        v22 := GenS1{F0: 0, F1: 932, F2: 409}
        var v23 int64 = 1
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v22
        _ = v23
    }
    if true {
        var v24 string = (("\n\t" + "") + "")
        var v25 float64 = 1e10
        _ = v24
        _ = v25
    }
    v21 -= 1
    v26 := GenS1{F0: -854, F1: 42, F2: 707}
    v27 := GenS1{F0: -2147483648, F1: 2147483647, F2: 889}
    if (100 < 127) {
        switch (("" + "") + "foo bar") {
        case "":
            if len(v1) > 0 {
                v1[0] = "a"
            }
        case "foo bar":
            v28 := v3["foo bar"]
            _ = v28
        case "hello":
            switch -460 {
            case -127:
                v21 = 669
            case 1:
                v21 -= (v21 & (v21 - 1))
            case -28:
                var v29 int16 = 1
                _ = v29
            default:
                var v30 int64 = 1
                _ = v30
            }
        }
        v21 += (^(v21))
        v31 := GenS1{F0: -2147483648, F1: 4294967295, F2: 2}
        _ = v31
    }
    v32 := GenS1{F0: 202, F1: 42, F2: 212}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v26
    _ = v27
    _ = v3
    _ = v32
    _ = v4
    _ = v9
}
