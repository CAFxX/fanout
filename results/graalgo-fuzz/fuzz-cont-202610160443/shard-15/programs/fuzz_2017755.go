package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16 }
func genFn0(p0 string) uint16 {
    if true {
        p0 = p0
        var v7 int64 = 647
        _ = v7
    } else {
        switch ("foo bar" + ((p0 + "") + p0)) {
        case "a":
            fmt.Println(p0)
        }
        v8 := []int{}
        _ = v8
    }
    var v9 int32 = 374
    _ = p0
    _ = v9
    return 1
}
func genFn1() int32 {
    for fi10 := 0; fi10 < 1; fi10++ {
        var v11 float32 = 3.14159
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 8; fi12++ {
        v13 := GenS1{F0: 0}
        for fi14 := 0; fi14 < 19; fi14++ {
            var v15 float64 = 1.5
            switch "a" {
            case "a":
                var v16 uint64 = 18446744073709551615
                _ = v16
            }
            var v17 uint = 18446744073709551615
            _ = fi14
            _ = v15
            _ = v17
        }
        _ = fi12
        _ = v13
    }
    v18 := GenS1{F0: 42}
    v19 := []int{}
    _ = v18
    _ = v19
    return 1
}

func main() {
    var v1 int16 = -1
    v2 := make(map[string]int)
    var v3 uint16 = 255
    v4 := []string{"\n\t"}
    var v5 uint64 = 42
    v6 := make(map[int]int)
    switch -1 {
    case 2147483647:
        if (true && (1 <= uint8(v3))) {
            v3 -= 1
        } else {
            for fi20 := 0; fi20 < 9; fi20++ {
                switch "hello" {
                case "a":
                    v1 = (v1 - v1)
                case "\n\t":
                    v3 += (^(v3))
                default:
                    var v21 int16 = -272
                    _ = v21
                }
                _ = fi20
            }
            v3 = v3
            if false {
                v1 += v1
            } else {
                if false {
                    for fi22 := 0; fi22 < 13; fi22++ {
                        v5 = 891
                        v1 += v1
                        _ = fi22
                    }
                    v3 = v3
                } else {
                    v2["hello"] = 168
                    v3 = (^(v3))
                }
                var v23 uint = 42
                v1 += v1
                _ = v23
            }
        }
    case -9223372036854775808:
        v24 := GenS1{F0: 1}
        _ = v24
    case 42:
        var v25 int64 = 1
        _ = v25
    }
    _ = len(v4)
    var v26 float64 = 0.0
    switch "hello" {
    case "foo bar":
        for fi27 := 0; fi27 < 16; fi27++ {
            v26 = (1e10 + (v26 + v26))
            var v28 uint16 = v3
            _ = fi27
            _ = v28
        }
    case "\n\t":
        fmt.Println(v3, v3)
    case "":
        v29 := v2[("a" + "a")]
        _ = v29
    }
    if len(v4) > 0 {
        v4[0] = "\n\t"
    }
    var v30 uint32 = 273
    _ = len(v2)
    if (2 < (v5 + 571)) {
        v6[0] = -2147483648
    } else {
        v31 := GenS1{F0: 649}
        v1 += (v1 * (v1 * (v1 % ((v1) | 1))))
        var v32 int16 = v1
        _ = v31
        _ = v32
    }
    if ((v3 * 1) > v3) {
        for fi33 := 0; fi33 < 6; fi33++ {
            var v34 float64 = v26
            if false {
                switch (fi33 | fi33) {
                case -71:
                    fmt.Println(v1)
                case 9223372036854775807:
                    var v35 string = ("" + "foo bar")
                    _ = v35
                case 636:
                    v36 := GenS2{F0: 0}
                    _ = v36
                }
                v37 := GenS2{F0: v1}
                _ = v37
            }
            _ = fi33
            _ = v34
        }
        fmt.Println(v1, v5)
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    _ = len(v4)
    var v38 uint16 = genFn0(((("" + (("" + "\n\t") + ("a" + "foo bar"))) + ("hello" + "foo bar")) + ("hello" + ((("hello" + "hello") + "a") + "hello"))))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v30
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
