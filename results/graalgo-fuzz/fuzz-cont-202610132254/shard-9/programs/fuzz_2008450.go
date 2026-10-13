package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 uint8; F2 uint32 }
type GenS4 struct { F0 int; F1 uint8; F2 int32 }
func genFn0(p0 int64, p1 int32) int8 {
    switch 42 {
    case 9223372036854775807:
        var v7 float32 = -2.25
        _ = v7
    case 42:
        var v8 uint64 = 401
        _ = v8
    default:
        p0 -= (^(p0))
    }
    v9 := []int{}
    v9 = append(v9, 9223372036854775807)
    var v10 float32 = (((0.0 / -2.25) + (float32(1e10) / (float32(3.14159) * 3.14159))) * -2.25)
    var v11 string = ""
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return 1
}

func main() {
    v1 := make(map[int]int)
    var v2 int32 = -517
    var v3 uint8 = 1
    var v4 int32 = (-(v2))
    var v5 int8 = 1
    var v6 uint8 = 2
    if false {
        v12 := GenS1{F0: 2, F1: 32767, F2: (v2 | (v4 & (v2 >> ((65535) & 63))))}
        fmt.Println(v2)
        var v13 float32 = 1.5
        _ = v12
        _ = v13
    } else {
        fmt.Println(v5)
        v14 := GenS1{F0: 1, F1: 155, F2: (v2 | v2)}
        _ = v14
    }
    for fi15 := 0; fi15 < 16; fi15++ {
        switch fi15 {
        case -895:
            var v16 int8 = (v5 - -128)
            _ = v16
        default:
            if ((v6 & 195) == 2) {
                if false {
                    var v17 uint64 = 143
                    _ = v17
                }
            } else {
                v18 := GenS2{F0: 51}
                if (fi15 >= (fi15 << ((42) & 63))) {
                    delete(v1, (fi15 | (fi15 ^ (fi15 ^ (-(fi15))))))
                    v6 = 127
                }
                _ = v18
            }
        }
        v19 := v1[0]
        var v20 float32 = -2.25
        _ = fi15
        _ = v19
        _ = v20
    }
    var v21 uint16 = 32767
    delete(v1, 9223372036854775807)
    switch "" {
    case "a":
        v2 += 2147483647
    }
    var v22 int32 = v2
    fmt.Println(v2)
    fmt.Println(v22, v2)
    switch 32767 {
    case -406:
        for fi23 := 0; fi23 < 4; fi23++ {
            v24 := []int{(fi23 / (((fi23 & (fi23 << ((189) & 63)))) | 1)), (fi23 % ((fi23) | 1))}
            for fi25 := 0; fi25 < 11; fi25++ {
                fmt.Println(v22)
                var v26 int = -9223372036854775808
                _ = fi25
                _ = v26
            }
            for fi27 := 0; fi27 < 20; fi27++ {
                v28 := GenS1{F0: 2, F1: 0, F2: v22}
                _ = fi27
                _ = v28
            }
            _ = fi23
            _ = v24
        }
    case -1:
        v21 += (v21 << ((255) & 63))
    case -9223372036854775808:
        v29 := v1[128]
        _ = v29
    default:
        var v30 uint = 816
        _ = v30
    }
    v6 = (^(v3))
    var v31 uint16 = (v21 / (((^(v21))) | 1))
    v32 := []int{1, int(v3), 1}
    for fi33 := 0; fi33 < 5; fi33++ {
        if false {
            v34 := GenS2{F0: v21}
            _ = v34
        }
        if true {
            for fi35 := 0; fi35 < 20; fi35++ {
                fmt.Println(v21, v21, fi33)
                if false {
                    v36 := GenS3{F0: (v6 << ((255) & 63)), F1: v3, F2: 0}
                    v37 := GenS2{F0: (v21 % (((v21 + 32767)) | 1))}
                    var v38 bool = (!(("foo bar" + "a") == ("hello" + "foo bar")))
                    _ = v36
                    _ = v37
                    _ = v38
                } else {
                    if false {
                        var v39 int = -660
                        v2 = (^(v4))
                        var v40 uint = 306
                        _ = v39
                        _ = v40
                    }
                    v4 -= -1
                }
                _ = fi35
            }
        }
        v41 := GenS4{F0: fi33, F1: 250, F2: (v4 | (^(v2)))}
        _ = fi33
        _ = v41
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
