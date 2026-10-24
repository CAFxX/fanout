package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 bool = false
    v3 := make(map[uint32]int)
    var v4 int8 = 100
    var v5 int8 = v4
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 int8 = (v5 ^ v4)
        if (!(!v2)) {
            for fi8 := 0; fi8 < 13; fi8++ {
                var v9 bool = true
                v10 := []int{(fi6 & 32767)}
                _ = fi8
                _ = v10
                _ = v9
            }
            var v11 uint = 65535
            if true {
                fmt.Println(v11)
            } else {
                for fi12 := 0; fi12 < 16; fi12++ {
                    switch "foo bar" {
                    case "":
                        v7 += (v5 % (((v5 >> (((v11 | (^(v11)))) & 63))) | 1))
                    case "foo bar":
                        v13 := []int{}
                        _ = v13
                    default:
                        var v14 bool = true
                        _ = v14
                    }
                    v4 = (v7 * (v4 % ((1) | 1)))
                    var v15 uint16 = 133
                    _ = fi12
                    _ = v15
                }
                var v16 bool = false
                var v17 bool = true
                _ = v16
                _ = v17
            }
            _ = v11
        } else {
            if v2 {
                if (2147483647 > 472) {
                    _ = len(v1)
                    v1[4294967295] = 0
                } else {
                    v2 = v2
                    for fi18 := 0; fi18 < 11; fi18++ {
                        v19 := GenS1{F0: 2}
                        if (0 < -2147483648) {
                            if (63 < 116) {
                                v20 := []int{(fi18 ^ 42)}
                                v2 = true
                                _ = v20
                            }
                            v2 = (127 > 18)
                            var v21 int16 = -424
                            v5 = v7
                            _ = v21
                        }
                        _ = fi18
                        _ = v19
                    }
                    switch (fi6 >> ((0) & 63)) {
                    case 42:
                        v7 += 100
                    default:
                        v22 := []int{(^(fi6))}
                        _ = v22
                    }
                }
            } else {
                v1[2147483648] = 2147483647
            }
            switch ("" + ("\n\t" + "foo bar")) {
            case "foo bar":
                var v23 float32 = ((float32(((((float64((1e10 + (3.14159 + 1e10))) + 1.5) / 1e10) * 3.14159) - 3.14159)) + 1e10) / (0.0 + (((((3.14159 * 0.1) + float32(((3.14159 + 1.5) - float64(3.14159)))) - -2.25) * ((((3.14159 + (3.14159 / 1.5)) + 0.0) + float32(0.0)) + ((0.1 / 1e10) - 0.0))) + 1e10)))
                _ = v23
            case "a":
                v24 := GenS1{F0: 0}
                _ = v24
            }
        }
        v25 := GenS1{F0: 18446744073709551615}
        _ = fi6
        _ = v25
        _ = v7
    }
    if (395 >= 2) {
        v5 += v5
        var v26 uint16 = 0
        _ = v26
    } else {
        var v27 uint8 = 1
        var v28 uint16 = 401
        _ = v27
        _ = v28
    }
    for fi29 := 0; fi29 < 13; fi29++ {
        for fi30 := 0; fi30 < 11; fi30++ {
            var v31 bool = true
            v2 = (65535 != 2147483647)
            _ = fi30
            _ = v31
        }
        v2 = (!(!v2))
        var v32 int8 = v5
        _ = fi29
        _ = v32
    }
    for fi33 := 0; fi33 < 10; fi33++ {
        v34 := GenS1{F0: 4294967295}
        fmt.Println(v4, v4)
        _ = fi33
        _ = v34
    }
    v35 := v3[2]
    v5 = 0
    for fi36 := 0; fi36 < 19; fi36++ {
        v37 := []int{int(v5), fi36, (fi36 + -456), v35}
        var v38 uint32 = 63
        v39 := v3[(v38 << ((255) & 63))]
        _ = fi36
        _ = v37
        _ = v38
        _ = v39
    }
    v40 := GenS1{F0: 2}
    for fi41 := 0; fi41 < 19; fi41++ {
        var v42 float64 = 0.1
        for fi43 := 0; fi43 < 3; fi43++ {
            v1[1] = 206
            _ = fi43
        }
        _ = fi41
        _ = v42
    }
    for fi44 := 0; fi44 < 19; fi44++ {
        v3[65535] = fi44
        var v45 int32 = -360
        v35 += fi44
        _ = fi44
        _ = v45
    }
    fmt.Println(v35)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v40
    _ = v5
}
