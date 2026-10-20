package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16 }
func genFn0(p0 int64, p1 string, p2 uint32) bool {
    var v7 int32 = 1
    var v8 bool = true
    v9 := GenS1{F0: 127}
    for fi10 := 0; fi10 < 9; fi10++ {
        v11 := []int{fi10}
        _ = fi10
        _ = v11
    }
    if v8 {
        for fi12 := 0; fi12 < 19; fi12++ {
            var v13 int32 = v7
            var v14 int16 = -32768
            v14 += v14
            _ = fi12
            _ = v13
            _ = v14
        }
        for fi15 := 0; fi15 < 4; fi15++ {
            for fi16 := 0; fi16 < 8; fi16++ {
                v7 = v7
                _ = fi16
            }
            var v17 float32 = (-2.25 * 3.14159)
            v7 -= int32(p2)
            _ = fi15
            _ = v17
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return v8
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := make(map[int]int)
    var v3 int64 = 9223372036854775807
    var v4 float64 = 0.0
    v5 := []int{42}
    v6 := make(map[uint32]int)
    var v18 int64 = (v3 >> ((1) & 63))
    if (!false) {
        v1 = (v1 % ((v1) | 1))
        if ("" < "a") {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            if (275 <= v1) {
                v18 = 1
                v1 -= v1
                if true {
                    v19 := v2[0]
                    for fi20 := 0; fi20 < 9; fi20++ {
                        v2[-1] = (fi20 - (v19 ^ (v19 % ((1) | 1))))
                        _ = fi20
                    }
                    _ = v19
                } else {
                    v3 -= v3
                }
            } else {
                if (0 < 4294967295) {
                    v1 = 0
                    if ((-1 < -1) || false) {
                        fmt.Println(v1)
                    }
                    v21 := GenS1{F0: 127}
                    _ = v21
                }
            }
        } else {
            if (!true) {
                v3 = (v18 ^ 440)
                v22 := v2[128]
                fmt.Println(v1)
                _ = v22
            }
            v23 := GenS1{F0: 42}
            _ = v23
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    } else {
        switch 32768 {
        case 4294967295:
            var v24 uint64 = v1
            _ = v24
        case 32767:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        default:
            delete(v6, 1)
        }
        var v25 float64 = 1.5
        _ = v25
    }
    if ((42 < 1) && ((false || (0.1 != ((v4 - (v4 - float64((1e10 - 0.1)))) - v4))) && ((2 >= 27) && (!true)))) {
        v3 = v3
    } else {
        var v26 string = ("\n\t" + "\n\t")
        v26 = (("hello" + "\n\t") + (v26 + v26))
        _ = v26
    }
    v3 = 42
    for fi27 := 0; fi27 < 14; fi27++ {
        var v28 int64 = (v3 >> ((2) & 63))
        fmt.Println(v3, v18)
        _ = fi27
        _ = v28
    }
    for fi29 := 0; fi29 < 4; fi29++ {
        v2[fi29] = -1
        if ((true || false) && (int64(fi29) == -232)) {
            v30 := GenS2{F0: 289}
            var v31 uint64 = v1
            var v32 int = (fi29 - (fi29 >> ((1) & 63)))
            _ = v30
            _ = v31
            _ = v32
        } else {
            switch fi29 {
            case 127:
                switch -66 {
                case -9223372036854775808:
                    if (125 != 42) {
                        v3 -= (v3 % ((v3) | 1))
                        v33 := GenS1{F0: -128}
                        _ = v33
                    }
                case 1:
                    fmt.Println(v3, fi29)
                }
            }
            if false {
                if false {
                    var v34 uint8 = 127
                    _ = v34
                }
                v1 = v1
                v35 := GenS1{F0: 68}
                _ = v35
            }
            var v36 bool = true
            _ = v36
        }
        _ = fi29
    }
    var v37 string = "foo bar"
    v38 := GenS1{F0: 1}
    if len(v5) > 0 {
        v5[0] = -927
    }
    if (((42 < 0) && false) && false) {
        v39 := v2[-1]
        v40 := GenS2{F0: 1000}
        _ = v39
        _ = v40
    }
    v4 = 3.14159
    v41 := GenS1{F0: 0}
    v37 = v37
    if (false || (true || true)) {
        v18 += (v18 | v18)
        fmt.Println(v4)
        if false {
            var v42 int32 = -1
            delete(v2, -443)
            var v43 uint = 4294967295
            _ = v42
            _ = v43
        } else {
            fmt.Println(v1)
            fmt.Println(v4, v4, v18)
            v2[-792] = 32767
        }
    }
    if len(v5) > 0 {
        v5[0] = 1
    }
    for fi44 := 0; fi44 < 14; fi44++ {
        var v45 int8 = 1
        _ = fi44
        _ = v45
    }
    delete(v2, 507)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
