package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]bool)
    v3 := []int{128, 2147483647, -9223372036854775808, -7}
    switch "" {
    case "hello":
        for fi4 := 0; fi4 < 7; fi4++ {
            fmt.Println(fi4)
            _ = fi4
        }
    case "":
        var v5 uint32 = 451
        _ = v5
    case "foo bar":
        if (1 == 1000) {
            if false {
                var v6 uint = 65535
                v6 += v6
                _ = v6
            } else {
                v1[4294967295] = ((true || false) || (-2147483648 > 189))
                switch "\n\t" {
                case "a":
                    _ = len(v2)
                default:
                    var v7 int8 = 122
                    _ = v7
                }
            }
            var v8 float64 = (((0.0 / 0.1) * 1.5) / 3.14159)
            switch (("\n\t" + "foo bar") + "\n\t") {
            case "foo bar":
                v9 := GenS1{F0: 18446744073709551615}
                _ = v9
            case "\n\t":
                var v10 uint32 = 741
                _ = v10
            default:
                switch 761 {
                case -1:
                    v1[0] = false
                case -2147483648:
                    v3 = append(v3, 234)
                case 128:
                    if false {
                        v8 = v8
                    }
                }
            }
            _ = v8
        }
    }
    var v11 uint32 = 65535
    for fi12 := 0; fi12 < 5; fi12++ {
        v11 -= v11
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi12
    }
    v13 := GenS1{F0: 1}
    v14 := GenS2{F0: 18446744073709551615}
    v11 = (v11 - (v11 << ((932) & 63)))
    var v15 string = ("hello" + "\n\t")
    for fi16 := 0; fi16 < 15; fi16++ {
        v15 = ""
        if false {
            var v17 int16 = 225
            var v18 uint32 = v11
            for fi19 := 0; fi19 < 4; fi19++ {
                v20 := GenS2{F0: 42}
                _ = fi19
                _ = v20
            }
            _ = v17
            _ = v18
        } else {
            _ = len(v3)
            var v21 float64 = -2.25
            _ = v21
        }
        _ = fi16
    }
    if false {
        for fi22 := 0; fi22 < 4; fi22++ {
            for fi23 := 0; fi23 < 17; fi23++ {
                _ = len(v3)
                _ = len(v1)
                _ = fi23
            }
            _ = fi22
        }
        if (0 < 0) {
            for fi24 := 0; fi24 < 18; fi24++ {
                v15 = v15
                _ = fi24
            }
            for fi25 := 0; fi25 < 8; fi25++ {
                v2[v11] = false
                v3 = append(v3, fi25)
                v15 = "hello"
                _ = fi25
            }
            _ = len(v1)
        }
    }
    var v26 float64 = 3.14159
    var v27 bool = ((((("foo bar" + "foo bar") + ("a" + "foo bar")) + v15) + (("\n\t" + "a") + (v15 + ("foo bar" + v15)))) == v15)
    fmt.Println(v15, v26, v11)
    if v27 {
        v27 = v27
    }
    v27 = v27
    if (!true) {
        if v27 {
            for fi28 := 0; fi28 < 1; fi28++ {
                var v29 float32 = 0.1
                _ = fi28
                _ = v29
            }
            v15 = v15
        }
        v30 := GenS2{F0: 65535}
        _ = v30
    }
    for fi31 := 0; fi31 < 8; fi31++ {
        var v32 string = "\n\t"
        _ = fi31
        _ = v32
    }
    var v33 int64 = 646
    var v34 int32 = -2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v33
    _ = v34
}
