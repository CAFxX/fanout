package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 bool }
type GenS2 struct { F0 uint16; F1 bool }
func genFn0(p0 string) uint32 {
    var v6 int16 = -1
    switch p0 {
    case "\n\t":
        p0 = ""
    case "foo bar":
        var v7 int32 = 2147483647
        _ = v7
    case "":
        var v8 uint = 65535
        _ = v8
    default:
        for fi9 := 0; fi9 < 3; fi9++ {
            fmt.Println(p0)
            p0 = p0
            v10 := []int{(fi9 & 4294967295), (fi9 >> ((42) & 63))}
            _ = fi9
            _ = v10
        }
    }
    _ = p0
    _ = v6
    return 752
}

func main() {
    v1 := []int{931, -2147483648, 1, -2147483648}
    var v2 int8 = 67
    v3 := []uint32{4294967295, 42, 4294967295, 2}
    v4 := []string{(("foo bar" + ("\n\t" + "\n\t")) + "\n\t"), "", (("foo bar" + (("hello" + "hello") + "")) + ("a" + "\n\t"))}
    var v5 int8 = v2
    var v11 int8 = 1
    switch 42 {
    case 1:
        if false {
            v3 = append(v3, 128)
            v11 += 73
            for fi12 := 0; fi12 < 12; fi12++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi12
            }
        } else {
            v3 = append(v3, 42)
            for fi13 := 0; fi13 < 14; fi13++ {
                v11 += 100
                v14 := GenS1{F0: 2147483647, F1: 42, F2: (1 <= 255)}
                v11 = (v5 % (((v2 >> ((447) & 63))) | 1))
                _ = fi13
                _ = v14
            }
            v15 := GenS1{F0: 1, F1: (v2 * 1), F2: (!false)}
            _ = v15
        }
    case 9223372036854775807:
        for fi16 := 0; fi16 < 3; fi16++ {
            if (false && true) {
                v2 = v5
            } else {
                _ = len(v1)
                fmt.Println(v2, v11, fi16)
                _ = len(v4)
            }
            if (!true) {
                v17 := GenS1{F0: -756, F1: 42, F2: true}
                fmt.Println(v2)
                v2 += 100
                _ = v17
            } else {
                var v18 int32 = 367
                v18 -= 2147483647
                _ = v18
            }
            if false {
                for fi19 := 0; fi19 < 14; fi19++ {
                    if len(v4) > 0 {
                        v4[0] = "a"
                    }
                    switch "foo bar" {
                    case "a":
                        v5 += (v11 >> ((176) & 63))
                    default:
                        v11 = (v2 * (v2 + (v11 >> ((3) & 63))))
                    }
                    _ = fi19
                }
                v20 := GenS1{F0: 1, F1: v11, F2: false}
                v2 = 1
                _ = v20
            } else {
                for fi21 := 0; fi21 < 7; fi21++ {
                    if true {
                        fmt.Println(v11, fi16, v2)
                        if (false && (-1 > 0)) {
                            v5 = 127
                            v3 = append(v3, 65535)
                            _ = len(v4)
                        } else {
                            if true {
                                var v22 uint = 42
                                _ = v22
                            }
                        }
                    }
                    var v23 string = ("" + "hello")
                    if (((v11 ^ v5) != (v5 & v2)) && (false && true)) {
                        v24 := GenS2{F0: 819, F1: (243 >= 2)}
                        var v25 int = -2147483648
                        v26 := GenS2{F0: 532, F1: (745 > 32767)}
                        _ = v24
                        _ = v25
                        _ = v26
                    } else {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    _ = fi21
                    _ = v23
                }
                var v27 int32 = -2147483648
                var v28 uint32 = 2
                _ = v27
                _ = v28
            }
            _ = fi16
        }
    }
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    v11 += 127
    v11 = (-(v5))
    fmt.Println(v2, v11, v11)
    if ((!true) || true) {
        v11 += (v5 % ((v11) | 1))
    }
    var v29 float64 = 1.5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
