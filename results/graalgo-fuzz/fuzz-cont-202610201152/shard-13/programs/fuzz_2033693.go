package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 uint64; F1 uint8 }
func genFn0(p0 int8, p1 uint32, p2 string) uint {
    if false {
        var v7 int64 = int64(p0)
        if true {
            var v8 string = "a"
            v9 := []int{}
            _ = v8
            _ = v9
        } else {
            if false {
                switch p2 {
                case "foo bar":
                    if true {
                        switch ("" + p2) {
                        case "hello":
                            var v10 int64 = 448
                            _ = v10
                        default:
                            p1 = 2
                        }
                    }
                case "":
                    var v11 float64 = (0.1 * -2.25)
                    _ = v11
                }
            } else {
                v12 := []int{}
                _ = v12
            }
            if (2 != 65535) {
                p0 = p0
                if false {
                    v13 := GenS1{F0: (p0 & (p0 + -96)), F1: 0}
                    _ = v13
                }
                var v14 int = 1
                _ = v14
            } else {
                var v15 float64 = -2.25
                _ = v15
            }
        }
        var v16 bool = ((true || true) || (2 > 42))
        _ = v16
        _ = v7
    }
    for fi17 := 0; fi17 < 16; fi17++ {
        fmt.Println(p0)
        _ = fi17
    }
    _ = p0
    _ = p1
    _ = p2
    return 18446744073709551615
}

func main() {
    v1 := []string{("" + "")}
    v2 := []int{40, 32767, 0, -9223372036854775808}
    var v3 uint64 = 2
    var v4 int32 = -125
    var v5 int64 = -1
    var v6 int16 = 32767
    var v18 int = -2147483648
    v19 := GenS2{F0: (^(v3)), F1: 127}
    v1 = append(v1, ((("a" + "\n\t") + "") + ("\n\t" + "hello")))
    var v20 bool = true
    var v21 int64 = v5
    v4 = v4
    fmt.Println(v4, v4)
    var v22 float32 = (0.0 * (((3.14159 + 3.14159) - (((((0.1 - 0.0) + (0.1 + 1.5)) + 1.5) + ((float32(-2.25) * (0.1 + 1.5)) + 1.5)) * -2.25)) - -2.25))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
