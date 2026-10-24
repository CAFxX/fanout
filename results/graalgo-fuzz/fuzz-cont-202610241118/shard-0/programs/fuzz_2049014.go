package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint }
func genFn0(p0 uint32) int64 {
    if false {
        switch ("foo bar" + "foo bar") {
        case "foo bar":
            var v6 uint16 = 32767
            _ = v6
        }
        v7 := []int{32767, 9223372036854775807, 0}
        _ = v7
    } else {
        for fi8 := 0; fi8 < 6; fi8++ {
            var v9 int32 = -1
            var v10 float32 = 1e10
            fmt.Println(v9)
            _ = fi8
            _ = v10
            _ = v9
        }
        p0 -= p0
        for fi11 := 0; fi11 < 13; fi11++ {
            for fi12 := 0; fi12 < 2; fi12++ {
                p0 = 114
                _ = fi12
            }
            var v13 bool = (!true)
            p0 -= 42
            _ = fi11
            _ = v13
        }
    }
    p0 -= 2147483647
    if (!(!(1e10 > ((0.1 * 1e10) * float32(3.14159))))) {
        p0 = p0
        fmt.Println(p0)
        if true {
            p0 += (p0 << ((255) & 63))
            fmt.Println(p0)
            fmt.Println(p0)
        }
    }
    _ = p0
    return -89
}

func main() {
    var v1 int = 127
    var v2 string = ("a" + ("foo bar" + "foo bar"))
    v3 := make(map[uint32]int)
    var v4 uint8 = 12
    v5 := []int{(-(v1)), (v1 % ((v1) | 1))}
    var v14 uint8 = 0
    if true {
        var v15 int = (v1 % ((v1) | 1))
        if (968 != 0) {
            v15 = (v15 - (v15 << ((4294967295) & 63)))
        } else {
            if (965 == 176) {
                var v16 float64 = -2.25
                _ = v16
            } else {
                fmt.Println(v4, v1, v2)
            }
        }
        _ = v15
    }
    v5 = append(v5, v1)
    var v17 float32 = 1.5
    var v18 int8 = -128
    for fi19 := 0; fi19 < 11; fi19++ {
        switch ("\n\t" + v2) {
        case "a":
            var v20 float64 = -2.25
            _ = v20
        case "":
            v21 := GenS1{F0: 1}
            _ = v21
        default:
            var v22 uint16 = 65535
            _ = v22
        }
        _ = fi19
    }
    v2 = (v2 + v2)
    var v23 int32 = 1
    if true {
        v2 = v2
        var v24 float64 = 1e10
        if true {
            var v25 int = (v1 & 0)
            v2 = "hello"
            if true {
                v25 = (v1 & v1)
            }
            _ = v25
        } else {
            for fi26 := 0; fi26 < 8; fi26++ {
                if (32767 != -950) {
                    v14 += v14
                }
                if true {
                    var v27 float32 = v17
                    switch "foo bar" {
                    case "foo bar":
                        v28 := GenS1{F0: 1}
                        _ = v28
                    case "hello":
                        v17 = v27
                    case "\n\t":
                        v24 += v24
                    default:
                        v29 := GenS1{F0: 1}
                        _ = v29
                    }
                    var v30 uint16 = 255
                    _ = v27
                    _ = v30
                }
                var v31 uint = 42
                _ = fi26
                _ = v31
            }
            if true {
                for fi32 := 0; fi32 < 4; fi32++ {
                    delete(v3, 4294967295)
                    if len(v5) > 0 {
                        v5[0] = 32768
                    }
                    _ = len(v5)
                    _ = fi32
                }
                v17 = 0.1
            } else {
                v33 := GenS2{F0: 1}
                for fi34 := 0; fi34 < 2; fi34++ {
                    if (!true) {
                        v3[4294967295] = 679
                        var v35 int8 = -1
                        _ = v35
                    }
                    _ = fi34
                }
                if true {
                    v24 = v24
                }
                _ = v33
            }
        }
        _ = v24
    }
    for fi36 := 0; fi36 < 8; fi36++ {
        var v37 float32 = float32((((1.5 + (0.1 + (1e10 * (1.5 + float64((v17 * 1.5)))))) * (0.0 * 0.1)) + -2.25))
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi36
        _ = v37
    }
    var v38 uint = 877
    var v39 int8 = 1
    var v40 uint16 = 55
    v41 := GenS2{F0: 444}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
