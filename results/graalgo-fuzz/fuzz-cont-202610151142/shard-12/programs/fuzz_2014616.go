package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 uint16; F2 uint }
func genFn0() uint {
    v6 := GenS1{F0: 169, F1: 1, F2: 42}
    v7 := []int{4294967295, 128}
    _ = len(v7)
    var v8 int = 32767
    var v9 uint = 65535
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 18446744073709551615
}

func main() {
    var v1 uint32 = 65535
    var v2 uint64 = 4294967295
    v3 := make(map[string]string)
    var v4 int = -868
    var v5 float32 = 3.14159
    v2 += (v2 / (((^(v2))) | 1))
    fmt.Println(v2, v4, v1)
    v10 := []int{v4, (v4 - (v4 / (((v4 & v4)) | 1)))}
    v11 := GenS2{F0: 255}
    v10 = append(v10, v4)
    v12 := GenS3{F0: 1, F1: 477, F2: genFn0()}
    if false {
        v13 := GenS2{F0: 206}
        v5 = (float32(3.14159) / 0.1)
        _ = v13
    } else {
        v5 += 0.1
        fmt.Println(v4)
        if true {
            if (true || false) {
                v10 = append(v10, (v4 ^ v4))
                for fi14 := 0; fi14 < 2; fi14++ {
                    var v15 bool = (100 != 100)
                    if (((v4 << ((2) & 63)) < -9223372036854775808) || (2 <= 427)) {
                        v1 = (^(v1))
                    }
                    v15 = v15
                    _ = fi14
                    _ = v15
                }
            }
            var v16 int32 = -2147483648
            _ = v16
        } else {
            _ = len(v3)
            v1 += v1
        }
    }
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
