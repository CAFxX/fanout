package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 string; F1 string; F2 uint64 }
func genFn0() uint32 {
    v7 := GenS1{F0: 42}
    var v8 int32 = 42
    v9 := []int{-1}
    var v10 string = ""
    v10 = "a"
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 0
}
func genFn1() uint32 {
    var v11 bool = (((0.1 + 1.5) / 1.5) <= -2.25)
    for fi12 := 0; fi12 < 17; fi12++ {
        var v13 uint16 = 65535
        _ = fi12
        _ = v13
    }
    fmt.Println(v11)
    _ = v11
    return genFn0()
}

func main() {
    v1 := []uint32{850, 2147483648, 2147483647, 788}
    var v2 int64 = -1
    var v3 uint16 = 463
    v4 := make(map[uint32]bool)
    var v5 int = 9223372036854775807
    var v6 uint = 2
    if (1000 > 1) {
        v3 = 255
        switch ("a" + "") {
        case "":
            v14 := v4[2147483648]
            _ = v14
        case "foo bar":
            var v15 float32 = ((1e10 - 1e10) * (float32((float64(0.0) - 3.14159)) - ((((0.1 * 1e10) + 1.5) * (1.5 * 0.1)) / 0.1)))
            _ = v15
        case "a":
            switch "" {
            case "\n\t":
                if (2 != uint64(v3)) {
                    v4[uint32(v6)] = ("" < "foo bar")
                    var v16 float64 = (3.14159 - ((((0.1 / ((0.0 + 0.0) - 3.14159)) * 3.14159) / (3.14159 * (1e10 + float64(1.5)))) / 1.0))
                    _ = v16
                }
            case "hello":
                if (true || ((!false) || true)) {
                    for fi17 := 0; fi17 < 20; fi17++ {
                        fmt.Println(v2)
                        fmt.Println(v6)
                        _ = fi17
                    }
                } else {
                    var v18 int8 = -1
                    v3 = 1
                    v3 -= (v3 % ((v3) | 1))
                    _ = v18
                }
            case "a":
                if ("a" == ("" + "hello")) {
                    delete(v4, 2147483648)
                } else {
                    var v19 float32 = 3.14159
                    _ = v19
                }
            }
        }
        var v20 uint64 = 2
        _ = v20
    } else {
        fmt.Println(v5, v2, v2)
        var v21 int64 = v2
        v22 := GenS1{F0: 3}
        _ = v21
        _ = v22
    }
    if true {
        v23 := GenS1{F0: 42}
        if true {
            switch 586 {
            case 55:
                var v24 bool = (255 > 124)
                _ = v24
            case -272:
                var v25 uint32 = 2147483648
                _ = v25
            case 128:
                if len(v1) > 0 {
                    v1[0] = genFn1()
                }
            }
            if true {
                if false {
                    if true {
                        v2 = (-(v2))
                        var v26 int32 = 1
                        var v27 int8 = 1
                        _ = v26
                        _ = v27
                    }
                    v2 += 9223372036854775807
                }
                for fi28 := 0; fi28 < 2; fi28++ {
                    v29 := GenS2{F0: 1, F1: 4294967295}
                    _ = fi28
                    _ = v29
                }
                for fi30 := 0; fi30 < 10; fi30++ {
                    v3 -= v3
                    _ = fi30
                }
            }
        } else {
            var v31 uint16 = (v3 << ((42) & 63))
            var v32 uint = v6
            _ = v31
            _ = v32
        }
        _ = v23
    }
    for fi33 := 0; fi33 < 20; fi33++ {
        for fi34 := 0; fi34 < 8; fi34++ {
            var v35 float32 = ((1.5 / 3.14159) / 1.0)
            v35 = 1e10
            v1 = append(v1, 2147483647)
            v3 = (v3 * (v3 * 42))
            if ((false || false) || (v2 <= (v2 % ((v2) | 1)))) {
                var v36 float64 = 0.0
                _ = v36
            } else {
                if (true || ((!(v3 == v3)) && (0.0 <= 1e10))) {
                    v37 := GenS2{F0: int32(v3), F1: 18446744073709551615}
                    v5 = int(v3)
                    _ = v37
                }
                v5 = (fi33 + 32768)
            }
            _ = fi34
            _ = v35
        }
        v3 = (^(v3))
        _ = fi33
    }
    if (true && (!(v2 == -1))) {
        if false {
            var v38 float32 = 1.5
            _ = v38
        }
        v2 = v2
        var v39 int8 = 1
        _ = v39
    }
    fmt.Println(v5)
    fmt.Println(v2, v3, v2)
    switch (-(v5)) {
    case 32768:
        var v40 int8 = 100
        _ = v40
    case 42:
        var v41 int16 = 1
        _ = v41
    case -25:
        v5 += (v5 ^ 128)
    default:
        var v42 string = "hello"
        _ = v42
    }
    fmt.Println(v3, v5)
    for fi43 := 0; fi43 < 4; fi43++ {
        v44 := GenS3{F0: "\n\t", F1: "foo bar", F2: 2}
        _ = fi43
        _ = v44
    }
    var v45 int32 = 42
    var v46 uint32 = 2
    v47 := GenS2{F0: v45, F1: 18446744073709551615}
    v6 += v6
    for fi48 := 0; fi48 < 19; fi48++ {
        var v49 uint32 = v46
        v50 := GenS2{F0: v45, F1: 939}
        _ = fi48
        _ = v49
        _ = v50
    }
    var v51 bool = (-793 == int16(v45))
    var v52 int64 = v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v51
    _ = v52
    _ = v6
}
