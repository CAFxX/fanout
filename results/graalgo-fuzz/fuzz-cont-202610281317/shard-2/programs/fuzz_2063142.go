package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 int }
type GenS2 struct { F0 uint32; F1 bool }

func main() {
    v1 := make(map[int]uint32)
    var v2 float64 = 0.1
    var v3 float32 = 3.14159
    v4 := []string{(("\n\t" + ("hello" + ((("" + "foo bar") + "foo bar") + "\n\t"))) + (("\n\t" + "a") + ((("" + "\n\t") + "foo bar") + "hello"))), (("\n\t" + "\n\t") + ("a" + ("a" + "")))}
    var v5 uint32 = 155
    var v6 string = ("hello" + ("hello" + ("" + ("\n\t" + ("" + "")))))
    if true {
        for fi7 := 0; fi7 < 7; fi7++ {
            var v8 string = ""
            switch v6 {
            case "a":
                var v9 uint16 = 827
                _ = v9
            case "\n\t":
                var v10 int64 = 42
                _ = v10
            case "":
                for fi11 := 0; fi11 < 19; fi11++ {
                    fmt.Println(v6, v2, v8)
                    var v12 int64 = 1
                    _ = fi11
                    _ = v12
                }
            default:
                var v13 uint32 = (v5 % ((v5) | 1))
                _ = v13
            }
            _ = fi7
            _ = v8
        }
        v6 = (v6 + (("\n\t" + (v6 + "foo bar")) + v6))
    } else {
        fmt.Println(v2)
        if true {
            for fi14 := 0; fi14 < 17; fi14++ {
                _ = len(v4)
                fmt.Println(v2)
                _ = fi14
            }
        } else {
            fmt.Println(v6)
        }
        v15 := GenS1{F0: v6, F1: 2147483647, F2: 0}
        _ = v15
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        _ = len(v1)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v5 = v5
        _ = fi16
    }
    v2 += float64(float32(float64(v3)))
    if false {
        switch 4294967295 {
        case 4294967295:
            var v17 uint64 = 477
            _ = v17
        case 323:
            v6 = ""
        }
        var v18 bool = (!false)
        _ = v18
    }
    v5 -= (v5 >> ((808) & 63))
    fmt.Println(v6)
    for fi19 := 0; fi19 < 3; fi19++ {
        v2 = v2
        var v20 bool = true
        _ = fi19
        _ = v20
    }
    v6 = "a"
    delete(v1, 128)
    for fi21 := 0; fi21 < 7; fi21++ {
        if len(v4) > 0 {
            v4[0] = v6
        }
        for fi22 := 0; fi22 < 3; fi22++ {
            var v23 string = (((v6 + v6) + (v6 + v6)) + v6)
            _ = fi22
            _ = v23
        }
        if true {
            fmt.Println(v6, v3, v2)
            _ = len(v4)
        } else {
            _ = len(v1)
        }
        _ = fi21
    }
    v5 -= (v5 | v5)
    if (!true) {
        var v24 uint64 = 0
        _ = v24
    }
    v4 = append(v4, "a")
    v3 = v3
    if true {
        v6 = "foo bar"
        v25 := GenS1{F0: ((((v6 + v6) + v6) + v6) + v6), F1: -9223372036854775808, F2: 9223372036854775807}
        switch (v6 + v6) {
        case "":
            if (((v2 - v2) * 3.14159) < v2) {
                if (2 == 32767) {
                    v4 = append(v4, ((v6 + v6) + "a"))
                    switch v6 {
                    case "foo bar":
                        v26 := GenS1{F0: v6, F1: -1, F2: 327}
                        _ = v26
                    default:
                        v1[-1] = (v5 ^ 980)
                    }
                }
                v1[9223372036854775807] = v5
            }
        case "a":
            v27 := GenS2{F0: v5, F1: true}
            _ = v27
        }
        _ = v25
    }
    v3 = float32(3.14159)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
