package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32; F1 uint16; F2 uint64 }
func genFn0(p0 string) string {
    p0 = p0
    switch p0 {
    case "hello":
        switch 1 {
        case 128:
            v5 := []int{}
            _ = v5
        default:
            for fi6 := 0; fi6 < 16; fi6++ {
                v7 := GenS1{F0: (false && ((4294967295 != 37) || (2 <= 4294967295)))}
                if true {
                    p0 = "hello"
                }
                var v8 string = p0
                _ = fi6
                _ = v7
                _ = v8
            }
        }
    case "\n\t":
        v9 := GenS1{F0: (105 <= 100)}
        _ = v9
    }
    var v10 float64 = 1e10
    if ((!true) || false) {
        var v11 uint8 = 201
        v11 += (v11 % (((v11 % ((v11) | 1))) | 1))
        for fi12 := 0; fi12 < 4; fi12++ {
            if (v10 <= v10) {
                if (86 < 1) {
                    fmt.Println(p0, v11)
                } else {
                    if false {
                        var v13 bool = (v11 >= (v11 ^ (^(v11))))
                        v10 += (v10 - (float64(3.14159) * (v10 + (v10 + ((3.14159 * (-2.25 - 1e10)) - v10)))))
                        _ = v13
                    } else {
                        v10 = 0.1
                    }
                    fmt.Println(v10)
                }
            } else {
                var v14 uint = 549
                v11 += 99
                _ = v14
            }
            v11 = 127
            switch -2147483648 {
            case 42:
                v11 = v11
            default:
                v15 := GenS1{F0: (false && (p0 == p0))}
                _ = v15
            }
            _ = fi12
        }
        _ = v11
    } else {
        v16 := GenS2{F0: 724, F1: -137}
        if (!(32767 <= 825)) {
            var v17 bool = (v10 <= (v10 / (v10 - (((v10 + float64(0.1)) + ((-2.25 - float64(0.0)) + v10)) - 1.5))))
            _ = v17
        } else {
            var v18 float64 = -2.25
            _ = v18
        }
        _ = v16
    }
    var v19 int32 = -2147483648
    _ = p0
    _ = v10
    _ = v19
    return p0
}
func genFn1(p0 int8, p1 uint) string {
    if false {
        v20 := GenS1{F0: true}
        p0 = p0
        _ = v20
    }
    var v21 uint = p1
    var v22 uint = (p1 & (p1 - v21))
    v23 := GenS3{F0: 0}
    for fi24 := 0; fi24 < 6; fi24++ {
        v25 := []int{}
        v22 -= p1
        if ((-2.25 - -2.25) >= -2.25) {
            switch ((("" + (("foo bar" + "a") + "a")) + (("foo bar" + "") + ("" + "a"))) + "") {
            case "":
                fmt.Println(v22)
            default:
                var v26 uint = v22
                _ = v26
            }
            v27 := GenS1{F0: ("a" <= (("" + "") + ("\n\t" + "\n\t")))}
            v21 = 42
            _ = v27
        }
        _ = fi24
        _ = v25
    }
    _ = p0
    _ = p1
    _ = v21
    _ = v22
    _ = v23
    return "hello"
}

func main() {
    v1 := []int{42}
    var v2 int32 = 0
    v3 := make(map[string]bool)
    v4 := []int{}
    var v28 uint = 2
    delete(v3, ("foo bar" + "a"))
    v28 += v28
    var v29 int8 = 100
    v30 := GenS4{F0: 968, F1: 65535, F2: 48}
    var v31 float32 = 0.1
    switch "foo bar" {
    case "":
        for fi32 := 0; fi32 < 5; fi32++ {
            fmt.Println(v28)
            var v33 int8 = (v29 * 42)
            _ = fi32
            _ = v33
        }
    case "a":
        v2 += (v2 << ((65535) & 63))
    default:
        v28 += (v28 - v28)
    }
    for fi34 := 0; fi34 < 18; fi34++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v35 bool = false
        _ = fi34
        _ = v35
    }
    v2 = v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
