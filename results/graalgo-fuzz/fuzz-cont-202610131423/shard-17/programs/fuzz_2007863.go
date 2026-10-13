package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 string }
type GenS2 struct { F0 bool; F1 uint32 }

func main() {
    v1 := []string{"hello"}
    v2 := make(map[int]string)
    var v3 float64 = ((0.0 - 1e10) / 1.0)
    var v4 int = -743
    v5 := make(map[int]string)
    var v6 string = "foo bar"
    if false {
        if (42 != -1) {
            var v7 int32 = 343
            var v8 uint16 = 810
            _ = v7
            _ = v8
        } else {
            if false {
                for fi9 := 0; fi9 < 4; fi9++ {
                    var v10 int = v4
                    if false {
                        fmt.Println(v3)
                    } else {
                        v10 -= v4
                        if (v6 <= v6) {
                            fmt.Println(v4, v3)
                        }
                    }
                    _ = fi9
                    _ = v10
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v11 := GenS1{F0: ("" + v6), F1: v6, F2: (("hello" + v6) + (v6 + v6))}
                _ = v11
            } else {
                if (!true) {
                    v1 = append(v1, v6)
                    v4 = 32768
                }
                fmt.Println(v6, v3)
                if true {
                    var v12 uint16 = 255
                    _ = v12
                }
            }
        }
        switch "hello" {
        case "foo bar":
            if len(v1) > 0 {
                v1[0] = ""
            }
        case "":
            var v13 int64 = 814
            _ = v13
        case "hello":
            v3 = (v3 / ((float64(3.14159) - (v3 - ((0.0 - (v3 + v3)) * (((v3 - v3) - v3) * v3)))) / v3))
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        if true {
            v14 := GenS1{F0: v6, F1: ("a" + v6), F2: "a"}
            _ = v14
        }
        v15 := GenS1{F0: ((v6 + v6) + v6), F1: v6, F2: "a"}
        v4 -= v4
        _ = v15
    }
    fmt.Println(v3)
    for fi16 := 0; fi16 < 20; fi16++ {
        for fi17 := 0; fi17 < 12; fi17++ {
            fmt.Println(fi17, v6, fi17)
            _ = fi17
        }
        _ = fi16
    }
    switch (v4 % ((v4) | 1)) {
    case -269:
        var v18 uint8 = 255
        _ = v18
    }
    var v19 bool = (143 < 65535)
    delete(v5, v4)
    if (-400 >= -2147483648) {
        v20 := GenS1{F0: v6, F1: v6, F2: v6}
        v4 -= (v4 - v4)
        _ = v20
    } else {
        var v21 uint = 65535
        if (!(!(255 > 32767))) {
            var v22 float32 = (float32(1e10) + -2.25)
            _ = v22
        }
        _ = v21
    }
    v23 := GenS2{F0: true, F1: 2}
    var v24 uint8 = 61
    var v25 int32 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
