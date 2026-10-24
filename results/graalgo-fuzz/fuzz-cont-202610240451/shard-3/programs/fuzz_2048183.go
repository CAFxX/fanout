package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 int32; F1 uint64; F2 int16 }
type GenS4 struct { F0 int16; F1 uint16 }
type GenS5 struct { F0 uint64; F1 uint64; F2 uint8 }
func genFn0(p0 string) int8 {
    var v4 uint8 = 0
    var v5 int8 = 1
    if (0 > 1) {
        var v6 bool = true
        fmt.Println(v5, v4)
        v7 := []int{618, 1}
        _ = v6
        _ = v7
    }
    if (1.5 != 0.1) {
        v4 = 128
        if (!false) {
            v8 := []int{}
            v9 := GenS1{F0: false, F1: 794}
            _ = v8
            _ = v9
        } else {
            v10 := []int{}
            fmt.Println(p0)
            if (1 != 65535) {
                if true {
                    v11 := GenS2{F0: true, F1: 42}
                    v12 := GenS3{F0: 952, F1: 4294967295, F2: 1000}
                    for fi13 := 0; fi13 < 17; fi13++ {
                        if true {
                            var v14 bool = true
                            v5 = v5
                            fmt.Println(v14, v5, fi13)
                            _ = v14
                        } else {
                            var v15 float64 = 1e10
                            _ = v15
                        }
                        switch ((p0 + "hello") + "hello") {
                        case "":
                            fmt.Println(fi13, v4)
                        default:
                            p0 = "foo bar"
                        }
                        v5 += (v5 - (v5 ^ (v5 % ((-67) | 1))))
                        _ = fi13
                    }
                    _ = v11
                    _ = v12
                }
                var v16 int16 = 259
                var v17 int32 = 1
                _ = v16
                _ = v17
            }
            _ = v10
        }
    }
    _ = p0
    _ = v4
    _ = v5
    return v5
}
func genFn1() int64 {
    if (42 >= -928) {
        var v18 uint16 = 1
        _ = v18
    }
    v19 := GenS4{F0: 0, F1: 1}
    var v20 int64 = -199
    if true {
        var v21 int32 = -1
        v22 := GenS1{F0: false, F1: 32767}
        _ = v21
        _ = v22
    } else {
        var v23 bool = (4294967295 >= 32768)
        v24 := []int{}
        v23 = false
        _ = v23
        _ = v24
    }
    var v25 int16 = -32768
    _ = v19
    _ = v20
    _ = v25
    return 9223372036854775807
}

func main() {
    var v1 int32 = 1
    var v2 int64 = -725
    v3 := []string{"", "foo bar"}
    v26 := GenS2{F0: (0.0 != 1e10), F1: 42}
    v3 = append(v3, "\n\t")
    var v27 uint64 = 597
    switch 2147483647 {
    case -736:
        if len(v3) > 0 {
            v3[0] = "a"
        }
    case 32767:
        if len(v3) > 0 {
            v3[0] = ("foo bar" + ("foo bar" + "foo bar"))
        }
    }
    for fi28 := 0; fi28 < 18; fi28++ {
        for fi29 := 0; fi29 < 2; fi29++ {
            var v30 bool = true
            _ = fi29
            _ = v30
        }
        _ = fi28
    }
    v1 = (v1 * (v1 >> ((65535) & 63)))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (!(!(true || ((genFn0("a") < 98) || true)))) {
        v31 := GenS1{F0: true, F1: 817}
        _ = v31
    } else {
        fmt.Println(v1, v27, v27)
        var v32 string = "hello"
        for fi33 := 0; fi33 < 15; fi33++ {
            v1 += v1
            v34 := GenS2{F0: false, F1: 32767}
            var v35 float32 = (3.14159 - ((-2.25 - 1.5) + 3.14159))
            _ = fi33
            _ = v34
            _ = v35
        }
        _ = v32
    }
    v2 = genFn1()
    v36 := GenS4{F0: 704, F1: 65535}
    for fi37 := 0; fi37 < 8; fi37++ {
        switch (fi37 | (fi37 & (fi37 * fi37))) {
        case -1:
            for fi38 := 0; fi38 < 16; fi38++ {
                var v39 int8 = -7
                var v40 string = "hello"
                if false {
                    if false {
                        if (false && false) {
                            v41 := GenS4{F0: 1000, F1: 228}
                            v40 = "foo bar"
                            v27 = (v27 | v27)
                            _ = v41
                        }
                        if false {
                            var v42 int16 = 822
                            v43 := GenS5{F0: (v27 * (v27 - (v27 >> ((42) & 63)))), F1: v27, F2: 90}
                            if false {
                                v2 += v2
                            }
                            _ = v42
                            _ = v43
                        } else {
                            fmt.Println(v40, v27)
                            v3 = append(v3, "")
                        }
                    }
                    var v44 int8 = (v39 | v39)
                    _ = v44
                }
                _ = fi38
                _ = v39
                _ = v40
            }
        }
        fmt.Println(v27, v1)
        _ = fi37
    }
    for fi45 := 0; fi45 < 8; fi45++ {
        var v46 int = fi45
        var v47 bool = (128 == 42)
        fmt.Println(v27)
        _ = fi45
        _ = v46
        _ = v47
    }
    switch "" {
    case "foo bar":
        for fi48 := 0; fi48 < 5; fi48++ {
            var v49 uint32 = 2
            v49 = v49
            _ = fi48
            _ = v49
        }
    case "\n\t":
        fmt.Println(v27, v2, v2)
    case "hello":
        v1 = (v1 << ((42) & 63))
    }
    for fi50 := 0; fi50 < 6; fi50++ {
        for fi51 := 0; fi51 < 13; fi51++ {
            _ = len(v3)
            var v52 string = ("" + "\n\t")
            _ = fi51
            _ = v52
        }
        var v53 float32 = (1e10 + (1.5 - (0.0 - 3.14159)))
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi50
        _ = v53
    }
    for fi54 := 0; fi54 < 18; fi54++ {
        v55 := GenS2{F0: true, F1: 42}
        _ = fi54
        _ = v55
    }
    var v56 int8 = 68
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v36
    _ = v56
}
