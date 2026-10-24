package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int32, p1 int64) uint {
    p1 += p1
    var v6 uint64 = 18446744073709551615
    _ = p0
    _ = p1
    _ = v6
    return 255
}

func main() {
    v1 := make(map[int]string)
    v2 := []string{(("foo bar" + "hello") + (("foo bar" + "\n\t") + "hello"))}
    var v3 int16 = -225
    v4 := []int{}
    v5 := make(map[int]string)
    v7 := GenS1{F0: 77}
    v4 = append(v4, 9223372036854775807)
    v3 -= (v3 << ((18446744073709551615) & 63))
    v3 = v3
    for fi8 := 0; fi8 < 7; fi8++ {
        v3 = 32767
        _ = fi8
    }
    var v9 int64 = -9223372036854775808
    v9 = v9
    v10 := GenS1{F0: 100}
    var v11 float32 = 0.0
    if false {
        fmt.Println(v9, v3, v11)
    } else {
        v2 = append(v2, "hello")
        v9 = 42
        for fi12 := 0; fi12 < 10; fi12++ {
            _ = len(v5)
            v13 := GenS2{F0: 4294967295}
            _ = fi12
            _ = v13
        }
    }
    var v14 string = "\n\t"
    for fi15 := 0; fi15 < 20; fi15++ {
        switch -460 {
        case -9223372036854775808:
            v3 = v3
        default:
            if true {
                if (true && (1 >= v9)) {
                    _ = len(v4)
                    var v16 bool = (true || (54 >= -35))
                    _ = v16
                }
            }
        }
        if (!true) {
            v14 = (v14 + ("hello" + v14))
            if (((0.0 == v11) && false) && false) {
                v17 := GenS1{F0: 42}
                v18 := GenS2{F0: 1}
                delete(v5, fi15)
                _ = v17
                _ = v18
            } else {
                v19 := GenS2{F0: 4294967295}
                if len(v4) > 0 {
                    v4[0] = -739
                }
                v11 -= 1e10
                _ = v19
            }
            _ = len(v4)
        } else {
            switch "a" {
            case "foo bar":
                var v20 string = v14
                _ = v20
            default:
                v14 = v14
            }
            v3 = (v3 >> ((65535) & 63))
        }
        _ = fi15
    }
    if true {
        v3 = (v3 % ((-32768) | 1))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v9
}
