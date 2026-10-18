package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8; F1 int32; F2 int }
type GenS4 struct { F0 int64; F1 uint64 }
func genFn0(p0 int16, p1 bool, p2 bool) uint32 {
    p1 = p2
    for fi5 := 0; fi5 < 12; fi5++ {
        v6 := GenS1{F0: ""}
        p1 = true
        p2 = (-1 != 32768)
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    return 481
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{"", (("\n\t" + (("a" + ((("" + "") + (("a" + "\n\t") + "")) + "")) + "")) + "hello")}
    var v3 float32 = -2.25
    var v4 bool = false
    var v7 uint = 18446744073709551615
    v8 := GenS1{F0: (("a" + (("" + "\n\t") + "\n\t")) + (("foo bar" + (("a" + ("hello" + (("hello" + "foo bar") + "foo bar"))) + "a")) + ""))}
    v9 := GenS1{F0: "a"}
    v3 = 0.1
    v10 := GenS1{F0: (("a" + "hello") + ("\n\t" + (("" + ((("foo bar" + (((("a" + "a") + ("hello" + "")) + ("foo bar" + "a")) + "a")) + (("" + (("hello" + "\n\t") + (("" + ("hello" + ("\n\t" + "foo bar"))) + ((("" + "foo bar") + "hello") + "")))) + "\n\t")) + "")) + "hello")))}
    if v4 {
        fmt.Println(v4)
    } else {
        v11 := GenS1{F0: "hello"}
        if len(v2) > 0 {
            v2[0] = "foo bar"
        }
        if false {
            switch ((("\n\t" + "hello") + "") + (("foo bar" + "hello") + "hello")) {
            case "":
                var v12 int16 = 32767
                _ = v12
            default:
                var v13 int64 = -1
                _ = v13
            }
        } else {
            var v14 uint32 = 638
            _ = v14
        }
        _ = v11
    }
    for fi15 := 0; fi15 < 9; fi15++ {
        if v4 {
            for fi16 := 0; fi16 < 3; fi16++ {
                v17 := GenS1{F0: "hello"}
                var v18 uint64 = 465
                v19 := GenS2{F0: "foo bar"}
                if (false || ((v7 & (v7 >> (((v7 & (v7 ^ 42))) & 63))) != 65535)) {
                    v20 := GenS2{F0: ("foo bar" + "\n\t")}
                    v18 -= (v18 + v18)
                    if v4 {
                        if (v3 == ((v3 * v3) * (v3 / v3))) {
                            v4 = (!v4)
                        } else {
                            switch fi15 {
                            case 4294967295:
                                v7 = 18446744073709551615
                            }
                            v21 := GenS3{F0: -75, F1: 5, F2: -9223372036854775808}
                            fmt.Println(fi16)
                            var v22 int16 = 0
                            _ = v21
                            _ = v22
                        }
                        var v23 float32 = -2.25
                        var v24 uint16 = 255
                        _ = v23
                        _ = v24
                    } else {
                        v4 = ((fi16 - (fi15 + fi16)) <= fi15)
                    }
                    _ = v20
                } else {
                    v18 -= 0
                }
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
            }
        } else {
            switch 0 {
            case 9223372036854775807:
                var v25 string = "hello"
                _ = v25
            case 2147483647:
                v26 := GenS4{F0: 691, F1: 4294967295}
                _ = v26
            case 0:
                v3 -= v3
            }
            if true {
                delete(v1, 0)
                fmt.Println(v3, v4)
            } else {
                var v27 int = (fi15 ^ (^(fi15)))
                _ = v27
            }
        }
        _ = fi15
    }
    _ = len(v2)
    v4 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
