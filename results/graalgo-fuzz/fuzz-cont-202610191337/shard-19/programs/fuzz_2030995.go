package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 bool; F1 string; F2 uint32 }

func main() {
    var v1 uint8 = 2
    v2 := []string{("" + ("a" + "hello"))}
    var v3 uint = 0
    var v4 bool = true
    v5 := make(map[uint32]int)
    if (!false) {
        v3 += v3
        if v4 {
            if v4 {
                switch ("\n\t" + "a") {
                case "\n\t":
                    var v6 int64 = -1
                    _ = v6
                }
                v1 = v1
                var v7 float64 = (1.5 * float64((-2.25 + (1e10 / 1e10))))
                _ = v7
            } else {
                if ("\n\t" != "hello") {
                    if (18446744073709551615 <= 766) {
                        var v8 int = 4294967295
                        for fi9 := 0; fi9 < 8; fi9++ {
                            for fi10 := 0; fi10 < 5; fi10++ {
                                var v11 string = ("hello" + "hello")
                                fmt.Println(fi9, v11, v8)
                                if len(v2) > 0 {
                                    v2[0] = v11
                                }
                                _ = fi10
                                _ = v11
                            }
                            var v12 float64 = 1.5
                            fmt.Println(v12, fi9)
                            _ = fi9
                            _ = v12
                        }
                        for fi13 := 0; fi13 < 1; fi13++ {
                            v4 = (v3 <= (^(v3)))
                            switch ((("" + "a") + (("" + (("foo bar" + (("\n\t" + "\n\t") + "hello")) + "\n\t")) + (("foo bar" + "foo bar") + "foo bar"))) + ("hello" + "hello")) {
                            case "hello":
                                v4 = true
                            case "a":
                                var v14 bool = ((-2147483648 < 2147483647) && (1e10 <= 3.14159))
                                _ = v14
                            case "foo bar":
                                v15 := GenS1{F0: uint16(v3)}
                                _ = v15
                            }
                            _ = fi13
                        }
                        _ = v8
                    } else {
                        v16 := GenS2{F0: ("hello" + "")}
                        if len(v2) > 0 {
                            v2[0] = "a"
                        }
                        fmt.Println(v1)
                        _ = v16
                    }
                    v17 := GenS3{F0: "", F1: 0}
                    v18 := GenS4{F0: ((v1 + v1) == (^(v1))), F1: "", F2: 1}
                    _ = v17
                    _ = v18
                }
            }
            var v19 int = 127
            _ = v19
        }
        if false {
            switch (("\n\t" + ("hello" + ((("" + "hello") + ("hello" + "hello")) + ("" + "a")))) + ("foo bar" + "")) {
            case "":
                v3 += v3
            case "foo bar":
                var v20 uint8 = v1
                _ = v20
            }
            var v21 int32 = 514
            var v22 int32 = -983
            _ = v21
            _ = v22
        } else {
            var v23 uint = v3
            _ = v23
        }
    }
    switch "hello" {
    case "a":
        fmt.Println(v4, v1)
    }
    if true {
        var v24 int64 = 648
        v4 = (3.14159 == (float64(0.1) - ((1.5 / 0.1) * (-2.25 + -2.25))))
        _ = v24
    } else {
        switch ((("hello" + ("" + "foo bar")) + "\n\t") + "hello") {
        case "":
            v25 := GenS1{F0: 691}
            _ = v25
        default:
            var v26 int16 = 0
            _ = v26
        }
    }
    if v4 {
        fmt.Println(v1)
        v5[0] = -962
    }
    v1 = 86
    v27 := GenS2{F0: (("hello" + (("" + ("a" + (("" + ("" + ("" + ("\n\t" + "\n\t")))) + ("hello" + "")))) + "hello")) + (("a" + "foo bar") + "hello"))}
    switch 128 {
    case -395:
        v5[1] = -1
    case -352:
        _ = len(v5)
    case 9223372036854775807:
        fmt.Println(v4, v3, v4)
    default:
        v28 := GenS2{F0: ""}
        _ = v28
    }
    fmt.Println(v3)
    switch "\n\t" {
    case "a":
        delete(v5, 429)
    case "":
        v29 := GenS2{F0: (("" + "foo bar") + "\n\t")}
        _ = v29
    }
    v1 -= v1
    for fi30 := 0; fi30 < 19; fi30++ {
        fmt.Println(fi30, v1)
        v1 = 42
        switch "hello" {
        case "foo bar":
            fmt.Println(v4, v3)
        }
        _ = fi30
    }
    v4 = (("hello" + "\n\t") == "a")
    fmt.Println(v4, v1)
    fmt.Println(v4, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
