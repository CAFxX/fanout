package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
func genFn0(p0 bool) string {
    var v7 bool = p0
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 int16 = -119
        if v7 {
            v9 = -534
            v10 := GenS1{F0: 166, F1: ("" + "foo bar")}
            _ = v10
        } else {
            v7 = (-1 < 0)
            if ((!v7) && (!(2 >= 636))) {
                var v11 int8 = 42
                for fi12 := 0; fi12 < 10; fi12++ {
                    v9 = v9
                    p0 = v7
                    _ = fi12
                }
                _ = v11
            }
        }
        v13 := []int{fi8, (fi8 + (fi8 & fi8))}
        _ = fi8
        _ = v13
        _ = v9
    }
    p0 = true
    _ = p0
    _ = v7
    return ("" + "")
}

func main() {
    var v1 uint = 65535
    var v2 uint16 = 32767
    v3 := make(map[uint32]uint32)
    var v4 bool = (2 != 0)
    v5 := []uint32{}
    var v6 uint = 730
    v2 -= v2
    v14 := v3[904]
    v14 = uint32(v2)
    v14 -= v14
    switch 32767 {
    case 405:
        v4 = v4
    case 32767:
        if true {
            switch "" {
            case "":
                var v15 uint32 = (v14 | uint32(v2))
                _ = v15
            case "a":
                for fi16 := 0; fi16 < 18; fi16++ {
                    var v17 uint8 = 255
                    _ = fi16
                    _ = v17
                }
            }
        } else {
            v18 := v3[v14]
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 14; fi19++ {
        v20 := GenS1{F0: 416, F1: ((("" + "hello") + "hello") + "foo bar")}
        v21 := GenS1{F0: v6, F1: "a"}
        _ = fi19
        _ = v20
        _ = v21
    }
    var v22 float32 = 1.5
    if (42 == 32768) {
        for fi23 := 0; fi23 < 2; fi23++ {
            var v24 uint8 = 0
            v6 = (^(v1))
            var v25 int16 = 1
            _ = fi23
            _ = v24
            _ = v25
        }
        fmt.Println(v22)
        var v26 bool = (v1 != (v1 << ((uint(v14)) & 63)))
        _ = v26
    }
    if v4 {
        for fi27 := 0; fi27 < 15; fi27++ {
            var v28 int16 = 1
            _ = fi27
            _ = v28
        }
        if v4 {
            fmt.Println(v22)
            var v29 int8 = -7
            _ = v29
        } else {
            var v30 int32 = 798
            var v31 bool = (!(v4 || (1e10 >= (float64((float32((((1.5 / 0.1) + ((1e10 / (0.1 / -2.25)) * 0.0)) - 3.14159)) - v22)) + 3.14159))))
            v30 += (v30 ^ v30)
            _ = v30
            _ = v31
        }
    }
    fmt.Println(v14, v14, v14)
    if false {
        fmt.Println(v6)
        v2 = 255
    } else {
        v22 = (((1.5 * v22) * (v22 / -2.25)) / -2.25)
        for fi32 := 0; fi32 < 12; fi32++ {
            fmt.Println(v4, v6)
            var v33 uint8 = 1
            _ = fi32
            _ = v33
        }
    }
    switch 128 {
    case 141:
        v4 = (!false)
    case 127:
        for fi34 := 0; fi34 < 17; fi34++ {
            if v4 {
                v35 := v3[v14]
                _ = len(v3)
                _ = v35
            }
            v6 += v6
            _ = fi34
        }
    default:
        var v36 int8 = -26
        _ = v36
    }
    v37 := GenS1{F0: 2, F1: "\n\t"}
    switch -232 {
    case -802:
        var v38 bool = (!((v1 % ((uint(v14)) | 1)) >= (v1 / ((0) | 1))))
        _ = v38
    case 2147483647:
        v14 -= 0
    default:
        for fi39 := 0; fi39 < 11; fi39++ {
            switch (^(fi39)) {
            case 0:
                var v40 uint64 = 1
                _ = v40
            case 779:
                for fi41 := 0; fi41 < 3; fi41++ {
                    var v42 uint = 65535
                    var v43 float64 = (float64(v22) - (((3.14159 / (3.14159 + (((1.5 / (1e10 - 0.0)) - (3.14159 / (((1.5 - 3.14159) * -2.25) / 1.5))) + float64(v22)))) - 1.5) + 1e10))
                    _ = fi41
                    _ = v42
                    _ = v43
                }
            case 127:
                v22 = -2.25
            default:
                switch ("foo bar" + "a") {
                case "hello":
                    var v44 uint8 = 242
                    _ = v44
                case "a":
                    v14 = (v14 & 422)
                default:
                    var v45 int32 = 2147483647
                    _ = v45
                }
            }
            if v4 {
                var v46 int32 = -678
                var v47 int = (fi39 & fi39)
                _ = v46
                _ = v47
            } else {
                v22 += (float32(0.1) + (float32(1.5) * v22))
                var v48 uint64 = 4294967295
                var v49 uint16 = 65535
                _ = v48
                _ = v49
            }
            var v50 int32 = -19
            _ = fi39
            _ = v50
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
