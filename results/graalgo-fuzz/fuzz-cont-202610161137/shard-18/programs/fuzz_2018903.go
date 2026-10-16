package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int16 }
type GenS2 struct { F0 uint32; F1 int64; F2 uint }
type GenS3 struct { F0 uint8; F1 int64; F2 bool }
type GenS4 struct { F0 uint8; F1 uint64; F2 int8 }
func genFn0() uint64 {
    var v6 int8 = 13
    for fi7 := 0; fi7 < 4; fi7++ {
        v6 = v6
        var v8 uint8 = 85
        v6 -= (v6 << ((65535) & 63))
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 0, F1: 169, F2: 1}
    if false {
        for fi10 := 0; fi10 < 13; fi10++ {
            v11 := []int{-1, fi10, -584, 523}
            for fi12 := 0; fi12 < 9; fi12++ {
                var v13 int8 = v6
                _ = fi12
                _ = v13
            }
            for fi14 := 0; fi14 < 19; fi14++ {
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                if ((((0.0 - 0.1) - 1.5) + 1e10) != ((0.1 / 1e10) - (0.1 * -2.25))) {
                    if len(v11) > 0 {
                        v11[0] = (fi14 - fi14)
                    }
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                    v6 -= (v6 ^ -97)
                } else {
                    var v15 float32 = 1e10
                    _ = v15
                }
                _ = fi14
            }
            _ = fi10
            _ = v11
        }
        var v16 float32 = -2.25
        var v17 uint8 = 127
        _ = v16
        _ = v17
    } else {
        v18 := []int{-494}
        var v19 int8 = (v6 * (v6 & v6))
        _ = v18
        _ = v19
    }
    _ = v6
    _ = v9
    return 42
}

func main() {
    var v1 bool = false
    var v2 int64 = 9223372036854775807
    v3 := make(map[int]bool)
    var v4 uint32 = 2
    var v5 int64 = v2
    var v20 float32 = (float32(float64(0.0)) * 1.5)
    if ((-424 > int32(v2)) || false) {
        v21 := []int{986, -45, 9223372036854775807}
        v22 := GenS2{F0: (v4 << ((18446744073709551615) & 63)), F1: (v2 - 858), F2: 774}
        for fi23 := 0; fi23 < 3; fi23++ {
            v1 = false
            _ = fi23
        }
        _ = v21
        _ = v22
    } else {
        if true {
            v2 = v2
            var v24 bool = (0 > 0)
            for fi25 := 0; fi25 < 7; fi25++ {
                var v26 float32 = (v20 + 0.1)
                _ = fi25
                _ = v26
            }
            _ = v24
        } else {
            v27 := GenS3{F0: 255, F1: (v2 % (((v2 ^ 0)) | 1)), F2: v1}
            v28 := []int{-1, 332, 127, -73}
            if ((32767 <= 1) || (!("" > ("hello" + "hello")))) {
                var v29 string = "hello"
                fmt.Println(v5, v2, v29)
                _ = v29
            }
            _ = v27
            _ = v28
        }
    }
    delete(v3, 542)
    delete(v3, 9223372036854775807)
    v30 := []int{127, int(v2), -1, -2147483648}
    if v1 {
        var v31 int = 32767
        switch ("a" + "\n\t") {
        case "a":
            var v32 string = ""
            _ = v32
        case "foo bar":
            var v33 float32 = v20
            _ = v33
        case "":
            for fi34 := 0; fi34 < 4; fi34++ {
                fmt.Println(v31, v2)
                switch (v31 / ((4294967295) | 1)) {
                case -9223372036854775808:
                    if true {
                        var v35 int16 = 1000
                        _ = v35
                    }
                default:
                    for fi36 := 0; fi36 < 7; fi36++ {
                        v31 -= 552
                        v37 := GenS3{F0: 127, F1: -9223372036854775808, F2: v1}
                        v38 := v3[(fi34 + (v31 | fi36))]
                        _ = fi36
                        _ = v37
                        _ = v38
                    }
                }
                _ = fi34
            }
        default:
            var v39 string = "hello"
            _ = v39
        }
        _ = v31
    }
    v30 = append(v30, 32767)
    v2 = v5
    v40 := GenS2{F0: v4, F1: -299, F2: 50}
    for fi41 := 0; fi41 < 20; fi41++ {
        if len(v30) >= 2 {
            v30 = v30[1:len(v30)]
        }
        _ = fi41
    }
    fmt.Println(v20, v5, v2)
    v2 += (v2 ^ 1)
    var v42 uint8 = 0
    if (!(473 > -674)) {
        var v43 float32 = v20
        var v44 string = ""
        v45 := GenS4{F0: (v42 + 42), F1: genFn0(), F2: 100}
        _ = v43
        _ = v44
        _ = v45
    } else {
        v46 := GenS3{F0: 211, F1: -1, F2: (774 > 2)}
        for fi47 := 0; fi47 < 6; fi47++ {
            if len(v30) > 0 {
                v30[0] = 32768
            }
            var v48 uint = 18446744073709551615
            var v49 uint64 = uint64(v42)
            _ = fi47
            _ = v48
            _ = v49
        }
        _ = v46
    }
    v1 = v1
    fmt.Println(v20, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v30
    _ = v4
    _ = v40
    _ = v42
    _ = v5
}
