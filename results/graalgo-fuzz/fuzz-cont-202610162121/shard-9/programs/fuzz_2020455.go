package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 uint }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 uint; F1 int8 }
func genFn0(p0 int64) int {
    fmt.Println(p0)
    p0 = (p0 >> ((255) & 63))
    for fi4 := 0; fi4 < 7; fi4++ {
        if false {
            p0 -= p0
        } else {
            v5 := []int{(^(fi4)), 128, (fi4 % ((fi4) | 1))}
            _ = v5
        }
        if (898 == 758) {
            for fi6 := 0; fi6 < 13; fi6++ {
                v7 := GenS1{F0: 2, F1: 0, F2: 18446744073709551615}
                _ = fi6
                _ = v7
            }
            var v8 bool = (-2.25 == 1.5)
            switch 32767 {
            case 32768:
                if ((-(fi4)) < 9223372036854775807) {
                    for fi9 := 0; fi9 < 18; fi9++ {
                        v10 := []int{fi9}
                        _ = fi9
                        _ = v10
                    }
                }
            case 2147483647:
                v11 := GenS2{F0: 909, F1: 128}
                _ = v11
            }
            _ = v8
        } else {
            if (-1 >= 42) {
                var v12 uint64 = 370
                v13 := []int{fi4, fi4, (-(fi4))}
                var v14 int8 = -1
                _ = v12
                _ = v13
                _ = v14
            }
            p0 = (p0 >> ((18446744073709551615) & 63))
        }
        _ = fi4
    }
    _ = p0
    return 128
}

func main() {
    var v1 uint32 = 42
    var v2 bool = false
    var v3 int32 = -1
    v3 += (v3 >> ((452) & 63))
    v15 := GenS2{F0: 477, F1: 127}
    v16 := GenS2{F0: 1, F1: 127}
    var v17 bool = v2
    v18 := []int{}
    fmt.Println(v3, v1, v1)
    switch -2147483648 {
    case 2147483647:
        switch 4294967295 {
        case 734:
            if len(v18) >= 2 {
                v18 = v18[1:len(v18)]
            }
        case -1:
            v2 = (true || true)
        case 127:
            var v19 uint16 = 32767
            _ = v19
        default:
            if ((((!(("a" + ("foo bar" + "a")) <= "foo bar")) && false) || false) && (816 >= 1)) {
                fmt.Println(v2, v3)
                if false {
                    v20 := GenS1{F0: 177, F1: 1, F2: 2}
                    fmt.Println(v3)
                    _ = v20
                }
                v1 = v1
            } else {
                fmt.Println(v1, v3, v3)
                v2 = false
            }
        }
    case 4294967295:
        v21 := GenS2{F0: 255, F1: 2}
        _ = v21
    default:
        v22 := GenS3{F0: 18446744073709551615, F1: 127}
        _ = v22
    }
    v17 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
