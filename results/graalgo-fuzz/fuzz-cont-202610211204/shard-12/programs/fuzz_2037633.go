package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 string; F1 uint }
func genFn0(p0 int) bool {
    for fi6 := 0; fi6 < 17; fi6++ {
        v7 := []int{(p0 << ((5) & 63)), (fi6 | fi6)}
        v7 = append(v7, 42)
        _ = fi6
        _ = v7
    }
    var v8 int = p0
    v9 := GenS1{F0: 4294967295, F1: 269}
    fmt.Println(p0, v8)
    _ = p0
    _ = v8
    _ = v9
    return true
}

func main() {
    var v1 uint = 639
    var v2 int8 = -35
    var v3 int32 = 0
    var v4 float32 = 1.5
    v5 := make(map[uint32]string)
    var v10 uint = 65535
    v2 += (v2 >> ((v1) & 63))
    v11 := GenS1{F0: 9223372036854775807, F1: 0}
    v1 -= (v1 | 625)
    var v12 uint32 = 2147483648
    switch 128 {
    case 127:
        var v13 int16 = 930
        _ = v13
    default:
        v4 = 0.1
    }
    var v14 uint16 = 255
    fmt.Println(v10, v2, v12)
    for fi15 := 0; fi15 < 2; fi15++ {
        v16 := GenS1{F0: fi15, F1: -1}
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 17; fi17++ {
        var v18 int64 = 9223372036854775807
        switch 2147483647 {
        case 9223372036854775807:
            delete(v5, 251)
        }
        switch 552 {
        case 127:
            fmt.Println(v14, v2, v1)
        case 1:
            if ((v18 ^ (v18 * v18)) > v18) {
                switch -335 {
                case 356:
                    for fi19 := 0; fi19 < 9; fi19++ {
                        v12 += (v12 >> ((255) & 63))
                        _ = fi19
                    }
                case 42:
                    var v20 string = "hello"
                    _ = v20
                case -233:
                    var v21 int8 = (v2 | 1)
                    _ = v21
                }
                for fi22 := 0; fi22 < 15; fi22++ {
                    v5[v12] = ("\n\t" + "a")
                    v23 := GenS1{F0: (fi17 ^ int(v2)), F1: 32767}
                    v24 := GenS1{F0: fi17, F1: -924}
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                if (((v14 & 1) >= (v14 / ((v14) | 1))) && (false && false)) {
                    v25 := []int{-2147483648, fi17}
                    _ = v25
                } else {
                    v12 += 2147483647
                    for fi26 := 0; fi26 < 19; fi26++ {
                        var v27 float32 = -2.25
                        v28 := v5[957]
                        var v29 uint = 713
                        _ = fi26
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                    v30 := GenS1{F0: (fi17 ^ 32768), F1: -1}
                    _ = v30
                }
            } else {
                v18 = 9223372036854775807
            }
        case 32768:
            v31 := GenS1{F0: fi17, F1: 1000}
            _ = v31
        }
        _ = fi17
        _ = v18
    }
    v12 = (v12 / ((v12) | 1))
    v32 := GenS1{F0: -60, F1: -1}
    if ((-(v3)) != v3) {
        v33 := GenS1{F0: 926, F1: 32767}
        _ = v33
    } else {
        var v34 int = -493
        fmt.Println(v10)
        _ = v34
    }
    var v35 uint16 = (v14 & (v14 & (^(v14))))
    var v36 uint8 = 255
    for fi37 := 0; fi37 < 11; fi37++ {
        v38 := GenS2{F0: 4294967295, F1: (v36 & (^(v36)))}
        _ = fi37
        _ = v38
    }
    v39 := GenS3{F0: "hello", F1: (v1 >> ((v1) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v32
    _ = v35
    _ = v36
    _ = v39
    _ = v4
    _ = v5
}
