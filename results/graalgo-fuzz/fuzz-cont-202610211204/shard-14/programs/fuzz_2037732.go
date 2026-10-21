package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 uint64, p1 int16, p2 bool) int32 {
    v6 := GenS1{F0: 856, F1: ("" + "\n\t"), F2: (p0 & p0)}
    if p2 {
        var v7 uint16 = 42
        var v8 uint8 = 165
        _ = v7
        _ = v8
    }
    p2 = p2
    p2 = p2
    p0 -= p0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -837
}

func main() {
    v1 := make(map[int]bool)
    var v2 string = ((("" + ((("" + "foo bar") + "foo bar") + "\n\t")) + "hello") + "\n\t")
    var v3 float32 = 1.5
    var v4 int32 = 523
    v5 := []uint32{65535, 2147483647}
    v3 -= (v3 - v3)
    var v9 int = 128
    for fi10 := 0; fi10 < 4; fi10++ {
        var v11 uint64 = 0
        switch -715 {
        case 42:
            for fi12 := 0; fi12 < 17; fi12++ {
                v13 := GenS1{F0: 0, F1: "foo bar", F2: v11}
                fmt.Println(fi12)
                _ = fi12
                _ = v13
            }
        case -9223372036854775808:
            if (!(false && ((71 < -50) || false))) {
                v11 = (v11 - v11)
                fmt.Println(v9)
                var v14 string = v2
                _ = v14
            } else {
                var v15 float32 = 0.1
                _ = v15
            }
        }
        _ = fi10
        _ = v11
    }
    switch (("hello" + (v2 + (v2 + "a"))) + ("a" + v2)) {
    case "a":
        switch v9 {
        case -266:
            var v16 string = ("foo bar" + "hello")
            _ = v16
        }
    case "hello":
        v17 := GenS1{F0: 65535, F1: (v2 + (v2 + v2)), F2: 1}
        _ = v17
    case "":
        if ((!false) || (!(0 > 42))) {
            fmt.Println(v2)
            for fi18 := 0; fi18 < 7; fi18++ {
                if (2147483647 <= 65535) {
                    var v19 float32 = v3
                    var v20 uint = 283
                    for fi21 := 0; fi21 < 16; fi21++ {
                        v22 := GenS1{F0: 255, F1: "", F2: 4294967295}
                        v19 = v3
                        var v23 string = (v2 + ((("hello" + v2) + "a") + "a"))
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                    _ = v19
                    _ = v20
                }
                fmt.Println(fi18)
                _ = fi18
            }
        } else {
            var v24 int = v9
            v25 := GenS2{F0: 18446744073709551615}
            _ = v24
            _ = v25
        }
    }
    for fi26 := 0; fi26 < 11; fi26++ {
        v27 := GenS2{F0: 42}
        _ = fi26
        _ = v27
    }
    v3 = (v3 - (((((v3 / (0.0 - v3)) * 0.0) * v3) + -2.25) + v3))
    switch "hello" {
    case "a":
        for fi28 := 0; fi28 < 12; fi28++ {
            v4 += (v4 % ((v4) | 1))
            var v29 int8 = -85
            _ = fi28
            _ = v29
        }
    case "":
        v30 := GenS1{F0: 0, F1: "hello", F2: 2}
        _ = v30
    case "\n\t":
        v4 = (v4 * genFn0(2, -1, (true || (!(false && true)))))
    default:
        for fi31 := 0; fi31 < 1; fi31++ {
            v32 := GenS3{F0: 758}
            v4 += 2147483647
            _ = fi31
            _ = v32
        }
    }
    var v33 uint = 932
    v4 = -1
    for fi34 := 0; fi34 < 12; fi34++ {
        switch v9 {
        case -1:
            var v35 uint8 = 1
            _ = v35
        case 614:
            v36 := v1[(v9 >> (((v33 / (((v33 & (v33 * v33))) | 1))) & 63))]
            _ = v36
        case -691:
            v2 = v2
        }
        switch v9 {
        case -1:
            var v37 int16 = 1000
            _ = v37
        case 2147483647:
            var v38 int8 = 0
            _ = v38
        case 42:
            var v39 uint8 = 1
            _ = v39
        }
        _ = fi34
    }
    delete(v1, (v9 << ((v33) & 63)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v9
}
