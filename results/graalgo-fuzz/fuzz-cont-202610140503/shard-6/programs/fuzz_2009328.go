package main

import (
    "fmt"
)

func genFn0() int8 {
    var v7 bool = true
    var v8 int64 = 589
    v9 := []int{127}
    _ = v7
    _ = v8
    _ = v9
    return -1
}

func main() {
    v1 := []uint32{65535, 2147483648}
    v2 := make(map[string]uint32)
    var v3 uint8 = 20
    var v4 uint64 = 0
    v5 := []uint32{uint32(v3)}
    var v6 uint16 = 32767
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v10 float32 = 1.5
    v5 = append(v5, 16)
    var v11 int8 = -128
    v10 -= float32((0.0 + 1e10))
    var v12 uint8 = 127
    var v13 uint64 = (v4 << ((4294967295) & 63))
    v14 := v2[(("foo bar" + "a") + "")]
    v14 -= v14
    if (1e10 <= 3.14159) {
        for fi15 := 0; fi15 < 13; fi15++ {
            fmt.Println(v11, v14)
            var v16 uint16 = v6
            _ = fi15
            _ = v16
        }
        v11 += v11
    } else {
        for fi17 := 0; fi17 < 8; fi17++ {
            for fi18 := 0; fi18 < 9; fi18++ {
                var v19 float32 = float32(0.1)
                _ = fi18
                _ = v19
            }
            _ = len(v2)
            if true {
                v20 := v2["a"]
                v2["\n\t"] = 74
                switch "a" {
                case "":
                    switch fi17 {
                    case 128:
                        if len(v5) > 0 {
                            v5[0] = v20
                        }
                    default:
                        fmt.Println(v20)
                    }
                case "a":
                    fmt.Println(v4)
                case "foo bar":
                    v11 = int8(v12)
                default:
                    if len(v1) > 0 {
                        v1[0] = 65535
                    }
                }
                _ = v20
            }
            _ = fi17
        }
        v13 = v13
        if len(v5) > 0 {
            v5[0] = v14
        }
    }
    v3 = v12
    v11 -= v11
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
