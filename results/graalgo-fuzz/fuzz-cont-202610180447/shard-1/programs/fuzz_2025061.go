package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32 }
func genFn0() int16 {
    v5 := GenS1{F0: 111}
    v6 := []int{-2147483648, 128}
    _ = v5
    _ = v6
    return 1000
}

func main() {
    var v1 bool = (0.1 <= float64((-2.25 * (0.1 * 0.0))))
    v2 := []int{42}
    var v3 uint8 = 128
    var v4 string = ((("foo bar" + (((("\n\t" + (((("hello" + ("" + "hello")) + (("hello" + ("hello" + (("" + "a") + "\n\t"))) + "")) + ("hello" + ("hello" + "hello"))) + "hello")) + (("\n\t" + "a") + "a")) + "hello") + "")) + "foo bar") + "hello")
    v3 -= 127
    if true {
        var v7 int8 = 127
        var v8 bool = (!v1)
        _ = v7
        _ = v8
    } else {
        for fi9 := 0; fi9 < 20; fi9++ {
            switch 9223372036854775807 {
            case 14:
                fmt.Println(v1, v1, v3)
            }
            var v10 uint16 = 42
            _ = fi9
            _ = v10
        }
        if v1 {
            v4 = ("foo bar" + (v4 + v4))
            for fi11 := 0; fi11 < 8; fi11++ {
                v3 = 255
                var v12 string = (v4 + "")
                var v13 uint8 = (v3 & (v3 / ((219) | 1)))
                _ = fi11
                _ = v12
                _ = v13
            }
        }
        v14 := GenS2{F0: -866}
        _ = v14
    }
    v2 = append(v2, 935)
    v15 := GenS1{F0: 2147483647}
    if v1 {
        var v16 uint16 = 42
        var v17 uint8 = v3
        _ = v16
        _ = v17
    } else {
        if (v1 || true) {
            var v18 bool = false
            var v19 uint16 = 42
            _ = v18
            _ = v19
        } else {
            switch (v4 + v4) {
            case "foo bar":
                v1 = (v4 > "\n\t")
            case "a":
                switch 42 {
                case 128:
                    var v20 uint16 = 0
                    _ = v20
                case 32768:
                    v4 = v4
                }
            case "\n\t":
                for fi21 := 0; fi21 < 3; fi21++ {
                    _ = len(v2)
                    v4 = "a"
                    for fi22 := 0; fi22 < 3; fi22++ {
                        if len(v2) > 0 {
                            v2[0] = (fi22 % ((fi22) | 1))
                        }
                        var v23 uint64 = 42
                        _ = fi22
                        _ = v23
                    }
                    _ = fi21
                }
            }
            v24 := GenS1{F0: 785}
            _ = v24
        }
    }
    v1 = v1
    if (v3 < (v3 ^ (v3 - 2))) {
        for fi25 := 0; fi25 < 12; fi25++ {
            if len(v2) > 0 {
                v2[0] = 2147483647
            }
            _ = fi25
        }
    } else {
        v3 = (v3 ^ (v3 % ((v3) | 1)))
    }
    v1 = v1
    for fi26 := 0; fi26 < 1; fi26++ {
        switch v4 {
        case "foo bar":
            switch "foo bar" {
            case "a":
                fmt.Println(fi26, v3, v3)
            case "":
                switch fi26 {
                case -1:
                    fmt.Println(v3, fi26, v4)
                default:
                    fmt.Println(v1, v3)
                }
            case "\n\t":
                if v1 {
                    if true {
                        v27 := GenS2{F0: 42}
                        for fi28 := 0; fi28 < 11; fi28++ {
                            fmt.Println(v3)
                            fmt.Println(fi28, v3, v4)
                            var v29 int16 = genFn0()
                            _ = fi28
                            _ = v29
                        }
                        v1 = v1
                        _ = v27
                    }
                    v2 = append(v2, fi26)
                    var v30 uint8 = (v3 >> ((42) & 63))
                    _ = v30
                } else {
                    fmt.Println(fi26, v4, v1)
                }
            }
        default:
            v1 = false
        }
        var v31 uint64 = 4294967295
        _ = fi26
        _ = v31
    }
    v2 = append(v2, 0)
    v1 = true
    if (-2147483648 > 841) {
        var v32 int64 = 0
        v2 = append(v2, 32768)
        _ = v32
    } else {
        if true {
            fmt.Println(v4, v4)
        }
        fmt.Println(v1, v3)
    }
    v4 = (v4 + (v4 + v4))
    for fi33 := 0; fi33 < 17; fi33++ {
        var v34 uint16 = 255
        var v35 uint8 = 2
        _ = fi33
        _ = v34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
