package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 int8; F2 int }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 int16 }
func genFn0(p0 uint, p1 int, p2 int16) int {
    var v7 float64 = 1.5
    var v8 uint32 = 65535
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return p1
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int64 = 899
    var v3 int32 = -100
    var v4 int16 = 32767
    var v5 string = ""
    var v6 uint16 = 395
    switch 128 {
    case 264:
        v9 := GenS1{F0: (v2 ^ v2)}
        _ = v9
    case 872:
        v10 := []int{-596, 32767, 128, 2147483647}
        _ = v10
    case 42:
        v3 = (v3 % ((v3) | 1))
    }
    if (true || (!(v6 <= v6))) {
        v5 = (((v5 + (v5 + v5)) + v5) + (v5 + ("\n\t" + "\n\t")))
        var v11 uint16 = v6
        _ = v11
    }
    if true {
        v1[1] = v5
        v12 := []int{-491, 964, 0}
        _ = v12
    } else {
        v5 = ((v5 + "foo bar") + "foo bar")
        switch (v5 + (v5 + "a")) {
        case "foo bar":
            v4 += v4
        case "":
            if (42 == 1) {
                v2 = 9223372036854775807
                v3 -= (v3 - -1)
                var v13 int32 = v3
                _ = v13
            } else {
                var v14 float32 = 3.14159
                var v15 int32 = v3
                if ((v5 + v5) != "a") {
                    v16 := GenS2{F0: (1 != -115), F1: 100, F2: 127}
                    v6 = (v6 ^ (^(v6)))
                    _ = v16
                } else {
                    var v17 uint64 = 4294967295
                    v18 := GenS1{F0: int64(v17)}
                    v19 := GenS3{F0: (151 != 45), F1: v17}
                    _ = v17
                    _ = v18
                    _ = v19
                }
                _ = v14
                _ = v15
            }
        case "hello":
            v4 = 32767
        default:
            var v20 bool = true
            _ = v20
        }
    }
    if false {
        var v21 int16 = v4
        _ = v21
    } else {
        var v22 uint32 = uint32(v6)
        switch -703 {
        case 695:
            if false {
                v5 = (v5 + "hello")
                v23 := GenS2{F0: (("foo bar" + ("hello" + v5)) > (v5 + v5)), F1: 1, F2: 461}
                var v24 uint16 = v6
                _ = v23
                _ = v24
            } else {
                for fi25 := 0; fi25 < 3; fi25++ {
                    var v26 int = (fi25 >> ((147) & 63))
                    fmt.Println(v4)
                    switch genFn0(111, genFn0(274, fi25, (v4 + v4)), v4) {
                    case 4294967295:
                        _ = len(v1)
                    }
                    _ = fi25
                    _ = v26
                }
                var v27 int32 = v3
                v2 = v2
                _ = v27
            }
        default:
            for fi28 := 0; fi28 < 14; fi28++ {
                switch v5 {
                case "foo bar":
                    v29 := []int{fi28, fi28, genFn0(18446744073709551615, fi28, v4), fi28}
                    _ = v29
                case "\n\t":
                    fmt.Println(v3, v6)
                default:
                    v6 = 65535
                }
                var v30 uint32 = 2
                fmt.Println(v3, v3)
                _ = fi28
                _ = v30
            }
        }
        switch (("\n\t" + v5) + v5) {
        case "foo bar":
            v4 -= (-(v4))
        default:
            switch (((v5 + v5) + v5) + ("a" + (v5 + (("\n\t" + "\n\t") + (v5 + v5))))) {
            case "foo bar":
                v31 := []int{4294967295}
                _ = v31
            case "hello":
                var v32 float64 = float64(((float32(0.1) / 1.5) / -2.25))
                _ = v32
            case "":
                delete(v1, (v22 & (v22 >> ((127) & 63))))
            default:
                fmt.Println(v3)
            }
        }
        _ = v22
    }
    if (42 != 1) {
        var v33 string = "a"
        _ = v33
    }
    fmt.Println(v3, v5, v3)
    v4 = (v4 + (v4 | v4))
    var v34 uint32 = 0
    var v35 int16 = (v4 + -1)
    delete(v1, v34)
    var v36 uint = 2
    v37 := []int{genFn0(v36, -1, v35)}
    var v38 int32 = (-(v3))
    v39 := GenS4{F0: v4}
    v37 = append(v37, genFn0(65535, 562, (^(v35))))
    if (32767 > 1) {
        var v40 uint16 = v6
        for fi41 := 0; fi41 < 5; fi41++ {
            v40 = v40
            if len(v37) >= 2 {
                v37 = v37[1:len(v37)]
            }
            _ = fi41
        }
        for fi42 := 0; fi42 < 18; fi42++ {
            v2 -= v2
            delete(v1, 2147483647)
            v43 := GenS2{F0: (true || true), F1: 1, F2: fi42}
            _ = fi42
            _ = v43
        }
        _ = v40
    } else {
        switch "hello" {
        case "a":
            for fi44 := 0; fi44 < 4; fi44++ {
                v5 = (v5 + v5)
                v2 += v2
                var v45 float32 = 1.5
                _ = fi44
                _ = v45
            }
        case "":
            v46 := GenS1{F0: v2}
            _ = v46
        case "\n\t":
            if (2147483647 == v34) {
                fmt.Println(v38, v3)
                _ = len(v37)
                if (!true) {
                    delete(v1, (v34 ^ (v34 - (v34 / ((v34) | 1)))))
                    var v47 uint64 = 1
                    if false {
                        if ((int(v2) < -339) || (1 > -37)) {
                            var v48 int32 = v3
                            _ = v48
                        } else {
                            v49 := GenS2{F0: false, F1: -1, F2: -2147483648}
                            if len(v37) >= 2 {
                                v37 = v37[1:len(v37)]
                            }
                            _ = v49
                        }
                        if (!(!false)) {
                            v5 = "hello"
                        }
                    }
                    _ = v47
                } else {
                    for fi50 := 0; fi50 < 7; fi50++ {
                        _ = len(v37)
                        var v51 uint = (v36 * (v36 & (v36 * (v36 << (((v36 - (v36 + 255))) & 63)))))
                        _ = fi50
                        _ = v51
                    }
                    if false {
                        var v52 int = genFn0(v36, -2147483648, 626)
                        _ = len(v1)
                        _ = v52
                    } else {
                        switch -1 {
                        case -393:
                            delete(v1, 719)
                        }
                    }
                }
            } else {
                delete(v1, uint32(v6))
                for fi53 := 0; fi53 < 8; fi53++ {
                    v37 = append(v37, 32767)
                    _ = fi53
                }
                if false {
                    for fi54 := 0; fi54 < 1; fi54++ {
                        v2 -= (v2 | (v2 % ((v2) | 1)))
                        if true {
                            var v55 int16 = (v35 - v4)
                            _ = v55
                        } else {
                            v34 = v34
                            if len(v37) >= 2 {
                                v37 = v37[1:len(v37)]
                            }
                        }
                        fmt.Println(fi54, fi54)
                        _ = fi54
                    }
                    _ = len(v1)
                    var v56 int16 = (^(v4))
                    _ = v56
                }
            }
        }
        var v57 uint8 = 128
        if true {
            v6 += (v6 & 255)
            var v58 int64 = -61
            _ = v58
        }
        _ = v57
    }
    var v59 int16 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v59
    _ = v6
}
