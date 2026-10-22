package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int }
type GenS2 struct { F0 int }
func genFn0(p0 bool, p1 int64) uint64 {
    v5 := GenS1{F0: 1, F1: "hello", F2: -721}
    var v6 uint64 = 42
    v7 := []int{}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return v6
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint64 = 0
    var v3 bool = (2147483647 >= -2147483648)
    var v4 int16 = 32767
    fmt.Println(v3, v2)
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 float32 = ((1.5 + ((0.1 - float32(-2.25)) * -2.25)) - 1e10)
        _ = fi8
        _ = v9
    }
    var v10 uint8 = 0
    var v11 bool = (3.14159 < 0.1)
    v1[219] = true
    var v12 int = 128
    for fi13 := 0; fi13 < 18; fi13++ {
        v10 = 42
        if ((v12 >= (fi13 + (v12 << ((4294967295) & 63)))) || (219 < int64(v12))) {
            switch ("a" + ((("a" + "a") + ("" + "hello")) + "")) {
            case "":
                v14 := GenS2{F0: fi13}
                _ = v14
            default:
                var v15 uint = 2
                _ = v15
            }
            for fi16 := 0; fi16 < 1; fi16++ {
                if v11 {
                    if v11 {
                        v17 := []int{1, int(v2), (fi16 ^ v12)}
                        v12 = (-(v12))
                        if (!v3) {
                            v18 := v1[65535]
                            v19 := GenS1{F0: 42, F1: "hello", F2: (fi13 >> ((65535) & 63))}
                            _ = v18
                            _ = v19
                        }
                        var v20 string = ("foo bar" + "a")
                        v21 := GenS1{F0: 163, F1: (("a" + v20) + "foo bar"), F2: -9223372036854775808}
                        _ = v17
                        _ = v20
                        _ = v21
                    }
                    var v22 bool = v3
                    switch fi16 {
                    case 0:
                        fmt.Println(v2, v11)
                    case -2147483648:
                        fmt.Println(v12)
                    default:
                        var v23 uint32 = 16
                        _ = v23
                    }
                    _ = v22
                }
                var v24 uint64 = 0
                _ = fi16
                _ = v24
            }
        }
        _ = fi13
    }
    var v25 float64 = ((3.14159 * -2.25) * float64((1e10 + ((((0.0 / 1.0) + (-2.25 / float32(1.5))) * -2.25) / ((1.5 - -2.25) + (1.5 / 0.1))))))
    v26 := GenS1{F0: 213, F1: ("" + "foo bar"), F2: -657}
    v4 += v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
