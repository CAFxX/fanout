package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint }
type GenS2 struct { F0 bool; F1 uint8; F2 uint }
func genFn0(p0 int64, p1 int32) uint16 {
    var v4 bool = (true || false)
    switch 4294967295 {
    case -2147483648:
        var v5 uint8 = 165
        _ = v5
    case 128:
        var v6 int = 42
        _ = v6
    }
    var v7 uint64 = 2
    _ = p0
    _ = p1
    _ = v4
    _ = v7
    return 42
}
func genFn1(p0 bool, p1 int32, p2 uint16) uint64 {
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 int = (-(fi8))
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 18; fi10++ {
        var v11 string = ""
        _ = fi10
        _ = v11
    }
    var v12 int16 = 821
    if ("a" > "foo bar") {
        for fi13 := 0; fi13 < 13; fi13++ {
            v14 := GenS1{F0: 4294967295, F1: genFn0(568, (p1 << ((797) & 63))), F2: 42}
            fmt.Println(v12, fi13)
            _ = fi13
            _ = v14
        }
        if false {
            if p0 {
                switch "hello" {
                case "\n\t":
                    fmt.Println(p1, v12)
                case "foo bar":
                    p2 = (p2 & (p2 * (p2 | p2)))
                }
                v15 := GenS2{F0: ((255 <= 0) && p0), F1: 2, F2: 389}
                _ = v15
            } else {
                v16 := []int{1, -1, 32768}
                switch -376 {
                case 32767:
                    var v17 uint64 = 161
                    _ = v17
                }
                v12 += v12
                _ = v16
            }
            v18 := []int{127, 638}
            fmt.Println(p1, p1)
            _ = v18
        } else {
            if p0 {
                if (!(p0 && p0)) {
                    fmt.Println(v12, p1)
                } else {
                    if (883 > 42) {
                        for fi19 := 0; fi19 < 5; fi19++ {
                            v12 += 42
                            v20 := []int{(fi19 - -9223372036854775808), (fi19 >> ((4294967295) & 63)), (fi19 << ((432) & 63))}
                            _ = fi19
                            _ = v20
                        }
                        switch 32767 {
                        case 4294967295:
                            v21 := []int{}
                            _ = v21
                        case 42:
                            var v22 uint = 255
                            _ = v22
                        case 2147483647:
                            if false {
                                fmt.Println(p1, v12, p0)
                            } else {
                                switch "" {
                                case "\n\t":
                                    var v23 int16 = (-(v12))
                                    _ = v23
                                case "a":
                                    for fi24 := 0; fi24 < 17; fi24++ {
                                        v12 -= (v12 >> ((4294967295) & 63))
                                        p2 += genFn0(-9223372036854775808, p1)
                                        _ = fi24
                                    }
                                }
                                p0 = (255 >= (p2 ^ p2))
                            }
                        default:
                            if (p1 != (p1 / (((p1 | 706)) | 1))) {
                                v12 += (v12 & v12)
                            } else {
                                var v25 string = ("" + (("a" + "foo bar") + ("" + "\n\t")))
                                _ = v25
                            }
                        }
                        p0 = true
                    }
                    fmt.Println(p0, p1, p1)
                    v26 := []int{560, 42}
                    _ = v26
                }
            }
            for fi27 := 0; fi27 < 5; fi27++ {
                for fi28 := 0; fi28 < 12; fi28++ {
                    var v29 uint16 = (p2 ^ (p2 - 255))
                    v30 := GenS2{F0: (v12 >= (v12 & v12)), F1: 127, F2: 1}
                    var v31 int = 128
                    _ = fi28
                    _ = v29
                    _ = v30
                    _ = v31
                }
                _ = fi27
            }
            var v32 int8 = 0
            _ = v32
        }
        p0 = (!true)
    }
    v33 := GenS2{F0: p0, F1: 128, F2: 18446744073709551615}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v33
    return 2
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 string = ("\n\t" + ((("hello" + (((("\n\t" + "a") + (("a" + "hello") + ("\n\t" + "\n\t"))) + ("hello" + "hello")) + ((("a" + ("hello" + "hello")) + ("" + "a")) + ("hello" + ("foo bar" + "foo bar"))))) + "hello") + "foo bar"))
    v3 := []string{v2, (v2 + "hello"), v2}
    _ = len(v3)
    v2 = ""
    if len(v3) > 0 {
        v3[0] = ((v2 + (v2 + v2)) + ("a" + ("foo bar" + v2)))
    }
    var v34 string = v2
    var v35 bool = (true && (611 <= 32767))
    v34 = "\n\t"
    v3 = append(v3, (((v34 + "hello") + v2) + "a"))
    fmt.Println(v1)
    var v36 int64 = v1
    v36 = 1
    v36 -= 893
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
}
