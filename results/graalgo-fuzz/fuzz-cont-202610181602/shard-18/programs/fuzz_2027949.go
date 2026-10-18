package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 uint8; F2 uint }
func genFn0(p0 int8) uint64 {
    p0 += -128
    if false {
        var v6 uint8 = 104
        v7 := GenS1{F0: 495}
        switch 32767 {
        case 42:
            switch 4294967295 {
            case 321:
                p0 = p0
            default:
                if ((127 > (v6 % ((v6) | 1))) && false) {
                    v6 = 181
                    v6 -= (v6 >> ((369) & 63))
                    p0 = (p0 % ((p0) | 1))
                }
            }
        }
        _ = v6
        _ = v7
    } else {
        p0 = (p0 % ((1) | 1))
    }
    _ = p0
    return 0
}

func main() {
    var v1 bool = true
    v2 := []string{(("foo bar" + "") + ("hello" + "\n\t"))}
    v3 := make(map[string]int)
    var v4 uint8 = 0
    var v5 int = 0
    if (42 != -128) {
        v4 -= 255
        fmt.Println(v4, v5)
    } else {
        switch (v5 & (v5 ^ v5)) {
        case 128:
            for fi8 := 0; fi8 < 16; fi8++ {
                for fi9 := 0; fi9 < 8; fi9++ {
                    if (128 > v5) {
                        fmt.Println(v5)
                    } else {
                        var v10 int = (v5 % (((fi8 & v5)) | 1))
                        v5 = (^(v10))
                        v2 = append(v2, ("hello" + "hello"))
                        _ = v10
                    }
                    if v1 {
                        v5 = (fi8 * (v5 - (fi8 << ((2) & 63))))
                    } else {
                        v5 = (v5 ^ -13)
                    }
                    _ = fi9
                }
                fmt.Println(v4, v1, fi8)
                for fi11 := 0; fi11 < 19; fi11++ {
                    v4 = v4
                    switch "\n\t" {
                    case "a":
                        v1 = v1
                    case "foo bar":
                        var v12 int32 = 42
                        _ = v12
                    }
                    _ = fi11
                }
                _ = fi8
            }
        case 4294967295:
            var v13 uint16 = 65535
            _ = v13
        default:
            v14 := GenS2{F0: ("hello" + (((("a" + "hello") + (("a" + "foo bar") + "a")) + ("hello" + "a")) + "\n\t")), F1: (v4 / ((v4) | 1)), F2: 1}
            _ = v14
        }
        for fi15 := 0; fi15 < 17; fi15++ {
            for fi16 := 0; fi16 < 13; fi16++ {
                switch "\n\t" {
                case "a":
                    var v17 uint32 = 2147483647
                    _ = v17
                }
                _ = len(v3)
                _ = fi16
            }
            v18 := GenS1{F0: 0}
            _ = fi15
            _ = v18
        }
        v4 += (v4 * v4)
    }
    fmt.Println(v5)
    v19 := GenS1{F0: 1}
    v20 := GenS1{F0: 550}
    v4 = v4
    var v21 uint64 = 1
    if (-481 >= 1000) {
        v21 -= (v21 + v21)
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        if (0 < 0) {
            if (32767 >= 255) {
                v4 += v4
            } else {
                v2 = append(v2, "hello")
                v5 -= (v5 & (^(fi22)))
                fmt.Println(fi22, v21, v1)
            }
            var v23 bool = (18446744073709551615 >= 4294967295)
            var v24 string = (("\n\t" + ("a" + "\n\t")) + "")
            _ = v23
            _ = v24
        }
        switch "a" {
        case "\n\t":
            for fi25 := 0; fi25 < 9; fi25++ {
                var v26 string = ((("foo bar" + "\n\t") + "a") + "a")
                if len(v2) > 0 {
                    v2[0] = v26
                }
                var v27 uint = 18446744073709551615
                _ = fi25
                _ = v26
                _ = v27
            }
        case "":
            for fi28 := 0; fi28 < 18; fi28++ {
                fmt.Println(v5, v4, v5)
                fmt.Println(v1)
                _ = len(v2)
                _ = fi28
            }
        case "foo bar":
            v5 -= (fi22 >> ((573) & 63))
        default:
            v5 -= v5
        }
        _ = fi22
    }
    var v29 uint = 255
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
