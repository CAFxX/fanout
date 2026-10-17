package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string }
func genFn0(p0 uint8, p1 uint32) int {
    var v7 uint16 = 0
    for fi8 := 0; fi8 < 18; fi8++ {
        v9 := []int{}
        var v10 bool = ((v7 / ((v7) | 1)) <= (^(v7)))
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 6; fi11++ {
        var v12 uint8 = (p0 - (p0 ^ 107))
        _ = fi11
        _ = v12
    }
    if (p1 >= (p1 * (p1 & (p1 & (p1 << ((18446744073709551615) & 63)))))) {
        for fi13 := 0; fi13 < 4; fi13++ {
            if (true || true) {
                switch "\n\t" {
                case "\n\t":
                    v7 = v7
                default:
                    for fi14 := 0; fi14 < 6; fi14++ {
                        fmt.Println(p1, v7)
                        switch 665 {
                        case -556:
                            v15 := []int{42, (^(fi14)), fi14}
                            _ = v15
                        case 32767:
                            switch ("foo bar" + "") {
                            case "hello":
                                v16 := []int{}
                                _ = v16
                            default:
                                v17 := []int{(^(fi13)), (fi14 & (-(fi13)))}
                                _ = v17
                            }
                        case -61:
                            var v18 int8 = 42
                            _ = v18
                        }
                        p0 = 1
                        _ = fi14
                    }
                }
            } else {
                v19 := []int{(fi13 ^ 32768), 32768, -2147483648}
                _ = v19
            }
            _ = fi13
        }
        for fi20 := 0; fi20 < 16; fi20++ {
            var v21 uint16 = (v7 << ((18446744073709551615) & 63))
            _ = fi20
            _ = v21
        }
        p0 = (p0 * p0)
    }
    _ = p0
    _ = p1
    _ = v7
    return 128
}
func genFn1() bool {
    var v22 int32 = -1
    if (42 <= -128) {
        for fi23 := 0; fi23 < 8; fi23++ {
            v24 := GenS1{F0: fi23}
            _ = fi23
            _ = v24
        }
    } else {
        var v25 int16 = 1
        _ = v25
    }
    _ = v22
    return true
}

func main() {
    var v1 uint64 = 0
    var v2 uint = 887
    var v3 bool = true
    var v4 string = "foo bar"
    var v5 float64 = (((0.1 + -2.25) / (((3.14159 / -2.25) + 1e10) + 3.14159)) / 1.5)
    var v6 string = ("hello" + v4)
    var v26 uint64 = v1
    for fi27 := 0; fi27 < 19; fi27++ {
        v28 := []int{(^(fi27)), 908}
        v29 := GenS2{F0: fi27}
        var v30 int8 = -1
        _ = fi27
        _ = v28
        _ = v29
        _ = v30
    }
    var v31 int64 = -343
    switch (v4 + ((v6 + "foo bar") + v4)) {
    case "hello":
        for fi32 := 0; fi32 < 10; fi32++ {
            v1 -= 900
            if (253 == 2147483647) {
                var v33 int64 = v31
                _ = v33
            } else {
                var v34 uint16 = 32767
                _ = v34
            }
            v35 := []int{fi32, 127}
            _ = fi32
            _ = v35
        }
    }
    v2 = uint(v1)
    v36 := GenS1{F0: 2147483647}
    v37 := GenS2{F0: -1}
    v1 = v1
    v38 := GenS1{F0: -9223372036854775808}
    var v39 int64 = (v31 >> (((v2 * (v2 & 65535))) & 63))
    for fi40 := 0; fi40 < 3; fi40++ {
        var v41 float32 = (3.14159 / (-2.25 * float32(v5)))
        var v42 float32 = 1e10
        if (42 > int32(v39)) {
            v26 -= 732
            var v43 int8 = 42
            _ = v43
        } else {
            v42 -= v41
        }
        _ = fi40
        _ = v41
        _ = v42
    }
    v44 := []int{32768, -1}
    var v45 uint8 = 1
    v45 -= v45
    for fi46 := 0; fi46 < 12; fi46++ {
        v26 -= v1
        _ = fi46
    }
    v3 = v3
    switch 128 {
    case -9223372036854775808:
        for fi47 := 0; fi47 < 6; fi47++ {
            var v48 float32 = 1e10
            var v49 int32 = 1
            var v50 int64 = (v31 ^ v31)
            _ = fi47
            _ = v48
            _ = v49
            _ = v50
        }
    case -589:
        switch "" {
        case "\n\t":
            if (3.14159 <= (float64(1e10) - v5)) {
                v5 = -2.25
            } else {
                for fi51 := 0; fi51 < 2; fi51++ {
                    v52 := GenS3{F0: v4}
                    _ = fi51
                    _ = v52
                }
                if len(v44) > 0 {
                    v44[0] = 9223372036854775807
                }
                v44 = append(v44, 128)
            }
        default:
            _ = len(v44)
        }
    default:
        if v3 {
            for fi53 := 0; fi53 < 13; fi53++ {
                if len(v44) >= 2 {
                    v44 = v44[1:len(v44)]
                }
                switch v6 {
                case "hello":
                    v2 -= v2
                }
                v31 += (v31 + -677)
                _ = fi53
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v31
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
