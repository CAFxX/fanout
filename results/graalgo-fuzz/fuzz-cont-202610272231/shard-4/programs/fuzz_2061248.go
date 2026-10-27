package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 int8; F2 uint }

func main() {
    var v1 string = ("foo bar" + "")
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]bool)
    v4 := []int{-876}
    for fi5 := 0; fi5 < 1; fi5++ {
        fmt.Println(fi5)
        _ = fi5
    }
    if (2 < 2) {
        for fi6 := 0; fi6 < 6; fi6++ {
            switch v1 {
            case "\n\t":
                var v7 int8 = -63
                _ = v7
            }
            if (-1 != 0) {
                if false {
                    fmt.Println(v1, fi6)
                } else {
                    v1 = v1
                    if true {
                        if len(v4) > 0 {
                            v4[0] = (fi6 & fi6)
                        }
                        v8 := GenS1{F0: 42}
                        _ = v8
                    }
                }
            } else {
                if false {
                    v9 := v3[65535]
                    var v10 int = fi6
                    for fi11 := 0; fi11 < 8; fi11++ {
                        var v12 uint = 18446744073709551615
                        v13 := GenS2{F0: 42, F1: int8(v12), F2: v12}
                        switch v10 {
                        case 456:
                            v4 = append(v4, (v10 & (fi11 * 42)))
                        }
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                    _ = v10
                    _ = v9
                } else {
                    fmt.Println(v1)
                    if false {
                        var v14 float64 = ((0.1 - 0.0) + -2.25)
                        _ = v14
                    } else {
                        v15 := GenS2{F0: 42, F1: 127, F2: 492}
                        _ = v15
                    }
                }
            }
            _ = fi6
        }
    } else {
        for fi16 := 0; fi16 < 7; fi16++ {
            v17 := GenS1{F0: 2147483647}
            var v18 uint64 = 4294967295
            _ = fi16
            _ = v17
            _ = v18
        }
    }
    switch -409 {
    case -1:
        var v19 float32 = 0.1
        _ = v19
    case 38:
        v20 := GenS2{F0: -1, F1: 127, F2: 332}
        _ = v20
    }
    v21 := v2[65535]
    v22 := GenS1{F0: -2147483648}
    var v23 int8 = -51
    v4 = append(v4, -567)
    switch 9223372036854775807 {
    case 241:
        _ = len(v4)
    case -9223372036854775808:
        if false {
            var v24 string = (v1 + v21)
            if (true && true) {
                var v25 uint = 899
                var v26 int8 = v23
                _ = v25
                _ = v26
            } else {
                var v27 uint16 = 42
                _ = v27
            }
            var v28 uint64 = 398
            _ = v24
            _ = v28
        } else {
            var v29 float32 = 1.5
            v23 = v23
            fmt.Println(v21, v1, v29)
            _ = v29
        }
    default:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    v21 = "foo bar"
    if false {
        var v30 int32 = 177
        switch (v1 + "hello") {
        case "foo bar":
            if false {
                v30 = v30
                v30 = v30
            }
        case "":
            var v31 int16 = -446
            _ = v31
        }
        for fi32 := 0; fi32 < 1; fi32++ {
            for fi33 := 0; fi33 < 11; fi33++ {
                v34 := GenS2{F0: 190, F1: (v23 * v23), F2: 42}
                v4 = append(v4, fi32)
                _ = fi33
                _ = v34
            }
            _ = fi32
        }
        _ = v30
    } else {
        v35 := GenS1{F0: -371}
        _ = v35
    }
    v36 := GenS1{F0: -733}
    v21 = v1
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v36
    _ = v4
}
