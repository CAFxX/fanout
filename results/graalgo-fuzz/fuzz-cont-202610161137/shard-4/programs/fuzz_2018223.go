package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 uint64; F2 uint32 }
type GenS3 struct { F0 uint16; F1 uint64 }
func genFn0(p0 bool) int8 {
    fmt.Println(p0)
    p0 = (((("\n\t" + "hello") + (("a" + "hello") + ("a" + ("hello" + ("\n\t" + "hello"))))) != "a") && true)
    var v6 float64 = 0.0
    _ = p0
    _ = v6
    return -38
}
func genFn1(p0 int32, p1 int32, p2 int16) uint16 {
    v7 := []int{-1, 986}
    for fi8 := 0; fi8 < 4; fi8++ {
        p2 -= (^(p2))
        _ = fi8
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        var v10 uint64 = 4294967295
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 65535
}

func main() {
    var v1 uint32 = 749
    v2 := []string{"", "\n\t", ("a" + "a"), ("hello" + ("foo bar" + ("hello" + ("hello" + (((("foo bar" + ("a" + ((("a" + ("hello" + ((("a" + ("" + ("hello" + "\n\t"))) + "\n\t") + "foo bar"))) + "a") + "hello"))) + "hello") + "hello") + "a")))))}
    var v3 string = ""
    var v4 uint32 = v1
    v5 := []int{4294967295}
    if len(v5) > 0 {
        v5[0] = 9223372036854775807
    }
    v1 = v4
    for fi11 := 0; fi11 < 18; fi11++ {
        if (!(42 == 42)) {
            v12 := GenS1{F0: 42}
            for fi13 := 0; fi13 < 1; fi13++ {
                var v14 int16 = -867
                v3 = (v3 + ((v3 + "foo bar") + "a"))
                _ = fi13
                _ = v14
            }
            if false {
                v4 += v4
            }
            _ = v12
        }
        for fi15 := 0; fi15 < 11; fi15++ {
            var v16 string = v3
            if true {
                if true {
                    v1 += v1
                    switch 104 {
                    case 0:
                        v3 = (v16 + v16)
                    case 32768:
                        if (254 > 1) {
                            if true {
                                v17 := GenS1{F0: 185}
                                fmt.Println(v1)
                                _ = v17
                            } else {
                                v2 = append(v2, "foo bar")
                            }
                            if len(v5) > 0 {
                                v5[0] = fi15
                            }
                        } else {
                            var v18 int32 = 2147483647
                            _ = v18
                        }
                    case 581:
                        fmt.Println(fi11, fi15)
                    default:
                        v4 = (v4 >> ((0) & 63))
                    }
                } else {
                    _ = len(v5)
                    v19 := GenS1{F0: 254}
                    var v20 float64 = (1e10 + float64(0.1))
                    v21 := GenS1{F0: 0}
                    switch "a" {
                    case "foo bar":
                        _ = len(v5)
                    case "a":
                        var v22 uint16 = 1
                        _ = v22
                    case "\n\t":
                        v4 = (v1 ^ 65535)
                    default:
                        v5 = append(v5, (fi11 + fi15))
                    }
                    _ = v19
                    _ = v20
                    _ = v21
                }
                v23 := GenS2{F0: 0, F1: 0, F2: 4294967295}
                var v24 float32 = (1e10 + 0.1)
                _ = v23
                _ = v24
            } else {
                v2 = append(v2, v16)
            }
            _ = fi15
            _ = v16
        }
        _ = fi11
    }
    v3 = (("a" + "\n\t") + (("" + "foo bar") + ""))
    var v25 bool = (3.14159 != 3.14159)
    for fi26 := 0; fi26 < 11; fi26++ {
        switch fi26 {
        case 2147483647:
            v1 += v4
        }
        v27 := GenS3{F0: 255, F1: 0}
        for fi28 := 0; fi28 < 5; fi28++ {
            v4 = 1
            _ = fi28
        }
        _ = fi26
        _ = v27
    }
    v29 := GenS1{F0: 16}
    var v30 uint32 = (v1 * 2147483648)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
