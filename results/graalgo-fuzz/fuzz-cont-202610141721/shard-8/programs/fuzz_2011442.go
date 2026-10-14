package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8 }
func genFn0() uint16 {
    v6 := []int{128}
    var v7 uint = 835
    _ = v6
    _ = v7
    return 32767
}
func genFn1(p0 int64, p1 uint8, p2 uint64) uint16 {
    for fi8 := 0; fi8 < 5; fi8++ {
        p2 -= (p2 * (p2 - p2))
        p2 += p2
        _ = fi8
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    return 826
}

func main() {
    var v1 float32 = (3.14159 / (((((float32(0.0) / (float32((3.14159 / 3.14159)) + (1e10 * 1.5))) + 3.14159) - (3.14159 * (1e10 / (3.14159 / 1.0)))) / 1.0) / 1.0))
    v2 := []string{}
    var v3 uint64 = 42
    var v4 float64 = -2.25
    v5 := make(map[string]int)
    for fi9 := 0; fi9 < 7; fi9++ {
        v10 := GenS1{F0: -1}
        _ = fi9
        _ = v10
    }
    v1 += (((v1 - ((v1 * -2.25) * (((v1 - ((0.1 / 1.0) - v1)) - v1) + ((v1 + (((0.0 - 1e10) / v1) + v1)) - v1)))) / (0.1 / 1e10)) + ((float32(float64(v1)) * (v1 + 1e10)) * v1))
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    v11 := v5["\n\t"]
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    v12 := GenS1{F0: 42}
    v11 -= (v11 >> ((18446744073709551615) & 63))
    var v13 int64 = -9223372036854775808
    var v14 uint16 = 0
    v15 := v5["a"]
    if false {
        if (v15 >= 128) {
            for fi16 := 0; fi16 < 2; fi16++ {
                if true {
                    v2 = append(v2, "foo bar")
                    fmt.Println(v3)
                    if (false && (!(0 <= (v13 & 358)))) {
                        switch (v11 >> ((847) & 63)) {
                        case 128:
                            v17 := GenS1{F0: 1}
                            _ = v17
                        case 9223372036854775807:
                            var v18 bool = ((v14 ^ v14) != v14)
                            _ = v18
                        case -1:
                            for fi19 := 0; fi19 < 5; fi19++ {
                                v2 = append(v2, ("\n\t" + "foo bar"))
                                var v20 string = ("\n\t" + "")
                                _ = fi19
                                _ = v20
                            }
                        }
                        v21 := GenS1{F0: 42}
                        _ = v21
                    }
                } else {
                    switch (fi16 << ((255) & 63)) {
                    case 4294967295:
                        fmt.Println(v3, v3)
                    case -9223372036854775808:
                        var v22 string = "a"
                        _ = v22
                    default:
                        for fi23 := 0; fi23 < 6; fi23++ {
                            v24 := GenS2{F0: fi23}
                            var v25 uint8 = 42
                            _ = fi23
                            _ = v24
                            _ = v25
                        }
                    }
                }
                v26 := GenS3{F0: 26}
                v13 = v13
                _ = fi16
                _ = v26
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        var v27 int8 = 0
        _ = v27
    }
    v4 += -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
