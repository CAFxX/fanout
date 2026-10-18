package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 int, p1 uint16) string {
    for fi5 := 0; fi5 < 17; fi5++ {
        fmt.Println(p1, fi5, p1)
        switch p0 {
        case 330:
            fmt.Println(fi5)
        default:
            for fi6 := 0; fi6 < 7; fi6++ {
                v7 := GenS1{F0: (2147483647 > 2147483647)}
                var v8 uint64 = 4294967295
                fmt.Println(v8, p1)
                _ = fi6
                _ = v7
                _ = v8
            }
        }
        _ = fi5
    }
    p1 = (p1 + 65535)
    _ = p0
    _ = p1
    return "\n\t"
}
func genFn1(p0 uint16, p1 int32) string {
    p0 -= (p0 & p0)
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    return (((("hello" + "\n\t") + "\n\t") + "hello") + ("foo bar" + "hello"))
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int = 4294967295
    var v3 uint64 = 42
    var v4 int16 = -1
    for fi9 := 0; fi9 < 3; fi9++ {
        fmt.Println(fi9)
        _ = fi9
    }
    for fi10 := 0; fi10 < 13; fi10++ {
        v4 -= v4
        _ = fi10
    }
    v11 := []int{}
    if true {
        v3 -= 1
        fmt.Println(v3, v2, v4)
    }
    v12 := GenS1{F0: true}
    switch "\n\t" {
    case "\n\t":
        var v13 int32 = 42
        _ = v13
    case "a":
        v2 += (v2 * v2)
    default:
        if false {
            if (v3 >= v3) {
                delete(v1, "")
                var v14 float64 = (1.5 * ((float64(1e10) - 1e10) / 1.0))
                var v15 uint16 = 1
                _ = v14
                _ = v15
            } else {
                if false {
                    var v16 int8 = -81
                    _ = v16
                } else {
                    for fi17 := 0; fi17 < 13; fi17++ {
                        if (((true && true) && (2147483648 < 2147483648)) || (1 >= 42)) {
                            var v18 string = ("hello" + ("foo bar" + "hello"))
                            var v19 float32 = 0.0
                            _ = v18
                            _ = v19
                        } else {
                            fmt.Println(v4, v4, v4)
                            v20 := GenS1{F0: false}
                            switch (fi17 - 705) {
                            case -768:
                                fmt.Println(v3, v3, v2)
                            case -423:
                                fmt.Println(v3, fi17)
                            case 128:
                                var v21 uint16 = uint16(v3)
                                _ = v21
                            }
                            _ = v20
                        }
                        v3 = v3
                        _ = fi17
                    }
                    v11 = append(v11, (v2 - v2))
                    if (-128 == 1) {
                        v4 += 32767
                    }
                }
                for fi22 := 0; fi22 < 15; fi22++ {
                    switch (("a" + "foo bar") + ("foo bar" + "\n\t")) {
                    case "\n\t":
                        if true {
                            v4 += (v4 * 628)
                            fmt.Println(v2)
                        }
                    case "a":
                        v3 = (v3 + (v3 ^ 4294967295))
                    case "":
                        v4 = 60
                    default:
                        fmt.Println(fi22)
                    }
                    if false {
                        var v23 int = fi22
                        _ = v23
                    }
                    for fi24 := 0; fi24 < 8; fi24++ {
                        var v25 int32 = 0
                        fmt.Println(fi24)
                        var v26 int64 = 9223372036854775807
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                    _ = fi22
                }
                if true {
                    var v27 int64 = 106
                    var v28 int16 = (v4 * (v4 - v4))
                    var v29 int64 = -26
                    _ = v27
                    _ = v28
                    _ = v29
                } else {
                    fmt.Println(v4, v4, v3)
                }
            }
            v2 = v2
        }
    }
    var v30 float32 = (3.14159 * 0.0)
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    if true {
        v31 := GenS1{F0: (true && (-1 < 2147483647))}
        _ = v31
    } else {
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
        if ((42 == 128) || true) {
            var v32 float32 = float32(float64(v30))
            var v33 uint16 = 664
            var v34 int64 = -9223372036854775808
            _ = v32
            _ = v33
            _ = v34
        }
        var v35 int32 = 2147483647
        _ = v35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v30
    _ = v4
}
