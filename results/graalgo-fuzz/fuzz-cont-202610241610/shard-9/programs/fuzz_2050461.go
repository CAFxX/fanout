package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 uint32 = 2
    var v2 float32 = ((((0.1 + float32(((1e10 + 0.0) + 1.5))) + 1e10) * -2.25) / -2.25)
    var v3 uint8 = 186
    fmt.Println(v3, v3, v1)
    v1 -= v1
    var v4 uint64 = 4294967295
    v5 := []int{42, 128, 42}
    var v6 bool = (!false)
    v7 := GenS1{F0: 2147483648}
    v2 = ((1.5 + 3.14159) / (v2 / ((((v2 * (v2 / v2)) + -2.25) - v2) / v2)))
    for fi8 := 0; fi8 < 10; fi8++ {
        v4 += 4294967295
        var v9 uint = 42
        _ = fi8
        _ = v9
    }
    v1 = v1
    if v6 {
        var v10 uint16 = 255
        switch 9223372036854775807 {
        case 32768:
            var v11 uint = uint(v4)
            _ = v11
        default:
            for fi12 := 0; fi12 < 9; fi12++ {
                var v13 float32 = 3.14159
                if ((v2 / v13) <= v13) {
                    for fi14 := 0; fi14 < 5; fi14++ {
                        var v15 int8 = 0
                        if (97 == 2147483647) {
                            v6 = (fi12 <= (fi14 >> ((42) & 63)))
                            if (v4 > v4) {
                                switch (fi12 / ((32767) | 1)) {
                                case -863:
                                    v16 := GenS1{F0: (v1 << ((2) & 63))}
                                    _ = v16
                                }
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                var v17 int32 = 2147483647
                                _ = v17
                            } else {
                                if v6 {
                                    v15 = 1
                                }
                                var v18 bool = false
                                var v19 uint8 = (v3 - 159)
                                _ = v18
                                _ = v19
                            }
                            v10 += 685
                        } else {
                            v4 = (^(v4))
                        }
                        v1 -= (^(v1))
                        _ = fi14
                        _ = v15
                    }
                } else {
                    var v20 int64 = 9223372036854775807
                    v21 := GenS1{F0: (v1 << ((42) & 63))}
                    _ = v20
                    _ = v21
                }
                for fi22 := 0; fi22 < 9; fi22++ {
                    v23 := GenS1{F0: v1}
                    v5 = append(v5, 127)
                    _ = fi22
                    _ = v23
                }
                _ = fi12
                _ = v13
            }
        }
        switch "a" {
        case "":
            for fi24 := 0; fi24 < 5; fi24++ {
                v10 -= (v10 >> ((4294967295) & 63))
                _ = fi24
            }
        case "a":
            v25 := GenS1{F0: v1}
            _ = v25
        }
        _ = v10
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
