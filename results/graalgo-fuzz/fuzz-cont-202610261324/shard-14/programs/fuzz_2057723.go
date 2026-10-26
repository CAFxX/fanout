package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string }
func genFn0() int {
    for fi6 := 0; fi6 < 17; fi6++ {
        var v7 int = fi6
        v7 -= -254
        if true {
            v7 += (v7 % (((-(fi6))) | 1))
            var v8 bool = false
            _ = v8
        }
        _ = fi6
        _ = v7
    }
    if (-322 == 42) {
        var v9 int = 127
        _ = v9
    }
    return -793
}

func main() {
    var v1 int16 = 658
    v2 := []int{}
    v3 := make(map[string]bool)
    v4 := []uint32{0, 185, 606, 2}
    var v5 uint8 = 87
    fmt.Println(v5)
    for fi10 := 0; fi10 < 8; fi10++ {
        if (!((^(fi10)) < fi10)) {
            for fi11 := 0; fi11 < 19; fi11++ {
                fmt.Println(v1)
                v5 = v5
                _ = fi11
            }
            _ = len(v4)
        }
        for fi12 := 0; fi12 < 5; fi12++ {
            delete(v3, "hello")
            if ((0.0 > 0.1) || true) {
                var v13 float32 = 3.14159
                _ = v13
            }
            _ = fi12
        }
        _ = fi10
    }
    var v14 float32 = 0.1
    v1 = (v1 * (v1 / (((v1 << ((0) & 63))) | 1)))
    if (v5 >= 207) {
        switch "a" {
        case "":
            var v15 uint8 = 1
            _ = v15
        case "a":
            if true {
                if false {
                    v14 = 0.0
                    delete(v3, "\n\t")
                }
                fmt.Println(v1, v1, v14)
                v16 := GenS1{F0: 0}
                _ = v16
            } else {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v14 = (0.0 + 0.0)
            }
        case "\n\t":
            v17 := GenS2{F0: ("foo bar" + "foo bar")}
            _ = v17
        default:
            var v18 string = ("a" + (("\n\t" + "\n\t") + "foo bar"))
            _ = v18
        }
    } else {
        v19 := GenS1{F0: v5}
        for fi20 := 0; fi20 < 15; fi20++ {
            var v21 bool = false
            _ = fi20
            _ = v21
        }
        _ = v19
    }
    v5 += 0
    var v22 uint = uint(v5)
    var v23 float32 = v14
    for fi24 := 0; fi24 < 6; fi24++ {
        v1 += 0
        v22 = v22
        for fi25 := 0; fi25 < 16; fi25++ {
            v14 = float32(((3.14159 / -2.25) / 1.5))
            var v26 float32 = (0.1 / v14)
            _ = fi25
            _ = v26
        }
        _ = fi24
    }
    fmt.Println(v14)
    switch genFn0() {
    case 32768:
        v14 += (0.0 / v14)
    default:
        if true {
            var v27 uint16 = 42
            _ = v27
        } else {
            if true {
                v23 = (v23 - float32(-2.25))
                v28 := GenS1{F0: (v5 >> ((497) & 63))}
                _ = v28
            }
            _ = len(v2)
            fmt.Println(v14, v22)
        }
    }
    fmt.Println(v5)
    for fi29 := 0; fi29 < 17; fi29++ {
        switch "hello" {
        case "":
            v30 := GenS1{F0: (v5 - v5)}
            _ = v30
        case "\n\t":
            var v31 uint64 = 0
            _ = v31
        case "foo bar":
            v32 := v3[((("foo bar" + ("a" + (("" + "") + ""))) + (("\n\t" + ("hello" + "foo bar")) + "a")) + "\n\t")]
            _ = v32
        default:
            v5 = (v5 - 1)
        }
        var v33 uint32 = 42
        var v34 int32 = 42
        _ = fi29
        _ = v33
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
