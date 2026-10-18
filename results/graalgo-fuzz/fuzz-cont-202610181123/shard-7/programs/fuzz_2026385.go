package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 uint64; F2 int }

func main() {
    var v1 int64 = -297
    var v2 uint64 = 751
    var v3 int = -1
    v4 := []string{("a" + ((("\n\t" + "") + (("hello" + "\n\t") + "")) + (("foo bar" + (((("a" + ("a" + "\n\t")) + ("hello" + (((((("foo bar" + ("hello" + "\n\t")) + "") + "hello") + ("a" + "a")) + ("hello" + "a")) + ""))) + "foo bar") + "a")) + (("a" + "\n\t") + ("hello" + "\n\t"))))), (("" + "") + ""), ("\n\t" + ""), ((("\n\t" + (("\n\t" + "\n\t") + "hello")) + "\n\t") + "a")}
    for fi5 := 0; fi5 < 8; fi5++ {
        v6 := GenS1{F0: (v1 | (v1 - v1))}
        var v7 int8 = 127
        if true {
            var v8 uint64 = 42
            _ = v8
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    fmt.Println(v2, v2)
    _ = len(v4)
    if (true || false) {
        if true {
            fmt.Println(v2, v3)
            fmt.Println(v3)
        } else {
            for fi9 := 0; fi9 < 8; fi9++ {
                var v10 int16 = -1
                if (4294967295 != 496) {
                    v11 := GenS1{F0: v1}
                    fmt.Println(v3)
                    _ = v11
                }
                var v12 int64 = v1
                _ = fi9
                _ = v10
                _ = v12
            }
            var v13 float32 = 1.5
            switch v3 {
            case -239:
                v2 = v2
            case 9223372036854775807:
                fmt.Println(v1, v13, v3)
            case 4294967295:
                if true {
                    if (65535 <= 0) {
                        v2 -= v2
                        switch ((("a" + "") + (("hello" + "foo bar") + ((("\n\t" + "foo bar") + "") + (("" + "\n\t") + "")))) + "") {
                        case "a":
                            v3 += v3
                        default:
                            v3 -= (v3 % ((v3) | 1))
                        }
                    }
                }
            }
            _ = v13
        }
        v2 = 4294967295
        for fi14 := 0; fi14 < 6; fi14++ {
            fmt.Println(v2)
            _ = fi14
        }
    } else {
        fmt.Println(v1)
        var v15 float32 = (float32((0.1 * 1e10)) - -2.25)
        var v16 bool = false
        _ = v15
        _ = v16
    }
    var v17 uint64 = v2
    v2 = 4294967295
    if true {
        switch v3 {
        case 32767:
            _ = len(v4)
        default:
            v18 := GenS2{F0: "hello", F1: (^(v2)), F2: (^(v3))}
            _ = v18
        }
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    for fi19 := 0; fi19 < 2; fi19++ {
        var v20 int64 = (v1 * -1)
        var v21 string = "a"
        _ = fi19
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 20; fi22++ {
        if (true || false) {
            if false {
                var v23 bool = false
                if (0.1 > 3.14159) {
                    v24 := GenS2{F0: "\n\t", F1: v2, F2: 32767}
                    v17 = (^(v2))
                    var v25 uint32 = 0
                    _ = v24
                    _ = v25
                }
                _ = v23
            } else {
                var v26 uint16 = 496
                _ = v26
            }
        } else {
            v27 := GenS2{F0: "a", F1: (^(v2)), F2: -9223372036854775808}
            _ = v27
        }
        var v28 string = ""
        v2 = 0
        _ = fi22
        _ = v28
    }
    v17 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
