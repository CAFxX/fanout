package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 int8 }
func genFn0(p0 int8, p1 uint16, p2 int64) int64 {
    for fi7 := 0; fi7 < 20; fi7++ {
        switch ("foo bar" + "a") {
        case "\n\t":
            if false {
                for fi8 := 0; fi8 < 1; fi8++ {
                    v9 := []int{fi8, fi8, (fi7 + (fi8 ^ fi8))}
                    var v10 uint8 = 35
                    if (-2.25 != 0.1) {
                        if (("" + "") != "a") {
                            var v11 int = fi8
                            var v12 int16 = 42
                            fmt.Println(v12)
                            _ = v11
                            _ = v12
                        } else {
                            var v13 int16 = 32767
                            fmt.Println(p1, fi8, p1)
                            switch (fi7 / (((fi8 % ((fi8) | 1))) | 1)) {
                            case -388:
                                if len(v9) >= 2 {
                                    v9 = v9[1:len(v9)]
                                }
                            case -9223372036854775808:
                                if true {
                                    if false {
                                        v10 += 1
                                        if false {
                                            var v14 int16 = (v13 + v13)
                                            _ = v14
                                        } else {
                                            switch fi7 {
                                            case -2147483648:
                                                v13 = v13
                                            }
                                        }
                                    }
                                    var v15 bool = (true || true)
                                    p1 += uint16(v10)
                                    switch "hello" {
                                    case "\n\t":
                                        var v16 bool = (2147483648 == 2147483647)
                                        _ = v16
                                    case "hello":
                                        p1 -= 618
                                    }
                                    switch (fi8 - 4294967295) {
                                    case 635:
                                        if v15 {
                                            p1 = 42
                                        } else {
                                            fmt.Println(v15)
                                        }
                                    default:
                                        v17 := GenS1{F0: 42}
                                        _ = v17
                                    }
                                    _ = v15
                                } else {
                                    switch -9223372036854775808 {
                                    case 676:
                                        var v18 int = (-(fi8))
                                        _ = v18
                                    default:
                                        fmt.Println(v10, p0, p2)
                                    }
                                    if ((fi7 / (((fi7 - fi7)) | 1)) < -9223372036854775808) {
                                        if false {
                                            var v19 uint8 = (v10 + v10)
                                            _ = v19
                                        } else {
                                            p1 = (^(p1))
                                        }
                                        var v20 int8 = p0
                                        _ = v20
                                    } else {
                                        var v21 int32 = 131
                                        fmt.Println(v21, v21, p0)
                                        var v22 int32 = (v21 & v21)
                                        _ = v21
                                        _ = v22
                                    }
                                    if (!false) {
                                        fmt.Println(p0, fi8, v10)
                                        p1 -= p1
                                        _ = len(v9)
                                    }
                                    if (int8(v10) != p0) {
                                        if true {
                                            fmt.Println(fi7)
                                            v23 := GenS1{F0: 237}
                                            _ = len(v9)
                                            _ = v23
                                        } else {
                                            _ = len(v9)
                                        }
                                        var v24 int32 = -486
                                        _ = v24
                                    }
                                }
                            case -631:
                                var v25 uint16 = (^(p1))
                                _ = v25
                            }
                            _ = v13
                        }
                        _ = len(v9)
                    }
                    _ = fi8
                    _ = v10
                    _ = v9
                }
            }
        case "":
            var v26 string = ""
            _ = v26
        case "foo bar":
            p1 -= (p1 % ((p1) | 1))
        default:
            v27 := []int{32767, (^(fi7))}
            _ = v27
        }
        _ = fi7
    }
    var v28 uint64 = 4294967295
    v29 := []int{0}
    _ = p0
    _ = p1
    _ = p2
    _ = v28
    _ = v29
    return p2
}

func main() {
    var v1 uint16 = 590
    var v2 uint16 = v1
    var v3 uint8 = 128
    var v4 int = 148
    v5 := make(map[int]string)
    var v6 int64 = 42
    var v30 string = ((((("a" + "") + (("foo bar" + ("hello" + "a")) + (("\n\t" + ("a" + "foo bar")) + "\n\t"))) + ((("hello" + "foo bar") + ("hello" + "\n\t")) + "\n\t")) + ((("a" + "foo bar") + "") + ("a" + ("foo bar" + "\n\t")))) + "hello")
    var v31 bool = false
    v2 = uint16(v3)
    _ = len(v5)
    v32 := GenS2{F0: 1000, F1: 42}
    var v33 int32 = 463
    v34 := v5[-550]
    var v35 int32 = v33
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
