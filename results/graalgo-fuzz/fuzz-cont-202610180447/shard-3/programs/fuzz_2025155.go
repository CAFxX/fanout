package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 int32; F1 int; F2 string }
type GenS3 struct { F0 string; F1 int32 }
func genFn0(p0 int16, p1 int32) int {
    v7 := GenS1{F0: false, F1: false}
    p1 = p1
    var v8 uint64 = 4294967295
    p0 -= (p0 & p0)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 32767
}
func genFn1(p0 int8, p1 uint64) int8 {
    p0 += p0
    var v9 int64 = -1
    _ = p0
    _ = p1
    _ = v9
    return (p0 / ((0) | 1))
}

func main() {
    v1 := []string{"\n\t", "hello", ("hello" + (("foo bar" + "a") + ("hello" + "a")))}
    var v2 uint64 = 1
    var v3 bool = false
    v4 := make(map[string]int)
    var v5 bool = false
    var v6 uint32 = 2147483647
    switch 4294967295 {
    case 2147483647:
        if v5 {
            switch -417 {
            case 0:
                fmt.Println(v6, v2)
            case 1:
                v6 -= 475
            case 128:
                delete(v4, ("\n\t" + ""))
            default:
                if true {
                    var v10 uint32 = v6
                    _ = v10
                }
            }
            var v11 int8 = genFn1(42, (v2 - (^(v2))))
            v11 = (v11 << ((255) & 63))
            _ = v11
        } else {
            for fi12 := 0; fi12 < 13; fi12++ {
                if v5 {
                    v3 = ((3.14159 / (1e10 - 0.1)) < 0.0)
                }
                if len(v1) > 0 {
                    v1[0] = "a"
                }
                _ = fi12
            }
            for fi13 := 0; fi13 < 18; fi13++ {
                if v5 {
                    fmt.Println(fi13)
                    var v14 float64 = 3.14159
                    _ = v14
                }
                v6 += 246
                _ = fi13
            }
        }
    case 582:
        v5 = v5
    case 556:
        if (189 <= 255) {
            var v15 bool = v5
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            for fi16 := 0; fi16 < 10; fi16++ {
                v5 = false
                for fi17 := 0; fi17 < 17; fi17++ {
                    var v18 uint64 = v2
                    fmt.Println(v15)
                    if false {
                        switch ("hello" + "") {
                        case "hello":
                            if false {
                                var v19 bool = (42 != -49)
                                if (65 < genFn1(-1, (v18 >> ((42) & 63)))) {
                                    switch (fi17 | fi17) {
                                    case 32767:
                                        v20 := GenS1{F0: (2 == 127), F1: false}
                                        _ = v20
                                    case -876:
                                        fmt.Println(v6, v18)
                                        v19 = v15
                                    }
                                } else {
                                    v21 := GenS1{F0: v3, F1: false}
                                    var v22 uint32 = v6
                                    _ = v21
                                    _ = v22
                                }
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = v19
                            }
                        case "":
                            var v23 float32 = 1e10
                            _ = v23
                        case "foo bar":
                            v4["a"] = -1
                        }
                        _ = len(v4)
                        v24 := GenS2{F0: 1, F1: (fi17 & -122), F2: ""}
                        _ = v24
                    }
                    _ = fi17
                    _ = v18
                }
                _ = fi16
            }
            _ = v15
        }
    default:
        var v25 int64 = 1
        _ = v25
    }
    v26 := v4[""]
    fmt.Println(v3, v3)
    for fi27 := 0; fi27 < 10; fi27++ {
        if v3 {
            if ((true && (!v5)) || (!(-631 > 2147483647))) {
                var v28 int = 42
                _ = v28
            }
        } else {
            var v29 int64 = -317
            _ = v29
        }
        _ = fi27
    }
    v2 = 656
    delete(v4, ("a" + "\n\t"))
    v5 = (42 <= -855)
    fmt.Println(v26, v3, v26)
    if ((0.1 + ((1e10 / float64(((float32(0.0) + 1.5) - 0.0))) * 1e10)) > float64(1.5)) {
        v3 = v3
    } else {
        var v30 uint16 = 0
        if true {
            var v31 uint = 430
            _ = v31
        } else {
            var v32 uint16 = 475
            for fi33 := 0; fi33 < 3; fi33++ {
                for fi34 := 0; fi34 < 15; fi34++ {
                    v35 := GenS3{F0: ("foo bar" + "hello"), F1: 42}
                    v26 += v26
                    v6 = (^(v6))
                    v36 := GenS3{F0: "hello", F1: -2147483648}
                    _ = fi34
                    _ = v35
                    _ = v36
                }
                _ = fi33
            }
            _ = v32
        }
        if v5 {
            var v37 uint64 = 145
            var v38 int8 = 127
            v39 := v4[(("foo bar" + "") + ((((("a" + "") + "foo bar") + "") + "hello") + ""))]
            _ = v37
            _ = v38
            _ = v39
        }
        _ = v30
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
