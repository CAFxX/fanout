package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 uint; F2 int8 }
func genFn0() int32 {
    for fi7 := 0; fi7 < 11; fi7++ {
        for fi8 := 0; fi8 < 10; fi8++ {
            var v9 int32 = 1
            if (true && (0.0 > 1.5)) {
                var v10 uint64 = 0
                _ = v10
            } else {
                if (!(true || (4294967295 >= 32768))) {
                    v11 := GenS1{F0: "hello"}
                    var v12 int16 = 296
                    _ = v11
                    _ = v12
                } else {
                    v13 := []int{}
                    _ = v13
                }
            }
            var v14 uint32 = 276
            _ = fi8
            _ = v14
            _ = v9
        }
        _ = fi7
    }
    if (!(true && true)) {
        if true {
            v15 := GenS1{F0: "foo bar"}
            var v16 int32 = 1
            _ = v15
            _ = v16
        }
    } else {
        for fi17 := 0; fi17 < 8; fi17++ {
            v18 := GenS1{F0: ""}
            var v19 uint64 = 1
            v20 := GenS1{F0: ((((("" + "hello") + "a") + "\n\t") + "a") + "hello")}
            _ = fi17
            _ = v18
            _ = v19
            _ = v20
        }
    }
    return -2147483648
}
func genFn1(p0 int16, p1 int64, p2 uint8) uint8 {
    p0 = p0
    switch -1 {
    case 127:
        var v21 uint32 = 487
        _ = v21
    case 9223372036854775807:
        p2 += 128
    default:
        v22 := []int{}
        _ = v22
    }
    v23 := []int{32768, -641}
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    return p2
}

func main() {
    v1 := make(map[int]string)
    var v2 int32 = -2147483648
    v3 := []uint32{0, 727, 995, 0}
    var v4 int32 = -1
    var v5 int8 = 0
    v6 := []string{("" + (("\n\t" + "a") + ((("" + "") + ("" + ((((("" + (("hello" + "hello") + ("" + (("hello" + "") + "hello")))) + ("\n\t" + ("foo bar" + "hello"))) + ("a" + "")) + "\n\t") + "a"))) + "foo bar"))), ""}
    v24 := GenS2{F0: 9223372036854775807, F1: 42, F2: 1}
    switch ((("" + "\n\t") + "") + (((("foo bar" + "hello") + "") + "hello") + ("foo bar" + "a"))) {
    case "foo bar":
        var v25 int64 = -347
        _ = v25
    case "\n\t":
        for fi26 := 0; fi26 < 19; fi26++ {
            v2 = 0
            var v27 int32 = (v2 - genFn0())
            delete(v1, (-(fi26)))
            _ = fi26
            _ = v27
        }
    }
    var v28 uint32 = 404
    var v29 int16 = 42
    v4 -= (v4 >> ((0) & 63))
    var v30 uint64 = 0
    var v31 float64 = 1e10
    v5 -= (^(v5))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v4 += v2
    v32 := GenS1{F0: "a"}
    var v33 uint32 = (^(v28))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
