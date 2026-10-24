package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int }
func genFn0(p0 uint, p1 int32, p2 uint32) uint16 {
    switch 127 {
    case 128:
        p2 += (p2 + p2)
    }
    p0 -= p0
    var v4 int8 = -37
    p0 += (p0 % ((uint(p2)) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 0
}
func genFn1(p0 int8, p1 uint32) uint32 {
    switch -956 {
    case -1:
        var v5 int8 = 100
        _ = v5
    case 1:
        var v6 int = -2147483648
        _ = v6
    case -600:
        for fi7 := 0; fi7 < 19; fi7++ {
            v8 := []int{}
            var v9 uint32 = (^(p1))
            _ = fi7
            _ = v8
            _ = v9
        }
    default:
        if false {
            v10 := GenS1{F0: p1}
            if true {
                p1 -= (p1 << ((964) & 63))
                for fi11 := 0; fi11 < 20; fi11++ {
                    switch -9223372036854775808 {
                    case -1:
                        var v12 float64 = (-2.25 + 1.5)
                        _ = v12
                    case 41:
                        if false {
                            p1 += 2
                            v13 := []int{128}
                            _ = v13
                        }
                    }
                    v14 := []int{fi11, (-(fi11)), -947, (fi11 * (^(fi11)))}
                    _ = fi11
                    _ = v14
                }
            } else {
                fmt.Println(p0)
                for fi15 := 0; fi15 < 7; fi15++ {
                    v16 := GenS1{F0: 65535}
                    _ = fi15
                    _ = v16
                }
            }
            if (1.5 < float64((0.0 * 1.5))) {
                p1 = 2147483648
                switch 1 {
                case 638:
                    v17 := []int{1}
                    _ = v17
                }
            } else {
                for fi18 := 0; fi18 < 17; fi18++ {
                    var v19 uint64 = 18446744073709551615
                    v20 := GenS1{F0: p1}
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                var v21 int16 = 935
                switch 0 {
                case 32768:
                    p0 += p0
                case 42:
                    switch -586 {
                    case -41:
                        v22 := GenS2{F0: 2}
                        _ = v22
                    case 128:
                        v23 := []int{1, 462, -131, 42}
                        _ = v23
                    }
                default:
                    p0 = p0
                }
                _ = v21
            }
            _ = v10
        } else {
            var v24 int64 = -9223372036854775808
            if ((-(p0)) == (p0 & p0)) {
                var v25 uint64 = 0
                p1 = (^(p1))
                switch -1 {
                case -9223372036854775808:
                    if true {
                        v26 := []int{32767, 4294967295}
                        switch "" {
                        case "hello":
                            p1 += uint32(v25)
                        }
                        _ = v26
                    }
                case 643:
                    fmt.Println(p0, v24)
                default:
                    fmt.Println(p1, p0)
                }
                _ = v25
            }
            _ = v24
        }
    }
    p0 -= -1
    _ = p0
    _ = p1
    return p1
}

func main() {
    var v1 uint32 = 42
    v2 := []uint32{v1, (v1 - 65535), 553}
    var v3 uint16 = 42
    fmt.Println(v3, v3)
    fmt.Println(v1)
    v27 := GenS3{F0: 0}
    v28 := GenS1{F0: v1}
    v3 = v3
    var v29 int16 = 1000
    v30 := GenS1{F0: v1}
    for fi31 := 0; fi31 < 11; fi31++ {
        v3 -= v3
        var v32 string = "hello"
        _ = fi31
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}
