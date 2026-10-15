package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := []string{}
    v2 := make(map[uint32]string)
    v3 := []int{2147483647, -9223372036854775808, 699}
    v4 := []uint32{65535}
    var v5 float64 = float64(1e10)
    v6 := v2[220]
    v7 := GenS1{F0: 704}
    v8 := GenS1{F0: 402}
    v9 := GenS1{F0: 2}
    var v10 bool = (833 == 0)
    if len(v3) > 0 {
        v3[0] = 128
    }
    if v10 {
        v10 = (true || (false || (4294967295 == 118)))
    }
    v11 := GenS1{F0: 2}
    for fi12 := 0; fi12 < 6; fi12++ {
        fmt.Println(v5, v6)
        if true {
            var v13 uint8 = 42
            v14 := v2[uint32(v13)]
            v15 := GenS1{F0: 18446744073709551615}
            _ = v13
            _ = v14
            _ = v15
        } else {
            for fi16 := 0; fi16 < 19; fi16++ {
                v17 := GenS1{F0: 0}
                var v18 int = fi12
                _ = fi16
                _ = v17
                _ = v18
            }
            if (127 == 82) {
                if (!v10) {
                    var v19 float32 = (1e10 - 3.14159)
                    v19 = v19
                    _ = v19
                } else {
                    fmt.Println(v10)
                }
                var v20 uint = 377
                switch "" {
                case "hello":
                    v6 = v6
                case "foo bar":
                    v10 = (0 != 348)
                default:
                    if v10 {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v21 string = ("a" + "")
                        _ = v21
                    }
                }
                _ = v20
            }
        }
        _ = fi12
    }
    for fi22 := 0; fi22 < 13; fi22++ {
        for fi23 := 0; fi23 < 17; fi23++ {
            var v24 float64 = v5
            var v25 uint16 = 255
            _ = fi23
            _ = v24
            _ = v25
        }
        _ = fi22
    }
    v1 = append(v1, (v6 + v6))
    v26 := GenS1{F0: 4294967295}
    var v27 bool = (!v10)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
