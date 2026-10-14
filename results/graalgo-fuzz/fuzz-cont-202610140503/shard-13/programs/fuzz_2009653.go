package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint32; F1 string; F2 uint16 }
func genFn0(p0 int8) uint16 {
    fmt.Println(p0)
    if (34 <= 1) {
        if (0.0 <= 0.0) {
            p0 = -55
        }
        var v6 int64 = -9223372036854775808
        _ = v6
    } else {
        for fi7 := 0; fi7 < 18; fi7++ {
            var v8 int16 = 32767
            _ = fi7
            _ = v8
        }
    }
    _ = p0
    return 441
}
func genFn1(p0 bool, p1 uint) int8 {
    if false {
        var v9 string = "a"
        var v10 int8 = 100
        _ = v10
        _ = v9
    } else {
        fmt.Println(p0, p1)
        p0 = (!true)
        var v11 uint8 = 127
        _ = v11
    }
    var v12 int = 4294967295
    fmt.Println(v12)
    switch 428 {
    case -9223372036854775808:
        if (-1 == 2147483647) {
            if true {
                if (("a" + "\n\t") == "foo bar") {
                    p0 = p0
                    fmt.Println(p0, v12)
                }
            } else {
                fmt.Println(p0, p1)
                p1 = p1
                switch "\n\t" {
                case "foo bar":
                    v13 := GenS1{F0: ("a" + "\n\t")}
                    _ = v13
                case "":
                    if p0 {
                        if ("" >= ("\n\t" + "hello")) {
                            if false {
                                var v14 uint32 = 2147483648
                                var v15 int8 = 42
                                _ = v14
                                _ = v15
                            }
                            fmt.Println(p0, v12, v12)
                        }
                    } else {
                        if p0 {
                            v16 := []int{(v12 & (v12 << (((p1 / ((p1) | 1))) & 63)))}
                            _ = v16
                        } else {
                            if ("foo bar" <= ("a" + "a")) {
                                var v17 uint = 2
                                var v18 int8 = -84
                                _ = v17
                                _ = v18
                            } else {
                                if (!((!p0) && (!false))) {
                                    v19 := []int{v12}
                                    v19 = append(v19, (^(v12)))
                                    var v20 bool = false
                                    _ = v19
                                    _ = v20
                                } else {
                                    var v21 uint16 = 42
                                    v22 := []int{}
                                    _ = v21
                                    _ = v22
                                }
                            }
                            p0 = true
                            var v23 string = "hello"
                            _ = v23
                        }
                        for fi24 := 0; fi24 < 20; fi24++ {
                            fmt.Println(fi24, v12)
                            p0 = (42 > -1)
                            v25 := GenS2{F0: (p1 - 4294967295), F1: 124}
                            _ = fi24
                            _ = v25
                        }
                    }
                case "a":
                    switch v12 {
                    case -275:
                        p0 = (v12 != (v12 << (((p1 << ((932) & 63))) & 63)))
                    }
                }
            }
            fmt.Println(p1, p0)
        } else {
            if (!p0) {
                var v26 uint16 = 65535
                for fi27 := 0; fi27 < 17; fi27++ {
                    p1 = p1
                    _ = fi27
                }
                p0 = (false || (-103 <= -128))
                _ = v26
            } else {
                p1 += p1
                v28 := GenS3{F0: 0}
                _ = v28
            }
            p0 = p0
            var v29 uint64 = 97
            _ = v29
        }
    default:
        var v30 int = v12
        _ = v30
    }
    for fi31 := 0; fi31 < 2; fi31++ {
        for fi32 := 0; fi32 < 19; fi32++ {
            v33 := GenS4{F0: 2147483647, F1: "", F2: 42}
            fmt.Println(v12, v12)
            if true {
                fmt.Println(v12, v12, v12)
            } else {
                p1 -= (p1 & (p1 % ((p1) | 1)))
            }
            p0 = p0
            fmt.Println(p1)
            _ = fi32
            _ = v33
        }
        v12 += (v12 - (fi31 >> ((p1) & 63)))
        _ = fi31
    }
    _ = p0
    _ = p1
    _ = v12
    return 42
}

func main() {
    var v1 uint = 2
    var v2 int32 = 861
    v3 := make(map[string]string)
    var v4 int8 = 105
    v5 := make(map[string]int)
    for fi34 := 0; fi34 < 13; fi34++ {
        if true {
            delete(v5, "a")
            v1 -= 255
        }
        _ = fi34
    }
    v35 := GenS3{F0: 2147483647}
    if (42 < (v1 % ((v1) | 1))) {
        fmt.Println(v2, v2, v2)
        v1 = v1
    } else {
        v36 := v3[(("foo bar" + "a") + "hello")]
        var v37 float32 = (-2.25 / 1.5)
        v38 := GenS1{F0: ("foo bar" + v36)}
        _ = v36
        _ = v37
        _ = v38
    }
    v1 -= v1
    v39 := []int{}
    var v40 uint32 = 0
    v2 -= (-(v2))
    v1 = (v1 << ((v1) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
