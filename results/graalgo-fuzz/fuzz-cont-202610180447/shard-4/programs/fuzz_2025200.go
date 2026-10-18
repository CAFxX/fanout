package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }

func main() {
    var v1 uint16 = 32767
    var v2 float32 = 0.0
    v3 := make(map[string]int)
    var v4 uint = 1
    var v5 bool = (false || true)
    var v6 bool = false
    for fi7 := 0; fi7 < 18; fi7++ {
        for fi8 := 0; fi8 < 17; fi8++ {
            if (9223372036854775807 <= fi7) {
                fmt.Println(v2, fi7)
                if ((v6 || (1e10 < v2)) && (((2147483647 >= 0) && v6) && (1 != 0))) {
                    v6 = (1 > 0)
                }
                v4 -= (v4 >> (((v4 << (((v4 << ((4294967295) & 63))) & 63))) & 63))
                fmt.Println(v1, v6, v2)
            } else {
                var v9 int = fi8
                _ = v9
            }
            v6 = false
            _ = fi8
        }
        _ = fi7
    }
    v4 -= v4
    v1 += v1
    for fi10 := 0; fi10 < 10; fi10++ {
        for fi11 := 0; fi11 < 4; fi11++ {
            switch fi11 {
            case 1:
                v12 := []int{}
                _ = v12
            case 32768:
                v13 := v3[""]
                _ = v13
            case 875:
                var v14 int = (^(fi11))
                _ = v14
            }
            v5 = ((-2147483648 == -575) && v5)
            var v15 int16 = 42
            _ = fi11
            _ = v15
        }
        v4 = (v4 - v4)
        v16 := GenS1{F0: v4, F1: 860}
        _ = fi10
        _ = v16
    }
    v6 = (v5 && true)
    v1 += v1
    if true {
        if v5 {
            var v17 int = -816
            switch v17 {
            case 4294967295:
                v1 = 32767
            default:
                v18 := GenS1{F0: (v4 * (^(v4))), F1: 331}
                _ = v18
            }
            v3[("" + (("a" + "a") + "hello"))] = v17
            _ = v17
        }
    }
    var v19 int32 = -256
    delete(v3, "hello")
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
