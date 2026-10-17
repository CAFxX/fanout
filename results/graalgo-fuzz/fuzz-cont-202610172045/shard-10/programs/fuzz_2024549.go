package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint16 }
func genFn0(p0 uint32, p1 uint) int32 {
    var v4 uint8 = 88
    for fi5 := 0; fi5 < 19; fi5++ {
        switch "a" {
        case "":
            p1 -= 895
        case "foo bar":
            for fi6 := 0; fi6 < 19; fi6++ {
                v4 = v4
                var v7 uint32 = 4
                v8 := []int{fi6, fi6}
                if false {
                    if true {
                        if true {
                            p0 += 146
                            fmt.Println(v7, v7, fi6)
                        } else {
                            v9 := GenS1{F0: (-683 > int64(p1))}
                            _ = len(v8)
                            v4 += v4
                            _ = v9
                        }
                    }
                }
                p1 -= 1
                if false {
                    v10 := GenS1{F0: (v7 == p0)}
                    _ = v10
                } else {
                    var v11 int32 = -721
                    v8 = append(v8, 128)
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    _ = v11
                }
                var v12 bool = true
                _ = fi6
                _ = v12
                _ = v7
                _ = v8
            }
        default:
            v13 := []int{fi5, fi5, fi5}
            _ = v13
        }
        _ = fi5
    }
    _ = p0
    _ = p1
    _ = v4
    return 277
}

func main() {
    var v1 uint = 0
    v2 := []string{("foo bar" + ("a" + ""))}
    var v3 float32 = 1e10
    v1 += v1
    v1 -= (v1 % (((v1 | (^(v1)))) | 1))
    var v14 int64 = 1
    if (false || false) {
        switch 42 {
        case 0:
            switch 154 {
            case 915:
                v15 := GenS1{F0: false}
                _ = v15
            default:
                v14 = (v14 + (v14 & v14))
            }
        case 2147483647:
            if (!(2 <= 2)) {
                v3 += v3
            }
        default:
            v14 = v14
        }
        v16 := GenS2{F0: 807, F1: 42, F2: 42}
        if (v14 <= (v14 * (v14 | v14))) {
            var v17 uint64 = 18446744073709551615
            if ((v3 > v3) && (v14 != (v14 & -384))) {
                var v18 uint = (^(v1))
                var v19 float64 = (1e10 * 1e10)
                v20 := GenS2{F0: 2147483647, F1: 1, F2: 476}
                _ = v18
                _ = v19
                _ = v20
            }
            v21 := GenS2{F0: 2147483648, F1: 2, F2: 234}
            _ = v17
            _ = v21
        }
        _ = v16
    }
    v14 -= -1
    var v22 uint64 = 0
    var v23 string = "\n\t"
    v2 = append(v2, "foo bar")
    var v24 int16 = int16(v14)
    for fi25 := 0; fi25 < 4; fi25++ {
        var v26 int8 = int8(v24)
        _ = fi25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
