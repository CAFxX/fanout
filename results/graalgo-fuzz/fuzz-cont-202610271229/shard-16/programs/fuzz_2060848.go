package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    v1 := make(map[int]int)
    var v2 bool = true
    var v3 uint8 = 67
    v4 := make(map[uint32]string)
    var v5 int32 = 0
    delete(v1, 127)
    var v6 uint64 = 0
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := GenS1{F0: 42}
        var v9 uint = 898
        for fi10 := 0; fi10 < 2; fi10++ {
            if ((-(fi10)) != 2147483647) {
                var v11 float32 = ((((0.1 / 1e10) / 0.1) / 1.0) * (0.1 - 1.5))
                var v12 int16 = int16(v5)
                var v13 uint16 = 255
                v11 = v11
                _ = v11
                _ = v12
                _ = v13
            }
            _ = len(v4)
            delete(v1, (^(fi10)))
            v3 -= 161
            _ = fi10
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    v5 = -1
    delete(v4, 2147483648)
    var v14 uint16 = 65535
    var v15 float64 = 1.5
    var v16 uint = 65535
    switch 32767 {
    case 2147483647:
        if (-100 >= 127) {
            v17 := v1[32767]
            for fi18 := 0; fi18 < 3; fi18++ {
                for fi19 := 0; fi19 < 16; fi19++ {
                    v20 := GenS1{F0: int8(v16)}
                    v21 := v4[2147483647]
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                _ = fi18
            }
            v16 += (^(v16))
            _ = v17
        } else {
            v6 = 0
        }
    case 982:
        v22 := GenS1{F0: -113}
        _ = v22
    }
    v23 := GenS1{F0: -128}
    for fi24 := 0; fi24 < 17; fi24++ {
        if (!(((9223372036854775807 < -1) || (v2 && (!false))) && (100 == 88))) {
            var v25 int32 = (v5 >> ((v16) & 63))
            var v26 uint16 = (v14 * (v14 + 2))
            v2 = (!(uint16(v16) > (v26 & (v14 | v26))))
            _ = v25
            _ = v26
        } else {
            if v2 {
                v6 = 4294967295
                for fi27 := 0; fi27 < 11; fi27++ {
                    switch fi24 {
                    case -755:
                        v28 := v4[65535]
                        _ = v28
                    default:
                        delete(v4, 970)
                    }
                    var v29 uint16 = (v14 - v14)
                    var v30 float64 = 1.5
                    _ = fi27
                    _ = v29
                    _ = v30
                }
            }
            v31 := []int{(fi24 << ((v16) & 63))}
            var v32 float64 = (v15 * v15)
            _ = v31
            _ = v32
        }
        v5 -= 2147483647
        _ = fi24
    }
    var v33 uint = (^(v16))
    v34 := []int{9223372036854775807}
    var v35 uint8 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
