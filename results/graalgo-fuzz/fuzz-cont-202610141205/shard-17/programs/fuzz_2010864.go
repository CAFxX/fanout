package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 int64 }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 bool }
func genFn0(p0 uint, p1 int64) uint32 {
    v7 := GenS1{F0: 42, F1: 42, F2: p1}
    if true {
        if false {
            v8 := GenS2{F0: (p1 << ((65535) & 63)), F1: 146}
            v9 := []int{32767, 319, 1, 2147483647}
            _ = v8
            _ = v9
        } else {
            var v10 int8 = 100
            for fi11 := 0; fi11 < 2; fi11++ {
                for fi12 := 0; fi12 < 16; fi12++ {
                    fmt.Println(fi12, p1)
                    switch "" {
                    case "\n\t":
                        v13 := []int{(^(fi12))}
                        _ = v13
                    case "hello":
                        var v14 int32 = 1
                        _ = v14
                    default:
                        p0 = 2
                    }
                    var v15 uint = p0
                    _ = fi12
                    _ = v15
                }
                p1 -= 659
                _ = fi11
            }
            _ = v10
        }
        var v16 float32 = ((3.14159 * 3.14159) + (0.1 * float32(1e10)))
        _ = v16
    }
    _ = p0
    _ = p1
    _ = v7
    return 42
}

func main() {
    v1 := []string{(("a" + "") + (("\n\t" + "foo bar") + "foo bar")), (((("hello" + "hello") + "foo bar") + (((("\n\t" + "a") + ("hello" + "")) + "hello") + (("foo bar" + "\n\t") + (((("a" + "") + "\n\t") + "") + "a")))) + ("a" + "a"))}
    var v2 uint16 = 758
    v3 := []string{("hello" + ("foo bar" + "\n\t"))}
    var v4 int64 = -515
    var v5 uint16 = (v2 << ((uint(v2)) & 63))
    v6 := []string{"\n\t", ("" + ("\n\t" + ("foo bar" + "a"))), ("" + "a"), ""}
    var v17 string = "a"
    var v18 uint16 = (^(v2))
    if ((!true) && false) {
        var v19 uint32 = 0
        _ = v19
    }
    fmt.Println(v5, v18)
    var v20 int16 = -1
    v21 := GenS3{F0: true}
    var v22 int16 = (v20 + -943)
    if true {
        fmt.Println(v18, v18)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
