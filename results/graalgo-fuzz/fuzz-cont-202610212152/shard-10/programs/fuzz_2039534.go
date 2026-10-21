package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 int8; F1 int8; F2 uint }
func genFn0(p0 uint64, p1 int8) uint8 {
    fmt.Println(p1, p0)
    fmt.Println(p0)
    if false {
        p1 -= (p1 >> ((18446744073709551615) & 63))
    }
    _ = p0
    _ = p1
    return 128
}
func genFn1() uint32 {
    var v6 int64 = 132
    fmt.Println(v6)
    v6 = 42
    fmt.Println(v6)
    v7 := GenS1{F0: 42, F1: (v6 >> ((255) & 63))}
    _ = v6
    _ = v7
    return 2
}

func main() {
    v1 := []string{("" + "foo bar")}
    var v2 int64 = -279
    var v3 int32 = 1
    var v4 int8 = -86
    var v5 uint8 = 255
    if false {
        var v8 uint32 = 916
        _ = v8
    } else {
        var v9 bool = false
        if v9 {
            fmt.Println(v2, v5)
            switch "foo bar" {
            case "":
                for fi10 := 0; fi10 < 16; fi10++ {
                    var v11 uint = 633
                    v4 += (-(v4))
                    for fi12 := 0; fi12 < 19; fi12++ {
                        fmt.Println(fi10)
                        if v9 {
                            v1 = append(v1, "hello")
                            v13 := GenS1{F0: 2, F1: 0}
                            var v14 int8 = (-(v4))
                            _ = v13
                            _ = v14
                        }
                        v15 := GenS1{F0: 2, F1: v2}
                        _ = fi12
                        _ = v15
                    }
                    _ = fi10
                    _ = v11
                }
            case "hello":
                for fi16 := 0; fi16 < 5; fi16++ {
                    if len(v1) > 0 {
                        v1[0] = "hello"
                    }
                    _ = fi16
                }
            case "foo bar":
                var v17 int64 = (v2 & (-(v2)))
                _ = v17
            }
        }
        if len(v1) > 0 {
            v1[0] = ("a" + ("" + ("a" + (("hello" + (((("\n\t" + ("a" + "foo bar")) + "foo bar") + (("\n\t" + (("hello" + "\n\t") + "foo bar")) + (("a" + (("hello" + ("a" + ("hello" + "foo bar"))) + "foo bar")) + "foo bar"))) + ("foo bar" + "\n\t"))) + ("\n\t" + "\n\t")))))
        }
        _ = v9
    }
    v5 = (v5 >> ((4294967295) & 63))
    switch (("a" + "a") + "foo bar") {
    case "\n\t":
        for fi18 := 0; fi18 < 6; fi18++ {
            _ = len(v1)
            for fi19 := 0; fi19 < 9; fi19++ {
                var v20 int64 = 1
                if len(v1) > 0 {
                    v1[0] = "\n\t"
                }
                switch fi18 {
                case 745:
                    v20 = 9223372036854775807
                }
                _ = fi19
                _ = v20
            }
            v5 -= v5
            _ = fi18
        }
    case "hello":
        switch ("\n\t" + "foo bar") {
        case "":
            var v21 uint = 740
            _ = v21
        default:
            for fi22 := 0; fi22 < 4; fi22++ {
                switch (((((("a" + ("\n\t" + ((("\n\t" + "foo bar") + ("foo bar" + ("" + "hello"))) + ""))) + ("\n\t" + (("foo bar" + ("\n\t" + "")) + ("hello" + "foo bar")))) + "\n\t") + "foo bar") + "a") + ((("a" + "hello") + "foo bar") + "a")) {
                case "hello":
                    if ((true && true) && true) {
                        var v23 uint64 = 4294967295
                        _ = v23
                    } else {
                        var v24 int64 = 0
                        v2 += v24
                        _ = v24
                    }
                default:
                    switch fi22 {
                    case -1:
                        for fi25 := 0; fi25 < 11; fi25++ {
                            if (!true) {
                                v26 := GenS1{F0: 638, F1: 1}
                                _ = v26
                            }
                            _ = fi25
                        }
                    default:
                        fmt.Println(v3)
                    }
                }
                v3 += v3
                switch fi22 {
                case 248:
                    var v27 int = (fi22 * (-(fi22)))
                    _ = v27
                default:
                    if true {
                        switch "hello" {
                        case "\n\t":
                            _ = len(v1)
                        case "foo bar":
                            v28 := GenS1{F0: 1, F1: (^(v2))}
                            _ = v28
                        case "":
                            fmt.Println(v4)
                        }
                        _ = len(v1)
                        var v29 string = "foo bar"
                        _ = v29
                    } else {
                        v4 += (v4 | (v4 & int8(v2)))
                        v30 := GenS1{F0: 699, F1: 942}
                        _ = v30
                    }
                }
                _ = fi22
            }
        }
    case "foo bar":
        _ = len(v1)
    default:
        v31 := GenS2{F0: (-(v4)), F1: -128, F2: 372}
        _ = v31
    }
    v2 = v2
    var v32 int64 = v2
    v3 = v3
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch "\n\t" {
    case "":
        v4 = (v4 | (v4 | (v4 * (-(v4)))))
    default:
        fmt.Println(v4, v5, v5)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
