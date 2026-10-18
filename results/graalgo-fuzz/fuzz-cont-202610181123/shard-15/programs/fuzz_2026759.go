package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 int16; F2 uint8 }
type GenS4 struct { F0 uint8; F1 string; F2 uint64 }
func genFn0(p0 uint32, p1 uint8) uint16 {
    var v6 int32 = 0
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 uint8 = 42
        p1 -= p1
        p1 -= (^(p1))
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 15; fi9++ {
        v10 := GenS1{F0: 32767}
        var v11 int16 = -1
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v6
    return 255
}

func main() {
    var v1 int32 = -1
    var v2 float32 = -2.25
    var v3 uint8 = 255
    v4 := make(map[int]string)
    v5 := make(map[string]int)
    if true {
        v12 := GenS1{F0: 65535}
        v3 -= v3
        _ = v12
    } else {
        if false {
            if true {
                v13 := []int{651}
                v14 := GenS1{F0: 32767}
                _ = v13
                _ = v14
            }
            switch "a" {
            case "a":
                switch 32768 {
                case 128:
                    if false {
                        var v15 uint = 255
                        _ = v15
                    }
                default:
                    var v16 float32 = v2
                    _ = v16
                }
            default:
                v1 = v1
            }
        }
        var v17 int16 = 0
        _ = v17
    }
    var v18 uint64 = 4294967295
    v19 := GenS2{F0: -936}
    v3 = (v3 * (^(v3)))
    v20 := GenS3{F0: 672, F1: -1, F2: v3}
    switch (("hello" + ("\n\t" + "foo bar")) + "a") {
    case "\n\t":
        v21 := []int{-9223372036854775808, 918, -503, 1}
        _ = v21
    case "a":
        v2 = (v2 - v2)
    }
    v22 := GenS3{F0: 255, F1: 652, F2: (v3 * v3)}
    if (true || true) {
        if ((-2.25 + 3.14159) > 3.14159) {
            var v23 uint32 = 1
            v3 = v3
            v24 := GenS4{F0: v3, F1: "\n\t", F2: v18}
            _ = v23
            _ = v24
        } else {
            var v25 float32 = (v2 + (v2 - (1e10 - 1.5)))
            _ = v25
        }
        if true {
            if ((!(v2 <= (v2 + v2))) || false) {
                v1 = (v1 + 42)
            }
            v18 = 2
            delete(v5, (("\n\t" + "") + ((("\n\t" + ("foo bar" + (("\n\t" + (("\n\t" + "") + ("a" + "foo bar"))) + "a"))) + "\n\t") + ((((("\n\t" + "a") + "\n\t") + "foo bar") + "hello") + ("hello" + "")))))
        }
    }
    v3 += (v3 << ((65535) & 63))
    v26 := GenS1{F0: genFn0(2147483648, (^(v3)))}
    delete(v4, -627)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
