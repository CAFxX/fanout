package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 bool; F2 uint64 }
func genFn0() bool {
    var v6 bool = false
    if ("\n\t" < "foo bar") {
        if true {
            for fi7 := 0; fi7 < 2; fi7++ {
                var v8 float64 = 1.5
                _ = fi7
                _ = v8
            }
        }
        var v9 uint = 22
        fmt.Println(v6)
        _ = v9
    } else {
        v10 := []int{1, 127, -2147483648}
        _ = v10
    }
    v11 := []int{425, -1, 859}
    for fi12 := 0; fi12 < 15; fi12++ {
        for fi13 := 0; fi13 < 8; fi13++ {
            if (((float32(3.14159) / 1.0) == float32((1.5 - float64((1e10 + (1e10 / float32(1e10))))))) || (42 == 23)) {
                if v6 {
                    fmt.Println(fi12, v6, fi13)
                    var v14 string = ("a" + ((("hello" + "") + "hello") + ("" + "hello")))
                    v15 := GenS1{F0: 0}
                    _ = v14
                    _ = v15
                } else {
                    v6 = (65535 != 233)
                    fmt.Println(fi12, v6)
                }
            }
            if v6 {
                switch "hello" {
                case "hello":
                    var v16 float32 = 1.5
                    _ = v16
                default:
                    v17 := GenS2{F0: 0}
                    _ = v17
                }
                switch -9223372036854775808 {
                case 4294967295:
                    var v18 int8 = 76
                    _ = v18
                default:
                    var v19 int16 = -32768
                    v19 += (-(v19))
                    _ = v19
                }
                var v20 uint16 = 1
                _ = v20
            }
            _ = fi13
        }
        _ = fi12
    }
    _ = v11
    _ = v6
    return v6
}
func genFn1(p0 uint32, p1 uint8) int32 {
    p0 = 2
    var v21 bool = (false || true)
    _ = p0
    _ = p1
    _ = v21
    return 754
}

func main() {
    var v1 int64 = 42
    v2 := make(map[string]int)
    var v3 int32 = -608
    v4 := make(map[int]bool)
    var v5 int8 = 100
    for fi22 := 0; fi22 < 12; fi22++ {
        var v23 float32 = -2.25
        v24 := GenS1{F0: (v5 << ((65535) & 63))}
        _ = fi22
        _ = v23
        _ = v24
    }
    switch -1 {
    case 1:
        v5 = (v5 | v5)
    case 42:
        for fi25 := 0; fi25 < 14; fi25++ {
            for fi26 := 0; fi26 < 3; fi26++ {
                v27 := []int{(^(fi26)), (^(fi26)), fi25}
                v5 = v5
                switch fi26 {
                case 261:
                    if (!((((("foo bar" + ("hello" + "a")) + "foo bar") + "") <= "hello") && ((float64((1.5 * 3.14159)) != ((3.14159 - (float64(-2.25) + (float64((1.5 / ((((float32(3.14159) / 1.0) / (1e10 / 0.1)) * 1.5) / (1e10 / 3.14159)))) * 1e10))) * (((((3.14159 * 3.14159) + (float64((0.0 * 0.0)) - (((((float64(float32(float64((0.0 * (1.5 + (1.5 / ((3.14159 / 1.0) * 3.14159))))))) / 1.0) - 0.0) + 0.1) + 0.1) * -2.25))) - -2.25) * (0.0 * 3.14159)) - 1.5))) && true))) {
                        v28 := GenS3{F0: (fi26 * (fi25 / ((fi25) | 1))), F1: true, F2: 0}
                        switch "\n\t" {
                        case "hello":
                            var v29 int = -1
                            _ = v29
                        default:
                            if (!(1 != 2147483648)) {
                                v5 = (-(v5))
                            }
                            v1 += v1
                            v5 = (-(v5))
                        }
                        _ = v28
                    }
                    if false {
                        if len(v27) > 0 {
                            v27[0] = fi26
                        }
                    }
                case 4294967295:
                    _ = len(v27)
                case 42:
                    if false {
                        v3 += int32(v1)
                    }
                default:
                    v27 = append(v27, fi25)
                }
                _ = fi26
                _ = v27
            }
            v30 := []int{2147483647, fi25}
            _ = fi25
            _ = v30
        }
    case 4294967295:
        for fi31 := 0; fi31 < 5; fi31++ {
            for fi32 := 0; fi32 < 4; fi32++ {
                if (false && false) {
                    var v33 int = (fi31 >> ((0) & 63))
                    _ = v33
                }
                _ = fi32
            }
            v5 -= -1
            var v34 int32 = (v3 * (v3 & (^(v3))))
            _ = fi31
            _ = v34
        }
    default:
        var v35 uint64 = 129
        _ = v35
    }
    var v36 uint16 = 2
    v37 := []int{42, 583, 9223372036854775807}
    var v38 uint = 18446744073709551615
    fmt.Println(v5, v3, v5)
    fmt.Println(v3, v3)
    v5 -= 127
    for fi39 := 0; fi39 < 9; fi39++ {
        switch ("\n\t" + "a") {
        case "\n\t":
            for fi40 := 0; fi40 < 6; fi40++ {
                switch (^(fi39)) {
                case 42:
                    delete(v2, "")
                case 4294967295:
                    var v41 int32 = (^(v3))
                    _ = v41
                default:
                    if len(v37) > 0 {
                        v37[0] = (^(fi39))
                    }
                }
                var v42 bool = (!(!((100 > (-(v5))) && (v5 == (v5 & v5)))))
                _ = fi40
                _ = v42
            }
        default:
            v43 := v2["a"]
            _ = v43
        }
        v3 -= -293
        v38 = (v38 & (v38 / ((v38) | 1)))
        _ = fi39
    }
    switch -9223372036854775808 {
    case 4294967295:
        v5 += v5
    case 449:
        for fi44 := 0; fi44 < 10; fi44++ {
            for fi45 := 0; fi45 < 1; fi45++ {
                if false {
                    _ = len(v2)
                    var v46 string = ("foo bar" + "a")
                    v36 = (^(v36))
                    _ = v46
                }
                _ = fi45
            }
            v3 -= v3
            _ = fi44
        }
    case 42:
        v1 = 115
    default:
        v5 = (^(v5))
    }
    var v47 bool = ((v38 + v38) == v38)
    v47 = v47
    _ = len(v2)
    v37 = append(v37, 1)
    var v48 float64 = 1.5
    v36 += (v36 >> ((2) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v47
    _ = v48
    _ = v5
}
