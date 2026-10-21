package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 bool; F1 int32 }
func genFn0(p0 uint8, p1 bool, p2 string) uint32 {
    switch (p2 + p2) {
    case "\n\t":
        var v4 uint64 = 953
        _ = v4
    }
    var v5 uint32 = 4294967295
    v5 += v5
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 1
}

func main() {
    var v1 int16 = 32767
    var v2 uint16 = 669
    var v3 bool = ((!(!(v2 <= v2))) && ((!(!true)) || false))
    var v6 float32 = (float32(1.5) * (3.14159 - (((((((float32(0.1) * float32(3.14159)) + (0.1 * ((1e10 / 1.0) - 1e10))) / 3.14159) / (((1.5 + 1.5) * 3.14159) + ((3.14159 / 1.0) * 1e10))) * float32(3.14159)) + ((1.5 - 1e10) + float32((0.0 / 1.0)))) / 1e10)))
    v1 += -1
    var v7 int32 = 1
    v8 := GenS1{F0: 18446744073709551615, F1: v3}
    v7 -= v7
    v9 := GenS2{F0: 84, F1: 32767}
    if (!(4294967295 >= int(v2))) {
        if false {
            if (!(!(9223372036854775807 >= 127))) {
                if ("hello" == "a") {
                    v10 := []int{0, -1, -867, 0}
                    _ = v10
                }
                v11 := []int{}
                v11 = append(v11, 0)
                _ = v11
            }
            v6 = -2.25
        }
        for fi12 := 0; fi12 < 19; fi12++ {
            if false {
                var v13 uint32 = 0
                v14 := GenS3{F0: (v7 < (-(v7))), F1: v7}
                _ = v13
                _ = v14
            } else {
                if (!v3) {
                    v15 := []int{32768, (fi12 >> ((65535) & 63))}
                    _ = v15
                }
            }
            fmt.Println(v7, v3)
            _ = fi12
        }
        var v16 float32 = ((v6 - ((float32((1.5 + float64(-2.25))) / (((v6 - 1e10) - v6) + ((v6 - float32(((1e10 - (1.5 / 3.14159)) + 0.0))) * v6))) - (v6 - 0.0))) * v6)
        _ = v16
    } else {
        v6 += v6
        v17 := []int{-9223372036854775808}
        _ = v17
    }
    var v18 float32 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
