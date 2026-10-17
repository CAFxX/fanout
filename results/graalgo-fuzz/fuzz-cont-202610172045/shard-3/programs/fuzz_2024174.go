package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 uint16 }
func genFn0(p0 uint8) uint16 {
    var v6 int32 = 98
    var v7 bool = false
    v6 += (v6 & (v6 << ((825) & 63)))
    v7 = v7
    v8 := []int{}
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 2
}

func main() {
    v1 := []string{""}
    var v2 string = ("hello" + "a")
    v3 := make(map[string]uint32)
    var v4 int32 = -2147483648
    v5 := []string{v2, (v2 + (v2 + v2))}
    fmt.Println(v2)
    if ((!false) || (false || ((!true) && true))) {
        _ = len(v5)
    }
    var v9 uint = 4294967295
    var v10 uint64 = 105
    var v11 int = -9223372036854775808
    var v12 uint32 = 1
    v13 := GenS1{F0: 0, F1: (v9 % (((^(v9))) | 1)), F2: 624}
    delete(v3, v2)
    for fi14 := 0; fi14 < 19; fi14++ {
        for fi15 := 0; fi15 < 14; fi15++ {
            var v16 bool = false
            var v17 float32 = 0.0
            _ = fi15
            _ = v16
            _ = v17
        }
        switch (-(fi14)) {
        case 721:
            for fi18 := 0; fi18 < 14; fi18++ {
                fmt.Println(fi18)
                _ = fi18
            }
        case -1:
            if (false && (!((v9 + (v9 << ((v9) & 63))) >= (v9 / (((v9 & (v9 / ((371) | 1)))) | 1))))) {
                v11 += 1
            } else {
                for fi19 := 0; fi19 < 19; fi19++ {
                    v4 -= v4
                    _ = fi19
                }
            }
        }
        _ = fi14
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
