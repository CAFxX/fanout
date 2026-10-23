package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int; F1 int8 }
type GenS5 struct { F0 uint8; F1 uint16 }
type GenS6 struct { F0 uint32; F1 string }
func genFn0() uint32 {
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 uint16 = 65535
        v9 := GenS1{F0: 255, F1: v8}
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := []int{-436, -1}
    var v11 int32 = -1
    for fi12 := 0; fi12 < 17; fi12++ {
        for fi13 := 0; fi13 < 19; fi13++ {
            switch "" {
            case "\n\t":
                v11 = 42
                var v14 int8 = 100
                if (633 != 2147483647) {
                    v15 := GenS2{F0: ((v11 | 264) >= 42)}
                    _ = v15
                } else {
                    var v16 int16 = -42
                    switch 32767 {
                    case -9223372036854775808:
                        v14 -= (-(v14))
                    case -1:
                        fmt.Println(v11, v11, v16)
                    case 2147483647:
                        v17 := GenS3{F0: 164}
                        _ = v17
                    default:
                        v18 := GenS4{F0: (^(fi13)), F1: v14}
                        _ = v18
                    }
                    var v19 float32 = 1e10
                    _ = v16
                    _ = v19
                }
                _ = v14
            case "foo bar":
                v20 := GenS3{F0: 32767}
                _ = v20
            default:
                fmt.Println(fi12, v11, v11)
            }
            v11 = v11
            if true {
                switch -557 {
                case 128:
                    var v21 int16 = 42
                    _ = v21
                }
                v22 := GenS5{F0: 128, F1: 65535}
                v11 -= (v11 % ((v11) | 1))
                v11 -= (v11 % (((-(v11))) | 1))
                var v23 float32 = float32(-2.25)
                _ = v22
                _ = v23
            } else {
                v11 = (v11 | v11)
                v24 := GenS2{F0: false}
                v11 = v11
                if ((true || (0.1 <= (1.5 * 0.1))) || false) {
                    v25 := GenS5{F0: 0, F1: 0}
                    var v26 float64 = float64(((float32(-2.25) * 0.0) - ((0.1 * float32(0.1)) / 1e10)))
                    _ = v25
                    _ = v26
                } else {
                    var v27 int8 = 127
                    _ = v27
                }
                var v28 int16 = -1
                _ = v24
                _ = v28
            }
            v11 = (v11 + v11)
            _ = fi13
        }
        v11 += (v11 << ((255) & 63))
        var v29 string = "foo bar"
        _ = fi12
        _ = v29
    }
    _ = v10
    _ = v11
    return 2
}
func genFn1(p0 int16, p1 int64, p2 string) uint16 {
    p2 = "\n\t"
    p0 = p0
    for fi30 := 0; fi30 < 10; fi30++ {
        var v31 uint = 288
        if false {
            var v32 uint16 = 255
            _ = v32
        } else {
            for fi33 := 0; fi33 < 18; fi33++ {
                var v34 uint16 = 2
                _ = fi33
                _ = v34
            }
        }
        v35 := []int{(fi30 + fi30)}
        _ = fi30
        _ = v31
        _ = v35
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    v1 := make(map[int]string)
    var v2 float32 = (((1e10 * 1e10) + 1e10) / 1e10)
    var v3 bool = (false && false)
    v4 := make(map[string]int)
    v5 := make(map[string]uint32)
    var v6 int8 = -1
    v36 := []int{127, -1, 32767}
    v37 := GenS6{F0: genFn0(), F1: ""}
    v6 -= v6
    var v38 float64 = 3.14159
    if (!v3) {
        v2 += ((((0.1 - 1e10) + 0.1) + ((((v2 * v2) * -2.25) - (v2 + (0.1 * v2))) + v2)) / 0.1)
    }
    fmt.Println(v2, v2, v2)
    var v39 int = 9223372036854775807
    var v40 uint16 = 65535
    if v3 {
        var v41 int32 = 533
        _ = v41
    } else {
        for fi42 := 0; fi42 < 19; fi42++ {
            var v43 uint = 255
            fmt.Println(v2, v40)
            v43 -= 65535
            _ = fi42
            _ = v43
        }
    }
    switch "\n\t" {
    case "hello":
        var v44 uint8 = 2
        _ = v44
    case "foo bar":
        var v45 string = ("hello" + ("" + "hello"))
        _ = v45
    default:
        if (!(v3 && ((v38 + v38) > v38))) {
            if true {
                var v46 int16 = 0
                var v47 uint64 = 0
                var v48 int32 = 143
                _ = v46
                _ = v47
                _ = v48
            } else {
                var v49 int = v39
                v6 -= v6
                for fi50 := 0; fi50 < 9; fi50++ {
                    v51 := v5["hello"]
                    _ = fi50
                    _ = v51
                }
                _ = v49
            }
            v52 := GenS2{F0: false}
            _ = v52
        } else {
            if (!v3) {
                switch (("a" + ((("hello" + ("" + "hello")) + "foo bar") + ("hello" + (("a" + "\n\t") + (("hello" + "\n\t") + "foo bar"))))) + "\n\t") {
                case "a":
                    v39 -= v39
                default:
                    var v53 int8 = (^(v6))
                    _ = v53
                }
            }
        }
    }
    v2 = v2
    v54 := GenS4{F0: (v39 | 9223372036854775807), F1: -18}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v54
    _ = v6
}
