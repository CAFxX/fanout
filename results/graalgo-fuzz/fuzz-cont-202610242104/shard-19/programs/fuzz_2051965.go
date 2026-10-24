package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8 }
func genFn0() int32 {
    var v4 uint8 = 1
    v5 := GenS1{F0: v4}
    v6 := []int{32768}
    if false {
        for fi7 := 0; fi7 < 17; fi7++ {
            switch (("" + "a") + "foo bar") {
            case "a":
                var v8 string = "\n\t"
                _ = v8
            }
            _ = fi7
        }
        var v9 float32 = 0.0
        for fi10 := 0; fi10 < 4; fi10++ {
            for fi11 := 0; fi11 < 14; fi11++ {
                fmt.Println(v4, v4)
                v4 = (v4 >> ((2) & 63))
                _ = fi11
            }
            fmt.Println(v4, v9, v9)
            for fi12 := 0; fi12 < 11; fi12++ {
                v6 = append(v6, fi10)
                _ = fi12
            }
            _ = fi10
        }
        _ = v9
    } else {
        v4 -= (v4 - (v4 * v4))
        if (!(!true)) {
            switch ("foo bar" + "foo bar") {
            case "\n\t":
                v13 := GenS2{F0: 116}
                _ = v13
            case "foo bar":
                if len(v6) > 0 {
                    v6[0] = 9223372036854775807
                }
            case "":
                switch 127 {
                case 536:
                    var v14 uint = 42
                    _ = v14
                default:
                    if len(v6) > 0 {
                        v6[0] = -1
                    }
                }
            }
        }
        switch "" {
        case "":
            fmt.Println(v4)
        case "foo bar":
            fmt.Println(v4)
        }
    }
    _ = v4
    _ = v5
    _ = v6
    return 2147483647
}
func genFn1() int16 {
    var v15 bool = false
    v16 := GenS1{F0: 2}
    fmt.Println(v15)
    v17 := []int{-220, 510}
    for fi18 := 0; fi18 < 7; fi18++ {
        if len(v17) > 0 {
            v17[0] = 32768
        }
        v19 := GenS2{F0: 0}
        v15 = v15
        _ = fi18
        _ = v19
    }
    _ = v15
    _ = v16
    _ = v17
    return -1
}

func main() {
    v1 := make(map[string]string)
    var v2 int = 4294967295
    v3 := []string{"", "foo bar"}
    fmt.Println(v2)
    for fi20 := 0; fi20 < 13; fi20++ {
        if (true && false) {
            v2 += fi20
            v2 += fi20
            fmt.Println(v2)
        } else {
            v2 += v2
        }
        fmt.Println(v2, v2)
        _ = fi20
    }
    _ = len(v1)
    v2 += v2
    v1["a"] = "a"
    for fi21 := 0; fi21 < 12; fi21++ {
        v2 = 128
        _ = fi21
    }
    delete(v1, "")
    if false {
        switch "\n\t" {
        case "":
            for fi22 := 0; fi22 < 5; fi22++ {
                v23 := GenS2{F0: 99}
                for fi24 := 0; fi24 < 18; fi24++ {
                    _ = len(v1)
                    v2 -= v2
                    _ = fi24
                }
                _ = fi22
                _ = v23
            }
        case "a":
            for fi25 := 0; fi25 < 18; fi25++ {
                switch fi25 {
                case 1:
                    v26 := GenS2{F0: -1}
                    _ = v26
                default:
                    v2 = (v2 / ((fi25) | 1))
                }
                _ = fi25
            }
        case "hello":
            var v27 int16 = 0
            _ = v27
        default:
            for fi28 := 0; fi28 < 6; fi28++ {
                var v29 int = -1
                _ = fi28
                _ = v29
            }
        }
        v1["foo bar"] = "foo bar"
    }
    var v30 int16 = genFn1()
    v2 -= 41
    v1[""] = ""
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
}
