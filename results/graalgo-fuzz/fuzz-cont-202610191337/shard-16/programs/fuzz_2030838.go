package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 string; F2 int32 }
func genFn0(p0 uint16, p1 uint8) string {
    p0 -= p0
    for fi5 := 0; fi5 < 6; fi5++ {
        switch "" {
        case "\n\t":
            p0 = p0
        case "a":
            if (!(!((-2147483648 < 599) && true))) {
                var v6 uint16 = 916
                var v7 float64 = 1.5
                _ = v6
                _ = v7
            }
        default:
            p1 = (p1 / ((p1) | 1))
        }
        if false {
            switch "hello" {
            case "a":
                if false {
                    p1 -= (p1 / ((p1) | 1))
                }
            case "hello":
                var v8 int32 = 50
                _ = v8
            case "":
                for fi9 := 0; fi9 < 11; fi9++ {
                    p1 = p1
                    _ = fi9
                }
            default:
                var v10 int8 = 127
                _ = v10
            }
            if false {
                p0 -= (p0 << ((42) & 63))
                v11 := GenS1{F0: 1000, F1: "hello"}
                for fi12 := 0; fi12 < 6; fi12++ {
                    if (65535 != 2) {
                        var v13 int8 = -1
                        var v14 int16 = 1
                        p0 -= 319
                        _ = v13
                        _ = v14
                    }
                    p1 = 2
                    p1 += 2
                    _ = fi12
                }
                _ = v11
            }
        } else {
            fmt.Println(p0, p1)
        }
        v15 := []int{(fi5 % ((1) | 1))}
        _ = fi5
        _ = v15
    }
    switch -225 {
    case 128:
        p0 -= p0
    case 1:
        var v16 int8 = -128
        _ = v16
    }
    _ = p0
    _ = p1
    return (("hello" + "a") + "\n\t")
}

func main() {
    var v1 int64 = 0
    var v2 uint64 = 42
    var v3 int8 = 1
    var v4 uint16 = 65535
    for fi17 := 0; fi17 < 13; fi17++ {
        v18 := GenS1{F0: 32767, F1: "hello"}
        _ = fi17
        _ = v18
    }
    if (true && (true && (v4 > (v4 / ((v4) | 1))))) {
        v4 = 255
        v19 := GenS2{F0: 65535}
        _ = v19
    }
    for fi20 := 0; fi20 < 5; fi20++ {
        for fi21 := 0; fi21 < 15; fi21++ {
            v22 := []int{128, fi20}
            var v23 string = ("hello" + ("a" + "\n\t"))
            _ = fi21
            _ = v22
            _ = v23
        }
        if (4294967295 < (v2 * (v2 << ((65535) & 63)))) {
            if false {
                var v24 int32 = 2147483647
                v1 -= -9223372036854775808
                for fi25 := 0; fi25 < 3; fi25++ {
                    v26 := []int{-695, fi25}
                    _ = fi25
                    _ = v26
                }
                _ = v24
            } else {
                var v27 float32 = 0.1
                fmt.Println(v4)
                v28 := GenS1{F0: -32768, F1: "\n\t"}
                _ = v27
                _ = v28
            }
            for fi29 := 0; fi29 < 14; fi29++ {
                v2 = (v2 * v2)
                v30 := []int{4294967295}
                var v31 float64 = 0.1
                _ = fi29
                _ = v30
                _ = v31
            }
            if false {
                v3 = (v3 % ((42) | 1))
                fmt.Println(v2)
            }
        } else {
            v2 -= 0
            var v32 uint16 = v4
            fmt.Println(fi20, fi20)
            _ = v32
        }
        _ = fi20
    }
    v33 := GenS3{F0: "", F1: ("hello" + "\n\t"), F2: 0}
    var v34 int = 766
    fmt.Println(v3, v4, v3)
    switch v34 {
    case 999:
        v2 = (^(v2))
    default:
        for fi35 := 0; fi35 < 13; fi35++ {
            var v36 float32 = -2.25
            _ = fi35
            _ = v36
        }
    }
    v2 = v2
    if false {
        var v37 int8 = (-(v3))
        _ = v37
    }
    v38 := GenS3{F0: "foo bar", F1: "", F2: -728}
    switch (v34 >> ((65535) & 63)) {
    case 2147483647:
        v39 := []int{v34, v34, 4294967295, (v34 >> ((65535) & 63))}
        _ = v39
    case -737:
        var v40 string = "hello"
        _ = v40
    }
    v4 += (^(v4))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v38
    _ = v4
}
