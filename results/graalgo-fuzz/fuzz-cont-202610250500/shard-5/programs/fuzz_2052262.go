package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 bool; F1 int16 }
func genFn0(p0 uint64, p1 uint64, p2 uint) uint64 {
    var v4 int8 = 49
    v4 = (-(v4))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return p1
}
func genFn1() int16 {
    if false {
        for fi5 := 0; fi5 < 20; fi5++ {
            var v6 bool = true
            var v7 bool = v6
            _ = fi5
            _ = v6
            _ = v7
        }
        v8 := []int{4294967295}
        v9 := GenS1{F0: -894}
        _ = v8
        _ = v9
    } else {
        var v10 int64 = -228
        _ = v10
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        for fi12 := 0; fi12 < 18; fi12++ {
            var v13 int64 = 1
            v14 := GenS2{F0: 2147483648}
            _ = fi12
            _ = v13
            _ = v14
        }
        var v15 int = (fi11 + (fi11 >> ((2) & 63)))
        v16 := []int{v15, -1, (-(fi11)), (v15 >> ((126) & 63))}
        _ = fi11
        _ = v15
        _ = v16
    }
    return -1
}

func main() {
    v1 := []uint32{0, 42, 766}
    var v2 uint16 = 0
    var v3 string = (("foo bar" + "") + ("a" + "hello"))
    var v17 bool = false
    for fi18 := 0; fi18 < 11; fi18++ {
        switch (v3 + v3) {
        case "a":
            switch 128 {
            case 124:
                for fi19 := 0; fi19 < 18; fi19++ {
                    var v20 float32 = (3.14159 + 1e10)
                    if ((float64(v20) < -2.25) || (0.0 < 0.1)) {
                        fmt.Println(v2, fi18)
                        var v21 uint32 = 377
                        var v22 int32 = -2147483648
                        switch (v3 + ("\n\t" + (v3 + ("" + "a")))) {
                        case "foo bar":
                            if (!true) {
                                v3 = (("" + (v3 + "foo bar")) + (((v3 + "a") + ("foo bar" + "hello")) + v3))
                            }
                        default:
                            var v23 string = ""
                            _ = v23
                        }
                        v24 := GenS3{F0: 54, F1: v2}
                        _ = v21
                        _ = v22
                        _ = v24
                    } else {
                        v20 -= v20
                        v3 = v3
                    }
                    v17 = (((true && (1 > 418)) && (-1 != 1)) || v17)
                    v2 = 460
                    _ = fi19
                    _ = v20
                }
            case 4294967295:
                _ = len(v1)
            }
        case "":
            v17 = (!(true || (!(v17 || (v3 >= v3)))))
        }
        if (-712 < -2147483648) {
            if v17 {
                v25 := GenS3{F0: 164, F1: (v2 ^ (v2 >> ((42) & 63)))}
                v3 = v3
                _ = v25
            }
            v3 = "a"
            v26 := GenS2{F0: 1}
            _ = v26
        } else {
            if ((1 >= -1) && true) {
                v1 = append(v1, 4294967295)
                var v27 uint = 0
                _ = v27
            }
            var v28 uint32 = 152
            _ = v28
        }
        _ = fi18
    }
    switch 32767 {
    case -9223372036854775808:
        _ = len(v1)
    case -648:
        var v29 bool = v17
        _ = v29
    case 127:
        for fi30 := 0; fi30 < 1; fi30++ {
            fmt.Println(fi30, v3)
            if (v17 && true) {
                v2 = v2
                var v31 string = ""
                v3 = ("a" + (v3 + ("" + "")))
                _ = v31
            } else {
                var v32 uint16 = v2
                v17 = false
                _ = v32
            }
            _ = fi30
        }
    }
    fmt.Println(v2, v2, v2)
    for fi33 := 0; fi33 < 13; fi33++ {
        v1 = append(v1, 114)
        for fi34 := 0; fi34 < 14; fi34++ {
            if false {
                var v35 uint64 = 2
                _ = v35
            }
            fmt.Println(fi34, v2, v3)
            _ = fi34
        }
        _ = fi33
    }
    var v36 string = v3
    var v37 int = 9223372036854775807
    fmt.Println(v2)
    if (2147483647 > 1) {
        var v38 uint = 7
        _ = v38
    } else {
        if (412 != int32(v37)) {
            v39 := GenS1{F0: (v37 + v37)}
            _ = len(v1)
            _ = v39
        } else {
            switch v36 {
            case "foo bar":
                switch "hello" {
                case "\n\t":
                    v40 := GenS2{F0: 1}
                    _ = v40
                case "a":
                    var v41 uint = 4294967295
                    _ = v41
                case "":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                default:
                    var v42 int8 = -128
                    _ = v42
                }
            case "a":
                v1 = append(v1, 0)
            case "hello":
                _ = len(v1)
            default:
                if (true || v17) {
                    fmt.Println(v3, v2)
                    v43 := GenS4{F0: true, F1: genFn1()}
                    _ = len(v1)
                    _ = v43
                }
            }
        }
        var v44 bool = v17
        _ = v44
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v36
    _ = v37
}
