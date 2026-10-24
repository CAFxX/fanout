package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 int8, p1 int, p2 uint16) int8 {
    fmt.Println(p2, p2, p2)
    if true {
        var v7 bool = ((!((-(p0)) != (-(p0)))) && (!(1 != -1)))
        _ = v7
    }
    fmt.Println(p0, p1, p0)
    var v8 int = p1
    var v9 int8 = (^(p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return p0
}

func main() {
    var v1 bool = false
    v2 := make(map[uint32]bool)
    v3 := make(map[int]int)
    v4 := []string{"a"}
    v5 := []string{(((("\n\t" + "") + "\n\t") + "") + "a"), "hello", "\n\t"}
    v6 := []string{(("" + ("foo bar" + ((("foo bar" + ("\n\t" + (("a" + ("hello" + (("foo bar" + "foo bar") + ""))) + (((("foo bar" + "\n\t") + (("a" + "\n\t") + "")) + ("" + ("hello" + "hello"))) + "a")))) + "") + (("\n\t" + ("foo bar" + "foo bar")) + ("" + "hello"))))) + (("foo bar" + "\n\t") + ("foo bar" + ""))), "foo bar", "hello"}
    switch -808 {
    case 42:
        switch "foo bar" {
        case "\n\t":
            fmt.Println(v1)
        case "a":
            v1 = (42 > 1)
        }
    case 4294967295:
        var v10 float64 = (((float64((1e10 * 1.5)) / (1e10 * 1e10)) - 0.0) / float64(-2.25))
        _ = v10
    }
    v1 = (42 != 4294967295)
    if v1 {
        for fi11 := 0; fi11 < 4; fi11++ {
            for fi12 := 0; fi12 < 4; fi12++ {
                if (!v1) {
                    var v13 bool = (float32(3.14159) != 0.1)
                    if v1 {
                        var v14 uint8 = 127
                        v13 = v13
                        var v15 float64 = (-2.25 * 0.0)
                        _ = v14
                        _ = v15
                    }
                    v1 = (255 == 147)
                    _ = v13
                } else {
                    fmt.Println(fi11, fi12)
                    var v16 float32 = ((0.1 + 0.0) * 1.5)
                    _ = v16
                }
                _ = fi12
            }
            v1 = v1
            _ = fi11
        }
        v17 := v2[247]
        if ("\n\t" >= "\n\t") {
            if true {
                v17 = (true && ((2 < 537) || (false && v1)))
                if (!v17) {
                    v17 = (2147483647 >= 2)
                } else {
                    v18 := GenS1{F0: -54}
                    _ = v18
                }
            }
            v1 = (0.1 <= -2.25)
        } else {
            fmt.Println(v17, v17)
            v19 := GenS1{F0: 32767}
            if len(v5) > 0 {
                v5[0] = (("\n\t" + "foo bar") + "\n\t")
            }
            _ = v19
        }
        _ = v17
    } else {
        if (65535 == 409) {
            v1 = v1
        } else {
            if v1 {
                v20 := GenS2{F0: 1, F1: 2}
                switch ("hello" + "") {
                case "foo bar":
                    v3[-646] = 128
                case "\n\t":
                    v1 = (838 < 2)
                default:
                    v1 = (!v1)
                }
                _ = v20
            } else {
                for fi21 := 0; fi21 < 15; fi21++ {
                    v22 := GenS1{F0: 42}
                    switch "\n\t" {
                    case "\n\t":
                        for fi23 := 0; fi23 < 15; fi23++ {
                            fmt.Println(fi21, fi21, fi23)
                            v1 = v1
                            _ = fi23
                        }
                    case "a":
                        v1 = true
                    default:
                        v1 = false
                    }
                    fmt.Println(v1, v1)
                    _ = fi21
                    _ = v22
                }
                var v24 int8 = 127
                _ = v24
            }
            for fi25 := 0; fi25 < 3; fi25++ {
                fmt.Println(fi25, v1)
                v26 := GenS3{F0: 632}
                for fi27 := 0; fi27 < 11; fi27++ {
                    var v28 bool = true
                    if v28 {
                        var v29 uint64 = 3
                        if ((1e10 + 1.5) < (3.14159 / 1.0)) {
                            v1 = v1
                        } else {
                            v28 = v1
                            if (-9223372036854775808 > -1) {
                                switch 0 {
                                case 4294967295:
                                    fmt.Println(v1, fi27)
                                case 771:
                                    fmt.Println(fi27, fi25, fi27)
                                default:
                                    v1 = (1 < 1)
                                }
                            } else {
                                v5 = append(v5, ("\n\t" + (((("hello" + "\n\t") + "") + "hello") + ("hello" + ("foo bar" + "\n\t")))))
                            }
                        }
                        v30 := GenS1{F0: 0}
                        _ = v29
                        _ = v30
                    }
                    var v31 uint32 = 2
                    _ = fi27
                    _ = v28
                    _ = v31
                }
                _ = fi25
                _ = v26
            }
        }
        for fi32 := 0; fi32 < 11; fi32++ {
            v1 = (2 != 255)
            v6 = append(v6, "hello")
            _ = fi32
        }
    }
    delete(v3, 128)
    v33 := GenS1{F0: 42}
    v34 := GenS2{F0: 42, F1: 65535}
    v1 = v1
    v1 = (-2.25 < (float32(3.14159) - 0.0))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
