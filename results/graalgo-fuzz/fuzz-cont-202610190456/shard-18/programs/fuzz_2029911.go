package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 string; F1 int32 }
func genFn0(p0 bool, p1 string, p2 uint8) bool {
    if (723 < 2) {
        var v7 uint32 = 1
        switch 32767 {
        case 128:
            v7 -= 834
        case 2147483647:
            v7 = 881
        }
        var v8 int64 = 42
        _ = v7
        _ = v8
    }
    p0 = (!(!false))
    switch p1 {
    case "a":
        var v9 uint16 = 353
        _ = v9
    case "\n\t":
        var v10 uint = 65535
        _ = v10
    case "foo bar":
        if (!false) {
            var v11 float64 = 0.1
            if true {
                var v12 int16 = -343
                _ = v12
            }
            switch (p1 + (p1 + (p1 + ((p1 + (p1 + (p1 + p1))) + p1)))) {
            case "foo bar":
                p1 = "a"
            default:
                for fi13 := 0; fi13 < 18; fi13++ {
                    var v14 uint32 = 2147483648
                    v15 := []int{}
                    _ = fi13
                    _ = v14
                    _ = v15
                }
            }
            _ = v11
        }
    default:
        v16 := GenS1{F0: -760, F1: 42}
        _ = v16
    }
    var v17 bool = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return true
}
func genFn1(p0 int32, p1 int) int {
    for fi18 := 0; fi18 < 19; fi18++ {
        v19 := GenS1{F0: p0, F1: -60}
        var v20 uint8 = 1
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 string = "\n\t"
    p1 -= p1
    var v22 int64 = -1
    _ = p0
    _ = p1
    _ = v21
    _ = v22
    return 32768
}

func main() {
    var v1 int32 = 42
    v2 := make(map[string]bool)
    var v3 int64 = -9223372036854775808
    var v4 int64 = v3
    v5 := make(map[int]int)
    v6 := []int{}
    if (0 < 127) {
        for fi23 := 0; fi23 < 4; fi23++ {
            var v24 uint32 = 65535
            var v25 int16 = 42
            _ = fi23
            _ = v24
            _ = v25
        }
    } else {
        v26 := GenS2{F0: ("hello" + ("" + ((("foo bar" + ("" + "")) + "\n\t") + (((("foo bar" + "hello") + (("foo bar" + ("a" + "foo bar")) + (("" + "a") + "\n\t"))) + "\n\t") + "\n\t")))), F1: v1}
        v4 += (v3 * (v3 - (v4 * (v4 << ((279) & 63)))))
        _ = v26
    }
    var v27 int32 = v1
    switch (("\n\t" + "") + "hello") {
    case "hello":
        v28 := GenS2{F0: "", F1: v1}
        _ = v28
    case "a":
        var v29 string = (("\n\t" + ("\n\t" + "")) + "hello")
        _ = v29
    case "foo bar":
        v30 := v5[genFn1(1, -1)]
        _ = v30
    default:
        var v31 float32 = float32(-2.25)
        _ = v31
    }
    switch 32768 {
    case 2147483647:
        fmt.Println(v27, v3, v27)
    default:
        var v32 float32 = ((1.5 * -2.25) * 1.5)
        _ = v32
    }
    switch 0 {
    case 42:
        var v33 float64 = -2.25
        _ = v33
    case 32768:
        var v34 int16 = 32767
        _ = v34
    default:
        for fi35 := 0; fi35 < 14; fi35++ {
            var v36 int32 = (^(v27))
            _ = fi35
            _ = v36
        }
    }
    if (((!(!true)) && false) && false) {
        var v37 uint = 292
        if len(v6) > 0 {
            v6[0] = 128
        }
        _ = v37
    } else {
        v4 = -9223372036854775808
        for fi38 := 0; fi38 < 19; fi38++ {
            fmt.Println(v3, v4)
            fmt.Println(v3, v1)
            _ = fi38
        }
    }
    v39 := GenS2{F0: ("foo bar" + ("\n\t" + "")), F1: (v1 & (-(v1)))}
    v40 := v2["hello"]
    v27 += 0
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
