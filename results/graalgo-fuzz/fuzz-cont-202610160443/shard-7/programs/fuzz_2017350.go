package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int }

func main() {
    v1 := []string{}
    var v2 float64 = 3.14159
    var v3 int16 = 42
    var v4 uint = 18446744073709551615
    v2 = v2
    v3 = (v3 * (v3 % ((42) | 1)))
    for fi5 := 0; fi5 < 3; fi5++ {
        for fi6 := 0; fi6 < 18; fi6++ {
            v4 -= v4
            if true {
                fmt.Println(fi6)
            }
            _ = fi6
        }
        if false {
            v7 := GenS1{F0: v4, F1: ("" + "")}
            switch (fi5 | (fi5 ^ fi5)) {
            case 32767:
                if (!(2 > 128)) {
                    v2 = v2
                    v3 -= v3
                    v8 := GenS2{F0: "foo bar"}
                    _ = v8
                }
            case 851:
                v9 := GenS2{F0: ("hello" + (((("hello" + ("foo bar" + "a")) + ((("\n\t" + "foo bar") + "\n\t") + "")) + ("a" + (("a" + "a") + "a"))) + ("a" + "foo bar")))}
                _ = v9
            default:
                var v10 uint16 = 42
                _ = v10
            }
            fmt.Println(v4)
            _ = v7
        }
        _ = fi5
    }
    for fi11 := 0; fi11 < 6; fi11++ {
        if len(v1) > 0 {
            v1[0] = ((("hello" + ("hello" + ("hello" + (("hello" + "hello") + "")))) + "") + (("" + ("a" + "\n\t")) + "hello"))
        }
        switch fi11 {
        case -1:
            v1 = append(v1, (("hello" + "") + "\n\t"))
        }
        var v12 uint = v4
        _ = fi11
        _ = v12
    }
    v13 := GenS3{F0: 4294967295}
    switch "a" {
    case "foo bar":
        v4 = 255
    case "a":
        v1 = append(v1, "foo bar")
    case "hello":
        v14 := GenS3{F0: -346}
        _ = v14
    default:
        var v15 int = 42
        _ = v15
    }
    v4 = v4
    switch "\n\t" {
    case "a":
        if (179 <= 2147483647) {
            if len(v1) > 0 {
                v1[0] = ""
            }
            var v16 uint16 = 255
            fmt.Println(v2)
            _ = v16
        }
    case "\n\t":
        if ("foo bar" > "foo bar") {
            if (0 >= 120) {
                fmt.Println(v4, v2)
                v1 = append(v1, (((("\n\t" + (("\n\t" + (("a" + "hello") + "foo bar")) + "")) + ("" + "hello")) + "foo bar") + ""))
                var v17 int = int(v4)
                _ = v17
            } else {
                var v18 string = "a"
                v3 -= (v3 / ((-1) | 1))
                _ = v18
            }
            if true {
                v4 -= v4
                if (true && false) {
                    v3 -= 334
                    fmt.Println(v2)
                } else {
                    for fi19 := 0; fi19 < 8; fi19++ {
                        var v20 int = fi19
                        var v21 bool = ((32768 <= (v20 | v20)) && true)
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                }
                var v22 int64 = -298
                _ = v22
            }
            var v23 int32 = 42
            _ = v23
        }
    }
    var v24 int16 = v3
    v2 = (float64(0.0) * ((v2 * v2) * v2))
    for fi25 := 0; fi25 < 16; fi25++ {
        v2 = float64((3.14159 / float32(v2)))
        v26 := GenS3{F0: (fi25 * 128)}
        _ = fi25
        _ = v26
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v24
    _ = v3
    _ = v4
}
