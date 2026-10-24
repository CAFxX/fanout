package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
func genFn0(p0 uint, p1 int64) int16 {
    fmt.Println(p0, p1)
    if (!true) {
        v6 := []int{2147483647, -1}
        for fi7 := 0; fi7 < 6; fi7++ {
            fmt.Println(p1, p0, p0)
            _ = fi7
        }
        p1 += (p1 % (((p1 / ((p1) | 1))) | 1))
        _ = v6
    }
    v8 := []int{0, -9223372036854775808, 0, 127}
    _ = p0
    _ = p1
    _ = v8
    return -248
}

func main() {
    var v1 int8 = 1
    var v2 uint32 = 424
    var v3 int = -9223372036854775808
    v4 := []int{(v3 | v3), v3, 127, (v3 << ((65535) & 63))}
    v5 := []string{"a", "a"}
    v2 -= (^(v2))
    if true {
        v9 := GenS1{F0: 42, F1: false}
        for fi10 := 0; fi10 < 20; fi10++ {
            v1 = v1
            _ = fi10
        }
        _ = v9
    } else {
        var v11 uint16 = 1
        for fi12 := 0; fi12 < 13; fi12++ {
            var v13 uint64 = 4294967295
            _ = fi12
            _ = v13
        }
        v14 := GenS1{F0: 9223372036854775807, F1: (((false && (!true)) && (42 <= 128)) || false)}
        _ = v11
        _ = v14
    }
    var v15 uint64 = 0
    v15 = 4294967295
    if true {
        v2 -= v2
    }
    if true {
        var v16 int = (v3 * (-(v3)))
        v17 := GenS1{F0: -1, F1: false}
        if len(v5) > 0 {
            v5[0] = ("hello" + "foo bar")
        }
        _ = v16
        _ = v17
    }
    v3 -= 2147483647
    if ("hello" <= "\n\t") {
        v4 = append(v4, (v3 | (-(v3))))
        for fi18 := 0; fi18 < 2; fi18++ {
            var v19 bool = (v1 >= -84)
            for fi20 := 0; fi20 < 15; fi20++ {
                v21 := GenS1{F0: 42, F1: v19}
                var v22 int32 = -1
                var v23 uint = 2
                _ = fi20
                _ = v21
                _ = v22
                _ = v23
            }
            _ = fi18
            _ = v19
        }
        for fi24 := 0; fi24 < 14; fi24++ {
            v1 += (v1 % ((69) | 1))
            v1 -= -127
            v25 := GenS1{F0: 9223372036854775807, F1: true}
            _ = fi24
            _ = v25
        }
    } else {
        v5 = append(v5, (((("hello" + (((("foo bar" + "foo bar") + "") + "hello") + ("\n\t" + "\n\t"))) + ((("foo bar" + ("hello" + (((("a" + "") + "\n\t") + ("foo bar" + (((("hello" + "\n\t") + "\n\t") + ("hello" + "hello")) + "a"))) + "\n\t"))) + "a") + ("\n\t" + "foo bar"))) + "a") + (("foo bar" + ("hello" + ("\n\t" + ("\n\t" + ("a" + ("foo bar" + "")))))) + "hello")))
    }
    for fi26 := 0; fi26 < 13; fi26++ {
        var v27 uint8 = 42
        if ("" == ("hello" + "a")) {
            if (!(1 >= v27)) {
                v28 := GenS1{F0: 9223372036854775807, F1: (true && ("" <= ("hello" + ("hello" + "foo bar"))))}
                _ = v28
            }
            var v29 int16 = 42
            _ = v29
        }
        switch fi26 {
        case -1:
            var v30 uint16 = 465
            _ = v30
        case 698:
            if (false && ("hello" == "foo bar")) {
                v31 := GenS1{F0: -9223372036854775808, F1: false}
                if (86 <= int64(v2)) {
                    for fi32 := 0; fi32 < 17; fi32++ {
                        v27 = 0
                        v3 = 127
                        v1 = (^(v1))
                        _ = fi32
                    }
                }
                _ = v31
            }
        default:
            var v33 int16 = -711
            _ = v33
        }
        _ = fi26
        _ = v27
    }
    for fi34 := 0; fi34 < 16; fi34++ {
        _ = len(v5)
        if false {
            _ = len(v5)
            if true {
                switch ("hello" + "hello") {
                case "\n\t":
                    v15 = 18446744073709551615
                }
                if (1 == 255) {
                    for fi35 := 0; fi35 < 3; fi35++ {
                        v1 -= (v1 + 100)
                        _ = fi35
                    }
                    var v36 bool = false
                    _ = v36
                }
            } else {
                v2 = v2
            }
            v3 = v3
        }
        switch -65 {
        case 32768:
            v3 = (^(fi34))
        case 128:
            fmt.Println(fi34)
        case 0:
            if false {
                v37 := GenS1{F0: -545, F1: (65535 != (v2 - v2))}
                v2 = v2
                _ = v37
            }
        default:
            if true {
                v15 += (v15 << ((462) & 63))
                var v38 bool = ((216 < 128) && (!false))
                _ = v38
            }
        }
        _ = fi34
    }
    var v39 uint16 = 32767
    var v40 uint8 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
