package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint16 }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 uint16 }
func genFn0(p0 bool, p1 uint8) uint32 {
    p0 = p0
    switch (("foo bar" + "") + "foo bar") {
    case "\n\t":
        p0 = (764 <= -1)
    default:
        p0 = p0
    }
    fmt.Println(p1)
    p1 -= (p1 ^ p1)
    _ = p0
    _ = p1
    return uint32(p1)
}

func main() {
    v1 := make(map[int]bool)
    v2 := []int{617, 4294967295}
    var v3 uint32 = 985
    v3 = (v3 * (^(v3)))
    var v4 int32 = 878
    var v5 float32 = ((0.1 * ((1e10 - 1.5) / 0.1)) + -2.25)
    v6 := GenS1{F0: ("a" + "foo bar"), F1: (!(((v5 * v5) - 0.1) <= (v5 * float32(1.5)))), F2: 745}
    if (42 <= 9223372036854775807) {
        if (((!(42 < -1)) || true) || (true && true)) {
            switch ("a" + ("\n\t" + "foo bar")) {
            case "foo bar":
                switch -1 {
                case 799:
                    switch "a" {
                    case "hello":
                        if true {
                            v3 -= v3
                            var v7 bool = (((795 != (v3 * (^(v3)))) || (189 > int64(v3))) && (32768 < 32767))
                            _ = v7
                        }
                    default:
                        v8 := GenS1{F0: (("a" + "hello") + ""), F1: (false && ((1 < 69) || (100 >= -123))), F2: 0}
                        _ = v8
                    }
                case -926:
                    if true {
                        if true {
                            v5 += float32(1.5)
                            if (false || (!(!false))) {
                                var v9 string = ""
                                _ = v9
                            } else {
                                for fi10 := 0; fi10 < 13; fi10++ {
                                    v4 = v4
                                    var v11 int8 = 100
                                    v11 = v11
                                    _ = fi10
                                    _ = v11
                                }
                                for fi12 := 0; fi12 < 16; fi12++ {
                                    if true {
                                        fmt.Println(v5, v4, v3)
                                    } else {
                                        if (-85 > 127) {
                                            v5 = -2.25
                                            v3 = v3
                                            var v13 uint64 = 42
                                            _ = v13
                                        }
                                        v2 = append(v2, fi12)
                                    }
                                    fmt.Println(fi12, fi12)
                                    v14 := GenS1{F0: "a", F1: (1 != 1), F2: 0}
                                    _ = fi12
                                    _ = v14
                                }
                            }
                            var v15 string = "a"
                            _ = v15
                        } else {
                            switch 0 {
                            case 127:
                                delete(v1, -2147483648)
                            }
                            for fi16 := 0; fi16 < 10; fi16++ {
                                switch ((("\n\t" + "foo bar") + ("" + ("a" + ("foo bar" + ((("foo bar" + "foo bar") + "a") + "hello"))))) + ("a" + ((("foo bar" + "\n\t") + "\n\t") + ("foo bar" + "\n\t")))) {
                                case "":
                                    fmt.Println(fi16)
                                case "a":
                                    v17 := GenS2{F0: -355, F1: int(v4)}
                                    _ = v17
                                case "hello":
                                    for fi18 := 0; fi18 < 10; fi18++ {
                                        if true {
                                            if len(v2) > 0 {
                                                v2[0] = fi16
                                            }
                                        }
                                        _ = fi18
                                    }
                                default:
                                    fmt.Println(fi16)
                                }
                                _ = fi16
                            }
                        }
                        var v19 uint64 = 627
                        for fi20 := 0; fi20 < 15; fi20++ {
                            v2 = append(v2, 823)
                            var v21 uint64 = 2
                            _ = fi20
                            _ = v21
                        }
                        _ = v19
                    }
                case 0:
                    if len(v2) > 0 {
                        v2[0] = 42
                    }
                }
            case "hello":
                v2 = append(v2, -9223372036854775808)
            case "":
                fmt.Println(v5, v5)
            default:
                v22 := GenS2{F0: -32768, F1: -2147483648}
                _ = v22
            }
        }
        v23 := GenS3{F0: 660}
        _ = v23
    }
    v24 := GenS3{F0: 975}
    v1[-1] = false
    for fi25 := 0; fi25 < 5; fi25++ {
        var v26 uint8 = 130
        switch (("\n\t" + "a") + (("" + "hello") + (("foo bar" + "\n\t") + "foo bar"))) {
        case "a":
            delete(v1, fi25)
        case "hello":
            switch ("" + ("hello" + "\n\t")) {
            case "hello":
                var v27 uint = 2
                _ = v27
            case "\n\t":
                var v28 int = (fi25 + (fi25 << ((4294967295) & 63)))
                _ = v28
            case "":
                var v29 uint8 = v26
                _ = v29
            }
        default:
            var v30 uint32 = (v3 - (^(v3)))
            _ = v30
        }
        _ = fi25
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
