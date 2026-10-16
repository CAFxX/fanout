package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 int16 }
type GenS2 struct { F0 int32; F1 string; F2 int }

func main() {
    var v1 bool = true
    v2 := []string{("foo bar" + ""), ("a" + "a"), "\n\t"}
    var v3 uint16 = 65535
    v4 := make(map[string]uint32)
    var v5 float64 = 0.0
    for fi6 := 0; fi6 < 11; fi6++ {
        v1 = v1
        var v7 uint = 957
        v7 += v7
        _ = fi6
        _ = v7
    }
    switch ("foo bar" + "\n\t") {
    case "\n\t":
        v5 = v5
    case "hello":
        fmt.Println(v5, v5, v1)
    }
    var v8 string = "foo bar"
    for fi9 := 0; fi9 < 4; fi9++ {
        if false {
            if len(v2) > 0 {
                v2[0] = v8
            }
            switch (v8 + (v8 + (v8 + ("a" + "a")))) {
            case "":
                if v1 {
                    v8 = ((v8 + ("hello" + v8)) + v8)
                    fmt.Println(v3, v1, v8)
                    fmt.Println(fi9, v3, v5)
                }
            case "\n\t":
                if true {
                    var v10 uint16 = v3
                    if v1 {
                        var v11 float64 = 3.14159
                        var v12 uint16 = v3
                        _ = v11
                        _ = v12
                    }
                    v2 = append(v2, "\n\t")
                    _ = v10
                }
            case "hello":
                v13 := GenS1{F0: 0, F1: (v3 + v3), F2: 32767}
                _ = v13
            default:
                var v14 uint64 = 291
                _ = v14
            }
        }
        _ = fi9
    }
    if (127 != 127) {
        var v15 int8 = -1
        for fi16 := 0; fi16 < 6; fi16++ {
            v2 = append(v2, "")
            var v17 int8 = v15
            _ = fi16
            _ = v17
        }
        var v18 int32 = 1
        _ = v15
        _ = v18
    }
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    if (2 < (^(v3))) {
        v3 += v3
        for fi19 := 0; fi19 < 7; fi19++ {
            var v20 float64 = 0.1
            _ = fi19
            _ = v20
        }
        v21 := GenS2{F0: 1, F1: (v8 + v8), F2: 1}
        _ = v21
    } else {
        var v22 int64 = 42
        _ = v22
    }
    v4[""] = 2147483648
    var v23 bool = true
    v5 -= v5
    v3 -= (v3 << ((0) & 63))
    v24 := GenS2{F0: 0, F1: v8, F2: 42}
    for fi25 := 0; fi25 < 19; fi25++ {
        if v23 {
            for fi26 := 0; fi26 < 16; fi26++ {
                v27 := GenS2{F0: 42, F1: ((v8 + v8) + ""), F2: fi25}
                fmt.Println(v1, fi26, v8)
                _ = fi26
                _ = v27
            }
            v5 = v5
        } else {
            v28 := v4[v8]
            fmt.Println(v23, v23)
            _ = v28
        }
        _ = fi25
    }
    var v29 bool = v23
    for fi30 := 0; fi30 < 13; fi30++ {
        for fi31 := 0; fi31 < 17; fi31++ {
            var v32 bool = true
            _ = fi31
            _ = v32
        }
        var v33 int64 = 26
        var v34 int = fi30
        _ = fi30
        _ = v33
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
