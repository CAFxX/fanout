package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 string; F1 bool; F2 int32 }
type GenS3 struct { F0 uint; F1 string }
func genFn0(p0 uint8) bool {
    if (false || false) {
        var v6 uint = 2
        for fi7 := 0; fi7 < 16; fi7++ {
            fmt.Println(v6, p0)
            var v8 uint64 = 466
            _ = fi7
            _ = v8
        }
        _ = v6
    } else {
        p0 -= (^(p0))
    }
    p0 = p0
    _ = p0
    return ((p0 + (^(p0))) <= (p0 >> ((125) & 63)))
}
func genFn1() int16 {
    for fi9 := 0; fi9 < 12; fi9++ {
        if true {
            var v10 int8 = 109
            _ = v10
        }
        _ = fi9
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        var v12 int = 4294967295
        v13 := GenS1{F0: (!("\n\t" > ("\n\t" + "foo bar"))), F1: 255}
        fmt.Println(v12, v12)
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 float32 = 0.0
    var v15 int32 = 0
    fmt.Println(v15)
    _ = v14
    _ = v15
    return -171
}

func main() {
    v1 := make(map[int]int)
    var v2 int = -2147483648
    v3 := make(map[int]uint32)
    v4 := []string{}
    v5 := make(map[uint32]string)
    v16 := GenS1{F0: (!(42 != 2)), F1: 4294967295}
    switch "hello" {
    case "hello":
        if (!true) {
            for fi17 := 0; fi17 < 20; fi17++ {
                v18 := GenS1{F0: (!(255 != 0)), F1: 435}
                if false {
                    var v19 uint16 = 1
                    var v20 uint = 65535
                    v21 := GenS1{F0: (!false), F1: (^(v20))}
                    _ = v19
                    _ = v20
                    _ = v21
                } else {
                    v2 -= (v2 % ((fi17) | 1))
                }
                _ = fi17
                _ = v18
            }
        } else {
            fmt.Println(v2)
            v22 := GenS2{F0: ("a" + "\n\t"), F1: true, F2: 42}
            _ = v22
        }
    default:
        for fi23 := 0; fi23 < 11; fi23++ {
            var v24 uint32 = 65535
            if len(v4) > 0 {
                v4[0] = ""
            }
            v2 = fi23
            _ = fi23
            _ = v24
        }
    }
    var v25 uint8 = 3
    fmt.Println(v25, v25)
    if len(v4) > 0 {
        v4[0] = "\n\t"
    }
    if (false && false) {
        var v26 int32 = -2147483648
        _ = len(v4)
        fmt.Println(v25, v25)
        _ = v26
    }
    v25 += 1
    fmt.Println(v2, v2)
    v4 = append(v4, (("foo bar" + (("foo bar" + "hello") + (("" + ("\n\t" + ("a" + "\n\t"))) + "foo bar"))) + ("foo bar" + "\n\t")))
    switch "hello" {
    case "hello":
        v2 += v2
    case "a":
        v27 := v1[(v2 + v2)]
        _ = v27
    default:
        var v28 bool = (!true)
        _ = v28
    }
    var v29 uint = 4294967295
    var v30 uint8 = v25
    fmt.Println(v30, v25, v25)
    if false {
        switch v2 {
        case 128:
            if (42 == 42) {
                fmt.Println(v29)
                var v31 uint = 18446744073709551615
                v32 := GenS1{F0: true, F1: v31}
                _ = v31
                _ = v32
            } else {
                v33 := GenS1{F0: false, F1: (v29 ^ 18446744073709551615)}
                v30 = v30
                _ = v33
            }
        case -2147483648:
            switch ("hello" + "a") {
            case "a":
                v2 -= -590
            case "\n\t":
                if true {
                    v34 := v5[1]
                    _ = v34
                } else {
                    fmt.Println(v29, v2, v29)
                    switch "\n\t" {
                    case "\n\t":
                        v35 := GenS2{F0: "a", F1: (-9223372036854775808 <= 9223372036854775807), F2: 1}
                        _ = v35
                    default:
                        v4 = append(v4, ("a" + "\n\t"))
                    }
                    v36 := GenS2{F0: ("\n\t" + ("a" + "foo bar")), F1: (true && ((false && ((^(v30)) <= v25)) || false)), F2: -2147483648}
                    _ = v36
                }
            case "":
                _ = len(v4)
            }
        }
        for fi37 := 0; fi37 < 14; fi37++ {
            for fi38 := 0; fi38 < 1; fi38++ {
                v30 = v30
                if (false || false) {
                    var v39 int16 = 234
                    var v40 int8 = 100
                    fmt.Println(v40, v25, v29)
                    _ = v39
                    _ = v40
                }
                v25 += (v25 << ((v29) & 63))
                if false {
                    var v41 int8 = 127
                    _ = v41
                }
                _ = fi38
            }
            if len(v4) > 0 {
                v4[0] = "a"
            }
            v42 := GenS3{F0: v29, F1: ("\n\t" + ("hello" + "foo bar"))}
            _ = fi37
            _ = v42
        }
        for fi43 := 0; fi43 < 19; fi43++ {
            v29 += 361
            v29 = v29
            _ = fi43
        }
    }
    var v44 uint16 = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v44
    _ = v5
}
