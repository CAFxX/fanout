package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8; F1 bool; F2 int32 }
func genFn0() uint32 {
    var v4 int64 = 913
    var v5 int64 = v4
    _ = v4
    _ = v5
    return 920
}
func genFn1(p0 int64, p1 uint, p2 uint16) int32 {
    fmt.Println(p1)
    if false {
        p1 = (p1 >> ((p1) & 63))
        v6 := []int{127, 128, -751}
        _ = v6
    } else {
        var v7 uint64 = 18446744073709551615
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return -404
}

func main() {
    var v1 float32 = 3.14159
    var v2 uint32 = 4294967295
    v3 := make(map[uint32]string)
    v1 -= float32(1e10)
    if (v1 == (v1 - v1)) {
        switch ("hello" + "\n\t") {
        case "foo bar":
            v2 -= (v2 & 0)
        default:
            _ = len(v3)
        }
        var v8 uint64 = 18446744073709551615
        v3[(v2 - v2)] = ""
        _ = v8
    }
    switch 759 {
    case 291:
        if false {
            _ = len(v3)
            var v9 string = "a"
            _ = v9
        } else {
            var v10 float32 = v1
            for fi11 := 0; fi11 < 5; fi11++ {
                v12 := GenS1{F0: 0, F1: -307, F2: 0}
                _ = fi11
                _ = v12
            }
            var v13 uint8 = 0
            _ = v10
            _ = v13
        }
    }
    v1 = float32(-2.25)
    for fi14 := 0; fi14 < 2; fi14++ {
        for fi15 := 0; fi15 < 13; fi15++ {
            if true {
                fmt.Println(v1, v1, fi15)
                delete(v3, v2)
                switch "a" {
                case "a":
                    if (32767 != 42) {
                        v1 -= (v1 * 0.0)
                        var v16 bool = false
                        switch 32768 {
                        case 4294967295:
                            var v17 string = "hello"
                            _ = v17
                        case 32767:
                            var v18 bool = (genFn0() != (v2 / ((genFn0()) | 1)))
                            _ = v18
                        case -495:
                            var v19 uint16 = 119
                            _ = v19
                        default:
                            var v20 float64 = ((3.14159 - (0.1 * (1e10 + 3.14159))) * 1.5)
                            _ = v20
                        }
                        v21 := GenS2{F0: 184}
                        v22 := []int{(fi15 & fi15), fi14, (fi15 << ((18446744073709551615) & 63))}
                        v1 = -2.25
                        _ = v16
                        _ = v21
                        _ = v22
                    } else {
                        v23 := GenS1{F0: 606, F1: -244, F2: 127}
                        v1 = ((float32(float64(1e10)) / (3.14159 - v1)) / (((v1 + 3.14159) * v1) + v1))
                        _ = v23
                    }
                    v24 := []int{}
                    _ = v24
                case "\n\t":
                    switch (fi14 ^ (-(fi15))) {
                    case 2147483647:
                        if true {
                            var v25 int8 = 41
                            v25 -= (v25 * 25)
                            var v26 uint16 = 32767
                            _ = v25
                            _ = v26
                        } else {
                            var v27 uint32 = v2
                            _ = len(v3)
                            _ = v27
                        }
                    default:
                        var v28 int8 = 85
                        var v29 int32 = 2147483647
                        v30 := GenS1{F0: (v29 * (v29 & (v29 >> ((698) & 63)))), F1: 0, F2: v28}
                        _ = v28
                        _ = v29
                        _ = v30
                    }
                case "":
                    v1 -= v1
                default:
                    v2 = genFn0()
                }
            }
            if false {
                if (true || true) {
                    v2 = genFn0()
                    var v31 uint64 = 2
                    _ = v31
                }
                fmt.Println(v1, fi14)
                switch ("foo bar" + ("foo bar" + "foo bar")) {
                case "\n\t":
                    var v32 int16 = -769
                    _ = v32
                case "":
                    fmt.Println(fi14, fi14, v2)
                default:
                    v33 := []int{}
                    _ = v33
                }
                if true {
                    var v34 int16 = 42
                    _ = v34
                }
            }
            _ = fi15
        }
        var v35 uint16 = 460
        _ = fi14
        _ = v35
    }
    fmt.Println(v1)
    if (18446744073709551615 < 1) {
        var v36 int16 = -537
        _ = v36
    }
    v37 := GenS3{F0: 0, F1: false, F2: -2147483648}
    var v38 int = -886
    v38 -= (v38 & (v38 & 2147483647))
    for fi39 := 0; fi39 < 7; fi39++ {
        var v40 uint = 472
        v41 := GenS1{F0: -1, F1: 32767, F2: 42}
        _ = fi39
        _ = v40
        _ = v41
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
}
