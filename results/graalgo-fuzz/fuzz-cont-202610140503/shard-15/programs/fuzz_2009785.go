package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 bool, p1 string, p2 uint16) uint32 {
    var v6 bool = (4294967295 >= -2147483648)
    var v7 int64 = 179
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 121
}
func genFn1(p0 int) int32 {
    v8 := GenS1{F0: (!false), F1: 0, F2: (!true)}
    var v9 int16 = -32768
    p0 = p0
    _ = p0
    _ = v8
    _ = v9
    return 2147483647
}

func main() {
    var v1 int32 = 222
    var v2 uint64 = 1
    v3 := make(map[uint32]string)
    var v4 int = -711
    v5 := make(map[uint32]bool)
    var v10 int32 = v1
    var v11 int16 = -1
    v12 := []int{0, v4, (v4 % (((v4 % (((v4 / (((v4 & 32768)) | 1))) | 1))) | 1))}
    if (true || true) {
        _ = len(v3)
        var v13 string = ""
        v14 := GenS2{F0: 472}
        _ = v13
        _ = v14
    } else {
        v15 := v3[42]
        _ = v15
    }
    v16 := GenS1{F0: (1 < v11), F1: 100, F2: (65535 < 79)}
    v17 := GenS2{F0: 65535}
    switch ("" + ("foo bar" + "\n\t")) {
    case "":
        for fi18 := 0; fi18 < 17; fi18++ {
            var v19 float32 = 3.14159
            switch v4 {
            case 898:
                var v20 int = (-(fi18))
                _ = v20
            default:
                v4 += fi18
            }
            _ = fi18
            _ = v19
        }
    }
    fmt.Println(v1, v11)
    v21 := GenS1{F0: false, F1: -1, F2: (!(!((0.1 * (0.1 / ((1e10 / 1.0) / (1.5 / (1.5 - -2.25))))) > (0.0 * ((float64((float32(-2.25) + ((1e10 / float32(3.14159)) / 0.1))) + 0.1) - 1e10)))))}
    var v22 bool = (591 == 2)
    v10 += (v1 - v1)
    v11 = (v11 >> ((4294967295) & 63))
    v12 = append(v12, (v4 / ((v4) | 1)))
    if (float64(1.5) != 0.0) {
        for fi23 := 0; fi23 < 8; fi23++ {
            v24 := GenS3{F0: v2}
            v25 := GenS2{F0: 1}
            _ = fi23
            _ = v24
            _ = v25
        }
    }
    fmt.Println(v10, v10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
