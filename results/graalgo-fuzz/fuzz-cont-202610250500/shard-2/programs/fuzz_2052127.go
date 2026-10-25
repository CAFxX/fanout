package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 string; F2 bool }
func genFn0(p0 int32, p1 uint32, p2 int16) string {
    v5 := GenS1{F0: 591}
    var v6 string = ("foo bar" + ("hello" + "a"))
    var v7 bool = false
    p1 = (p1 * (p1 / ((p1) | 1)))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return ("hello" + ((v6 + (v6 + "")) + ((v6 + v6) + (((("foo bar" + "\n\t") + ("hello" + "")) + ("hello" + v6)) + (v6 + v6)))))
}

func main() {
    var v1 int32 = 394
    v2 := make(map[int]bool)
    var v3 uint8 = 127
    var v4 int32 = (v1 + -1)
    v8 := v2[19]
    v3 = (v3 / ((v3) | 1))
    v1 -= (v1 | v4)
    var v9 uint64 = 929
    v10 := GenS1{F0: 42}
    var v11 uint32 = 1
    v12 := GenS2{F0: 0, F1: "hello", F2: v8}
    var v13 int64 = 0
    fmt.Println(v11, v11, v3)
    for fi14 := 0; fi14 < 9; fi14++ {
        v15 := []int{(fi14 & (^(fi14))), (fi14 - 9223372036854775807), -9223372036854775808}
        for fi16 := 0; fi16 < 3; fi16++ {
            delete(v2, (fi16 ^ fi16))
            v15 = append(v15, (-(fi16)))
            if v8 {
                if len(v15) > 0 {
                    v15[0] = (fi16 ^ (fi14 / ((fi14) | 1)))
                }
                fmt.Println(fi14, v8)
                v17 := GenS2{F0: (v13 & v13), F1: "", F2: (!((!(((v11 >> ((31) & 63)) <= (^(v11))) || true)) || false))}
                _ = v17
            } else {
                v8 = v8
                v15 = append(v15, fi14)
            }
            v11 += 4294967295
            fmt.Println(fi14, fi14)
            _ = fi16
        }
        v3 = (^(v3))
        _ = fi14
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
