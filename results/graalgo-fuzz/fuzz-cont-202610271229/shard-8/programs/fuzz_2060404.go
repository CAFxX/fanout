package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 string }
type GenS2 struct { F0 int }

func main() {
    v1 := []int{}
    v2 := make(map[int]bool)
    v3 := []uint32{42}
    if false {
        var v4 int8 = 90
        v5 := GenS1{F0: "a", F1: 270, F2: ("\n\t" + "a")}
        _ = v4
        _ = v5
    } else {
        if len(v3) > 0 {
            v3[0] = 2147483648
        }
        if (65535 >= 42) {
            var v6 float64 = (-2.25 * (0.1 + 1.5))
            switch "a" {
            case "":
                for fi7 := 0; fi7 < 19; fi7++ {
                    var v8 float32 = -2.25
                    _ = fi7
                    _ = v8
                }
            }
            _ = v6
        }
    }
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 int = (-(fi9))
        _ = fi9
        _ = v10
    }
    var v11 bool = (true || false)
    if true {
        switch -454 {
        case 2147483647:
            v11 = (132 == 255)
        case -2147483648:
            delete(v2, 2147483647)
        default:
            fmt.Println(v11)
        }
    } else {
        for fi12 := 0; fi12 < 5; fi12++ {
            v11 = (v11 || (-41 >= -1))
            switch "foo bar" {
            case "hello":
                v11 = ((3.14159 / 1.0) <= 0.1)
            case "":
                var v13 uint8 = 1
                _ = v13
            case "foo bar":
                fmt.Println(v11, fi12)
            }
            _ = fi12
        }
    }
    v2[42] = (!v11)
    v14 := GenS1{F0: ("hello" + (("" + "foo bar") + ("hello" + "\n\t"))), F1: 2147483647, F2: "\n\t"}
    v3 = append(v3, 814)
    v11 = (!(((0 < -32768) || true) || v11))
    v15 := GenS1{F0: "hello", F1: -261, F2: ""}
    if (9223372036854775807 >= -439) {
        v11 = ((false && false) && (2 >= 255))
        fmt.Println(v11)
    }
    v16 := GenS1{F0: "hello", F1: 498, F2: ("\n\t" + ("\n\t" + "\n\t"))}
    v11 = v11
    for fi17 := 0; fi17 < 19; fi17++ {
        var v18 int32 = 409
        _ = fi17
        _ = v18
    }
    switch "foo bar" {
    case "foo bar":
        if (-1 <= 0) {
            v19 := GenS1{F0: "", F1: 1, F2: (("foo bar" + "foo bar") + ("\n\t" + "\n\t"))}
            v20 := GenS2{F0: 42}
            for fi21 := 0; fi21 < 10; fi21++ {
                fmt.Println(v11, fi21)
                _ = fi21
            }
            _ = v19
            _ = v20
        }
    case "":
        if ((!true) || ("" == ("" + ("\n\t" + "")))) {
            if true {
                fmt.Println(v11)
                v11 = (432 < 18446744073709551615)
                v22 := v2[4294967295]
                _ = v22
            }
        }
    default:
        for fi23 := 0; fi23 < 17; fi23++ {
            v11 = (!true)
            _ = fi23
        }
    }
    var v24 int = -211
    v24 += (v24 | (v24 ^ (v24 % (((^(v24))) | 1))))
    for fi25 := 0; fi25 < 11; fi25++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi25
    }
    v24 += (v24 ^ (v24 / (((v24 - v24)) | 1)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v24
    _ = v3
}
