package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 string }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 int64 }
func genFn0(p0 int32, p1 int, p2 int64) uint64 {
    var v4 bool = (p2 >= p2)
    var v5 int = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 18446744073709551615
}

func main() {
    var v1 int = 334
    var v2 int8 = 124
    var v3 bool = (!((!true) && true))
    v3 = (0 < 4294967295)
    v2 = (v2 >> ((4294967295) & 63))
    switch "a" {
    case "a":
        v3 = (2 == 2)
    }
    fmt.Println(v2, v1)
    var v6 int16 = -1
    var v7 uint16 = 42
    v3 = v3
    v8 := GenS1{F0: 0, F1: (-(v6)), F2: ("hello" + ("hello" + "\n\t"))}
    if (!((v3 && (!false)) && v3)) {
        if (false || (v3 || (1.5 > (1e10 + ((-2.25 / -2.25) - float64(-2.25)))))) {
            var v9 uint32 = 65535
            v10 := GenS2{F0: v6, F1: v9}
            for fi11 := 0; fi11 < 18; fi11++ {
                var v12 uint8 = 79
                switch fi11 {
                case -2147483648:
                    if (!(1 != v7)) {
                        var v13 uint64 = genFn0(1, 32768, 890)
                        v14 := []int{fi11, v1}
                        _ = v13
                        _ = v14
                    }
                case 4294967295:
                    v15 := []int{fi11, 127, (v1 * -18), (v1 * (^(fi11)))}
                    _ = v15
                default:
                    v3 = v3
                }
                _ = fi11
                _ = v12
            }
            _ = v10
            _ = v9
        } else {
            for fi16 := 0; fi16 < 2; fi16++ {
                v17 := []int{}
                var v18 string = "hello"
                v19 := GenS3{F0: 0}
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
            }
            var v20 int64 = 42
            _ = v20
        }
    }
    v2 = (-(v2))
    v3 = ("a" > "foo bar")
    fmt.Println(v7, v7)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
