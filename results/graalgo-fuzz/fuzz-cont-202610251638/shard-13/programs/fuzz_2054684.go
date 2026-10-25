package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 int8 }
func genFn0(p0 uint16, p1 int8) uint32 {
    p0 = 255
    for fi5 := 0; fi5 < 5; fi5++ {
        p0 += 773
        v6 := []int{0, -1, (fi5 >> ((255) & 63))}
        for fi7 := 0; fi7 < 3; fi7++ {
            switch "foo bar" {
            case "\n\t":
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            case "":
                p1 -= -44
                var v8 int = (fi5 * -9223372036854775808)
                var v9 uint32 = 4294967295
                _ = v8
                _ = v9
            case "foo bar":
                p0 -= (p0 - (p0 << ((65535) & 63)))
            default:
                switch "\n\t" {
                case "hello":
                    fmt.Println(fi7, p1)
                case "":
                    switch fi5 {
                    case 708:
                        p0 = (p0 << ((255) & 63))
                        fmt.Println(fi7, fi5)
                        p1 -= p1
                    case 32767:
                        p0 = p0
                    case -9223372036854775808:
                        fmt.Println(fi7)
                    }
                }
            }
            fmt.Println(p1, p1, fi7)
            p1 += (p1 & p1)
            _ = fi7
        }
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    return 2
}
func genFn1(p0 uint8, p1 uint16) bool {
    switch 2147483647 {
    case 1:
        fmt.Println(p0)
    case -2147483648:
        v10 := []int{2147483647, -1, 42}
        _ = v10
    case 42:
        var v11 int64 = 831
        _ = v11
    }
    if false {
        var v12 uint = 42
        _ = v12
    } else {
        for fi13 := 0; fi13 < 20; fi13++ {
            fmt.Println(p1, fi13)
            var v14 float32 = 0.0
            _ = fi13
            _ = v14
        }
        if true {
            for fi15 := 0; fi15 < 15; fi15++ {
                if (int16(p0) > 42) {
                    var v16 int64 = 672
                    if (-2147483648 != 506) {
                        v17 := []int{}
                        v18 := GenS1{F0: 1}
                        _ = v17
                        _ = v18
                    } else {
                        var v19 bool = false
                        v20 := GenS2{F0: 18446744073709551615, F1: v19}
                        var v21 int = (fi15 / ((32768) | 1))
                        _ = v19
                        _ = v20
                        _ = v21
                    }
                    v22 := GenS3{F0: 100}
                    _ = v16
                    _ = v22
                }
                v23 := GenS1{F0: -32}
                _ = fi15
                _ = v23
            }
        } else {
            if false {
                fmt.Println(p0, p1)
                p1 = 65535
            } else {
                var v24 uint = 502
                switch 2147483647 {
                case 1:
                    v25 := []int{32768, -2147483648, 1, 127}
                    _ = v25
                }
                _ = v24
            }
            var v26 string = "a"
            var v27 uint16 = p1
            _ = v26
            _ = v27
        }
    }
    switch 32768 {
    case -2147483648:
        for fi28 := 0; fi28 < 13; fi28++ {
            var v29 int64 = -143
            var v30 uint64 = 42
            var v31 uint32 = 514
            _ = fi28
            _ = v29
            _ = v30
            _ = v31
        }
    case 1000:
        if false {
            p0 += (p0 ^ (p0 * (p0 + p0)))
        }
    default:
        for fi32 := 0; fi32 < 2; fi32++ {
            fmt.Println(fi32, fi32)
            _ = fi32
        }
    }
    p0 -= (p0 & (p0 ^ (p0 % ((p0) | 1))))
    for fi33 := 0; fi33 < 17; fi33++ {
        var v34 string = (((("a" + "hello") + "") + (("a" + "hello") + ("hello" + ("a" + "hello")))) + "\n\t")
        var v35 bool = (!true)
        fmt.Println(v34, p0, p1)
        _ = fi33
        _ = v34
        _ = v35
    }
    _ = p0
    _ = p1
    return true
}

func main() {
    var v1 int16 = -944
    var v2 int8 = 0
    v3 := []uint32{4294967295}
    v4 := []string{}
    v2 += (v2 >> ((4294967295) & 63))
    switch "a" {
    case "foo bar":
        fmt.Println(v1)
    }
    fmt.Println(v2)
    for fi36 := 0; fi36 < 3; fi36++ {
        v1 = v1
        _ = fi36
    }
    v37 := GenS2{F0: 65535, F1: (false || (true || (false || false)))}
    _ = len(v4)
    if len(v3) > 0 {
        v3[0] = 42
    }
    if ("hello" != ("hello" + ("a" + "foo bar"))) {
        var v38 int16 = (v1 >> ((599) & 63))
        fmt.Println(v2)
        var v39 bool = (969 <= 763)
        _ = v38
        _ = v39
    } else {
        v2 -= 100
        if (!(!(923 == 261))) {
            if false {
                var v40 int8 = 42
                _ = v40
            } else {
                if true {
                    var v41 uint = 1
                    _ = v41
                }
            }
            v2 = v2
        }
        if len(v4) > 0 {
            v4[0] = "foo bar"
        }
    }
    v2 = 127
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
}
