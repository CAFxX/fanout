package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 string; F2 int8 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 int) uint32 {
    p0 = (p0 | (p0 % (((p0 / (((p0 & 127)) | 1))) | 1)))
    var v7 int8 = 127
    switch p0 {
    case 2147483647:
        var v8 bool = (-32768 >= 689)
        _ = v8
    case 32768:
        var v9 uint = 309
        _ = v9
    case -9223372036854775808:
        var v10 int = p0
        _ = v10
    }
    var v11 int16 = 1000
    v12 := []int{(p0 + (p0 % ((680) | 1))), -2147483648}
    _ = p0
    _ = v11
    _ = v12
    _ = v7
    return 2147483647
}
func genFn1(p0 int64, p1 uint32, p2 int64) bool {
    p2 = 9223372036854775807
    for fi13 := 0; fi13 < 1; fi13++ {
        if false {
            var v14 int32 = 2147483647
            v15 := GenS1{F0: p1}
            fmt.Println(p1)
            _ = v14
            _ = v15
        } else {
            for fi16 := 0; fi16 < 17; fi16++ {
                var v17 uint32 = (p1 << ((220) & 63))
                if ((537 > 1) && false) {
                    fmt.Println(p2)
                    var v18 string = "a"
                    var v19 int8 = 1
                    v20 := GenS1{F0: 534}
                    fmt.Println(v18, p0)
                    _ = v18
                    _ = v19
                    _ = v20
                }
                var v21 uint64 = 514
                _ = fi16
                _ = v17
                _ = v21
            }
        }
        _ = fi13
    }
    p1 = (p1 << ((1) & 63))
    _ = p0
    _ = p1
    _ = p2
    return true
}

func main() {
    var v1 uint8 = 255
    var v2 int = -376
    v3 := make(map[int]uint32)
    var v4 int8 = 42
    var v5 uint8 = 1
    var v6 bool = true
    var v22 int8 = v4
    var v23 uint32 = 299
    var v24 uint = 2
    v25 := []int{v2, v2, 2147483647}
    v2 = 706
    for fi26 := 0; fi26 < 19; fi26++ {
        for fi27 := 0; fi27 < 11; fi27++ {
            var v28 int16 = 0
            v22 = (v22 - (-(v22)))
            var v29 uint = v24
            _ = fi27
            _ = v28
            _ = v29
        }
        if (434 == 173) {
            v22 = 0
            v30 := GenS1{F0: genFn0((fi26 >> ((v24) & 63)))}
            _ = v30
        } else {
            var v31 int8 = -97
            _ = v31
        }
        _ = fi26
    }
    if (("hello" < "\n\t") && v6) {
        v4 += v4
    } else {
        v32 := GenS1{F0: v23}
        var v33 int64 = 57
        _ = v32
        _ = v33
    }
    for fi34 := 0; fi34 < 1; fi34++ {
        v35 := v3[fi34]
        for fi36 := 0; fi36 < 9; fi36++ {
            if (float64((float32(1.5) * (1e10 + -2.25))) >= 1.5) {
                if ((v6 || false) && (v22 == v4)) {
                    var v37 uint8 = v1
                    v6 = true
                    v38 := GenS2{F0: 1000, F1: "foo bar", F2: v22}
                    _ = v37
                    _ = v38
                }
                v39 := GenS3{F0: 22}
                var v40 string = "foo bar"
                _ = v39
                _ = v40
            } else {
                v41 := v3[9223372036854775807]
                var v42 float32 = 1e10
                fmt.Println(v23)
                _ = v41
                _ = v42
            }
            v43 := GenS2{F0: 42, F1: ("\n\t" + ("a" + "")), F2: (v22 & (v4 + v4))}
            _ = fi36
            _ = v43
        }
        v23 += (v35 | (v23 + (v23 >> ((v24) & 63))))
        _ = fi34
        _ = v35
    }
    v22 = v22
    v5 -= 255
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
