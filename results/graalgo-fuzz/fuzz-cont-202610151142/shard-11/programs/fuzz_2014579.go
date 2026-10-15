package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int }
type GenS2 struct { F0 string }

func main() {
    var v1 uint16 = 956
    var v2 uint64 = 42
    var v3 float32 = 1e10
    switch -2147483648 {
    case 511:
        switch -635 {
        case 9223372036854775807:
            var v4 float32 = float32(-2.25)
            _ = v4
        case 0:
            switch ((("hello" + "") + ("foo bar" + ("\n\t" + "hello"))) + ("hello" + "")) {
            case "\n\t":
                fmt.Println(v1)
            case "foo bar":
                v5 := GenS1{F0: 405, F1: -21, F2: -1}
                _ = v5
            case "hello":
                for fi6 := 0; fi6 < 8; fi6++ {
                    for fi7 := 0; fi7 < 19; fi7++ {
                        v8 := GenS1{F0: -9223372036854775808, F1: -32768, F2: (fi6 * (-(fi6)))}
                        var v9 float32 = (1.5 + v3)
                        var v10 int64 = 517
                        v10 = (v10 * 757)
                        var v11 uint8 = 42
                        v9 = v3
                        var v12 int8 = 99
                        var v13 int = fi7
                        v11 -= 127
                        _ = fi7
                        _ = v10
                        _ = v11
                        _ = v12
                        _ = v13
                        _ = v8
                        _ = v9
                    }
                    _ = fi6
                }
            }
        }
    case -9223372036854775808:
        fmt.Println(v3)
    case -2147483648:
        v3 -= (-2.25 * 0.1)
    }
    fmt.Println(v2)
    if false {
        v14 := []int{127, 548, -2147483648, 210}
        _ = v14
    }
    v15 := []int{32768, -261, 1, 127}
    for fi16 := 0; fi16 < 11; fi16++ {
        fmt.Println(v3)
        if (true || true) {
            _ = len(v15)
            v1 -= v1
            switch "a" {
            case "\n\t":
                if len(v15) >= 2 {
                    v15 = v15[1:len(v15)]
                }
            case "foo bar":
                if true {
                    var v17 int = fi16
                    for fi18 := 0; fi18 < 18; fi18++ {
                        v19 := GenS1{F0: 9223372036854775807, F1: -32768, F2: fi18}
                        v2 = v2
                        v1 -= v1
                        if (true || (544 == 455)) {
                            var v20 uint64 = 389
                            v15 = append(v15, (fi16 * -2147483648))
                            v2 = v2
                            _ = v20
                        }
                        var v21 float64 = 0.1
                        _ = fi18
                        _ = v19
                        _ = v21
                    }
                    _ = v17
                } else {
                    v22 := GenS1{F0: 1, F1: 783, F2: fi16}
                    _ = v22
                }
            case "hello":
                var v23 string = (("\n\t" + "hello") + "")
                _ = v23
            }
        }
        _ = fi16
    }
    v15 = append(v15, -487)
    if len(v15) > 0 {
        v15[0] = -10
    }
    var v24 int64 = 0
    switch -2147483648 {
    case 1:
        var v25 int32 = 42
        _ = v25
    }
    if true {
        switch ("a" + (((("" + ("hello" + "\n\t")) + "a") + ("foo bar" + "\n\t")) + "")) {
        case "":
            v3 = (float32(3.14159) + (v3 / (v3 / 1.5)))
        default:
            v26 := GenS2{F0: ""}
            _ = v26
        }
        if (!(42 > 0)) {
            var v27 uint64 = (v2 % ((v2) | 1))
            v1 = v1
            v15 = append(v15, -1)
            _ = v27
        } else {
            fmt.Println(v1)
            var v28 int = 32768
            _ = v28
        }
    } else {
        v29 := GenS2{F0: "a"}
        if (-2147483648 >= -1) {
            v30 := GenS1{F0: v24, F1: 0, F2: 32767}
            fmt.Println(v24, v3, v2)
            _ = v30
        }
        _ = v29
    }
    v2 = v2
    var v31 bool = true
    _ = len(v15)
    for fi32 := 0; fi32 < 8; fi32++ {
        for fi33 := 0; fi33 < 10; fi33++ {
            var v34 string = ("foo bar" + "foo bar")
            var v35 uint64 = (^(v2))
            v24 = (v24 / ((v24) | 1))
            _ = fi33
            _ = v34
            _ = v35
        }
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
        _ = fi32
    }
    for fi36 := 0; fi36 < 8; fi36++ {
        var v37 uint = 65535
        _ = fi36
        _ = v37
    }
    for fi38 := 0; fi38 < 17; fi38++ {
        if false {
            var v39 string = "\n\t"
            fmt.Println(v31, v2, fi38)
            switch (((("a" + v39) + ("a" + v39)) + "foo bar") + (("\n\t" + "a") + (v39 + v39))) {
            case "foo bar":
                var v40 int16 = -1
                _ = v40
            case "\n\t":
                var v41 string = (("" + "") + ("a" + "a"))
                _ = v41
            case "a":
                var v42 int64 = v24
                _ = v42
            default:
                v43 := GenS2{F0: v39}
                _ = v43
            }
            _ = v39
        } else {
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
            v24 = (v24 % ((int64(v2)) | 1))
        }
        if len(v15) > 0 {
            v15[0] = fi38
        }
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
        _ = fi38
    }
    v24 = v24
    fmt.Println(v24, v31, v24)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v15
    _ = v2
    _ = v24
    _ = v3
    _ = v31
}
