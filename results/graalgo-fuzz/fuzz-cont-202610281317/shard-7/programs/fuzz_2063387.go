package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16; F1 int }
func genFn0(p0 int, p1 int32, p2 uint32) uint32 {
    p2 += p2
    if (!true) {
        p0 += (p0 << ((42) & 63))
        var v6 uint64 = 2
        var v7 float64 = (1e10 / 0.1)
        _ = v6
        _ = v7
    } else {
        p0 = p0
    }
    var v8 int64 = -494
    var v9 string = "a"
    switch (v9 + v9) {
    case "":
        fmt.Println(p2)
    case "\n\t":
        fmt.Println(p1, p1)
    default:
        var v10 int32 = p1
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return p2
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint = 2
    var v3 int64 = 9223372036854775807
    v4 := make(map[string]string)
    var v5 uint8 = 125
    var v11 float32 = (0.0 * ((float32((1e10 / 0.1)) * (0.1 / -2.25)) - (0.1 + 3.14159)))
    var v12 int32 = 42
    v13 := GenS1{F0: 65535}
    switch (("" + "\n\t") + "") {
    case "hello":
        if (300 >= 4294967295) {
            for fi14 := 0; fi14 < 14; fi14++ {
                fmt.Println(v3)
                var v15 uint16 = 32767
                _ = fi14
                _ = v15
            }
            fmt.Println(v3)
            for fi16 := 0; fi16 < 13; fi16++ {
                v17 := GenS1{F0: 32767}
                if true {
                    switch (fi16 / ((-171) | 1)) {
                    case -1:
                        fmt.Println(v11, v5)
                    case 32767:
                        var v18 int8 = 127
                        _ = v18
                    case 0:
                        for fi19 := 0; fi19 < 18; fi19++ {
                            v20 := GenS1{F0: 13}
                            switch (^(fi19)) {
                            case 610:
                                var v21 uint64 = 2
                                _ = v21
                            case -374:
                                v22 := []int{fi16, (fi19 / (((fi16 >> ((v2) & 63))) | 1))}
                                _ = v22
                            default:
                                switch "a" {
                                case "a":
                                    v23 := GenS1{F0: 868}
                                    _ = v23
                                }
                            }
                            v24 := GenS2{F0: 0}
                            _ = fi19
                            _ = v20
                            _ = v24
                        }
                    default:
                        switch "a" {
                        case "\n\t":
                            v25 := []int{(fi16 << ((4294967295) & 63)), fi16}
                            _ = v25
                        case "a":
                            var v26 int = (fi16 % ((fi16) | 1))
                            _ = v26
                        }
                    }
                } else {
                    var v27 int8 = 1
                    v2 += 434
                    v3 = (v3 << ((v2) & 63))
                    _ = v27
                }
                _ = fi16
                _ = v17
            }
        }
    case "\n\t":
        v11 = float32(0.0)
    default:
        v11 += v11
    }
    v28 := []int{-273, 127, 1}
    if ("foo bar" >= "\n\t") {
        v29 := v4[(("foo bar" + "hello") + "a")]
        var v30 string = v29
        v31 := GenS3{F0: -32768, F1: 9223372036854775807}
        _ = v29
        _ = v30
        _ = v31
    } else {
        for fi32 := 0; fi32 < 14; fi32++ {
            v33 := GenS2{F0: 32767}
            switch ("a" + "hello") {
            case "foo bar":
                switch "a" {
                case "a":
                    v3 += (v3 - (v3 - 738))
                }
            case "\n\t":
                switch fi32 {
                case 4294967295:
                    v11 -= float32(-2.25)
                case -9223372036854775808:
                    v34 := GenS3{F0: -238, F1: -369}
                    _ = v34
                case 1:
                    v35 := GenS3{F0: -32768, F1: (fi32 % (((-(fi32))) | 1))}
                    _ = v35
                default:
                    for fi36 := 0; fi36 < 10; fi36++ {
                        var v37 int64 = (v3 / (((v3 << (((v2 % ((0) | 1))) & 63))) | 1))
                        var v38 int8 = 1
                        var v39 string = (("foo bar" + ("foo bar" + (("\n\t" + (("a" + ("\n\t" + ("foo bar" + ""))) + "\n\t")) + "\n\t"))) + ((((("\n\t" + "\n\t") + "") + "\n\t") + "hello") + "a"))
                        _ = fi36
                        _ = v37
                        _ = v38
                        _ = v39
                    }
                }
            default:
                var v40 bool = (!("" != "a"))
                _ = v40
            }
            if false {
                if false {
                    v11 = -2.25
                } else {
                    v12 = -2147483648
                    var v41 int64 = v3
                    fmt.Println(fi32, v5)
                    _ = v41
                }
                var v42 bool = false
                _ = v42
            }
            _ = fi32
            _ = v33
        }
    }
    fmt.Println(v2, v3)
    switch -89 {
    case -549:
        if ((!true) || true) {
            v11 -= (v11 + v11)
            v2 = (^(v2))
        } else {
            v2 = (v2 % ((211) | 1))
            _ = len(v28)
            var v43 uint64 = 0
            _ = v43
        }
    default:
        fmt.Println(v11, v3)
    }
    v5 += v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
