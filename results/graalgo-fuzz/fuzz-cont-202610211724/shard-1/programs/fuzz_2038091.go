package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 uint }
func genFn0(p0 uint8) bool {
    var v5 int32 = 2147483647
    var v6 int64 = -549
    _ = p0
    _ = v5
    _ = v6
    return (32767 == 32767)
}

func main() {
    var v1 uint16 = 543
    var v2 int64 = -9223372036854775808
    var v3 uint16 = v1
    var v4 bool = true
    switch 1 {
    case 4294967295:
        switch "\n\t" {
        case "hello":
            if ((-92 < 2147483647) && (512 > 760)) {
                v2 = (v2 * v2)
                for fi7 := 0; fi7 < 20; fi7++ {
                    v8 := []int{}
                    v8 = append(v8, -9223372036854775808)
                    v9 := GenS1{F0: -2147483648, F1: 1, F2: 46}
                    _ = fi7
                    _ = v8
                    _ = v9
                }
            }
        case "":
            if v4 {
                for fi10 := 0; fi10 < 5; fi10++ {
                    var v11 int8 = 0
                    _ = fi10
                    _ = v11
                }
                if true {
                    fmt.Println(v4, v4)
                    var v12 string = "a"
                    _ = v12
                }
            }
        }
    case 1:
        if (v3 < 0) {
            v3 = v1
            v13 := []int{919, 4294967295, 2147483647}
            _ = v13
        } else {
            if (!v4) {
                v14 := GenS1{F0: 1, F1: (v2 * (^(v2))), F2: 4294967295}
                v4 = v4
                var v15 int32 = 226
                _ = v14
                _ = v15
            }
            switch ("a" + "\n\t") {
            case "foo bar":
                fmt.Println(v1)
            case "hello":
                v4 = v4
            default:
                v2 -= (v2 & (v2 & (v2 % (((v2 * v2)) | 1))))
            }
            for fi16 := 0; fi16 < 3; fi16++ {
                var v17 bool = v4
                fmt.Println(v2, v4, v4)
                _ = fi16
                _ = v17
            }
        }
    case 128:
        for fi18 := 0; fi18 < 19; fi18++ {
            v4 = (false || (!(370 <= 42)))
            var v19 int64 = (v2 * (v2 % (((v2 | (v2 ^ -1))) | 1)))
            _ = fi18
            _ = v19
        }
    }
    v20 := []int{4294967295, int(v1)}
    fmt.Println(v2)
    v1 = 255
    if false {
        switch "a" {
        case "":
            var v21 float64 = 1.5
            _ = v21
        case "foo bar":
            v20 = append(v20, -1)
        default:
            var v22 string = ((("foo bar" + "hello") + ("\n\t" + ("\n\t" + (("\n\t" + ((("foo bar" + (("" + ("a" + "a")) + ("" + "a"))) + "hello") + "foo bar")) + (("\n\t" + "\n\t") + "a"))))) + "foo bar")
            _ = v22
        }
        if len(v20) >= 2 {
            v20 = v20[1:len(v20)]
        }
    }
    if (true || (-109 <= -422)) {
        for fi23 := 0; fi23 < 4; fi23++ {
            fmt.Println(v2, v2)
            v4 = v4
            _ = fi23
        }
        if len(v20) > 0 {
            v20[0] = 42
        }
    } else {
        v24 := GenS1{F0: 0, F1: v2, F2: 2}
        v25 := GenS1{F0: 958, F1: v2, F2: 935}
        v3 = v3
        _ = v24
        _ = v25
    }
    if len(v20) > 0 {
        v20[0] = 128
    }
    fmt.Println(v3)
    for fi26 := 0; fi26 < 15; fi26++ {
        fmt.Println(v2)
        _ = fi26
    }
    if false {
        for fi27 := 0; fi27 < 14; fi27++ {
            _ = len(v20)
            _ = fi27
        }
    }
    var v28 int8 = 1
    if len(v20) >= 2 {
        v20 = v20[1:len(v20)]
    }
    var v29 uint = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
