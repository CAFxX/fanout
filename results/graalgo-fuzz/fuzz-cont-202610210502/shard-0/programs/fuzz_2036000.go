package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16 }
func genFn0(p0 string, p1 int32) int {
    fmt.Println(p1)
    fmt.Println(p0, p1)
    v7 := []int{130, 91, 9223372036854775807}
    v8 := GenS1{F0: ""}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -1
}

func main() {
    var v1 bool = true
    v2 := []string{("" + "\n\t"), "\n\t", "hello"}
    var v3 int64 = 675
    v4 := make(map[string]uint32)
    var v5 int = 32768
    v6 := make(map[string]string)
    var v9 float32 = -2.25
    if (373 > 2) {
        var v10 int16 = 42
        switch "a" {
        case "hello":
            var v11 uint = 4294967295
            _ = v11
        default:
            fmt.Println(v5)
        }
        _ = v10
    } else {
        for fi12 := 0; fi12 < 18; fi12++ {
            v13 := GenS1{F0: ("\n\t" + "a")}
            _ = fi12
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        var v15 uint16 = 1
        if false {
            var v16 uint = 583
            v5 += genFn0("\n\t", 1)
            _ = v16
        } else {
            for fi17 := 0; fi17 < 4; fi17++ {
                switch "" {
                case "hello":
                    v3 -= (v3 & (v3 / ((v3) | 1)))
                case "":
                    var v18 uint8 = 255
                    _ = v18
                case "\n\t":
                    _ = len(v2)
                    var v19 uint32 = 2147483648
                    v15 = 530
                    _ = v19
                }
                fmt.Println(v9, v15, v1)
                v9 = 1e10
                _ = fi17
            }
        }
        _ = fi14
        _ = v15
    }
    if (-1 < -509) {
        v20 := v4["hello"]
        if v1 {
            fmt.Println(v3, v20, v5)
            _ = len(v4)
            var v21 int8 = -128
            _ = v21
        } else {
            v20 -= v20
            v1 = v1
            if v1 {
                v20 = (v20 / ((v20) | 1))
                if v1 {
                    _ = len(v2)
                    v1 = true
                    fmt.Println(v3)
                }
            } else {
                for fi22 := 0; fi22 < 17; fi22++ {
                    if (!(!v1)) {
                        v23 := v6[((("hello" + "a") + "foo bar") + (("" + "") + "foo bar"))]
                        _ = v23
                    } else {
                        delete(v4, "hello")
                        v20 -= v20
                    }
                    var v24 uint64 = 627
                    _ = fi22
                    _ = v24
                }
                delete(v6, "a")
                for fi25 := 0; fi25 < 9; fi25++ {
                    v20 += 949
                    _ = fi25
                }
            }
        }
        _ = v20
    } else {
        _ = len(v2)
        v26 := GenS2{F0: 0}
        _ = v26
    }
    fmt.Println(v5, v9, v3)
    fmt.Println(v9, v9, v1)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v27 float32 = v9
    if v1 {
        for fi28 := 0; fi28 < 3; fi28++ {
            v5 += (fi28 ^ (fi28 >> ((255) & 63)))
            _ = fi28
        }
        v9 = v27
    }
    if (42 <= 728) {
        var v29 string = "hello"
        _ = v29
    }
    var v30 uint64 = 341
    v30 = (v30 & 529)
    v31 := GenS1{F0: "\n\t"}
    v5 -= (v5 / ((v5) | 1))
    var v32 int32 = 0
    if len(v2) > 0 {
        v2[0] = ("foo bar" + ("a" + ("foo bar" + "foo bar")))
    }
    fmt.Println(v5, v32)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
