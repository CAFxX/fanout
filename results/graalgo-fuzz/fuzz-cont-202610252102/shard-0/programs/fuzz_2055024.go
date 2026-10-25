package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 uint64; F1 uint }

func main() {
    var v1 uint8 = 128
    v2 := []string{"", "foo bar"}
    var v3 uint64 = 1
    v4 := make(map[string]int)
    v5 := make(map[string]int)
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 uint32 = 2147483647
        if true {
            if (((3.14159 / float64((1.5 * ((0.1 / 3.14159) - 0.1)))) * ((0.1 + 0.1) / 1e10)) < ((0.1 * 1e10) * (1e10 * float64(3.14159)))) {
                delete(v5, ("foo bar" + ("foo bar" + "hello")))
                var v8 uint64 = 111
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = v8
            } else {
                delete(v4, "foo bar")
                switch fi6 {
                case 2147483647:
                    v1 += v1
                case 42:
                    v9 := GenS1{F0: v3}
                    _ = v9
                case 316:
                    fmt.Println(v3, fi6, v1)
                }
                for fi10 := 0; fi10 < 20; fi10++ {
                    v7 = 2147483648
                    _ = fi10
                }
            }
        } else {
            var v11 string = "hello"
            v12 := GenS2{F0: 255, F1: (v3 >> ((uint(v3)) & 63))}
            _ = v11
            _ = v12
        }
        for fi13 := 0; fi13 < 4; fi13++ {
            var v14 int8 = 73
            var v15 bool = (1 >= 42)
            if true {
                var v16 int = 9223372036854775807
                _ = v16
            }
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi6
        _ = v7
    }
    fmt.Println(v3, v3)
    if true {
        v5["hello"] = 4294967295
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        var v18 bool = true
        _ = fi17
        _ = v18
    }
    fmt.Println(v3)
    v19 := GenS1{F0: (v3 << ((255) & 63))}
    switch -87 {
    case 0:
        fmt.Println(v1, v3)
    case 2147483647:
        v1 = v1
    }
    if ((!(!((96 == 14) && false))) || true) {
        for fi20 := 0; fi20 < 18; fi20++ {
            switch ("a" + ("foo bar" + (("a" + "") + "foo bar"))) {
            case "":
                for fi21 := 0; fi21 < 1; fi21++ {
                    v3 = v3
                    v22 := GenS3{F0: false, F1: (("\n\t" + "\n\t") + "foo bar")}
                    var v23 string = ((("hello" + "foo bar") + "foo bar") + "a")
                    _ = fi21
                    _ = v22
                    _ = v23
                }
            case "hello":
                v3 += (v3 >> ((4294967295) & 63))
            case "foo bar":
                fmt.Println(v3)
            }
            _ = fi20
        }
    }
    var v24 uint64 = (v3 | v3)
    var v25 uint = 0
    var v26 uint = v25
    v25 = 802
    v5[((("foo bar" + "\n\t") + ("hello" + ("\n\t" + ("hello" + "a")))) + "")] = 285
    fmt.Println(v24)
    v4["foo bar"] = 127
    for fi27 := 0; fi27 < 16; fi27++ {
        var v28 int8 = 127
        _ = fi27
        _ = v28
    }
    if false {
        for fi29 := 0; fi29 < 4; fi29++ {
            switch "\n\t" {
            case "foo bar":
                for fi30 := 0; fi30 < 17; fi30++ {
                    var v31 int32 = 1
                    v25 -= v26
                    v3 = 4294967295
                    var v32 float64 = (-2.25 / 1.5)
                    _ = fi30
                    _ = v31
                    _ = v32
                }
            case "":
                var v33 float64 = (float64(3.14159) + 0.0)
                _ = v33
            }
            var v34 int8 = 1
            _ = fi29
            _ = v34
        }
        for fi35 := 0; fi35 < 6; fi35++ {
            v36 := GenS4{F0: 815, F1: (v26 + v25)}
            _ = fi35
            _ = v36
        }
        var v37 int16 = -1
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
