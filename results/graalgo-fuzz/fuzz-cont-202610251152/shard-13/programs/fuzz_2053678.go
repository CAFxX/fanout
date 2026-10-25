package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 int8; F1 uint64; F2 int64 }

func main() {
    var v1 int = 657
    v2 := make(map[int]int)
    var v3 bool = true
    v4 := make(map[uint32]bool)
    var v5 int16 = 1
    var v6 int8 = -121
    _ = len(v2)
    for fi7 := 0; fi7 < 9; fi7++ {
        v8 := v4[1]
        if (!v3) {
            var v9 uint8 = 27
            _ = v9
        }
        v10 := []int{}
        _ = fi7
        _ = v10
        _ = v8
    }
    var v11 int = v1
    for fi12 := 0; fi12 < 8; fi12++ {
        switch 127 {
        case -2147483648:
            fmt.Println(v5, v3, v6)
        case -1:
            if (4294967295 == 100) {
                var v13 float64 = 3.14159
                var v14 int32 = -2147483648
                _ = v13
                _ = v14
            } else {
                if (false && v3) {
                    v4[2147483648] = ((975 < 65535) && v3)
                    v15 := GenS1{F0: 45}
                    _ = v15
                }
                var v16 uint16 = 186
                _ = v16
            }
        case -9223372036854775808:
            var v17 bool = v3
            _ = v17
        default:
            var v18 string = ((("foo bar" + "\n\t") + "foo bar") + "a")
            _ = v18
        }
        for fi19 := 0; fi19 < 16; fi19++ {
            v3 = ((1 != -1) || (1 == 2))
            switch (v11 % ((fi12) | 1)) {
            case -1:
                var v20 uint64 = 169
                _ = v20
            case 2147483647:
                fmt.Println(v11)
            case -463:
                switch (("" + ("" + "")) + ((("hello" + "a") + "a") + ("" + "\n\t"))) {
                case "a":
                    if ((1 <= -984) || true) {
                        var v21 uint64 = 254
                        _ = v21
                    } else {
                        v22 := GenS1{F0: 2}
                        var v23 float64 = float64((0.0 + ((float32((-2.25 / 1.0)) - 0.0) / float32(1e10))))
                        v3 = v3
                        _ = v22
                        _ = v23
                    }
                case "foo bar":
                    switch (v1 | (fi12 * (-(fi19)))) {
                    case -2147483648:
                        if ((42 <= 4294967295) && v3) {
                            var v24 int32 = int32(v6)
                            if (("\n\t" + ("a" + ("a" + "hello"))) == "foo bar") {
                                v25 := []int{fi19, fi19}
                                v25 = append(v25, -12)
                                v26 := GenS1{F0: 45}
                                if v3 {
                                    var v27 uint16 = 290
                                    _ = v27
                                } else {
                                    v28 := GenS2{F0: ((v24 % ((-2147483648) | 1)) != 1), F1: 913}
                                    if len(v25) >= 2 {
                                        v25 = v25[1:len(v25)]
                                    }
                                    v1 -= 42
                                    _ = v28
                                }
                                v2[fi12] = (v11 * (-(fi19)))
                                _ = v25
                                _ = v26
                            } else {
                                v29 := []int{(-(fi12)), (-(fi12))}
                                v24 = (v24 << ((18446744073709551615) & 63))
                                _ = len(v4)
                                _ = v29
                            }
                            if v3 {
                                var v30 uint16 = 2
                                switch ("a" + "\n\t") {
                                case "\n\t":
                                    var v31 float64 = 0.1
                                    _ = v31
                                default:
                                    v5 -= (v5 + (v5 % ((v5) | 1)))
                                }
                                _ = v30
                            }
                            if (float64(0.0) < (0.1 * float64(3.14159))) {
                                switch "a" {
                                case "":
                                    var v32 int32 = v24
                                    _ = v32
                                case "foo bar":
                                    switch fi12 {
                                    case -258:
                                        fmt.Println(v3, v11, fi19)
                                    case -9223372036854775808:
                                        var v33 int = v1
                                        _ = v33
                                    case 42:
                                        switch "" {
                                        case "\n\t":
                                            if (255 == 668) {
                                                v34 := GenS1{F0: 127}
                                                var v35 int64 = -1
                                                v35 -= v35
                                                _ = v34
                                                _ = v35
                                            }
                                        case "":
                                            v24 = v24
                                        case "hello":
                                            var v36 int32 = 42
                                            _ = v36
                                        default:
                                            var v37 float64 = float64(-2.25)
                                            _ = v37
                                        }
                                    }
                                    var v38 int = fi19
                                    v39 := GenS3{F0: v3, F1: ("hello" + "a")}
                                    _ = v38
                                    _ = v39
                                }
                                var v40 bool = (false || false)
                                _ = v40
                            }
                            v41 := []int{}
                            fmt.Println(fi19)
                            _ = v24
                            _ = v41
                        } else {
                            v3 = (float64(-2.25) >= (0.0 / (1e10 * -2.25)))
                        }
                    case 0:
                        v11 += 42
                    case 376:
                        var v42 uint8 = 1
                        _ = v42
                    default:
                        _ = len(v2)
                    }
                case "hello":
                    var v43 bool = (v5 == (v5 & v5))
                    _ = v43
                }
            default:
                fmt.Println(v5, v5)
            }
            fmt.Println(v5)
            _ = fi19
        }
        _ = fi12
    }
    switch (-(v11)) {
    case 2147483647:
        v1 = v11
    default:
        v44 := []int{(v11 / ((v11) | 1))}
        _ = v44
    }
    var v45 int = -311
    fmt.Println(v45, v6)
    v2[(v45 & v45)] = (v11 - (v45 ^ 1))
    v45 -= (-(v1))
    switch v1 {
    case 1:
        v46 := []int{v45, (^(v45)), (v45 | (v11 >> ((42) & 63))), 33}
        _ = v46
    case -9223372036854775808:
        if ((float32(1.5) + (1e10 * (0.0 / (0.1 / 1.5)))) >= (0.0 * -2.25)) {
            fmt.Println(v1, v3)
        }
    case 567:
        var v47 int16 = (^(v5))
        _ = v47
    default:
        switch v1 {
        case 1:
            delete(v4, 42)
        case 32767:
            v45 -= 9223372036854775807
        case 4294967295:
            var v48 uint64 = 540
            _ = v48
        }
    }
    if v3 {
        var v49 int32 = -851
        for fi50 := 0; fi50 < 7; fi50++ {
            v51 := v2[fi50]
            if (1 <= 65535) {
                for fi52 := 0; fi52 < 3; fi52++ {
                    fmt.Println(v11, fi50)
                    _ = fi52
                }
                v53 := []int{v45, (v1 >> ((42) & 63)), (v1 * v51), (fi50 ^ (v1 >> ((4294967295) & 63)))}
                _ = v53
            }
            v54 := GenS4{F0: int8(v49), F1: 4294967295, F2: -922}
            _ = fi50
            _ = v51
            _ = v54
        }
        _ = v49
    }
    v45 += 9223372036854775807
    if v3 {
        var v55 int16 = v5
        _ = v55
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v6
}
