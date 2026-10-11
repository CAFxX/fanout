package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 uint8, p1 uint64) uint16 {
    v7 := GenS1{F0: -406, F1: 255}
    if false {
        p1 -= p1
        v8 := []int{507, 32767, 42, 32767}
        var v9 uint = 124
        _ = v8
        _ = v9
    } else {
        v10 := []int{2147483647, 4294967295, 128, 128}
        _ = len(v10)
        if ((0.0 - ((-2.25 + 1e10) + (float32((3.14159 + 1e10)) * 0.1))) == 3.14159) {
            for fi11 := 0; fi11 < 9; fi11++ {
                v12 := GenS2{F0: 2, F1: 0}
                if true {
                    p0 -= 255
                    fmt.Println(p1, p0)
                } else {
                    var v13 uint8 = (p0 ^ p0)
                    _ = v13
                }
                v14 := GenS3{F0: 42}
                _ = fi11
                _ = v12
                _ = v14
            }
            switch "" {
            case "foo bar":
                var v15 float32 = 1.5
                _ = v15
            default:
                v16 := GenS1{F0: 735, F1: p0}
                _ = v16
            }
            v17 := GenS2{F0: 2147483647, F1: -323}
            _ = v17
        } else {
            for fi18 := 0; fi18 < 3; fi18++ {
                var v19 float32 = float32(0.1)
                _ = fi18
                _ = v19
            }
        }
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v7
    return uint16(p1)
}

func main() {
    var v1 uint = 65535
    v2 := make(map[int]string)
    v3 := []string{(("hello" + ("hello" + ("\n\t" + ("a" + (("hello" + "foo bar") + ""))))) + "foo bar"), "a", "", ""}
    var v4 float32 = (1e10 - 3.14159)
    v5 := make(map[string]uint32)
    var v6 uint = v1
    v1 += (v1 + v6)
    v20 := v2[608]
    if ((0.0 - -2.25) == 0.1) {
        v3 = append(v3, (v20 + ((v20 + "hello") + (v20 + v20))))
    }
    fmt.Println(v20, v20)
    fmt.Println(v4, v20, v20)
    v21 := GenS2{F0: 1, F1: 1}
    v6 = (v6 * (v6 << (((v1 - (v1 & (v6 << ((65535) & 63))))) & 63)))
    v20 = "\n\t"
    var v22 float64 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
