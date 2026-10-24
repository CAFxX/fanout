package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint8; F1 uint }
type GenS6 struct { F0 uint16 }
func genFn0(p0 uint8, p1 uint16, p2 int64) int8 {
    for fi6 := 0; fi6 < 12; fi6++ {
        v7 := []int{(fi6 >> ((348) & 63))}
        var v8 int64 = p2
        for fi9 := 0; fi9 < 7; fi9++ {
            if true {
                var v10 uint8 = p0
                _ = v10
            } else {
                v11 := GenS1{F0: (fi6 & (fi9 ^ 127)), F1: 0}
                if true {
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                }
                switch "foo bar" {
                case "":
                    var v12 int = fi9
                    _ = v12
                case "hello":
                    v13 := GenS1{F0: (fi6 + fi6), F1: 1}
                    p2 = v8
                    var v14 int8 = 113
                    _ = v13
                    _ = v14
                case "\n\t":
                    v15 := GenS2{F0: "a", F1: 1}
                    _ = v15
                default:
                    var v16 float64 = 0.1
                    _ = v16
                }
                _ = v11
            }
            p2 = p2
            var v17 int8 = 0
            _ = fi9
            _ = v17
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    if ((("foo bar" + "\n\t") != "a") && false) {
        fmt.Println(p1, p0)
        v18 := []int{}
        _ = v18
    } else {
        var v19 uint64 = 546
        _ = v19
    }
    v20 := GenS1{F0: -307, F1: 1}
    if false {
        fmt.Println(p0, p1, p1)
        p1 = 32767
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    return 124
}

func main() {
    v1 := make(map[string]bool)
    var v2 int = -856
    v3 := make(map[uint32]int)
    v4 := []string{("" + "foo bar"), "", "", ("" + ((("foo bar" + (("hello" + "a") + (("a" + "hello") + ("hello" + "foo bar")))) + (("" + "a") + ("" + "hello"))) + ""))}
    v5 := []uint32{65535, 0}
    for fi21 := 0; fi21 < 16; fi21++ {
        var v22 uint16 = 1
        delete(v1, (("foo bar" + "") + "hello"))
        var v23 string = ((("" + "") + ("a" + ("foo bar" + "foo bar"))) + ("a" + "foo bar"))
        _ = fi21
        _ = v22
        _ = v23
    }
    v2 = 9223372036854775807
    v24 := GenS3{F0: 0, F1: 382}
    if (533 > 565) {
        v2 = v2
        _ = len(v4)
        _ = len(v1)
    } else {
        v5 = append(v5, 276)
    }
    if false {
        if (!false) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v2 = 127
            if (true || false) {
                if true {
                    v25 := GenS4{F0: 100}
                    for fi26 := 0; fi26 < 18; fi26++ {
                        var v27 uint = 175
                        _ = fi26
                        _ = v27
                    }
                    _ = v25
                } else {
                    var v28 int = v2
                    if (!(32768 < (^(v2)))) {
                        var v29 float32 = 3.14159
                        if ((!(!(1 > 12))) || true) {
                            v29 = v29
                            v30 := GenS5{F0: 23, F1: 410}
                            _ = v30
                        } else {
                            var v31 int = v28
                            _ = v31
                        }
                        var v32 int32 = -2147483648
                        _ = v29
                        _ = v32
                    }
                    v33 := GenS6{F0: 506}
                    _ = v28
                    _ = v33
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
        } else {
            v1["\n\t"] = (10 != 78)
        }
    } else {
        v34 := GenS6{F0: 32767}
        fmt.Println(v2)
        for fi35 := 0; fi35 < 6; fi35++ {
            for fi36 := 0; fi36 < 13; fi36++ {
                v2 = fi35
                v2 = (fi36 ^ fi36)
                v2 += (fi35 << ((983) & 63))
                _ = fi36
            }
            _ = fi35
        }
        _ = v34
    }
    v37 := GenS4{F0: 42}
    var v38 uint32 = 451
    v38 = v38
    v2 += (v2 & (^(v2)))
    v38 = v38
    for fi39 := 0; fi39 < 4; fi39++ {
        var v40 int8 = 74
        _ = fi39
        _ = v40
    }
    var v41 float32 = -2.25
    fmt.Println(v2, v2, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v41
    _ = v5
}
